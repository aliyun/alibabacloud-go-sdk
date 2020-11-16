// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ModifyPolicyGroupRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Clipboard     *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive    *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect   *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark     *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
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

type PayOrderCallbackRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PayOrderCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s PayOrderCallbackRequest) GoString() string {
	return s.String()
}

func (s *PayOrderCallbackRequest) SetData(v string) *PayOrderCallbackRequest {
	s.Data = &v
	return s
}

type PayOrderCallbackResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s PayOrderCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s PayOrderCallbackResponse) GoString() string {
	return s.String()
}

func (s *PayOrderCallbackResponse) SetRequestId(v string) *PayOrderCallbackResponse {
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
	DesktopTypeId      *string `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty" require:"true"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty" require:"true"`
	CpuCount           *string `json:"CpuCount,omitempty" xml:"CpuCount,omitempty" require:"true"`
	GPUCount           *string `json:"GPUCount,omitempty" xml:"GPUCount,omitempty" require:"true"`
	MemorySize         *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty" require:"true"`
	SystemDiskSize     *string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskSize       *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
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

func (s *DescribeDesktopTypesResponseDesktopTypes) SetGPUCount(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.GPUCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseDesktopTypes) SetMemorySize(v string) *DescribeDesktopTypesResponseDesktopTypes {
	s.MemorySize = &v
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
	TrustPassword         *string                                               `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty" require:"true"`
	DomainName            *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName        *string                                               `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword        *string                                               `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	ADConnectors          []*DescribeDirectoriesResponseDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" require:"true" type:"Repeated"`
	DnsAddress            []*string                                             `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" require:"true" type:"Repeated"`
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

func (s *DescribeDirectoriesResponseDirectories) SetADConnectors(v []*DescribeDirectoriesResponseDirectoriesADConnectors) *DescribeDirectoriesResponseDirectories {
	s.ADConnectors = v
	return s
}

func (s *DescribeDirectoriesResponseDirectories) SetDnsAddress(v []*string) *DescribeDirectoriesResponseDirectories {
	s.DnsAddress = v
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
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponse) SetNextToken(v string) *DeleteDirectoriesResponse {
	s.NextToken = &v
	return s
}

func (s *DeleteDirectoriesResponse) SetRequestId(v string) *DeleteDirectoriesResponse {
	s.RequestId = &v
	return s
}

type ListDirectoryUsersRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
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
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
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
	VSwitchId            []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
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
	PolicyGroupId   *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	PolicyGroupType *string `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty" require:"true"`
	Clipboard       *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty" require:"true"`
	LocalDrive      *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty" require:"true"`
	UsbRedirect     *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty" require:"true"`
	Watermark       *string `json:"Watermark,omitempty" xml:"Watermark,omitempty" require:"true"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
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

type DoLogicalDeleteResourceRequest struct {
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DoLogicalDeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DoLogicalDeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DoLogicalDeleteResourceRequest) SetInvoker(v string) *DoLogicalDeleteResourceRequest {
	s.Invoker = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetPk(v string) *DoLogicalDeleteResourceRequest {
	s.Pk = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetBid(v string) *DoLogicalDeleteResourceRequest {
	s.Bid = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetHid(v int64) *DoLogicalDeleteResourceRequest {
	s.Hid = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetCountry(v string) *DoLogicalDeleteResourceRequest {
	s.Country = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetTaskIdentifier(v string) *DoLogicalDeleteResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetTaskExtraData(v string) *DoLogicalDeleteResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetGmtWakeup(v string) *DoLogicalDeleteResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetRegionId(v string) *DoLogicalDeleteResourceRequest {
	s.RegionId = &v
	return s
}

type DoLogicalDeleteResourceResponse struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty" require:"true"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty" require:"true"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty" require:"true"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty" require:"true"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty" require:"true"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty" require:"true"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty" require:"true"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty" require:"true"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty" require:"true"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DoLogicalDeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DoLogicalDeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DoLogicalDeleteResourceResponse) SetBid(v string) *DoLogicalDeleteResourceResponse {
	s.Bid = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetCountry(v string) *DoLogicalDeleteResourceResponse {
	s.Country = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetHid(v int64) *DoLogicalDeleteResourceResponse {
	s.Hid = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetInterrupt(v bool) *DoLogicalDeleteResourceResponse {
	s.Interrupt = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetInvoker(v string) *DoLogicalDeleteResourceResponse {
	s.Invoker = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetMessage(v string) *DoLogicalDeleteResourceResponse {
	s.Message = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetPk(v string) *DoLogicalDeleteResourceResponse {
	s.Pk = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetSuccess(v bool) *DoLogicalDeleteResourceResponse {
	s.Success = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetTaskExtraData(v string) *DoLogicalDeleteResourceResponse {
	s.TaskExtraData = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetTaskIdentifier(v string) *DoLogicalDeleteResourceResponse {
	s.TaskIdentifier = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetGmtWakeup(v string) *DoLogicalDeleteResourceResponse {
	s.GmtWakeup = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetRequestId(v string) *DoLogicalDeleteResourceResponse {
	s.RequestId = &v
	return s
}

type ModifyEntitlementRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
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
	DirectoryId        *string                                  `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" require:"true"`
	CreationTime       *string                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	DesktopId          *string                                  `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	DesktopStatus      *string                                  `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty" require:"true"`
	DesktopName        *string                                  `json:"DesktopName,omitempty" xml:"DesktopName,omitempty" require:"true"`
	ImageId            *string                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	DesktopType        *string                                  `json:"DesktopType,omitempty" xml:"DesktopType,omitempty" require:"true"`
	SystemDiskCategory *string                                  `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty" require:"true"`
	SystemDiskSize     *int                                     `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty" require:"true"`
	DataDiskCategory   *string                                  `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty" require:"true"`
	DataDiskSize       *string                                  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty" require:"true"`
	ConnectionStatus   *string                                  `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty" require:"true"`
	PolicyGroupId      *string                                  `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	Cpu                *int                                     `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory             *int64                                   `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	NetworkInterfaceId *int64                                   `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" require:"true"`
	ExpiredTime        *string                                  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty" require:"true"`
	ChargeType         *string                                  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty" require:"true"`
	Disks              []*DescribeDesktopsResponseDesktopsDisks `json:"Disks,omitempty" xml:"Disks,omitempty" require:"true" type:"Repeated"`
	EndUserIds         []*string                                `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" require:"true" type:"Repeated"`
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

func (s *DescribeDesktopsResponseDesktops) SetCreationTime(v string) *DescribeDesktopsResponseDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeDesktopsResponseDesktops) SetDesktopId(v string) *DescribeDesktopsResponseDesktops {
	s.DesktopId = &v
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

func (s *DescribeDesktopsResponseDesktops) SetConnectionStatus(v string) *DescribeDesktopsResponseDesktops {
	s.ConnectionStatus = &v
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

func (s *DescribeDesktopsResponseDesktops) SetNetworkInterfaceId(v int64) *DescribeDesktopsResponseDesktops {
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

func (s *DescribeDesktopsResponseDesktops) SetDisks(v []*DescribeDesktopsResponseDesktopsDisks) *DescribeDesktopsResponseDesktops {
	s.Disks = v
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
	Message   *int    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
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

func (s *ModifyDesktopsPolicyGroupResponseModifyResults) SetMessage(v int) *ModifyDesktopsPolicyGroupResponseModifyResults {
	s.Message = &v
	return s
}

type CreatePolicyGroupRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type DoPhysicalDeleteResourceRequest struct {
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DoPhysicalDeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DoPhysicalDeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DoPhysicalDeleteResourceRequest) SetInvoker(v string) *DoPhysicalDeleteResourceRequest {
	s.Invoker = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetPk(v string) *DoPhysicalDeleteResourceRequest {
	s.Pk = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetBid(v string) *DoPhysicalDeleteResourceRequest {
	s.Bid = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetHid(v int64) *DoPhysicalDeleteResourceRequest {
	s.Hid = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetCountry(v string) *DoPhysicalDeleteResourceRequest {
	s.Country = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetTaskIdentifier(v string) *DoPhysicalDeleteResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetTaskExtraData(v string) *DoPhysicalDeleteResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetGmtWakeup(v string) *DoPhysicalDeleteResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetRegionId(v string) *DoPhysicalDeleteResourceRequest {
	s.RegionId = &v
	return s
}

type DoPhysicalDeleteResourceResponse struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty" require:"true"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty" require:"true"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty" require:"true"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty" require:"true"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty" require:"true"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty" require:"true"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty" require:"true"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty" require:"true"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty" require:"true"`
}

func (s DoPhysicalDeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DoPhysicalDeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DoPhysicalDeleteResourceResponse) SetBid(v string) *DoPhysicalDeleteResourceResponse {
	s.Bid = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetCountry(v string) *DoPhysicalDeleteResourceResponse {
	s.Country = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetHid(v int64) *DoPhysicalDeleteResourceResponse {
	s.Hid = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetInterrupt(v bool) *DoPhysicalDeleteResourceResponse {
	s.Interrupt = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetInvoker(v string) *DoPhysicalDeleteResourceResponse {
	s.Invoker = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetMessage(v string) *DoPhysicalDeleteResourceResponse {
	s.Message = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetRequestId(v string) *DoPhysicalDeleteResourceResponse {
	s.RequestId = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetPk(v string) *DoPhysicalDeleteResourceResponse {
	s.Pk = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetSuccess(v bool) *DoPhysicalDeleteResourceResponse {
	s.Success = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetTaskExtraData(v string) *DoPhysicalDeleteResourceResponse {
	s.TaskExtraData = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetTaskIdentifier(v string) *DoPhysicalDeleteResourceResponse {
	s.TaskIdentifier = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetGmtWakeup(v string) *DoPhysicalDeleteResourceResponse {
	s.GmtWakeup = &v
	return s
}

type CreateADConnectorDirectoryRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DomainName     *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	DnsAddress     []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	VSwitchId      []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	DirectoryName  *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
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

type GetConnectionTicketRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password   *string `json:"Password,omitempty" xml:"Password,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetRegionId(v string) *GetConnectionTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetInstanceId(v string) *GetConnectionTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUserName(v string) *GetConnectionTicketRequest {
	s.UserName = &v
	return s
}

func (s *GetConnectionTicketRequest) SetPassword(v string) *GetConnectionTicketRequest {
	s.Password = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetDesktopId(v string) *GetConnectionTicketRequest {
	s.DesktopId = &v
	return s
}

type GetConnectionTicketResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Ticket     *string `json:"Ticket,omitempty" xml:"Ticket,omitempty" require:"true"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" require:"true"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetRequestId(v string) *GetConnectionTicketResponse {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponse) SetTicket(v string) *GetConnectionTicketResponse {
	s.Ticket = &v
	return s
}

func (s *GetConnectionTicketResponse) SetTaskId(v string) *GetConnectionTicketResponse {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponse) SetTaskStatus(v string) *GetConnectionTicketResponse {
	s.TaskStatus = &v
	return s
}

type ModifyDesktopPolicysRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId   []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true" type:"Repeated"`
	Clipboard   *string   `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive  *string   `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect *string   `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string   `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s ModifyDesktopPolicysRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopPolicysRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopPolicysRequest) SetRegionId(v string) *ModifyDesktopPolicysRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetDesktopId(v []*string) *ModifyDesktopPolicysRequest {
	s.DesktopId = v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetClipboard(v string) *ModifyDesktopPolicysRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetLocalDrive(v string) *ModifyDesktopPolicysRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetUsbRedirect(v string) *ModifyDesktopPolicysRequest {
	s.UsbRedirect = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetWatermark(v string) *ModifyDesktopPolicysRequest {
	s.Watermark = &v
	return s
}

type ModifyDesktopPolicysResponse struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Results   []*ModifyDesktopPolicysResponseResults `json:"Results,omitempty" xml:"Results,omitempty" require:"true" type:"Repeated"`
}

func (s ModifyDesktopPolicysResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopPolicysResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopPolicysResponse) SetRequestId(v string) *ModifyDesktopPolicysResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopPolicysResponse) SetResults(v []*ModifyDesktopPolicysResponseResults) *ModifyDesktopPolicysResponse {
	s.Results = v
	return s
}

type ModifyDesktopPolicysResponseResults struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s ModifyDesktopPolicysResponseResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopPolicysResponseResults) GoString() string {
	return s.String()
}

func (s *ModifyDesktopPolicysResponseResults) SetDesktopId(v string) *ModifyDesktopPolicysResponseResults {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopPolicysResponseResults) SetSuccess(v string) *ModifyDesktopPolicysResponseResults {
	s.Success = &v
	return s
}

func (s *ModifyDesktopPolicysResponseResults) SetCode(v string) *ModifyDesktopPolicysResponseResults {
	s.Code = &v
	return s
}

func (s *ModifyDesktopPolicysResponseResults) SetMessage(v string) *ModifyDesktopPolicysResponseResults {
	s.Message = &v
	return s
}

type DescribeBundlesRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	MaxResults *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserName   *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Category   *string   `json:"Category,omitempty" xml:"Category,omitempty"`
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

func (s *DescribeBundlesRequest) SetUserName(v string) *DescribeBundlesRequest {
	s.UserName = &v
	return s
}

func (s *DescribeBundlesRequest) SetCategory(v string) *DescribeBundlesRequest {
	s.Category = &v
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
	ImageId     *string                                `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	BundleId    *string                                `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	BundleType  *string                                `json:"BundleType,omitempty" xml:"BundleType,omitempty" require:"true"`
	BundleName  *string                                `json:"BundleName,omitempty" xml:"BundleName,omitempty" require:"true"`
	Description *string                                `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	DesktopType *string                                `json:"DesktopType,omitempty" xml:"DesktopType,omitempty" require:"true"`
	Disks       []*DescribeBundlesResponseBundlesDisks `json:"Disks,omitempty" xml:"Disks,omitempty" require:"true" type:"Repeated"`
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

type DoCheckResourceRequest struct {
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DoCheckResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DoCheckResourceRequest) GoString() string {
	return s.String()
}

func (s *DoCheckResourceRequest) SetInvoker(v string) *DoCheckResourceRequest {
	s.Invoker = &v
	return s
}

func (s *DoCheckResourceRequest) SetPk(v string) *DoCheckResourceRequest {
	s.Pk = &v
	return s
}

func (s *DoCheckResourceRequest) SetBid(v string) *DoCheckResourceRequest {
	s.Bid = &v
	return s
}

func (s *DoCheckResourceRequest) SetHid(v int64) *DoCheckResourceRequest {
	s.Hid = &v
	return s
}

func (s *DoCheckResourceRequest) SetCountry(v string) *DoCheckResourceRequest {
	s.Country = &v
	return s
}

func (s *DoCheckResourceRequest) SetTaskIdentifier(v string) *DoCheckResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *DoCheckResourceRequest) SetTaskExtraData(v string) *DoCheckResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *DoCheckResourceRequest) SetGmtWakeup(v string) *DoCheckResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *DoCheckResourceRequest) SetRegionId(v string) *DoCheckResourceRequest {
	s.RegionId = &v
	return s
}

type DoCheckResourceResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty" require:"true"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty" require:"true"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty" require:"true"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty" require:"true"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty" require:"true"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty" require:"true"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty" require:"true"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty" require:"true"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty" require:"true"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty" require:"true"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty" require:"true"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty" require:"true"`
}

func (s DoCheckResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DoCheckResourceResponse) GoString() string {
	return s.String()
}

func (s *DoCheckResourceResponse) SetRequestId(v string) *DoCheckResourceResponse {
	s.RequestId = &v
	return s
}

func (s *DoCheckResourceResponse) SetInterrupt(v bool) *DoCheckResourceResponse {
	s.Interrupt = &v
	return s
}

func (s *DoCheckResourceResponse) SetInvoker(v string) *DoCheckResourceResponse {
	s.Invoker = &v
	return s
}

func (s *DoCheckResourceResponse) SetPk(v string) *DoCheckResourceResponse {
	s.Pk = &v
	return s
}

func (s *DoCheckResourceResponse) SetBid(v string) *DoCheckResourceResponse {
	s.Bid = &v
	return s
}

func (s *DoCheckResourceResponse) SetHid(v int64) *DoCheckResourceResponse {
	s.Hid = &v
	return s
}

func (s *DoCheckResourceResponse) SetCountry(v string) *DoCheckResourceResponse {
	s.Country = &v
	return s
}

func (s *DoCheckResourceResponse) SetTaskIdentifier(v string) *DoCheckResourceResponse {
	s.TaskIdentifier = &v
	return s
}

func (s *DoCheckResourceResponse) SetTaskExtraData(v string) *DoCheckResourceResponse {
	s.TaskExtraData = &v
	return s
}

func (s *DoCheckResourceResponse) SetGmtWakeup(v string) *DoCheckResourceResponse {
	s.GmtWakeup = &v
	return s
}

func (s *DoCheckResourceResponse) SetSuccess(v bool) *DoCheckResourceResponse {
	s.Success = &v
	return s
}

func (s *DoCheckResourceResponse) SetMessage(v string) *DoCheckResourceResponse {
	s.Message = &v
	return s
}

func (s *DoCheckResourceResponse) SetLevel(v int64) *DoCheckResourceResponse {
	s.Level = &v
	return s
}

func (s *DoCheckResourceResponse) SetUrl(v string) *DoCheckResourceResponse {
	s.Url = &v
	return s
}

func (s *DoCheckResourceResponse) SetPrompt(v string) *DoCheckResourceResponse {
	s.Prompt = &v
	return s
}

type DescribeDesktopPolicysRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DesktopId  []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s DescribeDesktopPolicysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopPolicysRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopPolicysRequest) SetRegionId(v string) *DescribeDesktopPolicysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopPolicysRequest) SetDesktopId(v []*string) *DescribeDesktopPolicysRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopPolicysRequest) SetNextToken(v string) *DescribeDesktopPolicysRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopPolicysRequest) SetMaxResults(v int) *DescribeDesktopPolicysRequest {
	s.MaxResults = &v
	return s
}

type DescribeDesktopPolicysResponse struct {
	NextToken              *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	RequestId              *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DescribeDesktopPolicys []*DescribeDesktopPolicysResponseDescribeDesktopPolicys `json:"DescribeDesktopPolicys,omitempty" xml:"DescribeDesktopPolicys,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDesktopPolicysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopPolicysResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopPolicysResponse) SetNextToken(v string) *DescribeDesktopPolicysResponse {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopPolicysResponse) SetRequestId(v string) *DescribeDesktopPolicysResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopPolicysResponse) SetDescribeDesktopPolicys(v []*DescribeDesktopPolicysResponseDescribeDesktopPolicys) *DescribeDesktopPolicysResponse {
	s.DescribeDesktopPolicys = v
	return s
}

type DescribeDesktopPolicysResponseDescribeDesktopPolicys struct {
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty" require:"true"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty" require:"true"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty" require:"true"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty" require:"true"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" require:"true"`
}

func (s DescribeDesktopPolicysResponseDescribeDesktopPolicys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopPolicysResponseDescribeDesktopPolicys) GoString() string {
	return s.String()
}

func (s *DescribeDesktopPolicysResponseDescribeDesktopPolicys) SetClipboard(v string) *DescribeDesktopPolicysResponseDescribeDesktopPolicys {
	s.Clipboard = &v
	return s
}

func (s *DescribeDesktopPolicysResponseDescribeDesktopPolicys) SetLocalDrive(v string) *DescribeDesktopPolicysResponseDescribeDesktopPolicys {
	s.LocalDrive = &v
	return s
}

func (s *DescribeDesktopPolicysResponseDescribeDesktopPolicys) SetUsbRedirect(v string) *DescribeDesktopPolicysResponseDescribeDesktopPolicys {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeDesktopPolicysResponseDescribeDesktopPolicys) SetWatermark(v string) *DescribeDesktopPolicysResponseDescribeDesktopPolicys {
	s.Watermark = &v
	return s
}

func (s *DescribeDesktopPolicysResponseDescribeDesktopPolicys) SetDesktopId(v string) *DescribeDesktopPolicysResponseDescribeDesktopPolicys {
	s.DesktopId = &v
	return s
}

type CreateDesktopsRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	GroupId        *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BundleId       *string   `json:"BundleId,omitempty" xml:"BundleId,omitempty" require:"true"`
	SystemDiskSize *int      `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	DataDiskSize   *int      `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopName    *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	UserName       *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VpcId          *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Amount         *int      `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DirectoryId    *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId      []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" require:"true" type:"Repeated"`
	PolicyGroupId  *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" require:"true"`
	ChargeType     *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period         *int      `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit     *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay        *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
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

func (s *CreateDesktopsRequest) SetSystemDiskSize(v int) *CreateDesktopsRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateDesktopsRequest) SetDataDiskSize(v int) *CreateDesktopsRequest {
	s.DataDiskSize = &v
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
	Progress     *int    `json:"Progress,omitempty" xml:"Progress,omitempty" require:"true"`
	Size         *int    `json:"Size,omitempty" xml:"Size,omitempty" require:"true"`
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

func (s *DescribeImagesResponseImages) SetProgress(v int) *DescribeImagesResponseImages {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseImages) SetSize(v int) *DescribeImagesResponseImages {
	s.Size = &v
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

func (client *Client) PayOrderCallbackWithOptions(request *PayOrderCallbackRequest, runtime *util.RuntimeOptions) (_result *PayOrderCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PayOrderCallbackResponse{}
	_body, _err := client.DoRequest(tea.String("PayOrderCallback"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PayOrderCallback(request *PayOrderCallbackRequest) (_result *PayOrderCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PayOrderCallbackResponse{}
	_body, _err := client.PayOrderCallbackWithOptions(request, runtime)
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

func (client *Client) DoLogicalDeleteResourceWithOptions(request *DoLogicalDeleteResourceRequest, runtime *util.RuntimeOptions) (_result *DoLogicalDeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DoLogicalDeleteResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DoLogicalDeleteResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoLogicalDeleteResource(request *DoLogicalDeleteResourceRequest) (_result *DoLogicalDeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoLogicalDeleteResourceResponse{}
	_body, _err := client.DoLogicalDeleteResourceWithOptions(request, runtime)
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

func (client *Client) DoPhysicalDeleteResourceWithOptions(request *DoPhysicalDeleteResourceRequest, runtime *util.RuntimeOptions) (_result *DoPhysicalDeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DoPhysicalDeleteResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DoPhysicalDeleteResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoPhysicalDeleteResource(request *DoPhysicalDeleteResourceRequest) (_result *DoPhysicalDeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoPhysicalDeleteResourceResponse{}
	_body, _err := client.DoPhysicalDeleteResourceWithOptions(request, runtime)
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

func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.DoRequest(tea.String("GetConnectionTicket"), tea.String("HTTP"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopPolicysWithOptions(request *ModifyDesktopPolicysRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopPolicysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDesktopPolicysResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDesktopPolicys"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopPolicys(request *ModifyDesktopPolicysRequest) (_result *ModifyDesktopPolicysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopPolicysResponse{}
	_body, _err := client.ModifyDesktopPolicysWithOptions(request, runtime)
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

func (client *Client) DoCheckResourceWithOptions(request *DoCheckResourceRequest, runtime *util.RuntimeOptions) (_result *DoCheckResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DoCheckResourceResponse{}
	_body, _err := client.DoRequest(tea.String("DoCheckResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoCheckResource(request *DoCheckResourceRequest) (_result *DoCheckResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoCheckResourceResponse{}
	_body, _err := client.DoCheckResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopPolicysWithOptions(request *DescribeDesktopPolicysRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopPolicysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopPolicysResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktopPolicys"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopPolicys(request *DescribeDesktopPolicysRequest) (_result *DescribeDesktopPolicysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopPolicysResponse{}
	_body, _err := client.DescribeDesktopPolicysWithOptions(request, runtime)
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
