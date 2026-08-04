// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitWorkspaceSystemMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitWorkspaceSystemMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitWorkspaceSystemMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *InitWorkspaceSystemMcpServerResponseBody) *InitWorkspaceSystemMcpServerResponse
	GetBody() *InitWorkspaceSystemMcpServerResponseBody
}

type InitWorkspaceSystemMcpServerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitWorkspaceSystemMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitWorkspaceSystemMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s InitWorkspaceSystemMcpServerResponse) GoString() string {
	return s.String()
}

func (s *InitWorkspaceSystemMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitWorkspaceSystemMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitWorkspaceSystemMcpServerResponse) GetBody() *InitWorkspaceSystemMcpServerResponseBody {
	return s.Body
}

func (s *InitWorkspaceSystemMcpServerResponse) SetHeaders(v map[string]*string) *InitWorkspaceSystemMcpServerResponse {
	s.Headers = v
	return s
}

func (s *InitWorkspaceSystemMcpServerResponse) SetStatusCode(v int32) *InitWorkspaceSystemMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *InitWorkspaceSystemMcpServerResponse) SetBody(v *InitWorkspaceSystemMcpServerResponseBody) *InitWorkspaceSystemMcpServerResponse {
	s.Body = v
	return s
}

func (s *InitWorkspaceSystemMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
