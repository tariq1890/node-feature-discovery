/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

const (
	// FeatureLabelNs is the (default) namespace for feature labels.
	FeatureLabelNs = "feature.node.kubernetes.io"

	// FeatureLabelSubNsSuffix is the suffix for allowed feature label sub-namespaces.
	FeatureLabelSubNsSuffix = "." + FeatureLabelNs

	// ProfileLabelNs is the namespace for profile labels.
	ProfileLabelNs = "profile.node.kubernetes.io"

	// ProfileLabelSubNsSuffix is the suffix for allowed profile label sub-namespaces.
	ProfileLabelSubNsSuffix = "." + ProfileLabelNs

	// AnnotationNs namespace for all NFD-related annotations.
	AnnotationNs = "nfd.node.kubernetes.io"

	// ExtendedResourceAnnotation is the annotation that holds all extended resources managed by NFD.
	ExtendedResourceAnnotation = AnnotationNs + "/extended-resources"

	// FeatureLabelsAnnotation is the annotation that holds all feature labels managed by NFD.
	FeatureLabelsAnnotation = AnnotationNs + "/feature-labels"

	// MasterVersionAnnotation is the annotation that holds the version of nfd-master running on the node
	MasterVersionAnnotation = AnnotationNs + "/master.version"

	// WorkerVersionAnnotation is the annotation that holds the version of nfd-worker running on the node
	WorkerVersionAnnotation = AnnotationNs + "/worker.version"
)
