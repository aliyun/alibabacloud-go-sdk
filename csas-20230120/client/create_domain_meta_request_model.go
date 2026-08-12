// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListType(v string) *CreateDomainMetaRequest
	GetListType() *string
	SetName(v string) *CreateDomainMetaRequest
	GetName() *string
}

type CreateDomainMetaRequest struct {
	// The list type.
	//
	// example:
	//
	// la_domain_black_list
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The list name. Maximum length: 32 characters.
	//
	// example:
	//
	// office_domain
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDomainMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainMetaRequest) GetListType() *string {
	return s.ListType
}

func (s *CreateDomainMetaRequest) GetName() *string {
	return s.Name
}

func (s *CreateDomainMetaRequest) SetListType(v string) *CreateDomainMetaRequest {
	s.ListType = &v
	return s
}

func (s *CreateDomainMetaRequest) SetName(v string) *CreateDomainMetaRequest {
	s.Name = &v
	return s
}

func (s *CreateDomainMetaRequest) Validate() error {
	return dara.Validate(s)
}
