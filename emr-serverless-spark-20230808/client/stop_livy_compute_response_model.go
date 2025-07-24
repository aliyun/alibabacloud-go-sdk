// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLivyComputeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLivyComputeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLivyComputeResponse
	GetStatusCode() *int32
	SetBody(v *StopLivyComputeResponseBody) *StopLivyComputeResponse
	GetBody() *StopLivyComputeResponseBody
}

type StopLivyComputeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLivyComputeResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *StopLivyComputeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLivyComputeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLivyComputeResponse) GetBody() *StopLivyComputeResponseBody {
	return s.Body
}

func (s *StopLivyComputeResponse) SetHeaders(v map[string]*string) *StopLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *StopLivyComputeResponse) SetStatusCode(v int32) *StopLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLivyComputeResponse) SetBody(v *StopLivyComputeResponseBody) *StopLivyComputeResponse {
	s.Body = v
	return s
}

func (s *StopLivyComputeResponse) Validate() error {
	return dara.Validate(s)
}
