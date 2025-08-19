// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitCardVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitCardVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitCardVerifyResponse
	GetStatusCode() *int32
	SetBody(v *InitCardVerifyResponseBody) *InitCardVerifyResponse
	GetBody() *InitCardVerifyResponseBody
}

type InitCardVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitCardVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitCardVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s InitCardVerifyResponse) GoString() string {
	return s.String()
}

func (s *InitCardVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitCardVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitCardVerifyResponse) GetBody() *InitCardVerifyResponseBody {
	return s.Body
}

func (s *InitCardVerifyResponse) SetHeaders(v map[string]*string) *InitCardVerifyResponse {
	s.Headers = v
	return s
}

func (s *InitCardVerifyResponse) SetStatusCode(v int32) *InitCardVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *InitCardVerifyResponse) SetBody(v *InitCardVerifyResponseBody) *InitCardVerifyResponse {
	s.Body = v
	return s
}

func (s *InitCardVerifyResponse) Validate() error {
	return dara.Validate(s)
}
