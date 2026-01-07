// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataAgentWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataAgentWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataAgentWorkspaceResponseBody) *CreateDataAgentWorkspaceResponse
	GetBody() *CreateDataAgentWorkspaceResponseBody
}

type CreateDataAgentWorkspaceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataAgentWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataAgentWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataAgentWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataAgentWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataAgentWorkspaceResponse) GetBody() *CreateDataAgentWorkspaceResponseBody {
	return s.Body
}

func (s *CreateDataAgentWorkspaceResponse) SetHeaders(v map[string]*string) *CreateDataAgentWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataAgentWorkspaceResponse) SetStatusCode(v int32) *CreateDataAgentWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataAgentWorkspaceResponse) SetBody(v *CreateDataAgentWorkspaceResponseBody) *CreateDataAgentWorkspaceResponse {
	s.Body = v
	return s
}

func (s *CreateDataAgentWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
