// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckInstanceResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *ListCheckInstanceResultRequest
	GetCheckId() *int64
	SetCurrentPage(v int32) *ListCheckInstanceResultRequest
	GetCurrentPage() *int32
	SetInstanceIdKey(v string) *ListCheckInstanceResultRequest
	GetInstanceIdKey() *string
	SetInstanceIds(v []*string) *ListCheckInstanceResultRequest
	GetInstanceIds() []*string
	SetInstanceNameKey(v string) *ListCheckInstanceResultRequest
	GetInstanceNameKey() *string
	SetLang(v string) *ListCheckInstanceResultRequest
	GetLang() *string
	SetPageSize(v int32) *ListCheckInstanceResultRequest
	GetPageSize() *int32
	SetRegionIdKey(v string) *ListCheckInstanceResultRequest
	GetRegionIdKey() *string
	SetSortTypes(v []*string) *ListCheckInstanceResultRequest
	GetSortTypes() []*string
	SetStatuses(v []*string) *ListCheckInstanceResultRequest
	GetStatuses() []*string
}

type ListCheckInstanceResultRequest struct {
	// The ID of the check item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-uf64w4q6p9jti5gl****
	InstanceIdKey *string `json:"InstanceIdKey,omitempty" xml:"InstanceIdKey,omitempty"`
	// The instance IDs of cloud services.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The name of the instance.
	//
	// example:
	//
	// i-uf64w4q6p9jti5gl****
	InstanceNameKey *string `json:"InstanceNameKey,omitempty" xml:"InstanceNameKey,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-qingdao
	RegionIdKey *string `json:"RegionIdKey,omitempty" xml:"RegionIdKey,omitempty"`
	// The types of the conditions based on which the check items are sorted.
	SortTypes []*string `json:"SortTypes,omitempty" xml:"SortTypes,omitempty" type:"Repeated"`
	// The statuses of check items.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListCheckInstanceResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultRequest) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListCheckInstanceResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckInstanceResultRequest) GetInstanceIdKey() *string {
	return s.InstanceIdKey
}

func (s *ListCheckInstanceResultRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListCheckInstanceResultRequest) GetInstanceNameKey() *string {
	return s.InstanceNameKey
}

func (s *ListCheckInstanceResultRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCheckInstanceResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckInstanceResultRequest) GetRegionIdKey() *string {
	return s.RegionIdKey
}

func (s *ListCheckInstanceResultRequest) GetSortTypes() []*string {
	return s.SortTypes
}

func (s *ListCheckInstanceResultRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListCheckInstanceResultRequest) SetCheckId(v int64) *ListCheckInstanceResultRequest {
	s.CheckId = &v
	return s
}

func (s *ListCheckInstanceResultRequest) SetCurrentPage(v int32) *ListCheckInstanceResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckInstanceResultRequest) SetInstanceIdKey(v string) *ListCheckInstanceResultRequest {
	s.InstanceIdKey = &v
	return s
}

func (s *ListCheckInstanceResultRequest) SetInstanceIds(v []*string) *ListCheckInstanceResultRequest {
	s.InstanceIds = v
	return s
}

func (s *ListCheckInstanceResultRequest) SetInstanceNameKey(v string) *ListCheckInstanceResultRequest {
	s.InstanceNameKey = &v
	return s
}

func (s *ListCheckInstanceResultRequest) SetLang(v string) *ListCheckInstanceResultRequest {
	s.Lang = &v
	return s
}

func (s *ListCheckInstanceResultRequest) SetPageSize(v int32) *ListCheckInstanceResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckInstanceResultRequest) SetRegionIdKey(v string) *ListCheckInstanceResultRequest {
	s.RegionIdKey = &v
	return s
}

func (s *ListCheckInstanceResultRequest) SetSortTypes(v []*string) *ListCheckInstanceResultRequest {
	s.SortTypes = v
	return s
}

func (s *ListCheckInstanceResultRequest) SetStatuses(v []*string) *ListCheckInstanceResultRequest {
	s.Statuses = v
	return s
}

func (s *ListCheckInstanceResultRequest) Validate() error {
	return dara.Validate(s)
}
