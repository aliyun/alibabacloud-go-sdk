// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeAutoRenewAttributeResponseBodyItems) *DescribeAutoRenewAttributeResponseBody
	GetItems() *DescribeAutoRenewAttributeResponseBodyItems
	SetPageNumber(v int32) *DescribeAutoRenewAttributeResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeAutoRenewAttributeResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody
	GetTotalRecordCount() *int32
}

type DescribeAutoRenewAttributeResponseBody struct {
	Items *DescribeAutoRenewAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 51
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBody) GetItems() *DescribeAutoRenewAttributeResponseBodyItems {
	return s.Items
}

func (s *DescribeAutoRenewAttributeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoRenewAttributeResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoRenewAttributeResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeAutoRenewAttributeResponseBody) SetItems(v *DescribeAutoRenewAttributeResponseBodyItems) *DescribeAutoRenewAttributeResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetPageNumber(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetPageRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetTotalRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoRenewAttributeResponseBodyItems struct {
	AutoRenewAttribute []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute `json:"AutoRenewAttribute,omitempty" xml:"AutoRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAutoRenewAttributeResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItems) GetAutoRenewAttribute() []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	return s.AutoRenewAttribute
}

func (s *DescribeAutoRenewAttributeResponseBodyItems) SetAutoRenewAttribute(v []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) *DescribeAutoRenewAttributeResponseBodyItems {
	s.AutoRenewAttribute = v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItems) Validate() error {
	if s.AutoRenewAttribute != nil {
		for _, item := range s.AutoRenewAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute struct {
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	// example:
	//
	// cc-2ze57pg09*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// AutoRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GetAutoRenewEnabled() *bool {
	return s.AutoRenewEnabled
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetAutoRenewEnabled(v bool) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.AutoRenewEnabled = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDBClusterId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDuration(v int32) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetPeriodUnit(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRegionId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRenewalStatus(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) Validate() error {
	return dara.Validate(s)
}
