// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEditingProjectMaterialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialIds(v string) *SetEditingProjectMaterialsRequest
	GetMaterialIds() *string
	SetOwnerAccount(v string) *SetEditingProjectMaterialsRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetEditingProjectMaterialsRequest
	GetOwnerId() *string
	SetProjectId(v string) *SetEditingProjectMaterialsRequest
	GetProjectId() *string
	SetResourceOwnerAccount(v string) *SetEditingProjectMaterialsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SetEditingProjectMaterialsRequest
	GetResourceOwnerId() *string
}

type SetEditingProjectMaterialsRequest struct {
	// The ID of the media asset. You can specify IDs of media assets such as videos, images, or auxiliary media assets. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 9e3101bf24bf41c*****123318788ca
	MaterialIds  *string `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// fb2101bf24bf4*****754cb318787dc
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetEditingProjectMaterialsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *SetEditingProjectMaterialsRequest) GetMaterialIds() *string {
	return s.MaterialIds
}

func (s *SetEditingProjectMaterialsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetEditingProjectMaterialsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetEditingProjectMaterialsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SetEditingProjectMaterialsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetEditingProjectMaterialsRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SetEditingProjectMaterialsRequest) SetMaterialIds(v string) *SetEditingProjectMaterialsRequest {
	s.MaterialIds = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetOwnerAccount(v string) *SetEditingProjectMaterialsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetOwnerId(v string) *SetEditingProjectMaterialsRequest {
	s.OwnerId = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetProjectId(v string) *SetEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetResourceOwnerAccount(v string) *SetEditingProjectMaterialsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetResourceOwnerId(v string) *SetEditingProjectMaterialsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) Validate() error {
	return dara.Validate(s)
}
