// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListTopicRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListTopicRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTopicRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListTopicRequest
	GetRegionId() *string
	SetSortBy(v string) *ListTopicRequest
	GetSortBy() *string
	SetTopicName(v string) *ListTopicRequest
	GetTopicName() *string
	SetWorkspaceId(v int64) *ListTopicRequest
	GetWorkspaceId() *int64
}

type ListTopicRequest struct {
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// GmtModified
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// client_exposure_logs
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// example:
	//
	// 643168
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTopicRequest) GoString() string {
	return s.String()
}

func (s *ListTopicRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTopicRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTopicRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTopicRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTopicRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListTopicRequest) SetOrder(v string) *ListTopicRequest {
	s.Order = &v
	return s
}

func (s *ListTopicRequest) SetPageNumber(v int32) *ListTopicRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTopicRequest) SetPageSize(v int32) *ListTopicRequest {
	s.PageSize = &v
	return s
}

func (s *ListTopicRequest) SetRegionId(v string) *ListTopicRequest {
	s.RegionId = &v
	return s
}

func (s *ListTopicRequest) SetSortBy(v string) *ListTopicRequest {
	s.SortBy = &v
	return s
}

func (s *ListTopicRequest) SetTopicName(v string) *ListTopicRequest {
	s.TopicName = &v
	return s
}

func (s *ListTopicRequest) SetWorkspaceId(v int64) *ListTopicRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTopicRequest) Validate() error {
	return dara.Validate(s)
}
