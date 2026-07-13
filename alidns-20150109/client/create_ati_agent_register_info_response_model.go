// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAtiAgentRegisterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAtiAgentRegisterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAtiAgentRegisterInfoResponse
	GetStatusCode() *int32
	SetBody(v *CreateAtiAgentRegisterInfoResponseBody) *CreateAtiAgentRegisterInfoResponse
	GetBody() *CreateAtiAgentRegisterInfoResponseBody
}

type CreateAtiAgentRegisterInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAtiAgentRegisterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAtiAgentRegisterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiAgentRegisterInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateAtiAgentRegisterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAtiAgentRegisterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAtiAgentRegisterInfoResponse) GetBody() *CreateAtiAgentRegisterInfoResponseBody {
	return s.Body
}

func (s *CreateAtiAgentRegisterInfoResponse) SetHeaders(v map[string]*string) *CreateAtiAgentRegisterInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateAtiAgentRegisterInfoResponse) SetStatusCode(v int32) *CreateAtiAgentRegisterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAtiAgentRegisterInfoResponse) SetBody(v *CreateAtiAgentRegisterInfoResponseBody) *CreateAtiAgentRegisterInfoResponse {
	s.Body = v
	return s
}

func (s *CreateAtiAgentRegisterInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
