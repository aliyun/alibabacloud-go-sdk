// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCrossAccountVerifyTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CrossAccountVerifyTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CrossAccountVerifyTokenResponse
	GetStatusCode() *int32
	SetBody(v *CrossAccountVerifyTokenResponseBody) *CrossAccountVerifyTokenResponse
	GetBody() *CrossAccountVerifyTokenResponseBody
}

type CrossAccountVerifyTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CrossAccountVerifyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CrossAccountVerifyTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CrossAccountVerifyTokenResponse) GoString() string {
	return s.String()
}

func (s *CrossAccountVerifyTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CrossAccountVerifyTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CrossAccountVerifyTokenResponse) GetBody() *CrossAccountVerifyTokenResponseBody {
	return s.Body
}

func (s *CrossAccountVerifyTokenResponse) SetHeaders(v map[string]*string) *CrossAccountVerifyTokenResponse {
	s.Headers = v
	return s
}

func (s *CrossAccountVerifyTokenResponse) SetStatusCode(v int32) *CrossAccountVerifyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CrossAccountVerifyTokenResponse) SetBody(v *CrossAccountVerifyTokenResponseBody) *CrossAccountVerifyTokenResponse {
	s.Body = v
	return s
}

func (s *CrossAccountVerifyTokenResponse) Validate() error {
	return dara.Validate(s)
}
