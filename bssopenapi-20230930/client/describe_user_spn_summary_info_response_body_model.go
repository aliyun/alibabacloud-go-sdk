// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSpnSummaryInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceFamilyList(v []*string) *DescribeUserSpnSummaryInfoResponseBody
	GetInstanceFamilyList() []*string
	SetRegionList(v []*DescribeUserSpnSummaryInfoResponseBodyRegionList) *DescribeUserSpnSummaryInfoResponseBody
	GetRegionList() []*DescribeUserSpnSummaryInfoResponseBodyRegionList
	SetRequestId(v string) *DescribeUserSpnSummaryInfoResponseBody
	GetRequestId() *string
	SetSpnCodeAndTypeList(v []*DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) *DescribeUserSpnSummaryInfoResponseBody
	GetSpnCodeAndTypeList() []*DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList
}

type DescribeUserSpnSummaryInfoResponseBody struct {
	InstanceFamilyList []*string                                                   `json:"InstanceFamilyList,omitempty" xml:"InstanceFamilyList,omitempty" type:"Repeated"`
	RegionList         []*DescribeUserSpnSummaryInfoResponseBodyRegionList         `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpnCodeAndTypeList []*DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList `json:"SpnCodeAndTypeList,omitempty" xml:"SpnCodeAndTypeList,omitempty" type:"Repeated"`
}

func (s DescribeUserSpnSummaryInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoResponseBody) GetInstanceFamilyList() []*string {
	return s.InstanceFamilyList
}

func (s *DescribeUserSpnSummaryInfoResponseBody) GetRegionList() []*DescribeUserSpnSummaryInfoResponseBodyRegionList {
	return s.RegionList
}

func (s *DescribeUserSpnSummaryInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserSpnSummaryInfoResponseBody) GetSpnCodeAndTypeList() []*DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	return s.SpnCodeAndTypeList
}

func (s *DescribeUserSpnSummaryInfoResponseBody) SetInstanceFamilyList(v []*string) *DescribeUserSpnSummaryInfoResponseBody {
	s.InstanceFamilyList = v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBody) SetRegionList(v []*DescribeUserSpnSummaryInfoResponseBodyRegionList) *DescribeUserSpnSummaryInfoResponseBody {
	s.RegionList = v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBody) SetRequestId(v string) *DescribeUserSpnSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBody) SetSpnCodeAndTypeList(v []*DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) *DescribeUserSpnSummaryInfoResponseBody {
	s.SpnCodeAndTypeList = v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserSpnSummaryInfoResponseBodyRegionList struct {
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeUserSpnSummaryInfoResponseBodyRegionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoResponseBodyRegionList) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoResponseBodyRegionList) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeUserSpnSummaryInfoResponseBodyRegionList) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeUserSpnSummaryInfoResponseBodyRegionList) SetRegionCode(v string) *DescribeUserSpnSummaryInfoResponseBodyRegionList {
	s.RegionCode = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodyRegionList) SetRegionName(v string) *DescribeUserSpnSummaryInfoResponseBodyRegionList {
	s.RegionName = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodyRegionList) Validate() error {
	return dara.Validate(s)
}

type DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList struct {
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	SpnCommodityCode *string `json:"SpnCommodityCode,omitempty" xml:"SpnCommodityCode,omitempty"`
	SpnType          *string `json:"SpnType,omitempty" xml:"SpnType,omitempty"`
	SpnTypeName      *string `json:"SpnTypeName,omitempty" xml:"SpnTypeName,omitempty"`
}

func (s DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) GetSpnCommodityCode() *string {
	return s.SpnCommodityCode
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) GetSpnType() *string {
	return s.SpnType
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) GetSpnTypeName() *string {
	return s.SpnTypeName
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) SetProductCode(v string) *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	s.ProductCode = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) SetSpnCommodityCode(v string) *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	s.SpnCommodityCode = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) SetSpnType(v string) *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	s.SpnType = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) SetSpnTypeName(v string) *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList {
	s.SpnTypeName = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponseBodySpnCodeAndTypeList) Validate() error {
	return dara.Validate(s)
}
