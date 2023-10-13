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

type CancelSubscriptionBillRequest struct {
	// The type of the bill to which you want to cancel the subscription. Valid values: PartnerBillingItemDetailForBillingPeriod, PartnerBillingItemDetailMonthly, PartnerInstanceDetailForBillingPeriod, and PartnerInstanceDetailMonthly.
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s CancelSubscriptionBillRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelSubscriptionBillRequest) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillRequest) SetSubscribeType(v string) *CancelSubscriptionBillRequest {
	s.SubscribeType = &v
	return s
}

type CancelSubscriptionBillResponseBody struct {
	// The HTTP status code that is returned.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelSubscriptionBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelSubscriptionBillResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillResponseBody) SetCode(v string) *CancelSubscriptionBillResponseBody {
	s.Code = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetData(v bool) *CancelSubscriptionBillResponseBody {
	s.Data = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetMessage(v string) *CancelSubscriptionBillResponseBody {
	s.Message = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetRequestId(v string) *CancelSubscriptionBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSubscriptionBillResponseBody) SetSuccess(v bool) *CancelSubscriptionBillResponseBody {
	s.Success = &v
	return s
}

type CancelSubscriptionBillResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelSubscriptionBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelSubscriptionBillResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelSubscriptionBillResponse) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillResponse) SetHeaders(v map[string]*string) *CancelSubscriptionBillResponse {
	s.Headers = v
	return s
}

func (s *CancelSubscriptionBillResponse) SetStatusCode(v int32) *CancelSubscriptionBillResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSubscriptionBillResponse) SetBody(v *CancelSubscriptionBillResponseBody) *CancelSubscriptionBillResponse {
	s.Body = v
	return s
}

type CreateCustomerRequest struct {
	CustomerName     *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	CustomerSource   *string `json:"CustomerSource,omitempty" xml:"CustomerSource,omitempty"`
	CustomerSubTrade *string `json:"CustomerSubTrade,omitempty" xml:"CustomerSubTrade,omitempty"`
	CustomerTrade    *string `json:"CustomerTrade,omitempty" xml:"CustomerTrade,omitempty"`
	Nation           *string `json:"Nation,omitempty" xml:"Nation,omitempty"`
}

func (s CreateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequest) SetCustomerName(v string) *CreateCustomerRequest {
	s.CustomerName = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerSource(v string) *CreateCustomerRequest {
	s.CustomerSource = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerSubTrade(v string) *CreateCustomerRequest {
	s.CustomerSubTrade = &v
	return s
}

func (s *CreateCustomerRequest) SetCustomerTrade(v string) *CreateCustomerRequest {
	s.CustomerTrade = &v
	return s
}

func (s *CreateCustomerRequest) SetNation(v string) *CreateCustomerRequest {
	s.Nation = &v
	return s
}

type CreateCustomerResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBody) SetCode(v string) *CreateCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomerResponseBody) SetData(v bool) *CreateCustomerResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCustomerResponseBody) SetMessage(v string) *CreateCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomerResponseBody) SetRequestId(v string) *CreateCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerResponseBody) SetSuccess(v bool) *CreateCustomerResponseBody {
	s.Success = &v
	return s
}

type CreateCustomerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponse) SetHeaders(v map[string]*string) *CreateCustomerResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerResponse) SetStatusCode(v int32) *CreateCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerResponse) SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse {
	s.Body = v
	return s
}

type DeductOutstandingBalanceRequest struct {
	DeductAmount *string `json:"DeductAmount,omitempty" xml:"DeductAmount,omitempty"`
	Uid          *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DeductOutstandingBalanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeductOutstandingBalanceRequest) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceRequest) SetDeductAmount(v string) *DeductOutstandingBalanceRequest {
	s.DeductAmount = &v
	return s
}

func (s *DeductOutstandingBalanceRequest) SetUid(v int64) *DeductOutstandingBalanceRequest {
	s.Uid = &v
	return s
}

type DeductOutstandingBalanceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeductOutstandingBalanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeductOutstandingBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceResponseBody) SetCode(v string) *DeductOutstandingBalanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetMessage(v string) *DeductOutstandingBalanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetRequestId(v string) *DeductOutstandingBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetSuccess(v bool) *DeductOutstandingBalanceResponseBody {
	s.Success = &v
	return s
}

type DeductOutstandingBalanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeductOutstandingBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeductOutstandingBalanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeductOutstandingBalanceResponse) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceResponse) SetHeaders(v map[string]*string) *DeductOutstandingBalanceResponse {
	s.Headers = v
	return s
}

func (s *DeductOutstandingBalanceResponse) SetStatusCode(v int32) *DeductOutstandingBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeductOutstandingBalanceResponse) SetBody(v *DeductOutstandingBalanceResponseBody) *DeductOutstandingBalanceResponse {
	s.Body = v
	return s
}

type EditEndUserStatusRequest struct {
	CreditStatus *string `json:"CreditStatus,omitempty" xml:"CreditStatus,omitempty"`
	// uid
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditEndUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s EditEndUserStatusRequest) GoString() string {
	return s.String()
}

func (s *EditEndUserStatusRequest) SetCreditStatus(v string) *EditEndUserStatusRequest {
	s.CreditStatus = &v
	return s
}

func (s *EditEndUserStatusRequest) SetUid(v int64) *EditEndUserStatusRequest {
	s.Uid = &v
	return s
}

type EditEndUserStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditEndUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditEndUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *EditEndUserStatusResponseBody) SetCode(v string) *EditEndUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *EditEndUserStatusResponseBody) SetData(v string) *EditEndUserStatusResponseBody {
	s.Data = &v
	return s
}

func (s *EditEndUserStatusResponseBody) SetMessage(v string) *EditEndUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *EditEndUserStatusResponseBody) SetMsg(v string) *EditEndUserStatusResponseBody {
	s.Msg = &v
	return s
}

func (s *EditEndUserStatusResponseBody) SetRequestId(v string) *EditEndUserStatusResponseBody {
	s.RequestId = &v
	return s
}

type EditEndUserStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EditEndUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditEndUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s EditEndUserStatusResponse) GoString() string {
	return s.String()
}

func (s *EditEndUserStatusResponse) SetHeaders(v map[string]*string) *EditEndUserStatusResponse {
	s.Headers = v
	return s
}

func (s *EditEndUserStatusResponse) SetStatusCode(v int32) *EditEndUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *EditEndUserStatusResponse) SetBody(v *EditEndUserStatusResponseBody) *EditEndUserStatusResponse {
	s.Body = v
	return s
}

type EditNewBuyStatusRequest struct {
	NewBuyStatus *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	Uid          *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditNewBuyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s EditNewBuyStatusRequest) GoString() string {
	return s.String()
}

func (s *EditNewBuyStatusRequest) SetNewBuyStatus(v string) *EditNewBuyStatusRequest {
	s.NewBuyStatus = &v
	return s
}

func (s *EditNewBuyStatusRequest) SetUid(v int64) *EditNewBuyStatusRequest {
	s.Uid = &v
	return s
}

type EditNewBuyStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditNewBuyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditNewBuyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *EditNewBuyStatusResponseBody) SetCode(v string) *EditNewBuyStatusResponseBody {
	s.Code = &v
	return s
}

func (s *EditNewBuyStatusResponseBody) SetData(v string) *EditNewBuyStatusResponseBody {
	s.Data = &v
	return s
}

func (s *EditNewBuyStatusResponseBody) SetMessage(v string) *EditNewBuyStatusResponseBody {
	s.Message = &v
	return s
}

func (s *EditNewBuyStatusResponseBody) SetMsg(v string) *EditNewBuyStatusResponseBody {
	s.Msg = &v
	return s
}

func (s *EditNewBuyStatusResponseBody) SetRequestId(v string) *EditNewBuyStatusResponseBody {
	s.RequestId = &v
	return s
}

type EditNewBuyStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EditNewBuyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditNewBuyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s EditNewBuyStatusResponse) GoString() string {
	return s.String()
}

func (s *EditNewBuyStatusResponse) SetHeaders(v map[string]*string) *EditNewBuyStatusResponse {
	s.Headers = v
	return s
}

func (s *EditNewBuyStatusResponse) SetStatusCode(v int32) *EditNewBuyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *EditNewBuyStatusResponse) SetBody(v *EditNewBuyStatusResponseBody) *EditNewBuyStatusResponse {
	s.Body = v
	return s
}

type EditZeroCreditShutdownRequest struct {
	ShutdownPolicy *string `json:"ShutdownPolicy,omitempty" xml:"ShutdownPolicy,omitempty"`
	// uid
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditZeroCreditShutdownRequest) String() string {
	return tea.Prettify(s)
}

func (s EditZeroCreditShutdownRequest) GoString() string {
	return s.String()
}

func (s *EditZeroCreditShutdownRequest) SetShutdownPolicy(v string) *EditZeroCreditShutdownRequest {
	s.ShutdownPolicy = &v
	return s
}

func (s *EditZeroCreditShutdownRequest) SetUid(v int64) *EditZeroCreditShutdownRequest {
	s.Uid = &v
	return s
}

type EditZeroCreditShutdownResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditZeroCreditShutdownResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditZeroCreditShutdownResponseBody) GoString() string {
	return s.String()
}

func (s *EditZeroCreditShutdownResponseBody) SetCode(v string) *EditZeroCreditShutdownResponseBody {
	s.Code = &v
	return s
}

func (s *EditZeroCreditShutdownResponseBody) SetData(v string) *EditZeroCreditShutdownResponseBody {
	s.Data = &v
	return s
}

func (s *EditZeroCreditShutdownResponseBody) SetMessage(v string) *EditZeroCreditShutdownResponseBody {
	s.Message = &v
	return s
}

func (s *EditZeroCreditShutdownResponseBody) SetMsg(v string) *EditZeroCreditShutdownResponseBody {
	s.Msg = &v
	return s
}

func (s *EditZeroCreditShutdownResponseBody) SetRequestId(v string) *EditZeroCreditShutdownResponseBody {
	s.RequestId = &v
	return s
}

type EditZeroCreditShutdownResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EditZeroCreditShutdownResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditZeroCreditShutdownResponse) String() string {
	return tea.Prettify(s)
}

func (s EditZeroCreditShutdownResponse) GoString() string {
	return s.String()
}

func (s *EditZeroCreditShutdownResponse) SetHeaders(v map[string]*string) *EditZeroCreditShutdownResponse {
	s.Headers = v
	return s
}

func (s *EditZeroCreditShutdownResponse) SetStatusCode(v int32) *EditZeroCreditShutdownResponse {
	s.StatusCode = &v
	return s
}

func (s *EditZeroCreditShutdownResponse) SetBody(v *EditZeroCreditShutdownResponseBody) *EditZeroCreditShutdownResponse {
	s.Body = v
	return s
}

type GetAccountInfoRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Uid         *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserType    *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAccountInfoRequest) SetCurrentPage(v int32) *GetAccountInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAccountInfoRequest) SetPageSize(v int32) *GetAccountInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccountInfoRequest) SetUid(v int64) *GetAccountInfoRequest {
	s.Uid = &v
	return s
}

func (s *GetAccountInfoRequest) SetUserType(v string) *GetAccountInfoRequest {
	s.UserType = &v
	return s
}

type GetAccountInfoResponseBody struct {
	AccountInfoList *GetAccountInfoResponseBodyAccountInfoList `json:"AccountInfoList,omitempty" xml:"AccountInfoList,omitempty" type:"Struct"`
	Code            *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageInfo  *GetAccountInfoResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBody) SetAccountInfoList(v *GetAccountInfoResponseBodyAccountInfoList) *GetAccountInfoResponseBody {
	s.AccountInfoList = v
	return s
}

func (s *GetAccountInfoResponseBody) SetCode(v string) *GetAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetMessage(v string) *GetAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetPageInfo(v *GetAccountInfoResponseBodyPageInfo) *GetAccountInfoResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetAccountInfoResponseBody) SetRequestId(v string) *GetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountInfoResponseBody) SetSuccess(v bool) *GetAccountInfoResponseBody {
	s.Success = &v
	return s
}

type GetAccountInfoResponseBodyAccountInfoList struct {
	AccountInfo []*GetAccountInfoResponseBodyAccountInfoListAccountInfo `json:"AccountInfo,omitempty" xml:"AccountInfo,omitempty" type:"Repeated"`
}

func (s GetAccountInfoResponseBodyAccountInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBodyAccountInfoList) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyAccountInfoList) SetAccountInfo(v []*GetAccountInfoResponseBodyAccountInfoListAccountInfo) *GetAccountInfoResponseBodyAccountInfoList {
	s.AccountInfo = v
	return s
}

type GetAccountInfoResponseBodyAccountInfoListAccountInfo struct {
	AccountNickname        *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	AliyunId               *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	AssociationSuccessTime *string `json:"AssociationSuccessTime,omitempty" xml:"AssociationSuccessTime,omitempty"`
	Cid                    *int64  `json:"Cid,omitempty" xml:"Cid,omitempty"`
	Email                  *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile                 *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Remark                 *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SubAccountType         *int32  `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	Uid                    *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetAccountInfoResponseBodyAccountInfoListAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBodyAccountInfoListAccountInfo) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetAccountNickname(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.AccountNickname = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetAliyunId(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.AliyunId = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetAssociationSuccessTime(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.AssociationSuccessTime = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetCid(v int64) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Cid = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetEmail(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Email = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetMobile(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Mobile = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetRemark(v string) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Remark = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetSubAccountType(v int32) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.SubAccountType = &v
	return s
}

func (s *GetAccountInfoResponseBodyAccountInfoListAccountInfo) SetUid(v int64) *GetAccountInfoResponseBodyAccountInfoListAccountInfo {
	s.Uid = &v
	return s
}

type GetAccountInfoResponseBodyPageInfo struct {
	Page     *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAccountInfoResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponseBodyPageInfo) SetPage(v int32) *GetAccountInfoResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *GetAccountInfoResponseBodyPageInfo) SetPageSize(v int32) *GetAccountInfoResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetAccountInfoResponseBodyPageInfo) SetTotal(v int32) *GetAccountInfoResponseBodyPageInfo {
	s.Total = &v
	return s
}

type GetAccountInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponse) SetHeaders(v map[string]*string) *GetAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccountInfoResponse) SetStatusCode(v int32) *GetAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountInfoResponse) SetBody(v *GetAccountInfoResponseBody) *GetAccountInfoResponse {
	s.Body = v
	return s
}

type GetCreditInfoRequest struct {
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetCreditInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreditInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCreditInfoRequest) SetUid(v int64) *GetCreditInfoRequest {
	s.Uid = &v
	return s
}

type GetCreditInfoResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCreditInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCreditInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCreditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponseBody) SetCode(v string) *GetCreditInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetData(v *GetCreditInfoResponseBodyData) *GetCreditInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCreditInfoResponseBody) SetMessage(v string) *GetCreditInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetRequestId(v string) *GetCreditInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreditInfoResponseBody) SetSuccess(v bool) *GetCreditInfoResponseBody {
	s.Success = &v
	return s
}

type GetCreditInfoResponseBodyData struct {
	AccountStatus            *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AlarmThreshold           *string `json:"AlarmThreshold,omitempty" xml:"AlarmThreshold,omitempty"`
	AvailableCredit          *string `json:"AvailableCredit,omitempty" xml:"AvailableCredit,omitempty"`
	CreditLine               *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	OutstandingBalance       *string `json:"OutstandingBalance,omitempty" xml:"OutstandingBalance,omitempty"`
	ZeroCreditShutdownPolicy *string `json:"ZeroCreditShutdownPolicy,omitempty" xml:"ZeroCreditShutdownPolicy,omitempty"`
	NewBuyStatus             *string `json:"newBuyStatus,omitempty" xml:"newBuyStatus,omitempty"`
}

func (s GetCreditInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCreditInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponseBodyData) SetAccountStatus(v string) *GetCreditInfoResponseBodyData {
	s.AccountStatus = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetAlarmThreshold(v string) *GetCreditInfoResponseBodyData {
	s.AlarmThreshold = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetAvailableCredit(v string) *GetCreditInfoResponseBodyData {
	s.AvailableCredit = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetCreditLine(v string) *GetCreditInfoResponseBodyData {
	s.CreditLine = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetOutstandingBalance(v string) *GetCreditInfoResponseBodyData {
	s.OutstandingBalance = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetZeroCreditShutdownPolicy(v string) *GetCreditInfoResponseBodyData {
	s.ZeroCreditShutdownPolicy = &v
	return s
}

func (s *GetCreditInfoResponseBodyData) SetNewBuyStatus(v string) *GetCreditInfoResponseBodyData {
	s.NewBuyStatus = &v
	return s
}

type GetCreditInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCreditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCreditInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreditInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponse) SetHeaders(v map[string]*string) *GetCreditInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCreditInfoResponse) SetStatusCode(v int32) *GetCreditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreditInfoResponse) SetBody(v *GetCreditInfoResponseBody) *GetCreditInfoResponse {
	s.Body = v
	return s
}

type GetDailyBillRequest struct {
	BillOwner *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	BillType  *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// Billing date. Format YYYY-MM-DD
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetDailyBillRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDailyBillRequest) GoString() string {
	return s.String()
}

func (s *GetDailyBillRequest) SetBillOwner(v string) *GetDailyBillRequest {
	s.BillOwner = &v
	return s
}

func (s *GetDailyBillRequest) SetBillType(v string) *GetDailyBillRequest {
	s.BillType = &v
	return s
}

func (s *GetDailyBillRequest) SetDate(v string) *GetDailyBillRequest {
	s.Date = &v
	return s
}

type GetDailyBillResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDailyBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDailyBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDailyBillResponseBody) GoString() string {
	return s.String()
}

func (s *GetDailyBillResponseBody) SetCode(v string) *GetDailyBillResponseBody {
	s.Code = &v
	return s
}

func (s *GetDailyBillResponseBody) SetData(v *GetDailyBillResponseBodyData) *GetDailyBillResponseBody {
	s.Data = v
	return s
}

func (s *GetDailyBillResponseBody) SetMessage(v string) *GetDailyBillResponseBody {
	s.Message = &v
	return s
}

func (s *GetDailyBillResponseBody) SetRequestId(v string) *GetDailyBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDailyBillResponseBody) SetSuccess(v bool) *GetDailyBillResponseBody {
	s.Success = &v
	return s
}

type GetDailyBillResponseBodyData struct {
	BillLinkCSV  *string `json:"BillLinkCSV,omitempty" xml:"BillLinkCSV,omitempty"`
	BillLinkXLSX *string `json:"BillLinkXLSX,omitempty" xml:"BillLinkXLSX,omitempty"`
	BillOwner    *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	BillType     *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	SpendingTime *string `json:"SpendingTime,omitempty" xml:"SpendingTime,omitempty"`
}

func (s GetDailyBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDailyBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDailyBillResponseBodyData) SetBillLinkCSV(v string) *GetDailyBillResponseBodyData {
	s.BillLinkCSV = &v
	return s
}

func (s *GetDailyBillResponseBodyData) SetBillLinkXLSX(v string) *GetDailyBillResponseBodyData {
	s.BillLinkXLSX = &v
	return s
}

func (s *GetDailyBillResponseBodyData) SetBillOwner(v string) *GetDailyBillResponseBodyData {
	s.BillOwner = &v
	return s
}

func (s *GetDailyBillResponseBodyData) SetBillType(v string) *GetDailyBillResponseBodyData {
	s.BillType = &v
	return s
}

func (s *GetDailyBillResponseBodyData) SetSpendingTime(v string) *GetDailyBillResponseBodyData {
	s.SpendingTime = &v
	return s
}

type GetDailyBillResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDailyBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDailyBillResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDailyBillResponse) GoString() string {
	return s.String()
}

func (s *GetDailyBillResponse) SetHeaders(v map[string]*string) *GetDailyBillResponse {
	s.Headers = v
	return s
}

func (s *GetDailyBillResponse) SetStatusCode(v int32) *GetDailyBillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDailyBillResponse) SetBody(v *GetDailyBillResponseBody) *GetDailyBillResponse {
	s.Body = v
	return s
}

type GetInviteStatusRequest struct {
	// inviteId list
	InviteStatusList []*GetInviteStatusRequestInviteStatusList `json:"InviteStatusList,omitempty" xml:"InviteStatusList,omitempty" type:"Repeated"`
}

func (s GetInviteStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusRequest) GoString() string {
	return s.String()
}

func (s *GetInviteStatusRequest) SetInviteStatusList(v []*GetInviteStatusRequestInviteStatusList) *GetInviteStatusRequest {
	s.InviteStatusList = v
	return s
}

type GetInviteStatusRequestInviteStatusList struct {
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
}

func (s GetInviteStatusRequestInviteStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusRequestInviteStatusList) GoString() string {
	return s.String()
}

func (s *GetInviteStatusRequestInviteStatusList) SetInviteId(v int64) *GetInviteStatusRequestInviteStatusList {
	s.InviteId = &v
	return s
}

type GetInviteStatusResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetInviteStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInviteStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBody) SetCode(v string) *GetInviteStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetData(v *GetInviteStatusResponseBodyData) *GetInviteStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetInviteStatusResponseBody) SetMessage(v string) *GetInviteStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetRequestId(v string) *GetInviteStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetSuccess(v bool) *GetInviteStatusResponseBody {
	s.Success = &v
	return s
}

type GetInviteStatusResponseBodyData struct {
	InviteStatus []*GetInviteStatusResponseBodyDataInviteStatus `json:"InviteStatus,omitempty" xml:"InviteStatus,omitempty" type:"Repeated"`
}

func (s GetInviteStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyData) SetInviteStatus(v []*GetInviteStatusResponseBodyDataInviteStatus) *GetInviteStatusResponseBodyData {
	s.InviteStatus = v
	return s
}

type GetInviteStatusResponseBodyDataInviteStatus struct {
	Code             *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	InviteStatusList *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList `json:"InviteStatusList,omitempty" xml:"InviteStatusList,omitempty" type:"Struct"`
	Message          *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Success          *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInviteStatusResponseBodyDataInviteStatus) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponseBodyDataInviteStatus) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetCode(v string) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Code = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetInviteStatusList(v *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) *GetInviteStatusResponseBodyDataInviteStatus {
	s.InviteStatusList = v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetMessage(v string) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Message = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetSuccess(v bool) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Success = &v
	return s
}

type GetInviteStatusResponseBodyDataInviteStatusInviteStatusList struct {
	AssociationSuccessTime *string `json:"AssociationSuccessTime,omitempty" xml:"AssociationSuccessTime,omitempty"`
	Cid                    *int64  `json:"Cid,omitempty" xml:"Cid,omitempty"`
	GmtCreate              *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ParentId               *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Status                 *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubAccountType         *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	Uid                    *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetAssociationSuccessTime(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.AssociationSuccessTime = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetCid(v int64) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Cid = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetGmtCreate(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.GmtCreate = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetParentId(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.ParentId = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetStatus(v int32) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Status = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetSubAccountType(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.SubAccountType = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetUid(v int64) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Uid = &v
	return s
}

type GetInviteStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInviteStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInviteStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInviteStatusResponse) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponse) SetHeaders(v map[string]*string) *GetInviteStatusResponse {
	s.Headers = v
	return s
}

func (s *GetInviteStatusResponse) SetStatusCode(v int32) *GetInviteStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInviteStatusResponse) SetBody(v *GetInviteStatusResponseBody) *GetInviteStatusResponse {
	s.Body = v
	return s
}

type GetMonthlyBillRequest struct {
	// Bill Owner type.
	//
	//  Value range:
	//
	// 1: Master account
	//
	// 2: Sub account
	BillOwner *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	// Value Range:
	//
	// MonthlyInvoice
	//
	// MonthRefundInvoice
	//
	// MonthlySummary
	//
	// MonthlyInstanceAddAdjustBill
	//
	// MonthlyInstanceRefundBill
	//
	// MonthlyAddAdjustInvoce
	//
	// MonthlyRefundAdjustInvoce
	//
	// MonthlyInstanceConsumeV2
	//
	// MarginReportV2
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// Billing Month, Format is YYYY-MM
	Month *string `json:"Month,omitempty" xml:"Month,omitempty"`
}

func (s GetMonthlyBillRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMonthlyBillRequest) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillRequest) SetBillOwner(v string) *GetMonthlyBillRequest {
	s.BillOwner = &v
	return s
}

func (s *GetMonthlyBillRequest) SetBillType(v string) *GetMonthlyBillRequest {
	s.BillType = &v
	return s
}

func (s *GetMonthlyBillRequest) SetMonth(v string) *GetMonthlyBillRequest {
	s.Month = &v
	return s
}

type GetMonthlyBillResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMonthlyBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMonthlyBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMonthlyBillResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillResponseBody) SetCode(v string) *GetMonthlyBillResponseBody {
	s.Code = &v
	return s
}

func (s *GetMonthlyBillResponseBody) SetData(v *GetMonthlyBillResponseBodyData) *GetMonthlyBillResponseBody {
	s.Data = v
	return s
}

func (s *GetMonthlyBillResponseBody) SetMessage(v string) *GetMonthlyBillResponseBody {
	s.Message = &v
	return s
}

func (s *GetMonthlyBillResponseBody) SetRequestId(v string) *GetMonthlyBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonthlyBillResponseBody) SetSuccess(v bool) *GetMonthlyBillResponseBody {
	s.Success = &v
	return s
}

type GetMonthlyBillResponseBodyData struct {
	BillLinkCSV       *string `json:"BillLinkCSV,omitempty" xml:"BillLinkCSV,omitempty"`
	BillLinkXLSX      *string `json:"BillLinkXLSX,omitempty" xml:"BillLinkXLSX,omitempty"`
	BillOwner         *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	BillType          *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	InvoiceLink       *string `json:"InvoiceLink,omitempty" xml:"InvoiceLink,omitempty"`
	RefundInvoiceFlag *bool   `json:"RefundInvoiceFlag,omitempty" xml:"RefundInvoiceFlag,omitempty"`
	RefundInvoiceLink *string `json:"RefundInvoiceLink,omitempty" xml:"RefundInvoiceLink,omitempty"`
	SpendingTime      *string `json:"SpendingTime,omitempty" xml:"SpendingTime,omitempty"`
}

func (s GetMonthlyBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMonthlyBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillResponseBodyData) SetBillLinkCSV(v string) *GetMonthlyBillResponseBodyData {
	s.BillLinkCSV = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetBillLinkXLSX(v string) *GetMonthlyBillResponseBodyData {
	s.BillLinkXLSX = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetBillOwner(v string) *GetMonthlyBillResponseBodyData {
	s.BillOwner = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetBillType(v string) *GetMonthlyBillResponseBodyData {
	s.BillType = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetInvoiceLink(v string) *GetMonthlyBillResponseBodyData {
	s.InvoiceLink = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetRefundInvoiceFlag(v bool) *GetMonthlyBillResponseBodyData {
	s.RefundInvoiceFlag = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetRefundInvoiceLink(v string) *GetMonthlyBillResponseBodyData {
	s.RefundInvoiceLink = &v
	return s
}

func (s *GetMonthlyBillResponseBodyData) SetSpendingTime(v string) *GetMonthlyBillResponseBodyData {
	s.SpendingTime = &v
	return s
}

type GetMonthlyBillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMonthlyBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMonthlyBillResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonthlyBillResponse) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillResponse) SetHeaders(v map[string]*string) *GetMonthlyBillResponse {
	s.Headers = v
	return s
}

func (s *GetMonthlyBillResponse) SetStatusCode(v int32) *GetMonthlyBillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonthlyBillResponse) SetBody(v *GetMonthlyBillResponseBody) *GetMonthlyBillResponse {
	s.Body = v
	return s
}

type GetUnassociatedCustomerRequest struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetUnassociatedCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerRequest) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerRequest) SetCurrentPage(v int32) *GetUnassociatedCustomerRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetUnassociatedCustomerRequest) SetPageSize(v int32) *GetUnassociatedCustomerRequest {
	s.PageSize = &v
	return s
}

type GetUnassociatedCustomerResponseBody struct {
	Code           *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	InviteInfoList *GetUnassociatedCustomerResponseBodyInviteInfoList `json:"InviteInfoList,omitempty" xml:"InviteInfoList,omitempty" type:"Struct"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageInfo       *GetUnassociatedCustomerResponseBodyPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUnassociatedCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBody) SetCode(v string) *GetUnassociatedCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetInviteInfoList(v *GetUnassociatedCustomerResponseBodyInviteInfoList) *GetUnassociatedCustomerResponseBody {
	s.InviteInfoList = v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetMessage(v string) *GetUnassociatedCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetPageInfo(v *GetUnassociatedCustomerResponseBodyPageInfo) *GetUnassociatedCustomerResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetRequestId(v string) *GetUnassociatedCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetSuccess(v bool) *GetUnassociatedCustomerResponseBody {
	s.Success = &v
	return s
}

type GetUnassociatedCustomerResponseBodyInviteInfoList struct {
	InviteInfo []*GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo `json:"InviteInfo,omitempty" xml:"InviteInfo,omitempty" type:"Repeated"`
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoList) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoList) SetInviteInfo(v []*GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) *GetUnassociatedCustomerResponseBodyInviteInfoList {
	s.InviteInfo = v
	return s
}

type GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo struct {
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InviteId        *int64  `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetAccountNickname(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.AccountNickname = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetEmail(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.Email = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetGmtCreate(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetInviteId(v int64) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.InviteId = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetStatus(v int32) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.Status = &v
	return s
}

type GetUnassociatedCustomerResponseBodyPageInfo struct {
	Page     *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetUnassociatedCustomerResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetPage(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetPageSize(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetTotal(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.Total = &v
	return s
}

type GetUnassociatedCustomerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUnassociatedCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUnassociatedCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUnassociatedCustomerResponse) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponse) SetHeaders(v map[string]*string) *GetUnassociatedCustomerResponse {
	s.Headers = v
	return s
}

func (s *GetUnassociatedCustomerResponse) SetStatusCode(v int32) *GetUnassociatedCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnassociatedCustomerResponse) SetBody(v *GetUnassociatedCustomerResponseBody) *GetUnassociatedCustomerResponse {
	s.Body = v
	return s
}

type InviteSubAccountRequest struct {
	AccountInfoList []*InviteSubAccountRequestAccountInfoList `json:"AccountInfoList,omitempty" xml:"AccountInfoList,omitempty" type:"Repeated"`
}

func (s InviteSubAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountRequest) GoString() string {
	return s.String()
}

func (s *InviteSubAccountRequest) SetAccountInfoList(v []*InviteSubAccountRequestAccountInfoList) *InviteSubAccountRequest {
	s.AccountInfoList = v
	return s
}

type InviteSubAccountRequestAccountInfoList struct {
	AccountNickname          *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	CreditLine               *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	CustomerId               *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	EmailAddress             *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	NewBuyStatus             *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
	Remark                   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SubAccountType           *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	ZeroCreditShutdownPolicy *string `json:"ZeroCreditShutdownPolicy,omitempty" xml:"ZeroCreditShutdownPolicy,omitempty"`
}

func (s InviteSubAccountRequestAccountInfoList) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountRequestAccountInfoList) GoString() string {
	return s.String()
}

func (s *InviteSubAccountRequestAccountInfoList) SetAccountNickname(v string) *InviteSubAccountRequestAccountInfoList {
	s.AccountNickname = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetCreditLine(v string) *InviteSubAccountRequestAccountInfoList {
	s.CreditLine = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetCustomerId(v string) *InviteSubAccountRequestAccountInfoList {
	s.CustomerId = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetEmailAddress(v string) *InviteSubAccountRequestAccountInfoList {
	s.EmailAddress = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetNewBuyStatus(v string) *InviteSubAccountRequestAccountInfoList {
	s.NewBuyStatus = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetRemark(v string) *InviteSubAccountRequestAccountInfoList {
	s.Remark = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetSubAccountType(v string) *InviteSubAccountRequestAccountInfoList {
	s.SubAccountType = &v
	return s
}

func (s *InviteSubAccountRequestAccountInfoList) SetZeroCreditShutdownPolicy(v string) *InviteSubAccountRequestAccountInfoList {
	s.ZeroCreditShutdownPolicy = &v
	return s
}

type InviteSubAccountResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *InviteSubAccountResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteSubAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponseBody) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBody) SetCode(v string) *InviteSubAccountResponseBody {
	s.Code = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetMessage(v string) *InviteSubAccountResponseBody {
	s.Message = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetRequestId(v string) *InviteSubAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *InviteSubAccountResponseBody) SetResults(v *InviteSubAccountResponseBodyResults) *InviteSubAccountResponseBody {
	s.Results = v
	return s
}

func (s *InviteSubAccountResponseBody) SetSuccess(v bool) *InviteSubAccountResponseBody {
	s.Success = &v
	return s
}

type InviteSubAccountResponseBodyResults struct {
	Result []*InviteSubAccountResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s InviteSubAccountResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponseBodyResults) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResults) SetResult(v []*InviteSubAccountResponseBodyResultsResult) *InviteSubAccountResponseBodyResults {
	s.Result = v
	return s
}

type InviteSubAccountResponseBodyResultsResult struct {
	Code    *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Result  *InviteSubAccountResponseBodyResultsResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InviteSubAccountResponseBodyResultsResult) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResultsResult) SetCode(v string) *InviteSubAccountResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetMessage(v string) *InviteSubAccountResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetResult(v *InviteSubAccountResponseBodyResultsResultResult) *InviteSubAccountResponseBodyResultsResult {
	s.Result = v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResult) SetSuccess(v bool) *InviteSubAccountResponseBodyResultsResult {
	s.Success = &v
	return s
}

type InviteSubAccountResponseBodyResultsResultResult struct {
	Days     *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	InviteId *int64  `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
	RegUrl   *string `json:"RegUrl,omitempty" xml:"RegUrl,omitempty"`
}

func (s InviteSubAccountResponseBodyResultsResultResult) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponseBodyResultsResultResult) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetDays(v int32) *InviteSubAccountResponseBodyResultsResultResult {
	s.Days = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetInviteId(v int64) *InviteSubAccountResponseBodyResultsResultResult {
	s.InviteId = &v
	return s
}

func (s *InviteSubAccountResponseBodyResultsResultResult) SetRegUrl(v string) *InviteSubAccountResponseBodyResultsResultResult {
	s.RegUrl = &v
	return s
}

type InviteSubAccountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InviteSubAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InviteSubAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s InviteSubAccountResponse) GoString() string {
	return s.String()
}

func (s *InviteSubAccountResponse) SetHeaders(v map[string]*string) *InviteSubAccountResponse {
	s.Headers = v
	return s
}

func (s *InviteSubAccountResponse) SetStatusCode(v int32) *InviteSubAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteSubAccountResponse) SetBody(v *InviteSubAccountResponseBody) *InviteSubAccountResponse {
	s.Body = v
	return s
}

type ListCountriesResponseBody struct {
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCountriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCountriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCountriesResponseBody) SetCode(v string) *ListCountriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCountriesResponseBody) SetData(v []*string) *ListCountriesResponseBody {
	s.Data = v
	return s
}

func (s *ListCountriesResponseBody) SetMessage(v string) *ListCountriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCountriesResponseBody) SetRequestId(v string) *ListCountriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCountriesResponseBody) SetSuccess(v bool) *ListCountriesResponseBody {
	s.Success = &v
	return s
}

type ListCountriesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCountriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCountriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCountriesResponse) GoString() string {
	return s.String()
}

func (s *ListCountriesResponse) SetHeaders(v map[string]*string) *ListCountriesResponse {
	s.Headers = v
	return s
}

func (s *ListCountriesResponse) SetStatusCode(v int32) *ListCountriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCountriesResponse) SetBody(v *ListCountriesResponseBody) *ListCountriesResponse {
	s.Body = v
	return s
}

type ResendEmailRequest struct {
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
}

func (s ResendEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailRequest) GoString() string {
	return s.String()
}

func (s *ResendEmailRequest) SetInviteId(v int64) *ResendEmailRequest {
	s.InviteId = &v
	return s
}

type ResendEmailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResendEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendEmailResponseBody) SetCode(v string) *ResendEmailResponseBody {
	s.Code = &v
	return s
}

func (s *ResendEmailResponseBody) SetMessage(v string) *ResendEmailResponseBody {
	s.Message = &v
	return s
}

func (s *ResendEmailResponseBody) SetRequestId(v string) *ResendEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendEmailResponseBody) SetSuccess(v bool) *ResendEmailResponseBody {
	s.Success = &v
	return s
}

type ResendEmailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResendEmailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResendEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s ResendEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendEmailResponse) SetHeaders(v map[string]*string) *ResendEmailResponse {
	s.Headers = v
	return s
}

func (s *ResendEmailResponse) SetStatusCode(v int32) *ResendEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendEmailResponse) SetBody(v *ResendEmailResponseBody) *ResendEmailResponse {
	s.Body = v
	return s
}

type SetAccountInfoRequest struct {
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Uid             *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SetAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *SetAccountInfoRequest) SetAccountNickname(v string) *SetAccountInfoRequest {
	s.AccountNickname = &v
	return s
}

func (s *SetAccountInfoRequest) SetRemark(v string) *SetAccountInfoRequest {
	s.Remark = &v
	return s
}

func (s *SetAccountInfoRequest) SetUid(v int64) *SetAccountInfoRequest {
	s.Uid = &v
	return s
}

type SetAccountInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccountInfoResponseBody) SetCode(v string) *SetAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetMessage(v string) *SetAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetRequestId(v string) *SetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetSuccess(v bool) *SetAccountInfoResponseBody {
	s.Success = &v
	return s
}

type SetAccountInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *SetAccountInfoResponse) SetHeaders(v map[string]*string) *SetAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *SetAccountInfoResponse) SetStatusCode(v int32) *SetAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccountInfoResponse) SetBody(v *SetAccountInfoResponseBody) *SetAccountInfoResponse {
	s.Body = v
	return s
}

type SetCreditLineRequest struct {
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	Uid        *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SetCreditLineRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLineRequest) GoString() string {
	return s.String()
}

func (s *SetCreditLineRequest) SetCreditLine(v string) *SetCreditLineRequest {
	s.CreditLine = &v
	return s
}

func (s *SetCreditLineRequest) SetUid(v int64) *SetCreditLineRequest {
	s.Uid = &v
	return s
}

type SetCreditLineResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetCreditLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLineResponseBody) GoString() string {
	return s.String()
}

func (s *SetCreditLineResponseBody) SetCode(v string) *SetCreditLineResponseBody {
	s.Code = &v
	return s
}

func (s *SetCreditLineResponseBody) SetMessage(v string) *SetCreditLineResponseBody {
	s.Message = &v
	return s
}

func (s *SetCreditLineResponseBody) SetRequestId(v string) *SetCreditLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCreditLineResponseBody) SetSuccess(v bool) *SetCreditLineResponseBody {
	s.Success = &v
	return s
}

type SetCreditLineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetCreditLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCreditLineResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLineResponse) GoString() string {
	return s.String()
}

func (s *SetCreditLineResponse) SetHeaders(v map[string]*string) *SetCreditLineResponse {
	s.Headers = v
	return s
}

func (s *SetCreditLineResponse) SetStatusCode(v int32) *SetCreditLineResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCreditLineResponse) SetBody(v *SetCreditLineResponseBody) *SetCreditLineResponse {
	s.Body = v
	return s
}

type SetWarningThresholdRequest struct {
	Uid          *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	WarningValue *string `json:"WarningValue,omitempty" xml:"WarningValue,omitempty"`
}

func (s SetWarningThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s SetWarningThresholdRequest) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdRequest) SetUid(v int64) *SetWarningThresholdRequest {
	s.Uid = &v
	return s
}

func (s *SetWarningThresholdRequest) SetWarningValue(v string) *SetWarningThresholdRequest {
	s.WarningValue = &v
	return s
}

type SetWarningThresholdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetWarningThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetWarningThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdResponseBody) SetCode(v string) *SetWarningThresholdResponseBody {
	s.Code = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetMessage(v string) *SetWarningThresholdResponseBody {
	s.Message = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetRequestId(v string) *SetWarningThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWarningThresholdResponseBody) SetSuccess(v bool) *SetWarningThresholdResponseBody {
	s.Success = &v
	return s
}

type SetWarningThresholdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetWarningThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetWarningThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s SetWarningThresholdResponse) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdResponse) SetHeaders(v map[string]*string) *SetWarningThresholdResponse {
	s.Headers = v
	return s
}

func (s *SetWarningThresholdResponse) SetStatusCode(v int32) *SetWarningThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWarningThresholdResponse) SetBody(v *SetWarningThresholdResponseBody) *SetWarningThresholdResponse {
	s.Body = v
	return s
}

type SubscriptionBillRequest struct {
	// The start month from which the bills are pushed. Specify the value in the yyyy-MM format.
	//
	// After the subscription is generated, the system automatically pushes the bill data that is generated from the month that you specified to the current point in time. Data of up to six months can be pushed. The current month is included. If you subscribe to the bills for more than six months, the subscription is invalid.
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The file format of the bill. Valid values: csv and parquet.
	//
	// If you subscribe to the bills of multiple file formats, we recommend that you store the bills in different OSS buckets to prevent file overwriting.
	BillFormat *string `json:"BillFormat,omitempty" xml:"BillFormat,omitempty"`
	// The ID of the user to which the OSS bucket belongs.
	//
	// If you are an eco-partner of Alibaba Cloud and you need to push the bills to the OSS bucket of a subordinate partner account, you must set this parameter to the ID of the subordinate partner account and grant the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission to the subordinate partner account.
	//
	// If you are an eco-partner of Alibaba Cloud and you need to push the bills to the OSS bucket of your own account, your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
	BucketOwnerId *int64 `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
	// The name of the Object Storage Service (OSS) bucket in which you want to store the bills.
	SubscribeBucket *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	// The maximum rows in a single bill file. If the number of bill rows exceed the upper limit, the bill is automatically split into multiple files. The name of each split file is in the `uid_billType_billCycle_SquenceNo_fileNo` format.
	//
	// Files whose names are the same except for the fileNo field are of the same type and belong to the same billing cycle.
	SubscribeSegmentSize *int32 `json:"SubscribeSegmentSize,omitempty" xml:"SubscribeSegmentSize,omitempty"`
	// The type of the bill to which you want to subscribe. Valid values: PartnerBillingItemDetailForBillingPeriod, PartnerBillingItemDetailMonthly, PartnerInstanceDetailForBillingPeriod, and PartnerInstanceDetailMonthly.
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s SubscriptionBillRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionBillRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionBillRequest) SetBeginBillingCycle(v string) *SubscriptionBillRequest {
	s.BeginBillingCycle = &v
	return s
}

func (s *SubscriptionBillRequest) SetBillFormat(v string) *SubscriptionBillRequest {
	s.BillFormat = &v
	return s
}

func (s *SubscriptionBillRequest) SetBucketOwnerId(v int64) *SubscriptionBillRequest {
	s.BucketOwnerId = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeBucket(v string) *SubscriptionBillRequest {
	s.SubscribeBucket = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeSegmentSize(v int32) *SubscriptionBillRequest {
	s.SubscribeSegmentSize = &v
	return s
}

func (s *SubscriptionBillRequest) SetSubscribeType(v string) *SubscriptionBillRequest {
	s.SubscribeType = &v
	return s
}

type SubscriptionBillResponseBody struct {
	// The HTTP status code that is returned.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubscriptionBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionBillResponseBody) GoString() string {
	return s.String()
}

func (s *SubscriptionBillResponseBody) SetCode(v string) *SubscriptionBillResponseBody {
	s.Code = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetData(v bool) *SubscriptionBillResponseBody {
	s.Data = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetMessage(v string) *SubscriptionBillResponseBody {
	s.Message = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetRequestId(v string) *SubscriptionBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscriptionBillResponseBody) SetSuccess(v bool) *SubscriptionBillResponseBody {
	s.Success = &v
	return s
}

type SubscriptionBillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubscriptionBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubscriptionBillResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionBillResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionBillResponse) SetHeaders(v map[string]*string) *SubscriptionBillResponse {
	s.Headers = v
	return s
}

func (s *SubscriptionBillResponse) SetStatusCode(v int32) *SubscriptionBillResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscriptionBillResponse) SetBody(v *SubscriptionBillResponseBody) *SubscriptionBillResponse {
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
		"ap-northeast-1":              tea.String("agency.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("agency.aliyuncs.com"),
		"ap-south-1":                  tea.String("agency.aliyuncs.com"),
		"ap-southeast-2":              tea.String("agency.aliyuncs.com"),
		"ap-southeast-3":              tea.String("agency.aliyuncs.com"),
		"ap-southeast-5":              tea.String("agency.aliyuncs.com"),
		"cn-beijing":                  tea.String("agency.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("agency.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("agency.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("agency.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("agency.aliyuncs.com"),
		"cn-chengdu":                  tea.String("agency.aliyuncs.com"),
		"cn-edge-1":                   tea.String("agency.aliyuncs.com"),
		"cn-fujian":                   tea.String("agency.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("agency.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("agency.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("agency.aliyuncs.com"),
		"cn-hongkong":                 tea.String("agency.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("agency.aliyuncs.com"),
		"cn-huhehaote":                tea.String("agency.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("agency.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("agency.aliyuncs.com"),
		"cn-qingdao":                  tea.String("agency.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("agency.aliyuncs.com"),
		"cn-shanghai":                 tea.String("agency.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("agency.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("agency.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("agency.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("agency.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("agency.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("agency.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("agency.aliyuncs.com"),
		"cn-wuhan":                    tea.String("agency.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("agency.aliyuncs.com"),
		"cn-yushanfang":               tea.String("agency.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("agency.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("agency.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("agency.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("agency.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("agency.aliyuncs.com"),
		"eu-central-1":                tea.String("agency.aliyuncs.com"),
		"eu-west-1":                   tea.String("agency.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("agency.aliyuncs.com"),
		"me-east-1":                   tea.String("agency.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("agency.aliyuncs.com"),
		"us-east-1":                   tea.String("agency.aliyuncs.com"),
		"us-west-1":                   tea.String("agency.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("agency"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
 * You can call this operation to cancel the subscription to only one type of bill at a time.
 * After the subscription to a type of bill is canceled, bills of this type are no longer pushed to the specified Object Storage Service (OSS) bucket.
 * **This topic is published only on the international site (alibabacloud.com).
 *
 * @param request CancelSubscriptionBillRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelSubscriptionBillResponse
 */
func (client *Client) CancelSubscriptionBillWithOptions(request *CancelSubscriptionBillRequest, runtime *util.RuntimeOptions) (_result *CancelSubscriptionBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscribeType)) {
		query["SubscribeType"] = request.SubscribeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelSubscriptionBill"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelSubscriptionBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
 * You can call this operation to cancel the subscription to only one type of bill at a time.
 * After the subscription to a type of bill is canceled, bills of this type are no longer pushed to the specified Object Storage Service (OSS) bucket.
 * **This topic is published only on the international site (alibabacloud.com).
 *
 * @param request CancelSubscriptionBillRequest
 * @return CancelSubscriptionBillResponse
 */
func (client *Client) CancelSubscriptionBill(request *CancelSubscriptionBillRequest) (_result *CancelSubscriptionBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelSubscriptionBillResponse{}
	_body, _err := client.CancelSubscriptionBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomerWithOptions(request *CreateCustomerRequest, runtime *util.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerName)) {
		query["CustomerName"] = request.CustomerName
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerSource)) {
		query["CustomerSource"] = request.CustomerSource
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerSubTrade)) {
		query["CustomerSubTrade"] = request.CustomerSubTrade
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerTrade)) {
		query["CustomerTrade"] = request.CustomerTrade
	}

	if !tea.BoolValue(util.IsUnset(request.Nation)) {
		query["Nation"] = request.Nation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomer"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomer(request *CreateCustomerRequest) (_result *CreateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CreateCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeductOutstandingBalanceWithOptions(request *DeductOutstandingBalanceRequest, runtime *util.RuntimeOptions) (_result *DeductOutstandingBalanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeductAmount)) {
		query["DeductAmount"] = request.DeductAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeductOutstandingBalance"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeductOutstandingBalanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeductOutstandingBalance(request *DeductOutstandingBalanceRequest) (_result *DeductOutstandingBalanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeductOutstandingBalanceResponse{}
	_body, _err := client.DeductOutstandingBalanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditEndUserStatusWithOptions(request *EditEndUserStatusRequest, runtime *util.RuntimeOptions) (_result *EditEndUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditEndUserStatus"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditEndUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditEndUserStatus(request *EditEndUserStatusRequest) (_result *EditEndUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditEndUserStatusResponse{}
	_body, _err := client.EditEndUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditNewBuyStatusWithOptions(request *EditNewBuyStatusRequest, runtime *util.RuntimeOptions) (_result *EditNewBuyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewBuyStatus)) {
		query["NewBuyStatus"] = request.NewBuyStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditNewBuyStatus"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditNewBuyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditNewBuyStatus(request *EditNewBuyStatusRequest) (_result *EditNewBuyStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditNewBuyStatusResponse{}
	_body, _err := client.EditNewBuyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditZeroCreditShutdownWithOptions(request *EditZeroCreditShutdownRequest, runtime *util.RuntimeOptions) (_result *EditZeroCreditShutdownResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShutdownPolicy)) {
		query["ShutdownPolicy"] = request.ShutdownPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditZeroCreditShutdown"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditZeroCreditShutdownResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditZeroCreditShutdown(request *EditZeroCreditShutdownRequest) (_result *EditZeroCreditShutdownResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditZeroCreditShutdownResponse{}
	_body, _err := client.EditZeroCreditShutdownWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountInfoWithOptions(request *GetAccountInfoRequest, runtime *util.RuntimeOptions) (_result *GetAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountInfo"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountInfo(request *GetAccountInfoRequest) (_result *GetAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.GetAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCreditInfoWithOptions(request *GetCreditInfoRequest, runtime *util.RuntimeOptions) (_result *GetCreditInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCreditInfo"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCreditInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCreditInfo(request *GetCreditInfoRequest) (_result *GetCreditInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCreditInfoResponse{}
	_body, _err := client.GetCreditInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDailyBillWithOptions(request *GetDailyBillRequest, runtime *util.RuntimeOptions) (_result *GetDailyBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillOwner)) {
		query["BillOwner"] = request.BillOwner
	}

	if !tea.BoolValue(util.IsUnset(request.BillType)) {
		query["BillType"] = request.BillType
	}

	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["Date"] = request.Date
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDailyBill"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDailyBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDailyBill(request *GetDailyBillRequest) (_result *GetDailyBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDailyBillResponse{}
	_body, _err := client.GetDailyBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInviteStatusWithOptions(request *GetInviteStatusRequest, runtime *util.RuntimeOptions) (_result *GetInviteStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviteStatusList)) {
		query["InviteStatusList"] = request.InviteStatusList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInviteStatus"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInviteStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInviteStatus(request *GetInviteStatusRequest) (_result *GetInviteStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInviteStatusResponse{}
	_body, _err := client.GetInviteStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMonthlyBillWithOptions(request *GetMonthlyBillRequest, runtime *util.RuntimeOptions) (_result *GetMonthlyBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillOwner)) {
		query["BillOwner"] = request.BillOwner
	}

	if !tea.BoolValue(util.IsUnset(request.BillType)) {
		query["BillType"] = request.BillType
	}

	if !tea.BoolValue(util.IsUnset(request.Month)) {
		query["Month"] = request.Month
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMonthlyBill"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMonthlyBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMonthlyBill(request *GetMonthlyBillRequest) (_result *GetMonthlyBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMonthlyBillResponse{}
	_body, _err := client.GetMonthlyBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUnassociatedCustomerWithOptions(request *GetUnassociatedCustomerRequest, runtime *util.RuntimeOptions) (_result *GetUnassociatedCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUnassociatedCustomer"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUnassociatedCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUnassociatedCustomer(request *GetUnassociatedCustomerRequest) (_result *GetUnassociatedCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUnassociatedCustomerResponse{}
	_body, _err := client.GetUnassociatedCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InviteSubAccountWithOptions(request *InviteSubAccountRequest, runtime *util.RuntimeOptions) (_result *InviteSubAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountInfoList)) {
		query["AccountInfoList"] = request.AccountInfoList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InviteSubAccount"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InviteSubAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InviteSubAccount(request *InviteSubAccountRequest) (_result *InviteSubAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InviteSubAccountResponse{}
	_body, _err := client.InviteSubAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCountriesWithOptions(runtime *util.RuntimeOptions) (_result *ListCountriesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListCountries"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCountriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCountries() (_result *ListCountriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCountriesResponse{}
	_body, _err := client.ListCountriesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResendEmailWithOptions(request *ResendEmailRequest, runtime *util.RuntimeOptions) (_result *ResendEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviteId)) {
		query["InviteId"] = request.InviteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResendEmail"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResendEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResendEmail(request *ResendEmailRequest) (_result *ResendEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResendEmailResponse{}
	_body, _err := client.ResendEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAccountInfoWithOptions(request *SetAccountInfoRequest, runtime *util.RuntimeOptions) (_result *SetAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNickname)) {
		query["AccountNickname"] = request.AccountNickname
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAccountInfo"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAccountInfo(request *SetAccountInfoRequest) (_result *SetAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAccountInfoResponse{}
	_body, _err := client.SetAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCreditLineWithOptions(request *SetCreditLineRequest, runtime *util.RuntimeOptions) (_result *SetCreditLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreditLine)) {
		query["CreditLine"] = request.CreditLine
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetCreditLine"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetCreditLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCreditLine(request *SetCreditLineRequest) (_result *SetCreditLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCreditLineResponse{}
	_body, _err := client.SetCreditLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetWarningThresholdWithOptions(request *SetWarningThresholdRequest, runtime *util.RuntimeOptions) (_result *SetWarningThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.WarningValue)) {
		query["WarningValue"] = request.WarningValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetWarningThreshold"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetWarningThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetWarningThreshold(request *SetWarningThresholdRequest) (_result *SetWarningThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetWarningThresholdResponse{}
	_body, _err := client.SetWarningThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
 * *   You can call this operation to subscribe to only one type of bill at a time.
 * *   After the subscription to a type of bill is generated, the bill for the previous day is pushed on a daily basis from the next day. On the fifth day of each month, the full-data bill for the previous month is pushed.
 * *   A daily bill may be delayed. The delayed bill is pushed the next day after it is generated. The delayed bill may contain the bill data that is delayed until the previous day. We recommend that you query the full-data bill for the previous month at the beginning of each month.
 * *   Your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
 * *   The following file name formats are supported for bills:
 * ```
 * BillingItemDetailForBillingPeriod
 *
 * File name format of a daily bill: UID_PartnerBillingItemDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_BillingItemDetail_20190310_0001_01.
 *
 * File name format of a monthly full-data bill: UID_PartnerBillingItemDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetail_201903_0001_01.
 * InstanceDetailForBillingPeriod
 *
 *  File name format of a daily bill: UID_PartnerInstanceDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_InstanceDetail_20190310_0001_01.
 *
 * File name format of a monthly full-data bill: UID_PartnerInstanceDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetail_201903_1999-0001_01.
 * BillingItemDetailMonthly
 *
 * File name format of a daily bill: UID_PartnerBillingItemDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
 * InstanceDetailMonthly
 *
 * File name format of a daily bill: UID_PartnerInstanceDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
 * The fileNo field exists only when the number of bill rows reaches the maximum rows in a single bill file and the bill is split into multiple files.
 * ```
 * **This topic is published only on the international site (alibabacloud.com).
 *
 * @param request SubscriptionBillRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SubscriptionBillResponse
 */
func (client *Client) SubscriptionBillWithOptions(request *SubscriptionBillRequest, runtime *util.RuntimeOptions) (_result *SubscriptionBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginBillingCycle)) {
		query["BeginBillingCycle"] = request.BeginBillingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.BillFormat)) {
		query["BillFormat"] = request.BillFormat
	}

	if !tea.BoolValue(util.IsUnset(request.BucketOwnerId)) {
		query["BucketOwnerId"] = request.BucketOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeBucket)) {
		query["SubscribeBucket"] = request.SubscribeBucket
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeSegmentSize)) {
		query["SubscribeSegmentSize"] = request.SubscribeSegmentSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeType)) {
		query["SubscribeType"] = request.SubscribeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubscriptionBill"),
		Version:     tea.String("2022-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubscriptionBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
 * *   You can call this operation to subscribe to only one type of bill at a time.
 * *   After the subscription to a type of bill is generated, the bill for the previous day is pushed on a daily basis from the next day. On the fifth day of each month, the full-data bill for the previous month is pushed.
 * *   A daily bill may be delayed. The delayed bill is pushed the next day after it is generated. The delayed bill may contain the bill data that is delayed until the previous day. We recommend that you query the full-data bill for the previous month at the beginning of each month.
 * *   Your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
 * *   The following file name formats are supported for bills:
 * ```
 * BillingItemDetailForBillingPeriod
 *
 * File name format of a daily bill: UID_PartnerBillingItemDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_BillingItemDetail_20190310_0001_01.
 *
 * File name format of a monthly full-data bill: UID_PartnerBillingItemDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetail_201903_0001_01.
 * InstanceDetailForBillingPeriod
 *
 *  File name format of a daily bill: UID_PartnerInstanceDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_InstanceDetail_20190310_0001_01.
 *
 * File name format of a monthly full-data bill: UID_PartnerInstanceDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetail_201903_1999-0001_01.
 * BillingItemDetailMonthly
 *
 * File name format of a daily bill: UID_PartnerBillingItemDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
 * InstanceDetailMonthly
 *
 * File name format of a daily bill: UID_PartnerInstanceDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
 * The fileNo field exists only when the number of bill rows reaches the maximum rows in a single bill file and the bill is split into multiple files.
 * ```
 * **This topic is published only on the international site (alibabacloud.com).
 *
 * @param request SubscriptionBillRequest
 * @return SubscriptionBillResponse
 */
func (client *Client) SubscriptionBill(request *SubscriptionBillRequest) (_result *SubscriptionBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubscriptionBillResponse{}
	_body, _err := client.SubscriptionBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
