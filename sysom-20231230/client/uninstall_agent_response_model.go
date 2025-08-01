// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallAgentResponse
	GetStatusCode() *int32
	SetBody(v *UninstallAgentResponseBody) *UninstallAgentResponse
	GetBody() *UninstallAgentResponseBody
}

type UninstallAgentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentResponse) GoString() string {
	return s.String()
}

func (s *UninstallAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallAgentResponse) GetBody() *UninstallAgentResponseBody {
	return s.Body
}

func (s *UninstallAgentResponse) SetHeaders(v map[string]*string) *UninstallAgentResponse {
	s.Headers = v
	return s
}

func (s *UninstallAgentResponse) SetStatusCode(v int32) *UninstallAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallAgentResponse) SetBody(v *UninstallAgentResponseBody) *UninstallAgentResponse {
	s.Body = v
	return s
}

func (s *UninstallAgentResponse) Validate() error {
	return dara.Validate(s)
}
