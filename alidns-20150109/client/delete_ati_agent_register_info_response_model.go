// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAtiAgentRegisterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAtiAgentRegisterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAtiAgentRegisterInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAtiAgentRegisterInfoResponseBody) *DeleteAtiAgentRegisterInfoResponse
	GetBody() *DeleteAtiAgentRegisterInfoResponseBody
}

type DeleteAtiAgentRegisterInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAtiAgentRegisterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAtiAgentRegisterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAtiAgentRegisterInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteAtiAgentRegisterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAtiAgentRegisterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAtiAgentRegisterInfoResponse) GetBody() *DeleteAtiAgentRegisterInfoResponseBody {
	return s.Body
}

func (s *DeleteAtiAgentRegisterInfoResponse) SetHeaders(v map[string]*string) *DeleteAtiAgentRegisterInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponse) SetStatusCode(v int32) *DeleteAtiAgentRegisterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponse) SetBody(v *DeleteAtiAgentRegisterInfoResponseBody) *DeleteAtiAgentRegisterInfoResponse {
	s.Body = v
	return s
}

func (s *DeleteAtiAgentRegisterInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
