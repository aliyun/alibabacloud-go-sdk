// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkuSaleInfoListResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SkuSaleInfoListResult
	GetRequestId() *string
	SetSkuSaleInfos(v []*SkuSaleInfo) *SkuSaleInfoListResult
	GetSkuSaleInfos() []*SkuSaleInfo
}

type SkuSaleInfoListResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId    *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SkuSaleInfos []*SkuSaleInfo `json:"skuSaleInfos,omitempty" xml:"skuSaleInfos,omitempty" type:"Repeated"`
}

func (s SkuSaleInfoListResult) String() string {
	return dara.Prettify(s)
}

func (s SkuSaleInfoListResult) GoString() string {
	return s.String()
}

func (s *SkuSaleInfoListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *SkuSaleInfoListResult) GetSkuSaleInfos() []*SkuSaleInfo {
	return s.SkuSaleInfos
}

func (s *SkuSaleInfoListResult) SetRequestId(v string) *SkuSaleInfoListResult {
	s.RequestId = &v
	return s
}

func (s *SkuSaleInfoListResult) SetSkuSaleInfos(v []*SkuSaleInfo) *SkuSaleInfoListResult {
	s.SkuSaleInfos = v
	return s
}

func (s *SkuSaleInfoListResult) Validate() error {
	if s.SkuSaleInfos != nil {
		for _, item := range s.SkuSaleInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
