// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstallCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddInstallCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddInstallCodeResponse
	GetStatusCode() *int32
	SetBody(v *AddInstallCodeResponseBody) *AddInstallCodeResponse
	GetBody() *AddInstallCodeResponseBody
}

type AddInstallCodeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddInstallCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddInstallCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddInstallCodeResponse) GoString() string {
	return s.String()
}

func (s *AddInstallCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddInstallCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddInstallCodeResponse) GetBody() *AddInstallCodeResponseBody {
	return s.Body
}

func (s *AddInstallCodeResponse) SetHeaders(v map[string]*string) *AddInstallCodeResponse {
	s.Headers = v
	return s
}

func (s *AddInstallCodeResponse) SetStatusCode(v int32) *AddInstallCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInstallCodeResponse) SetBody(v *AddInstallCodeResponseBody) *AddInstallCodeResponse {
	s.Body = v
	return s
}

func (s *AddInstallCodeResponse) Validate() error {
	return dara.Validate(s)
}
