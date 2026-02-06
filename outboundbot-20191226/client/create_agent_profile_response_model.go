// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentProfileResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentProfileResponseBody) *CreateAgentProfileResponse
	GetBody() *CreateAgentProfileResponseBody
}

type CreateAgentProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentProfileResponse) GetBody() *CreateAgentProfileResponseBody {
	return s.Body
}

func (s *CreateAgentProfileResponse) SetHeaders(v map[string]*string) *CreateAgentProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentProfileResponse) SetStatusCode(v int32) *CreateAgentProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentProfileResponse) SetBody(v *CreateAgentProfileResponseBody) *CreateAgentProfileResponse {
	s.Body = v
	return s
}

func (s *CreateAgentProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
