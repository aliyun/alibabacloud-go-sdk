// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAndroidInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceModel(v []*DescribeAndroidInstancesResponseBodyInstanceModel) *DescribeAndroidInstancesResponseBody
	GetInstanceModel() []*DescribeAndroidInstancesResponseBodyInstanceModel
	SetNextToken(v string) *DescribeAndroidInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAndroidInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAndroidInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeAndroidInstancesResponseBody struct {
	// The cloud phone instances.
	InstanceModel []*DescribeAndroidInstancesResponseBodyInstanceModel `json:"InstanceModel,omitempty" xml:"InstanceModel,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kmma/xxE9WtwL/ADvZ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAndroidInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBody) GetInstanceModel() []*DescribeAndroidInstancesResponseBodyInstanceModel {
	return s.InstanceModel
}

func (s *DescribeAndroidInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAndroidInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAndroidInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAndroidInstancesResponseBody) SetInstanceModel(v []*DescribeAndroidInstancesResponseBodyInstanceModel) *DescribeAndroidInstancesResponseBody {
	s.InstanceModel = v
	return s
}

func (s *DescribeAndroidInstancesResponseBody) SetNextToken(v string) *DescribeAndroidInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBody) SetRequestId(v string) *DescribeAndroidInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBody) SetTotalCount(v int32) *DescribeAndroidInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBody) Validate() error {
	if s.InstanceModel != nil {
		for _, item := range s.InstanceModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAndroidInstancesResponseBodyInstanceModel struct {
	// The ID of the instance group.
	//
	// example:
	//
	// ag-ayyhomlal7po****
	AndroidInstanceGroupId *string `json:"AndroidInstanceGroupId,omitempty" xml:"AndroidInstanceGroupId,omitempty"`
	// The name of the instance group.
	//
	// example:
	//
	// AndroidInstanceGroupName
	AndroidInstanceGroupName *string `json:"AndroidInstanceGroupName,omitempty" xml:"AndroidInstanceGroupName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// acp-8at8h6ejkadjh****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// name
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	// The state of the instance.
	//
	// example:
	//
	// RUNNING
	AndroidInstanceStatus *string `json:"AndroidInstanceStatus,omitempty" xml:"AndroidInstanceStatus,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-i7yv6tkn7kh8dv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the physical instance.
	//
	// example:
	//
	// ai-9ey6io0q58rcd****
	AppInstanceId   *string                                                           `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppManagePolicy *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy `json:"AppManagePolicy,omitempty" xml:"AppManagePolicy,omitempty" type:"Struct"`
	// The ID of the user to whom the instance is assigned.
	//
	// example:
	//
	// test
	AuthorizedUserId     *string `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	BandwidthPackageId   *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	// The ID of the bound user.
	//
	// example:
	//
	// test
	BindUserId *string                                                     `json:"BindUserId,omitempty" xml:"BindUserId,omitempty"`
	BizTags    []*DescribeAndroidInstancesResponseBodyInstanceModelBizTags `json:"BizTags,omitempty" xml:"BizTags,omitempty" type:"Repeated"`
	// The billing method of the instance.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 4
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The disks.
	Disks              []*DescribeAndroidInstancesResponseBodyInstanceModelDisks       `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	DisplayConfig      *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty" type:"Struct"`
	DownBandwidthLimit *int32                                                          `json:"DownBandwidthLimit,omitempty" xml:"DownBandwidthLimit,omitempty"`
	// The cause of the instance data backup failure or restoration failure.
	//
	// example:
	//
	// FilePathNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2023-05-06 10:42:10
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the subscription instance group expires.
	//
	// example:
	//
	// 2024-07-15T02:03:33Z
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The time when the instance was modified.
	//
	// example:
	//
	// 2023-05-06 10:42:10
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The version of the image.
	//
	// example:
	//
	// 3.5.3.867
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// acp.basic.small
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetStatus *string `json:"InternetStatus,omitempty" xml:"InternetStatus,omitempty"`
	// The ID of the key pair.
	//
	// example:
	//
	// kp-5hh431emkpucs****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The IP address of the ENI.
	//
	// example:
	//
	// 192.168.22.48
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	NetworkInterfaceIpv6Address *string `json:"NetworkInterfaceIpv6Address,omitempty" xml:"NetworkInterfaceIpv6Address,omitempty"`
	NetworkType                 *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-shenzhen+dir-211620****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The ID of the persistent session.
	//
	// example:
	//
	// p-0btrd5zj8epo****
	PersistentAppInstanceId *string                                                         `json:"PersistentAppInstanceId,omitempty" xml:"PersistentAppInstanceId,omitempty"`
	PhoneDataInfo           *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo `json:"PhoneDataInfo,omitempty" xml:"PhoneDataInfo,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-0bszojpu0seql****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 10.32.1.41
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	PublicIpv6Address *string `json:"PublicIpv6Address,omitempty" xml:"PublicIpv6Address,omitempty"`
	QosRuleId         *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// The progress of instance data backup or restoration.
	//
	// example:
	//
	// 100
	Rate *int32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rendering type.
	//
	// example:
	//
	// local
	RenderingType *string `json:"RenderingType,omitempty" xml:"RenderingType,omitempty"`
	ServerStatus  *string `json:"ServerStatus,omitempty" xml:"ServerStatus,omitempty"`
	ServerType    *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The session status.
	//
	// Valid values:
	//
	// 	- disConnect: The session is disconnected.
	//
	// 	- connect: The session is connected.
	//
	// example:
	//
	// connect
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	StreamMode    *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	// The tags.
	Tags             []*DescribeAndroidInstancesResponseBodyInstanceModelTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UpBandwidthLimit *int32                                                   `json:"UpBandwidthLimit,omitempty" xml:"UpBandwidthLimit,omitempty"`
	VSwitchId        *string                                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId           *string                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModel) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAndroidInstanceGroupId() *string {
	return s.AndroidInstanceGroupId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAndroidInstanceGroupName() *string {
	return s.AndroidInstanceGroupName
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAndroidInstanceName() *string {
	return s.AndroidInstanceName
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAndroidInstanceStatus() *string {
	return s.AndroidInstanceStatus
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAppManagePolicy() *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy {
	return s.AppManagePolicy
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetAuthorizedUserId() *string {
	return s.AuthorizedUserId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetBindUserId() *string {
	return s.BindUserId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetBizTags() []*DescribeAndroidInstancesResponseBodyInstanceModelBizTags {
	return s.BizTags
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetDisks() []*DescribeAndroidInstancesResponseBodyInstanceModelDisks {
	return s.Disks
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetDisplayConfig() *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	return s.DisplayConfig
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetDownBandwidthLimit() *int32 {
	return s.DownBandwidthLimit
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetGmtExpired() *string {
	return s.GmtExpired
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetInternetStatus() *string {
	return s.InternetStatus
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetNetworkInterfaceIp() *string {
	return s.NetworkInterfaceIp
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetNetworkInterfaceIpv6Address() *string {
	return s.NetworkInterfaceIpv6Address
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetPersistentAppInstanceId() *string {
	return s.PersistentAppInstanceId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetPhoneDataInfo() *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo {
	return s.PhoneDataInfo
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetPublicIpv6Address() *string {
	return s.PublicIpv6Address
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetRate() *int32 {
	return s.Rate
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetRenderingType() *string {
	return s.RenderingType
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetServerStatus() *string {
	return s.ServerStatus
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetServerType() *string {
	return s.ServerType
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetStreamMode() *int32 {
	return s.StreamMode
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetSystemVersion() *string {
	return s.SystemVersion
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetTags() []*DescribeAndroidInstancesResponseBodyInstanceModelTags {
	return s.Tags
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetUpBandwidthLimit() *int32 {
	return s.UpBandwidthLimit
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceGroupId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceGroupName(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceGroupName = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceName(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceName = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAndroidInstanceStatus(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AndroidInstanceStatus = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAppInstanceGroupId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAppInstanceId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AppInstanceId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAppManagePolicy(v *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AppManagePolicy = v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetAuthorizedUserId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetBandwidthPackageId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetBandwidthPackageType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.BandwidthPackageType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetBindUserId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.BindUserId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetBizTags(v []*DescribeAndroidInstancesResponseBodyInstanceModelBizTags) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.BizTags = v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetChargeType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ChargeType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetCpu(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Cpu = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetDisks(v []*DescribeAndroidInstancesResponseBodyInstanceModelDisks) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Disks = v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetDisplayConfig(v *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.DisplayConfig = v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetDownBandwidthLimit(v int32) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.DownBandwidthLimit = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetErrorCode(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetGmtCreate(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetGmtExpired(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.GmtExpired = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetGmtModified(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetImageId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ImageId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetImageVersion(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ImageVersion = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetInstanceType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.InstanceType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetInternetStatus(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.InternetStatus = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetKeyPairId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.KeyPairId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetMemory(v int32) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Memory = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetNetworkInterfaceIp(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetNetworkInterfaceIpv6Address(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.NetworkInterfaceIpv6Address = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetNetworkType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.NetworkType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetOfficeSiteId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPersistentAppInstanceId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PersistentAppInstanceId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPhoneDataInfo(v *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PhoneDataInfo = v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPolicyGroupId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPublicIpAddress(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetPublicIpv6Address(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.PublicIpv6Address = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetQosRuleId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.QosRuleId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetRate(v int32) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Rate = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetRegionId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.RegionId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetRenderingType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.RenderingType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetServerStatus(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ServerStatus = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetServerType(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ServerType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetSessionStatus(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.SessionStatus = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetStreamMode(v int32) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.StreamMode = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetSystemVersion(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.SystemVersion = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetTags(v []*DescribeAndroidInstancesResponseBodyInstanceModelTags) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.Tags = v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetUpBandwidthLimit(v int32) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.UpBandwidthLimit = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetVSwitchId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) SetZoneId(v string) *DescribeAndroidInstancesResponseBodyInstanceModel {
	s.ZoneId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModel) Validate() error {
	if s.AppManagePolicy != nil {
		if err := s.AppManagePolicy.Validate(); err != nil {
			return err
		}
	}
	if s.BizTags != nil {
		for _, item := range s.BizTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Disks != nil {
		for _, item := range s.Disks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DisplayConfig != nil {
		if err := s.DisplayConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PhoneDataInfo != nil {
		if err := s.PhoneDataInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy struct {
	AppManagePolicyId   *string `json:"AppManagePolicyId,omitempty" xml:"AppManagePolicyId,omitempty"`
	AppManagePolicyName *string `json:"AppManagePolicyName,omitempty" xml:"AppManagePolicyName,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) GetAppManagePolicyId() *string {
	return s.AppManagePolicyId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) GetAppManagePolicyName() *string {
	return s.AppManagePolicyName
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) SetAppManagePolicyId(v string) *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy {
	s.AppManagePolicyId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) SetAppManagePolicyName(v string) *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy {
	s.AppManagePolicyName = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelAppManagePolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstancesResponseBodyInstanceModelBizTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelBizTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelBizTags) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelBizTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelBizTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelBizTags) SetKey(v string) *DescribeAndroidInstancesResponseBodyInstanceModelBizTags {
	s.Key = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelBizTags) SetValue(v string) *DescribeAndroidInstancesResponseBodyInstanceModelBizTags {
	s.Value = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelBizTags) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstancesResponseBodyInstanceModelDisks struct {
	// The disk size. Unit: GB.
	//
	// example:
	//
	// 32
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The type of the disk.
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelDisks) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisks) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisks) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisks) SetDiskSize(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisks) SetDiskType(v string) *DescribeAndroidInstancesResponseBodyInstanceModelDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig struct {
	Dpi              *int32  `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	Fps              *int32  `json:"Fps,omitempty" xml:"Fps,omitempty"`
	LockResolution   *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	ResolutionHeight *int32  `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	ResolutionWidth  *int32  `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) GetDpi() *int32 {
	return s.Dpi
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) GetFps() *int32 {
	return s.Fps
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) GetLockResolution() *string {
	return s.LockResolution
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetDpi(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.Dpi = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetFps(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.Fps = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetLockResolution(v string) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.LockResolution = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetResolutionHeight(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) SetResolutionWidth(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelDisplayConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo struct {
	PhoneDataId     *string `json:"PhoneDataId,omitempty" xml:"PhoneDataId,omitempty"`
	PhoneDataVolume *int32  `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) GetPhoneDataId() *string {
	return s.PhoneDataId
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) GetPhoneDataVolume() *int32 {
	return s.PhoneDataVolume
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) SetPhoneDataId(v string) *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo {
	s.PhoneDataId = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) SetPhoneDataVolume(v int32) *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo {
	s.PhoneDataVolume = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelPhoneDataInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstancesResponseBodyInstanceModelTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponseBodyInstanceModelTags) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelTags) SetKey(v string) *DescribeAndroidInstancesResponseBodyInstanceModelTags {
	s.Key = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelTags) SetValue(v string) *DescribeAndroidInstancesResponseBodyInstanceModelTags {
	s.Value = &v
	return s
}

func (s *DescribeAndroidInstancesResponseBodyInstanceModelTags) Validate() error {
	return dara.Validate(s)
}
