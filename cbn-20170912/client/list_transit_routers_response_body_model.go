// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRoutersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTransitRoutersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTransitRoutersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTransitRoutersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRoutersResponseBody
	GetTotalCount() *int32
	SetTransitRouters(v []*ListTransitRoutersResponseBodyTransitRouters) *ListTransitRoutersResponseBody
	GetTransitRouters() []*ListTransitRoutersResponseBodyTransitRouters
}

type ListTransitRoutersResponseBody struct {
	// The page number of the returned page.
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
	// The ID of the request.
	//
	// example:
	//
	// 68521297-5FA6-46CB-B4EB-658F1C68C8CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of transit routers.
	TransitRouters []*ListTransitRoutersResponseBodyTransitRouters `json:"TransitRouters,omitempty" xml:"TransitRouters,omitempty" type:"Repeated"`
}

func (s ListTransitRoutersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRoutersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTransitRoutersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTransitRoutersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRoutersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRoutersResponseBody) GetTransitRouters() []*ListTransitRoutersResponseBodyTransitRouters {
	return s.TransitRouters
}

func (s *ListTransitRoutersResponseBody) SetPageNumber(v int32) *ListTransitRoutersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTransitRoutersResponseBody) SetPageSize(v int32) *ListTransitRoutersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTransitRoutersResponseBody) SetRequestId(v string) *ListTransitRoutersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRoutersResponseBody) SetTotalCount(v int32) *ListTransitRoutersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRoutersResponseBody) SetTransitRouters(v []*ListTransitRoutersResponseBodyTransitRouters) *ListTransitRoutersResponseBody {
	s.TransitRouters = v
	return s
}

func (s *ListTransitRoutersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTransitRoutersResponseBodyTransitRouters struct {
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	//
	// example:
	//
	// 1210123456123456
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the transit router was created.
	//
	// The time follows the ISO8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-15T09:39Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the region where the transit router is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the transit router. Valid values:
	//
	// 	- **Creating**: The transit router is being created.
	//
	// 	- **Active**: The transit router is available.
	//
	// 	- **Modifying**: The transit router is being modified
	//
	// 	- **Deleting**: The transit router is being deleted.
	//
	// 	- **Upgrading**: The transit router is being upgraded.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether multicast is enabled for the transit router. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// false
	SupportMulticast *bool `json:"SupportMulticast,omitempty" xml:"SupportMulticast,omitempty"`
	// A list of tags.
	Tags []*ListTransitRoutersResponseBodyTransitRoutersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The CIDR blocks of the transit router.
	TransitRouterCidrList []*ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList `json:"TransitRouterCidrList,omitempty" xml:"TransitRouterCidrList,omitempty" type:"Repeated"`
	// The description of the transit router.
	//
	// example:
	//
	// testdesc
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The name of the transit router.
	//
	// example:
	//
	// testname
	TransitRouterName *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
	// The edition of the transit router. Valid values:
	//
	// 	- **Enterprise**: Enhance Edition
	//
	// 	- **Basic**: Basic Edition
	//
	// example:
	//
	// Enterprise
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTransitRoutersResponseBodyTransitRouters) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRoutersResponseBodyTransitRouters) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetSupportMulticast() *bool {
	return s.SupportMulticast
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetTags() []*ListTransitRoutersResponseBodyTransitRoutersTags {
	return s.Tags
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetTransitRouterCidrList() []*ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList {
	return s.TransitRouterCidrList
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetTransitRouterDescription() *string {
	return s.TransitRouterDescription
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetTransitRouterName() *string {
	return s.TransitRouterName
}

func (s *ListTransitRoutersResponseBodyTransitRouters) GetType() *string {
	return s.Type
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetAliUid(v int64) *ListTransitRoutersResponseBodyTransitRouters {
	s.AliUid = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetCenId(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.CenId = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetCreationTime(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetRegionId(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.RegionId = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetStatus(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.Status = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetSupportMulticast(v bool) *ListTransitRoutersResponseBodyTransitRouters {
	s.SupportMulticast = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetTags(v []*ListTransitRoutersResponseBodyTransitRoutersTags) *ListTransitRoutersResponseBodyTransitRouters {
	s.Tags = v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetTransitRouterCidrList(v []*ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) *ListTransitRoutersResponseBodyTransitRouters {
	s.TransitRouterCidrList = v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetTransitRouterDescription(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.TransitRouterDescription = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetTransitRouterId(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetTransitRouterName(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.TransitRouterName = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) SetType(v string) *ListTransitRoutersResponseBodyTransitRouters {
	s.Type = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRouters) Validate() error {
	return dara.Validate(s)
}

type ListTransitRoutersResponseBodyTransitRoutersTags struct {
	// The tag key.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRoutersResponseBodyTransitRoutersTags) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRoutersResponseBodyTransitRoutersTags) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTags) GetKey() *string {
	return s.Key
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTags) GetValue() *string {
	return s.Value
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTags) SetKey(v string) *ListTransitRoutersResponseBodyTransitRoutersTags {
	s.Key = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTags) SetValue(v string) *ListTransitRoutersResponseBodyTransitRoutersTags {
	s.Value = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTags) Validate() error {
	return dara.Validate(s)
}

type ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList struct {
	// The CIDR block of the transit router.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the CIDR block.
	//
	// example:
	//
	// CIDRdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the CIDR block.
	//
	// example:
	//
	// CIDRname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the system is allowed to automatically add a route to the route table of the transit router. Valid values:
	//
	// - **true**: yes
	//
	//   A value of **true*	- indicates that after you create a private VPN connection and create a route learning correlation for the private VPC connection, the system automatically adds the following route to the route table of the transit router that is in route learning correlation with the private VPN connection: A blackhole route whose destination CIDR block is the CIDR block of the transit router. The CIDR block of the transit router refers to the CIDR block from which gateway IP addresses are allocated to IPsec-VPN connections.
	//
	//
	//
	//   The blackhole route is advertised only to the route tables of virtual border routers (VBRs) that are connected to the transit router.
	//
	// - **false**: no
	//
	// example:
	//
	// true
	PublishCidrRoute *bool `json:"PublishCidrRoute,omitempty" xml:"PublishCidrRoute,omitempty"`
	// The ID of the CIDR block.
	//
	// example:
	//
	// cidr-46p5ceg21e8152****
	TransitRouterCidrId *string `json:"TransitRouterCidrId,omitempty" xml:"TransitRouterCidrId,omitempty"`
}

func (s ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) GetCidr() *string {
	return s.Cidr
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) GetDescription() *string {
	return s.Description
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) GetName() *string {
	return s.Name
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) GetPublishCidrRoute() *bool {
	return s.PublishCidrRoute
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) GetTransitRouterCidrId() *string {
	return s.TransitRouterCidrId
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) SetCidr(v string) *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList {
	s.Cidr = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) SetDescription(v string) *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList {
	s.Description = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) SetName(v string) *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList {
	s.Name = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) SetPublishCidrRoute(v bool) *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList {
	s.PublishCidrRoute = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) SetTransitRouterCidrId(v string) *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList {
	s.TransitRouterCidrId = &v
	return s
}

func (s *ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList) Validate() error {
	return dara.Validate(s)
}
