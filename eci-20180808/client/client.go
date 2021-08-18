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

type CreateImageCacheRequest struct {
	OwnerId                 *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string                                           `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64                                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount            *string                                           `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId                *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                  *string                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SecurityGroupId         *string                                           `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId               *string                                           `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ImageCacheName          *string                                           `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	EipInstanceId           *string                                           `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	ResourceGroupId         *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClientToken             *string                                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ImageCacheSize          *int32                                            `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	RetentionDays           *int32                                            `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	AutoMatchImageCache     *bool                                             `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	ImageRegistryCredential []*CreateImageCacheRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	Image                   []*string                                         `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	Tag                     []*CreateImageCacheRequestTag                     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Flash                   *bool                                             `json:"Flash,omitempty" xml:"Flash,omitempty"`
	AcrRegistryInfo         []*CreateImageCacheRequestAcrRegistryInfo         `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
}

func (s CreateImageCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequest) SetOwnerId(v int64) *CreateImageCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceOwnerAccount(v string) *CreateImageCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceOwnerId(v int64) *CreateImageCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateImageCacheRequest) SetOwnerAccount(v string) *CreateImageCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateImageCacheRequest) SetRegionId(v string) *CreateImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheRequest) SetZoneId(v string) *CreateImageCacheRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateImageCacheRequest) SetSecurityGroupId(v string) *CreateImageCacheRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateImageCacheRequest) SetVSwitchId(v string) *CreateImageCacheRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageCacheName(v string) *CreateImageCacheRequest {
	s.ImageCacheName = &v
	return s
}

func (s *CreateImageCacheRequest) SetEipInstanceId(v string) *CreateImageCacheRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceGroupId(v string) *CreateImageCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImageCacheRequest) SetClientToken(v string) *CreateImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageCacheSize(v int32) *CreateImageCacheRequest {
	s.ImageCacheSize = &v
	return s
}

func (s *CreateImageCacheRequest) SetRetentionDays(v int32) *CreateImageCacheRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateImageCacheRequest) SetAutoMatchImageCache(v bool) *CreateImageCacheRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageRegistryCredential(v []*CreateImageCacheRequestImageRegistryCredential) *CreateImageCacheRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *CreateImageCacheRequest) SetImage(v []*string) *CreateImageCacheRequest {
	s.Image = v
	return s
}

func (s *CreateImageCacheRequest) SetTag(v []*CreateImageCacheRequestTag) *CreateImageCacheRequest {
	s.Tag = v
	return s
}

func (s *CreateImageCacheRequest) SetFlash(v bool) *CreateImageCacheRequest {
	s.Flash = &v
	return s
}

func (s *CreateImageCacheRequest) SetAcrRegistryInfo(v []*CreateImageCacheRequestAcrRegistryInfo) *CreateImageCacheRequest {
	s.AcrRegistryInfo = v
	return s
}

type CreateImageCacheRequestImageRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateImageCacheRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetPassword(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetServer(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetUserName(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type CreateImageCacheRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageCacheRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestTag) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestTag) SetKey(v string) *CreateImageCacheRequestTag {
	s.Key = &v
	return s
}

func (s *CreateImageCacheRequestTag) SetValue(v string) *CreateImageCacheRequestTag {
	s.Value = &v
	return s
}

type CreateImageCacheRequestAcrRegistryInfo struct {
	Domain       []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateImageCacheRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetDomain(v []*string) *CreateImageCacheRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetInstanceName(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetInstanceId(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetRegionId(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type CreateImageCacheResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ImageCacheId     *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
}

func (s CreateImageCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponseBody) SetRequestId(v string) *CreateImageCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetImageCacheId(v string) *CreateImageCacheResponseBody {
	s.ImageCacheId = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetContainerGroupId(v string) *CreateImageCacheResponseBody {
	s.ContainerGroupId = &v
	return s
}

type CreateImageCacheResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponse) SetHeaders(v map[string]*string) *CreateImageCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateImageCacheResponse) SetBody(v *CreateImageCacheResponseBody) *CreateImageCacheResponse {
	s.Body = v
	return s
}

type DescribeContainerLogRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	ContainerName        *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tail                 *int32  `json:"Tail,omitempty" xml:"Tail,omitempty"`
	LastTime             *bool   `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	SinceSeconds         *int32  `json:"SinceSeconds,omitempty" xml:"SinceSeconds,omitempty"`
	LimitBytes           *int64  `json:"LimitBytes,omitempty" xml:"LimitBytes,omitempty"`
	Timestamps           *bool   `json:"Timestamps,omitempty" xml:"Timestamps,omitempty"`
}

func (s DescribeContainerLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogRequest) SetOwnerId(v int64) *DescribeContainerLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetResourceOwnerAccount(v string) *DescribeContainerLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerLogRequest) SetResourceOwnerId(v int64) *DescribeContainerLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetOwnerAccount(v string) *DescribeContainerLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerLogRequest) SetRegionId(v string) *DescribeContainerLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetContainerGroupId(v string) *DescribeContainerLogRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetContainerName(v string) *DescribeContainerLogRequest {
	s.ContainerName = &v
	return s
}

func (s *DescribeContainerLogRequest) SetStartTime(v string) *DescribeContainerLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerLogRequest) SetTail(v int32) *DescribeContainerLogRequest {
	s.Tail = &v
	return s
}

func (s *DescribeContainerLogRequest) SetLastTime(v bool) *DescribeContainerLogRequest {
	s.LastTime = &v
	return s
}

func (s *DescribeContainerLogRequest) SetSinceSeconds(v int32) *DescribeContainerLogRequest {
	s.SinceSeconds = &v
	return s
}

func (s *DescribeContainerLogRequest) SetLimitBytes(v int64) *DescribeContainerLogRequest {
	s.LimitBytes = &v
	return s
}

func (s *DescribeContainerLogRequest) SetTimestamps(v bool) *DescribeContainerLogRequest {
	s.Timestamps = &v
	return s
}

type DescribeContainerLogResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeContainerLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogResponseBody) SetRequestId(v string) *DescribeContainerLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerLogResponseBody) SetContainerName(v string) *DescribeContainerLogResponseBody {
	s.ContainerName = &v
	return s
}

func (s *DescribeContainerLogResponseBody) SetContent(v string) *DescribeContainerLogResponseBody {
	s.Content = &v
	return s
}

type DescribeContainerLogResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContainerLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogResponse) SetHeaders(v map[string]*string) *DescribeContainerLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerLogResponse) SetBody(v *DescribeContainerLogResponseBody) *DescribeContainerLogResponse {
	s.Body = v
	return s
}

type RestartContainerGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RestartContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupRequest) SetOwnerId(v int64) *RestartContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetResourceOwnerAccount(v string) *RestartContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartContainerGroupRequest) SetResourceOwnerId(v int64) *RestartContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetOwnerAccount(v string) *RestartContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartContainerGroupRequest) SetRegionId(v string) *RestartContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetContainerGroupId(v string) *RestartContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetClientToken(v string) *RestartContainerGroupRequest {
	s.ClientToken = &v
	return s
}

type RestartContainerGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupResponseBody) SetRequestId(v string) *RestartContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type RestartContainerGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupResponse) SetHeaders(v map[string]*string) *RestartContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *RestartContainerGroupResponse) SetBody(v *RestartContainerGroupResponseBody) *RestartContainerGroupResponse {
	s.Body = v
	return s
}

type DescribeImageCachesRequest struct {
	OwnerId              *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                          `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ImageCacheId         *string                          `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	ImageCacheName       *string                          `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	SnapshotId           *string                          `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Image                *string                          `json:"Image,omitempty" xml:"Image,omitempty"`
	ResourceGroupId      *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                  []*DescribeImageCachesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImageCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesRequest) SetOwnerId(v int64) *DescribeImageCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceOwnerAccount(v string) *DescribeImageCachesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceOwnerId(v int64) *DescribeImageCachesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetOwnerAccount(v string) *DescribeImageCachesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImageCachesRequest) SetRegionId(v string) *DescribeImageCachesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetImageCacheId(v string) *DescribeImageCachesRequest {
	s.ImageCacheId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetImageCacheName(v string) *DescribeImageCachesRequest {
	s.ImageCacheName = &v
	return s
}

func (s *DescribeImageCachesRequest) SetSnapshotId(v string) *DescribeImageCachesRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetImage(v string) *DescribeImageCachesRequest {
	s.Image = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceGroupId(v string) *DescribeImageCachesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetTag(v []*DescribeImageCachesRequestTag) *DescribeImageCachesRequest {
	s.Tag = v
	return s
}

type DescribeImageCachesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageCachesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesRequestTag) SetKey(v string) *DescribeImageCachesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeImageCachesRequestTag) SetValue(v string) *DescribeImageCachesRequestTag {
	s.Value = &v
	return s
}

type DescribeImageCachesResponseBody struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ImageCaches []*DescribeImageCachesResponseBodyImageCaches `json:"ImageCaches,omitempty" xml:"ImageCaches,omitempty" type:"Repeated"`
}

func (s DescribeImageCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBody) SetRequestId(v string) *DescribeImageCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageCachesResponseBody) SetImageCaches(v []*DescribeImageCachesResponseBodyImageCaches) *DescribeImageCachesResponseBody {
	s.ImageCaches = v
	return s
}

type DescribeImageCachesResponseBodyImageCaches struct {
	Images           []*string                                           `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	CreationTime     *string                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status           *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Progress         *string                                             `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ExpireDateTime   *string                                             `json:"ExpireDateTime,omitempty" xml:"ExpireDateTime,omitempty"`
	LastMatchedTime  *string                                             `json:"LastMatchedTime,omitempty" xml:"LastMatchedTime,omitempty"`
	ContainerGroupId *string                                             `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	Tags             []*DescribeImageCachesResponseBodyImageCachesTags   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Events           []*DescribeImageCachesResponseBodyImageCachesEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	ImageCacheId     *string                                             `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	RegionId         *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId       *string                                             `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	FlashSnapshotId  *string                                             `json:"FlashSnapshotId,omitempty" xml:"FlashSnapshotId,omitempty"`
	ResourceGroupId  *string                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ImageCacheSize   *int32                                              `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	ImageCacheName   *string                                             `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
}

func (s DescribeImageCachesResponseBodyImageCaches) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCaches) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImages(v []*string) *DescribeImageCachesResponseBodyImageCaches {
	s.Images = v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetCreationTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.CreationTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetStatus(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.Status = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetProgress(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.Progress = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetExpireDateTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ExpireDateTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetLastMatchedTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.LastMatchedTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetContainerGroupId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetTags(v []*DescribeImageCachesResponseBodyImageCachesTags) *DescribeImageCachesResponseBodyImageCaches {
	s.Tags = v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetEvents(v []*DescribeImageCachesResponseBodyImageCachesEvents) *DescribeImageCachesResponseBodyImageCaches {
	s.Events = v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetRegionId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.RegionId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetSnapshotId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetFlashSnapshotId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.FlashSnapshotId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetResourceGroupId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheSize(v int32) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheSize = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheName(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheName = &v
	return s
}

type DescribeImageCachesResponseBodyImageCachesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageCachesResponseBodyImageCachesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCachesTags) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCachesTags) SetKey(v string) *DescribeImageCachesResponseBodyImageCachesTags {
	s.Key = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesTags) SetValue(v string) *DescribeImageCachesResponseBodyImageCachesTags {
	s.Value = &v
	return s
}

type DescribeImageCachesResponseBodyImageCachesEvents struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	LastTimestamp  *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
}

func (s DescribeImageCachesResponseBodyImageCachesEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCachesEvents) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetType(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Type = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetLastTimestamp(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetMessage(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Message = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetName(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Name = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetCount(v int32) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Count = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetFirstTimestamp(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.FirstTimestamp = &v
	return s
}

type DescribeImageCachesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponse) SetHeaders(v map[string]*string) *DescribeImageCachesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageCachesResponse) SetBody(v *DescribeImageCachesResponseBody) *DescribeImageCachesResponse {
	s.Body = v
	return s
}

type DeleteContainerGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupRequest) SetOwnerId(v int64) *DeleteContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetResourceOwnerAccount(v string) *DeleteContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetResourceOwnerId(v int64) *DeleteContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetOwnerAccount(v string) *DeleteContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetRegionId(v string) *DeleteContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetContainerGroupId(v string) *DeleteContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetClientToken(v string) *DeleteContainerGroupRequest {
	s.ClientToken = &v
	return s
}

type DeleteContainerGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupResponseBody) SetRequestId(v string) *DeleteContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContainerGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupResponse) SetHeaders(v map[string]*string) *DeleteContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerGroupResponse) SetBody(v *DeleteContainerGroupResponseBody) *DeleteContainerGroupResponse {
	s.Body = v
	return s
}

type ListUsageRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsageRequest) GoString() string {
	return s.String()
}

func (s *ListUsageRequest) SetOwnerId(v int64) *ListUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *ListUsageRequest) SetResourceOwnerAccount(v string) *ListUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListUsageRequest) SetResourceOwnerId(v int64) *ListUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListUsageRequest) SetOwnerAccount(v string) *ListUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListUsageRequest) SetRegionId(v string) *ListUsageRequest {
	s.RegionId = &v
	return s
}

type ListUsageResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
}

func (s ListUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsageResponseBody) SetRequestId(v string) *ListUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsageResponseBody) SetAttributes(v map[string]interface{}) *ListUsageResponseBody {
	s.Attributes = v
	return s
}

type ListUsageResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsageResponse) GoString() string {
	return s.String()
}

func (s *ListUsageResponse) SetHeaders(v map[string]*string) *ListUsageResponse {
	s.Headers = v
	return s
}

func (s *ListUsageResponse) SetBody(v *ListUsageResponseBody) *ListUsageResponse {
	s.Body = v
	return s
}

type CreateContainerGroupRequest struct {
	DnsConfig                     *CreateContainerGroupRequestDnsConfig                 `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	SecurityContext               *CreateContainerGroupRequestSecurityContext           `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	OwnerId                       *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount          *string                                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId               *int64                                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount                  *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId                      *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                        *string                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SecurityGroupId               *string                                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId                     *string                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ContainerGroupName            *string                                               `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	RestartPolicy                 *string                                               `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	EipInstanceId                 *string                                               `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	Cpu                           *float32                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory                        *float32                                              `json:"Memory,omitempty" xml:"Memory,omitempty"`
	ResourceGroupId               *string                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DnsPolicy                     *string                                               `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	ClientToken                   *string                                               `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceType                  *string                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SlsEnable                     *bool                                                 `json:"SlsEnable,omitempty" xml:"SlsEnable,omitempty"`
	ImageSnapshotId               *string                                               `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	RamRoleName                   *string                                               `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	TerminationGracePeriodSeconds *int64                                                `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	AutoMatchImageCache           *bool                                                 `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	Ipv6AddressCount              *int32                                                `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	ActiveDeadlineSeconds         *int64                                                `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	SpotStrategy                  *string                                               `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SpotPriceLimit                *float32                                              `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	ScheduleStrategy              *string                                               `json:"ScheduleStrategy,omitempty" xml:"ScheduleStrategy,omitempty"`
	TenantVSwitchId               *string                                               `json:"TenantVSwitchId,omitempty" xml:"TenantVSwitchId,omitempty"`
	TenantSecurityGroupId         *string                                               `json:"TenantSecurityGroupId,omitempty" xml:"TenantSecurityGroupId,omitempty"`
	CorePattern                   *string                                               `json:"CorePattern,omitempty" xml:"CorePattern,omitempty"`
	ShareProcessNamespace         *bool                                                 `json:"ShareProcessNamespace,omitempty" xml:"ShareProcessNamespace,omitempty"`
	ProductOnEciMode              *string                                               `json:"ProductOnEciMode,omitempty" xml:"ProductOnEciMode,omitempty"`
	SecondaryENIPolicy            *string                                               `json:"SecondaryENIPolicy,omitempty" xml:"SecondaryENIPolicy,omitempty"`
	AutoCreateEip                 *bool                                                 `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	EipBandwidth                  *int32                                                `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EipISP                        *string                                               `json:"EipISP,omitempty" xml:"EipISP,omitempty"`
	EipCommonBandwidthPackage     *string                                               `json:"EipCommonBandwidthPackage,omitempty" xml:"EipCommonBandwidthPackage,omitempty"`
	HostName                      *string                                               `json:"HostName,omitempty" xml:"HostName,omitempty"`
	IngressBandwidth              *int64                                                `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	EgressBandwidth               *int64                                                `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	CpuOptionsCore                *int32                                                `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	CpuOptionsThreadsPerCore      *int32                                                `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	CpuOptionsNuma                *string                                               `json:"CpuOptionsNuma,omitempty" xml:"CpuOptionsNuma,omitempty"`
	EphemeralStorage              *int32                                                `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	Tag                           []*CreateContainerGroupRequestTag                     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ImageRegistryCredential       []*CreateContainerGroupRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	Container                     []*CreateContainerGroupRequestContainer               `json:"Container,omitempty" xml:"Container,omitempty" type:"Repeated"`
	Volume                        []*CreateContainerGroupRequestVolume                  `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Repeated"`
	InitContainer                 []*CreateContainerGroupRequestInitContainer           `json:"InitContainer,omitempty" xml:"InitContainer,omitempty" type:"Repeated"`
	HostAliase                    []*CreateContainerGroupRequestHostAliase              `json:"HostAliase,omitempty" xml:"HostAliase,omitempty" type:"Repeated"`
	Arn                           []*CreateContainerGroupRequestArn                     `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	NtpServer                     []*string                                             `json:"NtpServer,omitempty" xml:"NtpServer,omitempty" type:"Repeated"`
	AcrRegistryInfo               []*CreateContainerGroupRequestAcrRegistryInfo         `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	SpotDuration                  *int64                                                `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
}

func (s CreateContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequest) SetDnsConfig(v *CreateContainerGroupRequestDnsConfig) *CreateContainerGroupRequest {
	s.DnsConfig = v
	return s
}

func (s *CreateContainerGroupRequest) SetSecurityContext(v *CreateContainerGroupRequestSecurityContext) *CreateContainerGroupRequest {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequest) SetOwnerId(v int64) *CreateContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceOwnerAccount(v string) *CreateContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceOwnerId(v int64) *CreateContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetOwnerAccount(v string) *CreateContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRegionId(v string) *CreateContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetZoneId(v string) *CreateContainerGroupRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSecurityGroupId(v string) *CreateContainerGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetVSwitchId(v string) *CreateContainerGroupRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetContainerGroupName(v string) *CreateContainerGroupRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRestartPolicy(v string) *CreateContainerGroupRequest {
	s.RestartPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipInstanceId(v string) *CreateContainerGroupRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpu(v float32) *CreateContainerGroupRequest {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequest) SetMemory(v float32) *CreateContainerGroupRequest {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceGroupId(v string) *CreateContainerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetDnsPolicy(v string) *CreateContainerGroupRequest {
	s.DnsPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetClientToken(v string) *CreateContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateContainerGroupRequest) SetInstanceType(v string) *CreateContainerGroupRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSlsEnable(v bool) *CreateContainerGroupRequest {
	s.SlsEnable = &v
	return s
}

func (s *CreateContainerGroupRequest) SetImageSnapshotId(v string) *CreateContainerGroupRequest {
	s.ImageSnapshotId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRamRoleName(v string) *CreateContainerGroupRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetTerminationGracePeriodSeconds(v int64) *CreateContainerGroupRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequest) SetAutoMatchImageCache(v bool) *CreateContainerGroupRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIpv6AddressCount(v int32) *CreateContainerGroupRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetActiveDeadlineSeconds(v int64) *CreateContainerGroupRequest {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotStrategy(v string) *CreateContainerGroupRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotPriceLimit(v float32) *CreateContainerGroupRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateContainerGroupRequest) SetScheduleStrategy(v string) *CreateContainerGroupRequest {
	s.ScheduleStrategy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetTenantVSwitchId(v string) *CreateContainerGroupRequest {
	s.TenantVSwitchId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetTenantSecurityGroupId(v string) *CreateContainerGroupRequest {
	s.TenantSecurityGroupId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCorePattern(v string) *CreateContainerGroupRequest {
	s.CorePattern = &v
	return s
}

func (s *CreateContainerGroupRequest) SetShareProcessNamespace(v bool) *CreateContainerGroupRequest {
	s.ShareProcessNamespace = &v
	return s
}

func (s *CreateContainerGroupRequest) SetProductOnEciMode(v string) *CreateContainerGroupRequest {
	s.ProductOnEciMode = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSecondaryENIPolicy(v string) *CreateContainerGroupRequest {
	s.SecondaryENIPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetAutoCreateEip(v bool) *CreateContainerGroupRequest {
	s.AutoCreateEip = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipBandwidth(v int32) *CreateContainerGroupRequest {
	s.EipBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipISP(v string) *CreateContainerGroupRequest {
	s.EipISP = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipCommonBandwidthPackage(v string) *CreateContainerGroupRequest {
	s.EipCommonBandwidthPackage = &v
	return s
}

func (s *CreateContainerGroupRequest) SetHostName(v string) *CreateContainerGroupRequest {
	s.HostName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIngressBandwidth(v int64) *CreateContainerGroupRequest {
	s.IngressBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEgressBandwidth(v int64) *CreateContainerGroupRequest {
	s.EgressBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsCore(v int32) *CreateContainerGroupRequest {
	s.CpuOptionsCore = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsThreadsPerCore(v int32) *CreateContainerGroupRequest {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsNuma(v string) *CreateContainerGroupRequest {
	s.CpuOptionsNuma = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEphemeralStorage(v int32) *CreateContainerGroupRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *CreateContainerGroupRequest) SetTag(v []*CreateContainerGroupRequestTag) *CreateContainerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateContainerGroupRequest) SetImageRegistryCredential(v []*CreateContainerGroupRequestImageRegistryCredential) *CreateContainerGroupRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *CreateContainerGroupRequest) SetContainer(v []*CreateContainerGroupRequestContainer) *CreateContainerGroupRequest {
	s.Container = v
	return s
}

func (s *CreateContainerGroupRequest) SetVolume(v []*CreateContainerGroupRequestVolume) *CreateContainerGroupRequest {
	s.Volume = v
	return s
}

func (s *CreateContainerGroupRequest) SetInitContainer(v []*CreateContainerGroupRequestInitContainer) *CreateContainerGroupRequest {
	s.InitContainer = v
	return s
}

func (s *CreateContainerGroupRequest) SetHostAliase(v []*CreateContainerGroupRequestHostAliase) *CreateContainerGroupRequest {
	s.HostAliase = v
	return s
}

func (s *CreateContainerGroupRequest) SetArn(v []*CreateContainerGroupRequestArn) *CreateContainerGroupRequest {
	s.Arn = v
	return s
}

func (s *CreateContainerGroupRequest) SetNtpServer(v []*string) *CreateContainerGroupRequest {
	s.NtpServer = v
	return s
}

func (s *CreateContainerGroupRequest) SetAcrRegistryInfo(v []*CreateContainerGroupRequestAcrRegistryInfo) *CreateContainerGroupRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotDuration(v int64) *CreateContainerGroupRequest {
	s.SpotDuration = &v
	return s
}

type CreateContainerGroupRequestDnsConfig struct {
	NameServer []*string                                     `json:"NameServer,omitempty" xml:"NameServer,omitempty" type:"Repeated"`
	Search     []*string                                     `json:"Search,omitempty" xml:"Search,omitempty" type:"Repeated"`
	Option     []*CreateContainerGroupRequestDnsConfigOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestDnsConfig) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestDnsConfig) SetNameServer(v []*string) *CreateContainerGroupRequestDnsConfig {
	s.NameServer = v
	return s
}

func (s *CreateContainerGroupRequestDnsConfig) SetSearch(v []*string) *CreateContainerGroupRequestDnsConfig {
	s.Search = v
	return s
}

func (s *CreateContainerGroupRequestDnsConfig) SetOption(v []*CreateContainerGroupRequestDnsConfigOption) *CreateContainerGroupRequestDnsConfig {
	s.Option = v
	return s
}

type CreateContainerGroupRequestDnsConfigOption struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateContainerGroupRequestDnsConfigOption) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestDnsConfigOption) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestDnsConfigOption) SetValue(v string) *CreateContainerGroupRequestDnsConfigOption {
	s.Value = &v
	return s
}

func (s *CreateContainerGroupRequestDnsConfigOption) SetName(v string) *CreateContainerGroupRequestDnsConfigOption {
	s.Name = &v
	return s
}

type CreateContainerGroupRequestSecurityContext struct {
	Sysctl []*CreateContainerGroupRequestSecurityContextSysctl `json:"Sysctl,omitempty" xml:"Sysctl,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestSecurityContext) SetSysctl(v []*CreateContainerGroupRequestSecurityContextSysctl) *CreateContainerGroupRequestSecurityContext {
	s.Sysctl = v
	return s
}

type CreateContainerGroupRequestSecurityContextSysctl struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateContainerGroupRequestSecurityContextSysctl) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestSecurityContextSysctl) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestSecurityContextSysctl) SetValue(v string) *CreateContainerGroupRequestSecurityContextSysctl {
	s.Value = &v
	return s
}

func (s *CreateContainerGroupRequestSecurityContextSysctl) SetName(v string) *CreateContainerGroupRequestSecurityContextSysctl {
	s.Name = &v
	return s
}

type CreateContainerGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestTag) SetKey(v string) *CreateContainerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestTag) SetValue(v string) *CreateContainerGroupRequestTag {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestImageRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateContainerGroupRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetPassword(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetServer(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetUserName(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type CreateContainerGroupRequestContainer struct {
	ReadinessProbe                             *CreateContainerGroupRequestContainerReadinessProbe                               `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" require:"true" type:"Struct"`
	SecurityContext                            *CreateContainerGroupRequestContainerSecurityContext                              `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	LivenessProbe                              *CreateContainerGroupRequestContainerLivenessProbe                                `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" require:"true" type:"Struct"`
	EnvironmentVar                             []*CreateContainerGroupRequestContainerEnvironmentVar                             `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	Tty                                        *bool                                                                             `json:"Tty,omitempty" xml:"Tty,omitempty"`
	WorkingDir                                 *string                                                                           `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	Arg                                        []*string                                                                         `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Stdin                                      *bool                                                                             `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	VolumeMount                                []*CreateContainerGroupRequestContainerVolumeMount                                `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	ImagePullPolicy                            *string                                                                           `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	StdinOnce                                  *bool                                                                             `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	LifecyclePreStopHandlerTcpSocketPort       *int32                                                                            `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	LifecyclePostStartHandlerHttpGetScheme     *string                                                                           `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	Command                                    []*string                                                                         `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	LifecyclePostStartHandlerHttpGetHost       *string                                                                           `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	TerminationMessagePolicy                   *string                                                                           `json:"TerminationMessagePolicy,omitempty" xml:"TerminationMessagePolicy,omitempty"`
	LifecyclePostStartHandlerTcpSocketPort     *int32                                                                            `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	LifecyclePostStartHandlerHttpGetPath       *string                                                                           `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	LifecyclePostStartHandlerExec              []*string                                                                         `json:"LifecyclePostStartHandlerExec,omitempty" xml:"LifecyclePostStartHandlerExec,omitempty" type:"Repeated"`
	LifecyclePreStopHandlerHttpGetPath         *string                                                                           `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	Port                                       []*CreateContainerGroupRequestContainerPort                                       `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	TerminationMessagePath                     *string                                                                           `json:"TerminationMessagePath,omitempty" xml:"TerminationMessagePath,omitempty"`
	LifecyclePreStopHandlerHttpGetScheme       *string                                                                           `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	LifecyclePostStartHandlerTcpSocketHost     *string                                                                           `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	Gpu                                        *int32                                                                            `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	LifecyclePreStopHandlerExec                []*string                                                                         `json:"LifecyclePreStopHandlerExec,omitempty" xml:"LifecyclePreStopHandlerExec,omitempty" type:"Repeated"`
	Memory                                     *float32                                                                          `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                                       *string                                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	LifecyclePreStopHandlerHttpGetHost         *string                                                                           `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	LifecyclePreStopHandlerTcpSocketHost       *string                                                                           `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	Image                                      *string                                                                           `json:"Image,omitempty" xml:"Image,omitempty"`
	LifecyclePreStopHandlerHttpGetPort         *int32                                                                            `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	LifecyclePreStopHandlerHttpGetHttpHeader   []*CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader   `json:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	Cpu                                        *float32                                                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	LifecyclePostStartHandlerHttpGetPort       *int32                                                                            `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	LifecyclePostStartHandlerHttpGetHttpHeader []*CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader `json:"LifecyclePostStartHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePostStartHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainer) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainer) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainer) SetReadinessProbe(v *CreateContainerGroupRequestContainerReadinessProbe) *CreateContainerGroupRequestContainer {
	s.ReadinessProbe = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetSecurityContext(v *CreateContainerGroupRequestContainerSecurityContext) *CreateContainerGroupRequestContainer {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLivenessProbe(v *CreateContainerGroupRequestContainerLivenessProbe) *CreateContainerGroupRequestContainer {
	s.LivenessProbe = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetEnvironmentVar(v []*CreateContainerGroupRequestContainerEnvironmentVar) *CreateContainerGroupRequestContainer {
	s.EnvironmentVar = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTty(v bool) *CreateContainerGroupRequestContainer {
	s.Tty = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetWorkingDir(v string) *CreateContainerGroupRequestContainer {
	s.WorkingDir = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetArg(v []*string) *CreateContainerGroupRequestContainer {
	s.Arg = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetStdin(v bool) *CreateContainerGroupRequestContainer {
	s.Stdin = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetVolumeMount(v []*CreateContainerGroupRequestContainerVolumeMount) *CreateContainerGroupRequestContainer {
	s.VolumeMount = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetImagePullPolicy(v string) *CreateContainerGroupRequestContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetStdinOnce(v bool) *CreateContainerGroupRequestContainer {
	s.StdinOnce = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetScheme(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetCommand(v []*string) *CreateContainerGroupRequestContainer {
	s.Command = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTerminationMessagePolicy(v string) *CreateContainerGroupRequestContainer {
	s.TerminationMessagePolicy = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPath(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerExec(v []*string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerExec = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPath(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetPort(v []*CreateContainerGroupRequestContainerPort) *CreateContainerGroupRequestContainer {
	s.Port = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTerminationMessagePath(v string) *CreateContainerGroupRequestContainer {
	s.TerminationMessagePath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetScheme(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetGpu(v int32) *CreateContainerGroupRequestContainer {
	s.Gpu = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerExec(v []*string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerExec = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetMemory(v float32) *CreateContainerGroupRequestContainer {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetName(v string) *CreateContainerGroupRequestContainer {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetImage(v string) *CreateContainerGroupRequestContainer {
	s.Image = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHttpHeader(v []*CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHttpHeader = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetCpu(v float32) *CreateContainerGroupRequestContainer {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHttpHeader(v []*CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHttpHeader = v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbe struct {
	TimeoutSeconds      *int32                                                       `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	SuccessThreshold    *int32                                                       `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateContainerGroupRequestContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	HttpGet             *CreateContainerGroupRequestContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	PeriodSeconds       *int32                                                       `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	InitialDelaySeconds *int32                                                       `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	Exec                *CreateContainerGroupRequestContainerReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                       `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbe) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetTimeoutSeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetSuccessThreshold(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetTcpSocket(v *CreateContainerGroupRequestContainerReadinessProbeTcpSocket) *CreateContainerGroupRequestContainerReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetHttpGet(v *CreateContainerGroupRequestContainerReadinessProbeHttpGet) *CreateContainerGroupRequestContainerReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetPeriodSeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetInitialDelaySeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetExec(v *CreateContainerGroupRequestContainerReadinessProbeExec) *CreateContainerGroupRequestContainerReadinessProbe {
	s.Exec = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetFailureThreshold(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.FailureThreshold = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeTcpSocket) SetPort(v int32) *CreateContainerGroupRequestContainerReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeHttpGet struct {
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetScheme(v string) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetPath(v string) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetPort(v int32) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeExec) SetCommand(v []*string) *CreateContainerGroupRequestContainerReadinessProbeExec {
	s.Command = v
	return s
}

type CreateContainerGroupRequestContainerSecurityContext struct {
	Capability             *CreateContainerGroupRequestContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                          `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                         `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateContainerGroupRequestContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetCapability(v *CreateContainerGroupRequestContainerSecurityContextCapability) *CreateContainerGroupRequestContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateContainerGroupRequestContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetRunAsUser(v int64) *CreateContainerGroupRequestContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type CreateContainerGroupRequestContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerSecurityContextCapability) SetAdd(v []*string) *CreateContainerGroupRequestContainerSecurityContextCapability {
	s.Add = v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbe struct {
	PeriodSeconds       *int32                                                      `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	TcpSocket           *CreateContainerGroupRequestContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                      `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	SuccessThreshold    *int32                                                      `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	Exec                *CreateContainerGroupRequestContainerLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	HttpGet             *CreateContainerGroupRequestContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                      `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	TimeoutSeconds      *int32                                                      `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbe) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetPeriodSeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetTcpSocket(v *CreateContainerGroupRequestContainerLivenessProbeTcpSocket) *CreateContainerGroupRequestContainerLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetInitialDelaySeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetSuccessThreshold(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetExec(v *CreateContainerGroupRequestContainerLivenessProbeExec) *CreateContainerGroupRequestContainerLivenessProbe {
	s.Exec = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetHttpGet(v *CreateContainerGroupRequestContainerLivenessProbeHttpGet) *CreateContainerGroupRequestContainerLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetFailureThreshold(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetTimeoutSeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeTcpSocket) SetPort(v int32) *CreateContainerGroupRequestContainerLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeExec) SetCommand(v []*string) *CreateContainerGroupRequestContainerLivenessProbeExec {
	s.Command = v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeHttpGet struct {
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetScheme(v string) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetPort(v int32) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetPath(v string) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Path = &v
	return s
}

type CreateContainerGroupRequestContainerEnvironmentVar struct {
	FieldRef *CreateContainerGroupRequestContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                     `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetFieldRef(v *CreateContainerGroupRequestContainerEnvironmentVarFieldRef) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetKey(v string) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetValue(v string) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s CreateContainerGroupRequestContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerEnvironmentVarFieldRef) SetFieldPath(v string) *CreateContainerGroupRequestContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type CreateContainerGroupRequestContainerVolumeMount struct {
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateContainerGroupRequestContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetMountPropagation(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetMountPath(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetReadOnly(v bool) *CreateContainerGroupRequestContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetSubPath(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.SubPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetName(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.Name = &v
	return s
}

type CreateContainerGroupRequestContainerPort struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestContainerPort) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerPort) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerPort) SetProtocol(v string) *CreateContainerGroupRequestContainerPort {
	s.Protocol = &v
	return s
}

func (s *CreateContainerGroupRequestContainerPort) SetPort(v int32) *CreateContainerGroupRequestContainerPort {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetValue(v string) *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetName(v string) *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

type CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) SetValue(v string) *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) SetName(v string) *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

type CreateContainerGroupRequestVolume struct {
	DiskVolume       *CreateContainerGroupRequestVolumeDiskVolume       `json:"DiskVolume,omitempty" xml:"DiskVolume,omitempty" require:"true" type:"Struct"`
	NFSVolume        *CreateContainerGroupRequestVolumeNFSVolume        `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" require:"true" type:"Struct"`
	FlexVolume       *CreateContainerGroupRequestVolumeFlexVolume       `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" require:"true" type:"Struct"`
	HostPathVolume   *CreateContainerGroupRequestVolumeHostPathVolume   `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" require:"true" type:"Struct"`
	ConfigFileVolume *CreateContainerGroupRequestVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" xml:"ConfigFileVolume,omitempty" require:"true" type:"Struct"`
	EmptyDirVolume   *CreateContainerGroupRequestVolumeEmptyDirVolume   `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" require:"true" type:"Struct"`
	Type             *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	Name             *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateContainerGroupRequestVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolume) SetDiskVolume(v *CreateContainerGroupRequestVolumeDiskVolume) *CreateContainerGroupRequestVolume {
	s.DiskVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetNFSVolume(v *CreateContainerGroupRequestVolumeNFSVolume) *CreateContainerGroupRequestVolume {
	s.NFSVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetFlexVolume(v *CreateContainerGroupRequestVolumeFlexVolume) *CreateContainerGroupRequestVolume {
	s.FlexVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetHostPathVolume(v *CreateContainerGroupRequestVolumeHostPathVolume) *CreateContainerGroupRequestVolume {
	s.HostPathVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetConfigFileVolume(v *CreateContainerGroupRequestVolumeConfigFileVolume) *CreateContainerGroupRequestVolume {
	s.ConfigFileVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetEmptyDirVolume(v *CreateContainerGroupRequestVolumeEmptyDirVolume) *CreateContainerGroupRequestVolume {
	s.EmptyDirVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetType(v string) *CreateContainerGroupRequestVolume {
	s.Type = &v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetName(v string) *CreateContainerGroupRequestVolume {
	s.Name = &v
	return s
}

type CreateContainerGroupRequestVolumeDiskVolume struct {
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	FsType   *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
}

func (s CreateContainerGroupRequestVolumeDiskVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeDiskVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetDiskSize(v int32) *CreateContainerGroupRequestVolumeDiskVolume {
	s.DiskSize = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetFsType(v string) *CreateContainerGroupRequestVolumeDiskVolume {
	s.FsType = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetDiskId(v string) *CreateContainerGroupRequestVolumeDiskVolume {
	s.DiskId = &v
	return s
}

type CreateContainerGroupRequestVolumeNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s CreateContainerGroupRequestVolumeNFSVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeNFSVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetPath(v string) *CreateContainerGroupRequestVolumeNFSVolume {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetReadOnly(v bool) *CreateContainerGroupRequestVolumeNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetServer(v string) *CreateContainerGroupRequestVolumeNFSVolume {
	s.Server = &v
	return s
}

type CreateContainerGroupRequestVolumeFlexVolume struct {
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
}

func (s CreateContainerGroupRequestVolumeFlexVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeFlexVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetFsType(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.FsType = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetOptions(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.Options = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetDriver(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.Driver = &v
	return s
}

type CreateContainerGroupRequestVolumeHostPathVolume struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateContainerGroupRequestVolumeHostPathVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeHostPathVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeHostPathVolume) SetType(v string) *CreateContainerGroupRequestVolumeHostPathVolume {
	s.Type = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeHostPathVolume) SetPath(v string) *CreateContainerGroupRequestVolumeHostPathVolume {
	s.Path = &v
	return s
}

type CreateContainerGroupRequestVolumeConfigFileVolume struct {
	DefaultMode      *int32                                                               `json:"DefaultMode,omitempty" xml:"DefaultMode,omitempty"`
	ConfigFileToPath []*CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath `json:"ConfigFileToPath,omitempty" xml:"ConfigFileToPath,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestVolumeConfigFileVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeConfigFileVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolume) SetDefaultMode(v int32) *CreateContainerGroupRequestVolumeConfigFileVolume {
	s.DefaultMode = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolume) SetConfigFileToPath(v []*CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) *CreateContainerGroupRequestVolumeConfigFileVolume {
	s.ConfigFileToPath = v
	return s
}

type CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath struct {
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Mode    *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetPath(v string) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetMode(v int32) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Mode = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetContent(v string) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Content = &v
	return s
}

type CreateContainerGroupRequestVolumeEmptyDirVolume struct {
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
}

func (s CreateContainerGroupRequestVolumeEmptyDirVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeEmptyDirVolume) SetMedium(v string) *CreateContainerGroupRequestVolumeEmptyDirVolume {
	s.Medium = &v
	return s
}

type CreateContainerGroupRequestInitContainer struct {
	SecurityContext          *CreateContainerGroupRequestInitContainerSecurityContext  `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Image                    *string                                                   `json:"Image,omitempty" xml:"Image,omitempty"`
	VolumeMount              []*CreateContainerGroupRequestInitContainerVolumeMount    `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	Port                     []*CreateContainerGroupRequestInitContainerPort           `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	TerminationMessagePath   *string                                                   `json:"TerminationMessagePath,omitempty" xml:"TerminationMessagePath,omitempty"`
	EnvironmentVar           []*CreateContainerGroupRequestInitContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	ImagePullPolicy          *string                                                   `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	WorkingDir               *string                                                   `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	Cpu                      *float32                                                  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Arg                      []*string                                                 `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command                  []*string                                                 `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Gpu                      *int32                                                    `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory                   *float32                                                  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	TerminationMessagePolicy *string                                                   `json:"TerminationMessagePolicy,omitempty" xml:"TerminationMessagePolicy,omitempty"`
	Name                     *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateContainerGroupRequestInitContainer) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainer) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainer) SetSecurityContext(v *CreateContainerGroupRequestInitContainerSecurityContext) *CreateContainerGroupRequestInitContainer {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetImage(v string) *CreateContainerGroupRequestInitContainer {
	s.Image = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetVolumeMount(v []*CreateContainerGroupRequestInitContainerVolumeMount) *CreateContainerGroupRequestInitContainer {
	s.VolumeMount = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetPort(v []*CreateContainerGroupRequestInitContainerPort) *CreateContainerGroupRequestInitContainer {
	s.Port = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetTerminationMessagePath(v string) *CreateContainerGroupRequestInitContainer {
	s.TerminationMessagePath = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetEnvironmentVar(v []*CreateContainerGroupRequestInitContainerEnvironmentVar) *CreateContainerGroupRequestInitContainer {
	s.EnvironmentVar = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetImagePullPolicy(v string) *CreateContainerGroupRequestInitContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetWorkingDir(v string) *CreateContainerGroupRequestInitContainer {
	s.WorkingDir = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetCpu(v float32) *CreateContainerGroupRequestInitContainer {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetArg(v []*string) *CreateContainerGroupRequestInitContainer {
	s.Arg = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetCommand(v []*string) *CreateContainerGroupRequestInitContainer {
	s.Command = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetGpu(v int32) *CreateContainerGroupRequestInitContainer {
	s.Gpu = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetMemory(v float32) *CreateContainerGroupRequestInitContainer {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetTerminationMessagePolicy(v string) *CreateContainerGroupRequestInitContainer {
	s.TerminationMessagePolicy = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetName(v string) *CreateContainerGroupRequestInitContainer {
	s.Name = &v
	return s
}

type CreateContainerGroupRequestInitContainerSecurityContext struct {
	Capability             *CreateContainerGroupRequestInitContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                              `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                             `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetCapability(v *CreateContainerGroupRequestInitContainerSecurityContextCapability) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetRunAsUser(v int64) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type CreateContainerGroupRequestInitContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestInitContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerSecurityContextCapability) SetAdd(v []*string) *CreateContainerGroupRequestInitContainerSecurityContextCapability {
	s.Add = v
	return s
}

type CreateContainerGroupRequestInitContainerVolumeMount struct {
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetMountPropagation(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetMountPath(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetReadOnly(v bool) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetSubPath(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.SubPath = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetName(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.Name = &v
	return s
}

type CreateContainerGroupRequestInitContainerPort struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerPort) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerPort) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerPort) SetProtocol(v string) *CreateContainerGroupRequestInitContainerPort {
	s.Protocol = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerPort) SetPort(v int32) *CreateContainerGroupRequestInitContainerPort {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestInitContainerEnvironmentVar struct {
	FieldRef *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                         `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                         `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetFieldRef(v *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetKey(v string) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetValue(v string) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) SetFieldPath(v string) *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type CreateContainerGroupRequestHostAliase struct {
	Ip       *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Hostname []*string `json:"Hostname,omitempty" xml:"Hostname,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestHostAliase) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestHostAliase) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestHostAliase) SetIp(v string) *CreateContainerGroupRequestHostAliase {
	s.Ip = &v
	return s
}

func (s *CreateContainerGroupRequestHostAliase) SetHostname(v []*string) *CreateContainerGroupRequestHostAliase {
	s.Hostname = v
	return s
}

type CreateContainerGroupRequestArn struct {
	RoleArn       *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RoleType      *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
}

func (s CreateContainerGroupRequestArn) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestArn) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestArn) SetRoleArn(v string) *CreateContainerGroupRequestArn {
	s.RoleArn = &v
	return s
}

func (s *CreateContainerGroupRequestArn) SetRoleType(v string) *CreateContainerGroupRequestArn {
	s.RoleType = &v
	return s
}

func (s *CreateContainerGroupRequestArn) SetAssumeRoleFor(v string) *CreateContainerGroupRequestArn {
	s.AssumeRoleFor = &v
	return s
}

type CreateContainerGroupRequestAcrRegistryInfo struct {
	Domain       []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateContainerGroupRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetDomain(v []*string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetInstanceName(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetInstanceId(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetRegionId(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type CreateContainerGroupResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
}

func (s CreateContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupResponseBody) SetRequestId(v string) *CreateContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContainerGroupResponseBody) SetContainerGroupId(v string) *CreateContainerGroupResponseBody {
	s.ContainerGroupId = &v
	return s
}

type CreateContainerGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupResponse) SetHeaders(v map[string]*string) *CreateContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateContainerGroupResponse) SetBody(v *CreateContainerGroupResponseBody) *CreateContainerGroupResponse {
	s.Body = v
	return s
}

type DeleteImageCacheRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ImageCacheId         *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteImageCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheRequest) SetOwnerId(v int64) *DeleteImageCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetResourceOwnerAccount(v string) *DeleteImageCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteImageCacheRequest) SetResourceOwnerId(v int64) *DeleteImageCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetOwnerAccount(v string) *DeleteImageCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteImageCacheRequest) SetRegionId(v string) *DeleteImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetImageCacheId(v string) *DeleteImageCacheRequest {
	s.ImageCacheId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetClientToken(v string) *DeleteImageCacheRequest {
	s.ClientToken = &v
	return s
}

type DeleteImageCacheResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponseBody) SetRequestId(v string) *DeleteImageCacheResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageCacheResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponse) SetHeaders(v map[string]*string) *DeleteImageCacheResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageCacheResponse) SetBody(v *DeleteImageCacheResponseBody) *DeleteImageCacheResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupMetricRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeContainerGroupMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricRequest) SetOwnerId(v int64) *DescribeContainerGroupMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetOwnerAccount(v string) *DescribeContainerGroupMetricRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetRegionId(v string) *DescribeContainerGroupMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetContainerGroupId(v string) *DescribeContainerGroupMetricRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetStartTime(v string) *DescribeContainerGroupMetricRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetEndTime(v string) *DescribeContainerGroupMetricRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetPeriod(v string) *DescribeContainerGroupMetricRequest {
	s.Period = &v
	return s
}

type DescribeContainerGroupMetricResponseBody struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContainerGroupId *string                                            `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	Records          []*DescribeContainerGroupMetricResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBody) SetRequestId(v string) *DescribeContainerGroupMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBody) SetContainerGroupId(v string) *DescribeContainerGroupMetricResponseBody {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBody) SetRecords(v []*DescribeContainerGroupMetricResponseBodyRecords) *DescribeContainerGroupMetricResponseBody {
	s.Records = v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecords struct {
	Network    *DescribeContainerGroupMetricResponseBodyRecordsNetwork      `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	CPU        *DescribeContainerGroupMetricResponseBodyRecordsCPU          `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	Disk       []*DescribeContainerGroupMetricResponseBodyRecordsDisk       `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	Timestamp  *string                                                      `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Memory     *DescribeContainerGroupMetricResponseBodyRecordsMemory       `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Filesystem []*DescribeContainerGroupMetricResponseBodyRecordsFilesystem `json:"Filesystem,omitempty" xml:"Filesystem,omitempty" type:"Repeated"`
	Containers []*DescribeContainerGroupMetricResponseBodyRecordsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupMetricResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetNetwork(v *DescribeContainerGroupMetricResponseBodyRecordsNetwork) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Network = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetCPU(v *DescribeContainerGroupMetricResponseBodyRecordsCPU) *DescribeContainerGroupMetricResponseBodyRecords {
	s.CPU = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetDisk(v []*DescribeContainerGroupMetricResponseBodyRecordsDisk) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Disk = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetTimestamp(v string) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Timestamp = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetMemory(v *DescribeContainerGroupMetricResponseBodyRecordsMemory) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Memory = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetFilesystem(v []*DescribeContainerGroupMetricResponseBodyRecordsFilesystem) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Filesystem = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetContainers(v []*DescribeContainerGroupMetricResponseBodyRecordsContainers) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Containers = v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsNetwork struct {
	Interfaces []*DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces `json:"Interfaces,omitempty" xml:"Interfaces,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetwork) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetwork) SetInterfaces(v []*DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) *DescribeContainerGroupMetricResponseBodyRecordsNetwork {
	s.Interfaces = v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces struct {
	RxErrors  *int64  `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	TxDrops   *int64  `json:"TxDrops,omitempty" xml:"TxDrops,omitempty"`
	TxBytes   *int64  `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	RxPackets *int64  `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	TxPackets *int64  `json:"TxPackets,omitempty" xml:"TxPackets,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TxErrors  *int64  `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	RxBytes   *int64  `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	RxDrops   *int64  `json:"RxDrops,omitempty" xml:"RxDrops,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxErrors(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxErrors = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxDrops(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxDrops = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxPackets(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxPackets = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxPackets(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxPackets = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetName(v string) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxErrors(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxErrors = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxDrops(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxDrops = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsCPU struct {
	UsageNanoCores       *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
	Limit                *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	Load                 *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsCPU) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetUsageNanoCores(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.UsageNanoCores = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetLimit(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetUsageCoreNanoSeconds(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetLoad(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.Load = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsDisk struct {
	WriteBytes *int64  `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	WriteIO    *int64  `json:"WriteIO,omitempty" xml:"WriteIO,omitempty"`
	Device     *string `json:"Device,omitempty" xml:"Device,omitempty"`
	ReadIO     *int64  `json:"ReadIO,omitempty" xml:"ReadIO,omitempty"`
	ReadBytes  *int64  `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsDisk) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetWriteBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.WriteBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetWriteIO(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.WriteIO = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetDevice(v string) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.Device = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetReadIO(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.ReadIO = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetReadBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.ReadBytes = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsMemory struct {
	Rss            *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	UsageBytes     *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	WorkingSet     *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	Cache          *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsMemory) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetRss(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.Rss = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetUsageBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetWorkingSet(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.WorkingSet = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetAvailableBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetCache(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.Cache = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsFilesystem struct {
	Capacity  *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Available *int64  `json:"Available,omitempty" xml:"Available,omitempty"`
	FsName    *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	Usage     *int64  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Category  *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsFilesystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsFilesystem) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetCapacity(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Capacity = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetAvailable(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Available = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetFsName(v string) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.FsName = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetUsage(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Usage = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetCategory(v string) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Category = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainers struct {
	CPU    *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU    `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	Memory *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Name   *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetCPU(v *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.CPU = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetMemory(v *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.Memory = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetName(v string) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.Name = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainersCPU struct {
	UsageNanoCores       *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
	Limit                *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	Load                 *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetUsageNanoCores(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.UsageNanoCores = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetLimit(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetUsageCoreNanoSeconds(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetLoad(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.Load = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainersMemory struct {
	Rss            *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	UsageBytes     *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	WorkingSet     *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	Cache          *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetRss(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.Rss = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetUsageBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetWorkingSet(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.WorkingSet = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetAvailableBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetCache(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.Cache = &v
	return s
}

type DescribeContainerGroupMetricResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContainerGroupMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupMetricResponse) SetBody(v *DescribeContainerGroupMetricResponseBody) *DescribeContainerGroupMetricResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
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
	Zones          []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
	RecommendZones []*string `json:"RecommendZones,omitempty" xml:"RecommendZones,omitempty" type:"Repeated"`
	RegionEndpoint *string   `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*string) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRecommendZones(v []*string) *DescribeRegionsResponseBodyRegions {
	s.RecommendZones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
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

type UpdateContainerGroupRequest struct {
	DnsConfig               *UpdateContainerGroupRequestDnsConfig                 `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	OwnerId                 *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string                                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64                                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount            *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId                *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContainerGroupId        *string                                               `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	RestartPolicy           *string                                               `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	ClientToken             *string                                               `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Cpu                     *float32                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory                  *float32                                              `json:"Memory,omitempty" xml:"Memory,omitempty"`
	ResourceGroupId         *string                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                     []*UpdateContainerGroupRequestTag                     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Volume                  []*UpdateContainerGroupRequestVolume                  `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Repeated"`
	Container               []*UpdateContainerGroupRequestContainer               `json:"Container,omitempty" xml:"Container,omitempty" type:"Repeated"`
	InitContainer           []*UpdateContainerGroupRequestInitContainer           `json:"InitContainer,omitempty" xml:"InitContainer,omitempty" type:"Repeated"`
	ImageRegistryCredential []*UpdateContainerGroupRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequest) SetDnsConfig(v *UpdateContainerGroupRequestDnsConfig) *UpdateContainerGroupRequest {
	s.DnsConfig = v
	return s
}

func (s *UpdateContainerGroupRequest) SetOwnerId(v int64) *UpdateContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceOwnerAccount(v string) *UpdateContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceOwnerId(v int64) *UpdateContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetOwnerAccount(v string) *UpdateContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetRegionId(v string) *UpdateContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetContainerGroupId(v string) *UpdateContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetRestartPolicy(v string) *UpdateContainerGroupRequest {
	s.RestartPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetClientToken(v string) *UpdateContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetCpu(v float32) *UpdateContainerGroupRequest {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetMemory(v float32) *UpdateContainerGroupRequest {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceGroupId(v string) *UpdateContainerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetTag(v []*UpdateContainerGroupRequestTag) *UpdateContainerGroupRequest {
	s.Tag = v
	return s
}

func (s *UpdateContainerGroupRequest) SetVolume(v []*UpdateContainerGroupRequestVolume) *UpdateContainerGroupRequest {
	s.Volume = v
	return s
}

func (s *UpdateContainerGroupRequest) SetContainer(v []*UpdateContainerGroupRequestContainer) *UpdateContainerGroupRequest {
	s.Container = v
	return s
}

func (s *UpdateContainerGroupRequest) SetInitContainer(v []*UpdateContainerGroupRequestInitContainer) *UpdateContainerGroupRequest {
	s.InitContainer = v
	return s
}

func (s *UpdateContainerGroupRequest) SetImageRegistryCredential(v []*UpdateContainerGroupRequestImageRegistryCredential) *UpdateContainerGroupRequest {
	s.ImageRegistryCredential = v
	return s
}

type UpdateContainerGroupRequestDnsConfig struct {
	NameServer []*string                                     `json:"NameServer,omitempty" xml:"NameServer,omitempty" type:"Repeated"`
	Search     []*string                                     `json:"Search,omitempty" xml:"Search,omitempty" type:"Repeated"`
	Option     []*UpdateContainerGroupRequestDnsConfigOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestDnsConfig) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestDnsConfig) SetNameServer(v []*string) *UpdateContainerGroupRequestDnsConfig {
	s.NameServer = v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfig) SetSearch(v []*string) *UpdateContainerGroupRequestDnsConfig {
	s.Search = v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfig) SetOption(v []*UpdateContainerGroupRequestDnsConfigOption) *UpdateContainerGroupRequestDnsConfig {
	s.Option = v
	return s
}

type UpdateContainerGroupRequestDnsConfigOption struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateContainerGroupRequestDnsConfigOption) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestDnsConfigOption) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestDnsConfigOption) SetValue(v string) *UpdateContainerGroupRequestDnsConfigOption {
	s.Value = &v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfigOption) SetName(v string) *UpdateContainerGroupRequestDnsConfigOption {
	s.Name = &v
	return s
}

type UpdateContainerGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestTag) SetKey(v string) *UpdateContainerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestTag) SetValue(v string) *UpdateContainerGroupRequestTag {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestVolume struct {
	NFSVolume        *UpdateContainerGroupRequestVolumeNFSVolume        `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" require:"true" type:"Struct"`
	ConfigFileVolume *UpdateContainerGroupRequestVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" xml:"ConfigFileVolume,omitempty" require:"true" type:"Struct"`
	EmptyDirVolume   *UpdateContainerGroupRequestVolumeEmptyDirVolume   `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" require:"true" type:"Struct"`
	Type             *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	Name             *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateContainerGroupRequestVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolume) SetNFSVolume(v *UpdateContainerGroupRequestVolumeNFSVolume) *UpdateContainerGroupRequestVolume {
	s.NFSVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetConfigFileVolume(v *UpdateContainerGroupRequestVolumeConfigFileVolume) *UpdateContainerGroupRequestVolume {
	s.ConfigFileVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetEmptyDirVolume(v *UpdateContainerGroupRequestVolumeEmptyDirVolume) *UpdateContainerGroupRequestVolume {
	s.EmptyDirVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetType(v string) *UpdateContainerGroupRequestVolume {
	s.Type = &v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetName(v string) *UpdateContainerGroupRequestVolume {
	s.Name = &v
	return s
}

type UpdateContainerGroupRequestVolumeNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeNFSVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeNFSVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetPath(v string) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetServer(v string) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.Server = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetReadOnly(v bool) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.ReadOnly = &v
	return s
}

type UpdateContainerGroupRequestVolumeConfigFileVolume struct {
	ConfigFileToPath []*UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath `json:"ConfigFileToPath,omitempty" xml:"ConfigFileToPath,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolume) SetConfigFileToPath(v []*UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) *UpdateContainerGroupRequestVolumeConfigFileVolume {
	s.ConfigFileToPath = v
	return s
}

type UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath struct {
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetPath(v string) *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetContent(v string) *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Content = &v
	return s
}

type UpdateContainerGroupRequestVolumeEmptyDirVolume struct {
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeEmptyDirVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeEmptyDirVolume) SetMedium(v string) *UpdateContainerGroupRequestVolumeEmptyDirVolume {
	s.Medium = &v
	return s
}

type UpdateContainerGroupRequestContainer struct {
	ReadinessProbe                              *UpdateContainerGroupRequestContainerReadinessProbe                                `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" require:"true" type:"Struct"`
	SecurityContext                             *UpdateContainerGroupRequestContainerSecurityContext                               `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	LivenessProbe                               *UpdateContainerGroupRequestContainerLivenessProbe                                 `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" require:"true" type:"Struct"`
	EnvironmentVar                              []*UpdateContainerGroupRequestContainerEnvironmentVar                              `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	Tty                                         *bool                                                                              `json:"Tty,omitempty" xml:"Tty,omitempty"`
	WorkingDir                                  *string                                                                            `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	Arg                                         []*string                                                                          `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Stdin                                       *bool                                                                              `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	VolumeMount                                 []*UpdateContainerGroupRequestContainerVolumeMount                                 `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	ImagePullPolicy                             *string                                                                            `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	StdinOnce                                   *bool                                                                              `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	LifecyclePreStopHandlerTcpSocketPort        *int32                                                                             `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	LifecyclePostStartHandlerHttpGetScheme      *string                                                                            `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	Command                                     []*string                                                                          `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	LifecyclePostStartHandlerHttpGetHost        *string                                                                            `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	LifecyclePostStartHandlerTcpSocketPort      *int32                                                                             `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	LifecyclePostStartHandlerHttpGetPath        *string                                                                            `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	LifecyclePostStartHandlerExec               []*string                                                                          `json:"LifecyclePostStartHandlerExec,omitempty" xml:"LifecyclePostStartHandlerExec,omitempty" type:"Repeated"`
	LifecyclePreStopHandlerHttpGetPath          *string                                                                            `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	Port                                        []*UpdateContainerGroupRequestContainerPort                                        `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	LifecyclePreStopHandlerHttpGetScheme        *string                                                                            `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	LifecyclePostStartHandlerHttpGetHttpHeaders []*UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders `json:"LifecyclePostStartHandlerHttpGetHttpHeaders,omitempty" xml:"LifecyclePostStartHandlerHttpGetHttpHeaders,omitempty" type:"Repeated"`
	LifecyclePostStartHandlerTcpSocketHost      *string                                                                            `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	Gpu                                         *int32                                                                             `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	LifecyclePreStopHandlerExec                 []*string                                                                          `json:"LifecyclePreStopHandlerExec,omitempty" xml:"LifecyclePreStopHandlerExec,omitempty" type:"Repeated"`
	Memory                                      *float32                                                                           `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                                        *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	LifecyclePreStopHandlerHttpGetHost          *string                                                                            `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	LifecyclePreStopHandlerTcpSocketHost        *string                                                                            `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	Image                                       *string                                                                            `json:"Image,omitempty" xml:"Image,omitempty"`
	LifecyclePreStopHandlerHttpGetPort          *int32                                                                             `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	LifecyclePreStopHandlerHttpGetHttpHeader    []*UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader    `json:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	Cpu                                         *float32                                                                           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	LifecyclePostStartHandlerHttpGetPort        *int32                                                                             `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
}

func (s UpdateContainerGroupRequestContainer) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainer) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainer) SetReadinessProbe(v *UpdateContainerGroupRequestContainerReadinessProbe) *UpdateContainerGroupRequestContainer {
	s.ReadinessProbe = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetSecurityContext(v *UpdateContainerGroupRequestContainerSecurityContext) *UpdateContainerGroupRequestContainer {
	s.SecurityContext = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLivenessProbe(v *UpdateContainerGroupRequestContainerLivenessProbe) *UpdateContainerGroupRequestContainer {
	s.LivenessProbe = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetEnvironmentVar(v []*UpdateContainerGroupRequestContainerEnvironmentVar) *UpdateContainerGroupRequestContainer {
	s.EnvironmentVar = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetTty(v bool) *UpdateContainerGroupRequestContainer {
	s.Tty = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetWorkingDir(v string) *UpdateContainerGroupRequestContainer {
	s.WorkingDir = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetArg(v []*string) *UpdateContainerGroupRequestContainer {
	s.Arg = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetStdin(v bool) *UpdateContainerGroupRequestContainer {
	s.Stdin = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetVolumeMount(v []*UpdateContainerGroupRequestContainerVolumeMount) *UpdateContainerGroupRequestContainer {
	s.VolumeMount = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetImagePullPolicy(v string) *UpdateContainerGroupRequestContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetStdinOnce(v bool) *UpdateContainerGroupRequestContainer {
	s.StdinOnce = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetScheme(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetCommand(v []*string) *UpdateContainerGroupRequestContainer {
	s.Command = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPath(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerExec(v []*string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerExec = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPath(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetPort(v []*UpdateContainerGroupRequestContainerPort) *UpdateContainerGroupRequestContainer {
	s.Port = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetScheme(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHttpHeaders(v []*UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHttpHeaders = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetGpu(v int32) *UpdateContainerGroupRequestContainer {
	s.Gpu = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerExec(v []*string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerExec = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetMemory(v float32) *UpdateContainerGroupRequestContainer {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetName(v string) *UpdateContainerGroupRequestContainer {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetImage(v string) *UpdateContainerGroupRequestContainer {
	s.Image = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHttpHeader(v []*UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHttpHeader = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetCpu(v float32) *UpdateContainerGroupRequestContainer {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbe struct {
	TimeoutSeconds      *int32                                                       `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	SuccessThreshold    *int32                                                       `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	HttpGet             *UpdateContainerGroupRequestContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	PeriodSeconds       *int32                                                       `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	InitialDelaySeconds *int32                                                       `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	Exec                *UpdateContainerGroupRequestContainerReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                       `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbe) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetTimeoutSeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetSuccessThreshold(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetTcpSocket(v *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetHttpGet(v *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetPeriodSeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetInitialDelaySeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetExec(v *UpdateContainerGroupRequestContainerReadinessProbeExec) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.Exec = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetFailureThreshold(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.FailureThreshold = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) SetPort(v int32) *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeHttpGet struct {
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetScheme(v string) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetPath(v string) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetPort(v int32) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeExec) SetCommand(v []*string) *UpdateContainerGroupRequestContainerReadinessProbeExec {
	s.Command = v
	return s
}

type UpdateContainerGroupRequestContainerSecurityContext struct {
	Capability             *UpdateContainerGroupRequestContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                          `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                         `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s UpdateContainerGroupRequestContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetCapability(v *UpdateContainerGroupRequestContainerSecurityContextCapability) *UpdateContainerGroupRequestContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *UpdateContainerGroupRequestContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetRunAsUser(v int64) *UpdateContainerGroupRequestContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type UpdateContainerGroupRequestContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerSecurityContextCapability) SetAdd(v []*string) *UpdateContainerGroupRequestContainerSecurityContextCapability {
	s.Add = v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbe struct {
	PeriodSeconds       *int32                                                      `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	TcpSocket           *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                      `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	SuccessThreshold    *int32                                                      `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	Exec                *UpdateContainerGroupRequestContainerLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	HttpGet             *UpdateContainerGroupRequestContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                      `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	TimeoutSeconds      *int32                                                      `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbe) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetPeriodSeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetTcpSocket(v *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetInitialDelaySeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetSuccessThreshold(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetExec(v *UpdateContainerGroupRequestContainerLivenessProbeExec) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.Exec = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetHttpGet(v *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetFailureThreshold(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetTimeoutSeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) SetPort(v int32) *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeExec) SetCommand(v []*string) *UpdateContainerGroupRequestContainerLivenessProbeExec {
	s.Command = v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeHttpGet struct {
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetScheme(v string) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetPort(v int32) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetPath(v string) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Path = &v
	return s
}

type UpdateContainerGroupRequestContainerEnvironmentVar struct {
	FieldRef *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                     `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetFieldRef(v *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetKey(v string) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetValue(v string) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) SetFieldPath(v string) *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type UpdateContainerGroupRequestContainerVolumeMount struct {
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateContainerGroupRequestContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetMountPropagation(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetMountPath(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetReadOnly(v bool) *UpdateContainerGroupRequestContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetSubPath(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.SubPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetName(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.Name = &v
	return s
}

type UpdateContainerGroupRequestContainerPort struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestContainerPort) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerPort) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerPort) SetProtocol(v string) *UpdateContainerGroupRequestContainerPort {
	s.Protocol = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerPort) SetPort(v int32) *UpdateContainerGroupRequestContainerPort {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) SetValue(v string) *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders {
	s.Value = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) SetName(v string) *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders {
	s.Name = &v
	return s
}

type UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetValue(v string) *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetName(v string) *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

type UpdateContainerGroupRequestInitContainer struct {
	SecurityContext *UpdateContainerGroupRequestInitContainerSecurityContext  `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Image           *string                                                   `json:"Image,omitempty" xml:"Image,omitempty"`
	VolumeMount     []*UpdateContainerGroupRequestInitContainerVolumeMount    `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	Port            []*UpdateContainerGroupRequestInitContainerPort           `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	EnvironmentVar  []*UpdateContainerGroupRequestInitContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	ImagePullPolicy *string                                                   `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	StdinOnce       *bool                                                     `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	Cpu             *float32                                                  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Tty             *bool                                                     `json:"Tty,omitempty" xml:"Tty,omitempty"`
	WorkingDir      *string                                                   `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	Command         []*string                                                 `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Arg             []*string                                                 `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Gpu             *int32                                                    `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory          *float32                                                  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Stdin           *bool                                                     `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	Name            *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainer) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainer) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainer) SetSecurityContext(v *UpdateContainerGroupRequestInitContainerSecurityContext) *UpdateContainerGroupRequestInitContainer {
	s.SecurityContext = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetImage(v string) *UpdateContainerGroupRequestInitContainer {
	s.Image = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetVolumeMount(v []*UpdateContainerGroupRequestInitContainerVolumeMount) *UpdateContainerGroupRequestInitContainer {
	s.VolumeMount = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetPort(v []*UpdateContainerGroupRequestInitContainerPort) *UpdateContainerGroupRequestInitContainer {
	s.Port = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetEnvironmentVar(v []*UpdateContainerGroupRequestInitContainerEnvironmentVar) *UpdateContainerGroupRequestInitContainer {
	s.EnvironmentVar = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetImagePullPolicy(v string) *UpdateContainerGroupRequestInitContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetStdinOnce(v bool) *UpdateContainerGroupRequestInitContainer {
	s.StdinOnce = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetCpu(v float32) *UpdateContainerGroupRequestInitContainer {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetTty(v bool) *UpdateContainerGroupRequestInitContainer {
	s.Tty = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetWorkingDir(v string) *UpdateContainerGroupRequestInitContainer {
	s.WorkingDir = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetCommand(v []*string) *UpdateContainerGroupRequestInitContainer {
	s.Command = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetArg(v []*string) *UpdateContainerGroupRequestInitContainer {
	s.Arg = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetGpu(v int32) *UpdateContainerGroupRequestInitContainer {
	s.Gpu = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetMemory(v float32) *UpdateContainerGroupRequestInitContainer {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetStdin(v bool) *UpdateContainerGroupRequestInitContainer {
	s.Stdin = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetName(v string) *UpdateContainerGroupRequestInitContainer {
	s.Name = &v
	return s
}

type UpdateContainerGroupRequestInitContainerSecurityContext struct {
	Capability             *UpdateContainerGroupRequestInitContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                              `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                             `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetCapability(v *UpdateContainerGroupRequestInitContainerSecurityContextCapability) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetRunAsUser(v int64) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type UpdateContainerGroupRequestInitContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestInitContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContextCapability) SetAdd(v []*string) *UpdateContainerGroupRequestInitContainerSecurityContextCapability {
	s.Add = v
	return s
}

type UpdateContainerGroupRequestInitContainerVolumeMount struct {
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetMountPropagation(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetMountPath(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetReadOnly(v bool) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetSubPath(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.SubPath = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetName(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.Name = &v
	return s
}

type UpdateContainerGroupRequestInitContainerPort struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerPort) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerPort) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerPort) SetProtocol(v string) *UpdateContainerGroupRequestInitContainerPort {
	s.Protocol = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerPort) SetPort(v int32) *UpdateContainerGroupRequestInitContainerPort {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestInitContainerEnvironmentVar struct {
	FieldRef *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                         `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                         `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetFieldRef(v *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetKey(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetValue(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) SetFieldPath(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type UpdateContainerGroupRequestImageRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateContainerGroupRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetPassword(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetServer(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetUserName(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type UpdateContainerGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupResponseBody) SetRequestId(v string) *UpdateContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateContainerGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupResponse) SetHeaders(v map[string]*string) *UpdateContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateContainerGroupResponse) SetBody(v *UpdateContainerGroupResponseBody) *UpdateContainerGroupResponse {
	s.Body = v
	return s
}

type ExecContainerCommandRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	ContainerName        *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	Command              *string `json:"Command,omitempty" xml:"Command,omitempty"`
	TTY                  *bool   `json:"TTY,omitempty" xml:"TTY,omitempty"`
	Stdin                *bool   `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	Sync                 *bool   `json:"Sync,omitempty" xml:"Sync,omitempty"`
}

func (s ExecContainerCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandRequest) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandRequest) SetOwnerId(v int64) *ExecContainerCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetResourceOwnerAccount(v string) *ExecContainerCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExecContainerCommandRequest) SetResourceOwnerId(v int64) *ExecContainerCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetOwnerAccount(v string) *ExecContainerCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ExecContainerCommandRequest) SetRegionId(v string) *ExecContainerCommandRequest {
	s.RegionId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetContainerGroupId(v string) *ExecContainerCommandRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetContainerName(v string) *ExecContainerCommandRequest {
	s.ContainerName = &v
	return s
}

func (s *ExecContainerCommandRequest) SetCommand(v string) *ExecContainerCommandRequest {
	s.Command = &v
	return s
}

func (s *ExecContainerCommandRequest) SetTTY(v bool) *ExecContainerCommandRequest {
	s.TTY = &v
	return s
}

func (s *ExecContainerCommandRequest) SetStdin(v bool) *ExecContainerCommandRequest {
	s.Stdin = &v
	return s
}

func (s *ExecContainerCommandRequest) SetSync(v bool) *ExecContainerCommandRequest {
	s.Sync = &v
	return s
}

type ExecContainerCommandResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebSocketUri *string `json:"WebSocketUri,omitempty" xml:"WebSocketUri,omitempty"`
	HttpUrl      *string `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	SyncResponse *string `json:"SyncResponse,omitempty" xml:"SyncResponse,omitempty"`
}

func (s ExecContainerCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandResponseBody) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandResponseBody) SetRequestId(v string) *ExecContainerCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetWebSocketUri(v string) *ExecContainerCommandResponseBody {
	s.WebSocketUri = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetHttpUrl(v string) *ExecContainerCommandResponseBody {
	s.HttpUrl = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetSyncResponse(v string) *ExecContainerCommandResponseBody {
	s.SyncResponse = &v
	return s
}

type ExecContainerCommandResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecContainerCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecContainerCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandResponse) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandResponse) SetHeaders(v map[string]*string) *ExecContainerCommandResponse {
	s.Headers = v
	return s
}

func (s *ExecContainerCommandResponse) SetBody(v *ExecContainerCommandResponseBody) *ExecContainerCommandResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupPriceRequest struct {
	OwnerId              *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Cpu                  *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory               *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	InstanceType         *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotStrategy         *string  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	ZoneId               *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SpotPriceLimit       *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	EphemeralStorage     *int32   `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
}

func (s DescribeContainerGroupPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceRequest) SetOwnerId(v int64) *DescribeContainerGroupPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupPriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupPriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetOwnerAccount(v string) *DescribeContainerGroupPriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetRegionId(v string) *DescribeContainerGroupPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetCpu(v float32) *DescribeContainerGroupPriceRequest {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetMemory(v float32) *DescribeContainerGroupPriceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetInstanceType(v string) *DescribeContainerGroupPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetSpotStrategy(v string) *DescribeContainerGroupPriceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetZoneId(v string) *DescribeContainerGroupPriceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetSpotPriceLimit(v float32) *DescribeContainerGroupPriceRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetEphemeralStorage(v int32) *DescribeContainerGroupPriceRequest {
	s.EphemeralStorage = &v
	return s
}

type DescribeContainerGroupPriceResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PriceInfo *DescribeContainerGroupPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBody) SetRequestId(v string) *DescribeContainerGroupPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBody) SetPriceInfo(v *DescribeContainerGroupPriceResponseBodyPriceInfo) *DescribeContainerGroupPriceResponseBody {
	s.PriceInfo = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfo struct {
	SpotPrices *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices `json:"SpotPrices,omitempty" xml:"SpotPrices,omitempty" type:"Struct"`
	Price      *DescribeContainerGroupPriceResponseBodyPriceInfoPrice      `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules      *DescribeContainerGroupPriceResponseBodyPriceInfoRules      `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetSpotPrices(v *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.SpotPrices = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetPrice(v *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetRules(v *DescribeContainerGroupPriceResponseBodyPriceInfoRules) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices struct {
	SpotPrice []*DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) SetSpotPrice(v []*DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices {
	s.SpotPrice = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice struct {
	ZoneId       *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SpotPrice    *float32 `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty"`
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OriginPrice  *float32 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetZoneId(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.ZoneId = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetSpotPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.SpotPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetInstanceType(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetOriginPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.OriginPrice = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPrice struct {
	DiscountPrice *float32                                                          `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	TradePrice    *float32                                                          `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	OriginalPrice *float32                                                          `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	DetailInfos   *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Struct"`
	Currency      *string                                                           `json:"Currency,omitempty" xml:"Currency,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetDetailInfos(v *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.DetailInfos = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos struct {
	DetailInfo []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) SetDetailInfo(v []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos {
	s.DetailInfo = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo struct {
	Resource      *string                                                                          `json:"Resource,omitempty" xml:"Resource,omitempty"`
	DiscountPrice *float32                                                                         `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	TradePrice    *float32                                                                         `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	OriginalPrice *float32                                                                         `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Rules         *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetResource(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Resource = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetDiscountPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetTradePrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.TradePrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetOriginalPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetRules(v *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Rules = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules struct {
	Rule []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) SetRule(v []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules {
	s.Rule = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) SetDescription(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) SetRuleId(v int64) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule {
	s.RuleId = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoRules struct {
	Rule []*DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRules) SetRule(v []*DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) *DescribeContainerGroupPriceResponseBodyPriceInfoRules {
	s.Rule = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) SetDescription(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) SetRuleId(v int64) *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule {
	s.RuleId = &v
	return s
}

type DescribeContainerGroupPriceResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContainerGroupPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupPriceResponse) SetBody(v *DescribeContainerGroupPriceResponseBody) *DescribeContainerGroupPriceResponse {
	s.Body = v
	return s
}

type DescribeMultiContainerGroupMetricRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ContainerGroupIds    *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	MetricType           *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s DescribeMultiContainerGroupMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricRequest) SetOwnerId(v int64) *DescribeMultiContainerGroupMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceOwnerAccount(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceOwnerId(v int64) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetOwnerAccount(v string) *DescribeMultiContainerGroupMetricRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetRegionId(v string) *DescribeMultiContainerGroupMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetContainerGroupIds(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceGroupId(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetMetricType(v string) *DescribeMultiContainerGroupMetricRequest {
	s.MetricType = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBody struct {
	RequestId    *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MonitorDatas []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatas `json:"MonitorDatas,omitempty" xml:"MonitorDatas,omitempty" type:"Repeated"`
}

func (s DescribeMultiContainerGroupMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBody) SetRequestId(v string) *DescribeMultiContainerGroupMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBody) SetMonitorDatas(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) *DescribeMultiContainerGroupMetricResponseBody {
	s.MonitorDatas = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatas struct {
	Records          []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords        `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	ContainerInfos   []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos `json:"ContainerInfos,omitempty" xml:"ContainerInfos,omitempty" type:"Repeated"`
	ContainerGroupId *string                                                                    `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) SetRecords(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas {
	s.Records = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) SetContainerInfos(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas {
	s.ContainerInfos = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) SetContainerGroupId(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas {
	s.ContainerGroupId = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords struct {
	Network    *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork      `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	CPU        *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU          `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	Disk       []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk       `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	Timestamp  *string                                                                       `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Memory     *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory       `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Containers []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	Filesystem []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem `json:"Filesystem,omitempty" xml:"Filesystem,omitempty" type:"Repeated"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetNetwork(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Network = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetCPU(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.CPU = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetDisk(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Disk = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetTimestamp(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Timestamp = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetMemory(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Memory = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetContainers(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Containers = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetFilesystem(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Filesystem = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork struct {
	Interfaces []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces `json:"Interfaces,omitempty" xml:"Interfaces,omitempty" type:"Repeated"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) SetInterfaces(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork {
	s.Interfaces = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces struct {
	RxErrors  *int64  `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	TxDrops   *int64  `json:"TxDrops,omitempty" xml:"TxDrops,omitempty"`
	TxBytes   *int64  `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	RxPackets *int64  `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	TxPackets *int64  `json:"TxPackets,omitempty" xml:"TxPackets,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TxErrors  *int64  `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	RxBytes   *int64  `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	RxDrops   *int64  `json:"RxDrops,omitempty" xml:"RxDrops,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxDrops(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxDrops = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxPackets = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxPackets = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.Name = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxDrops(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxDrops = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU struct {
	UsageNanoCores       *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
	Limit                *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	Load                 *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetUsageNanoCores(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.UsageNanoCores = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetUsageCoreNanoSeconds(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetLoad(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.Load = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk struct {
	WriteBytes *int64  `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	Device     *string `json:"Device,omitempty" xml:"Device,omitempty"`
	WriteIo    *int64  `json:"WriteIo,omitempty" xml:"WriteIo,omitempty"`
	ReadBytes  *int64  `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	ReadIo     *int64  `json:"ReadIo,omitempty" xml:"ReadIo,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetWriteBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.WriteBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.Device = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetWriteIo(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.WriteIo = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetReadBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.ReadBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetReadIo(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.ReadIo = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory struct {
	Rss            *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	UsageBytes     *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	WorkingSet     *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	Cache          *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetRss(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.Rss = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetUsageBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetWorkingSet(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.WorkingSet = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetAvailableBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetCache(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.Cache = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers struct {
	CPU    *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU    `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	Memory *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Name   *string                                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetCPU(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.CPU = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetMemory(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.Memory = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.Name = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU struct {
	UsageNanoCores       *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
	Limit                *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	Load                 *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetUsageNanoCores(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.UsageNanoCores = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetUsageCoreNanoSeconds(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetLoad(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.Load = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory struct {
	Rss            *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	UsageBytes     *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	WorkingSet     *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	Cache          *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetRss(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.Rss = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetUsageBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetWorkingSet(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.WorkingSet = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetAvailableBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetCache(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.Cache = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem struct {
	Capacity  *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Available *int64  `json:"Available,omitempty" xml:"Available,omitempty"`
	FsName    *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	Usage     *int64  `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetCapacity(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Capacity = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetAvailable(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Available = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetFsName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.FsName = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetUsage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Usage = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos struct {
	Aliases        []*string                                                                                `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Repeated"`
	ContainerSpec  *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec    `json:"ContainerSpec,omitempty" xml:"ContainerSpec,omitempty" type:"Struct"`
	ContainerStats []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats `json:"ContainerStats,omitempty" xml:"ContainerStats,omitempty" type:"Repeated"`
	Labels         *string                                                                                  `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Namespace      *string                                                                                  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Name           *string                                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Id             *string                                                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) SetAliases(v []*string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos {
	s.Aliases = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) SetContainerSpec(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos {
	s.ContainerSpec = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) SetContainerStats(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos {
	s.ContainerStats = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) SetLabels(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos {
	s.Labels = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) SetNamespace(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos {
	s.Namespace = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos {
	s.Name = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos) SetId(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfos {
	s.Id = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec struct {
	CreationTime     *string                                                                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Image            *string                                                                                              `json:"Image,omitempty" xml:"Image,omitempty"`
	Labels           *string                                                                                              `json:"Labels,omitempty" xml:"Labels,omitempty"`
	HasCustomMetrics *bool                                                                                                `json:"HasCustomMetrics,omitempty" xml:"HasCustomMetrics,omitempty"`
	HasCpu           *bool                                                                                                `json:"HasCpu,omitempty" xml:"HasCpu,omitempty"`
	Envs             *string                                                                                              `json:"Envs,omitempty" xml:"Envs,omitempty"`
	HasDiskIo        *bool                                                                                                `json:"HasDiskIo,omitempty" xml:"HasDiskIo,omitempty"`
	HasFilesystem    *bool                                                                                                `json:"HasFilesystem,omitempty" xml:"HasFilesystem,omitempty"`
	HasNetwork       *bool                                                                                                `json:"HasNetwork,omitempty" xml:"HasNetwork,omitempty"`
	ContainerMemory  *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory `json:"ContainerMemory,omitempty" xml:"ContainerMemory,omitempty" type:"Struct"`
	HasMemory        *bool                                                                                                `json:"HasMemory,omitempty" xml:"HasMemory,omitempty"`
	ContainerCpu     *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu    `json:"ContainerCpu,omitempty" xml:"ContainerCpu,omitempty" type:"Struct"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetCreationTime(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.CreationTime = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetImage(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.Image = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetLabels(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.Labels = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetHasCustomMetrics(v bool) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.HasCustomMetrics = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetHasCpu(v bool) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.HasCpu = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetEnvs(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.Envs = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetHasDiskIo(v bool) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.HasDiskIo = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetHasFilesystem(v bool) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.HasFilesystem = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetHasNetwork(v bool) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.HasNetwork = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetContainerMemory(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.ContainerMemory = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetHasMemory(v bool) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.HasMemory = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec) SetContainerCpu(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpec {
	s.ContainerCpu = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory struct {
	Limit       *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	SwapLimit   *int64 `json:"SwapLimit,omitempty" xml:"SwapLimit,omitempty"`
	Reservation *int64 `json:"Reservation,omitempty" xml:"Reservation,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory) SetSwapLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory {
	s.SwapLimit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory) SetReservation(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerMemory {
	s.Reservation = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu struct {
	Quota    *int64  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	Mask     *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Limit    *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	MaxLimit *int64  `json:"MaxLimit,omitempty" xml:"MaxLimit,omitempty"`
	Period   *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu) SetQuota(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu {
	s.Quota = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu) SetMask(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu {
	s.Mask = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu) SetMaxLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu {
	s.MaxLimit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu) SetPeriod(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerSpecContainerCpu {
	s.Period = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats struct {
	NetworkStats     *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats       `json:"NetworkStats,omitempty" xml:"NetworkStats,omitempty" type:"Struct"`
	FsStats          []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats          `json:"FsStats,omitempty" xml:"FsStats,omitempty" type:"Repeated"`
	AcceleratorStats []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats `json:"AcceleratorStats,omitempty" xml:"AcceleratorStats,omitempty" type:"Repeated"`
	CpuStats         *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats           `json:"CpuStats,omitempty" xml:"CpuStats,omitempty" type:"Struct"`
	Timestamp        *string                                                                                                  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	MemoryStats      *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats        `json:"MemoryStats,omitempty" xml:"MemoryStats,omitempty" type:"Struct"`
	TaskStats        *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats          `json:"TaskStats,omitempty" xml:"TaskStats,omitempty" type:"Struct"`
	DiskIoStats      *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats        `json:"DiskIoStats,omitempty" xml:"DiskIoStats,omitempty" type:"Struct"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) SetNetworkStats(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats {
	s.NetworkStats = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) SetFsStats(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats {
	s.FsStats = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) SetAcceleratorStats(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats {
	s.AcceleratorStats = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) SetCpuStats(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats {
	s.CpuStats = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) SetTimestamp(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats {
	s.Timestamp = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) SetMemoryStats(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats {
	s.MemoryStats = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) SetTaskStats(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats {
	s.TaskStats = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats) SetDiskIoStats(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStats {
	s.DiskIoStats = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats struct {
	RxDropped      *int64                                                                                                             `json:"RxDropped,omitempty" xml:"RxDropped,omitempty"`
	Tcp            *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp              `json:"Tcp,omitempty" xml:"Tcp,omitempty" type:"Struct"`
	TxErrors       *int64                                                                                                             `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	RxErrors       *int64                                                                                                             `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	InterfaceStats []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats `json:"InterfaceStats,omitempty" xml:"InterfaceStats,omitempty" type:"Repeated"`
	Tcp6           *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6             `json:"Tcp6,omitempty" xml:"Tcp6,omitempty" type:"Struct"`
	TxDropped      *int64                                                                                                             `json:"TxDropped,omitempty" xml:"TxDropped,omitempty"`
	Udp            *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp              `json:"Udp,omitempty" xml:"Udp,omitempty" type:"Struct"`
	TxBytes        *int64                                                                                                             `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	Udp6           *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6             `json:"Udp6,omitempty" xml:"Udp6,omitempty" type:"Struct"`
	RxPackets      *int64                                                                                                             `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	Name           *string                                                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	RxBytes        *int64                                                                                                             `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	TxPackets      *int64                                                                                                             `json:"TxPackets,omitempty" xml:"TxPackets,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetRxDropped(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.RxDropped = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetTcp(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.Tcp = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetTxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.TxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetRxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.RxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetInterfaceStats(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.InterfaceStats = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetTcp6(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.Tcp6 = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetTxDropped(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.TxDropped = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetUdp(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.Udp = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetTxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.TxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetUdp6(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.Udp6 = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetRxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.RxPackets = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.Name = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetRxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.RxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats) SetTxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStats {
	s.TxPackets = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp struct {
	CloseWait   *int64 `json:"CloseWait,omitempty" xml:"CloseWait,omitempty"`
	FinWait2    *int64 `json:"FinWait2,omitempty" xml:"FinWait2,omitempty"`
	LastAck     *int64 `json:"LastAck,omitempty" xml:"LastAck,omitempty"`
	Closing     *int64 `json:"Closing,omitempty" xml:"Closing,omitempty"`
	Listen      *int64 `json:"Listen,omitempty" xml:"Listen,omitempty"`
	TimeWait    *int64 `json:"TimeWait,omitempty" xml:"TimeWait,omitempty"`
	Established *int64 `json:"Established,omitempty" xml:"Established,omitempty"`
	FinWait1    *int64 `json:"FinWait1,omitempty" xml:"FinWait1,omitempty"`
	Close       *int64 `json:"Close,omitempty" xml:"Close,omitempty"`
	SynRecv     *int64 `json:"SynRecv,omitempty" xml:"SynRecv,omitempty"`
	SynSent     *int64 `json:"SynSent,omitempty" xml:"SynSent,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetCloseWait(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.CloseWait = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetFinWait2(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.FinWait2 = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetLastAck(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.LastAck = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetClosing(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.Closing = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetListen(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.Listen = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetTimeWait(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.TimeWait = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetEstablished(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.Established = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetFinWait1(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.FinWait1 = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetClose(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.Close = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetSynRecv(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.SynRecv = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp) SetSynSent(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp {
	s.SynSent = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats struct {
	RxErrors  *int64  `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	RxDropped *int64  `json:"RxDropped,omitempty" xml:"RxDropped,omitempty"`
	TxDropped *int64  `json:"TxDropped,omitempty" xml:"TxDropped,omitempty"`
	TxBytes   *int64  `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	RxPackets *int64  `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	TxErrors  *int64  `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	TxPackets *int64  `json:"TxPackets,omitempty" xml:"TxPackets,omitempty"`
	RxBytes   *int64  `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetRxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.RxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetRxDropped(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.RxDropped = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetTxDropped(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.TxDropped = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetTxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.TxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetRxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.RxPackets = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetTxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.TxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetTxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.TxPackets = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetRxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.RxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsInterfaceStats {
	s.Name = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 struct {
	CloseWait   *int64 `json:"CloseWait,omitempty" xml:"CloseWait,omitempty"`
	FinWait2    *int64 `json:"FinWait2,omitempty" xml:"FinWait2,omitempty"`
	LastAck     *int64 `json:"LastAck,omitempty" xml:"LastAck,omitempty"`
	Closing     *int64 `json:"Closing,omitempty" xml:"Closing,omitempty"`
	Listen      *int64 `json:"Listen,omitempty" xml:"Listen,omitempty"`
	TimeWait    *int64 `json:"TimeWait,omitempty" xml:"TimeWait,omitempty"`
	Established *int64 `json:"Established,omitempty" xml:"Established,omitempty"`
	FinWait1    *int64 `json:"FinWait1,omitempty" xml:"FinWait1,omitempty"`
	Close       *int64 `json:"Close,omitempty" xml:"Close,omitempty"`
	SynRecv     *int64 `json:"SynRecv,omitempty" xml:"SynRecv,omitempty"`
	SynSent     *int64 `json:"SynSent,omitempty" xml:"SynSent,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetCloseWait(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.CloseWait = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetFinWait2(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.FinWait2 = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetLastAck(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.LastAck = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetClosing(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.Closing = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetListen(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.Listen = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetTimeWait(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.TimeWait = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetEstablished(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.Established = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetFinWait1(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.FinWait1 = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetClose(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.Close = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetSynRecv(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.SynRecv = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6) SetSynSent(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsTcp6 {
	s.SynSent = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp struct {
	TxQueued *int64 `json:"TxQueued,omitempty" xml:"TxQueued,omitempty"`
	Listen   *int64 `json:"Listen,omitempty" xml:"Listen,omitempty"`
	Dropped  *int64 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	RxQueued *int64 `json:"RxQueued,omitempty" xml:"RxQueued,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp) SetTxQueued(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp {
	s.TxQueued = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp) SetListen(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp {
	s.Listen = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp) SetDropped(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp {
	s.Dropped = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp) SetRxQueued(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp {
	s.RxQueued = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6 struct {
	TxQueued *int64 `json:"TxQueued,omitempty" xml:"TxQueued,omitempty"`
	Listen   *int64 `json:"Listen,omitempty" xml:"Listen,omitempty"`
	Dropped  *int64 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	RxQueued *int64 `json:"RxQueued,omitempty" xml:"RxQueued,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6) SetTxQueued(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6 {
	s.TxQueued = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6) SetListen(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6 {
	s.Listen = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6) SetDropped(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6 {
	s.Dropped = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6) SetRxQueued(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsNetworkStatsUdp6 {
	s.RxQueued = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats struct {
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ReadsMerged     *int64  `json:"ReadsMerged,omitempty" xml:"ReadsMerged,omitempty"`
	HasInodes       *bool   `json:"HasInodes,omitempty" xml:"HasInodes,omitempty"`
	ReadsCompleted  *int64  `json:"ReadsCompleted,omitempty" xml:"ReadsCompleted,omitempty"`
	WritesMerged    *int64  `json:"WritesMerged,omitempty" xml:"WritesMerged,omitempty"`
	InodesFree      *int64  `json:"InodesFree,omitempty" xml:"InodesFree,omitempty"`
	Available       *int64  `json:"Available,omitempty" xml:"Available,omitempty"`
	Usage           *int64  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Inodes          *int64  `json:"Inodes,omitempty" xml:"Inodes,omitempty"`
	BaseUsage       *int64  `json:"BaseUsage,omitempty" xml:"BaseUsage,omitempty"`
	SectorsRead     *int64  `json:"SectorsRead,omitempty" xml:"SectorsRead,omitempty"`
	WriteTime       *int64  `json:"WriteTime,omitempty" xml:"WriteTime,omitempty"`
	SectorsWritten  *int64  `json:"SectorsWritten,omitempty" xml:"SectorsWritten,omitempty"`
	ReadTime        *int64  `json:"ReadTime,omitempty" xml:"ReadTime,omitempty"`
	Limit           *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Device          *string `json:"Device,omitempty" xml:"Device,omitempty"`
	WritesCompleted *int64  `json:"WritesCompleted,omitempty" xml:"WritesCompleted,omitempty"`
	IoTime          *int64  `json:"IoTime,omitempty" xml:"IoTime,omitempty"`
	WeightedIoTime  *int64  `json:"WeightedIoTime,omitempty" xml:"WeightedIoTime,omitempty"`
	IoInProgress    *int64  `json:"IoInProgress,omitempty" xml:"IoInProgress,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetType(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.Type = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetReadsMerged(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.ReadsMerged = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetHasInodes(v bool) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.HasInodes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetReadsCompleted(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.ReadsCompleted = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetWritesMerged(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.WritesMerged = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetInodesFree(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.InodesFree = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetAvailable(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.Available = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetUsage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.Usage = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetInodes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.Inodes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetBaseUsage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.BaseUsage = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetSectorsRead(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.SectorsRead = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetWriteTime(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.WriteTime = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetSectorsWritten(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.SectorsWritten = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetReadTime(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.ReadTime = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.Device = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetWritesCompleted(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.WritesCompleted = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetIoTime(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.IoTime = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetWeightedIoTime(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.WeightedIoTime = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats) SetIoInProgress(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsFsStats {
	s.IoInProgress = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats struct {
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Minor       *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Temperature *int64  `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	PowerUsage  *int64  `json:"PowerUsage,omitempty" xml:"PowerUsage,omitempty"`
	MemoryTotal *int64  `json:"MemoryTotal,omitempty" xml:"MemoryTotal,omitempty"`
	Make        *string `json:"Make,omitempty" xml:"Make,omitempty"`
	DutyCycle   *int64  `json:"DutyCycle,omitempty" xml:"DutyCycle,omitempty"`
	MemoryUsed  *int64  `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetModel(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.Model = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetTemperature(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.Temperature = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetPowerUsage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.PowerUsage = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetMemoryTotal(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.MemoryTotal = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetMake(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.Make = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetDutyCycle(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.DutyCycle = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetMemoryUsed(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats) SetId(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsAcceleratorStats {
	s.Id = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats struct {
	CpuUsage    *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty" type:"Struct"`
	LoadAverage *int64                                                                                                 `json:"LoadAverage,omitempty" xml:"LoadAverage,omitempty"`
	CpuCFS      *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS   `json:"CpuCFS,omitempty" xml:"CpuCFS,omitempty" type:"Struct"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats) SetCpuUsage(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats {
	s.CpuUsage = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats) SetLoadAverage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats {
	s.LoadAverage = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats) SetCpuCFS(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStats {
	s.CpuCFS = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage struct {
	User         *int64   `json:"User,omitempty" xml:"User,omitempty"`
	PerCpuUsages []*int64 `json:"PerCpuUsages,omitempty" xml:"PerCpuUsages,omitempty" type:"Repeated"`
	Total        *int64   `json:"Total,omitempty" xml:"Total,omitempty"`
	System       *int64   `json:"System,omitempty" xml:"System,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage) SetUser(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage {
	s.User = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage) SetPerCpuUsages(v []*int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage {
	s.PerCpuUsages = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage) SetTotal(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage {
	s.Total = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage) SetSystem(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuUsage {
	s.System = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS struct {
	ThrottledTime    *int64 `json:"ThrottledTime,omitempty" xml:"ThrottledTime,omitempty"`
	Periods          *int64 `json:"Periods,omitempty" xml:"Periods,omitempty"`
	ThrottledPeriods *int64 `json:"ThrottledPeriods,omitempty" xml:"ThrottledPeriods,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS) SetThrottledTime(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS {
	s.ThrottledTime = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS) SetPeriods(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS {
	s.Periods = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS) SetThrottledPeriods(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsCpuStatsCpuCFS {
	s.ThrottledPeriods = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats struct {
	FailCnt          *int64                                                                                                            `json:"FailCnt,omitempty" xml:"FailCnt,omitempty"`
	MaxUsage         *int64                                                                                                            `json:"MaxUsage,omitempty" xml:"MaxUsage,omitempty"`
	HierarchicalData *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData `json:"HierarchicalData,omitempty" xml:"HierarchicalData,omitempty" type:"Struct"`
	Rss              *int64                                                                                                            `json:"Rss,omitempty" xml:"Rss,omitempty"`
	ContainerData    *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData    `json:"ContainerData,omitempty" xml:"ContainerData,omitempty" type:"Struct"`
	WorkingSet       *int64                                                                                                            `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
	Swap             *int64                                                                                                            `json:"Swap,omitempty" xml:"Swap,omitempty"`
	Cache            *int64                                                                                                            `json:"Cache,omitempty" xml:"Cache,omitempty"`
	Usage            *int64                                                                                                            `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetFailCnt(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.FailCnt = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetMaxUsage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.MaxUsage = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetHierarchicalData(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.HierarchicalData = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetRss(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.Rss = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetContainerData(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.ContainerData = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetWorkingSet(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.WorkingSet = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetSwap(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.Swap = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetCache(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.Cache = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats) SetUsage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStats {
	s.Usage = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData struct {
	PgmajFault *int64 `json:"PgmajFault,omitempty" xml:"PgmajFault,omitempty"`
	PgFault    *int64 `json:"PgFault,omitempty" xml:"PgFault,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData) SetPgmajFault(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData {
	s.PgmajFault = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData) SetPgFault(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsHierarchicalData {
	s.PgFault = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData struct {
	PgmajFault *int64 `json:"PgmajFault,omitempty" xml:"PgmajFault,omitempty"`
	PgFault    *int64 `json:"PgFault,omitempty" xml:"PgFault,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData) SetPgmajFault(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData {
	s.PgmajFault = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData) SetPgFault(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsMemoryStatsContainerData {
	s.PgFault = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats struct {
	NrUninterruptible *int64 `json:"NrUninterruptible,omitempty" xml:"NrUninterruptible,omitempty"`
	NrRunning         *int64 `json:"NrRunning,omitempty" xml:"NrRunning,omitempty"`
	NrSleeping        *int64 `json:"NrSleeping,omitempty" xml:"NrSleeping,omitempty"`
	NrIoWait          *int64 `json:"NrIoWait,omitempty" xml:"NrIoWait,omitempty"`
	NrStopped         *int64 `json:"NrStopped,omitempty" xml:"NrStopped,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats) SetNrUninterruptible(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats {
	s.NrUninterruptible = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats) SetNrRunning(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats {
	s.NrRunning = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats) SetNrSleeping(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats {
	s.NrSleeping = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats) SetNrIoWait(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats {
	s.NrIoWait = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats) SetNrStopped(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsTaskStats {
	s.NrStopped = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats struct {
	IoServiced     []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced     `json:"IoServiced,omitempty" xml:"IoServiced,omitempty" type:"Repeated"`
	IoServiceTime  []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime  `json:"IoServiceTime,omitempty" xml:"IoServiceTime,omitempty" type:"Repeated"`
	IoServiceBytes []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes `json:"IoServiceBytes,omitempty" xml:"IoServiceBytes,omitempty" type:"Repeated"`
	IoMerged       []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged       `json:"IoMerged,omitempty" xml:"IoMerged,omitempty" type:"Repeated"`
	Sectors        []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors        `json:"Sectors,omitempty" xml:"Sectors,omitempty" type:"Repeated"`
	IoQueued       []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued       `json:"IoQueued,omitempty" xml:"IoQueued,omitempty" type:"Repeated"`
	IoTime         []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime         `json:"IoTime,omitempty" xml:"IoTime,omitempty" type:"Repeated"`
	IoWaitTime     []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime     `json:"IoWaitTime,omitempty" xml:"IoWaitTime,omitempty" type:"Repeated"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) SetIoServiced(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats {
	s.IoServiced = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) SetIoServiceTime(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats {
	s.IoServiceTime = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) SetIoServiceBytes(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats {
	s.IoServiceBytes = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) SetIoMerged(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats {
	s.IoMerged = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) SetSectors(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats {
	s.Sectors = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) SetIoQueued(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats {
	s.IoQueued = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) SetIoTime(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats {
	s.IoTime = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats) SetIoWaitTime(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStats {
	s.IoWaitTime = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced struct {
	Stats  *string `json:"Stats,omitempty" xml:"Stats,omitempty"`
	Minor  *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Major  *int64  `json:"Major,omitempty" xml:"Major,omitempty"`
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced) SetStats(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced {
	s.Stats = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced) SetMajor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced {
	s.Major = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiced {
	s.Device = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime struct {
	Stats  *string `json:"Stats,omitempty" xml:"Stats,omitempty"`
	Minor  *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Major  *int64  `json:"Major,omitempty" xml:"Major,omitempty"`
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime) SetStats(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime {
	s.Stats = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime) SetMajor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime {
	s.Major = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceTime {
	s.Device = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes struct {
	Stats  *string `json:"Stats,omitempty" xml:"Stats,omitempty"`
	Minor  *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Major  *int64  `json:"Major,omitempty" xml:"Major,omitempty"`
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes) SetStats(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes {
	s.Stats = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes) SetMajor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes {
	s.Major = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoServiceBytes {
	s.Device = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged struct {
	Stats  *string `json:"Stats,omitempty" xml:"Stats,omitempty"`
	Minor  *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Major  *int64  `json:"Major,omitempty" xml:"Major,omitempty"`
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged) SetStats(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged {
	s.Stats = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged) SetMajor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged {
	s.Major = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoMerged {
	s.Device = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors struct {
	Stats  *string `json:"Stats,omitempty" xml:"Stats,omitempty"`
	Minor  *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Major  *int64  `json:"Major,omitempty" xml:"Major,omitempty"`
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors) SetStats(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors {
	s.Stats = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors) SetMajor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors {
	s.Major = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsSectors {
	s.Device = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued struct {
	Stats  *string `json:"Stats,omitempty" xml:"Stats,omitempty"`
	Minor  *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Major  *int64  `json:"Major,omitempty" xml:"Major,omitempty"`
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued) SetStats(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued {
	s.Stats = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued) SetMajor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued {
	s.Major = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoQueued {
	s.Device = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime struct {
	Stats  *string `json:"Stats,omitempty" xml:"Stats,omitempty"`
	Minor  *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Major  *int64  `json:"Major,omitempty" xml:"Major,omitempty"`
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime) SetStats(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime {
	s.Stats = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime) SetMajor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime {
	s.Major = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoTime {
	s.Device = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime struct {
	Stats  *string `json:"Stats,omitempty" xml:"Stats,omitempty"`
	Minor  *int64  `json:"Minor,omitempty" xml:"Minor,omitempty"`
	Major  *int64  `json:"Major,omitempty" xml:"Major,omitempty"`
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime) SetStats(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime {
	s.Stats = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime) SetMinor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime {
	s.Minor = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime) SetMajor(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime {
	s.Major = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasContainerInfosContainerStatsDiskIoStatsIoWaitTime {
	s.Device = &v
	return s
}

type DescribeMultiContainerGroupMetricResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMultiContainerGroupMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMultiContainerGroupMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponse) SetHeaders(v map[string]*string) *DescribeMultiContainerGroupMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponse) SetBody(v *DescribeMultiContainerGroupMetricResponseBody) *DescribeMultiContainerGroupMetricResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupsRequest struct {
	OwnerId              *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                               `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VSwitchId            *string                              `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NextToken            *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Limit                *int32                               `json:"Limit,omitempty" xml:"Limit,omitempty"`
	ContainerGroupIds    *string                              `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	ContainerGroupName   *string                              `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	Status               *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	ResourceGroupId      *string                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	WithEvent            *bool                                `json:"WithEvent,omitempty" xml:"WithEvent,omitempty"`
	Tag                  []*DescribeContainerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsRequest) SetOwnerId(v int64) *DescribeContainerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetOwnerAccount(v string) *DescribeContainerGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetRegionId(v string) *DescribeContainerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetZoneId(v string) *DescribeContainerGroupsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetVSwitchId(v string) *DescribeContainerGroupsRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetNextToken(v string) *DescribeContainerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetLimit(v int32) *DescribeContainerGroupsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetContainerGroupIds(v string) *DescribeContainerGroupsRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetContainerGroupName(v string) *DescribeContainerGroupsRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetStatus(v string) *DescribeContainerGroupsRequest {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceGroupId(v string) *DescribeContainerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetWithEvent(v bool) *DescribeContainerGroupsRequest {
	s.WithEvent = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetTag(v []*DescribeContainerGroupsRequestTag) *DescribeContainerGroupsRequest {
	s.Tag = v
	return s
}

type DescribeContainerGroupsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsRequestTag) SetKey(v string) *DescribeContainerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsRequestTag) SetValue(v string) *DescribeContainerGroupsRequestTag {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBody struct {
	TotalCount      *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken       *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContainerGroups []*DescribeContainerGroupsResponseBodyContainerGroups `json:"ContainerGroups,omitempty" xml:"ContainerGroups,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBody) SetTotalCount(v int32) *DescribeContainerGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetNextToken(v string) *DescribeContainerGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetRequestId(v string) *DescribeContainerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetContainerGroups(v []*DescribeContainerGroupsResponseBodyContainerGroups) *DescribeContainerGroupsResponseBody {
	s.ContainerGroups = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroups struct {
	Status                *string                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime          *string                                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	VpcId                 *string                                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	InternetIp            *string                                                               `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	TenantSecurityGroupId *string                                                               `json:"TenantSecurityGroupId,omitempty" xml:"TenantSecurityGroupId,omitempty"`
	SecurityGroupId       *string                                                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	HostAliases           []*DescribeContainerGroupsResponseBodyContainerGroupsHostAliases      `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	Tags                  []*DescribeContainerGroupsResponseBodyContainerGroupsTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Events                []*DescribeContainerGroupsResponseBodyContainerGroupsEvents           `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	SucceededTime         *string                                                               `json:"SucceededTime,omitempty" xml:"SucceededTime,omitempty"`
	SpotStrategy          *string                                                               `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	EphemeralStorage      *int32                                                                `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	TenantEniInstanceId   *string                                                               `json:"TenantEniInstanceId,omitempty" xml:"TenantEniInstanceId,omitempty"`
	Discount              *int32                                                                `json:"Discount,omitempty" xml:"Discount,omitempty"`
	RestartPolicy         *string                                                               `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	Memory                *float32                                                              `json:"Memory,omitempty" xml:"Memory,omitempty"`
	TenantVSwitchId       *string                                                               `json:"TenantVSwitchId,omitempty" xml:"TenantVSwitchId,omitempty"`
	Containers            []*DescribeContainerGroupsResponseBodyContainerGroupsContainers       `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	EniInstanceId         *string                                                               `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	InitContainers        []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainers   `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	ContainerGroupId      *string                                                               `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	TenantEniIp           *string                                                               `json:"TenantEniIp,omitempty" xml:"TenantEniIp,omitempty"`
	InstanceType          *string                                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IntranetIp            *string                                                               `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Ipv6Address           *string                                                               `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	RegionId              *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DnsConfig             *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig          `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	Volumes               []*DescribeContainerGroupsResponseBodyContainerGroupsVolumes          `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	RamRoleName           *string                                                               `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	VSwitchId             *string                                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Cpu                   *float32                                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ExpiredTime           *string                                                               `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ResourceGroupId       *string                                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId                *string                                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ContainerGroupName    *string                                                               `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	EciSecurityContext    *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext `json:"EciSecurityContext,omitempty" xml:"EciSecurityContext,omitempty" type:"Struct"`
	FailedTime            *string                                                               `json:"FailedTime,omitempty" xml:"FailedTime,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroups) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetCreationTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.CreationTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVpcId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.VpcId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInternetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InternetIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantSecurityGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantSecurityGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSecurityGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetHostAliases(v []*DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.HostAliases = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTags(v []*DescribeContainerGroupsResponseBodyContainerGroupsTags) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Tags = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEvents(v []*DescribeContainerGroupsResponseBodyContainerGroupsEvents) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Events = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSucceededTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SucceededTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSpotStrategy(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEphemeralStorage(v int32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EphemeralStorage = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantEniInstanceId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantEniInstanceId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetDiscount(v int32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Discount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRestartPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RestartPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantVSwitchId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantVSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainers(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainers) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Containers = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEniInstanceId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EniInstanceId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInitContainers(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InitContainers = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainerGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantEniIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantEniIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInstanceType(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetIntranetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.IntranetIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetIpv6Address(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRegionId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetDnsConfig(v *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.DnsConfig = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVolumes(v []*DescribeContainerGroupsResponseBodyContainerGroupsVolumes) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Volumes = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRamRoleName(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RamRoleName = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVSwitchId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetExpiredTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetResourceGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetZoneId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ZoneId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainerGroupName(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEciSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EciSecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetFailedTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.FailedTime = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsHostAliases struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	Ip        *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) SetHostnames(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases {
	s.Hostnames = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) SetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases {
	s.Ip = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsTags) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsTags) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsTags {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsTags) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsTags {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEvents struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	LastTimestamp  *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEvents) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Type = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetLastTimestamp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Count = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetFirstTimestamp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.FirstTimestamp = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainers struct {
	LivenessProbe   *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe     `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" type:"Struct"`
	Commands        []*string                                                                      `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	VolumeMounts    []*DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts    `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	Args            []*string                                                                      `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Image           *string                                                                        `json:"Image,omitempty" xml:"Image,omitempty"`
	Ports           []*DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts           `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	RestartCount    *int32                                                                         `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	ImagePullPolicy *string                                                                        `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	StdinOnce       *bool                                                                          `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	Cpu             *float32                                                                       `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	PreviousState   *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState     `json:"PreviousState,omitempty" xml:"PreviousState,omitempty" type:"Struct"`
	Tty             *bool                                                                          `json:"Tty,omitempty" xml:"Tty,omitempty"`
	WorkingDir      *string                                                                        `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	CurrentState    *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState      `json:"CurrentState,omitempty" xml:"CurrentState,omitempty" type:"Struct"`
	Ready           *bool                                                                          `json:"Ready,omitempty" xml:"Ready,omitempty"`
	Gpu             *int32                                                                         `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	SecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext   `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	Memory          *float32                                                                       `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Stdin           *bool                                                                          `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	Name            *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	EnvironmentVars []*DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	ReadinessProbe  *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe    `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetLivenessProbe(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.LivenessProbe = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCommands(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Commands = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetVolumeMounts(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetArgs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Args = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetImage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Image = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetPorts(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Ports = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetRestartCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.RestartCount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetImagePullPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetStdinOnce(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.StdinOnce = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetPreviousState(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.PreviousState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetTty(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Tty = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetWorkingDir(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.WorkingDir = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCurrentState(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.CurrentState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetReady(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Ready = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetGpu(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.SecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetStdin(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Stdin = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetEnvironmentVars(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetReadinessProbe(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.ReadinessProbe = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe struct {
	SuccessThreshold    *int32                                                                              `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	InitialDelaySeconds *int32                                                                              `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	FailureThreshold    *int32                                                                              `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	TimeoutSeconds      *int32                                                                              `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	TcpSocket           *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	Execs               []*string                                                                           `json:"Execs,omitempty" xml:"Execs,omitempty" type:"Repeated"`
	HttpGet             *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	PeriodSeconds       *int32                                                                              `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetSuccessThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetInitialDelaySeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetFailureThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetTimeoutSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetTcpSocket(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetExecs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.Execs = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetHttpGet(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetPeriodSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket struct {
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) SetHost(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet struct {
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetScheme(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts struct {
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetMountPropagation(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetMountPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.Name = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) SetProtocol(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Signal       *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	ExitCode     *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.State = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Reason = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Signal       *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	ExitCode     *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.State = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Reason = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext struct {
	ReadOnlyRootFilesystem *bool                                                                                  `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                                 `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
	Capability             *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetRunAsUser(v int64) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetCapability(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.Capability = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) SetAdds(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars struct {
	Key       *string                                                                               `json:"Key,omitempty" xml:"Key,omitempty"`
	Value     *string                                                                               `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueFrom *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom `json:"ValueFrom,omitempty" xml:"ValueFrom,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetValueFrom(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.ValueFrom = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom struct {
	FieldRef *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) SetFieldRef(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom {
	s.FieldRef = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) SetFieldPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef {
	s.FieldPath = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe struct {
	SuccessThreshold    *int32                                                                               `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	InitialDelaySeconds *int32                                                                               `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	FailureThreshold    *int32                                                                               `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	TimeoutSeconds      *int32                                                                               `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	TcpSocket           *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	Execs               []*string                                                                            `json:"Execs,omitempty" xml:"Execs,omitempty" type:"Repeated"`
	HttpGet             *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	PeriodSeconds       *int32                                                                               `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetSuccessThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetInitialDelaySeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetFailureThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetTimeoutSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetTcpSocket(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetExecs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.Execs = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetHttpGet(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetPeriodSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket struct {
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) SetHost(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet struct {
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetScheme(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainers struct {
	VolumeMounts    []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts    `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	Args            []*string                                                                          `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Image           *string                                                                            `json:"Image,omitempty" xml:"Image,omitempty"`
	Ports           []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts           `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	RestartCount    *int32                                                                             `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	ImagePullPolicy *string                                                                            `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	PreviousState   *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState     `json:"PreviousState,omitempty" xml:"PreviousState,omitempty" type:"Struct"`
	WorkingDir      *string                                                                            `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	Cpu             *float32                                                                           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CurrentState    *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState      `json:"CurrentState,omitempty" xml:"CurrentState,omitempty" type:"Struct"`
	Command         []*string                                                                          `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Ready           *bool                                                                              `json:"Ready,omitempty" xml:"Ready,omitempty"`
	Gpu             *int32                                                                             `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	SecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext   `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	Memory          *float32                                                                           `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name            *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	EnvironmentVars []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetVolumeMounts(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetArgs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Args = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetImage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Image = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetPorts(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Ports = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetRestartCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.RestartCount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetImagePullPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetPreviousState(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.PreviousState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetWorkingDir(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.WorkingDir = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCurrentState(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.CurrentState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCommand(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Command = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetReady(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Ready = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetGpu(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.SecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetEnvironmentVars(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.EnvironmentVars = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts struct {
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetMountPropagation(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetMountPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.Name = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) SetProtocol(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Signal       *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	ExitCode     *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.State = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Reason = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState struct {
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Signal       *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	ExitCode     *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.State = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Reason = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext struct {
	ReadOnlyRootFilesystem *bool                                                                                      `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                                     `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
	Capability             *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetRunAsUser(v int64) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetCapability(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.Capability = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) SetAdds(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars struct {
	Key       *string                                                                                   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value     *string                                                                                   `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueFrom *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom `json:"ValueFrom,omitempty" xml:"ValueFrom,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetValueFrom(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.ValueFrom = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom struct {
	FieldRef *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) SetFieldRef(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom {
	s.FieldRef = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) SetFieldPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef {
	s.FieldPath = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig struct {
	Searches    []*string                                                             `json:"Searches,omitempty" xml:"Searches,omitempty" type:"Repeated"`
	Options     []*DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	NameServers []*string                                                             `json:"NameServers,omitempty" xml:"NameServers,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetSearches(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.Searches = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetOptions(v []*DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.Options = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetNameServers(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.NameServers = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions {
	s.Value = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions {
	s.Name = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsVolumes struct {
	Type                              *string                                                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	DiskVolumeDiskId                  *string                                                                                       `json:"DiskVolumeDiskId,omitempty" xml:"DiskVolumeDiskId,omitempty"`
	NFSVolumeReadOnly                 *bool                                                                                         `json:"NFSVolumeReadOnly,omitempty" xml:"NFSVolumeReadOnly,omitempty"`
	ConfigFileVolumeConfigFileToPaths []*DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	FlexVolumeFsType                  *string                                                                                       `json:"FlexVolumeFsType,omitempty" xml:"FlexVolumeFsType,omitempty"`
	FlexVolumeDriver                  *string                                                                                       `json:"FlexVolumeDriver,omitempty" xml:"FlexVolumeDriver,omitempty"`
	DiskVolumeFsType                  *string                                                                                       `json:"DiskVolumeFsType,omitempty" xml:"DiskVolumeFsType,omitempty"`
	FlexVolumeOptions                 *string                                                                                       `json:"FlexVolumeOptions,omitempty" xml:"FlexVolumeOptions,omitempty"`
	NFSVolumeServer                   *string                                                                                       `json:"NFSVolumeServer,omitempty" xml:"NFSVolumeServer,omitempty"`
	NFSVolumePath                     *string                                                                                       `json:"NFSVolumePath,omitempty" xml:"NFSVolumePath,omitempty"`
	Name                              *string                                                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumes) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumes) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.Type = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetDiskVolumeDiskId(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.DiskVolumeDiskId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumeReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumeReadOnly = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetConfigFileVolumeConfigFileToPaths(v []*DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.ConfigFileVolumeConfigFileToPaths = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeFsType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeFsType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeDriver(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeDriver = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetDiskVolumeFsType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.DiskVolumeFsType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeOptions(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeOptions = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumeServer(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumeServer = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumePath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumePath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.Name = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths struct {
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths {
	s.Path = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) SetContent(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths {
	s.Content = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext struct {
	Sysctls []*DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls `json:"Sysctls,omitempty" xml:"Sysctls,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) SetSysctls(v []*DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext {
	s.Sysctls = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls {
	s.Value = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls {
	s.Name = &v
	return s
}

type DescribeContainerGroupsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContainerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupsResponse) SetBody(v *DescribeContainerGroupsResponseBody) *DescribeContainerGroupsResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("eci"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateImageCacheWithOptions(request *CreateImageCacheRequest, runtime *util.RuntimeOptions) (_result *CreateImageCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateImageCache"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageCache(request *CreateImageCacheRequest) (_result *CreateImageCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CreateImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerLogWithOptions(request *DescribeContainerLogRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContainerLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContainerLog"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerLog(request *DescribeContainerLogRequest) (_result *DescribeContainerLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerLogResponse{}
	_body, _err := client.DescribeContainerLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartContainerGroupWithOptions(request *RestartContainerGroupRequest, runtime *util.RuntimeOptions) (_result *RestartContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartContainerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartContainerGroup"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartContainerGroup(request *RestartContainerGroupRequest) (_result *RestartContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartContainerGroupResponse{}
	_body, _err := client.RestartContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageCachesWithOptions(request *DescribeImageCachesRequest, runtime *util.RuntimeOptions) (_result *DescribeImageCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImageCachesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImageCaches"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageCaches(request *DescribeImageCachesRequest) (_result *DescribeImageCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageCachesResponse{}
	_body, _err := client.DescribeImageCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContainerGroupWithOptions(request *DeleteContainerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteContainerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteContainerGroup"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) (_result *DeleteContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContainerGroupResponse{}
	_body, _err := client.DeleteContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsageWithOptions(request *ListUsageRequest, runtime *util.RuntimeOptions) (_result *ListUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUsageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUsage"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsage(request *ListUsageRequest) (_result *ListUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsageResponse{}
	_body, _err := client.ListUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateContainerGroupWithOptions(request *CreateContainerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateContainerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateContainerGroup"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateContainerGroup(request *CreateContainerGroupRequest) (_result *CreateContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateContainerGroupResponse{}
	_body, _err := client.CreateContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageCacheWithOptions(request *DeleteImageCacheRequest, runtime *util.RuntimeOptions) (_result *DeleteImageCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImageCache"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImageCache(request *DeleteImageCacheRequest) (_result *DeleteImageCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.DeleteImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupMetricWithOptions(request *DescribeContainerGroupMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContainerGroupMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContainerGroupMetric"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroupMetric(request *DescribeContainerGroupMetricRequest) (_result *DescribeContainerGroupMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupMetricResponse{}
	_body, _err := client.DescribeContainerGroupMetricWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateContainerGroupWithOptions(request *UpdateContainerGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateContainerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateContainerGroup"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContainerGroup(request *UpdateContainerGroupRequest) (_result *UpdateContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateContainerGroupResponse{}
	_body, _err := client.UpdateContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecContainerCommandWithOptions(request *ExecContainerCommandRequest, runtime *util.RuntimeOptions) (_result *ExecContainerCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecContainerCommandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecContainerCommand"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecContainerCommand(request *ExecContainerCommandRequest) (_result *ExecContainerCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecContainerCommandResponse{}
	_body, _err := client.ExecContainerCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupPriceWithOptions(request *DescribeContainerGroupPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContainerGroupPriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContainerGroupPrice"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroupPrice(request *DescribeContainerGroupPriceRequest) (_result *DescribeContainerGroupPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupPriceResponse{}
	_body, _err := client.DescribeContainerGroupPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMultiContainerGroupMetricWithOptions(request *DescribeMultiContainerGroupMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiContainerGroupMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMultiContainerGroupMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMultiContainerGroupMetric"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMultiContainerGroupMetric(request *DescribeMultiContainerGroupMetricRequest) (_result *DescribeMultiContainerGroupMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiContainerGroupMetricResponse{}
	_body, _err := client.DescribeMultiContainerGroupMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupsWithOptions(request *DescribeContainerGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContainerGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContainerGroups"), tea.String("2018-08-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (_result *DescribeContainerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupsResponse{}
	_body, _err := client.DescribeContainerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
