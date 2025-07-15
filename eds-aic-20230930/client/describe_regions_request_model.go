// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
	SetSaleMode(v string) *DescribeRegionsRequest
	GetSaleMode() *string
}

type DescribeRegionsRequest struct {
	// The display language of the console. Valid values:
	//
	// 	- cn: Simplified Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// en
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The sales mode.
	//
	// Valid values:
	//
	// 	- Instance: the instance group mode. [Default]
	//
	// 	- Node: the matrix mode. [Whitelist required]
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetSaleMode(v string) *DescribeRegionsRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
