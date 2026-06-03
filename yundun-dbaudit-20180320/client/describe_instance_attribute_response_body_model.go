// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody
	GetInstanceAttribute() *DescribeInstanceAttributeResponseBodyInstanceAttribute
	SetRequestId(v string) *DescribeInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeInstanceAttributeResponseBody struct {
	InstanceAttribute *DescribeInstanceAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Struct"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) GetInstanceAttribute() *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	return s.InstanceAttribute
}

func (s *DescribeInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) Validate() error {
	if s.InstanceAttribute != nil {
		if err := s.InstanceAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAttributeResponseBodyInstanceAttribute struct {
	AccessType          *int32    `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Description         *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsStatus           *string   `json:"EcsStatus,omitempty" xml:"EcsStatus,omitempty"`
	ExpireTime          *int64    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ImageVersionName    *string   `json:"ImageVersionName,omitempty" xml:"ImageVersionName,omitempty"`
	InstanceId          *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus      *string   `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InternetEndpoint    *string   `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	InternetIp          *string   `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IntranetEndpoint    *string   `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	IntranetIp          *string   `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	LicenseCode         *string   `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	NetworkType         *string   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Operatable          *bool     `json:"Operatable,omitempty" xml:"Operatable,omitempty"`
	PlanUpgradeStatus   *int32    `json:"PlanUpgradeStatus,omitempty" xml:"PlanUpgradeStatus,omitempty"`
	PlanUpgradeable     *bool     `json:"PlanUpgradeable,omitempty" xml:"PlanUpgradeable,omitempty"`
	PrivateWhiteList    []*string `json:"PrivateWhiteList,omitempty" xml:"PrivateWhiteList,omitempty" type:"Repeated"`
	PublicAccessControl *int32    `json:"PublicAccessControl,omitempty" xml:"PublicAccessControl,omitempty"`
	PublicWhiteList     []*string `json:"PublicWhiteList,omitempty" xml:"PublicWhiteList,omitempty" type:"Repeated"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Renewable           *bool     `json:"Renewable,omitempty" xml:"Renewable,omitempty"`
	SecurityGroupIds    []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SeriesCode          *string   `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	StartTime           *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UpgradeStatus       *int32    `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	Upgradeable         *bool     `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	VpcId               *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string   `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetAccessType() *int32 {
	return s.AccessType
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetEcsStatus() *string {
	return s.EcsStatus
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetImageVersionName() *string {
	return s.ImageVersionName
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetOperatable() *bool {
	return s.Operatable
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPlanUpgradeStatus() *int32 {
	return s.PlanUpgradeStatus
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPlanUpgradeable() *bool {
	return s.PlanUpgradeable
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPrivateWhiteList() []*string {
	return s.PrivateWhiteList
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPublicAccessControl() *int32 {
	return s.PublicAccessControl
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetPublicWhiteList() []*string {
	return s.PublicWhiteList
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetRenewable() *bool {
	return s.Renewable
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetSeriesCode() *string {
	return s.SeriesCode
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetUpgradeStatus() *int32 {
	return s.UpgradeStatus
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetUpgradeable() *bool {
	return s.Upgradeable
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetAccessType(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.AccessType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetEcsStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.EcsStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetExpireTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetImageVersionName(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ImageVersionName = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetIp(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetIp = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetIp(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetIp = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetLicenseCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetNetworkType(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetOperatable(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Operatable = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPlanUpgradeStatus(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PlanUpgradeStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPlanUpgradeable(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PlanUpgradeable = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPrivateWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PrivateWhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicAccessControl(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicAccessControl = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicWhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRenewable(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Renewable = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSecurityGroupIds(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSeriesCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStartTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetUpgradeStatus(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.UpgradeStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetUpgradeable(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Upgradeable = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVswitchId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) Validate() error {
	return dara.Validate(s)
}
