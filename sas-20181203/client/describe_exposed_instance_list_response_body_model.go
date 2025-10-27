// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExposedInstances(v []*DescribeExposedInstanceListResponseBodyExposedInstances) *DescribeExposedInstanceListResponseBody
	GetExposedInstances() []*DescribeExposedInstanceListResponseBodyExposedInstances
	SetPageInfo(v *DescribeExposedInstanceListResponseBodyPageInfo) *DescribeExposedInstanceListResponseBody
	GetPageInfo() *DescribeExposedInstanceListResponseBodyPageInfo
	SetRequestId(v string) *DescribeExposedInstanceListResponseBody
	GetRequestId() *string
}

type DescribeExposedInstanceListResponseBody struct {
	// The details of the exposures.
	ExposedInstances []*DescribeExposedInstanceListResponseBodyExposedInstances `json:"ExposedInstances,omitempty" xml:"ExposedInstances,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeExposedInstanceListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 598A4A61-ABA7-456B-8725-7378258276D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExposedInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponseBody) GetExposedInstances() []*DescribeExposedInstanceListResponseBodyExposedInstances {
	return s.ExposedInstances
}

func (s *DescribeExposedInstanceListResponseBody) GetPageInfo() *DescribeExposedInstanceListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeExposedInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExposedInstanceListResponseBody) SetExposedInstances(v []*DescribeExposedInstanceListResponseBodyExposedInstances) *DescribeExposedInstanceListResponseBody {
	s.ExposedInstances = v
	return s
}

func (s *DescribeExposedInstanceListResponseBody) SetPageInfo(v *DescribeExposedInstanceListResponseBodyPageInfo) *DescribeExposedInstanceListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeExposedInstanceListResponseBody) SetRequestId(v string) *DescribeExposedInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBody) Validate() error {
	if s.ExposedInstances != nil {
		for _, item := range s.ExposedInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeExposedInstanceListResponseBodyExposedInstances struct {
	// The number of high-severity vulnerabilities that are exposed on the Internet and can be exploited by attackers.
	//
	// example:
	//
	// 0
	AsapVulCount *int32 `json:"AsapVulCount,omitempty" xml:"AsapVulCount,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **0**: an ECS instance.
	//
	// 	- **1**: a SLB instance.
	//
	// 	- **2**: a NAT gateway.
	//
	// 	- **3**: an ApsaraDB RDS instance.
	//
	// 	- **4**: an ApsaraDB for MongoDB instance.
	//
	// 	- **5**: an ApsaraDB for Redis instance.
	//
	// 	- **6**: a container image.
	//
	// 	- **7**: a container.
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The JSON string that specifies the information about a database asset, which contains the following fields.
	//
	// 	- assetSubType: the asset subtype.
	//
	// 	- assetSubTypeName: the name of the asset subtype.
	//
	// 	- assetType: the type of the asset.
	//
	// 	- assetTypeName: the name of the asset type.
	//
	// 	- vendor: the service provider of the asset.
	//
	// example:
	//
	// {assetSubTypeName":"INSTANCE","assetType":3,"assetTypeName":"RDS","vendor":0}
	CloudAssetInfo *string `json:"CloudAssetInfo,omitempty" xml:"CloudAssetInfo,omitempty"`
	// The number of CSPM risks.
	//
	// example:
	//
	// 0
	CspmAlarmCount *int32 `json:"CspmAlarmCount,omitempty" xml:"CspmAlarmCount,omitempty"`
	// The number of weak password risks.
	//
	// example:
	//
	// 0
	ExploitHealthCount *int32 `json:"ExploitHealthCount,omitempty" xml:"ExploitHealthCount,omitempty"`
	// The server component that is exposed on the Internet.
	//
	// example:
	//
	// openssl,openssh
	ExposureComponent *string `json:"ExposureComponent,omitempty" xml:"ExposureComponent,omitempty"`
	// Expose component information list.
	ExposureComponentList []*DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList `json:"ExposureComponentList,omitempty" xml:"ExposureComponentList,omitempty" type:"Repeated"`
	// The public IP address that is exposed on the Internet.
	//
	// example:
	//
	// 116.12.XX.XX
	ExposureIp *string `json:"ExposureIp,omitempty" xml:"ExposureIp,omitempty"`
	// The port that is exposed on the Internet.
	//
	// example:
	//
	// 22
	ExposurePort *string `json:"ExposurePort,omitempty" xml:"ExposurePort,omitempty"`
	// The resource from which the asset is exposed. Valid values:
	//
	// 	- **INTERNET_IP**: the public IP address of an ECS instance.
	//
	// 	- **SLB**: the public IP address of a Server Load Balancer (SLB) instance.
	//
	// 	- **EIP**: an elastic IP address (EIP).
	//
	// 	- **DNAT**: the NAT gateway that connects to the Internet by using the Destination Network Address Translation (DNAT) feature.
	//
	// 	- **DB_CONNECTION**: the public endpoint of a database.
	//
	// example:
	//
	// INTERNET_IP
	ExposureType *string `json:"ExposureType,omitempty" xml:"ExposureType,omitempty"`
	// The ID of the instance to which the resource belongs. The valid values of this parameter vary based on the value of the ExposureType parameter.
	//
	// 	- If the value of the ExposureType parameter is **INTERNET_IP**, this parameter is empty.
	//
	// 	- If the value of the ExposureType parameter is **SLB**, the value of this parameter is the ID of the SLB instance.
	//
	// 	- If the value of the ExposureType parameter is **EIP**, the value of this parameter is the ID of the EIP.
	//
	// 	- If the value of the ExposureType parameter is **DNAT**, the value of this parameter is the ID of the NAT gateway.
	//
	// 	- If the value of the ExposureType parameter is **DB_CONNECTION**, the value of this parameter is the ID of the database.
	//
	// example:
	//
	// i-ew11313a****
	ExposureTypeId *string `json:"ExposureTypeId,omitempty" xml:"ExposureTypeId,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// 9469268
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the server group.
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// i-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// abc_centos7.2_005
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 116.12.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The number of medium-severity vulnerabilities that are exposed on the Internet and can be exploited by attackers.
	//
	// example:
	//
	// 0
	LaterVulCount *int32 `json:"LaterVulCount,omitempty" xml:"LaterVulCount,omitempty"`
	// The number of low-severity vulnerabilities that are exposed on the Internet and can be exploited by attackers.
	//
	// example:
	//
	// 0
	NntfVulCount *int32 `json:"NntfVulCount,omitempty" xml:"NntfVulCount,omitempty"`
	// The ID of the region in which the asset resides.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The total number of vulnerabilities that are exposed on the Internet and can be exploited by attackers.
	//
	// example:
	//
	// 0
	TotalVulCount *int32 `json:"TotalVulCount,omitempty" xml:"TotalVulCount,omitempty"`
	// The UUID of the server or the instance ID of the cloud service.
	//
	// example:
	//
	// dd803d9e-a337-4add-9c5b-7d503e08****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExposedInstanceListResponseBodyExposedInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceListResponseBodyExposedInstances) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetAsapVulCount() *int32 {
	return s.AsapVulCount
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetAssetType() *int32 {
	return s.AssetType
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetCloudAssetInfo() *string {
	return s.CloudAssetInfo
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetCspmAlarmCount() *int32 {
	return s.CspmAlarmCount
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetExploitHealthCount() *int32 {
	return s.ExploitHealthCount
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetExposureComponent() *string {
	return s.ExposureComponent
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetExposureComponentList() []*DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList {
	return s.ExposureComponentList
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetExposureIp() *string {
	return s.ExposureIp
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetExposurePort() *string {
	return s.ExposurePort
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetExposureType() *string {
	return s.ExposureType
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetExposureTypeId() *string {
	return s.ExposureTypeId
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetLaterVulCount() *int32 {
	return s.LaterVulCount
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetNntfVulCount() *int32 {
	return s.NntfVulCount
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetTotalVulCount() *int32 {
	return s.TotalVulCount
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetAsapVulCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.AsapVulCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetAssetType(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.AssetType = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetCloudAssetInfo(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.CloudAssetInfo = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetCspmAlarmCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.CspmAlarmCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExploitHealthCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExploitHealthCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureComponent(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureComponent = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureComponentList(v []*DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureComponentList = v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureIp(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureIp = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposurePort(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposurePort = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureType(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureType = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureTypeId(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureTypeId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetGroupId(v int64) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.GroupId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetGroupName(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.GroupName = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetInstanceId(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetInstanceName(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetInternetIp(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.InternetIp = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetIntranetIp(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetLaterVulCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.LaterVulCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetNntfVulCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.NntfVulCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetRegionId(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetTotalVulCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.TotalVulCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetUuid(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) Validate() error {
	if s.ExposureComponentList != nil {
		for _, item := range s.ExposureComponentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList struct {
	// Expose component type.
	//
	// example:
	//
	// system_service
	ComponentBizType *string `json:"ComponentBizType,omitempty" xml:"ComponentBizType,omitempty"`
	// Expose components.
	//
	// example:
	//
	// openssh
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// Expose component version.
	//
	// example:
	//
	// 8.7p1
	ComponentVersion *string `json:"ComponentVersion,omitempty" xml:"ComponentVersion,omitempty"`
	// Exposed port.
	//
	// example:
	//
	// 22
	ListenPort *string `json:"ListenPort,omitempty" xml:"ListenPort,omitempty"`
}

func (s DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) GetComponentBizType() *string {
	return s.ComponentBizType
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) GetComponentVersion() *string {
	return s.ComponentVersion
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) GetListenPort() *string {
	return s.ListenPort
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) SetComponentBizType(v string) *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList {
	s.ComponentBizType = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) SetComponentName(v string) *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList {
	s.ComponentName = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) SetComponentVersion(v string) *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList {
	s.ComponentVersion = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) SetListenPort(v string) *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList {
	s.ListenPort = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstancesExposureComponentList) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedInstanceListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExposedInstanceListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) SetCount(v int32) *DescribeExposedInstanceListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeExposedInstanceListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) SetPageSize(v int32) *DescribeExposedInstanceListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeExposedInstanceListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
