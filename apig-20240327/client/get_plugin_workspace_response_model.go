// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPluginWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPluginWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *GetPluginWorkspaceResponseBody) *GetPluginWorkspaceResponse
	GetBody() *GetPluginWorkspaceResponseBody
}

type GetPluginWorkspaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPluginWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPluginWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPluginWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetPluginWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPluginWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPluginWorkspaceResponse) GetBody() *GetPluginWorkspaceResponseBody {
	return s.Body
}

func (s *GetPluginWorkspaceResponse) SetHeaders(v map[string]*string) *GetPluginWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetPluginWorkspaceResponse) SetStatusCode(v int32) *GetPluginWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPluginWorkspaceResponse) SetBody(v *GetPluginWorkspaceResponseBody) *GetPluginWorkspaceResponse {
	s.Body = v
	return s
}

func (s *GetPluginWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
