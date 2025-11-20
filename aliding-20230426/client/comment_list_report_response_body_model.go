// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommentListReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComments(v []*CommentListReportResponseBodyComments) *CommentListReportResponseBody
	GetComments() []*CommentListReportResponseBodyComments
	SetHasMore(v bool) *CommentListReportResponseBody
	GetHasMore() *bool
	SetNextCursor(v int64) *CommentListReportResponseBody
	GetNextCursor() *int64
	SetRequestId(v string) *CommentListReportResponseBody
	GetRequestId() *string
}

type CommentListReportResponseBody struct {
	Comments []*CommentListReportResponseBodyComments `json:"comments,omitempty" xml:"comments,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 1568442466000
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CommentListReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportResponseBody) GoString() string {
	return s.String()
}

func (s *CommentListReportResponseBody) GetComments() []*CommentListReportResponseBodyComments {
	return s.Comments
}

func (s *CommentListReportResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *CommentListReportResponseBody) GetNextCursor() *int64 {
	return s.NextCursor
}

func (s *CommentListReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CommentListReportResponseBody) SetComments(v []*CommentListReportResponseBodyComments) *CommentListReportResponseBody {
	s.Comments = v
	return s
}

func (s *CommentListReportResponseBody) SetHasMore(v bool) *CommentListReportResponseBody {
	s.HasMore = &v
	return s
}

func (s *CommentListReportResponseBody) SetNextCursor(v int64) *CommentListReportResponseBody {
	s.NextCursor = &v
	return s
}

func (s *CommentListReportResponseBody) SetRequestId(v string) *CommentListReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommentListReportResponseBody) Validate() error {
	if s.Comments != nil {
		for _, item := range s.Comments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CommentListReportResponseBodyComments struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1678442466000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 012345
	Userid *string `json:"Userid,omitempty" xml:"Userid,omitempty"`
}

func (s CommentListReportResponseBodyComments) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportResponseBodyComments) GoString() string {
	return s.String()
}

func (s *CommentListReportResponseBodyComments) GetContent() *string {
	return s.Content
}

func (s *CommentListReportResponseBodyComments) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CommentListReportResponseBodyComments) GetUserid() *string {
	return s.Userid
}

func (s *CommentListReportResponseBodyComments) SetContent(v string) *CommentListReportResponseBodyComments {
	s.Content = &v
	return s
}

func (s *CommentListReportResponseBodyComments) SetCreateTime(v string) *CommentListReportResponseBodyComments {
	s.CreateTime = &v
	return s
}

func (s *CommentListReportResponseBodyComments) SetUserid(v string) *CommentListReportResponseBodyComments {
	s.Userid = &v
	return s
}

func (s *CommentListReportResponseBodyComments) Validate() error {
	return dara.Validate(s)
}
