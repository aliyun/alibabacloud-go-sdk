// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanAdditionalListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalTypes(v []*string) *ListVirusScanAdditionalListsRequest
	GetAdditionalTypes() []*string
	SetDevType(v string) *ListVirusScanAdditionalListsRequest
	GetDevType() *string
	SetListDetail(v string) *ListVirusScanAdditionalListsRequest
	GetListDetail() *string
	SetListIds(v []*string) *ListVirusScanAdditionalListsRequest
	GetListIds() []*string
	SetListType(v string) *ListVirusScanAdditionalListsRequest
	GetListType() *string
}

type ListVirusScanAdditionalListsRequest struct {
	AdditionalTypes []*string `json:"AdditionalTypes,omitempty" xml:"AdditionalTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// windows
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
	// example:
	//
	// .tmp
	ListDetail *string   `json:"ListDetail,omitempty" xml:"ListDetail,omitempty"`
	ListIds    []*string `json:"ListIds,omitempty" xml:"ListIds,omitempty" type:"Repeated"`
	// example:
	//
	// Whitelist
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
}

func (s ListVirusScanAdditionalListsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanAdditionalListsRequest) GoString() string {
	return s.String()
}

func (s *ListVirusScanAdditionalListsRequest) GetAdditionalTypes() []*string {
	return s.AdditionalTypes
}

func (s *ListVirusScanAdditionalListsRequest) GetDevType() *string {
	return s.DevType
}

func (s *ListVirusScanAdditionalListsRequest) GetListDetail() *string {
	return s.ListDetail
}

func (s *ListVirusScanAdditionalListsRequest) GetListIds() []*string {
	return s.ListIds
}

func (s *ListVirusScanAdditionalListsRequest) GetListType() *string {
	return s.ListType
}

func (s *ListVirusScanAdditionalListsRequest) SetAdditionalTypes(v []*string) *ListVirusScanAdditionalListsRequest {
	s.AdditionalTypes = v
	return s
}

func (s *ListVirusScanAdditionalListsRequest) SetDevType(v string) *ListVirusScanAdditionalListsRequest {
	s.DevType = &v
	return s
}

func (s *ListVirusScanAdditionalListsRequest) SetListDetail(v string) *ListVirusScanAdditionalListsRequest {
	s.ListDetail = &v
	return s
}

func (s *ListVirusScanAdditionalListsRequest) SetListIds(v []*string) *ListVirusScanAdditionalListsRequest {
	s.ListIds = v
	return s
}

func (s *ListVirusScanAdditionalListsRequest) SetListType(v string) *ListVirusScanAdditionalListsRequest {
	s.ListType = &v
	return s
}

func (s *ListVirusScanAdditionalListsRequest) Validate() error {
	return dara.Validate(s)
}
