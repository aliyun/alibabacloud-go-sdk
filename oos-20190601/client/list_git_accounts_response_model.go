// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGitAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGitAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListGitAccountsResponseBody) *ListGitAccountsResponse
	GetBody() *ListGitAccountsResponseBody
}

type ListGitAccountsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGitAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGitAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGitAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListGitAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGitAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGitAccountsResponse) GetBody() *ListGitAccountsResponseBody {
	return s.Body
}

func (s *ListGitAccountsResponse) SetHeaders(v map[string]*string) *ListGitAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListGitAccountsResponse) SetStatusCode(v int32) *ListGitAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGitAccountsResponse) SetBody(v *ListGitAccountsResponseBody) *ListGitAccountsResponse {
	s.Body = v
	return s
}

func (s *ListGitAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
