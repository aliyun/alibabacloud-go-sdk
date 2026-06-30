// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterPeerAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterPeerAttachmentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterPeerAttachmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterPeerAttachmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterPeerAttachmentsResponseBody
	GetTotalCount() *int32
	SetTransitRouterAttachments(v []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterPeerAttachmentsResponseBody
	GetTransitRouterAttachments() []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments
}

type ListTransitRouterPeerAttachmentsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// - If this parameter is empty, no more data is returned.
	//
	// - If a value is returned for this parameter, the value is the token that is used for the next query.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AA97AFA3-8E48-4BD7-9F3E-A9F6176018A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of inter-region connections.
	TransitRouterAttachments []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments `json:"TransitRouterAttachments,omitempty" xml:"TransitRouterAttachments,omitempty" type:"Repeated"`
}

func (s ListTransitRouterPeerAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) GetTransitRouterAttachments() []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	return s.TransitRouterAttachments
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetMaxResults(v int32) *ListTransitRouterPeerAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetNextToken(v string) *ListTransitRouterPeerAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetRequestId(v string) *ListTransitRouterPeerAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetTotalCount(v int32) *ListTransitRouterPeerAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) SetTransitRouterAttachments(v []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterPeerAttachmentsResponseBody {
	s.TransitRouterAttachments = v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBody) Validate() error {
	if s.TransitRouterAttachments != nil {
		for _, item := range s.TransitRouterAttachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments struct {
	// Indicates whether the Enterprise Edition transit router automatically advertises routes to the peer region.
	//
	// - **false*	- (default): No.
	//
	// - **true**: Yes.
	//
	// example:
	//
	// false
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The bandwidth of the inter-region connection. Unit: Mbit/s.
	//
	// - If **BandwidthType*	- is set to **BandwidthPackage**, this parameter indicates the bandwidth that is allocated to the inter-region connection.
	//
	// - If **BandwidthType*	- is set to **DataTransfer**, this parameter indicates the maximum bandwidth of the inter-region connection.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth allocation method. Valid values:
	//
	// - **BandwidthPackage**: Bandwidth is allocated from a bandwidth plan.
	//
	// - **DataTransfer**: The inter-region connection is not allocated a specific bandwidth and is billed on a pay-by-traffic basis.
	//
	// example:
	//
	// BandwidthPackage
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The ID of the bandwidth plan that is associated with the inter-region connection.
	//
	// example:
	//
	// cenbwp-3xrxupouolw5ou****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The time when the inter-region connection was created.
	//
	// The time is displayed in the ISO 8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-16T02:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The default link type.
	//
	// - **Gold*	- (default): Gold.
	//
	// - **Platinum**: Platinum.
	//
	// example:
	//
	// Gold
	DefaultLinkType *string `json:"DefaultLinkType,omitempty" xml:"DefaultLinkType,omitempty"`
	// The connected areas of the bandwidth plan.
	//
	// example:
	//
	// china_china
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	// The ID of the peer transit router.
	//
	// example:
	//
	// tr-m5eq27g6bndum7e88****
	PeerTransitRouterId *string `json:"PeerTransitRouterId,omitempty" xml:"PeerTransitRouterId,omitempty"`
	// The ID of the Alibaba Cloud account to which the peer transit router belongs.
	//
	// example:
	//
	// 253460731706911258
	PeerTransitRouterOwnerId *int64 `json:"PeerTransitRouterOwnerId,omitempty" xml:"PeerTransitRouterOwnerId,omitempty"`
	// The ID of the region where the peer transit router is deployed.
	//
	// example:
	//
	// cn-qingdao
	PeerTransitRouterRegionId *string `json:"PeerTransitRouterRegionId,omitempty" xml:"PeerTransitRouterRegionId,omitempty"`
	// The ID of the region where the Enterprise Edition transit router is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of resource that is associated with the connection.
	//
	// - **VPC**: virtual private cloud (VPC).
	//
	// - **CCN**: Cloud Connect Network (CCN) instance.
	//
	// - **VBR**: virtual border router (VBR).
	//
	// - **TR**: transit router.
	//
	// example:
	//
	// TR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the inter-region connection.
	//
	// - **Attached**: The connection is attached.
	//
	// - **Attaching**: The connection is being attached.
	//
	// - **Detaching**: The connection is being detached.
	//
	// - **Detached**: The connection is detached.
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the inter-region connection.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ID of the inter-region connection.
	//
	// example:
	//
	// tr-attach-5u4qbayfv2io5v****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The name of the inter-region connection.
	//
	// example:
	//
	// test
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetDefaultLinkType() *string {
	return s.DefaultLinkType
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetGeographicSpanId() *string {
	return s.GeographicSpanId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetPeerTransitRouterId() *string {
	return s.PeerTransitRouterId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetPeerTransitRouterOwnerId() *int64 {
	return s.PeerTransitRouterOwnerId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetPeerTransitRouterRegionId() *string {
	return s.PeerTransitRouterRegionId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetTags() []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags {
	return s.Tags
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetAutoPublishRouteEnabled(v bool) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetBandwidth(v int32) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.Bandwidth = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetBandwidthType(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.BandwidthType = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetCenBandwidthPackageId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetCenId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetCreationTime(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetDefaultLinkType(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.DefaultLinkType = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetGeographicSpanId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.GeographicSpanId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetPeerTransitRouterId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.PeerTransitRouterId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetPeerTransitRouterOwnerId(v int64) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.PeerTransitRouterOwnerId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetPeerTransitRouterRegionId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.PeerTransitRouterRegionId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetRegionId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetResourceType(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetStatus(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.Status = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTags(v []*ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.Tags = v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentDescription(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentName(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterId(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachments) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags struct {
	// The tag key.
	//
	// example:
	//
	// tag_A1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags) SetKey(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Key = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags) SetValue(v string) *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Value = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponseBodyTransitRouterAttachmentsTags) Validate() error {
	return dara.Validate(s)
}
