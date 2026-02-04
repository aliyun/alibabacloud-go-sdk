// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVpcAttachmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterVpcAttachmentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterVpcAttachmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterVpcAttachmentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterVpcAttachmentsResponseBody
	GetTotalCount() *int32
	SetTransitRouterAttachments(v []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterVpcAttachmentsResponseBody
	GetTransitRouterAttachments() []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments
}

type ListTransitRouterVpcAttachmentsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. Valid values:
	//
	// 	- If **NextToken*	- is returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// C97FF53F-3EF8-4883-B459-60E171924B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the VPC connection.
	TransitRouterAttachments []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments `json:"TransitRouterAttachments,omitempty" xml:"TransitRouterAttachments,omitempty" type:"Repeated"`
}

func (s ListTransitRouterVpcAttachmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) GetTransitRouterAttachments() []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	return s.TransitRouterAttachments
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetMaxResults(v int32) *ListTransitRouterVpcAttachmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetNextToken(v string) *ListTransitRouterVpcAttachmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetRequestId(v string) *ListTransitRouterVpcAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetTotalCount(v int32) *ListTransitRouterVpcAttachmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) SetTransitRouterAttachments(v []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) *ListTransitRouterVpcAttachmentsResponseBody {
	s.TransitRouterAttachments = v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBody) Validate() error {
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

type ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments struct {
	// Indicates whether the Enterprise Edition transit router can automatically advertise routes to the VPC. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The billing method of the VPC connection.
	//
	// Only **POSTPAY*	- may be returned, which indicates the default pay-as-you-go billing method.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the VPC connection was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-15T02:14Z
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ManagedService *string `json:"ManagedService,omitempty" xml:"ManagedService,omitempty"`
	// The entity that pays the fees of the network instance. Valid values:
	//
	// 	- **PayByCenOwner**: the Alibaba Cloud account that owns the CEN instance.
	//
	// 	- **PayByResourceOwner**: the Alibaba Cloud account that owns the network instance.
	//
	// example:
	//
	// PayByCenOwner
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The type of resource to which the transit router is connected.
	//
	// Only **VPC*	- may be returned, which indicates VPCs.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the VPC connection. Valid values:
	//
	// 	- **Attached**
	//
	// 	- **Attaching**
	//
	// 	- **Detaching**
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the VPC connection.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The VPC connection ID.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The name of the VPC connection.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The description of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The features of the VPC connection.
	TransitRouterVPCAttachmentOptions map[string]*string `json:"TransitRouterVPCAttachmentOptions,omitempty" xml:"TransitRouterVPCAttachmentOptions,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 1250123456123456
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// The primary and secondary zones, vSwitches, and ENIs of the VPC.
	ZoneMappings []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetManagedService() *string {
	return s.ManagedService
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetOrderType() *string {
	return s.OrderType
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetTags() []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags {
	return s.Tags
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetTransitRouterVPCAttachmentOptions() map[string]*string {
	return s.TransitRouterVPCAttachmentOptions
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetVpcId() *string {
	return s.VpcId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetZoneMappings() []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings {
	return s.ZoneMappings
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetAutoPublishRouteEnabled(v bool) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetCenId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetChargeType(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.ChargeType = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetCreationTime(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.CreationTime = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetManagedService(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.ManagedService = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetOrderType(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.OrderType = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetResourceType(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetStatus(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.Status = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTags(v []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.Tags = v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentDescription(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterAttachmentName(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetTransitRouterVPCAttachmentOptions(v map[string]*string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.TransitRouterVPCAttachmentOptions = v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetVpcId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.VpcId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetVpcOwnerId(v int64) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.VpcOwnerId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetVpcRegionId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.VpcRegionId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetZoneMappings(v []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.ZoneMappings = v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ZoneMappings != nil {
		for _, item := range s.ZoneMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags) SetKey(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Key = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags) SetValue(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags {
	s.Value = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings struct {
	// The ID of the ENI created by the Enterprise Edition transit router in the vSwitch.
	//
	// example:
	//
	// eni-bp149hmyaqegerml****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1a214sbus8z3b54****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) SetNetworkInterfaceId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) SetVSwitchId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) SetZoneId(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsZoneMappings) Validate() error {
	return dara.Validate(s)
}
