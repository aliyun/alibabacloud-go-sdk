// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListApmResponseBodyHeaders) *ListApmResponseBody
	GetHeaders() *ListApmResponseBodyHeaders
	SetRequestId(v string) *ListApmResponseBody
	GetRequestId() *string
	SetResult(v []*ListApmResponseBodyResult) *ListApmResponseBody
	GetResult() []*ListApmResponseBodyResult
}

type ListApmResponseBody struct {
	Headers *ListApmResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// example:
	//
	// 080D3399-76CF-519D-A540-2C44BC056EB7
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListApmResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListApmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApmResponseBody) GoString() string {
	return s.String()
}

func (s *ListApmResponseBody) GetHeaders() *ListApmResponseBodyHeaders {
	return s.Headers
}

func (s *ListApmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApmResponseBody) GetResult() []*ListApmResponseBodyResult {
	return s.Result
}

func (s *ListApmResponseBody) SetHeaders(v *ListApmResponseBodyHeaders) *ListApmResponseBody {
	s.Headers = v
	return s
}

func (s *ListApmResponseBody) SetRequestId(v string) *ListApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApmResponseBody) SetResult(v []*ListApmResponseBodyResult) *ListApmResponseBody {
	s.Result = v
	return s
}

func (s *ListApmResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApmResponseBodyHeaders struct {
	// example:
	//
	// 1
	XTotalCount *int64 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListApmResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListApmResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListApmResponseBodyHeaders) GetXTotalCount() *int64 {
	return s.XTotalCount
}

func (s *ListApmResponseBodyHeaders) SetXTotalCount(v int64) *ListApmResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListApmResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListApmResponseBodyResult struct {
	// example:
	//
	// 2021-11-16T07:15:51.967Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 1
	DeployedReplica *int64 `json:"deployedReplica,omitempty" xml:"deployedReplica,omitempty"`
	// example:
	//
	// APMtest
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// apm-cn-i7m2fuae****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1
	NodeAmount *int64 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// example:
	//
	// es-cn-i7m2fsfhc001x****
	OutputES *string `json:"outputES,omitempty" xml:"outputES,omitempty"`
	// example:
	//
	// elastic
	OutputESUserName *string `json:"outputESUserName,omitempty" xml:"outputESUserName,omitempty"`
	// example:
	//
	// 133071096032****
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 1
	Replica *int64 `json:"replica,omitempty" xml:"replica,omitempty"`
	// example:
	//
	// C1M2
	ResourceSpec *string `json:"resourceSpec,omitempty" xml:"resourceSpec,omitempty"`
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 7.10.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// vpc-bp1530vdhqkamm9s0****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// example:
	//
	// vsw-bp1j1mql6r9g5vfb4****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s ListApmResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListApmResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListApmResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListApmResponseBodyResult) GetDeployedReplica() *int64 {
	return s.DeployedReplica
}

func (s *ListApmResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListApmResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApmResponseBodyResult) GetNodeAmount() *int64 {
	return s.NodeAmount
}

func (s *ListApmResponseBodyResult) GetOutputES() *string {
	return s.OutputES
}

func (s *ListApmResponseBodyResult) GetOutputESUserName() *string {
	return s.OutputESUserName
}

func (s *ListApmResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListApmResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListApmResponseBodyResult) GetRegion() *string {
	return s.Region
}

func (s *ListApmResponseBodyResult) GetReplica() *int64 {
	return s.Replica
}

func (s *ListApmResponseBodyResult) GetResourceSpec() *string {
	return s.ResourceSpec
}

func (s *ListApmResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListApmResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *ListApmResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *ListApmResponseBodyResult) GetVsArea() *string {
	return s.VsArea
}

func (s *ListApmResponseBodyResult) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListApmResponseBodyResult) SetCreatedAt(v string) *ListApmResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListApmResponseBodyResult) SetDeployedReplica(v int64) *ListApmResponseBodyResult {
	s.DeployedReplica = &v
	return s
}

func (s *ListApmResponseBodyResult) SetDescription(v string) *ListApmResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListApmResponseBodyResult) SetInstanceId(v string) *ListApmResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListApmResponseBodyResult) SetNodeAmount(v int64) *ListApmResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *ListApmResponseBodyResult) SetOutputES(v string) *ListApmResponseBodyResult {
	s.OutputES = &v
	return s
}

func (s *ListApmResponseBodyResult) SetOutputESUserName(v string) *ListApmResponseBodyResult {
	s.OutputESUserName = &v
	return s
}

func (s *ListApmResponseBodyResult) SetOwnerId(v string) *ListApmResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListApmResponseBodyResult) SetPaymentType(v string) *ListApmResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *ListApmResponseBodyResult) SetRegion(v string) *ListApmResponseBodyResult {
	s.Region = &v
	return s
}

func (s *ListApmResponseBodyResult) SetReplica(v int64) *ListApmResponseBodyResult {
	s.Replica = &v
	return s
}

func (s *ListApmResponseBodyResult) SetResourceSpec(v string) *ListApmResponseBodyResult {
	s.ResourceSpec = &v
	return s
}

func (s *ListApmResponseBodyResult) SetStatus(v string) *ListApmResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListApmResponseBodyResult) SetVersion(v string) *ListApmResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListApmResponseBodyResult) SetVpcId(v string) *ListApmResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *ListApmResponseBodyResult) SetVsArea(v string) *ListApmResponseBodyResult {
	s.VsArea = &v
	return s
}

func (s *ListApmResponseBodyResult) SetVswitchId(v string) *ListApmResponseBodyResult {
	s.VswitchId = &v
	return s
}

func (s *ListApmResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
