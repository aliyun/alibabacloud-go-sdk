// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateMonitorGroupResponseBody) *CreateMonitorGroupResponse
	GetBody() *CreateMonitorGroupResponseBody
}

type CreateMonitorGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorGroupResponse) GetBody() *CreateMonitorGroupResponseBody {
	return s.Body
}

func (s *CreateMonitorGroupResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupResponse) SetStatusCode(v int32) *CreateMonitorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorGroupResponse) SetBody(v *CreateMonitorGroupResponseBody) *CreateMonitorGroupResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
