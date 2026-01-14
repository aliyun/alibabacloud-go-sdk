// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommodityPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommodityPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommodityPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommodityPriceResponseBody) *DescribeCommodityPriceResponse
	GetBody() *DescribeCommodityPriceResponseBody
}

type DescribeCommodityPriceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommodityPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommodityPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommodityPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommodityPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommodityPriceResponse) GetBody() *DescribeCommodityPriceResponseBody {
	return s.Body
}

func (s *DescribeCommodityPriceResponse) SetHeaders(v map[string]*string) *DescribeCommodityPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommodityPriceResponse) SetStatusCode(v int32) *DescribeCommodityPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommodityPriceResponse) SetBody(v *DescribeCommodityPriceResponseBody) *DescribeCommodityPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeCommodityPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
