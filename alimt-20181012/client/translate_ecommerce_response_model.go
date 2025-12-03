// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateECommerceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateECommerceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateECommerceResponse
	GetStatusCode() *int32
	SetBody(v *TranslateECommerceResponseBody) *TranslateECommerceResponse
	GetBody() *TranslateECommerceResponseBody
}

type TranslateECommerceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateECommerceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateECommerceResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateECommerceResponse) GoString() string {
	return s.String()
}

func (s *TranslateECommerceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateECommerceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateECommerceResponse) GetBody() *TranslateECommerceResponseBody {
	return s.Body
}

func (s *TranslateECommerceResponse) SetHeaders(v map[string]*string) *TranslateECommerceResponse {
	s.Headers = v
	return s
}

func (s *TranslateECommerceResponse) SetStatusCode(v int32) *TranslateECommerceResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateECommerceResponse) SetBody(v *TranslateECommerceResponseBody) *TranslateECommerceResponse {
	s.Body = v
	return s
}

func (s *TranslateECommerceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
