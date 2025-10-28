// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkspaceResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkspaceResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkspaceResourceResponseBody) *DeleteWorkspaceResourceResponse
	GetBody() *DeleteWorkspaceResourceResponseBody
}

type DeleteWorkspaceResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkspaceResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkspaceResourceResponse) GetBody() *DeleteWorkspaceResourceResponseBody {
	return s.Body
}

func (s *DeleteWorkspaceResourceResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceResourceResponse) SetStatusCode(v int32) *DeleteWorkspaceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceResourceResponse) SetBody(v *DeleteWorkspaceResourceResponseBody) *DeleteWorkspaceResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkspaceResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
