// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeProjectInfoResponseBody
	GetRequestId() *string
	SetResult(v *DescribeProjectInfoResponseBodyResult) *DescribeProjectInfoResponseBody
	GetResult() *DescribeProjectInfoResponseBodyResult
	SetSuccess(v bool) *DescribeProjectInfoResponseBody
	GetSuccess() *bool
}

type DescribeProjectInfoResponseBody struct {
	// example:
	//
	// ee3e1b3b-6c38-4bcf-be40-5a946cfda761
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeProjectInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectInfoResponseBody) GetResult() *DescribeProjectInfoResponseBodyResult {
	return s.Result
}

func (s *DescribeProjectInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProjectInfoResponseBody) SetRequestId(v string) *DescribeProjectInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectInfoResponseBody) SetResult(v *DescribeProjectInfoResponseBodyResult) *DescribeProjectInfoResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectInfoResponseBody) SetSuccess(v bool) *DescribeProjectInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProjectInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProjectInfoResponseBodyResult struct {
	// example:
	//
	// 3
	CurrentStepNo *int32 `json:"CurrentStepNo,omitempty" xml:"CurrentStepNo,omitempty"`
	// example:
	//
	// 27291111****
	CustomerAliUid *int64 `json:"CustomerAliUid,omitempty" xml:"CustomerAliUid,omitempty"`
	// example:
	//
	// 4
	FinalStepNo *int32 `json:"FinalStepNo,omitempty" xml:"FinalStepNo,omitempty"`
	// example:
	//
	// null
	FinishType *string `json:"FinishType,omitempty" xml:"FinishType,omitempty"`
	// example:
	//
	// 1588834324000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1620403200000
	GmtExpired *int64 `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// example:
	//
	// 1620403200000
	GmtFinished *int64 `json:"GmtFinished,omitempty" xml:"GmtFinished,omitempty"`
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2059111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// cmgj***055
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// yuncode****500001
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	ProductSkuName *string `json:"ProductSkuName,omitempty" xml:"ProductSkuName,omitempty"`
	// example:
	//
	// Starting
	ProjectStatus *string `json:"ProjectStatus,omitempty" xml:"ProjectStatus,omitempty"`
	// example:
	//
	// 45121111****
	SupplierAliUid *int64 `json:"SupplierAliUid,omitempty" xml:"SupplierAliUid,omitempty"`
	// example:
	//
	// 410
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// Public
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeProjectInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectInfoResponseBodyResult) GetCurrentStepNo() *int32 {
	return s.CurrentStepNo
}

func (s *DescribeProjectInfoResponseBodyResult) GetCustomerAliUid() *int64 {
	return s.CustomerAliUid
}

func (s *DescribeProjectInfoResponseBodyResult) GetFinalStepNo() *int32 {
	return s.FinalStepNo
}

func (s *DescribeProjectInfoResponseBodyResult) GetFinishType() *string {
	return s.FinishType
}

func (s *DescribeProjectInfoResponseBodyResult) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeProjectInfoResponseBodyResult) GetGmtExpired() *int64 {
	return s.GmtExpired
}

func (s *DescribeProjectInfoResponseBodyResult) GetGmtFinished() *int64 {
	return s.GmtFinished
}

func (s *DescribeProjectInfoResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeProjectInfoResponseBodyResult) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeProjectInfoResponseBodyResult) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeProjectInfoResponseBodyResult) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeProjectInfoResponseBodyResult) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribeProjectInfoResponseBodyResult) GetProductSkuName() *string {
	return s.ProductSkuName
}

func (s *DescribeProjectInfoResponseBodyResult) GetProjectStatus() *string {
	return s.ProjectStatus
}

func (s *DescribeProjectInfoResponseBodyResult) GetSupplierAliUid() *int64 {
	return s.SupplierAliUid
}

func (s *DescribeProjectInfoResponseBodyResult) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeProjectInfoResponseBodyResult) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeProjectInfoResponseBodyResult) SetCurrentStepNo(v int32) *DescribeProjectInfoResponseBodyResult {
	s.CurrentStepNo = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetCustomerAliUid(v int64) *DescribeProjectInfoResponseBodyResult {
	s.CustomerAliUid = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetFinalStepNo(v int32) *DescribeProjectInfoResponseBodyResult {
	s.FinalStepNo = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetFinishType(v string) *DescribeProjectInfoResponseBodyResult {
	s.FinishType = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetGmtCreate(v int64) *DescribeProjectInfoResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetGmtExpired(v int64) *DescribeProjectInfoResponseBodyResult {
	s.GmtExpired = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetGmtFinished(v int64) *DescribeProjectInfoResponseBodyResult {
	s.GmtFinished = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetInstanceId(v string) *DescribeProjectInfoResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetOrderId(v int64) *DescribeProjectInfoResponseBodyResult {
	s.OrderId = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetProductCode(v string) *DescribeProjectInfoResponseBodyResult {
	s.ProductCode = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetProductName(v string) *DescribeProjectInfoResponseBodyResult {
	s.ProductName = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetProductSkuCode(v string) *DescribeProjectInfoResponseBodyResult {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetProductSkuName(v string) *DescribeProjectInfoResponseBodyResult {
	s.ProductSkuName = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetProjectStatus(v string) *DescribeProjectInfoResponseBodyResult {
	s.ProjectStatus = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetSupplierAliUid(v int64) *DescribeProjectInfoResponseBodyResult {
	s.SupplierAliUid = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetTemplateId(v int64) *DescribeProjectInfoResponseBodyResult {
	s.TemplateId = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) SetTemplateType(v string) *DescribeProjectInfoResponseBodyResult {
	s.TemplateType = &v
	return s
}

func (s *DescribeProjectInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
