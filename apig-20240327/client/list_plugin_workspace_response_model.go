// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPluginWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPluginWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *ListPluginWorkspaceResponseBody) *ListPluginWorkspaceResponse
	GetBody() *ListPluginWorkspaceResponseBody
}

type ListPluginWorkspaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPluginWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPluginWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPluginWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *ListPluginWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPluginWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPluginWorkspaceResponse) GetBody() *ListPluginWorkspaceResponseBody {
	return s.Body
}

func (s *ListPluginWorkspaceResponse) SetHeaders(v map[string]*string) *ListPluginWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *ListPluginWorkspaceResponse) SetStatusCode(v int32) *ListPluginWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPluginWorkspaceResponse) SetBody(v *ListPluginWorkspaceResponseBody) *ListPluginWorkspaceResponse {
	s.Body = v
	return s
}

func (s *ListPluginWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
