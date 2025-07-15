// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetUserPasswordResponseBody) *ResetUserPasswordResponse
	GetBody() *ResetUserPasswordResponseBody
}

type ResetUserPasswordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetUserPasswordResponse) GetBody() *ResetUserPasswordResponseBody {
	return s.Body
}

func (s *ResetUserPasswordResponse) SetHeaders(v map[string]*string) *ResetUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetUserPasswordResponse) SetStatusCode(v int32) *ResetUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetUserPasswordResponse) SetBody(v *ResetUserPasswordResponseBody) *ResetUserPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetUserPasswordResponse) Validate() error {
	return dara.Validate(s)
}
