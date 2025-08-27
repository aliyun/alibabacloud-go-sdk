// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyListQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyListQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyListQueryResponse
	GetStatusCode() *int32
	SetBody(v *ApplyListQueryResponseBody) *ApplyListQueryResponse
	GetBody() *ApplyListQueryResponseBody
}

type ApplyListQueryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyListQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponse) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyListQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyListQueryResponse) GetBody() *ApplyListQueryResponseBody {
	return s.Body
}

func (s *ApplyListQueryResponse) SetHeaders(v map[string]*string) *ApplyListQueryResponse {
	s.Headers = v
	return s
}

func (s *ApplyListQueryResponse) SetStatusCode(v int32) *ApplyListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyListQueryResponse) SetBody(v *ApplyListQueryResponseBody) *ApplyListQueryResponse {
	s.Body = v
	return s
}

func (s *ApplyListQueryResponse) Validate() error {
	return dara.Validate(s)
}
