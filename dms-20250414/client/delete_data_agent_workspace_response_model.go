// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataAgentWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataAgentWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataAgentWorkspaceResponseBody) *DeleteDataAgentWorkspaceResponse
	GetBody() *DeleteDataAgentWorkspaceResponseBody
}

type DeleteDataAgentWorkspaceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataAgentWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataAgentWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataAgentWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataAgentWorkspaceResponse) GetBody() *DeleteDataAgentWorkspaceResponseBody {
	return s.Body
}

func (s *DeleteDataAgentWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteDataAgentWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataAgentWorkspaceResponse) SetStatusCode(v int32) *DeleteDataAgentWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataAgentWorkspaceResponse) SetBody(v *DeleteDataAgentWorkspaceResponseBody) *DeleteDataAgentWorkspaceResponse {
	s.Body = v
	return s
}

func (s *DeleteDataAgentWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
