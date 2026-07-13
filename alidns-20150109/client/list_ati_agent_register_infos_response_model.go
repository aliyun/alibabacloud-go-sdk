// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiAgentRegisterInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAtiAgentRegisterInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAtiAgentRegisterInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListAtiAgentRegisterInfosResponseBody) *ListAtiAgentRegisterInfosResponse
	GetBody() *ListAtiAgentRegisterInfosResponseBody
}

type ListAtiAgentRegisterInfosResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAtiAgentRegisterInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAtiAgentRegisterInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAtiAgentRegisterInfosResponse) GoString() string {
	return s.String()
}

func (s *ListAtiAgentRegisterInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAtiAgentRegisterInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAtiAgentRegisterInfosResponse) GetBody() *ListAtiAgentRegisterInfosResponseBody {
	return s.Body
}

func (s *ListAtiAgentRegisterInfosResponse) SetHeaders(v map[string]*string) *ListAtiAgentRegisterInfosResponse {
	s.Headers = v
	return s
}

func (s *ListAtiAgentRegisterInfosResponse) SetStatusCode(v int32) *ListAtiAgentRegisterInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponse) SetBody(v *ListAtiAgentRegisterInfosResponseBody) *ListAtiAgentRegisterInfosResponse {
	s.Body = v
	return s
}

func (s *ListAtiAgentRegisterInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
