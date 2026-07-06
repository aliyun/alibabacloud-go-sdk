// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAgentWithTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallAgentWithTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallAgentWithTypeResponse
	GetStatusCode() *int32
	SetBody(v *UninstallAgentWithTypeResponseBody) *UninstallAgentWithTypeResponse
	GetBody() *UninstallAgentWithTypeResponseBody
}

type UninstallAgentWithTypeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallAgentWithTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallAgentWithTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallAgentWithTypeResponse) GoString() string {
	return s.String()
}

func (s *UninstallAgentWithTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallAgentWithTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallAgentWithTypeResponse) GetBody() *UninstallAgentWithTypeResponseBody {
	return s.Body
}

func (s *UninstallAgentWithTypeResponse) SetHeaders(v map[string]*string) *UninstallAgentWithTypeResponse {
	s.Headers = v
	return s
}

func (s *UninstallAgentWithTypeResponse) SetStatusCode(v int32) *UninstallAgentWithTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallAgentWithTypeResponse) SetBody(v *UninstallAgentWithTypeResponseBody) *UninstallAgentWithTypeResponse {
	s.Body = v
	return s
}

func (s *UninstallAgentWithTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
