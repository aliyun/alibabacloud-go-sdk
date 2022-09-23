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

type CreateAccessGroupRequest struct {
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InputRegionId   *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s CreateAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupRequest) SetAccessGroupName(v string) *CreateAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessGroupRequest) SetDescription(v string) *CreateAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessGroupRequest) SetInputRegionId(v string) *CreateAccessGroupRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateAccessGroupRequest) SetNetworkType(v string) *CreateAccessGroupRequest {
	s.NetworkType = &v
	return s
}

type CreateAccessGroupResponseBody struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponseBody) SetAccessGroupId(v string) *CreateAccessGroupResponseBody {
	s.AccessGroupId = &v
	return s
}

func (s *CreateAccessGroupResponseBody) SetRequestId(v string) *CreateAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponse) SetHeaders(v map[string]*string) *CreateAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessGroupResponse) SetStatusCode(v int32) *CreateAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessGroupResponse) SetBody(v *CreateAccessGroupResponseBody) *CreateAccessGroupResponse {
	s.Body = v
	return s
}

type CreateAccessRuleRequest struct {
	AccessGroupId  *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InputRegionId  *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	NetworkSegment *string `json:"NetworkSegment,omitempty" xml:"NetworkSegment,omitempty"`
	Priority       *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RWAccessType   *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
}

func (s CreateAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleRequest) SetAccessGroupId(v string) *CreateAccessRuleRequest {
	s.AccessGroupId = &v
	return s
}

func (s *CreateAccessRuleRequest) SetDescription(v string) *CreateAccessRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessRuleRequest) SetInputRegionId(v string) *CreateAccessRuleRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateAccessRuleRequest) SetNetworkSegment(v string) *CreateAccessRuleRequest {
	s.NetworkSegment = &v
	return s
}

func (s *CreateAccessRuleRequest) SetPriority(v int32) *CreateAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateAccessRuleRequest) SetRWAccessType(v string) *CreateAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

type CreateAccessRuleResponseBody struct {
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleResponseBody) SetAccessRuleId(v string) *CreateAccessRuleResponseBody {
	s.AccessRuleId = &v
	return s
}

func (s *CreateAccessRuleResponseBody) SetRequestId(v string) *CreateAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleResponse) SetHeaders(v map[string]*string) *CreateAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessRuleResponse) SetStatusCode(v int32) *CreateAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessRuleResponse) SetBody(v *CreateAccessRuleResponseBody) *CreateAccessRuleResponse {
	s.Body = v
	return s
}

type CreateFileSystemRequest struct {
	DataRedundancyType           *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemName               *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	InputRegionId                *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	PartitionNumber              *int32  `json:"PartitionNumber,omitempty" xml:"PartitionNumber,omitempty"`
	ProtocolType                 *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ProvisionedThroughputInMiBps *int64  `json:"ProvisionedThroughputInMiBps,omitempty" xml:"ProvisionedThroughputInMiBps,omitempty"`
	SpaceCapacity                *int64  `json:"SpaceCapacity,omitempty" xml:"SpaceCapacity,omitempty"`
	StorageSetName               *string `json:"StorageSetName,omitempty" xml:"StorageSetName,omitempty"`
	StorageType                  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ThroughputMode               *string `json:"ThroughputMode,omitempty" xml:"ThroughputMode,omitempty"`
	ZoneId                       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequest) SetDataRedundancyType(v string) *CreateFileSystemRequest {
	s.DataRedundancyType = &v
	return s
}

func (s *CreateFileSystemRequest) SetDescription(v string) *CreateFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateFileSystemRequest) SetFileSystemName(v string) *CreateFileSystemRequest {
	s.FileSystemName = &v
	return s
}

func (s *CreateFileSystemRequest) SetInputRegionId(v string) *CreateFileSystemRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateFileSystemRequest) SetPartitionNumber(v int32) *CreateFileSystemRequest {
	s.PartitionNumber = &v
	return s
}

func (s *CreateFileSystemRequest) SetProtocolType(v string) *CreateFileSystemRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateFileSystemRequest) SetProvisionedThroughputInMiBps(v int64) *CreateFileSystemRequest {
	s.ProvisionedThroughputInMiBps = &v
	return s
}

func (s *CreateFileSystemRequest) SetSpaceCapacity(v int64) *CreateFileSystemRequest {
	s.SpaceCapacity = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageSetName(v string) *CreateFileSystemRequest {
	s.StorageSetName = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageType(v string) *CreateFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateFileSystemRequest) SetThroughputMode(v string) *CreateFileSystemRequest {
	s.ThroughputMode = &v
	return s
}

func (s *CreateFileSystemRequest) SetZoneId(v string) *CreateFileSystemRequest {
	s.ZoneId = &v
	return s
}

type CreateFileSystemResponseBody struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponseBody) SetFileSystemId(v string) *CreateFileSystemResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *CreateFileSystemResponseBody) SetRequestId(v string) *CreateFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type CreateFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponse) SetHeaders(v map[string]*string) *CreateFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CreateFileSystemResponse) SetStatusCode(v int32) *CreateFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileSystemResponse) SetBody(v *CreateFileSystemResponseBody) *CreateFileSystemResponse {
	s.Body = v
	return s
}

type CreateMountPointRequest struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	NetworkType   *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VSwitchId     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateMountPointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMountPointRequest) GoString() string {
	return s.String()
}

func (s *CreateMountPointRequest) SetAccessGroupId(v string) *CreateMountPointRequest {
	s.AccessGroupId = &v
	return s
}

func (s *CreateMountPointRequest) SetDescription(v string) *CreateMountPointRequest {
	s.Description = &v
	return s
}

func (s *CreateMountPointRequest) SetFileSystemId(v string) *CreateMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateMountPointRequest) SetInputRegionId(v string) *CreateMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateMountPointRequest) SetNetworkType(v string) *CreateMountPointRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateMountPointRequest) SetVSwitchId(v string) *CreateMountPointRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMountPointRequest) SetVpcId(v string) *CreateMountPointRequest {
	s.VpcId = &v
	return s
}

type CreateMountPointResponseBody struct {
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMountPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMountPointResponseBody) SetMountPointId(v string) *CreateMountPointResponseBody {
	s.MountPointId = &v
	return s
}

func (s *CreateMountPointResponseBody) SetRequestId(v string) *CreateMountPointResponseBody {
	s.RequestId = &v
	return s
}

type CreateMountPointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMountPointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMountPointResponse) GoString() string {
	return s.String()
}

func (s *CreateMountPointResponse) SetHeaders(v map[string]*string) *CreateMountPointResponse {
	s.Headers = v
	return s
}

func (s *CreateMountPointResponse) SetStatusCode(v int32) *CreateMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMountPointResponse) SetBody(v *CreateMountPointResponseBody) *CreateMountPointResponse {
	s.Body = v
	return s
}

type CreateUserGroupsMappingRequest struct {
	FileSystemId  *string                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	GroupNames    map[string]interface{} `json:"GroupNames,omitempty" xml:"GroupNames,omitempty"`
	InputRegionId *string                `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	UserName      *string                `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserGroupsMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupsMappingRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupsMappingRequest) SetFileSystemId(v string) *CreateUserGroupsMappingRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateUserGroupsMappingRequest) SetGroupNames(v map[string]interface{}) *CreateUserGroupsMappingRequest {
	s.GroupNames = v
	return s
}

func (s *CreateUserGroupsMappingRequest) SetInputRegionId(v string) *CreateUserGroupsMappingRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateUserGroupsMappingRequest) SetUserName(v string) *CreateUserGroupsMappingRequest {
	s.UserName = &v
	return s
}

type CreateUserGroupsMappingShrinkRequest struct {
	FileSystemId     *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	GroupNamesShrink *string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty"`
	InputRegionId    *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserGroupsMappingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupsMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupsMappingShrinkRequest) SetFileSystemId(v string) *CreateUserGroupsMappingShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateUserGroupsMappingShrinkRequest) SetGroupNamesShrink(v string) *CreateUserGroupsMappingShrinkRequest {
	s.GroupNamesShrink = &v
	return s
}

func (s *CreateUserGroupsMappingShrinkRequest) SetInputRegionId(v string) *CreateUserGroupsMappingShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateUserGroupsMappingShrinkRequest) SetUserName(v string) *CreateUserGroupsMappingShrinkRequest {
	s.UserName = &v
	return s
}

type CreateUserGroupsMappingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserGroupsMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupsMappingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupsMappingResponseBody) SetRequestId(v string) *CreateUserGroupsMappingResponseBody {
	s.RequestId = &v
	return s
}

type CreateUserGroupsMappingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserGroupsMappingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserGroupsMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupsMappingResponse) GoString() string {
	return s.String()
}

func (s *CreateUserGroupsMappingResponse) SetHeaders(v map[string]*string) *CreateUserGroupsMappingResponse {
	s.Headers = v
	return s
}

func (s *CreateUserGroupsMappingResponse) SetStatusCode(v int32) *CreateUserGroupsMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserGroupsMappingResponse) SetBody(v *CreateUserGroupsMappingResponseBody) *CreateUserGroupsMappingResponse {
	s.Body = v
	return s
}

type DeleteAccessGroupRequest struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s DeleteAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupRequest) SetAccessGroupId(v string) *DeleteAccessGroupRequest {
	s.AccessGroupId = &v
	return s
}

func (s *DeleteAccessGroupRequest) SetInputRegionId(v string) *DeleteAccessGroupRequest {
	s.InputRegionId = &v
	return s
}

type DeleteAccessGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupResponseBody) SetRequestId(v string) *DeleteAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupResponse) SetHeaders(v map[string]*string) *DeleteAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessGroupResponse) SetStatusCode(v int32) *DeleteAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessGroupResponse) SetBody(v *DeleteAccessGroupResponseBody) *DeleteAccessGroupResponse {
	s.Body = v
	return s
}

type DeleteAccessRuleRequest struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	AccessRuleId  *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s DeleteAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleRequest) SetAccessGroupId(v string) *DeleteAccessRuleRequest {
	s.AccessGroupId = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetAccessRuleId(v string) *DeleteAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetInputRegionId(v string) *DeleteAccessRuleRequest {
	s.InputRegionId = &v
	return s
}

type DeleteAccessRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleResponseBody) SetRequestId(v string) *DeleteAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleResponse) SetHeaders(v map[string]*string) *DeleteAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessRuleResponse) SetStatusCode(v int32) *DeleteAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessRuleResponse) SetBody(v *DeleteAccessRuleResponseBody) *DeleteAccessRuleResponse {
	s.Body = v
	return s
}

type DeleteFileSystemRequest struct {
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s DeleteFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemRequest) SetFileSystemId(v string) *DeleteFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFileSystemRequest) SetInputRegionId(v string) *DeleteFileSystemRequest {
	s.InputRegionId = &v
	return s
}

type DeleteFileSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemResponseBody) SetRequestId(v string) *DeleteFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileSystemResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileSystemResponse) SetHeaders(v map[string]*string) *DeleteFileSystemResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileSystemResponse) SetStatusCode(v int32) *DeleteFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileSystemResponse) SetBody(v *DeleteFileSystemResponseBody) *DeleteFileSystemResponse {
	s.Body = v
	return s
}

type DeleteMountPointRequest struct {
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	MountPointId  *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
}

func (s DeleteMountPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteMountPointRequest) SetFileSystemId(v string) *DeleteMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteMountPointRequest) SetInputRegionId(v string) *DeleteMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteMountPointRequest) SetMountPointId(v string) *DeleteMountPointRequest {
	s.MountPointId = &v
	return s
}

type DeleteMountPointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMountPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMountPointResponseBody) SetRequestId(v string) *DeleteMountPointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMountPointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMountPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMountPointResponse) GoString() string {
	return s.String()
}

func (s *DeleteMountPointResponse) SetHeaders(v map[string]*string) *DeleteMountPointResponse {
	s.Headers = v
	return s
}

func (s *DeleteMountPointResponse) SetStatusCode(v int32) *DeleteMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMountPointResponse) SetBody(v *DeleteMountPointResponseBody) *DeleteMountPointResponse {
	s.Body = v
	return s
}

type DeleteUserGroupsMappingRequest struct {
	FileSystemId  *string                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	GroupNames    map[string]interface{} `json:"GroupNames,omitempty" xml:"GroupNames,omitempty"`
	InputRegionId *string                `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	UserName      *string                `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteUserGroupsMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupsMappingRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupsMappingRequest) SetFileSystemId(v string) *DeleteUserGroupsMappingRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteUserGroupsMappingRequest) SetGroupNames(v map[string]interface{}) *DeleteUserGroupsMappingRequest {
	s.GroupNames = v
	return s
}

func (s *DeleteUserGroupsMappingRequest) SetInputRegionId(v string) *DeleteUserGroupsMappingRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteUserGroupsMappingRequest) SetUserName(v string) *DeleteUserGroupsMappingRequest {
	s.UserName = &v
	return s
}

type DeleteUserGroupsMappingShrinkRequest struct {
	FileSystemId     *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	GroupNamesShrink *string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty"`
	InputRegionId    *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteUserGroupsMappingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupsMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupsMappingShrinkRequest) SetFileSystemId(v string) *DeleteUserGroupsMappingShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteUserGroupsMappingShrinkRequest) SetGroupNamesShrink(v string) *DeleteUserGroupsMappingShrinkRequest {
	s.GroupNamesShrink = &v
	return s
}

func (s *DeleteUserGroupsMappingShrinkRequest) SetInputRegionId(v string) *DeleteUserGroupsMappingShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteUserGroupsMappingShrinkRequest) SetUserName(v string) *DeleteUserGroupsMappingShrinkRequest {
	s.UserName = &v
	return s
}

type DeleteUserGroupsMappingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserGroupsMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupsMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupsMappingResponseBody) SetRequestId(v string) *DeleteUserGroupsMappingResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserGroupsMappingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserGroupsMappingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserGroupsMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupsMappingResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupsMappingResponse) SetHeaders(v map[string]*string) *DeleteUserGroupsMappingResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupsMappingResponse) SetStatusCode(v int32) *DeleteUserGroupsMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupsMappingResponse) SetBody(v *DeleteUserGroupsMappingResponseBody) *DeleteUserGroupsMappingResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	InputRegionId  *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetInputRegionId(v string) *DescribeRegionsRequest {
	s.InputRegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
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

type GetAccessGroupRequest struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s GetAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *GetAccessGroupRequest) SetAccessGroupId(v string) *GetAccessGroupRequest {
	s.AccessGroupId = &v
	return s
}

func (s *GetAccessGroupRequest) SetInputRegionId(v string) *GetAccessGroupRequest {
	s.InputRegionId = &v
	return s
}

type GetAccessGroupResponseBody struct {
	AccessGroup *GetAccessGroupResponseBodyAccessGroup `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty" type:"Struct"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessGroupResponseBody) SetAccessGroup(v *GetAccessGroupResponseBodyAccessGroup) *GetAccessGroupResponseBody {
	s.AccessGroup = v
	return s
}

func (s *GetAccessGroupResponseBody) SetRequestId(v string) *GetAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessGroupResponseBodyAccessGroup struct {
	AccessGroupId   *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsDefault       *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	MountPointCount *int32  `json:"MountPointCount,omitempty" xml:"MountPointCount,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleCount       *int32  `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s GetAccessGroupResponseBodyAccessGroup) String() string {
	return tea.Prettify(s)
}

func (s GetAccessGroupResponseBodyAccessGroup) GoString() string {
	return s.String()
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetAccessGroupId(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.AccessGroupId = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetAccessGroupName(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.AccessGroupName = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetCreateTime(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.CreateTime = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetDescription(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.Description = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetIsDefault(v bool) *GetAccessGroupResponseBodyAccessGroup {
	s.IsDefault = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetMountPointCount(v int32) *GetAccessGroupResponseBodyAccessGroup {
	s.MountPointCount = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetNetworkType(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.NetworkType = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetRegionId(v string) *GetAccessGroupResponseBodyAccessGroup {
	s.RegionId = &v
	return s
}

func (s *GetAccessGroupResponseBodyAccessGroup) SetRuleCount(v int32) *GetAccessGroupResponseBodyAccessGroup {
	s.RuleCount = &v
	return s
}

type GetAccessGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *GetAccessGroupResponse) SetHeaders(v map[string]*string) *GetAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *GetAccessGroupResponse) SetStatusCode(v int32) *GetAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessGroupResponse) SetBody(v *GetAccessGroupResponseBody) *GetAccessGroupResponse {
	s.Body = v
	return s
}

type GetAccessRuleRequest struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	AccessRuleId  *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s GetAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAccessRuleRequest) SetAccessGroupId(v string) *GetAccessRuleRequest {
	s.AccessGroupId = &v
	return s
}

func (s *GetAccessRuleRequest) SetAccessRuleId(v string) *GetAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *GetAccessRuleRequest) SetInputRegionId(v string) *GetAccessRuleRequest {
	s.InputRegionId = &v
	return s
}

type GetAccessRuleResponseBody struct {
	AccessRule *GetAccessRuleResponseBodyAccessRule `json:"AccessRule,omitempty" xml:"AccessRule,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessRuleResponseBody) SetAccessRule(v *GetAccessRuleResponseBodyAccessRule) *GetAccessRuleResponseBody {
	s.AccessRule = v
	return s
}

func (s *GetAccessRuleResponseBody) SetRequestId(v string) *GetAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessRuleResponseBodyAccessRule struct {
	AccessGroupId  *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	AccessRuleId   *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NetworkSegment *string `json:"NetworkSegment,omitempty" xml:"NetworkSegment,omitempty"`
	Priority       *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RWAccessType   *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAccessRuleResponseBodyAccessRule) String() string {
	return tea.Prettify(s)
}

func (s GetAccessRuleResponseBodyAccessRule) GoString() string {
	return s.String()
}

func (s *GetAccessRuleResponseBodyAccessRule) SetAccessGroupId(v string) *GetAccessRuleResponseBodyAccessRule {
	s.AccessGroupId = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetAccessRuleId(v string) *GetAccessRuleResponseBodyAccessRule {
	s.AccessRuleId = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetCreateTime(v string) *GetAccessRuleResponseBodyAccessRule {
	s.CreateTime = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetDescription(v string) *GetAccessRuleResponseBodyAccessRule {
	s.Description = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetNetworkSegment(v string) *GetAccessRuleResponseBodyAccessRule {
	s.NetworkSegment = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetPriority(v int32) *GetAccessRuleResponseBodyAccessRule {
	s.Priority = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetRWAccessType(v string) *GetAccessRuleResponseBodyAccessRule {
	s.RWAccessType = &v
	return s
}

func (s *GetAccessRuleResponseBodyAccessRule) SetRegionId(v string) *GetAccessRuleResponseBodyAccessRule {
	s.RegionId = &v
	return s
}

type GetAccessRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAccessRuleResponse) SetHeaders(v map[string]*string) *GetAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAccessRuleResponse) SetStatusCode(v int32) *GetAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessRuleResponse) SetBody(v *GetAccessRuleResponseBody) *GetAccessRuleResponse {
	s.Body = v
	return s
}

type GetFileSystemRequest struct {
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s GetFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileSystemRequest) GoString() string {
	return s.String()
}

func (s *GetFileSystemRequest) SetFileSystemId(v string) *GetFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetFileSystemRequest) SetInputRegionId(v string) *GetFileSystemRequest {
	s.InputRegionId = &v
	return s
}

type GetFileSystemResponseBody struct {
	FileSystem *GetFileSystemResponseBodyFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileSystemResponseBody) SetFileSystem(v *GetFileSystemResponseBodyFileSystem) *GetFileSystemResponseBody {
	s.FileSystem = v
	return s
}

func (s *GetFileSystemResponseBody) SetRequestId(v string) *GetFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type GetFileSystemResponseBodyFileSystem struct {
	CreateTime                   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description                  *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId                 *string  `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemName               *string  `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	MeteringSpaceSize            *float32 `json:"MeteringSpaceSize,omitempty" xml:"MeteringSpaceSize,omitempty"`
	MountPointCount              *int64   `json:"MountPointCount,omitempty" xml:"MountPointCount,omitempty"`
	NumberOfDirectories          *int64   `json:"NumberOfDirectories,omitempty" xml:"NumberOfDirectories,omitempty"`
	NumberOfFiles                *int64   `json:"NumberOfFiles,omitempty" xml:"NumberOfFiles,omitempty"`
	ProtocolType                 *string  `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ProvisionedThroughputInMiBps *int64   `json:"ProvisionedThroughputInMiBps,omitempty" xml:"ProvisionedThroughputInMiBps,omitempty"`
	RegionId                     *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpaceCapacity                *int64   `json:"SpaceCapacity,omitempty" xml:"SpaceCapacity,omitempty"`
	StoragePackageId             *string  `json:"StoragePackageId,omitempty" xml:"StoragePackageId,omitempty"`
	StorageType                  *string  `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ThroughputMode               *string  `json:"ThroughputMode,omitempty" xml:"ThroughputMode,omitempty"`
	UsedSpaceSize                *float32 `json:"UsedSpaceSize,omitempty" xml:"UsedSpaceSize,omitempty"`
	Version                      *string  `json:"Version,omitempty" xml:"Version,omitempty"`
	ZoneId                       *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetFileSystemResponseBodyFileSystem) String() string {
	return tea.Prettify(s)
}

func (s GetFileSystemResponseBodyFileSystem) GoString() string {
	return s.String()
}

func (s *GetFileSystemResponseBodyFileSystem) SetCreateTime(v string) *GetFileSystemResponseBodyFileSystem {
	s.CreateTime = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetDescription(v string) *GetFileSystemResponseBodyFileSystem {
	s.Description = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetFileSystemId(v string) *GetFileSystemResponseBodyFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetFileSystemName(v string) *GetFileSystemResponseBodyFileSystem {
	s.FileSystemName = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetMeteringSpaceSize(v float32) *GetFileSystemResponseBodyFileSystem {
	s.MeteringSpaceSize = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetMountPointCount(v int64) *GetFileSystemResponseBodyFileSystem {
	s.MountPointCount = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetNumberOfDirectories(v int64) *GetFileSystemResponseBodyFileSystem {
	s.NumberOfDirectories = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetNumberOfFiles(v int64) *GetFileSystemResponseBodyFileSystem {
	s.NumberOfFiles = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetProtocolType(v string) *GetFileSystemResponseBodyFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetProvisionedThroughputInMiBps(v int64) *GetFileSystemResponseBodyFileSystem {
	s.ProvisionedThroughputInMiBps = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetRegionId(v string) *GetFileSystemResponseBodyFileSystem {
	s.RegionId = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetSpaceCapacity(v int64) *GetFileSystemResponseBodyFileSystem {
	s.SpaceCapacity = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetStoragePackageId(v string) *GetFileSystemResponseBodyFileSystem {
	s.StoragePackageId = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetStorageType(v string) *GetFileSystemResponseBodyFileSystem {
	s.StorageType = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetThroughputMode(v string) *GetFileSystemResponseBodyFileSystem {
	s.ThroughputMode = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetUsedSpaceSize(v float32) *GetFileSystemResponseBodyFileSystem {
	s.UsedSpaceSize = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetVersion(v string) *GetFileSystemResponseBodyFileSystem {
	s.Version = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetZoneId(v string) *GetFileSystemResponseBodyFileSystem {
	s.ZoneId = &v
	return s
}

type GetFileSystemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileSystemResponse) GoString() string {
	return s.String()
}

func (s *GetFileSystemResponse) SetHeaders(v map[string]*string) *GetFileSystemResponse {
	s.Headers = v
	return s
}

func (s *GetFileSystemResponse) SetStatusCode(v int32) *GetFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileSystemResponse) SetBody(v *GetFileSystemResponseBody) *GetFileSystemResponse {
	s.Body = v
	return s
}

type GetMountPointRequest struct {
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	MountPointId  *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
}

func (s GetMountPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMountPointRequest) GoString() string {
	return s.String()
}

func (s *GetMountPointRequest) SetFileSystemId(v string) *GetMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetMountPointRequest) SetInputRegionId(v string) *GetMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *GetMountPointRequest) SetMountPointId(v string) *GetMountPointRequest {
	s.MountPointId = &v
	return s
}

type GetMountPointResponseBody struct {
	MountPoint *GetMountPointResponseBodyMountPoint `json:"MountPoint,omitempty" xml:"MountPoint,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMountPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetMountPointResponseBody) SetMountPoint(v *GetMountPointResponseBodyMountPoint) *GetMountPointResponseBody {
	s.MountPoint = v
	return s
}

func (s *GetMountPointResponseBody) SetRequestId(v string) *GetMountPointResponseBody {
	s.RequestId = &v
	return s
}

type GetMountPointResponseBodyMountPoint struct {
	AccessGroupId    *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId     *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountPointDomain *string `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	MountPointId     *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetMountPointResponseBodyMountPoint) String() string {
	return tea.Prettify(s)
}

func (s GetMountPointResponseBodyMountPoint) GoString() string {
	return s.String()
}

func (s *GetMountPointResponseBodyMountPoint) SetAccessGroupId(v string) *GetMountPointResponseBodyMountPoint {
	s.AccessGroupId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetCreateTime(v string) *GetMountPointResponseBodyMountPoint {
	s.CreateTime = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetDescription(v string) *GetMountPointResponseBodyMountPoint {
	s.Description = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetFileSystemId(v string) *GetMountPointResponseBodyMountPoint {
	s.FileSystemId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetMountPointDomain(v string) *GetMountPointResponseBodyMountPoint {
	s.MountPointDomain = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetMountPointId(v string) *GetMountPointResponseBodyMountPoint {
	s.MountPointId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetNetworkType(v string) *GetMountPointResponseBodyMountPoint {
	s.NetworkType = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetRegionId(v string) *GetMountPointResponseBodyMountPoint {
	s.RegionId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetStatus(v string) *GetMountPointResponseBodyMountPoint {
	s.Status = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetVSwitchId(v string) *GetMountPointResponseBodyMountPoint {
	s.VSwitchId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetVpcId(v string) *GetMountPointResponseBodyMountPoint {
	s.VpcId = &v
	return s
}

type GetMountPointResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMountPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMountPointResponse) GoString() string {
	return s.String()
}

func (s *GetMountPointResponse) SetHeaders(v map[string]*string) *GetMountPointResponse {
	s.Headers = v
	return s
}

func (s *GetMountPointResponse) SetStatusCode(v int32) *GetMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMountPointResponse) SetBody(v *GetMountPointResponseBody) *GetMountPointResponse {
	s.Body = v
	return s
}

type GetRegionRequest struct {
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s GetRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegionRequest) GoString() string {
	return s.String()
}

func (s *GetRegionRequest) SetInputRegionId(v string) *GetRegionRequest {
	s.InputRegionId = &v
	return s
}

type GetRegionResponseBody struct {
	AvailableZones []*GetRegionResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionResponseBody) SetAvailableZones(v []*GetRegionResponseBodyAvailableZones) *GetRegionResponseBody {
	s.AvailableZones = v
	return s
}

func (s *GetRegionResponseBody) SetRequestId(v string) *GetRegionResponseBody {
	s.RequestId = &v
	return s
}

type GetRegionResponseBodyAvailableZones struct {
	Options []*GetRegionResponseBodyAvailableZonesOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	ZoneId  *string                                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetRegionResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s GetRegionResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *GetRegionResponseBodyAvailableZones) SetOptions(v []*GetRegionResponseBodyAvailableZonesOptions) *GetRegionResponseBodyAvailableZones {
	s.Options = v
	return s
}

func (s *GetRegionResponseBodyAvailableZones) SetZoneId(v string) *GetRegionResponseBodyAvailableZones {
	s.ZoneId = &v
	return s
}

type GetRegionResponseBodyAvailableZonesOptions struct {
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s GetRegionResponseBodyAvailableZonesOptions) String() string {
	return tea.Prettify(s)
}

func (s GetRegionResponseBodyAvailableZonesOptions) GoString() string {
	return s.String()
}

func (s *GetRegionResponseBodyAvailableZonesOptions) SetProtocolType(v string) *GetRegionResponseBodyAvailableZonesOptions {
	s.ProtocolType = &v
	return s
}

func (s *GetRegionResponseBodyAvailableZonesOptions) SetStorageType(v string) *GetRegionResponseBodyAvailableZonesOptions {
	s.StorageType = &v
	return s
}

type GetRegionResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegionResponse) GoString() string {
	return s.String()
}

func (s *GetRegionResponse) SetHeaders(v map[string]*string) *GetRegionResponse {
	s.Headers = v
	return s
}

func (s *GetRegionResponse) SetStatusCode(v int32) *GetRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionResponse) SetBody(v *GetRegionResponseBody) *GetRegionResponse {
	s.Body = v
	return s
}

type ListAccessGroupsRequest struct {
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	Limit         *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OrderBy       *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	StartOffset   *int32  `json:"StartOffset,omitempty" xml:"StartOffset,omitempty"`
}

func (s ListAccessGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessGroupsRequest) SetInputRegionId(v string) *ListAccessGroupsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListAccessGroupsRequest) SetLimit(v int32) *ListAccessGroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListAccessGroupsRequest) SetOrderBy(v string) *ListAccessGroupsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAccessGroupsRequest) SetOrderType(v string) *ListAccessGroupsRequest {
	s.OrderType = &v
	return s
}

func (s *ListAccessGroupsRequest) SetStartOffset(v int32) *ListAccessGroupsRequest {
	s.StartOffset = &v
	return s
}

type ListAccessGroupsResponseBody struct {
	AccessGroups []*ListAccessGroupsResponseBodyAccessGroups `json:"AccessGroups,omitempty" xml:"AccessGroups,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccessGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessGroupsResponseBody) SetAccessGroups(v []*ListAccessGroupsResponseBodyAccessGroups) *ListAccessGroupsResponseBody {
	s.AccessGroups = v
	return s
}

func (s *ListAccessGroupsResponseBody) SetRequestId(v string) *ListAccessGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessGroupsResponseBody) SetTotalCount(v int32) *ListAccessGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAccessGroupsResponseBodyAccessGroups struct {
	AccessGroupId   *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsDefault       *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	MountPointCount *int32  `json:"MountPointCount,omitempty" xml:"MountPointCount,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleCount       *int32  `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s ListAccessGroupsResponseBodyAccessGroups) String() string {
	return tea.Prettify(s)
}

func (s ListAccessGroupsResponseBodyAccessGroups) GoString() string {
	return s.String()
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetAccessGroupId(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.AccessGroupId = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetAccessGroupName(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.AccessGroupName = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetCreateTime(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.CreateTime = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetDescription(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.Description = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetIsDefault(v bool) *ListAccessGroupsResponseBodyAccessGroups {
	s.IsDefault = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetMountPointCount(v int32) *ListAccessGroupsResponseBodyAccessGroups {
	s.MountPointCount = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetNetworkType(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.NetworkType = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetRegionId(v string) *ListAccessGroupsResponseBodyAccessGroups {
	s.RegionId = &v
	return s
}

func (s *ListAccessGroupsResponseBodyAccessGroups) SetRuleCount(v int32) *ListAccessGroupsResponseBodyAccessGroups {
	s.RuleCount = &v
	return s
}

type ListAccessGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccessGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccessGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessGroupsResponse) SetHeaders(v map[string]*string) *ListAccessGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessGroupsResponse) SetStatusCode(v int32) *ListAccessGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessGroupsResponse) SetBody(v *ListAccessGroupsResponseBody) *ListAccessGroupsResponse {
	s.Body = v
	return s
}

type ListAccessRulesRequest struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	Limit         *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OrderBy       *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	StartOffset   *int32  `json:"StartOffset,omitempty" xml:"StartOffset,omitempty"`
}

func (s ListAccessRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAccessRulesRequest) SetAccessGroupId(v string) *ListAccessRulesRequest {
	s.AccessGroupId = &v
	return s
}

func (s *ListAccessRulesRequest) SetInputRegionId(v string) *ListAccessRulesRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListAccessRulesRequest) SetLimit(v int32) *ListAccessRulesRequest {
	s.Limit = &v
	return s
}

func (s *ListAccessRulesRequest) SetOrderBy(v string) *ListAccessRulesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAccessRulesRequest) SetOrderType(v string) *ListAccessRulesRequest {
	s.OrderType = &v
	return s
}

func (s *ListAccessRulesRequest) SetStartOffset(v int32) *ListAccessRulesRequest {
	s.StartOffset = &v
	return s
}

type ListAccessRulesResponseBody struct {
	AccessRules []*ListAccessRulesResponseBodyAccessRules `json:"AccessRules,omitempty" xml:"AccessRules,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccessRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessRulesResponseBody) SetAccessRules(v []*ListAccessRulesResponseBodyAccessRules) *ListAccessRulesResponseBody {
	s.AccessRules = v
	return s
}

func (s *ListAccessRulesResponseBody) SetRequestId(v string) *ListAccessRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessRulesResponseBody) SetTotalCount(v int32) *ListAccessRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAccessRulesResponseBodyAccessRules struct {
	AccessGroupId  *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	AccessRuleId   *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NetworkSegment *string `json:"NetworkSegment,omitempty" xml:"NetworkSegment,omitempty"`
	Priority       *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RWAccessType   *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAccessRulesResponseBodyAccessRules) String() string {
	return tea.Prettify(s)
}

func (s ListAccessRulesResponseBodyAccessRules) GoString() string {
	return s.String()
}

func (s *ListAccessRulesResponseBodyAccessRules) SetAccessGroupId(v string) *ListAccessRulesResponseBodyAccessRules {
	s.AccessGroupId = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetAccessRuleId(v string) *ListAccessRulesResponseBodyAccessRules {
	s.AccessRuleId = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetCreateTime(v string) *ListAccessRulesResponseBodyAccessRules {
	s.CreateTime = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetDescription(v string) *ListAccessRulesResponseBodyAccessRules {
	s.Description = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetNetworkSegment(v string) *ListAccessRulesResponseBodyAccessRules {
	s.NetworkSegment = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetPriority(v int32) *ListAccessRulesResponseBodyAccessRules {
	s.Priority = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetRWAccessType(v string) *ListAccessRulesResponseBodyAccessRules {
	s.RWAccessType = &v
	return s
}

func (s *ListAccessRulesResponseBodyAccessRules) SetRegionId(v string) *ListAccessRulesResponseBodyAccessRules {
	s.RegionId = &v
	return s
}

type ListAccessRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccessRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccessRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAccessRulesResponse) SetHeaders(v map[string]*string) *ListAccessRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAccessRulesResponse) SetStatusCode(v int32) *ListAccessRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessRulesResponse) SetBody(v *ListAccessRulesResponseBody) *ListAccessRulesResponse {
	s.Body = v
	return s
}

type ListFileSystemsRequest struct {
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	Limit         *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OrderBy       *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	StartOffset   *int32  `json:"StartOffset,omitempty" xml:"StartOffset,omitempty"`
}

func (s ListFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *ListFileSystemsRequest) SetInputRegionId(v string) *ListFileSystemsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListFileSystemsRequest) SetLimit(v int32) *ListFileSystemsRequest {
	s.Limit = &v
	return s
}

func (s *ListFileSystemsRequest) SetOrderBy(v string) *ListFileSystemsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListFileSystemsRequest) SetOrderType(v string) *ListFileSystemsRequest {
	s.OrderType = &v
	return s
}

func (s *ListFileSystemsRequest) SetStartOffset(v int32) *ListFileSystemsRequest {
	s.StartOffset = &v
	return s
}

type ListFileSystemsResponseBody struct {
	FileSystems []*ListFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileSystemsResponseBody) SetFileSystems(v []*ListFileSystemsResponseBodyFileSystems) *ListFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *ListFileSystemsResponseBody) SetRequestId(v string) *ListFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileSystemsResponseBody) SetTotalCount(v int32) *ListFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListFileSystemsResponseBodyFileSystems struct {
	CreateTime                   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description                  *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId                 *string  `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemName               *string  `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	MeteringSpaceSize            *float32 `json:"MeteringSpaceSize,omitempty" xml:"MeteringSpaceSize,omitempty"`
	MountPointCount              *int64   `json:"MountPointCount,omitempty" xml:"MountPointCount,omitempty"`
	NumberOfDirectories          *int64   `json:"NumberOfDirectories,omitempty" xml:"NumberOfDirectories,omitempty"`
	NumberOfFiles                *int64   `json:"NumberOfFiles,omitempty" xml:"NumberOfFiles,omitempty"`
	ProtocolType                 *string  `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ProvisionedThroughputInMiBps *int64   `json:"ProvisionedThroughputInMiBps,omitempty" xml:"ProvisionedThroughputInMiBps,omitempty"`
	RegionId                     *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpaceCapacity                *int64   `json:"SpaceCapacity,omitempty" xml:"SpaceCapacity,omitempty"`
	StoragePackageId             *string  `json:"StoragePackageId,omitempty" xml:"StoragePackageId,omitempty"`
	StorageType                  *string  `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ThroughputMode               *string  `json:"ThroughputMode,omitempty" xml:"ThroughputMode,omitempty"`
	UsedSpaceSize                *float32 `json:"UsedSpaceSize,omitempty" xml:"UsedSpaceSize,omitempty"`
	Version                      *string  `json:"Version,omitempty" xml:"Version,omitempty"`
	ZoneId                       *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListFileSystemsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *ListFileSystemsResponseBodyFileSystems) SetCreateTime(v string) *ListFileSystemsResponseBodyFileSystems {
	s.CreateTime = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetDescription(v string) *ListFileSystemsResponseBodyFileSystems {
	s.Description = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetFileSystemId(v string) *ListFileSystemsResponseBodyFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetFileSystemName(v string) *ListFileSystemsResponseBodyFileSystems {
	s.FileSystemName = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetMeteringSpaceSize(v float32) *ListFileSystemsResponseBodyFileSystems {
	s.MeteringSpaceSize = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetMountPointCount(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.MountPointCount = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetNumberOfDirectories(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.NumberOfDirectories = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetNumberOfFiles(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.NumberOfFiles = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetProtocolType(v string) *ListFileSystemsResponseBodyFileSystems {
	s.ProtocolType = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetProvisionedThroughputInMiBps(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.ProvisionedThroughputInMiBps = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetRegionId(v string) *ListFileSystemsResponseBodyFileSystems {
	s.RegionId = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetSpaceCapacity(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.SpaceCapacity = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetStoragePackageId(v string) *ListFileSystemsResponseBodyFileSystems {
	s.StoragePackageId = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetStorageType(v string) *ListFileSystemsResponseBodyFileSystems {
	s.StorageType = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetThroughputMode(v string) *ListFileSystemsResponseBodyFileSystems {
	s.ThroughputMode = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetUsedSpaceSize(v float32) *ListFileSystemsResponseBodyFileSystems {
	s.UsedSpaceSize = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetVersion(v string) *ListFileSystemsResponseBodyFileSystems {
	s.Version = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetZoneId(v string) *ListFileSystemsResponseBodyFileSystems {
	s.ZoneId = &v
	return s
}

type ListFileSystemsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *ListFileSystemsResponse) SetHeaders(v map[string]*string) *ListFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *ListFileSystemsResponse) SetStatusCode(v int32) *ListFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileSystemsResponse) SetBody(v *ListFileSystemsResponseBody) *ListFileSystemsResponse {
	s.Body = v
	return s
}

type ListMountPointsRequest struct {
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	Limit         *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OrderBy       *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	StartOffset   *int32  `json:"StartOffset,omitempty" xml:"StartOffset,omitempty"`
}

func (s ListMountPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMountPointsRequest) GoString() string {
	return s.String()
}

func (s *ListMountPointsRequest) SetFileSystemId(v string) *ListMountPointsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListMountPointsRequest) SetInputRegionId(v string) *ListMountPointsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListMountPointsRequest) SetLimit(v int32) *ListMountPointsRequest {
	s.Limit = &v
	return s
}

func (s *ListMountPointsRequest) SetOrderBy(v string) *ListMountPointsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMountPointsRequest) SetOrderType(v string) *ListMountPointsRequest {
	s.OrderType = &v
	return s
}

func (s *ListMountPointsRequest) SetStartOffset(v int32) *ListMountPointsRequest {
	s.StartOffset = &v
	return s
}

type ListMountPointsResponseBody struct {
	MountPoints []*ListMountPointsResponseBodyMountPoints `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMountPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMountPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMountPointsResponseBody) SetMountPoints(v []*ListMountPointsResponseBodyMountPoints) *ListMountPointsResponseBody {
	s.MountPoints = v
	return s
}

func (s *ListMountPointsResponseBody) SetRequestId(v string) *ListMountPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMountPointsResponseBody) SetTotalCount(v int32) *ListMountPointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListMountPointsResponseBodyMountPoints struct {
	AccessGroupId    *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId     *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountPointDomain *string `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	MountPointId     *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListMountPointsResponseBodyMountPoints) String() string {
	return tea.Prettify(s)
}

func (s ListMountPointsResponseBodyMountPoints) GoString() string {
	return s.String()
}

func (s *ListMountPointsResponseBodyMountPoints) SetAccessGroupId(v string) *ListMountPointsResponseBodyMountPoints {
	s.AccessGroupId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetCreateTime(v string) *ListMountPointsResponseBodyMountPoints {
	s.CreateTime = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetDescription(v string) *ListMountPointsResponseBodyMountPoints {
	s.Description = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetFileSystemId(v string) *ListMountPointsResponseBodyMountPoints {
	s.FileSystemId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetMountPointDomain(v string) *ListMountPointsResponseBodyMountPoints {
	s.MountPointDomain = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetMountPointId(v string) *ListMountPointsResponseBodyMountPoints {
	s.MountPointId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetNetworkType(v string) *ListMountPointsResponseBodyMountPoints {
	s.NetworkType = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetRegionId(v string) *ListMountPointsResponseBodyMountPoints {
	s.RegionId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetStatus(v string) *ListMountPointsResponseBodyMountPoints {
	s.Status = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetVSwitchId(v string) *ListMountPointsResponseBodyMountPoints {
	s.VSwitchId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetVpcId(v string) *ListMountPointsResponseBodyMountPoints {
	s.VpcId = &v
	return s
}

type ListMountPointsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMountPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMountPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMountPointsResponse) GoString() string {
	return s.String()
}

func (s *ListMountPointsResponse) SetHeaders(v map[string]*string) *ListMountPointsResponse {
	s.Headers = v
	return s
}

func (s *ListMountPointsResponse) SetStatusCode(v int32) *ListMountPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMountPointsResponse) SetBody(v *ListMountPointsResponseBody) *ListMountPointsResponse {
	s.Body = v
	return s
}

type ListUserGroupsMappingsRequest struct {
	FilesystemId  *string `json:"FilesystemId,omitempty" xml:"FilesystemId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	Limit         *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListUserGroupsMappingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsMappingsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsMappingsRequest) SetFilesystemId(v string) *ListUserGroupsMappingsRequest {
	s.FilesystemId = &v
	return s
}

func (s *ListUserGroupsMappingsRequest) SetInputRegionId(v string) *ListUserGroupsMappingsRequest {
	s.InputRegionId = &v
	return s
}

func (s *ListUserGroupsMappingsRequest) SetLimit(v int32) *ListUserGroupsMappingsRequest {
	s.Limit = &v
	return s
}

func (s *ListUserGroupsMappingsRequest) SetNextToken(v string) *ListUserGroupsMappingsRequest {
	s.NextToken = &v
	return s
}

type ListUserGroupsMappingsResponseBody struct {
	HasMore            *bool                                                   `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	NextToken          *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroupsMappings []*ListUserGroupsMappingsResponseBodyUserGroupsMappings `json:"UserGroupsMappings,omitempty" xml:"UserGroupsMappings,omitempty" type:"Repeated"`
}

func (s ListUserGroupsMappingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsMappingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsMappingsResponseBody) SetHasMore(v bool) *ListUserGroupsMappingsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListUserGroupsMappingsResponseBody) SetNextToken(v string) *ListUserGroupsMappingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserGroupsMappingsResponseBody) SetRequestId(v string) *ListUserGroupsMappingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsMappingsResponseBody) SetUserGroupsMappings(v []*ListUserGroupsMappingsResponseBodyUserGroupsMappings) *ListUserGroupsMappingsResponseBody {
	s.UserGroupsMappings = v
	return s
}

type ListUserGroupsMappingsResponseBodyUserGroupsMappings struct {
	GroupNames []*string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty" type:"Repeated"`
	UserName   *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUserGroupsMappingsResponseBodyUserGroupsMappings) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsMappingsResponseBodyUserGroupsMappings) GoString() string {
	return s.String()
}

func (s *ListUserGroupsMappingsResponseBodyUserGroupsMappings) SetGroupNames(v []*string) *ListUserGroupsMappingsResponseBodyUserGroupsMappings {
	s.GroupNames = v
	return s
}

func (s *ListUserGroupsMappingsResponseBodyUserGroupsMappings) SetUserName(v string) *ListUserGroupsMappingsResponseBodyUserGroupsMappings {
	s.UserName = &v
	return s
}

type ListUserGroupsMappingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserGroupsMappingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserGroupsMappingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsMappingsResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsMappingsResponse) SetHeaders(v map[string]*string) *ListUserGroupsMappingsResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsMappingsResponse) SetStatusCode(v int32) *ListUserGroupsMappingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsMappingsResponse) SetBody(v *ListUserGroupsMappingsResponseBody) *ListUserGroupsMappingsResponse {
	s.Body = v
	return s
}

type ModifyAccessGroupRequest struct {
	AccessGroupId   *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InputRegionId   *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s ModifyAccessGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupRequest) SetAccessGroupId(v string) *ModifyAccessGroupRequest {
	s.AccessGroupId = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetAccessGroupName(v string) *ModifyAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetDescription(v string) *ModifyAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetInputRegionId(v string) *ModifyAccessGroupRequest {
	s.InputRegionId = &v
	return s
}

type ModifyAccessGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupResponseBody) SetRequestId(v string) *ModifyAccessGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccessGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupResponse) SetHeaders(v map[string]*string) *ModifyAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessGroupResponse) SetStatusCode(v int32) *ModifyAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessGroupResponse) SetBody(v *ModifyAccessGroupResponseBody) *ModifyAccessGroupResponse {
	s.Body = v
	return s
}

type ModifyAccessRuleRequest struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	AccessRuleId  *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	Priority      *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RWAccessType  *string `json:"RWAccessType,omitempty" xml:"RWAccessType,omitempty"`
}

func (s ModifyAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleRequest) SetAccessGroupId(v string) *ModifyAccessRuleRequest {
	s.AccessGroupId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetAccessRuleId(v string) *ModifyAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetDescription(v string) *ModifyAccessRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetInputRegionId(v string) *ModifyAccessRuleRequest {
	s.InputRegionId = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetPriority(v int32) *ModifyAccessRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifyAccessRuleRequest) SetRWAccessType(v string) *ModifyAccessRuleRequest {
	s.RWAccessType = &v
	return s
}

type ModifyAccessRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleResponseBody) SetRequestId(v string) *ModifyAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleResponse) SetHeaders(v map[string]*string) *ModifyAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessRuleResponse) SetStatusCode(v int32) *ModifyAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessRuleResponse) SetBody(v *ModifyAccessRuleResponseBody) *ModifyAccessRuleResponse {
	s.Body = v
	return s
}

type ModifyFileSystemRequest struct {
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId                 *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemName               *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	InputRegionId                *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	ProvisionedThroughputInMiBps *int64  `json:"ProvisionedThroughputInMiBps,omitempty" xml:"ProvisionedThroughputInMiBps,omitempty"`
	SpaceCapacity                *int64  `json:"SpaceCapacity,omitempty" xml:"SpaceCapacity,omitempty"`
	ThroughputMode               *string `json:"ThroughputMode,omitempty" xml:"ThroughputMode,omitempty"`
}

func (s ModifyFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequest) SetDescription(v string) *ModifyFileSystemRequest {
	s.Description = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemId(v string) *ModifyFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemName(v string) *ModifyFileSystemRequest {
	s.FileSystemName = &v
	return s
}

func (s *ModifyFileSystemRequest) SetInputRegionId(v string) *ModifyFileSystemRequest {
	s.InputRegionId = &v
	return s
}

func (s *ModifyFileSystemRequest) SetProvisionedThroughputInMiBps(v int64) *ModifyFileSystemRequest {
	s.ProvisionedThroughputInMiBps = &v
	return s
}

func (s *ModifyFileSystemRequest) SetSpaceCapacity(v int64) *ModifyFileSystemRequest {
	s.SpaceCapacity = &v
	return s
}

func (s *ModifyFileSystemRequest) SetThroughputMode(v string) *ModifyFileSystemRequest {
	s.ThroughputMode = &v
	return s
}

type ModifyFileSystemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemResponseBody) SetRequestId(v string) *ModifyFileSystemResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFileSystemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFileSystemResponse) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemResponse) SetHeaders(v map[string]*string) *ModifyFileSystemResponse {
	s.Headers = v
	return s
}

func (s *ModifyFileSystemResponse) SetStatusCode(v int32) *ModifyFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFileSystemResponse) SetBody(v *ModifyFileSystemResponseBody) *ModifyFileSystemResponse {
	s.Body = v
	return s
}

type ModifyMountPointRequest struct {
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	MountPointId  *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyMountPointRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountPointRequest) GoString() string {
	return s.String()
}

func (s *ModifyMountPointRequest) SetAccessGroupId(v string) *ModifyMountPointRequest {
	s.AccessGroupId = &v
	return s
}

func (s *ModifyMountPointRequest) SetDescription(v string) *ModifyMountPointRequest {
	s.Description = &v
	return s
}

func (s *ModifyMountPointRequest) SetFileSystemId(v string) *ModifyMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyMountPointRequest) SetInputRegionId(v string) *ModifyMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *ModifyMountPointRequest) SetMountPointId(v string) *ModifyMountPointRequest {
	s.MountPointId = &v
	return s
}

func (s *ModifyMountPointRequest) SetStatus(v string) *ModifyMountPointRequest {
	s.Status = &v
	return s
}

type ModifyMountPointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMountPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMountPointResponseBody) SetRequestId(v string) *ModifyMountPointResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMountPointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMountPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMountPointResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMountPointResponse) GoString() string {
	return s.String()
}

func (s *ModifyMountPointResponse) SetHeaders(v map[string]*string) *ModifyMountPointResponse {
	s.Headers = v
	return s
}

func (s *ModifyMountPointResponse) SetStatusCode(v int32) *ModifyMountPointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMountPointResponse) SetBody(v *ModifyMountPointResponseBody) *ModifyMountPointResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dfs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateAccessGroupWithOptions(request *CreateAccessGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessGroup"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (_result *CreateAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CreateAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccessRuleWithOptions(request *CreateAccessRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkSegment)) {
		query["NetworkSegment"] = request.NetworkSegment
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RWAccessType)) {
		query["RWAccessType"] = request.RWAccessType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessRule"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessRule(request *CreateAccessRuleRequest) (_result *CreateAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CreateAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFileSystemWithOptions(request *CreateFileSystemRequest, runtime *util.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataRedundancyType)) {
		query["DataRedundancyType"] = request.DataRedundancyType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemName)) {
		query["FileSystemName"] = request.FileSystemName
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNumber)) {
		query["PartitionNumber"] = request.PartitionNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolType)) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionedThroughputInMiBps)) {
		query["ProvisionedThroughputInMiBps"] = request.ProvisionedThroughputInMiBps
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceCapacity)) {
		query["SpaceCapacity"] = request.SpaceCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSetName)) {
		query["StorageSetName"] = request.StorageSetName
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		query["StorageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.ThroughputMode)) {
		query["ThroughputMode"] = request.ThroughputMode
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFileSystem"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFileSystem(request *CreateFileSystemRequest) (_result *CreateFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CreateFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMountPointWithOptions(request *CreateMountPointRequest, runtime *util.RuntimeOptions) (_result *CreateMountPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMountPoint"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMountPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMountPoint(request *CreateMountPointRequest) (_result *CreateMountPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMountPointResponse{}
	_body, _err := client.CreateMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserGroupsMappingWithOptions(tmpReq *CreateUserGroupsMappingRequest, runtime *util.RuntimeOptions) (_result *CreateUserGroupsMappingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateUserGroupsMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GroupNames)) {
		request.GroupNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupNames, tea.String("GroupNames"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserGroupsMapping"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserGroupsMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserGroupsMapping(request *CreateUserGroupsMappingRequest) (_result *CreateUserGroupsMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserGroupsMappingResponse{}
	_body, _err := client.CreateUserGroupsMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessGroupWithOptions(request *DeleteAccessGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessGroup"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessGroup(request *DeleteAccessGroupRequest) (_result *DeleteAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.DeleteAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessRuleWithOptions(request *DeleteAccessRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessRule"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessRule(request *DeleteAccessRuleRequest) (_result *DeleteAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.DeleteAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileSystemWithOptions(request *DeleteFileSystemRequest, runtime *util.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFileSystem"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (_result *DeleteFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.DeleteFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMountPointWithOptions(request *DeleteMountPointRequest, runtime *util.RuntimeOptions) (_result *DeleteMountPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.MountPointId)) {
		query["MountPointId"] = request.MountPointId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMountPoint"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMountPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMountPoint(request *DeleteMountPointRequest) (_result *DeleteMountPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMountPointResponse{}
	_body, _err := client.DeleteMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserGroupsMappingWithOptions(tmpReq *DeleteUserGroupsMappingRequest, runtime *util.RuntimeOptions) (_result *DeleteUserGroupsMappingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteUserGroupsMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GroupNames)) {
		request.GroupNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupNames, tea.String("GroupNames"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserGroupsMapping"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserGroupsMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserGroupsMapping(request *DeleteUserGroupsMappingRequest) (_result *DeleteUserGroupsMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserGroupsMappingResponse{}
	_body, _err := client.DeleteUserGroupsMappingWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-06-20"),
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

func (client *Client) GetAccessGroupWithOptions(request *GetAccessGroupRequest, runtime *util.RuntimeOptions) (_result *GetAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessGroup"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessGroup(request *GetAccessGroupRequest) (_result *GetAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessGroupResponse{}
	_body, _err := client.GetAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessRuleWithOptions(request *GetAccessRuleRequest, runtime *util.RuntimeOptions) (_result *GetAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessRule"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessRule(request *GetAccessRuleRequest) (_result *GetAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessRuleResponse{}
	_body, _err := client.GetAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileSystemWithOptions(request *GetFileSystemRequest, runtime *util.RuntimeOptions) (_result *GetFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileSystem"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileSystem(request *GetFileSystemRequest) (_result *GetFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFileSystemResponse{}
	_body, _err := client.GetFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMountPointWithOptions(request *GetMountPointRequest, runtime *util.RuntimeOptions) (_result *GetMountPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.MountPointId)) {
		query["MountPointId"] = request.MountPointId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMountPoint"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMountPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMountPoint(request *GetMountPointRequest) (_result *GetMountPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMountPointResponse{}
	_body, _err := client.GetMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRegionWithOptions(request *GetRegionRequest, runtime *util.RuntimeOptions) (_result *GetRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegion"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegion(request *GetRegionRequest) (_result *GetRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRegionResponse{}
	_body, _err := client.GetRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccessGroupsWithOptions(request *ListAccessGroupsRequest, runtime *util.RuntimeOptions) (_result *ListAccessGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.StartOffset)) {
		query["StartOffset"] = request.StartOffset
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessGroups"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccessGroups(request *ListAccessGroupsRequest) (_result *ListAccessGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessGroupsResponse{}
	_body, _err := client.ListAccessGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccessRulesWithOptions(request *ListAccessRulesRequest, runtime *util.RuntimeOptions) (_result *ListAccessRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.StartOffset)) {
		query["StartOffset"] = request.StartOffset
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessRules"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccessRules(request *ListAccessRulesRequest) (_result *ListAccessRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessRulesResponse{}
	_body, _err := client.ListAccessRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFileSystemsWithOptions(request *ListFileSystemsRequest, runtime *util.RuntimeOptions) (_result *ListFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.StartOffset)) {
		query["StartOffset"] = request.StartOffset
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFileSystems"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFileSystemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFileSystems(request *ListFileSystemsRequest) (_result *ListFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFileSystemsResponse{}
	_body, _err := client.ListFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMountPointsWithOptions(request *ListMountPointsRequest, runtime *util.RuntimeOptions) (_result *ListMountPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.StartOffset)) {
		query["StartOffset"] = request.StartOffset
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMountPoints"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMountPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMountPoints(request *ListMountPointsRequest) (_result *ListMountPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMountPointsResponse{}
	_body, _err := client.ListMountPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserGroupsMappingsWithOptions(request *ListUserGroupsMappingsRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsMappingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilesystemId)) {
		query["FilesystemId"] = request.FilesystemId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserGroupsMappings"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserGroupsMappingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserGroupsMappings(request *ListUserGroupsMappingsRequest) (_result *ListUserGroupsMappingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupsMappingsResponse{}
	_body, _err := client.ListUserGroupsMappingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccessGroupWithOptions(request *ModifyAccessGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyAccessGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessGroupName)) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccessGroup"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (_result *ModifyAccessGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.ModifyAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccessRuleWithOptions(request *ModifyAccessRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessRuleId)) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RWAccessType)) {
		query["RWAccessType"] = request.RWAccessType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAccessRule"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccessRule(request *ModifyAccessRuleRequest) (_result *ModifyAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.ModifyAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFileSystemWithOptions(request *ModifyFileSystemRequest, runtime *util.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemName)) {
		query["FileSystemName"] = request.FileSystemName
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionedThroughputInMiBps)) {
		query["ProvisionedThroughputInMiBps"] = request.ProvisionedThroughputInMiBps
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceCapacity)) {
		query["SpaceCapacity"] = request.SpaceCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.ThroughputMode)) {
		query["ThroughputMode"] = request.ThroughputMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFileSystem"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (_result *ModifyFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.ModifyFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMountPointWithOptions(request *ModifyMountPointRequest, runtime *util.RuntimeOptions) (_result *ModifyMountPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessGroupId)) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.InputRegionId)) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.MountPointId)) {
		query["MountPointId"] = request.MountPointId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMountPoint"),
		Version:     tea.String("2018-06-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMountPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMountPoint(request *ModifyMountPointRequest) (_result *ModifyMountPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMountPointResponse{}
	_body, _err := client.ModifyMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
