// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDomainItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemIds(v []*int64) *BatchDeleteDomainItemsRequest
	GetItemIds() []*int64
	SetListId(v string) *BatchDeleteDomainItemsRequest
	GetListId() *string
	SetListType(v string) *BatchDeleteDomainItemsRequest
	GetListType() *string
}

type BatchDeleteDomainItemsRequest struct {
	// The IDs of domain name list entries.
	ItemIds []*int64 `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" type:"Repeated"`
	// The list ID. This is the unique business identifier used for policy references and CRUD operations.
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

func (s BatchDeleteDomainItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDomainItemsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDomainItemsRequest) GetItemIds() []*int64 {
	return s.ItemIds
}

func (s *BatchDeleteDomainItemsRequest) GetListId() *string {
	return s.ListId
}

func (s *BatchDeleteDomainItemsRequest) GetListType() *string {
	return s.ListType
}

func (s *BatchDeleteDomainItemsRequest) SetItemIds(v []*int64) *BatchDeleteDomainItemsRequest {
	s.ItemIds = v
	return s
}

func (s *BatchDeleteDomainItemsRequest) SetListId(v string) *BatchDeleteDomainItemsRequest {
	s.ListId = &v
	return s
}

func (s *BatchDeleteDomainItemsRequest) SetListType(v string) *BatchDeleteDomainItemsRequest {
	s.ListType = &v
	return s
}

func (s *BatchDeleteDomainItemsRequest) Validate() error {
	return dara.Validate(s)
}
