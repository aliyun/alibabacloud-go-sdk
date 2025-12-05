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
	// An array that consists of the bastion hosts returned.
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 61D36C55-AAFC-4678-8FAD-34FEF9E7182E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of bastion hosts that are queried.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	BandWidth *int64 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The remarks of the bastion host.
	//
	// example:
	//
	// Test API
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp when the bastion host expires. Unit: milliseconds.
	//
	// example:
	//
	// 1578326400000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The image version of the bastion host.
	//
	// example:
	//
	// 3.0.0
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The bastion host ID.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the bastion host. Valid values:
	//
	// 	- **PENDING**: The bastion host is not initialized.
	//
	// 	- **CREATING**: The bastion host is being created.
	//
	// 	- **RUNNING**: The bastion host is running.
	//
	// 	- **EXPIRED**: The bastion host expired.
	//
	// 	- **CREATE_FAILED**: The bastion host fails to be created.
	//
	// 	- **UPGRADING**: The configurations of the bastion host are being changed.
	//
	// 	- **UPGRADE_FAILED**: The configurations of the bastion host fail to be changed.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The public O\\&M address of the bastion host.
	//
	// example:
	//
	// drawvxalwb-public.bastionhost.aliyuncs.com
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The private O\\&M address of the bastion host.
	//
	// example:
	//
	// drawvxalwb.bastionhost.aliyuncs.com
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// Indicates whether the bastion host runs an earlier version. Valid values:
	//
	// 	- **true**: The bastion host runs V2 or V3.1.
	//
	// 	- **false**: The bastion host runs V3.2.
	//
	// example:
	//
	// false
	Legacy *bool `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	// The license code of the bastion host.
	//
	// example:
	//
	// bhah_ent_50_asset
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// The edition of the bastion host. Valid values:
	//
	// 	- **cloudbastion**: Basic Edition.
	//
	// 	- **cloudbastion_ha**: Enterprise Edition.
	//
	// example:
	//
	// cloudbastion_ha
	PlanCode *string `json:"PlanCode,omitempty" xml:"PlanCode,omitempty"`
	// Indicates whether the bastion host can be accessed from the Internet. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PublicNetworkAccess *bool `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	// The region ID of the bastion host.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the bastion host belongs.
	//
	// example:
	//
	// g-acfm26ougi****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveVswitchId  *string `json:"SlaveVswitchId,omitempty" xml:"SlaveVswitchId,omitempty"`
	// The timestamp when the bastion host is purchased or renewed. Unit: milliseconds.
	//
	// example:
	//
	// 1577681345000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the bastion host belongs.
	//
	// example:
	//
	// vpc-bp1c85tzgqu1bf5b****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch to which the bastion host belongs.
	//
	// example:
	//
	// vsw-bp1xfwzzfti0kjbf****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetBandWidth() *int64 {
	return s.BandWidth
}

func (s *DescribeInstancesResponseBodyInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstancesResponseBodyInstances) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeInstancesResponseBodyInstances) GetImageVersion() *string {
	return s.ImageVersion
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

func (s *DescribeInstancesResponseBodyInstances) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *DescribeInstancesResponseBodyInstances) GetLegacy() *bool {
	return s.Legacy
}

func (s *DescribeInstancesResponseBodyInstances) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeInstancesResponseBodyInstances) GetPlanCode() *string {
	return s.PlanCode
}

func (s *DescribeInstancesResponseBodyInstances) GetPublicNetworkAccess() *bool {
	return s.PublicNetworkAccess
}

func (s *DescribeInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesResponseBodyInstances) GetSlaveVswitchId() *string {
	return s.SlaveVswitchId
}

func (s *DescribeInstancesResponseBodyInstances) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeInstancesResponseBodyInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstances) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstancesResponseBodyInstances) SetBandWidth(v int64) *DescribeInstancesResponseBodyInstances {
	s.BandWidth = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDescription(v string) *DescribeInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetImageVersion(v string) *DescribeInstancesResponseBodyInstances {
	s.ImageVersion = &v
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

func (s *DescribeInstancesResponseBodyInstances) SetIntranetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetEndpoint = &v
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

func (s *DescribeInstancesResponseBodyInstances) SetPlanCode(v string) *DescribeInstancesResponseBodyInstances {
	s.PlanCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPublicNetworkAccess(v bool) *DescribeInstancesResponseBodyInstances {
	s.PublicNetworkAccess = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegionId(v string) *DescribeInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSlaveVswitchId(v string) *DescribeInstancesResponseBodyInstances {
	s.SlaveVswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
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
