// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMaskingProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopMaskingProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopMaskingProcessResponse
	GetStatusCode() *int32
	SetBody(v *StopMaskingProcessResponseBody) *StopMaskingProcessResponse
	GetBody() *StopMaskingProcessResponseBody
}

type StopMaskingProcessResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMaskingProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMaskingProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s StopMaskingProcessResponse) GoString() string {
	return s.String()
}

func (s *StopMaskingProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopMaskingProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopMaskingProcessResponse) GetBody() *StopMaskingProcessResponseBody {
	return s.Body
}

func (s *StopMaskingProcessResponse) SetHeaders(v map[string]*string) *StopMaskingProcessResponse {
	s.Headers = v
	return s
}

func (s *StopMaskingProcessResponse) SetStatusCode(v int32) *StopMaskingProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMaskingProcessResponse) SetBody(v *StopMaskingProcessResponseBody) *StopMaskingProcessResponse {
	s.Body = v
	return s
}

func (s *StopMaskingProcessResponse) Validate() error {
	return dara.Validate(s)
}
