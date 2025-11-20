// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommentListReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommentListReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommentListReportResponse
	GetStatusCode() *int32
	SetBody(v *CommentListReportResponseBody) *CommentListReportResponse
	GetBody() *CommentListReportResponseBody
}

type CommentListReportResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommentListReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommentListReportResponse) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportResponse) GoString() string {
	return s.String()
}

func (s *CommentListReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommentListReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommentListReportResponse) GetBody() *CommentListReportResponseBody {
	return s.Body
}

func (s *CommentListReportResponse) SetHeaders(v map[string]*string) *CommentListReportResponse {
	s.Headers = v
	return s
}

func (s *CommentListReportResponse) SetStatusCode(v int32) *CommentListReportResponse {
	s.StatusCode = &v
	return s
}

func (s *CommentListReportResponse) SetBody(v *CommentListReportResponseBody) *CommentListReportResponse {
	s.Body = v
	return s
}

func (s *CommentListReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
