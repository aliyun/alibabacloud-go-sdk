// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationMonitorResponseBody) *UpdateApplicationMonitorResponse
	GetBody() *UpdateApplicationMonitorResponseBody
}

type UpdateApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationMonitorResponse) GetBody() *UpdateApplicationMonitorResponseBody {
	return s.Body
}

func (s *UpdateApplicationMonitorResponse) SetHeaders(v map[string]*string) *UpdateApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationMonitorResponse) SetStatusCode(v int32) *UpdateApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationMonitorResponse) SetBody(v *UpdateApplicationMonitorResponseBody) *UpdateApplicationMonitorResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
