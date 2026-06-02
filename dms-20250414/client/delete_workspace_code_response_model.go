// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkspaceCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkspaceCodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkspaceCodeResponseBody) *DeleteWorkspaceCodeResponse
	GetBody() *DeleteWorkspaceCodeResponseBody
}

type DeleteWorkspaceCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceCodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkspaceCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkspaceCodeResponse) GetBody() *DeleteWorkspaceCodeResponseBody {
	return s.Body
}

func (s *DeleteWorkspaceCodeResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceCodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceCodeResponse) SetStatusCode(v int32) *DeleteWorkspaceCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceCodeResponse) SetBody(v *DeleteWorkspaceCodeResponseBody) *DeleteWorkspaceCodeResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkspaceCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
