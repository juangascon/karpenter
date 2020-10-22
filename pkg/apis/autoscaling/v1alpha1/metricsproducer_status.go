/*
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

import (
	v1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
)

// MetricsProducerStatus defines the observed state of the resource.
// +kubebuilder:subresource:status
type MetricsProducerStatus struct {
	// +optional
	PendingCapacity *PendingCapacityStatus `json:"pendingCapacity,omitempty"`
	// +optional
	Queue *QueueStatus `json:"queue,omitempty"`
	// +optional
	ReservedCapacity map[v1.ResourceName]string `json:"reservedCapacity,omitempty"`
	// +optional
	ScheduledCapacity *ScheduledCapacityStatus `json:"scheduledCapacity,omitempty"`
	// Conditions is the set of conditions required for the metrics producer to
	// successfully publish metrics to the metrics server
	// +optional
	Conditions apis.Conditions `json:"conditions,omitempty"`
}

const (
	// Calculable is a condition that refers to whether or not the Metrics
	// Producer is able to calculate a metric given the available data. This
	// will be false if no data is available, or if a mathmatical operation
	// desired by the producer results in an undefined value (i.e. div by 0).
	Calculable apis.ConditionType = "Calculable"
)

type PendingCapacityStatus struct {
}
type QueueStatus struct {
}

type ScheduledCapacityStatus struct {
}

// We use knative's libraries for ConditionSets to manage status conditions.
// Conditions are all of "true-happy" polarity. If any condition is false, the resource's "happiness" is false.
// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-conditions
// https://github.com/knative/serving/blob/f1582404be275d6eaaf89ccd908fb44aef9e48b5/vendor/knative.dev/pkg/apis/condition_set.go
func (m *MetricsProducer) StatusConditions() apis.ConditionManager {
	return apis.NewLivingConditionSet(
		Active,
		Calculable,
	).Manage(m)
}

func (m *MetricsProducer) GetConditions() apis.Conditions {
	return m.Status.Conditions
}

func (m *MetricsProducer) SetConditions(conditions apis.Conditions) {
	m.Status.Conditions = conditions
}