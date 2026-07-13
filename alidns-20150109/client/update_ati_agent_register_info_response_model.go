// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAtiAgentRegisterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAtiAgentRegisterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAtiAgentRegisterInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAtiAgentRegisterInfoResponseBody) *UpdateAtiAgentRegisterInfoResponse
	GetBody() *UpdateAtiAgentRegisterInfoResponseBody
}

type UpdateAtiAgentRegisterInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAtiAgentRegisterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAtiAgentRegisterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiAgentRegisterInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAtiAgentRegisterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAtiAgentRegisterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAtiAgentRegisterInfoResponse) GetBody() *UpdateAtiAgentRegisterInfoResponseBody {
	return s.Body
}

func (s *UpdateAtiAgentRegisterInfoResponse) SetHeaders(v map[string]*string) *UpdateAtiAgentRegisterInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponse) SetStatusCode(v int32) *UpdateAtiAgentRegisterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponse) SetBody(v *UpdateAtiAgentRegisterInfoResponseBody) *UpdateAtiAgentRegisterInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateAtiAgentRegisterInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
