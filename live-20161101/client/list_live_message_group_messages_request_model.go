// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListLiveMessageGroupMessagesRequest
	GetAppId() *string
	SetBeginTime(v int64) *ListLiveMessageGroupMessagesRequest
	GetBeginTime() *int64
	SetDataCenter(v string) *ListLiveMessageGroupMessagesRequest
	GetDataCenter() *string
	SetEndTime(v int64) *ListLiveMessageGroupMessagesRequest
	GetEndTime() *int64
	SetGroupId(v string) *ListLiveMessageGroupMessagesRequest
	GetGroupId() *string
	SetMsgType(v int64) *ListLiveMessageGroupMessagesRequest
	GetMsgType() *int64
	SetNextPageToken(v int64) *ListLiveMessageGroupMessagesRequest
	GetNextPageToken() *int64
	SetPageSize(v int32) *ListLiveMessageGroupMessagesRequest
	GetPageSize() *int32
	SetSortType(v int32) *ListLiveMessageGroupMessagesRequest
	GetSortType() *int32
}

type ListLiveMessageGroupMessagesRequest struct {
	// The ID of the interactive messaging application to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds. If you leave this parameter empty, the earliest available time is used.
	//
	// example:
	//
	// 1697783235
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds. If you leave this parameter empty, the latest available time is used.
	//
	// example:
	//
	// 1698301635
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the group to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The type of messages that you want to query. If you leave this parameter empty, all types of messages are queried.
	//
	// example:
	//
	// 1
	MsgType *int64 `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
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
	// The sort order based on the time when the messages were sent. Valid values:
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

func (s ListLiveMessageGroupMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupMessagesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListLiveMessageGroupMessagesRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListLiveMessageGroupMessagesRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListLiveMessageGroupMessagesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListLiveMessageGroupMessagesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListLiveMessageGroupMessagesRequest) GetMsgType() *int64 {
	return s.MsgType
}

func (s *ListLiveMessageGroupMessagesRequest) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListLiveMessageGroupMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveMessageGroupMessagesRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListLiveMessageGroupMessagesRequest) SetAppId(v string) *ListLiveMessageGroupMessagesRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) SetBeginTime(v int64) *ListLiveMessageGroupMessagesRequest {
	s.BeginTime = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) SetDataCenter(v string) *ListLiveMessageGroupMessagesRequest {
	s.DataCenter = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) SetEndTime(v int64) *ListLiveMessageGroupMessagesRequest {
	s.EndTime = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) SetGroupId(v string) *ListLiveMessageGroupMessagesRequest {
	s.GroupId = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) SetMsgType(v int64) *ListLiveMessageGroupMessagesRequest {
	s.MsgType = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) SetNextPageToken(v int64) *ListLiveMessageGroupMessagesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) SetPageSize(v int32) *ListLiveMessageGroupMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) SetSortType(v int32) *ListLiveMessageGroupMessagesRequest {
	s.SortType = &v
	return s
}

func (s *ListLiveMessageGroupMessagesRequest) Validate() error {
	return dara.Validate(s)
}
