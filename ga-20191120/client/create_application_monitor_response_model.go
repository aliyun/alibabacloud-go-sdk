// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationMonitorResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationMonitorResponseBody) *CreateApplicationMonitorResponse
	GetBody() *CreateApplicationMonitorResponseBody
}

type CreateApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationMonitorResponse) GetBody() *CreateApplicationMonitorResponseBody {
	return s.Body
}

func (s *CreateApplicationMonitorResponse) SetHeaders(v map[string]*string) *CreateApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationMonitorResponse) SetStatusCode(v int32) *CreateApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationMonitorResponse) SetBody(v *CreateApplicationMonitorResponseBody) *CreateApplicationMonitorResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
