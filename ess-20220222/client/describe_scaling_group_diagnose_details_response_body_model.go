// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupDiagnoseDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetails(v []*DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) *DescribeScalingGroupDiagnoseDetailsResponseBody
	GetDetails() []*DescribeScalingGroupDiagnoseDetailsResponseBodyDetails
	SetRequestId(v string) *DescribeScalingGroupDiagnoseDetailsResponseBody
	GetRequestId() *string
}

type DescribeScalingGroupDiagnoseDetailsResponseBody struct {
	// The diagnostic reports.
	Details []*DescribeScalingGroupDiagnoseDetailsResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// ID of the request
	//
	// example:
	//
	// 688B18B8-FB1E-42EB-A1ED-7F55B090****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScalingGroupDiagnoseDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDiagnoseDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBody) GetDetails() []*DescribeScalingGroupDiagnoseDetailsResponseBodyDetails {
	return s.Details
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBody) SetDetails(v []*DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) *DescribeScalingGroupDiagnoseDetailsResponseBody {
	s.Details = v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBody) SetRequestId(v string) *DescribeScalingGroupDiagnoseDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBody) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScalingGroupDiagnoseDetailsResponseBodyDetails struct {
	// The type of the diagnostic item. Valid values:
	//
	// 	- AccountArrearage: Checks whether your Alibaba Cloud account has overdue payments.
	//
	// 	- AccountNotEnoughBalance: Checks whether the balance of your Alibaba Cloud account is sufficient.
	//
	// 	- ElasticStrength: Checks whether the instance types that are specified in the scaling configuration are sufficient.
	//
	// 	- VSwitch: Checks whether the vSwitch is available. If the specified vSwitch is deleted, specify an existing vSwitch for the scaling group.
	//
	// 	- SecurityGroup: Checks whether the security group is available. If the specified security group is deleted, specify an existing security group for the scaling group.
	//
	// 	- KeyPair: Checks whether the key pair is available. If the specified key pair is deleted, specify another key pair for the scaling group.
	//
	// 	- SlbBackendServerQuota: Checks whether the number of ECS instances that are added to the default server group and the vServer groups of the SLB instances associated with the scaling group has reached the upper limit.
	//
	// 	- AlbBackendServerQuota: Checks whether the number of ECS instances that are attached to the ALB instances of the scaling group has reached the upper limit.
	//
	// 	- NlbBackendServerQuota: Checks whether the number of ECS instances that are attached to the NLB instances of the scaling group has reached the upper limit.
	//
	// example:
	//
	// SecurityGroup
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// The error code of the diagnostic item. Valid values:
	//
	// 	- VSwitchIdNotFound: The vSwitch does not exist.
	//
	// 	- SecurityGroupNotFound: The security group does not exist.
	//
	// 	- KeyPairNotFound: The key pair does not exist.
	//
	// 	- SlbBackendServerQuotaExceeded: The number of ECS instances that are added to the default server group and the vServer groups of the SLB instances associated with the scaling group has reached the upper limit.
	//
	// 	- AlbBackendServerQuotaExceeded: The number of ECS instances that are attached to the ALB instances of the scaling group has reached the upper limit.
	//
	// 	- NlbBackendServerQuotaExceeded: The number of ECS instances that are attached to the NLB instances of the scaling group has reached the upper limit.
	//
	// 	- AccountArrearage: Your account has an overdue payment.
	//
	// 	- AccountNotEnoughBalance: The balance of your Alibaba Cloud account is insufficient.
	//
	// 	- ElasticStrengthAlert: The inventory levels are lower than required.
	//
	// example:
	//
	// AccountArrearage
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The resource ID corresponding to the diagnostic result.
	//
	// example:
	//
	// sg-280ih****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Status of the diagnostic item. Possible values:
	//
	// - Normal: The diagnostic result is normal.
	//
	// - Warn: The diagnostic result is a warning.
	//
	// - Critical: The diagnostic result is critical.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) GetStatus() *string {
	return s.Status
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) SetDiagnoseType(v string) *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails {
	s.DiagnoseType = &v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) SetErrorCode(v string) *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails {
	s.ErrorCode = &v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) SetResourceId(v string) *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails {
	s.ResourceId = &v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) SetStatus(v string) *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails {
	s.Status = &v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponseBodyDetails) Validate() error {
	return dara.Validate(s)
}
