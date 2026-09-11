// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedPricingApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupportedPricingApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupportedPricingApisResponse
	GetStatusCode() *int32
	SetBody(v *ListSupportedPricingApisResponseBody) *ListSupportedPricingApisResponse
	GetBody() *ListSupportedPricingApisResponseBody
}

type ListSupportedPricingApisResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupportedPricingApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupportedPricingApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedPricingApisResponse) GoString() string {
	return s.String()
}

func (s *ListSupportedPricingApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupportedPricingApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupportedPricingApisResponse) GetBody() *ListSupportedPricingApisResponseBody {
	return s.Body
}

func (s *ListSupportedPricingApisResponse) SetHeaders(v map[string]*string) *ListSupportedPricingApisResponse {
	s.Headers = v
	return s
}

func (s *ListSupportedPricingApisResponse) SetStatusCode(v int32) *ListSupportedPricingApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportedPricingApisResponse) SetBody(v *ListSupportedPricingApisResponseBody) *ListSupportedPricingApisResponse {
	s.Body = v
	return s
}

func (s *ListSupportedPricingApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
