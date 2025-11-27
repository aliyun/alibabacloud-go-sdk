// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopRCInstanceResponseBody) *StopRCInstanceResponse
	GetBody() *StopRCInstanceResponseBody
}

type StopRCInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRCInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopRCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRCInstanceResponse) GetBody() *StopRCInstanceResponseBody {
	return s.Body
}

func (s *StopRCInstanceResponse) SetHeaders(v map[string]*string) *StopRCInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopRCInstanceResponse) SetStatusCode(v int32) *StopRCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRCInstanceResponse) SetBody(v *StopRCInstanceResponseBody) *StopRCInstanceResponse {
	s.Body = v
	return s
}

func (s *StopRCInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
