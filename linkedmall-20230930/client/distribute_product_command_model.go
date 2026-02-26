// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistributeProductCommand interface {
	dara.Model
	String() string
	GoString() string
	SetLmShopId(v string) *DistributeProductCommand
	GetLmShopId() *string
	SetProducts(v []*DistributionProduct) *DistributeProductCommand
	GetProducts() []*DistributionProduct
	SetRequestId(v string) *DistributeProductCommand
	GetRequestId() *string
	SetRequestTime(v string) *DistributeProductCommand
	GetRequestTime() *string
	SetRequestUser(v string) *DistributeProductCommand
	GetRequestUser() *string
}

type DistributeProductCommand struct {
	LmShopId  *string                `json:"lmShopId,omitempty" xml:"lmShopId,omitempty"`
	Products  []*DistributionProduct `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-12-01 10:01:00
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	RequestUser *string `json:"requestUser,omitempty" xml:"requestUser,omitempty"`
}

func (s DistributeProductCommand) String() string {
	return dara.Prettify(s)
}

func (s DistributeProductCommand) GoString() string {
	return s.String()
}

func (s *DistributeProductCommand) GetLmShopId() *string {
	return s.LmShopId
}

func (s *DistributeProductCommand) GetProducts() []*DistributionProduct {
	return s.Products
}

func (s *DistributeProductCommand) GetRequestId() *string {
	return s.RequestId
}

func (s *DistributeProductCommand) GetRequestTime() *string {
	return s.RequestTime
}

func (s *DistributeProductCommand) GetRequestUser() *string {
	return s.RequestUser
}

func (s *DistributeProductCommand) SetLmShopId(v string) *DistributeProductCommand {
	s.LmShopId = &v
	return s
}

func (s *DistributeProductCommand) SetProducts(v []*DistributionProduct) *DistributeProductCommand {
	s.Products = v
	return s
}

func (s *DistributeProductCommand) SetRequestId(v string) *DistributeProductCommand {
	s.RequestId = &v
	return s
}

func (s *DistributeProductCommand) SetRequestTime(v string) *DistributeProductCommand {
	s.RequestTime = &v
	return s
}

func (s *DistributeProductCommand) SetRequestUser(v string) *DistributeProductCommand {
	s.RequestUser = &v
	return s
}

func (s *DistributeProductCommand) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
