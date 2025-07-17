// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuyPayAsYouGoOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityType(v string) *BuyPayAsYouGoOrderRequest
	GetCommodityType() *string
	SetInsNum(v int32) *BuyPayAsYouGoOrderRequest
	GetInsNum() *int32
	SetTid(v int64) *BuyPayAsYouGoOrderRequest
	GetTid() *int64
	SetVersionType(v string) *BuyPayAsYouGoOrderRequest
	GetVersionType() *string
}

type BuyPayAsYouGoOrderRequest struct {
	// The type of the resource that you want to purchase.
	//
	// 	- **VersionType**: DMS that supports control modes
	//
	// 	- **SensitiveDataProtection**: DMS that supports sensitive data protection
	//
	// This parameter is required.
	//
	// example:
	//
	// SensitiveDataProtection
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	// The number of database instances that you want to use DMS to manage.
	//
	// > A quota can be used for only one database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	InsNum *int32 `json:"InsNum,omitempty" xml:"InsNum,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The control mode of DMS. Valid values:
	//
	// 	- **stand**: Stable Change
	//
	// 	- **safety**: Security Collaboration
	//
	// example:
	//
	// stand
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s BuyPayAsYouGoOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s BuyPayAsYouGoOrderRequest) GoString() string {
	return s.String()
}

func (s *BuyPayAsYouGoOrderRequest) GetCommodityType() *string {
	return s.CommodityType
}

func (s *BuyPayAsYouGoOrderRequest) GetInsNum() *int32 {
	return s.InsNum
}

func (s *BuyPayAsYouGoOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BuyPayAsYouGoOrderRequest) GetVersionType() *string {
	return s.VersionType
}

func (s *BuyPayAsYouGoOrderRequest) SetCommodityType(v string) *BuyPayAsYouGoOrderRequest {
	s.CommodityType = &v
	return s
}

func (s *BuyPayAsYouGoOrderRequest) SetInsNum(v int32) *BuyPayAsYouGoOrderRequest {
	s.InsNum = &v
	return s
}

func (s *BuyPayAsYouGoOrderRequest) SetTid(v int64) *BuyPayAsYouGoOrderRequest {
	s.Tid = &v
	return s
}

func (s *BuyPayAsYouGoOrderRequest) SetVersionType(v string) *BuyPayAsYouGoOrderRequest {
	s.VersionType = &v
	return s
}

func (s *BuyPayAsYouGoOrderRequest) Validate() error {
	return dara.Validate(s)
}
