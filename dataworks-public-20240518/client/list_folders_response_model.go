// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoldersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFoldersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFoldersResponse
	GetStatusCode() *int32
	SetBody(v *ListFoldersResponseBody) *ListFoldersResponse
	GetBody() *ListFoldersResponseBody
}

type ListFoldersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFoldersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFoldersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersResponse) GoString() string {
	return s.String()
}

func (s *ListFoldersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFoldersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFoldersResponse) GetBody() *ListFoldersResponseBody {
	return s.Body
}

func (s *ListFoldersResponse) SetHeaders(v map[string]*string) *ListFoldersResponse {
	s.Headers = v
	return s
}

func (s *ListFoldersResponse) SetStatusCode(v int32) *ListFoldersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFoldersResponse) SetBody(v *ListFoldersResponseBody) *ListFoldersResponse {
	s.Body = v
	return s
}

func (s *ListFoldersResponse) Validate() error {
	return dara.Validate(s)
}
