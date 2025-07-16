// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommentNum(v int64) *StatisticsReportResponseBody
	GetCommentNum() *int64
	SetCommentUserNum(v int64) *StatisticsReportResponseBody
	GetCommentUserNum() *int64
	SetLikeNum(v int64) *StatisticsReportResponseBody
	GetLikeNum() *int64
	SetReadNum(v int64) *StatisticsReportResponseBody
	GetReadNum() *int64
	SetRequestId(v string) *StatisticsReportResponseBody
	GetRequestId() *string
}

type StatisticsReportResponseBody struct {
	// example:
	//
	// 1
	CommentNum *int64 `json:"commentNum,omitempty" xml:"commentNum,omitempty"`
	// example:
	//
	// 1
	CommentUserNum *int64 `json:"commentUserNum,omitempty" xml:"commentUserNum,omitempty"`
	// example:
	//
	// 2
	LikeNum *int64 `json:"likeNum,omitempty" xml:"likeNum,omitempty"`
	// example:
	//
	// 3
	ReadNum *int64 `json:"readNum,omitempty" xml:"readNum,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StatisticsReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StatisticsReportResponseBody) GoString() string {
	return s.String()
}

func (s *StatisticsReportResponseBody) GetCommentNum() *int64 {
	return s.CommentNum
}

func (s *StatisticsReportResponseBody) GetCommentUserNum() *int64 {
	return s.CommentUserNum
}

func (s *StatisticsReportResponseBody) GetLikeNum() *int64 {
	return s.LikeNum
}

func (s *StatisticsReportResponseBody) GetReadNum() *int64 {
	return s.ReadNum
}

func (s *StatisticsReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StatisticsReportResponseBody) SetCommentNum(v int64) *StatisticsReportResponseBody {
	s.CommentNum = &v
	return s
}

func (s *StatisticsReportResponseBody) SetCommentUserNum(v int64) *StatisticsReportResponseBody {
	s.CommentUserNum = &v
	return s
}

func (s *StatisticsReportResponseBody) SetLikeNum(v int64) *StatisticsReportResponseBody {
	s.LikeNum = &v
	return s
}

func (s *StatisticsReportResponseBody) SetReadNum(v int64) *StatisticsReportResponseBody {
	s.ReadNum = &v
	return s
}

func (s *StatisticsReportResponseBody) SetRequestId(v string) *StatisticsReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *StatisticsReportResponseBody) Validate() error {
	return dara.Validate(s)
}
