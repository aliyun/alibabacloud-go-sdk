// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsListByTypeReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasMore(v bool) *StatisticsListByTypeReportResponseBody
	GetHasMore() *bool
	SetNextCursor(v int64) *StatisticsListByTypeReportResponseBody
	GetNextCursor() *int64
	SetRequestId(v string) *StatisticsListByTypeReportResponseBody
	GetRequestId() *string
	SetUseridList(v []*string) *StatisticsListByTypeReportResponseBody
	GetUseridList() []*string
}

type StatisticsListByTypeReportResponseBody struct {
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 12312131231
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId  *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	UseridList []*string `json:"useridList,omitempty" xml:"useridList,omitempty" type:"Repeated"`
}

func (s StatisticsListByTypeReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StatisticsListByTypeReportResponseBody) GoString() string {
	return s.String()
}

func (s *StatisticsListByTypeReportResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *StatisticsListByTypeReportResponseBody) GetNextCursor() *int64 {
	return s.NextCursor
}

func (s *StatisticsListByTypeReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StatisticsListByTypeReportResponseBody) GetUseridList() []*string {
	return s.UseridList
}

func (s *StatisticsListByTypeReportResponseBody) SetHasMore(v bool) *StatisticsListByTypeReportResponseBody {
	s.HasMore = &v
	return s
}

func (s *StatisticsListByTypeReportResponseBody) SetNextCursor(v int64) *StatisticsListByTypeReportResponseBody {
	s.NextCursor = &v
	return s
}

func (s *StatisticsListByTypeReportResponseBody) SetRequestId(v string) *StatisticsListByTypeReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *StatisticsListByTypeReportResponseBody) SetUseridList(v []*string) *StatisticsListByTypeReportResponseBody {
	s.UseridList = v
	return s
}

func (s *StatisticsListByTypeReportResponseBody) Validate() error {
	return dara.Validate(s)
}
