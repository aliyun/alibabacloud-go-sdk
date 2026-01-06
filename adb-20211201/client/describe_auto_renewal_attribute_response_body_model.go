// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRenewalAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeAutoRenewalAttributeResponseBodyItems) *DescribeAutoRenewalAttributeResponseBody
	GetItems() *DescribeAutoRenewalAttributeResponseBodyItems
	SetPageNumber(v int32) *DescribeAutoRenewalAttributeResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeAutoRenewalAttributeResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeAutoRenewalAttributeResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeAutoRenewalAttributeResponseBody
	GetTotalRecordCount() *int32
}

type DescribeAutoRenewalAttributeResponseBody struct {
	// The list of auto-renewal details.
	Items *DescribeAutoRenewalAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BA0F6761-7A8C-59F8-9624-FB56788C0EDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAutoRenewalAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewalAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewalAttributeResponseBody) GetItems() *DescribeAutoRenewalAttributeResponseBodyItems {
	return s.Items
}

func (s *DescribeAutoRenewalAttributeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoRenewalAttributeResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAutoRenewalAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoRenewalAttributeResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeAutoRenewalAttributeResponseBody) SetItems(v *DescribeAutoRenewalAttributeResponseBodyItems) *DescribeAutoRenewalAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBody) SetPageNumber(v int32) *DescribeAutoRenewalAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBody) SetPageRecordCount(v int32) *DescribeAutoRenewalAttributeResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBody) SetRequestId(v string) *DescribeAutoRenewalAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBody) SetTotalRecordCount(v int32) *DescribeAutoRenewalAttributeResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoRenewalAttributeResponseBodyItems struct {
	AutoRenewalAttribute []*DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute `json:"AutoRenewalAttribute,omitempty" xml:"AutoRenewalAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAutoRenewalAttributeResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewalAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewalAttributeResponseBodyItems) GetAutoRenewalAttribute() []*DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute {
	return s.AutoRenewalAttribute
}

func (s *DescribeAutoRenewalAttributeResponseBodyItems) SetAutoRenewalAttribute(v []*DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) *DescribeAutoRenewalAttributeResponseBodyItems {
	s.AutoRenewalAttribute = v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBodyItems) Validate() error {
	if s.AutoRenewalAttribute != nil {
		for _, item := range s.AutoRenewalAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute struct {
	// Indicates whether auto-renewal is enabled for the cluster. Valid values:
	//
	// 	- **true**: Enables.
	//
	// 	- **false**: Disables.
	//
	// example:
	//
	// true
	AutoRenewalEnabled *bool `json:"AutoRenewalEnabled,omitempty" xml:"AutoRenewalEnabled,omitempty"`
	// The auto-renewal duration.
	//
	// example:
	//
	// 1
	AutoRenewalPeriod *int64 `json:"AutoRenewalPeriod,omitempty" xml:"AutoRenewalPeriod,omitempty"`
	// The unit of auto-renewal duration. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// example:
	//
	// Year
	AutoRenewalPeriodUnit *string `json:"AutoRenewalPeriodUnit,omitempty" xml:"AutoRenewalPeriodUnit,omitempty"`
	// The renewal method. Valid values:
	//
	// 	- **AutoRenewal**: The cluster is automatically renewed.
	//
	// 	- **Normal**: The cluster is manually renewed. Before the cluster expires, the system sends you a reminder by SMS message.
	//
	// 	- **NotRenewal**: The cluster is not renewed. Reminders are only sent three days before cluster expiration.
	//
	// example:
	//
	// AutoRenewal
	AutoRenewalStatus *string `json:"AutoRenewalStatus,omitempty" xml:"AutoRenewalStatus,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-uf6485635fz8****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) GetAutoRenewalEnabled() *bool {
	return s.AutoRenewalEnabled
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) GetAutoRenewalPeriod() *int64 {
	return s.AutoRenewalPeriod
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) GetAutoRenewalPeriodUnit() *string {
	return s.AutoRenewalPeriodUnit
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) GetAutoRenewalStatus() *string {
	return s.AutoRenewalStatus
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) SetAutoRenewalEnabled(v bool) *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute {
	s.AutoRenewalEnabled = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) SetAutoRenewalPeriod(v int64) *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute {
	s.AutoRenewalPeriod = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) SetAutoRenewalPeriodUnit(v string) *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute {
	s.AutoRenewalPeriodUnit = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) SetAutoRenewalStatus(v string) *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute {
	s.AutoRenewalStatus = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) SetDBClusterId(v string) *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) SetRegionId(v string) *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeResponseBodyItemsAutoRenewalAttribute) Validate() error {
	return dara.Validate(s)
}
