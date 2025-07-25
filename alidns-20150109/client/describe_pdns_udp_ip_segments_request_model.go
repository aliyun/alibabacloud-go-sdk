// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsUdpIpSegmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePdnsUdpIpSegmentsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribePdnsUdpIpSegmentsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsUdpIpSegmentsRequest
	GetPageSize() *int64
}

type DescribePdnsUdpIpSegmentsRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePdnsUdpIpSegmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsUdpIpSegmentsRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsUdpIpSegmentsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsUdpIpSegmentsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsUdpIpSegmentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsUdpIpSegmentsRequest) SetLang(v string) *DescribePdnsUdpIpSegmentsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsRequest) SetPageNumber(v int64) *DescribePdnsUdpIpSegmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsRequest) SetPageSize(v int64) *DescribePdnsUdpIpSegmentsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsUdpIpSegmentsRequest) Validate() error {
	return dara.Validate(s)
}
