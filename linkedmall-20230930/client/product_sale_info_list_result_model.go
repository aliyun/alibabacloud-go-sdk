// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductSaleInfoListResult interface {
	dara.Model
	String() string
	GoString() string
	SetProductSaleInfos(v []*ProductSaleInfo) *ProductSaleInfoListResult
	GetProductSaleInfos() []*ProductSaleInfo
	SetRequestId(v string) *ProductSaleInfoListResult
	GetRequestId() *string
}

type ProductSaleInfoListResult struct {
	ProductSaleInfos []*ProductSaleInfo `json:"productSaleInfos,omitempty" xml:"productSaleInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ProductSaleInfoListResult) String() string {
	return dara.Prettify(s)
}

func (s ProductSaleInfoListResult) GoString() string {
	return s.String()
}

func (s *ProductSaleInfoListResult) GetProductSaleInfos() []*ProductSaleInfo {
	return s.ProductSaleInfos
}

func (s *ProductSaleInfoListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ProductSaleInfoListResult) SetProductSaleInfos(v []*ProductSaleInfo) *ProductSaleInfoListResult {
	s.ProductSaleInfos = v
	return s
}

func (s *ProductSaleInfoListResult) SetRequestId(v string) *ProductSaleInfoListResult {
	s.RequestId = &v
	return s
}

func (s *ProductSaleInfoListResult) Validate() error {
	if s.ProductSaleInfos != nil {
		for _, item := range s.ProductSaleInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
