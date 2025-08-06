// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteEditingProjectRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DeleteEditingProjectRequest
	GetOwnerId() *string
	SetProjectIds(v string) *DeleteEditingProjectRequest
	GetProjectIds() *string
	SetResourceOwnerAccount(v string) *DeleteEditingProjectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *DeleteEditingProjectRequest
	GetResourceOwnerId() *string
}

type DeleteEditingProjectRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the online editing project. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// fb2101bf24bf41*****cb318787dc
	ProjectIds           *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteEditingProjectRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteEditingProjectRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *DeleteEditingProjectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteEditingProjectRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DeleteEditingProjectRequest) SetOwnerAccount(v string) *DeleteEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteEditingProjectRequest) SetOwnerId(v string) *DeleteEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEditingProjectRequest) SetProjectIds(v string) *DeleteEditingProjectRequest {
	s.ProjectIds = &v
	return s
}

func (s *DeleteEditingProjectRequest) SetResourceOwnerAccount(v string) *DeleteEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteEditingProjectRequest) SetResourceOwnerId(v string) *DeleteEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
