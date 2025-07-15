// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisitorChatMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *ListVisitorChatMessagesRequest
	GetAccessChannelId() *string
	SetAccessToken(v string) *ListVisitorChatMessagesRequest
	GetAccessToken() *string
	SetEndTime(v string) *ListVisitorChatMessagesRequest
	GetEndTime() *string
	SetInstanceId(v string) *ListVisitorChatMessagesRequest
	GetInstanceId() *string
	SetNextPageToken(v string) *ListVisitorChatMessagesRequest
	GetNextPageToken() *string
	SetPageSize(v int32) *ListVisitorChatMessagesRequest
	GetPageSize() *int32
	SetSortOrder(v string) *ListVisitorChatMessagesRequest
	GetSortOrder() *string
	SetStartTime(v int64) *ListVisitorChatMessagesRequest
	GetStartTime() *int64
	SetVisitorId(v string) *ListVisitorChatMessagesRequest
	GetVisitorId() *string
}

type ListVisitorChatMessagesRequest struct {
	// example:
	//
	// cf584733-***-***-9699-cb77aa3b7aa6
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// example:
	//
	// 9XYGTGWtq2wXzVikKuip_zeVGl6O4VJ-l-*-*-JPofhap4P7fAevuE=
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// 1650316799000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1737193352340::7463707254.EAUNIT
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// DESC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// example:
	//
	// 1647325450000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// fcd020fe-****-1a272a174a7d
	VisitorId *string `json:"VisitorId,omitempty" xml:"VisitorId,omitempty"`
}

func (s ListVisitorChatMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVisitorChatMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListVisitorChatMessagesRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ListVisitorChatMessagesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListVisitorChatMessagesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListVisitorChatMessagesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVisitorChatMessagesRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListVisitorChatMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVisitorChatMessagesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListVisitorChatMessagesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListVisitorChatMessagesRequest) GetVisitorId() *string {
	return s.VisitorId
}

func (s *ListVisitorChatMessagesRequest) SetAccessChannelId(v string) *ListVisitorChatMessagesRequest {
	s.AccessChannelId = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) SetAccessToken(v string) *ListVisitorChatMessagesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) SetEndTime(v string) *ListVisitorChatMessagesRequest {
	s.EndTime = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) SetInstanceId(v string) *ListVisitorChatMessagesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) SetNextPageToken(v string) *ListVisitorChatMessagesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) SetPageSize(v int32) *ListVisitorChatMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) SetSortOrder(v string) *ListVisitorChatMessagesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) SetStartTime(v int64) *ListVisitorChatMessagesRequest {
	s.StartTime = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) SetVisitorId(v string) *ListVisitorChatMessagesRequest {
	s.VisitorId = &v
	return s
}

func (s *ListVisitorChatMessagesRequest) Validate() error {
	return dara.Validate(s)
}
