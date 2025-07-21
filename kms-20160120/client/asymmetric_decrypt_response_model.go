// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricDecryptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsymmetricDecryptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsymmetricDecryptResponse
	GetStatusCode() *int32
	SetBody(v *AsymmetricDecryptResponseBody) *AsymmetricDecryptResponse
	GetBody() *AsymmetricDecryptResponseBody
}

type AsymmetricDecryptResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsymmetricDecryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsymmetricDecryptResponse) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricDecryptResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsymmetricDecryptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsymmetricDecryptResponse) GetBody() *AsymmetricDecryptResponseBody {
	return s.Body
}

func (s *AsymmetricDecryptResponse) SetHeaders(v map[string]*string) *AsymmetricDecryptResponse {
	s.Headers = v
	return s
}

func (s *AsymmetricDecryptResponse) SetStatusCode(v int32) *AsymmetricDecryptResponse {
	s.StatusCode = &v
	return s
}

func (s *AsymmetricDecryptResponse) SetBody(v *AsymmetricDecryptResponseBody) *AsymmetricDecryptResponse {
	s.Body = v
	return s
}

func (s *AsymmetricDecryptResponse) Validate() error {
	return dara.Validate(s)
}
