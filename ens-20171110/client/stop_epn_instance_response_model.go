// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopEpnInstanceResponseBody) *StopEpnInstanceResponse
	GetBody() *StopEpnInstanceResponseBody
}

type StopEpnInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopEpnInstanceResponse) GetBody() *StopEpnInstanceResponseBody {
	return s.Body
}

func (s *StopEpnInstanceResponse) SetHeaders(v map[string]*string) *StopEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopEpnInstanceResponse) SetStatusCode(v int32) *StopEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopEpnInstanceResponse) SetBody(v *StopEpnInstanceResponseBody) *StopEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *StopEpnInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
