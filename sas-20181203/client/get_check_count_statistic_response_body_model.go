// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckCountStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCountStatisticDTO(v *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) *GetCheckCountStatisticResponseBody
	GetCheckCountStatisticDTO() *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO
	SetRequestId(v string) *GetCheckCountStatisticResponseBody
	GetRequestId() *string
}

type GetCheckCountStatisticResponseBody struct {
	// The risk item statistics.
	CheckCountStatisticDTO *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO `json:"CheckCountStatisticDTO,omitempty" xml:"CheckCountStatisticDTO,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCheckCountStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckCountStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckCountStatisticResponseBody) GetCheckCountStatisticDTO() *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO {
	return s.CheckCountStatisticDTO
}

func (s *GetCheckCountStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckCountStatisticResponseBody) SetCheckCountStatisticDTO(v *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) *GetCheckCountStatisticResponseBody {
	s.CheckCountStatisticDTO = v
	return s
}

func (s *GetCheckCountStatisticResponseBody) SetRequestId(v string) *GetCheckCountStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckCountStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCheckCountStatisticResponseBodyCheckCountStatisticDTO struct {
	// The risk item statistics.
	CheckCountStatisticItems []*GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems `json:"CheckCountStatisticItems,omitempty" xml:"CheckCountStatisticItems,omitempty" type:"Repeated"`
	// The type of the statistics. Valid values:
	//
	// 	- **user**: the top five users that are granted excessive permissions.
	//
	// 	- **role**: the top five roles that are granted excessive permissions.
	//
	// 	- **instance**: the top five cloud services on which risks are detected.
	//
	// example:
	//
	// instance
	StatisticType *string `json:"StatisticType,omitempty" xml:"StatisticType,omitempty"`
}

func (s GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) String() string {
	return dara.Prettify(s)
}

func (s GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) GoString() string {
	return s.String()
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) GetCheckCountStatisticItems() []*GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	return s.CheckCountStatisticItems
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) GetStatisticType() *string {
	return s.StatisticType
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) SetCheckCountStatisticItems(v []*GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO {
	s.CheckCountStatisticItems = v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) SetStatisticType(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO {
	s.StatisticType = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTO) Validate() error {
	return dara.Validate(s)
}

type GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems struct {
	// The number of the CPU cores used by the host instance.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The instance ID of the cloud service.
	//
	// example:
	//
	// i-wz9bpxyu6t74qn9g****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the asset.
	//
	// example:
	//
	// launch-advisor-2021****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The subtype of the cloud service.
	//
	// example:
	//
	// 0
	InstanceSubType *int32 `json:"InstanceSubType,omitempty" xml:"InstanceSubType,omitempty"`
	// The asset subtype of the cloud service. Valid values:
	//
	// 	- If **InstanceTypeName*	- is set to **ECS**, this parameter supports the following valid values:
	//
	//     	- **INSTANCE**
	//
	//     	- **DISK**
	//
	//     	- **SECURITY_GROUP**
	//
	// 	- If **InstanceTypeName*	- is set to **ACR**, this parameter supports the following valid values:
	//
	//     	- **REPOSITORY_ENTERPRISE**
	//
	//     	- **REPOSITORY_PERSON**
	//
	// 	- If **InstanceTypeName*	- is set to **RAM**, this parameter supports the following valid values:
	//
	//     	- **ALIAS**
	//
	//     	- **USER**
	//
	//     	- **POLICY**
	//
	//     	- **GROUP**
	//
	// 	- If **InstanceTypeName*	- is set to **WAF**, this parameter supports the following valid value:
	//
	//     	- **DOMAIN**
	//
	// 	- If **InstanceTypeName*	- is set to other values, this parameter supports the following valid values:
	//
	//     	- **INSTANCE**
	//
	// example:
	//
	// INSTANCE
	InstanceSubTypeName *string `json:"InstanceSubTypeName,omitempty" xml:"InstanceSubTypeName,omitempty"`
	// The asset type. Valid values:
	//
	// 	- **0**: Elastic Compute Service (ECS) instance.
	//
	// 	- **1**: Server Load Balancer (SLB) instance.
	//
	// 	- **2**: NAT gateway.
	//
	// 	- **3**: ApsaraDB RDS instance.
	//
	// 	- **4**: ApsaraDB for MongoDB (MongoDB) instance.
	//
	// 	- **5**: Tair (Redis OSS-compatible) (Tair) instance.
	//
	// 	- **6**: container image.
	//
	// 	- **7**: container.
	//
	// example:
	//
	// 0
	InstanceType *int32 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The asset type of the cloud service. Valid values:
	//
	// 	- **ECS**: ECS.
	//
	// 	- **SLB**: SLB.
	//
	// 	- **RDS**: ApsaraDB RDS.
	//
	// 	- **MONGODB**: ApsaraDB for MongoDB.
	//
	// 	- **KVSTORE**: Tair.
	//
	// 	- **ACR**: Container Registry.
	//
	// 	- **CSK**: Container Service for Kubernetes (ACK).
	//
	// 	- **VPC**: Virtual Private Cloud (VPC).
	//
	// 	- **ACTIONTRAIL**: ActionTrail.
	//
	// 	- **CDN**: Alibaba Cloud CDN (CDN).
	//
	// 	- **CAS**: Certificate Management Service (formerly SSL Certificates Service).
	//
	// 	- **RDC**: Alibaba Cloud DevOps.
	//
	// 	- **RAM**: Resource Access Management (RAM).
	//
	// 	- **DDOS**: Anti-DDoS.
	//
	// 	- **WAF**: Web Application Firewall (WAF).
	//
	// 	- **OSS**: Object Storage Service (OSS).
	//
	// 	- **POLARDB**: PolarDB.
	//
	// 	- **POSTGRESQL**: ApsaraDB RDS for PostgreSQL.
	//
	// 	- **MSE**: Microservices Engine (MSE).
	//
	// 	- **NAS**: File Storage NAS (NAS).
	//
	// 	- **SDDP**: Sensitive Data Discovery and Protection (SDDP).
	//
	// 	- **EIP**: Elastic IP Address (EIP).
	//
	// example:
	//
	// ECS
	InstanceTypeName *string `json:"InstanceTypeName,omitempty" xml:"InstanceTypeName,omitempty"`
	// The public IP address of the host instance.
	//
	// example:
	//
	// 1.2.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the host instance.
	//
	// example:
	//
	// 1.2.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The version of the operating system that the host instance runs.
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of risk items.
	//
	// example:
	//
	// 22
	RiskCount *int32 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The UUID of the host instance.
	//
	// example:
	//
	// c9107c04-942f-40c1-981a-f1c1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The cloud service provider. Valid values:
	//
	// 	- **ALIYUN**: Alibaba Cloud.
	//
	// 	- **TENCENT**: Tencent Cloud.
	//
	// 	- **MICROSOFT**: Microsoft Azure.
	//
	// 	- **AWS**: AWS.
	//
	// example:
	//
	// ALIYUN
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The ID of the VPC to which the host instance belongs.
	//
	// example:
	//
	// vpc-uf60agqq65bs98zoo****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) String() string {
	return dara.Prettify(s)
}

func (s GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GoString() string {
	return s.String()
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetCores() *int32 {
	return s.Cores
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetInstanceSubType() *int32 {
	return s.InstanceSubType
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetInstanceSubTypeName() *string {
	return s.InstanceSubTypeName
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetInstanceType() *int32 {
	return s.InstanceType
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetInstanceTypeName() *string {
	return s.InstanceTypeName
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetOs() *string {
	return s.Os
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetUuid() *string {
	return s.Uuid
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetVendor() *int32 {
	return s.Vendor
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetCores(v int32) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.Cores = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetInstanceId(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.InstanceId = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetInstanceName(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.InstanceName = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetInstanceSubType(v int32) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.InstanceSubType = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetInstanceSubTypeName(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.InstanceSubTypeName = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetInstanceType(v int32) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.InstanceType = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetInstanceTypeName(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.InstanceTypeName = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetInternetIp(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.InternetIp = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetIntranetIp(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.IntranetIp = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetOs(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.Os = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetRegionId(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.RegionId = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetRiskCount(v int32) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.RiskCount = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetUuid(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.Uuid = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetVendor(v int32) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.Vendor = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) SetVpcInstanceId(v string) *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems {
	s.VpcInstanceId = &v
	return s
}

func (s *GetCheckCountStatisticResponseBodyCheckCountStatisticDTOCheckCountStatisticItems) Validate() error {
	return dara.Validate(s)
}
