// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainItems(v []*ListDomainItemsResponseBodyDomainItems) *ListDomainItemsResponseBody
	GetDomainItems() []*ListDomainItemsResponseBodyDomainItems
	SetRequestId(v string) *ListDomainItemsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListDomainItemsResponseBody
	GetTotalNum() *int32
}

type ListDomainItemsResponseBody struct {
	DomainItems []*ListDomainItemsResponseBodyDomainItems `json:"DomainItems,omitempty" xml:"DomainItems,omitempty" type:"Repeated"`
	// example:
	//
	// 019F68B5-2D0D-5399-9BB2-D81C13C2C05D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 13
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDomainItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainItemsResponseBody) GetDomainItems() []*ListDomainItemsResponseBodyDomainItems {
	return s.DomainItems
}

func (s *ListDomainItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainItemsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListDomainItemsResponseBody) SetDomainItems(v []*ListDomainItemsResponseBodyDomainItems) *ListDomainItemsResponseBody {
	s.DomainItems = v
	return s
}

func (s *ListDomainItemsResponseBody) SetRequestId(v string) *ListDomainItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainItemsResponseBody) SetTotalNum(v int32) *ListDomainItemsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListDomainItemsResponseBody) Validate() error {
	if s.DomainItems != nil {
		for _, item := range s.DomainItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDomainItemsResponseBodyDomainItems struct {
	// example:
	//
	// 2026-08-01 10:20:30
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2026-08-01 10:20:30
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 499
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// example:
	//
	// www.example.com
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s ListDomainItemsResponseBodyDomainItems) String() string {
	return dara.Prettify(s)
}

func (s ListDomainItemsResponseBodyDomainItems) GoString() string {
	return s.String()
}

func (s *ListDomainItemsResponseBodyDomainItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDomainItemsResponseBodyDomainItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDomainItemsResponseBodyDomainItems) GetItemId() *int64 {
	return s.ItemId
}

func (s *ListDomainItemsResponseBodyDomainItems) GetItemValue() *string {
	return s.ItemValue
}

func (s *ListDomainItemsResponseBodyDomainItems) SetGmtCreate(v string) *ListDomainItemsResponseBodyDomainItems {
	s.GmtCreate = &v
	return s
}

func (s *ListDomainItemsResponseBodyDomainItems) SetGmtModified(v string) *ListDomainItemsResponseBodyDomainItems {
	s.GmtModified = &v
	return s
}

func (s *ListDomainItemsResponseBodyDomainItems) SetItemId(v int64) *ListDomainItemsResponseBodyDomainItems {
	s.ItemId = &v
	return s
}

func (s *ListDomainItemsResponseBodyDomainItems) SetItemValue(v string) *ListDomainItemsResponseBodyDomainItems {
	s.ItemValue = &v
	return s
}

func (s *ListDomainItemsResponseBodyDomainItems) Validate() error {
	return dara.Validate(s)
}
