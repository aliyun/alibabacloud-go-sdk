// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForModifyDesktopOversoldGroupSaleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePriceForModifyDesktopOversoldGroupSaleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePriceForModifyDesktopOversoldGroupSaleResponse
	GetStatusCode() *int32
	SetBody(v *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) *DescribePriceForModifyDesktopOversoldGroupSaleResponse
	GetBody() *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody
}

type DescribePriceForModifyDesktopOversoldGroupSaleResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePriceForModifyDesktopOversoldGroupSaleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForModifyDesktopOversoldGroupSaleResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponse) GetBody() *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody {
	return s.Body
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponse) SetHeaders(v map[string]*string) *DescribePriceForModifyDesktopOversoldGroupSaleResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponse) SetStatusCode(v int32) *DescribePriceForModifyDesktopOversoldGroupSaleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponse) SetBody(v *DescribePriceForModifyDesktopOversoldGroupSaleResponseBody) *DescribePriceForModifyDesktopOversoldGroupSaleResponse {
	s.Body = v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
