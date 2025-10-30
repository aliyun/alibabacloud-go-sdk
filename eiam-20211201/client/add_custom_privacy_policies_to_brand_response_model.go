// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomPrivacyPoliciesToBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomPrivacyPoliciesToBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomPrivacyPoliciesToBrandResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomPrivacyPoliciesToBrandResponseBody) *AddCustomPrivacyPoliciesToBrandResponse
	GetBody() *AddCustomPrivacyPoliciesToBrandResponseBody
}

type AddCustomPrivacyPoliciesToBrandResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomPrivacyPoliciesToBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomPrivacyPoliciesToBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomPrivacyPoliciesToBrandResponse) GoString() string {
	return s.String()
}

func (s *AddCustomPrivacyPoliciesToBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomPrivacyPoliciesToBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomPrivacyPoliciesToBrandResponse) GetBody() *AddCustomPrivacyPoliciesToBrandResponseBody {
	return s.Body
}

func (s *AddCustomPrivacyPoliciesToBrandResponse) SetHeaders(v map[string]*string) *AddCustomPrivacyPoliciesToBrandResponse {
	s.Headers = v
	return s
}

func (s *AddCustomPrivacyPoliciesToBrandResponse) SetStatusCode(v int32) *AddCustomPrivacyPoliciesToBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomPrivacyPoliciesToBrandResponse) SetBody(v *AddCustomPrivacyPoliciesToBrandResponseBody) *AddCustomPrivacyPoliciesToBrandResponse {
	s.Body = v
	return s
}

func (s *AddCustomPrivacyPoliciesToBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
