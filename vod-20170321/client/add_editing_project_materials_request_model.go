// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEditingProjectMaterialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialIds(v string) *AddEditingProjectMaterialsRequest
	GetMaterialIds() *string
	SetMaterialType(v string) *AddEditingProjectMaterialsRequest
	GetMaterialType() *string
	SetOwnerAccount(v string) *AddEditingProjectMaterialsRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *AddEditingProjectMaterialsRequest
	GetOwnerId() *string
	SetProjectId(v string) *AddEditingProjectMaterialsRequest
	GetProjectId() *string
	SetResourceOwnerAccount(v string) *AddEditingProjectMaterialsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *AddEditingProjectMaterialsRequest
	GetResourceOwnerId() *string
}

type AddEditingProjectMaterialsRequest struct {
	// Separate multiple material IDs with commas (,). You can specify up to 10 IDs.
	//
	// >  If you specify multiple materials, make sure that the materials are of the same type as specified in MaterialType.
	//
	// This parameter is required.
	//
	// example:
	//
	// d3251979f9fd41f2acb29ccda5a6f772
	MaterialIds *string `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty"`
	// The type of the material. Valid values:
	//
	// 	- **video**
	//
	// 	- **audio**
	//
	// 	- **image**
	//
	// This parameter is required.
	//
	// example:
	//
	// video
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// afa31b483b5c41609185de0e1b790579
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddEditingProjectMaterialsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsRequest) GetMaterialIds() *string {
	return s.MaterialIds
}

func (s *AddEditingProjectMaterialsRequest) GetMaterialType() *string {
	return s.MaterialType
}

func (s *AddEditingProjectMaterialsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddEditingProjectMaterialsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *AddEditingProjectMaterialsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *AddEditingProjectMaterialsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddEditingProjectMaterialsRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *AddEditingProjectMaterialsRequest) SetMaterialIds(v string) *AddEditingProjectMaterialsRequest {
	s.MaterialIds = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) SetMaterialType(v string) *AddEditingProjectMaterialsRequest {
	s.MaterialType = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) SetOwnerAccount(v string) *AddEditingProjectMaterialsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) SetOwnerId(v string) *AddEditingProjectMaterialsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) SetProjectId(v string) *AddEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) SetResourceOwnerAccount(v string) *AddEditingProjectMaterialsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) SetResourceOwnerId(v string) *AddEditingProjectMaterialsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) Validate() error {
	return dara.Validate(s)
}
