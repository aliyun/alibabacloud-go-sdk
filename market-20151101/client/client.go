// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateLicenseRequest struct {
	Identification *string `json:"Identification,omitempty" xml:"Identification,omitempty"`
	LicenseCode    *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
}

func (s ActivateLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActivateLicenseRequest) SetIdentification(v string) *ActivateLicenseRequest {
	s.Identification = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicenseCode(v string) *ActivateLicenseRequest {
	s.LicenseCode = &v
	return s
}

type ActivateLicenseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBody) SetRequestId(v string) *ActivateLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetSuccess(v bool) *ActivateLicenseResponseBody {
	s.Success = &v
	return s
}

type ActivateLicenseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActivateLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivateLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponse) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponse) SetHeaders(v map[string]*string) *ActivateLicenseResponse {
	s.Headers = v
	return s
}

func (s *ActivateLicenseResponse) SetStatusCode(v int32) *ActivateLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateLicenseResponse) SetBody(v *ActivateLicenseResponseBody) *ActivateLicenseResponse {
	s.Body = v
	return s
}

type AutoRenewInstanceRequest struct {
	AutoRenewCycle    *string `json:"AutoRenewCycle,omitempty" xml:"AutoRenewCycle,omitempty"`
	AutoRenewDuration *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	OrderBizId        *int64  `json:"OrderBizId,omitempty" xml:"OrderBizId,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AutoRenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AutoRenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *AutoRenewInstanceRequest) SetAutoRenewCycle(v string) *AutoRenewInstanceRequest {
	s.AutoRenewCycle = &v
	return s
}

func (s *AutoRenewInstanceRequest) SetAutoRenewDuration(v int32) *AutoRenewInstanceRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *AutoRenewInstanceRequest) SetOrderBizId(v int64) *AutoRenewInstanceRequest {
	s.OrderBizId = &v
	return s
}

func (s *AutoRenewInstanceRequest) SetOwnerId(v int64) *AutoRenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *AutoRenewInstanceRequest) SetType(v string) *AutoRenewInstanceRequest {
	s.Type = &v
	return s
}

type AutoRenewInstanceResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AutoRenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AutoRenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AutoRenewInstanceResponseBody) SetData(v bool) *AutoRenewInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *AutoRenewInstanceResponseBody) SetRequestId(v string) *AutoRenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AutoRenewInstanceResponseBody) SetSuccess(v bool) *AutoRenewInstanceResponseBody {
	s.Success = &v
	return s
}

type AutoRenewInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AutoRenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AutoRenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AutoRenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *AutoRenewInstanceResponse) SetHeaders(v map[string]*string) *AutoRenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *AutoRenewInstanceResponse) SetStatusCode(v int32) *AutoRenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AutoRenewInstanceResponse) SetBody(v *AutoRenewInstanceResponseBody) *AutoRenewInstanceResponse {
	s.Body = v
	return s
}

type CreateOrderRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Commodity   *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	OrderSouce  *string `json:"OrderSouce,omitempty" xml:"OrderSouce,omitempty"`
	OrderType   *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetClientToken(v string) *CreateOrderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOrderRequest) SetCommodity(v string) *CreateOrderRequest {
	s.Commodity = &v
	return s
}

func (s *CreateOrderRequest) SetOrderSouce(v string) *CreateOrderRequest {
	s.OrderSouce = &v
	return s
}

func (s *CreateOrderRequest) SetOrderType(v string) *CreateOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateOrderRequest) SetOwnerId(v string) *CreateOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOrderRequest) SetPaymentType(v string) *CreateOrderRequest {
	s.PaymentType = &v
	return s
}

type CreateOrderResponseBody struct {
	InstanceIds *CreateOrderResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	OrderId     *string                             `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetInstanceIds(v *CreateOrderResponseBodyInstanceIds) *CreateOrderResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateOrderResponseBody) SetOrderId(v string) *CreateOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrderResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s CreateOrderResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyInstanceIds) SetInstanceId(v []*string) *CreateOrderResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type CreateOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetStatusCode(v int32) *CreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type CrossAccountVerifyTokenRequest struct {
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CrossAccountVerifyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CrossAccountVerifyTokenRequest) GoString() string {
	return s.String()
}

func (s *CrossAccountVerifyTokenRequest) SetToken(v string) *CrossAccountVerifyTokenRequest {
	s.Token = &v
	return s
}

type CrossAccountVerifyTokenResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CrossAccountVerifyTokenResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CrossAccountVerifyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CrossAccountVerifyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CrossAccountVerifyTokenResponseBody) SetCode(v string) *CrossAccountVerifyTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) SetMessage(v string) *CrossAccountVerifyTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) SetRequestId(v string) *CrossAccountVerifyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) SetResult(v *CrossAccountVerifyTokenResponseBodyResult) *CrossAccountVerifyTokenResponseBody {
	s.Result = v
	return s
}

func (s *CrossAccountVerifyTokenResponseBody) SetSuccess(v bool) *CrossAccountVerifyTokenResponseBody {
	s.Success = &v
	return s
}

type CrossAccountVerifyTokenResponseBodyResult struct {
	AuthRoles []*string `json:"AuthRoles,omitempty" xml:"AuthRoles,omitempty" type:"Repeated"`
	AuthTime  *int64    `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
	Name      *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Uid       *string   `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CrossAccountVerifyTokenResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CrossAccountVerifyTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CrossAccountVerifyTokenResponseBodyResult) SetAuthRoles(v []*string) *CrossAccountVerifyTokenResponseBodyResult {
	s.AuthRoles = v
	return s
}

func (s *CrossAccountVerifyTokenResponseBodyResult) SetAuthTime(v int64) *CrossAccountVerifyTokenResponseBodyResult {
	s.AuthTime = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBodyResult) SetName(v string) *CrossAccountVerifyTokenResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CrossAccountVerifyTokenResponseBodyResult) SetUid(v string) *CrossAccountVerifyTokenResponseBodyResult {
	s.Uid = &v
	return s
}

type CrossAccountVerifyTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CrossAccountVerifyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CrossAccountVerifyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CrossAccountVerifyTokenResponse) GoString() string {
	return s.String()
}

func (s *CrossAccountVerifyTokenResponse) SetHeaders(v map[string]*string) *CrossAccountVerifyTokenResponse {
	s.Headers = v
	return s
}

func (s *CrossAccountVerifyTokenResponse) SetStatusCode(v int32) *CrossAccountVerifyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CrossAccountVerifyTokenResponse) SetBody(v *CrossAccountVerifyTokenResponseBody) *CrossAccountVerifyTokenResponse {
	s.Body = v
	return s
}

type DescribeApiMeteringRequest struct {
	PageNum     *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	Type        *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeApiMeteringRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMeteringRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiMeteringRequest) SetPageNum(v int32) *DescribeApiMeteringRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeApiMeteringRequest) SetProductCode(v string) *DescribeApiMeteringRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeApiMeteringRequest) SetType(v int32) *DescribeApiMeteringRequest {
	s.Type = &v
	return s
}

type DescribeApiMeteringResponseBody struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	// fatal
	Fatal      *bool                                    `json:"Fatal,omitempty" xml:"Fatal,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int64                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*DescribeApiMeteringResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success    *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Version    *string                                  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeApiMeteringResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMeteringResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiMeteringResponseBody) SetCode(v string) *DescribeApiMeteringResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetCount(v int64) *DescribeApiMeteringResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetFatal(v bool) *DescribeApiMeteringResponseBody {
	s.Fatal = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetMessage(v string) *DescribeApiMeteringResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetPageNumber(v int64) *DescribeApiMeteringResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetPageSize(v int64) *DescribeApiMeteringResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetRequestId(v string) *DescribeApiMeteringResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetResult(v []*DescribeApiMeteringResponseBodyResult) *DescribeApiMeteringResponseBody {
	s.Result = v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetSuccess(v bool) *DescribeApiMeteringResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetVersion(v string) *DescribeApiMeteringResponseBody {
	s.Version = &v
	return s
}

type DescribeApiMeteringResponseBodyResult struct {
	AliyunPk    *int64  `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	TotalQuota  *int64  `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	TotalUsage  *int64  `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	Unit        *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeApiMeteringResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMeteringResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeApiMeteringResponseBodyResult) SetAliyunPk(v int64) *DescribeApiMeteringResponseBodyResult {
	s.AliyunPk = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetProductCode(v string) *DescribeApiMeteringResponseBodyResult {
	s.ProductCode = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetProductName(v string) *DescribeApiMeteringResponseBodyResult {
	s.ProductName = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetTotalQuota(v int64) *DescribeApiMeteringResponseBodyResult {
	s.TotalQuota = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetTotalUsage(v int64) *DescribeApiMeteringResponseBodyResult {
	s.TotalUsage = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetUnit(v string) *DescribeApiMeteringResponseBodyResult {
	s.Unit = &v
	return s
}

type DescribeApiMeteringResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiMeteringResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiMeteringResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMeteringResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiMeteringResponse) SetHeaders(v map[string]*string) *DescribeApiMeteringResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiMeteringResponse) SetStatusCode(v int32) *DescribeApiMeteringResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiMeteringResponse) SetBody(v *DescribeApiMeteringResponseBody) *DescribeApiMeteringResponse {
	s.Body = v
	return s
}

type DescribeCurrentNodeInfoRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCurrentNodeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCurrentNodeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCurrentNodeInfoRequest) SetInstanceId(v string) *DescribeCurrentNodeInfoRequest {
	s.InstanceId = &v
	return s
}

type DescribeCurrentNodeInfoResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeCurrentNodeInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCurrentNodeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCurrentNodeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCurrentNodeInfoResponseBody) SetRequestId(v string) *DescribeCurrentNodeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBody) SetResult(v *DescribeCurrentNodeInfoResponseBodyResult) *DescribeCurrentNodeInfoResponseBody {
	s.Result = v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBody) SetSuccess(v bool) *DescribeCurrentNodeInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeCurrentNodeInfoResponseBodyResult struct {
	AllowRollbackNode *bool   `json:"AllowRollbackNode,omitempty" xml:"AllowRollbackNode,omitempty"`
	AutoFinishNode    *bool   `json:"AutoFinishNode,omitempty" xml:"AutoFinishNode,omitempty"`
	FinalStepNo       *int32  `json:"FinalStepNo,omitempty" xml:"FinalStepNo,omitempty"`
	GmtExpired        *int64  `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	GmtFinished       *int64  `json:"GmtFinished,omitempty" xml:"GmtFinished,omitempty"`
	GmtStart          *int64  `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	NeedAttachment    *bool   `json:"NeedAttachment,omitempty" xml:"NeedAttachment,omitempty"`
	NextNodeId        *int64  `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	NodeId            *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName          *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeStatus        *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	OperatorRole      *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	ParentNodeId      *int64  `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	PreviousNodeId    *int64  `json:"PreviousNodeId,omitempty" xml:"PreviousNodeId,omitempty"`
	StepNo            *int32  `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
	TemplateForm      *string `json:"TemplateForm,omitempty" xml:"TemplateForm,omitempty"`
}

func (s DescribeCurrentNodeInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeCurrentNodeInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetAllowRollbackNode(v bool) *DescribeCurrentNodeInfoResponseBodyResult {
	s.AllowRollbackNode = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetAutoFinishNode(v bool) *DescribeCurrentNodeInfoResponseBodyResult {
	s.AutoFinishNode = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetFinalStepNo(v int32) *DescribeCurrentNodeInfoResponseBodyResult {
	s.FinalStepNo = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetGmtExpired(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.GmtExpired = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetGmtFinished(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.GmtFinished = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetGmtStart(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.GmtStart = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNeedAttachment(v bool) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NeedAttachment = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNextNodeId(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NextNodeId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNodeId(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NodeId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNodeName(v string) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NodeName = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetNodeStatus(v string) *DescribeCurrentNodeInfoResponseBodyResult {
	s.NodeStatus = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetOperatorRole(v string) *DescribeCurrentNodeInfoResponseBodyResult {
	s.OperatorRole = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetParentNodeId(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.ParentNodeId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetPreviousNodeId(v int64) *DescribeCurrentNodeInfoResponseBodyResult {
	s.PreviousNodeId = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetStepNo(v int32) *DescribeCurrentNodeInfoResponseBodyResult {
	s.StepNo = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponseBodyResult) SetTemplateForm(v string) *DescribeCurrentNodeInfoResponseBodyResult {
	s.TemplateForm = &v
	return s
}

type DescribeCurrentNodeInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCurrentNodeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCurrentNodeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCurrentNodeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCurrentNodeInfoResponse) SetHeaders(v map[string]*string) *DescribeCurrentNodeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCurrentNodeInfoResponse) SetStatusCode(v int32) *DescribeCurrentNodeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponse) SetBody(v *DescribeCurrentNodeInfoResponseBody) *DescribeCurrentNodeInfoResponse {
	s.Body = v
	return s
}

type DescribeDistributionProductsRequest struct {
	Filter     []*DescribeDistributionProductsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	PageNumber *int64                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDistributionProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsRequest) SetFilter(v []*DescribeDistributionProductsRequestFilter) *DescribeDistributionProductsRequest {
	s.Filter = v
	return s
}

func (s *DescribeDistributionProductsRequest) SetPageNumber(v int64) *DescribeDistributionProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDistributionProductsRequest) SetPageSize(v int64) *DescribeDistributionProductsRequest {
	s.PageSize = &v
	return s
}

type DescribeDistributionProductsRequestFilter struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDistributionProductsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsRequestFilter) SetKey(v string) *DescribeDistributionProductsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeDistributionProductsRequestFilter) SetValue(v string) *DescribeDistributionProductsRequestFilter {
	s.Value = &v
	return s
}

type DescribeDistributionProductsResponseBody struct {
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results    []*DescribeDistributionProductsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDistributionProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsResponseBody) SetPageNumber(v int32) *DescribeDistributionProductsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDistributionProductsResponseBody) SetPageSize(v int32) *DescribeDistributionProductsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDistributionProductsResponseBody) SetRequestId(v string) *DescribeDistributionProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDistributionProductsResponseBody) SetResults(v []*DescribeDistributionProductsResponseBodyResults) *DescribeDistributionProductsResponseBody {
	s.Results = v
	return s
}

func (s *DescribeDistributionProductsResponseBody) SetTotalCount(v int32) *DescribeDistributionProductsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDistributionProductsResponseBodyResults struct {
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	FirstCategoryName  *string `json:"FirstCategoryName,omitempty" xml:"FirstCategoryName,omitempty"`
	ImageUrl           *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Price              *string `json:"Price,omitempty" xml:"Price,omitempty"`
	Score              *string `json:"Score,omitempty" xml:"Score,omitempty"`
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
	ShortDescription   *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	SubmissionRadio    *string `json:"SubmissionRadio,omitempty" xml:"SubmissionRadio,omitempty"`
	SupplierName       *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	SupplierUId        *string `json:"SupplierUId,omitempty" xml:"SupplierUId,omitempty"`
	TradeCount         *string `json:"TradeCount,omitempty" xml:"TradeCount,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserCommentCount   *string `json:"UserCommentCount,omitempty" xml:"UserCommentCount,omitempty"`
}

func (s DescribeDistributionProductsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsResponseBodyResults) SetCode(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetFirstCategoryName(v string) *DescribeDistributionProductsResponseBodyResults {
	s.FirstCategoryName = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetImageUrl(v string) *DescribeDistributionProductsResponseBodyResults {
	s.ImageUrl = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetName(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Name = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetPrice(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Price = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetScore(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Score = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetSecondCategoryName(v string) *DescribeDistributionProductsResponseBodyResults {
	s.SecondCategoryName = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetShortDescription(v string) *DescribeDistributionProductsResponseBodyResults {
	s.ShortDescription = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetSubmissionRadio(v string) *DescribeDistributionProductsResponseBodyResults {
	s.SubmissionRadio = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetSupplierName(v string) *DescribeDistributionProductsResponseBodyResults {
	s.SupplierName = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetSupplierUId(v string) *DescribeDistributionProductsResponseBodyResults {
	s.SupplierUId = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetTradeCount(v string) *DescribeDistributionProductsResponseBodyResults {
	s.TradeCount = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetType(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Type = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetUserCommentCount(v string) *DescribeDistributionProductsResponseBodyResults {
	s.UserCommentCount = &v
	return s
}

type DescribeDistributionProductsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDistributionProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDistributionProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsResponse) SetHeaders(v map[string]*string) *DescribeDistributionProductsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistributionProductsResponse) SetStatusCode(v int32) *DescribeDistributionProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistributionProductsResponse) SetBody(v *DescribeDistributionProductsResponseBody) *DescribeDistributionProductsResponse {
	s.Body = v
	return s
}

type DescribeDistributionProductsLinkRequest struct {
	Codes []*string `json:"Codes,omitempty" xml:"Codes,omitempty" type:"Repeated"`
}

func (s DescribeDistributionProductsLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsLinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkRequest) SetCodes(v []*string) *DescribeDistributionProductsLinkRequest {
	s.Codes = v
	return s
}

type DescribeDistributionProductsLinkShrinkRequest struct {
	CodesShrink *string `json:"Codes,omitempty" xml:"Codes,omitempty"`
}

func (s DescribeDistributionProductsLinkShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsLinkShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkShrinkRequest) SetCodesShrink(v string) *DescribeDistributionProductsLinkShrinkRequest {
	s.CodesShrink = &v
	return s
}

type DescribeDistributionProductsLinkResponseBody struct {
	// Id of the request
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*DescribeDistributionProductsLinkResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success    *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDistributionProductsLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkResponseBody) SetRequestId(v string) *DescribeDistributionProductsLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBody) SetResult(v []*DescribeDistributionProductsLinkResponseBodyResult) *DescribeDistributionProductsLinkResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBody) SetSuccess(v bool) *DescribeDistributionProductsLinkResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBody) SetTotalCount(v int64) *DescribeDistributionProductsLinkResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDistributionProductsLinkResponseBodyResult struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDistributionProductsLinkResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsLinkResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) SetCode(v string) *DescribeDistributionProductsLinkResponseBodyResult {
	s.Code = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) SetName(v string) *DescribeDistributionProductsLinkResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponseBodyResult) SetUrl(v string) *DescribeDistributionProductsLinkResponseBodyResult {
	s.Url = &v
	return s
}

type DescribeDistributionProductsLinkResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDistributionProductsLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDistributionProductsLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDistributionProductsLinkResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkResponse) SetHeaders(v map[string]*string) *DescribeDistributionProductsLinkResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistributionProductsLinkResponse) SetStatusCode(v int32) *DescribeDistributionProductsLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponse) SetBody(v *DescribeDistributionProductsLinkResponseBody) *DescribeDistributionProductsLinkResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderType  *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRequest) SetOrderType(v string) *DescribeInstanceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeInstanceRequest) SetOwnerId(v int64) *DescribeInstanceRequest {
	s.OwnerId = &v
	return s
}

type DescribeInstanceResponseBody struct {
	AppJson        *string                                     `json:"AppJson,omitempty" xml:"AppJson,omitempty"`
	AutoRenewal    *string                                     `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	BeganOn        *int64                                      `json:"BeganOn,omitempty" xml:"BeganOn,omitempty"`
	ComponentJson  *string                                     `json:"ComponentJson,omitempty" xml:"ComponentJson,omitempty"`
	Constraints    *string                                     `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	CreatedOn      *int64                                      `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	EndOn          *int64                                      `json:"EndOn,omitempty" xml:"EndOn,omitempty"`
	ExtendJson     *string                                     `json:"ExtendJson,omitempty" xml:"ExtendJson,omitempty"`
	HostJson       *string                                     `json:"HostJson,omitempty" xml:"HostJson,omitempty"`
	InstanceId     *int64                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsTrial        *bool                                       `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	Modules        *DescribeInstanceResponseBodyModules        `json:"Modules,omitempty" xml:"Modules,omitempty" type:"Struct"`
	OrderId        *int64                                      `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductCode    *string                                     `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName    *string                                     `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductSkuCode *string                                     `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	ProductType    *string                                     `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RelationalData *DescribeInstanceResponseBodyRelationalData `json:"RelationalData,omitempty" xml:"RelationalData,omitempty" type:"Struct"`
	Status         *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName   *string                                     `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetAppJson(v string) *DescribeInstanceResponseBody {
	s.AppJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetAutoRenewal(v string) *DescribeInstanceResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetBeganOn(v int64) *DescribeInstanceResponseBody {
	s.BeganOn = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetComponentJson(v string) *DescribeInstanceResponseBody {
	s.ComponentJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetConstraints(v string) *DescribeInstanceResponseBody {
	s.Constraints = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreatedOn(v int64) *DescribeInstanceResponseBody {
	s.CreatedOn = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEndOn(v int64) *DescribeInstanceResponseBody {
	s.EndOn = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetExtendJson(v string) *DescribeInstanceResponseBody {
	s.ExtendJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetHostJson(v string) *DescribeInstanceResponseBody {
	s.HostJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v int64) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsTrial(v bool) *DescribeInstanceResponseBody {
	s.IsTrial = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModules(v *DescribeInstanceResponseBodyModules) *DescribeInstanceResponseBody {
	s.Modules = v
	return s
}

func (s *DescribeInstanceResponseBody) SetOrderId(v int64) *DescribeInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProductCode(v string) *DescribeInstanceResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProductName(v string) *DescribeInstanceResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProductSkuCode(v string) *DescribeInstanceResponseBody {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProductType(v string) *DescribeInstanceResponseBody {
	s.ProductType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRelationalData(v *DescribeInstanceResponseBodyRelationalData) *DescribeInstanceResponseBody {
	s.RelationalData = v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v string) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetSupplierName(v string) *DescribeInstanceResponseBody {
	s.SupplierName = &v
	return s
}

type DescribeInstanceResponseBodyModules struct {
	Module []*DescribeInstanceResponseBodyModulesModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyModules) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyModules) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModules) SetModule(v []*DescribeInstanceResponseBodyModulesModule) *DescribeInstanceResponseBodyModules {
	s.Module = v
	return s
}

type DescribeInstanceResponseBodyModulesModule struct {
	Code       *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Id         *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties *DescribeInstanceResponseBodyModulesModuleProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
}

func (s DescribeInstanceResponseBodyModulesModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModule) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModule) SetCode(v string) *DescribeInstanceResponseBodyModulesModule {
	s.Code = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModule) SetId(v string) *DescribeInstanceResponseBodyModulesModule {
	s.Id = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModule) SetName(v string) *DescribeInstanceResponseBodyModulesModule {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModule) SetProperties(v *DescribeInstanceResponseBodyModulesModuleProperties) *DescribeInstanceResponseBodyModulesModule {
	s.Properties = v
	return s
}

type DescribeInstanceResponseBodyModulesModuleProperties struct {
	Property []*DescribeInstanceResponseBodyModulesModulePropertiesProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyModulesModuleProperties) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModuleProperties) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModuleProperties) SetProperty(v []*DescribeInstanceResponseBodyModulesModulePropertiesProperty) *DescribeInstanceResponseBodyModulesModuleProperties {
	s.Property = v
	return s
}

type DescribeInstanceResponseBodyModulesModulePropertiesProperty struct {
	DisplayUnit    *string                                                                    `json:"DisplayUnit,omitempty" xml:"DisplayUnit,omitempty"`
	Key            *string                                                                    `json:"Key,omitempty" xml:"Key,omitempty"`
	Name           *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	PropertyValues *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Struct"`
	ShowType       *string                                                                    `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesProperty) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesProperty) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetDisplayUnit(v string) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.DisplayUnit = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetKey(v string) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.Key = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetName(v string) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetPropertyValues(v *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.PropertyValues = v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetShowType(v string) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.ShowType = &v
	return s
}

type DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues struct {
	PropertyValue []*DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) SetPropertyValue(v []*DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues {
	s.PropertyValue = v
	return s
}

type DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Max         *string `json:"Max,omitempty" xml:"Max,omitempty"`
	Min         *string `json:"Min,omitempty" xml:"Min,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Step        *string `json:"Step,omitempty" xml:"Step,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetDisplayName(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.DisplayName = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetMax(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Max = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetMin(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Min = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetRemark(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetStep(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Step = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetType(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetValue(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Value = &v
	return s
}

type DescribeInstanceResponseBodyRelationalData struct {
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s DescribeInstanceResponseBodyRelationalData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyRelationalData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyRelationalData) SetServiceStatus(v string) *DescribeInstanceResponseBodyRelationalData {
	s.ServiceStatus = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	Codes       *string `json:"Codes,omitempty" xml:"Codes,omitempty"`
	ExceptCodes *string `json:"ExceptCodes,omitempty" xml:"ExceptCodes,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetCodes(v string) *DescribeInstancesRequest {
	s.Codes = &v
	return s
}

func (s *DescribeInstancesRequest) SetExceptCodes(v string) *DescribeInstancesRequest {
	s.ExceptCodes = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetProductType(v string) *DescribeInstancesRequest {
	s.ProductType = &v
	return s
}

type DescribeInstancesResponseBody struct {
	InstanceItems *DescribeInstancesResponseBodyInstanceItems `json:"InstanceItems,omitempty" xml:"InstanceItems,omitempty" type:"Struct"`
	PageNumber    *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstanceItems(v *DescribeInstancesResponseBodyInstanceItems) *DescribeInstancesResponseBody {
	s.InstanceItems = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstancesResponseBodyInstanceItems struct {
	InstanceItem []*DescribeInstancesResponseBodyInstanceItemsInstanceItem `json:"InstanceItem,omitempty" xml:"InstanceItem,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstanceItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstanceItems) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstanceItems) SetInstanceItem(v []*DescribeInstancesResponseBodyInstanceItemsInstanceItem) *DescribeInstancesResponseBodyInstanceItems {
	s.InstanceItem = v
	return s
}

type DescribeInstancesResponseBodyInstanceItemsInstanceItem struct {
	ApiJson        *string `json:"ApiJson,omitempty" xml:"ApiJson,omitempty"`
	AppJson        *string `json:"AppJson,omitempty" xml:"AppJson,omitempty"`
	BeganOn        *int64  `json:"BeganOn,omitempty" xml:"BeganOn,omitempty"`
	CreatedOn      *int64  `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	EndOn          *int64  `json:"EndOn,omitempty" xml:"EndOn,omitempty"`
	ExtendJson     *string `json:"ExtendJson,omitempty" xml:"ExtendJson,omitempty"`
	HostJson       *string `json:"HostJson,omitempty" xml:"HostJson,omitempty"`
	IdaasJson      *string `json:"IdaasJson,omitempty" xml:"IdaasJson,omitempty"`
	ImageJson      *string `json:"ImageJson,omitempty" xml:"ImageJson,omitempty"`
	InstanceId     *int64  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId        *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductCode    *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName    *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	ProductType    *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName   *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeInstancesResponseBodyInstanceItemsInstanceItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstanceItemsInstanceItem) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetApiJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ApiJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetAppJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.AppJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetBeganOn(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.BeganOn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetCreatedOn(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.CreatedOn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetEndOn(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.EndOn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetExtendJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ExtendJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetHostJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.HostJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetIdaasJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.IdaasJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetImageJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ImageJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetInstanceId(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetOrderId(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.OrderId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetProductCode(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetProductName(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ProductName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetProductSkuCode(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetProductType(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ProductType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetStatus(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetSupplierName(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.SupplierName = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeLicenseRequest struct {
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
}

func (s DescribeLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLicenseRequest) GoString() string {
	return s.String()
}

func (s *DescribeLicenseRequest) SetLicenseCode(v string) *DescribeLicenseRequest {
	s.LicenseCode = &v
	return s
}

type DescribeLicenseResponseBody struct {
	License   *DescribeLicenseResponseBodyLicense `json:"License,omitempty" xml:"License,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBody) SetLicense(v *DescribeLicenseResponseBodyLicense) *DescribeLicenseResponseBody {
	s.License = v
	return s
}

func (s *DescribeLicenseResponseBody) SetRequestId(v string) *DescribeLicenseResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLicenseResponseBodyLicense struct {
	ActivateTime  *string                                        `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	CreateTime    *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredTime   *string                                        `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ExtendArray   *DescribeLicenseResponseBodyLicenseExtendArray `json:"ExtendArray,omitempty" xml:"ExtendArray,omitempty" type:"Struct"`
	ExtendInfo    *DescribeLicenseResponseBodyLicenseExtendInfo  `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty" type:"Struct"`
	InstanceId    *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LicenseCode   *string                                        `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	LicenseStatus *string                                        `json:"LicenseStatus,omitempty" xml:"LicenseStatus,omitempty"`
	ProductCode   *string                                        `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName   *string                                        `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductSkuId  *string                                        `json:"ProductSkuId,omitempty" xml:"ProductSkuId,omitempty"`
	SupplierName  *string                                        `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeLicenseResponseBodyLicense) String() string {
	return tea.Prettify(s)
}

func (s DescribeLicenseResponseBodyLicense) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBodyLicense) SetActivateTime(v string) *DescribeLicenseResponseBodyLicense {
	s.ActivateTime = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetCreateTime(v string) *DescribeLicenseResponseBodyLicense {
	s.CreateTime = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetExpiredTime(v string) *DescribeLicenseResponseBodyLicense {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetExtendArray(v *DescribeLicenseResponseBodyLicenseExtendArray) *DescribeLicenseResponseBodyLicense {
	s.ExtendArray = v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetExtendInfo(v *DescribeLicenseResponseBodyLicenseExtendInfo) *DescribeLicenseResponseBodyLicense {
	s.ExtendInfo = v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetInstanceId(v string) *DescribeLicenseResponseBodyLicense {
	s.InstanceId = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetLicenseCode(v string) *DescribeLicenseResponseBodyLicense {
	s.LicenseCode = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetLicenseStatus(v string) *DescribeLicenseResponseBodyLicense {
	s.LicenseStatus = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetProductCode(v string) *DescribeLicenseResponseBodyLicense {
	s.ProductCode = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetProductName(v string) *DescribeLicenseResponseBodyLicense {
	s.ProductName = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetProductSkuId(v string) *DescribeLicenseResponseBodyLicense {
	s.ProductSkuId = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetSupplierName(v string) *DescribeLicenseResponseBodyLicense {
	s.SupplierName = &v
	return s
}

type DescribeLicenseResponseBodyLicenseExtendArray struct {
	LicenseAttribute []*DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute `json:"LicenseAttribute,omitempty" xml:"LicenseAttribute,omitempty" type:"Repeated"`
}

func (s DescribeLicenseResponseBodyLicenseExtendArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeLicenseResponseBodyLicenseExtendArray) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBodyLicenseExtendArray) SetLicenseAttribute(v []*DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) *DescribeLicenseResponseBodyLicenseExtendArray {
	s.LicenseAttribute = v
	return s
}

type DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) SetCode(v string) *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute {
	s.Code = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) SetValue(v string) *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute {
	s.Value = &v
	return s
}

type DescribeLicenseResponseBodyLicenseExtendInfo struct {
	AccountQuantity *int64  `json:"AccountQuantity,omitempty" xml:"AccountQuantity,omitempty"`
	AliUid          *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile          *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s DescribeLicenseResponseBodyLicenseExtendInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLicenseResponseBodyLicenseExtendInfo) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) SetAccountQuantity(v int64) *DescribeLicenseResponseBodyLicenseExtendInfo {
	s.AccountQuantity = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) SetAliUid(v int64) *DescribeLicenseResponseBodyLicenseExtendInfo {
	s.AliUid = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) SetEmail(v string) *DescribeLicenseResponseBodyLicenseExtendInfo {
	s.Email = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) SetMobile(v string) *DescribeLicenseResponseBodyLicenseExtendInfo {
	s.Mobile = &v
	return s
}

type DescribeLicenseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLicenseResponse) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponse) SetHeaders(v map[string]*string) *DescribeLicenseResponse {
	s.Headers = v
	return s
}

func (s *DescribeLicenseResponse) SetStatusCode(v int32) *DescribeLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLicenseResponse) SetBody(v *DescribeLicenseResponseBody) *DescribeLicenseResponse {
	s.Body = v
	return s
}

type DescribeOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DescribeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrderRequest) SetOrderId(v string) *DescribeOrderRequest {
	s.OrderId = &v
	return s
}

type DescribeOrderResponseBody struct {
	AccountQuantity     *int64                                       `json:"AccountQuantity,omitempty" xml:"AccountQuantity,omitempty"`
	AliUid              *int64                                       `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Components          map[string]interface{}                       `json:"Components,omitempty" xml:"Components,omitempty"`
	CouponPrice         *float32                                     `json:"CouponPrice,omitempty" xml:"CouponPrice,omitempty"`
	CreatedOn           *int64                                       `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	InstanceIds         *DescribeOrderResponseBodyInstanceIds        `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	OrderId             *int64                                       `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderStatus         *string                                      `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderType           *string                                      `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OriginalPrice       *float32                                     `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PaidOn              *int64                                       `json:"PaidOn,omitempty" xml:"PaidOn,omitempty"`
	PayStatus           *string                                      `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	PaymentPrice        *float32                                     `json:"PaymentPrice,omitempty" xml:"PaymentPrice,omitempty"`
	PeriodType          *string                                      `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	ProductCode         *string                                      `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName         *string                                      `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductSkuCode      *string                                      `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	Quantity            *int32                                       `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	RequestId           *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupplierCompanyName *string                                      `json:"SupplierCompanyName,omitempty" xml:"SupplierCompanyName,omitempty"`
	SupplierTelephones  *DescribeOrderResponseBodySupplierTelephones `json:"SupplierTelephones,omitempty" xml:"SupplierTelephones,omitempty" type:"Struct"`
	TotalPrice          *float32                                     `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrderResponseBody) SetAccountQuantity(v int64) *DescribeOrderResponseBody {
	s.AccountQuantity = &v
	return s
}

func (s *DescribeOrderResponseBody) SetAliUid(v int64) *DescribeOrderResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeOrderResponseBody) SetComponents(v map[string]interface{}) *DescribeOrderResponseBody {
	s.Components = v
	return s
}

func (s *DescribeOrderResponseBody) SetCouponPrice(v float32) *DescribeOrderResponseBody {
	s.CouponPrice = &v
	return s
}

func (s *DescribeOrderResponseBody) SetCreatedOn(v int64) *DescribeOrderResponseBody {
	s.CreatedOn = &v
	return s
}

func (s *DescribeOrderResponseBody) SetInstanceIds(v *DescribeOrderResponseBodyInstanceIds) *DescribeOrderResponseBody {
	s.InstanceIds = v
	return s
}

func (s *DescribeOrderResponseBody) SetOrderId(v int64) *DescribeOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeOrderResponseBody) SetOrderStatus(v string) *DescribeOrderResponseBody {
	s.OrderStatus = &v
	return s
}

func (s *DescribeOrderResponseBody) SetOrderType(v string) *DescribeOrderResponseBody {
	s.OrderType = &v
	return s
}

func (s *DescribeOrderResponseBody) SetOriginalPrice(v float32) *DescribeOrderResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeOrderResponseBody) SetPaidOn(v int64) *DescribeOrderResponseBody {
	s.PaidOn = &v
	return s
}

func (s *DescribeOrderResponseBody) SetPayStatus(v string) *DescribeOrderResponseBody {
	s.PayStatus = &v
	return s
}

func (s *DescribeOrderResponseBody) SetPaymentPrice(v float32) *DescribeOrderResponseBody {
	s.PaymentPrice = &v
	return s
}

func (s *DescribeOrderResponseBody) SetPeriodType(v string) *DescribeOrderResponseBody {
	s.PeriodType = &v
	return s
}

func (s *DescribeOrderResponseBody) SetProductCode(v string) *DescribeOrderResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeOrderResponseBody) SetProductName(v string) *DescribeOrderResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeOrderResponseBody) SetProductSkuCode(v string) *DescribeOrderResponseBody {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeOrderResponseBody) SetQuantity(v int32) *DescribeOrderResponseBody {
	s.Quantity = &v
	return s
}

func (s *DescribeOrderResponseBody) SetRequestId(v string) *DescribeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOrderResponseBody) SetSupplierCompanyName(v string) *DescribeOrderResponseBody {
	s.SupplierCompanyName = &v
	return s
}

func (s *DescribeOrderResponseBody) SetSupplierTelephones(v *DescribeOrderResponseBodySupplierTelephones) *DescribeOrderResponseBody {
	s.SupplierTelephones = v
	return s
}

func (s *DescribeOrderResponseBody) SetTotalPrice(v float32) *DescribeOrderResponseBody {
	s.TotalPrice = &v
	return s
}

type DescribeOrderResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeOrderResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeOrderResponseBodyInstanceIds) SetInstanceId(v []*string) *DescribeOrderResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type DescribeOrderResponseBodySupplierTelephones struct {
	Telephone []*string `json:"Telephone,omitempty" xml:"Telephone,omitempty" type:"Repeated"`
}

func (s DescribeOrderResponseBodySupplierTelephones) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderResponseBodySupplierTelephones) GoString() string {
	return s.String()
}

func (s *DescribeOrderResponseBodySupplierTelephones) SetTelephone(v []*string) *DescribeOrderResponseBodySupplierTelephones {
	s.Telephone = v
	return s
}

type DescribeOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrderResponse) SetHeaders(v map[string]*string) *DescribeOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrderResponse) SetStatusCode(v int32) *DescribeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOrderResponse) SetBody(v *DescribeOrderResponseBody) *DescribeOrderResponse {
	s.Body = v
	return s
}

type DescribePriceRequest struct {
	Commodity *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetCommodity(v string) *DescribePriceRequest {
	s.Commodity = &v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

type DescribePriceResponseBody struct {
	Coupons           *DescribePriceResponseBodyCoupons        `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	Currency          *string                                  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	Cuxiao            *bool                                    `json:"Cuxiao,omitempty" xml:"Cuxiao,omitempty"`
	Cycle             *string                                  `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	DiscountPrice     *float32                                 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	Duration          *int32                                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExpressionCode    *string                                  `json:"ExpressionCode,omitempty" xml:"ExpressionCode,omitempty"`
	ExpressionMessage *string                                  `json:"ExpressionMessage,omitempty" xml:"ExpressionMessage,omitempty"`
	InfoTitle         *string                                  `json:"InfoTitle,omitempty" xml:"InfoTitle,omitempty"`
	OriginalPrice     *float32                                 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	ProductCode       *string                                  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PromotionRules    *DescribePriceResponseBodyPromotionRules `json:"PromotionRules,omitempty" xml:"PromotionRules,omitempty" type:"Struct"`
	TradePrice        *float32                                 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) SetCoupons(v *DescribePriceResponseBodyCoupons) *DescribePriceResponseBody {
	s.Coupons = v
	return s
}

func (s *DescribePriceResponseBody) SetCurrency(v string) *DescribePriceResponseBody {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBody) SetCuxiao(v bool) *DescribePriceResponseBody {
	s.Cuxiao = &v
	return s
}

func (s *DescribePriceResponseBody) SetCycle(v string) *DescribePriceResponseBody {
	s.Cycle = &v
	return s
}

func (s *DescribePriceResponseBody) SetDiscountPrice(v float32) *DescribePriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBody) SetDuration(v int32) *DescribePriceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribePriceResponseBody) SetExpressionCode(v string) *DescribePriceResponseBody {
	s.ExpressionCode = &v
	return s
}

func (s *DescribePriceResponseBody) SetExpressionMessage(v string) *DescribePriceResponseBody {
	s.ExpressionMessage = &v
	return s
}

func (s *DescribePriceResponseBody) SetInfoTitle(v string) *DescribePriceResponseBody {
	s.InfoTitle = &v
	return s
}

func (s *DescribePriceResponseBody) SetOriginalPrice(v float32) *DescribePriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBody) SetProductCode(v string) *DescribePriceResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribePriceResponseBody) SetPromotionRules(v *DescribePriceResponseBodyPromotionRules) *DescribePriceResponseBody {
	s.PromotionRules = v
	return s
}

func (s *DescribePriceResponseBody) SetTradePrice(v float32) *DescribePriceResponseBody {
	s.TradePrice = &v
	return s
}

type DescribePriceResponseBodyCoupons struct {
	Coupon []*DescribePriceResponseBodyCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyCoupons) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyCoupons) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyCoupons) SetCoupon(v []*DescribePriceResponseBodyCouponsCoupon) *DescribePriceResponseBodyCoupons {
	s.Coupon = v
	return s
}

type DescribePriceResponseBodyCouponsCoupon struct {
	CanPromFee       *float32 `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	CouponDesc       *string  `json:"CouponDesc,omitempty" xml:"CouponDesc,omitempty"`
	CouponName       *string  `json:"CouponName,omitempty" xml:"CouponName,omitempty"`
	CouponOptionCode *string  `json:"CouponOptionCode,omitempty" xml:"CouponOptionCode,omitempty"`
	CouponOptionNo   *string  `json:"CouponOptionNo,omitempty" xml:"CouponOptionNo,omitempty"`
	IsSelected       *bool    `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
}

func (s DescribePriceResponseBodyCouponsCoupon) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCanPromFee(v float32) *DescribePriceResponseBodyCouponsCoupon {
	s.CanPromFee = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCouponDesc(v string) *DescribePriceResponseBodyCouponsCoupon {
	s.CouponDesc = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCouponName(v string) *DescribePriceResponseBodyCouponsCoupon {
	s.CouponName = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCouponOptionCode(v string) *DescribePriceResponseBodyCouponsCoupon {
	s.CouponOptionCode = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCouponOptionNo(v string) *DescribePriceResponseBodyCouponsCoupon {
	s.CouponOptionNo = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetIsSelected(v bool) *DescribePriceResponseBodyCouponsCoupon {
	s.IsSelected = &v
	return s
}

type DescribePriceResponseBodyPromotionRules struct {
	PromotionRule []*DescribePriceResponseBodyPromotionRulesPromotionRule `json:"PromotionRule,omitempty" xml:"PromotionRule,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPromotionRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPromotionRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPromotionRules) SetPromotionRule(v []*DescribePriceResponseBodyPromotionRulesPromotionRule) *DescribePriceResponseBodyPromotionRules {
	s.PromotionRule = v
	return s
}

type DescribePriceResponseBodyPromotionRulesPromotionRule struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePriceResponseBodyPromotionRulesPromotionRule) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPromotionRulesPromotionRule) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) SetName(v string) *DescribePriceResponseBodyPromotionRulesPromotionRule {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) SetRuleId(v string) *DescribePriceResponseBodyPromotionRulesPromotionRule {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) SetTitle(v string) *DescribePriceResponseBodyPromotionRulesPromotionRule {
	s.Title = &v
	return s
}

type DescribePriceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetStatusCode(v int32) *DescribePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

type DescribeProductRequest struct {
	// AliUid
	AliUid     *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	QueryDraft *bool   `json:"QueryDraft,omitempty" xml:"QueryDraft,omitempty"`
}

func (s DescribeProductRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductRequest) SetAliUid(v string) *DescribeProductRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeProductRequest) SetCode(v string) *DescribeProductRequest {
	s.Code = &v
	return s
}

func (s *DescribeProductRequest) SetQueryDraft(v bool) *DescribeProductRequest {
	s.QueryDraft = &v
	return s
}

type DescribeProductResponseBody struct {
	AuditFailMsg     *string                                   `json:"AuditFailMsg,omitempty" xml:"AuditFailMsg,omitempty"`
	AuditStatus      *string                                   `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuditTime        *int64                                    `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	Code             *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Description      *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	FrontCategoryId  *int64                                    `json:"FrontCategoryId,omitempty" xml:"FrontCategoryId,omitempty"`
	GmtCreated       *int64                                    `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified      *int64                                    `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name             *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PicUrl           *string                                   `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	ProductExtras    *DescribeProductResponseBodyProductExtras `json:"ProductExtras,omitempty" xml:"ProductExtras,omitempty" type:"Struct"`
	ProductSkus      *DescribeProductResponseBodyProductSkus   `json:"ProductSkus,omitempty" xml:"ProductSkus,omitempty" type:"Struct"`
	RequestId        *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Score            *float32                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	ShopInfo         *DescribeProductResponseBodyShopInfo      `json:"ShopInfo,omitempty" xml:"ShopInfo,omitempty" type:"Struct"`
	ShortDescription *string                                   `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	Status           *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierPk       *int64                                    `json:"SupplierPk,omitempty" xml:"SupplierPk,omitempty"`
	Type             *string                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	UseCount         *int64                                    `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s DescribeProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBody) SetAuditFailMsg(v string) *DescribeProductResponseBody {
	s.AuditFailMsg = &v
	return s
}

func (s *DescribeProductResponseBody) SetAuditStatus(v string) *DescribeProductResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *DescribeProductResponseBody) SetAuditTime(v int64) *DescribeProductResponseBody {
	s.AuditTime = &v
	return s
}

func (s *DescribeProductResponseBody) SetCode(v string) *DescribeProductResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProductResponseBody) SetDescription(v string) *DescribeProductResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeProductResponseBody) SetFrontCategoryId(v int64) *DescribeProductResponseBody {
	s.FrontCategoryId = &v
	return s
}

func (s *DescribeProductResponseBody) SetGmtCreated(v int64) *DescribeProductResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *DescribeProductResponseBody) SetGmtModified(v int64) *DescribeProductResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeProductResponseBody) SetName(v string) *DescribeProductResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBody) SetPicUrl(v string) *DescribeProductResponseBody {
	s.PicUrl = &v
	return s
}

func (s *DescribeProductResponseBody) SetProductExtras(v *DescribeProductResponseBodyProductExtras) *DescribeProductResponseBody {
	s.ProductExtras = v
	return s
}

func (s *DescribeProductResponseBody) SetProductSkus(v *DescribeProductResponseBodyProductSkus) *DescribeProductResponseBody {
	s.ProductSkus = v
	return s
}

func (s *DescribeProductResponseBody) SetRequestId(v string) *DescribeProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductResponseBody) SetScore(v float32) *DescribeProductResponseBody {
	s.Score = &v
	return s
}

func (s *DescribeProductResponseBody) SetShopInfo(v *DescribeProductResponseBodyShopInfo) *DescribeProductResponseBody {
	s.ShopInfo = v
	return s
}

func (s *DescribeProductResponseBody) SetShortDescription(v string) *DescribeProductResponseBody {
	s.ShortDescription = &v
	return s
}

func (s *DescribeProductResponseBody) SetStatus(v string) *DescribeProductResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeProductResponseBody) SetSupplierPk(v int64) *DescribeProductResponseBody {
	s.SupplierPk = &v
	return s
}

func (s *DescribeProductResponseBody) SetType(v string) *DescribeProductResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeProductResponseBody) SetUseCount(v int64) *DescribeProductResponseBody {
	s.UseCount = &v
	return s
}

type DescribeProductResponseBodyProductExtras struct {
	ProductExtra []*DescribeProductResponseBodyProductExtrasProductExtra `json:"ProductExtra,omitempty" xml:"ProductExtra,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductExtras) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductExtras) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductExtras) SetProductExtra(v []*DescribeProductResponseBodyProductExtrasProductExtra) *DescribeProductResponseBodyProductExtras {
	s.ProductExtra = v
	return s
}

type DescribeProductResponseBodyProductExtrasProductExtra struct {
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Label  *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Order  *int32  `json:"Order,omitempty" xml:"Order,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeProductResponseBodyProductExtrasProductExtra) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductExtrasProductExtra) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetKey(v string) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Key = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetLabel(v string) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Label = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetOrder(v int32) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Order = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetType(v string) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Type = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetValues(v string) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Values = &v
	return s
}

type DescribeProductResponseBodyProductSkus struct {
	ProductSku []*DescribeProductResponseBodyProductSkusProductSku `json:"ProductSku,omitempty" xml:"ProductSku,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkus) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkus) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkus) SetProductSku(v []*DescribeProductResponseBodyProductSkusProductSku) *DescribeProductResponseBodyProductSkus {
	s.ProductSku = v
	return s
}

type DescribeProductResponseBodyProductSkusProductSku struct {
	ChargeType   *string                                                       `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Code         *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Constraints  *string                                                       `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	Hidden       *bool                                                         `json:"Hidden,omitempty" xml:"Hidden,omitempty"`
	Modules      *DescribeProductResponseBodyProductSkusProductSkuModules      `json:"Modules,omitempty" xml:"Modules,omitempty" type:"Struct"`
	Name         *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderPeriods *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods `json:"OrderPeriods,omitempty" xml:"OrderPeriods,omitempty" type:"Struct"`
}

func (s DescribeProductResponseBodyProductSkusProductSku) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSku) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetChargeType(v string) *DescribeProductResponseBodyProductSkusProductSku {
	s.ChargeType = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetCode(v string) *DescribeProductResponseBodyProductSkusProductSku {
	s.Code = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetConstraints(v string) *DescribeProductResponseBodyProductSkusProductSku {
	s.Constraints = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetHidden(v bool) *DescribeProductResponseBodyProductSkusProductSku {
	s.Hidden = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetModules(v *DescribeProductResponseBodyProductSkusProductSkuModules) *DescribeProductResponseBodyProductSkusProductSku {
	s.Modules = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetName(v string) *DescribeProductResponseBodyProductSkusProductSku {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetOrderPeriods(v *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) *DescribeProductResponseBodyProductSkusProductSku {
	s.OrderPeriods = v
	return s
}

type DescribeProductResponseBodyProductSkusProductSkuModules struct {
	Module []*DescribeProductResponseBodyProductSkusProductSkuModulesModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModules) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModules) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModules) SetModule(v []*DescribeProductResponseBodyProductSkusProductSkuModulesModule) *DescribeProductResponseBodyProductSkusProductSkuModules {
	s.Module = v
	return s
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModule struct {
	Code       *string                                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Id         *string                                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModule) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) SetCode(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	s.Code = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) SetId(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	s.Id = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) SetName(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) SetProperties(v *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) *DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	s.Properties = v
	return s
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties struct {
	Property []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) SetProperty(v []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties {
	s.Property = v
	return s
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty struct {
	DisplayUnit    *string                                                                                        `json:"DisplayUnit,omitempty" xml:"DisplayUnit,omitempty"`
	Key            *string                                                                                        `json:"Key,omitempty" xml:"Key,omitempty"`
	Name           *string                                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	PropertyValues *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Struct"`
	ShowType       *string                                                                                        `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetDisplayUnit(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.DisplayUnit = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetKey(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.Key = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetName(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetPropertyValues(v *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.PropertyValues = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetShowType(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.ShowType = &v
	return s
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues struct {
	PropertyValue []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) SetPropertyValue(v []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues {
	s.PropertyValue = v
	return s
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Max         *string `json:"Max,omitempty" xml:"Max,omitempty"`
	Min         *string `json:"Min,omitempty" xml:"Min,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Step        *string `json:"Step,omitempty" xml:"Step,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetDisplayName(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.DisplayName = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetMax(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Max = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetMin(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Min = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetRemark(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Remark = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetStep(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Step = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetType(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Type = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetValue(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Value = &v
	return s
}

type DescribeProductResponseBodyProductSkusProductSkuOrderPeriods struct {
	OrderPeriod []*DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod `json:"OrderPeriod,omitempty" xml:"OrderPeriod,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) SetOrderPeriod(v []*DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods {
	s.OrderPeriod = v
	return s
}

type DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) SetName(v string) *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) SetPeriodType(v string) *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod {
	s.PeriodType = &v
	return s
}

type DescribeProductResponseBodyShopInfo struct {
	Emails     *string                                        `json:"Emails,omitempty" xml:"Emails,omitempty"`
	Id         *int64                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Telephones *DescribeProductResponseBodyShopInfoTelephones `json:"Telephones,omitempty" xml:"Telephones,omitempty" type:"Struct"`
	WangWangs  *DescribeProductResponseBodyShopInfoWangWangs  `json:"WangWangs,omitempty" xml:"WangWangs,omitempty" type:"Struct"`
}

func (s DescribeProductResponseBodyShopInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyShopInfo) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyShopInfo) SetEmails(v string) *DescribeProductResponseBodyShopInfo {
	s.Emails = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) SetId(v int64) *DescribeProductResponseBodyShopInfo {
	s.Id = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) SetName(v string) *DescribeProductResponseBodyShopInfo {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) SetTelephones(v *DescribeProductResponseBodyShopInfoTelephones) *DescribeProductResponseBodyShopInfo {
	s.Telephones = v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) SetWangWangs(v *DescribeProductResponseBodyShopInfoWangWangs) *DescribeProductResponseBodyShopInfo {
	s.WangWangs = v
	return s
}

type DescribeProductResponseBodyShopInfoTelephones struct {
	Telephone []*string `json:"Telephone,omitempty" xml:"Telephone,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyShopInfoTelephones) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyShopInfoTelephones) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyShopInfoTelephones) SetTelephone(v []*string) *DescribeProductResponseBodyShopInfoTelephones {
	s.Telephone = v
	return s
}

type DescribeProductResponseBodyShopInfoWangWangs struct {
	WangWang []*DescribeProductResponseBodyShopInfoWangWangsWangWang `json:"WangWang,omitempty" xml:"WangWang,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyShopInfoWangWangs) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyShopInfoWangWangs) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyShopInfoWangWangs) SetWangWang(v []*DescribeProductResponseBodyShopInfoWangWangsWangWang) *DescribeProductResponseBodyShopInfoWangWangs {
	s.WangWang = v
	return s
}

type DescribeProductResponseBodyShopInfoWangWangsWangWang struct {
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeProductResponseBodyShopInfoWangWangsWangWang) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponseBodyShopInfoWangWangsWangWang) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyShopInfoWangWangsWangWang) SetRemark(v string) *DescribeProductResponseBodyShopInfoWangWangsWangWang {
	s.Remark = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfoWangWangsWangWang) SetUserName(v string) *DescribeProductResponseBodyShopInfoWangWangsWangWang {
	s.UserName = &v
	return s
}

type DescribeProductResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductResponse) SetHeaders(v map[string]*string) *DescribeProductResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductResponse) SetStatusCode(v int32) *DescribeProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductResponse) SetBody(v *DescribeProductResponseBody) *DescribeProductResponse {
	s.Body = v
	return s
}

type DescribeProductsRequest struct {
	Filter     []*DescribeProductsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchTerm *string                          `json:"SearchTerm,omitempty" xml:"SearchTerm,omitempty"`
}

func (s DescribeProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductsRequest) SetFilter(v []*DescribeProductsRequestFilter) *DescribeProductsRequest {
	s.Filter = v
	return s
}

func (s *DescribeProductsRequest) SetPageNumber(v int32) *DescribeProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProductsRequest) SetPageSize(v int32) *DescribeProductsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProductsRequest) SetSearchTerm(v string) *DescribeProductsRequest {
	s.SearchTerm = &v
	return s
}

type DescribeProductsRequestFilter struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeProductsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeProductsRequestFilter) SetKey(v string) *DescribeProductsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeProductsRequestFilter) SetValue(v string) *DescribeProductsRequestFilter {
	s.Value = &v
	return s
}

type DescribeProductsResponseBody struct {
	PageNumber   *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductItems *DescribeProductsResponseBodyProductItems `json:"ProductItems,omitempty" xml:"ProductItems,omitempty" type:"Struct"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBody) SetPageNumber(v int32) *DescribeProductsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProductsResponseBody) SetPageSize(v int32) *DescribeProductsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProductsResponseBody) SetProductItems(v *DescribeProductsResponseBodyProductItems) *DescribeProductsResponseBody {
	s.ProductItems = v
	return s
}

func (s *DescribeProductsResponseBody) SetRequestId(v string) *DescribeProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductsResponseBody) SetTotalCount(v int32) *DescribeProductsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeProductsResponseBodyProductItems struct {
	ProductItem []*DescribeProductsResponseBodyProductItemsProductItem `json:"ProductItem,omitempty" xml:"ProductItem,omitempty" type:"Repeated"`
}

func (s DescribeProductsResponseBodyProductItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponseBodyProductItems) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyProductItems) SetProductItem(v []*DescribeProductsResponseBodyProductItemsProductItem) *DescribeProductsResponseBodyProductItems {
	s.ProductItem = v
	return s
}

type DescribeProductsResponseBodyProductItemsProductItem struct {
	CategoryId       *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DeliveryDate     *string `json:"DeliveryDate,omitempty" xml:"DeliveryDate,omitempty"`
	DeliveryWay      *string `json:"DeliveryWay,omitempty" xml:"DeliveryWay,omitempty"`
	ImageUrl         *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationSystem  *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	PriceInfo        *string `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty"`
	Score            *string `json:"Score,omitempty" xml:"Score,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	SuggestedPrice   *string `json:"SuggestedPrice,omitempty" xml:"SuggestedPrice,omitempty"`
	SupplierId       *int64  `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
	SupplierName     *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TargetUrl        *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	WarrantyDate     *string `json:"WarrantyDate,omitempty" xml:"WarrantyDate,omitempty"`
}

func (s DescribeProductsResponseBodyProductItemsProductItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponseBodyProductItemsProductItem) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetCategoryId(v int64) *DescribeProductsResponseBodyProductItemsProductItem {
	s.CategoryId = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetCode(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.Code = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetDeliveryDate(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.DeliveryDate = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetDeliveryWay(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.DeliveryWay = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetImageUrl(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.ImageUrl = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetName(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.Name = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetOperationSystem(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.OperationSystem = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetPriceInfo(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.PriceInfo = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetScore(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.Score = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetShortDescription(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.ShortDescription = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetSuggestedPrice(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.SuggestedPrice = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetSupplierId(v int64) *DescribeProductsResponseBodyProductItemsProductItem {
	s.SupplierId = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetSupplierName(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.SupplierName = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetTags(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.Tags = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetTargetUrl(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.TargetUrl = &v
	return s
}

func (s *DescribeProductsResponseBodyProductItemsProductItem) SetWarrantyDate(v string) *DescribeProductsResponseBodyProductItemsProductItem {
	s.WarrantyDate = &v
	return s
}

type DescribeProductsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponse) SetHeaders(v map[string]*string) *DescribeProductsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductsResponse) SetStatusCode(v int32) *DescribeProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductsResponse) SetBody(v *DescribeProductsResponseBody) *DescribeProductsResponse {
	s.Body = v
	return s
}

type DescribeProjectAttachmentsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProjectAttachmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectAttachmentsRequest) SetInstanceId(v string) *DescribeProjectAttachmentsRequest {
	s.InstanceId = &v
	return s
}

type DescribeProjectAttachmentsResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeProjectAttachmentsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectAttachmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAttachmentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectAttachmentsResponseBody) SetRequestId(v string) *DescribeProjectAttachmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBody) SetResult(v []*DescribeProjectAttachmentsResponseBodyResult) *DescribeProjectAttachmentsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectAttachmentsResponseBody) SetSuccess(v bool) *DescribeProjectAttachmentsResponseBody {
	s.Success = &v
	return s
}

type DescribeProjectAttachmentsResponseBodyResult struct {
	AttachmentToken    *string `json:"AttachmentToken,omitempty" xml:"AttachmentToken,omitempty"`
	AttachmentType     *string `json:"AttachmentType,omitempty" xml:"AttachmentType,omitempty"`
	Content            *string `json:"Content,omitempty" xml:"Content,omitempty"`
	FileLink           *string `json:"FileLink,omitempty" xml:"FileLink,omitempty"`
	FileLinkGmtExpired *int64  `json:"FileLinkGmtExpired,omitempty" xml:"FileLinkGmtExpired,omitempty"`
	FileName           *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize           *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileSuffix         *string `json:"FileSuffix,omitempty" xml:"FileSuffix,omitempty"`
	GmtCreate          *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	NodeId             *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName           *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Operator           *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName       *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	OperatorRole       *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	StepNo             *int32  `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
}

func (s DescribeProjectAttachmentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAttachmentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetAttachmentToken(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.AttachmentToken = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetAttachmentType(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.AttachmentType = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetContent(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileLink(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileLink = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileLinkGmtExpired(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileLinkGmtExpired = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileName(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileName = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileSize(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetFileSuffix(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.FileSuffix = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetGmtCreate(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetNodeId(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.NodeId = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetNodeName(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.NodeName = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetOperator(v int64) *DescribeProjectAttachmentsResponseBodyResult {
	s.Operator = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetOperatorName(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetOperatorRole(v string) *DescribeProjectAttachmentsResponseBodyResult {
	s.OperatorRole = &v
	return s
}

func (s *DescribeProjectAttachmentsResponseBodyResult) SetStepNo(v int32) *DescribeProjectAttachmentsResponseBodyResult {
	s.StepNo = &v
	return s
}

type DescribeProjectAttachmentsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProjectAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectAttachmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectAttachmentsResponse) SetHeaders(v map[string]*string) *DescribeProjectAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectAttachmentsResponse) SetStatusCode(v int32) *DescribeProjectAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectAttachmentsResponse) SetBody(v *DescribeProjectAttachmentsResponseBody) *DescribeProjectAttachmentsResponse {
	s.Body = v
	return s
}

type DescribeProjectInfoRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProjectInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectInfoRequest) SetInstanceId(v string) *DescribeProjectInfoRequest {
	s.InstanceId = &v
	return s
}

type DescribeProjectInfoResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeProjectInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectInfoResponseBody) GoString() string {
	return s.String()
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

type DescribeProjectInfoResponseBodyResult struct {
	CurrentStepNo  *int32  `json:"CurrentStepNo,omitempty" xml:"CurrentStepNo,omitempty"`
	CustomerAliUid *int64  `json:"CustomerAliUid,omitempty" xml:"CustomerAliUid,omitempty"`
	FinalStepNo    *int32  `json:"FinalStepNo,omitempty" xml:"FinalStepNo,omitempty"`
	FinishType     *string `json:"FinishType,omitempty" xml:"FinishType,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtExpired     *int64  `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	GmtFinished    *int64  `json:"GmtFinished,omitempty" xml:"GmtFinished,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId        *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductCode    *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName    *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	ProductSkuName *string `json:"ProductSkuName,omitempty" xml:"ProductSkuName,omitempty"`
	ProjectStatus  *string `json:"ProjectStatus,omitempty" xml:"ProjectStatus,omitempty"`
	SupplierAliUid *int64  `json:"SupplierAliUid,omitempty" xml:"SupplierAliUid,omitempty"`
	TemplateId     *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeProjectInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectInfoResponseBodyResult) GoString() string {
	return s.String()
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

type DescribeProjectInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectInfoResponse) SetHeaders(v map[string]*string) *DescribeProjectInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectInfoResponse) SetStatusCode(v int32) *DescribeProjectInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectInfoResponse) SetBody(v *DescribeProjectInfoResponseBody) *DescribeProjectInfoResponse {
	s.Body = v
	return s
}

type DescribeProjectMessagesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageIndex  *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
}

func (s DescribeProjectMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMessagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectMessagesRequest) SetInstanceId(v string) *DescribeProjectMessagesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProjectMessagesRequest) SetPageIndex(v int32) *DescribeProjectMessagesRequest {
	s.PageIndex = &v
	return s
}

type DescribeProjectMessagesResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*DescribeProjectMessagesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProjectMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectMessagesResponseBody) SetRequestId(v string) *DescribeProjectMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectMessagesResponseBody) SetResult(v []*DescribeProjectMessagesResponseBodyResult) *DescribeProjectMessagesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectMessagesResponseBody) SetSuccess(v bool) *DescribeProjectMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProjectMessagesResponseBody) SetTotalCount(v int64) *DescribeProjectMessagesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeProjectMessagesResponseBodyResult struct {
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Operator     *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
}

func (s DescribeProjectMessagesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMessagesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectMessagesResponseBodyResult) SetContent(v string) *DescribeProjectMessagesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) SetGmtCreate(v int64) *DescribeProjectMessagesResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) SetOperator(v int64) *DescribeProjectMessagesResponseBodyResult {
	s.Operator = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) SetOperatorName(v string) *DescribeProjectMessagesResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *DescribeProjectMessagesResponseBodyResult) SetOperatorRole(v string) *DescribeProjectMessagesResponseBodyResult {
	s.OperatorRole = &v
	return s
}

type DescribeProjectMessagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProjectMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMessagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectMessagesResponse) SetHeaders(v map[string]*string) *DescribeProjectMessagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectMessagesResponse) SetStatusCode(v int32) *DescribeProjectMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectMessagesResponse) SetBody(v *DescribeProjectMessagesResponseBody) *DescribeProjectMessagesResponse {
	s.Body = v
	return s
}

type DescribeProjectNodesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProjectNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectNodesRequest) SetInstanceId(v string) *DescribeProjectNodesRequest {
	s.InstanceId = &v
	return s
}

type DescribeProjectNodesResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeProjectNodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectNodesResponseBody) SetRequestId(v string) *DescribeProjectNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectNodesResponseBody) SetResult(v []*DescribeProjectNodesResponseBodyResult) *DescribeProjectNodesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectNodesResponseBody) SetSuccess(v bool) *DescribeProjectNodesResponseBody {
	s.Success = &v
	return s
}

type DescribeProjectNodesResponseBodyResult struct {
	AllowRollbackNode *bool   `json:"AllowRollbackNode,omitempty" xml:"AllowRollbackNode,omitempty"`
	AutoFinishNode    *bool   `json:"AutoFinishNode,omitempty" xml:"AutoFinishNode,omitempty"`
	FinalStepNo       *int32  `json:"FinalStepNo,omitempty" xml:"FinalStepNo,omitempty"`
	GmtExpired        *int64  `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	GmtFinished       *int64  `json:"GmtFinished,omitempty" xml:"GmtFinished,omitempty"`
	GmtStart          *int64  `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	NeedAttachment    *bool   `json:"NeedAttachment,omitempty" xml:"NeedAttachment,omitempty"`
	NextNodeId        *int64  `json:"NextNodeId,omitempty" xml:"NextNodeId,omitempty"`
	NodeId            *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName          *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeStatus        *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	OperatorRole      *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
	ParentNodeId      *int64  `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	PreviousNodeId    *int64  `json:"PreviousNodeId,omitempty" xml:"PreviousNodeId,omitempty"`
	StepNo            *int32  `json:"StepNo,omitempty" xml:"StepNo,omitempty"`
	TemplateForm      *string `json:"TemplateForm,omitempty" xml:"TemplateForm,omitempty"`
}

func (s DescribeProjectNodesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectNodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectNodesResponseBodyResult) SetAllowRollbackNode(v bool) *DescribeProjectNodesResponseBodyResult {
	s.AllowRollbackNode = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetAutoFinishNode(v bool) *DescribeProjectNodesResponseBodyResult {
	s.AutoFinishNode = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetFinalStepNo(v int32) *DescribeProjectNodesResponseBodyResult {
	s.FinalStepNo = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetGmtExpired(v int64) *DescribeProjectNodesResponseBodyResult {
	s.GmtExpired = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetGmtFinished(v int64) *DescribeProjectNodesResponseBodyResult {
	s.GmtFinished = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetGmtStart(v int64) *DescribeProjectNodesResponseBodyResult {
	s.GmtStart = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNeedAttachment(v bool) *DescribeProjectNodesResponseBodyResult {
	s.NeedAttachment = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNextNodeId(v int64) *DescribeProjectNodesResponseBodyResult {
	s.NextNodeId = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNodeId(v int64) *DescribeProjectNodesResponseBodyResult {
	s.NodeId = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNodeName(v string) *DescribeProjectNodesResponseBodyResult {
	s.NodeName = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetNodeStatus(v string) *DescribeProjectNodesResponseBodyResult {
	s.NodeStatus = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetOperatorRole(v string) *DescribeProjectNodesResponseBodyResult {
	s.OperatorRole = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetParentNodeId(v int64) *DescribeProjectNodesResponseBodyResult {
	s.ParentNodeId = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetPreviousNodeId(v int64) *DescribeProjectNodesResponseBodyResult {
	s.PreviousNodeId = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetStepNo(v int32) *DescribeProjectNodesResponseBodyResult {
	s.StepNo = &v
	return s
}

func (s *DescribeProjectNodesResponseBodyResult) SetTemplateForm(v string) *DescribeProjectNodesResponseBodyResult {
	s.TemplateForm = &v
	return s
}

type DescribeProjectNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProjectNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectNodesResponse) SetHeaders(v map[string]*string) *DescribeProjectNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectNodesResponse) SetStatusCode(v int32) *DescribeProjectNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectNodesResponse) SetBody(v *DescribeProjectNodesResponseBody) *DescribeProjectNodesResponse {
	s.Body = v
	return s
}

type DescribeProjectOperateLogsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProjectOperateLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectOperateLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectOperateLogsRequest) SetInstanceId(v string) *DescribeProjectOperateLogsRequest {
	s.InstanceId = &v
	return s
}

type DescribeProjectOperateLogsResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeProjectOperateLogsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectOperateLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectOperateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectOperateLogsResponseBody) SetRequestId(v string) *DescribeProjectOperateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBody) SetResult(v []*DescribeProjectOperateLogsResponseBodyResult) *DescribeProjectOperateLogsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeProjectOperateLogsResponseBody) SetSuccess(v bool) *DescribeProjectOperateLogsResponseBody {
	s.Success = &v
	return s
}

type DescribeProjectOperateLogsResponseBodyResult struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Operator     *int64  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	OperatorRole *string `json:"OperatorRole,omitempty" xml:"OperatorRole,omitempty"`
}

func (s DescribeProjectOperateLogsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectOperateLogsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetDescription(v string) *DescribeProjectOperateLogsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetGmtCreate(v int64) *DescribeProjectOperateLogsResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetOperator(v int64) *DescribeProjectOperateLogsResponseBodyResult {
	s.Operator = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetOperatorName(v string) *DescribeProjectOperateLogsResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *DescribeProjectOperateLogsResponseBodyResult) SetOperatorRole(v string) *DescribeProjectOperateLogsResponseBodyResult {
	s.OperatorRole = &v
	return s
}

type DescribeProjectOperateLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeProjectOperateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectOperateLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectOperateLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectOperateLogsResponse) SetHeaders(v map[string]*string) *DescribeProjectOperateLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectOperateLogsResponse) SetStatusCode(v int32) *DescribeProjectOperateLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectOperateLogsResponse) SetBody(v *DescribeProjectOperateLogsResponseBody) *DescribeProjectOperateLogsResponse {
	s.Body = v
	return s
}

type FinishCurrentProjectNodeRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId       *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TemplateForm *string `json:"TemplateForm,omitempty" xml:"TemplateForm,omitempty"`
}

func (s FinishCurrentProjectNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishCurrentProjectNodeRequest) GoString() string {
	return s.String()
}

func (s *FinishCurrentProjectNodeRequest) SetInstanceId(v string) *FinishCurrentProjectNodeRequest {
	s.InstanceId = &v
	return s
}

func (s *FinishCurrentProjectNodeRequest) SetNodeId(v int64) *FinishCurrentProjectNodeRequest {
	s.NodeId = &v
	return s
}

func (s *FinishCurrentProjectNodeRequest) SetRemark(v string) *FinishCurrentProjectNodeRequest {
	s.Remark = &v
	return s
}

func (s *FinishCurrentProjectNodeRequest) SetTemplateForm(v string) *FinishCurrentProjectNodeRequest {
	s.TemplateForm = &v
	return s
}

type FinishCurrentProjectNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FinishCurrentProjectNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishCurrentProjectNodeResponseBody) GoString() string {
	return s.String()
}

func (s *FinishCurrentProjectNodeResponseBody) SetRequestId(v string) *FinishCurrentProjectNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishCurrentProjectNodeResponseBody) SetResult(v bool) *FinishCurrentProjectNodeResponseBody {
	s.Result = &v
	return s
}

func (s *FinishCurrentProjectNodeResponseBody) SetSuccess(v bool) *FinishCurrentProjectNodeResponseBody {
	s.Success = &v
	return s
}

type FinishCurrentProjectNodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FinishCurrentProjectNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishCurrentProjectNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishCurrentProjectNodeResponse) GoString() string {
	return s.String()
}

func (s *FinishCurrentProjectNodeResponse) SetHeaders(v map[string]*string) *FinishCurrentProjectNodeResponse {
	s.Headers = v
	return s
}

func (s *FinishCurrentProjectNodeResponse) SetStatusCode(v int32) *FinishCurrentProjectNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishCurrentProjectNodeResponse) SetBody(v *FinishCurrentProjectNodeResponseBody) *FinishCurrentProjectNodeResponse {
	s.Body = v
	return s
}

type PauseProjectRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s PauseProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseProjectRequest) GoString() string {
	return s.String()
}

func (s *PauseProjectRequest) SetInstanceId(v string) *PauseProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *PauseProjectRequest) SetNodeId(v int64) *PauseProjectRequest {
	s.NodeId = &v
	return s
}

func (s *PauseProjectRequest) SetRemark(v string) *PauseProjectRequest {
	s.Remark = &v
	return s
}

type PauseProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PauseProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PauseProjectResponseBody) SetRequestId(v string) *PauseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseProjectResponseBody) SetResult(v bool) *PauseProjectResponseBody {
	s.Result = &v
	return s
}

func (s *PauseProjectResponseBody) SetSuccess(v bool) *PauseProjectResponseBody {
	s.Success = &v
	return s
}

type PauseProjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PauseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PauseProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseProjectResponse) GoString() string {
	return s.String()
}

func (s *PauseProjectResponse) SetHeaders(v map[string]*string) *PauseProjectResponse {
	s.Headers = v
	return s
}

func (s *PauseProjectResponse) SetStatusCode(v int32) *PauseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseProjectResponse) SetBody(v *PauseProjectResponseBody) *PauseProjectResponse {
	s.Body = v
	return s
}

type PushMeteringDataRequest struct {
	Metering *string `json:"Metering,omitempty" xml:"Metering,omitempty"`
}

func (s PushMeteringDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataRequest) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequest) SetMetering(v string) *PushMeteringDataRequest {
	s.Metering = &v
	return s
}

type PushMeteringDataResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushMeteringDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponseBody) SetRequestId(v string) *PushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMeteringDataResponseBody) SetSuccess(v bool) *PushMeteringDataResponseBody {
	s.Success = &v
	return s
}

type PushMeteringDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushMeteringDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushMeteringDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PushMeteringDataResponse) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponse) SetHeaders(v map[string]*string) *PushMeteringDataResponse {
	s.Headers = v
	return s
}

func (s *PushMeteringDataResponse) SetStatusCode(v int32) *PushMeteringDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMeteringDataResponse) SetBody(v *PushMeteringDataResponseBody) *PushMeteringDataResponse {
	s.Body = v
	return s
}

type ResumeProjectRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ResumeProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeProjectRequest) GoString() string {
	return s.String()
}

func (s *ResumeProjectRequest) SetInstanceId(v string) *ResumeProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *ResumeProjectRequest) SetNodeId(v int64) *ResumeProjectRequest {
	s.NodeId = &v
	return s
}

func (s *ResumeProjectRequest) SetRemark(v string) *ResumeProjectRequest {
	s.Remark = &v
	return s
}

type ResumeProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeProjectResponseBody) SetRequestId(v string) *ResumeProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeProjectResponseBody) SetResult(v bool) *ResumeProjectResponseBody {
	s.Result = &v
	return s
}

func (s *ResumeProjectResponseBody) SetSuccess(v bool) *ResumeProjectResponseBody {
	s.Success = &v
	return s
}

type ResumeProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeProjectResponse) GoString() string {
	return s.String()
}

func (s *ResumeProjectResponse) SetHeaders(v map[string]*string) *ResumeProjectResponse {
	s.Headers = v
	return s
}

func (s *ResumeProjectResponse) SetStatusCode(v int32) *ResumeProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeProjectResponse) SetBody(v *ResumeProjectResponseBody) *ResumeProjectResponse {
	s.Body = v
	return s
}

type RollbackCurrentProjectNodeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeId     *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s RollbackCurrentProjectNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackCurrentProjectNodeRequest) GoString() string {
	return s.String()
}

func (s *RollbackCurrentProjectNodeRequest) SetInstanceId(v string) *RollbackCurrentProjectNodeRequest {
	s.InstanceId = &v
	return s
}

func (s *RollbackCurrentProjectNodeRequest) SetNodeId(v int64) *RollbackCurrentProjectNodeRequest {
	s.NodeId = &v
	return s
}

func (s *RollbackCurrentProjectNodeRequest) SetRemark(v string) *RollbackCurrentProjectNodeRequest {
	s.Remark = &v
	return s
}

type RollbackCurrentProjectNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackCurrentProjectNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackCurrentProjectNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackCurrentProjectNodeResponseBody) SetRequestId(v string) *RollbackCurrentProjectNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackCurrentProjectNodeResponseBody) SetResult(v bool) *RollbackCurrentProjectNodeResponseBody {
	s.Result = &v
	return s
}

func (s *RollbackCurrentProjectNodeResponseBody) SetSuccess(v bool) *RollbackCurrentProjectNodeResponseBody {
	s.Success = &v
	return s
}

type RollbackCurrentProjectNodeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollbackCurrentProjectNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackCurrentProjectNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackCurrentProjectNodeResponse) GoString() string {
	return s.String()
}

func (s *RollbackCurrentProjectNodeResponse) SetHeaders(v map[string]*string) *RollbackCurrentProjectNodeResponse {
	s.Headers = v
	return s
}

func (s *RollbackCurrentProjectNodeResponse) SetStatusCode(v int32) *RollbackCurrentProjectNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackCurrentProjectNodeResponse) SetBody(v *RollbackCurrentProjectNodeResponseBody) *RollbackCurrentProjectNodeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":           tea.String("market.aliyuncs.com"),
		"ap-northeast-1":        tea.String("market.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("market.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("market.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("market.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("market.ap-southeast-1.aliyuncs.com"),
		"cn-beijing":            tea.String("market.aliyuncs.com"),
		"cn-chengdu":            tea.String("market.aliyuncs.com"),
		"cn-hongkong":           tea.String("market.aliyuncs.com"),
		"cn-huhehaote":          tea.String("market.aliyuncs.com"),
		"cn-qingdao":            tea.String("market.aliyuncs.com"),
		"cn-shanghai":           tea.String("market.aliyuncs.com"),
		"cn-shenzhen":           tea.String("market.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("market.aliyuncs.com"),
		"eu-central-1":          tea.String("market.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("market.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("market.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("market.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             tea.String("market.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("market.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("market.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("market.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("market.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("market"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActivateLicenseWithOptions(request *ActivateLicenseRequest, runtime *util.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identification)) {
		query["Identification"] = request.Identification
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseCode)) {
		query["LicenseCode"] = request.LicenseCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateLicense"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivateLicense(request *ActivateLicenseRequest) (_result *ActivateLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.ActivateLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AutoRenewInstanceWithOptions(request *AutoRenewInstanceRequest, runtime *util.RuntimeOptions) (_result *AutoRenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewCycle)) {
		body["AutoRenewCycle"] = request.AutoRenewCycle
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		body["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBizId)) {
		body["OrderBizId"] = request.OrderBizId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		body["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AutoRenewInstance"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AutoRenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AutoRenewInstance(request *AutoRenewInstanceRequest) (_result *AutoRenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AutoRenewInstanceResponse{}
	_body, _err := client.AutoRenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Commodity)) {
		query["Commodity"] = request.Commodity
	}

	if !tea.BoolValue(util.IsUnset(request.OrderSouce)) {
		query["OrderSouce"] = request.OrderSouce
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrder"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CrossAccountVerifyTokenWithOptions(request *CrossAccountVerifyTokenRequest, runtime *util.RuntimeOptions) (_result *CrossAccountVerifyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CrossAccountVerifyToken"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CrossAccountVerifyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CrossAccountVerifyToken(request *CrossAccountVerifyTokenRequest) (_result *CrossAccountVerifyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CrossAccountVerifyTokenResponse{}
	_body, _err := client.CrossAccountVerifyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiMeteringWithOptions(request *DescribeApiMeteringRequest, runtime *util.RuntimeOptions) (_result *DescribeApiMeteringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiMetering"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiMeteringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiMetering(request *DescribeApiMeteringRequest) (_result *DescribeApiMeteringResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiMeteringResponse{}
	_body, _err := client.DescribeApiMeteringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCurrentNodeInfoWithOptions(request *DescribeCurrentNodeInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeCurrentNodeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCurrentNodeInfo"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCurrentNodeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCurrentNodeInfo(request *DescribeCurrentNodeInfoRequest) (_result *DescribeCurrentNodeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCurrentNodeInfoResponse{}
	_body, _err := client.DescribeCurrentNodeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDistributionProductsWithOptions(request *DescribeDistributionProductsRequest, runtime *util.RuntimeOptions) (_result *DescribeDistributionProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDistributionProducts"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDistributionProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDistributionProducts(request *DescribeDistributionProductsRequest) (_result *DescribeDistributionProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDistributionProductsResponse{}
	_body, _err := client.DescribeDistributionProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDistributionProductsLinkWithOptions(tmpReq *DescribeDistributionProductsLinkRequest, runtime *util.RuntimeOptions) (_result *DescribeDistributionProductsLinkResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeDistributionProductsLinkShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Codes)) {
		request.CodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Codes, tea.String("Codes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodesShrink)) {
		query["Codes"] = request.CodesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDistributionProductsLink"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDistributionProductsLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDistributionProductsLink(request *DescribeDistributionProductsLinkRequest) (_result *DescribeDistributionProductsLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDistributionProductsLinkResponse{}
	_body, _err := client.DescribeDistributionProductsLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Codes)) {
		query["Codes"] = request.Codes
	}

	if !tea.BoolValue(util.IsUnset(request.ExceptCodes)) {
		query["ExceptCodes"] = request.ExceptCodes
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLicenseWithOptions(request *DescribeLicenseRequest, runtime *util.RuntimeOptions) (_result *DescribeLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LicenseCode)) {
		query["LicenseCode"] = request.LicenseCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLicense"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLicense(request *DescribeLicenseRequest) (_result *DescribeLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLicenseResponse{}
	_body, _err := client.DescribeLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOrderWithOptions(request *DescribeOrderRequest, runtime *util.RuntimeOptions) (_result *DescribeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOrder"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOrder(request *DescribeOrderRequest) (_result *DescribeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOrderResponse{}
	_body, _err := client.DescribeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Commodity)) {
		query["Commodity"] = request.Commodity
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePrice"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProductWithOptions(request *DescribeProductRequest, runtime *util.RuntimeOptions) (_result *DescribeProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDraft)) {
		query["QueryDraft"] = request.QueryDraft
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProduct"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProduct(request *DescribeProductRequest) (_result *DescribeProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProductResponse{}
	_body, _err := client.DescribeProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProductsWithOptions(request *DescribeProductsRequest, runtime *util.RuntimeOptions) (_result *DescribeProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchTerm)) {
		query["SearchTerm"] = request.SearchTerm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProducts"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProducts(request *DescribeProductsRequest) (_result *DescribeProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProductsResponse{}
	_body, _err := client.DescribeProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectAttachmentsWithOptions(request *DescribeProjectAttachmentsRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectAttachmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProjectAttachments"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectAttachmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProjectAttachments(request *DescribeProjectAttachmentsRequest) (_result *DescribeProjectAttachmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectAttachmentsResponse{}
	_body, _err := client.DescribeProjectAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectInfoWithOptions(request *DescribeProjectInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProjectInfo"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProjectInfo(request *DescribeProjectInfoRequest) (_result *DescribeProjectInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectInfoResponse{}
	_body, _err := client.DescribeProjectInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectMessagesWithOptions(request *DescribeProjectMessagesRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProjectMessages"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProjectMessages(request *DescribeProjectMessagesRequest) (_result *DescribeProjectMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectMessagesResponse{}
	_body, _err := client.DescribeProjectMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * **
 * **
 *
 * @param request DescribeProjectNodesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeProjectNodesResponse
 */
func (client *Client) DescribeProjectNodesWithOptions(request *DescribeProjectNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProjectNodes"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * **
 * **
 *
 * @param request DescribeProjectNodesRequest
 * @return DescribeProjectNodesResponse
 */
func (client *Client) DescribeProjectNodes(request *DescribeProjectNodesRequest) (_result *DescribeProjectNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectNodesResponse{}
	_body, _err := client.DescribeProjectNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectOperateLogsWithOptions(request *DescribeProjectOperateLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectOperateLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProjectOperateLogs"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectOperateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProjectOperateLogs(request *DescribeProjectOperateLogsRequest) (_result *DescribeProjectOperateLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectOperateLogsResponse{}
	_body, _err := client.DescribeProjectOperateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishCurrentProjectNodeWithOptions(request *FinishCurrentProjectNodeRequest, runtime *util.RuntimeOptions) (_result *FinishCurrentProjectNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateForm)) {
		query["TemplateForm"] = request.TemplateForm
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FinishCurrentProjectNode"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishCurrentProjectNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishCurrentProjectNode(request *FinishCurrentProjectNodeRequest) (_result *FinishCurrentProjectNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FinishCurrentProjectNodeResponse{}
	_body, _err := client.FinishCurrentProjectNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PauseProjectWithOptions(request *PauseProjectRequest, runtime *util.RuntimeOptions) (_result *PauseProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PauseProject"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PauseProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseProject(request *PauseProjectRequest) (_result *PauseProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseProjectResponse{}
	_body, _err := client.PauseProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushMeteringDataWithOptions(request *PushMeteringDataRequest, runtime *util.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Metering)) {
		query["Metering"] = request.Metering
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushMeteringData"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushMeteringData(request *PushMeteringDataRequest) (_result *PushMeteringDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.PushMeteringDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeProjectWithOptions(request *ResumeProjectRequest, runtime *util.RuntimeOptions) (_result *ResumeProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeProject"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeProject(request *ResumeProjectRequest) (_result *ResumeProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeProjectResponse{}
	_body, _err := client.ResumeProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackCurrentProjectNodeWithOptions(request *RollbackCurrentProjectNodeRequest, runtime *util.RuntimeOptions) (_result *RollbackCurrentProjectNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackCurrentProjectNode"),
		Version:     tea.String("2015-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackCurrentProjectNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackCurrentProjectNode(request *RollbackCurrentProjectNodeRequest) (_result *RollbackCurrentProjectNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackCurrentProjectNodeResponse{}
	_body, _err := client.RollbackCurrentProjectNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
