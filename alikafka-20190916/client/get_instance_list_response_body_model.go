// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetInstanceListResponseBody
	GetCode() *int32
	SetInstanceList(v *GetInstanceListResponseBodyInstanceList) *GetInstanceListResponseBody
	GetInstanceList() *GetInstanceListResponseBodyInstanceList
	SetMessage(v string) *GetInstanceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceListResponseBody
	GetSuccess() *bool
}

type GetInstanceListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code         *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceList *GetInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// 4B6D821D-7F67-4CAA-9E13-A5A997C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetInstanceListResponseBody) GetInstanceList() *GetInstanceListResponseBodyInstanceList {
	return s.InstanceList
}

func (s *GetInstanceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceListResponseBody) SetCode(v int32) *GetInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceListResponseBody) SetInstanceList(v *GetInstanceListResponseBodyInstanceList) *GetInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *GetInstanceListResponseBody) SetMessage(v string) *GetInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceListResponseBody) SetRequestId(v string) *GetInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceListResponseBody) SetSuccess(v bool) *GetInstanceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceListResponseBody) Validate() error {
	if s.InstanceList != nil {
		if err := s.InstanceList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceListResponseBodyInstanceList struct {
	InstanceVO []*GetInstanceListResponseBodyInstanceListInstanceVO `json:"InstanceVO,omitempty" xml:"InstanceVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceList) GetInstanceVO() []*GetInstanceListResponseBodyInstanceListInstanceVO {
	return s.InstanceVO
}

func (s *GetInstanceListResponseBodyInstanceList) SetInstanceVO(v []*GetInstanceListResponseBodyInstanceListInstanceVO) *GetInstanceListResponseBodyInstanceList {
	s.InstanceVO = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceList) Validate() error {
	if s.InstanceVO != nil {
		for _, item := range s.InstanceVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceListResponseBodyInstanceListInstanceVO struct {
	AllConfig                   *string                                                                       `json:"AllConfig,omitempty" xml:"AllConfig,omitempty"`
	AutoCreateGroupEnable       *bool                                                                         `json:"AutoCreateGroupEnable,omitempty" xml:"AutoCreateGroupEnable,omitempty"`
	AutoCreateTopicEnable       *bool                                                                         `json:"AutoCreateTopicEnable,omitempty" xml:"AutoCreateTopicEnable,omitempty"`
	BackupZoneId                *string                                                                       `json:"BackupZoneId,omitempty" xml:"BackupZoneId,omitempty"`
	ConfluentConfig             *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig             `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	ConfluentInstanceComponents *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents `json:"ConfluentInstanceComponents,omitempty" xml:"ConfluentInstanceComponents,omitempty" type:"Struct"`
	CreateTime                  *int64                                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultPartitionNum         *int32                                                                        `json:"DefaultPartitionNum,omitempty" xml:"DefaultPartitionNum,omitempty"`
	DeployType                  *int32                                                                        `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DiskSize                    *int32                                                                        `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType                    *int32                                                                        `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DomainEndpoint              *string                                                                       `json:"DomainEndpoint,omitempty" xml:"DomainEndpoint,omitempty"`
	EipMax                      *int32                                                                        `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	EndPoint                    *string                                                                       `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	ExpiredTime                 *int64                                                                        `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceId                  *string                                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoMax                       *int32                                                                        `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxRead                   *int32                                                                        `json:"IoMaxRead,omitempty" xml:"IoMaxRead,omitempty"`
	IoMaxSpec                   *string                                                                       `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	IoMaxWrite                  *int32                                                                        `json:"IoMaxWrite,omitempty" xml:"IoMaxWrite,omitempty"`
	KmsKeyId                    *string                                                                       `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	MsgRetain                   *int32                                                                        `json:"MsgRetain,omitempty" xml:"MsgRetain,omitempty"`
	Name                        *string                                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	PaidType                    *int32                                                                        `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	RecommendedPartitionCount   *int32                                                                        `json:"RecommendedPartitionCount,omitempty" xml:"RecommendedPartitionCount,omitempty"`
	RegionId                    *string                                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReservedPublishCapacity     *int32                                                                        `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	ReservedSubscribeCapacity   *int32                                                                        `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
	ResourceGroupId             *string                                                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SaslDomainEndpoint          *string                                                                       `json:"SaslDomainEndpoint,omitempty" xml:"SaslDomainEndpoint,omitempty"`
	SaslEndPoint                *string                                                                       `json:"SaslEndPoint,omitempty" xml:"SaslEndPoint,omitempty"`
	ScheduledRetirement         *bool                                                                         `json:"ScheduledRetirement,omitempty" xml:"ScheduledRetirement,omitempty"`
	SecurityGroup               *string                                                                       `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	Series                      *string                                                                       `json:"Series,omitempty" xml:"Series,omitempty"`
	ServiceStatus               *int32                                                                        `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	SpecType                    *string                                                                       `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	SslDomainEndpoint           *string                                                                       `json:"SslDomainEndpoint,omitempty" xml:"SslDomainEndpoint,omitempty"`
	SslEndPoint                 *string                                                                       `json:"SslEndPoint,omitempty" xml:"SslEndPoint,omitempty"`
	StandardZoneId              *string                                                                       `json:"StandardZoneId,omitempty" xml:"StandardZoneId,omitempty"`
	Tags                        *GetInstanceListResponseBodyInstanceListInstanceVOTags                        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TopicNumLimit               *int32                                                                        `json:"TopicNumLimit,omitempty" xml:"TopicNumLimit,omitempty"`
	UpgradeServiceDetailInfo    *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo    `json:"UpgradeServiceDetailInfo,omitempty" xml:"UpgradeServiceDetailInfo,omitempty" type:"Struct"`
	UsedGroupCount              *int32                                                                        `json:"UsedGroupCount,omitempty" xml:"UsedGroupCount,omitempty"`
	UsedPartitionCount          *int32                                                                        `json:"UsedPartitionCount,omitempty" xml:"UsedPartitionCount,omitempty"`
	UsedTopicCount              *int32                                                                        `json:"UsedTopicCount,omitempty" xml:"UsedTopicCount,omitempty"`
	VSwitchId                   *string                                                                       `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchIds                  *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds                  `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	ViewInstanceStatusCode      *int32                                                                        `json:"ViewInstanceStatusCode,omitempty" xml:"ViewInstanceStatusCode,omitempty"`
	VpcId                       *string                                                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcSaslDomainEndpoint       *string                                                                       `json:"VpcSaslDomainEndpoint,omitempty" xml:"VpcSaslDomainEndpoint,omitempty"`
	VpcSaslEndPoint             *string                                                                       `json:"VpcSaslEndPoint,omitempty" xml:"VpcSaslEndPoint,omitempty"`
	ZoneId                      *string                                                                       `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVO) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetAllConfig() *string {
	return s.AllConfig
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetAutoCreateGroupEnable() *bool {
	return s.AutoCreateGroupEnable
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetAutoCreateTopicEnable() *bool {
	return s.AutoCreateTopicEnable
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetBackupZoneId() *string {
	return s.BackupZoneId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetConfluentConfig() *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	return s.ConfluentConfig
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetConfluentInstanceComponents() *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents {
	return s.ConfluentInstanceComponents
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDefaultPartitionNum() *int32 {
	return s.DefaultPartitionNum
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDeployType() *int32 {
	return s.DeployType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDiskType() *int32 {
	return s.DiskType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDomainEndpoint() *string {
	return s.DomainEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetEipMax() *int32 {
	return s.EipMax
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetEndPoint() *string {
	return s.EndPoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetIoMax() *int32 {
	return s.IoMax
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetIoMaxRead() *int32 {
	return s.IoMaxRead
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetIoMaxWrite() *int32 {
	return s.IoMaxWrite
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetMsgRetain() *int32 {
	return s.MsgRetain
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetName() *string {
	return s.Name
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetPaidType() *int32 {
	return s.PaidType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetRecommendedPartitionCount() *int32 {
	return s.RecommendedPartitionCount
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetReservedPublishCapacity() *int32 {
	return s.ReservedPublishCapacity
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetReservedSubscribeCapacity() *int32 {
	return s.ReservedSubscribeCapacity
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSaslDomainEndpoint() *string {
	return s.SaslDomainEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSaslEndPoint() *string {
	return s.SaslEndPoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetScheduledRetirement() *bool {
	return s.ScheduledRetirement
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSeries() *string {
	return s.Series
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSpecType() *string {
	return s.SpecType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSslDomainEndpoint() *string {
	return s.SslDomainEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSslEndPoint() *string {
	return s.SslEndPoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetStandardZoneId() *string {
	return s.StandardZoneId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetTags() *GetInstanceListResponseBodyInstanceListInstanceVOTags {
	return s.Tags
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetTopicNumLimit() *int32 {
	return s.TopicNumLimit
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetUpgradeServiceDetailInfo() *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo {
	return s.UpgradeServiceDetailInfo
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetUsedGroupCount() *int32 {
	return s.UsedGroupCount
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetUsedPartitionCount() *int32 {
	return s.UsedPartitionCount
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetUsedTopicCount() *int32 {
	return s.UsedTopicCount
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVSwitchIds() *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds {
	return s.VSwitchIds
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetViewInstanceStatusCode() *int32 {
	return s.ViewInstanceStatusCode
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVpcSaslDomainEndpoint() *string {
	return s.VpcSaslDomainEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVpcSaslEndPoint() *string {
	return s.VpcSaslEndPoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAllConfig(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AllConfig = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAutoCreateGroupEnable(v bool) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AutoCreateGroupEnable = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAutoCreateTopicEnable(v bool) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AutoCreateTopicEnable = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetBackupZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.BackupZoneId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetConfluentConfig(v *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ConfluentConfig = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetConfluentInstanceComponents(v *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ConfluentInstanceComponents = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetCreateTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDefaultPartitionNum(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DefaultPartitionNum = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDeployType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DeployType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDiskSize(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DiskSize = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDiskType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DiskType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetEipMax(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.EipMax = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.EndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetExpiredTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetInstanceId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMax(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMax = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxRead(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxRead = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxSpec(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxSpec = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxWrite(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxWrite = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetKmsKeyId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.KmsKeyId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetMsgRetain(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.MsgRetain = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetName(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Name = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetPaidType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.PaidType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetRecommendedPartitionCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.RecommendedPartitionCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetRegionId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.RegionId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetReservedPublishCapacity(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetReservedSubscribeCapacity(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ReservedSubscribeCapacity = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetResourceGroupId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSaslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SaslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSaslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SaslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetScheduledRetirement(v bool) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ScheduledRetirement = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSecurityGroup(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SecurityGroup = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSeries(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Series = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetServiceStatus(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ServiceStatus = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSpecType(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SpecType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetStandardZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.StandardZoneId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetTags(v *GetInstanceListResponseBodyInstanceListInstanceVOTags) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Tags = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetTopicNumLimit(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.TopicNumLimit = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUpgradeServiceDetailInfo(v *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UpgradeServiceDetailInfo = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedGroupCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedGroupCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedPartitionCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedPartitionCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedTopicCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedTopicCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVSwitchId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVSwitchIds(v *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VSwitchIds = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetViewInstanceStatusCode(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ViewInstanceStatusCode = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcSaslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcSaslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcSaslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcSaslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ZoneId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) Validate() error {
	if s.ConfluentConfig != nil {
		if err := s.ConfluentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ConfluentInstanceComponents != nil {
		if err := s.ConfluentInstanceComponents.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.UpgradeServiceDetailInfo != nil {
		if err := s.UpgradeServiceDetailInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchIds != nil {
		if err := s.VSwitchIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig struct {
	ConnectCU             *int32                                                                    `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	ConnectReplica        *int32                                                                    `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	ControlCenterCU       *int32                                                                    `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	ControlCenterReplica  *int32                                                                    `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	ControlCenterStorage  *int32                                                                    `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	KafkaCU               *int32                                                                    `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	KafkaReplica          *int32                                                                    `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	KafkaRestProxyCU      *int32                                                                    `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	KafkaRestProxyReplica *int32                                                                    `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	KafkaStorage          *int32                                                                    `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	KsqlCU                *int32                                                                    `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	KsqlList              *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList `json:"KsqlList,omitempty" xml:"KsqlList,omitempty" type:"Struct"`
	KsqlReplica           *int32                                                                    `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	KsqlStorage           *int32                                                                    `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	SchemaRegistryCU      *int32                                                                    `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	SchemaRegistryReplica *int32                                                                    `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	ZooKeeperCU           *int32                                                                    `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	ZooKeeperReplica      *int32                                                                    `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	ZooKeeperStorage      *int32                                                                    `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetConnectCU() *int32 {
	return s.ConnectCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetConnectReplica() *int32 {
	return s.ConnectReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetControlCenterCU() *int32 {
	return s.ControlCenterCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetControlCenterReplica() *int32 {
	return s.ControlCenterReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetControlCenterStorage() *int32 {
	return s.ControlCenterStorage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaCU() *int32 {
	return s.KafkaCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaReplica() *int32 {
	return s.KafkaReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaRestProxyCU() *int32 {
	return s.KafkaRestProxyCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaRestProxyReplica() *int32 {
	return s.KafkaRestProxyReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaStorage() *int32 {
	return s.KafkaStorage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKsqlCU() *int32 {
	return s.KsqlCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKsqlList() *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList {
	return s.KsqlList
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKsqlReplica() *int32 {
	return s.KsqlReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKsqlStorage() *int32 {
	return s.KsqlStorage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetSchemaRegistryCU() *int32 {
	return s.SchemaRegistryCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetSchemaRegistryReplica() *int32 {
	return s.SchemaRegistryReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetZooKeeperCU() *int32 {
	return s.ZooKeeperCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetZooKeeperReplica() *int32 {
	return s.ZooKeeperReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetZooKeeperStorage() *int32 {
	return s.ZooKeeperStorage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetConnectCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetConnectReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaRestProxyCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaRestProxyReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlList(v *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlList = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetSchemaRegistryCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetSchemaRegistryReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) Validate() error {
	if s.KsqlList != nil {
		if err := s.KsqlList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList struct {
	ConfluentInstanceComponentResourceVO []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO `json:"ConfluentInstanceComponentResourceVO,omitempty" xml:"ConfluentInstanceComponentResourceVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList) GetConfluentInstanceComponentResourceVO() []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO {
	return s.ConfluentInstanceComponentResourceVO
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList) SetConfluentInstanceComponentResourceVO(v []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList {
	s.ConfluentInstanceComponentResourceVO = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlList) Validate() error {
	if s.ConfluentInstanceComponentResourceVO != nil {
		for _, item := range s.ConfluentInstanceComponentResourceVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO struct {
	Cu         *int32  `json:"Cu,omitempty" xml:"Cu,omitempty"`
	InternalId *string `json:"InternalId,omitempty" xml:"InternalId,omitempty"`
	Replica    *int32  `json:"Replica,omitempty" xml:"Replica,omitempty"`
	Storage    *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) GetCu() *int32 {
	return s.Cu
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) GetInternalId() *string {
	return s.InternalId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) GetReplica() *int32 {
	return s.Replica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) GetStorage() *int32 {
	return s.Storage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) GetType() *string {
	return s.Type
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) SetCu(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO {
	s.Cu = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) SetInternalId(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO {
	s.InternalId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) SetReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO {
	s.Replica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) SetStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO {
	s.Storage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) SetType(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO {
	s.Type = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfigKsqlListConfluentInstanceComponentResourceVO) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents struct {
	ConfluentInstanceComponentVO []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO `json:"ConfluentInstanceComponentVO,omitempty" xml:"ConfluentInstanceComponentVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) GetConfluentInstanceComponentVO() []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	return s.ConfluentInstanceComponentVO
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) SetConfluentInstanceComponentVO(v []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents {
	s.ConfluentInstanceComponentVO = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) Validate() error {
	if s.ConfluentInstanceComponentVO != nil {
		for _, item := range s.ConfluentInstanceComponentVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO struct {
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	DeployModule  *string `json:"DeployModule,omitempty" xml:"DeployModule,omitempty"`
	PubEndpoint   *string `json:"PubEndpoint,omitempty" xml:"PubEndpoint,omitempty"`
	VpcEndpoint   *string `json:"VpcEndpoint,omitempty" xml:"VpcEndpoint,omitempty"`
	InternalId    *string `json:"internalId,omitempty" xml:"internalId,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetComponentType() *string {
	return s.ComponentType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetDeployModule() *string {
	return s.DeployModule
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetPubEndpoint() *string {
	return s.PubEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetVpcEndpoint() *string {
	return s.VpcEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetInternalId() *string {
	return s.InternalId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetComponentType(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.ComponentType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetDeployModule(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.DeployModule = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetPubEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.PubEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetVpcEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.VpcEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetInternalId(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.InternalId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOTags struct {
	TagVO []*GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO `json:"TagVO,omitempty" xml:"TagVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTags) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTags) GetTagVO() []*GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO {
	return s.TagVO
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTags) SetTagVO(v []*GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) *GetInstanceListResponseBodyInstanceListInstanceVOTags {
	s.TagVO = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTags) Validate() error {
	if s.TagVO != nil {
		for _, item := range s.TagVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) GetKey() *string {
	return s.Key
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) GetValue() *string {
	return s.Value
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) SetKey(v string) *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO {
	s.Key = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) SetValue(v string) *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO {
	s.Value = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo struct {
	Current2OpenSourceVersion *string `json:"Current2OpenSourceVersion,omitempty" xml:"Current2OpenSourceVersion,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) GetCurrent2OpenSourceVersion() *string {
	return s.Current2OpenSourceVersion
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) SetCurrent2OpenSourceVersion(v string) *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo {
	s.Current2OpenSourceVersion = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds struct {
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) SetVSwitchIds(v []*string) *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds {
	s.VSwitchIds = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) Validate() error {
	return dara.Validate(s)
}
