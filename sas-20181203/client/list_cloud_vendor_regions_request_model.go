// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudVendorRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListCloudVendorRegionsRequest
	GetLang() *string
	SetVendor(v string) *ListCloudVendorRegionsRequest
	GetVendor() *string
}

type ListCloudVendorRegionsRequest struct {
	// The language of the content in the request and response messages. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The service provider of cloud assets. Valid values:
	//
	// 	- **Tencent**, **HUAWEICLOUD**, **Azure**, and **AWS**: other service providers of cloud assets.
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListCloudVendorRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudVendorRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudVendorRegionsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCloudVendorRegionsRequest) GetVendor() *string {
	return s.Vendor
}

func (s *ListCloudVendorRegionsRequest) SetLang(v string) *ListCloudVendorRegionsRequest {
	s.Lang = &v
	return s
}

func (s *ListCloudVendorRegionsRequest) SetVendor(v string) *ListCloudVendorRegionsRequest {
	s.Vendor = &v
	return s
}

func (s *ListCloudVendorRegionsRequest) Validate() error {
	return dara.Validate(s)
}
