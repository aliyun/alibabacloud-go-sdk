// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAtiAgentRegisterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeAtiAgentRegisterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeAtiAgentRegisterInfoResponse
	GetStatusCode() *int32
	SetBody(v *RevokeAtiAgentRegisterInfoResponseBody) *RevokeAtiAgentRegisterInfoResponse
	GetBody() *RevokeAtiAgentRegisterInfoResponseBody
}

type RevokeAtiAgentRegisterInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeAtiAgentRegisterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeAtiAgentRegisterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeAtiAgentRegisterInfoResponse) GoString() string {
	return s.String()
}

func (s *RevokeAtiAgentRegisterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeAtiAgentRegisterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeAtiAgentRegisterInfoResponse) GetBody() *RevokeAtiAgentRegisterInfoResponseBody {
	return s.Body
}

func (s *RevokeAtiAgentRegisterInfoResponse) SetHeaders(v map[string]*string) *RevokeAtiAgentRegisterInfoResponse {
	s.Headers = v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponse) SetStatusCode(v int32) *RevokeAtiAgentRegisterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponse) SetBody(v *RevokeAtiAgentRegisterInfoResponseBody) *RevokeAtiAgentRegisterInfoResponse {
	s.Body = v
	return s
}

func (s *RevokeAtiAgentRegisterInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
