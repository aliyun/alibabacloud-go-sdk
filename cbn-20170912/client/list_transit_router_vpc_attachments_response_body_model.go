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
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results.
	//
	// - If this parameter is empty, all results have been returned.
	//
	// - If a value is returned for **NextToken**, it is the token to start the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C97FF53F-3EF8-4883-B459-60E171924B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of VPC connections.
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
	// Specifies whether the Enterprise Edition transit router automatically advertises routes to the VPC.
	//
	// - **false**: Routes are not automatically advertised.
	//
	// - **true**: Routes are automatically advertised.
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
	// The value is always **POSTPAY**, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the VPC connection was created.
	//
	// The time is in the `YYYY-MM-DDThh:mmZ` format and in UTC.
	//
	// example:
	//
	// 2021-06-15T02:14Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The cloud service to which the resource belongs.
	//
	// example:
	//
	// SAS
	ManagedService *string `json:"ManagedService,omitempty" xml:"ManagedService,omitempty"`
	// A collection of feature attributes.
	Options *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// Specifies who pays for the network instance. Valid values:
	//
	// - **PayByCenOwner**: The account that owns the CEN instance pays the fees.
	//
	// - **PayByResourceOwner**: The account that owns the network instance pays the fees.
	//
	// example:
	//
	// PayByCenOwner
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The type of resource to which the connection is attached.
	//
	// The value is always **VPC**, which indicates a VPC.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the VPC connection.
	//
	// - **Attached**: The connection is established.
	//
	// - **Attaching**: The connection is being created.
	//
	// - **Detaching**: The connection is being deleted.
	//
	// example:
	//
	// Attached
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of tags.
	Tags []*ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the VPC connection.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ID of the VPC connection.
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
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The feature attributes of the VPC connection. This parameter is deprecated. We recommend that you use the Options parameter instead.
	TransitRouterVPCAttachmentOptions map[string]*string `json:"TransitRouterVPCAttachmentOptions,omitempty" xml:"TransitRouterVPCAttachmentOptions,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the account that owns the VPC.
	//
	// example:
	//
	// 1250123456123456
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	// The ID of the region where the VPC is deployed.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// The zone mappings of the VPC connection. This includes the vSwitches and elastic network interfaces (ENIs) in the associated VPC.
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

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) GetOptions() *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions {
	return s.Options
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

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments) SetOptions(v *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachments {
	s.Options = v
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
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
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

type ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions struct {
	// Specifies whether appliance mode is enabled.
	//
	// - **disable*	- (default): Appliance mode is disabled.
	//
	// - **enable**: Appliance mode is enabled.
	//
	// example:
	//
	// enable
	ApplianceModeSupport *string `json:"ApplianceModeSupport,omitempty" xml:"ApplianceModeSupport,omitempty"`
	// Specifies whether IPv6 is enabled.
	//
	// - **disable*	- (default): IPv6 is disabled.
	//
	// - **enable**: IPv6 is enabled.
	//
	// example:
	//
	// enable
	Ipv6Support *string `json:"Ipv6Support,omitempty" xml:"Ipv6Support,omitempty"`
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions) GetApplianceModeSupport() *string {
	return s.ApplianceModeSupport
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions) GetIpv6Support() *string {
	return s.Ipv6Support
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions) SetApplianceModeSupport(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions {
	s.ApplianceModeSupport = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions) SetIpv6Support(v string) *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions {
	s.Ipv6Support = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponseBodyTransitRouterAttachmentsOptions) Validate() error {
	return dara.Validate(s)
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
	// The ID of the ENI that the Enterprise Edition transit router creates in the vSwitch.
	//
	// example:
	//
	// eni-bp149hmyaqegerml****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1a214sbus8z3b54****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone.
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
