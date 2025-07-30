// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrecisionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrecisionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrecisionTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListPrecisionTaskResponseBody) *ListPrecisionTaskResponse
	GetBody() *ListPrecisionTaskResponseBody
}

type ListPrecisionTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrecisionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrecisionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrecisionTaskResponse) GetBody() *ListPrecisionTaskResponseBody {
	return s.Body
}

func (s *ListPrecisionTaskResponse) SetHeaders(v map[string]*string) *ListPrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *ListPrecisionTaskResponse) SetStatusCode(v int32) *ListPrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrecisionTaskResponse) SetBody(v *ListPrecisionTaskResponseBody) *ListPrecisionTaskResponse {
	s.Body = v
	return s
}

func (s *ListPrecisionTaskResponse) Validate() error {
	return dara.Validate(s)
}
