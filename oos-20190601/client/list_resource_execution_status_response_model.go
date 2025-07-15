// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExecutionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceExecutionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceExecutionStatusResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceExecutionStatusResponseBody) *ListResourceExecutionStatusResponse
	GetBody() *ListResourceExecutionStatusResponseBody
}

type ListResourceExecutionStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceExecutionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceExecutionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExecutionStatusResponse) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceExecutionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceExecutionStatusResponse) GetBody() *ListResourceExecutionStatusResponseBody {
	return s.Body
}

func (s *ListResourceExecutionStatusResponse) SetHeaders(v map[string]*string) *ListResourceExecutionStatusResponse {
	s.Headers = v
	return s
}

func (s *ListResourceExecutionStatusResponse) SetStatusCode(v int32) *ListResourceExecutionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceExecutionStatusResponse) SetBody(v *ListResourceExecutionStatusResponseBody) *ListResourceExecutionStatusResponse {
	s.Body = v
	return s
}

func (s *ListResourceExecutionStatusResponse) Validate() error {
	return dara.Validate(s)
}
