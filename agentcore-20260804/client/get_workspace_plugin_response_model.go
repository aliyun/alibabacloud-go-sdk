// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspacePluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspacePluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspacePluginResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkspacePluginResponseBody) *GetWorkspacePluginResponse
	GetBody() *GetWorkspacePluginResponseBody
}

type GetWorkspacePluginResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspacePluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspacePluginResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacePluginResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspacePluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspacePluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspacePluginResponse) GetBody() *GetWorkspacePluginResponseBody {
	return s.Body
}

func (s *GetWorkspacePluginResponse) SetHeaders(v map[string]*string) *GetWorkspacePluginResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspacePluginResponse) SetStatusCode(v int32) *GetWorkspacePluginResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspacePluginResponse) SetBody(v *GetWorkspacePluginResponseBody) *GetWorkspacePluginResponse {
	s.Body = v
	return s
}

func (s *GetWorkspacePluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
