// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCens(v *DescribeCensResponseBodyCens) *DescribeCensResponseBody
	GetCens() *DescribeCensResponseBodyCens
	SetPageNumber(v int32) *DescribeCensResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCensResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCensResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCensResponseBody
	GetTotalCount() *int32
}

type DescribeCensResponseBody struct {
	// The information about the CEN instance.
	Cens *DescribeCensResponseBodyCens `json:"Cens,omitempty" xml:"Cens,omitempty" type:"Struct"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BFA6822-240E-4E27-B4C8-AA400EF7474D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBody) GetCens() *DescribeCensResponseBodyCens {
	return s.Cens
}

func (s *DescribeCensResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCensResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCensResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCensResponseBody) SetCens(v *DescribeCensResponseBodyCens) *DescribeCensResponseBody {
	s.Cens = v
	return s
}

func (s *DescribeCensResponseBody) SetPageNumber(v int32) *DescribeCensResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensResponseBody) SetPageSize(v int32) *DescribeCensResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCensResponseBody) SetRequestId(v string) *DescribeCensResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCensResponseBody) SetTotalCount(v int32) *DescribeCensResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCensResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCensResponseBodyCens struct {
	Cen []*DescribeCensResponseBodyCensCen `json:"Cen,omitempty" xml:"Cen,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCens) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBodyCens) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCens) GetCen() []*DescribeCensResponseBodyCensCen {
	return s.Cen
}

func (s *DescribeCensResponseBodyCens) SetCen(v []*DescribeCensResponseBodyCensCen) *DescribeCensResponseBodyCens {
	s.Cen = v
	return s
}

func (s *DescribeCensResponseBodyCens) Validate() error {
	return dara.Validate(s)
}

type DescribeCensResponseBodyCensCen struct {
	// The IDs of the bandwidth plans that are associated with the CEN instance.
	CenBandwidthPackageIds *DescribeCensResponseBodyCensCenCenBandwidthPackageIds `json:"CenBandwidthPackageIds,omitempty" xml:"CenBandwidthPackageIds,omitempty" type:"Struct"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-0xyeagctz5sfg9****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the CEN instance was created.
	//
	// The time follows the ISO8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-22T07:44Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the CEN instance.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether IPv6 is enabled for the CEN instance.
	//
	// 	- **ENABLE**
	//
	// 	- **DISABLED**
	//
	// example:
	//
	// DISABLED
	Ipv6Level *string `json:"Ipv6Level,omitempty" xml:"Ipv6Level,omitempty"`
	// The CEN instance name.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The level of CIDR block overlapping.
	//
	// **REDUCED**: Overlapped CIDR blocks are allowed. This value specifies that CIDR blocks can overlap but CIDR blocks cannot be duplicates.
	//
	// example:
	//
	// REDUCED
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The ID of the resource group to which the CEN instance belongs.
	//
	// example:
	//
	// rg-acfm3unpnuw****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the CEN instance.
	//
	// 	- **Creating**
	//
	// 	- **Active**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of the tags that are added to the CEN instance.
	Tags *DescribeCensResponseBodyCensCenTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeCensResponseBodyCensCen) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBodyCensCen) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCen) GetCenBandwidthPackageIds() *DescribeCensResponseBodyCensCenCenBandwidthPackageIds {
	return s.CenBandwidthPackageIds
}

func (s *DescribeCensResponseBodyCensCen) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCensResponseBodyCensCen) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeCensResponseBodyCensCen) GetDescription() *string {
	return s.Description
}

func (s *DescribeCensResponseBodyCensCen) GetIpv6Level() *string {
	return s.Ipv6Level
}

func (s *DescribeCensResponseBodyCensCen) GetName() *string {
	return s.Name
}

func (s *DescribeCensResponseBodyCensCen) GetProtectionLevel() *string {
	return s.ProtectionLevel
}

func (s *DescribeCensResponseBodyCensCen) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCensResponseBodyCensCen) GetStatus() *string {
	return s.Status
}

func (s *DescribeCensResponseBodyCensCen) GetTags() *DescribeCensResponseBodyCensCenTags {
	return s.Tags
}

func (s *DescribeCensResponseBodyCensCen) SetCenBandwidthPackageIds(v *DescribeCensResponseBodyCensCenCenBandwidthPackageIds) *DescribeCensResponseBodyCensCen {
	s.CenBandwidthPackageIds = v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetCenId(v string) *DescribeCensResponseBodyCensCen {
	s.CenId = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetCreationTime(v string) *DescribeCensResponseBodyCensCen {
	s.CreationTime = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetDescription(v string) *DescribeCensResponseBodyCensCen {
	s.Description = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetIpv6Level(v string) *DescribeCensResponseBodyCensCen {
	s.Ipv6Level = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetName(v string) *DescribeCensResponseBodyCensCen {
	s.Name = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetProtectionLevel(v string) *DescribeCensResponseBodyCensCen {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetResourceGroupId(v string) *DescribeCensResponseBodyCensCen {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetStatus(v string) *DescribeCensResponseBodyCensCen {
	s.Status = &v
	return s
}

func (s *DescribeCensResponseBodyCensCen) SetTags(v *DescribeCensResponseBodyCensCenTags) *DescribeCensResponseBodyCensCen {
	s.Tags = v
	return s
}

func (s *DescribeCensResponseBodyCensCen) Validate() error {
	return dara.Validate(s)
}

type DescribeCensResponseBodyCensCenCenBandwidthPackageIds struct {
	CenBandwidthPackageId []*string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCensCenCenBandwidthPackageIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBodyCensCenCenBandwidthPackageIds) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCenCenBandwidthPackageIds) GetCenBandwidthPackageId() []*string {
	return s.CenBandwidthPackageId
}

func (s *DescribeCensResponseBodyCensCenCenBandwidthPackageIds) SetCenBandwidthPackageId(v []*string) *DescribeCensResponseBodyCensCenCenBandwidthPackageIds {
	s.CenBandwidthPackageId = v
	return s
}

func (s *DescribeCensResponseBodyCensCenCenBandwidthPackageIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCensResponseBodyCensCenTags struct {
	Tag []*DescribeCensResponseBodyCensCenTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCensCenTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBodyCensCenTags) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCenTags) GetTag() []*DescribeCensResponseBodyCensCenTagsTag {
	return s.Tag
}

func (s *DescribeCensResponseBodyCensCenTags) SetTag(v []*DescribeCensResponseBodyCensCenTagsTag) *DescribeCensResponseBodyCensCenTags {
	s.Tag = v
	return s
}

func (s *DescribeCensResponseBodyCensCenTags) Validate() error {
	return dara.Validate(s)
}

type DescribeCensResponseBodyCensCenTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCensResponseBodyCensCenTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBodyCensCenTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensCenTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCensResponseBodyCensCenTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCensResponseBodyCensCenTagsTag) SetKey(v string) *DescribeCensResponseBodyCensCenTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeCensResponseBodyCensCenTagsTag) SetValue(v string) *DescribeCensResponseBodyCensCenTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeCensResponseBodyCensCenTagsTag) Validate() error {
	return dara.Validate(s)
}
