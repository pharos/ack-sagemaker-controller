// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitoringScheduleSpec defines the desired state of MonitoringSchedule.
//
// A schedule for a model monitoring job. For information about model monitor,
// see Amazon SageMaker Model Monitor (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html).
type MonitoringScheduleSpec struct {
	// The configuration object that specifies the monitoring schedule and defines
	// the monitoring job.
	// +kubebuilder:validation:Required
	MonitoringScheduleConfig *MonitoringScheduleConfig `json:"monitoringScheduleConfig"`
	// The name of the monitoring schedule. The name must be unique within an AWS
	// Region within an AWS account.
	// +kubebuilder:validation:Required
	MonitoringScheduleName *string `json:"monitoringScheduleName"`
	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL)
	// in the AWS Billing and Cost Management User Guide.
	Tags []*Tag `json:"tags,omitempty"`
}

// MonitoringScheduleStatus defines the observed state of MonitoringSchedule
type MonitoringScheduleStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The time at which the monitoring job was created.
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// A string, up to one KB in size, that contains the reason a monitoring job
	// failed, if it failed.
	FailureReason *string `json:"failureReason,omitempty"`
	// The time at which the monitoring job was last modified.
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	// Describes metadata on the last execution to run, if there was one.
	LastMonitoringExecutionSummary *MonitoringExecutionSummary `json:"lastMonitoringExecutionSummary,omitempty"`
	// The status of an monitoring job.
	MonitoringScheduleStatus *string `json:"monitoringScheduleStatus,omitempty"`
}

// MonitoringSchedule is the Schema for the MonitoringSchedules API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FAILURE-REASON",type=string,priority=1,JSONPath=`.status.failureReason`
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.monitoringScheduleStatus`
type MonitoringSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringScheduleSpec   `json:"spec,omitempty"`
	Status            MonitoringScheduleStatus `json:"status,omitempty"`
}

// MonitoringScheduleList contains a list of MonitoringSchedule
// +kubebuilder:object:root=true
type MonitoringScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringSchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringSchedule{}, &MonitoringScheduleList{})
}
