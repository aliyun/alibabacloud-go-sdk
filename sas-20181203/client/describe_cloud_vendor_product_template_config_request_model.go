// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudVendorProductTemplateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCloudVendorProductTemplateConfigRequest
	GetLang() *string
	SetVendor(v string) *DescribeCloudVendorProductTemplateConfigRequest
	GetVendor() *string
}

type DescribeCloudVendorProductTemplateConfigRequest struct {
	// Set the language type for request and response messages, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Cloud asset vendor. Values:
	//
	// - **CHAITIN**: Chaitin Technology
	//
	// - **FORTINET**: Fortinet
	//
	// - **THREATBOOK**: ThreatBook
	//
	// example:
	//
	// CHAITIN
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeCloudVendorProductTemplateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorProductTemplateConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorProductTemplateConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCloudVendorProductTemplateConfigRequest) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeCloudVendorProductTemplateConfigRequest) SetLang(v string) *DescribeCloudVendorProductTemplateConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCloudVendorProductTemplateConfigRequest) SetVendor(v string) *DescribeCloudVendorProductTemplateConfigRequest {
	s.Vendor = &v
	return s
}

func (s *DescribeCloudVendorProductTemplateConfigRequest) Validate() error {
	return dara.Validate(s)
}
