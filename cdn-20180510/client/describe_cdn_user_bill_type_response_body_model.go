// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillTypeData(v *DescribeCdnUserBillTypeResponseBodyBillTypeData) *DescribeCdnUserBillTypeResponseBody
	GetBillTypeData() *DescribeCdnUserBillTypeResponseBodyBillTypeData
	SetRequestId(v string) *DescribeCdnUserBillTypeResponseBody
	GetRequestId() *string
}

type DescribeCdnUserBillTypeResponseBody struct {
	BillTypeData *DescribeCdnUserBillTypeResponseBodyBillTypeData `json:"BillTypeData,omitempty" xml:"BillTypeData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnUserBillTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBody) GetBillTypeData() *DescribeCdnUserBillTypeResponseBodyBillTypeData {
	return s.BillTypeData
}

func (s *DescribeCdnUserBillTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnUserBillTypeResponseBody) SetBillTypeData(v *DescribeCdnUserBillTypeResponseBodyBillTypeData) *DescribeCdnUserBillTypeResponseBody {
	s.BillTypeData = v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBody) SetRequestId(v string) *DescribeCdnUserBillTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBody) Validate() error {
	if s.BillTypeData != nil {
		if err := s.BillTypeData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnUserBillTypeResponseBodyBillTypeData struct {
	BillTypeDataItem []*DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem `json:"BillTypeDataItem,omitempty" xml:"BillTypeDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeData) GetBillTypeDataItem() []*DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	return s.BillTypeDataItem
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeData) SetBillTypeDataItem(v []*DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) *DescribeCdnUserBillTypeResponseBodyBillTypeData {
	s.BillTypeDataItem = v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeData) Validate() error {
	if s.BillTypeDataItem != nil {
		for _, item := range s.BillTypeDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem struct {
	BillType     *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Dimension    *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetBillType() *string {
	return s.BillType
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetProduct() *string {
	return s.Product
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillType(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillingCycle(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillingCycle = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetDimension(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Dimension = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetEndTime(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetProduct(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Product = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetStartTime(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) Validate() error {
	return dara.Validate(s)
}
