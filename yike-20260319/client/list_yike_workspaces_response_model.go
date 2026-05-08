// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeWorkspacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListYikeWorkspacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListYikeWorkspacesResponse
	GetStatusCode() *int32
	SetBody(v *ListYikeWorkspacesResponseBody) *ListYikeWorkspacesResponse
	GetBody() *ListYikeWorkspacesResponseBody
}

type ListYikeWorkspacesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListYikeWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListYikeWorkspacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListYikeWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListYikeWorkspacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListYikeWorkspacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListYikeWorkspacesResponse) GetBody() *ListYikeWorkspacesResponseBody {
	return s.Body
}

func (s *ListYikeWorkspacesResponse) SetHeaders(v map[string]*string) *ListYikeWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListYikeWorkspacesResponse) SetStatusCode(v int32) *ListYikeWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListYikeWorkspacesResponse) SetBody(v *ListYikeWorkspacesResponseBody) *ListYikeWorkspacesResponse {
	s.Body = v
	return s
}

func (s *ListYikeWorkspacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
