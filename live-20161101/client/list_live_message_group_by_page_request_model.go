// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListLiveMessageGroupByPageRequest
	GetAppId() *string
	SetDataCenter(v string) *ListLiveMessageGroupByPageRequest
	GetDataCenter() *string
	SetGroupStatus(v int32) *ListLiveMessageGroupByPageRequest
	GetGroupStatus() *int32
	SetPageNumber(v int32) *ListLiveMessageGroupByPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLiveMessageGroupByPageRequest
	GetPageSize() *int32
	SetSortType(v int32) *ListLiveMessageGroupByPageRequest
	GetSortType() *int32
}

type ListLiveMessageGroupByPageRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. This value must be the same as the data center specified in [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html). Currently supported data centers are Shanghai (cn-shanghai) and Singapore (ap-southeast-1).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The group status. Default value: 0. Valid values:
	//
	// - 0: Queries all groups.
	//
	// - 1: Queries groups that are not deleted.
	//
	// - 2: Queries deleted groups.
	//
	// example:
	//
	// 1
	GroupStatus *int32 `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	// The page number. Valid values: [1,10000\\].
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Valid values: [1,50\\].
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort order, based on group creation time. Valid values:
	//
	// - 1: Ascending order.
	//
	// - 2: Descending order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListLiveMessageGroupByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupByPageRequest) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupByPageRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListLiveMessageGroupByPageRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListLiveMessageGroupByPageRequest) GetGroupStatus() *int32 {
	return s.GroupStatus
}

func (s *ListLiveMessageGroupByPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLiveMessageGroupByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveMessageGroupByPageRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListLiveMessageGroupByPageRequest) SetAppId(v string) *ListLiveMessageGroupByPageRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveMessageGroupByPageRequest) SetDataCenter(v string) *ListLiveMessageGroupByPageRequest {
	s.DataCenter = &v
	return s
}

func (s *ListLiveMessageGroupByPageRequest) SetGroupStatus(v int32) *ListLiveMessageGroupByPageRequest {
	s.GroupStatus = &v
	return s
}

func (s *ListLiveMessageGroupByPageRequest) SetPageNumber(v int32) *ListLiveMessageGroupByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLiveMessageGroupByPageRequest) SetPageSize(v int32) *ListLiveMessageGroupByPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveMessageGroupByPageRequest) SetSortType(v int32) *ListLiveMessageGroupByPageRequest {
	s.SortType = &v
	return s
}

func (s *ListLiveMessageGroupByPageRequest) Validate() error {
	return dara.Validate(s)
}
