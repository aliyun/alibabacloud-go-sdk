// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDomainItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainItems(v []*string) *BatchCreateDomainItemsRequest
	GetDomainItems() []*string
	SetListId(v string) *BatchCreateDomainItemsRequest
	GetListId() *string
	SetListType(v string) *BatchCreateDomainItemsRequest
	GetListType() *string
}

type BatchCreateDomainItemsRequest struct {
	// The domain name list.
	DomainItems []*string `json:"DomainItems,omitempty" xml:"DomainItems,omitempty" type:"Repeated"`
	// The list ID. This is a unique business identifier used for policy references and add, delete, and modify operations.
	//
	// example:
	//
	// ladl-61aae0c0ba715e3b
	ListId *string `json:"ListId,omitempty" xml:"ListId,omitempty"`
	// The list type (Blacklist/Whitelist).
	//
	// example:
	//
	// la_domain_white_list
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
}

func (s BatchCreateDomainItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDomainItemsRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateDomainItemsRequest) GetDomainItems() []*string {
	return s.DomainItems
}

func (s *BatchCreateDomainItemsRequest) GetListId() *string {
	return s.ListId
}

func (s *BatchCreateDomainItemsRequest) GetListType() *string {
	return s.ListType
}

func (s *BatchCreateDomainItemsRequest) SetDomainItems(v []*string) *BatchCreateDomainItemsRequest {
	s.DomainItems = v
	return s
}

func (s *BatchCreateDomainItemsRequest) SetListId(v string) *BatchCreateDomainItemsRequest {
	s.ListId = &v
	return s
}

func (s *BatchCreateDomainItemsRequest) SetListType(v string) *BatchCreateDomainItemsRequest {
	s.ListType = &v
	return s
}

func (s *BatchCreateDomainItemsRequest) Validate() error {
	return dara.Validate(s)
}
