// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListId(v string) *UpdateDomainMetaRequest
	GetListId() *string
	SetListType(v string) *UpdateDomainMetaRequest
	GetListType() *string
	SetName(v string) *UpdateDomainMetaRequest
	GetName() *string
}

type UpdateDomainMetaRequest struct {
	// The list ID. This is a unique business identifier used for policy references and add, delete, or modify operations.
	//
	// example:
	//
	// ladl-6f1exxxxx6ab59
	ListId *string `json:"ListId,omitempty" xml:"ListId,omitempty"`
	// The list type.
	//
	// example:
	//
	// la_domain_white_list
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The list name. Maximum length: 32 characters.
	//
	// example:
	//
	// new_office_domain
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDomainMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainMetaRequest) GetListId() *string {
	return s.ListId
}

func (s *UpdateDomainMetaRequest) GetListType() *string {
	return s.ListType
}

func (s *UpdateDomainMetaRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDomainMetaRequest) SetListId(v string) *UpdateDomainMetaRequest {
	s.ListId = &v
	return s
}

func (s *UpdateDomainMetaRequest) SetListType(v string) *UpdateDomainMetaRequest {
	s.ListType = &v
	return s
}

func (s *UpdateDomainMetaRequest) SetName(v string) *UpdateDomainMetaRequest {
	s.Name = &v
	return s
}

func (s *UpdateDomainMetaRequest) Validate() error {
	return dara.Validate(s)
}
