// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddEntriesToAclRequest struct {
	AclEntries  []*AddEntriesToAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	AclId       *string                             `json:"AclId,omitempty" xml:"AclId,omitempty"`
	ClientToken *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool                               `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId    *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddEntriesToAclRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclRequest) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequest) SetAclEntries(v []*AddEntriesToAclRequestAclEntries) *AddEntriesToAclRequest {
	s.AclEntries = v
	return s
}

func (s *AddEntriesToAclRequest) SetAclId(v string) *AddEntriesToAclRequest {
	s.AclId = &v
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

func (s *AddEntriesToAclRequest) SetRegionId(v string) *AddEntriesToAclRequest {
	s.RegionId = &v
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
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEntriesToAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclResponseBody) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponseBody) SetAclId(v string) *AddEntriesToAclResponseBody {
	s.AclId = &v
	return s
}

func (s *AddEntriesToAclResponseBody) SetRequestId(v string) *AddEntriesToAclResponseBody {
	s.RequestId = &v
	return s
}

type AddEntriesToAclResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddEntriesToAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddEntriesToAclResponse) SetStatusCode(v int32) *AddEntriesToAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEntriesToAclResponse) SetBody(v *AddEntriesToAclResponseBody) *AddEntriesToAclResponse {
	s.Body = v
	return s
}

type AssociateAclsWithListenerRequest struct {
	AclIds      []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	AclType     *string   `json:"AclType,omitempty" xml:"AclType,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ListenerId  *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateAclsWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerRequest) SetAclIds(v []*string) *AssociateAclsWithListenerRequest {
	s.AclIds = v
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

func (s *AssociateAclsWithListenerRequest) SetListenerId(v string) *AssociateAclsWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetRegionId(v string) *AssociateAclsWithListenerRequest {
	s.RegionId = &v
	return s
}

type AssociateAclsWithListenerResponseBody struct {
	AclIds     []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	ListenerId *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAclsWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponseBody) SetAclIds(v []*string) *AssociateAclsWithListenerResponseBody {
	s.AclIds = v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetListenerId(v string) *AssociateAclsWithListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetRequestId(v string) *AssociateAclsWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type AssociateAclsWithListenerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateAclsWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AssociateAclsWithListenerResponse) SetStatusCode(v int32) *AssociateAclsWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateAclsWithListenerResponse) SetBody(v *AssociateAclsWithListenerResponseBody) *AssociateAclsWithListenerResponse {
	s.Body = v
	return s
}

type AssociateAdditionalCertificatesWithListenerRequest struct {
	AcceleratorId *string                                                           `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Certificates  []*AssociateAdditionalCertificatesWithListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	ClientToken   *string                                                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ListenerId    *string                                                           `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId      *string                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetAcceleratorId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetCertificates(v []*AssociateAdditionalCertificatesWithListenerRequestCertificates) *AssociateAdditionalCertificatesWithListenerRequest {
	s.Certificates = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetRegionId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.RegionId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerRequestCertificates struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) SetDomain(v string) *AssociateAdditionalCertificatesWithListenerRequestCertificates {
	s.Domain = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) SetId(v string) *AssociateAdditionalCertificatesWithListenerRequestCertificates {
	s.Id = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerResponseBody struct {
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetRequestId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateAdditionalCertificatesWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateAdditionalCertificatesWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponse) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetHeaders(v map[string]*string) *AssociateAdditionalCertificatesWithListenerResponse {
	s.Headers = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetStatusCode(v int32) *AssociateAdditionalCertificatesWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetBody(v *AssociateAdditionalCertificatesWithListenerResponseBody) *AssociateAdditionalCertificatesWithListenerResponse {
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
	GaId      *string `json:"GaId,omitempty" xml:"GaId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *AttachDdosToAcceleratorResponseBody) SetGaId(v string) *AttachDdosToAcceleratorResponseBody {
	s.GaId = &v
	return s
}

func (s *AttachDdosToAcceleratorResponseBody) SetRequestId(v string) *AttachDdosToAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type AttachDdosToAcceleratorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachDdosToAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachDdosToAcceleratorResponse) SetStatusCode(v int32) *AttachDdosToAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDdosToAcceleratorResponse) SetBody(v *AttachDdosToAcceleratorResponseBody) *AttachDdosToAcceleratorResponse {
	s.Body = v
	return s
}

type AttachLogStoreToEndpointGroupRequest struct {
	AcceleratorId    *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken      *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	ListenerId       *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlsLogStoreName  *string   `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	SlsProjectName   *string   `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	SlsRegionId      *string   `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
}

func (s AttachLogStoreToEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachLogStoreToEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachLogStoreToEndpointGroupRequest) SetAcceleratorId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetClientToken(v string) *AttachLogStoreToEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetEndpointGroupIds(v []*string) *AttachLogStoreToEndpointGroupRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetListenerId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetRegionId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsLogStoreName(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsLogStoreName = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsProjectName(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsProjectName = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupRequest) SetSlsRegionId(v string) *AttachLogStoreToEndpointGroupRequest {
	s.SlsRegionId = &v
	return s
}

type AttachLogStoreToEndpointGroupResponseBody struct {
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachLogStoreToEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachLogStoreToEndpointGroupResponse) SetStatusCode(v int32) *AttachLogStoreToEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupResponse) SetBody(v *AttachLogStoreToEndpointGroupResponseBody) *AttachLogStoreToEndpointGroupResponse {
	s.Body = v
	return s
}

type BandwidthPackageAddAcceleratorRequest struct {
	AcceleratorId      *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BandwidthPackageAddAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageAddAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *BandwidthPackageAddAcceleratorRequest) SetAcceleratorId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorRequest) SetBandwidthPackageId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorRequest) SetRegionId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.RegionId = &v
	return s
}

type BandwidthPackageAddAcceleratorResponseBody struct {
	Accelerators       []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	BandwidthPackageId *string   `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BandwidthPackageAddAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageAddAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetAccelerators(v []*string) *BandwidthPackageAddAcceleratorResponseBody {
	s.Accelerators = v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetBandwidthPackageId(v string) *BandwidthPackageAddAcceleratorResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponseBody) SetRequestId(v string) *BandwidthPackageAddAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type BandwidthPackageAddAcceleratorResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BandwidthPackageAddAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BandwidthPackageAddAcceleratorResponse) SetStatusCode(v int32) *BandwidthPackageAddAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorResponse) SetBody(v *BandwidthPackageAddAcceleratorResponseBody) *BandwidthPackageAddAcceleratorResponse {
	s.Body = v
	return s
}

type BandwidthPackageRemoveAcceleratorRequest struct {
	AcceleratorId      *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BandwidthPackageRemoveAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageRemoveAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetAcceleratorId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetBandwidthPackageId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetRegionId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.RegionId = &v
	return s
}

type BandwidthPackageRemoveAcceleratorResponseBody struct {
	Accelerators       []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	BandwidthPackageId *string   `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BandwidthPackageRemoveAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BandwidthPackageRemoveAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetAccelerators(v []*string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.Accelerators = v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetBandwidthPackageId(v string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponseBody) SetRequestId(v string) *BandwidthPackageRemoveAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type BandwidthPackageRemoveAcceleratorResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BandwidthPackageRemoveAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BandwidthPackageRemoveAcceleratorResponse) SetStatusCode(v int32) *BandwidthPackageRemoveAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorResponse) SetBody(v *BandwidthPackageRemoveAcceleratorResponseBody) *BandwidthPackageRemoveAcceleratorResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetClientToken(v string) *ChangeResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type ConfigEndpointProbeRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Enable          *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointType    *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	ProbePort       *string `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
	ProbeProtocol   *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigEndpointProbeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigEndpointProbeRequest) GoString() string {
	return s.String()
}

func (s *ConfigEndpointProbeRequest) SetClientToken(v string) *ConfigEndpointProbeRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEnable(v string) *ConfigEndpointProbeRequest {
	s.Enable = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEndpoint(v string) *ConfigEndpointProbeRequest {
	s.Endpoint = &v
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

func (s *ConfigEndpointProbeRequest) SetProbePort(v string) *ConfigEndpointProbeRequest {
	s.ProbePort = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetProbeProtocol(v string) *ConfigEndpointProbeRequest {
	s.ProbeProtocol = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetRegionId(v string) *ConfigEndpointProbeRequest {
	s.RegionId = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigEndpointProbeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigEndpointProbeResponse) SetStatusCode(v int32) *ConfigEndpointProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigEndpointProbeResponse) SetBody(v *ConfigEndpointProbeResponseBody) *ConfigEndpointProbeResponse {
	s.Body = v
	return s
}

type CreateAcceleratorRequest struct {
	AutoPay              *bool                                `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool                                `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration    *int32                               `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoUseCoupon        *string                              `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	BandwidthBillingType *string                              `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	ClientToken          *string                              `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Duration             *int32                               `json:"Duration,omitempty" xml:"Duration,omitempty"`
	IpSetConfig          *CreateAcceleratorRequestIpSetConfig `json:"IpSetConfig,omitempty" xml:"IpSetConfig,omitempty" type:"Struct"`
	Name                 *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	PricingCycle         *string                              `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId             *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Spec                 *string                              `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorRequest) SetAutoPay(v bool) *CreateAcceleratorRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAcceleratorRequest) SetAutoRenew(v bool) *CreateAcceleratorRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAcceleratorRequest) SetAutoRenewDuration(v int32) *CreateAcceleratorRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateAcceleratorRequest) SetAutoUseCoupon(v string) *CreateAcceleratorRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateAcceleratorRequest) SetBandwidthBillingType(v string) *CreateAcceleratorRequest {
	s.BandwidthBillingType = &v
	return s
}

func (s *CreateAcceleratorRequest) SetClientToken(v string) *CreateAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAcceleratorRequest) SetDuration(v int32) *CreateAcceleratorRequest {
	s.Duration = &v
	return s
}

func (s *CreateAcceleratorRequest) SetIpSetConfig(v *CreateAcceleratorRequestIpSetConfig) *CreateAcceleratorRequest {
	s.IpSetConfig = v
	return s
}

func (s *CreateAcceleratorRequest) SetName(v string) *CreateAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *CreateAcceleratorRequest) SetPricingCycle(v string) *CreateAcceleratorRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateAcceleratorRequest) SetRegionId(v string) *CreateAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAcceleratorRequest) SetResourceGroupId(v string) *CreateAcceleratorRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAcceleratorRequest) SetSpec(v string) *CreateAcceleratorRequest {
	s.Spec = &v
	return s
}

type CreateAcceleratorRequestIpSetConfig struct {
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s CreateAcceleratorRequestIpSetConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAcceleratorRequestIpSetConfig) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorRequestIpSetConfig) SetAccessMode(v string) *CreateAcceleratorRequestIpSetConfig {
	s.AccessMode = &v
	return s
}

type CreateAcceleratorResponseBody struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorResponseBody) SetAcceleratorId(v string) *CreateAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateAcceleratorResponseBody) SetOrderId(v string) *CreateAcceleratorResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAcceleratorResponseBody) SetRequestId(v string) *CreateAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type CreateAcceleratorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAcceleratorResponse) SetStatusCode(v int32) *CreateAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAcceleratorResponse) SetBody(v *CreateAcceleratorResponseBody) *CreateAcceleratorResponse {
	s.Body = v
	return s
}

type CreateAclRequest struct {
	AclEntries       []*CreateAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	AclName          *string                       `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AddressIPVersion *string                       `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	ClientToken      *string                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool                         `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId         *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId  *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateAclRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRequest) SetAclEntries(v []*CreateAclRequestAclEntries) *CreateAclRequest {
	s.AclEntries = v
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

func (s *CreateAclRequest) SetClientToken(v string) *CreateAclRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAclRequest) SetDryRun(v bool) *CreateAclRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAclRequest) SetRegionId(v string) *CreateAclRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAclRequest) SetResourceGroupId(v string) *CreateAclRequest {
	s.ResourceGroupId = &v
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
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclResponseBody) SetAclId(v string) *CreateAclResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateAclResponseBody) SetRequestId(v string) *CreateAclResponseBody {
	s.RequestId = &v
	return s
}

type CreateAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAclResponse) SetStatusCode(v int32) *CreateAclResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAclResponse) SetBody(v *CreateAclResponseBody) *CreateAclResponse {
	s.Body = v
	return s
}

type CreateApplicationMonitorRequest struct {
	AcceleratorId   *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DetectEnable    *bool   `json:"DetectEnable,omitempty" xml:"DetectEnable,omitempty"`
	DetectThreshold *int32  `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	DetectTimes     *int32  `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	OptionsJson     *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SilenceTime     *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateApplicationMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationMonitorRequest) SetAcceleratorId(v string) *CreateApplicationMonitorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetAddress(v string) *CreateApplicationMonitorRequest {
	s.Address = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetClientToken(v string) *CreateApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetDetectEnable(v bool) *CreateApplicationMonitorRequest {
	s.DetectEnable = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetDetectThreshold(v int32) *CreateApplicationMonitorRequest {
	s.DetectThreshold = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetDetectTimes(v int32) *CreateApplicationMonitorRequest {
	s.DetectTimes = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetListenerId(v string) *CreateApplicationMonitorRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetOptionsJson(v string) *CreateApplicationMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetRegionId(v string) *CreateApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetSilenceTime(v int32) *CreateApplicationMonitorRequest {
	s.SilenceTime = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetTaskName(v string) *CreateApplicationMonitorRequest {
	s.TaskName = &v
	return s
}

type CreateApplicationMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateApplicationMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationMonitorResponseBody) SetRequestId(v string) *CreateApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationMonitorResponseBody) SetTaskId(v string) *CreateApplicationMonitorResponseBody {
	s.TaskId = &v
	return s
}

type CreateApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationMonitorResponse) SetHeaders(v map[string]*string) *CreateApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationMonitorResponse) SetStatusCode(v int32) *CreateApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationMonitorResponse) SetBody(v *CreateApplicationMonitorResponseBody) *CreateApplicationMonitorResponse {
	s.Body = v
	return s
}

type CreateBandwidthPackageRequest struct {
	AutoPay                *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew              *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration      *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoUseCoupon          *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	Bandwidth              *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType          *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	BillingType            *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CbnGeographicRegionIdA *string `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	CbnGeographicRegionIdB *string `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	ChargeType             *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Duration               *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle           *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Ratio                  *int32  `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId        *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Type                   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageRequest) SetAutoPay(v bool) *CreateBandwidthPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetAutoRenew(v bool) *CreateBandwidthPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetAutoRenewDuration(v int32) *CreateBandwidthPackageRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetAutoUseCoupon(v string) *CreateBandwidthPackageRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBandwidth(v int32) *CreateBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBandwidthType(v string) *CreateBandwidthPackageRequest {
	s.BandwidthType = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetBillingType(v string) *CreateBandwidthPackageRequest {
	s.BillingType = &v
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

func (s *CreateBandwidthPackageRequest) SetChargeType(v string) *CreateBandwidthPackageRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetClientToken(v string) *CreateBandwidthPackageRequest {
	s.ClientToken = &v
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

func (s *CreateBandwidthPackageRequest) SetRatio(v int32) *CreateBandwidthPackageRequest {
	s.Ratio = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetRegionId(v string) *CreateBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetResourceGroupId(v string) *CreateBandwidthPackageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateBandwidthPackageRequest) SetType(v string) *CreateBandwidthPackageRequest {
	s.Type = &v
	return s
}

type CreateBandwidthPackageResponseBody struct {
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *CreateBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateBandwidthPackageResponseBody) SetOrderId(v string) *CreateBandwidthPackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateBandwidthPackageResponseBody) SetRequestId(v string) *CreateBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

type CreateBandwidthPackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateBandwidthPackageResponse) SetStatusCode(v int32) *CreateBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBandwidthPackageResponse) SetBody(v *CreateBandwidthPackageResponseBody) *CreateBandwidthPackageResponse {
	s.Body = v
	return s
}

type CreateBasicAccelerateIpRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId       *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicAccelerateIpRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpRequest) SetAcceleratorId(v string) *CreateBasicAccelerateIpRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpRequest) SetClientToken(v string) *CreateBasicAccelerateIpRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicAccelerateIpRequest) SetIpSetId(v string) *CreateBasicAccelerateIpRequest {
	s.IpSetId = &v
	return s
}

func (s *CreateBasicAccelerateIpRequest) SetRegionId(v string) *CreateBasicAccelerateIpRequest {
	s.RegionId = &v
	return s
}

type CreateBasicAccelerateIpResponseBody struct {
	AccelerateIpAddress *string `json:"AccelerateIpAddress,omitempty" xml:"AccelerateIpAddress,omitempty"`
	AccelerateIpId      *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId       *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	IpSetId             *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateBasicAccelerateIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpResponseBody) SetAccelerateIpAddress(v string) *CreateBasicAccelerateIpResponseBody {
	s.AccelerateIpAddress = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetAccelerateIpId(v string) *CreateBasicAccelerateIpResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetAcceleratorId(v string) *CreateBasicAccelerateIpResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetIpSetId(v string) *CreateBasicAccelerateIpResponseBody {
	s.IpSetId = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetRequestId(v string) *CreateBasicAccelerateIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicAccelerateIpResponseBody) SetState(v string) *CreateBasicAccelerateIpResponseBody {
	s.State = &v
	return s
}

type CreateBasicAccelerateIpResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBasicAccelerateIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBasicAccelerateIpResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpResponse) SetHeaders(v map[string]*string) *CreateBasicAccelerateIpResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicAccelerateIpResponse) SetStatusCode(v int32) *CreateBasicAccelerateIpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicAccelerateIpResponse) SetBody(v *CreateBasicAccelerateIpResponseBody) *CreateBasicAccelerateIpResponse {
	s.Body = v
	return s
}

type CreateBasicAccelerateIpEndpointRelationRequest struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId  *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationRequest) SetAccelerateIpId(v string) *CreateBasicAccelerateIpEndpointRelationRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationRequest) SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationRequest) SetClientToken(v string) *CreateBasicAccelerateIpEndpointRelationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationRequest) SetEndpointId(v string) *CreateBasicAccelerateIpEndpointRelationRequest {
	s.EndpointId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationRequest) SetRegionId(v string) *CreateBasicAccelerateIpEndpointRelationRequest {
	s.RegionId = &v
	return s
}

type CreateBasicAccelerateIpEndpointRelationResponseBody struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId  *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State          *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetAccelerateIpId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetEndpointId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetRequestId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetState(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.State = &v
	return s
}

type CreateBasicAccelerateIpEndpointRelationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBasicAccelerateIpEndpointRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBasicAccelerateIpEndpointRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) SetHeaders(v map[string]*string) *CreateBasicAccelerateIpEndpointRelationResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) SetStatusCode(v int32) *CreateBasicAccelerateIpEndpointRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponse) SetBody(v *CreateBasicAccelerateIpEndpointRelationResponseBody) *CreateBasicAccelerateIpEndpointRelationResponse {
	s.Body = v
	return s
}

type CreateBasicAccelerateIpEndpointRelationsRequest struct {
	AccelerateIpEndpointRelations []*CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations `json:"AccelerateIpEndpointRelations,omitempty" xml:"AccelerateIpEndpointRelations,omitempty" type:"Repeated"`
	AcceleratorId                 *string                                                                         `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken                   *string                                                                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId                      *string                                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationsRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) SetAccelerateIpEndpointRelations(v []*CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) *CreateBasicAccelerateIpEndpointRelationsRequest {
	s.AccelerateIpEndpointRelations = v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) SetClientToken(v string) *CreateBasicAccelerateIpEndpointRelationsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) SetRegionId(v string) *CreateBasicAccelerateIpEndpointRelationsRequest {
	s.RegionId = &v
	return s
}

type CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) SetAccelerateIpId(v string) *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations {
	s.AccelerateIpId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) SetEndpointId(v string) *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations {
	s.EndpointId = &v
	return s
}

type CreateBasicAccelerateIpEndpointRelationsResponseBody struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponseBody) SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationsResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponseBody) SetRequestId(v string) *CreateBasicAccelerateIpEndpointRelationsResponseBody {
	s.RequestId = &v
	return s
}

type CreateBasicAccelerateIpEndpointRelationsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBasicAccelerateIpEndpointRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBasicAccelerateIpEndpointRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationsResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) SetHeaders(v map[string]*string) *CreateBasicAccelerateIpEndpointRelationsResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) SetStatusCode(v int32) *CreateBasicAccelerateIpEndpointRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponse) SetBody(v *CreateBasicAccelerateIpEndpointRelationsResponseBody) *CreateBasicAccelerateIpEndpointRelationsResponse {
	s.Body = v
	return s
}

type CreateBasicAcceleratorRequest struct {
	AutoPay              *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration    *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoUseCoupon        *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Duration             *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateBasicAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicAcceleratorRequest) SetAutoPay(v bool) *CreateBasicAcceleratorRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetAutoRenew(v bool) *CreateBasicAcceleratorRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetAutoRenewDuration(v int32) *CreateBasicAcceleratorRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetAutoUseCoupon(v string) *CreateBasicAcceleratorRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetBandwidthBillingType(v string) *CreateBasicAcceleratorRequest {
	s.BandwidthBillingType = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetClientToken(v string) *CreateBasicAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetDuration(v int32) *CreateBasicAcceleratorRequest {
	s.Duration = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetPricingCycle(v string) *CreateBasicAcceleratorRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetRegionId(v string) *CreateBasicAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBasicAcceleratorRequest) SetResourceGroupId(v string) *CreateBasicAcceleratorRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateBasicAcceleratorResponseBody struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicAcceleratorResponseBody) SetAcceleratorId(v string) *CreateBasicAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAcceleratorResponseBody) SetOrderId(v string) *CreateBasicAcceleratorResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateBasicAcceleratorResponseBody) SetRequestId(v string) *CreateBasicAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type CreateBasicAcceleratorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBasicAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBasicAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicAcceleratorResponse) SetHeaders(v map[string]*string) *CreateBasicAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicAcceleratorResponse) SetStatusCode(v int32) *CreateBasicAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicAcceleratorResponse) SetBody(v *CreateBasicAcceleratorResponseBody) *CreateBasicAcceleratorResponse {
	s.Body = v
	return s
}

type CreateBasicEndpointRequest struct {
	AcceleratorId          *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointAddress        *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointGroupId        *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointSubAddress     *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointZoneId         *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointRequest) SetAcceleratorId(v string) *CreateBasicEndpointRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetClientToken(v string) *CreateBasicEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointAddress(v string) *CreateBasicEndpointRequest {
	s.EndpointAddress = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointGroupId(v string) *CreateBasicEndpointRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointSubAddress(v string) *CreateBasicEndpointRequest {
	s.EndpointSubAddress = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointSubAddressType(v string) *CreateBasicEndpointRequest {
	s.EndpointSubAddressType = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointType(v string) *CreateBasicEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointZoneId(v string) *CreateBasicEndpointRequest {
	s.EndpointZoneId = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetName(v string) *CreateBasicEndpointRequest {
	s.Name = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetRegionId(v string) *CreateBasicEndpointRequest {
	s.RegionId = &v
	return s
}

type CreateBasicEndpointResponseBody struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId      *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointResponseBody) SetEndpointGroupId(v string) *CreateBasicEndpointResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointResponseBody) SetEndpointId(v string) *CreateBasicEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateBasicEndpointResponseBody) SetRequestId(v string) *CreateBasicEndpointResponseBody {
	s.RequestId = &v
	return s
}

type CreateBasicEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBasicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBasicEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointResponse) SetHeaders(v map[string]*string) *CreateBasicEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicEndpointResponse) SetStatusCode(v int32) *CreateBasicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicEndpointResponse) SetBody(v *CreateBasicEndpointResponseBody) *CreateBasicEndpointResponse {
	s.Body = v
	return s
}

type CreateBasicEndpointGroupRequest struct {
	AcceleratorId       *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointAddress     *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointSubAddress  *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointType        *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointGroupRequest) SetAcceleratorId(v string) *CreateBasicEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetClientToken(v string) *CreateBasicEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetDescription(v string) *CreateBasicEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetEndpointAddress(v string) *CreateBasicEndpointGroupRequest {
	s.EndpointAddress = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetEndpointGroupRegion(v string) *CreateBasicEndpointGroupRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetEndpointSubAddress(v string) *CreateBasicEndpointGroupRequest {
	s.EndpointSubAddress = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetEndpointType(v string) *CreateBasicEndpointGroupRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetName(v string) *CreateBasicEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetRegionId(v string) *CreateBasicEndpointGroupRequest {
	s.RegionId = &v
	return s
}

type CreateBasicEndpointGroupResponseBody struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointGroupResponseBody) SetEndpointGroupId(v string) *CreateBasicEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointGroupResponseBody) SetRequestId(v string) *CreateBasicEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateBasicEndpointGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBasicEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBasicEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointGroupResponse) SetHeaders(v map[string]*string) *CreateBasicEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicEndpointGroupResponse) SetStatusCode(v int32) *CreateBasicEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicEndpointGroupResponse) SetBody(v *CreateBasicEndpointGroupResponseBody) *CreateBasicEndpointGroupResponse {
	s.Body = v
	return s
}

type CreateBasicEndpointsRequest struct {
	AcceleratorId   *string                                 `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken     *string                                 `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string                                 `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	Endpoints       []*CreateBasicEndpointsRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	RegionId        *string                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointsRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsRequest) SetAcceleratorId(v string) *CreateBasicEndpointsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicEndpointsRequest) SetClientToken(v string) *CreateBasicEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicEndpointsRequest) SetEndpointGroupId(v string) *CreateBasicEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointsRequest) SetEndpoints(v []*CreateBasicEndpointsRequestEndpoints) *CreateBasicEndpointsRequest {
	s.Endpoints = v
	return s
}

func (s *CreateBasicEndpointsRequest) SetRegionId(v string) *CreateBasicEndpointsRequest {
	s.RegionId = &v
	return s
}

type CreateBasicEndpointsRequestEndpoints struct {
	EndpointAddress        *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointSubAddress     *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointZoneId         *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateBasicEndpointsRequestEndpoints) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointsRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointAddress(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointAddress = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointSubAddress(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointSubAddress = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointSubAddressType(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointSubAddressType = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointType(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointType = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointZoneId(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointZoneId = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetName(v string) *CreateBasicEndpointsRequestEndpoints {
	s.Name = &v
	return s
}

type CreateBasicEndpointsResponseBody struct {
	EndpointGroupId *string                                      `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	Endpoints       []*CreateBasicEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsResponseBody) SetEndpointGroupId(v string) *CreateBasicEndpointsResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointsResponseBody) SetEndpoints(v []*CreateBasicEndpointsResponseBodyEndpoints) *CreateBasicEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *CreateBasicEndpointsResponseBody) SetRequestId(v string) *CreateBasicEndpointsResponseBody {
	s.RequestId = &v
	return s
}

type CreateBasicEndpointsResponseBodyEndpoints struct {
	EndpointAddress    *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointId         *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointSubAddress *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointType       *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s CreateBasicEndpointsResponseBodyEndpoints) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) SetEndpointAddress(v string) *CreateBasicEndpointsResponseBodyEndpoints {
	s.EndpointAddress = &v
	return s
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) SetEndpointId(v string) *CreateBasicEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) SetEndpointSubAddress(v string) *CreateBasicEndpointsResponseBodyEndpoints {
	s.EndpointSubAddress = &v
	return s
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) SetEndpointType(v string) *CreateBasicEndpointsResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

type CreateBasicEndpointsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBasicEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBasicEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicEndpointsResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsResponse) SetHeaders(v map[string]*string) *CreateBasicEndpointsResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicEndpointsResponse) SetStatusCode(v int32) *CreateBasicEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicEndpointsResponse) SetBody(v *CreateBasicEndpointsResponseBody) *CreateBasicEndpointsResponse {
	s.Body = v
	return s
}

type CreateBasicIpSetRequest struct {
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	AcceleratorId      *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Bandwidth          *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IspType            *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicIpSetRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicIpSetRequest) SetAccelerateRegionId(v string) *CreateBasicIpSetRequest {
	s.AccelerateRegionId = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetAcceleratorId(v string) *CreateBasicIpSetRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetBandwidth(v int64) *CreateBasicIpSetRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetClientToken(v string) *CreateBasicIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetIspType(v string) *CreateBasicIpSetRequest {
	s.IspType = &v
	return s
}

func (s *CreateBasicIpSetRequest) SetRegionId(v string) *CreateBasicIpSetRequest {
	s.RegionId = &v
	return s
}

type CreateBasicIpSetResponseBody struct {
	IpSetId   *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicIpSetResponseBody) SetIpSetId(v string) *CreateBasicIpSetResponseBody {
	s.IpSetId = &v
	return s
}

func (s *CreateBasicIpSetResponseBody) SetRequestId(v string) *CreateBasicIpSetResponseBody {
	s.RequestId = &v
	return s
}

type CreateBasicIpSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBasicIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBasicIpSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBasicIpSetResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicIpSetResponse) SetHeaders(v map[string]*string) *CreateBasicIpSetResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicIpSetResponse) SetStatusCode(v int32) *CreateBasicIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicIpSetResponse) SetBody(v *CreateBasicIpSetResponseBody) *CreateBasicIpSetResponse {
	s.Body = v
	return s
}

type CreateCustomRoutingEndpointGroupDestinationsRequest struct {
	ClientToken               *string                                                                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationConfigurations []*CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	DryRun                    *bool                                                                           `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointGroupId           *string                                                                         `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RegionId                  *string                                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetClientToken(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetDestinationConfigurations(v []*CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.DestinationConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetDryRun(v bool) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

type CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations struct {
	FromPort  *int32    `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	ToPort    *int32    `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetFromPort(v int32) *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetProtocols(v []*string) *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.Protocols = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetToPort(v int32) *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.ToPort = &v
	return s
}

type CreateCustomRoutingEndpointGroupDestinationsResponseBody struct {
	DestinationIds []*string `json:"DestinationIds,omitempty" xml:"DestinationIds,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponseBody) SetDestinationIds(v []*string) *CreateCustomRoutingEndpointGroupDestinationsResponseBody {
	s.DestinationIds = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *CreateCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *CreateCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *CreateCustomRoutingEndpointGroupDestinationsResponseBody) *CreateCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

type CreateCustomRoutingEndpointGroupsRequest struct {
	AcceleratorId               *string                                                                `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken                 *string                                                                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                      *bool                                                                  `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointGroupConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	ListenerId                  *string                                                                `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId                    *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetAcceleratorId(v string) *CreateCustomRoutingEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetClientToken(v string) *CreateCustomRoutingEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetDryRun(v bool) *CreateCustomRoutingEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetEndpointGroupConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) *CreateCustomRoutingEndpointGroupsRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetListenerId(v string) *CreateCustomRoutingEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetRegionId(v string) *CreateCustomRoutingEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations struct {
	Description               *string                                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	EndpointConfigurations    []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations    `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupRegion       *string                                                                                         `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	Name                      *string                                                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetDescription(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.Description = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetDestinationConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.DestinationConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetName(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.Name = &v
	return s
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations struct {
	FromPort  *int32    `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	ToPort    *int32    `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) SetFromPort(v int32) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) SetProtocols(v []*string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations {
	s.Protocols = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) SetToPort(v int32) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations {
	s.ToPort = &v
	return s
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	Endpoint                *string                                                                                                          `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	PolicyConfigurations    []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	TrafficToEndpointPolicy *string                                                                                                          `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	Type                    *string                                                                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetPolicyConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.PolicyConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetTrafficToEndpointPolicy(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations struct {
	Address    *string                                                                                                                    `json:"Address,omitempty" xml:"Address,omitempty"`
	PortRanges []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) SetAddress(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) SetPortRanges(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	s.PortRanges = v
	return s
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) SetFromPort(v int32) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) SetToPort(v int32) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

type CreateCustomRoutingEndpointGroupsResponseBody struct {
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsResponseBody) SetEndpointGroupIds(v []*string) *CreateCustomRoutingEndpointGroupsResponseBody {
	s.EndpointGroupIds = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsResponseBody) SetRequestId(v string) *CreateCustomRoutingEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomRoutingEndpointGroupsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomRoutingEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomRoutingEndpointGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsResponse) SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsResponse) SetStatusCode(v int32) *CreateCustomRoutingEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsResponse) SetBody(v *CreateCustomRoutingEndpointGroupsResponseBody) *CreateCustomRoutingEndpointGroupsResponse {
	s.Body = v
	return s
}

type CreateCustomRoutingEndpointTrafficPoliciesRequest struct {
	ClientToken          *string                                                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointId           *string                                                                  `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	PolicyConfigurations []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	RegionId             *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) SetClientToken(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointId(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointId = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) SetPolicyConfigurations(v []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) *CreateCustomRoutingEndpointTrafficPoliciesRequest {
	s.PolicyConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) SetRegionId(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest {
	s.RegionId = &v
	return s
}

type CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations struct {
	Address    *string                                                                            `json:"Address,omitempty" xml:"Address,omitempty"`
	PortRanges []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetAddress(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetPortRanges(v []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.PortRanges = v
	return s
}

type CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) SetFromPort(v int32) *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) SetToPort(v int32) *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

type CreateCustomRoutingEndpointTrafficPoliciesResponseBody struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) SetPolicyIds(v []*string) *CreateCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.PolicyIds = v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) SetRequestId(v string) *CreateCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomRoutingEndpointTrafficPoliciesResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomRoutingEndpointTrafficPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointTrafficPoliciesResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) SetStatusCode(v int32) *CreateCustomRoutingEndpointTrafficPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) SetBody(v *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) *CreateCustomRoutingEndpointTrafficPoliciesResponse {
	s.Body = v
	return s
}

type CreateCustomRoutingEndpointsRequest struct {
	ClientToken            *string                                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointConfigurations []*CreateCustomRoutingEndpointsRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupId        *string                                                      `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RegionId               *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCustomRoutingEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointsRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsRequest) SetClientToken(v string) *CreateCustomRoutingEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequest) SetEndpointConfigurations(v []*CreateCustomRoutingEndpointsRequestEndpointConfigurations) *CreateCustomRoutingEndpointsRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointsRequest) SetEndpointGroupId(v string) *CreateCustomRoutingEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequest) SetRegionId(v string) *CreateCustomRoutingEndpointsRequest {
	s.RegionId = &v
	return s
}

type CreateCustomRoutingEndpointsRequestEndpointConfigurations struct {
	Endpoint                *string                                                                          `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	PolicyConfigurations    []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	TrafficToEndpointPolicy *string                                                                          `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	Type                    *string                                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) SetEndpoint(v string) *CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) SetPolicyConfigurations(v []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) *CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.PolicyConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) SetTrafficToEndpointPolicy(v string) *CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) SetType(v string) *CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.Type = &v
	return s
}

type CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations struct {
	Address    *string                                                                                    `json:"Address,omitempty" xml:"Address,omitempty"`
	PortRanges []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) SetAddress(v string) *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) SetPortRanges(v []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	s.PortRanges = v
	return s
}

type CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) SetFromPort(v int32) *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) SetToPort(v int32) *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

type CreateCustomRoutingEndpointsResponseBody struct {
	EndpointIds []*string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomRoutingEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsResponseBody) SetEndpointIds(v []*string) *CreateCustomRoutingEndpointsResponseBody {
	s.EndpointIds = v
	return s
}

func (s *CreateCustomRoutingEndpointsResponseBody) SetRequestId(v string) *CreateCustomRoutingEndpointsResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomRoutingEndpointsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomRoutingEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomRoutingEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomRoutingEndpointsResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsResponse) SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointsResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoutingEndpointsResponse) SetStatusCode(v int32) *CreateCustomRoutingEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoutingEndpointsResponse) SetBody(v *CreateCustomRoutingEndpointsResponseBody) *CreateCustomRoutingEndpointsResponse {
	s.Body = v
	return s
}

type CreateEndpointGroupRequest struct {
	AcceleratorId              *string                                             `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken                *string                                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointConfigurations     []*CreateEndpointGroupRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupRegion        *string                                             `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointGroupType          *string                                             `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointRequestProtocol    *string                                             `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	HealthCheckEnabled         *bool                                               `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckIntervalSeconds *int32                                              `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath            *string                                             `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort            *int32                                              `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol        *string                                             `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	ListenerId                 *string                                             `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Name                       *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	PortOverrides              []*CreateEndpointGroupRequestPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	RegionId                   *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ThresholdCount             *int32                                              `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	TrafficPercentage          *int32                                              `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s CreateEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequest) SetAcceleratorId(v string) *CreateEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetClientToken(v string) *CreateEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetDescription(v string) *CreateEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointConfigurations(v []*CreateEndpointGroupRequestEndpointConfigurations) *CreateEndpointGroupRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointGroupRegion(v string) *CreateEndpointGroupRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointGroupType(v string) *CreateEndpointGroupRequest {
	s.EndpointGroupType = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointRequestProtocol(v string) *CreateEndpointGroupRequest {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckEnabled(v bool) *CreateEndpointGroupRequest {
	s.HealthCheckEnabled = &v
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

func (s *CreateEndpointGroupRequest) SetListenerId(v string) *CreateEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetName(v string) *CreateEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetPortOverrides(v []*CreateEndpointGroupRequestPortOverrides) *CreateEndpointGroupRequest {
	s.PortOverrides = v
	return s
}

func (s *CreateEndpointGroupRequest) SetRegionId(v string) *CreateEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetThresholdCount(v int32) *CreateEndpointGroupRequest {
	s.ThresholdCount = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetTrafficPercentage(v int32) *CreateEndpointGroupRequest {
	s.TrafficPercentage = &v
	return s
}

type CreateEndpointGroupRequestEndpointConfigurations struct {
	EnableClientIPPreservation *bool   `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	Endpoint                   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Type                       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight                     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateEndpointGroupRequestEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetEnableClientIPPreservation(v bool) *CreateEndpointGroupRequestEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetEndpoint(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetType(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetWeight(v int32) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Weight = &v
	return s
}

type CreateEndpointGroupRequestPortOverrides struct {
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s CreateEndpointGroupRequestPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupRequestPortOverrides) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequestPortOverrides) SetEndpointPort(v int32) *CreateEndpointGroupRequestPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *CreateEndpointGroupRequestPortOverrides) SetListenerPort(v int32) *CreateEndpointGroupRequestPortOverrides {
	s.ListenerPort = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateEndpointGroupResponse) SetStatusCode(v int32) *CreateEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEndpointGroupResponse) SetBody(v *CreateEndpointGroupResponseBody) *CreateEndpointGroupResponse {
	s.Body = v
	return s
}

type CreateEndpointGroupsRequest struct {
	AcceleratorId               *string                                                   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken                 *string                                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                      *bool                                                     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointGroupConfigurations []*CreateEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	ListenerId                  *string                                                   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId                    *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEndpointGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequest) SetAcceleratorId(v string) *CreateEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetClientToken(v string) *CreateEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetDryRun(v bool) *CreateEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetEndpointGroupConfigurations(v []*CreateEndpointGroupsRequestEndpointGroupConfigurations) *CreateEndpointGroupsRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *CreateEndpointGroupsRequest) SetListenerId(v string) *CreateEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetRegionId(v string) *CreateEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

type CreateEndpointGroupsRequestEndpointGroupConfigurations struct {
	EnableClientIPPreservationProxyProtocol *bool                                                                           `json:"EnableClientIPPreservationProxyProtocol,omitempty" xml:"EnableClientIPPreservationProxyProtocol,omitempty"`
	EnableClientIPPreservationToa           *bool                                                                           `json:"EnableClientIPPreservationToa,omitempty" xml:"EnableClientIPPreservationToa,omitempty"`
	EndpointConfigurations                  []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupDescription                *string                                                                         `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	EndpointGroupName                       *string                                                                         `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	EndpointGroupRegion                     *string                                                                         `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointGroupType                       *string                                                                         `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointRequestProtocol                 *string                                                                         `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	HealthCheckEnabled                      *bool                                                                           `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckIntervalSeconds              *int64                                                                          `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath                         *string                                                                         `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort                         *int64                                                                          `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol                     *string                                                                         `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	PortOverrides                           []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	ThresholdCount                          *int64                                                                          `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	TrafficPercentage                       *int64                                                                          `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEnableClientIPPreservationProxyProtocol(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EnableClientIPPreservationProxyProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEnableClientIPPreservationToa(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EnableClientIPPreservationToa = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupDescription(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupDescription = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupName(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupName = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupType(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupType = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointRequestProtocol(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckEnabled(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckIntervalSeconds(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPath(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckProtocol(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetPortOverrides(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.PortOverrides = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetThresholdCount(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.ThresholdCount = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetTrafficPercentage(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.TrafficPercentage = &v
	return s
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *int64  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetWeight(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Weight = &v
	return s
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides struct {
	EndpointPort *int64 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetEndpointPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetListenerPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.ListenerPort = &v
	return s
}

type CreateEndpointGroupsResponseBody struct {
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEndpointGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsResponseBody) SetEndpointGroupIds(v []*string) *CreateEndpointGroupsResponseBody {
	s.EndpointGroupIds = v
	return s
}

func (s *CreateEndpointGroupsResponseBody) SetRequestId(v string) *CreateEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

type CreateEndpointGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEndpointGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsResponse) SetHeaders(v map[string]*string) *CreateEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *CreateEndpointGroupsResponse) SetStatusCode(v int32) *CreateEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEndpointGroupsResponse) SetBody(v *CreateEndpointGroupsResponseBody) *CreateEndpointGroupsResponse {
	s.Body = v
	return s
}

type CreateForwardingRulesRequest struct {
	AcceleratorId   *string                                        `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken     *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForwardingRules []*CreateForwardingRulesRequestForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	ListenerId      *string                                        `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId        *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateForwardingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequest) SetAcceleratorId(v string) *CreateForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetClientToken(v string) *CreateForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetForwardingRules(v []*CreateForwardingRulesRequestForwardingRules) *CreateForwardingRulesRequest {
	s.ForwardingRules = v
	return s
}

func (s *CreateForwardingRulesRequest) SetListenerId(v string) *CreateForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetRegionId(v string) *CreateForwardingRulesRequest {
	s.RegionId = &v
	return s
}

type CreateForwardingRulesRequestForwardingRules struct {
	ForwardingRuleName *string                                                      `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	Priority           *int32                                                       `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleActions        []*CreateForwardingRulesRequestForwardingRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	RuleConditions     []*CreateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	RuleDirection      *string                                                      `json:"RuleDirection,omitempty" xml:"RuleDirection,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRules) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRules) SetForwardingRuleName(v string) *CreateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetPriority(v int32) *CreateForwardingRulesRequestForwardingRules {
	s.Priority = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleActions(v []*CreateForwardingRulesRequestForwardingRulesRuleActions) *CreateForwardingRulesRequestForwardingRules {
	s.RuleActions = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleConditions(v []*CreateForwardingRulesRequestForwardingRulesRuleConditions) *CreateForwardingRulesRequestForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleDirection(v string) *CreateForwardingRulesRequestForwardingRules {
	s.RuleDirection = &v
	return s
}

type CreateForwardingRulesRequestForwardingRulesRuleActions struct {
	ForwardGroupConfig *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	Order              *int32                                                                    `json:"Order,omitempty" xml:"Order,omitempty"`
	RuleActionType     *string                                                                   `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	RuleActionValue    *string                                                                   `json:"RuleActionValue,omitempty" xml:"RuleActionValue,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetForwardGroupConfig(v *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetOrder(v int32) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionType(v string) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionValue(v string) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionValue = &v
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

type CreateForwardingRulesRequestForwardingRulesRuleConditions struct {
	HostConfig         *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	PathConfig         *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	RuleConditionType  *string                                                              `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	RuleConditionValue *string                                                              `json:"RuleConditionValue,omitempty" xml:"RuleConditionValue,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetHostConfig(v *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetPathConfig(v *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionType(v string) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionValue(v string) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionValue = &v
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

type CreateForwardingRulesResponseBody struct {
	ForwardingRules []*CreateForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateForwardingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesResponseBody) SetForwardingRules(v []*CreateForwardingRulesResponseBodyForwardingRules) *CreateForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

func (s *CreateForwardingRulesResponseBody) SetRequestId(v string) *CreateForwardingRulesResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateForwardingRulesResponse) SetStatusCode(v int32) *CreateForwardingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateForwardingRulesResponse) SetBody(v *CreateForwardingRulesResponseBody) *CreateForwardingRulesResponse {
	s.Body = v
	return s
}

type CreateIpSetsRequest struct {
	AccelerateRegion []*CreateIpSetsRequestAccelerateRegion `json:"AccelerateRegion,omitempty" xml:"AccelerateRegion,omitempty" type:"Repeated"`
	AcceleratorId    *string                                `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken      *string                                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId         *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateIpSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpSetsRequest) GoString() string {
	return s.String()
}

func (s *CreateIpSetsRequest) SetAccelerateRegion(v []*CreateIpSetsRequestAccelerateRegion) *CreateIpSetsRequest {
	s.AccelerateRegion = v
	return s
}

func (s *CreateIpSetsRequest) SetAcceleratorId(v string) *CreateIpSetsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateIpSetsRequest) SetClientToken(v string) *CreateIpSetsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpSetsRequest) SetRegionId(v string) *CreateIpSetsRequest {
	s.RegionId = &v
	return s
}

type CreateIpSetsRequestAccelerateRegion struct {
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	IpVersion          *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IspType            *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
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

func (s *CreateIpSetsRequestAccelerateRegion) SetBandwidth(v int32) *CreateIpSetsRequestAccelerateRegion {
	s.Bandwidth = &v
	return s
}

func (s *CreateIpSetsRequestAccelerateRegion) SetIpVersion(v string) *CreateIpSetsRequestAccelerateRegion {
	s.IpVersion = &v
	return s
}

func (s *CreateIpSetsRequestAccelerateRegion) SetIspType(v string) *CreateIpSetsRequestAccelerateRegion {
	s.IspType = &v
	return s
}

type CreateIpSetsResponseBody struct {
	AcceleratorId *string                           `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	IpSets        []*CreateIpSetsResponseBodyIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
	RequestId     *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpSetsResponseBody) SetAcceleratorId(v string) *CreateIpSetsResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateIpSetsResponseBody) SetIpSets(v []*CreateIpSetsResponseBodyIpSets) *CreateIpSetsResponseBody {
	s.IpSets = v
	return s
}

func (s *CreateIpSetsResponseBody) SetRequestId(v string) *CreateIpSetsResponseBody {
	s.RequestId = &v
	return s
}

type CreateIpSetsResponseBodyIpSets struct {
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	IpSetId            *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	IspType            *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
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

func (s *CreateIpSetsResponseBodyIpSets) SetIspType(v string) *CreateIpSetsResponseBodyIpSets {
	s.IspType = &v
	return s
}

type CreateIpSetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateIpSetsResponse) SetStatusCode(v int32) *CreateIpSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpSetsResponse) SetBody(v *CreateIpSetsResponseBody) *CreateIpSetsResponse {
	s.Body = v
	return s
}

type CreateListenerRequest struct {
	AcceleratorId                            *string                                                          `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Certificates                             []*CreateListenerRequestCertificates                             `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	ClientAffinity                           *string                                                          `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	ClientToken                              *string                                                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CustomRoutingEndpointGroupConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurations `json:"CustomRoutingEndpointGroupConfigurations,omitempty" xml:"CustomRoutingEndpointGroupConfigurations,omitempty" type:"Repeated"`
	Description                              *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointGroupConfigurations              []*CreateListenerRequestEndpointGroupConfigurations              `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	Name                                     *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges                               []*CreateListenerRequestPortRanges                               `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol                                 *string                                                          `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ProxyProtocol                            *bool                                                            `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	RegionId                                 *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityPolicyId                         *string                                                          `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	Type                                     *string                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	XForwardedForConfig                      *CreateListenerRequestXForwardedForConfig                        `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s CreateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) SetAcceleratorId(v string) *CreateListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateListenerRequest) SetCertificates(v []*CreateListenerRequestCertificates) *CreateListenerRequest {
	s.Certificates = v
	return s
}

func (s *CreateListenerRequest) SetClientAffinity(v string) *CreateListenerRequest {
	s.ClientAffinity = &v
	return s
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetCustomRoutingEndpointGroupConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurations) *CreateListenerRequest {
	s.CustomRoutingEndpointGroupConfigurations = v
	return s
}

func (s *CreateListenerRequest) SetDescription(v string) *CreateListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateListenerRequest) SetEndpointGroupConfigurations(v []*CreateListenerRequestEndpointGroupConfigurations) *CreateListenerRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *CreateListenerRequest) SetName(v string) *CreateListenerRequest {
	s.Name = &v
	return s
}

func (s *CreateListenerRequest) SetPortRanges(v []*CreateListenerRequestPortRanges) *CreateListenerRequest {
	s.PortRanges = v
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

func (s *CreateListenerRequest) SetRegionId(v string) *CreateListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateListenerRequest) SetSecurityPolicyId(v string) *CreateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateListenerRequest) SetType(v string) *CreateListenerRequest {
	s.Type = &v
	return s
}

func (s *CreateListenerRequest) SetXForwardedForConfig(v *CreateListenerRequestXForwardedForConfig) *CreateListenerRequest {
	s.XForwardedForConfig = v
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

type CreateListenerRequestCustomRoutingEndpointGroupConfigurations struct {
	Description               *string                                                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationConfigurations []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	EndpointConfigurations    []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations    `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupRegion       *string                                                                                   `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	Name                      *string                                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetDescription(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.Description = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetDestinationConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.DestinationConfigurations = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurations) SetName(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurations {
	s.Name = &v
	return s
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations struct {
	FromPort  *int32    `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	ToPort    *int32    `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) SetFromPort(v int32) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations {
	s.FromPort = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) SetProtocols(v []*string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations {
	s.Protocols = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations) SetToPort(v int32) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsDestinationConfigurations {
	s.ToPort = &v
	return s
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations struct {
	Endpoint                *string                                                                                                    `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	PolicyConfigurations    []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	TrafficToEndpointPolicy *string                                                                                                    `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	Type                    *string                                                                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) SetPolicyConfigurations(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	s.PolicyConfigurations = v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) SetTrafficToEndpointPolicy(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations struct {
	Address    *string                                                                                                              `json:"Address,omitempty" xml:"Address,omitempty"`
	PortRanges []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) SetAddress(v string) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) SetPortRanges(v []*CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	s.PortRanges = v
	return s
}

type CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) SetFromPort(v int32) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) SetToPort(v int32) *CreateListenerRequestCustomRoutingEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

type CreateListenerRequestEndpointGroupConfigurations struct {
	EnableClientIPPreservationProxyProtocol *bool                                                                     `json:"EnableClientIPPreservationProxyProtocol,omitempty" xml:"EnableClientIPPreservationProxyProtocol,omitempty"`
	EnableClientIPPreservationToa           *bool                                                                     `json:"EnableClientIPPreservationToa,omitempty" xml:"EnableClientIPPreservationToa,omitempty"`
	EndpointConfigurations                  []*CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupDescription                *string                                                                   `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	EndpointGroupName                       *string                                                                   `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	EndpointGroupRegion                     *string                                                                   `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointGroupType                       *string                                                                   `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointRequestProtocol                 *string                                                                   `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	HealthCheckEnabled                      *bool                                                                     `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckIntervalSeconds              *int64                                                                    `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath                         *string                                                                   `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort                         *int64                                                                    `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol                     *string                                                                   `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	PortOverrides                           []*CreateListenerRequestEndpointGroupConfigurationsPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	ThresholdCount                          *int64                                                                    `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	TrafficPercentage                       *int64                                                                    `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s CreateListenerRequestEndpointGroupConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEnableClientIPPreservationProxyProtocol(v bool) *CreateListenerRequestEndpointGroupConfigurations {
	s.EnableClientIPPreservationProxyProtocol = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEnableClientIPPreservationToa(v bool) *CreateListenerRequestEndpointGroupConfigurations {
	s.EnableClientIPPreservationToa = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointGroupDescription(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointGroupDescription = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointGroupName(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointGroupName = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointGroupType(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointGroupType = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetEndpointRequestProtocol(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckEnabled(v bool) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckIntervalSeconds(v int64) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckPath(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckPort(v int64) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetHealthCheckProtocol(v string) *CreateListenerRequestEndpointGroupConfigurations {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetPortOverrides(v []*CreateListenerRequestEndpointGroupConfigurationsPortOverrides) *CreateListenerRequestEndpointGroupConfigurations {
	s.PortOverrides = v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetThresholdCount(v int64) *CreateListenerRequestEndpointGroupConfigurations {
	s.ThresholdCount = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurations) SetTrafficPercentage(v int64) *CreateListenerRequestEndpointGroupConfigurations {
	s.TrafficPercentage = &v
	return s
}

type CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *int64  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations) SetWeight(v int64) *CreateListenerRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Weight = &v
	return s
}

type CreateListenerRequestEndpointGroupConfigurationsPortOverrides struct {
	EndpointPort *int64 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s CreateListenerRequestEndpointGroupConfigurationsPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestEndpointGroupConfigurationsPortOverrides) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestEndpointGroupConfigurationsPortOverrides) SetEndpointPort(v int64) *CreateListenerRequestEndpointGroupConfigurationsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *CreateListenerRequestEndpointGroupConfigurationsPortOverrides) SetListenerPort(v int64) *CreateListenerRequestEndpointGroupConfigurationsPortOverrides {
	s.ListenerPort = &v
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

type CreateListenerRequestXForwardedForConfig struct {
	XForwardedForGaApEnabled  *bool `json:"XForwardedForGaApEnabled,omitempty" xml:"XForwardedForGaApEnabled,omitempty"`
	XForwardedForGaIdEnabled  *bool `json:"XForwardedForGaIdEnabled,omitempty" xml:"XForwardedForGaIdEnabled,omitempty"`
	XForwardedForPortEnabled  *bool `json:"XForwardedForPortEnabled,omitempty" xml:"XForwardedForPortEnabled,omitempty"`
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	XRealIpEnabled            *bool `json:"XRealIpEnabled,omitempty" xml:"XRealIpEnabled,omitempty"`
}

func (s CreateListenerRequestXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForGaApEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForGaApEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForGaIdEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForGaIdEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForPortEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForPortEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXRealIpEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XRealIpEnabled = &v
	return s
}

type CreateListenerResponseBody struct {
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBody) SetListenerId(v string) *CreateListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *CreateListenerResponseBody) SetRequestId(v string) *CreateListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateListenerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateListenerResponse) SetStatusCode(v int32) *CreateListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateListenerResponse) SetBody(v *CreateListenerResponseBody) *CreateListenerResponse {
	s.Body = v
	return s
}

type CreateSpareIpsRequest struct {
	AcceleratorId *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpareIps      []*string `json:"SpareIps,omitempty" xml:"SpareIps,omitempty" type:"Repeated"`
}

func (s CreateSpareIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSpareIpsRequest) GoString() string {
	return s.String()
}

func (s *CreateSpareIpsRequest) SetAcceleratorId(v string) *CreateSpareIpsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateSpareIpsRequest) SetClientToken(v string) *CreateSpareIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSpareIpsRequest) SetDryRun(v bool) *CreateSpareIpsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSpareIpsRequest) SetRegionId(v string) *CreateSpareIpsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSpareIpsRequest) SetSpareIps(v []*string) *CreateSpareIpsRequest {
	s.SpareIps = v
	return s
}

type CreateSpareIpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSpareIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSpareIpsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSpareIpsResponseBody) SetRequestId(v string) *CreateSpareIpsResponseBody {
	s.RequestId = &v
	return s
}

type CreateSpareIpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSpareIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSpareIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSpareIpsResponse) GoString() string {
	return s.String()
}

func (s *CreateSpareIpsResponse) SetHeaders(v map[string]*string) *CreateSpareIpsResponse {
	s.Headers = v
	return s
}

func (s *CreateSpareIpsResponse) SetStatusCode(v int32) *CreateSpareIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSpareIpsResponse) SetBody(v *CreateSpareIpsResponseBody) *CreateSpareIpsResponse {
	s.Body = v
	return s
}

type DeleteAcceleratorRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DeleteAcceleratorRequest) SetAcceleratorId(v string) *DeleteAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteAcceleratorRequest) SetRegionId(v string) *DeleteAcceleratorRequest {
	s.RegionId = &v
	return s
}

type DeleteAcceleratorResponseBody struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAcceleratorResponseBody) SetAcceleratorId(v string) *DeleteAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteAcceleratorResponseBody) SetRequestId(v string) *DeleteAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAcceleratorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAcceleratorResponse) SetStatusCode(v int32) *DeleteAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAcceleratorResponse) SetBody(v *DeleteAcceleratorResponseBody) *DeleteAcceleratorResponse {
	s.Body = v
	return s
}

type DeleteAclRequest struct {
	AclId       *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRequest) GoString() string {
	return s.String()
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

func (s *DeleteAclRequest) SetRegionId(v string) *DeleteAclRequest {
	s.RegionId = &v
	return s
}

type DeleteAclResponseBody struct {
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclResponseBody) SetAclId(v string) *DeleteAclResponseBody {
	s.AclId = &v
	return s
}

func (s *DeleteAclResponseBody) SetRequestId(v string) *DeleteAclResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAclResponse) SetStatusCode(v int32) *DeleteAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAclResponse) SetBody(v *DeleteAclResponseBody) *DeleteAclResponse {
	s.Body = v
	return s
}

type DeleteApplicationMonitorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteApplicationMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationMonitorRequest) SetClientToken(v string) *DeleteApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteApplicationMonitorRequest) SetRegionId(v string) *DeleteApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApplicationMonitorRequest) SetTaskId(v string) *DeleteApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

type DeleteApplicationMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationMonitorResponseBody) SetRequestId(v string) *DeleteApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationMonitorResponse) SetHeaders(v map[string]*string) *DeleteApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationMonitorResponse) SetStatusCode(v int32) *DeleteApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationMonitorResponse) SetBody(v *DeleteApplicationMonitorResponseBody) *DeleteApplicationMonitorResponse {
	s.Body = v
	return s
}

type DeleteBandwidthPackageRequest struct {
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageRequest) SetBandwidthPackageId(v string) *DeleteBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetClientToken(v string) *DeleteBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetRegionId(v string) *DeleteBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

type DeleteBandwidthPackageResponseBody struct {
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *DeleteBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *DeleteBandwidthPackageResponseBody) SetRequestId(v string) *DeleteBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBandwidthPackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteBandwidthPackageResponse) SetStatusCode(v int32) *DeleteBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBandwidthPackageResponse) SetBody(v *DeleteBandwidthPackageResponseBody) *DeleteBandwidthPackageResponse {
	s.Body = v
	return s
}

type DeleteBasicAccelerateIpRequest struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBasicAccelerateIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAccelerateIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpRequest) SetAccelerateIpId(v string) *DeleteBasicAccelerateIpRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *DeleteBasicAccelerateIpRequest) SetClientToken(v string) *DeleteBasicAccelerateIpRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicAccelerateIpRequest) SetRegionId(v string) *DeleteBasicAccelerateIpRequest {
	s.RegionId = &v
	return s
}

type DeleteBasicAccelerateIpResponseBody struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicAccelerateIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAccelerateIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpResponseBody) SetAccelerateIpId(v string) *DeleteBasicAccelerateIpResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *DeleteBasicAccelerateIpResponseBody) SetRequestId(v string) *DeleteBasicAccelerateIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBasicAccelerateIpResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBasicAccelerateIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBasicAccelerateIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAccelerateIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpResponse) SetHeaders(v map[string]*string) *DeleteBasicAccelerateIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicAccelerateIpResponse) SetStatusCode(v int32) *DeleteBasicAccelerateIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicAccelerateIpResponse) SetBody(v *DeleteBasicAccelerateIpResponseBody) *DeleteBasicAccelerateIpResponse {
	s.Body = v
	return s
}

type DeleteBasicAccelerateIpEndpointRelationRequest struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId  *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBasicAccelerateIpEndpointRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAccelerateIpEndpointRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetAccelerateIpId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetAcceleratorId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetClientToken(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetEndpointId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetRegionId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.RegionId = &v
	return s
}

type DeleteBasicAccelerateIpEndpointRelationResponseBody struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId  *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State          *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DeleteBasicAccelerateIpEndpointRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAccelerateIpEndpointRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetAccelerateIpId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetAcceleratorId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetEndpointId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetRequestId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetState(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.State = &v
	return s
}

type DeleteBasicAccelerateIpEndpointRelationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBasicAccelerateIpEndpointRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBasicAccelerateIpEndpointRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAccelerateIpEndpointRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) SetHeaders(v map[string]*string) *DeleteBasicAccelerateIpEndpointRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) SetStatusCode(v int32) *DeleteBasicAccelerateIpEndpointRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponse) SetBody(v *DeleteBasicAccelerateIpEndpointRelationResponseBody) *DeleteBasicAccelerateIpEndpointRelationResponse {
	s.Body = v
	return s
}

type DeleteBasicAcceleratorRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBasicAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicAcceleratorRequest) SetAcceleratorId(v string) *DeleteBasicAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteBasicAcceleratorRequest) SetRegionId(v string) *DeleteBasicAcceleratorRequest {
	s.RegionId = &v
	return s
}

type DeleteBasicAcceleratorResponseBody struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicAcceleratorResponseBody) SetAcceleratorId(v string) *DeleteBasicAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteBasicAcceleratorResponseBody) SetRequestId(v string) *DeleteBasicAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBasicAcceleratorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBasicAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBasicAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicAcceleratorResponse) SetHeaders(v map[string]*string) *DeleteBasicAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicAcceleratorResponse) SetStatusCode(v int32) *DeleteBasicAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicAcceleratorResponse) SetBody(v *DeleteBasicAcceleratorResponseBody) *DeleteBasicAcceleratorResponse {
	s.Body = v
	return s
}

type DeleteBasicEndpointRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId      *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBasicEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointRequest) SetClientToken(v string) *DeleteBasicEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicEndpointRequest) SetEndpointGroupId(v string) *DeleteBasicEndpointRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DeleteBasicEndpointRequest) SetEndpointId(v string) *DeleteBasicEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteBasicEndpointRequest) SetRegionId(v string) *DeleteBasicEndpointRequest {
	s.RegionId = &v
	return s
}

type DeleteBasicEndpointResponseBody struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointResponseBody) SetEndpointId(v string) *DeleteBasicEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DeleteBasicEndpointResponseBody) SetRequestId(v string) *DeleteBasicEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBasicEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBasicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBasicEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointResponse) SetHeaders(v map[string]*string) *DeleteBasicEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicEndpointResponse) SetStatusCode(v int32) *DeleteBasicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicEndpointResponse) SetBody(v *DeleteBasicEndpointResponseBody) *DeleteBasicEndpointResponse {
	s.Body = v
	return s
}

type DeleteBasicEndpointGroupRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s DeleteBasicEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointGroupRequest) SetClientToken(v string) *DeleteBasicEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicEndpointGroupRequest) SetEndpointGroupId(v string) *DeleteBasicEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

type DeleteBasicEndpointGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointGroupResponseBody) SetRequestId(v string) *DeleteBasicEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBasicEndpointGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBasicEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBasicEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointGroupResponse) SetHeaders(v map[string]*string) *DeleteBasicEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicEndpointGroupResponse) SetStatusCode(v int32) *DeleteBasicEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicEndpointGroupResponse) SetBody(v *DeleteBasicEndpointGroupResponseBody) *DeleteBasicEndpointGroupResponse {
	s.Body = v
	return s
}

type DeleteBasicIpSetRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId     *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBasicIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicIpSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicIpSetRequest) SetClientToken(v string) *DeleteBasicIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicIpSetRequest) SetIpSetId(v string) *DeleteBasicIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *DeleteBasicIpSetRequest) SetRegionId(v string) *DeleteBasicIpSetRequest {
	s.RegionId = &v
	return s
}

type DeleteBasicIpSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicIpSetResponseBody) SetRequestId(v string) *DeleteBasicIpSetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBasicIpSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBasicIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBasicIpSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBasicIpSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicIpSetResponse) SetHeaders(v map[string]*string) *DeleteBasicIpSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicIpSetResponse) SetStatusCode(v int32) *DeleteBasicIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicIpSetResponse) SetBody(v *DeleteBasicIpSetResponseBody) *DeleteBasicIpSetResponse {
	s.Body = v
	return s
}

type DeleteCustomRoutingEndpointGroupDestinationsRequest struct {
	ClientToken     *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationIds  []*string `json:"DestinationIds,omitempty" xml:"DestinationIds,omitempty" type:"Repeated"`
	DryRun          *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointGroupId *string   `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetClientToken(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetDestinationIds(v []*string) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.DestinationIds = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetDryRun(v bool) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *DeleteCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

type DeleteCustomRoutingEndpointGroupDestinationsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *DeleteCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *DeleteCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *DeleteCustomRoutingEndpointGroupDestinationsResponseBody) *DeleteCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

type DeleteCustomRoutingEndpointGroupsRequest struct {
	AcceleratorId    *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken      *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetAcceleratorId(v string) *DeleteCustomRoutingEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetClientToken(v string) *DeleteCustomRoutingEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetDryRun(v bool) *DeleteCustomRoutingEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetEndpointGroupIds(v []*string) *DeleteCustomRoutingEndpointGroupsRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsRequest) SetRegionId(v string) *DeleteCustomRoutingEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

type DeleteCustomRoutingEndpointGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupsResponseBody) SetRequestId(v string) *DeleteCustomRoutingEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomRoutingEndpointGroupsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomRoutingEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomRoutingEndpointGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) SetStatusCode(v int32) *DeleteCustomRoutingEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupsResponse) SetBody(v *DeleteCustomRoutingEndpointGroupsResponseBody) *DeleteCustomRoutingEndpointGroupsResponse {
	s.Body = v
	return s
}

type DeleteCustomRoutingEndpointTrafficPoliciesRequest struct {
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointId  *string   `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	PolicyIds   []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) SetClientToken(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) SetPolicyIds(v []*string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesRequest) SetRegionId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesRequest {
	s.RegionId = &v
	return s
}

type DeleteCustomRoutingEndpointTrafficPoliciesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) SetRequestId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomRoutingEndpointTrafficPoliciesResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointTrafficPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) SetStatusCode(v int32) *DeleteCustomRoutingEndpointTrafficPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) SetBody(v *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) *DeleteCustomRoutingEndpointTrafficPoliciesResponse {
	s.Body = v
	return s
}

type DeleteCustomRoutingEndpointsRequest struct {
	ClientToken     *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string   `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointIds     []*string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty" type:"Repeated"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomRoutingEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointsRequest) SetClientToken(v string) *DeleteCustomRoutingEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomRoutingEndpointsRequest) SetEndpointGroupId(v string) *DeleteCustomRoutingEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointsRequest) SetEndpointIds(v []*string) *DeleteCustomRoutingEndpointsRequest {
	s.EndpointIds = v
	return s
}

func (s *DeleteCustomRoutingEndpointsRequest) SetRegionId(v string) *DeleteCustomRoutingEndpointsRequest {
	s.RegionId = &v
	return s
}

type DeleteCustomRoutingEndpointsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomRoutingEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointsResponseBody) SetRequestId(v string) *DeleteCustomRoutingEndpointsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomRoutingEndpointsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomRoutingEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomRoutingEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomRoutingEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointsResponse) SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoutingEndpointsResponse) SetStatusCode(v int32) *DeleteCustomRoutingEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoutingEndpointsResponse) SetBody(v *DeleteCustomRoutingEndpointsResponseBody) *DeleteCustomRoutingEndpointsResponse {
	s.Body = v
	return s
}

type DeleteEndpointGroupRequest struct {
	AcceleratorId   *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s DeleteEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupRequest) SetAcceleratorId(v string) *DeleteEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteEndpointGroupRequest) SetClientToken(v string) *DeleteEndpointGroupRequest {
	s.ClientToken = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteEndpointGroupResponse) SetStatusCode(v int32) *DeleteEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEndpointGroupResponse) SetBody(v *DeleteEndpointGroupResponseBody) *DeleteEndpointGroupResponse {
	s.Body = v
	return s
}

type DeleteEndpointGroupsRequest struct {
	ClientToken      *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEndpointGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupsRequest) SetClientToken(v string) *DeleteEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEndpointGroupsRequest) SetDryRun(v bool) *DeleteEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteEndpointGroupsRequest) SetEndpointGroupIds(v []*string) *DeleteEndpointGroupsRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *DeleteEndpointGroupsRequest) SetRegionId(v string) *DeleteEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

type DeleteEndpointGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEndpointGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupsResponseBody) SetRequestId(v string) *DeleteEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEndpointGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEndpointGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupsResponse) SetHeaders(v map[string]*string) *DeleteEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEndpointGroupsResponse) SetStatusCode(v int32) *DeleteEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEndpointGroupsResponse) SetBody(v *DeleteEndpointGroupsResponseBody) *DeleteEndpointGroupsResponse {
	s.Body = v
	return s
}

type DeleteForwardingRulesRequest struct {
	AcceleratorId     *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken       *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForwardingRuleIds []*string `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	ListenerId        *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteForwardingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesRequest) SetAcceleratorId(v string) *DeleteForwardingRulesRequest {
	s.AcceleratorId = &v
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

func (s *DeleteForwardingRulesRequest) SetListenerId(v string) *DeleteForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *DeleteForwardingRulesRequest) SetRegionId(v string) *DeleteForwardingRulesRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteForwardingRulesResponse) SetStatusCode(v int32) *DeleteForwardingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteForwardingRulesResponse) SetBody(v *DeleteForwardingRulesResponseBody) *DeleteForwardingRulesResponse {
	s.Body = v
	return s
}

type DeleteIpSetRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId       *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpSetRequest) SetAcceleratorId(v string) *DeleteIpSetRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteIpSetRequest) SetClientToken(v string) *DeleteIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpSetRequest) SetIpSetId(v string) *DeleteIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *DeleteIpSetRequest) SetRegionId(v string) *DeleteIpSetRequest {
	s.RegionId = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteIpSetResponse) SetStatusCode(v int32) *DeleteIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpSetResponse) SetBody(v *DeleteIpSetResponseBody) *DeleteIpSetResponse {
	s.Body = v
	return s
}

type DeleteIpSetsRequest struct {
	IpSetIds []*string `json:"IpSetIds,omitempty" xml:"IpSetIds,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpSetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpSetsRequest) SetIpSetIds(v []*string) *DeleteIpSetsRequest {
	s.IpSetIds = v
	return s
}

func (s *DeleteIpSetsRequest) SetRegionId(v string) *DeleteIpSetsRequest {
	s.RegionId = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteIpSetsResponse) SetStatusCode(v int32) *DeleteIpSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpSetsResponse) SetBody(v *DeleteIpSetsResponseBody) *DeleteIpSetsResponse {
	s.Body = v
	return s
}

type DeleteListenerRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ListenerId    *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) SetAcceleratorId(v string) *DeleteListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteListenerResponse) SetStatusCode(v int32) *DeleteListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteListenerResponse) SetBody(v *DeleteListenerResponseBody) *DeleteListenerResponse {
	s.Body = v
	return s
}

type DeleteSpareIpsRequest struct {
	AcceleratorId *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpareIps      []*string `json:"SpareIps,omitempty" xml:"SpareIps,omitempty" type:"Repeated"`
}

func (s DeleteSpareIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpareIpsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpareIpsRequest) SetAcceleratorId(v string) *DeleteSpareIpsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteSpareIpsRequest) SetClientToken(v string) *DeleteSpareIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSpareIpsRequest) SetDryRun(v bool) *DeleteSpareIpsRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteSpareIpsRequest) SetRegionId(v string) *DeleteSpareIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSpareIpsRequest) SetSpareIps(v []*string) *DeleteSpareIpsRequest {
	s.SpareIps = v
	return s
}

type DeleteSpareIpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSpareIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpareIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpareIpsResponseBody) SetRequestId(v string) *DeleteSpareIpsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSpareIpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSpareIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSpareIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpareIpsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpareIpsResponse) SetHeaders(v map[string]*string) *DeleteSpareIpsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpareIpsResponse) SetStatusCode(v int32) *DeleteSpareIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSpareIpsResponse) SetBody(v *DeleteSpareIpsResponseBody) *DeleteSpareIpsResponse {
	s.Body = v
	return s
}

type DescribeAcceleratorRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorRequest) SetAcceleratorId(v string) *DescribeAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeAcceleratorRequest) SetRegionId(v string) *DescribeAcceleratorRequest {
	s.RegionId = &v
	return s
}

type DescribeAcceleratorResponseBody struct {
	AcceleratorId               *string                                                     `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	BandwidthBillingType        *string                                                     `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	BasicBandwidthPackage       *DescribeAcceleratorResponseBodyBasicBandwidthPackage       `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	CenId                       *string                                                     `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CreateTime                  *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossDomainBandwidthPackage *DescribeAcceleratorResponseBodyCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	CrossPrivateState           *string                                                     `json:"CrossPrivateState,omitempty" xml:"CrossPrivateState,omitempty"`
	DdosId                      *string                                                     `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	Description                 *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	DnsName                     *string                                                     `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	ExpiredTime                 *int64                                                      `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceChargeType          *string                                                     `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	IpSetConfig                 *DescribeAcceleratorResponseBodyIpSetConfig                 `json:"IpSetConfig,omitempty" xml:"IpSetConfig,omitempty" type:"Struct"`
	Name                        *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId                    *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId                   *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId             *string                                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondDnsName               *string                                                     `json:"SecondDnsName,omitempty" xml:"SecondDnsName,omitempty"`
	Spec                        *string                                                     `json:"Spec,omitempty" xml:"Spec,omitempty"`
	State                       *string                                                     `json:"State,omitempty" xml:"State,omitempty"`
	Tags                        []*DescribeAcceleratorResponseBodyTags                      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBody) SetAcceleratorId(v string) *DescribeAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetBandwidthBillingType(v string) *DescribeAcceleratorResponseBody {
	s.BandwidthBillingType = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetBasicBandwidthPackage(v *DescribeAcceleratorResponseBodyBasicBandwidthPackage) *DescribeAcceleratorResponseBody {
	s.BasicBandwidthPackage = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetCenId(v string) *DescribeAcceleratorResponseBody {
	s.CenId = &v
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

func (s *DescribeAcceleratorResponseBody) SetCrossPrivateState(v string) *DescribeAcceleratorResponseBody {
	s.CrossPrivateState = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDdosId(v string) *DescribeAcceleratorResponseBody {
	s.DdosId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDescription(v string) *DescribeAcceleratorResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetDnsName(v string) *DescribeAcceleratorResponseBody {
	s.DnsName = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetExpiredTime(v int64) *DescribeAcceleratorResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetInstanceChargeType(v string) *DescribeAcceleratorResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetIpSetConfig(v *DescribeAcceleratorResponseBodyIpSetConfig) *DescribeAcceleratorResponseBody {
	s.IpSetConfig = v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetName(v string) *DescribeAcceleratorResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetRegionId(v string) *DescribeAcceleratorResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetRequestId(v string) *DescribeAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetResourceGroupId(v string) *DescribeAcceleratorResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetSecondDnsName(v string) *DescribeAcceleratorResponseBody {
	s.SecondDnsName = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetSpec(v string) *DescribeAcceleratorResponseBody {
	s.Spec = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetState(v string) *DescribeAcceleratorResponseBody {
	s.State = &v
	return s
}

func (s *DescribeAcceleratorResponseBody) SetTags(v []*DescribeAcceleratorResponseBodyTags) *DescribeAcceleratorResponseBody {
	s.Tags = v
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

type DescribeAcceleratorResponseBodyIpSetConfig struct {
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s DescribeAcceleratorResponseBodyIpSetConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyIpSetConfig) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyIpSetConfig) SetAccessMode(v string) *DescribeAcceleratorResponseBodyIpSetConfig {
	s.AccessMode = &v
	return s
}

type DescribeAcceleratorResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAcceleratorResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorResponseBodyTags) SetKey(v string) *DescribeAcceleratorResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeAcceleratorResponseBodyTags) SetValue(v string) *DescribeAcceleratorResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeAcceleratorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAcceleratorResponse) SetStatusCode(v int32) *DescribeAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAcceleratorResponse) SetBody(v *DescribeAcceleratorResponseBody) *DescribeAcceleratorResponse {
	s.Body = v
	return s
}

type DescribeAcceleratorAutoRenewAttributeRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAcceleratorAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorAutoRenewAttributeRequest) SetAcceleratorId(v string) *DescribeAcceleratorAutoRenewAttributeRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeRequest) SetRegionId(v string) *DescribeAcceleratorAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

type DescribeAcceleratorAutoRenewAttributeResponseBody struct {
	AcceleratorId     *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AutoRenew         *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	RenewalStatus     *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAcceleratorAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetAcceleratorId(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetAutoRenew(v bool) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetAutoRenewDuration(v int32) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.AutoRenewDuration = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetRenewalStatus(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeAcceleratorAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAcceleratorAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAcceleratorAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAcceleratorAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAcceleratorAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeAcceleratorAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeAcceleratorAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeResponse) SetBody(v *DescribeAcceleratorAutoRenewAttributeResponseBody) *DescribeAcceleratorAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type DescribeApplicationMonitorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeApplicationMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMonitorRequest) SetClientToken(v string) *DescribeApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeApplicationMonitorRequest) SetRegionId(v string) *DescribeApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationMonitorRequest) SetTaskId(v string) *DescribeApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

type DescribeApplicationMonitorResponseBody struct {
	AcceleratorId   *string                                              `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Address         *string                                              `json:"Address,omitempty" xml:"Address,omitempty"`
	DetectEnable    *bool                                                `json:"DetectEnable,omitempty" xml:"DetectEnable,omitempty"`
	DetectThreshold *int32                                               `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	DetectTimes     *int32                                               `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	IspCityList     []*DescribeApplicationMonitorResponseBodyIspCityList `json:"IspCityList,omitempty" xml:"IspCityList,omitempty" type:"Repeated"`
	ListenerId      *string                                              `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	OptionsJson     *string                                              `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	RegionId        *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SilenceTime     *int32                                               `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	TaskId          *string                                              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName        *string                                              `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeApplicationMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMonitorResponseBody) SetAcceleratorId(v string) *DescribeApplicationMonitorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetAddress(v string) *DescribeApplicationMonitorResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetDetectEnable(v bool) *DescribeApplicationMonitorResponseBody {
	s.DetectEnable = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetDetectThreshold(v int32) *DescribeApplicationMonitorResponseBody {
	s.DetectThreshold = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetDetectTimes(v int32) *DescribeApplicationMonitorResponseBody {
	s.DetectTimes = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetIspCityList(v []*DescribeApplicationMonitorResponseBodyIspCityList) *DescribeApplicationMonitorResponseBody {
	s.IspCityList = v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetListenerId(v string) *DescribeApplicationMonitorResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetOptionsJson(v string) *DescribeApplicationMonitorResponseBody {
	s.OptionsJson = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetRegionId(v string) *DescribeApplicationMonitorResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetRequestId(v string) *DescribeApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetSilenceTime(v int32) *DescribeApplicationMonitorResponseBody {
	s.SilenceTime = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetTaskId(v string) *DescribeApplicationMonitorResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBody) SetTaskName(v string) *DescribeApplicationMonitorResponseBody {
	s.TaskName = &v
	return s
}

type DescribeApplicationMonitorResponseBodyIspCityList struct {
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	Isp      *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	IspName  *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeApplicationMonitorResponseBodyIspCityList) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationMonitorResponseBodyIspCityList) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) SetCity(v string) *DescribeApplicationMonitorResponseBodyIspCityList {
	s.City = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) SetCityName(v string) *DescribeApplicationMonitorResponseBodyIspCityList {
	s.CityName = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) SetIsp(v string) *DescribeApplicationMonitorResponseBodyIspCityList {
	s.Isp = &v
	return s
}

func (s *DescribeApplicationMonitorResponseBodyIspCityList) SetIspName(v string) *DescribeApplicationMonitorResponseBodyIspCityList {
	s.IspName = &v
	return s
}

type DescribeApplicationMonitorResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApplicationMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMonitorResponse) SetHeaders(v map[string]*string) *DescribeApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationMonitorResponse) SetStatusCode(v int32) *DescribeApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationMonitorResponse) SetBody(v *DescribeApplicationMonitorResponseBody) *DescribeApplicationMonitorResponse {
	s.Body = v
	return s
}

type DescribeBandwidthPackageRequest struct {
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageRequest) SetBandwidthPackageId(v string) *DescribeBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeBandwidthPackageRequest) SetRegionId(v string) *DescribeBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

type DescribeBandwidthPackageResponseBody struct {
	Accelerators           []*string                                   `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	Bandwidth              *int32                                      `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageId     *string                                     `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthType          *string                                     `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	BillingType            *string                                     `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CbnGeographicRegionIdA *string                                     `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	CbnGeographicRegionIdB *string                                     `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	ChargeType             *string                                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime             *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime            *string                                     `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Name                   *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Ratio                  *int32                                      `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	RegionId               *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId              *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId        *string                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State                  *string                                     `json:"State,omitempty" xml:"State,omitempty"`
	Tags                   []*DescribeBandwidthPackageResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type                   *string                                     `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeBandwidthPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageResponseBody) SetAccelerators(v []*string) *DescribeBandwidthPackageResponseBody {
	s.Accelerators = v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidth(v int32) *DescribeBandwidthPackageResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *DescribeBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidthType(v string) *DescribeBandwidthPackageResponseBody {
	s.BandwidthType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBillingType(v string) *DescribeBandwidthPackageResponseBody {
	s.BillingType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetCbnGeographicRegionIdA(v string) *DescribeBandwidthPackageResponseBody {
	s.CbnGeographicRegionIdA = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetCbnGeographicRegionIdB(v string) *DescribeBandwidthPackageResponseBody {
	s.CbnGeographicRegionIdB = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetChargeType(v string) *DescribeBandwidthPackageResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetCreateTime(v string) *DescribeBandwidthPackageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetDescription(v string) *DescribeBandwidthPackageResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetExpiredTime(v string) *DescribeBandwidthPackageResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetName(v string) *DescribeBandwidthPackageResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRatio(v int32) *DescribeBandwidthPackageResponseBody {
	s.Ratio = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRegionId(v string) *DescribeBandwidthPackageResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRequestId(v string) *DescribeBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetResourceGroupId(v string) *DescribeBandwidthPackageResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetState(v string) *DescribeBandwidthPackageResponseBody {
	s.State = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetTags(v []*DescribeBandwidthPackageResponseBodyTags) *DescribeBandwidthPackageResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetType(v string) *DescribeBandwidthPackageResponseBody {
	s.Type = &v
	return s
}

type DescribeBandwidthPackageResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBandwidthPackageResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageResponseBodyTags) SetKey(v string) *DescribeBandwidthPackageResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBodyTags) SetValue(v string) *DescribeBandwidthPackageResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeBandwidthPackageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeBandwidthPackageResponse) SetStatusCode(v int32) *DescribeBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandwidthPackageResponse) SetBody(v *DescribeBandwidthPackageResponseBody) *DescribeBandwidthPackageResponse {
	s.Body = v
	return s
}

type DescribeBandwidthPackageAutoRenewAttributeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBandwidthPackageAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageAutoRenewAttributeRequest) SetInstanceId(v string) *DescribeBandwidthPackageAutoRenewAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeRequest) SetRegionId(v string) *DescribeBandwidthPackageAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

type DescribeBandwidthPackageAutoRenewAttributeResponseBody struct {
	AutoRenew         *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RenewalStatus     *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBandwidthPackageAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetAutoRenew(v bool) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetAutoRenewDuration(v int32) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.AutoRenewDuration = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetInstanceId(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetRenewalStatus(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBandwidthPackageAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBandwidthPackageAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBandwidthPackageAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBandwidthPackageAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeBandwidthPackageAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeBandwidthPackageAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) SetBody(v *DescribeBandwidthPackageAutoRenewAttributeResponseBody) *DescribeBandwidthPackageAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type DescribeCustomRoutingEndPointTrafficPolicyRequest struct {
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) SetPolicyId(v string) *DescribeCustomRoutingEndPointTrafficPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyRequest) SetRegionId(v string) *DescribeCustomRoutingEndPointTrafficPolicyRequest {
	s.RegionId = &v
	return s
}

type DescribeCustomRoutingEndPointTrafficPolicyResponseBody struct {
	AcceleratorId   *string                                                             `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Address         *string                                                             `json:"Address,omitempty" xml:"Address,omitempty"`
	Endpoint        *string                                                             `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EndpointGroupId *string                                                             `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId      *string                                                             `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ListenerId      *string                                                             `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	PolicyId        *string                                                             `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PortRanges      []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	RequestId       *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetAcceleratorId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetAddress(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetEndpoint(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetEndpointGroupId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetEndpointId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetListenerId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetPolicyId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetPortRanges(v []*DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.PortRanges = v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) SetRequestId(v string) *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) SetFromPort(v int32) *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges {
	s.FromPort = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges) SetToPort(v int32) *DescribeCustomRoutingEndPointTrafficPolicyResponseBodyPortRanges {
	s.ToPort = &v
	return s
}

type DescribeCustomRoutingEndPointTrafficPolicyResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomRoutingEndPointTrafficPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) SetHeaders(v map[string]*string) *DescribeCustomRoutingEndPointTrafficPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) SetStatusCode(v int32) *DescribeCustomRoutingEndPointTrafficPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) SetBody(v *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) *DescribeCustomRoutingEndPointTrafficPolicyResponse {
	s.Body = v
	return s
}

type DescribeCustomRoutingEndpointRequest struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomRoutingEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointRequest) SetEndpointId(v string) *DescribeCustomRoutingEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointRequest) SetRegionId(v string) *DescribeCustomRoutingEndpointRequest {
	s.RegionId = &v
	return s
}

type DescribeCustomRoutingEndpointResponseBody struct {
	AcceleratorId           *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Endpoint                *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EndpointGroupId         *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId              *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ListenerId              *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCustomRoutingEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetAcceleratorId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetEndpoint(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetEndpointId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetListenerId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetRequestId(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetTrafficToEndpointPolicy(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponseBody) SetType(v string) *DescribeCustomRoutingEndpointResponseBody {
	s.Type = &v
	return s
}

type DescribeCustomRoutingEndpointResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomRoutingEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomRoutingEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointResponse) SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomRoutingEndpointResponse) SetStatusCode(v int32) *DescribeCustomRoutingEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponse) SetBody(v *DescribeCustomRoutingEndpointResponseBody) *DescribeCustomRoutingEndpointResponse {
	s.Body = v
	return s
}

type DescribeCustomRoutingEndpointGroupRequest struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupRequest) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupRequest) SetRegionId(v string) *DescribeCustomRoutingEndpointGroupRequest {
	s.RegionId = &v
	return s
}

type DescribeCustomRoutingEndpointGroupResponseBody struct {
	AcceleratorId                  *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccessLogSwitch                *string   `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	Description                    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableAccessLog                *bool     `json:"EnableAccessLog,omitempty" xml:"EnableAccessLog,omitempty"`
	EndpointGroupId                *string   `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupIpList            []*string `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	EndpointGroupRegion            *string   `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointGroupUnconfirmedIpList []*string `json:"EndpointGroupUnconfirmedIpList,omitempty" xml:"EndpointGroupUnconfirmedIpList,omitempty" type:"Repeated"`
	ListenerId                     *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Name                           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId                      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsLogStoreName                *string   `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	SlsProjectName                 *string   `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	SlsRegion                      *string   `json:"SlsRegion,omitempty" xml:"SlsRegion,omitempty"`
	State                          *string   `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetAcceleratorId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetAccessLogSwitch(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.AccessLogSwitch = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetDescription(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEnableAccessLog(v bool) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EnableAccessLog = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEndpointGroupIpList(v []*string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EndpointGroupIpList = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEndpointGroupRegion(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EndpointGroupRegion = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetEndpointGroupUnconfirmedIpList(v []*string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.EndpointGroupUnconfirmedIpList = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetListenerId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetName(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetRequestId(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetSlsLogStoreName(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.SlsLogStoreName = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetSlsProjectName(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.SlsProjectName = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetSlsRegion(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.SlsRegion = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponseBody) SetState(v string) *DescribeCustomRoutingEndpointGroupResponseBody {
	s.State = &v
	return s
}

type DescribeCustomRoutingEndpointGroupResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomRoutingEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomRoutingEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupResponse) SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponse) SetStatusCode(v int32) *DescribeCustomRoutingEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponse) SetBody(v *DescribeCustomRoutingEndpointGroupResponseBody) *DescribeCustomRoutingEndpointGroupResponse {
	s.Body = v
	return s
}

type DescribeCustomRoutingEndpointGroupDestinationsRequest struct {
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) SetDestinationId(v string) *DescribeCustomRoutingEndpointGroupDestinationsRequest {
	s.DestinationId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *DescribeCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

type DescribeCustomRoutingEndpointGroupDestinationsResponseBody struct {
	AcceleratorId   *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	DestinationId   *string   `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	EndpointGroupId *string   `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	FromPort        *int32    `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ListenerId      *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Protocols       []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	RequestId       *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ToPort          *int32    `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetAcceleratorId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetDestinationId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.DestinationId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetEndpointGroupId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetFromPort(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.FromPort = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetListenerId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetProtocols(v []*string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.Protocols = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) SetToPort(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	s.ToPort = &v
	return s
}

type DescribeCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) *DescribeCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

type DescribeEndpointGroupRequest struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupRequest) SetEndpointGroupId(v string) *DescribeEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeEndpointGroupRequest) SetRegionId(v string) *DescribeEndpointGroupRequest {
	s.RegionId = &v
	return s
}

type DescribeEndpointGroupResponseBody struct {
	AcceleratorId                  *string                                                    `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccessLogSwitch                *string                                                    `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	Description                    *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableAccessLog                *bool                                                      `json:"EnableAccessLog,omitempty" xml:"EnableAccessLog,omitempty"`
	EndpointConfigurations         []*DescribeEndpointGroupResponseBodyEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupId                *string                                                    `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupIpList            []*string                                                  `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	EndpointGroupRegion            *string                                                    `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointGroupType              *string                                                    `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointGroupUnconfirmedIpList []*string                                                  `json:"EndpointGroupUnconfirmedIpList,omitempty" xml:"EndpointGroupUnconfirmedIpList,omitempty" type:"Repeated"`
	EndpointRequestProtocol        *string                                                    `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	ForwardingRuleIds              []*string                                                  `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	HealthCheckEnabled             *bool                                                      `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckIntervalSeconds     *int32                                                     `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath                *string                                                    `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort                *int32                                                     `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol            *string                                                    `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	ListenerId                     *string                                                    `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Name                           *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	PortOverrides                  []*DescribeEndpointGroupResponseBodyPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	RequestId                      *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsLogStoreName                *string                                                    `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	SlsProjectName                 *string                                                    `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	SlsRegion                      *string                                                    `json:"SlsRegion,omitempty" xml:"SlsRegion,omitempty"`
	State                          *string                                                    `json:"State,omitempty" xml:"State,omitempty"`
	Tags                           []*DescribeEndpointGroupResponseBodyTags                   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ThresholdCount                 *int32                                                     `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	TotalCount                     *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrafficPercentage              *int32                                                     `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s DescribeEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBody) SetAcceleratorId(v string) *DescribeEndpointGroupResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetAccessLogSwitch(v string) *DescribeEndpointGroupResponseBody {
	s.AccessLogSwitch = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetDescription(v string) *DescribeEndpointGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEnableAccessLog(v bool) *DescribeEndpointGroupResponseBody {
	s.EnableAccessLog = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointConfigurations(v []*DescribeEndpointGroupResponseBodyEndpointConfigurations) *DescribeEndpointGroupResponseBody {
	s.EndpointConfigurations = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupId(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupIpList(v []*string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupIpList = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupRegion(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupRegion = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupType(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupType = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupUnconfirmedIpList(v []*string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupUnconfirmedIpList = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointRequestProtocol(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetForwardingRuleIds(v []*string) *DescribeEndpointGroupResponseBody {
	s.ForwardingRuleIds = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckEnabled(v bool) *DescribeEndpointGroupResponseBody {
	s.HealthCheckEnabled = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckIntervalSeconds(v int32) *DescribeEndpointGroupResponseBody {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckPath(v string) *DescribeEndpointGroupResponseBody {
	s.HealthCheckPath = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckPort(v int32) *DescribeEndpointGroupResponseBody {
	s.HealthCheckPort = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckProtocol(v string) *DescribeEndpointGroupResponseBody {
	s.HealthCheckProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetListenerId(v string) *DescribeEndpointGroupResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetName(v string) *DescribeEndpointGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetPortOverrides(v []*DescribeEndpointGroupResponseBodyPortOverrides) *DescribeEndpointGroupResponseBody {
	s.PortOverrides = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetRequestId(v string) *DescribeEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsLogStoreName(v string) *DescribeEndpointGroupResponseBody {
	s.SlsLogStoreName = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsProjectName(v string) *DescribeEndpointGroupResponseBody {
	s.SlsProjectName = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsRegion(v string) *DescribeEndpointGroupResponseBody {
	s.SlsRegion = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetState(v string) *DescribeEndpointGroupResponseBody {
	s.State = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetTags(v []*DescribeEndpointGroupResponseBodyTags) *DescribeEndpointGroupResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetThresholdCount(v int32) *DescribeEndpointGroupResponseBody {
	s.ThresholdCount = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetTotalCount(v int32) *DescribeEndpointGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetTrafficPercentage(v int32) *DescribeEndpointGroupResponseBody {
	s.TrafficPercentage = &v
	return s
}

type DescribeEndpointGroupResponseBodyEndpointConfigurations struct {
	EnableClientIPPreservation *bool   `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	Endpoint                   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ProbePort                  *int32  `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
	ProbeProtocol              *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	Type                       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight                     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetEnableClientIPPreservation(v bool) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.EnableClientIPPreservation = &v
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

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetProbeProtocol(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.ProbeProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetType(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetWeight(v int32) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Weight = &v
	return s
}

type DescribeEndpointGroupResponseBodyPortOverrides struct {
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyPortOverrides) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) SetEndpointPort(v int32) *DescribeEndpointGroupResponseBodyPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) SetListenerPort(v int32) *DescribeEndpointGroupResponseBodyPortOverrides {
	s.ListenerPort = &v
	return s
}

type DescribeEndpointGroupResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyTags) SetKey(v string) *DescribeEndpointGroupResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyTags) SetValue(v string) *DescribeEndpointGroupResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeEndpointGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeEndpointGroupResponse) SetStatusCode(v int32) *DescribeEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointGroupResponse) SetBody(v *DescribeEndpointGroupResponseBody) *DescribeEndpointGroupResponse {
	s.Body = v
	return s
}

type DescribeIpSetRequest struct {
	IpSetId  *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpSetRequest) SetIpSetId(v string) *DescribeIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *DescribeIpSetRequest) SetRegionId(v string) *DescribeIpSetRequest {
	s.RegionId = &v
	return s
}

type DescribeIpSetResponseBody struct {
	AccelerateRegionId *string   `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	AcceleratorId      *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Bandwidth          *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	IpAddressList      []*string `json:"IpAddressList,omitempty" xml:"IpAddressList,omitempty" type:"Repeated"`
	IpSetId            *string   `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	IpVersion          *string   `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IspType            *string   `json:"IspType,omitempty" xml:"IspType,omitempty"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State              *string   `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpSetResponseBody) SetAccelerateRegionId(v string) *DescribeIpSetResponseBody {
	s.AccelerateRegionId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetAcceleratorId(v string) *DescribeIpSetResponseBody {
	s.AcceleratorId = &v
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

func (s *DescribeIpSetResponseBody) SetIpSetId(v string) *DescribeIpSetResponseBody {
	s.IpSetId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetIpVersion(v string) *DescribeIpSetResponseBody {
	s.IpVersion = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetIspType(v string) *DescribeIpSetResponseBody {
	s.IspType = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetRequestId(v string) *DescribeIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpSetResponseBody) SetState(v string) *DescribeIpSetResponseBody {
	s.State = &v
	return s
}

type DescribeIpSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeIpSetResponse) SetStatusCode(v int32) *DescribeIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpSetResponse) SetBody(v *DescribeIpSetResponseBody) *DescribeIpSetResponse {
	s.Body = v
	return s
}

type DescribeListenerRequest struct {
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerRequest) GoString() string {
	return s.String()
}

func (s *DescribeListenerRequest) SetListenerId(v string) *DescribeListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DescribeListenerRequest) SetRegionId(v string) *DescribeListenerRequest {
	s.RegionId = &v
	return s
}

type DescribeListenerResponseBody struct {
	AcceleratorId       *string                                          `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AclType             *string                                          `json:"AclType,omitempty" xml:"AclType,omitempty"`
	BackendPorts        []*DescribeListenerResponseBodyBackendPorts      `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	Certificates        []*DescribeListenerResponseBodyCertificates      `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	ClientAffinity      *string                                          `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	CreateTime          *string                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description         *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	ListenerId          *string                                          `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Name                *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges          []*DescribeListenerResponseBodyPortRanges        `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol            *string                                          `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ProxyProtocol       *bool                                            `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	RelatedAcls         []*DescribeListenerResponseBodyRelatedAcls       `json:"RelatedAcls,omitempty" xml:"RelatedAcls,omitempty" type:"Repeated"`
	RequestId           *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityPolicyId    *string                                          `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	State               *string                                          `json:"State,omitempty" xml:"State,omitempty"`
	Type                *string                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	XForwardedForConfig *DescribeListenerResponseBodyXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s DescribeListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBody) SetAcceleratorId(v string) *DescribeListenerResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetAclType(v string) *DescribeListenerResponseBody {
	s.AclType = &v
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

func (s *DescribeListenerResponseBody) SetClientAffinity(v string) *DescribeListenerResponseBody {
	s.ClientAffinity = &v
	return s
}

func (s *DescribeListenerResponseBody) SetCreateTime(v string) *DescribeListenerResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeListenerResponseBody) SetDescription(v string) *DescribeListenerResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeListenerResponseBody) SetListenerId(v string) *DescribeListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetName(v string) *DescribeListenerResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeListenerResponseBody) SetPortRanges(v []*DescribeListenerResponseBodyPortRanges) *DescribeListenerResponseBody {
	s.PortRanges = v
	return s
}

func (s *DescribeListenerResponseBody) SetProtocol(v string) *DescribeListenerResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeListenerResponseBody) SetProxyProtocol(v bool) *DescribeListenerResponseBody {
	s.ProxyProtocol = &v
	return s
}

func (s *DescribeListenerResponseBody) SetRelatedAcls(v []*DescribeListenerResponseBodyRelatedAcls) *DescribeListenerResponseBody {
	s.RelatedAcls = v
	return s
}

func (s *DescribeListenerResponseBody) SetRequestId(v string) *DescribeListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetSecurityPolicyId(v string) *DescribeListenerResponseBody {
	s.SecurityPolicyId = &v
	return s
}

func (s *DescribeListenerResponseBody) SetState(v string) *DescribeListenerResponseBody {
	s.State = &v
	return s
}

func (s *DescribeListenerResponseBody) SetType(v string) *DescribeListenerResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeListenerResponseBody) SetXForwardedForConfig(v *DescribeListenerResponseBodyXForwardedForConfig) *DescribeListenerResponseBody {
	s.XForwardedForConfig = v
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
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeListenerResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyCertificates) SetId(v string) *DescribeListenerResponseBodyCertificates {
	s.Id = &v
	return s
}

func (s *DescribeListenerResponseBodyCertificates) SetType(v string) *DescribeListenerResponseBodyCertificates {
	s.Type = &v
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

type DescribeListenerResponseBodyXForwardedForConfig struct {
	XForwardedForGaApEnabled  *bool `json:"XForwardedForGaApEnabled,omitempty" xml:"XForwardedForGaApEnabled,omitempty"`
	XForwardedForGaIdEnabled  *bool `json:"XForwardedForGaIdEnabled,omitempty" xml:"XForwardedForGaIdEnabled,omitempty"`
	XForwardedForPortEnabled  *bool `json:"XForwardedForPortEnabled,omitempty" xml:"XForwardedForPortEnabled,omitempty"`
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	XRealIpEnabled            *bool `json:"XRealIpEnabled,omitempty" xml:"XRealIpEnabled,omitempty"`
}

func (s DescribeListenerResponseBodyXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeListenerResponseBodyXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXForwardedForGaApEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XForwardedForGaApEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXForwardedForGaIdEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XForwardedForGaIdEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXForwardedForPortEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XForwardedForPortEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *DescribeListenerResponseBodyXForwardedForConfig) SetXRealIpEnabled(v bool) *DescribeListenerResponseBodyXForwardedForConfig {
	s.XRealIpEnabled = &v
	return s
}

type DescribeListenerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeListenerResponse) SetStatusCode(v int32) *DescribeListenerResponse {
	s.StatusCode = &v
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
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachDdosFromAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachDdosFromAcceleratorResponse) SetStatusCode(v int32) *DetachDdosFromAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDdosFromAcceleratorResponse) SetBody(v *DetachDdosFromAcceleratorResponseBody) *DetachDdosFromAcceleratorResponse {
	s.Body = v
	return s
}

type DetachLogStoreFromEndpointGroupRequest struct {
	AcceleratorId    *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken      *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	ListenerId       *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachLogStoreFromEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachLogStoreFromEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetAcceleratorId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetClientToken(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetEndpointGroupIds(v []*string) *DetachLogStoreFromEndpointGroupRequest {
	s.EndpointGroupIds = v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetListenerId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupRequest) SetRegionId(v string) *DetachLogStoreFromEndpointGroupRequest {
	s.RegionId = &v
	return s
}

type DetachLogStoreFromEndpointGroupResponseBody struct {
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachLogStoreFromEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachLogStoreFromEndpointGroupResponse) SetStatusCode(v int32) *DetachLogStoreFromEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupResponse) SetBody(v *DetachLogStoreFromEndpointGroupResponseBody) *DetachLogStoreFromEndpointGroupResponse {
	s.Body = v
	return s
}

type DetectApplicationMonitorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetectApplicationMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *DetectApplicationMonitorRequest) SetClientToken(v string) *DetectApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectApplicationMonitorRequest) SetRegionId(v string) *DetectApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DetectApplicationMonitorRequest) SetTaskId(v string) *DetectApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

type DetectApplicationMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectApplicationMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DetectApplicationMonitorResponseBody) SetRequestId(v string) *DetectApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

type DetectApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectApplicationMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DetectApplicationMonitorResponse) SetHeaders(v map[string]*string) *DetectApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DetectApplicationMonitorResponse) SetStatusCode(v int32) *DetectApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectApplicationMonitorResponse) SetBody(v *DetectApplicationMonitorResponseBody) *DetectApplicationMonitorResponse {
	s.Body = v
	return s
}

type DisableApplicationMonitorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DisableApplicationMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationMonitorRequest) SetClientToken(v string) *DisableApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableApplicationMonitorRequest) SetRegionId(v string) *DisableApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DisableApplicationMonitorRequest) SetTaskId(v string) *DisableApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

type DisableApplicationMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationMonitorResponseBody) SetRequestId(v string) *DisableApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

type DisableApplicationMonitorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableApplicationMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationMonitorResponse) SetHeaders(v map[string]*string) *DisableApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationMonitorResponse) SetStatusCode(v int32) *DisableApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationMonitorResponse) SetBody(v *DisableApplicationMonitorResponseBody) *DisableApplicationMonitorResponse {
	s.Body = v
	return s
}

type DissociateAclsFromListenerRequest struct {
	AclIds      []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ListenerId  *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DissociateAclsFromListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerRequest) SetAclIds(v []*string) *DissociateAclsFromListenerRequest {
	s.AclIds = v
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

func (s *DissociateAclsFromListenerRequest) SetListenerId(v string) *DissociateAclsFromListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetRegionId(v string) *DissociateAclsFromListenerRequest {
	s.RegionId = &v
	return s
}

type DissociateAclsFromListenerResponseBody struct {
	AclIds     []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	ListenerId *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateAclsFromListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponseBody) SetAclIds(v []*string) *DissociateAclsFromListenerResponseBody {
	s.AclIds = v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetListenerId(v string) *DissociateAclsFromListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetRequestId(v string) *DissociateAclsFromListenerResponseBody {
	s.RequestId = &v
	return s
}

type DissociateAclsFromListenerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DissociateAclsFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DissociateAclsFromListenerResponse) SetStatusCode(v int32) *DissociateAclsFromListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateAclsFromListenerResponse) SetBody(v *DissociateAclsFromListenerResponseBody) *DissociateAclsFromListenerResponse {
	s.Body = v
	return s
}

type DissociateAdditionalCertificatesFromListenerRequest struct {
	AcceleratorId *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Domains       []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	ListenerId    *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetAcceleratorId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetClientToken(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetDomains(v []*string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.Domains = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetListenerId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetRegionId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.RegionId = &v
	return s
}

type DissociateAdditionalCertificatesFromListenerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) SetRequestId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody {
	s.RequestId = &v
	return s
}

type DissociateAdditionalCertificatesFromListenerResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DissociateAdditionalCertificatesFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateAdditionalCertificatesFromListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponse) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetHeaders(v map[string]*string) *DissociateAdditionalCertificatesFromListenerResponse {
	s.Headers = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetStatusCode(v int32) *DissociateAdditionalCertificatesFromListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetBody(v *DissociateAdditionalCertificatesFromListenerResponseBody) *DissociateAdditionalCertificatesFromListenerResponse {
	s.Body = v
	return s
}

type EnableApplicationMonitorRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s EnableApplicationMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *EnableApplicationMonitorRequest) SetClientToken(v string) *EnableApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableApplicationMonitorRequest) SetRegionId(v string) *EnableApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *EnableApplicationMonitorRequest) SetTaskId(v string) *EnableApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

type EnableApplicationMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *EnableApplicationMonitorResponseBody) SetRequestId(v string) *EnableApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

type EnableApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableApplicationMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *EnableApplicationMonitorResponse) SetHeaders(v map[string]*string) *EnableApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *EnableApplicationMonitorResponse) SetStatusCode(v int32) *EnableApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableApplicationMonitorResponse) SetBody(v *EnableApplicationMonitorResponseBody) *EnableApplicationMonitorResponse {
	s.Body = v
	return s
}

type GetAclRequest struct {
	AclId    *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAclRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAclRequest) GoString() string {
	return s.String()
}

func (s *GetAclRequest) SetAclId(v string) *GetAclRequest {
	s.AclId = &v
	return s
}

func (s *GetAclRequest) SetRegionId(v string) *GetAclRequest {
	s.RegionId = &v
	return s
}

type GetAclResponseBody struct {
	AclEntries       []*GetAclResponseBodyAclEntries       `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	AclId            *string                               `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName          *string                               `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AclStatus        *string                               `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AddressIPVersion *string                               `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	RelatedListeners []*GetAclResponseBodyRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId  *string                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags             []*GetAclResponseBodyTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetAclResponseBody) SetAclEntries(v []*GetAclResponseBodyAclEntries) *GetAclResponseBody {
	s.AclEntries = v
	return s
}

func (s *GetAclResponseBody) SetAclId(v string) *GetAclResponseBody {
	s.AclId = &v
	return s
}

func (s *GetAclResponseBody) SetAclName(v string) *GetAclResponseBody {
	s.AclName = &v
	return s
}

func (s *GetAclResponseBody) SetAclStatus(v string) *GetAclResponseBody {
	s.AclStatus = &v
	return s
}

func (s *GetAclResponseBody) SetAddressIPVersion(v string) *GetAclResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *GetAclResponseBody) SetRelatedListeners(v []*GetAclResponseBodyRelatedListeners) *GetAclResponseBody {
	s.RelatedListeners = v
	return s
}

func (s *GetAclResponseBody) SetRequestId(v string) *GetAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAclResponseBody) SetResourceGroupId(v string) *GetAclResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetAclResponseBody) SetTags(v []*GetAclResponseBodyTags) *GetAclResponseBody {
	s.Tags = v
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
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AclType       *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	ListenerId    *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s GetAclResponseBodyRelatedListeners) String() string {
	return tea.Prettify(s)
}

func (s GetAclResponseBodyRelatedListeners) GoString() string {
	return s.String()
}

func (s *GetAclResponseBodyRelatedListeners) SetAcceleratorId(v string) *GetAclResponseBodyRelatedListeners {
	s.AcceleratorId = &v
	return s
}

func (s *GetAclResponseBodyRelatedListeners) SetAclType(v string) *GetAclResponseBodyRelatedListeners {
	s.AclType = &v
	return s
}

func (s *GetAclResponseBodyRelatedListeners) SetListenerId(v string) *GetAclResponseBodyRelatedListeners {
	s.ListenerId = &v
	return s
}

type GetAclResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAclResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetAclResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetAclResponseBodyTags) SetKey(v string) *GetAclResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetAclResponseBodyTags) SetValue(v string) *GetAclResponseBodyTags {
	s.Value = &v
	return s
}

type GetAclResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAclResponse) SetStatusCode(v int32) *GetAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAclResponse) SetBody(v *GetAclResponseBody) *GetAclResponse {
	s.Body = v
	return s
}

type GetBasicAccelerateIpRequest struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicAccelerateIpRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpRequest) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpRequest) SetAccelerateIpId(v string) *GetBasicAccelerateIpRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *GetBasicAccelerateIpRequest) SetClientToken(v string) *GetBasicAccelerateIpRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicAccelerateIpRequest) SetRegionId(v string) *GetBasicAccelerateIpRequest {
	s.RegionId = &v
	return s
}

type GetBasicAccelerateIpResponseBody struct {
	AccelerateIpAddress *string `json:"AccelerateIpAddress,omitempty" xml:"AccelerateIpAddress,omitempty"`
	AccelerateIpId      *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId       *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	IpSetId             *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicAccelerateIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpResponseBody) SetAccelerateIpAddress(v string) *GetBasicAccelerateIpResponseBody {
	s.AccelerateIpAddress = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetAccelerateIpId(v string) *GetBasicAccelerateIpResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetAcceleratorId(v string) *GetBasicAccelerateIpResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetIpSetId(v string) *GetBasicAccelerateIpResponseBody {
	s.IpSetId = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetRequestId(v string) *GetBasicAccelerateIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicAccelerateIpResponseBody) SetState(v string) *GetBasicAccelerateIpResponseBody {
	s.State = &v
	return s
}

type GetBasicAccelerateIpResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBasicAccelerateIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBasicAccelerateIpResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpResponse) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpResponse) SetHeaders(v map[string]*string) *GetBasicAccelerateIpResponse {
	s.Headers = v
	return s
}

func (s *GetBasicAccelerateIpResponse) SetStatusCode(v int32) *GetBasicAccelerateIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicAccelerateIpResponse) SetBody(v *GetBasicAccelerateIpResponseBody) *GetBasicAccelerateIpResponse {
	s.Body = v
	return s
}

type GetBasicAccelerateIpEndpointRelationRequest struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId  *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicAccelerateIpEndpointRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpEndpointRelationRequest) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetAccelerateIpId(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetAcceleratorId(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetClientToken(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetEndpointId(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.EndpointId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetRegionId(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.RegionId = &v
	return s
}

type GetBasicAccelerateIpEndpointRelationResponseBody struct {
	AccelerateIpId         *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	EndpointAddress        *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName           *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	EndpointSubAddress     *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointZoneId         *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	IpAddress              *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State                  *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicAccelerateIpEndpointRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpEndpointRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetAccelerateIpId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointAddress = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointName(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointName = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointSubAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointSubAddress = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointSubAddressType(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointSubAddressType = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointType(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointZoneId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointZoneId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetIpAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.IpAddress = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetRequestId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetState(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.State = &v
	return s
}

type GetBasicAccelerateIpEndpointRelationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBasicAccelerateIpEndpointRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBasicAccelerateIpEndpointRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpEndpointRelationResponse) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) SetHeaders(v map[string]*string) *GetBasicAccelerateIpEndpointRelationResponse {
	s.Headers = v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) SetStatusCode(v int32) *GetBasicAccelerateIpEndpointRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponse) SetBody(v *GetBasicAccelerateIpEndpointRelationResponseBody) *GetBasicAccelerateIpEndpointRelationResponse {
	s.Body = v
	return s
}

type GetBasicAccelerateIpIdleCountRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId     *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicAccelerateIpIdleCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpIdleCountRequest) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpIdleCountRequest) SetClientToken(v string) *GetBasicAccelerateIpIdleCountRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountRequest) SetIpSetId(v string) *GetBasicAccelerateIpIdleCountRequest {
	s.IpSetId = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountRequest) SetRegionId(v string) *GetBasicAccelerateIpIdleCountRequest {
	s.RegionId = &v
	return s
}

type GetBasicAccelerateIpIdleCountResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetBasicAccelerateIpIdleCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpIdleCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpIdleCountResponseBody) SetRequestId(v string) *GetBasicAccelerateIpIdleCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountResponseBody) SetTotalCount(v int64) *GetBasicAccelerateIpIdleCountResponseBody {
	s.TotalCount = &v
	return s
}

type GetBasicAccelerateIpIdleCountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBasicAccelerateIpIdleCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBasicAccelerateIpIdleCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAccelerateIpIdleCountResponse) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpIdleCountResponse) SetHeaders(v map[string]*string) *GetBasicAccelerateIpIdleCountResponse {
	s.Headers = v
	return s
}

func (s *GetBasicAccelerateIpIdleCountResponse) SetStatusCode(v int32) *GetBasicAccelerateIpIdleCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicAccelerateIpIdleCountResponse) SetBody(v *GetBasicAccelerateIpIdleCountResponseBody) *GetBasicAccelerateIpIdleCountResponse {
	s.Body = v
	return s
}

type GetBasicAcceleratorRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorRequest) SetAcceleratorId(v string) *GetBasicAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAcceleratorRequest) SetRegionId(v string) *GetBasicAcceleratorRequest {
	s.RegionId = &v
	return s
}

type GetBasicAcceleratorResponseBody struct {
	AcceleratorId               *string                                                     `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	BandwidthBillingType        *string                                                     `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	BasicBandwidthPackage       *GetBasicAcceleratorResponseBodyBasicBandwidthPackage       `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	BasicEndpointGroupId        *string                                                     `json:"BasicEndpointGroupId,omitempty" xml:"BasicEndpointGroupId,omitempty"`
	BasicIpSetId                *string                                                     `json:"BasicIpSetId,omitempty" xml:"BasicIpSetId,omitempty"`
	CenId                       *string                                                     `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CreateTime                  *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossDomainBandwidthPackage *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	CrossPrivateState           *string                                                     `json:"CrossPrivateState,omitempty" xml:"CrossPrivateState,omitempty"`
	Description                 *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime                 *int64                                                      `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceChargeType          *string                                                     `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	Name                        *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId                    *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId                   *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId             *string                                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State                       *string                                                     `json:"State,omitempty" xml:"State,omitempty"`
	Tags                        []*GetBasicAcceleratorResponseBodyTags                      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetBasicAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponseBody) SetAcceleratorId(v string) *GetBasicAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetBandwidthBillingType(v string) *GetBasicAcceleratorResponseBody {
	s.BandwidthBillingType = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetBasicBandwidthPackage(v *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) *GetBasicAcceleratorResponseBody {
	s.BasicBandwidthPackage = v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetBasicEndpointGroupId(v string) *GetBasicAcceleratorResponseBody {
	s.BasicEndpointGroupId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetBasicIpSetId(v string) *GetBasicAcceleratorResponseBody {
	s.BasicIpSetId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCenId(v string) *GetBasicAcceleratorResponseBody {
	s.CenId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCreateTime(v int64) *GetBasicAcceleratorResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCrossDomainBandwidthPackage(v *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) *GetBasicAcceleratorResponseBody {
	s.CrossDomainBandwidthPackage = v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCrossPrivateState(v string) *GetBasicAcceleratorResponseBody {
	s.CrossPrivateState = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetDescription(v string) *GetBasicAcceleratorResponseBody {
	s.Description = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetExpiredTime(v int64) *GetBasicAcceleratorResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetInstanceChargeType(v string) *GetBasicAcceleratorResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetName(v string) *GetBasicAcceleratorResponseBody {
	s.Name = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetRegionId(v string) *GetBasicAcceleratorResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetRequestId(v string) *GetBasicAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetResourceGroupId(v string) *GetBasicAcceleratorResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetState(v string) *GetBasicAcceleratorResponseBody {
	s.State = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetTags(v []*GetBasicAcceleratorResponseBodyTags) *GetBasicAcceleratorResponseBody {
	s.Tags = v
	return s
}

type GetBasicAcceleratorResponseBodyBasicBandwidthPackage struct {
	Bandwidth     *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetBasicAcceleratorResponseBodyBasicBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAcceleratorResponseBodyBasicBandwidthPackage) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) SetBandwidth(v int32) *GetBasicAcceleratorResponseBodyBasicBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) SetBandwidthType(v string) *GetBasicAcceleratorResponseBodyBasicBandwidthPackage {
	s.BandwidthType = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) SetInstanceId(v string) *GetBasicAcceleratorResponseBodyBasicBandwidthPackage {
	s.InstanceId = &v
	return s
}

type GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage struct {
	Bandwidth  *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) SetBandwidth(v int32) *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) SetInstanceId(v string) *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage {
	s.InstanceId = &v
	return s
}

type GetBasicAcceleratorResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetBasicAcceleratorResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAcceleratorResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponseBodyTags) SetKey(v string) *GetBasicAcceleratorResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyTags) SetValue(v string) *GetBasicAcceleratorResponseBodyTags {
	s.Value = &v
	return s
}

type GetBasicAcceleratorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBasicAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBasicAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponse) SetHeaders(v map[string]*string) *GetBasicAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *GetBasicAcceleratorResponse) SetStatusCode(v int32) *GetBasicAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicAcceleratorResponse) SetBody(v *GetBasicAcceleratorResponseBody) *GetBasicAcceleratorResponse {
	s.Body = v
	return s
}

type GetBasicEndpointRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointId  *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointRequest) SetClientToken(v string) *GetBasicEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicEndpointRequest) SetEndpointId(v string) *GetBasicEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *GetBasicEndpointRequest) SetRegionId(v string) *GetBasicEndpointRequest {
	s.RegionId = &v
	return s
}

type GetBasicEndpointResponseBody struct {
	AcceleratorId          *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndPointId             *string `json:"EndPointId,omitempty" xml:"EndPointId,omitempty"`
	EndpointAddress        *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointGroupId        *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointSubAddress     *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointZoneId         *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State                  *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointResponseBody) SetAcceleratorId(v string) *GetBasicEndpointResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndPointId(v string) *GetBasicEndpointResponseBody {
	s.EndPointId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointAddress(v string) *GetBasicEndpointResponseBody {
	s.EndpointAddress = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointGroupId(v string) *GetBasicEndpointResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointSubAddress(v string) *GetBasicEndpointResponseBody {
	s.EndpointSubAddress = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointSubAddressType(v string) *GetBasicEndpointResponseBody {
	s.EndpointSubAddressType = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointType(v string) *GetBasicEndpointResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointZoneId(v string) *GetBasicEndpointResponseBody {
	s.EndpointZoneId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetName(v string) *GetBasicEndpointResponseBody {
	s.Name = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetRequestId(v string) *GetBasicEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetState(v string) *GetBasicEndpointResponseBody {
	s.State = &v
	return s
}

type GetBasicEndpointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBasicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBasicEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointResponse) SetHeaders(v map[string]*string) *GetBasicEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetBasicEndpointResponse) SetStatusCode(v int32) *GetBasicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicEndpointResponse) SetBody(v *GetBasicEndpointResponseBody) *GetBasicEndpointResponse {
	s.Body = v
	return s
}

type GetBasicEndpointGroupRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointGroupRequest) SetClientToken(v string) *GetBasicEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicEndpointGroupRequest) SetEndpointGroupId(v string) *GetBasicEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *GetBasicEndpointGroupRequest) SetRegionId(v string) *GetBasicEndpointGroupRequest {
	s.RegionId = &v
	return s
}

type GetBasicEndpointGroupResponseBody struct {
	AcceleratorId       *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointAddress     *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointGroupId     *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointSubAddress  *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointType        *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointGroupResponseBody) SetAcceleratorId(v string) *GetBasicEndpointGroupResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetDescription(v string) *GetBasicEndpointGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointAddress(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointAddress = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointGroupId(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointGroupRegion(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointGroupRegion = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointSubAddress(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointSubAddress = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointType(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetName(v string) *GetBasicEndpointGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetRequestId(v string) *GetBasicEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetState(v string) *GetBasicEndpointGroupResponseBody {
	s.State = &v
	return s
}

type GetBasicEndpointGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBasicEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBasicEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointGroupResponse) SetHeaders(v map[string]*string) *GetBasicEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *GetBasicEndpointGroupResponse) SetStatusCode(v int32) *GetBasicEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicEndpointGroupResponse) SetBody(v *GetBasicEndpointGroupResponseBody) *GetBasicEndpointGroupResponse {
	s.Body = v
	return s
}

type GetBasicIpSetRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId     *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicIpSetRequest) GoString() string {
	return s.String()
}

func (s *GetBasicIpSetRequest) SetClientToken(v string) *GetBasicIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicIpSetRequest) SetIpSetId(v string) *GetBasicIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *GetBasicIpSetRequest) SetRegionId(v string) *GetBasicIpSetRequest {
	s.RegionId = &v
	return s
}

type GetBasicIpSetResponseBody struct {
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	AcceleratorId      *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Bandwidth          *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	IpAddress          *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	IpSetId            *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	IpVersion          *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IspType            *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBasicIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicIpSetResponseBody) SetAccelerateRegionId(v string) *GetBasicIpSetResponseBody {
	s.AccelerateRegionId = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetAcceleratorId(v string) *GetBasicIpSetResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetBandwidth(v int64) *GetBasicIpSetResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetIpAddress(v string) *GetBasicIpSetResponseBody {
	s.IpAddress = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetIpSetId(v string) *GetBasicIpSetResponseBody {
	s.IpSetId = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetIpVersion(v string) *GetBasicIpSetResponseBody {
	s.IpVersion = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetIspType(v string) *GetBasicIpSetResponseBody {
	s.IspType = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetRequestId(v string) *GetBasicIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicIpSetResponseBody) SetState(v string) *GetBasicIpSetResponseBody {
	s.State = &v
	return s
}

type GetBasicIpSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBasicIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBasicIpSetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicIpSetResponse) GoString() string {
	return s.String()
}

func (s *GetBasicIpSetResponse) SetHeaders(v map[string]*string) *GetBasicIpSetResponse {
	s.Headers = v
	return s
}

func (s *GetBasicIpSetResponse) SetStatusCode(v int32) *GetBasicIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicIpSetResponse) SetBody(v *GetBasicIpSetResponseBody) *GetBasicIpSetResponse {
	s.Body = v
	return s
}

type GetHealthStatusRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ListenerId    *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *GetHealthStatusRequest) SetAcceleratorId(v string) *GetHealthStatusRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetHealthStatusRequest) SetClientToken(v string) *GetHealthStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHealthStatusRequest) SetDryRun(v bool) *GetHealthStatusRequest {
	s.DryRun = &v
	return s
}

func (s *GetHealthStatusRequest) SetListenerId(v string) *GetHealthStatusRequest {
	s.ListenerId = &v
	return s
}

func (s *GetHealthStatusRequest) SetRegionId(v string) *GetHealthStatusRequest {
	s.RegionId = &v
	return s
}

type GetHealthStatusResponseBody struct {
	EndpointGroups []*GetHealthStatusResponseBodyEndpointGroups `json:"EndpointGroups,omitempty" xml:"EndpointGroups,omitempty" type:"Repeated"`
	HealthStatus   *string                                      `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	ListenerId     *string                                      `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthStatusResponseBody) SetEndpointGroups(v []*GetHealthStatusResponseBodyEndpointGroups) *GetHealthStatusResponseBody {
	s.EndpointGroups = v
	return s
}

func (s *GetHealthStatusResponseBody) SetHealthStatus(v string) *GetHealthStatusResponseBody {
	s.HealthStatus = &v
	return s
}

func (s *GetHealthStatusResponseBody) SetListenerId(v string) *GetHealthStatusResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetHealthStatusResponseBody) SetRequestId(v string) *GetHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetHealthStatusResponseBodyEndpointGroups struct {
	EndpointGroupId   *string                                               `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupType *string                                               `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	Endpoints         []*GetHealthStatusResponseBodyEndpointGroupsEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	ForwardingRuleIds []*string                                             `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	HealthStatus      *string                                               `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
}

func (s GetHealthStatusResponseBodyEndpointGroups) String() string {
	return tea.Prettify(s)
}

func (s GetHealthStatusResponseBodyEndpointGroups) GoString() string {
	return s.String()
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetEndpointGroupId(v string) *GetHealthStatusResponseBodyEndpointGroups {
	s.EndpointGroupId = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetEndpointGroupType(v string) *GetHealthStatusResponseBodyEndpointGroups {
	s.EndpointGroupType = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetEndpoints(v []*GetHealthStatusResponseBodyEndpointGroupsEndpoints) *GetHealthStatusResponseBodyEndpointGroups {
	s.Endpoints = v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetForwardingRuleIds(v []*string) *GetHealthStatusResponseBodyEndpointGroups {
	s.ForwardingRuleIds = v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroups) SetHealthStatus(v string) *GetHealthStatusResponseBodyEndpointGroups {
	s.HealthStatus = &v
	return s
}

type GetHealthStatusResponseBodyEndpointGroupsEndpoints struct {
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EndpointId   *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	HealthDetail *string `json:"HealthDetail,omitempty" xml:"HealthDetail,omitempty"`
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	Port         *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetHealthStatusResponseBodyEndpointGroupsEndpoints) String() string {
	return tea.Prettify(s)
}

func (s GetHealthStatusResponseBodyEndpointGroupsEndpoints) GoString() string {
	return s.String()
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetAddress(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.Address = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetEndpointId(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.EndpointId = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetHealthDetail(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.HealthDetail = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetHealthStatus(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.HealthStatus = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetPort(v int64) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.Port = &v
	return s
}

func (s *GetHealthStatusResponseBodyEndpointGroupsEndpoints) SetType(v string) *GetHealthStatusResponseBodyEndpointGroupsEndpoints {
	s.Type = &v
	return s
}

type GetHealthStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *GetHealthStatusResponse) SetHeaders(v map[string]*string) *GetHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *GetHealthStatusResponse) SetStatusCode(v int32) *GetHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHealthStatusResponse) SetBody(v *GetHealthStatusResponseBody) *GetHealthStatusResponse {
	s.Body = v
	return s
}

type GetSpareIpRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpareIp       *string `json:"SpareIp,omitempty" xml:"SpareIp,omitempty"`
}

func (s GetSpareIpRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpareIpRequest) GoString() string {
	return s.String()
}

func (s *GetSpareIpRequest) SetAcceleratorId(v string) *GetSpareIpRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetSpareIpRequest) SetClientToken(v string) *GetSpareIpRequest {
	s.ClientToken = &v
	return s
}

func (s *GetSpareIpRequest) SetDryRun(v bool) *GetSpareIpRequest {
	s.DryRun = &v
	return s
}

func (s *GetSpareIpRequest) SetRegionId(v string) *GetSpareIpRequest {
	s.RegionId = &v
	return s
}

func (s *GetSpareIpRequest) SetSpareIp(v string) *GetSpareIpRequest {
	s.SpareIp = &v
	return s
}

type GetSpareIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetSpareIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpareIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpareIpResponseBody) SetRequestId(v string) *GetSpareIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpareIpResponseBody) SetState(v string) *GetSpareIpResponseBody {
	s.State = &v
	return s
}

type GetSpareIpResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSpareIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpareIpResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpareIpResponse) GoString() string {
	return s.String()
}

func (s *GetSpareIpResponse) SetHeaders(v map[string]*string) *GetSpareIpResponse {
	s.Headers = v
	return s
}

func (s *GetSpareIpResponse) SetStatusCode(v int32) *GetSpareIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpareIpResponse) SetBody(v *GetSpareIpResponseBody) *GetSpareIpResponse {
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
	Areas     []*ListAccelerateAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccelerateAreasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccelerateAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponseBody) SetAreas(v []*ListAccelerateAreasResponseBodyAreas) *ListAccelerateAreasResponseBody {
	s.Areas = v
	return s
}

func (s *ListAccelerateAreasResponseBody) SetRequestId(v string) *ListAccelerateAreasResponseBody {
	s.RequestId = &v
	return s
}

type ListAccelerateAreasResponseBodyAreas struct {
	AreaId     *string                                           `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	LocalName  *string                                           `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionList []*ListAccelerateAreasResponseBodyAreasRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s ListAccelerateAreasResponseBodyAreas) String() string {
	return tea.Prettify(s)
}

func (s ListAccelerateAreasResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasResponseBodyAreas) SetAreaId(v string) *ListAccelerateAreasResponseBodyAreas {
	s.AreaId = &v
	return s
}

func (s *ListAccelerateAreasResponseBodyAreas) SetLocalName(v string) *ListAccelerateAreasResponseBodyAreas {
	s.LocalName = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccelerateAreasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAccelerateAreasResponse) SetStatusCode(v int32) *ListAccelerateAreasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccelerateAreasResponse) SetBody(v *ListAccelerateAreasResponseBody) *ListAccelerateAreasResponse {
	s.Body = v
	return s
}

type ListAcceleratorsRequest struct {
	AcceleratorId   *string                       `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	PageNumber      *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State           *string                       `json:"State,omitempty" xml:"State,omitempty"`
	Tag             []*ListAcceleratorsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAcceleratorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsRequest) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsRequest) SetAcceleratorId(v string) *ListAcceleratorsRequest {
	s.AcceleratorId = &v
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

func (s *ListAcceleratorsRequest) SetRegionId(v string) *ListAcceleratorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAcceleratorsRequest) SetResourceGroupId(v string) *ListAcceleratorsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAcceleratorsRequest) SetState(v string) *ListAcceleratorsRequest {
	s.State = &v
	return s
}

func (s *ListAcceleratorsRequest) SetTag(v []*ListAcceleratorsRequestTag) *ListAcceleratorsRequest {
	s.Tag = v
	return s
}

type ListAcceleratorsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAcceleratorsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsRequestTag) SetKey(v string) *ListAcceleratorsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAcceleratorsRequestTag) SetValue(v string) *ListAcceleratorsRequestTag {
	s.Value = &v
	return s
}

type ListAcceleratorsResponseBody struct {
	Accelerators []*ListAcceleratorsResponseBodyAccelerators `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	PageNumber   *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAcceleratorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBody) SetAccelerators(v []*ListAcceleratorsResponseBodyAccelerators) *ListAcceleratorsResponseBody {
	s.Accelerators = v
	return s
}

func (s *ListAcceleratorsResponseBody) SetPageNumber(v int32) *ListAcceleratorsResponseBody {
	s.PageNumber = &v
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

func (s *ListAcceleratorsResponseBody) SetTotalCount(v int32) *ListAcceleratorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAcceleratorsResponseBodyAccelerators struct {
	AcceleratorId               *string                                                              `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Bandwidth                   *int32                                                               `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthBillingType        *string                                                              `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	BasicBandwidthPackage       *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage       `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	CenId                       *string                                                              `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CreateTime                  *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossDomainBandwidthPackage *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	DdosId                      *string                                                              `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	Description                 *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DnsName                     *string                                                              `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	ExpiredTime                 *int64                                                               `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceChargeType          *string                                                              `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	IpSetConfig                 *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig                 `json:"IpSetConfig,omitempty" xml:"IpSetConfig,omitempty" type:"Struct"`
	Name                        *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId                    *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId             *string                                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondDnsName               *string                                                              `json:"SecondDnsName,omitempty" xml:"SecondDnsName,omitempty"`
	Spec                        *string                                                              `json:"Spec,omitempty" xml:"Spec,omitempty"`
	State                       *string                                                              `json:"State,omitempty" xml:"State,omitempty"`
	Tags                        []*ListAcceleratorsResponseBodyAcceleratorsTags                      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type                        *string                                                              `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAcceleratorsResponseBodyAccelerators) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAccelerators) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetAcceleratorId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.AcceleratorId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetBandwidth(v int32) *ListAcceleratorsResponseBodyAccelerators {
	s.Bandwidth = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetBandwidthBillingType(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.BandwidthBillingType = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetBasicBandwidthPackage(v *ListAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) *ListAcceleratorsResponseBodyAccelerators {
	s.BasicBandwidthPackage = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCenId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.CenId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCreateTime(v int64) *ListAcceleratorsResponseBodyAccelerators {
	s.CreateTime = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetCrossDomainBandwidthPackage(v *ListAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) *ListAcceleratorsResponseBodyAccelerators {
	s.CrossDomainBandwidthPackage = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDdosId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.DdosId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDescription(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Description = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetDnsName(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.DnsName = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetExpiredTime(v int64) *ListAcceleratorsResponseBodyAccelerators {
	s.ExpiredTime = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetInstanceChargeType(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetIpSetConfig(v *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) *ListAcceleratorsResponseBodyAccelerators {
	s.IpSetConfig = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetName(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Name = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetRegionId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.RegionId = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetResourceGroupId(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.ResourceGroupId = &v
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

func (s *ListAcceleratorsResponseBodyAccelerators) SetTags(v []*ListAcceleratorsResponseBodyAcceleratorsTags) *ListAcceleratorsResponseBodyAccelerators {
	s.Tags = v
	return s
}

func (s *ListAcceleratorsResponseBodyAccelerators) SetType(v string) *ListAcceleratorsResponseBodyAccelerators {
	s.Type = &v
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

type ListAcceleratorsResponseBodyAcceleratorsIpSetConfig struct {
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig) SetAccessMode(v string) *ListAcceleratorsResponseBodyAcceleratorsIpSetConfig {
	s.AccessMode = &v
	return s
}

type ListAcceleratorsResponseBodyAcceleratorsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAcceleratorsResponseBodyAcceleratorsTags) String() string {
	return tea.Prettify(s)
}

func (s ListAcceleratorsResponseBodyAcceleratorsTags) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponseBodyAcceleratorsTags) SetKey(v string) *ListAcceleratorsResponseBodyAcceleratorsTags {
	s.Key = &v
	return s
}

func (s *ListAcceleratorsResponseBodyAcceleratorsTags) SetValue(v string) *ListAcceleratorsResponseBodyAcceleratorsTags {
	s.Value = &v
	return s
}

type ListAcceleratorsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAcceleratorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAcceleratorsResponse) SetStatusCode(v int32) *ListAcceleratorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAcceleratorsResponse) SetBody(v *ListAcceleratorsResponseBody) *ListAcceleratorsResponse {
	s.Body = v
	return s
}

type ListAclsRequest struct {
	AclIds          []*string             `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	AclName         *string               `json:"AclName,omitempty" xml:"AclName,omitempty"`
	ClientToken     *string               `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MaxResults      *int32                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId        *string               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListAclsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAclsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclsRequest) GoString() string {
	return s.String()
}

func (s *ListAclsRequest) SetAclIds(v []*string) *ListAclsRequest {
	s.AclIds = v
	return s
}

func (s *ListAclsRequest) SetAclName(v string) *ListAclsRequest {
	s.AclName = &v
	return s
}

func (s *ListAclsRequest) SetClientToken(v string) *ListAclsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAclsRequest) SetMaxResults(v int32) *ListAclsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAclsRequest) SetNextToken(v string) *ListAclsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAclsRequest) SetRegionId(v string) *ListAclsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAclsRequest) SetResourceGroupId(v string) *ListAclsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAclsRequest) SetTag(v []*ListAclsRequestTag) *ListAclsRequest {
	s.Tag = v
	return s
}

type ListAclsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAclsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListAclsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAclsRequestTag) SetKey(v string) *ListAclsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAclsRequestTag) SetValue(v string) *ListAclsRequestTag {
	s.Value = &v
	return s
}

type ListAclsResponseBody struct {
	Acls       []*ListAclsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	MaxResults *int32                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAclsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBody) SetAcls(v []*ListAclsResponseBodyAcls) *ListAclsResponseBody {
	s.Acls = v
	return s
}

func (s *ListAclsResponseBody) SetMaxResults(v int32) *ListAclsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAclsResponseBody) SetNextToken(v string) *ListAclsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAclsResponseBody) SetRequestId(v string) *ListAclsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclsResponseBody) SetTotalCount(v int32) *ListAclsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAclsResponseBodyAcls struct {
	AclId            *string                         `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName          *string                         `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AclStatus        *string                         `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AddressIPVersion *string                         `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	ResourceGroupId  *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags             []*ListAclsResponseBodyAclsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListAclsResponseBodyAcls) SetAclStatus(v string) *ListAclsResponseBodyAcls {
	s.AclStatus = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAddressIPVersion(v string) *ListAclsResponseBodyAcls {
	s.AddressIPVersion = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetResourceGroupId(v string) *ListAclsResponseBodyAcls {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetTags(v []*ListAclsResponseBodyAclsTags) *ListAclsResponseBodyAcls {
	s.Tags = v
	return s
}

type ListAclsResponseBodyAclsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAclsResponseBodyAclsTags) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBodyAclsTags) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBodyAclsTags) SetKey(v string) *ListAclsResponseBodyAclsTags {
	s.Key = &v
	return s
}

func (s *ListAclsResponseBodyAclsTags) SetValue(v string) *ListAclsResponseBodyAclsTags {
	s.Value = &v
	return s
}

type ListAclsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAclsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAclsResponse) SetStatusCode(v int32) *ListAclsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclsResponse) SetBody(v *ListAclsResponseBody) *ListAclsResponse {
	s.Body = v
	return s
}

type ListApplicationMonitorRequest struct {
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s ListApplicationMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorRequest) SetPageNumber(v int32) *ListApplicationMonitorRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationMonitorRequest) SetPageSize(v int32) *ListApplicationMonitorRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMonitorRequest) SetRegionId(v string) *ListApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationMonitorRequest) SetSearchValue(v string) *ListApplicationMonitorRequest {
	s.SearchValue = &v
	return s
}

type ListApplicationMonitorResponseBody struct {
	ApplicationMonitors []*ListApplicationMonitorResponseBodyApplicationMonitors `json:"ApplicationMonitors,omitempty" xml:"ApplicationMonitors,omitempty" type:"Repeated"`
	PageNumber          *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorResponseBody) SetApplicationMonitors(v []*ListApplicationMonitorResponseBodyApplicationMonitors) *ListApplicationMonitorResponseBody {
	s.ApplicationMonitors = v
	return s
}

func (s *ListApplicationMonitorResponseBody) SetPageNumber(v int32) *ListApplicationMonitorResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationMonitorResponseBody) SetPageSize(v int32) *ListApplicationMonitorResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMonitorResponseBody) SetRequestId(v string) *ListApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationMonitorResponseBody) SetTotalCount(v int32) *ListApplicationMonitorResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationMonitorResponseBodyApplicationMonitors struct {
	AcceleratorId   *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	DetectEnable    *bool   `json:"DetectEnable,omitempty" xml:"DetectEnable,omitempty"`
	DetectThreshold *int32  `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	DetectTimes     *int32  `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	OptionsJson     *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	SilenceTime     *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListApplicationMonitorResponseBodyApplicationMonitors) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationMonitorResponseBodyApplicationMonitors) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetAcceleratorId(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.AcceleratorId = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetAddress(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.Address = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetDetectEnable(v bool) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.DetectEnable = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetDetectThreshold(v int32) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.DetectThreshold = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetDetectTimes(v int32) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.DetectTimes = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetListenerId(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.ListenerId = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetOptionsJson(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.OptionsJson = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetSilenceTime(v int32) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.SilenceTime = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetState(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.State = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetTaskId(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.TaskId = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetTaskName(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.TaskName = &v
	return s
}

type ListApplicationMonitorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorResponse) SetHeaders(v map[string]*string) *ListApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationMonitorResponse) SetStatusCode(v int32) *ListApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationMonitorResponse) SetBody(v *ListApplicationMonitorResponseBody) *ListApplicationMonitorResponse {
	s.Body = v
	return s
}

type ListApplicationMonitorDetectResultRequest struct {
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListApplicationMonitorDetectResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationMonitorDetectResultRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorDetectResultRequest) SetBeginTime(v int64) *ListApplicationMonitorDetectResultRequest {
	s.BeginTime = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetEndTime(v int64) *ListApplicationMonitorDetectResultRequest {
	s.EndTime = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetPageNumber(v int32) *ListApplicationMonitorDetectResultRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetPageSize(v int32) *ListApplicationMonitorDetectResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetRegionId(v string) *ListApplicationMonitorDetectResultRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultRequest) SetTaskId(v string) *ListApplicationMonitorDetectResultRequest {
	s.TaskId = &v
	return s
}

type ListApplicationMonitorDetectResultResponseBody struct {
	ApplicationMonitorDetectResultList []*ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList `json:"ApplicationMonitorDetectResultList,omitempty" xml:"ApplicationMonitorDetectResultList,omitempty" type:"Repeated"`
	PageNumber                         *int32                                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                           *int32                                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                          *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                         *int32                                                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationMonitorDetectResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationMonitorDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetApplicationMonitorDetectResultList(v []*ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) *ListApplicationMonitorDetectResultResponseBody {
	s.ApplicationMonitorDetectResultList = v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetPageNumber(v int32) *ListApplicationMonitorDetectResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetPageSize(v int32) *ListApplicationMonitorDetectResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetRequestId(v string) *ListApplicationMonitorDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBody) SetTotalCount(v int32) *ListApplicationMonitorDetectResultResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Detail        *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	DetectTime    *string `json:"DetectTime,omitempty" xml:"DetectTime,omitempty"`
	DiagStatus    *string `json:"DiagStatus,omitempty" xml:"DiagStatus,omitempty"`
	ListenerId    *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Port          *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol      *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	StatusCode    *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetAcceleratorId(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.AcceleratorId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetContent(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.Content = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetDetail(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.Detail = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetDetectTime(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.DetectTime = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetDiagStatus(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.DiagStatus = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetListenerId(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.ListenerId = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetPort(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.Port = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetProtocol(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.Protocol = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetStatusCode(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList) SetTaskId(v string) *ListApplicationMonitorDetectResultResponseBodyApplicationMonitorDetectResultList {
	s.TaskId = &v
	return s
}

type ListApplicationMonitorDetectResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApplicationMonitorDetectResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationMonitorDetectResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationMonitorDetectResultResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorDetectResultResponse) SetHeaders(v map[string]*string) *ListApplicationMonitorDetectResultResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationMonitorDetectResultResponse) SetStatusCode(v int32) *ListApplicationMonitorDetectResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationMonitorDetectResultResponse) SetBody(v *ListApplicationMonitorDetectResultResponseBody) *ListApplicationMonitorDetectResultResponse {
	s.Body = v
	return s
}

type ListAvailableAccelerateAreasRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAvailableAccelerateAreasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableAccelerateAreasRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasRequest) SetAcceleratorId(v string) *ListAvailableAccelerateAreasRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListAvailableAccelerateAreasRequest) SetRegionId(v string) *ListAvailableAccelerateAreasRequest {
	s.RegionId = &v
	return s
}

type ListAvailableAccelerateAreasResponseBody struct {
	Areas     []*ListAvailableAccelerateAreasResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAvailableAccelerateAreasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponseBody) SetAreas(v []*ListAvailableAccelerateAreasResponseBodyAreas) *ListAvailableAccelerateAreasResponseBody {
	s.Areas = v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBody) SetRequestId(v string) *ListAvailableAccelerateAreasResponseBody {
	s.RequestId = &v
	return s
}

type ListAvailableAccelerateAreasResponseBodyAreas struct {
	AreaId     *string                                                    `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	LocalName  *string                                                    `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionList []*ListAvailableAccelerateAreasResponseBodyAreasRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s ListAvailableAccelerateAreasResponseBodyAreas) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableAccelerateAreasResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) SetAreaId(v string) *ListAvailableAccelerateAreasResponseBodyAreas {
	s.AreaId = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponseBodyAreas) SetLocalName(v string) *ListAvailableAccelerateAreasResponseBodyAreas {
	s.LocalName = &v
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAvailableAccelerateAreasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAvailableAccelerateAreasResponse) SetStatusCode(v int32) *ListAvailableAccelerateAreasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableAccelerateAreasResponse) SetBody(v *ListAvailableAccelerateAreasResponseBody) *ListAvailableAccelerateAreasResponse {
	s.Body = v
	return s
}

type ListAvailableBusiRegionsRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAvailableBusiRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBusiRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsRequest) SetAcceleratorId(v string) *ListAvailableBusiRegionsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListAvailableBusiRegionsRequest) SetRegionId(v string) *ListAvailableBusiRegionsRequest {
	s.RegionId = &v
	return s
}

type ListAvailableBusiRegionsResponseBody struct {
	Regions   []*ListAvailableBusiRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAvailableBusiRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableBusiRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsResponseBody) SetRegions(v []*ListAvailableBusiRegionsResponseBodyRegions) *ListAvailableBusiRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListAvailableBusiRegionsResponseBody) SetRequestId(v string) *ListAvailableBusiRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListAvailableBusiRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	Pop       *bool   `json:"Pop,omitempty" xml:"Pop,omitempty"`
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

func (s *ListAvailableBusiRegionsResponseBodyRegions) SetPop(v bool) *ListAvailableBusiRegionsResponseBodyRegions {
	s.Pop = &v
	return s
}

func (s *ListAvailableBusiRegionsResponseBodyRegions) SetRegionId(v string) *ListAvailableBusiRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListAvailableBusiRegionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAvailableBusiRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListAvailableBusiRegionsResponse) SetStatusCode(v int32) *ListAvailableBusiRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableBusiRegionsResponse) SetBody(v *ListAvailableBusiRegionsResponseBody) *ListAvailableBusiRegionsResponse {
	s.Body = v
	return s
}

type ListBandwidthPackagesRequest struct {
	BandwidthPackageId *string                            `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	PageNumber         *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId           *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId    *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State              *string                            `json:"State,omitempty" xml:"State,omitempty"`
	Tag                []*ListBandwidthPackagesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Type               *string                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBandwidthPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesRequest) SetBandwidthPackageId(v string) *ListBandwidthPackagesRequest {
	s.BandwidthPackageId = &v
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

func (s *ListBandwidthPackagesRequest) SetRegionId(v string) *ListBandwidthPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetResourceGroupId(v string) *ListBandwidthPackagesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetState(v string) *ListBandwidthPackagesRequest {
	s.State = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetTag(v []*ListBandwidthPackagesRequestTag) *ListBandwidthPackagesRequest {
	s.Tag = v
	return s
}

func (s *ListBandwidthPackagesRequest) SetType(v string) *ListBandwidthPackagesRequest {
	s.Type = &v
	return s
}

type ListBandwidthPackagesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBandwidthPackagesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesRequestTag) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesRequestTag) SetKey(v string) *ListBandwidthPackagesRequestTag {
	s.Key = &v
	return s
}

func (s *ListBandwidthPackagesRequestTag) SetValue(v string) *ListBandwidthPackagesRequestTag {
	s.Value = &v
	return s
}

type ListBandwidthPackagesResponseBody struct {
	BandwidthPackages []*ListBandwidthPackagesResponseBodyBandwidthPackages `json:"BandwidthPackages,omitempty" xml:"BandwidthPackages,omitempty" type:"Repeated"`
	PageNumber        *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBandwidthPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponseBody) SetBandwidthPackages(v []*ListBandwidthPackagesResponseBodyBandwidthPackages) *ListBandwidthPackagesResponseBody {
	s.BandwidthPackages = v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetPageNumber(v int32) *ListBandwidthPackagesResponseBody {
	s.PageNumber = &v
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

func (s *ListBandwidthPackagesResponseBody) SetTotalCount(v int32) *ListBandwidthPackagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListBandwidthPackagesResponseBodyBandwidthPackages struct {
	Accelerators           []*string                                                 `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	Bandwidth              *int32                                                    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageId     *string                                                   `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthType          *string                                                   `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	BillingType            *string                                                   `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CbnGeographicRegionIdA *string                                                   `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	CbnGeographicRegionIdB *string                                                   `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	ChargeType             *string                                                   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime             *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime            *string                                                   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Name                   *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Ratio                  *int32                                                    `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	RegionId               *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId        *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State                  *string                                                   `json:"State,omitempty" xml:"State,omitempty"`
	Tags                   []*ListBandwidthPackagesResponseBodyBandwidthPackagesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type                   *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackages) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackages) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetAccelerators(v []*string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Accelerators = v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidth(v int32) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Bandwidth = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidthPackageId(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidthType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BandwidthType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBillingType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BillingType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCbnGeographicRegionIdA(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CbnGeographicRegionIdA = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCbnGeographicRegionIdB(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CbnGeographicRegionIdB = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetChargeType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.ChargeType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCreateTime(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CreateTime = &v
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

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetName(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Name = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetRatio(v int32) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Ratio = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetRegionId(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetResourceGroupId(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetState(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.State = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetTags(v []*ListBandwidthPackagesResponseBodyBandwidthPackagesTags) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Tags = v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Type = &v
	return s
}

type ListBandwidthPackagesResponseBodyBandwidthPackagesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackagesTags) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackagesTags) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackagesTags) SetKey(v string) *ListBandwidthPackagesResponseBodyBandwidthPackagesTags {
	s.Key = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackagesTags) SetValue(v string) *ListBandwidthPackagesResponseBodyBandwidthPackagesTags {
	s.Value = &v
	return s
}

type ListBandwidthPackagesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBandwidthPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListBandwidthPackagesResponse) SetStatusCode(v int32) *ListBandwidthPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBandwidthPackagesResponse) SetBody(v *ListBandwidthPackagesResponseBody) *ListBandwidthPackagesResponse {
	s.Body = v
	return s
}

type ListBandwidthackagesRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBandwidthackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthackagesRequest) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesRequest) SetPageNumber(v int32) *ListBandwidthackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthackagesRequest) SetPageSize(v int32) *ListBandwidthackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBandwidthackagesRequest) SetRegionId(v string) *ListBandwidthackagesRequest {
	s.RegionId = &v
	return s
}

type ListBandwidthackagesResponseBody struct {
	BandwidthPackages []*ListBandwidthackagesResponseBodyBandwidthPackages `json:"BandwidthPackages,omitempty" xml:"BandwidthPackages,omitempty" type:"Repeated"`
	PageNumber        *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBandwidthackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponseBody) SetBandwidthPackages(v []*ListBandwidthackagesResponseBodyBandwidthPackages) *ListBandwidthackagesResponseBody {
	s.BandwidthPackages = v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetPageNumber(v int32) *ListBandwidthackagesResponseBody {
	s.PageNumber = &v
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

func (s *ListBandwidthackagesResponseBody) SetTotalCount(v int32) *ListBandwidthackagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListBandwidthackagesResponseBodyBandwidthPackages struct {
	Accelerators       []*string                                                `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	Bandwidth          *int32                                                   `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageId *string                                                  `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	ChargeType         *string                                                  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime         *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description        *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime        *string                                                  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Name               *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId           *string                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId    *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State              *string                                                  `json:"State,omitempty" xml:"State,omitempty"`
	Tags               []*ListBandwidthackagesResponseBodyBandwidthPackagesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListBandwidthackagesResponseBodyBandwidthPackages) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthackagesResponseBodyBandwidthPackages) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetAccelerators(v []*string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Accelerators = v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetBandwidth(v int32) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Bandwidth = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetBandwidthPackageId(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetChargeType(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.ChargeType = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetCreateTime(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.CreateTime = &v
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

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetName(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Name = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetRegionId(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetResourceGroupId(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetState(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.State = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetTags(v []*ListBandwidthackagesResponseBodyBandwidthPackagesTags) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Tags = v
	return s
}

type ListBandwidthackagesResponseBodyBandwidthPackagesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBandwidthackagesResponseBodyBandwidthPackagesTags) String() string {
	return tea.Prettify(s)
}

func (s ListBandwidthackagesResponseBodyBandwidthPackagesTags) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackagesTags) SetKey(v string) *ListBandwidthackagesResponseBodyBandwidthPackagesTags {
	s.Key = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackagesTags) SetValue(v string) *ListBandwidthackagesResponseBodyBandwidthPackagesTags {
	s.Value = &v
	return s
}

type ListBandwidthackagesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBandwidthackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListBandwidthackagesResponse) SetStatusCode(v int32) *ListBandwidthackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBandwidthackagesResponse) SetBody(v *ListBandwidthackagesResponseBody) *ListBandwidthackagesResponse {
	s.Body = v
	return s
}

type ListBasicAccelerateIpEndpointRelationsRequest struct {
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId  *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointId     *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBasicAccelerateIpEndpointRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpEndpointRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetAccelerateIpId(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetAcceleratorId(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetClientToken(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetEndpointId(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetMaxResults(v int32) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetNextToken(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsRequest) SetRegionId(v string) *ListBasicAccelerateIpEndpointRelationsRequest {
	s.RegionId = &v
	return s
}

type ListBasicAccelerateIpEndpointRelationsResponseBody struct {
	AccelerateIpEndpointRelations []*ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations `json:"AccelerateIpEndpointRelations,omitempty" xml:"AccelerateIpEndpointRelations,omitempty" type:"Repeated"`
	MaxResults                    *string                                                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                     *string                                                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                     *string                                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                    *string                                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBasicAccelerateIpEndpointRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpEndpointRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetAccelerateIpEndpointRelations(v []*ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.AccelerateIpEndpointRelations = v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetMaxResults(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetNextToken(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetRequestId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBody) SetTotalCount(v string) *ListBasicAccelerateIpEndpointRelationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations struct {
	AccelerateIpId         *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	EndpointAddress        *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName           *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	EndpointSubAddress     *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointZoneId         *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	IpAddress              *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	State                  *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetAccelerateIpId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.AccelerateIpId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointAddress(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointAddress = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointName(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointName = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointSubAddress(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointSubAddress = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointSubAddressType(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointSubAddressType = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointType(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointType = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetEndpointZoneId(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.EndpointZoneId = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetIpAddress(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.IpAddress = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations) SetState(v string) *ListBasicAccelerateIpEndpointRelationsResponseBodyAccelerateIpEndpointRelations {
	s.State = &v
	return s
}

type ListBasicAccelerateIpEndpointRelationsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBasicAccelerateIpEndpointRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBasicAccelerateIpEndpointRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpEndpointRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) SetHeaders(v map[string]*string) *ListBasicAccelerateIpEndpointRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) SetStatusCode(v int32) *ListBasicAccelerateIpEndpointRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicAccelerateIpEndpointRelationsResponse) SetBody(v *ListBasicAccelerateIpEndpointRelationsResponseBody) *ListBasicAccelerateIpEndpointRelationsResponse {
	s.Body = v
	return s
}

type ListBasicAccelerateIpsRequest struct {
	AccelerateIpAddress *string `json:"AccelerateIpAddress,omitempty" xml:"AccelerateIpAddress,omitempty"`
	AccelerateIpId      *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId             *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	MaxResults          *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBasicAccelerateIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpsRequest) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsRequest) SetAccelerateIpAddress(v string) *ListBasicAccelerateIpsRequest {
	s.AccelerateIpAddress = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetAccelerateIpId(v string) *ListBasicAccelerateIpsRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetClientToken(v string) *ListBasicAccelerateIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetIpSetId(v string) *ListBasicAccelerateIpsRequest {
	s.IpSetId = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetMaxResults(v int32) *ListBasicAccelerateIpsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetNextToken(v string) *ListBasicAccelerateIpsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBasicAccelerateIpsRequest) SetRegionId(v string) *ListBasicAccelerateIpsRequest {
	s.RegionId = &v
	return s
}

type ListBasicAccelerateIpsResponseBody struct {
	AccelerateIps []*ListBasicAccelerateIpsResponseBodyAccelerateIps `json:"AccelerateIps,omitempty" xml:"AccelerateIps,omitempty" type:"Repeated"`
	MaxResults    *int32                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBasicAccelerateIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsResponseBody) SetAccelerateIps(v []*ListBasicAccelerateIpsResponseBodyAccelerateIps) *ListBasicAccelerateIpsResponseBody {
	s.AccelerateIps = v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) SetMaxResults(v int32) *ListBasicAccelerateIpsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) SetNextToken(v string) *ListBasicAccelerateIpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) SetRequestId(v string) *ListBasicAccelerateIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBody) SetTotalCount(v int32) *ListBasicAccelerateIpsResponseBody {
	s.TotalCount = &v
	return s
}

type ListBasicAccelerateIpsResponseBodyAccelerateIps struct {
	AccelerateIpAddress *string                                                  `json:"AccelerateIpAddress,omitempty" xml:"AccelerateIpAddress,omitempty"`
	AccelerateIpId      *string                                                  `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	AcceleratorId       *string                                                  `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Endpoint            *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	IpSetId             *string                                                  `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	State               *string                                                  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListBasicAccelerateIpsResponseBodyAccelerateIps) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpsResponseBodyAccelerateIps) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetAccelerateIpAddress(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.AccelerateIpAddress = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetAccelerateIpId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.AccelerateIpId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetAcceleratorId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetEndpoint(v *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.Endpoint = v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetIpSetId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.IpSetId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIps) SetState(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIps {
	s.State = &v
	return s
}

type ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint struct {
	EndPointId             *string `json:"EndPointId,omitempty" xml:"EndPointId,omitempty"`
	EndpointAddress        *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointGroupId        *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointSubAddress     *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointZoneId         *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	State                  *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetEndPointId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.EndPointId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetEndpointAddress(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.EndpointAddress = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetEndpointGroupId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.EndpointGroupId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetEndpointSubAddress(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.EndpointSubAddress = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetEndpointSubAddressType(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.EndpointSubAddressType = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetEndpointType(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.EndpointType = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetEndpointZoneId(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.EndpointZoneId = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetName(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.Name = &v
	return s
}

func (s *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint) SetState(v string) *ListBasicAccelerateIpsResponseBodyAccelerateIpsEndpoint {
	s.State = &v
	return s
}

type ListBasicAccelerateIpsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBasicAccelerateIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBasicAccelerateIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAccelerateIpsResponse) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsResponse) SetHeaders(v map[string]*string) *ListBasicAccelerateIpsResponse {
	s.Headers = v
	return s
}

func (s *ListBasicAccelerateIpsResponse) SetStatusCode(v int32) *ListBasicAccelerateIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicAccelerateIpsResponse) SetBody(v *ListBasicAccelerateIpsResponseBody) *ListBasicAccelerateIpsResponse {
	s.Body = v
	return s
}

type ListBasicAcceleratorsRequest struct {
	AcceleratorId   *string                            `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	PageNumber      *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State           *string                            `json:"State,omitempty" xml:"State,omitempty"`
	Tag             []*ListBasicAcceleratorsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListBasicAcceleratorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAcceleratorsRequest) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsRequest) SetAcceleratorId(v string) *ListBasicAcceleratorsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetPageNumber(v int32) *ListBasicAcceleratorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetPageSize(v int32) *ListBasicAcceleratorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetRegionId(v string) *ListBasicAcceleratorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetResourceGroupId(v string) *ListBasicAcceleratorsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetState(v string) *ListBasicAcceleratorsRequest {
	s.State = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetTag(v []*ListBasicAcceleratorsRequestTag) *ListBasicAcceleratorsRequest {
	s.Tag = v
	return s
}

type ListBasicAcceleratorsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBasicAcceleratorsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAcceleratorsRequestTag) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsRequestTag) SetKey(v string) *ListBasicAcceleratorsRequestTag {
	s.Key = &v
	return s
}

func (s *ListBasicAcceleratorsRequestTag) SetValue(v string) *ListBasicAcceleratorsRequestTag {
	s.Value = &v
	return s
}

type ListBasicAcceleratorsResponseBody struct {
	Accelerators []*ListBasicAcceleratorsResponseBodyAccelerators `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	PageNumber   *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBasicAcceleratorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBody) SetAccelerators(v []*ListBasicAcceleratorsResponseBodyAccelerators) *ListBasicAcceleratorsResponseBody {
	s.Accelerators = v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) SetPageNumber(v int32) *ListBasicAcceleratorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) SetPageSize(v int32) *ListBasicAcceleratorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) SetRequestId(v string) *ListBasicAcceleratorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) SetTotalCount(v int32) *ListBasicAcceleratorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListBasicAcceleratorsResponseBodyAccelerators struct {
	AcceleratorId               *string                                                                   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	BandwidthBillingType        *string                                                                   `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	BasicBandwidthPackage       *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage       `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	BasicEndpointGroupId        *string                                                                   `json:"BasicEndpointGroupId,omitempty" xml:"BasicEndpointGroupId,omitempty"`
	BasicIpSetId                *string                                                                   `json:"BasicIpSetId,omitempty" xml:"BasicIpSetId,omitempty"`
	CreateTime                  *int64                                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossDomainBandwidthPackage *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	Description                 *string                                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime                 *int64                                                                    `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceChargeType          *string                                                                   `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	Name                        *string                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId                    *string                                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId             *string                                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	State                       *string                                                                   `json:"State,omitempty" xml:"State,omitempty"`
	Tags                        []*ListBasicAcceleratorsResponseBodyAcceleratorsTags                      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type                        *string                                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBasicAcceleratorsResponseBodyAccelerators) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBodyAccelerators) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetAcceleratorId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetBandwidthBillingType(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.BandwidthBillingType = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetBasicBandwidthPackage(v *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.BasicBandwidthPackage = v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetBasicEndpointGroupId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.BasicEndpointGroupId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetBasicIpSetId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.BasicIpSetId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetCreateTime(v int64) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.CreateTime = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetCrossDomainBandwidthPackage(v *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.CrossDomainBandwidthPackage = v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetDescription(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.Description = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetExpiredTime(v int64) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.ExpiredTime = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetInstanceChargeType(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.InstanceChargeType = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetName(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.Name = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetRegionId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.RegionId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetResourceGroupId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetState(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.State = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetTags(v []*ListBasicAcceleratorsResponseBodyAcceleratorsTags) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.Tags = v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetType(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.Type = &v
	return s
}

type ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage struct {
	Bandwidth     *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetBandwidth(v int32) *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetBandwidthType(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.BandwidthType = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetInstanceId(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.InstanceId = &v
	return s
}

type ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage struct {
	Bandwidth  *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) SetBandwidth(v int32) *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) SetInstanceId(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	s.InstanceId = &v
	return s
}

type ListBasicAcceleratorsResponseBodyAcceleratorsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsTags) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsTags) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsTags) SetKey(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsTags {
	s.Key = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsTags) SetValue(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsTags {
	s.Value = &v
	return s
}

type ListBasicAcceleratorsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBasicAcceleratorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBasicAcceleratorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBasicAcceleratorsResponse) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponse) SetHeaders(v map[string]*string) *ListBasicAcceleratorsResponse {
	s.Headers = v
	return s
}

func (s *ListBasicAcceleratorsResponse) SetStatusCode(v int32) *ListBasicAcceleratorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicAcceleratorsResponse) SetBody(v *ListBasicAcceleratorsResponseBody) *ListBasicAcceleratorsResponse {
	s.Body = v
	return s
}

type ListBasicEndpointsRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId      *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointType    *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBasicEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBasicEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListBasicEndpointsRequest) SetClientToken(v string) *ListBasicEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetEndpointGroupId(v string) *ListBasicEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetEndpointId(v string) *ListBasicEndpointsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetEndpointType(v string) *ListBasicEndpointsRequest {
	s.EndpointType = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetMaxResults(v int32) *ListBasicEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetName(v string) *ListBasicEndpointsRequest {
	s.Name = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetNextToken(v string) *ListBasicEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBasicEndpointsRequest) SetRegionId(v string) *ListBasicEndpointsRequest {
	s.RegionId = &v
	return s
}

type ListBasicEndpointsResponseBody struct {
	Endpoints  []*ListBasicEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	MaxResults *string                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBasicEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBasicEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicEndpointsResponseBody) SetEndpoints(v []*ListBasicEndpointsResponseBodyEndpoints) *ListBasicEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListBasicEndpointsResponseBody) SetMaxResults(v string) *ListBasicEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBasicEndpointsResponseBody) SetNextToken(v string) *ListBasicEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBasicEndpointsResponseBody) SetRequestId(v string) *ListBasicEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBasicEndpointsResponseBody) SetTotalCount(v string) *ListBasicEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListBasicEndpointsResponseBodyEndpoints struct {
	AcceleratorId          *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointAddress        *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointGroupId        *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointSubAddress     *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	EndpointZoneId         *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	State                  *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListBasicEndpointsResponseBodyEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListBasicEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetAcceleratorId(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointAddress(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointAddress = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointGroupId(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointGroupId = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointId(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointSubAddress(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointSubAddress = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointSubAddressType(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointSubAddressType = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointType(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetEndpointZoneId(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.EndpointZoneId = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetName(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.Name = &v
	return s
}

func (s *ListBasicEndpointsResponseBodyEndpoints) SetState(v string) *ListBasicEndpointsResponseBodyEndpoints {
	s.State = &v
	return s
}

type ListBasicEndpointsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBasicEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBasicEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBasicEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListBasicEndpointsResponse) SetHeaders(v map[string]*string) *ListBasicEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListBasicEndpointsResponse) SetStatusCode(v int32) *ListBasicEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicEndpointsResponse) SetBody(v *ListBasicEndpointsResponseBody) *ListBasicEndpointsResponse {
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
	Regions   []*ListBusiRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBusiRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBusiRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsResponseBody) SetRegions(v []*ListBusiRegionsResponseBodyRegions) *ListBusiRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListBusiRegionsResponseBody) SetRequestId(v string) *ListBusiRegionsResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBusiRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListBusiRegionsResponse) SetStatusCode(v int32) *ListBusiRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBusiRegionsResponse) SetBody(v *ListBusiRegionsResponseBody) *ListBusiRegionsResponse {
	s.Body = v
	return s
}

type ListCustomRoutingEndpointGroupDestinationsRequest struct {
	AcceleratorId   *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointGroupId *string   `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	FromPort        *int32    `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ListenerId      *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	PageNumber      *string   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Protocols       []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ToPort          *int32    `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetFromPort(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.FromPort = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetListenerId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetPageNumber(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetPageSize(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetProtocols(v []*string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.Protocols = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsRequest) SetToPort(v int32) *ListCustomRoutingEndpointGroupDestinationsRequest {
	s.ToPort = &v
	return s
}

type ListCustomRoutingEndpointGroupDestinationsResponseBody struct {
	Destinations []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations `json:"Destinations,omitempty" xml:"Destinations,omitempty" type:"Repeated"`
	PageNumber   *int32                                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetDestinations(v []*ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.Destinations = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetPageNumber(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetPageSize(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBody) SetTotalCount(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations struct {
	AcceleratorId   *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	DestinationId   *string   `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	EndpointGroupId *string   `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	FromPort        *int32    `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ListenerId      *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Protocols       []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	ToPort          *int32    `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetDestinationId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.DestinationId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetFromPort(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.FromPort = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetListenerId(v string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetProtocols(v []*string) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.Protocols = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations) SetToPort(v int32) *ListCustomRoutingEndpointGroupDestinationsResponseBodyDestinations {
	s.ToPort = &v
	return s
}

type ListCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *ListCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *ListCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *ListCustomRoutingEndpointGroupDestinationsResponseBody) *ListCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

type ListCustomRoutingEndpointGroupsRequest struct {
	AcceleratorId   *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingEndpointGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetListenerId(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetPageNumber(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetPageSize(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsRequest) SetRegionId(v string) *ListCustomRoutingEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

type ListCustomRoutingEndpointGroupsResponseBody struct {
	EndpointGroups []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups `json:"EndpointGroups,omitempty" xml:"EndpointGroups,omitempty" type:"Repeated"`
	PageNumber     *int32                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingEndpointGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetEndpointGroups(v []*ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) *ListCustomRoutingEndpointGroupsResponseBody {
	s.EndpointGroups = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetPageNumber(v int32) *ListCustomRoutingEndpointGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetPageSize(v int32) *ListCustomRoutingEndpointGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetRequestId(v string) *ListCustomRoutingEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBody) SetTotalCount(v int32) *ListCustomRoutingEndpointGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups struct {
	AcceleratorId                  *string   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Description                    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointGroupId                *string   `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupIpList            []*string `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	EndpointGroupRegion            *string   `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointGroupUnconfirmedIpList []*string `json:"EndpointGroupUnconfirmedIpList,omitempty" xml:"EndpointGroupUnconfirmedIpList,omitempty" type:"Repeated"`
	ListenerId                     *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Name                           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	State                          *string   `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetAcceleratorId(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetDescription(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.Description = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupId(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupIpList(v []*string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupIpList = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupRegion(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupUnconfirmedIpList(v []*string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupUnconfirmedIpList = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetListenerId(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetName(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.Name = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups) SetState(v string) *ListCustomRoutingEndpointGroupsResponseBodyEndpointGroups {
	s.State = &v
	return s
}

type ListCustomRoutingEndpointGroupsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomRoutingEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomRoutingEndpointGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupsResponse) SetHeaders(v map[string]*string) *ListCustomRoutingEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponse) SetStatusCode(v int32) *ListCustomRoutingEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupsResponse) SetBody(v *ListCustomRoutingEndpointGroupsResponseBody) *ListCustomRoutingEndpointGroupsResponse {
	s.Body = v
	return s
}

type ListCustomRoutingEndpointTrafficPoliciesRequest struct {
	AcceleratorId   *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId      *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetAcceleratorId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetAddress(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.Address = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointGroupId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetListenerId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetPageNumber(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetPageSize(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesRequest) SetRegionId(v string) *ListCustomRoutingEndpointTrafficPoliciesRequest {
	s.RegionId = &v
	return s
}

type ListCustomRoutingEndpointTrafficPoliciesResponseBody struct {
	PageNumber *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Policies   []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	RequestId  *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetPageNumber(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetPageSize(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetPolicies(v []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetRequestId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBody) SetTotalCount(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies struct {
	AcceleratorId   *string                                                                   `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Address         *string                                                                   `json:"Address,omitempty" xml:"Address,omitempty"`
	EndpointGroupId *string                                                                   `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId      *string                                                                   `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ListenerId      *string                                                                   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	PolicyId        *string                                                                   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PortRanges      []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetAcceleratorId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetAddress(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.Address = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetEndpointGroupId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetEndpointId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetListenerId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetPolicyId(v string) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies) SetPortRanges(v []*ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPolicies {
	s.PortRanges = v
	return s
}

type ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) SetFromPort(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges {
	s.FromPort = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges) SetToPort(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponseBodyPoliciesPortRanges {
	s.ToPort = &v
	return s
}

type ListCustomRoutingEndpointTrafficPoliciesResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomRoutingEndpointTrafficPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) SetHeaders(v map[string]*string) *ListCustomRoutingEndpointTrafficPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) SetStatusCode(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) SetBody(v *ListCustomRoutingEndpointTrafficPoliciesResponseBody) *ListCustomRoutingEndpointTrafficPoliciesResponse {
	s.Body = v
	return s
}

type ListCustomRoutingEndpointsRequest struct {
	AcceleratorId   *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsRequest) SetAcceleratorId(v string) *ListCustomRoutingEndpointsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetEndpointGroupId(v string) *ListCustomRoutingEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetListenerId(v string) *ListCustomRoutingEndpointsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetPageNumber(v string) *ListCustomRoutingEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetPageSize(v string) *ListCustomRoutingEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointsRequest) SetRegionId(v string) *ListCustomRoutingEndpointsRequest {
	s.RegionId = &v
	return s
}

type ListCustomRoutingEndpointsResponseBody struct {
	Endpoints  []*ListCustomRoutingEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsResponseBody) SetEndpoints(v []*ListCustomRoutingEndpointsResponseBodyEndpoints) *ListCustomRoutingEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) SetPageNumber(v int32) *ListCustomRoutingEndpointsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) SetPageSize(v int32) *ListCustomRoutingEndpointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) SetRequestId(v string) *ListCustomRoutingEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBody) SetTotalCount(v int32) *ListCustomRoutingEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomRoutingEndpointsResponseBodyEndpoints struct {
	AcceleratorId           *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Endpoint                *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EndpointGroupId         *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId              *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ListenerId              *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCustomRoutingEndpointsResponseBodyEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetAcceleratorId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetEndpoint(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.Endpoint = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetEndpointGroupId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetEndpointId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetListenerId(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetTrafficToEndpointPolicy(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponseBodyEndpoints) SetType(v string) *ListCustomRoutingEndpointsResponseBodyEndpoints {
	s.Type = &v
	return s
}

type ListCustomRoutingEndpointsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomRoutingEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomRoutingEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsResponse) SetHeaders(v map[string]*string) *ListCustomRoutingEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingEndpointsResponse) SetStatusCode(v int32) *ListCustomRoutingEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponse) SetBody(v *ListCustomRoutingEndpointsResponseBody) *ListCustomRoutingEndpointsResponse {
	s.Body = v
	return s
}

type ListCustomRoutingPortMappingsRequest struct {
	AcceleratorId   *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingPortMappingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsRequest) SetAcceleratorId(v string) *ListCustomRoutingPortMappingsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetEndpointGroupId(v string) *ListCustomRoutingPortMappingsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetListenerId(v string) *ListCustomRoutingPortMappingsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetPageNumber(v string) *ListCustomRoutingPortMappingsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetPageSize(v string) *ListCustomRoutingPortMappingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingPortMappingsRequest) SetRegionId(v string) *ListCustomRoutingPortMappingsRequest {
	s.RegionId = &v
	return s
}

type ListCustomRoutingPortMappingsResponseBody struct {
	PageNumber   *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PortMappings []*ListCustomRoutingPortMappingsResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	RequestId    *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingPortMappingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetPageNumber(v int32) *ListCustomRoutingPortMappingsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetPageSize(v int32) *ListCustomRoutingPortMappingsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetPortMappings(v []*ListCustomRoutingPortMappingsResponseBodyPortMappings) *ListCustomRoutingPortMappingsResponseBody {
	s.PortMappings = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetRequestId(v string) *ListCustomRoutingPortMappingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBody) SetTotalCount(v int32) *ListCustomRoutingPortMappingsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomRoutingPortMappingsResponseBodyPortMappings struct {
	AcceleratorId            *string                                                                        `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AcceleratorPort          *int32                                                                         `json:"AcceleratorPort,omitempty" xml:"AcceleratorPort,omitempty"`
	DestinationSocketAddress *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress `json:"DestinationSocketAddress,omitempty" xml:"DestinationSocketAddress,omitempty" type:"Struct"`
	DestinationTrafficState  *string                                                                        `json:"DestinationTrafficState,omitempty" xml:"DestinationTrafficState,omitempty"`
	EndpointGroupId          *string                                                                        `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupRegion      *string                                                                        `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointId               *string                                                                        `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ListenerId               *string                                                                        `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Protocols                []*string                                                                      `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Vswitch                  *string                                                                        `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
}

func (s ListCustomRoutingPortMappingsResponseBodyPortMappings) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsResponseBodyPortMappings) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetAcceleratorId(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetAcceleratorPort(v int32) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.AcceleratorPort = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetDestinationSocketAddress(v *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.DestinationSocketAddress = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetDestinationTrafficState(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.DestinationTrafficState = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetEndpointGroupId(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetEndpointGroupRegion(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetEndpointId(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetListenerId(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetProtocols(v []*string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.Protocols = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappings) SetVswitch(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappings {
	s.Vswitch = &v
	return s
}

type ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress struct {
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) SetIpAddress(v string) *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress {
	s.IpAddress = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress) SetPort(v int32) *ListCustomRoutingPortMappingsResponseBodyPortMappingsDestinationSocketAddress {
	s.Port = &v
	return s
}

type ListCustomRoutingPortMappingsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomRoutingPortMappingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomRoutingPortMappingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsResponse) SetHeaders(v map[string]*string) *ListCustomRoutingPortMappingsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingPortMappingsResponse) SetStatusCode(v int32) *ListCustomRoutingPortMappingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingPortMappingsResponse) SetBody(v *ListCustomRoutingPortMappingsResponseBody) *ListCustomRoutingPortMappingsResponse {
	s.Body = v
	return s
}

type ListCustomRoutingPortMappingsByDestinationRequest struct {
	DestinationAddress *string `json:"DestinationAddress,omitempty" xml:"DestinationAddress,omitempty"`
	EndpointId         *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	PageNumber         *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationRequest) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetDestinationAddress(v string) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.DestinationAddress = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetEndpointId(v string) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetPageNumber(v string) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetPageSize(v string) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationRequest) SetRegionId(v string) *ListCustomRoutingPortMappingsByDestinationRequest {
	s.RegionId = &v
	return s
}

type ListCustomRoutingPortMappingsByDestinationResponseBody struct {
	PageNumber   *int32                                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PortMappings []*ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	RequestId    *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetPageNumber(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetPageSize(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetPortMappings(v []*ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.PortMappings = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetRequestId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBody) SetTotalCount(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings struct {
	AcceleratorId            *string                                                                                     `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AcceleratorPort          *int32                                                                                      `json:"AcceleratorPort,omitempty" xml:"AcceleratorPort,omitempty"`
	DestinationSocketAddress *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress `json:"DestinationSocketAddress,omitempty" xml:"DestinationSocketAddress,omitempty" type:"Struct"`
	DestinationTrafficState  *string                                                                                     `json:"DestinationTrafficState,omitempty" xml:"DestinationTrafficState,omitempty"`
	EndpointGroupId          *string                                                                                     `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupRegion      *string                                                                                     `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointId               *string                                                                                     `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ListenerId               *string                                                                                     `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Protocols                []*string                                                                                   `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Vswitch                  *string                                                                                     `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetAcceleratorId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.AcceleratorId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetAcceleratorPort(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.AcceleratorPort = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetDestinationSocketAddress(v *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.DestinationSocketAddress = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetDestinationTrafficState(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.DestinationTrafficState = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetEndpointGroupId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.EndpointGroupId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetEndpointGroupRegion(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetEndpointId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.EndpointId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetListenerId(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.ListenerId = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetProtocols(v []*string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.Protocols = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings) SetVswitch(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappings {
	s.Vswitch = &v
	return s
}

type ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress struct {
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) SetIpAddress(v string) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress {
	s.IpAddress = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress) SetPort(v int32) *ListCustomRoutingPortMappingsByDestinationResponseBodyPortMappingsDestinationSocketAddress {
	s.Port = &v
	return s
}

type ListCustomRoutingPortMappingsByDestinationResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomRoutingPortMappingsByDestinationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomRoutingPortMappingsByDestinationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomRoutingPortMappingsByDestinationResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) SetHeaders(v map[string]*string) *ListCustomRoutingPortMappingsByDestinationResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) SetStatusCode(v int32) *ListCustomRoutingPortMappingsByDestinationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingPortMappingsByDestinationResponse) SetBody(v *ListCustomRoutingPortMappingsByDestinationResponseBody) *ListCustomRoutingPortMappingsByDestinationResponse {
	s.Body = v
	return s
}

type ListEndpointGroupsRequest struct {
	AcceleratorId     *string                         `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccessLogSwitch   *string                         `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	EndpointGroupId   *string                         `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupType *string                         `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	ListenerId        *string                         `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	PageNumber        *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId          *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag               []*ListEndpointGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListEndpointGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsRequest) SetAcceleratorId(v string) *ListEndpointGroupsRequest {
	s.AcceleratorId = &v
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

func (s *ListEndpointGroupsRequest) SetEndpointGroupType(v string) *ListEndpointGroupsRequest {
	s.EndpointGroupType = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetListenerId(v string) *ListEndpointGroupsRequest {
	s.ListenerId = &v
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

func (s *ListEndpointGroupsRequest) SetRegionId(v string) *ListEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetTag(v []*ListEndpointGroupsRequestTag) *ListEndpointGroupsRequest {
	s.Tag = v
	return s
}

type ListEndpointGroupsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEndpointGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsRequestTag) SetKey(v string) *ListEndpointGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListEndpointGroupsRequestTag) SetValue(v string) *ListEndpointGroupsRequestTag {
	s.Value = &v
	return s
}

type ListEndpointGroupsResponseBody struct {
	EndpointGroups []*ListEndpointGroupsResponseBodyEndpointGroups `json:"EndpointGroups,omitempty" xml:"EndpointGroups,omitempty" type:"Repeated"`
	PageNumber     *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEndpointGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBody) SetEndpointGroups(v []*ListEndpointGroupsResponseBodyEndpointGroups) *ListEndpointGroupsResponseBody {
	s.EndpointGroups = v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetPageNumber(v int32) *ListEndpointGroupsResponseBody {
	s.PageNumber = &v
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

func (s *ListEndpointGroupsResponseBody) SetTotalCount(v int32) *ListEndpointGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEndpointGroupsResponseBodyEndpointGroups struct {
	AcceleratorId                  *string                                                               `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	Description                    *string                                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointConfigurations         []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupId                *string                                                               `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupIpList            []*string                                                             `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	EndpointGroupRegion            *string                                                               `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointGroupType              *string                                                               `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointGroupUnconfirmedIpList []*string                                                             `json:"EndpointGroupUnconfirmedIpList,omitempty" xml:"EndpointGroupUnconfirmedIpList,omitempty" type:"Repeated"`
	EndpointRequestProtocol        *string                                                               `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	ForwardingRuleIds              []*string                                                             `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	HealthCheckEnabled             *bool                                                                 `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckIntervalSeconds     *int32                                                                `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath                *string                                                               `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort                *int32                                                                `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol            *string                                                               `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	ListenerId                     *string                                                               `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Name                           *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	PortOverrides                  []*ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	State                          *string                                                               `json:"State,omitempty" xml:"State,omitempty"`
	Tags                           []*ListEndpointGroupsResponseBodyEndpointGroupsTags                   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ThresholdCount                 *int32                                                                `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	TrafficPercentage              *int32                                                                `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroups) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroups) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetAcceleratorId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.AcceleratorId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetDescription(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.Description = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointConfigurations(v []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointConfigurations = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupIpList(v []*string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupIpList = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupRegion(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupType(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupType = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupUnconfirmedIpList(v []*string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupUnconfirmedIpList = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointRequestProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetForwardingRuleIds(v []*string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ForwardingRuleIds = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckEnabled(v bool) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckEnabled = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckIntervalSeconds(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckPath(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckPath = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckPort = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetListenerId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ListenerId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetName(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.Name = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetPortOverrides(v []*ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.PortOverrides = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetState(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.State = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetTags(v []*ListEndpointGroupsResponseBodyEndpointGroupsTags) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.Tags = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetThresholdCount(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ThresholdCount = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetTrafficPercentage(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.TrafficPercentage = &v
	return s
}

type ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations struct {
	EnableClientIPPreservation *bool   `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	Endpoint                   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EndpointId                 *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ProbePort                  *int32  `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
	ProbeProtocol              *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	Type                       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight                     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEnableClientIPPreservation(v bool) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEndpoint(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEndpointId(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.EndpointId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetProbePort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.ProbePort = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetProbeProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.ProbeProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetType(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetWeight(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Weight = &v
	return s
}

type ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides struct {
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) SetEndpointPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) SetListenerPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides {
	s.ListenerPort = &v
	return s
}

type ListEndpointGroupsResponseBodyEndpointGroupsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsTags) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsTags) SetKey(v string) *ListEndpointGroupsResponseBodyEndpointGroupsTags {
	s.Key = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsTags) SetValue(v string) *ListEndpointGroupsResponseBodyEndpointGroupsTags {
	s.Value = &v
	return s
}

type ListEndpointGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListEndpointGroupsResponse) SetStatusCode(v int32) *ListEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEndpointGroupsResponse) SetBody(v *ListEndpointGroupsResponseBody) *ListEndpointGroupsResponse {
	s.Body = v
	return s
}

type ListForwardingRulesRequest struct {
	AcceleratorId    *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	ListenerId       *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListForwardingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesRequest) SetAcceleratorId(v string) *ListForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetClientToken(v string) *ListForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListForwardingRulesRequest) SetForwardingRuleId(v string) *ListForwardingRulesRequest {
	s.ForwardingRuleId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetListenerId(v string) *ListForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetMaxResults(v int32) *ListForwardingRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListForwardingRulesRequest) SetNextToken(v string) *ListForwardingRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListForwardingRulesRequest) SetRegionId(v string) *ListForwardingRulesRequest {
	s.RegionId = &v
	return s
}

type ListForwardingRulesResponseBody struct {
	ForwardingRules []*ListForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	MaxResults      *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListForwardingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBody) SetForwardingRules(v []*ListForwardingRulesResponseBodyForwardingRules) *ListForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

func (s *ListForwardingRulesResponseBody) SetMaxResults(v int32) *ListForwardingRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetNextToken(v string) *ListForwardingRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetRequestId(v string) *ListForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListForwardingRulesResponseBody) SetTotalCount(v int32) *ListForwardingRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListForwardingRulesResponseBodyForwardingRules struct {
	ForwardingRuleDirection *string                                                         `json:"ForwardingRuleDirection,omitempty" xml:"ForwardingRuleDirection,omitempty"`
	ForwardingRuleId        *string                                                         `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	ForwardingRuleName      *string                                                         `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	ForwardingRuleStatus    *string                                                         `json:"ForwardingRuleStatus,omitempty" xml:"ForwardingRuleStatus,omitempty"`
	ListenerId              *string                                                         `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Priority                *int32                                                          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleActions             []*ListForwardingRulesResponseBodyForwardingRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	RuleConditions          []*ListForwardingRulesResponseBodyForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
}

func (s ListForwardingRulesResponseBodyForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetForwardingRuleDirection(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleDirection = &v
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

func (s *ListForwardingRulesResponseBodyForwardingRules) SetListenerId(v string) *ListForwardingRulesResponseBodyForwardingRules {
	s.ListenerId = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetPriority(v int32) *ListForwardingRulesResponseBodyForwardingRules {
	s.Priority = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetRuleActions(v []*ListForwardingRulesResponseBodyForwardingRulesRuleActions) *ListForwardingRulesResponseBodyForwardingRules {
	s.RuleActions = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRules) SetRuleConditions(v []*ListForwardingRulesResponseBodyForwardingRulesRuleConditions) *ListForwardingRulesResponseBodyForwardingRules {
	s.RuleConditions = v
	return s
}

type ListForwardingRulesResponseBodyForwardingRulesRuleActions struct {
	ForwardGroupConfig *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	Order              *int32                                                                       `json:"Order,omitempty" xml:"Order,omitempty"`
	RuleActionType     *string                                                                      `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	RuleActionValue    *string                                                                      `json:"RuleActionValue,omitempty" xml:"RuleActionValue,omitempty"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetForwardGroupConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleActionsForwardGroupConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetOrder(v int32) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetRuleActionType(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleActions) SetRuleActionValue(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleActions {
	s.RuleActionValue = &v
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

type ListForwardingRulesResponseBodyForwardingRulesRuleConditions struct {
	HostConfig         *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	PathConfig         *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	RuleConditionType  *string                                                                 `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	RuleConditionValue *string                                                                 `json:"RuleConditionValue,omitempty" xml:"RuleConditionValue,omitempty"`
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s ListForwardingRulesResponseBodyForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetHostConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsHostConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetPathConfig(v *ListForwardingRulesResponseBodyForwardingRulesRuleConditionsPathConfig) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetRuleConditionType(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *ListForwardingRulesResponseBodyForwardingRulesRuleConditions) SetRuleConditionValue(v string) *ListForwardingRulesResponseBodyForwardingRulesRuleConditions {
	s.RuleConditionValue = &v
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

type ListForwardingRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListForwardingRulesResponse) SetStatusCode(v int32) *ListForwardingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListForwardingRulesResponse) SetBody(v *ListForwardingRulesResponseBody) *ListForwardingRulesResponse {
	s.Body = v
	return s
}

type ListIpSetsRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIpSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIpSetsRequest) GoString() string {
	return s.String()
}

func (s *ListIpSetsRequest) SetAcceleratorId(v string) *ListIpSetsRequest {
	s.AcceleratorId = &v
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

func (s *ListIpSetsRequest) SetRegionId(v string) *ListIpSetsRequest {
	s.RegionId = &v
	return s
}

type ListIpSetsResponseBody struct {
	IpSets     []*ListIpSetsResponseBodyIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpSetsResponseBody) SetIpSets(v []*ListIpSetsResponseBodyIpSets) *ListIpSetsResponseBody {
	s.IpSets = v
	return s
}

func (s *ListIpSetsResponseBody) SetPageNumber(v int32) *ListIpSetsResponseBody {
	s.PageNumber = &v
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

func (s *ListIpSetsResponseBody) SetTotalCount(v int32) *ListIpSetsResponseBody {
	s.TotalCount = &v
	return s
}

type ListIpSetsResponseBodyIpSets struct {
	AccelerateRegionId *string   `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	Bandwidth          *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	IpAddressList      []*string `json:"IpAddressList,omitempty" xml:"IpAddressList,omitempty" type:"Repeated"`
	IpSetId            *string   `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	IpVersion          *string   `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	IspType            *string   `json:"IspType,omitempty" xml:"IspType,omitempty"`
	State              *string   `json:"State,omitempty" xml:"State,omitempty"`
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

func (s *ListIpSetsResponseBodyIpSets) SetBandwidth(v int32) *ListIpSetsResponseBodyIpSets {
	s.Bandwidth = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpAddressList(v []*string) *ListIpSetsResponseBodyIpSets {
	s.IpAddressList = v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpSetId(v string) *ListIpSetsResponseBodyIpSets {
	s.IpSetId = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIpVersion(v string) *ListIpSetsResponseBodyIpSets {
	s.IpVersion = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetIspType(v string) *ListIpSetsResponseBodyIpSets {
	s.IspType = &v
	return s
}

func (s *ListIpSetsResponseBodyIpSets) SetState(v string) *ListIpSetsResponseBodyIpSets {
	s.State = &v
	return s
}

type ListIpSetsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListIpSetsResponse) SetStatusCode(v int32) *ListIpSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpSetsResponse) SetBody(v *ListIpSetsResponseBody) *ListIpSetsResponse {
	s.Body = v
	return s
}

type ListListenerCertificatesRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ListenerId    *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Role          *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListListenerCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesRequest) SetAcceleratorId(v string) *ListListenerCertificatesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetListenerId(v string) *ListListenerCertificatesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetMaxResults(v int32) *ListListenerCertificatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetNextToken(v string) *ListListenerCertificatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetRegionId(v string) *ListListenerCertificatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetRole(v string) *ListListenerCertificatesRequest {
	s.Role = &v
	return s
}

type ListListenerCertificatesResponseBody struct {
	Certificates []*ListListenerCertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	MaxResults   *int32                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenerCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBody) SetCertificates(v []*ListListenerCertificatesResponseBodyCertificates) *ListListenerCertificatesResponseBody {
	s.Certificates = v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetMaxResults(v int32) *ListListenerCertificatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetNextToken(v string) *ListListenerCertificatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetRequestId(v string) *ListListenerCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetTotalCount(v int32) *ListListenerCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenerCertificatesResponseBodyCertificates struct {
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	IsDefault     *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListListenerCertificatesResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateId(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetDomain(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.Domain = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetIsDefault(v bool) *ListListenerCertificatesResponseBodyCertificates {
	s.IsDefault = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetState(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.State = &v
	return s
}

type ListListenerCertificatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListListenerCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenerCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponse) SetHeaders(v map[string]*string) *ListListenerCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListListenerCertificatesResponse) SetStatusCode(v int32) *ListListenerCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenerCertificatesResponse) SetBody(v *ListListenerCertificatesResponseBody) *ListListenerCertificatesResponse {
	s.Body = v
	return s
}

type ListListenersRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListListenersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) SetAcceleratorId(v string) *ListListenersRequest {
	s.AcceleratorId = &v
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

func (s *ListListenersRequest) SetRegionId(v string) *ListListenersRequest {
	s.RegionId = &v
	return s
}

type ListListenersResponseBody struct {
	Listeners  []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetPageNumber(v int32) *ListListenersResponseBody {
	s.PageNumber = &v
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

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenersResponseBodyListeners struct {
	AcceleratorId       *string                                                `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	BackendPorts        []*ListListenersResponseBodyListenersBackendPorts      `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	Certificates        []*ListListenersResponseBodyListenersCertificates      `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	ClientAffinity      *string                                                `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	CreateTime          *int64                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description         *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ListenerId          *string                                                `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Name                *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges          []*ListListenersResponseBodyListenersPortRanges        `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol            *string                                                `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ProxyProtocol       *bool                                                  `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	SecurityPolicyId    *string                                                `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	State               *string                                                `json:"State,omitempty" xml:"State,omitempty"`
	Type                *string                                                `json:"Type,omitempty" xml:"Type,omitempty"`
	XForwardedForConfig *ListListenersResponseBodyListenersXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) SetAcceleratorId(v string) *ListListenersResponseBodyListeners {
	s.AcceleratorId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetBackendPorts(v []*ListListenersResponseBodyListenersBackendPorts) *ListListenersResponseBodyListeners {
	s.BackendPorts = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCertificates(v []*ListListenersResponseBodyListenersCertificates) *ListListenersResponseBodyListeners {
	s.Certificates = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetClientAffinity(v string) *ListListenersResponseBodyListeners {
	s.ClientAffinity = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCreateTime(v int64) *ListListenersResponseBodyListeners {
	s.CreateTime = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetDescription(v string) *ListListenersResponseBodyListeners {
	s.Description = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetName(v string) *ListListenersResponseBodyListeners {
	s.Name = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetPortRanges(v []*ListListenersResponseBodyListenersPortRanges) *ListListenersResponseBodyListeners {
	s.PortRanges = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetProtocol(v string) *ListListenersResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetProxyProtocol(v bool) *ListListenersResponseBodyListeners {
	s.ProxyProtocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetSecurityPolicyId(v string) *ListListenersResponseBodyListeners {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetState(v string) *ListListenersResponseBodyListeners {
	s.State = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetType(v string) *ListListenersResponseBodyListeners {
	s.Type = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetXForwardedForConfig(v *ListListenersResponseBodyListenersXForwardedForConfig) *ListListenersResponseBodyListeners {
	s.XForwardedForConfig = v
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

type ListListenersResponseBodyListenersCertificates struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListListenersResponseBodyListenersCertificates) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersCertificates) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersCertificates) SetId(v string) *ListListenersResponseBodyListenersCertificates {
	s.Id = &v
	return s
}

func (s *ListListenersResponseBodyListenersCertificates) SetType(v string) *ListListenersResponseBodyListenersCertificates {
	s.Type = &v
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

type ListListenersResponseBodyListenersXForwardedForConfig struct {
	XForwardedForGaApEnabled  *bool `json:"XForwardedForGaApEnabled,omitempty" xml:"XForwardedForGaApEnabled,omitempty"`
	XForwardedForGaIdEnabled  *bool `json:"XForwardedForGaIdEnabled,omitempty" xml:"XForwardedForGaIdEnabled,omitempty"`
	XForwardedForPortEnabled  *bool `json:"XForwardedForPortEnabled,omitempty" xml:"XForwardedForPortEnabled,omitempty"`
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	XRealIpEnabled            *bool `json:"XRealIpEnabled,omitempty" xml:"XRealIpEnabled,omitempty"`
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForGaApEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForGaApEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForGaIdEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForGaIdEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForPortEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForPortEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXRealIpEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XRealIpEnabled = &v
	return s
}

type ListListenersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListListenersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListListenersResponse) SetStatusCode(v int32) *ListListenersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenersResponse) SetBody(v *ListListenersResponseBody) *ListListenersResponse {
	s.Body = v
	return s
}

type ListSpareIpsRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSpareIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpareIpsRequest) GoString() string {
	return s.String()
}

func (s *ListSpareIpsRequest) SetAcceleratorId(v string) *ListSpareIpsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListSpareIpsRequest) SetClientToken(v string) *ListSpareIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSpareIpsRequest) SetDryRun(v bool) *ListSpareIpsRequest {
	s.DryRun = &v
	return s
}

func (s *ListSpareIpsRequest) SetRegionId(v string) *ListSpareIpsRequest {
	s.RegionId = &v
	return s
}

type ListSpareIpsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpareIps  []*ListSpareIpsResponseBodySpareIps `json:"SpareIps,omitempty" xml:"SpareIps,omitempty" type:"Repeated"`
}

func (s ListSpareIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpareIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpareIpsResponseBody) SetRequestId(v string) *ListSpareIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpareIpsResponseBody) SetSpareIps(v []*ListSpareIpsResponseBodySpareIps) *ListSpareIpsResponseBody {
	s.SpareIps = v
	return s
}

type ListSpareIpsResponseBodySpareIps struct {
	SpareIp *string `json:"SpareIp,omitempty" xml:"SpareIp,omitempty"`
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListSpareIpsResponseBodySpareIps) String() string {
	return tea.Prettify(s)
}

func (s ListSpareIpsResponseBodySpareIps) GoString() string {
	return s.String()
}

func (s *ListSpareIpsResponseBodySpareIps) SetSpareIp(v string) *ListSpareIpsResponseBodySpareIps {
	s.SpareIp = &v
	return s
}

func (s *ListSpareIpsResponseBodySpareIps) SetState(v string) *ListSpareIpsResponseBodySpareIps {
	s.State = &v
	return s
}

type ListSpareIpsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSpareIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSpareIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpareIpsResponse) GoString() string {
	return s.String()
}

func (s *ListSpareIpsResponse) SetHeaders(v map[string]*string) *ListSpareIpsResponse {
	s.Headers = v
	return s
}

func (s *ListSpareIpsResponse) SetStatusCode(v int32) *ListSpareIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSpareIpsResponse) SetBody(v *ListSpareIpsResponseBody) *ListSpareIpsResponse {
	s.Body = v
	return s
}

type ListSystemSecurityPoliciesRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSystemSecurityPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesRequest) SetPageNumber(v int32) *ListSystemSecurityPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSystemSecurityPoliciesRequest) SetPageSize(v int32) *ListSystemSecurityPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSystemSecurityPoliciesRequest) SetRegionId(v string) *ListSystemSecurityPoliciesRequest {
	s.RegionId = &v
	return s
}

type ListSystemSecurityPoliciesResponseBody struct {
	PageNumber       *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityPolicies []*ListSystemSecurityPoliciesResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
	TotalCount       *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSystemSecurityPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBody) SetPageNumber(v int32) *ListSystemSecurityPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetPageSize(v int32) *ListSystemSecurityPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetRequestId(v string) *ListSystemSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetSecurityPolicies(v []*ListSystemSecurityPoliciesResponseBodySecurityPolicies) *ListSystemSecurityPoliciesResponseBody {
	s.SecurityPolicies = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetTotalCount(v int32) *ListSystemSecurityPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListSystemSecurityPoliciesResponseBodySecurityPolicies struct {
	Ciphers          []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	SecurityPolicyId *string   `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	TlsVersions      []*string `json:"TlsVersions,omitempty" xml:"TlsVersions,omitempty" type:"Repeated"`
}

func (s ListSystemSecurityPoliciesResponseBodySecurityPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetCiphers(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.Ciphers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetTlsVersions(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.TlsVersions = v
	return s
}

type ListSystemSecurityPoliciesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSystemSecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSystemSecurityPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponse) SetHeaders(v map[string]*string) *ListSystemSecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponse) SetStatusCode(v int32) *ListSystemSecurityPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponse) SetBody(v *ListSystemSecurityPoliciesResponseBody) *ListSystemSecurityPoliciesResponse {
	s.Body = v
	return s
}

type RemoveEntriesFromAclRequest struct {
	AclEntries  []*RemoveEntriesFromAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	AclId       *string                                  `json:"AclId,omitempty" xml:"AclId,omitempty"`
	ClientToken *string                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool                                    `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId    *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveEntriesFromAclRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclRequest) SetAclEntries(v []*RemoveEntriesFromAclRequestAclEntries) *RemoveEntriesFromAclRequest {
	s.AclEntries = v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetAclId(v string) *RemoveEntriesFromAclRequest {
	s.AclId = &v
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

func (s *RemoveEntriesFromAclRequest) SetRegionId(v string) *RemoveEntriesFromAclRequest {
	s.RegionId = &v
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
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveEntriesFromAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponseBody) SetAclId(v string) *RemoveEntriesFromAclResponseBody {
	s.AclId = &v
	return s
}

func (s *RemoveEntriesFromAclResponseBody) SetRequestId(v string) *RemoveEntriesFromAclResponseBody {
	s.RequestId = &v
	return s
}

type RemoveEntriesFromAclResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveEntriesFromAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveEntriesFromAclResponse) SetStatusCode(v int32) *RemoveEntriesFromAclResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveEntriesFromAclResponse) SetBody(v *RemoveEntriesFromAclResponseBody) *RemoveEntriesFromAclResponse {
	s.Body = v
	return s
}

type ReplaceBandwidthPackageRequest struct {
	AcceleratorId            *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	BandwidthPackageId       *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TargetBandwidthPackageId *string `json:"TargetBandwidthPackageId,omitempty" xml:"TargetBandwidthPackageId,omitempty"`
}

func (s ReplaceBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *ReplaceBandwidthPackageRequest) SetAcceleratorId(v string) *ReplaceBandwidthPackageRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ReplaceBandwidthPackageRequest) SetBandwidthPackageId(v string) *ReplaceBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ReplaceBandwidthPackageRequest) SetRegionId(v string) *ReplaceBandwidthPackageRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReplaceBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReplaceBandwidthPackageResponse) SetStatusCode(v int32) *ReplaceBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceBandwidthPackageResponse) SetBody(v *ReplaceBandwidthPackageResponseBody) *ReplaceBandwidthPackageResponse {
	s.Body = v
	return s
}

type UpdateAcceleratorRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AutoPay       *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Spec          *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorRequest) SetAcceleratorId(v string) *UpdateAcceleratorRequest {
	s.AcceleratorId = &v
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

func (s *UpdateAcceleratorRequest) SetClientToken(v string) *UpdateAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetDescription(v string) *UpdateAcceleratorRequest {
	s.Description = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetName(v string) *UpdateAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetRegionId(v string) *UpdateAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAcceleratorRequest) SetSpec(v string) *UpdateAcceleratorRequest {
	s.Spec = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAcceleratorResponse) SetStatusCode(v int32) *UpdateAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAcceleratorResponse) SetBody(v *UpdateAcceleratorResponseBody) *UpdateAcceleratorResponse {
	s.Body = v
	return s
}

type UpdateAcceleratorAutoRenewAttributeRequest struct {
	AcceleratorId     *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AutoRenew         *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RenewalStatus     *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s UpdateAcceleratorAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetAcceleratorId(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetAutoRenew(v bool) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetAutoRenewDuration(v int32) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetClientToken(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetName(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetRegionId(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetRenewalStatus(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

type UpdateAcceleratorAutoRenewAttributeResponseBody struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAcceleratorAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorAutoRenewAttributeResponseBody) SetAcceleratorId(v string) *UpdateAcceleratorAutoRenewAttributeResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeResponseBody) SetRequestId(v string) *UpdateAcceleratorAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAcceleratorAutoRenewAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAcceleratorAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAcceleratorAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *UpdateAcceleratorAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) SetStatusCode(v int32) *UpdateAcceleratorAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeResponse) SetBody(v *UpdateAcceleratorAutoRenewAttributeResponseBody) *UpdateAcceleratorAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type UpdateAcceleratorConfirmRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAcceleratorConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorConfirmRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorConfirmRequest) SetAcceleratorId(v string) *UpdateAcceleratorConfirmRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorConfirmRequest) SetRegionId(v string) *UpdateAcceleratorConfirmRequest {
	s.RegionId = &v
	return s
}

type UpdateAcceleratorConfirmResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAcceleratorConfirmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorConfirmResponseBody) SetRequestId(v string) *UpdateAcceleratorConfirmResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAcceleratorConfirmResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAcceleratorConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAcceleratorConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAcceleratorConfirmResponse) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorConfirmResponse) SetHeaders(v map[string]*string) *UpdateAcceleratorConfirmResponse {
	s.Headers = v
	return s
}

func (s *UpdateAcceleratorConfirmResponse) SetStatusCode(v int32) *UpdateAcceleratorConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAcceleratorConfirmResponse) SetBody(v *UpdateAcceleratorConfirmResponseBody) *UpdateAcceleratorConfirmResponse {
	s.Body = v
	return s
}

type UpdateAclAttributeRequest struct {
	AclId       *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName     *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAclAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeRequest) GoString() string {
	return s.String()
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

func (s *UpdateAclAttributeRequest) SetRegionId(v string) *UpdateAclAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateAclAttributeResponseBody struct {
	AclId     *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAclAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponseBody) SetAclId(v string) *UpdateAclAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *UpdateAclAttributeResponseBody) SetRequestId(v string) *UpdateAclAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAclAttributeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAclAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateAclAttributeResponse) SetStatusCode(v int32) *UpdateAclAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAclAttributeResponse) SetBody(v *UpdateAclAttributeResponseBody) *UpdateAclAttributeResponse {
	s.Body = v
	return s
}

type UpdateAdditionalCertificateWithListenerRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DryRun        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ListenerId    *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAdditionalCertificateWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdditionalCertificateWithListenerRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetAcceleratorId(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetCertificateId(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.CertificateId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetClientToken(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetDomain(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.Domain = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetDryRun(v bool) *UpdateAdditionalCertificateWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetListenerId(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerRequest) SetRegionId(v string) *UpdateAdditionalCertificateWithListenerRequest {
	s.RegionId = &v
	return s
}

type UpdateAdditionalCertificateWithListenerResponseBody struct {
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAdditionalCertificateWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdditionalCertificateWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdditionalCertificateWithListenerResponseBody) SetListenerId(v string) *UpdateAdditionalCertificateWithListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerResponseBody) SetRequestId(v string) *UpdateAdditionalCertificateWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAdditionalCertificateWithListenerResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAdditionalCertificateWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAdditionalCertificateWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdditionalCertificateWithListenerResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdditionalCertificateWithListenerResponse) SetHeaders(v map[string]*string) *UpdateAdditionalCertificateWithListenerResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerResponse) SetStatusCode(v int32) *UpdateAdditionalCertificateWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerResponse) SetBody(v *UpdateAdditionalCertificateWithListenerResponseBody) *UpdateAdditionalCertificateWithListenerResponse {
	s.Body = v
	return s
}

type UpdateApplicationMonitorRequest struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DetectEnable    *bool   `json:"DetectEnable,omitempty" xml:"DetectEnable,omitempty"`
	DetectThreshold *int32  `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	DetectTimes     *int32  `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	ListenerId      *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	OptionsJson     *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SilenceTime     *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateApplicationMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationMonitorRequest) SetAddress(v string) *UpdateApplicationMonitorRequest {
	s.Address = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetClientToken(v string) *UpdateApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetDetectEnable(v bool) *UpdateApplicationMonitorRequest {
	s.DetectEnable = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetDetectThreshold(v int32) *UpdateApplicationMonitorRequest {
	s.DetectThreshold = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetDetectTimes(v int32) *UpdateApplicationMonitorRequest {
	s.DetectTimes = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetListenerId(v string) *UpdateApplicationMonitorRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetOptionsJson(v string) *UpdateApplicationMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetRegionId(v string) *UpdateApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetSilenceTime(v int32) *UpdateApplicationMonitorRequest {
	s.SilenceTime = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetTaskId(v string) *UpdateApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetTaskName(v string) *UpdateApplicationMonitorRequest {
	s.TaskName = &v
	return s
}

type UpdateApplicationMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationMonitorResponseBody) SetRequestId(v string) *UpdateApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApplicationMonitorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateApplicationMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationMonitorResponse) SetHeaders(v map[string]*string) *UpdateApplicationMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationMonitorResponse) SetStatusCode(v int32) *UpdateApplicationMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationMonitorResponse) SetBody(v *UpdateApplicationMonitorResponseBody) *UpdateApplicationMonitorResponse {
	s.Body = v
	return s
}

type UpdateBandwidthPackagaAutoRenewAttributeRequest struct {
	AutoRenew         *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RenewalStatus     *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s UpdateBandwidthPackagaAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBandwidthPackagaAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetAutoRenew(v bool) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetAutoRenewDuration(v int32) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetClientToken(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetInstanceId(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetName(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetRegionId(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetRenewalStatus(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

type UpdateBandwidthPackagaAutoRenewAttributeResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBandwidthPackagaAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBandwidthPackagaAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) SetInstanceId(v string) *UpdateBandwidthPackagaAutoRenewAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) SetRequestId(v string) *UpdateBandwidthPackagaAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBandwidthPackagaAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBandwidthPackagaAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBandwidthPackagaAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBandwidthPackagaAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *UpdateBandwidthPackagaAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) SetStatusCode(v int32) *UpdateBandwidthPackagaAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponse) SetBody(v *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) *UpdateBandwidthPackagaAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type UpdateBandwidthPackageRequest struct {
	AutoPay            *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon      *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthType      *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBandwidthPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackageRequest) SetAutoPay(v bool) *UpdateBandwidthPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetAutoUseCoupon(v bool) *UpdateBandwidthPackageRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidth(v int32) *UpdateBandwidthPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidthPackageId(v string) *UpdateBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetBandwidthType(v string) *UpdateBandwidthPackageRequest {
	s.BandwidthType = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetDescription(v string) *UpdateBandwidthPackageRequest {
	s.Description = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetName(v string) *UpdateBandwidthPackageRequest {
	s.Name = &v
	return s
}

func (s *UpdateBandwidthPackageRequest) SetRegionId(v string) *UpdateBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

type UpdateBandwidthPackageResponseBody struct {
	BandwidthPackage *string `json:"BandwidthPackage,omitempty" xml:"BandwidthPackage,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *UpdateBandwidthPackageResponseBody) SetName(v string) *UpdateBandwidthPackageResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateBandwidthPackageResponseBody) SetRequestId(v string) *UpdateBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBandwidthPackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateBandwidthPackageResponse) SetStatusCode(v int32) *UpdateBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBandwidthPackageResponse) SetBody(v *UpdateBandwidthPackageResponseBody) *UpdateBandwidthPackageResponse {
	s.Body = v
	return s
}

type UpdateBasicAcceleratorRequest struct {
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBasicAcceleratorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicAcceleratorRequest) SetAcceleratorId(v string) *UpdateBasicAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) SetClientToken(v string) *UpdateBasicAcceleratorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) SetDescription(v string) *UpdateBasicAcceleratorRequest {
	s.Description = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) SetName(v string) *UpdateBasicAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *UpdateBasicAcceleratorRequest) SetRegionId(v string) *UpdateBasicAcceleratorRequest {
	s.RegionId = &v
	return s
}

type UpdateBasicAcceleratorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBasicAcceleratorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicAcceleratorResponseBody) SetRequestId(v string) *UpdateBasicAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBasicAcceleratorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBasicAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBasicAcceleratorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicAcceleratorResponse) SetHeaders(v map[string]*string) *UpdateBasicAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicAcceleratorResponse) SetStatusCode(v int32) *UpdateBasicAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicAcceleratorResponse) SetBody(v *UpdateBasicAcceleratorResponseBody) *UpdateBasicAcceleratorResponse {
	s.Body = v
	return s
}

type UpdateBasicEndpointRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId      *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBasicEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointRequest) SetClientToken(v string) *UpdateBasicEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBasicEndpointRequest) SetEndpointGroupId(v string) *UpdateBasicEndpointRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateBasicEndpointRequest) SetEndpointId(v string) *UpdateBasicEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateBasicEndpointRequest) SetName(v string) *UpdateBasicEndpointRequest {
	s.Name = &v
	return s
}

func (s *UpdateBasicEndpointRequest) SetRegionId(v string) *UpdateBasicEndpointRequest {
	s.RegionId = &v
	return s
}

type UpdateBasicEndpointResponseBody struct {
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointId      *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBasicEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointResponseBody) SetEndpointGroupId(v string) *UpdateBasicEndpointResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateBasicEndpointResponseBody) SetEndpointId(v string) *UpdateBasicEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *UpdateBasicEndpointResponseBody) SetName(v string) *UpdateBasicEndpointResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateBasicEndpointResponseBody) SetRequestId(v string) *UpdateBasicEndpointResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBasicEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBasicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBasicEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointResponse) SetHeaders(v map[string]*string) *UpdateBasicEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicEndpointResponse) SetStatusCode(v int32) *UpdateBasicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicEndpointResponse) SetBody(v *UpdateBasicEndpointResponseBody) *UpdateBasicEndpointResponse {
	s.Body = v
	return s
}

type UpdateBasicEndpointGroupRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointAddress    *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	EndpointGroupId    *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointSubAddress *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	EndpointType       *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBasicEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointGroupRequest) SetClientToken(v string) *UpdateBasicEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetDescription(v string) *UpdateBasicEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetEndpointAddress(v string) *UpdateBasicEndpointGroupRequest {
	s.EndpointAddress = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetEndpointGroupId(v string) *UpdateBasicEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetEndpointSubAddress(v string) *UpdateBasicEndpointGroupRequest {
	s.EndpointSubAddress = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetEndpointType(v string) *UpdateBasicEndpointGroupRequest {
	s.EndpointType = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetName(v string) *UpdateBasicEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetRegionId(v string) *UpdateBasicEndpointGroupRequest {
	s.RegionId = &v
	return s
}

type UpdateBasicEndpointGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBasicEndpointGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointGroupResponseBody) SetRequestId(v string) *UpdateBasicEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBasicEndpointGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBasicEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBasicEndpointGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointGroupResponse) SetHeaders(v map[string]*string) *UpdateBasicEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicEndpointGroupResponse) SetStatusCode(v int32) *UpdateBasicEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicEndpointGroupResponse) SetBody(v *UpdateBasicEndpointGroupResponseBody) *UpdateBasicEndpointGroupResponse {
	s.Body = v
	return s
}

type UpdateBasicIpSetRequest struct {
	Bandwidth   *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId     *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBasicIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicIpSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicIpSetRequest) SetBandwidth(v int32) *UpdateBasicIpSetRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateBasicIpSetRequest) SetClientToken(v string) *UpdateBasicIpSetRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBasicIpSetRequest) SetIpSetId(v string) *UpdateBasicIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *UpdateBasicIpSetRequest) SetRegionId(v string) *UpdateBasicIpSetRequest {
	s.RegionId = &v
	return s
}

type UpdateBasicIpSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBasicIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicIpSetResponseBody) SetRequestId(v string) *UpdateBasicIpSetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBasicIpSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBasicIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBasicIpSetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBasicIpSetResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicIpSetResponse) SetHeaders(v map[string]*string) *UpdateBasicIpSetResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicIpSetResponse) SetStatusCode(v int32) *UpdateBasicIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicIpSetResponse) SetBody(v *UpdateBasicIpSetResponseBody) *UpdateBasicIpSetResponse {
	s.Body = v
	return s
}

type UpdateCustomRoutingEndpointGroupAttributeRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetClientToken(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetDescription(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetName(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetRegionId(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateCustomRoutingEndpointGroupAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponseBody) SetRequestId(v string) *UpdateCustomRoutingEndpointGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCustomRoutingEndpointGroupAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCustomRoutingEndpointGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomRoutingEndpointGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) SetStatusCode(v int32) *UpdateCustomRoutingEndpointGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) SetBody(v *UpdateCustomRoutingEndpointGroupAttributeResponseBody) *UpdateCustomRoutingEndpointGroupAttributeResponse {
	s.Body = v
	return s
}

type UpdateCustomRoutingEndpointGroupDestinationsRequest struct {
	ClientToken               *string                                                                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationConfigurations []*UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	DryRun                    *bool                                                                           `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointGroupId           *string                                                                         `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RegionId                  *string                                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetClientToken(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetDestinationConfigurations(v []*UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.DestinationConfigurations = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetDryRun(v bool) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

type UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations struct {
	DestinationId *string   `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	FromPort      *int32    `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	Protocols     []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	ToPort        *int32    `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetDestinationId(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.DestinationId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetFromPort(v int32) *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.FromPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetProtocols(v []*string) *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.Protocols = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetToPort(v int32) *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.ToPort = &v
	return s
}

type UpdateCustomRoutingEndpointGroupDestinationsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *UpdateCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *UpdateCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *UpdateCustomRoutingEndpointGroupDestinationsResponseBody) *UpdateCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

type UpdateCustomRoutingEndpointTrafficPoliciesRequest struct {
	ClientToken          *string                                                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointId           *string                                                                  `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	PolicyConfigurations []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	RegionId             *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) SetClientToken(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) SetPolicyConfigurations(v []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) *UpdateCustomRoutingEndpointTrafficPoliciesRequest {
	s.PolicyConfigurations = v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequest) SetRegionId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequest {
	s.RegionId = &v
	return s
}

type UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations struct {
	Address    *string                                                                            `json:"Address,omitempty" xml:"Address,omitempty"`
	PolicyId   *string                                                                            `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PortRanges []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetAddress(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetPolicyId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.PolicyId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetPortRanges(v []*UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.PortRanges = v
	return s
}

type UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges struct {
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) SetFromPort(v int32) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) SetToPort(v int32) *UpdateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

type UpdateCustomRoutingEndpointTrafficPoliciesResponseBody struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) SetPolicyIds(v []*string) *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.PolicyIds = v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) SetRequestId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCustomRoutingEndpointTrafficPoliciesResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointTrafficPoliciesResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) SetStatusCode(v int32) *UpdateCustomRoutingEndpointTrafficPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) SetBody(v *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) *UpdateCustomRoutingEndpointTrafficPoliciesResponse {
	s.Body = v
	return s
}

type UpdateCustomRoutingEndpointsRequest struct {
	ClientToken            *string                                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointConfigurations []*UpdateCustomRoutingEndpointsRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupId        *string                                                      `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	RegionId               *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateCustomRoutingEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsRequest) SetClientToken(v string) *UpdateCustomRoutingEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequest) SetEndpointConfigurations(v []*UpdateCustomRoutingEndpointsRequestEndpointConfigurations) *UpdateCustomRoutingEndpointsRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequest) SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequest) SetRegionId(v string) *UpdateCustomRoutingEndpointsRequest {
	s.RegionId = &v
	return s
}

type UpdateCustomRoutingEndpointsRequestEndpointConfigurations struct {
	EndpointId              *string                                                                          `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	PolicyConfigurations    []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	TrafficToEndpointPolicy *string                                                                          `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) SetEndpointId(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.EndpointId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) SetPolicyConfigurations(v []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) *UpdateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.PolicyConfigurations = v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) SetTrafficToEndpointPolicy(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.TrafficToEndpointPolicy = &v
	return s
}

type UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations struct {
	Address    *string                                                                                    `json:"Address,omitempty" xml:"Address,omitempty"`
	PortRanges []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) SetAddress(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) SetPortRanges(v []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	s.PortRanges = v
	return s
}

type UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges struct {
	FromPort *string `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	ToPort   *string `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) SetFromPort(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) SetToPort(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

type UpdateCustomRoutingEndpointsResponseBody struct {
	EndpointIds []*string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomRoutingEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsResponseBody) SetEndpointIds(v []*string) *UpdateCustomRoutingEndpointsResponseBody {
	s.EndpointIds = v
	return s
}

func (s *UpdateCustomRoutingEndpointsResponseBody) SetRequestId(v string) *UpdateCustomRoutingEndpointsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCustomRoutingEndpointsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCustomRoutingEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomRoutingEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsResponse) SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointsResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoutingEndpointsResponse) SetStatusCode(v int32) *UpdateCustomRoutingEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsResponse) SetBody(v *UpdateCustomRoutingEndpointsResponseBody) *UpdateCustomRoutingEndpointsResponse {
	s.Body = v
	return s
}

type UpdateEndpointGroupRequest struct {
	ClientToken                *string                                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointConfigurations     []*UpdateEndpointGroupRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupId            *string                                             `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupRegion        *string                                             `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointRequestProtocol    *string                                             `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	HealthCheckEnabled         *bool                                               `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckIntervalSeconds *int32                                              `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath            *string                                             `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort            *int32                                              `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol        *string                                             `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	Name                       *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	PortOverrides              []*UpdateEndpointGroupRequestPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	RegionId                   *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ThresholdCount             *int32                                              `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	TrafficPercentage          *int32                                              `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s UpdateEndpointGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequest) SetClientToken(v string) *UpdateEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetDescription(v string) *UpdateEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointConfigurations(v []*UpdateEndpointGroupRequestEndpointConfigurations) *UpdateEndpointGroupRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointGroupId(v string) *UpdateEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointGroupRegion(v string) *UpdateEndpointGroupRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointRequestProtocol(v string) *UpdateEndpointGroupRequest {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckEnabled(v bool) *UpdateEndpointGroupRequest {
	s.HealthCheckEnabled = &v
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

func (s *UpdateEndpointGroupRequest) SetName(v string) *UpdateEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetPortOverrides(v []*UpdateEndpointGroupRequestPortOverrides) *UpdateEndpointGroupRequest {
	s.PortOverrides = v
	return s
}

func (s *UpdateEndpointGroupRequest) SetRegionId(v string) *UpdateEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetThresholdCount(v int32) *UpdateEndpointGroupRequest {
	s.ThresholdCount = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetTrafficPercentage(v int32) *UpdateEndpointGroupRequest {
	s.TrafficPercentage = &v
	return s
}

type UpdateEndpointGroupRequestEndpointConfigurations struct {
	EnableClientIPPreservation *bool   `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	Endpoint                   *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Type                       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight                     *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateEndpointGroupRequestEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetEnableClientIPPreservation(v bool) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetEndpoint(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetType(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetWeight(v int32) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Weight = &v
	return s
}

type UpdateEndpointGroupRequestPortOverrides struct {
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s UpdateEndpointGroupRequestPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupRequestPortOverrides) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequestPortOverrides) SetEndpointPort(v int32) *UpdateEndpointGroupRequestPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *UpdateEndpointGroupRequestPortOverrides) SetListenerPort(v int32) *UpdateEndpointGroupRequestPortOverrides {
	s.ListenerPort = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateEndpointGroupResponse) SetStatusCode(v int32) *UpdateEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEndpointGroupResponse) SetBody(v *UpdateEndpointGroupResponseBody) *UpdateEndpointGroupResponse {
	s.Body = v
	return s
}

type UpdateEndpointGroupAttributeRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEndpointGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupAttributeRequest) SetClientToken(v string) *UpdateEndpointGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetDescription(v string) *UpdateEndpointGroupAttributeRequest {
	s.Description = &v
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

func (s *UpdateEndpointGroupAttributeRequest) SetRegionId(v string) *UpdateEndpointGroupAttributeRequest {
	s.RegionId = &v
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEndpointGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateEndpointGroupAttributeResponse) SetStatusCode(v int32) *UpdateEndpointGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEndpointGroupAttributeResponse) SetBody(v *UpdateEndpointGroupAttributeResponseBody) *UpdateEndpointGroupAttributeResponse {
	s.Body = v
	return s
}

type UpdateEndpointGroupsRequest struct {
	ClientToken                 *string                                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                      *bool                                                     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointGroupConfigurations []*UpdateEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	ListenerId                  *string                                                   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId                    *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEndpointGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsRequest) SetClientToken(v string) *UpdateEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEndpointGroupsRequest) SetDryRun(v bool) *UpdateEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEndpointGroupsRequest) SetEndpointGroupConfigurations(v []*UpdateEndpointGroupsRequestEndpointGroupConfigurations) *UpdateEndpointGroupsRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *UpdateEndpointGroupsRequest) SetListenerId(v string) *UpdateEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateEndpointGroupsRequest) SetRegionId(v string) *UpdateEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

type UpdateEndpointGroupsRequestEndpointGroupConfigurations struct {
	EnableClientIPPreservationProxyProtocol *bool                                                                           `json:"EnableClientIPPreservationProxyProtocol,omitempty" xml:"EnableClientIPPreservationProxyProtocol,omitempty"`
	EnableClientIPPreservationToa           *bool                                                                           `json:"EnableClientIPPreservationToa,omitempty" xml:"EnableClientIPPreservationToa,omitempty"`
	EndpointConfigurations                  []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupDescription                *string                                                                         `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	EndpointGroupId                         *string                                                                         `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupName                       *string                                                                         `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	EndpointRequestProtocol                 *string                                                                         `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	HealthCheckEnabled                      *bool                                                                           `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckIntervalSeconds              *int64                                                                          `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath                         *string                                                                         `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort                         *int64                                                                          `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol                     *string                                                                         `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	PortOverrides                           []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides          `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	ThresholdCount                          *int64                                                                          `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	TrafficPercentage                       *int64                                                                          `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurations) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEnableClientIPPreservationProxyProtocol(v bool) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EnableClientIPPreservationProxyProtocol = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEnableClientIPPreservationToa(v bool) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EnableClientIPPreservationToa = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupDescription(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupDescription = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupId(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupName(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupName = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointRequestProtocol(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckEnabled(v bool) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckEnabled = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckIntervalSeconds(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPath(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPort(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPort = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckProtocol(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetPortOverrides(v []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.PortOverrides = v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetThresholdCount(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.ThresholdCount = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetTrafficPercentage(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.TrafficPercentage = &v
	return s
}

type UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *int64  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetWeight(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Weight = &v
	return s
}

type UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides struct {
	EndpointPort *int64 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetEndpointPort(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetListenerPort(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.ListenerPort = &v
	return s
}

type UpdateEndpointGroupsResponseBody struct {
	EndpointGroupIds []*string `json:"EndpointGroupIds,omitempty" xml:"EndpointGroupIds,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEndpointGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsResponseBody) SetEndpointGroupIds(v []*string) *UpdateEndpointGroupsResponseBody {
	s.EndpointGroupIds = v
	return s
}

func (s *UpdateEndpointGroupsResponseBody) SetRequestId(v string) *UpdateEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEndpointGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEndpointGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsResponse) SetHeaders(v map[string]*string) *UpdateEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *UpdateEndpointGroupsResponse) SetStatusCode(v int32) *UpdateEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEndpointGroupsResponse) SetBody(v *UpdateEndpointGroupsResponseBody) *UpdateEndpointGroupsResponse {
	s.Body = v
	return s
}

type UpdateForwardingRulesRequest struct {
	AcceleratorId   *string                                        `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	ClientToken     *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForwardingRules []*UpdateForwardingRulesRequestForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	ListenerId      *string                                        `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId        *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateForwardingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequest) SetAcceleratorId(v string) *UpdateForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetClientToken(v string) *UpdateForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetForwardingRules(v []*UpdateForwardingRulesRequestForwardingRules) *UpdateForwardingRulesRequest {
	s.ForwardingRules = v
	return s
}

func (s *UpdateForwardingRulesRequest) SetListenerId(v string) *UpdateForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateForwardingRulesRequest) SetRegionId(v string) *UpdateForwardingRulesRequest {
	s.RegionId = &v
	return s
}

type UpdateForwardingRulesRequestForwardingRules struct {
	ForwardingRuleId   *string                                                      `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	ForwardingRuleName *string                                                      `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	Priority           *int32                                                       `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RuleActions        []*UpdateForwardingRulesRequestForwardingRulesRuleActions    `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	RuleConditions     []*UpdateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	RuleDirection      *string                                                      `json:"RuleDirection,omitempty" xml:"RuleDirection,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRules) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetForwardingRuleId(v string) *UpdateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetForwardingRuleName(v string) *UpdateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetPriority(v int32) *UpdateForwardingRulesRequestForwardingRules {
	s.Priority = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetRuleActions(v []*UpdateForwardingRulesRequestForwardingRulesRuleActions) *UpdateForwardingRulesRequestForwardingRules {
	s.RuleActions = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetRuleConditions(v []*UpdateForwardingRulesRequestForwardingRulesRuleConditions) *UpdateForwardingRulesRequestForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRules) SetRuleDirection(v string) *UpdateForwardingRulesRequestForwardingRules {
	s.RuleDirection = &v
	return s
}

type UpdateForwardingRulesRequestForwardingRulesRuleActions struct {
	ForwardGroupConfig *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	Order              *int32                                                                    `json:"Order,omitempty" xml:"Order,omitempty"`
	RuleActionType     *string                                                                   `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	RuleActionValue    *string                                                                   `json:"RuleActionValue,omitempty" xml:"RuleActionValue,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetForwardGroupConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetOrder(v int32) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionType(v string) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionValue(v string) *UpdateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionValue = &v
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

type UpdateForwardingRulesRequestForwardingRulesRuleConditions struct {
	HostConfig         *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	PathConfig         *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	RuleConditionType  *string                                                              `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	RuleConditionValue *string                                                              `json:"RuleConditionValue,omitempty" xml:"RuleConditionValue,omitempty"`
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateForwardingRulesRequestForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetHostConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetPathConfig(v *UpdateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionType(v string) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *UpdateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionValue(v string) *UpdateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionValue = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateForwardingRulesResponse) SetStatusCode(v int32) *UpdateForwardingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateForwardingRulesResponse) SetBody(v *UpdateForwardingRulesResponseBody) *UpdateForwardingRulesResponse {
	s.Body = v
	return s
}

type UpdateIpSetRequest struct {
	Bandwidth   *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IpSetId     *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpSetRequest) SetBandwidth(v int32) *UpdateIpSetRequest {
	s.Bandwidth = &v
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

func (s *UpdateIpSetRequest) SetRegionId(v string) *UpdateIpSetRequest {
	s.RegionId = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateIpSetResponse) SetStatusCode(v int32) *UpdateIpSetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpSetResponse) SetBody(v *UpdateIpSetResponseBody) *UpdateIpSetResponse {
	s.Body = v
	return s
}

type UpdateIpSetsRequest struct {
	IpSets   []*UpdateIpSetsRequestIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
	RegionId *string                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateIpSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIpSetsRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpSetsRequest) SetIpSets(v []*UpdateIpSetsRequestIpSets) *UpdateIpSetsRequest {
	s.IpSets = v
	return s
}

func (s *UpdateIpSetsRequest) SetRegionId(v string) *UpdateIpSetsRequest {
	s.RegionId = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIpSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateIpSetsResponse) SetStatusCode(v int32) *UpdateIpSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpSetsResponse) SetBody(v *UpdateIpSetsResponseBody) *UpdateIpSetsResponse {
	s.Body = v
	return s
}

type UpdateListenerRequest struct {
	BackendPorts        []*UpdateListenerRequestBackendPorts      `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	Certificates        []*UpdateListenerRequestCertificates      `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	ClientAffinity      *string                                   `json:"ClientAffinity,omitempty" xml:"ClientAffinity,omitempty"`
	ClientToken         *string                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description         *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	ListenerId          *string                                   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	Name                *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges          []*UpdateListenerRequestPortRanges        `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol            *string                                   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	ProxyProtocol       *string                                   `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	RegionId            *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityPolicyId    *string                                   `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	XForwardedForConfig *UpdateListenerRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s UpdateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequest) SetBackendPorts(v []*UpdateListenerRequestBackendPorts) *UpdateListenerRequest {
	s.BackendPorts = v
	return s
}

func (s *UpdateListenerRequest) SetCertificates(v []*UpdateListenerRequestCertificates) *UpdateListenerRequest {
	s.Certificates = v
	return s
}

func (s *UpdateListenerRequest) SetClientAffinity(v string) *UpdateListenerRequest {
	s.ClientAffinity = &v
	return s
}

func (s *UpdateListenerRequest) SetClientToken(v string) *UpdateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerRequest) SetDescription(v string) *UpdateListenerRequest {
	s.Description = &v
	return s
}

func (s *UpdateListenerRequest) SetListenerId(v string) *UpdateListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerRequest) SetName(v string) *UpdateListenerRequest {
	s.Name = &v
	return s
}

func (s *UpdateListenerRequest) SetPortRanges(v []*UpdateListenerRequestPortRanges) *UpdateListenerRequest {
	s.PortRanges = v
	return s
}

func (s *UpdateListenerRequest) SetProtocol(v string) *UpdateListenerRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateListenerRequest) SetProxyProtocol(v string) *UpdateListenerRequest {
	s.ProxyProtocol = &v
	return s
}

func (s *UpdateListenerRequest) SetRegionId(v string) *UpdateListenerRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateListenerRequest) SetSecurityPolicyId(v string) *UpdateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateListenerRequest) SetXForwardedForConfig(v *UpdateListenerRequestXForwardedForConfig) *UpdateListenerRequest {
	s.XForwardedForConfig = v
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

type UpdateListenerRequestXForwardedForConfig struct {
	XForwardedForGaApEnabled  *bool `json:"XForwardedForGaApEnabled,omitempty" xml:"XForwardedForGaApEnabled,omitempty"`
	XForwardedForGaIdEnabled  *bool `json:"XForwardedForGaIdEnabled,omitempty" xml:"XForwardedForGaIdEnabled,omitempty"`
	XForwardedForPortEnabled  *bool `json:"XForwardedForPortEnabled,omitempty" xml:"XForwardedForPortEnabled,omitempty"`
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	XRealIpEnabled            *bool `json:"XRealIpEnabled,omitempty" xml:"XRealIpEnabled,omitempty"`
}

func (s UpdateListenerRequestXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXForwardedForGaApEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XForwardedForGaApEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXForwardedForGaIdEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XForwardedForGaIdEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXForwardedForPortEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XForwardedForPortEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *UpdateListenerRequestXForwardedForConfig) SetXRealIpEnabled(v bool) *UpdateListenerRequestXForwardedForConfig {
	s.XRealIpEnabled = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateListenerResponse) SetStatusCode(v int32) *UpdateListenerResponse {
	s.StatusCode = &v
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclEntries)) {
		query["AclEntries"] = request.AclEntries
	}

	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddEntriesToAcl"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclIds)) {
		query["AclIds"] = request.AclIds
	}

	if !tea.BoolValue(util.IsUnset(request.AclType)) {
		query["AclType"] = request.AclType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateAclsWithListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) AssociateAdditionalCertificatesWithListenerWithOptions(request *AssociateAdditionalCertificatesWithListenerRequest, runtime *util.RuntimeOptions) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateAdditionalCertificatesWithListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateAdditionalCertificatesWithListener(request *AssociateAdditionalCertificatesWithListenerRequest) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.AssociateAdditionalCertificatesWithListenerWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosId)) {
		query["DdosId"] = request.DdosId
	}

	if !tea.BoolValue(util.IsUnset(request.DdosRegionId)) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachDdosToAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachDdosToAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupIds)) {
		query["EndpointGroupIds"] = request.EndpointGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SlsLogStoreName)) {
		query["SlsLogStoreName"] = request.SlsLogStoreName
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProjectName)) {
		query["SlsProjectName"] = request.SlsProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SlsRegionId)) {
		query["SlsRegionId"] = request.SlsRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachLogStoreToEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachLogStoreToEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BandwidthPackageAddAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BandwidthPackageAddAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BandwidthPackageRemoveAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BandwidthPackageRemoveAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.ProbePort)) {
		query["ProbePort"] = request.ProbePort
	}

	if !tea.BoolValue(util.IsUnset(request.ProbeProtocol)) {
		query["ProbeProtocol"] = request.ProbeProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigEndpointProbe"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigEndpointProbeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUseCoupon)) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthBillingType)) {
		query["BandwidthBillingType"] = request.BandwidthBillingType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetConfig)) {
		query["IpSetConfig"] = request.IpSetConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		query["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclEntries)) {
		query["AclEntries"] = request.AclEntries
	}

	if !tea.BoolValue(util.IsUnset(request.AclName)) {
		query["AclName"] = request.AclName
	}

	if !tea.BoolValue(util.IsUnset(request.AddressIPVersion)) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAcl"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateApplicationMonitorWithOptions(request *CreateApplicationMonitorRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DetectEnable)) {
		query["DetectEnable"] = request.DetectEnable
	}

	if !tea.BoolValue(util.IsUnset(request.DetectThreshold)) {
		query["DetectThreshold"] = request.DetectThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.DetectTimes)) {
		query["DetectTimes"] = request.DetectTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.OptionsJson)) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplicationMonitor"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplicationMonitor(request *CreateApplicationMonitorRequest) (_result *CreateApplicationMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationMonitorResponse{}
	_body, _err := client.CreateApplicationMonitorWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUseCoupon)) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthType)) {
		query["BandwidthType"] = request.BandwidthType
	}

	if !tea.BoolValue(util.IsUnset(request.BillingType)) {
		query["BillingType"] = request.BillingType
	}

	if !tea.BoolValue(util.IsUnset(request.CbnGeographicRegionIdA)) {
		query["CbnGeographicRegionIdA"] = request.CbnGeographicRegionIdA
	}

	if !tea.BoolValue(util.IsUnset(request.CbnGeographicRegionIdB)) {
		query["CbnGeographicRegionIdB"] = request.CbnGeographicRegionIdB
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Ratio)) {
		query["Ratio"] = request.Ratio
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBandwidthPackage"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateBasicAccelerateIpWithOptions(request *CreateBasicAccelerateIpRequest, runtime *util.RuntimeOptions) (_result *CreateBasicAccelerateIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBasicAccelerateIp"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBasicAccelerateIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBasicAccelerateIp(request *CreateBasicAccelerateIpRequest) (_result *CreateBasicAccelerateIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBasicAccelerateIpResponse{}
	_body, _err := client.CreateBasicAccelerateIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBasicAccelerateIpEndpointRelationWithOptions(request *CreateBasicAccelerateIpEndpointRelationRequest, runtime *util.RuntimeOptions) (_result *CreateBasicAccelerateIpEndpointRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateIpId)) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBasicAccelerateIpEndpointRelation"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBasicAccelerateIpEndpointRelation(request *CreateBasicAccelerateIpEndpointRelationRequest) (_result *CreateBasicAccelerateIpEndpointRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.CreateBasicAccelerateIpEndpointRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBasicAccelerateIpEndpointRelationsWithOptions(request *CreateBasicAccelerateIpEndpointRelationsRequest, runtime *util.RuntimeOptions) (_result *CreateBasicAccelerateIpEndpointRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateIpEndpointRelations)) {
		query["AccelerateIpEndpointRelations"] = request.AccelerateIpEndpointRelations
	}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBasicAccelerateIpEndpointRelations"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBasicAccelerateIpEndpointRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBasicAccelerateIpEndpointRelations(request *CreateBasicAccelerateIpEndpointRelationsRequest) (_result *CreateBasicAccelerateIpEndpointRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBasicAccelerateIpEndpointRelationsResponse{}
	_body, _err := client.CreateBasicAccelerateIpEndpointRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBasicAcceleratorWithOptions(request *CreateBasicAcceleratorRequest, runtime *util.RuntimeOptions) (_result *CreateBasicAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUseCoupon)) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthBillingType)) {
		query["BandwidthBillingType"] = request.BandwidthBillingType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBasicAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBasicAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBasicAccelerator(request *CreateBasicAcceleratorRequest) (_result *CreateBasicAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBasicAcceleratorResponse{}
	_body, _err := client.CreateBasicAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBasicEndpointWithOptions(request *CreateBasicEndpointRequest, runtime *util.RuntimeOptions) (_result *CreateBasicEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointAddress)) {
		query["EndpointAddress"] = request.EndpointAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointSubAddress)) {
		query["EndpointSubAddress"] = request.EndpointSubAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointSubAddressType)) {
		query["EndpointSubAddressType"] = request.EndpointSubAddressType
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointZoneId)) {
		query["EndpointZoneId"] = request.EndpointZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBasicEndpoint"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBasicEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBasicEndpoint(request *CreateBasicEndpointRequest) (_result *CreateBasicEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBasicEndpointResponse{}
	_body, _err := client.CreateBasicEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBasicEndpointGroupWithOptions(request *CreateBasicEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *CreateBasicEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointAddress)) {
		query["EndpointAddress"] = request.EndpointAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupRegion)) {
		query["EndpointGroupRegion"] = request.EndpointGroupRegion
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointSubAddress)) {
		query["EndpointSubAddress"] = request.EndpointSubAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBasicEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBasicEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBasicEndpointGroup(request *CreateBasicEndpointGroupRequest) (_result *CreateBasicEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBasicEndpointGroupResponse{}
	_body, _err := client.CreateBasicEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBasicEndpointsWithOptions(request *CreateBasicEndpointsRequest, runtime *util.RuntimeOptions) (_result *CreateBasicEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoints)) {
		query["Endpoints"] = request.Endpoints
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBasicEndpoints"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBasicEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBasicEndpoints(request *CreateBasicEndpointsRequest) (_result *CreateBasicEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBasicEndpointsResponse{}
	_body, _err := client.CreateBasicEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBasicIpSetWithOptions(request *CreateBasicIpSetRequest, runtime *util.RuntimeOptions) (_result *CreateBasicIpSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateRegionId)) {
		query["AccelerateRegionId"] = request.AccelerateRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IspType)) {
		query["IspType"] = request.IspType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBasicIpSet"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBasicIpSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBasicIpSet(request *CreateBasicIpSetRequest) (_result *CreateBasicIpSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBasicIpSetResponse{}
	_body, _err := client.CreateBasicIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomRoutingEndpointGroupDestinationsWithOptions(request *CreateCustomRoutingEndpointGroupDestinationsRequest, runtime *util.RuntimeOptions) (_result *CreateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationConfigurations)) {
		query["DestinationConfigurations"] = request.DestinationConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomRoutingEndpointGroupDestinations"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomRoutingEndpointGroupDestinations(request *CreateCustomRoutingEndpointGroupDestinationsRequest) (_result *CreateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CreateCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomRoutingEndpointGroupsWithOptions(request *CreateCustomRoutingEndpointGroupsRequest, runtime *util.RuntimeOptions) (_result *CreateCustomRoutingEndpointGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupConfigurations)) {
		query["EndpointGroupConfigurations"] = request.EndpointGroupConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomRoutingEndpointGroups"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomRoutingEndpointGroups(request *CreateCustomRoutingEndpointGroupsRequest) (_result *CreateCustomRoutingEndpointGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.CreateCustomRoutingEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomRoutingEndpointTrafficPoliciesWithOptions(request *CreateCustomRoutingEndpointTrafficPoliciesRequest, runtime *util.RuntimeOptions) (_result *CreateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyConfigurations)) {
		query["PolicyConfigurations"] = request.PolicyConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomRoutingEndpointTrafficPolicies"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomRoutingEndpointTrafficPolicies(request *CreateCustomRoutingEndpointTrafficPoliciesRequest) (_result *CreateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CreateCustomRoutingEndpointTrafficPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomRoutingEndpointsWithOptions(request *CreateCustomRoutingEndpointsRequest, runtime *util.RuntimeOptions) (_result *CreateCustomRoutingEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointConfigurations)) {
		query["EndpointConfigurations"] = request.EndpointConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomRoutingEndpoints"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomRoutingEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomRoutingEndpoints(request *CreateCustomRoutingEndpointsRequest) (_result *CreateCustomRoutingEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomRoutingEndpointsResponse{}
	_body, _err := client.CreateCustomRoutingEndpointsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointConfigurations)) {
		query["EndpointConfigurations"] = request.EndpointConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupRegion)) {
		query["EndpointGroupRegion"] = request.EndpointGroupRegion
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupType)) {
		query["EndpointGroupType"] = request.EndpointGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointRequestProtocol)) {
		query["EndpointRequestProtocol"] = request.EndpointRequestProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckEnabled)) {
		query["HealthCheckEnabled"] = request.HealthCheckEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckIntervalSeconds)) {
		query["HealthCheckIntervalSeconds"] = request.HealthCheckIntervalSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckPath)) {
		query["HealthCheckPath"] = request.HealthCheckPath
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckPort)) {
		query["HealthCheckPort"] = request.HealthCheckPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckProtocol)) {
		query["HealthCheckProtocol"] = request.HealthCheckProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PortOverrides)) {
		query["PortOverrides"] = request.PortOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdCount)) {
		query["ThresholdCount"] = request.ThresholdCount
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficPercentage)) {
		query["TrafficPercentage"] = request.TrafficPercentage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateEndpointGroupsWithOptions(request *CreateEndpointGroupsRequest, runtime *util.RuntimeOptions) (_result *CreateEndpointGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupConfigurations)) {
		query["EndpointGroupConfigurations"] = request.EndpointGroupConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEndpointGroups"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEndpointGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEndpointGroups(request *CreateEndpointGroupsRequest) (_result *CreateEndpointGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEndpointGroupsResponse{}
	_body, _err := client.CreateEndpointGroupsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardingRules)) {
		query["ForwardingRules"] = request.ForwardingRules
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateForwardingRules"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateForwardingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateRegion)) {
		query["AccelerateRegion"] = request.AccelerateRegion
	}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpSets"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIpSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientAffinity)) {
		query["ClientAffinity"] = request.ClientAffinity
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CustomRoutingEndpointGroupConfigurations)) {
		query["CustomRoutingEndpointGroupConfigurations"] = request.CustomRoutingEndpointGroupConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupConfigurations)) {
		query["EndpointGroupConfigurations"] = request.EndpointGroupConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PortRanges)) {
		query["PortRanges"] = request.PortRanges
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocol)) {
		query["ProxyProtocol"] = request.ProxyProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedForConfig)) {
		query["XForwardedForConfig"] = request.XForwardedForConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateSpareIpsWithOptions(request *CreateSpareIpsRequest, runtime *util.RuntimeOptions) (_result *CreateSpareIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpareIps)) {
		query["SpareIps"] = request.SpareIps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSpareIps"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSpareIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSpareIps(request *CreateSpareIpsRequest) (_result *CreateSpareIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSpareIpsResponse{}
	_body, _err := client.CreateSpareIpsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAcl"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteApplicationMonitorWithOptions(request *DeleteApplicationMonitorRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplicationMonitor"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplicationMonitor(request *DeleteApplicationMonitorRequest) (_result *DeleteApplicationMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationMonitorResponse{}
	_body, _err := client.DeleteApplicationMonitorWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBandwidthPackage"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteBasicAccelerateIpWithOptions(request *DeleteBasicAccelerateIpRequest, runtime *util.RuntimeOptions) (_result *DeleteBasicAccelerateIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateIpId)) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBasicAccelerateIp"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBasicAccelerateIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBasicAccelerateIp(request *DeleteBasicAccelerateIpRequest) (_result *DeleteBasicAccelerateIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBasicAccelerateIpResponse{}
	_body, _err := client.DeleteBasicAccelerateIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBasicAccelerateIpEndpointRelationWithOptions(request *DeleteBasicAccelerateIpEndpointRelationRequest, runtime *util.RuntimeOptions) (_result *DeleteBasicAccelerateIpEndpointRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateIpId)) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBasicAccelerateIpEndpointRelation"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBasicAccelerateIpEndpointRelation(request *DeleteBasicAccelerateIpEndpointRelationRequest) (_result *DeleteBasicAccelerateIpEndpointRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.DeleteBasicAccelerateIpEndpointRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBasicAcceleratorWithOptions(request *DeleteBasicAcceleratorRequest, runtime *util.RuntimeOptions) (_result *DeleteBasicAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBasicAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBasicAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBasicAccelerator(request *DeleteBasicAcceleratorRequest) (_result *DeleteBasicAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBasicAcceleratorResponse{}
	_body, _err := client.DeleteBasicAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBasicEndpointWithOptions(request *DeleteBasicEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteBasicEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBasicEndpoint"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBasicEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBasicEndpoint(request *DeleteBasicEndpointRequest) (_result *DeleteBasicEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBasicEndpointResponse{}
	_body, _err := client.DeleteBasicEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBasicEndpointGroupWithOptions(request *DeleteBasicEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteBasicEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBasicEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBasicEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBasicEndpointGroup(request *DeleteBasicEndpointGroupRequest) (_result *DeleteBasicEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBasicEndpointGroupResponse{}
	_body, _err := client.DeleteBasicEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBasicIpSetWithOptions(request *DeleteBasicIpSetRequest, runtime *util.RuntimeOptions) (_result *DeleteBasicIpSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBasicIpSet"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBasicIpSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBasicIpSet(request *DeleteBasicIpSetRequest) (_result *DeleteBasicIpSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBasicIpSetResponse{}
	_body, _err := client.DeleteBasicIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomRoutingEndpointGroupDestinationsWithOptions(request *DeleteCustomRoutingEndpointGroupDestinationsRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationIds)) {
		query["DestinationIds"] = request.DestinationIds
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomRoutingEndpointGroupDestinations"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomRoutingEndpointGroupDestinations(request *DeleteCustomRoutingEndpointGroupDestinationsRequest) (_result *DeleteCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.DeleteCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomRoutingEndpointGroupsWithOptions(request *DeleteCustomRoutingEndpointGroupsRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomRoutingEndpointGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupIds)) {
		query["EndpointGroupIds"] = request.EndpointGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomRoutingEndpointGroups"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomRoutingEndpointGroups(request *DeleteCustomRoutingEndpointGroupsRequest) (_result *DeleteCustomRoutingEndpointGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.DeleteCustomRoutingEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomRoutingEndpointTrafficPoliciesWithOptions(request *DeleteCustomRoutingEndpointTrafficPoliciesRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyIds)) {
		query["PolicyIds"] = request.PolicyIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomRoutingEndpointTrafficPolicies"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomRoutingEndpointTrafficPolicies(request *DeleteCustomRoutingEndpointTrafficPoliciesRequest) (_result *DeleteCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.DeleteCustomRoutingEndpointTrafficPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomRoutingEndpointsWithOptions(request *DeleteCustomRoutingEndpointsRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomRoutingEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointIds)) {
		query["EndpointIds"] = request.EndpointIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomRoutingEndpoints"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomRoutingEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomRoutingEndpoints(request *DeleteCustomRoutingEndpointsRequest) (_result *DeleteCustomRoutingEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomRoutingEndpointsResponse{}
	_body, _err := client.DeleteCustomRoutingEndpointsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteEndpointGroupsWithOptions(request *DeleteEndpointGroupsRequest, runtime *util.RuntimeOptions) (_result *DeleteEndpointGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupIds)) {
		query["EndpointGroupIds"] = request.EndpointGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEndpointGroups"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEndpointGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEndpointGroups(request *DeleteEndpointGroupsRequest) (_result *DeleteEndpointGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEndpointGroupsResponse{}
	_body, _err := client.DeleteEndpointGroupsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardingRuleIds)) {
		query["ForwardingRuleIds"] = request.ForwardingRuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteForwardingRules"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteForwardingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpSet"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpSetIds)) {
		query["IpSetIds"] = request.IpSetIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpSets"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteSpareIpsWithOptions(request *DeleteSpareIpsRequest, runtime *util.RuntimeOptions) (_result *DeleteSpareIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpareIps)) {
		query["SpareIps"] = request.SpareIps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSpareIps"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSpareIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSpareIps(request *DeleteSpareIpsRequest) (_result *DeleteSpareIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSpareIpsResponse{}
	_body, _err := client.DeleteSpareIpsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeAcceleratorAutoRenewAttributeWithOptions(request *DescribeAcceleratorAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeAcceleratorAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAcceleratorAutoRenewAttribute"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAcceleratorAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAcceleratorAutoRenewAttribute(request *DescribeAcceleratorAutoRenewAttributeRequest) (_result *DescribeAcceleratorAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAcceleratorAutoRenewAttributeResponse{}
	_body, _err := client.DescribeAcceleratorAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicationMonitorWithOptions(request *DescribeApplicationMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeApplicationMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationMonitor"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicationMonitor(request *DescribeApplicationMonitorRequest) (_result *DescribeApplicationMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApplicationMonitorResponse{}
	_body, _err := client.DescribeApplicationMonitorWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBandwidthPackage"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeBandwidthPackageAutoRenewAttributeWithOptions(request *DescribeBandwidthPackageAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeBandwidthPackageAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBandwidthPackageAutoRenewAttribute"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBandwidthPackageAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBandwidthPackageAutoRenewAttribute(request *DescribeBandwidthPackageAutoRenewAttributeRequest) (_result *DescribeBandwidthPackageAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBandwidthPackageAutoRenewAttributeResponse{}
	_body, _err := client.DescribeBandwidthPackageAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomRoutingEndPointTrafficPolicyWithOptions(request *DescribeCustomRoutingEndPointTrafficPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomRoutingEndPointTrafficPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomRoutingEndPointTrafficPolicy"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomRoutingEndPointTrafficPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomRoutingEndPointTrafficPolicy(request *DescribeCustomRoutingEndPointTrafficPolicyRequest) (_result *DescribeCustomRoutingEndPointTrafficPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomRoutingEndPointTrafficPolicyResponse{}
	_body, _err := client.DescribeCustomRoutingEndPointTrafficPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomRoutingEndpointWithOptions(request *DescribeCustomRoutingEndpointRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomRoutingEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomRoutingEndpoint"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomRoutingEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomRoutingEndpoint(request *DescribeCustomRoutingEndpointRequest) (_result *DescribeCustomRoutingEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomRoutingEndpointResponse{}
	_body, _err := client.DescribeCustomRoutingEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomRoutingEndpointGroupWithOptions(request *DescribeCustomRoutingEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomRoutingEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomRoutingEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomRoutingEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomRoutingEndpointGroup(request *DescribeCustomRoutingEndpointGroupRequest) (_result *DescribeCustomRoutingEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomRoutingEndpointGroupResponse{}
	_body, _err := client.DescribeCustomRoutingEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomRoutingEndpointGroupDestinationsWithOptions(request *DescribeCustomRoutingEndpointGroupDestinationsRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationId)) {
		query["DestinationId"] = request.DestinationId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomRoutingEndpointGroupDestinations"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomRoutingEndpointGroupDestinations(request *DescribeCustomRoutingEndpointGroupDestinationsRequest) (_result *DescribeCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.DescribeCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpSet"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIpSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachDdosFromAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachDdosFromAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupIds)) {
		query["EndpointGroupIds"] = request.EndpointGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachLogStoreFromEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachLogStoreFromEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DetectApplicationMonitorWithOptions(request *DetectApplicationMonitorRequest, runtime *util.RuntimeOptions) (_result *DetectApplicationMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectApplicationMonitor"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectApplicationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectApplicationMonitor(request *DetectApplicationMonitorRequest) (_result *DetectApplicationMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectApplicationMonitorResponse{}
	_body, _err := client.DetectApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableApplicationMonitorWithOptions(request *DisableApplicationMonitorRequest, runtime *util.RuntimeOptions) (_result *DisableApplicationMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableApplicationMonitor"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableApplicationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableApplicationMonitor(request *DisableApplicationMonitorRequest) (_result *DisableApplicationMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableApplicationMonitorResponse{}
	_body, _err := client.DisableApplicationMonitorWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclIds)) {
		query["AclIds"] = request.AclIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateAclsFromListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DissociateAdditionalCertificatesFromListenerWithOptions(request *DissociateAdditionalCertificatesFromListenerRequest, runtime *util.RuntimeOptions) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Domains)) {
		query["Domains"] = request.Domains
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateAdditionalCertificatesFromListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateAdditionalCertificatesFromListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateAdditionalCertificatesFromListener(request *DissociateAdditionalCertificatesFromListenerRequest) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateAdditionalCertificatesFromListenerResponse{}
	_body, _err := client.DissociateAdditionalCertificatesFromListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableApplicationMonitorWithOptions(request *EnableApplicationMonitorRequest, runtime *util.RuntimeOptions) (_result *EnableApplicationMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableApplicationMonitor"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableApplicationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableApplicationMonitor(request *EnableApplicationMonitorRequest) (_result *EnableApplicationMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableApplicationMonitorResponse{}
	_body, _err := client.EnableApplicationMonitorWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAcl"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetBasicAccelerateIpWithOptions(request *GetBasicAccelerateIpRequest, runtime *util.RuntimeOptions) (_result *GetBasicAccelerateIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateIpId)) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBasicAccelerateIp"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicAccelerateIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBasicAccelerateIp(request *GetBasicAccelerateIpRequest) (_result *GetBasicAccelerateIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBasicAccelerateIpResponse{}
	_body, _err := client.GetBasicAccelerateIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBasicAccelerateIpEndpointRelationWithOptions(request *GetBasicAccelerateIpEndpointRelationRequest, runtime *util.RuntimeOptions) (_result *GetBasicAccelerateIpEndpointRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateIpId)) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBasicAccelerateIpEndpointRelation"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBasicAccelerateIpEndpointRelation(request *GetBasicAccelerateIpEndpointRelationRequest) (_result *GetBasicAccelerateIpEndpointRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBasicAccelerateIpEndpointRelationResponse{}
	_body, _err := client.GetBasicAccelerateIpEndpointRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBasicAccelerateIpIdleCountWithOptions(request *GetBasicAccelerateIpIdleCountRequest, runtime *util.RuntimeOptions) (_result *GetBasicAccelerateIpIdleCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBasicAccelerateIpIdleCount"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicAccelerateIpIdleCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBasicAccelerateIpIdleCount(request *GetBasicAccelerateIpIdleCountRequest) (_result *GetBasicAccelerateIpIdleCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBasicAccelerateIpIdleCountResponse{}
	_body, _err := client.GetBasicAccelerateIpIdleCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBasicAcceleratorWithOptions(request *GetBasicAcceleratorRequest, runtime *util.RuntimeOptions) (_result *GetBasicAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBasicAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBasicAccelerator(request *GetBasicAcceleratorRequest) (_result *GetBasicAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBasicAcceleratorResponse{}
	_body, _err := client.GetBasicAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBasicEndpointWithOptions(request *GetBasicEndpointRequest, runtime *util.RuntimeOptions) (_result *GetBasicEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBasicEndpoint"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBasicEndpoint(request *GetBasicEndpointRequest) (_result *GetBasicEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBasicEndpointResponse{}
	_body, _err := client.GetBasicEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBasicEndpointGroupWithOptions(request *GetBasicEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *GetBasicEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBasicEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBasicEndpointGroup(request *GetBasicEndpointGroupRequest) (_result *GetBasicEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBasicEndpointGroupResponse{}
	_body, _err := client.GetBasicEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBasicIpSetWithOptions(request *GetBasicIpSetRequest, runtime *util.RuntimeOptions) (_result *GetBasicIpSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBasicIpSet"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBasicIpSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBasicIpSet(request *GetBasicIpSetRequest) (_result *GetBasicIpSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBasicIpSetResponse{}
	_body, _err := client.GetBasicIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHealthStatusWithOptions(request *GetHealthStatusRequest, runtime *util.RuntimeOptions) (_result *GetHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHealthStatus"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHealthStatus(request *GetHealthStatusRequest) (_result *GetHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHealthStatusResponse{}
	_body, _err := client.GetHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpareIpWithOptions(request *GetSpareIpRequest, runtime *util.RuntimeOptions) (_result *GetSpareIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpareIp)) {
		query["SpareIp"] = request.SpareIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSpareIp"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpareIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpareIp(request *GetSpareIpRequest) (_result *GetSpareIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpareIpResponse{}
	_body, _err := client.GetSpareIpWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccelerateAreas"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccelerateAreasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccelerators"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAcceleratorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclIds)) {
		query["AclIds"] = request.AclIds
	}

	if !tea.BoolValue(util.IsUnset(request.AclName)) {
		query["AclName"] = request.AclName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAcls"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAclsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListApplicationMonitorWithOptions(request *ListApplicationMonitorRequest, runtime *util.RuntimeOptions) (_result *ListApplicationMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationMonitor"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplicationMonitor(request *ListApplicationMonitorRequest) (_result *ListApplicationMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationMonitorResponse{}
	_body, _err := client.ListApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationMonitorDetectResultWithOptions(request *ListApplicationMonitorDetectResultRequest, runtime *util.RuntimeOptions) (_result *ListApplicationMonitorDetectResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationMonitorDetectResult"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationMonitorDetectResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplicationMonitorDetectResult(request *ListApplicationMonitorDetectResultRequest) (_result *ListApplicationMonitorDetectResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationMonitorDetectResultResponse{}
	_body, _err := client.ListApplicationMonitorDetectResultWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableAccelerateAreas"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableAccelerateAreasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableBusiRegions"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableBusiRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListBandwidthPackagesWithOptions(request *ListBandwidthPackagesRequest, runtime *util.RuntimeOptions) (_result *ListBandwidthPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBandwidthPackages"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBandwidthPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListBandwidthackagesWithOptions(request *ListBandwidthackagesRequest, runtime *util.RuntimeOptions) (_result *ListBandwidthackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBandwidthackages"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBandwidthackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListBasicAccelerateIpEndpointRelationsWithOptions(request *ListBasicAccelerateIpEndpointRelationsRequest, runtime *util.RuntimeOptions) (_result *ListBasicAccelerateIpEndpointRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateIpId)) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBasicAccelerateIpEndpointRelations"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBasicAccelerateIpEndpointRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBasicAccelerateIpEndpointRelations(request *ListBasicAccelerateIpEndpointRelationsRequest) (_result *ListBasicAccelerateIpEndpointRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBasicAccelerateIpEndpointRelationsResponse{}
	_body, _err := client.ListBasicAccelerateIpEndpointRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBasicAccelerateIpsWithOptions(request *ListBasicAccelerateIpsRequest, runtime *util.RuntimeOptions) (_result *ListBasicAccelerateIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccelerateIpAddress)) {
		query["AccelerateIpAddress"] = request.AccelerateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.AccelerateIpId)) {
		query["AccelerateIpId"] = request.AccelerateIpId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBasicAccelerateIps"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBasicAccelerateIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBasicAccelerateIps(request *ListBasicAccelerateIpsRequest) (_result *ListBasicAccelerateIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBasicAccelerateIpsResponse{}
	_body, _err := client.ListBasicAccelerateIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBasicAcceleratorsWithOptions(request *ListBasicAcceleratorsRequest, runtime *util.RuntimeOptions) (_result *ListBasicAcceleratorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBasicAccelerators"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBasicAcceleratorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBasicAccelerators(request *ListBasicAcceleratorsRequest) (_result *ListBasicAcceleratorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBasicAcceleratorsResponse{}
	_body, _err := client.ListBasicAcceleratorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBasicEndpointsWithOptions(request *ListBasicEndpointsRequest, runtime *util.RuntimeOptions) (_result *ListBasicEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBasicEndpoints"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBasicEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBasicEndpoints(request *ListBasicEndpointsRequest) (_result *ListBasicEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBasicEndpointsResponse{}
	_body, _err := client.ListBasicEndpointsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBusiRegions"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBusiRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListCustomRoutingEndpointGroupDestinationsWithOptions(request *ListCustomRoutingEndpointGroupDestinationsRequest, runtime *util.RuntimeOptions) (_result *ListCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.FromPort)) {
		query["FromPort"] = request.FromPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Protocols)) {
		query["Protocols"] = request.Protocols
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ToPort)) {
		query["ToPort"] = request.ToPort
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomRoutingEndpointGroupDestinations"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomRoutingEndpointGroupDestinations(request *ListCustomRoutingEndpointGroupDestinationsRequest) (_result *ListCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.ListCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomRoutingEndpointGroupsWithOptions(request *ListCustomRoutingEndpointGroupsRequest, runtime *util.RuntimeOptions) (_result *ListCustomRoutingEndpointGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomRoutingEndpointGroups"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomRoutingEndpointGroups(request *ListCustomRoutingEndpointGroupsRequest) (_result *ListCustomRoutingEndpointGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomRoutingEndpointGroupsResponse{}
	_body, _err := client.ListCustomRoutingEndpointGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomRoutingEndpointTrafficPoliciesWithOptions(request *ListCustomRoutingEndpointTrafficPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomRoutingEndpointTrafficPolicies"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomRoutingEndpointTrafficPolicies(request *ListCustomRoutingEndpointTrafficPoliciesRequest) (_result *ListCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.ListCustomRoutingEndpointTrafficPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomRoutingEndpointsWithOptions(request *ListCustomRoutingEndpointsRequest, runtime *util.RuntimeOptions) (_result *ListCustomRoutingEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomRoutingEndpoints"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomRoutingEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomRoutingEndpoints(request *ListCustomRoutingEndpointsRequest) (_result *ListCustomRoutingEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomRoutingEndpointsResponse{}
	_body, _err := client.ListCustomRoutingEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomRoutingPortMappingsWithOptions(request *ListCustomRoutingPortMappingsRequest, runtime *util.RuntimeOptions) (_result *ListCustomRoutingPortMappingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomRoutingPortMappings"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomRoutingPortMappingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomRoutingPortMappings(request *ListCustomRoutingPortMappingsRequest) (_result *ListCustomRoutingPortMappingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomRoutingPortMappingsResponse{}
	_body, _err := client.ListCustomRoutingPortMappingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomRoutingPortMappingsByDestinationWithOptions(request *ListCustomRoutingPortMappingsByDestinationRequest, runtime *util.RuntimeOptions) (_result *ListCustomRoutingPortMappingsByDestinationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationAddress)) {
		query["DestinationAddress"] = request.DestinationAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomRoutingPortMappingsByDestination"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomRoutingPortMappingsByDestinationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomRoutingPortMappingsByDestination(request *ListCustomRoutingPortMappingsByDestinationRequest) (_result *ListCustomRoutingPortMappingsByDestinationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomRoutingPortMappingsByDestinationResponse{}
	_body, _err := client.ListCustomRoutingPortMappingsByDestinationWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLogSwitch)) {
		query["AccessLogSwitch"] = request.AccessLogSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupType)) {
		query["EndpointGroupType"] = request.EndpointGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEndpointGroups"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEndpointGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardingRuleId)) {
		query["ForwardingRuleId"] = request.ForwardingRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListForwardingRules"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListForwardingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIpSets"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIpSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListListenerCertificatesWithOptions(request *ListListenerCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListListenerCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListenerCertificates"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListenerCertificates(request *ListListenerCertificatesRequest) (_result *ListListenerCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.ListListenerCertificatesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListeners"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListSpareIpsWithOptions(request *ListSpareIpsRequest, runtime *util.RuntimeOptions) (_result *ListSpareIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSpareIps"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSpareIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSpareIps(request *ListSpareIpsRequest) (_result *ListSpareIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSpareIpsResponse{}
	_body, _err := client.ListSpareIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSystemSecurityPoliciesWithOptions(request *ListSystemSecurityPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListSystemSecurityPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSystemSecurityPolicies"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSystemSecurityPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSystemSecurityPolicies(request *ListSystemSecurityPoliciesRequest) (_result *ListSystemSecurityPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSystemSecurityPoliciesResponse{}
	_body, _err := client.ListSystemSecurityPoliciesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclEntries)) {
		query["AclEntries"] = request.AclEntries
	}

	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveEntriesFromAcl"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetBandwidthPackageId)) {
		query["TargetBandwidthPackageId"] = request.TargetBandwidthPackageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplaceBandwidthPackage"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplaceBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUseCoupon)) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		query["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateAcceleratorAutoRenewAttributeWithOptions(request *UpdateAcceleratorAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateAcceleratorAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RenewalStatus)) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAcceleratorAutoRenewAttribute"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAcceleratorAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAcceleratorAutoRenewAttribute(request *UpdateAcceleratorAutoRenewAttributeRequest) (_result *UpdateAcceleratorAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAcceleratorAutoRenewAttributeResponse{}
	_body, _err := client.UpdateAcceleratorAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAcceleratorConfirmWithOptions(request *UpdateAcceleratorConfirmRequest, runtime *util.RuntimeOptions) (_result *UpdateAcceleratorConfirmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAcceleratorConfirm"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAcceleratorConfirmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAcceleratorConfirm(request *UpdateAcceleratorConfirmRequest) (_result *UpdateAcceleratorConfirmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAcceleratorConfirmResponse{}
	_body, _err := client.UpdateAcceleratorConfirmWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclId)) {
		query["AclId"] = request.AclId
	}

	if !tea.BoolValue(util.IsUnset(request.AclName)) {
		query["AclName"] = request.AclName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAclAttribute"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateAdditionalCertificateWithListenerWithOptions(request *UpdateAdditionalCertificateWithListenerRequest, runtime *util.RuntimeOptions) (_result *UpdateAdditionalCertificateWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAdditionalCertificateWithListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAdditionalCertificateWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAdditionalCertificateWithListener(request *UpdateAdditionalCertificateWithListenerRequest) (_result *UpdateAdditionalCertificateWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAdditionalCertificateWithListenerResponse{}
	_body, _err := client.UpdateAdditionalCertificateWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationMonitorWithOptions(request *UpdateApplicationMonitorRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DetectEnable)) {
		query["DetectEnable"] = request.DetectEnable
	}

	if !tea.BoolValue(util.IsUnset(request.DetectThreshold)) {
		query["DetectThreshold"] = request.DetectThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.DetectTimes)) {
		query["DetectTimes"] = request.DetectTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.OptionsJson)) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationMonitor"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicationMonitor(request *UpdateApplicationMonitorRequest) (_result *UpdateApplicationMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationMonitorResponse{}
	_body, _err := client.UpdateApplicationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBandwidthPackagaAutoRenewAttributeWithOptions(request *UpdateBandwidthPackagaAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateBandwidthPackagaAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RenewalStatus)) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBandwidthPackagaAutoRenewAttribute"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBandwidthPackagaAutoRenewAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBandwidthPackagaAutoRenewAttribute(request *UpdateBandwidthPackagaAutoRenewAttributeRequest) (_result *UpdateBandwidthPackagaAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBandwidthPackagaAutoRenewAttributeResponse{}
	_body, _err := client.UpdateBandwidthPackagaAutoRenewAttributeWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoUseCoupon)) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		query["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthType)) {
		query["BandwidthType"] = request.BandwidthType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBandwidthPackage"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBandwidthPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateBasicAcceleratorWithOptions(request *UpdateBasicAcceleratorRequest, runtime *util.RuntimeOptions) (_result *UpdateBasicAcceleratorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBasicAccelerator"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBasicAcceleratorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBasicAccelerator(request *UpdateBasicAcceleratorRequest) (_result *UpdateBasicAcceleratorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBasicAcceleratorResponse{}
	_body, _err := client.UpdateBasicAcceleratorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBasicEndpointWithOptions(request *UpdateBasicEndpointRequest, runtime *util.RuntimeOptions) (_result *UpdateBasicEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBasicEndpoint"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBasicEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBasicEndpoint(request *UpdateBasicEndpointRequest) (_result *UpdateBasicEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBasicEndpointResponse{}
	_body, _err := client.UpdateBasicEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBasicEndpointGroupWithOptions(request *UpdateBasicEndpointGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateBasicEndpointGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointAddress)) {
		query["EndpointAddress"] = request.EndpointAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointSubAddress)) {
		query["EndpointSubAddress"] = request.EndpointSubAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBasicEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBasicEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBasicEndpointGroup(request *UpdateBasicEndpointGroupRequest) (_result *UpdateBasicEndpointGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBasicEndpointGroupResponse{}
	_body, _err := client.UpdateBasicEndpointGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBasicIpSetWithOptions(request *UpdateBasicIpSetRequest, runtime *util.RuntimeOptions) (_result *UpdateBasicIpSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBasicIpSet"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBasicIpSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBasicIpSet(request *UpdateBasicIpSetRequest) (_result *UpdateBasicIpSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBasicIpSetResponse{}
	_body, _err := client.UpdateBasicIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomRoutingEndpointGroupAttributeWithOptions(request *UpdateCustomRoutingEndpointGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomRoutingEndpointGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomRoutingEndpointGroupAttribute"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomRoutingEndpointGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomRoutingEndpointGroupAttribute(request *UpdateCustomRoutingEndpointGroupAttributeRequest) (_result *UpdateCustomRoutingEndpointGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomRoutingEndpointGroupAttributeResponse{}
	_body, _err := client.UpdateCustomRoutingEndpointGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomRoutingEndpointGroupDestinationsWithOptions(request *UpdateCustomRoutingEndpointGroupDestinationsRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationConfigurations)) {
		query["DestinationConfigurations"] = request.DestinationConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomRoutingEndpointGroupDestinations"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomRoutingEndpointGroupDestinations(request *UpdateCustomRoutingEndpointGroupDestinationsRequest) (_result *UpdateCustomRoutingEndpointGroupDestinationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomRoutingEndpointGroupDestinationsResponse{}
	_body, _err := client.UpdateCustomRoutingEndpointGroupDestinationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomRoutingEndpointTrafficPoliciesWithOptions(request *UpdateCustomRoutingEndpointTrafficPoliciesRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyConfigurations)) {
		query["PolicyConfigurations"] = request.PolicyConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomRoutingEndpointTrafficPolicies"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomRoutingEndpointTrafficPolicies(request *UpdateCustomRoutingEndpointTrafficPoliciesRequest) (_result *UpdateCustomRoutingEndpointTrafficPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomRoutingEndpointTrafficPoliciesResponse{}
	_body, _err := client.UpdateCustomRoutingEndpointTrafficPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomRoutingEndpointsWithOptions(request *UpdateCustomRoutingEndpointsRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomRoutingEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointConfigurations)) {
		query["EndpointConfigurations"] = request.EndpointConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomRoutingEndpoints"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomRoutingEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomRoutingEndpoints(request *UpdateCustomRoutingEndpointsRequest) (_result *UpdateCustomRoutingEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomRoutingEndpointsResponse{}
	_body, _err := client.UpdateCustomRoutingEndpointsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointConfigurations)) {
		query["EndpointConfigurations"] = request.EndpointConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupRegion)) {
		query["EndpointGroupRegion"] = request.EndpointGroupRegion
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointRequestProtocol)) {
		query["EndpointRequestProtocol"] = request.EndpointRequestProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckEnabled)) {
		query["HealthCheckEnabled"] = request.HealthCheckEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckIntervalSeconds)) {
		query["HealthCheckIntervalSeconds"] = request.HealthCheckIntervalSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckPath)) {
		query["HealthCheckPath"] = request.HealthCheckPath
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckPort)) {
		query["HealthCheckPort"] = request.HealthCheckPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckProtocol)) {
		query["HealthCheckProtocol"] = request.HealthCheckProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PortOverrides)) {
		query["PortOverrides"] = request.PortOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdCount)) {
		query["ThresholdCount"] = request.ThresholdCount
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficPercentage)) {
		query["TrafficPercentage"] = request.TrafficPercentage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEndpointGroup"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEndpointGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupId)) {
		query["EndpointGroupId"] = request.EndpointGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEndpointGroupAttribute"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEndpointGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateEndpointGroupsWithOptions(request *UpdateEndpointGroupsRequest, runtime *util.RuntimeOptions) (_result *UpdateEndpointGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointGroupConfigurations)) {
		query["EndpointGroupConfigurations"] = request.EndpointGroupConfigurations
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEndpointGroups"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEndpointGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEndpointGroups(request *UpdateEndpointGroupsRequest) (_result *UpdateEndpointGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEndpointGroupsResponse{}
	_body, _err := client.UpdateEndpointGroupsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardingRules)) {
		query["ForwardingRules"] = request.ForwardingRules
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateForwardingRules"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateForwardingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.IpSetId)) {
		query["IpSetId"] = request.IpSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpSet"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpSets)) {
		query["IpSets"] = request.IpSets
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIpSets"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIpSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackendPorts)) {
		query["BackendPorts"] = request.BackendPorts
	}

	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientAffinity)) {
		query["ClientAffinity"] = request.ClientAffinity
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PortRanges)) {
		query["PortRanges"] = request.PortRanges
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocol)) {
		query["ProxyProtocol"] = request.ProxyProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedForConfig)) {
		query["XForwardedForConfig"] = request.XForwardedForConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
