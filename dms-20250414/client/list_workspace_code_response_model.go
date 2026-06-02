// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkspaceCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkspaceCodeResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkspaceCodeResponseBody) *ListWorkspaceCodeResponse
	GetBody() *ListWorkspaceCodeResponseBody
}

type ListWorkspaceCodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceCodeResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkspaceCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkspaceCodeResponse) GetBody() *ListWorkspaceCodeResponseBody {
	return s.Body
}

func (s *ListWorkspaceCodeResponse) SetHeaders(v map[string]*string) *ListWorkspaceCodeResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceCodeResponse) SetStatusCode(v int32) *ListWorkspaceCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceCodeResponse) SetBody(v *ListWorkspaceCodeResponseBody) *ListWorkspaceCodeResponse {
	s.Body = v
	return s
}

func (s *ListWorkspaceCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
