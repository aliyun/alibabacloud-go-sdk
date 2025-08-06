// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectMaterialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialType(v string) *GetEditingProjectMaterialsRequest
	GetMaterialType() *string
	SetOwnerAccount(v string) *GetEditingProjectMaterialsRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetEditingProjectMaterialsRequest
	GetOwnerId() *string
	SetProjectId(v string) *GetEditingProjectMaterialsRequest
	GetProjectId() *string
	SetResourceOwnerAccount(v string) *GetEditingProjectMaterialsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetEditingProjectMaterialsRequest
	GetResourceOwnerId() *string
	SetType(v string) *GetEditingProjectMaterialsRequest
	GetType() *string
}

type GetEditingProjectMaterialsRequest struct {
	// The type of the material. Valid values:
	//
	// 	- **video**
	//
	// 	- **audio**
	//
	// 	- **image**
	//
	// example:
	//
	// video
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the online editing project. You can use one of the following methods to obtain the ID of the online editing project:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Production Center*	- > **Video Editing*	- to obtain the ID of the specified online editing project.
	//
	// 	- Call the **AddEditingProject*	- operation. The value of the response parameter ProjectId is the ID of the specified online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1982361011093374****
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the material. Valid values:
	//
	// 	- **video**
	//
	// 	- **audio**
	//
	// example:
	//
	// video
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetEditingProjectMaterialsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsRequest) GetMaterialType() *string {
	return s.MaterialType
}

func (s *GetEditingProjectMaterialsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetEditingProjectMaterialsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetEditingProjectMaterialsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetEditingProjectMaterialsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetEditingProjectMaterialsRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetEditingProjectMaterialsRequest) GetType() *string {
	return s.Type
}

func (s *GetEditingProjectMaterialsRequest) SetMaterialType(v string) *GetEditingProjectMaterialsRequest {
	s.MaterialType = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetOwnerAccount(v string) *GetEditingProjectMaterialsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetOwnerId(v string) *GetEditingProjectMaterialsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetProjectId(v string) *GetEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetResourceOwnerAccount(v string) *GetEditingProjectMaterialsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetResourceOwnerId(v string) *GetEditingProjectMaterialsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetType(v string) *GetEditingProjectMaterialsRequest {
	s.Type = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) Validate() error {
	return dara.Validate(s)
}
