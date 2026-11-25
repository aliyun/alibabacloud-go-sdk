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
	// Details of the diagnostic report.
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
	// Item type for diagnostics. Possible values:
	//
	// - AccountArrearage: Checks if the user\\"s account is in arrears.
	//
	// - AccountNotEnoughBalance: Checks the account balance.
	//
	// - ElasticStrength: Checks the inventory health of instance types corresponding to the scaling group configuration.
	//
	// - VSwitch: Checks if the switch is available, for example, whether it has been deleted.
	//
	// - SecurityGroup: Checks if the security group is available, for example, whether it has been deleted.
	//
	// - KeyPair: Checks if the key pair is available, for example, whether it has been deleted.
	//
	// - SlbBackendServerQuota: Checks if the number of ECS instances mounted to the default SLB group and virtual server group exceeds the quota.
	//
	// - AlbBackendServerQuota: Checks if the number of ECS instances mounted to the ALB group exceeds the quota.
	//
	// - NlbBackendServerQuota: Checks if the number of ECS instances mounted to the NLB group exceeds the quota.
	//
	// example:
	//
	// SecurityGroup
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// Error code of the diagnostic item. Possible values:
	//
	// - VSwitchIdNotFound: The VSwitch does not exist.
	//
	// - SecurityGroupNotFound: The security group does not exist.
	//
	// - KeyPairNotFound: The key pair does not exist.
	//
	// - SlbBackendServerQuotaExceeded: The number of ECS instances mounted on the backend of the SLB default group and virtual server group exceeds the quota.
	//
	// - AlbBackendServerQuotaExceeded: The number of ECS instances mounted on the backend of the ALB group exceeds the quota.
	//
	// - NlbBackendServerQuotaExceeded: The number of ECS instances mounted on the backend of the NLB group exceeds the quota.
	//
	// - AccountArrearage: The account is in arrears.
	//
	// - AccountNotEnoughBalance: The account balance is insufficient.
	//
	// - ElasticStrengthAlert: The inventory health is poor.
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
