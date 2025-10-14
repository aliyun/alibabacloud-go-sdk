// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorGroupInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorGroupInstancesResponse
	GetStatusCode() *int32
	SetBody(v *CreateMonitorGroupInstancesResponseBody) *CreateMonitorGroupInstancesResponse
	GetBody() *CreateMonitorGroupInstancesResponseBody
}

type CreateMonitorGroupInstancesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitorGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorGroupInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorGroupInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorGroupInstancesResponse) GetBody() *CreateMonitorGroupInstancesResponseBody {
	return s.Body
}

func (s *CreateMonitorGroupInstancesResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupInstancesResponse) SetStatusCode(v int32) *CreateMonitorGroupInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorGroupInstancesResponse) SetBody(v *CreateMonitorGroupInstancesResponseBody) *CreateMonitorGroupInstancesResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorGroupInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
