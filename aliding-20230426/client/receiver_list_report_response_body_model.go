// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiverListReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *ReceiverListReportResponseBody
	GetHasMore() *bool
	SetNextCursor(v int64) *ReceiverListReportResponseBody
	GetNextCursor() *int64
	SetRequestId(v string) *ReceiverListReportResponseBody
	GetRequestId() *string
	SetUseridList(v []*string) *ReceiverListReportResponseBody
	GetUseridList() []*string
}

type ReceiverListReportResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 132131312312
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId  *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	UseridList []*string `json:"useridList,omitempty" xml:"useridList,omitempty" type:"Repeated"`
}

func (s ReceiverListReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReceiverListReportResponseBody) GoString() string {
	return s.String()
}

func (s *ReceiverListReportResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ReceiverListReportResponseBody) GetNextCursor() *int64 {
	return s.NextCursor
}

func (s *ReceiverListReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReceiverListReportResponseBody) GetUseridList() []*string {
	return s.UseridList
}

func (s *ReceiverListReportResponseBody) SetHasMore(v bool) *ReceiverListReportResponseBody {
	s.HasMore = &v
	return s
}

func (s *ReceiverListReportResponseBody) SetNextCursor(v int64) *ReceiverListReportResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ReceiverListReportResponseBody) SetRequestId(v string) *ReceiverListReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReceiverListReportResponseBody) SetUseridList(v []*string) *ReceiverListReportResponseBody {
	s.UseridList = v
	return s
}

func (s *ReceiverListReportResponseBody) Validate() error {
	return dara.Validate(s)
}
