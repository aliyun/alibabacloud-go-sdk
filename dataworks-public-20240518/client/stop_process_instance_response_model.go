// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopProcessInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopProcessInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopProcessInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopProcessInstanceResponseBody) *StopProcessInstanceResponse
	GetBody() *StopProcessInstanceResponseBody
}

type StopProcessInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopProcessInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopProcessInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopProcessInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopProcessInstanceResponse) GetBody() *StopProcessInstanceResponseBody {
	return s.Body
}

func (s *StopProcessInstanceResponse) SetHeaders(v map[string]*string) *StopProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopProcessInstanceResponse) SetStatusCode(v int32) *StopProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopProcessInstanceResponse) SetBody(v *StopProcessInstanceResponseBody) *StopProcessInstanceResponse {
	s.Body = v
	return s
}

func (s *StopProcessInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
