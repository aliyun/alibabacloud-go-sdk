// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsInResourceDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccountsInResourceDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccountsInResourceDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *ListAccountsInResourceDirectoryResponseBody) *ListAccountsInResourceDirectoryResponse
	GetBody() *ListAccountsInResourceDirectoryResponseBody
}

type ListAccountsInResourceDirectoryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountsInResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountsInResourceDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsInResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ListAccountsInResourceDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccountsInResourceDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccountsInResourceDirectoryResponse) GetBody() *ListAccountsInResourceDirectoryResponseBody {
	return s.Body
}

func (s *ListAccountsInResourceDirectoryResponse) SetHeaders(v map[string]*string) *ListAccountsInResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ListAccountsInResourceDirectoryResponse) SetStatusCode(v int32) *ListAccountsInResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponse) SetBody(v *ListAccountsInResourceDirectoryResponseBody) *ListAccountsInResourceDirectoryResponse {
	s.Body = v
	return s
}

func (s *ListAccountsInResourceDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
