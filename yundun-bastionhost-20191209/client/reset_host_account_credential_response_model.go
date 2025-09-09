// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetHostAccountCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetHostAccountCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetHostAccountCredentialResponse
	GetStatusCode() *int32
	SetBody(v *ResetHostAccountCredentialResponseBody) *ResetHostAccountCredentialResponse
	GetBody() *ResetHostAccountCredentialResponseBody
}

type ResetHostAccountCredentialResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetHostAccountCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetHostAccountCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetHostAccountCredentialResponse) GoString() string {
	return s.String()
}

func (s *ResetHostAccountCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetHostAccountCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetHostAccountCredentialResponse) GetBody() *ResetHostAccountCredentialResponseBody {
	return s.Body
}

func (s *ResetHostAccountCredentialResponse) SetHeaders(v map[string]*string) *ResetHostAccountCredentialResponse {
	s.Headers = v
	return s
}

func (s *ResetHostAccountCredentialResponse) SetStatusCode(v int32) *ResetHostAccountCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetHostAccountCredentialResponse) SetBody(v *ResetHostAccountCredentialResponseBody) *ResetHostAccountCredentialResponse {
	s.Body = v
	return s
}

func (s *ResetHostAccountCredentialResponse) Validate() error {
	return dara.Validate(s)
}
