// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllAdminResponse
	GetStatusCode() *int32
	SetBody(v *ListAllAdminResponseBody) *ListAllAdminResponse
	GetBody() *ListAllAdminResponseBody
}

type ListAllAdminResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllAdminResponse) GoString() string {
	return s.String()
}

func (s *ListAllAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllAdminResponse) GetBody() *ListAllAdminResponseBody {
	return s.Body
}

func (s *ListAllAdminResponse) SetHeaders(v map[string]*string) *ListAllAdminResponse {
	s.Headers = v
	return s
}

func (s *ListAllAdminResponse) SetStatusCode(v int32) *ListAllAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllAdminResponse) SetBody(v *ListAllAdminResponseBody) *ListAllAdminResponse {
	s.Body = v
	return s
}

func (s *ListAllAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
