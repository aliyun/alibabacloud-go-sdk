// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetErResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetErResponseBody
	GetCode() *int32
	SetContent(v *GetErResponseBodyContent) *GetErResponseBody
	GetContent() *GetErResponseBodyContent
	SetMessage(v string) *GetErResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetErResponseBody
	GetRequestId() *string
}

type GetErResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Content *GetErResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// Information returned when the call fails
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 308DE9D2-03A6-5B44-A369-67B75D1EE091
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetErResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErResponseBody) GoString() string {
	return s.String()
}

func (s *GetErResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetErResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetErResponseBody) GetContent() *GetErResponseBodyContent {
	return s.Content
}

func (s *GetErResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetErResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErResponseBody) SetAccessDeniedDetail(v string) *GetErResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetErResponseBody) SetCode(v int32) *GetErResponseBody {
	s.Code = &v
	return s
}

func (s *GetErResponseBody) SetContent(v *GetErResponseBodyContent) *GetErResponseBody {
	s.Content = v
	return s
}

func (s *GetErResponseBody) SetMessage(v string) *GetErResponseBody {
	s.Message = &v
	return s
}

func (s *GetErResponseBody) SetRequestId(v string) *GetErResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetErResponseBodyContent struct {
	// The time when the data address was created.
	//
	// example:
	//
	// 1644283112720
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Network instance information list
	ErAttachments []*GetErResponseBodyContentErAttachments `json:"ErAttachments,omitempty" xml:"ErAttachments,omitempty" type:"Repeated"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Lingjun HUB Instance Name
	//
	// example:
	//
	// er-heyuan-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The list of route entry information.
	ErRouteEntrys []*GetErResponseBodyContentErRouteEntrys `json:"ErRouteEntrys,omitempty" xml:"ErRouteEntrys,omitempty" type:"Repeated"`
	// routing policy information list
	ErRouteMaps []*GetErResponseBodyContentErRouteMaps `json:"ErRouteMaps,omitempty" xml:"ErRouteMaps,omitempty" type:"Repeated"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1627545952000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary Zone
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzlki4ehfse4y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// List of Tags
	Tags []*GetErResponseBodyContentTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s GetErResponseBodyContent) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetErResponseBodyContent) GetDescription() *string {
	return s.Description
}

func (s *GetErResponseBodyContent) GetErAttachments() []*GetErResponseBodyContentErAttachments {
	return s.ErAttachments
}

func (s *GetErResponseBodyContent) GetErId() *string {
	return s.ErId
}

func (s *GetErResponseBodyContent) GetErName() *string {
	return s.ErName
}

func (s *GetErResponseBodyContent) GetErRouteEntrys() []*GetErResponseBodyContentErRouteEntrys {
	return s.ErRouteEntrys
}

func (s *GetErResponseBodyContent) GetErRouteMaps() []*GetErResponseBodyContentErRouteMaps {
	return s.ErRouteMaps
}

func (s *GetErResponseBodyContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetErResponseBodyContent) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *GetErResponseBodyContent) GetMessage() *string {
	return s.Message
}

func (s *GetErResponseBodyContent) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErResponseBodyContent) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetErResponseBodyContent) GetStatus() *string {
	return s.Status
}

func (s *GetErResponseBodyContent) GetTags() []*GetErResponseBodyContentTags {
	return s.Tags
}

func (s *GetErResponseBodyContent) GetTenantId() *string {
	return s.TenantId
}

func (s *GetErResponseBodyContent) SetCreateTime(v string) *GetErResponseBodyContent {
	s.CreateTime = &v
	return s
}

func (s *GetErResponseBodyContent) SetDescription(v string) *GetErResponseBodyContent {
	s.Description = &v
	return s
}

func (s *GetErResponseBodyContent) SetErAttachments(v []*GetErResponseBodyContentErAttachments) *GetErResponseBodyContent {
	s.ErAttachments = v
	return s
}

func (s *GetErResponseBodyContent) SetErId(v string) *GetErResponseBodyContent {
	s.ErId = &v
	return s
}

func (s *GetErResponseBodyContent) SetErName(v string) *GetErResponseBodyContent {
	s.ErName = &v
	return s
}

func (s *GetErResponseBodyContent) SetErRouteEntrys(v []*GetErResponseBodyContentErRouteEntrys) *GetErResponseBodyContent {
	s.ErRouteEntrys = v
	return s
}

func (s *GetErResponseBodyContent) SetErRouteMaps(v []*GetErResponseBodyContentErRouteMaps) *GetErResponseBodyContent {
	s.ErRouteMaps = v
	return s
}

func (s *GetErResponseBodyContent) SetGmtModified(v string) *GetErResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *GetErResponseBodyContent) SetMasterZoneId(v string) *GetErResponseBodyContent {
	s.MasterZoneId = &v
	return s
}

func (s *GetErResponseBodyContent) SetMessage(v string) *GetErResponseBodyContent {
	s.Message = &v
	return s
}

func (s *GetErResponseBodyContent) SetRegionId(v string) *GetErResponseBodyContent {
	s.RegionId = &v
	return s
}

func (s *GetErResponseBodyContent) SetResourceGroupId(v string) *GetErResponseBodyContent {
	s.ResourceGroupId = &v
	return s
}

func (s *GetErResponseBodyContent) SetStatus(v string) *GetErResponseBodyContent {
	s.Status = &v
	return s
}

func (s *GetErResponseBodyContent) SetTags(v []*GetErResponseBodyContentTags) *GetErResponseBodyContent {
	s.Tags = v
	return s
}

func (s *GetErResponseBodyContent) SetTenantId(v string) *GetErResponseBodyContent {
	s.TenantId = &v
	return s
}

func (s *GetErResponseBodyContent) Validate() error {
	if s.ErAttachments != nil {
		for _, item := range s.ErAttachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ErRouteEntrys != nil {
		for _, item := range s.ErRouteEntrys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ErRouteMaps != nil {
		for _, item := range s.ErRouteMaps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	return nil
}

type GetErResponseBodyContentErAttachments struct {
	// Cross-account
	//
	// example:
	//
	// false
	Across *bool `json:"Across,omitempty" xml:"Across,omitempty"`
	// Receive all routes automatically
	//
	// example:
	//
	// true
	AutoReceiveAllRoute *bool `json:"AutoReceiveAllRoute,omitempty" xml:"AutoReceiveAllRoute,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1644283112720
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The connection ID of the Lingjun HUB network instance.
	//
	// example:
	//
	// er-attachment-f32hxfsu
	ErAttachmentId *string `json:"ErAttachmentId,omitempty" xml:"ErAttachmentId,omitempty"`
	// Network Instance Name
	//
	// example:
	//
	// fudan-egpu
	ErAttachmentName *string `json:"ErAttachmentName,omitempty" xml:"ErAttachmentName,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1649303733000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// vpd-kkopgtne
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// zhijiao
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Instance type: VPD and VCC
	//
	// Valid value:
	//
	// 	- VCC: Lingjun Connection.
	//
	// 	- VPD: Lingjun network segment.
	//
	// example:
	//
	// VPD
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The synchronized region where the ECS instances are deployed.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmzzka6bnjvbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// xxxxxxxx
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErResponseBodyContentErAttachments) String() string {
	return dara.Prettify(s)
}

func (s GetErResponseBodyContentErAttachments) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContentErAttachments) GetAcross() *bool {
	return s.Across
}

func (s *GetErResponseBodyContentErAttachments) GetAutoReceiveAllRoute() *bool {
	return s.AutoReceiveAllRoute
}

func (s *GetErResponseBodyContentErAttachments) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetErResponseBodyContentErAttachments) GetErAttachmentId() *string {
	return s.ErAttachmentId
}

func (s *GetErResponseBodyContentErAttachments) GetErAttachmentName() *string {
	return s.ErAttachmentName
}

func (s *GetErResponseBodyContentErAttachments) GetErId() *string {
	return s.ErId
}

func (s *GetErResponseBodyContentErAttachments) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetErResponseBodyContentErAttachments) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetErResponseBodyContentErAttachments) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetErResponseBodyContentErAttachments) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetErResponseBodyContentErAttachments) GetMessage() *string {
	return s.Message
}

func (s *GetErResponseBodyContentErAttachments) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErResponseBodyContentErAttachments) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetErResponseBodyContentErAttachments) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *GetErResponseBodyContentErAttachments) GetStatus() *string {
	return s.Status
}

func (s *GetErResponseBodyContentErAttachments) GetTenantId() *string {
	return s.TenantId
}

func (s *GetErResponseBodyContentErAttachments) SetAcross(v bool) *GetErResponseBodyContentErAttachments {
	s.Across = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetAutoReceiveAllRoute(v bool) *GetErResponseBodyContentErAttachments {
	s.AutoReceiveAllRoute = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetCreateTime(v string) *GetErResponseBodyContentErAttachments {
	s.CreateTime = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetErAttachmentId(v string) *GetErResponseBodyContentErAttachments {
	s.ErAttachmentId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetErAttachmentName(v string) *GetErResponseBodyContentErAttachments {
	s.ErAttachmentName = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetErId(v string) *GetErResponseBodyContentErAttachments {
	s.ErId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetGmtModified(v string) *GetErResponseBodyContentErAttachments {
	s.GmtModified = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetInstanceId(v string) *GetErResponseBodyContentErAttachments {
	s.InstanceId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetInstanceName(v string) *GetErResponseBodyContentErAttachments {
	s.InstanceName = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetInstanceType(v string) *GetErResponseBodyContentErAttachments {
	s.InstanceType = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetMessage(v string) *GetErResponseBodyContentErAttachments {
	s.Message = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetRegionId(v string) *GetErResponseBodyContentErAttachments {
	s.RegionId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetResourceGroupId(v string) *GetErResponseBodyContentErAttachments {
	s.ResourceGroupId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetResourceTenantId(v string) *GetErResponseBodyContentErAttachments {
	s.ResourceTenantId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetStatus(v string) *GetErResponseBodyContentErAttachments {
	s.Status = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) SetTenantId(v string) *GetErResponseBodyContentErAttachments {
	s.TenantId = &v
	return s
}

func (s *GetErResponseBodyContentErAttachments) Validate() error {
	return dara.Validate(s)
}

type GetErResponseBodyContentErRouteEntrys struct {
	// Destination CIDR Block
	//
	// example:
	//
	// 10.0.0.0/9
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// er-rte-xnmsd2kl
	ErRouteEntryId *string `json:"ErRouteEntryId,omitempty" xml:"ErRouteEntryId,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 1623317089000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Next Hop Instance
	//
	// example:
	//
	// vcc-xxkmggkw
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// Next Hop Instance Type
	//
	// example:
	//
	// VCC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmyoj5mg3w54y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tenant to which the resource belongs.
	//
	// example:
	//
	// 1620939556166277
	ResourceTenantId *string `json:"ResourceTenantId,omitempty" xml:"ResourceTenantId,omitempty"`
	// Route type
	//
	// example:
	//
	// System
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 1620939556166277
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetErResponseBodyContentErRouteEntrys) String() string {
	return dara.Prettify(s)
}

func (s GetErResponseBodyContentErRouteEntrys) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContentErRouteEntrys) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *GetErResponseBodyContentErRouteEntrys) GetErId() *string {
	return s.ErId
}

func (s *GetErResponseBodyContentErRouteEntrys) GetErRouteEntryId() *string {
	return s.ErRouteEntryId
}

func (s *GetErResponseBodyContentErRouteEntrys) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetErResponseBodyContentErRouteEntrys) GetNextHopId() *string {
	return s.NextHopId
}

func (s *GetErResponseBodyContentErRouteEntrys) GetNextHopType() *string {
	return s.NextHopType
}

func (s *GetErResponseBodyContentErRouteEntrys) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErResponseBodyContentErRouteEntrys) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetErResponseBodyContentErRouteEntrys) GetResourceTenantId() *string {
	return s.ResourceTenantId
}

func (s *GetErResponseBodyContentErRouteEntrys) GetRouteType() *string {
	return s.RouteType
}

func (s *GetErResponseBodyContentErRouteEntrys) GetStatus() *string {
	return s.Status
}

func (s *GetErResponseBodyContentErRouteEntrys) GetTenantId() *string {
	return s.TenantId
}

func (s *GetErResponseBodyContentErRouteEntrys) SetDestinationCidrBlock(v string) *GetErResponseBodyContentErRouteEntrys {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetErId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.ErId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetErRouteEntryId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.ErRouteEntryId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetGmtModified(v string) *GetErResponseBodyContentErRouteEntrys {
	s.GmtModified = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetNextHopId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.NextHopId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetNextHopType(v string) *GetErResponseBodyContentErRouteEntrys {
	s.NextHopType = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetRegionId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.RegionId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetResourceGroupId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.ResourceGroupId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetResourceTenantId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.ResourceTenantId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetRouteType(v string) *GetErResponseBodyContentErRouteEntrys {
	s.RouteType = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetStatus(v string) *GetErResponseBodyContentErRouteEntrys {
	s.Status = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) SetTenantId(v string) *GetErResponseBodyContentErRouteEntrys {
	s.TenantId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteEntrys) Validate() error {
	return dara.Validate(s)
}

type GetErResponseBodyContentErRouteMaps struct {
	// Policy behavior
	//
	// Valid value:
	//
	// 	- deny: rejects the.
	//
	// 	- permit: The allows.
	//
	// example:
	//
	// permit
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when the data address was created.
	//
	// example:
	//
	// 1645766599809
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Policy description
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Destination CIDR Block
	//
	// example:
	//
	// 10.0.0.0/8
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// routing policy ID
	//
	// example:
	//
	// er-rmap-xkslnmsr
	ErRouteMapId *string `json:"ErRouteMapId,omitempty" xml:"ErRouteMapId,omitempty"`
	// The name of the routing policy.
	//
	// example:
	//
	// route-map-name
	ErRouteMapName *string `json:"ErRouteMapName,omitempty" xml:"ErRouteMapName,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 1623899444000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// vpd-sdkd2gkx
	ReceptionInstanceId *string `json:"ReceptionInstanceId,omitempty" xml:"ReceptionInstanceId,omitempty"`
	// The name of the destination instance.
	//
	// example:
	//
	// Reception-name
	ReceptionInstanceName *string `json:"ReceptionInstanceName,omitempty" xml:"ReceptionInstanceName,omitempty"`
	// The tenant to which the destination instance belongs.
	//
	// example:
	//
	// 1620939556166277
	ReceptionInstanceOwner *string `json:"ReceptionInstanceOwner,omitempty" xml:"ReceptionInstanceOwner,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// VPD
	ReceptionInstanceType *string `json:"ReceptionInstanceType,omitempty" xml:"ReceptionInstanceType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmzaq3ypaqkdy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Policy sequence number (1001-2000)
	//
	// example:
	//
	// 1001
	RouteMapNum *int32 `json:"RouteMapNum,omitempty" xml:"RouteMapNum,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// XXQGPROD-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// vpd-xmglsymg
	TransmissionInstanceId *string `json:"TransmissionInstanceId,omitempty" xml:"TransmissionInstanceId,omitempty"`
	// Source instance name
	//
	// example:
	//
	// test-transmission
	TransmissionInstanceName *string `json:"TransmissionInstanceName,omitempty" xml:"TransmissionInstanceName,omitempty"`
	// The tenant to which the source instance belongs.
	//
	// example:
	//
	// 1620939556166277
	TransmissionInstanceOwner *string `json:"TransmissionInstanceOwner,omitempty" xml:"TransmissionInstanceOwner,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// VPD
	TransmissionInstanceType *string `json:"TransmissionInstanceType,omitempty" xml:"TransmissionInstanceType,omitempty"`
}

func (s GetErResponseBodyContentErRouteMaps) String() string {
	return dara.Prettify(s)
}

func (s GetErResponseBodyContentErRouteMaps) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContentErRouteMaps) GetAction() *string {
	return s.Action
}

func (s *GetErResponseBodyContentErRouteMaps) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetErResponseBodyContentErRouteMaps) GetDescription() *string {
	return s.Description
}

func (s *GetErResponseBodyContentErRouteMaps) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *GetErResponseBodyContentErRouteMaps) GetErId() *string {
	return s.ErId
}

func (s *GetErResponseBodyContentErRouteMaps) GetErRouteMapId() *string {
	return s.ErRouteMapId
}

func (s *GetErResponseBodyContentErRouteMaps) GetErRouteMapName() *string {
	return s.ErRouteMapName
}

func (s *GetErResponseBodyContentErRouteMaps) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetErResponseBodyContentErRouteMaps) GetMessage() *string {
	return s.Message
}

func (s *GetErResponseBodyContentErRouteMaps) GetReceptionInstanceId() *string {
	return s.ReceptionInstanceId
}

func (s *GetErResponseBodyContentErRouteMaps) GetReceptionInstanceName() *string {
	return s.ReceptionInstanceName
}

func (s *GetErResponseBodyContentErRouteMaps) GetReceptionInstanceOwner() *string {
	return s.ReceptionInstanceOwner
}

func (s *GetErResponseBodyContentErRouteMaps) GetReceptionInstanceType() *string {
	return s.ReceptionInstanceType
}

func (s *GetErResponseBodyContentErRouteMaps) GetRegionId() *string {
	return s.RegionId
}

func (s *GetErResponseBodyContentErRouteMaps) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetErResponseBodyContentErRouteMaps) GetRouteMapNum() *int32 {
	return s.RouteMapNum
}

func (s *GetErResponseBodyContentErRouteMaps) GetStatus() *string {
	return s.Status
}

func (s *GetErResponseBodyContentErRouteMaps) GetTenantId() *string {
	return s.TenantId
}

func (s *GetErResponseBodyContentErRouteMaps) GetTransmissionInstanceId() *string {
	return s.TransmissionInstanceId
}

func (s *GetErResponseBodyContentErRouteMaps) GetTransmissionInstanceName() *string {
	return s.TransmissionInstanceName
}

func (s *GetErResponseBodyContentErRouteMaps) GetTransmissionInstanceOwner() *string {
	return s.TransmissionInstanceOwner
}

func (s *GetErResponseBodyContentErRouteMaps) GetTransmissionInstanceType() *string {
	return s.TransmissionInstanceType
}

func (s *GetErResponseBodyContentErRouteMaps) SetAction(v string) *GetErResponseBodyContentErRouteMaps {
	s.Action = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetCreateTime(v string) *GetErResponseBodyContentErRouteMaps {
	s.CreateTime = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetDescription(v string) *GetErResponseBodyContentErRouteMaps {
	s.Description = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetDestinationCidrBlock(v string) *GetErResponseBodyContentErRouteMaps {
	s.DestinationCidrBlock = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetErId(v string) *GetErResponseBodyContentErRouteMaps {
	s.ErId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetErRouteMapId(v string) *GetErResponseBodyContentErRouteMaps {
	s.ErRouteMapId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetErRouteMapName(v string) *GetErResponseBodyContentErRouteMaps {
	s.ErRouteMapName = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetGmtModified(v string) *GetErResponseBodyContentErRouteMaps {
	s.GmtModified = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetMessage(v string) *GetErResponseBodyContentErRouteMaps {
	s.Message = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetReceptionInstanceId(v string) *GetErResponseBodyContentErRouteMaps {
	s.ReceptionInstanceId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetReceptionInstanceName(v string) *GetErResponseBodyContentErRouteMaps {
	s.ReceptionInstanceName = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetReceptionInstanceOwner(v string) *GetErResponseBodyContentErRouteMaps {
	s.ReceptionInstanceOwner = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetReceptionInstanceType(v string) *GetErResponseBodyContentErRouteMaps {
	s.ReceptionInstanceType = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetRegionId(v string) *GetErResponseBodyContentErRouteMaps {
	s.RegionId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetResourceGroupId(v string) *GetErResponseBodyContentErRouteMaps {
	s.ResourceGroupId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetRouteMapNum(v int32) *GetErResponseBodyContentErRouteMaps {
	s.RouteMapNum = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetStatus(v string) *GetErResponseBodyContentErRouteMaps {
	s.Status = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTenantId(v string) *GetErResponseBodyContentErRouteMaps {
	s.TenantId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTransmissionInstanceId(v string) *GetErResponseBodyContentErRouteMaps {
	s.TransmissionInstanceId = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTransmissionInstanceName(v string) *GetErResponseBodyContentErRouteMaps {
	s.TransmissionInstanceName = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTransmissionInstanceOwner(v string) *GetErResponseBodyContentErRouteMaps {
	s.TransmissionInstanceOwner = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) SetTransmissionInstanceType(v string) *GetErResponseBodyContentErRouteMaps {
	s.TransmissionInstanceType = &v
	return s
}

func (s *GetErResponseBodyContentErRouteMaps) Validate() error {
	return dara.Validate(s)
}

type GetErResponseBodyContentTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetErResponseBodyContentTags) String() string {
	return dara.Prettify(s)
}

func (s GetErResponseBodyContentTags) GoString() string {
	return s.String()
}

func (s *GetErResponseBodyContentTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetErResponseBodyContentTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetErResponseBodyContentTags) SetTagKey(v string) *GetErResponseBodyContentTags {
	s.TagKey = &v
	return s
}

func (s *GetErResponseBodyContentTags) SetTagValue(v string) *GetErResponseBodyContentTags {
	s.TagValue = &v
	return s
}

func (s *GetErResponseBodyContentTags) Validate() error {
	return dara.Validate(s)
}
