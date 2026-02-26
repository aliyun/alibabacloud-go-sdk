// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShopPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ShopPageResult
	GetRequestId() *string
	SetShopList(v []*ShopPageDataResult) *ShopPageResult
	GetShopList() []*ShopPageDataResult
	SetTotal(v int32) *ShopPageResult
	GetTotal() *int32
}

type ShopPageResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ShopList  []*ShopPageDataResult `json:"shopList,omitempty" xml:"shopList,omitempty" type:"Repeated"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ShopPageResult) String() string {
	return dara.Prettify(s)
}

func (s ShopPageResult) GoString() string {
	return s.String()
}

func (s *ShopPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ShopPageResult) GetShopList() []*ShopPageDataResult {
	return s.ShopList
}

func (s *ShopPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *ShopPageResult) SetRequestId(v string) *ShopPageResult {
	s.RequestId = &v
	return s
}

func (s *ShopPageResult) SetShopList(v []*ShopPageDataResult) *ShopPageResult {
	s.ShopList = v
	return s
}

func (s *ShopPageResult) SetTotal(v int32) *ShopPageResult {
	s.Total = &v
	return s
}

func (s *ShopPageResult) Validate() error {
	if s.ShopList != nil {
		for _, item := range s.ShopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
