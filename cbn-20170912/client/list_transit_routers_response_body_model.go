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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. For more information about paging, see the related parameter descriptions.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
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
	// The list of transit router instances.
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
	if s.TransitRouters != nil {
		for _, item := range s.TransitRouters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRoutersResponseBodyTransitRouters struct {
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	//
	// example:
	//
	// 1210123456123456
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the transit router instance was created.
	//
	// The time is displayed in UTC in the `YYYY-MM-DDThh:mmZ` format.
	//
	// example:
	//
	// 2021-03-15T09:39Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The region ID of the transit router instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the transit router instance. Valid values:
	//
	// - **Creating**: being created.
	//
	// - **Active**: active.
	//
	// - **Modifying**: being modified.
	//
	// - **Deleting**: being deleted.
	//
	// - **Upgrading**: being upgraded.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the multicast feature is enabled for the transit router instance. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	SupportMulticast *bool `json:"SupportMulticast,omitempty" xml:"SupportMulticast,omitempty"`
	// The list of tags.
	Tags []*ListTransitRoutersResponseBodyTransitRoutersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The list of transit router CIDR blocks.
	TransitRouterCidrList []*ListTransitRoutersResponseBodyTransitRoutersTransitRouterCidrList `json:"TransitRouterCidrList,omitempty" xml:"TransitRouterCidrList,omitempty" type:"Repeated"`
	// The description of the transit router instance.
	//
	// example:
	//
	// testdesc
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	// The transit router instance ID.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The name of the transit router instance.
	//
	// example:
	//
	// testname
	TransitRouterName *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
	// The type of the transit router instance. Valid values:
	//
	// - **Enterprise**: Enterprise Edition transit router.
	//
	// - **Basic**: Basic Edition transit router.
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TransitRouterCidrList != nil {
		for _, item := range s.TransitRouterCidrList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	// The transit router CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the transit router CIDR block.
	//
	// example:
	//
	// CIDRdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the transit router CIDR block.
	//
	// example:
	//
	// CIDRname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the system is allowed to automatically add a route for the transit router CIDR block to the transit router route table. Valid values:
	//
	// - **true**: allowed.
	//
	//      If the value is **true**, after you create a VPN connection of the private gateway type and create a route learning relationship for the VPN connection, the system automatically adds the following route entry to the transit router route table that has a route learning relationship with the VPN connection:
	//
	//   A blackhole route whose destination CIDR block is the transit router CIDR block from which a gateway IP address is allocated to the IPsec connection.
	//
	//
	//
	//   The blackhole route is propagated only to the route tables of VBR instances under the transit router.
	//
	// - **false**: not allowed.
	//
	// example:
	//
	// true
	PublishCidrRoute *bool `json:"PublishCidrRoute,omitempty" xml:"PublishCidrRoute,omitempty"`
	// The ID of the transit router CIDR block.
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
