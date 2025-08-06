// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectMaterialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialIds(v string) *DeleteEditingProjectMaterialsRequest
	GetMaterialIds() *string
	SetMaterialType(v string) *DeleteEditingProjectMaterialsRequest
	GetMaterialType() *string
	SetOwnerAccount(v string) *DeleteEditingProjectMaterialsRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DeleteEditingProjectMaterialsRequest
	GetOwnerId() *string
	SetProjectId(v string) *DeleteEditingProjectMaterialsRequest
	GetProjectId() *string
	SetResourceOwnerAccount(v string) *DeleteEditingProjectMaterialsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *DeleteEditingProjectMaterialsRequest
	GetResourceOwnerId() *string
}

type DeleteEditingProjectMaterialsRequest struct {
	// The material ID. Separate multiple material IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 9e3101bf24bf41c*****123318788ca
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
	// 198236101*****1093374
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteEditingProjectMaterialsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectMaterialsRequest) GetMaterialIds() *string {
	return s.MaterialIds
}

func (s *DeleteEditingProjectMaterialsRequest) GetMaterialType() *string {
	return s.MaterialType
}

func (s *DeleteEditingProjectMaterialsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteEditingProjectMaterialsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteEditingProjectMaterialsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteEditingProjectMaterialsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteEditingProjectMaterialsRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DeleteEditingProjectMaterialsRequest) SetMaterialIds(v string) *DeleteEditingProjectMaterialsRequest {
	s.MaterialIds = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetMaterialType(v string) *DeleteEditingProjectMaterialsRequest {
	s.MaterialType = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetOwnerAccount(v string) *DeleteEditingProjectMaterialsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetOwnerId(v string) *DeleteEditingProjectMaterialsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetProjectId(v string) *DeleteEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetResourceOwnerAccount(v string) *DeleteEditingProjectMaterialsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetResourceOwnerId(v string) *DeleteEditingProjectMaterialsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) Validate() error {
	return dara.Validate(s)
}
