// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePurchasedApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePurchasedApisResponse
	GetStatusCode() *int32
	SetBody(v *DescribePurchasedApisResponseBody) *DescribePurchasedApisResponse
	GetBody() *DescribePurchasedApisResponseBody
}

type DescribePurchasedApisResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedApisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApisResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePurchasedApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePurchasedApisResponse) GetBody() *DescribePurchasedApisResponseBody {
	return s.Body
}

func (s *DescribePurchasedApisResponse) SetHeaders(v map[string]*string) *DescribePurchasedApisResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedApisResponse) SetStatusCode(v int32) *DescribePurchasedApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedApisResponse) SetBody(v *DescribePurchasedApisResponseBody) *DescribePurchasedApisResponse {
	s.Body = v
	return s
}

func (s *DescribePurchasedApisResponse) Validate() error {
	return dara.Validate(s)
}
