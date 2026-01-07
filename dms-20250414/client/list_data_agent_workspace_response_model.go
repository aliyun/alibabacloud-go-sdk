// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentWorkspaceResponseBody) *ListDataAgentWorkspaceResponse
	GetBody() *ListDataAgentWorkspaceResponseBody
}

type ListDataAgentWorkspaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentWorkspaceResponse) GetBody() *ListDataAgentWorkspaceResponseBody {
	return s.Body
}

func (s *ListDataAgentWorkspaceResponse) SetHeaders(v map[string]*string) *ListDataAgentWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentWorkspaceResponse) SetStatusCode(v int32) *ListDataAgentWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentWorkspaceResponse) SetBody(v *ListDataAgentWorkspaceResponseBody) *ListDataAgentWorkspaceResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
