// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeApmResponseBody
	GetRequestId() *string
	SetResult(v *DescribeApmResponseBodyResult) *DescribeApmResponseBody
	GetResult() *DescribeApmResponseBodyResult
}

type DescribeApmResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 97597D87-8066-5493-B692-5C50DA236D68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result.
	Result *DescribeApmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeApmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApmResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApmResponseBody) GetResult() *DescribeApmResponseBodyResult {
	return s.Result
}

func (s *DescribeApmResponseBody) SetRequestId(v string) *DescribeApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApmResponseBody) SetResult(v *DescribeApmResponseBodyResult) *DescribeApmResponseBody {
	s.Result = v
	return s
}

func (s *DescribeApmResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApmResponseBodyResult struct {
	// Instance domain.
	//
	// example:
	//
	// apm-cn-i7m2fuae****.apm.elasticsearch.aliyuncs.com:8200
	ApmServerDomain *string `json:"apmServerDomain,omitempty" xml:"apmServerDomain,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2021-11-16T07:15:51.967Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Deployed replica count.
	//
	// example:
	//
	// 1
	DeployedReplica *int64 `json:"deployedReplica,omitempty" xml:"deployedReplica,omitempty"`
	// Instance name.
	//
	// example:
	//
	// APMtest
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Instance expiration time.
	//
	// example:
	//
	// 4792752000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// apm-cn-i7m2fuae****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Number of nodes.
	//
	// example:
	//
	// 1
	NodeAmount *int64 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// Associated Elasticsearch instance.
	//
	// example:
	//
	// es-cn-i7m2fsfhc001x****
	OutputES *string `json:"outputES,omitempty" xml:"outputES,omitempty"`
	// Username of the associated Elasticsearch instance.
	//
	// example:
	//
	// elastic
	OutputESUserName *string `json:"outputESUserName,omitempty" xml:"outputESUserName,omitempty"`
	// User account ID.
	//
	// example:
	//
	// 133071096032****
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// Payment method, with the following values:
	//
	// - postpaid: Pay-as-you-go.
	//
	// - prepaid: Subscription.
	//
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// Region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Replica count.
	//
	// example:
	//
	// 1
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// Specification, with the following values:
	//
	// - C1M2 (1 core, 2 GB)
	//
	// - C2M4 (2 cores, 4 GB)
	//
	// example:
	//
	// C1M2
	ResourceSpec *string `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	// Lifecycle status, with the following values:
	//
	// - CREATING: Creating.
	//
	// - ACTIVATING: Activating.
	//
	// - ACTIVE: Active.
	//
	// - INACTIVE: Frozen.
	//
	// - INVALID: Invalid.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Version information.
	//
	// example:
	//
	// 7.10.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// VPC ID.
	//
	// example:
	//
	// vpc-bp1530vdhqkamm9s0****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// Availability zone where the switch is located.
	//
	// example:
	//
	// cn-hangzhou-i
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// Switch ID.
	//
	// example:
	//
	// vsw-bp1j1mql6r9g5vfb4****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s DescribeApmResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeApmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeApmResponseBodyResult) GetApmServerDomain() *string {
	return s.ApmServerDomain
}

func (s *DescribeApmResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeApmResponseBodyResult) GetDeployedReplica() *int64 {
	return s.DeployedReplica
}

func (s *DescribeApmResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeApmResponseBodyResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeApmResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApmResponseBodyResult) GetNodeAmount() *int64 {
	return s.NodeAmount
}

func (s *DescribeApmResponseBodyResult) GetOutputES() *string {
	return s.OutputES
}

func (s *DescribeApmResponseBodyResult) GetOutputESUserName() *string {
	return s.OutputESUserName
}

func (s *DescribeApmResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeApmResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeApmResponseBodyResult) GetRegion() *string {
	return s.Region
}

func (s *DescribeApmResponseBodyResult) GetReplica() *int32 {
	return s.Replica
}

func (s *DescribeApmResponseBodyResult) GetResourceSpec() *string {
	return s.ResourceSpec
}

func (s *DescribeApmResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeApmResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *DescribeApmResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeApmResponseBodyResult) GetVsArea() *string {
	return s.VsArea
}

func (s *DescribeApmResponseBodyResult) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeApmResponseBodyResult) SetApmServerDomain(v string) *DescribeApmResponseBodyResult {
	s.ApmServerDomain = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetCreatedAt(v string) *DescribeApmResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetDeployedReplica(v int64) *DescribeApmResponseBodyResult {
	s.DeployedReplica = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetDescription(v string) *DescribeApmResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetEndTime(v int64) *DescribeApmResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetInstanceId(v string) *DescribeApmResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetNodeAmount(v int64) *DescribeApmResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetOutputES(v string) *DescribeApmResponseBodyResult {
	s.OutputES = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetOutputESUserName(v string) *DescribeApmResponseBodyResult {
	s.OutputESUserName = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetOwnerId(v string) *DescribeApmResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetPaymentType(v string) *DescribeApmResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetRegion(v string) *DescribeApmResponseBodyResult {
	s.Region = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetReplica(v int32) *DescribeApmResponseBodyResult {
	s.Replica = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetResourceSpec(v string) *DescribeApmResponseBodyResult {
	s.ResourceSpec = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetStatus(v string) *DescribeApmResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetVersion(v string) *DescribeApmResponseBodyResult {
	s.Version = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetVpcId(v string) *DescribeApmResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetVsArea(v string) *DescribeApmResponseBodyResult {
	s.VsArea = &v
	return s
}

func (s *DescribeApmResponseBodyResult) SetVswitchId(v string) *DescribeApmResponseBodyResult {
	s.VswitchId = &v
	return s
}

func (s *DescribeApmResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
