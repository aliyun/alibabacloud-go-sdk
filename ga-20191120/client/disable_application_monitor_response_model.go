// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationMonitorResponseBody) *DisableApplicationMonitorResponse
	GetBody() *DisableApplicationMonitorResponseBody
}

type DisableApplicationMonitorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationMonitorResponse) GetBody() *DisableApplicationMonitorResponseBody {
	return s.Body
}

func (s *DisableApplicationMonitorResponse) SetHeaders(v map[string]*string) *DisableApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationMonitorResponse) SetStatusCode(v int32) *DisableApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationMonitorResponse) SetBody(v *DisableApplicationMonitorResponseBody) *DisableApplicationMonitorResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
