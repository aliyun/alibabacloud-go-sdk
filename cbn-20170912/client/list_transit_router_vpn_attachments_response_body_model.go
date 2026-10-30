// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVpnAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterVpnAttachmentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterVpnAttachmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterVpnAttachmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterVpnAttachmentsResponseBody
	GetTotalCount() *int32
	SetTransitRouterAttachments(v []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterVpnAttachmentsResponseBody
	GetTransitRouterAttachments() []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments
}

type ListTransitRouterVpnAttachmentsResponseBody struct {
	MaxResults               *int32                                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                *string                                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount               *int32                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TransitRouterAttachments []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments `json:"TransitRouterAttachments,omitempty" xml:"TransitRouterAttachments,omitempty" type:"Repeated"`
}

func (s ListTransitRouterVpnAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) GetTransitRouterAttachments() []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	return s.TransitRouterAttachments
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) SetMaxResults(v int32) *ListTransitRouterVpnAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) SetNextToken(v string) *ListTransitRouterVpnAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) SetRequestId(v string) *ListTransitRouterVpnAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) SetTotalCount(v int32) *ListTransitRouterVpnAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) SetTransitRouterAttachments(v []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterVpnAttachmentsResponseBody {
	s.TransitRouterAttachments = v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBody) Validate() error {
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

type ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments struct {
	AutoPublishRouteEnabled            *bool                                                                       `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	CenId                              *string                                                                     `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ChargeType                         *string                                                                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreationTime                       *string                                                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	OrderType                          *string                                                                     `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	ResourceType                       *string                                                                     `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                             *string                                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                               []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TransitRouterAttachmentDescription *string                                                                     `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	TransitRouterAttachmentId          *string                                                                     `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterAttachmentName        *string                                                                     `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	TransitRouterId                    *string                                                                     `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	VpnId                              *string                                                                     `json:"VpnId,omitempty" xml:"VpnId,omitempty"`
	VpnOwnerId                         *int64                                                                      `json:"VpnOwnerId,omitempty" xml:"VpnOwnerId,omitempty"`
	VpnRegionId                        *string                                                                     `json:"VpnRegionId,omitempty" xml:"VpnRegionId,omitempty"`
	Zones                              []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetOrderType() *string {
	return s.OrderType
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetTags() []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags {
	return s.Tags
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetVpnId() *string {
	return s.VpnId
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetVpnOwnerId() *int64 {
	return s.VpnOwnerId
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetVpnRegionId() *string {
	return s.VpnRegionId
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) GetZones() []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones {
	return s.Zones
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetAutoPublishRouteEnabled(v bool) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetCenId(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetChargeType(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.ChargeType = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetCreationTime(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetOrderType(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.OrderType = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetResourceType(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetStatus(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.Status = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetTags(v []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.Tags = v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentDescription(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentId(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentName(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterId(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetVpnId(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.VpnId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetVpnOwnerId(v int64) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.VpnOwnerId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetVpnRegionId(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.VpnRegionId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) SetZones(v []*ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments {
	s.Zones = v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachments) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags) SetKey(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Key = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags) SetValue(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Value = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsTags) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones) SetZoneId(v string) *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones {
	s.ZoneId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponseBodyTransitRouterAttachmentsZones) Validate() error {
	return dara.Validate(s)
}
