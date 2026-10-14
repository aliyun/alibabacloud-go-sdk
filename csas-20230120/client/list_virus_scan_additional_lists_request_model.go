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
	// Filters by matching dimension. Duplicate values are not allowed. If this parameter is not specified, all matching dimensions are returned.
	AdditionalTypes []*string `json:"AdditionalTypes,omitempty" xml:"AdditionalTypes,omitempty" type:"Repeated"`
	// The operating system type for which the list takes effect. Valid values:
	//
	// - **windows**: Windows.
	//
	// - **macOS**: macOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// windows
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
	// The content of the list entry, using exact match. This parameter is used to locate a specific list entry by content and does not support fuzzy match.
	//
	// example:
	//
	// .tmp
	ListDetail *string `json:"ListDetail,omitempty" xml:"ListDetail,omitempty"`
	// Filters by list entry ID. Duplicate values are not allowed. If this parameter is not specified, no filtering by list entry ID is applied.
	ListIds []*string `json:"ListIds,omitempty" xml:"ListIds,omitempty" type:"Repeated"`
	// Filters by list type. If this parameter is not specified, both blacklists and whitelists are returned. Valid values:
	//
	// - **Blacklist**: Blacklist. Files that match are directly identified as virus files.
	//
	// - **Whitelist**: Whitelist. Files that match are excluded from virus detection.
	//
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
