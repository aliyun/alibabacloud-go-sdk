// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RollbackApplicationRequest
	GetAppId() *string
	SetAutoEnableApplicationScalingRule(v string) *RollbackApplicationRequest
	GetAutoEnableApplicationScalingRule() *string
	SetBatchWaitTime(v int32) *RollbackApplicationRequest
	GetBatchWaitTime() *int32
	SetMinReadyInstanceRatio(v int32) *RollbackApplicationRequest
	GetMinReadyInstanceRatio() *int32
	SetMinReadyInstances(v int32) *RollbackApplicationRequest
	GetMinReadyInstances() *int32
	SetUpdateStrategy(v string) *RollbackApplicationRequest
	GetUpdateStrategy() *string
	SetVersionId(v string) *RollbackApplicationRequest
	GetVersionId() *string
}

type RollbackApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to automatically enable an auto scaling policy for the application. Take note of the following rules:
	//
	// 	- **true**: turns on Logon-free Sharing
	//
	// 	- **false**: turns off Logon-free Sharing
	//
	// example:
	//
	// true
	AutoEnableApplicationScalingRule *string `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	// The wait time between batches. Unit: seconds.
	//
	// example:
	//
	// 10
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The percentage of the minimum number of available instances. Take note of the following rules:
	//
	// 	- If you set the value to **-1**, the minimum number of available instances is not determined based on this parameter. Default value: -1.
	//
	// 	- If you set the value to a number **from 0 to 100**, the minimum number of available instances is calculated by using the following formula: Current number of instances × (Value of MinReadyInstanceRatio × 100%). The value is the nearest integer rounded up from the calculated result. For example, if the percentage is set to **50**% and five instances are available, the minimum number of available instances is 3.
	//
	// > When both **MinReadyInstance*	- and **MinReadyInstanceRatio*	- are specified and **MinReadyInstanceRatio*	- is set to a number from 0 to 100, the value of **MinReadyInstanceRatio*	- takes precedence.*	- For example, if **MinReadyInstances*	- is set to **5, and **MinReadyInstanceRatio*	- is set to **50**, the minimum number of available instances is set to the nearest integer rounded up from the calculated result of the following formula: Current number of instances × **50%**.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of available instances. Take note of the following rules:
	//
	// 	- If you set the value to **0**, business interruptions occur when the application is updated.
	//
	// 	- If you set the value to \\*\\*-1\\*\\*, the minimum number of available instances is automatically set to a system-recommended value. The value is the nearest integer to which the calculated result of the following formula is rounded up: Current number of instances × 25%. For example, if five instances are available, the minimum number of available instances is calculated by using the following formula: 5 × 25% = 1.25. In this case, the minimum number of available instances is 2.
	//
	// > Make sure that at least one instance is available during application deployment and rollback to prevent business interruptions.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// The deployment policy. If the minimum number of available instances is 1, the value of the **UpdateStrategy*	- parameter is an empty string (""). If the minimum number of available instances is larger than 1, specify this parameter based on your requirements. Examples:
	//
	// 	- Perform canary release for one instance and release the remaining instances in two batches automatically with a one-minute interval between the deployment of each instance:
	//
	//     `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`
	//
	// 	- Perform canary release for one instance and release the remaining instances in two batches manually:
	//
	//     `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`
	//
	// 	- Release the instances in two batches automatically with no interval between the deployment of each instance:
	//
	//     `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// The following table describes the parameters that are used in the preceding statements.
	//
	// 	- **type**: the type of the release policy. Valid values: **GrayBatchUpdate*	- and **BatchUpdate**.
	//
	// 	- **batchUpdate**: the phased release policy.
	//
	//     	- **batch**: the number of release batches.
	//
	//     	- **releaseType**: the processing method for the batches. Valid values: **auto*	- and **manual**.
	//
	//     	- **batchWaitTime**: the interval between release batches. Unit: seconds.
	//
	// 	- **grayUpdate**: the number of release batches in the phased release after a canary release. This parameter is returned only if the **type*	- parameter is set to **GrayBatchUpdate**.
	//
	// example:
	//
	// {"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}
	UpdateStrategy *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	// The ID of the application version. Call the [ListAppVersions](https://help.aliyun.com/document_detail/162054.html) operation to obtain the version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0026ff7f-2b57-4127-bdd0-9bf202bb9****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s RollbackApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackApplicationRequest) GoString() string {
	return s.String()
}

func (s *RollbackApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RollbackApplicationRequest) GetAutoEnableApplicationScalingRule() *string {
	return s.AutoEnableApplicationScalingRule
}

func (s *RollbackApplicationRequest) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *RollbackApplicationRequest) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *RollbackApplicationRequest) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *RollbackApplicationRequest) GetUpdateStrategy() *string {
	return s.UpdateStrategy
}

func (s *RollbackApplicationRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *RollbackApplicationRequest) SetAppId(v string) *RollbackApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RollbackApplicationRequest) SetAutoEnableApplicationScalingRule(v string) *RollbackApplicationRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *RollbackApplicationRequest) SetBatchWaitTime(v int32) *RollbackApplicationRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *RollbackApplicationRequest) SetMinReadyInstanceRatio(v int32) *RollbackApplicationRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *RollbackApplicationRequest) SetMinReadyInstances(v int32) *RollbackApplicationRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *RollbackApplicationRequest) SetUpdateStrategy(v string) *RollbackApplicationRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *RollbackApplicationRequest) SetVersionId(v string) *RollbackApplicationRequest {
	s.VersionId = &v
	return s
}

func (s *RollbackApplicationRequest) Validate() error {
	return dara.Validate(s)
}
