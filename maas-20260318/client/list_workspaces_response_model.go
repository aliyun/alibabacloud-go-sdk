// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkspacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkspacesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse
	GetBody() *ListWorkspacesResponseBody
}

type ListWorkspacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkspacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkspacesResponse) GetBody() *ListWorkspacesResponseBody {
	return s.Body
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetStatusCode(v int32) *ListWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

func (s *ListWorkspacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
