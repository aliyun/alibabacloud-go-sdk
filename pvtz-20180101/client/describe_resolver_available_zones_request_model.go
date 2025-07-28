// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverAvailableZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAzId(v string) *DescribeResolverAvailableZonesRequest
	GetAzId() *string
	SetLang(v string) *DescribeResolverAvailableZonesRequest
	GetLang() *string
	SetResolverRegionId(v string) *DescribeResolverAvailableZonesRequest
	GetResolverRegionId() *string
}

type DescribeResolverAvailableZonesRequest struct {
	// The zone ID.
	//
	// example:
	//
	// cn-zhangjiakou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	ResolverRegionId *string `json:"ResolverRegionId,omitempty" xml:"ResolverRegionId,omitempty"`
}

func (s DescribeResolverAvailableZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverAvailableZonesRequest) GetAzId() *string {
	return s.AzId
}

func (s *DescribeResolverAvailableZonesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeResolverAvailableZonesRequest) GetResolverRegionId() *string {
	return s.ResolverRegionId
}

func (s *DescribeResolverAvailableZonesRequest) SetAzId(v string) *DescribeResolverAvailableZonesRequest {
	s.AzId = &v
	return s
}

func (s *DescribeResolverAvailableZonesRequest) SetLang(v string) *DescribeResolverAvailableZonesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverAvailableZonesRequest) SetResolverRegionId(v string) *DescribeResolverAvailableZonesRequest {
	s.ResolverRegionId = &v
	return s
}

func (s *DescribeResolverAvailableZonesRequest) Validate() error {
	return dara.Validate(s)
}
