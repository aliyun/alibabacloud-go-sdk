// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyQueryResponse
	GetStatusCode() *int32
	SetBody(v *ApplyQueryResponseBody) *ApplyQueryResponse
	GetBody() *ApplyQueryResponseBody
}

type ApplyQueryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyQueryResponse) GetBody() *ApplyQueryResponseBody {
	return s.Body
}

func (s *ApplyQueryResponse) SetHeaders(v map[string]*string) *ApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *ApplyQueryResponse) SetStatusCode(v int32) *ApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyQueryResponse) SetBody(v *ApplyQueryResponseBody) *ApplyQueryResponse {
	s.Body = v
	return s
}

func (s *ApplyQueryResponse) Validate() error {
	return dara.Validate(s)
}
