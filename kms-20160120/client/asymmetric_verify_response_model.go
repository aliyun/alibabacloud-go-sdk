// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsymmetricVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsymmetricVerifyResponse
	GetStatusCode() *int32
	SetBody(v *AsymmetricVerifyResponseBody) *AsymmetricVerifyResponse
	GetBody() *AsymmetricVerifyResponseBody
}

type AsymmetricVerifyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsymmetricVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsymmetricVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricVerifyResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsymmetricVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsymmetricVerifyResponse) GetBody() *AsymmetricVerifyResponseBody {
	return s.Body
}

func (s *AsymmetricVerifyResponse) SetHeaders(v map[string]*string) *AsymmetricVerifyResponse {
	s.Headers = v
	return s
}

func (s *AsymmetricVerifyResponse) SetStatusCode(v int32) *AsymmetricVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *AsymmetricVerifyResponse) SetBody(v *AsymmetricVerifyResponseBody) *AsymmetricVerifyResponse {
	s.Body = v
	return s
}

func (s *AsymmetricVerifyResponse) Validate() error {
	return dara.Validate(s)
}
