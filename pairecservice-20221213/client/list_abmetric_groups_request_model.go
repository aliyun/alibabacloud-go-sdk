// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABMetricGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListABMetricGroupsRequest
	GetInstanceId() *string
	SetOrder(v string) *ListABMetricGroupsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListABMetricGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListABMetricGroupsRequest
	GetPageSize() *int32
	SetRealtime(v bool) *ListABMetricGroupsRequest
	GetRealtime() *bool
	SetSceneId(v string) *ListABMetricGroupsRequest
	GetSceneId() *string
	SetSortBy(v string) *ListABMetricGroupsRequest
	GetSortBy() *string
}

type ListABMetricGroupsRequest struct {
	// The instance ID. You can call the [ListInstances](https://help.aliyun.com/document_detail/2411819.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sort order. Valid values:
	//
	// - ASC: ascending.
	//
	// - DESC: descending.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to return only real-time A/B metric groups.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// The scene ID. You can call the [ListScenes](\\(~~2402581~~\\)) operation to obtain this ID.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The field to sort the results by. Valid values:
	//
	// - GmtCreateTime: creation time.
	//
	// - GmtModifiedTime: modification time.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListABMetricGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListABMetricGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListABMetricGroupsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListABMetricGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListABMetricGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListABMetricGroupsRequest) GetRealtime() *bool {
	return s.Realtime
}

func (s *ListABMetricGroupsRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListABMetricGroupsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListABMetricGroupsRequest) SetInstanceId(v string) *ListABMetricGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetOrder(v string) *ListABMetricGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetPageNumber(v int32) *ListABMetricGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetPageSize(v int32) *ListABMetricGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetRealtime(v bool) *ListABMetricGroupsRequest {
	s.Realtime = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetSceneId(v string) *ListABMetricGroupsRequest {
	s.SceneId = &v
	return s
}

func (s *ListABMetricGroupsRequest) SetSortBy(v string) *ListABMetricGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListABMetricGroupsRequest) Validate() error {
	return dara.Validate(s)
}
