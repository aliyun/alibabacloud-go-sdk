// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupsMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserGroupsMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserGroupsMappingResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserGroupsMappingResponseBody) *CreateUserGroupsMappingResponse
	GetBody() *CreateUserGroupsMappingResponseBody
}

type CreateUserGroupsMappingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserGroupsMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserGroupsMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupsMappingResponse) GoString() string {
	return s.String()
}

func (s *CreateUserGroupsMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserGroupsMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserGroupsMappingResponse) GetBody() *CreateUserGroupsMappingResponseBody {
	return s.Body
}

func (s *CreateUserGroupsMappingResponse) SetHeaders(v map[string]*string) *CreateUserGroupsMappingResponse {
	s.Headers = v
	return s
}

func (s *CreateUserGroupsMappingResponse) SetStatusCode(v int32) *CreateUserGroupsMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserGroupsMappingResponse) SetBody(v *CreateUserGroupsMappingResponseBody) *CreateUserGroupsMappingResponse {
	s.Body = v
	return s
}

func (s *CreateUserGroupsMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
