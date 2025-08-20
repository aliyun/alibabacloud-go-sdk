// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackGroupDriftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetectStackGroupDriftRequest
	GetClientToken() *string
	SetOperationPreferences(v map[string]interface{}) *DetectStackGroupDriftRequest
	GetOperationPreferences() map[string]interface{}
	SetRegionId(v string) *DetectStackGroupDriftRequest
	GetRegionId() *string
	SetStackGroupName(v string) *DetectStackGroupDriftRequest
	GetStackGroupName() *string
}

type DetectStackGroupDriftRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The value can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The operation settings, in JSON format. The following fields are supported:
	//
	// 	- FailureToleranceCount
	//
	// The maximum number of stack group operation failures that can occur. In a stack group operation, if the total number of failures does not exceed the FailureToleranceCount value, the operation succeeds. Otherwise, the operation fails.
	//
	// If FailureToleranceCount is not specified, the default value 0 is used. You can specify one of FailureToleranceCount or FailureTolerancePercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 0 to 20.
	//
	// 	- FailureTolerancePercentage
	//
	// The percentage of stack group operation failures that can occur. In a stack group operation, if the percentage of failures does not exceed the FailureTolerancePercentage value, the operation succeeds. Otherwise, the operation fails.
	//
	// You can specify one of FailureToleranceCount or FailureTolerancePercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 0 to 100.
	//
	// 	- MaxConcurrentCount
	//
	// The maximum number of target accounts in which a drift detection operation can be performed at a time.
	//
	// You can specify one of MaxConcurrentCount or MaxConcurrentPercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 1 to 20.
	//
	// 	- MaxConcurrentPercentage
	//
	// The maximum percentage of target accounts in which a drift detection operation can be performed at a time.
	//
	// You can specify one of MaxConcurrentCount or MaxConcurrentPercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// {"FailureToleranceCount": 1, "MaxConcurrentCount": 2}
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035#doc-api-ROS-DescribeRegions~~ "Queries the DescribeRegions list of a region.") operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique in a region.
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s DetectStackGroupDriftRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectStackGroupDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetectStackGroupDriftRequest) GetOperationPreferences() map[string]interface{} {
	return s.OperationPreferences
}

func (s *DetectStackGroupDriftRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetectStackGroupDriftRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *DetectStackGroupDriftRequest) SetClientToken(v string) *DetectStackGroupDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackGroupDriftRequest) SetOperationPreferences(v map[string]interface{}) *DetectStackGroupDriftRequest {
	s.OperationPreferences = v
	return s
}

func (s *DetectStackGroupDriftRequest) SetRegionId(v string) *DetectStackGroupDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackGroupDriftRequest) SetStackGroupName(v string) *DetectStackGroupDriftRequest {
	s.StackGroupName = &v
	return s
}

func (s *DetectStackGroupDriftRequest) Validate() error {
	return dara.Validate(s)
}
