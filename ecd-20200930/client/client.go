// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDirectorySsoStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
}

func (s GetDirectorySsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusRequest) SetRegionId(v string) *GetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetDirectorySsoStatusRequest) SetDirectoryId(v string) *GetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

type GetDirectorySsoStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SsoStatus *bool   `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty" require:"true"`
}

func (s GetDirectorySsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusResponse) SetRequestId(v string) *GetDirectorySsoStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetDirectorySsoStatusResponse) SetSsoStatus(v bool) *GetDirectorySsoStatusResponse {
	s.SsoStatus = &v
	return s
}

type SetDirectorySsoStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	EnableSso   *bool   `json:"EnableSso,omitempty" xml:"EnableSso,omitempty" require:"true"`
}

func (s SetDirectorySsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusRequest) SetRegionId(v string) *SetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetDirectoryId(v string) *SetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetEnableSso(v bool) *SetDirectorySsoStatusRequest {
	s.EnableSso = &v
	return s
}

type SetDirectorySsoStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetDirectorySsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusResponse) SetRequestId(v string) *SetDirectorySsoStatusResponse {
	s.RequestId = &v
	return s
}

type GetSpMetadataRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
}

func (s GetSpMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataRequest) GoString() string {
	return s.String()
}

func (s *GetSpMetadataRequest) SetRegionId(v string) *GetSpMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *GetSpMetadataRequest) SetDirectoryId(v string) *GetSpMetadataRequest {
	s.DirectoryId = &v
	return s
}

type GetSpMetadataResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SpMetadata *string `json:"SpMetadata,omitempty" xml:"SpMetadata,omitempty" require:"true"`
}

func (s GetSpMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataResponse) GoString() string {
	return s.String()
}

func (s *GetSpMetadataResponse) SetRequestId(v string) *GetSpMetadataResponse {
	s.RequestId = &v
	return s
}

func (s *GetSpMetadataResponse) SetSpMetadata(v string) *GetSpMetadataResponse {
	s.SpMetadata = &v
	return s
}

type SetIdpMetadataRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	IdpMetadata *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty" require:"true"`
}

func (s SetIdpMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataRequest) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataRequest) SetRegionId(v string) *SetIdpMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetDirectoryId(v string) *SetIdpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetIdpMetadata(v string) *SetIdpMetadataRequest {
	s.IdpMetadata = &v
	return s
}

type SetIdpMetadataResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IdpEntityId *string `json:"IdpEntityId,omitempty" xml:"IdpEntityId,omitempty" require:"true"`
}

func (s SetIdpMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataResponse) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataResponse) SetRequestId(v string) *SetIdpMetadataResponse {
	s.RequestId = &v
	return s
}

func (s *SetIdpMetadataResponse) SetIdpEntityId(v string) *SetIdpMetadataResponse {
	s.IdpEntityId = &v
	return s
}

type RebuildDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s RebuildDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsRequest) SetRegionId(v string) *RebuildDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebuildDesktopsRequest) SetDesktopId(v []*string) *RebuildDesktopsRequest {
	s.DesktopId = v
	return s
}

type RebuildDesktopsResponse struct {
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RebuildResults []*RebuildDesktopsResponseRebuildResults `json:"RebuildResults,omitempty" xml:"RebuildResults,omitempty" require:"true" type:"Repeated"`
}

func (s RebuildDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponse) SetRequestId(v string) *RebuildDesktopsResponse {
	s.RequestId = &v
	return s
}

func (s *RebuildDesktopsResponse) SetRebuildResults(v []*RebuildDesktopsResponseRebuildResults) *RebuildDesktopsResponse {
	s.RebuildResults = v
	return s
}

type RebuildDesktopsResponseRebuildResults struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s RebuildDesktopsResponseRebuildResults) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponseRebuildResults) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponseRebuildResults) SetDesktopId(v string) *RebuildDesktopsResponseRebuildResults {
	s.DesktopId = &v
	return s
}

func (s *RebuildDesktopsResponseRebuildResults) SetCode(v string) *RebuildDesktopsResponseRebuildResults {
	s.Code = &v
	return s
}

func (s *RebuildDesktopsResponseRebuildResults) SetMessage(v string) *RebuildDesktopsResponseRebuildResults {
	s.Message = &v
	return s
}

type ModifyBundleRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	BundleName  *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleRequest) GoString() string {
	return s.String()
}

func (s *ModifyBundleRequest) SetRegionId(v string) *ModifyBundleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBundleRequest) SetBundleId(v string) *ModifyBundleRequest {
	s.BundleId = &v
	return s
}

func (s *ModifyBundleRequest) SetImageId(v string) *ModifyBundleRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyBundleRequest) SetBundleName(v string) *ModifyBundleRequest {
	s.BundleName = &v
	return s
}

func (s *ModifyBundleRequest) SetDescription(v string) *ModifyBundleRequest {
	s.Description = &v
	return s
}

type ModifyBundleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleResponse) GoString() string {
	return s.String()
}

func (s *ModifyBundleResponse) SetRequestId(v string) *ModifyBundleResponse {
	s.RequestId = &v
	return s
}

type UnlockVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s UnlockVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceRequest) SetRegionId(v string) *UnlockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockVirtualMFADeviceRequest) SetSerialNumber(v string) *UnlockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type UnlockVirtualMFADeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnlockVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceResponse) SetRequestId(v string) *UnlockVirtualMFADeviceResponse {
	s.RequestId = &v
	return s
}

type DescribeVirtualMFADevicesRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults  *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	DirectoryId *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	EndUserId   []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
}

func (s DescribeVirtualMFADevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesRequest) SetRegionId(v string) *DescribeVirtualMFADevicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetMaxResults(v int) *DescribeVirtualMFADevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetNextToken(v string) *DescribeVirtualMFADevicesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetDirectoryId(v string) *DescribeVirtualMFADevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetEndUserId(v []*string) *DescribeVirtualMFADevicesRequest {
	s.EndUserId = v
	return s
}

type DescribeVirtualMFADevicesResponse struct {
	NextToken         *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VirtualMFADevices []*DescribeVirtualMFADevicesResponseVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeVirtualMFADevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponse) SetNextToken(v string) *DescribeVirtualMFADevicesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) SetRequestId(v string) *DescribeVirtualMFADevicesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) SetVirtualMFADevices(v []*DescribeVirtualMFADevicesResponseVirtualMFADevices) *DescribeVirtualMFADevicesResponse {
	s.VirtualMFADevices = v
	return s
}

type DescribeVirtualMFADevicesResponseVirtualMFADevices struct {
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	GmtEnabled       *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty" require:"true"`
	GmtUnlock        *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty" require:"true"`
	ConsecutiveFails *int    `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty" require:"true"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s DescribeVirtualMFADevicesResponseVirtualMFADevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetDirectoryId(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetEndUserId(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.EndUserId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetGmtEnabled(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.GmtEnabled = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetGmtUnlock(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetConsecutiveFails(v int) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.ConsecutiveFails = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetSerialNumber(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.SerialNumber = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseVirtualMFADevices) SetStatus(v string) *DescribeVirtualMFADevicesResponseVirtualMFADevices {
	s.Status = &v
	return s
}

type LockVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s LockVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceRequest) SetRegionId(v string) *LockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *LockVirtualMFADeviceRequest) SetSerialNumber(v string) *LockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type LockVirtualMFADeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s LockVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceResponse) SetRequestId(v string) *LockVirtualMFADeviceResponse {
	s.RequestId = &v
	return s
}

type DeleteVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) SetRegionId(v string) *DeleteVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type DeleteVirtualMFADeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponse) SetRequestId(v string) *DeleteVirtualMFADeviceResponse {
	s.RequestId = &v
	return s
}

type ModifyADConnectorDirectoryRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId         *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
}

func (s ModifyADConnectorDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryRequest) SetRegionId(v string) *ModifyADConnectorDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDirectoryId(v string) *ModifyADConnectorDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainName(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainUserName(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainUserName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainPassword(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainPassword = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDnsAddress(v []*string) *ModifyADConnectorDirectoryRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDirectoryName(v string) *ModifyADConnectorDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetSubDomainDnsAddress(v []*string) *ModifyADConnectorDirectoryRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetSubDomainName(v string) *ModifyADConnectorDirectoryRequest {
	s.SubDomainName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetMfaEnabled(v bool) *ModifyADConnectorDirectoryRequest {
	s.MfaEnabled = &v
	return s
}

type ModifyADConnectorDirectoryResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyADConnectorDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryResponse) SetRequestId(v string) *ModifyADConnectorDirectoryResponse {
	s.RequestId = &v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	MaxResults   *int                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetMaxResults(v int) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken    *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	TagResources []*ListTagResourcesResponseTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetRequestId(v string) *ListTagResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponse) SetNextToken(v string) *ListTagResourcesResponse {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponse) SetTagResources(v []*ListTagResourcesResponseTagResources) *ListTagResourcesResponse {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseTagResources struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s ListTagResourcesResponseTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResources) SetResourceType(v string) *ListTagResourcesResponseTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetResourceId(v string) *ListTagResourcesResponseTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetTagKey(v string) *ListTagResourcesResponseTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetTagValue(v string) *ListTagResourcesResponseTagResources {
	s.TagValue = &v
	return s
}

type UntagResourcesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetRequestId(v string) *UntagResourcesResponse {
	s.RequestId = &v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetRequestId(v string) *TagResourcesResponse {
	s.RequestId = &v
	return s
}

type ModifyDesktopNameRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	NewDesktopName *string `json:"NewDesktopName,omitempty" xml:"NewDesktopName,omitempty" require:"true"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
}

func (s ModifyDesktopNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameRequest) SetRegionId(v string) *ModifyDesktopNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetNewDesktopName(v string) *ModifyDesktopNameRequest {
	s.NewDesktopName = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetDesktopId(v string) *ModifyDesktopNameRequest {
	s.DesktopId = &v
	return s
}

type ModifyDesktopNameResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDesktopNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameResponse) SetRequestId(v string) *ModifyDesktopNameResponse {
	s.RequestId = &v
	return s
}

type RunCommandRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Type            *string   `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	CommandContent  *string   `json:"CommandContent,omitempty" xml:"CommandContent,omitempty" require:"true"`
	Timeout         *int64    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	DesktopId       []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
	ContentEncoding *string   `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int64) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetDesktopId(v []*string) *RunCommandRequest {
	s.DesktopId = v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

type RunCommandResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty" require:"true"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetRequestId(v string) *RunCommandResponse {
	s.RequestId = &v
	return s
}

func (s *RunCommandResponse) SetInvokeId(v string) *RunCommandResponse {
	s.InvokeId = &v
	return s
}

type StopInvocationRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InvokeId  *string   `json:"InvokeId,omitempty" xml:"InvokeId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s StopInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) SetRegionId(v string) *StopInvocationRequest {
	s.RegionId = &v
	return s
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetDesktopId(v []*string) *StopInvocationRequest {
	s.DesktopId = v
	return s
}

type StopInvocationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopInvocationResponse) SetRequestId(v string) *StopInvocationResponse {
	s.RequestId = &v
	return s
}

type DescribeInvocationsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InvokeId        *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	CommandType     *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	InvokeStatus    *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	IncludeOutput   *bool   `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	MaxResults      *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetCommandType(v string) *DescribeInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetDesktopId(v string) *DescribeInvocationsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetMaxResults(v int) *DescribeInvocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNextToken(v string) *DescribeInvocationsRequest {
	s.NextToken = &v
	return s
}

type DescribeInvocationsResponse struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Invocations []*DescribeInvocationsResponseInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetRequestId(v string) *DescribeInvocationsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponse) SetNextToken(v string) *DescribeInvocationsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsResponse) SetInvocations(v []*DescribeInvocationsResponseInvocations) *DescribeInvocationsResponse {
	s.Invocations = v
	return s
}

type DescribeInvocationsResponseInvocations struct {
	InvokeId         *string                                                 `json:"InvokeId,omitempty" xml:"InvokeId,omitempty" require:"true"`
	CreationTime     *string                                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	CommandType      *string                                                 `json:"CommandType,omitempty" xml:"CommandType,omitempty" require:"true"`
	CommandContent   *string                                                 `json:"CommandContent,omitempty" xml:"CommandContent,omitempty" require:"true"`
	InvocationStatus *string                                                 `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty" require:"true"`
	InvokeDesktops   []*DescribeInvocationsResponseInvocationsInvokeDesktops `json:"InvokeDesktops,omitempty" xml:"InvokeDesktops,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInvocationsResponseInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseInvocations) SetInvokeId(v string) *DescribeInvocationsResponseInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetCreationTime(v string) *DescribeInvocationsResponseInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetCommandType(v string) *DescribeInvocationsResponseInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetCommandContent(v string) *DescribeInvocationsResponseInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetInvocationStatus(v string) *DescribeInvocationsResponseInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseInvocations) SetInvokeDesktops(v []*DescribeInvocationsResponseInvocationsInvokeDesktops) *DescribeInvocationsResponseInvocations {
	s.InvokeDesktops = v
	return s
}

type DescribeInvocationsResponseInvocationsInvokeDesktops struct {
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	Repeats          *int    `json:"Repeats,omitempty" xml:"Repeats,omitempty" require:"true"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty" require:"true"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	StopTime         *string `json:"StopTime,omitempty" xml:"StopTime,omitempty" require:"true"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	Output           *string `json:"Output,omitempty" xml:"Output,omitempty" require:"true"`
	ExitCode         *int64  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty" require:"true"`
	Dropped          *int    `json:"Dropped,omitempty" xml:"Dropped,omitempty" require:"true"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	ErrorInfo        *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" require:"true"`
}

func (s DescribeInvocationsResponseInvocationsInvokeDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseInvocationsInvokeDesktops) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetDesktopId(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetRepeats(v int) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetInvocationStatus(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetCreationTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetStartTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetStopTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetFinishTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetUpdateTime(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetOutput(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetExitCode(v int64) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetDropped(v int) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetErrorCode(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseInvocationsInvokeDesktops) SetErrorInfo(v string) *DescribeInvocationsResponseInvocationsInvokeDesktops {
	s.ErrorInfo = &v
	return s
}

type DescribeZonesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

type DescribeZonesResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Zones     []*DescribeZonesResponseZones `json:"Zones,omitempty" xml:"Zones,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetRequestId(v string) *DescribeZonesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponse) SetZones(v []*DescribeZonesResponseZones) *DescribeZonesResponse {
	s.Zones = v
	return s
}

type DescribeZonesResponseZones struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
}

func (s DescribeZonesResponseZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseZones) SetZoneId(v string) *DescribeZonesResponseZones {
	s.ZoneId = &v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
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

type DescribeRegionsResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Regions   []*DescribeRegionsResponseRegions `json:"Regions,omitempty" xml:"Regions,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetRequestId(v string) *DescribeRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponse) SetRegions(v []*DescribeRegionsResponseRegions) *DescribeRegionsResponse {
	s.Regions = v
	return s
}

type DescribeRegionsResponseRegions struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty" require:"true"`
}

func (s DescribeRegionsResponseRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegions) SetRegionId(v string) *DescribeRegionsResponseRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseRegions {
	s.RegionEndpoint = &v
	return s
}

type DescribeClientEventsRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EndUserId   *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopIp   *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EventType   *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults  *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeClientEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsRequest) SetRegionId(v string) *DescribeClientEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndUserId(v string) *DescribeClientEventsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopId(v string) *DescribeClientEventsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopIp(v string) *DescribeClientEventsRequest {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDirectoryId(v string) *DescribeClientEventsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEventType(v string) *DescribeClientEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsRequest) SetStartTime(v string) *DescribeClientEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndTime(v string) *DescribeClientEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClientEventsRequest) SetMaxResults(v int) *DescribeClientEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeClientEventsRequest) SetNextToken(v string) *DescribeClientEventsRequest {
	s.NextToken = &v
	return s
}

type DescribeClientEventsResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Events    []*DescribeClientEventsResponseEvents `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeClientEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponse) SetRequestId(v string) *DescribeClientEventsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeClientEventsResponse) SetNextToken(v string) *DescribeClientEventsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeClientEventsResponse) SetEvents(v []*DescribeClientEventsResponseEvents) *DescribeClientEventsResponse {
	s.Events = v
	return s
}

type DescribeClientEventsResponseEvents struct {
	EventId       *string `json:"EventId,omitempty" xml:"EventId,omitempty" require:"true"`
	EventType     *string `json:"EventType,omitempty" xml:"EventType,omitempty" require:"true"`
	EventTime     *string `json:"EventTime,omitempty" xml:"EventTime,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty" require:"true"`
	EndUserId     *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopIp     *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty" require:"true"`
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty" require:"true"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty" require:"true"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty" require:"true"`
	BytesSend     *string `json:"BytesSend,omitempty" xml:"BytesSend,omitempty" require:"true"`
	BytesReceived *string `json:"BytesReceived,omitempty" xml:"BytesReceived,omitempty" require:"true"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeClientEventsResponseEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponseEvents) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseEvents) SetEventId(v string) *DescribeClientEventsResponseEvents {
	s.EventId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetEventType(v string) *DescribeClientEventsResponseEvents {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetEventTime(v string) *DescribeClientEventsResponseEvents {
	s.EventTime = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetRegionId(v string) *DescribeClientEventsResponseEvents {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetAliUid(v string) *DescribeClientEventsResponseEvents {
	s.AliUid = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetEndUserId(v string) *DescribeClientEventsResponseEvents {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetDesktopId(v string) *DescribeClientEventsResponseEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetDesktopIp(v string) *DescribeClientEventsResponseEvents {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetDirectoryId(v string) *DescribeClientEventsResponseEvents {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetDirectoryType(v string) *DescribeClientEventsResponseEvents {
	s.DirectoryType = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetClientOS(v string) *DescribeClientEventsResponseEvents {
	s.ClientOS = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetClientVersion(v string) *DescribeClientEventsResponseEvents {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetClientIp(v string) *DescribeClientEventsResponseEvents {
	s.ClientIp = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetBytesSend(v string) *DescribeClientEventsResponseEvents {
	s.BytesSend = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetBytesReceived(v string) *DescribeClientEventsResponseEvents {
	s.BytesReceived = &v
	return s
}

func (s *DescribeClientEventsResponseEvents) SetStatus(v string) *DescribeClientEventsResponseEvents {
	s.Status = &v
	return s
}

type ResetSnapshotRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" require:"true"`
}

func (s ResetSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ResetSnapshotRequest) SetRegionId(v string) *ResetSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSnapshotId(v string) *ResetSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type ResetSnapshotResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ResetSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponse) SetRequestId(v string) *ResetSnapshotResponse {
	s.RequestId = &v
	return s
}

type DeleteSnapshotRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SnapshotId []*string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v []*string) *DeleteSnapshotRequest {
	s.SnapshotId = v
	return s
}

type DeleteSnapshotResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetRequestId(v string) *DeleteSnapshotResponse {
	s.RequestId = &v
	return s
}

type CreateSnapshotRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty" require:"true"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetDesktopId(v string) *CreateSnapshotRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetSourceDiskType(v string) *CreateSnapshotRequest {
	s.SourceDiskType = &v
	return s
}

type CreateSnapshotResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetRequestId(v string) *CreateSnapshotResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponse) SetSnapshotId(v string) *CreateSnapshotResponse {
	s.SnapshotId = &v
	return s
}

type DescribeSnapshotsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	MaxResults *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDesktopId(v string) *DescribeSnapshotsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotId(v string) *DescribeSnapshotsRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

type DescribeSnapshotsResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Snapshots []*DescribeSnapshotsResponseSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) SetRequestId(v string) *DescribeSnapshotsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetNextToken(v string) *DescribeSnapshotsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetSnapshots(v []*DescribeSnapshotsResponseSnapshots) *DescribeSnapshotsResponse {
	s.Snapshots = v
	return s
}

type DescribeSnapshotsResponseSnapshots struct {
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" require:"true"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	SnapshotType   *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty" require:"true"`
	SourceDiskSize *string `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty" require:"true"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty" require:"true"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty" require:"true"`
	RemainTime     *int    `json:"RemainTime,omitempty" xml:"RemainTime,omitempty" require:"true"`
}

func (s DescribeSnapshotsResponseSnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseSnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseSnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseSnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetDesktopId(v string) *DescribeSnapshotsResponseSnapshots {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetSnapshotName(v string) *DescribeSnapshotsResponseSnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetDescription(v string) *DescribeSnapshotsResponseSnapshots {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetSnapshotType(v string) *DescribeSnapshotsResponseSnapshots {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetSourceDiskSize(v string) *DescribeSnapshotsResponseSnapshots {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetSourceDiskType(v string) *DescribeSnapshotsResponseSnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetStatus(v string) *DescribeSnapshotsResponseSnapshots {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetCreationTime(v string) *DescribeSnapshotsResponseSnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetProgress(v string) *DescribeSnapshotsResponseSnapshots {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseSnapshots) SetRemainTime(v int) *DescribeSnapshotsResponseSnapshots {
	s.RemainTime = &v
	return s
}

type RenewDesktopsRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId  []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
	Period     *int      `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay    *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
}

func (s RenewDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RenewDesktopsRequest) SetRegionId(v string) *RenewDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RenewDesktopsRequest) SetDesktopId(v []*string) *RenewDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RenewDesktopsRequest) SetPeriod(v int) *RenewDesktopsRequest {
	s.Period = &v
	return s
}

func (s *RenewDesktopsRequest) SetPeriodUnit(v string) *RenewDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDesktopsRequest) SetAutoPay(v bool) *RenewDesktopsRequest {
	s.AutoPay = &v
	return s
}

type RenewDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
}

func (s RenewDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponse) SetRequestId(v string) *RenewDesktopsResponse {
	s.RequestId = &v
	return s
}

func (s *RenewDesktopsResponse) SetOrderId(v string) *RenewDesktopsResponse {
	s.OrderId = &v
	return s
}

type StopDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s StopDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

type StopDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponse) SetRequestId(v string) *StopDesktopsResponse {
	s.RequestId = &v
	return s
}

type StartDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s StartDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

type StartDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StartDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponse) SetRequestId(v string) *StartDesktopsResponse {
	s.RequestId = &v
	return s
}

type ModifyPolicyGroupRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PolicyGroupId         *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Clipboard             *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive            *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect           *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark             *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	WatermarkType         *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkCustomText   *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
}

func (s ModifyPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequest) SetRegionId(v string) *ModifyPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetName(v string) *ModifyPolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClipboard(v string) *ModifyPolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLocalDrive(v string) *ModifyPolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetUsbRedirect(v string) *ModifyPolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermark(v string) *ModifyPolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkType(v string) *ModifyPolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkCustomText(v string) *ModifyPolicyGroupRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkTransparency(v string) *ModifyPolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

type ModifyPolicyGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponse) SetRequestId(v string) *ModifyPolicyGroupResponse {
	s.RequestId = &v
	return s
}

type DescribeDesktopTypesRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopTypeId      *string `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
}

func (s DescribeDesktopTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesRequest) SetRegionId(v string) *DescribeDesktopTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopTypeId(v string) *DescribeDesktopTypesRequest {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetInstanceTypeFamily(v string) *DescribeDesktopTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

type DescribeDesktopTypesResponse struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DesktopTypes []*DescribeDesktopTypesResponseDesktopTypes `json:"DesktopTypes,omitempty" xml:"DesktopTypes,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponse) SetRequestId(v string) *DescribeDesktopTypesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopTypesResponse) SetDesktopTypes(v []*DescribeDesktopTypesResponseDesktopTypes) *DescribeDesktopTypesResponse {
	s.DesktopTypes = v
	return s
}

type DescribeDesktopTypesResponseDesktopTypes struct {
	DesktopTypeId      *string                                                  `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty" require:"true"`
	InstanceTypeFamily *string                                                  `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty" require:"true"`
	CpuCount           *string                                                  `json:"CpuCount,omitempty" xml:"CpuCount,omitempty" require:"true"`
	MemorySize         *string                                                  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty" require:"true"`
	GpuCount           *float32                                                 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty" require:"true"`
	GpuSpec            *string                                                  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty" require:"true"`
	SystemDiskSize     *string                                                  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskSize       *string                                                  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	AllowDiskSize      []*DescribeDesktopTypesResponseDesktopTypesAllowDiskSize `json:"AllowDiskSize,omitempty" xml:"AllowDiskSize,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopTypesResponseDesktopTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseDesktopTypes) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetDesktopTypeId(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetInstanceTypeFamily(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetCpuCount(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetMemorySize(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetGpuCount(v float32) *DescribeDesktopTypesResponseDesktopTypes {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetGpuSpec(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetSystemDiskSize(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetDataDiskSize(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetAllowDiskSize(v []*DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) *DescribeDesktopTypesResponseDesktopTypes {
	s.AllowDiskSize = v
	return s
}

type DescribeDesktopTypesResponseDesktopTypesAllowDiskSize struct {
	SystemDiskSize *int `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskSize   *int `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
}

func (s DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) SetSystemDiskSize(v int) *DescribeDesktopTypesResponseDesktopTypesAllowDiskSize {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypesAllowDiskSize) SetDataDiskSize(v int) *DescribeDesktopTypesResponseDesktopTypesAllowDiskSize {
	s.DataDiskSize = &v
	return s
}

type DescribeDirectoriesRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryType *string   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	DirectoryId   []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	MaxResults    *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryType(v string) *DescribeDirectoriesRequest {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetMaxResults(v int) *DescribeDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNextToken(v string) *DescribeDirectoriesRequest {
	s.NextToken = &v
	return s
}

type DescribeDirectoriesResponse struct {
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Directories []*DescribeDirectoriesResponseDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) SetNextToken(v string) *DescribeDirectoriesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesResponse) SetRequestId(v string) *DescribeDirectoriesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDirectoriesResponse) SetDirectories(v []*DescribeDirectoriesResponseDirectories) *DescribeDirectoriesResponse {
	s.Directories = v
	return s
}

type DescribeDirectoriesResponseDirectories struct {
	DirectoryId           *string                                               `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	Status                *string                                               `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DirectoryType         *string                                               `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
	CreationTime          *string                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	Name                  *string                                               `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	VpcId                 *string                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	CustomSecurityGroupId *string                                               `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty" require:"true"`
	DnsUserName           *string                                               `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty" require:"true"`
	EnableInternetAccess  *bool                                                 `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty" require:"true"`
	EnableAdminAccess     *bool                                                 `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty" require:"true"`
	DesktopAccessType     *string                                               `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty" require:"true"`
	DesktopVpcEndpoint    *string                                               `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty" require:"true"`
	TrustPassword         *string                                               `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty" require:"true"`
	DomainName            *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName        *string                                               `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword        *string                                               `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	SubDomainName         *string                                               `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty" require:"true"`
	MfaEnabled            *bool                                                 `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty" require:"true"`
	SsoEnabled            *bool                                                 `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty" require:"true"`
	ADConnectors          []*DescribeDirectoriesResponseDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" require:"true" type:"Repeated"`
	Logs                  []*DescribeDirectoriesResponseDirectoriesLogs         `json:"Logs,omitempty" xml:"Logs,omitempty" require:"true" type:"Repeated"`
	DnsAddress            []*string                                             `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" require:"true" type:"Repeated"`
	SubDnsAddress         []*string                                             `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" require:"true" type:"Repeated"`
	VSwitchIds            []*string                                             `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDirectoriesResponseDirectories) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetStatus(v string) *DescribeDirectoriesResponseDirectories {
	s.Status = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetCreationTime(v string) *DescribeDirectoriesResponseDirectories {
	s.CreationTime = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetName(v string) *DescribeDirectoriesResponseDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetVpcId(v string) *DescribeDirectoriesResponseDirectories {
	s.VpcId = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetCustomSecurityGroupId(v string) *DescribeDirectoriesResponseDirectories {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDnsUserName(v string) *DescribeDirectoriesResponseDirectories {
	s.DnsUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetEnableInternetAccess(v bool) *DescribeDirectoriesResponseDirectories {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetEnableAdminAccess(v bool) *DescribeDirectoriesResponseDirectories {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDesktopVpcEndpoint(v string) *DescribeDirectoriesResponseDirectories {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetTrustPassword(v string) *DescribeDirectoriesResponseDirectories {
	s.TrustPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDomainName(v string) *DescribeDirectoriesResponseDirectories {
	s.DomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDomainUserName(v string) *DescribeDirectoriesResponseDirectories {
	s.DomainUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDomainPassword(v string) *DescribeDirectoriesResponseDirectories {
	s.DomainPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetSubDomainName(v string) *DescribeDirectoriesResponseDirectories {
	s.SubDomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetMfaEnabled(v bool) *DescribeDirectoriesResponseDirectories {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetSsoEnabled(v bool) *DescribeDirectoriesResponseDirectories {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetADConnectors(v []*DescribeDirectoriesResponseDirectoriesADConnectors) *DescribeDirectoriesResponseDirectories {
	s.ADConnectors = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetLogs(v []*DescribeDirectoriesResponseDirectoriesLogs) *DescribeDirectoriesResponseDirectories {
	s.Logs = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDnsAddress(v []*string) *DescribeDirectoriesResponseDirectories {
	s.DnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetSubDnsAddress(v []*string) *DescribeDirectoriesResponseDirectories {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetVSwitchIds(v []*string) *DescribeDirectoriesResponseDirectories {
	s.VSwitchIds = v
	return s
}

type DescribeDirectoriesResponseDirectoriesADConnectors struct {
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty" require:"true"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty" require:"true"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" require:"true"`
}

func (s DescribeDirectoriesResponseDirectoriesADConnectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseDirectoriesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseDirectoriesADConnectors) SetADConnectorAddress(v string) *DescribeDirectoriesResponseDirectoriesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesADConnectors) SetVSwitchId(v string) *DescribeDirectoriesResponseDirectoriesADConnectors {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesADConnectors) SetConnectorStatus(v string) *DescribeDirectoriesResponseDirectoriesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesADConnectors) SetNetworkInterfaceId(v string) *DescribeDirectoriesResponseDirectoriesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

type DescribeDirectoriesResponseDirectoriesLogs struct {
	Level     *string `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty" require:"true"`
	Step      *string `json:"Step,omitempty" xml:"Step,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s DescribeDirectoriesResponseDirectoriesLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseDirectoriesLogs) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseDirectoriesLogs) SetLevel(v string) *DescribeDirectoriesResponseDirectoriesLogs {
	s.Level = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesLogs) SetTimeStamp(v string) *DescribeDirectoriesResponseDirectoriesLogs {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesLogs) SetStep(v string) *DescribeDirectoriesResponseDirectoriesLogs {
	s.Step = &v
	return s
}

func (s *DescribeDirectoriesResponseDirectoriesLogs) SetMessage(v string) *DescribeDirectoriesResponseDirectoriesLogs {
	s.Message = &v
	return s
}

type DeleteDirectoriesRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesRequest) SetRegionId(v string) *DeleteDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDirectoriesRequest) SetDirectoryId(v []*string) *DeleteDirectoriesRequest {
	s.DirectoryId = v
	return s
}

type DeleteDirectoriesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponse) SetRequestId(v string) *DeleteDirectoriesResponse {
	s.RequestId = &v
	return s
}

type ListDirectoryUsersRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Filter      *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDirectoryUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersRequest) SetRegionId(v string) *ListDirectoryUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetFilter(v string) *ListDirectoryUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetDirectoryId(v string) *ListDirectoryUsersRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetNextToken(v string) *ListDirectoryUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetMaxResults(v int) *ListDirectoryUsersRequest {
	s.MaxResults = &v
	return s
}

type ListDirectoryUsersResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Users     []*ListDirectoryUsersResponseUsers `json:"Users,omitempty" xml:"Users,omitempty" require:"true" type:"Repeated"`
}

func (s ListDirectoryUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponse) SetRequestId(v string) *ListDirectoryUsersResponse {
	s.RequestId = &v
	return s
}

func (s *ListDirectoryUsersResponse) SetNextToken(v string) *ListDirectoryUsersResponse {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersResponse) SetUsers(v []*ListDirectoryUsersResponseUsers) *ListDirectoryUsersResponse {
	s.Users = v
	return s
}

type ListDirectoryUsersResponseUsers struct {
	EndUser *string `json:"EndUser,omitempty" xml:"EndUser,omitempty" require:"true"`
}

func (s ListDirectoryUsersResponseUsers) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponseUsers) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponseUsers) SetEndUser(v string) *ListDirectoryUsersResponseUsers {
	s.EndUser = &v
	return s
}

type CreateImageRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	ImageName   *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageRequest) SetDesktopId(v string) *CreateImageRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

type CreateImageResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetRequestId(v string) *CreateImageResponse {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponse) SetImageId(v string) *CreateImageResponse {
	s.ImageId = &v
	return s
}

type CreateRAMDirectoryRequest struct {
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryName        *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	EnableAdminAccess    *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType    *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	VSwitchId            []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true" type:"Repeated"`
}

func (s CreateRAMDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryRequest) SetRegionId(v string) *CreateRAMDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetDirectoryName(v string) *CreateRAMDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableInternetAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableAdminAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetDesktopAccessType(v string) *CreateRAMDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetVSwitchId(v []*string) *CreateRAMDirectoryRequest {
	s.VSwitchId = v
	return s
}

type CreateRAMDirectoryResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
}

func (s CreateRAMDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryResponse) SetRequestId(v string) *CreateRAMDirectoryResponse {
	s.RequestId = &v
	return s
}

func (s *CreateRAMDirectoryResponse) SetDirectoryId(v string) *CreateRAMDirectoryResponse {
	s.DirectoryId = &v
	return s
}

type DeletePolicyGroupsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true" type:"Repeated"`
}

func (s DeletePolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsRequest) SetRegionId(v string) *DeletePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DeletePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

type DeletePolicyGroupsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeletePolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsResponse) SetRequestId(v string) *DeletePolicyGroupsResponse {
	s.RequestId = &v
	return s
}

type DescribePolicyGroupsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults    *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
}

func (s DescribePolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsRequest) SetRegionId(v string) *DescribePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetMaxResults(v int) *DescribePolicyGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetNextToken(v string) *DescribePolicyGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DescribePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

type DescribePolicyGroupsResponse struct {
	NextToken            *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId            *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DescribePolicyGroups []*DescribePolicyGroupsResponseDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponse) SetNextToken(v string) *DescribePolicyGroupsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsResponse) SetRequestId(v string) *DescribePolicyGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyGroupsResponse) SetDescribePolicyGroups(v []*DescribePolicyGroupsResponseDescribePolicyGroups) *DescribePolicyGroupsResponse {
	s.DescribePolicyGroups = v
	return s
}

type DescribePolicyGroupsResponseDescribePolicyGroups struct {
	PolicyGroupId         *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	PolicyGroupType       *string `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty" require:"true"`
	Clipboard             *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty" require:"true"`
	LocalDrive            *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty" require:"true"`
	UsbRedirect           *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty" require:"true"`
	Watermark             *string `json:"Watermark,omitempty" xml:"Watermark,omitempty" require:"true"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	WatermarkType         *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty" require:"true"`
	WatermarkCustomText   *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty" require:"true"`
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty" require:"true"`
	PolicyStatus          *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty" require:"true"`
	EdsCount              *int    `json:"EdsCount,omitempty" xml:"EdsCount,omitempty" require:"true"`
}

func (s DescribePolicyGroupsResponseDescribePolicyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseDescribePolicyGroups) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetPolicyGroupId(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetPolicyGroupType(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.PolicyGroupType = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetClipboard(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.Clipboard = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetLocalDrive(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.LocalDrive = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetUsbRedirect(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.UsbRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetWatermark(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.Watermark = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetName(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.Name = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetWatermarkType(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.WatermarkType = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetWatermarkCustomText(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.WatermarkCustomText = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetWatermarkTransparency(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.WatermarkTransparency = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetPolicyStatus(v string) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.PolicyStatus = &v
	return s
}

func (s *DescribePolicyGroupsResponseDescribePolicyGroups) SetEdsCount(v int) *DescribePolicyGroupsResponseDescribePolicyGroups {
	s.EdsCount = &v
	return s
}

type DeleteDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsRequest) SetRegionId(v string) *DeleteDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDesktopsRequest) SetDesktopId(v []*string) *DeleteDesktopsRequest {
	s.DesktopId = v
	return s
}

type DeleteDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsResponse) SetRequestId(v string) *DeleteDesktopsResponse {
	s.RequestId = &v
	return s
}

type ModifyImageAttributeRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyImageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) SetRegionId(v string) *ModifyImageAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetName(v string) *ModifyImageAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetDescription(v string) *ModifyImageAttributeRequest {
	s.Description = &v
	return s
}

type ModifyImageAttributeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyImageAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponse) SetRequestId(v string) *ModifyImageAttributeResponse {
	s.RequestId = &v
	return s
}

type ModifyEntitlementRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyEntitlementRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementRequest) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementRequest) SetRegionId(v string) *ModifyEntitlementRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEntitlementRequest) SetDesktopId(v string) *ModifyEntitlementRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyEntitlementRequest) SetEndUserId(v []*string) *ModifyEntitlementRequest {
	s.EndUserId = v
	return s
}

type ModifyEntitlementResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyEntitlementResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementResponse) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementResponse) SetRequestId(v string) *ModifyEntitlementResponse {
	s.RequestId = &v
	return s
}

type DeleteBundlesRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesRequest) GoString() string {
	return s.String()
}

func (s *DeleteBundlesRequest) SetRegionId(v string) *DeleteBundlesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBundlesRequest) SetBundleId(v []*string) *DeleteBundlesRequest {
	s.BundleId = v
	return s
}

type DeleteBundlesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesResponse) GoString() string {
	return s.String()
}

func (s *DeleteBundlesResponse) SetRequestId(v string) *DeleteBundlesResponse {
	s.RequestId = &v
	return s
}

type DescribeDesktopsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	GroupId       *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	DesktopStatus *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	MaxResults    *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserName      *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	DesktopName   *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DirectoryId   *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	EndUserId     []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
}

func (s DescribeDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsRequest) SetRegionId(v string) *DescribeDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetGroupId(v string) *DescribeDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopStatus(v string) *DescribeDesktopsRequest {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsRequest) SetMaxResults(v int) *DescribeDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsRequest) SetNextToken(v string) *DescribeDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsRequest) SetUserName(v string) *DescribeDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopName(v string) *DescribeDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDirectoryId(v string) *DescribeDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetPolicyGroupId(v string) *DescribeDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopId(v []*string) *DescribeDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopsRequest) SetEndUserId(v []*string) *DescribeDesktopsRequest {
	s.EndUserId = v
	return s
}

type DescribeDesktopsResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	Desktops  []*DescribeDesktopsResponseDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponse) SetRequestId(v string) *DescribeDesktopsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsResponse) SetNextToken(v string) *DescribeDesktopsResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsResponse) SetDesktops(v []*DescribeDesktopsResponseDesktops) *DescribeDesktopsResponse {
	s.Desktops = v
	return s
}

type DescribeDesktopsResponseDesktops struct {
	DirectoryId        *string                                     `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	DirectoryType      *string                                     `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
	CreationTime       *string                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	DesktopId          *string                                     `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	BundleId           *string                                     `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	DesktopStatus      *string                                     `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty" require:"true"`
	DesktopName        *string                                     `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
	ImageId            *string                                     `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	DesktopType        *string                                     `json:"DesktopType,omitempty" xml:"DesktopType,omitempty" require:"true"`
	SystemDiskCategory *string                                     `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" require:"true"`
	SystemDiskSize     *int                                        `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskCategory   *string                                     `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty" require:"true"`
	DataDiskSize       *string                                     `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	PolicyGroupId      *string                                     `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	Cpu                *int                                        `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory             *int64                                      `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	NetworkInterfaceId *string                                     `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" require:"true"`
	ExpiredTime        *string                                     `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" require:"true"`
	ChargeType         *string                                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty" require:"true"`
	GpuCount           *float32                                    `json:"GpuCount,omitempty" xml:"GpuCount,omitempty" require:"true"`
	GpuSpec            *string                                     `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty" require:"true"`
	StartTime          *string                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	ConnectionStatus   *string                                     `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty" require:"true"`
	OsType             *string                                     `json:"OsType,omitempty" xml:"OsType,omitempty" require:"true"`
	Disks              []*DescribeDesktopsResponseDesktopsDisks    `json:"Disks,omitempty" xml:"Disks,omitempty" require:"true" type:"Repeated"`
	Tags               []*DescribeDesktopsResponseDesktopsTags     `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Repeated"`
	Sessions           []*DescribeDesktopsResponseDesktopsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" require:"true" type:"Repeated"`
	EndUserIds         []*string                                   `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopsResponseDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseDesktops) SetDirectoryId(v string) *DescribeDesktopsResponseDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDirectoryType(v string) *DescribeDesktopsResponseDesktops {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetCreationTime(v string) *DescribeDesktopsResponseDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopId(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetBundleId(v string) *DescribeDesktopsResponseDesktops {
	s.BundleId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopStatus(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopName(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetImageId(v string) *DescribeDesktopsResponseDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopType(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetSystemDiskCategory(v string) *DescribeDesktopsResponseDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetSystemDiskSize(v int) *DescribeDesktopsResponseDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDataDiskCategory(v string) *DescribeDesktopsResponseDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDataDiskSize(v string) *DescribeDesktopsResponseDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetPolicyGroupId(v string) *DescribeDesktopsResponseDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetCpu(v int) *DescribeDesktopsResponseDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetMemory(v int64) *DescribeDesktopsResponseDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetNetworkInterfaceId(v string) *DescribeDesktopsResponseDesktops {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetExpiredTime(v string) *DescribeDesktopsResponseDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetChargeType(v string) *DescribeDesktopsResponseDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetGpuCount(v float32) *DescribeDesktopsResponseDesktops {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetGpuSpec(v string) *DescribeDesktopsResponseDesktops {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetStartTime(v string) *DescribeDesktopsResponseDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetConnectionStatus(v string) *DescribeDesktopsResponseDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetOsType(v string) *DescribeDesktopsResponseDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDisks(v []*DescribeDesktopsResponseDesktopsDisks) *DescribeDesktopsResponseDesktops {
	s.Disks = v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetTags(v []*DescribeDesktopsResponseDesktopsTags) *DescribeDesktopsResponseDesktops {
	s.Tags = v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetSessions(v []*DescribeDesktopsResponseDesktopsSessions) *DescribeDesktopsResponseDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetEndUserIds(v []*string) *DescribeDesktopsResponseDesktops {
	s.EndUserIds = v
	return s
}

type DescribeDesktopsResponseDesktopsDisks struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty" require:"true"`
	DiskSize *int    `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" require:"true"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty" require:"true"`
}

func (s DescribeDesktopsResponseDesktopsDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseDesktopsDisks) SetDiskId(v string) *DescribeDesktopsResponseDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktopsDisks) SetDiskSize(v int) *DescribeDesktopsResponseDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseDesktopsDisks) SetDiskType(v string) *DescribeDesktopsResponseDesktopsDisks {
	s.DiskType = &v
	return s
}

type DescribeDesktopsResponseDesktopsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeDesktopsResponseDesktopsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseDesktopsTags) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseDesktopsTags) SetKey(v string) *DescribeDesktopsResponseDesktopsTags {
	s.Key = &v
	return s
}

func (s *DescribeDesktopsResponseDesktopsTags) SetValue(v string) *DescribeDesktopsResponseDesktopsTags {
	s.Value = &v
	return s
}

type DescribeDesktopsResponseDesktopsSessions struct {
	EndUserId         *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true"`
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty" require:"true"`
}

func (s DescribeDesktopsResponseDesktopsSessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseDesktopsSessions) SetEndUserId(v string) *DescribeDesktopsResponseDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsResponseDesktopsSessions) SetEstablishmentTime(v string) *DescribeDesktopsResponseDesktopsSessions {
	s.EstablishmentTime = &v
	return s
}

type RebootDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s RebootDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

type RebootDesktopsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RebootDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponse) SetRequestId(v string) *RebootDesktopsResponse {
	s.RequestId = &v
	return s
}

type CreateBundleRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	DesktopType     *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty" require:"true"`
	RootDiskSizeGib *int    `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty" require:"true"`
	BundleName      *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserDiskSizeGib []*int  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty" require:"true" type:"Repeated"`
}

func (s CreateBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleRequest) GoString() string {
	return s.String()
}

func (s *CreateBundleRequest) SetRegionId(v string) *CreateBundleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBundleRequest) SetImageId(v string) *CreateBundleRequest {
	s.ImageId = &v
	return s
}

func (s *CreateBundleRequest) SetDesktopType(v string) *CreateBundleRequest {
	s.DesktopType = &v
	return s
}

func (s *CreateBundleRequest) SetRootDiskSizeGib(v int) *CreateBundleRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *CreateBundleRequest) SetBundleName(v string) *CreateBundleRequest {
	s.BundleName = &v
	return s
}

func (s *CreateBundleRequest) SetDescription(v string) *CreateBundleRequest {
	s.Description = &v
	return s
}

func (s *CreateBundleRequest) SetUserDiskSizeGib(v []*int) *CreateBundleRequest {
	s.UserDiskSizeGib = v
	return s
}

type CreateBundleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BundleId  *string `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
}

func (s CreateBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleResponse) GoString() string {
	return s.String()
}

func (s *CreateBundleResponse) SetRequestId(v string) *CreateBundleResponse {
	s.RequestId = &v
	return s
}

func (s *CreateBundleResponse) SetBundleId(v string) *CreateBundleResponse {
	s.BundleId = &v
	return s
}

type ModifyDesktopsPolicyGroupRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyDesktopsPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupRequest) SetRegionId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetDesktopId(v []*string) *ModifyDesktopsPolicyGroupRequest {
	s.DesktopId = v
	return s
}

type ModifyDesktopsPolicyGroupResponse struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ModifyResults []*ModifyDesktopsPolicyGroupResponseModifyResults `json:"ModifyResults,omitempty" xml:"ModifyResults,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyDesktopsPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponse) SetRequestId(v string) *ModifyDesktopsPolicyGroupResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponse) SetModifyResults(v []*ModifyDesktopsPolicyGroupResponseModifyResults) *ModifyDesktopsPolicyGroupResponse {
	s.ModifyResults = v
	return s
}

type ModifyDesktopsPolicyGroupResponseModifyResults struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s ModifyDesktopsPolicyGroupResponseModifyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponseModifyResults) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponseModifyResults) SetDesktopId(v string) *ModifyDesktopsPolicyGroupResponseModifyResults {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseModifyResults) SetCode(v string) *ModifyDesktopsPolicyGroupResponseModifyResults {
	s.Code = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseModifyResults) SetMessage(v string) *ModifyDesktopsPolicyGroupResponseModifyResults {
	s.Message = &v
	return s
}

type CreatePolicyGroupRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Clipboard             *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive            *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect           *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark             *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WatermarkType         *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkCustomText   *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
}

func (s CreatePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequest) SetRegionId(v string) *CreatePolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetClipboard(v string) *CreatePolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLocalDrive(v string) *CreatePolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetUsbRedirect(v string) *CreatePolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermark(v string) *CreatePolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetName(v string) *CreatePolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkType(v string) *CreatePolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkCustomText(v string) *CreatePolicyGroupRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkTransparency(v string) *CreatePolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

type CreatePolicyGroupResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
}

func (s CreatePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponse) SetRequestId(v string) *CreatePolicyGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyGroupResponse) SetPolicyGroupId(v string) *CreatePolicyGroupResponse {
	s.PolicyGroupId = &v
	return s
}

type CreateADConnectorDirectoryRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" require:"true" type:"Repeated"`
	VSwitchId           []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true" type:"Repeated"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	EnableAdminAccess   *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType   *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
}

func (s CreateADConnectorDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryRequest) SetRegionId(v string) *CreateADConnectorDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainName(v string) *CreateADConnectorDirectoryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainUserName(v string) *CreateADConnectorDirectoryRequest {
	s.DomainUserName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainPassword(v string) *CreateADConnectorDirectoryRequest {
	s.DomainPassword = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDnsAddress(v []*string) *CreateADConnectorDirectoryRequest {
	s.DnsAddress = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetVSwitchId(v []*string) *CreateADConnectorDirectoryRequest {
	s.VSwitchId = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDirectoryName(v string) *CreateADConnectorDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetEnableAdminAccess(v bool) *CreateADConnectorDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDesktopAccessType(v string) *CreateADConnectorDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSubDomainDnsAddress(v []*string) *CreateADConnectorDirectoryRequest {
	s.SubDomainDnsAddress = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSubDomainName(v string) *CreateADConnectorDirectoryRequest {
	s.SubDomainName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetMfaEnabled(v bool) *CreateADConnectorDirectoryRequest {
	s.MfaEnabled = &v
	return s
}

type CreateADConnectorDirectoryResponse struct {
	DirectoryId   *string                                           `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TrustPassword *string                                           `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty" require:"true"`
	AdConnectors  []*CreateADConnectorDirectoryResponseAdConnectors `json:"AdConnectors,omitempty" xml:"AdConnectors,omitempty" require:"true" type:"Repeated"`
}

func (s CreateADConnectorDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponse) SetDirectoryId(v string) *CreateADConnectorDirectoryResponse {
	s.DirectoryId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetRequestId(v string) *CreateADConnectorDirectoryResponse {
	s.RequestId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetTrustPassword(v string) *CreateADConnectorDirectoryResponse {
	s.TrustPassword = &v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetAdConnectors(v []*CreateADConnectorDirectoryResponseAdConnectors) *CreateADConnectorDirectoryResponse {
	s.AdConnectors = v
	return s
}

type CreateADConnectorDirectoryResponseAdConnectors struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
}

func (s CreateADConnectorDirectoryResponseAdConnectors) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponseAdConnectors) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponseAdConnectors) SetAddress(v string) *CreateADConnectorDirectoryResponseAdConnectors {
	s.Address = &v
	return s
}

type DescribeBundlesRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	BundleId   []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
	BundleType *string   `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
}

func (s DescribeBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBundlesRequest) SetRegionId(v string) *DescribeBundlesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBundlesRequest) SetMaxResults(v int) *DescribeBundlesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeBundlesRequest) SetNextToken(v string) *DescribeBundlesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesRequest) SetBundleId(v []*string) *DescribeBundlesRequest {
	s.BundleId = v
	return s
}

func (s *DescribeBundlesRequest) SetBundleType(v string) *DescribeBundlesRequest {
	s.BundleType = &v
	return s
}

type DescribeBundlesResponse struct {
	NextToken *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Bundles   []*DescribeBundlesResponseBundles `json:"Bundles,omitempty" xml:"Bundles,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponse) SetNextToken(v string) *DescribeBundlesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesResponse) SetRequestId(v string) *DescribeBundlesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBundlesResponse) SetBundles(v []*DescribeBundlesResponseBundles) *DescribeBundlesResponse {
	s.Bundles = v
	return s
}

type DescribeBundlesResponseBundles struct {
	ImageId              *string                                             `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	BundleId             *string                                             `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	BundleType           *string                                             `json:"BundleType,omitempty" xml:"BundleType,omitempty" require:"true"`
	BundleName           *string                                             `json:"BundleName,omitempty" xml:"BundleName,omitempty" require:"true"`
	Description          *string                                             `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	DesktopType          *string                                             `json:"DesktopType,omitempty" xml:"DesktopType,omitempty" require:"true"`
	Disks                []*DescribeBundlesResponseBundlesDisks              `json:"Disks,omitempty" xml:"Disks,omitempty" require:"true" type:"Repeated"`
	DesktopTypeAttribute *DescribeBundlesResponseBundlesDesktopTypeAttribute `json:"DesktopTypeAttribute,omitempty" xml:"DesktopTypeAttribute,omitempty" require:"true" type:"Struct"`
}

func (s DescribeBundlesResponseBundles) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBundles) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBundles) SetImageId(v string) *DescribeBundlesResponseBundles {
	s.ImageId = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetBundleId(v string) *DescribeBundlesResponseBundles {
	s.BundleId = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetBundleType(v string) *DescribeBundlesResponseBundles {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetBundleName(v string) *DescribeBundlesResponseBundles {
	s.BundleName = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetDescription(v string) *DescribeBundlesResponseBundles {
	s.Description = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetDesktopType(v string) *DescribeBundlesResponseBundles {
	s.DesktopType = &v
	return s
}

func (s *DescribeBundlesResponseBundles) SetDisks(v []*DescribeBundlesResponseBundlesDisks) *DescribeBundlesResponseBundles {
	s.Disks = v
	return s
}

func (s *DescribeBundlesResponseBundles) SetDesktopTypeAttribute(v *DescribeBundlesResponseBundlesDesktopTypeAttribute) *DescribeBundlesResponseBundles {
	s.DesktopTypeAttribute = v
	return s
}

type DescribeBundlesResponseBundlesDisks struct {
	DiskSize *int    `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" require:"true"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty" require:"true"`
}

func (s DescribeBundlesResponseBundlesDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBundlesDisks) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBundlesDisks) SetDiskSize(v int) *DescribeBundlesResponseBundlesDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeBundlesResponseBundlesDisks) SetDiskType(v string) *DescribeBundlesResponseBundlesDisks {
	s.DiskType = &v
	return s
}

type DescribeBundlesResponseBundlesDesktopTypeAttribute struct {
	CpuCount   *int     `json:"CpuCount,omitempty" xml:"CpuCount,omitempty" require:"true"`
	MemorySize *int     `json:"MemorySize,omitempty" xml:"MemorySize,omitempty" require:"true"`
	GpuCount   *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty" require:"true"`
	GpuSpec    *string  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty" require:"true"`
}

func (s DescribeBundlesResponseBundlesDesktopTypeAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBundlesDesktopTypeAttribute) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBundlesDesktopTypeAttribute) SetCpuCount(v int) *DescribeBundlesResponseBundlesDesktopTypeAttribute {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBundlesDesktopTypeAttribute) SetMemorySize(v int) *DescribeBundlesResponseBundlesDesktopTypeAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeBundlesResponseBundlesDesktopTypeAttribute) SetGpuCount(v float32) *DescribeBundlesResponseBundlesDesktopTypeAttribute {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBundlesDesktopTypeAttribute) SetGpuSpec(v string) *DescribeBundlesResponseBundlesDesktopTypeAttribute {
	s.GpuSpec = &v
	return s
}

type DeleteImagesRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ImageId  []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesRequest) SetRegionId(v string) *DeleteImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImagesRequest) SetImageId(v []*string) *DeleteImagesRequest {
	s.ImageId = v
	return s
}

type DeleteImagesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponse) SetRequestId(v string) *DeleteImagesResponse {
	s.RequestId = &v
	return s
}

type CreateDesktopsRequest struct {
	RegionId      *string                     `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	GroupId       *string                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BundleId      *string                     `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	DesktopName   *string                     `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	UserName      *string                     `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VpcId         *string                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Amount        *int                        `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DirectoryId   *string                     `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId     []*string                   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true" type:"Repeated"`
	PolicyGroupId *string                     `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	ChargeType    *string                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period        *int                        `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit    *string                     `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay       *bool                       `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Tag           []*CreateDesktopsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	AutoRenew     *bool                       `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
}

func (s CreateDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequest) SetRegionId(v string) *CreateDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetGroupId(v string) *CreateDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetBundleId(v string) *CreateDesktopsRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopName(v string) *CreateDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *CreateDesktopsRequest) SetUserName(v string) *CreateDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *CreateDesktopsRequest) SetVpcId(v string) *CreateDesktopsRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDesktopsRequest) SetAmount(v int) *CreateDesktopsRequest {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsRequest) SetDirectoryId(v string) *CreateDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopsRequest) SetEndUserId(v []*string) *CreateDesktopsRequest {
	s.EndUserId = v
	return s
}

func (s *CreateDesktopsRequest) SetPolicyGroupId(v string) *CreateDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetChargeType(v string) *CreateDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriod(v int) *CreateDesktopsRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriodUnit(v string) *CreateDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopsRequest) SetAutoPay(v bool) *CreateDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopsRequest) SetTag(v []*CreateDesktopsRequestTag) *CreateDesktopsRequest {
	s.Tag = v
	return s
}

func (s *CreateDesktopsRequest) SetAutoRenew(v bool) *CreateDesktopsRequest {
	s.AutoRenew = &v
	return s
}

type CreateDesktopsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDesktopsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestTag) SetKey(v string) *CreateDesktopsRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDesktopsRequestTag) SetValue(v string) *CreateDesktopsRequestTag {
	s.Value = &v
	return s
}

type CreateDesktopsResponse struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OrderId   *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponse) SetRequestId(v string) *CreateDesktopsResponse {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopsResponse) SetOrderId(v string) *CreateDesktopsResponse {
	s.OrderId = &v
	return s
}

func (s *CreateDesktopsResponse) SetDesktopId(v []*string) *CreateDesktopsResponse {
	s.DesktopId = v
	return s
}

type DescribeImagesRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults  *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ImageType   *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageStatus *string   `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty"`
	ImageId     []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	GpuCategory *bool     `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) SetRegionId(v string) *DescribeImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImagesRequest) SetMaxResults(v int) *DescribeImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagesRequest) SetNextToken(v string) *DescribeImagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesRequest) SetImageType(v string) *DescribeImagesRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesRequest) SetImageStatus(v string) *DescribeImagesRequest {
	s.ImageStatus = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v []*string) *DescribeImagesRequest {
	s.ImageId = v
	return s
}

func (s *DescribeImagesRequest) SetGpuCategory(v bool) *DescribeImagesRequest {
	s.GpuCategory = &v
	return s
}

type DescribeImagesResponse struct {
	NextToken *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Images    []*DescribeImagesResponseImages `json:"Images,omitempty" xml:"Images,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetNextToken(v string) *DescribeImagesResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesResponse) SetRequestId(v string) *DescribeImagesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponse) SetImages(v []*DescribeImagesResponseImages) *DescribeImagesResponse {
	s.Images = v
	return s
}

type DescribeImagesResponseImages struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	ImageType    *string `json:"ImageType,omitempty" xml:"ImageType,omitempty" require:"true"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty" require:"true"`
	Size         *int    `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
	DataDiskSize *int    `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	GpuCategory  *bool   `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	OsType       *string `json:"OsType,omitempty" xml:"OsType,omitempty" require:"true"`
}

func (s DescribeImagesResponseImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseImages) SetCreationTime(v string) *DescribeImagesResponseImages {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseImages) SetImageId(v string) *DescribeImagesResponseImages {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseImages) SetImageType(v string) *DescribeImagesResponseImages {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesResponseImages) SetName(v string) *DescribeImagesResponseImages {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseImages) SetProgress(v string) *DescribeImagesResponseImages {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseImages) SetSize(v int) *DescribeImagesResponseImages {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseImages) SetDataDiskSize(v int) *DescribeImagesResponseImages {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeImagesResponseImages) SetGpuCategory(v bool) *DescribeImagesResponseImages {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesResponseImages) SetStatus(v string) *DescribeImagesResponseImages {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseImages) SetDescription(v string) *DescribeImagesResponseImages {
	s.Description = &v
	return s
}

func (s *DescribeImagesResponseImages) SetOsType(v string) *DescribeImagesResponseImages {
	s.OsType = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetDirectorySsoStatusWithOptions(request *GetDirectorySsoStatusRequest, runtime *util.RuntimeOptions) (_result *GetDirectorySsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDirectorySsoStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetDirectorySsoStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDirectorySsoStatus(request *GetDirectorySsoStatusRequest) (_result *GetDirectorySsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectorySsoStatusResponse{}
	_body, _err := client.GetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDirectorySsoStatusWithOptions(request *SetDirectorySsoStatusRequest, runtime *util.RuntimeOptions) (_result *SetDirectorySsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.DoRequest(tea.String("SetDirectorySsoStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDirectorySsoStatus(request *SetDirectorySsoStatusRequest) (_result *SetDirectorySsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.SetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpMetadataWithOptions(request *GetSpMetadataRequest, runtime *util.RuntimeOptions) (_result *GetSpMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.DoRequest(tea.String("GetSpMetadata"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpMetadata(request *GetSpMetadataRequest) (_result *GetSpMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.GetSpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetIdpMetadataWithOptions(request *SetIdpMetadataRequest, runtime *util.RuntimeOptions) (_result *SetIdpMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.DoRequest(tea.String("SetIdpMetadata"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetIdpMetadata(request *SetIdpMetadataRequest) (_result *SetIdpMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.SetIdpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebuildDesktopsWithOptions(request *RebuildDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebuildDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("RebuildDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebuildDesktops(request *RebuildDesktopsRequest) (_result *RebuildDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.RebuildDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBundleWithOptions(request *ModifyBundleRequest, runtime *util.RuntimeOptions) (_result *ModifyBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyBundleResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyBundle"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBundle(request *ModifyBundleRequest) (_result *ModifyBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBundleResponse{}
	_body, _err := client.ModifyBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockVirtualMFADeviceWithOptions(request *UnlockVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("UnlockVirtualMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockVirtualMFADevice(request *UnlockVirtualMFADeviceRequest) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.UnlockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVirtualMFADevicesWithOptions(request *DescribeVirtualMFADevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeVirtualMFADevices"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVirtualMFADevices(request *DescribeVirtualMFADevicesRequest) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.DescribeVirtualMFADevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockVirtualMFADeviceWithOptions(request *LockVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *LockVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("LockVirtualMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockVirtualMFADevice(request *LockVirtualMFADeviceRequest) (_result *LockVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.LockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteVirtualMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyADConnectorDirectoryWithOptions(request *ModifyADConnectorDirectoryRequest, runtime *util.RuntimeOptions) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyADConnectorDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyADConnectorDirectory(request *ModifyADConnectorDirectoryRequest) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.ModifyADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("UntagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("TagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopNameWithOptions(request *ModifyDesktopNameRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDesktopName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopName(request *ModifyDesktopNameRequest) (_result *ModifyDesktopNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.ModifyDesktopNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCommandWithOptions(request *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RunCommandResponse{}
	_body, _err := client.DoRequest(tea.String("RunCommand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInvocationWithOptions(request *StopInvocationRequest, runtime *util.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopInvocationResponse{}
	_body, _err := client.DoRequest(tea.String("StopInvocation"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInvocation(request *StopInvocationRequest) (_result *StopInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInvocationResponse{}
	_body, _err := client.StopInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInvocations"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeZones"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
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
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRegions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeClientEventsWithOptions(request *DescribeClientEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeClientEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeClientEvents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClientEvents(request *DescribeClientEventsRequest) (_result *DescribeClientEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.DescribeClientEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSnapshotWithOptions(request *ResetSnapshotRequest, runtime *util.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.DoRequest(tea.String("ResetSnapshot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSnapshot(request *ResetSnapshotRequest) (_result *ResetSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.ResetSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteSnapshot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.DoRequest(tea.String("CreateSnapshot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSnapshots"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewDesktopsWithOptions(request *RenewDesktopsRequest, runtime *util.RuntimeOptions) (_result *RenewDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("RenewDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewDesktops(request *RenewDesktopsRequest) (_result *RenewDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.RenewDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *util.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("StopDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDesktops(request *StopDesktopsRequest) (_result *StopDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDesktopsResponse{}
	_body, _err := client.StopDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *util.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("StartDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDesktops(request *StartDesktopsRequest) (_result *StartDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDesktopsResponse{}
	_body, _err := client.StartDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPolicyGroupWithOptions(request *ModifyPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyPolicyGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (_result *ModifyPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.ModifyPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopTypesWithOptions(request *DescribeDesktopTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktopTypes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopTypes(request *DescribeDesktopTypesRequest) (_result *DescribeDesktopTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.DescribeDesktopTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDirectories"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDirectoriesWithOptions(request *DeleteDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDirectories"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDirectories(request *DeleteDirectoriesRequest) (_result *DeleteDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.DeleteDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDirectoryUsersWithOptions(request *ListDirectoryUsersRequest, runtime *util.RuntimeOptions) (_result *ListDirectoryUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.DoRequest(tea.String("ListDirectoryUsers"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDirectoryUsers(request *ListDirectoryUsersRequest) (_result *ListDirectoryUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.ListDirectoryUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoRequest(tea.String("CreateImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRAMDirectoryWithOptions(request *CreateRAMDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateRAMDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("CreateRAMDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRAMDirectory(request *CreateRAMDirectoryRequest) (_result *CreateRAMDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.CreateRAMDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePolicyGroupsWithOptions(request *DeletePolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePolicyGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePolicyGroups(request *DeletePolicyGroupsRequest) (_result *DeletePolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.DeletePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePolicyGroupsWithOptions(request *DescribePolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePolicyGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePolicyGroups(request *DescribePolicyGroupsRequest) (_result *DescribePolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.DescribePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDesktopsWithOptions(request *DeleteDesktopsRequest, runtime *util.RuntimeOptions) (_result *DeleteDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDesktops(request *DeleteDesktopsRequest) (_result *DeleteDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.DeleteDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageAttributeWithOptions(request *ModifyImageAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyImageAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (_result *ModifyImageAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.ModifyImageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEntitlementWithOptions(request *ModifyEntitlementRequest, runtime *util.RuntimeOptions) (_result *ModifyEntitlementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyEntitlement"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEntitlement(request *ModifyEntitlementRequest) (_result *ModifyEntitlementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.ModifyEntitlementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBundlesWithOptions(request *DeleteBundlesRequest, runtime *util.RuntimeOptions) (_result *DeleteBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteBundles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBundles(request *DeleteBundlesRequest) (_result *DeleteBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.DeleteBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopsWithOptions(request *DescribeDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktops(request *DescribeDesktopsRequest) (_result *DescribeDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DescribeDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("RebootDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootDesktops(request *RebootDesktopsRequest) (_result *RebootDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.RebootDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBundleWithOptions(request *CreateBundleRequest, runtime *util.RuntimeOptions) (_result *CreateBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateBundleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateBundle"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBundle(request *CreateBundleRequest) (_result *CreateBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBundleResponse{}
	_body, _err := client.CreateBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopsPolicyGroupWithOptions(request *ModifyDesktopsPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDesktopsPolicyGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopsPolicyGroup(request *ModifyDesktopsPolicyGroupRequest) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.ModifyDesktopsPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePolicyGroupWithOptions(request *CreatePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePolicyGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (_result *CreatePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CreatePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateADConnectorDirectoryWithOptions(request *CreateADConnectorDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateADConnectorDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.DoRequest(tea.String("CreateADConnectorDirectory"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateADConnectorDirectory(request *CreateADConnectorDirectoryRequest) (_result *CreateADConnectorDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.CreateADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBundlesWithOptions(request *DescribeBundlesRequest, runtime *util.RuntimeOptions) (_result *DescribeBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBundles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBundles(request *DescribeBundlesRequest) (_result *DescribeBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.DescribeBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImagesWithOptions(request *DeleteImagesRequest, runtime *util.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteImages"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImages(request *DeleteImagesRequest) (_result *DeleteImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DeleteImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopsWithOptions(request *CreateDesktopsRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktops(request *CreateDesktopsRequest) (_result *CreateDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.CreateDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeImages"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
