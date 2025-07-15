// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCens(v []*DescribeCensResponseBodyCens) *DescribeCensResponseBody
	GetCens() []*DescribeCensResponseBodyCens
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
	// Details of the CEN instances.
	Cens []*DescribeCensResponseBodyCens `json:"Cens,omitempty" xml:"Cens,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of CEN instances returned.
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

func (s *DescribeCensResponseBody) GetCens() []*DescribeCensResponseBodyCens {
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

func (s *DescribeCensResponseBody) SetCens(v []*DescribeCensResponseBodyCens) *DescribeCensResponseBody {
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
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-3gwy16dojz1m65****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the CEN instance was created.
	//
	// example:
	//
	// 2021-06-16T08:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the CEN instance.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IPv6 level.
	//
	// >  IPv6 is not supported.
	//
	// Valid value:
	//
	// 	- DISABLED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DISABLED
	Ipv6Level *string `json:"Ipv6Level,omitempty" xml:"Ipv6Level,omitempty"`
	// The name of the CEN instance.
	//
	// example:
	//
	// testCen
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The bandwidth plans that are bound to the CEN instance.
	PackageIds []*DescribeCensResponseBodyCensPackageIds `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// The tolerated level of CIDR block conflict.
	//
	// Valid value:
	//
	// 	- REDUCED: CIDR block conflicts are allowed, but the conflicting CIDR blocks cannot be identical.
	//
	// example:
	//
	// REDUCED
	ProtectionLevel *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The status of the CEN instance.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Active
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the CEN instance.
	Tags []*DescribeCensResponseBodyCensTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCens) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBodyCens) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCens) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCensResponseBodyCens) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeCensResponseBodyCens) GetDescription() *string {
	return s.Description
}

func (s *DescribeCensResponseBodyCens) GetIpv6Level() *string {
	return s.Ipv6Level
}

func (s *DescribeCensResponseBodyCens) GetName() *string {
	return s.Name
}

func (s *DescribeCensResponseBodyCens) GetPackageIds() []*DescribeCensResponseBodyCensPackageIds {
	return s.PackageIds
}

func (s *DescribeCensResponseBodyCens) GetProtectionLevel() *string {
	return s.ProtectionLevel
}

func (s *DescribeCensResponseBodyCens) GetStatus() *string {
	return s.Status
}

func (s *DescribeCensResponseBodyCens) GetTags() []*DescribeCensResponseBodyCensTags {
	return s.Tags
}

func (s *DescribeCensResponseBodyCens) SetCenId(v string) *DescribeCensResponseBodyCens {
	s.CenId = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetCreationTime(v string) *DescribeCensResponseBodyCens {
	s.CreationTime = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetDescription(v string) *DescribeCensResponseBodyCens {
	s.Description = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetIpv6Level(v string) *DescribeCensResponseBodyCens {
	s.Ipv6Level = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetName(v string) *DescribeCensResponseBodyCens {
	s.Name = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetPackageIds(v []*DescribeCensResponseBodyCensPackageIds) *DescribeCensResponseBodyCens {
	s.PackageIds = v
	return s
}

func (s *DescribeCensResponseBodyCens) SetProtectionLevel(v string) *DescribeCensResponseBodyCens {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetStatus(v string) *DescribeCensResponseBodyCens {
	s.Status = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetTags(v []*DescribeCensResponseBodyCensTags) *DescribeCensResponseBodyCens {
	s.Tags = v
	return s
}

func (s *DescribeCensResponseBodyCens) Validate() error {
	return dara.Validate(s)
}

type DescribeCensResponseBodyCensPackageIds struct {
	// The ID of the bandwidth plan that is bound to the CEN instance.
	//
	// example:
	//
	// cenbwp-4c2zaavbvh5f42****
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
}

func (s DescribeCensResponseBodyCensPackageIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBodyCensPackageIds) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensPackageIds) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeCensResponseBodyCensPackageIds) SetPackageId(v string) *DescribeCensResponseBodyCensPackageIds {
	s.PackageId = &v
	return s
}

func (s *DescribeCensResponseBodyCensPackageIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCensResponseBodyCensTags struct {
	// The key of the tag.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCensResponseBodyCensTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBodyCensTags) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensTags) GetKey() *string {
	return s.Key
}

func (s *DescribeCensResponseBodyCensTags) GetValue() *string {
	return s.Value
}

func (s *DescribeCensResponseBodyCensTags) SetKey(v string) *DescribeCensResponseBodyCensTags {
	s.Key = &v
	return s
}

func (s *DescribeCensResponseBodyCensTags) SetValue(v string) *DescribeCensResponseBodyCensTags {
	s.Value = &v
	return s
}

func (s *DescribeCensResponseBodyCensTags) Validate() error {
	return dara.Validate(s)
}
