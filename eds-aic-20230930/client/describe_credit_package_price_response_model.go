// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackagePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCreditPackagePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCreditPackagePriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCreditPackagePriceResponseBody) *DescribeCreditPackagePriceResponse
	GetBody() *DescribeCreditPackagePriceResponseBody
}

type DescribeCreditPackagePriceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCreditPackagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCreditPackagePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackagePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackagePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCreditPackagePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCreditPackagePriceResponse) GetBody() *DescribeCreditPackagePriceResponseBody {
	return s.Body
}

func (s *DescribeCreditPackagePriceResponse) SetHeaders(v map[string]*string) *DescribeCreditPackagePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreditPackagePriceResponse) SetStatusCode(v int32) *DescribeCreditPackagePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCreditPackagePriceResponse) SetBody(v *DescribeCreditPackagePriceResponseBody) *DescribeCreditPackagePriceResponse {
	s.Body = v
	return s
}

func (s *DescribeCreditPackagePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
