// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedFoldersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedFoldersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedFoldersResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedFoldersResponseBody) *ListAuthorizedFoldersResponse
	GetBody() *ListAuthorizedFoldersResponseBody
}

type ListAuthorizedFoldersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedFoldersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedFoldersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedFoldersResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedFoldersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedFoldersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedFoldersResponse) GetBody() *ListAuthorizedFoldersResponseBody {
	return s.Body
}

func (s *ListAuthorizedFoldersResponse) SetHeaders(v map[string]*string) *ListAuthorizedFoldersResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedFoldersResponse) SetStatusCode(v int32) *ListAuthorizedFoldersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedFoldersResponse) SetBody(v *ListAuthorizedFoldersResponseBody) *ListAuthorizedFoldersResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedFoldersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
