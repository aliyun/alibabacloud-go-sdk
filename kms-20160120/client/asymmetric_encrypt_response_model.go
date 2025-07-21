// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricEncryptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsymmetricEncryptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsymmetricEncryptResponse
	GetStatusCode() *int32
	SetBody(v *AsymmetricEncryptResponseBody) *AsymmetricEncryptResponse
	GetBody() *AsymmetricEncryptResponseBody
}

type AsymmetricEncryptResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsymmetricEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsymmetricEncryptResponse) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricEncryptResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsymmetricEncryptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsymmetricEncryptResponse) GetBody() *AsymmetricEncryptResponseBody {
	return s.Body
}

func (s *AsymmetricEncryptResponse) SetHeaders(v map[string]*string) *AsymmetricEncryptResponse {
	s.Headers = v
	return s
}

func (s *AsymmetricEncryptResponse) SetStatusCode(v int32) *AsymmetricEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *AsymmetricEncryptResponse) SetBody(v *AsymmetricEncryptResponseBody) *AsymmetricEncryptResponse {
	s.Body = v
	return s
}

func (s *AsymmetricEncryptResponse) Validate() error {
	return dara.Validate(s)
}
