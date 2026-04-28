// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportUserResponse
	GetStatusCode() *int32
	SetBody(v *User) *ImportUserResponse
	GetBody() *User
}

type ImportUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *User              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportUserResponse) GoString() string {
	return s.String()
}

func (s *ImportUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportUserResponse) GetBody() *User {
	return s.Body
}

func (s *ImportUserResponse) SetHeaders(v map[string]*string) *ImportUserResponse {
	s.Headers = v
	return s
}

func (s *ImportUserResponse) SetStatusCode(v int32) *ImportUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportUserResponse) SetBody(v *User) *ImportUserResponse {
	s.Body = v
	return s
}

func (s *ImportUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
