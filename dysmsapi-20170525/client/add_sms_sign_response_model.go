// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmsSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSmsSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSmsSignResponse
	GetStatusCode() *int32
	SetBody(v *AddSmsSignResponseBody) *AddSmsSignResponse
	GetBody() *AddSmsSignResponseBody
}

type AddSmsSignResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSmsSignResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSmsSignResponse) GoString() string {
	return s.String()
}

func (s *AddSmsSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSmsSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSmsSignResponse) GetBody() *AddSmsSignResponseBody {
	return s.Body
}

func (s *AddSmsSignResponse) SetHeaders(v map[string]*string) *AddSmsSignResponse {
	s.Headers = v
	return s
}

func (s *AddSmsSignResponse) SetStatusCode(v int32) *AddSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSmsSignResponse) SetBody(v *AddSmsSignResponseBody) *AddSmsSignResponse {
	s.Body = v
	return s
}

func (s *AddSmsSignResponse) Validate() error {
	return dara.Validate(s)
}
