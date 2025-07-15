// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListLiveMessageGroupsRequest
	GetAppId() *string
	SetDataCenter(v string) *ListLiveMessageGroupsRequest
	GetDataCenter() *string
	SetGroupStatus(v int32) *ListLiveMessageGroupsRequest
	GetGroupStatus() *int32
	SetNextPageToken(v int64) *ListLiveMessageGroupsRequest
	GetNextPageToken() *int64
	SetSortType(v int32) *ListLiveMessageGroupsRequest
	GetSortType() *int32
}

type ListLiveMessageGroupsRequest struct {
	// The application ID.
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
	// The status of the groups to query. Default value: 0. Valid values:
	//
	// 	- 0: all groups
	//
	// 	- 1: existing groups
	//
	// 	- 2: deleted groups
	//
	// example:
	//
	// 1
	GroupStatus *int32 `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	// The starting page number for the query. If you leave this parameter empty, the query starts from the first page.
	//
	// example:
	//
	// 1001
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The sort order based on the time when the groups were created. Valid values:
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

func (s ListLiveMessageGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListLiveMessageGroupsRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListLiveMessageGroupsRequest) GetGroupStatus() *int32 {
	return s.GroupStatus
}

func (s *ListLiveMessageGroupsRequest) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListLiveMessageGroupsRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListLiveMessageGroupsRequest) SetAppId(v string) *ListLiveMessageGroupsRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveMessageGroupsRequest) SetDataCenter(v string) *ListLiveMessageGroupsRequest {
	s.DataCenter = &v
	return s
}

func (s *ListLiveMessageGroupsRequest) SetGroupStatus(v int32) *ListLiveMessageGroupsRequest {
	s.GroupStatus = &v
	return s
}

func (s *ListLiveMessageGroupsRequest) SetNextPageToken(v int64) *ListLiveMessageGroupsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListLiveMessageGroupsRequest) SetSortType(v int32) *ListLiveMessageGroupsRequest {
	s.SortType = &v
	return s
}

func (s *ListLiveMessageGroupsRequest) Validate() error {
	return dara.Validate(s)
}
