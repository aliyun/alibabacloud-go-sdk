// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTouchListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTouchListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTouchListResponse
	GetStatusCode() *int32
	SetBody(v *QueryTouchListResponseBody) *QueryTouchListResponse
	GetBody() *QueryTouchListResponseBody
}

type QueryTouchListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTouchListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTouchListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTouchListResponse) GoString() string {
	return s.String()
}

func (s *QueryTouchListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTouchListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTouchListResponse) GetBody() *QueryTouchListResponseBody {
	return s.Body
}

func (s *QueryTouchListResponse) SetHeaders(v map[string]*string) *QueryTouchListResponse {
	s.Headers = v
	return s
}

func (s *QueryTouchListResponse) SetStatusCode(v int32) *QueryTouchListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTouchListResponse) SetBody(v *QueryTouchListResponseBody) *QueryTouchListResponse {
	s.Body = v
	return s
}

func (s *QueryTouchListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
