// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListId(v string) *DeleteDomainMetaRequest
	GetListId() *string
	SetListType(v string) *DeleteDomainMetaRequest
	GetListType() *string
}

type DeleteDomainMetaRequest struct {
	// The list ID, which is a unique business identifier used for policy references and create, update, and delete operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// ladl-6f1exxxxx6ab59
	ListId *string `json:"ListId,omitempty" xml:"ListId,omitempty"`
	// The list type. Valid values:
	//
	// - la_domain_white_list: domain name whitelist.
	//
	// - la_domain_black_list: domain name blacklist.
	//
	// This parameter is required.
	//
	// example:
	//
	// la_domain_white_list
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
}

func (s DeleteDomainMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainMetaRequest) GetListId() *string {
	return s.ListId
}

func (s *DeleteDomainMetaRequest) GetListType() *string {
	return s.ListType
}

func (s *DeleteDomainMetaRequest) SetListId(v string) *DeleteDomainMetaRequest {
	s.ListId = &v
	return s
}

func (s *DeleteDomainMetaRequest) SetListType(v string) *DeleteDomainMetaRequest {
	s.ListType = &v
	return s
}

func (s *DeleteDomainMetaRequest) Validate() error {
	return dara.Validate(s)
}
