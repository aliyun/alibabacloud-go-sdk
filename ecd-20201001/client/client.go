// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeDirectoriesRequest struct {
	DirectoryType *string   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty" require:"true"`
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
	CustomSecurityGroupId *string                                               `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty" require:"true"`
	DnsUserName           *string                                               `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty" require:"true"`
	EnableInternetAccess  *bool                                                 `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty" require:"true"`
	TrustPassword         *string                                               `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty" require:"true"`
	DomainName            *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty" require:"true"`
	DomainUserName        *string                                               `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty" require:"true"`
	DomainPassword        *string                                               `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty" require:"true"`
	ADConnectors          []*DescribeDirectoriesResponseDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" require:"true" type:"Repeated"`
	DnsAddress            []*string                                             `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" require:"true" type:"Repeated"`
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

type DescribeDirectoriesResponseDirectoriesADConnectors struct {
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty" require:"true"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty" require:"true"`
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

type DeleteDirectoriesRequest struct {
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
}

func (s DeleteDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesRequest) GoString() string {
	return s.String()
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

type DescribeDesktopsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DirectoryId   *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupId       *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	DesktopStatus *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	MaxResults    *int      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserName      *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	DesktopName   *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
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

func (s *DescribeDesktopsRequest) SetDirectoryId(v string) *DescribeDesktopsRequest {
	s.DirectoryId = &v
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

func (s *DescribeDesktopsRequest) SetDesktopId(v []*string) *DescribeDesktopsRequest {
	s.DesktopId = v
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
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" require:"true"`
	Ticket     *string `json:"Ticket,omitempty" xml:"Ticket,omitempty" require:"true"`
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

func (s *GetConnectionTicketResponse) SetTaskId(v string) *GetConnectionTicketResponse {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponse) SetTaskStatus(v string) *GetConnectionTicketResponse {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponse) SetTicket(v string) *GetConnectionTicketResponse {
	s.Ticket = &v
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

func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDirectories"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-10-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_body, _err := client.DoRequest(tea.String("DeleteDirectories"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-10-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeDesktopsWithOptions(request *DescribeDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-10-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_body, _err := client.DoRequest(tea.String("RebootDesktops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-10-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.DoRequest(tea.String("GetConnectionTicket"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-10-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
