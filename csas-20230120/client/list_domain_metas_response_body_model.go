// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainMetasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainMetas(v []*ListDomainMetasResponseBodyDomainMetas) *ListDomainMetasResponseBody
	GetDomainMetas() []*ListDomainMetasResponseBodyDomainMetas
	SetRequestId(v string) *ListDomainMetasResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListDomainMetasResponseBody
	GetTotalNum() *int32
}

type ListDomainMetasResponseBody struct {
	// The list of domain name lists.
	DomainMetas []*ListDomainMetasResponseBodyDomainMetas `json:"DomainMetas,omitempty" xml:"DomainMetas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D2788E14-8C9F-5FE8-B72F-5ABD033AA27E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of lists that match the specified conditions.
	//
	// example:
	//
	// 34
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDomainMetasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainMetasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainMetasResponseBody) GetDomainMetas() []*ListDomainMetasResponseBodyDomainMetas {
	return s.DomainMetas
}

func (s *ListDomainMetasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainMetasResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListDomainMetasResponseBody) SetDomainMetas(v []*ListDomainMetasResponseBodyDomainMetas) *ListDomainMetasResponseBody {
	s.DomainMetas = v
	return s
}

func (s *ListDomainMetasResponseBody) SetRequestId(v string) *ListDomainMetasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainMetasResponseBody) SetTotalNum(v int32) *ListDomainMetasResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListDomainMetasResponseBody) Validate() error {
	if s.DomainMetas != nil {
		for _, item := range s.DomainMetas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainMetasResponseBodyDomainMetas struct {
	// The time when the list was created.
	//
	// example:
	//
	// 2026-08-01 10:20:30
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the list was last modified.
	//
	// example:
	//
	// 2026-08-02 15:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The number of domain name entries in the list.
	//
	// example:
	//
	// 128
	ItemCount *int64 `json:"ItemCount,omitempty" xml:"ItemCount,omitempty"`
	// The list ID, which is a unique business identifier used for policy references and CRUD operations.
	//
	// example:
	//
	// ladl-8acxxxa0f2a7daf9
	ListId *string `json:"ListId,omitempty" xml:"ListId,omitempty"`
	// The list type.
	//
	// example:
	//
	// la_domain_white_list
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The list name.
	//
	// example:
	//
	// OfficeDomainWhitelist
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// RS_ladl-xxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ListDomainMetasResponseBodyDomainMetas) String() string {
	return dara.Prettify(s)
}

func (s ListDomainMetasResponseBodyDomainMetas) GoString() string {
	return s.String()
}

func (s *ListDomainMetasResponseBodyDomainMetas) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDomainMetasResponseBodyDomainMetas) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDomainMetasResponseBodyDomainMetas) GetItemCount() *int64 {
	return s.ItemCount
}

func (s *ListDomainMetasResponseBodyDomainMetas) GetListId() *string {
	return s.ListId
}

func (s *ListDomainMetasResponseBodyDomainMetas) GetListType() *string {
	return s.ListType
}

func (s *ListDomainMetasResponseBodyDomainMetas) GetName() *string {
	return s.Name
}

func (s *ListDomainMetasResponseBodyDomainMetas) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListDomainMetasResponseBodyDomainMetas) SetGmtCreate(v string) *ListDomainMetasResponseBodyDomainMetas {
	s.GmtCreate = &v
	return s
}

func (s *ListDomainMetasResponseBodyDomainMetas) SetGmtModified(v string) *ListDomainMetasResponseBodyDomainMetas {
	s.GmtModified = &v
	return s
}

func (s *ListDomainMetasResponseBodyDomainMetas) SetItemCount(v int64) *ListDomainMetasResponseBodyDomainMetas {
	s.ItemCount = &v
	return s
}

func (s *ListDomainMetasResponseBodyDomainMetas) SetListId(v string) *ListDomainMetasResponseBodyDomainMetas {
	s.ListId = &v
	return s
}

func (s *ListDomainMetasResponseBodyDomainMetas) SetListType(v string) *ListDomainMetasResponseBodyDomainMetas {
	s.ListType = &v
	return s
}

func (s *ListDomainMetasResponseBodyDomainMetas) SetName(v string) *ListDomainMetasResponseBodyDomainMetas {
	s.Name = &v
	return s
}

func (s *ListDomainMetasResponseBodyDomainMetas) SetResourceId(v string) *ListDomainMetasResponseBodyDomainMetas {
	s.ResourceId = &v
	return s
}

func (s *ListDomainMetasResponseBodyDomainMetas) Validate() error {
	return dara.Validate(s)
}
