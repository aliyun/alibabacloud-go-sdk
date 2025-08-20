// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLoginPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetLoginPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetLoginPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetLoginPasswordResponseBody) *ResetLoginPasswordResponse
	GetBody() *ResetLoginPasswordResponseBody
}

type ResetLoginPasswordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetLoginPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetLoginPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetLoginPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetLoginPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetLoginPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetLoginPasswordResponse) GetBody() *ResetLoginPasswordResponseBody {
	return s.Body
}

func (s *ResetLoginPasswordResponse) SetHeaders(v map[string]*string) *ResetLoginPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetLoginPasswordResponse) SetStatusCode(v int32) *ResetLoginPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetLoginPasswordResponse) SetBody(v *ResetLoginPasswordResponseBody) *ResetLoginPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetLoginPasswordResponse) Validate() error {
	return dara.Validate(s)
}
