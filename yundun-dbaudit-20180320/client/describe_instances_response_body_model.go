// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() []*DescribeInstancesResponseBodyInstances
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeInstancesResponseBody
	GetTotalCount() *int64
}

type DescribeInstancesResponseBody struct {
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() []*DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsStatus         *string `json:"EcsStatus,omitempty" xml:"EcsStatus,omitempty"`
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ImageVersionName  *string `json:"ImageVersionName,omitempty" xml:"ImageVersionName,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus    *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InternetEndpoint  *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	InternetIp        *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IntranetEndpoint  *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	IntranetIp        *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Legacy            *bool   `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	LicenseCode       *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Operatable        *bool   `json:"Operatable,omitempty" xml:"Operatable,omitempty"`
	PlanUpgradeStatus *int32  `json:"PlanUpgradeStatus,omitempty" xml:"PlanUpgradeStatus,omitempty"`
	PlanUpgradeable   *bool   `json:"PlanUpgradeable,omitempty" xml:"PlanUpgradeable,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Renewable         *bool   `json:"Renewable,omitempty" xml:"Renewable,omitempty"`
	SeriesCode        *string `json:"SeriesCode,omitempty" xml:"SeriesCode,omitempty"`
	StartTime         *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UpgradeStatus     *int32  `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	Upgradeable       *bool   `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId         *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstancesResponseBodyInstances) GetEcsStatus() *string {
	return s.EcsStatus
}

func (s *DescribeInstancesResponseBodyInstances) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeInstancesResponseBodyInstances) GetImageVersionName() *string {
	return s.ImageVersionName
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstancesResponseBodyInstances) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *DescribeInstancesResponseBodyInstances) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeInstancesResponseBodyInstances) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *DescribeInstancesResponseBodyInstances) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeInstancesResponseBodyInstances) GetLegacy() *bool {
	return s.Legacy
}

func (s *DescribeInstancesResponseBodyInstances) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeInstancesResponseBodyInstances) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstancesResponseBodyInstances) GetOperatable() *bool {
	return s.Operatable
}

func (s *DescribeInstancesResponseBodyInstances) GetPlanUpgradeStatus() *int32 {
	return s.PlanUpgradeStatus
}

func (s *DescribeInstancesResponseBodyInstances) GetPlanUpgradeable() *bool {
	return s.PlanUpgradeable
}

func (s *DescribeInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstances) GetRenewable() *bool {
	return s.Renewable
}

func (s *DescribeInstancesResponseBodyInstances) GetSeriesCode() *string {
	return s.SeriesCode
}

func (s *DescribeInstancesResponseBodyInstances) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeInstancesResponseBodyInstances) GetUpgradeStatus() *int32 {
	return s.UpgradeStatus
}

func (s *DescribeInstancesResponseBodyInstances) GetUpgradeable() *bool {
	return s.Upgradeable
}

func (s *DescribeInstancesResponseBodyInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstances) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstancesResponseBodyInstances) SetDescription(v string) *DescribeInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEcsStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.EcsStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetImageVersionName(v string) *DescribeInstancesResponseBodyInstances {
	s.ImageVersionName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInternetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInternetIp(v string) *DescribeInstancesResponseBodyInstances {
	s.InternetIp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIntranetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIntranetIp(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLegacy(v bool) *DescribeInstancesResponseBodyInstances {
	s.Legacy = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLicenseCode(v string) *DescribeInstancesResponseBodyInstances {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetNetworkType(v string) *DescribeInstancesResponseBodyInstances {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetOperatable(v bool) *DescribeInstancesResponseBodyInstances {
	s.Operatable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPlanUpgradeStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.PlanUpgradeStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPlanUpgradeable(v bool) *DescribeInstancesResponseBodyInstances {
	s.PlanUpgradeable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegionId(v string) *DescribeInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRenewable(v bool) *DescribeInstancesResponseBodyInstances {
	s.Renewable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSeriesCode(v string) *DescribeInstancesResponseBodyInstances {
	s.SeriesCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUpgradeStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.UpgradeStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUpgradeable(v bool) *DescribeInstancesResponseBodyInstances {
	s.Upgradeable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVswitchId(v string) *DescribeInstancesResponseBodyInstances {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
