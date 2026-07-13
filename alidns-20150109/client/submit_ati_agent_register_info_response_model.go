// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAtiAgentRegisterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAtiAgentRegisterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAtiAgentRegisterInfoResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAtiAgentRegisterInfoResponseBody) *SubmitAtiAgentRegisterInfoResponse
	GetBody() *SubmitAtiAgentRegisterInfoResponseBody
}

type SubmitAtiAgentRegisterInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAtiAgentRegisterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAtiAgentRegisterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAtiAgentRegisterInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitAtiAgentRegisterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAtiAgentRegisterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAtiAgentRegisterInfoResponse) GetBody() *SubmitAtiAgentRegisterInfoResponseBody {
	return s.Body
}

func (s *SubmitAtiAgentRegisterInfoResponse) SetHeaders(v map[string]*string) *SubmitAtiAgentRegisterInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponse) SetStatusCode(v int32) *SubmitAtiAgentRegisterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponse) SetBody(v *SubmitAtiAgentRegisterInfoResponseBody) *SubmitAtiAgentRegisterInfoResponse {
	s.Body = v
	return s
}

func (s *SubmitAtiAgentRegisterInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
