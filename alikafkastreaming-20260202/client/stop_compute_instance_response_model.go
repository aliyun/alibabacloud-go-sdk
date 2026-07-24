// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopComputeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopComputeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopComputeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopComputeInstanceResponseBody) *StopComputeInstanceResponse
	GetBody() *StopComputeInstanceResponseBody
}

type StopComputeInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopComputeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopComputeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopComputeInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopComputeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopComputeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopComputeInstanceResponse) GetBody() *StopComputeInstanceResponseBody {
	return s.Body
}

func (s *StopComputeInstanceResponse) SetHeaders(v map[string]*string) *StopComputeInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopComputeInstanceResponse) SetStatusCode(v int32) *StopComputeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopComputeInstanceResponse) SetBody(v *StopComputeInstanceResponseBody) *StopComputeInstanceResponse {
	s.Body = v
	return s
}

func (s *StopComputeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
