// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsymmetricSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsymmetricSignResponse
	GetStatusCode() *int32
	SetBody(v *AsymmetricSignResponseBody) *AsymmetricSignResponse
	GetBody() *AsymmetricSignResponseBody
}

type AsymmetricSignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsymmetricSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsymmetricSignResponse) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricSignResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsymmetricSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsymmetricSignResponse) GetBody() *AsymmetricSignResponseBody {
	return s.Body
}

func (s *AsymmetricSignResponse) SetHeaders(v map[string]*string) *AsymmetricSignResponse {
	s.Headers = v
	return s
}

func (s *AsymmetricSignResponse) SetStatusCode(v int32) *AsymmetricSignResponse {
	s.StatusCode = &v
	return s
}

func (s *AsymmetricSignResponse) SetBody(v *AsymmetricSignResponseBody) *AsymmetricSignResponse {
	s.Body = v
	return s
}

func (s *AsymmetricSignResponse) Validate() error {
	return dara.Validate(s)
}
