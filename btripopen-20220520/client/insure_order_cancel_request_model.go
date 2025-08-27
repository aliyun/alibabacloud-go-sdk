// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *InsureOrderCancelRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureOrderCancelRequest
	GetBuyerName() *string
	SetIsvName(v string) *InsureOrderCancelRequest
	GetIsvName() *string
	SetSupplierCode(v string) *InsureOrderCancelRequest
	GetSupplierCode() *string
}

type InsureOrderCancelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2000310301
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// example:
	//
	// open12igetbis4o07v10B1TlOWcM00
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// fliggy
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s InsureOrderCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCancelRequest) GoString() string {
	return s.String()
}

func (s *InsureOrderCancelRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderCancelRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureOrderCancelRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureOrderCancelRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureOrderCancelRequest) SetBtripUserId(v string) *InsureOrderCancelRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderCancelRequest) SetBuyerName(v string) *InsureOrderCancelRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureOrderCancelRequest) SetIsvName(v string) *InsureOrderCancelRequest {
	s.IsvName = &v
	return s
}

func (s *InsureOrderCancelRequest) SetSupplierCode(v string) *InsureOrderCancelRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureOrderCancelRequest) Validate() error {
	return dara.Validate(s)
}
