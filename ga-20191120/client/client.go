// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddEntriesToAclRequest struct {
	RegionId    *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AclId       *string                             `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclEntries  []*AddEntriesToAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	ClientToken *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool                               `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s AddEntriesToAclRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclRequest) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequest) SetRegionId(v string) *AddEntriesToAclRequest {
	s.RegionId = &v
	return s
}

func (s *AddEntriesToAclRequest) SetAclId(v string) *AddEntriesToAclRequest {
	s.AclId = &v
	return s
}

func (s *AddEntriesToAclRequest) SetAclEntries(v []*AddEntriesToAclRequestAclEntries) *AddEntriesToAclRequest {
	s.AclEntries = v
	return s
}

func (s *AddEntriesToAclRequest) SetClientToken(v string) *AddEntriesToAclRequest {
	s.ClientToken = &v
	return s
}

func (s *AddEntriesToAclRequest) SetDryRun(v bool) *AddEntriesToAclRequest {
	s.DryRun = &v
	return s
}

type AddEntriesToAclRequestAclEntries struct {
	Entry            *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	EntryDescription *string `json:"EntryDescription,omitempty" xml:"EntryDescription,omitempty"`
}

func (s AddEntriesToAclRequestAclEntries) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequestAclEntries) SetEntry(v string) *AddEntriesToAclRequestAclEntries {
	s.Entry = &v
	return s
}

func (s *AddEntriesToAclRequestAclEntries) SetEntryDescription(v string) *AddEntriesToAclRequestAclEntries {
	s.EntryDescription = &v
	return s
}

type AddEntriesToAclResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
}

func (s AddEntriesToAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclResponseBody) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponseBody) SetRequestId(v string) *AddEntriesToAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEntriesToAclResponseBody) SetAclId(v string) *AddEntriesToAclResponseBody {
	s.AclId = &v
	return s
}

type AddEntriesToAclResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEntriesToAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEntriesToAclResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclResponse) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponse) SetHeaders(v map[string]*string) *AddEntriesToAclResponse {
	s.Headers = v
	return s
}

func (s *AddEntriesToAclResponse) SetBody(v *AddEntriesToAclResponseBody) *AddEntriesToAclResponse {
	s.Body = v
	return s
}

type AssociateAclsWithListenerRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AclIds      []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	ListenerId  *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	AclType     *string   `json:"AclType,omitempty" xml:"AclType,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s AssociateAclsWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerRequest) SetRegionId(v string) *AssociateAclsWithListenerRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetAclIds(v []*string) *AssociateAclsWithListenerRequest {
	s.AclIds = v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetListenerId(v string) *AssociateAclsWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetAclType(v string) *AssociateAclsWithListenerRequest {
	s.AclType = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetClientToken(v string) *AssociateAclsWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetDryRun(v bool) *AssociateAclsWithListenerRequest {
	s.DryRun = &v
	return s
}

type AssociateAclsWithListenerResponseBody struct {
	// Id of the request
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclIds     []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	ListenerId *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s AssociateAclsWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponseBody) SetRequestId(v string) *AssociateAclsWithListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetAclIds(v []*string) *AssociateAclsWithListenerResponseBody {
	s.AclIds = v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetListenerId(v string) *AssociateAclsWithListenerResponseBody {
	s.ListenerId = &v
	return s
}

type AssociateAclsWithListenerResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateAclsWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateAclsWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerResponse) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponse) SetHeaders(v map[string]*string) *AssociateAclsWithListenerResponse {
	s.Headers = v
	return s
}

func (s *AssociateAclsWithListenerResponse) SetBody(v *AssociateAclsWithListenerResponseBody) *AssociateAclsWithListenerResponse {
	s.Body = v
	return s
}

type AttachDdosToAcceleratorRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	DdosId        *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	DdosRegionId  *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachDdosToAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachDdosToAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *AttachDdosToAcceleratorRequest) SetAcceleratorId(v string) *AttachDdosToAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequest) SetDdosId(v string) *AttachDdosToAcceleratorRequest {
	s.DdosId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequest) SetDdosRegionId(v string) *AttachDdosToAcceleratorRequest {
	s.DdosRegionId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequest) SetRegionId(v string) *AttachDdosToAcceleratorRequest {
	s.RegionId = &v
	return s
}

type AttachDdosToAcceleratorResponseBody struct {
	DdosId    *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GaId      *string `json:"GaId,omitempty" xml:"GaId,omitempty"`
}

func (s AttachDdosToAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachDdosToAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDdosToAcceleratorResponseBody) SetDdosId(v string) *AttachDdosToAcceleratorResponseBody {
	s.DdosId = &v
	return s
}

func (s *AttachDdosToAcceleratorResponseBody) SetRequestId(v string) *AttachDdosToAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDdosToAcceleratorResponseBody) SetGaId(v string) *AttachDdosToAcceleratorResponseBody {
	s.GaId = &v
	return s
}

type AttachDdosToAcceleratorResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachDdosToAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachDdosToAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachDdosToAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *AttachDdosToAcceleratorResponse) SetHeaders(v map[string]*string) *AttachDdosToAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *AttachDdosToAcceleratorResponse) SetBody(v *AttachDdosToAcceleratorResponseBody) *AttachDdosToAcceleratorResponse {
	s.Body = v
	return s
}

type AttachLogStoreToEndpointGroupRequest struct {
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlsProjectName   *string   `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	SlsLogStoreName  *string   `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	AcceleratorId    *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId       *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	SlsRegionId      *string   `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	ClientToken      *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AttachLogStoreToEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachLogStoreToEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachLogStoreToEndpointGroupRequest) SetRegionId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsProjectName(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsProjectName = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsLogStoreName(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsLogStoreName = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetAcceleratorId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetListenerId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsRegionId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsRegionId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetEndpointGroupIds(v []*string) *AttachLogStoreToEndpointGroupRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetClientToken(v string) *AttachLogStoreToEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

type AttachLogStoreToEndpointGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachLogStoreToEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachLogStoreToEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLogStoreToEndpointGroupResponseBody) SetRequestId(v string) *AttachLogStoreToEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

type AttachLogStoreToEndpointGroupResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachLogStoreToEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachLogStoreToEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachLogStoreToEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachLogStoreToEndpointGroupResponse) SetHeaders(v map[string]*string) *AttachLogStoreToEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachLogStoreToEndpointGroupResponse) SetBody(v *AttachLogStoreToEndpointGroupResponseBody) *AttachLogStoreToEndpointGroupResponse {
	s.Body = v
	return s
}

type BandwidthPackageAddAcceleratorRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	AcceleratorId      *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s BandwidthPackageAddAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageAddAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *BandwidthPackageAddAcceleratorRequest) SetRegionId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorRequest) SetBandwidthPackageId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorRequest) SetAcceleratorId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

type BandwidthPackageAddAcceleratorResponseBody struct {
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Accelerators       []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	BandwidthPackageId *string   `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
}

func (s BandwidthPackageAddAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageAddAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetRequestId(v string) *BandwidthPackageAddAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetAccelerators(v []*string) *BandwidthPackageAddAcceleratorResponseBody {
	s.Accelerators = v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetBandwidthPackageId(v string) *BandwidthPackageAddAcceleratorResponseBody {
	s.BandwidthPackageId = &v
	return s
}

type BandwidthPackageAddAcceleratorResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BandwidthPackageAddAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BandwidthPackageAddAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageAddAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *BandwidthPackageAddAcceleratorResponse) SetHeaders(v map[string]*string) *BandwidthPackageAddAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponse) SetBody(v *BandwidthPackageAddAcceleratorResponseBody) *BandwidthPackageAddAcceleratorResponse {
	s.Body = v
	return s
}

type BandwidthPackageRemoveAcceleratorRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	AcceleratorId      *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s BandwidthPackageRemoveAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageRemoveAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetRegionId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetBandwidthPackageId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetAcceleratorId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

type BandwidthPackageRemoveAcceleratorResponseBody struct {
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Accelerators       []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	BandwidthPackageId *string   `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
}

func (s BandwidthPackageRemoveAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageRemoveAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetRequestId(v string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetAccelerators(v []*string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.Accelerators = v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetBandwidthPackageId(v string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.BandwidthPackageId = &v
	return s
}

type BandwidthPackageRemoveAcceleratorResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BandwidthPackageRemoveAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BandwidthPackageRemoveAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageRemoveAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *BandwidthPackageRemoveAcceleratorResponse) SetHeaders(v map[string]*string) *BandwidthPackageRemoveAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponse) SetBody(v *BandwidthPackageRemoveAcceleratorResponseBody) *BandwidthPackageRemoveAcceleratorResponse {
	s.Body = v
	return s
}

type ConfigEndpointProbeRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointType    *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ProbeProtocol   *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	ProbePort       *string `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
	Enable          *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ConfigEndpointProbeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigEndpointProbeRequest) GoString() string {
	return s.String()
}

func (s *ConfigEndpointProbeRequest) SetRegionId(v string) *ConfigEndpointProbeRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetClientToken(v string) *ConfigEndpointProbeRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEndpointGroupId(v string) *ConfigEndpointProbeRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEndpointType(v string) *ConfigEndpointProbeRequest {
	s.EndpointType = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEndpoint(v string) *ConfigEndpointProbeRequest {
	s.Endpoint = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetProbeProtocol(v string) *ConfigEndpointProbeRequest {
	s.ProbeProtocol = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetProbePort(v string) *ConfigEndpointProbeRequest {
	s.ProbePort = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEnable(v string) *ConfigEndpointProbeRequest {
	s.Enable = &v
	return s
}

type ConfigEndpointProbeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigEndpointProbeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigEndpointProbeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigEndpointProbeResponseBody) SetRequestId(v string) *ConfigEndpointProbeResponseBody {
	s.RequestId = &v
	return s
}

type ConfigEndpointProbeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigEndpointProbeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigEndpointProbeResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigEndpointProbeResponse) GoString() string {
	return s.String()
}

func (s *ConfigEndpointProbeResponse) SetHeaders(v map[string]*string) *ConfigEndpointProbeResponse {
	s.Headers = v
	return s
}

func (s *ConfigEndpointProbeResponse) SetBody(v *ConfigEndpointProbeResponseBody) *ConfigEndpointProbeResponse {
	s.Body = v
	return s
}

type CreateAcceleratorRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Spec          *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	AutoPay       *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
}

func (s CreateAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorRequest) SetRegionId(v string) *CreateAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAcceleratorRequest) SetClientToken(v string) *CreateAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAcceleratorRequest) SetName(v string) *CreateAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *CreateAcceleratorRequest) SetDuration(v int32) *CreateAcceleratorRequest {
	s.Duration = &v
	return s
}

func (s *CreateAcceleratorRequest) SetPricingCycle(v string) *CreateAcceleratorRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateAcceleratorRequest) SetSpec(v string) *CreateAcceleratorRequest {
	s.Spec = &v
	return s
}

func (s *CreateAcceleratorRequest) SetAutoPay(v bool) *CreateAcceleratorRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAcceleratorRequest) SetAutoUseCoupon(v string) *CreateAcceleratorRequest {
	s.AutoUseCoupon = &v
	return s
}

type CreateAcceleratorResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s CreateAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorResponseBody) SetRequestId(v string) *CreateAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAcceleratorResponseBody) SetOrderId(v string) *CreateAcceleratorResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAcceleratorResponseBody) SetAcceleratorId(v string) *CreateAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

type CreateAcceleratorResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorResponse) SetHeaders(v map[string]*string) *CreateAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *CreateAcceleratorResponse) SetBody(v *CreateAcceleratorResponseBody) *CreateAcceleratorResponse {
	s.Body = v
	return s
}

type CreateAclRequest struct {
	RegionId         *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AclName          *string                       `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AddressIPVersion *string                       `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AclEntries       []*CreateAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	ClientToken      *string                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool                         `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s CreateAclRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRequest) SetRegionId(v string) *CreateAclRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAclRequest) SetAclName(v string) *CreateAclRequest {
	s.AclName = &v
	return s
}

func (s *CreateAclRequest) SetAddressIPVersion(v string) *CreateAclRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateAclRequest) SetAclEntries(v []*CreateAclRequestAclEntries) *CreateAclRequest {
	s.AclEntries = v
	return s
}

func (s *CreateAclRequest) SetClientToken(v string) *CreateAclRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAclRequest) SetDryRun(v bool) *CreateAclRequest {
	s.DryRun = &v
	return s
}

type CreateAclRequestAclEntries struct {
	Entry            *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	EntryDescription *string `json:"EntryDescription,omitempty" xml:"EntryDescription,omitempty"`
}

func (s CreateAclRequestAclEntries) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *CreateAclRequestAclEntries) SetEntry(v string) *CreateAclRequestAclEntries {
	s.Entry = &v
	return s
}

func (s *CreateAclRequestAclEntries) SetEntryDescription(v string) *CreateAclRequestAclEntries {
	s.EntryDescription = &v
	return s
}

type CreateAclResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
}

func (s CreateAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclResponseBody) SetRequestId(v string) *CreateAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAclResponseBody) SetAclId(v string) *CreateAclResponseBody {
	s.AclId = &v
	return s
}

type CreateAclResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAclResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponse) GoString() string {
	return s.String()
}

func (s *CreateAclResponse) SetHeaders(v map[string]*string) *CreateAclResponse {
	s.Headers = v
	return s
}

func (s *CreateAclResponse) SetBody(v *CreateAclResponseBody) *CreateAclResponse {
	s.Body = v
	return s
}

type CreateBandwidthPackageRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Bandwidth              *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Duration               *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle           *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	AutoPay                *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Type                   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	BandwidthType          *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	AutoUseCoupon          *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	Ratio                  *int32  `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	BillingType            *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	ChargeType             *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CbnGeographicRegionIdA *string `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	CbnGeographicRegionIdB *string `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
}

func (s CreateBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageRequest) SetRegionId(v string) *CreateBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBandwidth(v int32) *CreateBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetDuration(v string) *CreateBandwidthPackageRequest {
	s.Duration = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetPricingCycle(v string) *CreateBandwidthPackageRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetAutoPay(v bool) *CreateBandwidthPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetClientToken(v string) *CreateBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetType(v string) *CreateBandwidthPackageRequest {
	s.Type = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBandwidthType(v string) *CreateBandwidthPackageRequest {
	s.BandwidthType = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetAutoUseCoupon(v string) *CreateBandwidthPackageRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetRatio(v int32) *CreateBandwidthPackageRequest {
	s.Ratio = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBillingType(v string) *CreateBandwidthPackageRequest {
	s.BillingType = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetChargeType(v string) *CreateBandwidthPackageRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetCbnGeographicRegionIdA(v string) *CreateBandwidthPackageRequest {
	s.CbnGeographicRegionIdA = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetCbnGeographicRegionIdB(v string) *CreateBandwidthPackageRequest {
	s.CbnGeographicRegionIdB = &v
	return s
}

type CreateBandwidthPackageResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageResponseBody) SetRequestId(v string) *CreateBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *CreateBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateBandwidthPackageResponseBody) SetOrderId(v string) *CreateBandwidthPackageResponseBody {
	s.OrderId = &v
	return s
}

type CreateBandwidthPackageResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageResponse) SetHeaders(v map[string]*string) *CreateBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateBandwidthPackageResponse) SetBody(v *CreateBandwidthPackageResponseBody) *CreateBandwidthPackageResponse {
	s.Body = v
	return s
}

type CreateEndpointGroupRequest struct {
	RegionId                   *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken                *string                                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AcceleratorId              *string                                             `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Name                       *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Description                *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointGroupRegion        *string                                             `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	ListenerId                 *string                                             `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	TrafficPercentage          *int32                                              `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
	HealthCheckIntervalSeconds *int32                                              `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath            *string                                             `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort            *int32                                              `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol        *string                                             `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	ThresholdCount             *int32                                              `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	EndpointConfigurations     []*CreateEndpointGroupRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointRequestProtocol    *string                                             `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	EndpointGroupType          *string                                             `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	PortOverrides              []*CreateEndpointGroupRequestPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
}

func (s CreateEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequest) SetRegionId(v string) *CreateEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetClientToken(v string) *CreateEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetAcceleratorId(v string) *CreateEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetName(v string) *CreateEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetDescription(v string) *CreateEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointGroupRegion(v string) *CreateEndpointGroupRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetListenerId(v string) *CreateEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetTrafficPercentage(v int32) *CreateEndpointGroupRequest {
	s.TrafficPercentage = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckIntervalSeconds(v int32) *CreateEndpointGroupRequest {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckPath(v string) *CreateEndpointGroupRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckPort(v int32) *CreateEndpointGroupRequest {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckProtocol(v string) *CreateEndpointGroupRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetThresholdCount(v int32) *CreateEndpointGroupRequest {
	s.ThresholdCount = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointConfigurations(v []*CreateEndpointGroupRequestEndpointConfigurations) *CreateEndpointGroupRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointRequestProtocol(v string) *CreateEndpointGroupRequest {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointGroupType(v string) *CreateEndpointGroupRequest {
	s.EndpointGroupType = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetPortOverrides(v []*CreateEndpointGroupRequestPortOverrides) *CreateEndpointGroupRequest {
	s.PortOverrides = v
	return s
}

type CreateEndpointGroupRequestEndpointConfigurations struct {
	Type                       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EnableClientIPPreservation *bool   `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	Weight                     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Endpoint                   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
}

func (s CreateEndpointGroupRequestEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetType(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetEnableClientIPPreservation(v bool) *CreateEndpointGroupRequestEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetWeight(v int32) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetEndpoint(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Endpoint = &v
	return s
}

type CreateEndpointGroupRequestPortOverrides struct {
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
}

func (s CreateEndpointGroupRequestPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupRequestPortOverrides) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequestPortOverrides) SetListenerPort(v int32) *CreateEndpointGroupRequestPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *CreateEndpointGroupRequestPortOverrides) SetEndpointPort(v int32) *CreateEndpointGroupRequestPortOverrides {
	s.EndpointPort = &v
	return s
}

type CreateEndpointGroupResponseBody struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupResponseBody) SetEndpointGroupId(v string) *CreateEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateEndpointGroupResponseBody) SetRequestId(v string) *CreateEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateEndpointGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupResponse) SetHeaders(v map[string]*string) *CreateEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateEndpointGroupResponse) SetBody(v *CreateEndpointGroupResponseBody) *CreateEndpointGroupResponse {
	s.Body = v
	return s
}

type CreateForwardingRulesRequest struct {
	RegionId        *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken     *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AcceleratorId   *string                                        `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId      *string                                        `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	ForwardingRules []*CreateForwardingRulesRequestForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequest) SetRegionId(v string) *CreateForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetClientToken(v string) *CreateForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetAcceleratorId(v string) *CreateForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetListenerId(v string) *CreateForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetForwardingRules(v []*CreateForwardingRulesRequestForwardingRules) *CreateForwardingRulesRequest {
	s.ForwardingRules = v
	return s
}

type CreateForwardingRulesRequestForwardingRules struct {
	Priority           *int32                                                       `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleConditions     []*CreateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	RuleActions        []*CreateForwardingRulesRequestForwardingRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	ForwardingRuleName *string                                                      `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRules) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRules) SetPriority(v int32) *CreateForwardingRulesRequestForwardingRules {
	s.Priority = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleConditions(v []*CreateForwardingRulesRequestForwardingRulesRuleConditions) *CreateForwardingRulesRequestForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleActions(v []*CreateForwardingRulesRequestForwardingRulesRuleActions) *CreateForwardingRulesRequestForwardingRules {
	s.RuleActions = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetForwardingRuleName(v string) *CreateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

type CreateForwardingRulesRequestForwardingRulesRuleConditions struct {
	RuleConditionType *string                                                              `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	PathConfig        *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	HostConfig        *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionType(v string) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetPathConfig(v *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetHostConfig(v *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

type CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) SetValues(v []*string) *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

type CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) SetValues(v []*string) *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

type CreateForwardingRulesRequestForwardingRulesRuleActions struct {
	Order              *int32                                                                    `json:"Order,omitempty" xml:"Order,omitempty"`
	RuleActionType     *string                                                                   `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	ForwardGroupConfig *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetOrder(v int32) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionType(v string) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetForwardGroupConfig(v *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

type CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig struct {
	ServerGroupTuples []*CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) SetEndpointGroupId(v string) *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.EndpointGroupId = &v
	return s
}

type CreateForwardingRulesResponseBody struct {
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ForwardingRules []*CreateForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesResponseBody) SetRequestId(v string) *CreateForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateForwardingRulesResponseBody) SetForwardingRules(v []*CreateForwardingRulesResponseBodyForwardingRules) *CreateForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

type CreateForwardingRulesResponseBodyForwardingRules struct {
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
}

func (s CreateForwardingRulesResponseBodyForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesResponseBodyForwardingRules) SetForwardingRuleId(v string) *CreateForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

type CreateForwardingRulesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateForwardingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesResponse) SetHeaders(v map[string]*string) *CreateForwardingRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateForwardingRulesResponse) SetBody(v *CreateForwardingRulesResponseBody) *CreateForwardingRulesResponse {
	s.Body = v
	return s
}

type CreateIpSetsRequest struct {
	RegionId         *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken      *string                                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AcceleratorId    *string                                `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccelerateRegion []*CreateIpSetsRequestAccelerateRegion `json:"AccelerateRegion,omitempty" xml:"AccelerateRegion,omitempty" type:"Repeated"`
}

func (s CreateIpSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpSetsRequest) GoString() string {
	return s.String()
}

func (s *CreateIpSetsRequest) SetRegionId(v string) *CreateIpSetsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpSetsRequest) SetClientToken(v string) *CreateIpSetsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpSetsRequest) SetAcceleratorId(v string) *CreateIpSetsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateIpSetsRequest) SetAccelerateRegion(v []*CreateIpSetsRequestAccelerateRegion) *CreateIpSetsRequest {
	s.AccelerateRegion = v
	return s
}

type CreateIpSetsRequestAccelerateRegion struct {
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	IpVersion          *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s CreateIpSetsRequestAccelerateRegion) String() string {
	return tea.Prettify(s)
}

func (s CreateIpSetsRequestAccelerateRegion) GoString() string {
	return s.String()
}

func (s *CreateIpSetsRequestAccelerateRegion) SetAccelerateRegionId(v string) *CreateIpSetsRequestAccelerateRegion {
	s.AccelerateRegionId = &v
	return s
}

func (s *CreateIpSetsRequestAccelerateRegion) SetIpVersion(v string) *CreateIpSetsRequestAccelerateRegion {
	s.IpVersion = &v
	return s
}

func (s *CreateIpSetsRequestAccelerateRegion) SetBandwidth(v int32) *CreateIpSetsRequestAccelerateRegion {
	s.Bandwidth = &v
	return s
}

type CreateIpSetsResponseBody struct {
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpSets        []*CreateIpSetsResponseBodyIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
	AcceleratorId *string                           `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s CreateIpSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpSetsResponseBody) SetRequestId(v string) *CreateIpSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpSetsResponseBody) SetIpSets(v []*CreateIpSetsResponseBodyIpSets) *CreateIpSetsResponseBody {
	s.IpSets = v
	return s
}

func (s *CreateIpSetsResponseBody) SetAcceleratorId(v string) *CreateIpSetsResponseBody {
	s.AcceleratorId = &v
	return s
}

type CreateIpSetsResponseBodyIpSets struct {
	AccelerateRegionId *string   `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	Bandwidth          *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	IpSetId            *string   `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	IpList             []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
}

func (s CreateIpSetsResponseBodyIpSets) String() string {
	return tea.Prettify(s)
}

func (s CreateIpSetsResponseBodyIpSets) GoString() string {
	return s.String()
}

func (s *CreateIpSetsResponseBodyIpSets) SetAccelerateRegionId(v string) *CreateIpSetsResponseBodyIpSets {
	s.AccelerateRegionId = &v
	return s
}

func (s *CreateIpSetsResponseBodyIpSets) SetBandwidth(v int32) *CreateIpSetsResponseBodyIpSets {
	s.Bandwidth = &v
	return s
}

func (s *CreateIpSetsResponseBodyIpSets) SetIpSetId(v string) *CreateIpSetsResponseBodyIpSets {
	s.IpSetId = &v
	return s
}

func (s *CreateIpSetsResponseBodyIpSets) SetIpList(v []*string) *CreateIpSetsResponseBodyIpSets {
	s.IpList = v
	return s
}

type CreateIpSetsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIpSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpSetsResponse) GoString() string {
	return s.String()
}

func (s *CreateIpSetsResponse) SetHeaders(v map[string]*string) *CreateIpSetsResponse {
	s.Headers = v
	return s
}

func (s *CreateIpSetsResponse) SetBody(v *CreateIpSetsResponseBody) *CreateIpSetsResponse {
	s.Body = v
	return s
}

type CreateListenerRequest struct {
	RegionId       *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken    *string                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AcceleratorId  *string                              `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Name           *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientAffinity *string                              `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	Protocol       *string                              `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ProxyProtocol  *bool                                `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	PortRanges     []*CreateListenerRequestPortRanges   `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Certificates   []*CreateListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
}

func (s CreateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) SetRegionId(v string) *CreateListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetAcceleratorId(v string) *CreateListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateListenerRequest) SetName(v string) *CreateListenerRequest {
	s.Name = &v
	return s
}

func (s *CreateListenerRequest) SetDescription(v string) *CreateListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateListenerRequest) SetClientAffinity(v string) *CreateListenerRequest {
	s.ClientAffinity = &v
	return s
}

func (s *CreateListenerRequest) SetProtocol(v string) *CreateListenerRequest {
	s.Protocol = &v
	return s
}

func (s *CreateListenerRequest) SetProxyProtocol(v bool) *CreateListenerRequest {
	s.ProxyProtocol = &v
	return s
}

func (s *CreateListenerRequest) SetPortRanges(v []*CreateListenerRequestPortRanges) *CreateListenerRequest {
	s.PortRanges = v
	return s
}

func (s *CreateListenerRequest) SetCertificates(v []*CreateListenerRequestCertificates) *CreateListenerRequest {
	s.Certificates = v
	return s
}

type CreateListenerRequestPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateListenerRequestPortRanges) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestPortRanges) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestPortRanges) SetFromPort(v int32) *CreateListenerRequestPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateListenerRequestPortRanges) SetToPort(v int32) *CreateListenerRequestPortRanges {
	s.ToPort = &v
	return s
}

type CreateListenerRequestCertificates struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCertificates) SetId(v string) *CreateListenerRequestCertificates {
	s.Id = &v
	return s
}

type CreateListenerResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s CreateListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBody) SetRequestId(v string) *CreateListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateListenerResponseBody) SetListenerId(v string) *CreateListenerResponseBody {
	s.ListenerId = &v
	return s
}

type CreateListenerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateListenerResponse) SetHeaders(v map[string]*string) *CreateListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateListenerResponse) SetBody(v *CreateListenerResponseBody) *CreateListenerResponse {
	s.Body = v
	return s
}

type DeleteAcceleratorRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s DeleteAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DeleteAcceleratorRequest) SetRegionId(v string) *DeleteAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAcceleratorRequest) SetAcceleratorId(v string) *DeleteAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

type DeleteAcceleratorResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s DeleteAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAcceleratorResponseBody) SetRequestId(v string) *DeleteAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAcceleratorResponseBody) SetAcceleratorId(v string) *DeleteAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

type DeleteAcceleratorResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DeleteAcceleratorResponse) SetHeaders(v map[string]*string) *DeleteAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DeleteAcceleratorResponse) SetBody(v *DeleteAcceleratorResponseBody) *DeleteAcceleratorResponse {
	s.Body = v
	return s
}

type DeleteAclRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AclId       *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclRequest) SetRegionId(v string) *DeleteAclRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAclRequest) SetAclId(v string) *DeleteAclRequest {
	s.AclId = &v
	return s
}

func (s *DeleteAclRequest) SetClientToken(v string) *DeleteAclRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAclRequest) SetDryRun(v bool) *DeleteAclRequest {
	s.DryRun = &v
	return s
}

type DeleteAclResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
}

func (s DeleteAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclResponseBody) SetRequestId(v string) *DeleteAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAclResponseBody) SetAclId(v string) *DeleteAclResponseBody {
	s.AclId = &v
	return s
}

type DeleteAclResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclResponse) SetHeaders(v map[string]*string) *DeleteAclResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclResponse) SetBody(v *DeleteAclResponseBody) *DeleteAclResponse {
	s.Body = v
	return s
}

type DeleteBandwidthPackageRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageRequest) SetRegionId(v string) *DeleteBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetBandwidthPackageId(v string) *DeleteBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetClientToken(v string) *DeleteBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

type DeleteBandwidthPackageResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
}

func (s DeleteBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageResponseBody) SetRequestId(v string) *DeleteBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *DeleteBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

type DeleteBandwidthPackageResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageResponse) SetHeaders(v map[string]*string) *DeleteBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *DeleteBandwidthPackageResponse) SetBody(v *DeleteBandwidthPackageResponseBody) *DeleteBandwidthPackageResponse {
	s.Body = v
	return s
}

type DeleteEndpointGroupRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AcceleratorId   *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s DeleteEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupRequest) SetClientToken(v string) *DeleteEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEndpointGroupRequest) SetAcceleratorId(v string) *DeleteEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteEndpointGroupRequest) SetEndpointGroupId(v string) *DeleteEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

type DeleteEndpointGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupResponseBody) SetRequestId(v string) *DeleteEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEndpointGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupResponse) SetHeaders(v map[string]*string) *DeleteEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteEndpointGroupResponse) SetBody(v *DeleteEndpointGroupResponseBody) *DeleteEndpointGroupResponse {
	s.Body = v
	return s
}

type DeleteForwardingRulesRequest struct {
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken       *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForwardingRuleIds []*string `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	AcceleratorId     *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId        *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DeleteForwardingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesRequest) SetRegionId(v string) *DeleteForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteForwardingRulesRequest) SetClientToken(v string) *DeleteForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteForwardingRulesRequest) SetForwardingRuleIds(v []*string) *DeleteForwardingRulesRequest {
	s.ForwardingRuleIds = v
	return s
}

func (s *DeleteForwardingRulesRequest) SetAcceleratorId(v string) *DeleteForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteForwardingRulesRequest) SetListenerId(v string) *DeleteForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

type DeleteForwardingRulesResponseBody struct {
	ForwardingRules []*DeleteForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteForwardingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesResponseBody) SetForwardingRules(v []*DeleteForwardingRulesResponseBodyForwardingRules) *DeleteForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

func (s *DeleteForwardingRulesResponseBody) SetRequestId(v string) *DeleteForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteForwardingRulesResponseBodyForwardingRules struct {
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
}

func (s DeleteForwardingRulesResponseBodyForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s DeleteForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesResponseBodyForwardingRules) SetForwardingRuleId(v string) *DeleteForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

type DeleteForwardingRulesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteForwardingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteForwardingRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesResponse) SetHeaders(v map[string]*string) *DeleteForwardingRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteForwardingRulesResponse) SetBody(v *DeleteForwardingRulesResponseBody) *DeleteForwardingRulesResponse {
	s.Body = v
	return s
}

type DeleteIpSetRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	IpSetId       *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
}

func (s DeleteIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpSetRequest) SetRegionId(v string) *DeleteIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpSetRequest) SetClientToken(v string) *DeleteIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpSetRequest) SetAcceleratorId(v string) *DeleteIpSetRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteIpSetRequest) SetIpSetId(v string) *DeleteIpSetRequest {
	s.IpSetId = &v
	return s
}

type DeleteIpSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpSetResponseBody) SetRequestId(v string) *DeleteIpSetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpSetResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIpSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpSetResponse) SetHeaders(v map[string]*string) *DeleteIpSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpSetResponse) SetBody(v *DeleteIpSetResponseBody) *DeleteIpSetResponse {
	s.Body = v
	return s
}

type DeleteIpSetsRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IpSetIds []*string `json:"IpSetIds,omitempty" xml:"IpSetIds,omitempty" type:"Repeated"`
}

func (s DeleteIpSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpSetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpSetsRequest) SetRegionId(v string) *DeleteIpSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpSetsRequest) SetIpSetIds(v []*string) *DeleteIpSetsRequest {
	s.IpSetIds = v
	return s
}

type DeleteIpSetsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpSetsResponseBody) SetRequestId(v string) *DeleteIpSetsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpSetsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIpSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpSetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpSetsResponse) SetHeaders(v map[string]*string) *DeleteIpSetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpSetsResponse) SetBody(v *DeleteIpSetsResponseBody) *DeleteIpSetsResponse {
	s.Body = v
	return s
}

type DeleteListenerRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId    *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteListenerRequest) SetAcceleratorId(v string) *DeleteListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteListenerRequest) SetListenerId(v string) *DeleteListenerRequest {
	s.ListenerId = &v
	return s
}

type DeleteListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponseBody) SetRequestId(v string) *DeleteListenerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteListenerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponse) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponse) SetHeaders(v map[string]*string) *DeleteListenerResponse {
	s.Headers = v
	return s
}

func (s *DeleteListenerResponse) SetBody(v *DeleteListenerResponseBody) *DeleteListenerResponse {
	s.Body = v
	return s
}

type DescribeAcceleratorRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s DescribeAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorRequest) SetRegionId(v string) *DescribeAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAcceleratorRequest) SetAcceleratorId(v string) *DescribeAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

type DescribeAcceleratorResponseBody struct {
	DdosId                      *string                                                     `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	DnsName                     *string                                                     `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	Description                 *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId                   *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceChargeType          *string                                                     `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	CreateTime                  *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossDomainBandwidthPackage *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	SecondDnsName               *string                                                     `json:"SecondDnsName,omitempty" xml:"SecondDnsName,omitempty"`
	Name                        *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	BasicBandwidthPackage       *DescribeAcceleratorResponseBodyBasicBandwidthPackage       `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	State                       *string                                                     `json:"State,omitempty" xml:"State,omitempty"`
	ExpiredTime                 *int64                                                      `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	CenId                       *string                                                     `json:"CenId,omitempty" xml:"CenId,omitempty"`
	RegionId                    *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Spec                        *string                                                     `json:"Spec,omitempty" xml:"Spec,omitempty"`
	AcceleratorId               *string                                                     `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s DescribeAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBody) SetDdosId(v string) *DescribeAcceleratorResponseBody {
	s.DdosId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDnsName(v string) *DescribeAcceleratorResponseBody {
	s.DnsName = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDescription(v string) *DescribeAcceleratorResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetRequestId(v string) *DescribeAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetInstanceChargeType(v string) *DescribeAcceleratorResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCreateTime(v int64) *DescribeAcceleratorResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCrossDomainBandwidthPackage(v *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) *DescribeAcceleratorResponseBody {
	s.CrossDomainBandwidthPackage = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetSecondDnsName(v string) *DescribeAcceleratorResponseBody {
	s.SecondDnsName = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetName(v string) *DescribeAcceleratorResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetBasicBandwidthPackage(v *DescribeAcceleratorResponseBodyBasicBandwidthPackage) *DescribeAcceleratorResponseBody {
	s.BasicBandwidthPackage = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetState(v string) *DescribeAcceleratorResponseBody {
	s.State = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetExpiredTime(v int64) *DescribeAcceleratorResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCenId(v string) *DescribeAcceleratorResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetRegionId(v string) *DescribeAcceleratorResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetSpec(v string) *DescribeAcceleratorResponseBody {
	s.Spec = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetAcceleratorId(v string) *DescribeAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

type DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage struct {
	Bandwidth  *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) SetBandwidth(v int32) *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage) SetInstanceId(v string) *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage {
	s.InstanceId = &v
	return s
}

type DescribeAcceleratorResponseBodyBasicBandwidthPackage struct {
	Bandwidth     *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeAcceleratorResponseBodyBasicBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyBasicBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) SetBandwidth(v int32) *DescribeAcceleratorResponseBodyBasicBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) SetBandwidthType(v string) *DescribeAcceleratorResponseBodyBasicBandwidthPackage {
	s.BandwidthType = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyBasicBandwidthPackage) SetInstanceId(v string) *DescribeAcceleratorResponseBodyBasicBandwidthPackage {
	s.InstanceId = &v
	return s
}

type DescribeAcceleratorResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponse) SetHeaders(v map[string]*string) *DescribeAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DescribeAcceleratorResponse) SetBody(v *DescribeAcceleratorResponseBody) *DescribeAcceleratorResponse {
	s.Body = v
	return s
}

type DescribeBandwidthPackageRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
}

func (s DescribeBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageRequest) SetRegionId(v string) *DescribeBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthPackageRequest) SetBandwidthPackageId(v string) *DescribeBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

type DescribeBandwidthPackageResponseBody struct {
	CbnGeographicRegionIdB *string   `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	CbnGeographicRegionIdA *string   `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	Description            *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId              *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime             *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name                   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	BandwidthType          *string   `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	Type                   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Accelerators           []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	State                  *string   `json:"State,omitempty" xml:"State,omitempty"`
	ChargeType             *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Bandwidth              *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ExpiredTime            *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	BandwidthPackageId     *string   `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	RegionId               *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BillingType            *string   `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	Ratio                  *int32    `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s DescribeBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageResponseBody) SetCbnGeographicRegionIdB(v string) *DescribeBandwidthPackageResponseBody {
	s.CbnGeographicRegionIdB = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetCbnGeographicRegionIdA(v string) *DescribeBandwidthPackageResponseBody {
	s.CbnGeographicRegionIdA = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetDescription(v string) *DescribeBandwidthPackageResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRequestId(v string) *DescribeBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetCreateTime(v string) *DescribeBandwidthPackageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetName(v string) *DescribeBandwidthPackageResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidthType(v string) *DescribeBandwidthPackageResponseBody {
	s.BandwidthType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetType(v string) *DescribeBandwidthPackageResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetAccelerators(v []*string) *DescribeBandwidthPackageResponseBody {
	s.Accelerators = v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetState(v string) *DescribeBandwidthPackageResponseBody {
	s.State = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetChargeType(v string) *DescribeBandwidthPackageResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidth(v int32) *DescribeBandwidthPackageResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetExpiredTime(v string) *DescribeBandwidthPackageResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *DescribeBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRegionId(v string) *DescribeBandwidthPackageResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBillingType(v string) *DescribeBandwidthPackageResponseBody {
	s.BillingType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRatio(v int32) *DescribeBandwidthPackageResponseBody {
	s.Ratio = &v
	return s
}

type DescribeBandwidthPackageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageResponse) SetHeaders(v map[string]*string) *DescribeBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwidthPackageResponse) SetBody(v *DescribeBandwidthPackageResponseBody) *DescribeBandwidthPackageResponse {
	s.Body = v
	return s
}

type DescribeEndpointGroupRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s DescribeEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupRequest) SetRegionId(v string) *DescribeEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEndpointGroupRequest) SetEndpointGroupId(v string) *DescribeEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

type DescribeEndpointGroupResponseBody struct {
	HealthCheckIntervalSeconds *int32                                                     `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	TrafficPercentage          *int32                                                     `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
	EndpointGroupId            *string                                                    `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	Description                *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId                  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HealthCheckPath            *string                                                    `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	ThresholdCount             *int32                                                     `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	Name                       *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	EndpointGroupRegion        *string                                                    `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	TotalCount                 *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	State                      *string                                                    `json:"State,omitempty" xml:"State,omitempty"`
	HealthCheckProtocol        *string                                                    `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	HealthCheckPort            *int32                                                     `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	EndpointConfigurations     []*DescribeEndpointGroupResponseBodyEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	PortOverrides              []*DescribeEndpointGroupResponseBodyPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	EndpointRequestProtocol    *string                                                    `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	EndpointGroupType          *string                                                    `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	ForwardingRuleIds          []*string                                                  `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	AcceleratorId              *string                                                    `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId                 *string                                                    `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	SlsRegion                  *string                                                    `json:"SlsRegion,omitempty" xml:"SlsRegion,omitempty"`
	SlsProjectName             *string                                                    `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	SlsLogStoreName            *string                                                    `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	AccessLogSwitch            *string                                                    `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	EnableAccessLog            *bool                                                      `json:"EnableAccessLog,omitempty" xml:"EnableAccessLog,omitempty"`
}

func (s DescribeEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckIntervalSeconds(v int32) *DescribeEndpointGroupResponseBody {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetTrafficPercentage(v int32) *DescribeEndpointGroupResponseBody {
	s.TrafficPercentage = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupId(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetDescription(v string) *DescribeEndpointGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetRequestId(v string) *DescribeEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckPath(v string) *DescribeEndpointGroupResponseBody {
	s.HealthCheckPath = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetThresholdCount(v int32) *DescribeEndpointGroupResponseBody {
	s.ThresholdCount = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetName(v string) *DescribeEndpointGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupRegion(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupRegion = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetTotalCount(v int32) *DescribeEndpointGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetState(v string) *DescribeEndpointGroupResponseBody {
	s.State = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckProtocol(v string) *DescribeEndpointGroupResponseBody {
	s.HealthCheckProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckPort(v int32) *DescribeEndpointGroupResponseBody {
	s.HealthCheckPort = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointConfigurations(v []*DescribeEndpointGroupResponseBodyEndpointConfigurations) *DescribeEndpointGroupResponseBody {
	s.EndpointConfigurations = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetPortOverrides(v []*DescribeEndpointGroupResponseBodyPortOverrides) *DescribeEndpointGroupResponseBody {
	s.PortOverrides = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointRequestProtocol(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupType(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupType = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetForwardingRuleIds(v []*string) *DescribeEndpointGroupResponseBody {
	s.ForwardingRuleIds = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetAcceleratorId(v string) *DescribeEndpointGroupResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetListenerId(v string) *DescribeEndpointGroupResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsRegion(v string) *DescribeEndpointGroupResponseBody {
	s.SlsRegion = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsProjectName(v string) *DescribeEndpointGroupResponseBody {
	s.SlsProjectName = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsLogStoreName(v string) *DescribeEndpointGroupResponseBody {
	s.SlsLogStoreName = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetAccessLogSwitch(v string) *DescribeEndpointGroupResponseBody {
	s.AccessLogSwitch = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEnableAccessLog(v bool) *DescribeEndpointGroupResponseBody {
	s.EnableAccessLog = &v
	return s
}

type DescribeEndpointGroupResponseBodyEndpointConfigurations struct {
	Type                       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EnableClientIPPreservation *bool   `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	Weight                     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	ProbeProtocol              *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	Endpoint                   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ProbePort                  *int32  `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetType(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetEnableClientIPPreservation(v bool) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetWeight(v int32) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetProbeProtocol(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.ProbeProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetEndpoint(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetProbePort(v int32) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.ProbePort = &v
	return s
}

type DescribeEndpointGroupResponseBodyPortOverrides struct {
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyPortOverrides) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) SetListenerPort(v int32) *DescribeEndpointGroupResponseBodyPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) SetEndpointPort(v int32) *DescribeEndpointGroupResponseBodyPortOverrides {
	s.EndpointPort = &v
	return s
}

type DescribeEndpointGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponse) SetHeaders(v map[string]*string) *DescribeEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointGroupResponse) SetBody(v *DescribeEndpointGroupResponseBody) *DescribeEndpointGroupResponse {
	s.Body = v
	return s
}

type DescribeIpSetRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IpSetId  *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
}

func (s DescribeIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpSetRequest) SetRegionId(v string) *DescribeIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIpSetRequest) SetIpSetId(v string) *DescribeIpSetRequest {
	s.IpSetId = &v
	return s
}

type DescribeIpSetResponseBody struct {
	IpSetId            *string   `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpVersion          *string   `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	State              *string   `json:"State,omitempty" xml:"State,omitempty"`
	Bandwidth          *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	IpAddressList      []*string `json:"IpAddressList,omitempty" xml:"IpAddressList,omitempty" type:"Repeated"`
	AccelerateRegionId *string   `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	AcceleratorId      *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s DescribeIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpSetResponseBody) SetIpSetId(v string) *DescribeIpSetResponseBody {
	s.IpSetId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetRequestId(v string) *DescribeIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetIpVersion(v string) *DescribeIpSetResponseBody {
	s.IpVersion = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetState(v string) *DescribeIpSetResponseBody {
	s.State = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetBandwidth(v int32) *DescribeIpSetResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetIpAddressList(v []*string) *DescribeIpSetResponseBody {
	s.IpAddressList = v
	return s
}

func (s *DescribeIpSetResponseBody) SetAccelerateRegionId(v string) *DescribeIpSetResponseBody {
	s.AccelerateRegionId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetAcceleratorId(v string) *DescribeIpSetResponseBody {
	s.AcceleratorId = &v
	return s
}

type DescribeIpSetResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpSetResponse) SetHeaders(v map[string]*string) *DescribeIpSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpSetResponse) SetBody(v *DescribeIpSetResponseBody) *DescribeIpSetResponse {
	s.Body = v
	return s
}

type DescribeListenerRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DescribeListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerRequest) GoString() string {
	return s.String()
}

func (s *DescribeListenerRequest) SetRegionId(v string) *DescribeListenerRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeListenerRequest) SetListenerId(v string) *DescribeListenerRequest {
	s.ListenerId = &v
	return s
}

type DescribeListenerResponseBody struct {
	Description    *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State          *string                                     `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime     *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PortRanges     []*DescribeListenerResponseBodyPortRanges   `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	BackendPorts   []*DescribeListenerResponseBodyBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	Certificates   []*DescribeListenerResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	Protocol       *string                                     `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ListenerId     *string                                     `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	ClientAffinity *string                                     `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	Name           *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RelatedAcls    []*DescribeListenerResponseBodyRelatedAcls  `json:"RelatedAcls,omitempty" xml:"RelatedAcls,omitempty" type:"Repeated"`
	AclType        *string                                     `json:"AclType,omitempty" xml:"AclType,omitempty"`
	AcceleratorId  *string                                     `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ProxyProtocol  *bool                                       `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
}

func (s DescribeListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBody) SetDescription(v string) *DescribeListenerResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeListenerResponseBody) SetRequestId(v string) *DescribeListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetState(v string) *DescribeListenerResponseBody {
	s.State = &v
	return s
}

func (s *DescribeListenerResponseBody) SetCreateTime(v string) *DescribeListenerResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeListenerResponseBody) SetPortRanges(v []*DescribeListenerResponseBodyPortRanges) *DescribeListenerResponseBody {
	s.PortRanges = v
	return s
}

func (s *DescribeListenerResponseBody) SetBackendPorts(v []*DescribeListenerResponseBodyBackendPorts) *DescribeListenerResponseBody {
	s.BackendPorts = v
	return s
}

func (s *DescribeListenerResponseBody) SetCertificates(v []*DescribeListenerResponseBodyCertificates) *DescribeListenerResponseBody {
	s.Certificates = v
	return s
}

func (s *DescribeListenerResponseBody) SetProtocol(v string) *DescribeListenerResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeListenerResponseBody) SetListenerId(v string) *DescribeListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetClientAffinity(v string) *DescribeListenerResponseBody {
	s.ClientAffinity = &v
	return s
}

func (s *DescribeListenerResponseBody) SetName(v string) *DescribeListenerResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeListenerResponseBody) SetRelatedAcls(v []*DescribeListenerResponseBodyRelatedAcls) *DescribeListenerResponseBody {
	s.RelatedAcls = v
	return s
}

func (s *DescribeListenerResponseBody) SetAclType(v string) *DescribeListenerResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeListenerResponseBody) SetAcceleratorId(v string) *DescribeListenerResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetProxyProtocol(v bool) *DescribeListenerResponseBody {
	s.ProxyProtocol = &v
	return s
}

type DescribeListenerResponseBodyPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s DescribeListenerResponseBodyPortRanges) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponseBodyPortRanges) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyPortRanges) SetFromPort(v int32) *DescribeListenerResponseBodyPortRanges {
	s.FromPort = &v
	return s
}

func (s *DescribeListenerResponseBodyPortRanges) SetToPort(v int32) *DescribeListenerResponseBodyPortRanges {
	s.ToPort = &v
	return s
}

type DescribeListenerResponseBodyBackendPorts struct {
	FromPort *string `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *string `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s DescribeListenerResponseBodyBackendPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponseBodyBackendPorts) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyBackendPorts) SetFromPort(v string) *DescribeListenerResponseBodyBackendPorts {
	s.FromPort = &v
	return s
}

func (s *DescribeListenerResponseBodyBackendPorts) SetToPort(v string) *DescribeListenerResponseBodyBackendPorts {
	s.ToPort = &v
	return s
}

type DescribeListenerResponseBodyCertificates struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeListenerResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyCertificates) SetType(v string) *DescribeListenerResponseBodyCertificates {
	s.Type = &v
	return s
}

func (s *DescribeListenerResponseBodyCertificates) SetId(v string) *DescribeListenerResponseBodyCertificates {
	s.Id = &v
	return s
}

type DescribeListenerResponseBodyRelatedAcls struct {
	AclId  *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeListenerResponseBodyRelatedAcls) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponseBodyRelatedAcls) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyRelatedAcls) SetAclId(v string) *DescribeListenerResponseBodyRelatedAcls {
	s.AclId = &v
	return s
}

func (s *DescribeListenerResponseBodyRelatedAcls) SetStatus(v string) *DescribeListenerResponseBodyRelatedAcls {
	s.Status = &v
	return s
}

type DescribeListenerResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponse) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponse) SetHeaders(v map[string]*string) *DescribeListenerResponse {
	s.Headers = v
	return s
}

func (s *DescribeListenerResponse) SetBody(v *DescribeListenerResponseBody) *DescribeListenerResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DetachDdosFromAcceleratorRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachDdosFromAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachDdosFromAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DetachDdosFromAcceleratorRequest) SetAcceleratorId(v string) *DetachDdosFromAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DetachDdosFromAcceleratorRequest) SetRegionId(v string) *DetachDdosFromAcceleratorRequest {
	s.RegionId = &v
	return s
}

type DetachDdosFromAcceleratorResponseBody struct {
	DdosId    *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachDdosFromAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachDdosFromAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDdosFromAcceleratorResponseBody) SetDdosId(v string) *DetachDdosFromAcceleratorResponseBody {
	s.DdosId = &v
	return s
}

func (s *DetachDdosFromAcceleratorResponseBody) SetRequestId(v string) *DetachDdosFromAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type DetachDdosFromAcceleratorResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachDdosFromAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachDdosFromAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachDdosFromAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DetachDdosFromAcceleratorResponse) SetHeaders(v map[string]*string) *DetachDdosFromAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DetachDdosFromAcceleratorResponse) SetBody(v *DetachDdosFromAcceleratorResponseBody) *DetachDdosFromAcceleratorResponse {
	s.Body = v
	return s
}

type DetachLogStoreFromEndpointGroupRequest struct {
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceleratorId    *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId       *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	ClientToken      *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DetachLogStoreFromEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachLogStoreFromEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetRegionId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetAcceleratorId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetListenerId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetEndpointGroupIds(v []*string) *DetachLogStoreFromEndpointGroupRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetClientToken(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

type DetachLogStoreFromEndpointGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachLogStoreFromEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachLogStoreFromEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLogStoreFromEndpointGroupResponseBody) SetRequestId(v string) *DetachLogStoreFromEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

type DetachLogStoreFromEndpointGroupResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachLogStoreFromEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachLogStoreFromEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachLogStoreFromEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachLogStoreFromEndpointGroupResponse) SetHeaders(v map[string]*string) *DetachLogStoreFromEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachLogStoreFromEndpointGroupResponse) SetBody(v *DetachLogStoreFromEndpointGroupResponseBody) *DetachLogStoreFromEndpointGroupResponse {
	s.Body = v
	return s
}

type DissociateAclsFromListenerRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AclIds      []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	ListenerId  *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DissociateAclsFromListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerRequest) SetRegionId(v string) *DissociateAclsFromListenerRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetAclIds(v []*string) *DissociateAclsFromListenerRequest {
	s.AclIds = v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetListenerId(v string) *DissociateAclsFromListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetClientToken(v string) *DissociateAclsFromListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetDryRun(v bool) *DissociateAclsFromListenerRequest {
	s.DryRun = &v
	return s
}

type DissociateAclsFromListenerResponseBody struct {
	// Id of the request
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclIds     []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	ListenerId *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DissociateAclsFromListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponseBody) SetRequestId(v string) *DissociateAclsFromListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetAclIds(v []*string) *DissociateAclsFromListenerResponseBody {
	s.AclIds = v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetListenerId(v string) *DissociateAclsFromListenerResponseBody {
	s.ListenerId = &v
	return s
}

type DissociateAclsFromListenerResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissociateAclsFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateAclsFromListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerResponse) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponse) SetHeaders(v map[string]*string) *DissociateAclsFromListenerResponse {
	s.Headers = v
	return s
}

func (s *DissociateAclsFromListenerResponse) SetBody(v *DissociateAclsFromListenerResponseBody) *DissociateAclsFromListenerResponse {
	s.Body = v
	return s
}

type GetAclRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AclId    *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
}

func (s GetAclRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAclRequest) GoString() string {
	return s.String()
}

func (s *GetAclRequest) SetRegionId(v string) *GetAclRequest {
	s.RegionId = &v
	return s
}

func (s *GetAclRequest) SetAclId(v string) *GetAclRequest {
	s.AclId = &v
	return s
}

type GetAclResponseBody struct {
	// Id of the request
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclId            *string                               `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AddressIPVersion *string                               `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AclStatus        *string                               `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclEntries       []*GetAclResponseBodyAclEntries       `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	RelatedListeners []*GetAclResponseBodyRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
	AclName          *string                               `json:"AclName,omitempty" xml:"AclName,omitempty"`
}

func (s GetAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetAclResponseBody) SetRequestId(v string) *GetAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAclResponseBody) SetAclId(v string) *GetAclResponseBody {
	s.AclId = &v
	return s
}

func (s *GetAclResponseBody) SetAddressIPVersion(v string) *GetAclResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *GetAclResponseBody) SetAclStatus(v string) *GetAclResponseBody {
	s.AclStatus = &v
	return s
}

func (s *GetAclResponseBody) SetAclEntries(v []*GetAclResponseBodyAclEntries) *GetAclResponseBody {
	s.AclEntries = v
	return s
}

func (s *GetAclResponseBody) SetRelatedListeners(v []*GetAclResponseBodyRelatedListeners) *GetAclResponseBody {
	s.RelatedListeners = v
	return s
}

func (s *GetAclResponseBody) SetAclName(v string) *GetAclResponseBody {
	s.AclName = &v
	return s
}

type GetAclResponseBodyAclEntries struct {
	Entry            *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	EntryDescription *string `json:"EntryDescription,omitempty" xml:"EntryDescription,omitempty"`
}

func (s GetAclResponseBodyAclEntries) String() string {
	return tea.Prettify(s)
}

func (s GetAclResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *GetAclResponseBodyAclEntries) SetEntry(v string) *GetAclResponseBodyAclEntries {
	s.Entry = &v
	return s
}

func (s *GetAclResponseBodyAclEntries) SetEntryDescription(v string) *GetAclResponseBodyAclEntries {
	s.EntryDescription = &v
	return s
}

type GetAclResponseBodyRelatedListeners struct {
	ListenerId    *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	AclType       *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s GetAclResponseBodyRelatedListeners) String() string {
	return tea.Prettify(s)
}

func (s GetAclResponseBodyRelatedListeners) GoString() string {
	return s.String()
}

func (s *GetAclResponseBodyRelatedListeners) SetListenerId(v string) *GetAclResponseBodyRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *GetAclResponseBodyRelatedListeners) SetAclType(v string) *GetAclResponseBodyRelatedListeners {
	s.AclType = &v
	return s
}

func (s *GetAclResponseBodyRelatedListeners) SetAcceleratorId(v string) *GetAclResponseBodyRelatedListeners {
	s.AcceleratorId = &v
	return s
}

type GetAclResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAclResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAclResponse) GoString() string {
	return s.String()
}

func (s *GetAclResponse) SetHeaders(v map[string]*string) *GetAclResponse {
	s.Headers = v
	return s
}

func (s *GetAclResponse) SetBody(v *GetAclResponseBody) *GetAclResponse {
	s.Body = v
	return s
}

type ListAccelerateAreasRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAccelerateAreasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccelerateAreasRequest) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasRequest) SetRegionId(v string) *ListAccelerateAreasRequest {
	s.RegionId = &v
	return s
}

type ListAccelerateAreasResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Areas     []*ListAccelerateAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
}

func (s ListAccelerateAreasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccelerateAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponseBody) SetRequestId(v string) *ListAccelerateAreasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccelerateAreasResponseBody) SetAreas(v []*ListAccelerateAreasResponseBodyAreas) *ListAccelerateAreasResponseBody {
	s.Areas = v
	return s
}

type ListAccelerateAreasResponseBodyAreas struct {
	LocalName  *string                                           `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	AreaId     *string                                           `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	RegionList []*ListAccelerateAreasResponseBodyAreasRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s ListAccelerateAreasResponseBodyAreas) String() string {
	return tea.Prettify(s)
}

func (s ListAccelerateAreasResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponseBodyAreas) SetLocalName(v string) *ListAccelerateAreasResponseBodyAreas {
	s.LocalName = &v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreas) SetAreaId(v string) *ListAccelerateAreasResponseBodyAreas {
	s.AreaId = &v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreas) SetRegionList(v []*ListAccelerateAreasResponseBodyAreasRegionList) *ListAccelerateAreasResponseBodyAreas {
	s.RegionList = v
	return s
}

type ListAccelerateAreasResponseBodyAreasRegionList struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAccelerateAreasResponseBodyAreasRegionList) String() string {
	return tea.Prettify(s)
}

func (s ListAccelerateAreasResponseBodyAreasRegionList) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponseBodyAreasRegionList) SetLocalName(v string) *ListAccelerateAreasResponseBodyAreasRegionList {
	s.LocalName = &v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreasRegionList) SetRegionId(v string) *ListAccelerateAreasResponseBodyAreasRegionList {
	s.RegionId = &v
	return s
}

type ListAccelerateAreasResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAccelerateAreasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccelerateAreasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccelerateAreasResponse) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponse) SetHeaders(v map[string]*string) *ListAccelerateAreasResponse {
	s.Headers = v
	return s
}

func (s *ListAccelerateAreasResponse) SetBody(v *ListAccelerateAreasResponseBody) *ListAccelerateAreasResponse {
	s.Body = v
	return s
}

type ListAcceleratorsRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s ListAcceleratorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsRequest) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsRequest) SetRegionId(v string) *ListAcceleratorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAcceleratorsRequest) SetPageNumber(v int32) *ListAcceleratorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAcceleratorsRequest) SetPageSize(v int32) *ListAcceleratorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAcceleratorsRequest) SetAcceleratorId(v string) *ListAcceleratorsRequest {
	s.AcceleratorId = &v
	return s
}

type ListAcceleratorsResponseBody struct {
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize     *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Accelerators []*ListAcceleratorsResponseBodyAccelerators `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	PageNumber   *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListAcceleratorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBody) SetTotalCount(v int32) *ListAcceleratorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAcceleratorsResponseBody) SetPageSize(v int32) *ListAcceleratorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAcceleratorsResponseBody) SetRequestId(v string) *ListAcceleratorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAcceleratorsResponseBody) SetAccelerators(v []*ListAcceleratorsResponseBodyAccelerators) *ListAcceleratorsResponseBody {
	s.Accelerators = v
	return s
}

func (s *ListAcceleratorsResponseBody) SetPageNumber(v int32) *ListAcceleratorsResponseBody {
	s.PageNumber = &v
	return s
}

type ListAcceleratorsResponseBodyAccelerators struct {
	DnsName                     *string                                                              `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	Type                        *string                                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	SecondDnsName               *string                                                              `json:"SecondDnsName,omitempty" xml:"SecondDnsName,omitempty"`
	Spec                        *string                                                              `json:"Spec,omitempty" xml:"Spec,omitempty"`
	State                       *string                                                              `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime                  *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CenId                       *string                                                              `json:"CenId,omitempty" xml:"CenId,omitempty"`
	DdosId                      *string                                                              `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	BasicBandwidthPackage       *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage       `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	RegionId                    *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceChargeType          *string                                                              `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	AcceleratorId               *string                                                              `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Description                 *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Bandwidth                   *int32                                                               `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ExpiredTime                 *int64                                                               `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Name                        *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	CrossDomainBandwidthPackage *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
}

func (s ListAcceleratorsResponseBodyAccelerators) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAccelerators) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDnsName(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.DnsName = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetType(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Type = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetSecondDnsName(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.SecondDnsName = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetSpec(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Spec = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetState(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.State = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCreateTime(v int64) *ListAcceleratorsResponseBodyAccelerators {
	s.CreateTime = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCenId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.CenId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDdosId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.DdosId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetBasicBandwidthPackage(v *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) *ListAcceleratorsResponseBodyAccelerators {
	s.BasicBandwidthPackage = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetRegionId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.RegionId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetInstanceChargeType(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetAcceleratorId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.AcceleratorId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDescription(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Description = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetBandwidth(v int32) *ListAcceleratorsResponseBodyAccelerators {
	s.Bandwidth = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetExpiredTime(v int64) *ListAcceleratorsResponseBodyAccelerators {
	s.ExpiredTime = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetName(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Name = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCrossDomainBandwidthPackage(v *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) *ListAcceleratorsResponseBodyAccelerators {
	s.CrossDomainBandwidthPackage = v
	return s
}

type ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage struct {
	Bandwidth     *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetBandwidth(v int32) *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetBandwidthType(v string) *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.BandwidthType = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetInstanceId(v string) *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.InstanceId = &v
	return s
}

type ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage struct {
	Bandwidth  *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) SetBandwidth(v int32) *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) SetInstanceId(v string) *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	s.InstanceId = &v
	return s
}

type ListAcceleratorsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAcceleratorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAcceleratorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponse) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponse) SetHeaders(v map[string]*string) *ListAcceleratorsResponse {
	s.Headers = v
	return s
}

func (s *ListAcceleratorsResponse) SetBody(v *ListAcceleratorsResponseBody) *ListAcceleratorsResponse {
	s.Body = v
	return s
}

type ListAclsRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AclIds      []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	AclName     *string   `json:"AclName,omitempty" xml:"AclName,omitempty"`
	NextToken   *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListAclsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclsRequest) GoString() string {
	return s.String()
}

func (s *ListAclsRequest) SetRegionId(v string) *ListAclsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAclsRequest) SetClientToken(v string) *ListAclsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAclsRequest) SetAclIds(v []*string) *ListAclsRequest {
	s.AclIds = v
	return s
}

func (s *ListAclsRequest) SetAclName(v string) *ListAclsRequest {
	s.AclName = &v
	return s
}

func (s *ListAclsRequest) SetNextToken(v string) *ListAclsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAclsRequest) SetMaxResults(v int32) *ListAclsRequest {
	s.MaxResults = &v
	return s
}

type ListAclsResponseBody struct {
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken  *string                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Acls       []*ListAclsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
}

func (s ListAclsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBody) SetRequestId(v string) *ListAclsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclsResponseBody) SetTotalCount(v int32) *ListAclsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAclsResponseBody) SetNextToken(v string) *ListAclsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAclsResponseBody) SetMaxResults(v int32) *ListAclsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAclsResponseBody) SetAcls(v []*ListAclsResponseBodyAcls) *ListAclsResponseBody {
	s.Acls = v
	return s
}

type ListAclsResponseBodyAcls struct {
	AclId            *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName          *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AclStatus        *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
}

func (s ListAclsResponseBodyAcls) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBodyAcls) SetAclId(v string) *ListAclsResponseBodyAcls {
	s.AclId = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAclName(v string) *ListAclsResponseBodyAcls {
	s.AclName = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAddressIPVersion(v string) *ListAclsResponseBodyAcls {
	s.AddressIPVersion = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAclStatus(v string) *ListAclsResponseBodyAcls {
	s.AclStatus = &v
	return s
}

type ListAclsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAclsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAclsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponse) GoString() string {
	return s.String()
}

func (s *ListAclsResponse) SetHeaders(v map[string]*string) *ListAclsResponse {
	s.Headers = v
	return s
}

func (s *ListAclsResponse) SetBody(v *ListAclsResponseBody) *ListAclsResponse {
	s.Body = v
	return s
}

type ListAvailableAccelerateAreasRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s ListAvailableAccelerateAreasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableAccelerateAreasRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasRequest) SetRegionId(v string) *ListAvailableAccelerateAreasRequest {
	s.RegionId = &v
	return s
}

func (s *ListAvailableAccelerateAreasRequest) SetAcceleratorId(v string) *ListAvailableAccelerateAreasRequest {
	s.AcceleratorId = &v
	return s
}

type ListAvailableAccelerateAreasResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Areas     []*ListAvailableAccelerateAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
}

func (s ListAvailableAccelerateAreasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponseBody) SetRequestId(v string) *ListAvailableAccelerateAreasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBody) SetAreas(v []*ListAvailableAccelerateAreasResponseBodyAreas) *ListAvailableAccelerateAreasResponseBody {
	s.Areas = v
	return s
}

type ListAvailableAccelerateAreasResponseBodyAreas struct {
	LocalName  *string                                                    `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	AreaId     *string                                                    `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	RegionList []*ListAvailableAccelerateAreasResponseBodyAreasRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s ListAvailableAccelerateAreasResponseBodyAreas) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) SetLocalName(v string) *ListAvailableAccelerateAreasResponseBodyAreas {
	s.LocalName = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) SetAreaId(v string) *ListAvailableAccelerateAreasResponseBodyAreas {
	s.AreaId = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) SetRegionList(v []*ListAvailableAccelerateAreasResponseBodyAreasRegionList) *ListAvailableAccelerateAreasResponseBodyAreas {
	s.RegionList = v
	return s
}

type ListAvailableAccelerateAreasResponseBodyAreasRegionList struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAvailableAccelerateAreasResponseBodyAreasRegionList) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponseBodyAreasRegionList) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) SetLocalName(v string) *ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	s.LocalName = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreasRegionList) SetRegionId(v string) *ListAvailableAccelerateAreasResponseBodyAreasRegionList {
	s.RegionId = &v
	return s
}

type ListAvailableAccelerateAreasResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAvailableAccelerateAreasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableAccelerateAreasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponse) SetHeaders(v map[string]*string) *ListAvailableAccelerateAreasResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableAccelerateAreasResponse) SetBody(v *ListAvailableAccelerateAreasResponseBody) *ListAvailableAccelerateAreasResponse {
	s.Body = v
	return s
}

type ListAvailableBusiRegionsRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s ListAvailableBusiRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBusiRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsRequest) SetRegionId(v string) *ListAvailableBusiRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAvailableBusiRegionsRequest) SetAcceleratorId(v string) *ListAvailableBusiRegionsRequest {
	s.AcceleratorId = &v
	return s
}

type ListAvailableBusiRegionsResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*ListAvailableBusiRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s ListAvailableBusiRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBusiRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsResponseBody) SetRequestId(v string) *ListAvailableBusiRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableBusiRegionsResponseBody) SetRegions(v []*ListAvailableBusiRegionsResponseBodyRegions) *ListAvailableBusiRegionsResponseBody {
	s.Regions = v
	return s
}

type ListAvailableBusiRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAvailableBusiRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBusiRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) SetLocalName(v string) *ListAvailableBusiRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) SetRegionId(v string) *ListAvailableBusiRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListAvailableBusiRegionsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAvailableBusiRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableBusiRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBusiRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsResponse) SetHeaders(v map[string]*string) *ListAvailableBusiRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableBusiRegionsResponse) SetBody(v *ListAvailableBusiRegionsResponseBody) *ListAvailableBusiRegionsResponse {
	s.Body = v
	return s
}

type ListBandwidthackagesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBandwidthackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthackagesRequest) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesRequest) SetRegionId(v string) *ListBandwidthackagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthackagesRequest) SetPageNumber(v int32) *ListBandwidthackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthackagesRequest) SetPageSize(v int32) *ListBandwidthackagesRequest {
	s.PageSize = &v
	return s
}

type ListBandwidthackagesResponseBody struct {
	TotalCount        *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize          *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	BandwidthPackages []*ListBandwidthackagesResponseBodyBandwidthPackages `json:"BandwidthPackages,omitempty" xml:"BandwidthPackages,omitempty" type:"Repeated"`
}

func (s ListBandwidthackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponseBody) SetTotalCount(v int32) *ListBandwidthackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetPageSize(v int32) *ListBandwidthackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetRequestId(v string) *ListBandwidthackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetPageNumber(v int32) *ListBandwidthackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetBandwidthPackages(v []*ListBandwidthackagesResponseBodyBandwidthPackages) *ListBandwidthackagesResponseBody {
	s.BandwidthPackages = v
	return s
}

type ListBandwidthackagesResponseBodyBandwidthPackages struct {
	BandwidthPackageId *string   `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	Bandwidth          *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description        *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime        *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	State              *string   `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime         *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType         *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Accelerators       []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	Name               *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBandwidthackagesResponseBodyBandwidthPackages) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthackagesResponseBodyBandwidthPackages) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetBandwidthPackageId(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetBandwidth(v int32) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Bandwidth = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetDescription(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Description = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetExpiredTime(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.ExpiredTime = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetState(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.State = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetCreateTime(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.CreateTime = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetChargeType(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.ChargeType = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetAccelerators(v []*string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Accelerators = v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetName(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Name = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetRegionId(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.RegionId = &v
	return s
}

type ListBandwidthackagesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBandwidthackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBandwidthackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthackagesResponse) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponse) SetHeaders(v map[string]*string) *ListBandwidthackagesResponse {
	s.Headers = v
	return s
}

func (s *ListBandwidthackagesResponse) SetBody(v *ListBandwidthackagesResponseBody) *ListBandwidthackagesResponse {
	s.Body = v
	return s
}

type ListBandwidthPackagesRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
}

func (s ListBandwidthPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesRequest) SetRegionId(v string) *ListBandwidthPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetPageNumber(v int32) *ListBandwidthPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetPageSize(v int32) *ListBandwidthPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetState(v string) *ListBandwidthPackagesRequest {
	s.State = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetType(v string) *ListBandwidthPackagesRequest {
	s.Type = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetBandwidthPackageId(v string) *ListBandwidthPackagesRequest {
	s.BandwidthPackageId = &v
	return s
}

type ListBandwidthPackagesResponseBody struct {
	TotalCount        *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize          *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	BandwidthPackages []*ListBandwidthPackagesResponseBodyBandwidthPackages `json:"BandwidthPackages,omitempty" xml:"BandwidthPackages,omitempty" type:"Repeated"`
}

func (s ListBandwidthPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponseBody) SetTotalCount(v int32) *ListBandwidthPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetPageSize(v int32) *ListBandwidthPackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetRequestId(v string) *ListBandwidthPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetPageNumber(v int32) *ListBandwidthPackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetBandwidthPackages(v []*ListBandwidthPackagesResponseBodyBandwidthPackages) *ListBandwidthPackagesResponseBody {
	s.BandwidthPackages = v
	return s
}

type ListBandwidthPackagesResponseBodyBandwidthPackages struct {
	Type                   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	BandwidthType          *string   `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	State                  *string   `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime             *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType             *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	RegionId               *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CbnGeographicRegionIdA *string   `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	BandwidthPackageId     *string   `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	Bandwidth              *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description            *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime            *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Accelerators           []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	CbnGeographicRegionIdB *string   `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	Name                   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	BillingType            *string   `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	Ratio                  *int32    `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackages) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackages) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Type = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidthType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BandwidthType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetState(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.State = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCreateTime(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CreateTime = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetChargeType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.ChargeType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetRegionId(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCbnGeographicRegionIdA(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CbnGeographicRegionIdA = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidthPackageId(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidth(v int32) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Bandwidth = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetDescription(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Description = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetExpiredTime(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.ExpiredTime = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetAccelerators(v []*string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Accelerators = v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCbnGeographicRegionIdB(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CbnGeographicRegionIdB = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetName(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Name = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBillingType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BillingType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetRatio(v int32) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Ratio = &v
	return s
}

type ListBandwidthPackagesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBandwidthPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBandwidthPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponse) SetHeaders(v map[string]*string) *ListBandwidthPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListBandwidthPackagesResponse) SetBody(v *ListBandwidthPackagesResponseBody) *ListBandwidthPackagesResponse {
	s.Body = v
	return s
}

type ListBusiRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBusiRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBusiRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsRequest) SetRegionId(v string) *ListBusiRegionsRequest {
	s.RegionId = &v
	return s
}

type ListBusiRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*ListBusiRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s ListBusiRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBusiRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsResponseBody) SetRequestId(v string) *ListBusiRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBusiRegionsResponseBody) SetRegions(v []*ListBusiRegionsResponseBodyRegions) *ListBusiRegionsResponseBody {
	s.Regions = v
	return s
}

type ListBusiRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBusiRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListBusiRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsResponseBodyRegions) SetLocalName(v string) *ListBusiRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListBusiRegionsResponseBodyRegions) SetRegionId(v string) *ListBusiRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListBusiRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBusiRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBusiRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBusiRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsResponse) SetHeaders(v map[string]*string) *ListBusiRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListBusiRegionsResponse) SetBody(v *ListBusiRegionsResponseBody) *ListBusiRegionsResponse {
	s.Body = v
	return s
}

type ListEndpointGroupsRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AcceleratorId     *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId        *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	AccessLogSwitch   *string `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	EndpointGroupId   *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s ListEndpointGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsRequest) SetRegionId(v string) *ListEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetPageNumber(v int32) *ListEndpointGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetPageSize(v int32) *ListEndpointGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetAcceleratorId(v string) *ListEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetListenerId(v string) *ListEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetEndpointGroupType(v string) *ListEndpointGroupsRequest {
	s.EndpointGroupType = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetAccessLogSwitch(v string) *ListEndpointGroupsRequest {
	s.AccessLogSwitch = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetEndpointGroupId(v string) *ListEndpointGroupsRequest {
	s.EndpointGroupId = &v
	return s
}

type ListEndpointGroupsResponseBody struct {
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize       *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	EndpointGroups []*ListEndpointGroupsResponseBodyEndpointGroups `json:"EndpointGroups,omitempty" xml:"EndpointGroups,omitempty" type:"Repeated"`
}

func (s ListEndpointGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBody) SetTotalCount(v int32) *ListEndpointGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetPageSize(v int32) *ListEndpointGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetRequestId(v string) *ListEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetPageNumber(v int32) *ListEndpointGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetEndpointGroups(v []*ListEndpointGroupsResponseBodyEndpointGroups) *ListEndpointGroupsResponseBody {
	s.EndpointGroups = v
	return s
}

type ListEndpointGroupsResponseBodyEndpointGroups struct {
	EndpointGroupId            *string                                                               `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupIpList        []*string                                                             `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	State                      *string                                                               `json:"State,omitempty" xml:"State,omitempty"`
	HealthCheckPath            *string                                                               `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	EndpointGroupRegion        *string                                                               `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	HealthCheckIntervalSeconds *int32                                                                `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	TrafficPercentage          *int32                                                                `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
	HealthCheckProtocol        *string                                                               `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	ThresholdCount             *int32                                                                `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	ListenerId                 *string                                                               `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	AcceleratorId              *string                                                               `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointConfigurations     []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	PortOverrides              []*ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	ForwardingRuleIds          []*string                                                             `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	EndpointGroupType          *string                                                               `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointRequestProtocol    *string                                                               `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	Description                *string                                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                       *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	HealthCheckPort            *int32                                                                `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroups) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroups) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupIpList(v []*string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupIpList = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetState(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.State = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckPath(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckPath = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupRegion(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckIntervalSeconds(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetTrafficPercentage(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.TrafficPercentage = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetThresholdCount(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ThresholdCount = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetListenerId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ListenerId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetAcceleratorId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.AcceleratorId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointConfigurations(v []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointConfigurations = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetPortOverrides(v []*ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.PortOverrides = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetForwardingRuleIds(v []*string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ForwardingRuleIds = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupType(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupType = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointRequestProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetDescription(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.Description = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetName(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.Name = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckPort = &v
	return s
}

type ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations struct {
	Type                       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EnableClientIPPreservation *bool   `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	Weight                     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	ProbeProtocol              *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	Endpoint                   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ProbePort                  *int32  `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetType(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEnableClientIPPreservation(v bool) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetWeight(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetProbeProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.ProbeProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEndpoint(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetProbePort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.ProbePort = &v
	return s
}

type ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides struct {
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) SetListenerPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) SetEndpointPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides {
	s.EndpointPort = &v
	return s
}

type ListEndpointGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEndpointGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponse) SetHeaders(v map[string]*string) *ListEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListEndpointGroupsResponse) SetBody(v *ListEndpointGroupsResponseBody) *ListEndpointGroupsResponse {
	s.Body = v
	return s
}

type ListForwardingRulesRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ListenerId       *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	AcceleratorId    *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListForwardingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesRequest) SetRegionId(v string) *ListForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetClientToken(v string) *ListForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListForwardingRulesRequest) SetListenerId(v string) *ListForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetAcceleratorId(v string) *ListForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetForwardingRuleId(v string) *ListForwardingRulesRequest {
	s.ForwardingRuleId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetNextToken(v string) *ListForwardingRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListForwardingRulesRequest) SetMaxResults(v int32) *ListForwardingRulesRequest {
	s.MaxResults = &v
	return s
}

type ListForwardingRulesResponseBody struct {
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken       *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults      *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ForwardingRules []*ListForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBody) SetRequestId(v string) *ListForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetTotalCount(v int32) *ListForwardingRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetNextToken(v string) *ListForwardingRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetMaxResults(v int32) *ListForwardingRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetForwardingRules(v []*ListForwardingRulesResponseBodyForwardingRules) *ListForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

type ListForwardingRulesResponseBodyForwardingRules struct {
	Priority             *int32                                                          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ForwardingRuleId     *string                                                         `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	ForwardingRuleName   *string                                                         `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	ForwardingRuleStatus *string                                                         `json:"ForwardingRuleStatus,omitempty" xml:"ForwardingRuleStatus,omitempty"`
	RuleConditions       []*ListForwardingRulesResponseBodyForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	RuleActions          []*ListForwardingRulesResponseBodyForwardingRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	ListenerId           *string                                                         `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s ListForwardingRulesResponseBodyForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetPriority(v int32) *ListForwardingRulesResponseBodyForwardingRules {
	s.Priority = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetForwardingRuleId(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetForwardingRuleName(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetForwardingRuleStatus(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleStatus = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetRuleConditions(v []*ListForwardingRulesResponseBodyForwardingRulesRuleConditions) *ListForwardingRulesResponseBodyForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetRuleActions(v []*ListForwardingRulesResponseBodyForwardingRulesRuleActions) *ListForwardingRulesResponseBodyForwardingRules {
	s.RuleActions = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetListenerId(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ListenerId = &v
	return s
}

type ListForwardingRulesResponseBodyForwardingRulesRuleConditions struct {
	RuleConditionType *string                                                                 `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	PathConfig        *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	HostConfig        *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetRuleConditionType(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetPathConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetHostConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

type ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) SetValues(v []*string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

type ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) SetValues(v []*string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

type ListForwardingRulesResponseBodyForwardingRulesRuleActions struct {
	Order              *int32                                                                       `json:"Order,omitempty" xml:"Order,omitempty"`
	RuleActionType     *string                                                                      `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	ForwardGroupConfig *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetOrder(v int32) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetRuleActionType(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetForwardGroupConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

type ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig struct {
	ServerGroupTuples []*ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) SetEndpointGroupId(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.EndpointGroupId = &v
	return s
}

type ListForwardingRulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListForwardingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponse) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponse) SetHeaders(v map[string]*string) *ListForwardingRulesResponse {
	s.Headers = v
	return s
}

func (s *ListForwardingRulesResponse) SetBody(v *ListForwardingRulesResponseBody) *ListForwardingRulesResponse {
	s.Body = v
	return s
}

type ListIpSetsRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s ListIpSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpSetsRequest) GoString() string {
	return s.String()
}

func (s *ListIpSetsRequest) SetRegionId(v string) *ListIpSetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpSetsRequest) SetPageNumber(v int32) *ListIpSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIpSetsRequest) SetPageSize(v int32) *ListIpSetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIpSetsRequest) SetAcceleratorId(v string) *ListIpSetsRequest {
	s.AcceleratorId = &v
	return s
}

type ListIpSetsResponseBody struct {
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	IpSets     []*ListIpSetsResponseBodyIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
}

func (s ListIpSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpSetsResponseBody) SetTotalCount(v int32) *ListIpSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpSetsResponseBody) SetPageSize(v int32) *ListIpSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIpSetsResponseBody) SetRequestId(v string) *ListIpSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpSetsResponseBody) SetPageNumber(v int32) *ListIpSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIpSetsResponseBody) SetIpSets(v []*ListIpSetsResponseBodyIpSets) *ListIpSetsResponseBody {
	s.IpSets = v
	return s
}

type ListIpSetsResponseBodyIpSets struct {
	AccelerateRegionId *string   `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	IpVersion          *string   `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Bandwidth          *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	State              *string   `json:"State,omitempty" xml:"State,omitempty"`
	IpSetId            *string   `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	IpAddressList      []*string `json:"IpAddressList,omitempty" xml:"IpAddressList,omitempty" type:"Repeated"`
}

func (s ListIpSetsResponseBodyIpSets) String() string {
	return tea.Prettify(s)
}

func (s ListIpSetsResponseBodyIpSets) GoString() string {
	return s.String()
}

func (s *ListIpSetsResponseBodyIpSets) SetAccelerateRegionId(v string) *ListIpSetsResponseBodyIpSets {
	s.AccelerateRegionId = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpVersion(v string) *ListIpSetsResponseBodyIpSets {
	s.IpVersion = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetBandwidth(v int32) *ListIpSetsResponseBodyIpSets {
	s.Bandwidth = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetState(v string) *ListIpSetsResponseBodyIpSets {
	s.State = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpSetId(v string) *ListIpSetsResponseBodyIpSets {
	s.IpSetId = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpAddressList(v []*string) *ListIpSetsResponseBodyIpSets {
	s.IpAddressList = v
	return s
}

type ListIpSetsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIpSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIpSetsResponse) GoString() string {
	return s.String()
}

func (s *ListIpSetsResponse) SetHeaders(v map[string]*string) *ListIpSetsResponse {
	s.Headers = v
	return s
}

func (s *ListIpSetsResponse) SetBody(v *ListIpSetsResponseBody) *ListIpSetsResponse {
	s.Body = v
	return s
}

type ListListenersRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s ListListenersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) SetRegionId(v string) *ListListenersRequest {
	s.RegionId = &v
	return s
}

func (s *ListListenersRequest) SetPageNumber(v int32) *ListListenersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListListenersRequest) SetPageSize(v int32) *ListListenersRequest {
	s.PageSize = &v
	return s
}

func (s *ListListenersRequest) SetAcceleratorId(v string) *ListListenersRequest {
	s.AcceleratorId = &v
	return s
}

type ListListenersResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Listeners  []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetPageSize(v int32) *ListListenersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetPageNumber(v int32) *ListListenersResponseBody {
	s.PageNumber = &v
	return s
}

type ListListenersResponseBodyListeners struct {
	Certificates   []*ListListenersResponseBodyListenersCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	BackendPorts   []*ListListenersResponseBodyListenersBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	ListenerId     *string                                           `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Description    *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	State          *string                                           `json:"State,omitempty" xml:"State,omitempty"`
	ClientAffinity *string                                           `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	Protocol       *string                                           `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	CreateTime     *int64                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PortRanges     []*ListListenersResponseBodyListenersPortRanges   `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Name           *string                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	ProxyProtocol  *bool                                             `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	AcceleratorId  *string                                           `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) SetCertificates(v []*ListListenersResponseBodyListenersCertificates) *ListListenersResponseBodyListeners {
	s.Certificates = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetBackendPorts(v []*ListListenersResponseBodyListenersBackendPorts) *ListListenersResponseBodyListeners {
	s.BackendPorts = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetDescription(v string) *ListListenersResponseBodyListeners {
	s.Description = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetState(v string) *ListListenersResponseBodyListeners {
	s.State = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetClientAffinity(v string) *ListListenersResponseBodyListeners {
	s.ClientAffinity = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetProtocol(v string) *ListListenersResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCreateTime(v int64) *ListListenersResponseBodyListeners {
	s.CreateTime = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetPortRanges(v []*ListListenersResponseBodyListenersPortRanges) *ListListenersResponseBodyListeners {
	s.PortRanges = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetName(v string) *ListListenersResponseBodyListeners {
	s.Name = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetProxyProtocol(v bool) *ListListenersResponseBodyListeners {
	s.ProxyProtocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetAcceleratorId(v string) *ListListenersResponseBodyListeners {
	s.AcceleratorId = &v
	return s
}

type ListListenersResponseBodyListenersCertificates struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListListenersResponseBodyListenersCertificates) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersCertificates) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersCertificates) SetType(v string) *ListListenersResponseBodyListenersCertificates {
	s.Type = &v
	return s
}

func (s *ListListenersResponseBodyListenersCertificates) SetId(v string) *ListListenersResponseBodyListenersCertificates {
	s.Id = &v
	return s
}

type ListListenersResponseBodyListenersBackendPorts struct {
	FromPort *string `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *string `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListListenersResponseBodyListenersBackendPorts) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersBackendPorts) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersBackendPorts) SetFromPort(v string) *ListListenersResponseBodyListenersBackendPorts {
	s.FromPort = &v
	return s
}

func (s *ListListenersResponseBodyListenersBackendPorts) SetToPort(v string) *ListListenersResponseBodyListenersBackendPorts {
	s.ToPort = &v
	return s
}

type ListListenersResponseBodyListenersPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListListenersResponseBodyListenersPortRanges) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersPortRanges) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersPortRanges) SetFromPort(v int32) *ListListenersResponseBodyListenersPortRanges {
	s.FromPort = &v
	return s
}

func (s *ListListenersResponseBodyListenersPortRanges) SetToPort(v int32) *ListListenersResponseBodyListenersPortRanges {
	s.ToPort = &v
	return s
}

type ListListenersResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListListenersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponse) GoString() string {
	return s.String()
}

func (s *ListListenersResponse) SetHeaders(v map[string]*string) *ListListenersResponse {
	s.Headers = v
	return s
}

func (s *ListListenersResponse) SetBody(v *ListListenersResponseBody) *ListListenersResponse {
	s.Body = v
	return s
}

type RemoveEntriesFromAclRequest struct {
	RegionId    *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AclId       *string                                  `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclEntries  []*RemoveEntriesFromAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	ClientToken *string                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool                                    `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s RemoveEntriesFromAclRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclRequest) SetRegionId(v string) *RemoveEntriesFromAclRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetAclId(v string) *RemoveEntriesFromAclRequest {
	s.AclId = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetAclEntries(v []*RemoveEntriesFromAclRequestAclEntries) *RemoveEntriesFromAclRequest {
	s.AclEntries = v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetClientToken(v string) *RemoveEntriesFromAclRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetDryRun(v bool) *RemoveEntriesFromAclRequest {
	s.DryRun = &v
	return s
}

type RemoveEntriesFromAclRequestAclEntries struct {
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s RemoveEntriesFromAclRequestAclEntries) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclRequestAclEntries) SetEntry(v string) *RemoveEntriesFromAclRequestAclEntries {
	s.Entry = &v
	return s
}

type RemoveEntriesFromAclResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
}

func (s RemoveEntriesFromAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponseBody) SetRequestId(v string) *RemoveEntriesFromAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEntriesFromAclResponseBody) SetAclId(v string) *RemoveEntriesFromAclResponseBody {
	s.AclId = &v
	return s
}

type RemoveEntriesFromAclResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveEntriesFromAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveEntriesFromAclResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclResponse) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponse) SetHeaders(v map[string]*string) *RemoveEntriesFromAclResponse {
	s.Headers = v
	return s
}

func (s *RemoveEntriesFromAclResponse) SetBody(v *RemoveEntriesFromAclResponseBody) *RemoveEntriesFromAclResponse {
	s.Body = v
	return s
}

type ReplaceBandwidthPackageRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BandwidthPackageId       *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	TargetBandwidthPackageId *string `json:"TargetBandwidthPackageId,omitempty" xml:"TargetBandwidthPackageId,omitempty"`
}

func (s ReplaceBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *ReplaceBandwidthPackageRequest) SetRegionId(v string) *ReplaceBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *ReplaceBandwidthPackageRequest) SetBandwidthPackageId(v string) *ReplaceBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ReplaceBandwidthPackageRequest) SetTargetBandwidthPackageId(v string) *ReplaceBandwidthPackageRequest {
	s.TargetBandwidthPackageId = &v
	return s
}

type ReplaceBandwidthPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceBandwidthPackageResponseBody) SetRequestId(v string) *ReplaceBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

type ReplaceBandwidthPackageResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplaceBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *ReplaceBandwidthPackageResponse) SetHeaders(v map[string]*string) *ReplaceBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *ReplaceBandwidthPackageResponse) SetBody(v *ReplaceBandwidthPackageResponseBody) *ReplaceBandwidthPackageResponse {
	s.Body = v
	return s
}

type UpdateAcceleratorRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Spec          *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	AutoPay       *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
}

func (s UpdateAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorRequest) SetRegionId(v string) *UpdateAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetClientToken(v string) *UpdateAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetName(v string) *UpdateAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetDescription(v string) *UpdateAcceleratorRequest {
	s.Description = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetAcceleratorId(v string) *UpdateAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetSpec(v string) *UpdateAcceleratorRequest {
	s.Spec = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetAutoPay(v bool) *UpdateAcceleratorRequest {
	s.AutoPay = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetAutoUseCoupon(v bool) *UpdateAcceleratorRequest {
	s.AutoUseCoupon = &v
	return s
}

type UpdateAcceleratorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorResponseBody) SetRequestId(v string) *UpdateAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAcceleratorResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorResponse) SetHeaders(v map[string]*string) *UpdateAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *UpdateAcceleratorResponse) SetBody(v *UpdateAcceleratorResponseBody) *UpdateAcceleratorResponse {
	s.Body = v
	return s
}

type UpdateAclAttributeRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AclId       *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName     *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s UpdateAclAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeRequest) SetRegionId(v string) *UpdateAclAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetAclId(v string) *UpdateAclAttributeRequest {
	s.AclId = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetAclName(v string) *UpdateAclAttributeRequest {
	s.AclName = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetClientToken(v string) *UpdateAclAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetDryRun(v bool) *UpdateAclAttributeRequest {
	s.DryRun = &v
	return s
}

type UpdateAclAttributeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
}

func (s UpdateAclAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponseBody) SetRequestId(v string) *UpdateAclAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAclAttributeResponseBody) SetAclId(v string) *UpdateAclAttributeResponseBody {
	s.AclId = &v
	return s
}

type UpdateAclAttributeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAclAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAclAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponse) SetHeaders(v map[string]*string) *UpdateAclAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAclAttributeResponse) SetBody(v *UpdateAclAttributeResponseBody) *UpdateAclAttributeResponse {
	s.Body = v
	return s
}

type UpdateBandwidthPackageRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType      *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	AutoPay            *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon      *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
}

func (s UpdateBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackageRequest) SetRegionId(v string) *UpdateBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidthPackageId(v string) *UpdateBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetName(v string) *UpdateBandwidthPackageRequest {
	s.Name = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetDescription(v string) *UpdateBandwidthPackageRequest {
	s.Description = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidth(v int32) *UpdateBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidthType(v string) *UpdateBandwidthPackageRequest {
	s.BandwidthType = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetAutoPay(v bool) *UpdateBandwidthPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetAutoUseCoupon(v bool) *UpdateBandwidthPackageRequest {
	s.AutoUseCoupon = &v
	return s
}

type UpdateBandwidthPackageResponseBody struct {
	BandwidthPackage *string `json:"BandwidthPackage,omitempty" xml:"BandwidthPackage,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackageResponseBody) SetBandwidthPackage(v string) *UpdateBandwidthPackageResponseBody {
	s.BandwidthPackage = &v
	return s
}

func (s *UpdateBandwidthPackageResponseBody) SetDescription(v string) *UpdateBandwidthPackageResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateBandwidthPackageResponseBody) SetRequestId(v string) *UpdateBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBandwidthPackageResponseBody) SetName(v string) *UpdateBandwidthPackageResponseBody {
	s.Name = &v
	return s
}

type UpdateBandwidthPackageResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBandwidthPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackageResponse) SetHeaders(v map[string]*string) *UpdateBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *UpdateBandwidthPackageResponse) SetBody(v *UpdateBandwidthPackageResponseBody) *UpdateBandwidthPackageResponse {
	s.Body = v
	return s
}

type UpdateEndpointGroupRequest struct {
	RegionId                   *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken                *string                                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId            *string                                             `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	Name                       *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Description                *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointGroupRegion        *string                                             `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	TrafficPercentage          *int32                                              `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
	HealthCheckIntervalSeconds *int32                                              `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath            *string                                             `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort            *int32                                              `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol        *string                                             `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	ThresholdCount             *int32                                              `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	EndpointConfigurations     []*UpdateEndpointGroupRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointRequestProtocol    *string                                             `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	PortOverrides              []*UpdateEndpointGroupRequestPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
}

func (s UpdateEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequest) SetRegionId(v string) *UpdateEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetClientToken(v string) *UpdateEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointGroupId(v string) *UpdateEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetName(v string) *UpdateEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetDescription(v string) *UpdateEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointGroupRegion(v string) *UpdateEndpointGroupRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetTrafficPercentage(v int32) *UpdateEndpointGroupRequest {
	s.TrafficPercentage = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckIntervalSeconds(v int32) *UpdateEndpointGroupRequest {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckPath(v string) *UpdateEndpointGroupRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckPort(v int32) *UpdateEndpointGroupRequest {
	s.HealthCheckPort = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckProtocol(v string) *UpdateEndpointGroupRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetThresholdCount(v int32) *UpdateEndpointGroupRequest {
	s.ThresholdCount = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointConfigurations(v []*UpdateEndpointGroupRequestEndpointConfigurations) *UpdateEndpointGroupRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointRequestProtocol(v string) *UpdateEndpointGroupRequest {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetPortOverrides(v []*UpdateEndpointGroupRequestPortOverrides) *UpdateEndpointGroupRequest {
	s.PortOverrides = v
	return s
}

type UpdateEndpointGroupRequestEndpointConfigurations struct {
	Type                       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EnableClientIPPreservation *bool   `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	Weight                     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Endpoint                   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
}

func (s UpdateEndpointGroupRequestEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetType(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetEnableClientIPPreservation(v bool) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetWeight(v int32) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetEndpoint(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Endpoint = &v
	return s
}

type UpdateEndpointGroupRequestPortOverrides struct {
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
}

func (s UpdateEndpointGroupRequestPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupRequestPortOverrides) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequestPortOverrides) SetListenerPort(v int32) *UpdateEndpointGroupRequestPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *UpdateEndpointGroupRequestPortOverrides) SetEndpointPort(v int32) *UpdateEndpointGroupRequestPortOverrides {
	s.EndpointPort = &v
	return s
}

type UpdateEndpointGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupResponseBody) SetRequestId(v string) *UpdateEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEndpointGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupResponse) SetHeaders(v map[string]*string) *UpdateEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateEndpointGroupResponse) SetBody(v *UpdateEndpointGroupResponseBody) *UpdateEndpointGroupResponse {
	s.Body = v
	return s
}

type UpdateEndpointGroupAttributeRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateEndpointGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupAttributeRequest) SetRegionId(v string) *UpdateEndpointGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetClientToken(v string) *UpdateEndpointGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetEndpointGroupId(v string) *UpdateEndpointGroupAttributeRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetName(v string) *UpdateEndpointGroupAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetDescription(v string) *UpdateEndpointGroupAttributeRequest {
	s.Description = &v
	return s
}

type UpdateEndpointGroupAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEndpointGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupAttributeResponseBody) SetRequestId(v string) *UpdateEndpointGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEndpointGroupAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEndpointGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEndpointGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateEndpointGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEndpointGroupAttributeResponse) SetBody(v *UpdateEndpointGroupAttributeResponseBody) *UpdateEndpointGroupAttributeResponse {
	s.Body = v
	return s
}

type UpdateForwardingRulesRequest struct {
	RegionId        *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken     *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AcceleratorId   *string                                        `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId      *string                                        `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	ForwardingRules []*UpdateForwardingRulesRequestForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
}

func (s UpdateForwardingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequest) SetRegionId(v string) *UpdateForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetClientToken(v string) *UpdateForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetAcceleratorId(v string) *UpdateForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetListenerId(v string) *UpdateForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetForwardingRules(v []*UpdateForwardingRulesRequestForwardingRules) *UpdateForwardingRulesRequest {
	s.ForwardingRules = v
	return s
}

type UpdateForwardingRulesRequestForwardingRules struct {
	Priority           *int32                                                       `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleConditions     []*UpdateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	RuleActions        []*UpdateForwardingRulesRequestForwardingRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	ForwardingRuleId   *string                                                      `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	ForwardingRuleName *string                                                      `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRules) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetPriority(v int32) *UpdateForwardingRulesRequestForwardingRules {
	s.Priority = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetRuleConditions(v []*UpdateForwardingRulesRequestForwardingRulesRuleConditions) *UpdateForwardingRulesRequestForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetRuleActions(v []*UpdateForwardingRulesRequestForwardingRulesRuleActions) *UpdateForwardingRulesRequestForwardingRules {
	s.RuleActions = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetForwardingRuleId(v string) *UpdateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetForwardingRuleName(v string) *UpdateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

type UpdateForwardingRulesRequestForwardingRulesRuleConditions struct {
	RuleConditionType *string                                                              `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	PathConfig        *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	HostConfig        *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionType(v string) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetPathConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetHostConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

type UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) SetValues(v []*string) *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

type UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) SetValues(v []*string) *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

type UpdateForwardingRulesRequestForwardingRulesRuleActions struct {
	Order              *int32                                                                    `json:"Order,omitempty" xml:"Order,omitempty"`
	RuleActionType     *string                                                                   `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	ForwardGroupConfig *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetOrder(v int32) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionType(v string) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetForwardGroupConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

type UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig struct {
	ServerGroupTuples []*UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) SetEndpointGroupId(v string) *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.EndpointGroupId = &v
	return s
}

type UpdateForwardingRulesResponseBody struct {
	ForwardingRules []*UpdateForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateForwardingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesResponseBody) SetForwardingRules(v []*UpdateForwardingRulesResponseBodyForwardingRules) *UpdateForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

func (s *UpdateForwardingRulesResponseBody) SetRequestId(v string) *UpdateForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateForwardingRulesResponseBodyForwardingRules struct {
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
}

func (s UpdateForwardingRulesResponseBodyForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesResponseBodyForwardingRules) SetForwardingRuleId(v string) *UpdateForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

type UpdateForwardingRulesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateForwardingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesResponse) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesResponse) SetHeaders(v map[string]*string) *UpdateForwardingRulesResponse {
	s.Headers = v
	return s
}

func (s *UpdateForwardingRulesResponse) SetBody(v *UpdateForwardingRulesResponseBody) *UpdateForwardingRulesResponse {
	s.Body = v
	return s
}

type UpdateIpSetRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId     *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	Bandwidth   *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s UpdateIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpSetRequest) SetRegionId(v string) *UpdateIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpSetRequest) SetClientToken(v string) *UpdateIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpSetRequest) SetIpSetId(v string) *UpdateIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *UpdateIpSetRequest) SetBandwidth(v int32) *UpdateIpSetRequest {
	s.Bandwidth = &v
	return s
}

type UpdateIpSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpSetResponseBody) SetRequestId(v string) *UpdateIpSetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpSetResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIpSetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpSetResponse) SetHeaders(v map[string]*string) *UpdateIpSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpSetResponse) SetBody(v *UpdateIpSetResponseBody) *UpdateIpSetResponse {
	s.Body = v
	return s
}

type UpdateIpSetsRequest struct {
	RegionId *string                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IpSets   []*UpdateIpSetsRequestIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
}

func (s UpdateIpSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetsRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsRequest) SetRegionId(v string) *UpdateIpSetsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpSetsRequest) SetIpSets(v []*UpdateIpSetsRequestIpSets) *UpdateIpSetsRequest {
	s.IpSets = v
	return s
}

type UpdateIpSetsRequestIpSets struct {
	Bandwidth *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	IpSetId   *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
}

func (s UpdateIpSetsRequestIpSets) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetsRequestIpSets) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsRequestIpSets) SetBandwidth(v int32) *UpdateIpSetsRequestIpSets {
	s.Bandwidth = &v
	return s
}

func (s *UpdateIpSetsRequestIpSets) SetIpSetId(v string) *UpdateIpSetsRequestIpSets {
	s.IpSetId = &v
	return s
}

type UpdateIpSetsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsResponseBody) SetRequestId(v string) *UpdateIpSetsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIpSetsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIpSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetsResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsResponse) SetHeaders(v map[string]*string) *UpdateIpSetsResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpSetsResponse) SetBody(v *UpdateIpSetsResponseBody) *UpdateIpSetsResponse {
	s.Body = v
	return s
}

type UpdateListenerRequest struct {
	RegionId       *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken    *string                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name           *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	ClientAffinity *string                              `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	Protocol       *string                              `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ListenerId     *string                              `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	ProxyProtocol  *string                              `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	PortRanges     []*UpdateListenerRequestPortRanges   `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Certificates   []*UpdateListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	BackendPorts   []*UpdateListenerRequestBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
}

func (s UpdateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequest) SetRegionId(v string) *UpdateListenerRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateListenerRequest) SetClientToken(v string) *UpdateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerRequest) SetName(v string) *UpdateListenerRequest {
	s.Name = &v
	return s
}

func (s *UpdateListenerRequest) SetDescription(v string) *UpdateListenerRequest {
	s.Description = &v
	return s
}

func (s *UpdateListenerRequest) SetClientAffinity(v string) *UpdateListenerRequest {
	s.ClientAffinity = &v
	return s
}

func (s *UpdateListenerRequest) SetProtocol(v string) *UpdateListenerRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateListenerRequest) SetListenerId(v string) *UpdateListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerRequest) SetProxyProtocol(v string) *UpdateListenerRequest {
	s.ProxyProtocol = &v
	return s
}

func (s *UpdateListenerRequest) SetPortRanges(v []*UpdateListenerRequestPortRanges) *UpdateListenerRequest {
	s.PortRanges = v
	return s
}

func (s *UpdateListenerRequest) SetCertificates(v []*UpdateListenerRequestCertificates) *UpdateListenerRequest {
	s.Certificates = v
	return s
}

func (s *UpdateListenerRequest) SetBackendPorts(v []*UpdateListenerRequestBackendPorts) *UpdateListenerRequest {
	s.BackendPorts = v
	return s
}

type UpdateListenerRequestPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateListenerRequestPortRanges) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerRequestPortRanges) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequestPortRanges) SetFromPort(v int32) *UpdateListenerRequestPortRanges {
	s.FromPort = &v
	return s
}

func (s *UpdateListenerRequestPortRanges) SetToPort(v int32) *UpdateListenerRequestPortRanges {
	s.ToPort = &v
	return s
}

type UpdateListenerRequestCertificates struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequestCertificates) SetId(v string) *UpdateListenerRequestCertificates {
	s.Id = &v
	return s
}

type UpdateListenerRequestBackendPorts struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateListenerRequestBackendPorts) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerRequestBackendPorts) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequestBackendPorts) SetFromPort(v int32) *UpdateListenerRequestBackendPorts {
	s.FromPort = &v
	return s
}

func (s *UpdateListenerRequestBackendPorts) SetToPort(v int32) *UpdateListenerRequestBackendPorts {
	s.ToPort = &v
	return s
}

type UpdateListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerResponseBody) SetRequestId(v string) *UpdateListenerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateListenerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerResponse) SetHeaders(v map[string]*string) *UpdateListenerResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerResponse) SetBody(v *UpdateListenerResponseBody) *UpdateListenerResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ga"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEntriesToAclWithOptions(request *AddEntriesToAclRequest, runtime *util.RuntimeOptions) (_result *AddEntriesToAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddEntriesToAcl"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEntriesToAcl(request *AddEntriesToAclRequest) (_result *AddEntriesToAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.AddEntriesToAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateAclsWithListenerWithOptions(request *AssociateAclsWithListenerRequest, runtime *util.RuntimeOptions) (_result *AssociateAclsWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateAclsWithListener"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateAclsWithListener(request *AssociateAclsWithListenerRequest) (_result *AssociateAclsWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.AssociateAclsWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachDdosToAcceleratorWithOptions(request *AttachDdosToAcceleratorRequest, runtime *util.RuntimeOptions) (_result *AttachDdosToAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachDdosToAcceleratorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachDdosToAccelerator"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachDdosToAccelerator(request *AttachDdosToAcceleratorRequest) (_result *AttachDdosToAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachDdosToAcceleratorResponse{}
	_body, _err := client.AttachDdosToAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachLogStoreToEndpointGroupWithOptions(request *AttachLogStoreToEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *AttachLogStoreToEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachLogStoreToEndpointGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachLogStoreToEndpointGroup"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachLogStoreToEndpointGroup(request *AttachLogStoreToEndpointGroupRequest) (_result *AttachLogStoreToEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachLogStoreToEndpointGroupResponse{}
	_body, _err := client.AttachLogStoreToEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BandwidthPackageAddAcceleratorWithOptions(request *BandwidthPackageAddAcceleratorRequest, runtime *util.RuntimeOptions) (_result *BandwidthPackageAddAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BandwidthPackageAddAcceleratorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BandwidthPackageAddAccelerator"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BandwidthPackageAddAccelerator(request *BandwidthPackageAddAcceleratorRequest) (_result *BandwidthPackageAddAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BandwidthPackageAddAcceleratorResponse{}
	_body, _err := client.BandwidthPackageAddAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BandwidthPackageRemoveAcceleratorWithOptions(request *BandwidthPackageRemoveAcceleratorRequest, runtime *util.RuntimeOptions) (_result *BandwidthPackageRemoveAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BandwidthPackageRemoveAcceleratorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BandwidthPackageRemoveAccelerator"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BandwidthPackageRemoveAccelerator(request *BandwidthPackageRemoveAcceleratorRequest) (_result *BandwidthPackageRemoveAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BandwidthPackageRemoveAcceleratorResponse{}
	_body, _err := client.BandwidthPackageRemoveAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigEndpointProbeWithOptions(request *ConfigEndpointProbeRequest, runtime *util.RuntimeOptions) (_result *ConfigEndpointProbeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigEndpointProbeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigEndpointProbe"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigEndpointProbe(request *ConfigEndpointProbeRequest) (_result *ConfigEndpointProbeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigEndpointProbeResponse{}
	_body, _err := client.ConfigEndpointProbeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAcceleratorWithOptions(request *CreateAcceleratorRequest, runtime *util.RuntimeOptions) (_result *CreateAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAcceleratorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAccelerator"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccelerator(request *CreateAcceleratorRequest) (_result *CreateAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAcceleratorResponse{}
	_body, _err := client.CreateAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAclWithOptions(request *CreateAclRequest, runtime *util.RuntimeOptions) (_result *CreateAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAcl"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAcl(request *CreateAclRequest) (_result *CreateAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAclResponse{}
	_body, _err := client.CreateAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBandwidthPackageWithOptions(request *CreateBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *CreateBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBandwidthPackage"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBandwidthPackage(request *CreateBandwidthPackageRequest) (_result *CreateBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBandwidthPackageResponse{}
	_body, _err := client.CreateBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEndpointGroupWithOptions(request *CreateEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *CreateEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEndpointGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEndpointGroup"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEndpointGroup(request *CreateEndpointGroupRequest) (_result *CreateEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEndpointGroupResponse{}
	_body, _err := client.CreateEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateForwardingRulesWithOptions(request *CreateForwardingRulesRequest, runtime *util.RuntimeOptions) (_result *CreateForwardingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateForwardingRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateForwardingRules"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateForwardingRules(request *CreateForwardingRulesRequest) (_result *CreateForwardingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateForwardingRulesResponse{}
	_body, _err := client.CreateForwardingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIpSetsWithOptions(request *CreateIpSetsRequest, runtime *util.RuntimeOptions) (_result *CreateIpSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIpSetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIpSets"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIpSets(request *CreateIpSetsRequest) (_result *CreateIpSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpSetsResponse{}
	_body, _err := client.CreateIpSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateListenerWithOptions(request *CreateListenerRequest, runtime *util.RuntimeOptions) (_result *CreateListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateListener"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateListener(request *CreateListenerRequest) (_result *CreateListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateListenerResponse{}
	_body, _err := client.CreateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAcceleratorWithOptions(request *DeleteAcceleratorRequest, runtime *util.RuntimeOptions) (_result *DeleteAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAcceleratorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAccelerator"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccelerator(request *DeleteAcceleratorRequest) (_result *DeleteAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAcceleratorResponse{}
	_body, _err := client.DeleteAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAclWithOptions(request *DeleteAclRequest, runtime *util.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAcl"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAcl(request *DeleteAclRequest) (_result *DeleteAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAclResponse{}
	_body, _err := client.DeleteAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBandwidthPackageWithOptions(request *DeleteBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *DeleteBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBandwidthPackage"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBandwidthPackage(request *DeleteBandwidthPackageRequest) (_result *DeleteBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBandwidthPackageResponse{}
	_body, _err := client.DeleteBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEndpointGroupWithOptions(request *DeleteEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEndpointGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEndpointGroup"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEndpointGroup(request *DeleteEndpointGroupRequest) (_result *DeleteEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEndpointGroupResponse{}
	_body, _err := client.DeleteEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteForwardingRulesWithOptions(request *DeleteForwardingRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteForwardingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteForwardingRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteForwardingRules"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteForwardingRules(request *DeleteForwardingRulesRequest) (_result *DeleteForwardingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteForwardingRulesResponse{}
	_body, _err := client.DeleteForwardingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIpSetWithOptions(request *DeleteIpSetRequest, runtime *util.RuntimeOptions) (_result *DeleteIpSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteIpSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteIpSet"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIpSet(request *DeleteIpSetRequest) (_result *DeleteIpSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpSetResponse{}
	_body, _err := client.DeleteIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIpSetsWithOptions(request *DeleteIpSetsRequest, runtime *util.RuntimeOptions) (_result *DeleteIpSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteIpSetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteIpSets"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIpSets(request *DeleteIpSetsRequest) (_result *DeleteIpSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpSetsResponse{}
	_body, _err := client.DeleteIpSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteListenerWithOptions(request *DeleteListenerRequest, runtime *util.RuntimeOptions) (_result *DeleteListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteListener"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteListener(request *DeleteListenerRequest) (_result *DeleteListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteListenerResponse{}
	_body, _err := client.DeleteListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAcceleratorWithOptions(request *DescribeAcceleratorRequest, runtime *util.RuntimeOptions) (_result *DescribeAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAcceleratorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccelerator"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccelerator(request *DescribeAcceleratorRequest) (_result *DescribeAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAcceleratorResponse{}
	_body, _err := client.DescribeAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBandwidthPackageWithOptions(request *DescribeBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *DescribeBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBandwidthPackage"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBandwidthPackage(request *DescribeBandwidthPackageRequest) (_result *DescribeBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBandwidthPackageResponse{}
	_body, _err := client.DescribeBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEndpointGroupWithOptions(request *DescribeEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEndpointGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEndpointGroup"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEndpointGroup(request *DescribeEndpointGroupRequest) (_result *DescribeEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndpointGroupResponse{}
	_body, _err := client.DescribeEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpSetWithOptions(request *DescribeIpSetRequest, runtime *util.RuntimeOptions) (_result *DescribeIpSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpSet"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpSet(request *DescribeIpSetRequest) (_result *DescribeIpSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpSetResponse{}
	_body, _err := client.DescribeIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeListenerWithOptions(request *DescribeListenerRequest, runtime *util.RuntimeOptions) (_result *DescribeListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeListener"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeListener(request *DescribeListenerRequest) (_result *DescribeListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeListenerResponse{}
	_body, _err := client.DescribeListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachDdosFromAcceleratorWithOptions(request *DetachDdosFromAcceleratorRequest, runtime *util.RuntimeOptions) (_result *DetachDdosFromAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachDdosFromAcceleratorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachDdosFromAccelerator"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachDdosFromAccelerator(request *DetachDdosFromAcceleratorRequest) (_result *DetachDdosFromAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachDdosFromAcceleratorResponse{}
	_body, _err := client.DetachDdosFromAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachLogStoreFromEndpointGroupWithOptions(request *DetachLogStoreFromEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *DetachLogStoreFromEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachLogStoreFromEndpointGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachLogStoreFromEndpointGroup"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachLogStoreFromEndpointGroup(request *DetachLogStoreFromEndpointGroupRequest) (_result *DetachLogStoreFromEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachLogStoreFromEndpointGroupResponse{}
	_body, _err := client.DetachLogStoreFromEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissociateAclsFromListenerWithOptions(request *DissociateAclsFromListenerRequest, runtime *util.RuntimeOptions) (_result *DissociateAclsFromListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DissociateAclsFromListener"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateAclsFromListener(request *DissociateAclsFromListenerRequest) (_result *DissociateAclsFromListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.DissociateAclsFromListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAclWithOptions(request *GetAclRequest, runtime *util.RuntimeOptions) (_result *GetAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAcl"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAcl(request *GetAclRequest) (_result *GetAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAclResponse{}
	_body, _err := client.GetAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccelerateAreasWithOptions(request *ListAccelerateAreasRequest, runtime *util.RuntimeOptions) (_result *ListAccelerateAreasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAccelerateAreasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAccelerateAreas"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccelerateAreas(request *ListAccelerateAreasRequest) (_result *ListAccelerateAreasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccelerateAreasResponse{}
	_body, _err := client.ListAccelerateAreasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAcceleratorsWithOptions(request *ListAcceleratorsRequest, runtime *util.RuntimeOptions) (_result *ListAcceleratorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAcceleratorsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAccelerators"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccelerators(request *ListAcceleratorsRequest) (_result *ListAcceleratorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAcceleratorsResponse{}
	_body, _err := client.ListAcceleratorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAclsWithOptions(request *ListAclsRequest, runtime *util.RuntimeOptions) (_result *ListAclsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAclsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAcls"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAcls(request *ListAclsRequest) (_result *ListAclsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAclsResponse{}
	_body, _err := client.ListAclsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableAccelerateAreasWithOptions(request *ListAvailableAccelerateAreasRequest, runtime *util.RuntimeOptions) (_result *ListAvailableAccelerateAreasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAvailableAccelerateAreasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAvailableAccelerateAreas"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableAccelerateAreas(request *ListAvailableAccelerateAreasRequest) (_result *ListAvailableAccelerateAreasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableAccelerateAreasResponse{}
	_body, _err := client.ListAvailableAccelerateAreasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableBusiRegionsWithOptions(request *ListAvailableBusiRegionsRequest, runtime *util.RuntimeOptions) (_result *ListAvailableBusiRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAvailableBusiRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAvailableBusiRegions"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableBusiRegions(request *ListAvailableBusiRegionsRequest) (_result *ListAvailableBusiRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableBusiRegionsResponse{}
	_body, _err := client.ListAvailableBusiRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBandwidthackagesWithOptions(request *ListBandwidthackagesRequest, runtime *util.RuntimeOptions) (_result *ListBandwidthackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBandwidthackagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBandwidthackages"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBandwidthackages(request *ListBandwidthackagesRequest) (_result *ListBandwidthackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBandwidthackagesResponse{}
	_body, _err := client.ListBandwidthackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBandwidthPackagesWithOptions(request *ListBandwidthPackagesRequest, runtime *util.RuntimeOptions) (_result *ListBandwidthPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBandwidthPackagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBandwidthPackages"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBandwidthPackages(request *ListBandwidthPackagesRequest) (_result *ListBandwidthPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBandwidthPackagesResponse{}
	_body, _err := client.ListBandwidthPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBusiRegionsWithOptions(request *ListBusiRegionsRequest, runtime *util.RuntimeOptions) (_result *ListBusiRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBusiRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBusiRegions"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBusiRegions(request *ListBusiRegionsRequest) (_result *ListBusiRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBusiRegionsResponse{}
	_body, _err := client.ListBusiRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEndpointGroupsWithOptions(request *ListEndpointGroupsRequest, runtime *util.RuntimeOptions) (_result *ListEndpointGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListEndpointGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEndpointGroups"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEndpointGroups(request *ListEndpointGroupsRequest) (_result *ListEndpointGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEndpointGroupsResponse{}
	_body, _err := client.ListEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListForwardingRulesWithOptions(request *ListForwardingRulesRequest, runtime *util.RuntimeOptions) (_result *ListForwardingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListForwardingRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListForwardingRules"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListForwardingRules(request *ListForwardingRulesRequest) (_result *ListForwardingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListForwardingRulesResponse{}
	_body, _err := client.ListForwardingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIpSetsWithOptions(request *ListIpSetsRequest, runtime *util.RuntimeOptions) (_result *ListIpSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIpSetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIpSets"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIpSets(request *ListIpSetsRequest) (_result *ListIpSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIpSetsResponse{}
	_body, _err := client.ListIpSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenersWithOptions(request *ListListenersRequest, runtime *util.RuntimeOptions) (_result *ListListenersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListListenersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListListeners"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListeners(request *ListListenersRequest) (_result *ListListenersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenersResponse{}
	_body, _err := client.ListListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveEntriesFromAclWithOptions(request *RemoveEntriesFromAclRequest, runtime *util.RuntimeOptions) (_result *RemoveEntriesFromAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveEntriesFromAcl"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveEntriesFromAcl(request *RemoveEntriesFromAclRequest) (_result *RemoveEntriesFromAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.RemoveEntriesFromAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceBandwidthPackageWithOptions(request *ReplaceBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *ReplaceBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReplaceBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReplaceBandwidthPackage"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceBandwidthPackage(request *ReplaceBandwidthPackageRequest) (_result *ReplaceBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplaceBandwidthPackageResponse{}
	_body, _err := client.ReplaceBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAcceleratorWithOptions(request *UpdateAcceleratorRequest, runtime *util.RuntimeOptions) (_result *UpdateAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAcceleratorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAccelerator"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAccelerator(request *UpdateAcceleratorRequest) (_result *UpdateAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAcceleratorResponse{}
	_body, _err := client.UpdateAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAclAttributeWithOptions(request *UpdateAclAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateAclAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAclAttribute"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAclAttribute(request *UpdateAclAttributeRequest) (_result *UpdateAclAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.UpdateAclAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBandwidthPackageWithOptions(request *UpdateBandwidthPackageRequest, runtime *util.RuntimeOptions) (_result *UpdateBandwidthPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateBandwidthPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateBandwidthPackage"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBandwidthPackage(request *UpdateBandwidthPackageRequest) (_result *UpdateBandwidthPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBandwidthPackageResponse{}
	_body, _err := client.UpdateBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEndpointGroupWithOptions(request *UpdateEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEndpointGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEndpointGroup"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEndpointGroup(request *UpdateEndpointGroupRequest) (_result *UpdateEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEndpointGroupResponse{}
	_body, _err := client.UpdateEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEndpointGroupAttributeWithOptions(request *UpdateEndpointGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateEndpointGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEndpointGroupAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEndpointGroupAttribute"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEndpointGroupAttribute(request *UpdateEndpointGroupAttributeRequest) (_result *UpdateEndpointGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEndpointGroupAttributeResponse{}
	_body, _err := client.UpdateEndpointGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateForwardingRulesWithOptions(request *UpdateForwardingRulesRequest, runtime *util.RuntimeOptions) (_result *UpdateForwardingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateForwardingRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateForwardingRules"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateForwardingRules(request *UpdateForwardingRulesRequest) (_result *UpdateForwardingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateForwardingRulesResponse{}
	_body, _err := client.UpdateForwardingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIpSetWithOptions(request *UpdateIpSetRequest, runtime *util.RuntimeOptions) (_result *UpdateIpSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateIpSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateIpSet"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIpSet(request *UpdateIpSetRequest) (_result *UpdateIpSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpSetResponse{}
	_body, _err := client.UpdateIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIpSetsWithOptions(request *UpdateIpSetsRequest, runtime *util.RuntimeOptions) (_result *UpdateIpSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateIpSetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateIpSets"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIpSets(request *UpdateIpSetsRequest) (_result *UpdateIpSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIpSetsResponse{}
	_body, _err := client.UpdateIpSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateListenerWithOptions(request *UpdateListenerRequest, runtime *util.RuntimeOptions) (_result *UpdateListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateListener"), tea.String("2019-11-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateListener(request *UpdateListenerRequest) (_result *UpdateListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateListenerResponse{}
	_body, _err := client.UpdateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
