// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetEditingProjectRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetEditingProjectRequest
	GetOwnerId() *string
	SetProjectId(v string) *GetEditingProjectRequest
	GetProjectId() *string
	SetResourceOwnerAccount(v string) *GetEditingProjectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetEditingProjectRequest
	GetResourceOwnerId() *string
}

type GetEditingProjectRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// fb2101bf24b27*****54cb318787dc
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetEditingProjectRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetEditingProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetEditingProjectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetEditingProjectRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetEditingProjectRequest) SetOwnerAccount(v string) *GetEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEditingProjectRequest) SetOwnerId(v string) *GetEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEditingProjectRequest) SetProjectId(v string) *GetEditingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectRequest) SetResourceOwnerAccount(v string) *GetEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEditingProjectRequest) SetResourceOwnerId(v string) *GetEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
