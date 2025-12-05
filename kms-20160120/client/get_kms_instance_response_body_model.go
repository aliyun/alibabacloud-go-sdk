// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKmsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKmsInstance(v *GetKmsInstanceResponseBodyKmsInstance) *GetKmsInstanceResponseBody
	GetKmsInstance() *GetKmsInstanceResponseBodyKmsInstance
	SetRequestId(v string) *GetKmsInstanceResponseBody
	GetRequestId() *string
}

type GetKmsInstanceResponseBody struct {
	// The details of the KMS instance.
	KmsInstance *GetKmsInstanceResponseBodyKmsInstance `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 46b4a94a-57d2-44b4-9810-1e87d31abb33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKmsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponseBody) GetKmsInstance() *GetKmsInstanceResponseBodyKmsInstance {
	return s.KmsInstance
}

func (s *GetKmsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKmsInstanceResponseBody) SetKmsInstance(v *GetKmsInstanceResponseBodyKmsInstance) *GetKmsInstanceResponseBody {
	s.KmsInstance = v
	return s
}

func (s *GetKmsInstanceResponseBody) SetRequestId(v string) *GetKmsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKmsInstanceResponseBody) Validate() error {
	if s.KmsInstance != nil {
		if err := s.KmsInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKmsInstanceResponseBodyKmsInstance struct {
	// A list of associated VPCs.
	//
	// >  If your self-managed applications are deployed in multiple VPCs in the same region, you can associate VPCs with the KMS instance beyond the VPC that you specify when you enable the KMS instance. The VPCs can belong to the same Alibaba Cloud account or different Alibaba Cloud accounts. After the configuration is complete, self-managed applications in the VPCs can access the specified KMS instance.
	BindVpcs *GetKmsInstanceResponseBodyKmsInstanceBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Struct"`
	// The content of the certificate authority (CA) certificate of the KMS instance.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\r\\nMIIDuzCCAqOgAwIBAgIJALTKwWAjvbMiMA0GCSqGSIb3DQEBCwUAMHQxCzAJBgNV****-----END CERTIFICATE-----
	CaCertificateChainPem *string `json:"CaCertificateChainPem,omitempty" xml:"CaCertificateChainPem,omitempty"`
	ChargeType            *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the KMS instance is created.
	//
	// example:
	//
	// 2023-09-05T12:44:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the KMS instance.
	//
	// example:
	//
	// 2023-10-05T16:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The ID of the KMS instance.
	//
	// example:
	//
	// kst-bjj62f5ba3dnpb6v8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the KMS instance.
	//
	// example:
	//
	// kst-bjj62f5ba3dnpb6v8****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of keys that can be created for the KMS instance.
	//
	// example:
	//
	// 1000
	KeyNum         *int64  `json:"KeyNum,omitempty" xml:"KeyNum,omitempty"`
	Log            *int64  `json:"Log,omitempty" xml:"Log,omitempty"`
	LogStorage     *int64  `json:"LogStorage,omitempty" xml:"LogStorage,omitempty"`
	ProductType    *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	SaleStatus     *string `json:"SaleStatus,omitempty" xml:"SaleStatus,omitempty"`
	// The number of secrets that can be created for the KMS instance.
	//
	// example:
	//
	// 10
	SecretNum *string `json:"SecretNum,omitempty" xml:"SecretNum,omitempty"`
	// The computing performance of the KMS instance.
	//
	// example:
	//
	// 1000
	Spec *int64 `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The time when the KMS instance is enabled.
	//
	// example:
	//
	// 2023-09-05T12:44:19Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The status of the KMS instance. Valid values:
	//
	// 	- Uninitialized: The KMS instance is not enabled.
	//
	// 	- Connecting: The KMS instance is being connected.
	//
	// 	- Connected: The KMS instance is enabled.
	//
	// 	- Disconnected: The KMS instance is disconnected.
	//
	// 	- Error: The KMS instance is abnormal.
	//
	// example:
	//
	// Connected
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The virtual private cloud (VPC) with which the KMS instance is associated.
	//
	// example:
	//
	// vpc-bp19z7cwmltad5dff****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The access management quota for the KMS instance.
	//
	// example:
	//
	// 5
	VpcNum *int64 `json:"VpcNum,omitempty" xml:"VpcNum,omitempty"`
	// The vSwitch in the VPC.
	//
	// example:
	//
	// vsw-bp1i512amda6d10a0****
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
	// The zone with which the KMS instance is associated.
	//
	// example:
	//
	// "cn-hangzhou-k",       "cn-hangzhou-j"
	ZoneIds []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s GetKmsInstanceResponseBodyKmsInstance) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceResponseBodyKmsInstance) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetBindVpcs() *GetKmsInstanceResponseBodyKmsInstanceBindVpcs {
	return s.BindVpcs
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetCaCertificateChainPem() *string {
	return s.CaCertificateChainPem
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetEndDate() *string {
	return s.EndDate
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetKeyNum() *int64 {
	return s.KeyNum
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetLog() *int64 {
	return s.Log
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetLogStorage() *int64 {
	return s.LogStorage
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetProductType() *string {
	return s.ProductType
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetSaleStatus() *string {
	return s.SaleStatus
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetSecretNum() *string {
	return s.SecretNum
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetSpec() *int64 {
	return s.Spec
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetStartDate() *string {
	return s.StartDate
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetStatus() *string {
	return s.Status
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetVpcNum() *int64 {
	return s.VpcNum
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *GetKmsInstanceResponseBodyKmsInstance) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetBindVpcs(v *GetKmsInstanceResponseBodyKmsInstanceBindVpcs) *GetKmsInstanceResponseBodyKmsInstance {
	s.BindVpcs = v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetCaCertificateChainPem(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.CaCertificateChainPem = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetChargeType(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.ChargeType = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetCreateTime(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.CreateTime = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetEndDate(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.EndDate = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetInstanceId(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.InstanceId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetInstanceName(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.InstanceName = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetKeyNum(v int64) *GetKmsInstanceResponseBodyKmsInstance {
	s.KeyNum = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetLog(v int64) *GetKmsInstanceResponseBodyKmsInstance {
	s.Log = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetLogStorage(v int64) *GetKmsInstanceResponseBodyKmsInstance {
	s.LogStorage = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetProductType(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.ProductType = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetProductVersion(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.ProductVersion = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetSaleStatus(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.SaleStatus = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetSecretNum(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.SecretNum = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetSpec(v int64) *GetKmsInstanceResponseBodyKmsInstance {
	s.Spec = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetStartDate(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.StartDate = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetStatus(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.Status = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetVpcId(v string) *GetKmsInstanceResponseBodyKmsInstance {
	s.VpcId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetVpcNum(v int64) *GetKmsInstanceResponseBodyKmsInstance {
	s.VpcNum = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetVswitchIds(v []*string) *GetKmsInstanceResponseBodyKmsInstance {
	s.VswitchIds = v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) SetZoneIds(v []*string) *GetKmsInstanceResponseBodyKmsInstance {
	s.ZoneIds = v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstance) Validate() error {
	if s.BindVpcs != nil {
		if err := s.BindVpcs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKmsInstanceResponseBodyKmsInstanceBindVpcs struct {
	BindVpc []*GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc `json:"BindVpc,omitempty" xml:"BindVpc,omitempty" type:"Repeated"`
}

func (s GetKmsInstanceResponseBodyKmsInstanceBindVpcs) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceResponseBodyKmsInstanceBindVpcs) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcs) GetBindVpc() []*GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	return s.BindVpc
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcs) SetBindVpc(v []*GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) *GetKmsInstanceResponseBodyKmsInstanceBindVpcs {
	s.BindVpc = v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcs) Validate() error {
	if s.BindVpc != nil {
		for _, item := range s.BindVpc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc struct {
	// The region to which the VPC belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch in the VPC.
	//
	// example:
	//
	// vsw-bp1i512amhdje10f1****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp19z7djuhtad5dff****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 190325303126****
	VpcOwnerId *string `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) GetRegionId() *string {
	return s.RegionId
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) GetVpcOwnerId() *string {
	return s.VpcOwnerId
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) SetRegionId(v string) *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	s.RegionId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) SetVSwitchId(v string) *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	s.VSwitchId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) SetVpcId(v string) *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	s.VpcId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) SetVpcOwnerId(v string) *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc {
	s.VpcOwnerId = &v
	return s
}

func (s *GetKmsInstanceResponseBodyKmsInstanceBindVpcsBindVpc) Validate() error {
	return dara.Validate(s)
}
