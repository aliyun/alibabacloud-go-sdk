// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkspaceResponseBody) *DeleteWorkspaceResponse
	GetBody() *DeleteWorkspaceResponseBody
}

type DeleteWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkspaceResponse) GetBody() *DeleteWorkspaceResponseBody {
	return s.Body
}

func (s *DeleteWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceResponse) SetStatusCode(v int32) *DeleteWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceResponse) SetBody(v *DeleteWorkspaceResponseBody) *DeleteWorkspaceResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
