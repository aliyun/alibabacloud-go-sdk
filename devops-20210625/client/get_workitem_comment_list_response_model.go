// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkitemCommentListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkitemCommentListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkitemCommentListResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkitemCommentListResponseBody) *GetWorkitemCommentListResponse
	GetBody() *GetWorkitemCommentListResponseBody
}

type GetWorkitemCommentListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkitemCommentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkitemCommentListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkitemCommentListResponse) GoString() string {
	return s.String()
}

func (s *GetWorkitemCommentListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkitemCommentListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkitemCommentListResponse) GetBody() *GetWorkitemCommentListResponseBody {
	return s.Body
}

func (s *GetWorkitemCommentListResponse) SetHeaders(v map[string]*string) *GetWorkitemCommentListResponse {
	s.Headers = v
	return s
}

func (s *GetWorkitemCommentListResponse) SetStatusCode(v int32) *GetWorkitemCommentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkitemCommentListResponse) SetBody(v *GetWorkitemCommentListResponseBody) *GetWorkitemCommentListResponse {
	s.Body = v
	return s
}

func (s *GetWorkitemCommentListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
