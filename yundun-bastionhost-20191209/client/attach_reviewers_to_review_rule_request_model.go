// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachReviewersToReviewRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AttachReviewersToReviewRuleRequest
	GetInstanceId() *string
	SetPrincipals(v []*AttachReviewersToReviewRuleRequestPrincipals) *AttachReviewersToReviewRuleRequest
	GetPrincipals() []*AttachReviewersToReviewRuleRequestPrincipals
	SetRegionId(v string) *AttachReviewersToReviewRuleRequest
	GetRegionId() *string
	SetReviewRuleId(v int64) *AttachReviewersToReviewRuleRequest
	GetReviewRuleId() *int64
}

type AttachReviewersToReviewRuleRequest struct {
	// The ID of the Bastionhost instance.
	//
	// > You can call the DescribeInstances operation to obtain the Bastionhost instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2ve7h00a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The information about the auditors to be associated with the audit policy.
	Principals []*AttachReviewersToReviewRuleRequestPrincipals `json:"Principals,omitempty" xml:"Principals,omitempty" type:"Repeated"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the audit policy.
	//
	// > You can invoke the ListReviewRules operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewRuleId *int64 `json:"ReviewRuleId,omitempty" xml:"ReviewRuleId,omitempty"`
}

func (s AttachReviewersToReviewRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachReviewersToReviewRuleRequest) GoString() string {
	return s.String()
}

func (s *AttachReviewersToReviewRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachReviewersToReviewRuleRequest) GetPrincipals() []*AttachReviewersToReviewRuleRequestPrincipals {
	return s.Principals
}

func (s *AttachReviewersToReviewRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachReviewersToReviewRuleRequest) GetReviewRuleId() *int64 {
	return s.ReviewRuleId
}

func (s *AttachReviewersToReviewRuleRequest) SetInstanceId(v string) *AttachReviewersToReviewRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachReviewersToReviewRuleRequest) SetPrincipals(v []*AttachReviewersToReviewRuleRequestPrincipals) *AttachReviewersToReviewRuleRequest {
	s.Principals = v
	return s
}

func (s *AttachReviewersToReviewRuleRequest) SetRegionId(v string) *AttachReviewersToReviewRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AttachReviewersToReviewRuleRequest) SetReviewRuleId(v int64) *AttachReviewersToReviewRuleRequest {
	s.ReviewRuleId = &v
	return s
}

func (s *AttachReviewersToReviewRuleRequest) Validate() error {
	if s.Principals != nil {
		for _, item := range s.Principals {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachReviewersToReviewRuleRequestPrincipals struct {
	// The ID of the RAM user or the ARN of the RAM role of the auditor to be associated. If PrincipalType is set to RamUser, specify the RAM user ID. If PrincipalType is set to RamRole, specify the ARN of the RAM role.
	//
	// example:
	//
	// 160****9214
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The display name of the RAM user or RAM role of the associated auditor.
	//
	// example:
	//
	// name
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the auditor to be associated. Valid values:
	//
	// - **RamUser**: Resource Access Management (RAM) user.
	//
	// - **RamRole**: RAM role.
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s AttachReviewersToReviewRuleRequestPrincipals) String() string {
	return dara.Prettify(s)
}

func (s AttachReviewersToReviewRuleRequestPrincipals) GoString() string {
	return s.String()
}

func (s *AttachReviewersToReviewRuleRequestPrincipals) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *AttachReviewersToReviewRuleRequestPrincipals) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *AttachReviewersToReviewRuleRequestPrincipals) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *AttachReviewersToReviewRuleRequestPrincipals) SetPrincipalId(v string) *AttachReviewersToReviewRuleRequestPrincipals {
	s.PrincipalId = &v
	return s
}

func (s *AttachReviewersToReviewRuleRequestPrincipals) SetPrincipalName(v string) *AttachReviewersToReviewRuleRequestPrincipals {
	s.PrincipalName = &v
	return s
}

func (s *AttachReviewersToReviewRuleRequestPrincipals) SetPrincipalType(v string) *AttachReviewersToReviewRuleRequestPrincipals {
	s.PrincipalType = &v
	return s
}

func (s *AttachReviewersToReviewRuleRequestPrincipals) Validate() error {
	return dara.Validate(s)
}
