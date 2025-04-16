// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CalMcpToolRequest struct {
	Args           *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Authorization  *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Server         *string `json:"Server,omitempty" xml:"Server,omitempty"`
	SessionId      *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Tool           *string `json:"Tool,omitempty" xml:"Tool,omitempty"`
}

func (s CalMcpToolRequest) String() string {
	return tea.Prettify(s)
}

func (s CalMcpToolRequest) GoString() string {
	return s.String()
}

func (s *CalMcpToolRequest) SetArgs(v string) *CalMcpToolRequest {
	s.Args = &v
	return s
}

func (s *CalMcpToolRequest) SetAuthorization(v string) *CalMcpToolRequest {
	s.Authorization = &v
	return s
}

func (s *CalMcpToolRequest) SetExternalUserId(v string) *CalMcpToolRequest {
	s.ExternalUserId = &v
	return s
}

func (s *CalMcpToolRequest) SetName(v string) *CalMcpToolRequest {
	s.Name = &v
	return s
}

func (s *CalMcpToolRequest) SetServer(v string) *CalMcpToolRequest {
	s.Server = &v
	return s
}

func (s *CalMcpToolRequest) SetSessionId(v string) *CalMcpToolRequest {
	s.SessionId = &v
	return s
}

func (s *CalMcpToolRequest) SetTool(v string) *CalMcpToolRequest {
	s.Tool = &v
	return s
}

type CalMcpToolResponseBody struct {
	Code           *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CalMcpToolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CalMcpToolResponseBody) GoString() string {
	return s.String()
}

func (s *CalMcpToolResponseBody) SetCode(v string) *CalMcpToolResponseBody {
	s.Code = &v
	return s
}

func (s *CalMcpToolResponseBody) SetData(v interface{}) *CalMcpToolResponseBody {
	s.Data = v
	return s
}

func (s *CalMcpToolResponseBody) SetHttpStatusCode(v int32) *CalMcpToolResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CalMcpToolResponseBody) SetMessage(v string) *CalMcpToolResponseBody {
	s.Message = &v
	return s
}

func (s *CalMcpToolResponseBody) SetRequestId(v string) *CalMcpToolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CalMcpToolResponseBody) SetSuccess(v bool) *CalMcpToolResponseBody {
	s.Success = &v
	return s
}

type CalMcpToolResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CalMcpToolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CalMcpToolResponse) String() string {
	return tea.Prettify(s)
}

func (s CalMcpToolResponse) GoString() string {
	return s.String()
}

func (s *CalMcpToolResponse) SetHeaders(v map[string]*string) *CalMcpToolResponse {
	s.Headers = v
	return s
}

func (s *CalMcpToolResponse) SetStatusCode(v int32) *CalMcpToolResponse {
	s.StatusCode = &v
	return s
}

func (s *CalMcpToolResponse) SetBody(v *CalMcpToolResponseBody) *CalMcpToolResponse {
	s.Body = v
	return s
}

type CreateMcpSessionRequest struct {
	Authorization  *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	SessionId      *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s CreateMcpSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMcpSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpSessionRequest) SetAuthorization(v string) *CreateMcpSessionRequest {
	s.Authorization = &v
	return s
}

func (s *CreateMcpSessionRequest) SetExternalUserId(v string) *CreateMcpSessionRequest {
	s.ExternalUserId = &v
	return s
}

func (s *CreateMcpSessionRequest) SetSessionId(v string) *CreateMcpSessionRequest {
	s.SessionId = &v
	return s
}

type CreateMcpSessionResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateMcpSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcpSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMcpSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpSessionResponseBody) SetCode(v string) *CreateMcpSessionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMcpSessionResponseBody) SetData(v *CreateMcpSessionResponseBodyData) *CreateMcpSessionResponseBody {
	s.Data = v
	return s
}

func (s *CreateMcpSessionResponseBody) SetHttpStatusCode(v int32) *CreateMcpSessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMcpSessionResponseBody) SetMessage(v string) *CreateMcpSessionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMcpSessionResponseBody) SetRequestId(v string) *CreateMcpSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpSessionResponseBody) SetSuccess(v bool) *CreateMcpSessionResponseBody {
	s.Success = &v
	return s
}

type CreateMcpSessionResponseBodyData struct {
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	ErrMsg        *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ResourceId    *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceUrl   *string `json:"ResourceUrl,omitempty" xml:"ResourceUrl,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcpSessionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateMcpSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMcpSessionResponseBodyData) SetAppInstanceId(v string) *CreateMcpSessionResponseBodyData {
	s.AppInstanceId = &v
	return s
}

func (s *CreateMcpSessionResponseBodyData) SetErrMsg(v string) *CreateMcpSessionResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *CreateMcpSessionResponseBodyData) SetResourceId(v string) *CreateMcpSessionResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *CreateMcpSessionResponseBodyData) SetResourceUrl(v string) *CreateMcpSessionResponseBodyData {
	s.ResourceUrl = &v
	return s
}

func (s *CreateMcpSessionResponseBodyData) SetSessionId(v string) *CreateMcpSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateMcpSessionResponseBodyData) SetSuccess(v bool) *CreateMcpSessionResponseBodyData {
	s.Success = &v
	return s
}

type CreateMcpSessionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcpSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcpSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMcpSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateMcpSessionResponse) SetHeaders(v map[string]*string) *CreateMcpSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateMcpSessionResponse) SetStatusCode(v int32) *CreateMcpSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcpSessionResponse) SetBody(v *CreateMcpSessionResponseBody) *CreateMcpSessionResponse {
	s.Body = v
	return s
}

type DescribeDesktopsRequest struct {
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	DesktopName   *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DirectoryId   *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GroupId       *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MaxResults    *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId  *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsRequest) SetDesktopId(v []*string) *DescribeDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopName(v string) *DescribeDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopStatus(v string) *DescribeDesktopsRequest {
	s.DesktopStatus = &v
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

func (s *DescribeDesktopsRequest) SetMaxResults(v int32) *DescribeDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsRequest) SetNextToken(v string) *DescribeDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOfficeSiteId(v string) *DescribeDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetRegionId(v string) *DescribeDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetUserName(v string) *DescribeDesktopsRequest {
	s.UserName = &v
	return s
}

type DescribeDesktopsResponseBody struct {
	Desktops  []*DescribeDesktopsResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
	NextToken *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBody) SetDesktops(v []*DescribeDesktopsResponseBodyDesktops) *DescribeDesktopsResponseBody {
	s.Desktops = v
	return s
}

func (s *DescribeDesktopsResponseBody) SetNextToken(v string) *DescribeDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetRequestId(v string) *DescribeDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDesktopsResponseBodyDesktops struct {
	ConnectionStatus   *string                                      `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	Cpu                *int32                                       `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime       *string                                      `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataDiskCategory   *string                                      `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskSize       *string                                      `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopId          *string                                      `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName        *string                                      `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus      *string                                      `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DesktopType        *string                                      `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	DirectoryId        *string                                      `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Disks              []*DescribeDesktopsResponseBodyDesktopsDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	EndUserIds         []*string                                    `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	ImageId            *string                                      `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	LastStartTime      *string                                      `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	Memory             *int64                                       `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetworkInterfaceId *string                                      `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OfficeSiteId       *string                                      `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PolicyGroupId      *string                                      `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	SystemDiskCategory *string                                      `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskSize     *int32                                       `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktops) SetConnectionStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCpu(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCreationTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDataDiskCategory(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDataDiskSize(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDirectoryId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDisks(v []*DescribeDesktopsResponseBodyDesktopsDisks) *DescribeDesktopsResponseBodyDesktops {
	s.Disks = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetEndUserIds(v []*string) *DescribeDesktopsResponseBodyDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetImageId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetLastStartTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.LastStartTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetMemory(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetNetworkInterfaceId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSystemDiskCategory(v string) *DescribeDesktopsResponseBodyDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSystemDiskSize(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.SystemDiskSize = &v
	return s
}

type DescribeDesktopsResponseBodyDesktopsDisks struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskId(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskSize(v int32) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskType(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskType = &v
	return s
}

type DescribeDesktopsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponse) SetHeaders(v map[string]*string) *DescribeDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopsResponse) SetStatusCode(v int32) *DescribeDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopsResponse) SetBody(v *DescribeDesktopsResponseBody) *DescribeDesktopsResponse {
	s.Body = v
	return s
}

type DescribeDirectoriesRequest struct {
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	// This parameter is required.
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryType(v string) *DescribeDirectoriesRequest {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetMaxResults(v int32) *DescribeDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNextToken(v string) *DescribeDirectoriesRequest {
	s.NextToken = &v
	return s
}

type DescribeDirectoriesResponseBody struct {
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetNextToken(v string) *DescribeDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectories struct {
	ADConnectors          []*DescribeDirectoriesResponseBodyDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" type:"Repeated"`
	CreationTime          *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CustomSecurityGroupId *string                                                   `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty"`
	DirectoryId           *string                                                   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryType         *string                                                   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	DnsAddress            []*string                                                 `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DnsUserName           *string                                                   `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty"`
	DomainName            *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainPassword        *string                                                   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DomainUserName        *string                                                   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	EnableInternetAccess  *bool                                                     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	Name                  *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Status                *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	TrustPassword         *string                                                   `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetADConnectors(v []*DescribeDirectoriesResponseBodyDirectoriesADConnectors) *DescribeDirectoriesResponseBodyDirectories {
	s.ADConnectors = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCreationTime(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CreationTime = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCustomSecurityGroupId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsAddress(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableInternetAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetStatus(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Status = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetTrustPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.TrustPassword = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectoriesADConnectors struct {
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty"`
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectoriesADConnectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectoriesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetADConnectorAddress(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetConnectorStatus(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetVSwitchId(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.VSwitchId = &v
	return s
}

type DescribeDirectoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) SetHeaders(v map[string]*string) *DescribeDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoriesResponse) SetStatusCode(v int32) *DescribeDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDirectoriesResponse) SetBody(v *DescribeDirectoriesResponseBody) *DescribeDirectoriesResponse {
	s.Body = v
	return s
}

type GetConnectionTicketRequest struct {
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UserName             *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetClientOS(v string) *GetConnectionTicketRequest {
	s.ClientOS = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientVersion(v string) *GetConnectionTicketRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetDesktopId(v string) *GetConnectionTicketRequest {
	s.DesktopId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetInstanceId(v string) *GetConnectionTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetOwnerId(v int64) *GetConnectionTicketRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetPassword(v string) *GetConnectionTicketRequest {
	s.Password = &v
	return s
}

func (s *GetConnectionTicketRequest) SetRegionId(v string) *GetConnectionTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerAccount(v string) *GetConnectionTicketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerId(v int64) *GetConnectionTicketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUserName(v string) *GetConnectionTicketRequest {
	s.UserName = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Ticket     *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetHeaders(v map[string]*string) *GetConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionTicketResponse) SetStatusCode(v int32) *GetConnectionTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionTicketResponse) SetBody(v *GetConnectionTicketResponseBody) *GetConnectionTicketResponse {
	s.Body = v
	return s
}

type GetMcpResourceRequest struct {
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetMcpResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMcpResourceRequest) GoString() string {
	return s.String()
}

func (s *GetMcpResourceRequest) SetAuthorization(v string) *GetMcpResourceRequest {
	s.Authorization = &v
	return s
}

func (s *GetMcpResourceRequest) SetSessionId(v string) *GetMcpResourceRequest {
	s.SessionId = &v
	return s
}

type GetMcpResourceResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMcpResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcpResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMcpResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpResourceResponseBody) SetCode(v string) *GetMcpResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpResourceResponseBody) SetData(v *GetMcpResourceResponseBodyData) *GetMcpResourceResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpResourceResponseBody) SetHttpStatusCode(v int32) *GetMcpResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpResourceResponseBody) SetMessage(v string) *GetMcpResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpResourceResponseBody) SetRequestId(v string) *GetMcpResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpResourceResponseBody) SetSuccess(v bool) *GetMcpResourceResponseBody {
	s.Success = &v
	return s
}

type GetMcpResourceResponseBodyData struct {
	ResourceUrl *string `json:"ResourceUrl,omitempty" xml:"ResourceUrl,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetMcpResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMcpResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpResourceResponseBodyData) SetResourceUrl(v string) *GetMcpResourceResponseBodyData {
	s.ResourceUrl = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetSessionId(v string) *GetMcpResourceResponseBodyData {
	s.SessionId = &v
	return s
}

type GetMcpResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcpResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcpResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMcpResourceResponse) GoString() string {
	return s.String()
}

func (s *GetMcpResourceResponse) SetHeaders(v map[string]*string) *GetMcpResourceResponse {
	s.Headers = v
	return s
}

func (s *GetMcpResourceResponse) SetStatusCode(v int32) *GetMcpResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcpResourceResponse) SetBody(v *GetMcpResourceResponseBody) *GetMcpResourceResponse {
	s.Body = v
	return s
}

type ListMcpToolsRequest struct {
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListMcpToolsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMcpToolsRequest) GoString() string {
	return s.String()
}

func (s *ListMcpToolsRequest) SetAuthorization(v string) *ListMcpToolsRequest {
	s.Authorization = &v
	return s
}

type ListMcpToolsResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMcpToolsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMcpToolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpToolsResponseBody) SetCode(v string) *ListMcpToolsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetData(v string) *ListMcpToolsResponseBody {
	s.Data = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetHttpStatusCode(v int32) *ListMcpToolsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetMessage(v string) *ListMcpToolsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetRequestId(v string) *ListMcpToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpToolsResponseBody) SetSuccess(v bool) *ListMcpToolsResponseBody {
	s.Success = &v
	return s
}

type ListMcpToolsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcpToolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcpToolsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMcpToolsResponse) GoString() string {
	return s.String()
}

func (s *ListMcpToolsResponse) SetHeaders(v map[string]*string) *ListMcpToolsResponse {
	s.Headers = v
	return s
}

func (s *ListMcpToolsResponse) SetStatusCode(v int32) *ListMcpToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcpToolsResponse) SetBody(v *ListMcpToolsResponseBody) *ListMcpToolsResponse {
	s.Body = v
	return s
}

type RebootDesktopsRequest struct {
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebootDesktopsRequest) SetClientOS(v string) *RebootDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *RebootDesktopsRequest) SetClientVersion(v string) *RebootDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

type RebootDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponseBody) SetRequestId(v string) *RebootDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type RebootDesktopsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponse) SetHeaders(v map[string]*string) *RebootDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebootDesktopsResponse) SetStatusCode(v int32) *RebootDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootDesktopsResponse) SetBody(v *RebootDesktopsResponseBody) *RebootDesktopsResponse {
	s.Body = v
	return s
}

type ReleaseMcpSessionRequest struct {
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ReleaseMcpSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseMcpSessionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseMcpSessionRequest) SetAuthorization(v string) *ReleaseMcpSessionRequest {
	s.Authorization = &v
	return s
}

func (s *ReleaseMcpSessionRequest) SetSessionId(v string) *ReleaseMcpSessionRequest {
	s.SessionId = &v
	return s
}

type ReleaseMcpSessionResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseMcpSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseMcpSessionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseMcpSessionResponseBody) SetCode(v string) *ReleaseMcpSessionResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseMcpSessionResponseBody) SetHttpStatusCode(v int32) *ReleaseMcpSessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseMcpSessionResponseBody) SetMessage(v string) *ReleaseMcpSessionResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseMcpSessionResponseBody) SetRequestId(v string) *ReleaseMcpSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseMcpSessionResponseBody) SetSuccess(v bool) *ReleaseMcpSessionResponseBody {
	s.Success = &v
	return s
}

type ReleaseMcpSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseMcpSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseMcpSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseMcpSessionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseMcpSessionResponse) SetHeaders(v map[string]*string) *ReleaseMcpSessionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseMcpSessionResponse) SetStatusCode(v int32) *ReleaseMcpSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseMcpSessionResponse) SetBody(v *ReleaseMcpSessionResponseBody) *ReleaseMcpSessionResponse {
	s.Body = v
	return s
}

type StartDesktopsRequest struct {
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) SetClientOS(v string) *StartDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *StartDesktopsRequest) SetClientVersion(v string) *StartDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

type StartDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponseBody) SetRequestId(v string) *StartDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StartDesktopsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponse) SetHeaders(v map[string]*string) *StartDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StartDesktopsResponse) SetStatusCode(v int32) *StartDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDesktopsResponse) SetBody(v *StartDesktopsResponseBody) *StartDesktopsResponse {
	s.Body = v
	return s
}

type StopDesktopsRequest struct {
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) SetClientOS(v string) *StopDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *StopDesktopsRequest) SetClientVersion(v string) *StopDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

type StopDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponseBody) SetRequestId(v string) *StopDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StopDesktopsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponse) SetHeaders(v map[string]*string) *StopDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StopDesktopsResponse) SetStatusCode(v int32) *StopDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDesktopsResponse) SetBody(v *StopDesktopsResponseBody) *StopDesktopsResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
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

// Summary:
//
// 调用mcp工具
//
// @param request - CalMcpToolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CalMcpToolResponse
func (client *Client) CalMcpToolWithOptions(request *CalMcpToolRequest, runtime *util.RuntimeOptions) (_result *CalMcpToolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Args)) {
		body["Args"] = request.Args
	}

	if !tea.BoolValue(util.IsUnset(request.Authorization)) {
		body["Authorization"] = request.Authorization
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUserId)) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Server)) {
		body["Server"] = request.Server
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tool)) {
		body["Tool"] = request.Tool
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CalMcpTool"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CalMcpToolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用mcp工具
//
// @param request - CalMcpToolRequest
//
// @return CalMcpToolResponse
func (client *Client) CalMcpTool(request *CalMcpToolRequest) (_result *CalMcpToolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CalMcpToolResponse{}
	_body, _err := client.CalMcpToolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建 mcp session
//
// @param request - CreateMcpSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpSessionResponse
func (client *Client) CreateMcpSessionWithOptions(request *CreateMcpSessionRequest, runtime *util.RuntimeOptions) (_result *CreateMcpSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Authorization)) {
		body["Authorization"] = request.Authorization
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUserId)) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMcpSession"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMcpSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 mcp session
//
// @param request - CreateMcpSessionRequest
//
// @return CreateMcpSessionResponse
func (client *Client) CreateMcpSession(request *CreateMcpSessionRequest) (_result *CreateMcpSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMcpSessionResponse{}
	_body, _err := client.CreateMcpSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopsResponse
func (client *Client) DescribeDesktopsWithOptions(request *DescribeDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopName)) {
		query["DesktopName"] = request.DesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopStatus)) {
		query["DesktopStatus"] = request.DesktopStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDesktops"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDesktopsRequest
//
// @return DescribeDesktopsResponse
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

// @param request - DescribeDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirectoriesResponse
func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryType)) {
		query["DirectoryType"] = request.DirectoryType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDirectories"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDirectoriesRequest
//
// @return DescribeDirectoriesResponse
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

// @param request - GetConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionTicket"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetConnectionTicketRequest
//
// @return GetConnectionTicketResponse
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

// Summary:
//
// # GetMcpResource
//
// @param request - GetMcpResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpResourceResponse
func (client *Client) GetMcpResourceWithOptions(request *GetMcpResourceRequest, runtime *util.RuntimeOptions) (_result *GetMcpResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Authorization)) {
		body["Authorization"] = request.Authorization
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMcpResource"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMcpResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMcpResource
//
// @param request - GetMcpResourceRequest
//
// @return GetMcpResourceResponse
func (client *Client) GetMcpResource(request *GetMcpResourceRequest) (_result *GetMcpResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMcpResourceResponse{}
	_body, _err := client.GetMcpResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工具列表
//
// @param request - ListMcpToolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpToolsResponse
func (client *Client) ListMcpToolsWithOptions(request *ListMcpToolsRequest, runtime *util.RuntimeOptions) (_result *ListMcpToolsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Authorization)) {
		body["Authorization"] = request.Authorization
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMcpTools"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMcpToolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工具列表
//
// @param request - ListMcpToolsRequest
//
// @return ListMcpToolsResponse
func (client *Client) ListMcpTools(request *ListMcpToolsRequest) (_result *ListMcpToolsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMcpToolsResponse{}
	_body, _err := client.ListMcpToolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RebootDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootDesktopsResponse
func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootDesktops"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RebootDesktopsRequest
//
// @return RebootDesktopsResponse
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

// Summary:
//
// 创建 mcp session
//
// @param request - ReleaseMcpSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseMcpSessionResponse
func (client *Client) ReleaseMcpSessionWithOptions(request *ReleaseMcpSessionRequest, runtime *util.RuntimeOptions) (_result *ReleaseMcpSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Authorization)) {
		body["Authorization"] = request.Authorization
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseMcpSession"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseMcpSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 mcp session
//
// @param request - ReleaseMcpSessionRequest
//
// @return ReleaseMcpSessionResponse
func (client *Client) ReleaseMcpSession(request *ReleaseMcpSessionRequest) (_result *ReleaseMcpSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseMcpSessionResponse{}
	_body, _err := client.ReleaseMcpSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDesktopsResponse
func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *util.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDesktops"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartDesktopsRequest
//
// @return StartDesktopsResponse
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

// @param request - StopDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDesktopsResponse
func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *util.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDesktops"),
		Version:     tea.String("2020-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopDesktopsRequest
//
// @return StopDesktopsResponse
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
