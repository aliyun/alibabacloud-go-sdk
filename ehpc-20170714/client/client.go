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

type AddNodesRequest struct {
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ComputeSpotPriceLimit *string `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ComputeSpotStrategy   *string `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	Count                 *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	ImageId               *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias       *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
}

func (s AddNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddNodesRequest) GoString() string {
	return s.String()
}

func (s *AddNodesRequest) SetClusterId(v string) *AddNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotPriceLimit(v string) *AddNodesRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotStrategy(v string) *AddNodesRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *AddNodesRequest) SetCount(v int32) *AddNodesRequest {
	s.Count = &v
	return s
}

func (s *AddNodesRequest) SetImageId(v string) *AddNodesRequest {
	s.ImageId = &v
	return s
}

func (s *AddNodesRequest) SetImageOwnerAlias(v string) *AddNodesRequest {
	s.ImageOwnerAlias = &v
	return s
}

type AddNodesResponseBody struct {
	InstanceIds *AddNodesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	RequestId   *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddNodesResponseBody) SetInstanceIds(v *AddNodesResponseBodyInstanceIds) *AddNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *AddNodesResponseBody) SetRequestId(v string) *AddNodesResponseBody {
	s.RequestId = &v
	return s
}

type AddNodesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s AddNodesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *AddNodesResponseBodyInstanceIds) SetInstanceId(v []*string) *AddNodesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type AddNodesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponse) GoString() string {
	return s.String()
}

func (s *AddNodesResponse) SetHeaders(v map[string]*string) *AddNodesResponse {
	s.Headers = v
	return s
}

func (s *AddNodesResponse) SetStatusCode(v int32) *AddNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddNodesResponse) SetBody(v *AddNodesResponseBody) *AddNodesResponse {
	s.Body = v
	return s
}

type AddUsersRequest struct {
	ClusterId       *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ReleaseInstance *bool                  `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
	User            []*AddUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s AddUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUsersRequest) GoString() string {
	return s.String()
}

func (s *AddUsersRequest) SetClusterId(v string) *AddUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *AddUsersRequest) SetReleaseInstance(v bool) *AddUsersRequest {
	s.ReleaseInstance = &v
	return s
}

func (s *AddUsersRequest) SetUser(v []*AddUsersRequestUser) *AddUsersRequest {
	s.User = v
	return s
}

type AddUsersRequestUser struct {
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s AddUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s AddUsersRequestUser) GoString() string {
	return s.String()
}

func (s *AddUsersRequestUser) SetGroup(v string) *AddUsersRequestUser {
	s.Group = &v
	return s
}

func (s *AddUsersRequestUser) SetName(v string) *AddUsersRequestUser {
	s.Name = &v
	return s
}

func (s *AddUsersRequestUser) SetPassword(v string) *AddUsersRequestUser {
	s.Password = &v
	return s
}

type AddUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersResponseBody) SetRequestId(v string) *AddUsersResponseBody {
	s.RequestId = &v
	return s
}

type AddUsersResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUsersResponse) GoString() string {
	return s.String()
}

func (s *AddUsersResponse) SetHeaders(v map[string]*string) *AddUsersResponse {
	s.Headers = v
	return s
}

func (s *AddUsersResponse) SetStatusCode(v int32) *AddUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUsersResponse) SetBody(v *AddUsersResponseBody) *AddUsersResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	EcsOrder              *CreateClusterRequestEcsOrder      `json:"EcsOrder,omitempty" xml:"EcsOrder,omitempty" type:"Struct"`
	AccountType           *string                            `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Application           []*CreateClusterRequestApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
	ComputeSpotPriceLimit *string                            `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ComputeSpotStrategy   *string                            `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	Description           *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsChargeType         *string                            `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	EhpcVersion           *string                            `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	HaEnable              *bool                              `json:"HaEnable,omitempty" xml:"HaEnable,omitempty"`
	ImageId               *string                            `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias       *string                            `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	KeyPairName           *string                            `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	Name                  *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	OsTag                 *string                            `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Password              *string                            `json:"Password,omitempty" xml:"Password,omitempty"`
	RemoteDirectory       *string                            `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	SccClusterId          *string                            `json:"SccClusterId,omitempty" xml:"SccClusterId,omitempty"`
	SchedulerType         *string                            `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	SecurityGroupId       *string                            `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName     *string                            `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	VSwitchId             *string                            `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VolumeId              *string                            `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint      *string                            `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol        *string                            `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType            *string                            `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VpcId                 *string                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                *string                            `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetEcsOrder(v *CreateClusterRequestEcsOrder) *CreateClusterRequest {
	s.EcsOrder = v
	return s
}

func (s *CreateClusterRequest) SetAccountType(v string) *CreateClusterRequest {
	s.AccountType = &v
	return s
}

func (s *CreateClusterRequest) SetApplication(v []*CreateClusterRequestApplication) *CreateClusterRequest {
	s.Application = v
	return s
}

func (s *CreateClusterRequest) SetComputeSpotPriceLimit(v string) *CreateClusterRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *CreateClusterRequest) SetComputeSpotStrategy(v string) *CreateClusterRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetEcsChargeType(v string) *CreateClusterRequest {
	s.EcsChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetEhpcVersion(v string) *CreateClusterRequest {
	s.EhpcVersion = &v
	return s
}

func (s *CreateClusterRequest) SetHaEnable(v bool) *CreateClusterRequest {
	s.HaEnable = &v
	return s
}

func (s *CreateClusterRequest) SetImageId(v string) *CreateClusterRequest {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequest) SetImageOwnerAlias(v string) *CreateClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateClusterRequest) SetKeyPairName(v string) *CreateClusterRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetOsTag(v string) *CreateClusterRequest {
	s.OsTag = &v
	return s
}

func (s *CreateClusterRequest) SetPassword(v string) *CreateClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateClusterRequest) SetRemoteDirectory(v string) *CreateClusterRequest {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateClusterRequest) SetSccClusterId(v string) *CreateClusterRequest {
	s.SccClusterId = &v
	return s
}

func (s *CreateClusterRequest) SetSchedulerType(v string) *CreateClusterRequest {
	s.SchedulerType = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupName(v string) *CreateClusterRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeId(v string) *CreateClusterRequest {
	s.VolumeId = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeMountpoint(v string) *CreateClusterRequest {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeProtocol(v string) *CreateClusterRequest {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeType(v string) *CreateClusterRequest {
	s.VolumeType = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

type CreateClusterRequestEcsOrder struct {
	Compute *CreateClusterRequestEcsOrderCompute `json:"Compute,omitempty" xml:"Compute,omitempty" require:"true" type:"Struct"`
	Login   *CreateClusterRequestEcsOrderLogin   `json:"Login,omitempty" xml:"Login,omitempty" require:"true" type:"Struct"`
	Manager *CreateClusterRequestEcsOrderManager `json:"Manager,omitempty" xml:"Manager,omitempty" require:"true" type:"Struct"`
}

func (s CreateClusterRequestEcsOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrder) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrder) SetCompute(v *CreateClusterRequestEcsOrderCompute) *CreateClusterRequestEcsOrder {
	s.Compute = v
	return s
}

func (s *CreateClusterRequestEcsOrder) SetLogin(v *CreateClusterRequestEcsOrderLogin) *CreateClusterRequestEcsOrder {
	s.Login = v
	return s
}

func (s *CreateClusterRequestEcsOrder) SetManager(v *CreateClusterRequestEcsOrderManager) *CreateClusterRequestEcsOrder {
	s.Manager = v
	return s
}

type CreateClusterRequestEcsOrderCompute struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderCompute) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderCompute) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderCompute) SetCount(v int32) *CreateClusterRequestEcsOrderCompute {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderCompute) SetInstanceType(v string) *CreateClusterRequestEcsOrderCompute {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestEcsOrderLogin struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderLogin) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderLogin) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderLogin) SetCount(v int32) *CreateClusterRequestEcsOrderLogin {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderLogin) SetInstanceType(v string) *CreateClusterRequestEcsOrderLogin {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestEcsOrderManager struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderManager) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderManager) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderManager) SetCount(v int32) *CreateClusterRequestEcsOrderManager {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderManager) SetInstanceType(v string) *CreateClusterRequestEcsOrderManager {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestApplication struct {
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateClusterRequestApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestApplication) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestApplication) SetTag(v string) *CreateClusterRequestApplication {
	s.Tag = &v
	return s
}

type CreateClusterResponseBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateJobTemplateRequest struct {
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreateJobTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequest) SetArrayRequest(v string) *CreateJobTemplateRequest {
	s.ArrayRequest = &v
	return s
}

func (s *CreateJobTemplateRequest) SetCommandLine(v string) *CreateJobTemplateRequest {
	s.CommandLine = &v
	return s
}

func (s *CreateJobTemplateRequest) SetName(v string) *CreateJobTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateJobTemplateRequest) SetPackagePath(v string) *CreateJobTemplateRequest {
	s.PackagePath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetPriority(v int32) *CreateJobTemplateRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobTemplateRequest) SetReRunable(v bool) *CreateJobTemplateRequest {
	s.ReRunable = &v
	return s
}

func (s *CreateJobTemplateRequest) SetRegionId(v string) *CreateJobTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateJobTemplateRequest) SetRunasUser(v string) *CreateJobTemplateRequest {
	s.RunasUser = &v
	return s
}

func (s *CreateJobTemplateRequest) SetStderrRedirectPath(v string) *CreateJobTemplateRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetStdoutRedirectPath(v string) *CreateJobTemplateRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetVariables(v string) *CreateJobTemplateRequest {
	s.Variables = &v
	return s
}

type CreateJobTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateJobTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponseBody) SetRequestId(v string) *CreateJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetTemplateId(v string) *CreateJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateJobTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponse) SetHeaders(v map[string]*string) *CreateJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateJobTemplateResponse) SetStatusCode(v int32) *CreateJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobTemplateResponse) SetBody(v *CreateJobTemplateResponseBody) *CreateJobTemplateResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ReleaseInstance *string `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetReleaseInstance(v string) *DeleteClusterRequest {
	s.ReleaseInstance = &v
	return s
}

type DeleteClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteJobTemplatesRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Templates *string `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s DeleteJobTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesRequest) SetRegionId(v string) *DeleteJobTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteJobTemplatesRequest) SetTemplates(v string) *DeleteJobTemplatesRequest {
	s.Templates = &v
	return s
}

type DeleteJobTemplatesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesResponseBody) SetRequestId(v string) *DeleteJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesResponse) SetHeaders(v map[string]*string) *DeleteJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobTemplatesResponse) SetStatusCode(v int32) *DeleteJobTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobTemplatesResponse) SetBody(v *DeleteJobTemplatesResponseBody) *DeleteJobTemplatesResponse {
	s.Body = v
	return s
}

type DeleteJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s DeleteJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) SetClusterId(v string) *DeleteJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteJobsRequest) SetJobs(v string) *DeleteJobsRequest {
	s.Jobs = &v
	return s
}

type DeleteJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponse) SetHeaders(v map[string]*string) *DeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobsResponse) SetStatusCode(v int32) *DeleteJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobsResponse) SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse {
	s.Body = v
	return s
}

type DeleteNodesRequest struct {
	ClusterId       *string                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance        []*DeleteNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
	ReleaseInstance *bool                         `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
}

func (s DeleteNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequest) SetClusterId(v string) *DeleteNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodesRequest) SetInstance(v []*DeleteNodesRequestInstance) *DeleteNodesRequest {
	s.Instance = v
	return s
}

func (s *DeleteNodesRequest) SetReleaseInstance(v bool) *DeleteNodesRequest {
	s.ReleaseInstance = &v
	return s
}

type DeleteNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequestInstance) SetId(v string) *DeleteNodesRequestInstance {
	s.Id = &v
	return s
}

type DeleteNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponseBody) SetRequestId(v string) *DeleteNodesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponse) SetHeaders(v map[string]*string) *DeleteNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodesResponse) SetStatusCode(v int32) *DeleteNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodesResponse) SetBody(v *DeleteNodesResponseBody) *DeleteNodesResponse {
	s.Body = v
	return s
}

type DeleteUsersRequest struct {
	ClusterId *string                   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*DeleteUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DeleteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequest) SetClusterId(v string) *DeleteUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersRequest) SetUser(v []*DeleteUsersRequestUser) *DeleteUsersRequest {
	s.User = v
	return s
}

type DeleteUsersRequestUser struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequestUser) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequestUser) SetName(v string) *DeleteUsersRequestUser {
	s.Name = &v
	return s
}

type DeleteUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponseBody) SetRequestId(v string) *DeleteUsersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponse) SetHeaders(v map[string]*string) *DeleteUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteUsersResponse) SetStatusCode(v int32) *DeleteUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUsersResponse) SetBody(v *DeleteUsersResponseBody) *DeleteUsersResponse {
	s.Body = v
	return s
}

type DescribeClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterResponseBody struct {
	ClusterInfo *DescribeClusterResponseBodyClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) SetClusterInfo(v *DescribeClusterResponseBodyClusterInfo) *DescribeClusterResponseBody {
	s.ClusterInfo = v
	return s
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterResponseBodyClusterInfo struct {
	AccountType      *string                                             `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Applications     *DescribeClusterResponseBodyClusterInfoApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	ClientVersion    *string                                             `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CreateTime       *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsChargeType    *string                                             `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	EcsInfo          *DescribeClusterResponseBodyClusterInfoEcsInfo      `json:"EcsInfo,omitempty" xml:"EcsInfo,omitempty" type:"Struct"`
	HaEnable         *bool                                               `json:"HaEnable,omitempty" xml:"HaEnable,omitempty"`
	Id               *string                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId          *string                                             `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias  *string                                             `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	KeyPairName      *string                                             `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	Name             *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	OsTag            *string                                             `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	RegionId         *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemoteDirectory  *string                                             `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	SccClusterId     *string                                             `json:"SccClusterId,omitempty" xml:"SccClusterId,omitempty"`
	SchedulerType    *string                                             `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	SecurityGroupId  *string                                             `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Status           *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId        *string                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VolumeId         *string                                             `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint *string                                             `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol   *string                                             `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType       *string                                             `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfo) SetAccountType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.AccountType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetApplications(v *DescribeClusterResponseBodyClusterInfoApplications) *DescribeClusterResponseBodyClusterInfo {
	s.Applications = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetClientVersion(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ClientVersion = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetCreateTime(v string) *DescribeClusterResponseBodyClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetDescription(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Description = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetEcsChargeType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.EcsChargeType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetEcsInfo(v *DescribeClusterResponseBodyClusterInfoEcsInfo) *DescribeClusterResponseBodyClusterInfo {
	s.EcsInfo = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetHaEnable(v bool) *DescribeClusterResponseBodyClusterInfo {
	s.HaEnable = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Id = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageOwnerAlias(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetKeyPairName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.KeyPairName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetOsTag(v string) *DescribeClusterResponseBodyClusterInfo {
	s.OsTag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRegionId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRemoteDirectory(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSccClusterId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SccClusterId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSchedulerType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SchedulerType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSecurityGroupId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetStatus(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVSwitchId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeMountpoint(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeProtocol(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoApplications struct {
	ApplicationInfo []*DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterResponseBodyClusterInfoApplications) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoApplications) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoApplications) SetApplicationInfo(v []*DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) *DescribeClusterResponseBodyClusterInfoApplications {
	s.ApplicationInfo = v
	return s
}

type DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag     *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetName(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetTag(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Tag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo) SetVersion(v string) *DescribeClusterResponseBodyClusterInfoApplicationsApplicationInfo {
	s.Version = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfo struct {
	Compute *DescribeClusterResponseBodyClusterInfoEcsInfoCompute `json:"Compute,omitempty" xml:"Compute,omitempty" type:"Struct"`
	Login   *DescribeClusterResponseBodyClusterInfoEcsInfoLogin   `json:"Login,omitempty" xml:"Login,omitempty" type:"Struct"`
	Manager *DescribeClusterResponseBodyClusterInfoEcsInfoManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetCompute(v *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Compute = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetLogin(v *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Login = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetManager(v *DescribeClusterResponseBodyClusterInfoEcsInfoManager) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Manager = v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoCompute struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoCompute) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoCompute) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoCompute {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoCompute {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoLogin struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoLogin) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoLogin) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoLogin {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoLogin {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoManager struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoManager) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoManager) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoManager) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoManager {
	s.Count = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoManager) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoManager {
	s.InstanceType = &v
	return s
}

type DescribeClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetStatusCode(v int32) *DescribeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

type EditJobTemplateRequest struct {
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s EditJobTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *EditJobTemplateRequest) SetArrayRequest(v string) *EditJobTemplateRequest {
	s.ArrayRequest = &v
	return s
}

func (s *EditJobTemplateRequest) SetCommandLine(v string) *EditJobTemplateRequest {
	s.CommandLine = &v
	return s
}

func (s *EditJobTemplateRequest) SetName(v string) *EditJobTemplateRequest {
	s.Name = &v
	return s
}

func (s *EditJobTemplateRequest) SetPackagePath(v string) *EditJobTemplateRequest {
	s.PackagePath = &v
	return s
}

func (s *EditJobTemplateRequest) SetPriority(v int32) *EditJobTemplateRequest {
	s.Priority = &v
	return s
}

func (s *EditJobTemplateRequest) SetReRunable(v bool) *EditJobTemplateRequest {
	s.ReRunable = &v
	return s
}

func (s *EditJobTemplateRequest) SetRegionId(v string) *EditJobTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *EditJobTemplateRequest) SetRunasUser(v string) *EditJobTemplateRequest {
	s.RunasUser = &v
	return s
}

func (s *EditJobTemplateRequest) SetStderrRedirectPath(v string) *EditJobTemplateRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *EditJobTemplateRequest) SetStdoutRedirectPath(v string) *EditJobTemplateRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *EditJobTemplateRequest) SetTemplateId(v string) *EditJobTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *EditJobTemplateRequest) SetVariables(v string) *EditJobTemplateRequest {
	s.Variables = &v
	return s
}

type EditJobTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s EditJobTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *EditJobTemplateResponseBody) SetRequestId(v string) *EditJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditJobTemplateResponseBody) SetTemplateId(v string) *EditJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type EditJobTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EditJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditJobTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *EditJobTemplateResponse) SetHeaders(v map[string]*string) *EditJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *EditJobTemplateResponse) SetStatusCode(v int32) *EditJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *EditJobTemplateResponse) SetBody(v *EditJobTemplateResponseBody) *EditJobTemplateResponse {
	s.Body = v
	return s
}

type GetAutoScaleConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigRequest) SetClusterId(v string) *GetAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

type GetAutoScaleConfigResponseBody struct {
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType             *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EnableAutoGrow          *bool   `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	EnableAutoShrink        *bool   `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	ExcludeNodes            *string `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	ExtraNodesGrowRatio     *int32  `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	GrowIntervalInMinutes   *int32  `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	GrowRatio               *int32  `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	GrowTimeoutInMinutes    *int32  `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	MaxNodesInCluster       *int32  `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShrinkIdleTimes         *int32  `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	ShrinkIntervalInMinutes *int32  `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	Uid                     *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBody) SetClusterId(v string) *GetAutoScaleConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetClusterType(v string) *GetAutoScaleConfigResponseBody {
	s.ClusterType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetEnableAutoGrow(v bool) *GetAutoScaleConfigResponseBody {
	s.EnableAutoGrow = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetEnableAutoShrink(v bool) *GetAutoScaleConfigResponseBody {
	s.EnableAutoShrink = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetExcludeNodes(v string) *GetAutoScaleConfigResponseBody {
	s.ExcludeNodes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetExtraNodesGrowRatio(v int32) *GetAutoScaleConfigResponseBody {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowIntervalInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowRatio(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowRatio = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowTimeoutInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetMaxNodesInCluster(v int32) *GetAutoScaleConfigResponseBody {
	s.MaxNodesInCluster = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetRequestId(v string) *GetAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetShrinkIdleTimes(v int32) *GetAutoScaleConfigResponseBody {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetShrinkIntervalInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetUid(v string) *GetAutoScaleConfigResponseBody {
	s.Uid = &v
	return s
}

type GetAutoScaleConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponse) SetHeaders(v map[string]*string) *GetAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScaleConfigResponse) SetStatusCode(v int32) *GetAutoScaleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScaleConfigResponse) SetBody(v *GetAutoScaleConfigResponseBody) *GetAutoScaleConfigResponse {
	s.Body = v
	return s
}

type ListClusterLogsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterLogsRequest) SetClusterId(v string) *ListClusterLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterLogsRequest) SetPageNumber(v int32) *ListClusterLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterLogsRequest) SetPageSize(v int32) *ListClusterLogsRequest {
	s.PageSize = &v
	return s
}

type ListClusterLogsResponseBody struct {
	ClusterId  *string                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Logs       *ListClusterLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClusterLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBody) SetClusterId(v string) *ListClusterLogsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetLogs(v *ListClusterLogsResponseBodyLogs) *ListClusterLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListClusterLogsResponseBody) SetPageNumber(v int32) *ListClusterLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetPageSize(v int32) *ListClusterLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetRequestId(v string) *ListClusterLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetTotalCount(v int32) *ListClusterLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListClusterLogsResponseBodyLogs struct {
	LogInfo []*ListClusterLogsResponseBodyLogsLogInfo `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
}

func (s ListClusterLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBodyLogs) SetLogInfo(v []*ListClusterLogsResponseBodyLogsLogInfo) *ListClusterLogsResponseBodyLogs {
	s.LogInfo = v
	return s
}

type ListClusterLogsResponseBodyLogsLogInfo struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Level      *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Operation  *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s ListClusterLogsResponseBodyLogsLogInfo) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBodyLogsLogInfo) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetCreateTime(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.CreateTime = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetLevel(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Level = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetMessage(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Message = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogsLogInfo) SetOperation(v string) *ListClusterLogsResponseBodyLogsLogInfo {
	s.Operation = &v
	return s
}

type ListClusterLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponse) SetHeaders(v map[string]*string) *ListClusterLogsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterLogsResponse) SetStatusCode(v int32) *ListClusterLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterLogsResponse) SetBody(v *ListClusterLogsResponseBody) *ListClusterLogsResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetPageNumber(v int32) *ListClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

type ListClustersResponseBody struct {
	Clusters   *ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v *ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v int32) *ListClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

type ListClustersResponseBodyClusters struct {
	ClusterInfoSimple []*ListClustersResponseBodyClustersClusterInfoSimple `json:"ClusterInfoSimple,omitempty" xml:"ClusterInfoSimple,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterInfoSimple(v []*ListClustersResponseBodyClustersClusterInfoSimple) *ListClustersResponseBodyClusters {
	s.ClusterInfoSimple = v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimple struct {
	AccountType     *string                                                          `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Computes        *ListClustersResponseBodyClustersClusterInfoSimpleComputes       `json:"Computes,omitempty" xml:"Computes,omitempty" type:"Struct"`
	Count           *int32                                                           `json:"Count,omitempty" xml:"Count,omitempty"`
	CreateTime      *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *string                                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId         *string                                                          `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias *string                                                          `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	InstanceType    *string                                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LoginNodes      *string                                                          `json:"LoginNodes,omitempty" xml:"LoginNodes,omitempty"`
	Managers        *ListClustersResponseBodyClustersClusterInfoSimpleManagers       `json:"Managers,omitempty" xml:"Managers,omitempty" type:"Struct"`
	Name            *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	OsTag           *string                                                          `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	RegionId        *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SchedulerType   *string                                                          `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	Status          *string                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalResources  *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	UsedResources   *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	ZoneId          *string                                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimple) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimple) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetAccountType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.AccountType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetComputes(v *ListClustersResponseBodyClustersClusterInfoSimpleComputes) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Computes = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Count = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetCreateTime(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetDescription(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Id = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetImageId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ImageId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetImageOwnerAlias(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetInstanceType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.InstanceType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetLoginNodes(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.LoginNodes = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetManagers(v *ListClustersResponseBodyClustersClusterInfoSimpleManagers) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Managers = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetName(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetOsTag(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.OsTag = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetRegionId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.RegionId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetSchedulerType(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.SchedulerType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetStatus(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetTotalResources(v *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.TotalResources = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetUsedResources(v *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.UsedResources = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimple) SetZoneId(v string) *ListClustersResponseBodyClustersClusterInfoSimple {
	s.ZoneId = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleComputes struct {
	ExceptionCount *int32 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	NormalCount    *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	Total          *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleComputes) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleComputes) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetExceptionCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.ExceptionCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetNormalCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.NormalCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleComputes) SetTotal(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleComputes {
	s.Total = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleManagers struct {
	ExceptionCount *int32 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	NormalCount    *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	Total          *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleManagers) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleManagers) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetExceptionCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.ExceptionCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetNormalCount(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.NormalCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleManagers) SetTotal(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleManagers {
	s.Total = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetCpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetGpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources) SetMemory(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleTotalResources {
	s.Memory = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoSimpleUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetCpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetGpu(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources) SetMemory(v int32) *ListClustersResponseBodyClustersClusterInfoSimpleUsedResources {
	s.Memory = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListCurrentClientVersionResponseBody struct {
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCurrentClientVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCurrentClientVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListCurrentClientVersionResponseBody) SetClientVersion(v string) *ListCurrentClientVersionResponseBody {
	s.ClientVersion = &v
	return s
}

func (s *ListCurrentClientVersionResponseBody) SetRequestId(v string) *ListCurrentClientVersionResponseBody {
	s.RequestId = &v
	return s
}

type ListCurrentClientVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCurrentClientVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCurrentClientVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCurrentClientVersionResponse) GoString() string {
	return s.String()
}

func (s *ListCurrentClientVersionResponse) SetHeaders(v map[string]*string) *ListCurrentClientVersionResponse {
	s.Headers = v
	return s
}

func (s *ListCurrentClientVersionResponse) SetStatusCode(v int32) *ListCurrentClientVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCurrentClientVersionResponse) SetBody(v *ListCurrentClientVersionResponseBody) *ListCurrentClientVersionResponse {
	s.Body = v
	return s
}

type ListCustomImagesRequest struct {
	BaseOsTag       *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomImagesRequest) SetBaseOsTag(v string) *ListCustomImagesRequest {
	s.BaseOsTag = &v
	return s
}

func (s *ListCustomImagesRequest) SetImageOwnerAlias(v string) *ListCustomImagesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListCustomImagesRequest) SetRegionId(v string) *ListCustomImagesRequest {
	s.RegionId = &v
	return s
}

type ListCustomImagesResponseBody struct {
	Images    *ListCustomImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCustomImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBody) SetImages(v *ListCustomImagesResponseBodyImages) *ListCustomImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListCustomImagesResponseBody) SetRequestId(v string) *ListCustomImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListCustomImagesResponseBodyImages struct {
	ImageInfo []*ListCustomImagesResponseBodyImagesImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Repeated"`
}

func (s ListCustomImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImages) SetImageInfo(v []*ListCustomImagesResponseBodyImagesImageInfo) *ListCustomImagesResponseBodyImages {
	s.ImageInfo = v
	return s
}

type ListCustomImagesResponseBodyImagesImageInfo struct {
	BaseOsTag       *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty" type:"Struct"`
	Description     *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId         *string                                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName       *string                                               `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageOwnerAlias *string                                               `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesImageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesImageInfo) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetBaseOsTag(v *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) *ListCustomImagesResponseBodyImagesImageInfo {
	s.BaseOsTag = v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetDescription(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.Description = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageId(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageId = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageName(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageName = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfo) SetImageOwnerAlias(v string) *ListCustomImagesResponseBodyImagesImageInfo {
	s.ImageOwnerAlias = &v
	return s
}

type ListCustomImagesResponseBodyImagesImageInfoBaseOsTag struct {
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetArchitecture(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Architecture = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetOsTag(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetPlatform(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Platform = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag) SetVersion(v string) *ListCustomImagesResponseBodyImagesImageInfoBaseOsTag {
	s.Version = &v
	return s
}

type ListCustomImagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponse) SetHeaders(v map[string]*string) *ListCustomImagesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomImagesResponse) SetStatusCode(v int32) *ListCustomImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomImagesResponse) SetBody(v *ListCustomImagesResponseBody) *ListCustomImagesResponse {
	s.Body = v
	return s
}

type ListImagesResponseBody struct {
	OsTags    *ListImagesResponseBodyOsTags `json:"OsTags,omitempty" xml:"OsTags,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetOsTags(v *ListImagesResponseBodyOsTags) *ListImagesResponseBody {
	s.OsTags = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyOsTags struct {
	OsInfo []*ListImagesResponseBodyOsTagsOsInfo `json:"OsInfo,omitempty" xml:"OsInfo,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBodyOsTags) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyOsTags) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyOsTags) SetOsInfo(v []*ListImagesResponseBodyOsTagsOsInfo) *ListImagesResponseBodyOsTags {
	s.OsInfo = v
	return s
}

type ListImagesResponseBodyOsTagsOsInfo struct {
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListImagesResponseBodyOsTagsOsInfo) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyOsTagsOsInfo) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetArchitecture(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Architecture = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetOsTag(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.OsTag = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetPlatform(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyOsTagsOsInfo) SetVersion(v string) *ListImagesResponseBodyOsTagsOsInfo {
	s.Version = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListJobTemplatesRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListJobTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesRequest) SetName(v string) *ListJobTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageNumber(v int32) *ListJobTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageSize(v int32) *ListJobTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobTemplatesRequest) SetRegionId(v string) *ListJobTemplatesRequest {
	s.RegionId = &v
	return s
}

type ListJobTemplatesResponseBody struct {
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates  *ListJobTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBody) SetPageNumber(v int32) *ListJobTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetPageSize(v int32) *ListJobTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetRequestId(v string) *ListJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetTemplates(v *ListJobTemplatesResponseBodyTemplates) *ListJobTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListJobTemplatesResponseBody) SetTotalCount(v int32) *ListJobTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobTemplatesResponseBodyTemplates struct {
	JobTemplates []*ListJobTemplatesResponseBodyTemplatesJobTemplates `json:"JobTemplates,omitempty" xml:"JobTemplates,omitempty" type:"Repeated"`
}

func (s ListJobTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyTemplates) SetJobTemplates(v []*ListJobTemplatesResponseBodyTemplatesJobTemplates) *ListJobTemplatesResponseBodyTemplates {
	s.JobTemplates = v
	return s
}

type ListJobTemplatesResponseBodyTemplatesJobTemplates struct {
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ListJobTemplatesResponseBodyTemplatesJobTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBodyTemplatesJobTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetArrayRequest(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetCommandLine(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.CommandLine = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetId(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Id = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetName(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Name = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetPackagePath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.PackagePath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetPriority(v int32) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Priority = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetReRunable(v bool) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.ReRunable = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetRunasUser(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.RunasUser = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetStderrRedirectPath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.StderrRedirectPath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetStdoutRedirectPath(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.StdoutRedirectPath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplatesJobTemplates) SetVariables(v string) *ListJobTemplatesResponseBodyTemplatesJobTemplates {
	s.Variables = &v
	return s
}

type ListJobTemplatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponse) SetHeaders(v map[string]*string) *ListJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListJobTemplatesResponse) SetStatusCode(v int32) *ListJobTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobTemplatesResponse) SetBody(v *ListJobTemplatesResponseBody) *ListJobTemplatesResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rerunable  *string `json:"Rerunable,omitempty" xml:"Rerunable,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetClusterId(v string) *ListJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsRequest) SetOwner(v string) *ListJobsRequest {
	s.Owner = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetRerunable(v string) *ListJobsRequest {
	s.Rerunable = &v
	return s
}

func (s *ListJobsRequest) SetState(v string) *ListJobsRequest {
	s.State = &v
	return s
}

type ListJobsResponseBody struct {
	Jobs       *ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Struct"`
	PageNumber *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobs(v *ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponseBodyJobs struct {
	JobInfo []*ListJobsResponseBodyJobsJobInfo `json:"JobInfo,omitempty" xml:"JobInfo,omitempty" type:"Repeated"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) SetJobInfo(v []*ListJobsResponseBodyJobsJobInfo) *ListJobsResponseBodyJobs {
	s.JobInfo = v
	return s
}

type ListJobsResponseBodyJobsJobInfo struct {
	ArrayRequest   *string                                   `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	Comment        *string                                   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Id             *string                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	LastModifyTime *string                                   `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	Name           *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner          *string                                   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Priority       *int32                                    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Resources      *ListJobsResponseBodyJobsJobInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	ShellPath      *string                                   `json:"ShellPath,omitempty" xml:"ShellPath,omitempty"`
	StartTime      *string                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State          *string                                   `json:"State,omitempty" xml:"State,omitempty"`
	Stderr         *string                                   `json:"Stderr,omitempty" xml:"Stderr,omitempty"`
	Stdout         *string                                   `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
	SubmitTime     *string                                   `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s ListJobsResponseBodyJobsJobInfo) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobInfo) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobInfo) SetArrayRequest(v string) *ListJobsResponseBodyJobsJobInfo {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetComment(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Comment = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetId(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Id = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetLastModifyTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.LastModifyTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetName(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetOwner(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Owner = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetPriority(v int32) *ListJobsResponseBodyJobsJobInfo {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetResources(v *ListJobsResponseBodyJobsJobInfoResources) *ListJobsResponseBodyJobsJobInfo {
	s.Resources = v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetShellPath(v string) *ListJobsResponseBodyJobsJobInfo {
	s.ShellPath = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStartTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetState(v string) *ListJobsResponseBodyJobsJobInfo {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStderr(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Stderr = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetStdout(v string) *ListJobsResponseBodyJobsJobInfo {
	s.Stdout = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfo) SetSubmitTime(v string) *ListJobsResponseBodyJobsJobInfo {
	s.SubmitTime = &v
	return s
}

type ListJobsResponseBodyJobsJobInfoResources struct {
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Nodes *int32 `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ListJobsResponseBodyJobsJobInfoResources) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobInfoResources) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobInfoResources) SetCores(v int32) *ListJobsResponseBodyJobsJobInfoResources {
	s.Cores = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobInfoResources) SetNodes(v int32) *ListJobsResponseBodyJobsJobInfoResources {
	s.Nodes = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostName   *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetClusterId(v string) *ListNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesRequest) SetHostName(v string) *ListNodesRequest {
	s.HostName = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetRole(v string) *ListNodesRequest {
	s.Role = &v
	return s
}

type ListNodesResponseBody struct {
	Nodes      *ListNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	PageNumber *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetNodes(v *ListNodesResponseBodyNodes) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBody) SetPageNumber(v int32) *ListNodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBody) SetPageSize(v int32) *ListNodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetTotalCount(v int32) *ListNodesResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodesResponseBodyNodes struct {
	NodeInfo []*ListNodesResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodes) SetNodeInfo(v []*ListNodesResponseBodyNodesNodeInfo) *ListNodesResponseBodyNodes {
	s.NodeInfo = v
	return s
}

type ListNodesResponseBodyNodesNodeInfo struct {
	AddTime         *string                                           `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	CreatedByEhpc   *bool                                             `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	Expired         *bool                                             `json:"Expired,omitempty" xml:"Expired,omitempty"`
	ExpiredTime     *string                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Id              *string                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId         *string                                           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias *string                                           `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	LockReason      *string                                           `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	RegionId        *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Role            *string                                           `json:"Role,omitempty" xml:"Role,omitempty"`
	SpotStrategy    *string                                           `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Status          *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalResources  *ListNodesResponseBodyNodesNodeInfoTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	UsedResources   *ListNodesResponseBodyNodesNodeInfoUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyNodesNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetAddTime(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.AddTime = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetCreatedByEhpc(v bool) *ListNodesResponseBodyNodesNodeInfo {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetExpired(v bool) *ListNodesResponseBodyNodesNodeInfo {
	s.Expired = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetExpiredTime(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetImageId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ImageId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetImageOwnerAlias(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetLockReason(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.LockReason = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetRegionId(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.RegionId = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetRole(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Role = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetSpotStrategy(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetStatus(v string) *ListNodesResponseBodyNodesNodeInfo {
	s.Status = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetTotalResources(v *ListNodesResponseBodyNodesNodeInfoTotalResources) *ListNodesResponseBodyNodesNodeInfo {
	s.TotalResources = v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfo) SetUsedResources(v *ListNodesResponseBodyNodesNodeInfoUsedResources) *ListNodesResponseBodyNodesNodeInfo {
	s.UsedResources = v
	return s
}

type ListNodesResponseBodyNodesNodeInfoTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesNodeInfoTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfoTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetCpu(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetGpu(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoTotalResources) SetMemory(v int32) *ListNodesResponseBodyNodesNodeInfoTotalResources {
	s.Memory = &v
	return s
}

type ListNodesResponseBodyNodesNodeInfoUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesNodeInfoUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesNodeInfoUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetCpu(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetGpu(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesNodeInfoUsedResources) SetMemory(v int32) *ListNodesResponseBodyNodesNodeInfoUsedResources {
	s.Memory = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListNodesNoPagingRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostName     *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OnlyDetached *bool   `json:"OnlyDetached,omitempty" xml:"OnlyDetached,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListNodesNoPagingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingRequest) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingRequest) SetClusterId(v string) *ListNodesNoPagingRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetHostName(v string) *ListNodesNoPagingRequest {
	s.HostName = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetOnlyDetached(v bool) *ListNodesNoPagingRequest {
	s.OnlyDetached = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetRole(v string) *ListNodesNoPagingRequest {
	s.Role = &v
	return s
}

type ListNodesNoPagingResponseBody struct {
	Nodes      *ListNodesNoPagingResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesNoPagingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBody) SetNodes(v *ListNodesNoPagingResponseBodyNodes) *ListNodesNoPagingResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetPageNumber(v int32) *ListNodesNoPagingResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetPageSize(v int32) *ListNodesNoPagingResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetRequestId(v string) *ListNodesNoPagingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetTotalCount(v int32) *ListNodesNoPagingResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodesNoPagingResponseBodyNodes struct {
	NodeInfo []*ListNodesNoPagingResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s ListNodesNoPagingResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodes) SetNodeInfo(v []*ListNodesNoPagingResponseBodyNodesNodeInfo) *ListNodesNoPagingResponseBodyNodes {
	s.NodeInfo = v
	return s
}

type ListNodesNoPagingResponseBodyNodesNodeInfo struct {
	AddTime         *string                                                   `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	CreatedByEhpc   *bool                                                     `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	Expired         *bool                                                     `json:"Expired,omitempty" xml:"Expired,omitempty"`
	ExpiredTime     *string                                                   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Id              *string                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageId         *string                                                   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOwnerAlias *string                                                   `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	LockReason      *string                                                   `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	RegionId        *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Role            *string                                                   `json:"Role,omitempty" xml:"Role,omitempty"`
	SpotStrategy    *string                                                   `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Status          *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalResources  *ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	UsedResources   *ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetAddTime(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.AddTime = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetCreatedByEhpc(v bool) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetExpired(v bool) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.Expired = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetExpiredTime(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetId(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.Id = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetImageId(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.ImageId = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetImageOwnerAlias(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetLockReason(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.LockReason = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetRegionId(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.RegionId = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetRole(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.Role = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetSpotStrategy(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetStatus(v string) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.Status = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetTotalResources(v *ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.TotalResources = v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfo) SetUsedResources(v *ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources) *ListNodesNoPagingResponseBodyNodesNodeInfo {
	s.UsedResources = v
	return s
}

type ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources) SetCpu(v int32) *ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources) SetGpu(v int32) *ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources) SetMemory(v int32) *ListNodesNoPagingResponseBodyNodesNodeInfoTotalResources {
	s.Memory = &v
	return s
}

type ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources) SetCpu(v int32) *ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources) SetGpu(v int32) *ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources) SetMemory(v int32) *ListNodesNoPagingResponseBodyNodesNodeInfoUsedResources {
	s.Memory = &v
	return s
}

type ListNodesNoPagingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodesNoPagingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesNoPagingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponse) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponse) SetHeaders(v map[string]*string) *ListNodesNoPagingResponse {
	s.Headers = v
	return s
}

func (s *ListNodesNoPagingResponse) SetStatusCode(v int32) *ListNodesNoPagingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesNoPagingResponse) SetBody(v *ListNodesNoPagingResponseBody) *ListNodesNoPagingResponse {
	s.Body = v
	return s
}

type ListPreferredEcsTypesRequest struct {
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListPreferredEcsTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesRequest) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesRequest) SetSpotStrategy(v string) *ListPreferredEcsTypesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ListPreferredEcsTypesRequest) SetZoneId(v string) *ListPreferredEcsTypesRequest {
	s.ZoneId = &v
	return s
}

type ListPreferredEcsTypesResponseBody struct {
	RequestId           *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Series              *ListPreferredEcsTypesResponseBodySeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Struct"`
	SupportSpotInstance *bool                                    `json:"SupportSpotInstance,omitempty" xml:"SupportSpotInstance,omitempty"`
}

func (s ListPreferredEcsTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBody) SetRequestId(v string) *ListPreferredEcsTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBody) SetSeries(v *ListPreferredEcsTypesResponseBodySeries) *ListPreferredEcsTypesResponseBody {
	s.Series = v
	return s
}

func (s *ListPreferredEcsTypesResponseBody) SetSupportSpotInstance(v bool) *ListPreferredEcsTypesResponseBody {
	s.SupportSpotInstance = &v
	return s
}

type ListPreferredEcsTypesResponseBodySeries struct {
	SeriesInfo []*ListPreferredEcsTypesResponseBodySeriesSeriesInfo `json:"SeriesInfo,omitempty" xml:"SeriesInfo,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeries) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeries) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeries) SetSeriesInfo(v []*ListPreferredEcsTypesResponseBodySeriesSeriesInfo) *ListPreferredEcsTypesResponseBodySeries {
	s.SeriesInfo = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfo struct {
	Roles      *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
	SeriesId   *string                                                 `json:"SeriesId,omitempty" xml:"SeriesId,omitempty"`
	SeriesName *string                                                 `json:"SeriesName,omitempty" xml:"SeriesName,omitempty"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfo) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetRoles(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.Roles = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetSeriesId(v string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.SeriesId = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfo) SetSeriesName(v string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfo {
	s.SeriesName = &v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles struct {
	Compute *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute `json:"Compute,omitempty" xml:"Compute,omitempty" type:"Struct"`
	Login   *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin   `json:"Login,omitempty" xml:"Login,omitempty" type:"Struct"`
	Manager *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetCompute(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Compute = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetLogin(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Login = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles) SetManager(v *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRoles {
	s.Manager = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesCompute {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesLogin {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager struct {
	InstanceTypeId []*string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager) SetInstanceTypeId(v []*string) *ListPreferredEcsTypesResponseBodySeriesSeriesInfoRolesManager {
	s.InstanceTypeId = v
	return s
}

type ListPreferredEcsTypesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPreferredEcsTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPreferredEcsTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponse) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponse) SetHeaders(v map[string]*string) *ListPreferredEcsTypesResponse {
	s.Headers = v
	return s
}

func (s *ListPreferredEcsTypesResponse) SetStatusCode(v int32) *ListPreferredEcsTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPreferredEcsTypesResponse) SetBody(v *ListPreferredEcsTypesResponseBody) *ListPreferredEcsTypesResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	Regions   *ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v *ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	RegionInfo []*ListRegionsResponseBodyRegionsRegionInfo `json:"RegionInfo,omitempty" xml:"RegionInfo,omitempty" type:"Repeated"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetRegionInfo(v []*ListRegionsResponseBodyRegionsRegionInfo) *ListRegionsResponseBodyRegions {
	s.RegionInfo = v
	return s
}

type ListRegionsResponseBodyRegionsRegionInfo struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegionsRegionInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegionsRegionInfo) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegionsRegionInfo) SetLocalName(v string) *ListRegionsResponseBodyRegionsRegionInfo {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegionsRegionInfo) SetRegionId(v string) *ListRegionsResponseBodyRegionsRegionInfo {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSoftwaresRequest struct {
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
}

func (s ListSoftwaresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwaresRequest) SetEhpcVersion(v string) *ListSoftwaresRequest {
	s.EhpcVersion = &v
	return s
}

type ListSoftwaresResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Softwares *ListSoftwaresResponseBodySoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Struct"`
}

func (s ListSoftwaresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBody) SetRequestId(v string) *ListSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetSoftwares(v *ListSoftwaresResponseBodySoftwares) *ListSoftwaresResponseBody {
	s.Softwares = v
	return s
}

type ListSoftwaresResponseBodySoftwares struct {
	SoftwareInfo []*ListSoftwaresResponseBodySoftwaresSoftwareInfo `json:"SoftwareInfo,omitempty" xml:"SoftwareInfo,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodySoftwares) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwares) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwares) SetSoftwareInfo(v []*ListSoftwaresResponseBodySoftwaresSoftwareInfo) *ListSoftwaresResponseBodySoftwares {
	s.SoftwareInfo = v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfo struct {
	AccountType      *string                                                     `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountVersion   *string                                                     `json:"AccountVersion,omitempty" xml:"AccountVersion,omitempty"`
	Applications     *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	EhpcVersion      *string                                                     `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	OsTag            *string                                                     `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	SchedulerType    *string                                                     `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	SchedulerVersion *string                                                     `json:"SchedulerVersion,omitempty" xml:"SchedulerVersion,omitempty"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfo) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetAccountType(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.AccountType = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetAccountVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.AccountVersion = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetApplications(v *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.Applications = v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetEhpcVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.EhpcVersion = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetOsTag(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.OsTag = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetSchedulerType(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.SchedulerType = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfo) SetSchedulerVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfo {
	s.SchedulerVersion = &v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications struct {
	ApplicationInfo []*ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications) SetApplicationInfo(v []*ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplications {
	s.ApplicationInfo = v
	return s
}

type ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	Tag      *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetName(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Name = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetRequired(v bool) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Required = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetTag(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Tag = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo) SetVersion(v string) *ListSoftwaresResponseBodySoftwaresSoftwareInfoApplicationsApplicationInfo {
	s.Version = &v
	return s
}

type ListSoftwaresResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSoftwaresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponse) SetHeaders(v map[string]*string) *ListSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *ListSoftwaresResponse) SetStatusCode(v int32) *ListSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSoftwaresResponse) SetBody(v *ListSoftwaresResponseBody) *ListSoftwaresResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetClusterId(v string) *ListUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	PageNumber *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	UserInfo []*ListUsersResponseBodyUsersUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetUserInfo(v []*ListUsersResponseBodyUsersUserInfo) *ListUsersResponseBodyUsers {
	s.UserInfo = v
	return s
}

type ListUsersResponseBodyUsersUserInfo struct {
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUsersResponseBodyUsersUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsersUserInfo) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUserInfo) SetAddTime(v string) *ListUsersResponseBodyUsersUserInfo {
	s.AddTime = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetGroup(v string) *ListUsersResponseBodyUsersUserInfo {
	s.Group = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserInfo) SetName(v string) *ListUsersResponseBodyUsersUserInfo {
	s.Name = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListVolumesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesRequest) GoString() string {
	return s.String()
}

func (s *ListVolumesRequest) SetPageNumber(v int32) *ListVolumesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVolumesRequest) SetPageSize(v int32) *ListVolumesRequest {
	s.PageSize = &v
	return s
}

type ListVolumesResponseBody struct {
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Volumes    *ListVolumesResponseBodyVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Struct"`
}

func (s ListVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBody) SetPageNumber(v int32) *ListVolumesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVolumesResponseBody) SetPageSize(v int32) *ListVolumesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVolumesResponseBody) SetRequestId(v string) *ListVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVolumesResponseBody) SetTotalCount(v int32) *ListVolumesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVolumesResponseBody) SetVolumes(v *ListVolumesResponseBodyVolumes) *ListVolumesResponseBody {
	s.Volumes = v
	return s
}

type ListVolumesResponseBodyVolumes struct {
	VolumeInfo []*ListVolumesResponseBodyVolumesVolumeInfo `json:"VolumeInfo,omitempty" xml:"VolumeInfo,omitempty" type:"Repeated"`
}

func (s ListVolumesResponseBodyVolumes) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumes) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumes) SetVolumeInfo(v []*ListVolumesResponseBodyVolumesVolumeInfo) *ListVolumesResponseBodyVolumes {
	s.VolumeInfo = v
	return s
}

type ListVolumesResponseBodyVolumesVolumeInfo struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName      *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemoteDirectory  *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	VolumeId         *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	VolumeProtocol   *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeType       *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s ListVolumesResponseBodyVolumesVolumeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumesVolumeInfo) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetClusterId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.ClusterId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetClusterName(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.ClusterName = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetRegionId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.RegionId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetRemoteDirectory(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeId(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeMountpoint(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeProtocol(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesVolumeInfo) SetVolumeType(v string) *ListVolumesResponseBodyVolumesVolumeInfo {
	s.VolumeType = &v
	return s
}

type ListVolumesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponse) GoString() string {
	return s.String()
}

func (s *ListVolumesResponse) SetHeaders(v map[string]*string) *ListVolumesResponse {
	s.Headers = v
	return s
}

func (s *ListVolumesResponse) SetStatusCode(v int32) *ListVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVolumesResponse) SetBody(v *ListVolumesResponseBody) *ListVolumesResponse {
	s.Body = v
	return s
}

type ModifyClusterAttributesRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyClusterAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesRequest) SetClusterId(v string) *ModifyClusterAttributesRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetDescription(v string) *ModifyClusterAttributesRequest {
	s.Description = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetName(v string) *ModifyClusterAttributesRequest {
	s.Name = &v
	return s
}

type ModifyClusterAttributesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesResponseBody) SetRequestId(v string) *ModifyClusterAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterAttributesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyClusterAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesResponse) SetHeaders(v map[string]*string) *ModifyClusterAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterAttributesResponse) SetStatusCode(v int32) *ModifyClusterAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterAttributesResponse) SetBody(v *ModifyClusterAttributesResponseBody) *ModifyClusterAttributesResponse {
	s.Body = v
	return s
}

type ModifyUserGroupsRequest struct {
	ClusterId *string                        `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*ModifyUserGroupsRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ModifyUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsRequest) SetClusterId(v string) *ModifyUserGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserGroupsRequest) SetUser(v []*ModifyUserGroupsRequestUser) *ModifyUserGroupsRequest {
	s.User = v
	return s
}

type ModifyUserGroupsRequestUser struct {
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyUserGroupsRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsRequestUser) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsRequestUser) SetGroup(v string) *ModifyUserGroupsRequestUser {
	s.Group = &v
	return s
}

func (s *ModifyUserGroupsRequestUser) SetName(v string) *ModifyUserGroupsRequestUser {
	s.Name = &v
	return s
}

type ModifyUserGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsResponseBody) SetRequestId(v string) *ModifyUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsResponse) SetHeaders(v map[string]*string) *ModifyUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserGroupsResponse) SetStatusCode(v int32) *ModifyUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserGroupsResponse) SetBody(v *ModifyUserGroupsResponseBody) *ModifyUserGroupsResponse {
	s.Body = v
	return s
}

type ModifyUserPasswordsRequest struct {
	ClusterId *string                           `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*ModifyUserPasswordsRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ModifyUserPasswordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsRequest) SetClusterId(v string) *ModifyUserPasswordsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserPasswordsRequest) SetUser(v []*ModifyUserPasswordsRequestUser) *ModifyUserPasswordsRequest {
	s.User = v
	return s
}

type ModifyUserPasswordsRequestUser struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ModifyUserPasswordsRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsRequestUser) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsRequestUser) SetName(v string) *ModifyUserPasswordsRequestUser {
	s.Name = &v
	return s
}

func (s *ModifyUserPasswordsRequestUser) SetPassword(v string) *ModifyUserPasswordsRequestUser {
	s.Password = &v
	return s
}

type ModifyUserPasswordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserPasswordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsResponseBody) SetRequestId(v string) *ModifyUserPasswordsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserPasswordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyUserPasswordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserPasswordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsResponse) SetHeaders(v map[string]*string) *ModifyUserPasswordsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserPasswordsResponse) SetStatusCode(v int32) *ModifyUserPasswordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserPasswordsResponse) SetBody(v *ModifyUserPasswordsResponseBody) *ModifyUserPasswordsResponse {
	s.Body = v
	return s
}

type RerunJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s RerunJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsRequest) GoString() string {
	return s.String()
}

func (s *RerunJobsRequest) SetClusterId(v string) *RerunJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *RerunJobsRequest) SetJobs(v string) *RerunJobsRequest {
	s.Jobs = &v
	return s
}

type RerunJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RerunJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsResponseBody) GoString() string {
	return s.String()
}

func (s *RerunJobsResponseBody) SetRequestId(v string) *RerunJobsResponseBody {
	s.RequestId = &v
	return s
}

type RerunJobsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RerunJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RerunJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsResponse) GoString() string {
	return s.String()
}

func (s *RerunJobsResponse) SetHeaders(v map[string]*string) *RerunJobsResponse {
	s.Headers = v
	return s
}

func (s *RerunJobsResponse) SetStatusCode(v int32) *RerunJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunJobsResponse) SetBody(v *RerunJobsResponseBody) *RerunJobsResponse {
	s.Body = v
	return s
}

type ResetNodesRequest struct {
	ClusterId *string                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*ResetNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ResetNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesRequest) GoString() string {
	return s.String()
}

func (s *ResetNodesRequest) SetClusterId(v string) *ResetNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ResetNodesRequest) SetInstance(v []*ResetNodesRequestInstance) *ResetNodesRequest {
	s.Instance = v
	return s
}

type ResetNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ResetNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *ResetNodesRequestInstance) SetId(v string) *ResetNodesRequestInstance {
	s.Id = &v
	return s
}

type ResetNodesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNodesResponseBody) SetRequestId(v string) *ResetNodesResponseBody {
	s.RequestId = &v
	return s
}

type ResetNodesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesResponse) GoString() string {
	return s.String()
}

func (s *ResetNodesResponse) SetHeaders(v map[string]*string) *ResetNodesResponse {
	s.Headers = v
	return s
}

func (s *ResetNodesResponse) SetStatusCode(v int32) *ResetNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetNodesResponse) SetBody(v *ResetNodesResponseBody) *ResetNodesResponse {
	s.Body = v
	return s
}

type SetAutoScaleConfigRequest struct {
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	EnableAutoGrow          *bool   `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	EnableAutoShrink        *bool   `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	ExcludeNodes            *string `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	ExtraNodesGrowRatio     *int32  `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	GrowIntervalInMinutes   *int32  `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	GrowRatio               *int32  `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	GrowTimeoutInMinutes    *int32  `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	MaxNodesInCluster       *int32  `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	ShrinkIdleTimes         *int32  `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	ShrinkIntervalInMinutes *int32  `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
}

func (s SetAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequest) SetClusterId(v string) *SetAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetEnableAutoGrow(v bool) *SetAutoScaleConfigRequest {
	s.EnableAutoGrow = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetEnableAutoShrink(v bool) *SetAutoScaleConfigRequest {
	s.EnableAutoShrink = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetExcludeNodes(v string) *SetAutoScaleConfigRequest {
	s.ExcludeNodes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetExtraNodesGrowRatio(v int32) *SetAutoScaleConfigRequest {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowIntervalInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowRatio(v int32) *SetAutoScaleConfigRequest {
	s.GrowRatio = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowTimeoutInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetMaxNodesInCluster(v int32) *SetAutoScaleConfigRequest {
	s.MaxNodesInCluster = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetShrinkIdleTimes(v int32) *SetAutoScaleConfigRequest {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetShrinkIntervalInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.ShrinkIntervalInMinutes = &v
	return s
}

type SetAutoScaleConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigResponseBody) SetRequestId(v string) *SetAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetAutoScaleConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigResponse) SetHeaders(v map[string]*string) *SetAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *SetAutoScaleConfigResponse) SetStatusCode(v int32) *SetAutoScaleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAutoScaleConfigResponse) SetBody(v *SetAutoScaleConfigResponseBody) *SetAutoScaleConfigResponse {
	s.Body = v
	return s
}

type SetJobUserRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RunasUser         *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
}

func (s SetJobUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SetJobUserRequest) GoString() string {
	return s.String()
}

func (s *SetJobUserRequest) SetClusterId(v string) *SetJobUserRequest {
	s.ClusterId = &v
	return s
}

func (s *SetJobUserRequest) SetRunasUser(v string) *SetJobUserRequest {
	s.RunasUser = &v
	return s
}

func (s *SetJobUserRequest) SetRunasUserPassword(v string) *SetJobUserRequest {
	s.RunasUserPassword = &v
	return s
}

type SetJobUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetJobUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetJobUserResponseBody) GoString() string {
	return s.String()
}

func (s *SetJobUserResponseBody) SetRequestId(v string) *SetJobUserResponseBody {
	s.RequestId = &v
	return s
}

type SetJobUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetJobUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetJobUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SetJobUserResponse) GoString() string {
	return s.String()
}

func (s *SetJobUserResponse) SetHeaders(v map[string]*string) *SetJobUserResponse {
	s.Headers = v
	return s
}

func (s *SetJobUserResponse) SetStatusCode(v int32) *SetJobUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SetJobUserResponse) SetBody(v *SetJobUserResponseBody) *SetJobUserResponse {
	s.Body = v
	return s
}

type StopJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s StopJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobsRequest) GoString() string {
	return s.String()
}

func (s *StopJobsRequest) SetClusterId(v string) *StopJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *StopJobsRequest) SetJobs(v string) *StopJobsRequest {
	s.Jobs = &v
	return s
}

type StopJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobsResponseBody) SetRequestId(v string) *StopJobsResponseBody {
	s.RequestId = &v
	return s
}

type StopJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponse) GoString() string {
	return s.String()
}

func (s *StopJobsResponse) SetHeaders(v map[string]*string) *StopJobsResponse {
	s.Headers = v
	return s
}

func (s *StopJobsResponse) SetStatusCode(v int32) *StopJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobsResponse) SetBody(v *StopJobsResponseBody) *StopJobsResponse {
	s.Body = v
	return s
}

type SubmitJobRequest struct {
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword  *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s SubmitJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitJobRequest) SetArrayRequest(v string) *SubmitJobRequest {
	s.ArrayRequest = &v
	return s
}

func (s *SubmitJobRequest) SetClusterId(v string) *SubmitJobRequest {
	s.ClusterId = &v
	return s
}

func (s *SubmitJobRequest) SetCommandLine(v string) *SubmitJobRequest {
	s.CommandLine = &v
	return s
}

func (s *SubmitJobRequest) SetName(v string) *SubmitJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitJobRequest) SetPackagePath(v string) *SubmitJobRequest {
	s.PackagePath = &v
	return s
}

func (s *SubmitJobRequest) SetPriority(v int32) *SubmitJobRequest {
	s.Priority = &v
	return s
}

func (s *SubmitJobRequest) SetReRunable(v bool) *SubmitJobRequest {
	s.ReRunable = &v
	return s
}

func (s *SubmitJobRequest) SetRunasUser(v string) *SubmitJobRequest {
	s.RunasUser = &v
	return s
}

func (s *SubmitJobRequest) SetRunasUserPassword(v string) *SubmitJobRequest {
	s.RunasUserPassword = &v
	return s
}

func (s *SubmitJobRequest) SetStderrRedirectPath(v string) *SubmitJobRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *SubmitJobRequest) SetStdoutRedirectPath(v string) *SubmitJobRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *SubmitJobRequest) SetVariables(v string) *SubmitJobRequest {
	s.Variables = &v
	return s
}

type SubmitJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitJobResponseBody) SetJobId(v string) *SubmitJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitJobResponseBody) SetRequestId(v string) *SubmitJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitJobResponse) SetHeaders(v map[string]*string) *SubmitJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitJobResponse) SetStatusCode(v int32) *SubmitJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitJobResponse) SetBody(v *SubmitJobResponseBody) *SubmitJobResponse {
	s.Body = v
	return s
}

type UpgradeClientRequest struct {
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpgradeClientRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClientRequest) SetClientVersion(v string) *UpgradeClientRequest {
	s.ClientVersion = &v
	return s
}

func (s *UpgradeClientRequest) SetClusterId(v string) *UpgradeClientRequest {
	s.ClusterId = &v
	return s
}

type UpgradeClientResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponseBody) SetRequestId(v string) *UpgradeClientResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeClientResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeClientResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponse) SetHeaders(v map[string]*string) *UpgradeClientResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClientResponse) SetStatusCode(v int32) *UpgradeClientResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeClientResponse) SetBody(v *UpgradeClientResponseBody) *UpgradeClientResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ehpc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddNodesWithOptions(request *AddNodesRequest, runtime *util.RuntimeOptions) (_result *AddNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddNodes"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddNodes(request *AddNodesRequest) (_result *AddNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddNodesResponse{}
	_body, _err := client.AddNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUsersWithOptions(request *AddUsersRequest, runtime *util.RuntimeOptions) (_result *AddUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUsers"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUsers(request *AddUsersRequest) (_result *AddUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUsersResponse{}
	_body, _err := client.AddUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobTemplateWithOptions(request *CreateJobTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateJobTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJobTemplate"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobTemplate(request *CreateJobTemplateRequest) (_result *CreateJobTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobTemplateResponse{}
	_body, _err := client.CreateJobTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobTemplatesWithOptions(request *DeleteJobTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteJobTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobTemplates"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobTemplates(request *DeleteJobTemplatesRequest) (_result *DeleteJobTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobTemplatesResponse{}
	_body, _err := client.DeleteJobTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobsWithOptions(request *DeleteJobsRequest, runtime *util.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobs"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobs(request *DeleteJobsRequest) (_result *DeleteJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobsResponse{}
	_body, _err := client.DeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNodesWithOptions(request *DeleteNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNodes"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNodes(request *DeleteNodesRequest) (_result *DeleteNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodesResponse{}
	_body, _err := client.DeleteNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUsersWithOptions(request *DeleteUsersRequest, runtime *util.RuntimeOptions) (_result *DeleteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUsers"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUsers(request *DeleteUsersRequest) (_result *DeleteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUsersResponse{}
	_body, _err := client.DeleteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCluster"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditJobTemplateWithOptions(request *EditJobTemplateRequest, runtime *util.RuntimeOptions) (_result *EditJobTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditJobTemplate"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditJobTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditJobTemplate(request *EditJobTemplateRequest) (_result *EditJobTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditJobTemplateResponse{}
	_body, _err := client.EditJobTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAutoScaleConfigWithOptions(request *GetAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *GetAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutoScaleConfig"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAutoScaleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAutoScaleConfig(request *GetAutoScaleConfigRequest) (_result *GetAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoScaleConfigResponse{}
	_body, _err := client.GetAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterLogsWithOptions(request *ListClusterLogsRequest, runtime *util.RuntimeOptions) (_result *ListClusterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterLogs"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterLogs(request *ListClusterLogsRequest) (_result *ListClusterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterLogsResponse{}
	_body, _err := client.ListClusterLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCurrentClientVersionWithOptions(runtime *util.RuntimeOptions) (_result *ListCurrentClientVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListCurrentClientVersion"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCurrentClientVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCurrentClientVersion() (_result *ListCurrentClientVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCurrentClientVersionResponse{}
	_body, _err := client.ListCurrentClientVersionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomImagesWithOptions(request *ListCustomImagesRequest, runtime *util.RuntimeOptions) (_result *ListCustomImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomImages"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomImages(request *ListCustomImagesRequest) (_result *ListCustomImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.ListCustomImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages() (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobTemplatesWithOptions(request *ListJobTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListJobTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobTemplates"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobTemplates(request *ListJobTemplatesRequest) (_result *ListJobTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobTemplatesResponse{}
	_body, _err := client.ListJobTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(request *ListJobsRequest, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesNoPagingWithOptions(request *ListNodesNoPagingRequest, runtime *util.RuntimeOptions) (_result *ListNodesNoPagingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodesNoPaging"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesNoPagingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesNoPaging(request *ListNodesNoPagingRequest) (_result *ListNodesNoPagingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesNoPagingResponse{}
	_body, _err := client.ListNodesNoPagingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPreferredEcsTypesWithOptions(request *ListPreferredEcsTypesRequest, runtime *util.RuntimeOptions) (_result *ListPreferredEcsTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPreferredEcsTypes"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPreferredEcsTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPreferredEcsTypes(request *ListPreferredEcsTypesRequest) (_result *ListPreferredEcsTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPreferredEcsTypesResponse{}
	_body, _err := client.ListPreferredEcsTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSoftwaresWithOptions(request *ListSoftwaresRequest, runtime *util.RuntimeOptions) (_result *ListSoftwaresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSoftwares"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSoftwares(request *ListSoftwaresRequest) (_result *ListSoftwaresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.ListSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVolumesWithOptions(request *ListVolumesRequest, runtime *util.RuntimeOptions) (_result *ListVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVolumes"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVolumes(request *ListVolumesRequest) (_result *ListVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVolumesResponse{}
	_body, _err := client.ListVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterAttributesWithOptions(request *ModifyClusterAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyClusterAttributes"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyClusterAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterAttributes(request *ModifyClusterAttributesRequest) (_result *ModifyClusterAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterAttributesResponse{}
	_body, _err := client.ModifyClusterAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserGroupsWithOptions(request *ModifyUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ModifyUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserGroups"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserGroups(request *ModifyUserGroupsRequest) (_result *ModifyUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserGroupsResponse{}
	_body, _err := client.ModifyUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserPasswordsWithOptions(request *ModifyUserPasswordsRequest, runtime *util.RuntimeOptions) (_result *ModifyUserPasswordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserPasswords"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserPasswordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserPasswords(request *ModifyUserPasswordsRequest) (_result *ModifyUserPasswordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserPasswordsResponse{}
	_body, _err := client.ModifyUserPasswordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RerunJobsWithOptions(request *RerunJobsRequest, runtime *util.RuntimeOptions) (_result *RerunJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RerunJobs"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RerunJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RerunJobs(request *RerunJobsRequest) (_result *RerunJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RerunJobsResponse{}
	_body, _err := client.RerunJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetNodesWithOptions(request *ResetNodesRequest, runtime *util.RuntimeOptions) (_result *ResetNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetNodes"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetNodes(request *ResetNodesRequest) (_result *ResetNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetNodesResponse{}
	_body, _err := client.ResetNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAutoScaleConfigWithOptions(request *SetAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *SetAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAutoScaleConfig"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAutoScaleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAutoScaleConfig(request *SetAutoScaleConfigRequest) (_result *SetAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAutoScaleConfigResponse{}
	_body, _err := client.SetAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetJobUserWithOptions(request *SetJobUserRequest, runtime *util.RuntimeOptions) (_result *SetJobUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetJobUser"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetJobUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetJobUser(request *SetJobUserRequest) (_result *SetJobUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetJobUserResponse{}
	_body, _err := client.SetJobUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobsWithOptions(request *StopJobsRequest, runtime *util.RuntimeOptions) (_result *StopJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopJobs"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJobs(request *StopJobsRequest) (_result *StopJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopJobsResponse{}
	_body, _err := client.StopJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitJobWithOptions(request *SubmitJobRequest, runtime *util.RuntimeOptions) (_result *SubmitJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitJob"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitJob(request *SubmitJobRequest) (_result *SubmitJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitJobResponse{}
	_body, _err := client.SubmitJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClientWithOptions(request *UpgradeClientRequest, runtime *util.RuntimeOptions) (_result *UpgradeClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeClient"),
		Version:     tea.String("2017-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeClient(request *UpgradeClientRequest) (_result *UpgradeClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClientResponse{}
	_body, _err := client.UpgradeClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
