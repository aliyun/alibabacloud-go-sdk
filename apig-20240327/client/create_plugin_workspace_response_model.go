// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePluginWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePluginWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *CreatePluginWorkspaceResponseBody) *CreatePluginWorkspaceResponse
	GetBody() *CreatePluginWorkspaceResponseBody
}

type CreatePluginWorkspaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePluginWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePluginWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreatePluginWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePluginWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePluginWorkspaceResponse) GetBody() *CreatePluginWorkspaceResponseBody {
	return s.Body
}

func (s *CreatePluginWorkspaceResponse) SetHeaders(v map[string]*string) *CreatePluginWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreatePluginWorkspaceResponse) SetStatusCode(v int32) *CreatePluginWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePluginWorkspaceResponse) SetBody(v *CreatePluginWorkspaceResponseBody) *CreatePluginWorkspaceResponse {
	s.Body = v
	return s
}

func (s *CreatePluginWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
