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

package training_job

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.TrainingJob{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeTrainingJobOutput
	resp, err = rm.sdkapi.DescribeTrainingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeTrainingJob", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ValidationException" && strings.HasPrefix(awsErr.Message(), "Requested resource not found") {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AlgorithmSpecification != nil {
		f0 := &svcapitypes.AlgorithmSpecification{}
		if resp.AlgorithmSpecification.AlgorithmName != nil {
			f0.AlgorithmName = resp.AlgorithmSpecification.AlgorithmName
		}
		if resp.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries != nil {
			f0.EnableSageMakerMetricsTimeSeries = resp.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries
		}
		if resp.AlgorithmSpecification.MetricDefinitions != nil {
			f0f2 := []*svcapitypes.MetricDefinition{}
			for _, f0f2iter := range resp.AlgorithmSpecification.MetricDefinitions {
				f0f2elem := &svcapitypes.MetricDefinition{}
				if f0f2iter.Name != nil {
					f0f2elem.Name = f0f2iter.Name
				}
				if f0f2iter.Regex != nil {
					f0f2elem.Regex = f0f2iter.Regex
				}
				f0f2 = append(f0f2, f0f2elem)
			}
			f0.MetricDefinitions = f0f2
		}
		if resp.AlgorithmSpecification.TrainingImage != nil {
			f0.TrainingImage = resp.AlgorithmSpecification.TrainingImage
		}
		if resp.AlgorithmSpecification.TrainingInputMode != nil {
			f0.TrainingInputMode = resp.AlgorithmSpecification.TrainingInputMode
		}
		ko.Spec.AlgorithmSpecification = f0
	} else {
		ko.Spec.AlgorithmSpecification = nil
	}
	if resp.CheckpointConfig != nil {
		f3 := &svcapitypes.CheckpointConfig{}
		if resp.CheckpointConfig.LocalPath != nil {
			f3.LocalPath = resp.CheckpointConfig.LocalPath
		}
		if resp.CheckpointConfig.S3Uri != nil {
			f3.S3URI = resp.CheckpointConfig.S3Uri
		}
		ko.Spec.CheckpointConfig = f3
	} else {
		ko.Spec.CheckpointConfig = nil
	}
	if resp.DebugHookConfig != nil {
		f5 := &svcapitypes.DebugHookConfig{}
		if resp.DebugHookConfig.CollectionConfigurations != nil {
			f5f0 := []*svcapitypes.CollectionConfiguration{}
			for _, f5f0iter := range resp.DebugHookConfig.CollectionConfigurations {
				f5f0elem := &svcapitypes.CollectionConfiguration{}
				if f5f0iter.CollectionName != nil {
					f5f0elem.CollectionName = f5f0iter.CollectionName
				}
				if f5f0iter.CollectionParameters != nil {
					f5f0elemf1 := map[string]*string{}
					for f5f0elemf1key, f5f0elemf1valiter := range f5f0iter.CollectionParameters {
						var f5f0elemf1val string
						f5f0elemf1val = *f5f0elemf1valiter
						f5f0elemf1[f5f0elemf1key] = &f5f0elemf1val
					}
					f5f0elem.CollectionParameters = f5f0elemf1
				}
				f5f0 = append(f5f0, f5f0elem)
			}
			f5.CollectionConfigurations = f5f0
		}
		if resp.DebugHookConfig.HookParameters != nil {
			f5f1 := map[string]*string{}
			for f5f1key, f5f1valiter := range resp.DebugHookConfig.HookParameters {
				var f5f1val string
				f5f1val = *f5f1valiter
				f5f1[f5f1key] = &f5f1val
			}
			f5.HookParameters = f5f1
		}
		if resp.DebugHookConfig.LocalPath != nil {
			f5.LocalPath = resp.DebugHookConfig.LocalPath
		}
		if resp.DebugHookConfig.S3OutputPath != nil {
			f5.S3OutputPath = resp.DebugHookConfig.S3OutputPath
		}
		ko.Spec.DebugHookConfig = f5
	} else {
		ko.Spec.DebugHookConfig = nil
	}
	if resp.DebugRuleConfigurations != nil {
		f6 := []*svcapitypes.DebugRuleConfiguration{}
		for _, f6iter := range resp.DebugRuleConfigurations {
			f6elem := &svcapitypes.DebugRuleConfiguration{}
			if f6iter.InstanceType != nil {
				f6elem.InstanceType = f6iter.InstanceType
			}
			if f6iter.LocalPath != nil {
				f6elem.LocalPath = f6iter.LocalPath
			}
			if f6iter.RuleConfigurationName != nil {
				f6elem.RuleConfigurationName = f6iter.RuleConfigurationName
			}
			if f6iter.RuleEvaluatorImage != nil {
				f6elem.RuleEvaluatorImage = f6iter.RuleEvaluatorImage
			}
			if f6iter.RuleParameters != nil {
				f6elemf4 := map[string]*string{}
				for f6elemf4key, f6elemf4valiter := range f6iter.RuleParameters {
					var f6elemf4val string
					f6elemf4val = *f6elemf4valiter
					f6elemf4[f6elemf4key] = &f6elemf4val
				}
				f6elem.RuleParameters = f6elemf4
			}
			if f6iter.S3OutputPath != nil {
				f6elem.S3OutputPath = f6iter.S3OutputPath
			}
			if f6iter.VolumeSizeInGB != nil {
				f6elem.VolumeSizeInGB = f6iter.VolumeSizeInGB
			}
			f6 = append(f6, f6elem)
		}
		ko.Spec.DebugRuleConfigurations = f6
	} else {
		ko.Spec.DebugRuleConfigurations = nil
	}
	if resp.DebugRuleEvaluationStatuses != nil {
		f7 := []*svcapitypes.DebugRuleEvaluationStatus{}
		for _, f7iter := range resp.DebugRuleEvaluationStatuses {
			f7elem := &svcapitypes.DebugRuleEvaluationStatus{}
			if f7iter.LastModifiedTime != nil {
				f7elem.LastModifiedTime = &metav1.Time{*f7iter.LastModifiedTime}
			}
			if f7iter.RuleConfigurationName != nil {
				f7elem.RuleConfigurationName = f7iter.RuleConfigurationName
			}
			if f7iter.RuleEvaluationJobArn != nil {
				f7elem.RuleEvaluationJobARN = f7iter.RuleEvaluationJobArn
			}
			if f7iter.RuleEvaluationStatus != nil {
				f7elem.RuleEvaluationStatus = f7iter.RuleEvaluationStatus
			}
			if f7iter.StatusDetails != nil {
				f7elem.StatusDetails = f7iter.StatusDetails
			}
			f7 = append(f7, f7elem)
		}
		ko.Status.DebugRuleEvaluationStatuses = f7
	} else {
		ko.Status.DebugRuleEvaluationStatuses = nil
	}
	if resp.EnableInterContainerTrafficEncryption != nil {
		ko.Spec.EnableInterContainerTrafficEncryption = resp.EnableInterContainerTrafficEncryption
	} else {
		ko.Spec.EnableInterContainerTrafficEncryption = nil
	}
	if resp.EnableManagedSpotTraining != nil {
		ko.Spec.EnableManagedSpotTraining = resp.EnableManagedSpotTraining
	} else {
		ko.Spec.EnableManagedSpotTraining = nil
	}
	if resp.EnableNetworkIsolation != nil {
		ko.Spec.EnableNetworkIsolation = resp.EnableNetworkIsolation
	} else {
		ko.Spec.EnableNetworkIsolation = nil
	}
	if resp.Environment != nil {
		f11 := map[string]*string{}
		for f11key, f11valiter := range resp.Environment {
			var f11val string
			f11val = *f11valiter
			f11[f11key] = &f11val
		}
		ko.Spec.Environment = f11
	} else {
		ko.Spec.Environment = nil
	}
	if resp.ExperimentConfig != nil {
		f12 := &svcapitypes.ExperimentConfig{}
		if resp.ExperimentConfig.ExperimentName != nil {
			f12.ExperimentName = resp.ExperimentConfig.ExperimentName
		}
		if resp.ExperimentConfig.TrialComponentDisplayName != nil {
			f12.TrialComponentDisplayName = resp.ExperimentConfig.TrialComponentDisplayName
		}
		if resp.ExperimentConfig.TrialName != nil {
			f12.TrialName = resp.ExperimentConfig.TrialName
		}
		ko.Spec.ExperimentConfig = f12
	} else {
		ko.Spec.ExperimentConfig = nil
	}
	if resp.FailureReason != nil {
		ko.Status.FailureReason = resp.FailureReason
	} else {
		ko.Status.FailureReason = nil
	}
	if resp.HyperParameters != nil {
		f15 := map[string]*string{}
		for f15key, f15valiter := range resp.HyperParameters {
			var f15val string
			f15val = *f15valiter
			f15[f15key] = &f15val
		}
		ko.Spec.HyperParameters = f15
	} else {
		ko.Spec.HyperParameters = nil
	}
	if resp.InputDataConfig != nil {
		f16 := []*svcapitypes.Channel{}
		for _, f16iter := range resp.InputDataConfig {
			f16elem := &svcapitypes.Channel{}
			if f16iter.ChannelName != nil {
				f16elem.ChannelName = f16iter.ChannelName
			}
			if f16iter.CompressionType != nil {
				f16elem.CompressionType = f16iter.CompressionType
			}
			if f16iter.ContentType != nil {
				f16elem.ContentType = f16iter.ContentType
			}
			if f16iter.DataSource != nil {
				f16elemf3 := &svcapitypes.DataSource{}
				if f16iter.DataSource.FileSystemDataSource != nil {
					f16elemf3f0 := &svcapitypes.FileSystemDataSource{}
					if f16iter.DataSource.FileSystemDataSource.DirectoryPath != nil {
						f16elemf3f0.DirectoryPath = f16iter.DataSource.FileSystemDataSource.DirectoryPath
					}
					if f16iter.DataSource.FileSystemDataSource.FileSystemAccessMode != nil {
						f16elemf3f0.FileSystemAccessMode = f16iter.DataSource.FileSystemDataSource.FileSystemAccessMode
					}
					if f16iter.DataSource.FileSystemDataSource.FileSystemId != nil {
						f16elemf3f0.FileSystemID = f16iter.DataSource.FileSystemDataSource.FileSystemId
					}
					if f16iter.DataSource.FileSystemDataSource.FileSystemType != nil {
						f16elemf3f0.FileSystemType = f16iter.DataSource.FileSystemDataSource.FileSystemType
					}
					f16elemf3.FileSystemDataSource = f16elemf3f0
				}
				if f16iter.DataSource.S3DataSource != nil {
					f16elemf3f1 := &svcapitypes.S3DataSource{}
					if f16iter.DataSource.S3DataSource.AttributeNames != nil {
						f16elemf3f1f0 := []*string{}
						for _, f16elemf3f1f0iter := range f16iter.DataSource.S3DataSource.AttributeNames {
							var f16elemf3f1f0elem string
							f16elemf3f1f0elem = *f16elemf3f1f0iter
							f16elemf3f1f0 = append(f16elemf3f1f0, &f16elemf3f1f0elem)
						}
						f16elemf3f1.AttributeNames = f16elemf3f1f0
					}
					if f16iter.DataSource.S3DataSource.S3DataDistributionType != nil {
						f16elemf3f1.S3DataDistributionType = f16iter.DataSource.S3DataSource.S3DataDistributionType
					}
					if f16iter.DataSource.S3DataSource.S3DataType != nil {
						f16elemf3f1.S3DataType = f16iter.DataSource.S3DataSource.S3DataType
					}
					if f16iter.DataSource.S3DataSource.S3Uri != nil {
						f16elemf3f1.S3URI = f16iter.DataSource.S3DataSource.S3Uri
					}
					f16elemf3.S3DataSource = f16elemf3f1
				}
				f16elem.DataSource = f16elemf3
			}
			if f16iter.InputMode != nil {
				f16elem.InputMode = f16iter.InputMode
			}
			if f16iter.RecordWrapperType != nil {
				f16elem.RecordWrapperType = f16iter.RecordWrapperType
			}
			if f16iter.ShuffleConfig != nil {
				f16elemf6 := &svcapitypes.ShuffleConfig{}
				if f16iter.ShuffleConfig.Seed != nil {
					f16elemf6.Seed = f16iter.ShuffleConfig.Seed
				}
				f16elem.ShuffleConfig = f16elemf6
			}
			f16 = append(f16, f16elem)
		}
		ko.Spec.InputDataConfig = f16
	} else {
		ko.Spec.InputDataConfig = nil
	}
	if resp.ModelArtifacts != nil {
		f19 := &svcapitypes.ModelArtifacts{}
		if resp.ModelArtifacts.S3ModelArtifacts != nil {
			f19.S3ModelArtifacts = resp.ModelArtifacts.S3ModelArtifacts
		}
		ko.Status.ModelArtifacts = f19
	} else {
		ko.Status.ModelArtifacts = nil
	}
	if resp.OutputDataConfig != nil {
		f20 := &svcapitypes.OutputDataConfig{}
		if resp.OutputDataConfig.KmsKeyId != nil {
			f20.KMSKeyID = resp.OutputDataConfig.KmsKeyId
		}
		if resp.OutputDataConfig.S3OutputPath != nil {
			f20.S3OutputPath = resp.OutputDataConfig.S3OutputPath
		}
		ko.Spec.OutputDataConfig = f20
	} else {
		ko.Spec.OutputDataConfig = nil
	}
	if resp.ProfilerConfig != nil {
		f21 := &svcapitypes.ProfilerConfig{}
		if resp.ProfilerConfig.ProfilingIntervalInMilliseconds != nil {
			f21.ProfilingIntervalInMilliseconds = resp.ProfilerConfig.ProfilingIntervalInMilliseconds
		}
		if resp.ProfilerConfig.ProfilingParameters != nil {
			f21f1 := map[string]*string{}
			for f21f1key, f21f1valiter := range resp.ProfilerConfig.ProfilingParameters {
				var f21f1val string
				f21f1val = *f21f1valiter
				f21f1[f21f1key] = &f21f1val
			}
			f21.ProfilingParameters = f21f1
		}
		if resp.ProfilerConfig.S3OutputPath != nil {
			f21.S3OutputPath = resp.ProfilerConfig.S3OutputPath
		}
		ko.Spec.ProfilerConfig = f21
	} else {
		ko.Spec.ProfilerConfig = nil
	}
	if resp.ProfilerRuleConfigurations != nil {
		f22 := []*svcapitypes.ProfilerRuleConfiguration{}
		for _, f22iter := range resp.ProfilerRuleConfigurations {
			f22elem := &svcapitypes.ProfilerRuleConfiguration{}
			if f22iter.InstanceType != nil {
				f22elem.InstanceType = f22iter.InstanceType
			}
			if f22iter.LocalPath != nil {
				f22elem.LocalPath = f22iter.LocalPath
			}
			if f22iter.RuleConfigurationName != nil {
				f22elem.RuleConfigurationName = f22iter.RuleConfigurationName
			}
			if f22iter.RuleEvaluatorImage != nil {
				f22elem.RuleEvaluatorImage = f22iter.RuleEvaluatorImage
			}
			if f22iter.RuleParameters != nil {
				f22elemf4 := map[string]*string{}
				for f22elemf4key, f22elemf4valiter := range f22iter.RuleParameters {
					var f22elemf4val string
					f22elemf4val = *f22elemf4valiter
					f22elemf4[f22elemf4key] = &f22elemf4val
				}
				f22elem.RuleParameters = f22elemf4
			}
			if f22iter.S3OutputPath != nil {
				f22elem.S3OutputPath = f22iter.S3OutputPath
			}
			if f22iter.VolumeSizeInGB != nil {
				f22elem.VolumeSizeInGB = f22iter.VolumeSizeInGB
			}
			f22 = append(f22, f22elem)
		}
		ko.Spec.ProfilerRuleConfigurations = f22
	} else {
		ko.Spec.ProfilerRuleConfigurations = nil
	}
	if resp.ProfilerRuleEvaluationStatuses != nil {
		f23 := []*svcapitypes.ProfilerRuleEvaluationStatus{}
		for _, f23iter := range resp.ProfilerRuleEvaluationStatuses {
			f23elem := &svcapitypes.ProfilerRuleEvaluationStatus{}
			if f23iter.LastModifiedTime != nil {
				f23elem.LastModifiedTime = &metav1.Time{*f23iter.LastModifiedTime}
			}
			if f23iter.RuleConfigurationName != nil {
				f23elem.RuleConfigurationName = f23iter.RuleConfigurationName
			}
			if f23iter.RuleEvaluationJobArn != nil {
				f23elem.RuleEvaluationJobARN = f23iter.RuleEvaluationJobArn
			}
			if f23iter.RuleEvaluationStatus != nil {
				f23elem.RuleEvaluationStatus = f23iter.RuleEvaluationStatus
			}
			if f23iter.StatusDetails != nil {
				f23elem.StatusDetails = f23iter.StatusDetails
			}
			f23 = append(f23, f23elem)
		}
		ko.Status.ProfilerRuleEvaluationStatuses = f23
	} else {
		ko.Status.ProfilerRuleEvaluationStatuses = nil
	}
	if resp.ResourceConfig != nil {
		f25 := &svcapitypes.ResourceConfig{}
		if resp.ResourceConfig.InstanceCount != nil {
			f25.InstanceCount = resp.ResourceConfig.InstanceCount
		}
		if resp.ResourceConfig.InstanceType != nil {
			f25.InstanceType = resp.ResourceConfig.InstanceType
		}
		if resp.ResourceConfig.VolumeKmsKeyId != nil {
			f25.VolumeKMSKeyID = resp.ResourceConfig.VolumeKmsKeyId
		}
		if resp.ResourceConfig.VolumeSizeInGB != nil {
			f25.VolumeSizeInGB = resp.ResourceConfig.VolumeSizeInGB
		}
		ko.Spec.ResourceConfig = f25
	} else {
		ko.Spec.ResourceConfig = nil
	}
	if resp.RoleArn != nil {
		ko.Spec.RoleARN = resp.RoleArn
	} else {
		ko.Spec.RoleARN = nil
	}
	if resp.SecondaryStatus != nil {
		ko.Status.SecondaryStatus = resp.SecondaryStatus
	} else {
		ko.Status.SecondaryStatus = nil
	}
	if resp.StoppingCondition != nil {
		f29 := &svcapitypes.StoppingCondition{}
		if resp.StoppingCondition.MaxRuntimeInSeconds != nil {
			f29.MaxRuntimeInSeconds = resp.StoppingCondition.MaxRuntimeInSeconds
		}
		if resp.StoppingCondition.MaxWaitTimeInSeconds != nil {
			f29.MaxWaitTimeInSeconds = resp.StoppingCondition.MaxWaitTimeInSeconds
		}
		ko.Spec.StoppingCondition = f29
	} else {
		ko.Spec.StoppingCondition = nil
	}
	if resp.TensorBoardOutputConfig != nil {
		f30 := &svcapitypes.TensorBoardOutputConfig{}
		if resp.TensorBoardOutputConfig.LocalPath != nil {
			f30.LocalPath = resp.TensorBoardOutputConfig.LocalPath
		}
		if resp.TensorBoardOutputConfig.S3OutputPath != nil {
			f30.S3OutputPath = resp.TensorBoardOutputConfig.S3OutputPath
		}
		ko.Spec.TensorBoardOutputConfig = f30
	} else {
		ko.Spec.TensorBoardOutputConfig = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.TrainingJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.TrainingJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.TrainingJobName != nil {
		ko.Spec.TrainingJobName = resp.TrainingJobName
	} else {
		ko.Spec.TrainingJobName = nil
	}
	if resp.TrainingJobStatus != nil {
		ko.Status.TrainingJobStatus = resp.TrainingJobStatus
	} else {
		ko.Status.TrainingJobStatus = nil
	}
	if resp.VpcConfig != nil {
		f38 := &svcapitypes.VPCConfig{}
		if resp.VpcConfig.SecurityGroupIds != nil {
			f38f0 := []*string{}
			for _, f38f0iter := range resp.VpcConfig.SecurityGroupIds {
				var f38f0elem string
				f38f0elem = *f38f0iter
				f38f0 = append(f38f0, &f38f0elem)
			}
			f38.SecurityGroupIDs = f38f0
		}
		if resp.VpcConfig.Subnets != nil {
			f38f1 := []*string{}
			for _, f38f1iter := range resp.VpcConfig.Subnets {
				var f38f1elem string
				f38f1elem = *f38f1iter
				f38f1 = append(f38f1, &f38f1elem)
			}
			f38.Subnets = f38f1
		}
		ko.Spec.VPCConfig = f38
	} else {
		ko.Spec.VPCConfig = nil
	}

	rm.setStatusDefaults(ko)
	rm.customSetOutput(&resource{ko})
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.TrainingJobName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeTrainingJobInput, error) {
	res := &svcsdk.DescribeTrainingJobInput{}

	if r.ko.Spec.TrainingJobName != nil {
		res.SetTrainingJobName(*r.ko.Spec.TrainingJobName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateTrainingJobOutput
	_ = resp
	resp, err = rm.sdkapi.CreateTrainingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateTrainingJob", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.TrainingJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.TrainingJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateTrainingJobInput, error) {
	res := &svcsdk.CreateTrainingJobInput{}

	if r.ko.Spec.AlgorithmSpecification != nil {
		f0 := &svcsdk.AlgorithmSpecification{}
		if r.ko.Spec.AlgorithmSpecification.AlgorithmName != nil {
			f0.SetAlgorithmName(*r.ko.Spec.AlgorithmSpecification.AlgorithmName)
		}
		if r.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries != nil {
			f0.SetEnableSageMakerMetricsTimeSeries(*r.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries)
		}
		if r.ko.Spec.AlgorithmSpecification.MetricDefinitions != nil {
			f0f2 := []*svcsdk.MetricDefinition{}
			for _, f0f2iter := range r.ko.Spec.AlgorithmSpecification.MetricDefinitions {
				f0f2elem := &svcsdk.MetricDefinition{}
				if f0f2iter.Name != nil {
					f0f2elem.SetName(*f0f2iter.Name)
				}
				if f0f2iter.Regex != nil {
					f0f2elem.SetRegex(*f0f2iter.Regex)
				}
				f0f2 = append(f0f2, f0f2elem)
			}
			f0.SetMetricDefinitions(f0f2)
		}
		if r.ko.Spec.AlgorithmSpecification.TrainingImage != nil {
			f0.SetTrainingImage(*r.ko.Spec.AlgorithmSpecification.TrainingImage)
		}
		if r.ko.Spec.AlgorithmSpecification.TrainingInputMode != nil {
			f0.SetTrainingInputMode(*r.ko.Spec.AlgorithmSpecification.TrainingInputMode)
		}
		res.SetAlgorithmSpecification(f0)
	}
	if r.ko.Spec.CheckpointConfig != nil {
		f1 := &svcsdk.CheckpointConfig{}
		if r.ko.Spec.CheckpointConfig.LocalPath != nil {
			f1.SetLocalPath(*r.ko.Spec.CheckpointConfig.LocalPath)
		}
		if r.ko.Spec.CheckpointConfig.S3URI != nil {
			f1.SetS3Uri(*r.ko.Spec.CheckpointConfig.S3URI)
		}
		res.SetCheckpointConfig(f1)
	}
	if r.ko.Spec.DebugHookConfig != nil {
		f2 := &svcsdk.DebugHookConfig{}
		if r.ko.Spec.DebugHookConfig.CollectionConfigurations != nil {
			f2f0 := []*svcsdk.CollectionConfiguration{}
			for _, f2f0iter := range r.ko.Spec.DebugHookConfig.CollectionConfigurations {
				f2f0elem := &svcsdk.CollectionConfiguration{}
				if f2f0iter.CollectionName != nil {
					f2f0elem.SetCollectionName(*f2f0iter.CollectionName)
				}
				if f2f0iter.CollectionParameters != nil {
					f2f0elemf1 := map[string]*string{}
					for f2f0elemf1key, f2f0elemf1valiter := range f2f0iter.CollectionParameters {
						var f2f0elemf1val string
						f2f0elemf1val = *f2f0elemf1valiter
						f2f0elemf1[f2f0elemf1key] = &f2f0elemf1val
					}
					f2f0elem.SetCollectionParameters(f2f0elemf1)
				}
				f2f0 = append(f2f0, f2f0elem)
			}
			f2.SetCollectionConfigurations(f2f0)
		}
		if r.ko.Spec.DebugHookConfig.HookParameters != nil {
			f2f1 := map[string]*string{}
			for f2f1key, f2f1valiter := range r.ko.Spec.DebugHookConfig.HookParameters {
				var f2f1val string
				f2f1val = *f2f1valiter
				f2f1[f2f1key] = &f2f1val
			}
			f2.SetHookParameters(f2f1)
		}
		if r.ko.Spec.DebugHookConfig.LocalPath != nil {
			f2.SetLocalPath(*r.ko.Spec.DebugHookConfig.LocalPath)
		}
		if r.ko.Spec.DebugHookConfig.S3OutputPath != nil {
			f2.SetS3OutputPath(*r.ko.Spec.DebugHookConfig.S3OutputPath)
		}
		res.SetDebugHookConfig(f2)
	}
	if r.ko.Spec.DebugRuleConfigurations != nil {
		f3 := []*svcsdk.DebugRuleConfiguration{}
		for _, f3iter := range r.ko.Spec.DebugRuleConfigurations {
			f3elem := &svcsdk.DebugRuleConfiguration{}
			if f3iter.InstanceType != nil {
				f3elem.SetInstanceType(*f3iter.InstanceType)
			}
			if f3iter.LocalPath != nil {
				f3elem.SetLocalPath(*f3iter.LocalPath)
			}
			if f3iter.RuleConfigurationName != nil {
				f3elem.SetRuleConfigurationName(*f3iter.RuleConfigurationName)
			}
			if f3iter.RuleEvaluatorImage != nil {
				f3elem.SetRuleEvaluatorImage(*f3iter.RuleEvaluatorImage)
			}
			if f3iter.RuleParameters != nil {
				f3elemf4 := map[string]*string{}
				for f3elemf4key, f3elemf4valiter := range f3iter.RuleParameters {
					var f3elemf4val string
					f3elemf4val = *f3elemf4valiter
					f3elemf4[f3elemf4key] = &f3elemf4val
				}
				f3elem.SetRuleParameters(f3elemf4)
			}
			if f3iter.S3OutputPath != nil {
				f3elem.SetS3OutputPath(*f3iter.S3OutputPath)
			}
			if f3iter.VolumeSizeInGB != nil {
				f3elem.SetVolumeSizeInGB(*f3iter.VolumeSizeInGB)
			}
			f3 = append(f3, f3elem)
		}
		res.SetDebugRuleConfigurations(f3)
	}
	if r.ko.Spec.EnableInterContainerTrafficEncryption != nil {
		res.SetEnableInterContainerTrafficEncryption(*r.ko.Spec.EnableInterContainerTrafficEncryption)
	}
	if r.ko.Spec.EnableManagedSpotTraining != nil {
		res.SetEnableManagedSpotTraining(*r.ko.Spec.EnableManagedSpotTraining)
	}
	if r.ko.Spec.EnableNetworkIsolation != nil {
		res.SetEnableNetworkIsolation(*r.ko.Spec.EnableNetworkIsolation)
	}
	if r.ko.Spec.Environment != nil {
		f7 := map[string]*string{}
		for f7key, f7valiter := range r.ko.Spec.Environment {
			var f7val string
			f7val = *f7valiter
			f7[f7key] = &f7val
		}
		res.SetEnvironment(f7)
	}
	if r.ko.Spec.ExperimentConfig != nil {
		f8 := &svcsdk.ExperimentConfig{}
		if r.ko.Spec.ExperimentConfig.ExperimentName != nil {
			f8.SetExperimentName(*r.ko.Spec.ExperimentConfig.ExperimentName)
		}
		if r.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil {
			f8.SetTrialComponentDisplayName(*r.ko.Spec.ExperimentConfig.TrialComponentDisplayName)
		}
		if r.ko.Spec.ExperimentConfig.TrialName != nil {
			f8.SetTrialName(*r.ko.Spec.ExperimentConfig.TrialName)
		}
		res.SetExperimentConfig(f8)
	}
	if r.ko.Spec.HyperParameters != nil {
		f9 := map[string]*string{}
		for f9key, f9valiter := range r.ko.Spec.HyperParameters {
			var f9val string
			f9val = *f9valiter
			f9[f9key] = &f9val
		}
		res.SetHyperParameters(f9)
	}
	if r.ko.Spec.InputDataConfig != nil {
		f10 := []*svcsdk.Channel{}
		for _, f10iter := range r.ko.Spec.InputDataConfig {
			f10elem := &svcsdk.Channel{}
			if f10iter.ChannelName != nil {
				f10elem.SetChannelName(*f10iter.ChannelName)
			}
			if f10iter.CompressionType != nil {
				f10elem.SetCompressionType(*f10iter.CompressionType)
			}
			if f10iter.ContentType != nil {
				f10elem.SetContentType(*f10iter.ContentType)
			}
			if f10iter.DataSource != nil {
				f10elemf3 := &svcsdk.DataSource{}
				if f10iter.DataSource.FileSystemDataSource != nil {
					f10elemf3f0 := &svcsdk.FileSystemDataSource{}
					if f10iter.DataSource.FileSystemDataSource.DirectoryPath != nil {
						f10elemf3f0.SetDirectoryPath(*f10iter.DataSource.FileSystemDataSource.DirectoryPath)
					}
					if f10iter.DataSource.FileSystemDataSource.FileSystemAccessMode != nil {
						f10elemf3f0.SetFileSystemAccessMode(*f10iter.DataSource.FileSystemDataSource.FileSystemAccessMode)
					}
					if f10iter.DataSource.FileSystemDataSource.FileSystemID != nil {
						f10elemf3f0.SetFileSystemId(*f10iter.DataSource.FileSystemDataSource.FileSystemID)
					}
					if f10iter.DataSource.FileSystemDataSource.FileSystemType != nil {
						f10elemf3f0.SetFileSystemType(*f10iter.DataSource.FileSystemDataSource.FileSystemType)
					}
					f10elemf3.SetFileSystemDataSource(f10elemf3f0)
				}
				if f10iter.DataSource.S3DataSource != nil {
					f10elemf3f1 := &svcsdk.S3DataSource{}
					if f10iter.DataSource.S3DataSource.AttributeNames != nil {
						f10elemf3f1f0 := []*string{}
						for _, f10elemf3f1f0iter := range f10iter.DataSource.S3DataSource.AttributeNames {
							var f10elemf3f1f0elem string
							f10elemf3f1f0elem = *f10elemf3f1f0iter
							f10elemf3f1f0 = append(f10elemf3f1f0, &f10elemf3f1f0elem)
						}
						f10elemf3f1.SetAttributeNames(f10elemf3f1f0)
					}
					if f10iter.DataSource.S3DataSource.S3DataDistributionType != nil {
						f10elemf3f1.SetS3DataDistributionType(*f10iter.DataSource.S3DataSource.S3DataDistributionType)
					}
					if f10iter.DataSource.S3DataSource.S3DataType != nil {
						f10elemf3f1.SetS3DataType(*f10iter.DataSource.S3DataSource.S3DataType)
					}
					if f10iter.DataSource.S3DataSource.S3URI != nil {
						f10elemf3f1.SetS3Uri(*f10iter.DataSource.S3DataSource.S3URI)
					}
					f10elemf3.SetS3DataSource(f10elemf3f1)
				}
				f10elem.SetDataSource(f10elemf3)
			}
			if f10iter.InputMode != nil {
				f10elem.SetInputMode(*f10iter.InputMode)
			}
			if f10iter.RecordWrapperType != nil {
				f10elem.SetRecordWrapperType(*f10iter.RecordWrapperType)
			}
			if f10iter.ShuffleConfig != nil {
				f10elemf6 := &svcsdk.ShuffleConfig{}
				if f10iter.ShuffleConfig.Seed != nil {
					f10elemf6.SetSeed(*f10iter.ShuffleConfig.Seed)
				}
				f10elem.SetShuffleConfig(f10elemf6)
			}
			f10 = append(f10, f10elem)
		}
		res.SetInputDataConfig(f10)
	}
	if r.ko.Spec.OutputDataConfig != nil {
		f11 := &svcsdk.OutputDataConfig{}
		if r.ko.Spec.OutputDataConfig.KMSKeyID != nil {
			f11.SetKmsKeyId(*r.ko.Spec.OutputDataConfig.KMSKeyID)
		}
		if r.ko.Spec.OutputDataConfig.S3OutputPath != nil {
			f11.SetS3OutputPath(*r.ko.Spec.OutputDataConfig.S3OutputPath)
		}
		res.SetOutputDataConfig(f11)
	}
	if r.ko.Spec.ProfilerConfig != nil {
		f12 := &svcsdk.ProfilerConfig{}
		if r.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds != nil {
			f12.SetProfilingIntervalInMilliseconds(*r.ko.Spec.ProfilerConfig.ProfilingIntervalInMilliseconds)
		}
		if r.ko.Spec.ProfilerConfig.ProfilingParameters != nil {
			f12f1 := map[string]*string{}
			for f12f1key, f12f1valiter := range r.ko.Spec.ProfilerConfig.ProfilingParameters {
				var f12f1val string
				f12f1val = *f12f1valiter
				f12f1[f12f1key] = &f12f1val
			}
			f12.SetProfilingParameters(f12f1)
		}
		if r.ko.Spec.ProfilerConfig.S3OutputPath != nil {
			f12.SetS3OutputPath(*r.ko.Spec.ProfilerConfig.S3OutputPath)
		}
		res.SetProfilerConfig(f12)
	}
	if r.ko.Spec.ProfilerRuleConfigurations != nil {
		f13 := []*svcsdk.ProfilerRuleConfiguration{}
		for _, f13iter := range r.ko.Spec.ProfilerRuleConfigurations {
			f13elem := &svcsdk.ProfilerRuleConfiguration{}
			if f13iter.InstanceType != nil {
				f13elem.SetInstanceType(*f13iter.InstanceType)
			}
			if f13iter.LocalPath != nil {
				f13elem.SetLocalPath(*f13iter.LocalPath)
			}
			if f13iter.RuleConfigurationName != nil {
				f13elem.SetRuleConfigurationName(*f13iter.RuleConfigurationName)
			}
			if f13iter.RuleEvaluatorImage != nil {
				f13elem.SetRuleEvaluatorImage(*f13iter.RuleEvaluatorImage)
			}
			if f13iter.RuleParameters != nil {
				f13elemf4 := map[string]*string{}
				for f13elemf4key, f13elemf4valiter := range f13iter.RuleParameters {
					var f13elemf4val string
					f13elemf4val = *f13elemf4valiter
					f13elemf4[f13elemf4key] = &f13elemf4val
				}
				f13elem.SetRuleParameters(f13elemf4)
			}
			if f13iter.S3OutputPath != nil {
				f13elem.SetS3OutputPath(*f13iter.S3OutputPath)
			}
			if f13iter.VolumeSizeInGB != nil {
				f13elem.SetVolumeSizeInGB(*f13iter.VolumeSizeInGB)
			}
			f13 = append(f13, f13elem)
		}
		res.SetProfilerRuleConfigurations(f13)
	}
	if r.ko.Spec.ResourceConfig != nil {
		f14 := &svcsdk.ResourceConfig{}
		if r.ko.Spec.ResourceConfig.InstanceCount != nil {
			f14.SetInstanceCount(*r.ko.Spec.ResourceConfig.InstanceCount)
		}
		if r.ko.Spec.ResourceConfig.InstanceType != nil {
			f14.SetInstanceType(*r.ko.Spec.ResourceConfig.InstanceType)
		}
		if r.ko.Spec.ResourceConfig.VolumeKMSKeyID != nil {
			f14.SetVolumeKmsKeyId(*r.ko.Spec.ResourceConfig.VolumeKMSKeyID)
		}
		if r.ko.Spec.ResourceConfig.VolumeSizeInGB != nil {
			f14.SetVolumeSizeInGB(*r.ko.Spec.ResourceConfig.VolumeSizeInGB)
		}
		res.SetResourceConfig(f14)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}
	if r.ko.Spec.StoppingCondition != nil {
		f16 := &svcsdk.StoppingCondition{}
		if r.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil {
			f16.SetMaxRuntimeInSeconds(*r.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
		}
		if r.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds != nil {
			f16.SetMaxWaitTimeInSeconds(*r.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds)
		}
		res.SetStoppingCondition(f16)
	}
	if r.ko.Spec.Tags != nil {
		f17 := []*svcsdk.Tag{}
		for _, f17iter := range r.ko.Spec.Tags {
			f17elem := &svcsdk.Tag{}
			if f17iter.Key != nil {
				f17elem.SetKey(*f17iter.Key)
			}
			if f17iter.Value != nil {
				f17elem.SetValue(*f17iter.Value)
			}
			f17 = append(f17, f17elem)
		}
		res.SetTags(f17)
	}
	if r.ko.Spec.TensorBoardOutputConfig != nil {
		f18 := &svcsdk.TensorBoardOutputConfig{}
		if r.ko.Spec.TensorBoardOutputConfig.LocalPath != nil {
			f18.SetLocalPath(*r.ko.Spec.TensorBoardOutputConfig.LocalPath)
		}
		if r.ko.Spec.TensorBoardOutputConfig.S3OutputPath != nil {
			f18.SetS3OutputPath(*r.ko.Spec.TensorBoardOutputConfig.S3OutputPath)
		}
		res.SetTensorBoardOutputConfig(f18)
	}
	if r.ko.Spec.TrainingJobName != nil {
		res.SetTrainingJobName(*r.ko.Spec.TrainingJobName)
	}
	if r.ko.Spec.VPCConfig != nil {
		f20 := &svcsdk.VpcConfig{}
		if r.ko.Spec.VPCConfig.SecurityGroupIDs != nil {
			f20f0 := []*string{}
			for _, f20f0iter := range r.ko.Spec.VPCConfig.SecurityGroupIDs {
				var f20f0elem string
				f20f0elem = *f20f0iter
				f20f0 = append(f20f0, &f20f0elem)
			}
			f20.SetSecurityGroupIds(f20f0)
		}
		if r.ko.Spec.VPCConfig.Subnets != nil {
			f20f1 := []*string{}
			for _, f20f1iter := range r.ko.Spec.VPCConfig.Subnets {
				var f20f1elem string
				f20f1elem = *f20f1iter
				f20f1 = append(f20f1, &f20f1elem)
			}
			f20.SetSubnets(f20f1)
		}
		res.SetVpcConfig(f20)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	latestStatus := r.ko.Status.TrainingJobStatus
	if latestStatus != nil {
		if *latestStatus == svcsdk.TrainingJobStatusStopping {
			return r, requeueWaitWhileDeleting
		}

		// Call StopTrainingJob only if the job is InProgress, otherwise just
		// return nil to mark the resource Unmanaged
		if *latestStatus != svcsdk.TrainingJobStatusInProgress {
			return r, err
		}
	}
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.StopTrainingJobOutput
	_ = resp
	resp, err = rm.sdkapi.StopTrainingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "StopTrainingJob", err)

	if err == nil {
		if observed, err := rm.sdkFind(ctx, r); err != ackerr.NotFound {
			if err != nil {
				return nil, err
			}
			r.SetStatus(observed)
			return r, requeueWaitWhileDeleting
		}
	}

	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.StopTrainingJobInput, error) {
	res := &svcsdk.StopTrainingJobInput{}

	if r.ko.Spec.TrainingJobName != nil {
		res.SetTrainingJobName(*r.ko.Spec.TrainingJobName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.TrainingJob,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "ResourceNotFound",
		"ResourceInUse",
		"InvalidParameterCombination",
		"InvalidParameterValue",
		"MissingParameter":
		return true
	default:
		return false
	}
}
