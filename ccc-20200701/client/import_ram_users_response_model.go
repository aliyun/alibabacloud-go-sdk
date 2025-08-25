// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRamUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportRamUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportRamUsersResponse
	GetStatusCode() *int32
	SetBody(v *ImportRamUsersResponseBody) *ImportRamUsersResponse
	GetBody() *ImportRamUsersResponseBody
}

type ImportRamUsersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportRamUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportRamUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportRamUsersResponse) GoString() string {
	return s.String()
}

func (s *ImportRamUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportRamUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportRamUsersResponse) GetBody() *ImportRamUsersResponseBody {
	return s.Body
}

func (s *ImportRamUsersResponse) SetHeaders(v map[string]*string) *ImportRamUsersResponse {
	s.Headers = v
	return s
}

func (s *ImportRamUsersResponse) SetStatusCode(v int32) *ImportRamUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportRamUsersResponse) SetBody(v *ImportRamUsersResponseBody) *ImportRamUsersResponse {
	s.Body = v
	return s
}

func (s *ImportRamUsersResponse) Validate() error {
	return dara.Validate(s)
}
