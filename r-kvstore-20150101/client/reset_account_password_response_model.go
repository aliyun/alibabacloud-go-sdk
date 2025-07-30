// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAccountPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAccountPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse
	GetBody() *ResetAccountPasswordResponseBody
}

type ResetAccountPasswordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAccountPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAccountPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAccountPasswordResponse) GetBody() *ResetAccountPasswordResponseBody {
	return s.Body
}

func (s *ResetAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountPasswordResponse) SetStatusCode(v int32) *ResetAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountPasswordResponse) SetBody(v *ResetAccountPasswordResponseBody) *ResetAccountPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetAccountPasswordResponse) Validate() error {
	return dara.Validate(s)
}
