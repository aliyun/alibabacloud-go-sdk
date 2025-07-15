// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListLiveMessageGroupUsersRequest
	GetAppId() *string
	SetDataCenter(v string) *ListLiveMessageGroupUsersRequest
	GetDataCenter() *string
	SetGroupId(v string) *ListLiveMessageGroupUsersRequest
	GetGroupId() *string
	SetNextPageToken(v int64) *ListLiveMessageGroupUsersRequest
	GetNextPageToken() *int64
	SetPageSize(v int32) *ListLiveMessageGroupUsersRequest
	GetPageSize() *int32
	SetSortType(v int32) *ListLiveMessageGroupUsersRequest
	GetSortType() *int32
}

type ListLiveMessageGroupUsersRequest struct {
	// The ID of the interactive messaging application to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the group to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The starting page number for the query. If you leave this parameter empty, the query starts from the first page.
	//
	// example:
	//
	// 1
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The number of entries per page. Valid values: **10 to 50**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort order based on the time when the users joined the group. Valid values:
	//
	// 	- 1: ascending order
	//
	// 	- 2: descending order
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListLiveMessageGroupUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupUsersRequest) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupUsersRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListLiveMessageGroupUsersRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListLiveMessageGroupUsersRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListLiveMessageGroupUsersRequest) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListLiveMessageGroupUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveMessageGroupUsersRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListLiveMessageGroupUsersRequest) SetAppId(v string) *ListLiveMessageGroupUsersRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveMessageGroupUsersRequest) SetDataCenter(v string) *ListLiveMessageGroupUsersRequest {
	s.DataCenter = &v
	return s
}

func (s *ListLiveMessageGroupUsersRequest) SetGroupId(v string) *ListLiveMessageGroupUsersRequest {
	s.GroupId = &v
	return s
}

func (s *ListLiveMessageGroupUsersRequest) SetNextPageToken(v int64) *ListLiveMessageGroupUsersRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListLiveMessageGroupUsersRequest) SetPageSize(v int32) *ListLiveMessageGroupUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveMessageGroupUsersRequest) SetSortType(v int32) *ListLiveMessageGroupUsersRequest {
	s.SortType = &v
	return s
}

func (s *ListLiveMessageGroupUsersRequest) Validate() error {
	return dara.Validate(s)
}
