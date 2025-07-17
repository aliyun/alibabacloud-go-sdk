// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministrators(v []*string) *UpdateMetaCollectionRequest
	GetAdministrators() []*string
	SetDescription(v string) *UpdateMetaCollectionRequest
	GetDescription() *string
	SetId(v string) *UpdateMetaCollectionRequest
	GetId() *string
	SetName(v string) *UpdateMetaCollectionRequest
	GetName() *string
}

type UpdateMetaCollectionRequest struct {
	// The collection administrator IDs. This parameter is available only for data albums. The administrator must be an account within the same tenant.
	Administrators []*string `json:"Administrators,omitempty" xml:"Administrators,omitempty" type:"Repeated"`
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

func (s UpdateMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaCollectionRequest) GetAdministrators() []*string {
	return s.Administrators
}

func (s *UpdateMetaCollectionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMetaCollectionRequest) GetId() *string {
	return s.Id
}

func (s *UpdateMetaCollectionRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMetaCollectionRequest) SetAdministrators(v []*string) *UpdateMetaCollectionRequest {
	s.Administrators = v
	return s
}

func (s *UpdateMetaCollectionRequest) SetDescription(v string) *UpdateMetaCollectionRequest {
	s.Description = &v
	return s
}

func (s *UpdateMetaCollectionRequest) SetId(v string) *UpdateMetaCollectionRequest {
	s.Id = &v
	return s
}

func (s *UpdateMetaCollectionRequest) SetName(v string) *UpdateMetaCollectionRequest {
	s.Name = &v
	return s
}

func (s *UpdateMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
