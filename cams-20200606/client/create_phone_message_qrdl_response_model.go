// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhoneMessageQrdlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePhoneMessageQrdlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePhoneMessageQrdlResponse
	GetStatusCode() *int32
	SetBody(v *CreatePhoneMessageQrdlResponseBody) *CreatePhoneMessageQrdlResponse
	GetBody() *CreatePhoneMessageQrdlResponseBody
}

type CreatePhoneMessageQrdlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePhoneMessageQrdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePhoneMessageQrdlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePhoneMessageQrdlResponse) GoString() string {
	return s.String()
}

func (s *CreatePhoneMessageQrdlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePhoneMessageQrdlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePhoneMessageQrdlResponse) GetBody() *CreatePhoneMessageQrdlResponseBody {
	return s.Body
}

func (s *CreatePhoneMessageQrdlResponse) SetHeaders(v map[string]*string) *CreatePhoneMessageQrdlResponse {
	s.Headers = v
	return s
}

func (s *CreatePhoneMessageQrdlResponse) SetStatusCode(v int32) *CreatePhoneMessageQrdlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponse) SetBody(v *CreatePhoneMessageQrdlResponseBody) *CreatePhoneMessageQrdlResponse {
	s.Body = v
	return s
}

func (s *CreatePhoneMessageQrdlResponse) Validate() error {
	return dara.Validate(s)
}
