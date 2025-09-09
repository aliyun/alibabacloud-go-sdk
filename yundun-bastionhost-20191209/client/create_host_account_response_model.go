// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHostAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHostAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateHostAccountResponseBody) *CreateHostAccountResponse
	GetBody() *CreateHostAccountResponseBody
}

type CreateHostAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHostAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateHostAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHostAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHostAccountResponse) GetBody() *CreateHostAccountResponseBody {
	return s.Body
}

func (s *CreateHostAccountResponse) SetHeaders(v map[string]*string) *CreateHostAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateHostAccountResponse) SetStatusCode(v int32) *CreateHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostAccountResponse) SetBody(v *CreateHostAccountResponseBody) *CreateHostAccountResponse {
	s.Body = v
	return s
}

func (s *CreateHostAccountResponse) Validate() error {
	return dara.Validate(s)
}
