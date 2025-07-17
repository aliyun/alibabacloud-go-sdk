// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCollectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministratorsShrink(v string) *UpdateMetaCollectionShrinkRequest
	GetAdministratorsShrink() *string
	SetDescription(v string) *UpdateMetaCollectionShrinkRequest
	GetDescription() *string
	SetId(v string) *UpdateMetaCollectionShrinkRequest
	GetId() *string
	SetName(v string) *UpdateMetaCollectionShrinkRequest
	GetName() *string
}

type UpdateMetaCollectionShrinkRequest struct {
	// The collection administrator IDs. This parameter is available only for data albums. The administrator must be an account within the same tenant.
	AdministratorsShrink *string `json:"Administrators,omitempty" xml:"Administrators,omitempty"`
	// example:
	//
	// new comment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The collection ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// category.123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// new_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMetaCollectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCollectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaCollectionShrinkRequest) GetAdministratorsShrink() *string {
	return s.AdministratorsShrink
}

func (s *UpdateMetaCollectionShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMetaCollectionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateMetaCollectionShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMetaCollectionShrinkRequest) SetAdministratorsShrink(v string) *UpdateMetaCollectionShrinkRequest {
	s.AdministratorsShrink = &v
	return s
}

func (s *UpdateMetaCollectionShrinkRequest) SetDescription(v string) *UpdateMetaCollectionShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateMetaCollectionShrinkRequest) SetId(v string) *UpdateMetaCollectionShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateMetaCollectionShrinkRequest) SetName(v string) *UpdateMetaCollectionShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateMetaCollectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
