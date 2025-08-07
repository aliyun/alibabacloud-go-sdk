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
	// The renewal information about the clusters.
	Items *DescribeAutoRenewAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 65D7ACE6-4A61-4B6E-B357-8CB24A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute struct {
	// Indicates whether the auto-renewal feature is enabled. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The renewal duration.
	//
	// example:
	//
	// 4
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The unit of the duration. Valid values:
	//
	// 	- Year
	//
	// 	- Month
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The renewal status. Valid values:
	//
	// 	- AutoRenewal: The cluster is automatically renewed.
	//
	// 	- Normal: The cluster is manually renewed. The system sends a text message to remind you before the cluster expires.
	//
	// 	- NotRenewal: The cluster is not renewed. The system does not send a reminder for expiration but only sends a text message three days before the cluster expires to remind you that the cluster is not renewed.
	//
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
