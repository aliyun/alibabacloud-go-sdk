// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *MoveResourceGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *MoveResourceGroupRequest
	GetDryRun() *bool
	SetNewResourceGroupId(v string) *MoveResourceGroupRequest
	GetNewResourceGroupId() *string
	SetOwnerAccount(v string) *MoveResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MoveResourceGroupRequest
	GetOwnerId() *int64
	SetResourceId(v string) *MoveResourceGroupRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *MoveResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MoveResourceGroupRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *MoveResourceGroupRequest
	GetResourceType() *string
}

type MoveResourceGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e456****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false:*	- performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the resource group to which you want to move the CEN instance or bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzqqurtfg****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-nye53d7p3hzyu4****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **CEN**: CEN instance
	//
	// 	- **bandwidthpackage**: bandwidth plan
	//
	// This parameter is required.
	//
	// example:
	//
	// cen
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *MoveResourceGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *MoveResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *MoveResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MoveResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MoveResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *MoveResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MoveResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MoveResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *MoveResourceGroupRequest) SetClientToken(v string) *MoveResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *MoveResourceGroupRequest) SetDryRun(v bool) *MoveResourceGroupRequest {
	s.DryRun = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetOwnerAccount(v string) *MoveResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MoveResourceGroupRequest) SetOwnerId(v int64) *MoveResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceOwnerAccount(v string) *MoveResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceOwnerId(v int64) *MoveResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
