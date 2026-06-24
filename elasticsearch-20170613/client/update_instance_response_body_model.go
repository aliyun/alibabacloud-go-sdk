// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceResponseBody
	GetRequestId() *string
	SetResult(v *UpdateInstanceResponseBodyResult) *UpdateInstanceResponseBody
	GetResult() *UpdateInstanceResponseBodyResult
}

type UpdateInstanceResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *UpdateInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResponseBody) GetResult() *UpdateInstanceResponseBodyResult {
	return s.Result
}

func (s *UpdateInstanceResponseBody) SetCode(v string) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetResult(v *UpdateInstanceResponseBodyResult) *UpdateInstanceResponseBody {
	s.Result = v
	return s
}

func (s *UpdateInstanceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceResponseBodyResult struct {
	// The time when the instance was created.
	//
	// example:
	//
	// 2018-07-13T03:58:07.253Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 实例的私网访问域名。
	//
	// example:
	//
	// es-cn-abc.elasticsearch.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The instance version.
	//
	// example:
	//
	// 5.5.3_with_X-Pack
	EsVersion *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// es-cn-abc
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The Kibana node configuration.
	KibanaConfiguration *UpdateInstanceResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	// The master node configuration.
	MasterConfiguration *UpdateInstanceResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	// The number of data nodes.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// The data node configuration.
	NodeSpec *UpdateInstanceResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// - prepaid: subscription
	//
	// - postpaid: pay-as-you-go.
	//
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The status of the instance. Valid values:
	//
	// - active: Normal
	//
	// - activating: Taking effect
	//
	// - inactive: Frozen
	//
	// - invalid: Invalid.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateInstanceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *UpdateInstanceResponseBodyResult) GetEsVersion() *string {
	return s.EsVersion
}

func (s *UpdateInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceResponseBodyResult) GetKibanaConfiguration() *UpdateInstanceResponseBodyResultKibanaConfiguration {
	return s.KibanaConfiguration
}

func (s *UpdateInstanceResponseBodyResult) GetMasterConfiguration() *UpdateInstanceResponseBodyResultMasterConfiguration {
	return s.MasterConfiguration
}

func (s *UpdateInstanceResponseBodyResult) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *UpdateInstanceResponseBodyResult) GetNodeSpec() *UpdateInstanceResponseBodyResultNodeSpec {
	return s.NodeSpec
}

func (s *UpdateInstanceResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *UpdateInstanceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *UpdateInstanceResponseBodyResult) SetCreatedAt(v string) *UpdateInstanceResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetDescription(v string) *UpdateInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetDomain(v string) *UpdateInstanceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetEsVersion(v string) *UpdateInstanceResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetInstanceId(v string) *UpdateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetKibanaConfiguration(v *UpdateInstanceResponseBodyResultKibanaConfiguration) *UpdateInstanceResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetMasterConfiguration(v *UpdateInstanceResponseBodyResultMasterConfiguration) *UpdateInstanceResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetNodeAmount(v int32) *UpdateInstanceResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetNodeSpec(v *UpdateInstanceResponseBodyResultNodeSpec) *UpdateInstanceResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetPaymentType(v string) *UpdateInstanceResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) SetStatus(v string) *UpdateInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateInstanceResponseBodyResult) Validate() error {
	if s.KibanaConfiguration != nil {
		if err := s.KibanaConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.MasterConfiguration != nil {
		if err := s.MasterConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSpec != nil {
		if err := s.NodeSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceResponseBodyResultKibanaConfiguration struct {
	// The number of nodes.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The storage size of the node.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node. You can ignore this parameter.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// 节点规格。
	//
	// example:
	//
	// elasticsearch.n4.small
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdateInstanceResponseBodyResultKibanaConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) SetAmount(v int32) *UpdateInstanceResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) SetDisk(v int32) *UpdateInstanceResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) SetDiskType(v string) *UpdateInstanceResponseBodyResultKibanaConfiguration {
	s.DiskType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) SetSpec(v string) *UpdateInstanceResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultKibanaConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceResponseBodyResultMasterConfiguration struct {
	// The number of nodes.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The storage size of the node. Unit: GB.
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node. Only cloud_ssd (standard SSD) is supported.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// 节点规格。
	//
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdateInstanceResponseBodyResultMasterConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) SetAmount(v int32) *UpdateInstanceResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) SetDisk(v int32) *UpdateInstanceResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) SetDiskType(v string) *UpdateInstanceResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) SetSpec(v string) *UpdateInstanceResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultMasterConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceResponseBodyResultNodeSpec struct {
	// The storage size of the node. Unit: GB.
	//
	// example:
	//
	// 40
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The storage type of the node. Valid values:
	//
	// - cloud_ssd: standard SSD
	//
	// - cloud_efficiency: ultra disk.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// 节点规格。
	//
	// example:
	//
	// elasticsearch.sn2ne.xlarge
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdateInstanceResponseBodyResultNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) SetDisk(v int32) *UpdateInstanceResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) SetDiskType(v string) *UpdateInstanceResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) SetSpec(v string) *UpdateInstanceResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

func (s *UpdateInstanceResponseBodyResultNodeSpec) Validate() error {
	return dara.Validate(s)
}
