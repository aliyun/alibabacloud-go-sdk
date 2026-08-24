// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanAdditionalListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalLists(v []*ListVirusScanAdditionalListsResponseBodyAdditionalLists) *ListVirusScanAdditionalListsResponseBody
	GetAdditionalLists() []*ListVirusScanAdditionalListsResponseBodyAdditionalLists
	SetRequestId(v string) *ListVirusScanAdditionalListsResponseBody
	GetRequestId() *string
}

type ListVirusScanAdditionalListsResponseBody struct {
	AdditionalLists []*ListVirusScanAdditionalListsResponseBodyAdditionalLists `json:"AdditionalLists,omitempty" xml:"AdditionalLists,omitempty" type:"Repeated"`
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVirusScanAdditionalListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanAdditionalListsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanAdditionalListsResponseBody) GetAdditionalLists() []*ListVirusScanAdditionalListsResponseBodyAdditionalLists {
	return s.AdditionalLists
}

func (s *ListVirusScanAdditionalListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanAdditionalListsResponseBody) SetAdditionalLists(v []*ListVirusScanAdditionalListsResponseBodyAdditionalLists) *ListVirusScanAdditionalListsResponseBody {
	s.AdditionalLists = v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBody) SetRequestId(v string) *ListVirusScanAdditionalListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBody) Validate() error {
	if s.AdditionalLists != nil {
		for _, item := range s.AdditionalLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanAdditionalListsResponseBodyAdditionalLists struct {
	// example:
	//
	// FileSuffix
	AdditionalType *string                                                         `json:"AdditionalType,omitempty" xml:"AdditionalType,omitempty"`
	Lists          []*ListVirusScanAdditionalListsResponseBodyAdditionalListsLists `json:"Lists,omitempty" xml:"Lists,omitempty" type:"Repeated"`
}

func (s ListVirusScanAdditionalListsResponseBodyAdditionalLists) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanAdditionalListsResponseBodyAdditionalLists) GoString() string {
	return s.String()
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalLists) GetAdditionalType() *string {
	return s.AdditionalType
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalLists) GetLists() []*ListVirusScanAdditionalListsResponseBodyAdditionalListsLists {
	return s.Lists
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalLists) SetAdditionalType(v string) *ListVirusScanAdditionalListsResponseBodyAdditionalLists {
	s.AdditionalType = &v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalLists) SetLists(v []*ListVirusScanAdditionalListsResponseBodyAdditionalListsLists) *ListVirusScanAdditionalListsResponseBodyAdditionalLists {
	s.Lists = v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalLists) Validate() error {
	if s.Lists != nil {
		for _, item := range s.Lists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanAdditionalListsResponseBodyAdditionalListsLists struct {
	ListDetail []*ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail `json:"ListDetail,omitempty" xml:"ListDetail,omitempty" type:"Repeated"`
	// example:
	//
	// Whitelist
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
}

func (s ListVirusScanAdditionalListsResponseBodyAdditionalListsLists) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanAdditionalListsResponseBodyAdditionalListsLists) GoString() string {
	return s.String()
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsLists) GetListDetail() []*ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail {
	return s.ListDetail
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsLists) GetListType() *string {
	return s.ListType
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsLists) SetListDetail(v []*ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) *ListVirusScanAdditionalListsResponseBodyAdditionalListsLists {
	s.ListDetail = v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsLists) SetListType(v string) *ListVirusScanAdditionalListsResponseBodyAdditionalListsLists {
	s.ListType = &v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsLists) Validate() error {
	if s.ListDetail != nil {
		for _, item := range s.ListDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail struct {
	// example:
	//
	// 2026-08-21 10:24:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// .tmp
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// additional-list-4d7b1e9a6c38****
	ListId *string `json:"ListId,omitempty" xml:"ListId,omitempty"`
}

func (s ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) GoString() string {
	return s.String()
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) GetDetail() *string {
	return s.Detail
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) GetListId() *string {
	return s.ListId
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) SetCreateTime(v string) *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail {
	s.CreateTime = &v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) SetDetail(v string) *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail {
	s.Detail = &v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) SetListId(v string) *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail {
	s.ListId = &v
	return s
}

func (s *ListVirusScanAdditionalListsResponseBodyAdditionalListsListsListDetail) Validate() error {
	return dara.Validate(s)
}
