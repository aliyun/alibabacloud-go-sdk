// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnOrderCommodityCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnOrderCommodityCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnOrderCommodityCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnOrderCommodityCodeResponseBody) *DescribeCdnOrderCommodityCodeResponse
	GetBody() *DescribeCdnOrderCommodityCodeResponseBody
}

type DescribeCdnOrderCommodityCodeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnOrderCommodityCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnOrderCommodityCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnOrderCommodityCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnOrderCommodityCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnOrderCommodityCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnOrderCommodityCodeResponse) GetBody() *DescribeCdnOrderCommodityCodeResponseBody {
	return s.Body
}

func (s *DescribeCdnOrderCommodityCodeResponse) SetHeaders(v map[string]*string) *DescribeCdnOrderCommodityCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnOrderCommodityCodeResponse) SetStatusCode(v int32) *DescribeCdnOrderCommodityCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnOrderCommodityCodeResponse) SetBody(v *DescribeCdnOrderCommodityCodeResponseBody) *DescribeCdnOrderCommodityCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnOrderCommodityCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
