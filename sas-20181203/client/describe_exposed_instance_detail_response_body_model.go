// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExposedChains(v []*DescribeExposedInstanceDetailResponseBodyExposedChains) *DescribeExposedInstanceDetailResponseBody
	GetExposedChains() []*DescribeExposedInstanceDetailResponseBodyExposedChains
	SetRequestId(v string) *DescribeExposedInstanceDetailResponseBody
	GetRequestId() *string
}

type DescribeExposedInstanceDetailResponseBody struct {
	// The list of exposure details of the server or database.
	ExposedChains []*DescribeExposedInstanceDetailResponseBodyExposedChains `json:"ExposedChains,omitempty" xml:"ExposedChains,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C590482B-54A7-4273-8115-9DBE2DE46B26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExposedInstanceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBody) GetExposedChains() []*DescribeExposedInstanceDetailResponseBodyExposedChains {
	return s.ExposedChains
}

func (s *DescribeExposedInstanceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExposedInstanceDetailResponseBody) SetExposedChains(v []*DescribeExposedInstanceDetailResponseBodyExposedChains) *DescribeExposedInstanceDetailResponseBody {
	s.ExposedChains = v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBody) SetRequestId(v string) *DescribeExposedInstanceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedInstanceDetailResponseBodyExposedChains struct {
	// The information about all vulnerabilities on the server.
	AllVulList []*DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList `json:"AllVulList,omitempty" xml:"AllVulList,omitempty" type:"Repeated"`
	// The list of configuration risks.
	CspmRiskList []*DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList `json:"CspmRiskList,omitempty" xml:"CspmRiskList,omitempty" type:"Repeated"`
	// The server component that is exposed on the Internet.
	//
	// example:
	//
	// openssl,openssh
	ExposureComponent *string `json:"ExposureComponent,omitempty" xml:"ExposureComponent,omitempty"`
	// The IP address of the server or the public endpoint of the database.
	//
	// example:
	//
	// 47.99.XX.XX
	ExposureIp *string `json:"ExposureIp,omitempty" xml:"ExposureIp,omitempty"`
	// The port that is exposed on the Internet.
	//
	// example:
	//
	// 22
	ExposurePort *string `json:"ExposurePort,omitempty" xml:"ExposurePort,omitempty"`
	// The resource from which the server or database is exposed. Valid values:
	//
	// 	- **INTERNET_IP**: the public IP address of an Elastic Compute Service (ECS) instance.
	//
	// 	- **SLB**: the public IP address of a Server Load Balancer (SLB) instance.
	//
	// 	- **EIP**: an elastic IP address (EIP).
	//
	// 	- **DNAT**: the Network Address Translation (NAT) gateway that connects to the Internet by using the Destination Network Address Translation (DNAT) feature
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
	// eip-bp1bkgowzam49rld3****
	ExposureTypeId *string `json:"ExposureTypeId,omitempty" xml:"ExposureTypeId,omitempty"`
	// The server group to which the server belongs.
	//
	// example:
	//
	// sg-bp1iw5enua6gf5i2xr7z
	GroupNo *string `json:"GroupNo,omitempty" xml:"GroupNo,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp116qem8npvchqc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// worker-k8s-for-cs-c929ee2a145214f89a8b248005be5****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 47.99.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The information about the vulnerabilities that are exposed on the Internet and can be exploited by attackers.
	RealVulList []*DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList `json:"RealVulList,omitempty" xml:"RealVulList,omitempty" type:"Repeated"`
	// The region ID.
	//
	// >  For information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The UUID of the server or the instance ID of the database.
	//
	// example:
	//
	// 4f9ce097-4a7d-48fe-baef-6960e5b6****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChains) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChains) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetAllVulList() []*DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	return s.AllVulList
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetCspmRiskList() []*DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	return s.CspmRiskList
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetExposureComponent() *string {
	return s.ExposureComponent
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetExposureIp() *string {
	return s.ExposureIp
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetExposurePort() *string {
	return s.ExposurePort
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetExposureType() *string {
	return s.ExposureType
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetExposureTypeId() *string {
	return s.ExposureTypeId
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetGroupNo() *string {
	return s.GroupNo
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetRealVulList() []*DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	return s.RealVulList
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetAllVulList(v []*DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.AllVulList = v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetCspmRiskList(v []*DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.CspmRiskList = v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposureComponent(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposureComponent = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposureIp(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposureIp = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposurePort(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposurePort = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposureType(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposureType = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposureTypeId(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposureTypeId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetGroupNo(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.GroupNo = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetInstanceId(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.InstanceId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetInstanceName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.InstanceName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetInternetIp(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.InternetIp = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetIntranetIp(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.IntranetIp = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetRealVulList(v []*DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.RealVulList = v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetRegionId(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.RegionId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetUuid(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2022:0274-Important: polkit pkexec Local Privilege Escalation Vulnerability(CVE-2021-4034)
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// SCA:ACSV-2020-052801
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority to fix the vulnerability. Valid values:
	//
	// 	- **asap**: high
	//
	// 	- **later**: medium
	//
	// 	- **nntf**: low
	//
	// >  We recommend that you fix the vulnerabilities that have the **high*	- priority at the earliest opportunity.
	//
	// example:
	//
	// asap
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerabilities
	//
	// 	- **sys**: Windows system vulnerabilities
	//
	// 	- **cms**: Web-CMS vulnerabilities
	//
	// 	- **app**: application vulnerabilities
	//
	// 	- **emg**: urgent vulnerabilities
	//
	// 	- **sca**: middleware vulnerabilities
	//
	// example:
	//
	// sca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 4f9ce097-4a7d-48fe-baef-6960e5b6****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) GetName() *string {
	return s.Name
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) GetType() *string {
	return s.Type
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetAliasName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.AliasName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.Name = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetNecessity(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.Necessity = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetType(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.Type = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetUuid(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList struct {
	// The subtype of the cloud asset. Valid values:
	//
	// 	- **0**: Elastic Compute Service (ECS).
	//
	//     	- **100**: instance.
	//
	// 	- **3**: ApsaraDB RDS.
	//
	//     	- **0**: instance.
	//
	// 	- **4**: ApsaraDB for MongoDB (MongoDB).
	//
	//     	- **0**: instance.
	//
	// 	- **5**: ApsaraDB for Redis (Redis).
	//
	//     	- **0**: instance.
	//
	// example:
	//
	// 100
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The subtype name of the cloud asset. Valid values:
	//
	// 	- **INSTANCE**: MongoDB instance, Apsara DB for RDS instance, and ApsaraDB for Redis instance.
	//
	// 	- **ECS_INSTANCE**: ECS instance.
	//
	// example:
	//
	// INSTANCE
	AssetSubTypeName *string `json:"AssetSubTypeName,omitempty" xml:"AssetSubTypeName,omitempty"`
	// The instance type. Valid values:
	//
	// 	- 0: an ECS instance.
	//
	// 	- 3: an ApsaraDB RDS instance.
	//
	// 	- 4: an ApsaraDB for MongoDB instance.
	//
	// 	- 5: an ApsaraDB for Redis instance.
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The name of the cloud asset type. Valid values:
	//
	// 	- **ECS**
	//
	// 	- **RDS**
	//
	// 	- **KVSTORE**
	//
	// 	- **MONGODB**
	//
	// example:
	//
	// ECS
	AssetTypeName *string `json:"AssetTypeName,omitempty" xml:"AssetTypeName,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// Create Alert Rule
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp14ggqzi9k6ocfb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **HIGH**
	//
	// 	- **MEDIUM**
	//
	// 	- **LOW**
	//
	// example:
	//
	// HIGH
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The type of the cloud asset by source. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetAssetSubTypeName() *string {
	return s.AssetSubTypeName
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetAssetType() *int32 {
	return s.AssetType
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetAssetTypeName() *string {
	return s.AssetTypeName
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetCheckName() *string {
	return s.CheckName
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) GetVendor() *int32 {
	return s.Vendor
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetAssetSubType(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.AssetSubType = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetAssetSubTypeName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.AssetSubTypeName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetAssetType(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.AssetType = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetAssetTypeName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.AssetTypeName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetCheckName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.CheckName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetInstanceId(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.InstanceId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetRegionId(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.RegionId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetRiskLevel(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) SetVendor(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList {
	s.Vendor = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsCspmRiskList) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2022:0274-Important: polkit pkexec Local Privilege Escalation Vulnerability(CVE-2021-4034)
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// SCA:ACSV-2020-052801
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority to fix the vulnerability. Valid values:
	//
	// 	- **asap**: high
	//
	// 	- **later**: medium
	//
	// 	- **nntf**: low
	//
	// >  We recommend that you fix the vulnerabilities that have the **high*	- priority at the earliest opportunity.
	//
	// example:
	//
	// asap
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerabilities
	//
	// 	- **sys**: Windows system vulnerabilities
	//
	// 	- **cms**: Web-CMS vulnerabilities
	//
	// 	- **app**: application vulnerabilities
	//
	// 	- **emg**: urgent vulnerabilities
	//
	// 	- **sca**: middleware vulnerabilities
	//
	// example:
	//
	// sca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 4f9ce097-4a7d-48fe-baef-6960e5b6****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) GetName() *string {
	return s.Name
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) GetType() *string {
	return s.Type
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetAliasName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.AliasName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.Name = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetNecessity(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.Necessity = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetType(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.Type = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetUuid(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) Validate() error {
	return dara.Validate(s)
}
