// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCollectorResponse
	GetStatusCode() *int32
	SetBody(v *StopCollectorResponseBody) *StopCollectorResponse
	GetBody() *StopCollectorResponseBody
}

type StopCollectorResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCollectorResponse) GoString() string {
	return s.String()
}

func (s *StopCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCollectorResponse) GetBody() *StopCollectorResponseBody {
	return s.Body
}

func (s *StopCollectorResponse) SetHeaders(v map[string]*string) *StopCollectorResponse {
	s.Headers = v
	return s
}

func (s *StopCollectorResponse) SetStatusCode(v int32) *StopCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCollectorResponse) SetBody(v *StopCollectorResponseBody) *StopCollectorResponse {
	s.Body = v
	return s
}

func (s *StopCollectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
