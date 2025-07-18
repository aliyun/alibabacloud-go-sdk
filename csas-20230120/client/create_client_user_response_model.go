// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClientUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClientUserResponse
	GetStatusCode() *int32
	SetBody(v *CreateClientUserResponseBody) *CreateClientUserResponse
	GetBody() *CreateClientUserResponseBody
}

type CreateClientUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClientUserResponse) GoString() string {
	return s.String()
}

func (s *CreateClientUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClientUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClientUserResponse) GetBody() *CreateClientUserResponseBody {
	return s.Body
}

func (s *CreateClientUserResponse) SetHeaders(v map[string]*string) *CreateClientUserResponse {
	s.Headers = v
	return s
}

func (s *CreateClientUserResponse) SetStatusCode(v int32) *CreateClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientUserResponse) SetBody(v *CreateClientUserResponseBody) *CreateClientUserResponse {
	s.Body = v
	return s
}

func (s *CreateClientUserResponse) Validate() error {
	return dara.Validate(s)
}
