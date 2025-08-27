// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTripTaskExecuteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyTripTaskExecuteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyTripTaskExecuteResponse
	GetStatusCode() *int32
	SetBody(v *ApplyTripTaskExecuteResponseBody) *ApplyTripTaskExecuteResponse
	GetBody() *ApplyTripTaskExecuteResponseBody
}

type ApplyTripTaskExecuteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyTripTaskExecuteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyTripTaskExecuteResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyTripTaskExecuteResponse) GoString() string {
	return s.String()
}

func (s *ApplyTripTaskExecuteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyTripTaskExecuteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyTripTaskExecuteResponse) GetBody() *ApplyTripTaskExecuteResponseBody {
	return s.Body
}

func (s *ApplyTripTaskExecuteResponse) SetHeaders(v map[string]*string) *ApplyTripTaskExecuteResponse {
	s.Headers = v
	return s
}

func (s *ApplyTripTaskExecuteResponse) SetStatusCode(v int32) *ApplyTripTaskExecuteResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyTripTaskExecuteResponse) SetBody(v *ApplyTripTaskExecuteResponseBody) *ApplyTripTaskExecuteResponse {
	s.Body = v
	return s
}

func (s *ApplyTripTaskExecuteResponse) Validate() error {
	return dara.Validate(s)
}
