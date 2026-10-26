// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachProjectManagersToProjectAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v int64) *AttachProjectManagersToProjectAuthorizationRequest
	GetAuthorizationId() *int64
	SetInstanceId(v string) *AttachProjectManagersToProjectAuthorizationRequest
	GetInstanceId() *string
	SetPrincipals(v []*AttachProjectManagersToProjectAuthorizationRequestPrincipals) *AttachProjectManagersToProjectAuthorizationRequest
	GetPrincipals() []*AttachProjectManagersToProjectAuthorizationRequestPrincipals
	SetRegionId(v string) *AttachProjectManagersToProjectAuthorizationRequest
	GetRegionId() *string
}

type AttachProjectManagersToProjectAuthorizationRequest struct {
	// The project authorization ID.
	//
	// > You can call the ListProjectAuthorizations operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The information about the project administrators specified in the project authorization.
	Principals []*AttachProjectManagersToProjectAuthorizationRequestPrincipals `json:"Principals,omitempty" xml:"Principals,omitempty" type:"Repeated"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachProjectManagersToProjectAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachProjectManagersToProjectAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) GetPrincipals() []*AttachProjectManagersToProjectAuthorizationRequestPrincipals {
	return s.Principals
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) SetAuthorizationId(v int64) *AttachProjectManagersToProjectAuthorizationRequest {
	s.AuthorizationId = &v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) SetInstanceId(v string) *AttachProjectManagersToProjectAuthorizationRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) SetPrincipals(v []*AttachProjectManagersToProjectAuthorizationRequestPrincipals) *AttachProjectManagersToProjectAuthorizationRequest {
	s.Principals = v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) SetRegionId(v string) *AttachProjectManagersToProjectAuthorizationRequest {
	s.RegionId = &v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationRequest) Validate() error {
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

type AttachProjectManagersToProjectAuthorizationRequestPrincipals struct {
	// When PrincipalType is set to RamUser, specify the ID of the RAM user to be designated as the project administrator. When PrincipalType is set to RamRole, specify the ARN of the RAM role to be designated as the project administrator.
	//
	// example:
	//
	// 160****2519
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The display name of the RAM user or RAM role designated as the project administrator.
	//
	// example:
	//
	// name
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the project administrator. Valid values:
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

func (s AttachProjectManagersToProjectAuthorizationRequestPrincipals) String() string {
	return dara.Prettify(s)
}

func (s AttachProjectManagersToProjectAuthorizationRequestPrincipals) GoString() string {
	return s.String()
}

func (s *AttachProjectManagersToProjectAuthorizationRequestPrincipals) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *AttachProjectManagersToProjectAuthorizationRequestPrincipals) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *AttachProjectManagersToProjectAuthorizationRequestPrincipals) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *AttachProjectManagersToProjectAuthorizationRequestPrincipals) SetPrincipalId(v string) *AttachProjectManagersToProjectAuthorizationRequestPrincipals {
	s.PrincipalId = &v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationRequestPrincipals) SetPrincipalName(v string) *AttachProjectManagersToProjectAuthorizationRequestPrincipals {
	s.PrincipalName = &v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationRequestPrincipals) SetPrincipalType(v string) *AttachProjectManagersToProjectAuthorizationRequestPrincipals {
	s.PrincipalType = &v
	return s
}

func (s *AttachProjectManagersToProjectAuthorizationRequestPrincipals) Validate() error {
	return dara.Validate(s)
}
