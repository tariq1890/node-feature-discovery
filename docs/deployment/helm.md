---
title: "Helm"
layout: default
sort: 3
---

# Deployment with Helm
{: .no_toc}

## Table of contents
{: .no_toc .text-delta}

1. TOC
{:toc}

---

Node Feature Discovery Helm chart allow to easily deploy and manage NFD.

## Prerequisites

[Helm package manager](https://helm.sh/) should be installed.

## Deployment

To install the latest stable version:

```bash
export NFD_NS=node-feature-discovery
helm repo add nfd https://kubernetes-sigs.github.io/node-feature-discovery/charts
helm repo update
helm install nfd/node-feature-discovery --namespace $NFD_NS --create-namespace --generate-name
```

To install the latest development version you need to clone the NFD Git
repository and install from there.

```bash
git clone https://github.com/kubernetes-sigs/node-feature-discovery/
cd node-feature-discovery/deployment/helm
export NFD_NS=node-feature-discovery
helm install node-feature-discovery ./node-feature-discovery/ --namespace $NFD_NS --create-namespace
```

See the [configuration](#configuration) section below for instructions how to
alter the deployment parameters.

In order to deploy the [minimal](image-variants#minimal) image you need to
override the image tag:

```bash
helm install node-feature-discovery ./node-feature-discovery/ --set image.tag={{ site.release }}-minimal --namespace $NFD_NS --create-namespace
```

## Configuration

You can override values from `values.yaml` and provide a file with custom values:

```bash
export NFD_NS=node-feature-discovery
helm install nfd/node-feature-discovery -f <path/to/custom/values.yaml> --namespace $NFD_NS --create-namespace
```

To specify each parameter separately you can provide them to helm install command:

```bash
export NFD_NS=node-feature-discovery
helm install nfd/node-feature-discovery --set nameOverride=NFDinstance --set master.replicaCount=2 --namespace $NFD_NS --create-namespace
```

## Uninstalling the chart

To uninstall the `node-feature-discovery` deployment:

```bash
export NFD_NS=node-feature-discovery
helm uninstall node-feature-discovery --namespace $NFD_NS
```

The command removes all the Kubernetes components associated with the chart and
deletes the release.

## Chart parameters

In order to tailor the deployment of the Node Feature Discovery to your cluster needs
We have introduced the following Chart parameters.

### General parameters

| Name | Type | Default | description |
| ---- | ---- | ------- | ----------- |
| `image.repository` | string | `{{ site.container_image | split: ":" | first }}` | NFD image repository |
| `image.tag` | string | `{{ site.release }}` | NFD image tag |
| `image.pullPolicy` | string | `Always` | Image pull policy |
| `imagePullSecrets` | list | [] | ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. [More info](https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod) |
| `nameOverride` | string |  | Override the name of the chart |
| `fullnameOverride` | string |  | Override a default fully qualified app name |
| `nodeFeatureRule.createCRD` | bool | true | Specifies whether to create the NodeFeatureRule CRD |
| `tls.enable` | bool | false | Specifies whether to use TLS for communications between components |
| `tls.certManager` | bool | false | If enabled, requires [cert-manager](https://cert-manager.io/docs/) to be installed and will automatically create the required TLS certificates |

### Master pod parameters

| Name                        | Type    | Default                                 | description                                                                                                                              |
|-----------------------------|---------|-----------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| `master.*`                  | dict    |                                         | NFD master deployment configuration                                                                                                      |
| `master.instance`           | string  |                                         | Instance name. Used to separate annotation namespaces for multiple parallel deployments                                                  |
| `master.extraLabelNs`       | array   | []                                      | List of allowed extra label namespaces                                                                                                   |
| `master.resourceLabels`     | array   | []                                      | List of labels to be registered as extended resources                                                                                          |
| `master.featureRulesController` | bool | null                                   | Specifies whether the controller for processing of NodeFeatureRule objects is enabled. If not set, controller will be enabled if `master.instance` is empty. |
| `master.replicaCount`       | integer | 1                                       | Number of desired pods. This is a pointer to distinguish between explicit zero and not specified                                         |
| `master.podSecurityContext` | dict    | {}                                      | [PodSecurityContext](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod) holds pod-level security attributes and common container settings |
| `master.securityContext`    | dict    | {}                                      | Container [security settings](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container)|
| `master.serviceAccount.create` | bool | true                                    | Specifies whether a service account should be created
| `master.serviceAccount.annotations` | dict | {}                                 | Annotations to add to the service account
| `master.serviceAccount.name` | string |                                         | The name of the service account to use. If not set and create is true, a name is generated using the fullname template
| `master.rbac.create`        | bool    | true                                    | Specifies whether to create [RBAC][rbac] configuration for nfd-master
| `master.service.type`       | string  | ClusterIP                               | NFD master service type                                                                                                                  |
| `master.service.port`       | integer | 8080                                    | NFD master service port                                                                                                                  |
| `master.resources`          | dict    | {}                                      | NFD master pod [resources management](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)                    |
| `master.nodeSelector`       | dict    | {}                                      | NFD master pod [node selector](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector)                    |
| `master.tolerations`        | dict    | _Scheduling to master node is disabled_ | NFD master pod [tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/)                              |
| `master.annotations`        | dict    | {}                                      | NFD master pod [annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)                                |
| `master.affinity`           | dict    |                                         | NFD master pod required [node affinity](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes-using-node-affinity/) |
| `master.deploymentAnnotations` | dict | {}                                      | NFD master deployment [annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/) |

### Worker pod parameters

| Name | Type | Default | description |
| ---- | ---- | ------- | ----------- |
| `worker.*` | dict |  | NFD worker daemonset configuration |
| `worker.config` | dict |  | NFD worker [configuration](../reference/worker-configuration-reference) |
| `worker.podSecurityContext` | dict | {} | [PodSecurityContext](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod) holds pod-level security attributes and common container settings |
| `worker.securityContext` | dict | {} | Container [security settings](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container) |
| `worker.serviceAccount.create`      | bool   | true | Specifies whether a service account for nfd-worker should be created
| `worker.serviceAccount.annotations` | dict   | {}   | Annotations to add to the service account for nfd-worker
| `worker.serviceAccount.name`        | string |      | The name of the service account to use for nfd-worker. If not set and create is true, a name is generated using the fullname template (suffixed with `-worker`)
| `worker.mountUsrSrc` | bool | false | Specifies whether to allow users to mount the hostpath /user/src. Does not work on systems without /usr/src AND a read-only /usr |
| `worker.resources` | dict | {} | NFD worker pod [resources management](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) |
| `worker.nodeSelector` | dict | {} | NFD worker pod [node selector](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector) |
| `worker.tolerations` | dict | {} | NFD worker pod [node tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) |
| `worker.priorityClassName` | string |  | NFD worker pod [priority class](https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/) |
| `worker.annotations` | dict | {} | NFD worker pod [annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/) |
| `worker.daemonsetAnnotations` | dict | {} | NFD worker daemonset [annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/) |

### Topology updater parameters

| Name                                          | Type   | Default | description                                                                                                                                                                                           |
|-----------------------------------------------|--------|---------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `topologyUpdater.*`                           | dict   |         | NFD Topology Updater configuration                                                                                                                                                                    |
| `topologyUpdater.enable`                      | bool   | false   | Specifies whether the NFD Topology Updater should be created                                                                                                                                          |
| `topologyUpdater.createCRDs`                  | bool   | false   | Specifies whether the NFD Topology Updater CRDs should be created                                                                                                                                     |
| `topologyUpdater.serviceAccount.create`       | bool   | true    | Specifies whether the service account for topology updater should be created                                                                                                                          |
| `topologyUpdater.serviceAccount.annotations`  | dict   | {}      | Annotations to add to the service account for topology updater                                                                                                                                        |
| `topologyUpdater.serviceAccount.name`         | string |         | The name of the service account for topology updater to use. If not set and create is true, a name is generated using the fullname template and `-topology-updater` suffix                            |
| `topologyUpdater.rbac.create`                 | bool   | false   | Specifies whether to create [RBAC][rbac] configuration for topology updater                                                                                                                           |
| `topologyUpdater.kubeletConfigPath`           | string | ""      | Specifies the kubelet config host path                                                                                                                                                                |
| `topologyUpdater.kubeletPodResourcesSockPath` | string | ""      | Specifies the kubelet sock path to read pod resources                                                                                                                                                 |
| `topologyUpdater.updateInterval`              | string | 60s     | Time to sleep between CR updates. Non-positive value implies no CR update.                                                                                                                            |
| `topologyUpdater.watchNamespace`              | string | `*`     | Namespace to watch pods, `*` for all namespaces                                                                                                                                                       |
| `topologyUpdater.podSecurityContext`          | dict   | {}      | [PodSecurityContext](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod) holds pod-level security attributes and common container settings |
| `topologyUpdater.securityContext`             | dict   | {}      | Container [security settings](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container)                                                    |
| `topologyUpdater.resources`                   | dict   | {}      | Topology updater pod [resources management](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)                                                                           |
| `topologyUpdater.nodeSelector`                | dict   | {}      | Topology updater pod [node selector](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector)                                                                           |
| `topologyUpdater.tolerations`                 | dict   | {}      | Topology updater pod [node tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/)                                                                                |
| `topologyUpdater.annotations`                 | dict   | {}      | Topology updater pod [annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)                                                                                    |
| `topologyUpdater.affinity`                    | dict   | {}      | Topology updater pod [affinity](https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes-using-node-affinity/)                                                                      |
| `topologyUpdater.config`                      | dict   |         | [configuration](../reference/topology-updater-configuration-reference)                                                                                                                                |

<!-- Links -->
[rbac]: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
