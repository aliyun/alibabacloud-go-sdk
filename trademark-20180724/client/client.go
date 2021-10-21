// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryTradeProduceListRequest struct {
	RegisterNumber *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	PageNum        *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PreOrderId     *string `json:"PreOrderId,omitempty" xml:"PreOrderId,omitempty"`
	BuyerStatus    *int32  `json:"BuyerStatus,omitempty" xml:"BuyerStatus,omitempty"`
	SortOrder      *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	SortFiled      *string `json:"SortFiled,omitempty" xml:"SortFiled,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s QueryTradeProduceListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceListRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceListRequest) SetRegisterNumber(v string) *QueryTradeProduceListRequest {
	s.RegisterNumber = &v
	return s
}

func (s *QueryTradeProduceListRequest) SetPageNum(v int32) *QueryTradeProduceListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTradeProduceListRequest) SetPageSize(v int32) *QueryTradeProduceListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTradeProduceListRequest) SetPreOrderId(v string) *QueryTradeProduceListRequest {
	s.PreOrderId = &v
	return s
}

func (s *QueryTradeProduceListRequest) SetBuyerStatus(v int32) *QueryTradeProduceListRequest {
	s.BuyerStatus = &v
	return s
}

func (s *QueryTradeProduceListRequest) SetSortOrder(v string) *QueryTradeProduceListRequest {
	s.SortOrder = &v
	return s
}

func (s *QueryTradeProduceListRequest) SetSortFiled(v string) *QueryTradeProduceListRequest {
	s.SortFiled = &v
	return s
}

func (s *QueryTradeProduceListRequest) SetBizId(v string) *QueryTradeProduceListRequest {
	s.BizId = &v
	return s
}

type QueryTradeProduceListResponseBody struct {
	CurrentPageNum *int32                                 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                                 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                                 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryTradeProduceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTradeProduceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceListResponseBody) SetCurrentPageNum(v int32) *QueryTradeProduceListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTradeProduceListResponseBody) SetTotalPageNum(v int32) *QueryTradeProduceListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTradeProduceListResponseBody) SetRequestId(v string) *QueryTradeProduceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTradeProduceListResponseBody) SetPageSize(v int32) *QueryTradeProduceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTradeProduceListResponseBody) SetTotalItemNum(v int32) *QueryTradeProduceListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTradeProduceListResponseBody) SetData(v *QueryTradeProduceListResponseBodyData) *QueryTradeProduceListResponseBody {
	s.Data = v
	return s
}

type QueryTradeProduceListResponseBodyData struct {
	TradeProduces []*QueryTradeProduceListResponseBodyDataTradeProduces `json:"TradeProduces,omitempty" xml:"TradeProduces,omitempty" type:"Repeated"`
}

func (s QueryTradeProduceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceListResponseBodyData) SetTradeProduces(v []*QueryTradeProduceListResponseBodyDataTradeProduces) *QueryTradeProduceListResponseBodyData {
	s.TradeProduces = v
	return s
}

type QueryTradeProduceListResponseBodyDataTradeProduces struct {
	UpdateTime     *int64   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	PreAmount      *float32 `json:"PreAmount,omitempty" xml:"PreAmount,omitempty"`
	CreateTime     *int64   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId         *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizId          *string  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Icon           *string  `json:"Icon,omitempty" xml:"Icon,omitempty"`
	BuyerStatus    *int32   `json:"BuyerStatus,omitempty" xml:"BuyerStatus,omitempty"`
	Source         *int32   `json:"Source,omitempty" xml:"Source,omitempty"`
	OperateNote    *string  `json:"OperateNote,omitempty" xml:"OperateNote,omitempty"`
	PreOrderId     *string  `json:"PreOrderId,omitempty" xml:"PreOrderId,omitempty"`
	AllowCancel    *bool    `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	RegisterNumber *string  `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	Classification *string  `json:"Classification,omitempty" xml:"Classification,omitempty"`
	FinalAmount    *float32 `json:"FinalAmount,omitempty" xml:"FinalAmount,omitempty"`
	FailReason     *int32   `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
}

func (s QueryTradeProduceListResponseBodyDataTradeProduces) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceListResponseBodyDataTradeProduces) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetUpdateTime(v int64) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.UpdateTime = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetPreAmount(v float32) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.PreAmount = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetCreateTime(v int64) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.CreateTime = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetUserId(v string) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.UserId = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetBizId(v string) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.BizId = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetIcon(v string) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.Icon = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetBuyerStatus(v int32) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.BuyerStatus = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetSource(v int32) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.Source = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetOperateNote(v string) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.OperateNote = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetPreOrderId(v string) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.PreOrderId = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetAllowCancel(v bool) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.AllowCancel = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetRegisterNumber(v string) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.RegisterNumber = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetClassification(v string) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.Classification = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetFinalAmount(v float32) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.FinalAmount = &v
	return s
}

func (s *QueryTradeProduceListResponseBodyDataTradeProduces) SetFailReason(v int32) *QueryTradeProduceListResponseBodyDataTradeProduces {
	s.FailReason = &v
	return s
}

type QueryTradeProduceListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTradeProduceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTradeProduceListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceListResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceListResponse) SetHeaders(v map[string]*string) *QueryTradeProduceListResponse {
	s.Headers = v
	return s
}

func (s *QueryTradeProduceListResponse) SetBody(v *QueryTradeProduceListResponseBody) *QueryTradeProduceListResponse {
	s.Body = v
	return s
}

type QueryTrademarkMonitorResultsRequest struct {
	RuleId             *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ActionType         *int32  `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	ProcedureStatus    *int32  `json:"ProcedureStatus,omitempty" xml:"ProcedureStatus,omitempty"`
	TmName             *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	ApplyYear          *string `json:"ApplyYear,omitempty" xml:"ApplyYear,omitempty"`
	RegistrationNumber *string `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	Classification     *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	PageNum            *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryTrademarkMonitorResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorResultsRequest) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorResultsRequest) SetRuleId(v int64) *QueryTrademarkMonitorResultsRequest {
	s.RuleId = &v
	return s
}

func (s *QueryTrademarkMonitorResultsRequest) SetActionType(v int32) *QueryTrademarkMonitorResultsRequest {
	s.ActionType = &v
	return s
}

func (s *QueryTrademarkMonitorResultsRequest) SetProcedureStatus(v int32) *QueryTrademarkMonitorResultsRequest {
	s.ProcedureStatus = &v
	return s
}

func (s *QueryTrademarkMonitorResultsRequest) SetTmName(v string) *QueryTrademarkMonitorResultsRequest {
	s.TmName = &v
	return s
}

func (s *QueryTrademarkMonitorResultsRequest) SetApplyYear(v string) *QueryTrademarkMonitorResultsRequest {
	s.ApplyYear = &v
	return s
}

func (s *QueryTrademarkMonitorResultsRequest) SetRegistrationNumber(v string) *QueryTrademarkMonitorResultsRequest {
	s.RegistrationNumber = &v
	return s
}

func (s *QueryTrademarkMonitorResultsRequest) SetClassification(v string) *QueryTrademarkMonitorResultsRequest {
	s.Classification = &v
	return s
}

func (s *QueryTrademarkMonitorResultsRequest) SetPageNum(v int32) *QueryTrademarkMonitorResultsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTrademarkMonitorResultsRequest) SetPageSize(v int32) *QueryTrademarkMonitorResultsRequest {
	s.PageSize = &v
	return s
}

type QueryTrademarkMonitorResultsResponseBody struct {
	NextPage       *bool                                         `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrePage        *bool                                         `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	TotalItemNum   *int32                                        `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	CurrentPageNum *int32                                        `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                                        `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	PageSize       *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data           *QueryTrademarkMonitorResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTrademarkMonitorResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorResultsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorResultsResponseBody) SetNextPage(v bool) *QueryTrademarkMonitorResultsResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBody) SetRequestId(v string) *QueryTrademarkMonitorResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBody) SetPrePage(v bool) *QueryTrademarkMonitorResultsResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBody) SetTotalItemNum(v int32) *QueryTrademarkMonitorResultsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBody) SetCurrentPageNum(v int32) *QueryTrademarkMonitorResultsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBody) SetTotalPageNum(v int32) *QueryTrademarkMonitorResultsResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBody) SetPageSize(v int32) *QueryTrademarkMonitorResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBody) SetData(v *QueryTrademarkMonitorResultsResponseBodyData) *QueryTrademarkMonitorResultsResponseBody {
	s.Data = v
	return s
}

type QueryTrademarkMonitorResultsResponseBodyData struct {
	TmMonitorResult []*QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult `json:"TmMonitorResult,omitempty" xml:"TmMonitorResult,omitempty" type:"Repeated"`
}

func (s QueryTrademarkMonitorResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorResultsResponseBodyData) SetTmMonitorResult(v []*QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) *QueryTrademarkMonitorResultsResponseBodyData {
	s.TmMonitorResult = v
	return s
}

type QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult struct {
	TmProcedureStatusDesc *string `json:"TmProcedureStatusDesc,omitempty" xml:"TmProcedureStatusDesc,omitempty"`
	WuxiaoEndDate         *string `json:"WuxiaoEndDate,omitempty" xml:"WuxiaoEndDate,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	OwnerEnName           *string `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	TmUid                 *string `json:"TmUid,omitempty" xml:"TmUid,omitempty"`
	OwnerName             *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	DataUpdateTime        *int64  `json:"DataUpdateTime,omitempty" xml:"DataUpdateTime,omitempty"`
	ChesanEndDate         *string `json:"ChesanEndDate,omitempty" xml:"ChesanEndDate,omitempty"`
	XuzhanEndDate         *string `json:"XuzhanEndDate,omitempty" xml:"XuzhanEndDate,omitempty"`
	RuleId                *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RegistrationNumber    *string `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	TmName                *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmImage               *string `json:"TmImage,omitempty" xml:"TmImage,omitempty"`
	DataCreateTime        *int64  `json:"DataCreateTime,omitempty" xml:"DataCreateTime,omitempty"`
	YiyiEndDate           *string `json:"YiyiEndDate,omitempty" xml:"YiyiEndDate,omitempty"`
	Classification        *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	ApplyDate             *string `json:"ApplyDate,omitempty" xml:"ApplyDate,omitempty"`
}

func (s QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetTmProcedureStatusDesc(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.TmProcedureStatusDesc = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetWuxiaoEndDate(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.WuxiaoEndDate = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetUserId(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.UserId = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetOwnerEnName(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.OwnerEnName = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetTmUid(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.TmUid = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetOwnerName(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.OwnerName = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetDataUpdateTime(v int64) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.DataUpdateTime = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetChesanEndDate(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.ChesanEndDate = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetXuzhanEndDate(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.XuzhanEndDate = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetRuleId(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.RuleId = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetRegistrationNumber(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.RegistrationNumber = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetTmName(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.TmName = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetTmImage(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.TmImage = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetDataCreateTime(v int64) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.DataCreateTime = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetYiyiEndDate(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.YiyiEndDate = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetClassification(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.Classification = &v
	return s
}

func (s *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult) SetApplyDate(v string) *QueryTrademarkMonitorResultsResponseBodyDataTmMonitorResult {
	s.ApplyDate = &v
	return s
}

type QueryTrademarkMonitorResultsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTrademarkMonitorResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTrademarkMonitorResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorResultsResponse) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorResultsResponse) SetHeaders(v map[string]*string) *QueryTrademarkMonitorResultsResponse {
	s.Headers = v
	return s
}

func (s *QueryTrademarkMonitorResultsResponse) SetBody(v *QueryTrademarkMonitorResultsResponseBody) *QueryTrademarkMonitorResultsResponse {
	s.Body = v
	return s
}

type CancelTradeOrderRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s CancelTradeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelTradeOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelTradeOrderRequest) SetBizId(v string) *CancelTradeOrderRequest {
	s.BizId = &v
	return s
}

type CancelTradeOrderResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s CancelTradeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelTradeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTradeOrderResponseBody) SetErrorMsg(v string) *CancelTradeOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CancelTradeOrderResponseBody) SetRequestId(v string) *CancelTradeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTradeOrderResponseBody) SetSuccess(v bool) *CancelTradeOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CancelTradeOrderResponseBody) SetErrorCode(v string) *CancelTradeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

type CancelTradeOrderResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelTradeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelTradeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelTradeOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelTradeOrderResponse) SetHeaders(v map[string]*string) *CancelTradeOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelTradeOrderResponse) SetBody(v *CancelTradeOrderResponseBody) *CancelTradeOrderResponse {
	s.Body = v
	return s
}

type DeleteTmMonitorRuleRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteTmMonitorRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTmMonitorRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteTmMonitorRuleRequest) SetId(v int64) *DeleteTmMonitorRuleRequest {
	s.Id = &v
	return s
}

type DeleteTmMonitorRuleResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s DeleteTmMonitorRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTmMonitorRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTmMonitorRuleResponseBody) SetErrorMsg(v string) *DeleteTmMonitorRuleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteTmMonitorRuleResponseBody) SetRequestId(v string) *DeleteTmMonitorRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTmMonitorRuleResponseBody) SetSuccess(v bool) *DeleteTmMonitorRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTmMonitorRuleResponseBody) SetErrorCode(v string) *DeleteTmMonitorRuleResponseBody {
	s.ErrorCode = &v
	return s
}

type DeleteTmMonitorRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTmMonitorRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTmMonitorRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTmMonitorRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteTmMonitorRuleResponse) SetHeaders(v map[string]*string) *DeleteTmMonitorRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteTmMonitorRuleResponse) SetBody(v *DeleteTmMonitorRuleResponseBody) *DeleteTmMonitorRuleResponse {
	s.Body = v
	return s
}

type UploadNotaryDataRequest struct {
	NotaryType    *int32  `json:"NotaryType,omitempty" xml:"NotaryType,omitempty"`
	BizOrderNo    *string `json:"BizOrderNo,omitempty" xml:"BizOrderNo,omitempty"`
	UploadContext *string `json:"UploadContext,omitempty" xml:"UploadContext,omitempty"`
}

func (s UploadNotaryDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadNotaryDataRequest) GoString() string {
	return s.String()
}

func (s *UploadNotaryDataRequest) SetNotaryType(v int32) *UploadNotaryDataRequest {
	s.NotaryType = &v
	return s
}

func (s *UploadNotaryDataRequest) SetBizOrderNo(v string) *UploadNotaryDataRequest {
	s.BizOrderNo = &v
	return s
}

func (s *UploadNotaryDataRequest) SetUploadContext(v string) *UploadNotaryDataRequest {
	s.UploadContext = &v
	return s
}

type UploadNotaryDataResponseBody struct {
	UserAuthUrl *string `json:"UserAuthUrl,omitempty" xml:"UserAuthUrl,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadNotaryDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadNotaryDataResponseBody) GoString() string {
	return s.String()
}

func (s *UploadNotaryDataResponseBody) SetUserAuthUrl(v string) *UploadNotaryDataResponseBody {
	s.UserAuthUrl = &v
	return s
}

func (s *UploadNotaryDataResponseBody) SetRequestId(v string) *UploadNotaryDataResponseBody {
	s.RequestId = &v
	return s
}

type UploadNotaryDataResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadNotaryDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadNotaryDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadNotaryDataResponse) GoString() string {
	return s.String()
}

func (s *UploadNotaryDataResponse) SetHeaders(v map[string]*string) *UploadNotaryDataResponse {
	s.Headers = v
	return s
}

func (s *UploadNotaryDataResponse) SetBody(v *UploadNotaryDataResponseBody) *UploadNotaryDataResponse {
	s.Body = v
	return s
}

type CopyApplicantRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CopyApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyApplicantRequest) GoString() string {
	return s.String()
}

func (s *CopyApplicantRequest) SetId(v int64) *CopyApplicantRequest {
	s.Id = &v
	return s
}

type CopyApplicantResponseBody struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *CopyApplicantResponseBody) SetId(v int64) *CopyApplicantResponseBody {
	s.Id = &v
	return s
}

func (s *CopyApplicantResponseBody) SetRequestId(v string) *CopyApplicantResponseBody {
	s.RequestId = &v
	return s
}

type CopyApplicantResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CopyApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopyApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyApplicantResponse) GoString() string {
	return s.String()
}

func (s *CopyApplicantResponse) SetHeaders(v map[string]*string) *CopyApplicantResponse {
	s.Headers = v
	return s
}

func (s *CopyApplicantResponse) SetBody(v *CopyApplicantResponseBody) *CopyApplicantResponse {
	s.Body = v
	return s
}

type PartnerUpdateTrademarkNameRequest struct {
	// tmIcon
	TmIcon *string `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	// aliyunKp
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// type
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// bizId
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// callerType
	CallerType *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	// callerParentId
	CallerParentId *int64 `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	// tmComment
	TmComment *string `json:"TmComment,omitempty" xml:"TmComment,omitempty"`
	// tmName
	TmName *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	// bid
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	EventSceneType *string `json:"EventSceneType,omitempty" xml:"EventSceneType,omitempty"`
}

func (s PartnerUpdateTrademarkNameRequest) String() string {
	return tea.Prettify(s)
}

func (s PartnerUpdateTrademarkNameRequest) GoString() string {
	return s.String()
}

func (s *PartnerUpdateTrademarkNameRequest) SetTmIcon(v string) *PartnerUpdateTrademarkNameRequest {
	s.TmIcon = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetAliyunKp(v string) *PartnerUpdateTrademarkNameRequest {
	s.AliyunKp = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetType(v int32) *PartnerUpdateTrademarkNameRequest {
	s.Type = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetBizId(v string) *PartnerUpdateTrademarkNameRequest {
	s.BizId = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetCallerType(v string) *PartnerUpdateTrademarkNameRequest {
	s.CallerType = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetCallerParentId(v int64) *PartnerUpdateTrademarkNameRequest {
	s.CallerParentId = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetTmComment(v string) *PartnerUpdateTrademarkNameRequest {
	s.TmComment = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetTmName(v string) *PartnerUpdateTrademarkNameRequest {
	s.TmName = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetBid(v string) *PartnerUpdateTrademarkNameRequest {
	s.Bid = &v
	return s
}

func (s *PartnerUpdateTrademarkNameRequest) SetEventSceneType(v string) *PartnerUpdateTrademarkNameRequest {
	s.EventSceneType = &v
	return s
}

type PartnerUpdateTrademarkNameResponseBody struct {
	// allowRetry
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// errorMsg
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// dynamicCode
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// errorCode
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// dynamicMessage
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// appName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s PartnerUpdateTrademarkNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PartnerUpdateTrademarkNameResponseBody) GoString() string {
	return s.String()
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetAllowRetry(v bool) *PartnerUpdateTrademarkNameResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetRequestId(v string) *PartnerUpdateTrademarkNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetErrorMsg(v string) *PartnerUpdateTrademarkNameResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetHttpStatusCode(v int32) *PartnerUpdateTrademarkNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetDynamicCode(v string) *PartnerUpdateTrademarkNameResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetErrorCode(v string) *PartnerUpdateTrademarkNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetDynamicMessage(v string) *PartnerUpdateTrademarkNameResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetSuccess(v bool) *PartnerUpdateTrademarkNameResponseBody {
	s.Success = &v
	return s
}

func (s *PartnerUpdateTrademarkNameResponseBody) SetAppName(v string) *PartnerUpdateTrademarkNameResponseBody {
	s.AppName = &v
	return s
}

type PartnerUpdateTrademarkNameResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PartnerUpdateTrademarkNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PartnerUpdateTrademarkNameResponse) String() string {
	return tea.Prettify(s)
}

func (s PartnerUpdateTrademarkNameResponse) GoString() string {
	return s.String()
}

func (s *PartnerUpdateTrademarkNameResponse) SetHeaders(v map[string]*string) *PartnerUpdateTrademarkNameResponse {
	s.Headers = v
	return s
}

func (s *PartnerUpdateTrademarkNameResponse) SetBody(v *PartnerUpdateTrademarkNameResponseBody) *PartnerUpdateTrademarkNameResponse {
	s.Body = v
	return s
}

type QueryIntentionDetailRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s QueryIntentionDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryIntentionDetailRequest) SetBizId(v string) *QueryIntentionDetailRequest {
	s.BizId = &v
	return s
}

type QueryIntentionDetailResponseBody struct {
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime     *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	RelationBizId  *string `json:"RelationBizId,omitempty" xml:"RelationBizId,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PartnerMobile  *string `json:"PartnerMobile,omitempty" xml:"PartnerMobile,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Mobile         *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	RegisterNumber *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryIntentionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIntentionDetailResponseBody) SetStatus(v int32) *QueryIntentionDetailResponseBody {
	s.Status = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetType(v int32) *QueryIntentionDetailResponseBody {
	s.Type = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetUpdateTime(v int64) *QueryIntentionDetailResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetRelationBizId(v string) *QueryIntentionDetailResponseBody {
	s.RelationBizId = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetCreateTime(v int64) *QueryIntentionDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetUserId(v string) *QueryIntentionDetailResponseBody {
	s.UserId = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetBizId(v string) *QueryIntentionDetailResponseBody {
	s.BizId = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetPartnerMobile(v string) *QueryIntentionDetailResponseBody {
	s.PartnerMobile = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetRequestId(v string) *QueryIntentionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetDescription(v string) *QueryIntentionDetailResponseBody {
	s.Description = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetMobile(v string) *QueryIntentionDetailResponseBody {
	s.Mobile = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetRegisterNumber(v string) *QueryIntentionDetailResponseBody {
	s.RegisterNumber = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetClassification(v string) *QueryIntentionDetailResponseBody {
	s.Classification = &v
	return s
}

func (s *QueryIntentionDetailResponseBody) SetUserName(v string) *QueryIntentionDetailResponseBody {
	s.UserName = &v
	return s
}

type QueryIntentionDetailResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryIntentionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryIntentionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryIntentionDetailResponse) SetHeaders(v map[string]*string) *QueryIntentionDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryIntentionDetailResponse) SetBody(v *QueryIntentionDetailResponseBody) *QueryIntentionDetailResponse {
	s.Body = v
	return s
}

type QueryIntentionPriceRequest struct {
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Channel        *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s QueryIntentionPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryIntentionPriceRequest) SetIntentionBizId(v string) *QueryIntentionPriceRequest {
	s.IntentionBizId = &v
	return s
}

func (s *QueryIntentionPriceRequest) SetChannel(v string) *QueryIntentionPriceRequest {
	s.Channel = &v
	return s
}

type QueryIntentionPriceResponseBody struct {
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                               `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                               `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryIntentionPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryIntentionPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIntentionPriceResponseBody) SetCurrentPageNum(v int32) *QueryIntentionPriceResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryIntentionPriceResponseBody) SetTotalPageNum(v int32) *QueryIntentionPriceResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryIntentionPriceResponseBody) SetRequestId(v string) *QueryIntentionPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIntentionPriceResponseBody) SetPageSize(v int32) *QueryIntentionPriceResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryIntentionPriceResponseBody) SetTotalItemNum(v int32) *QueryIntentionPriceResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryIntentionPriceResponseBody) SetData(v *QueryIntentionPriceResponseBodyData) *QueryIntentionPriceResponseBody {
	s.Data = v
	return s
}

type QueryIntentionPriceResponseBodyData struct {
	TmProduces []*QueryIntentionPriceResponseBodyDataTmProduces `json:"TmProduces,omitempty" xml:"TmProduces,omitempty" type:"Repeated"`
}

func (s QueryIntentionPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryIntentionPriceResponseBodyData) SetTmProduces(v []*QueryIntentionPriceResponseBodyDataTmProduces) *QueryIntentionPriceResponseBodyData {
	s.TmProduces = v
	return s
}

type QueryIntentionPriceResponseBodyDataTmProduces struct {
	Type                *int32                                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	Status              *int32                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	OrderPrice          *float32                                                          `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	UpdateTime          *int64                                                            `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	MaterialName        *string                                                           `json:"MaterialName,omitempty" xml:"MaterialName,omitempty"`
	CreateTime          *int64                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	BizId               *string                                                           `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ServicePrice        *float32                                                          `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	TmIcon              *string                                                           `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	TmName              *string                                                           `json:"TmName,omitempty" xml:"TmName,omitempty"`
	MaterialId          *string                                                           `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	SupplementId        *int64                                                            `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	LoaUrl              *string                                                           `json:"LoaUrl,omitempty" xml:"LoaUrl,omitempty"`
	TmNumber            *string                                                           `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	Note                *string                                                           `json:"Note,omitempty" xml:"Note,omitempty"`
	SupplementStatus    *int32                                                            `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	TotalPrice          *float32                                                          `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	ThirdClassification *QueryIntentionPriceResponseBodyDataTmProducesThirdClassification `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty" type:"Struct"`
	FirstClassification *QueryIntentionPriceResponseBodyDataTmProducesFirstClassification `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
}

func (s QueryIntentionPriceResponseBodyDataTmProduces) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionPriceResponseBodyDataTmProduces) GoString() string {
	return s.String()
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetType(v int32) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.Type = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetStatus(v int32) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.Status = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetOrderPrice(v float32) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.OrderPrice = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetUpdateTime(v int64) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.UpdateTime = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetMaterialName(v string) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.MaterialName = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetCreateTime(v int64) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.CreateTime = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetBizId(v string) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.BizId = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetServicePrice(v float32) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.ServicePrice = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetTmIcon(v string) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.TmIcon = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetTmName(v string) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.TmName = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetMaterialId(v string) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.MaterialId = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetSupplementId(v int64) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.SupplementId = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetLoaUrl(v string) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.LoaUrl = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetTmNumber(v string) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.TmNumber = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetNote(v string) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.Note = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetSupplementStatus(v int32) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.SupplementStatus = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetTotalPrice(v float32) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.TotalPrice = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetThirdClassification(v *QueryIntentionPriceResponseBodyDataTmProducesThirdClassification) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.ThirdClassification = v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProduces) SetFirstClassification(v *QueryIntentionPriceResponseBodyDataTmProducesFirstClassification) *QueryIntentionPriceResponseBodyDataTmProduces {
	s.FirstClassification = v
	return s
}

type QueryIntentionPriceResponseBodyDataTmProducesThirdClassification struct {
	ThirdClassifications []*QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications `json:"ThirdClassifications,omitempty" xml:"ThirdClassifications,omitempty" type:"Repeated"`
}

func (s QueryIntentionPriceResponseBodyDataTmProducesThirdClassification) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionPriceResponseBodyDataTmProducesThirdClassification) GoString() string {
	return s.String()
}

func (s *QueryIntentionPriceResponseBodyDataTmProducesThirdClassification) SetThirdClassifications(v []*QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications) *QueryIntentionPriceResponseBodyDataTmProducesThirdClassification {
	s.ThirdClassifications = v
	return s
}

type QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications struct {
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
}

func (s QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications) GoString() string {
	return s.String()
}

func (s *QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications) SetClassificationName(v string) *QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications {
	s.ClassificationName = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications) SetClassificationCode(v string) *QueryIntentionPriceResponseBodyDataTmProducesThirdClassificationThirdClassifications {
	s.ClassificationCode = &v
	return s
}

type QueryIntentionPriceResponseBodyDataTmProducesFirstClassification struct {
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
}

func (s QueryIntentionPriceResponseBodyDataTmProducesFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionPriceResponseBodyDataTmProducesFirstClassification) GoString() string {
	return s.String()
}

func (s *QueryIntentionPriceResponseBodyDataTmProducesFirstClassification) SetClassificationName(v string) *QueryIntentionPriceResponseBodyDataTmProducesFirstClassification {
	s.ClassificationName = &v
	return s
}

func (s *QueryIntentionPriceResponseBodyDataTmProducesFirstClassification) SetClassificationCode(v string) *QueryIntentionPriceResponseBodyDataTmProducesFirstClassification {
	s.ClassificationCode = &v
	return s
}

type QueryIntentionPriceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryIntentionPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryIntentionPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryIntentionPriceResponse) SetHeaders(v map[string]*string) *QueryIntentionPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryIntentionPriceResponse) SetBody(v *QueryIntentionPriceResponseBody) *QueryIntentionPriceResponse {
	s.Body = v
	return s
}

type QueryOfficialFileCustomListRequest struct {
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
}

func (s QueryOfficialFileCustomListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFileCustomListRequest) GoString() string {
	return s.String()
}

func (s *QueryOfficialFileCustomListRequest) SetPageSize(v int32) *QueryOfficialFileCustomListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOfficialFileCustomListRequest) SetPageNum(v int32) *QueryOfficialFileCustomListRequest {
	s.PageNum = &v
	return s
}

type QueryOfficialFileCustomListResponseBody struct {
	CurrentPageNum *int32                                       `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                                       `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                                       `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryOfficialFileCustomListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryOfficialFileCustomListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFileCustomListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOfficialFileCustomListResponseBody) SetCurrentPageNum(v int32) *QueryOfficialFileCustomListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBody) SetTotalPageNum(v int32) *QueryOfficialFileCustomListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBody) SetRequestId(v string) *QueryOfficialFileCustomListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBody) SetPageSize(v int32) *QueryOfficialFileCustomListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBody) SetTotalItemNum(v int32) *QueryOfficialFileCustomListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBody) SetData(v *QueryOfficialFileCustomListResponseBodyData) *QueryOfficialFileCustomListResponseBody {
	s.Data = v
	return s
}

type QueryOfficialFileCustomListResponseBodyData struct {
	CustomList []*QueryOfficialFileCustomListResponseBodyDataCustomList `json:"CustomList,omitempty" xml:"CustomList,omitempty" type:"Repeated"`
}

func (s QueryOfficialFileCustomListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFileCustomListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOfficialFileCustomListResponseBodyData) SetCustomList(v []*QueryOfficialFileCustomListResponseBodyDataCustomList) *QueryOfficialFileCustomListResponseBodyData {
	s.CustomList = v
	return s
}

type QueryOfficialFileCustomListResponseBodyDataCustomList struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ExpireTime      *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	DownloadUrl     *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndAcceptTime   *int64  `json:"EndAcceptTime,omitempty" xml:"EndAcceptTime,omitempty"`
	StartAcceptTime *int64  `json:"StartAcceptTime,omitempty" xml:"StartAcceptTime,omitempty"`
}

func (s QueryOfficialFileCustomListResponseBodyDataCustomList) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFileCustomListResponseBodyDataCustomList) GoString() string {
	return s.String()
}

func (s *QueryOfficialFileCustomListResponseBodyDataCustomList) SetStatus(v string) *QueryOfficialFileCustomListResponseBodyDataCustomList {
	s.Status = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBodyDataCustomList) SetExpireTime(v int64) *QueryOfficialFileCustomListResponseBodyDataCustomList {
	s.ExpireTime = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBodyDataCustomList) SetRemark(v string) *QueryOfficialFileCustomListResponseBodyDataCustomList {
	s.Remark = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBodyDataCustomList) SetDownloadUrl(v string) *QueryOfficialFileCustomListResponseBodyDataCustomList {
	s.DownloadUrl = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBodyDataCustomList) SetCreateTime(v int64) *QueryOfficialFileCustomListResponseBodyDataCustomList {
	s.CreateTime = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBodyDataCustomList) SetEndAcceptTime(v int64) *QueryOfficialFileCustomListResponseBodyDataCustomList {
	s.EndAcceptTime = &v
	return s
}

func (s *QueryOfficialFileCustomListResponseBodyDataCustomList) SetStartAcceptTime(v int64) *QueryOfficialFileCustomListResponseBodyDataCustomList {
	s.StartAcceptTime = &v
	return s
}

type QueryOfficialFileCustomListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOfficialFileCustomListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOfficialFileCustomListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOfficialFileCustomListResponse) GoString() string {
	return s.String()
}

func (s *QueryOfficialFileCustomListResponse) SetHeaders(v map[string]*string) *QueryOfficialFileCustomListResponse {
	s.Headers = v
	return s
}

func (s *QueryOfficialFileCustomListResponse) SetBody(v *QueryOfficialFileCustomListResponseBody) *QueryOfficialFileCustomListResponse {
	s.Body = v
	return s
}

type CheckTrademarkIconRequest struct {
	TrademarkIconOssKey *string `json:"TrademarkIconOssKey,omitempty" xml:"TrademarkIconOssKey,omitempty"`
	EventSceneType      *int32  `json:"EventSceneType,omitempty" xml:"EventSceneType,omitempty"`
}

func (s CheckTrademarkIconRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkIconRequest) GoString() string {
	return s.String()
}

func (s *CheckTrademarkIconRequest) SetTrademarkIconOssKey(v string) *CheckTrademarkIconRequest {
	s.TrademarkIconOssKey = &v
	return s
}

func (s *CheckTrademarkIconRequest) SetEventSceneType(v int32) *CheckTrademarkIconRequest {
	s.EventSceneType = &v
	return s
}

type CheckTrademarkIconResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckTrademarkIconResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkIconResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTrademarkIconResponseBody) SetMessage(v string) *CheckTrademarkIconResponseBody {
	s.Message = &v
	return s
}

func (s *CheckTrademarkIconResponseBody) SetRequestId(v string) *CheckTrademarkIconResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTrademarkIconResponseBody) SetResult(v string) *CheckTrademarkIconResponseBody {
	s.Result = &v
	return s
}

type CheckTrademarkIconResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckTrademarkIconResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckTrademarkIconResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkIconResponse) GoString() string {
	return s.String()
}

func (s *CheckTrademarkIconResponse) SetHeaders(v map[string]*string) *CheckTrademarkIconResponse {
	s.Headers = v
	return s
}

func (s *CheckTrademarkIconResponse) SetBody(v *CheckTrademarkIconResponseBody) *CheckTrademarkIconResponse {
	s.Body = v
	return s
}

type QuerySupplementDetailRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QuerySupplementDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplementDetailRequest) GoString() string {
	return s.String()
}

func (s *QuerySupplementDetailRequest) SetId(v int64) *QuerySupplementDetailRequest {
	s.Id = &v
	return s
}

type QuerySupplementDetailResponseBody struct {
	OperateTime           *int64                                             `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	SerialNumber          *string                                            `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status                *int32                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                  *int32                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	SbjDeadTime           *int64                                             `json:"SbjDeadTime,omitempty" xml:"SbjDeadTime,omitempty"`
	AcceptDeadTime        *int64                                             `json:"AcceptDeadTime,omitempty" xml:"AcceptDeadTime,omitempty"`
	SendTime              *int64                                             `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	AcceptTime            *int64                                             `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	RequestId             *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TmNumber              *string                                            `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	UploadFileTemplateUrl *string                                            `json:"UploadFileTemplateUrl,omitempty" xml:"UploadFileTemplateUrl,omitempty"`
	Content               *string                                            `json:"Content,omitempty" xml:"Content,omitempty"`
	Id                    *int64                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	FileTemplateUrls      *QuerySupplementDetailResponseBodyFileTemplateUrls `json:"FileTemplateUrls,omitempty" xml:"FileTemplateUrls,omitempty" type:"Struct"`
}

func (s QuerySupplementDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplementDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySupplementDetailResponseBody) SetOperateTime(v int64) *QuerySupplementDetailResponseBody {
	s.OperateTime = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetSerialNumber(v string) *QuerySupplementDetailResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetStatus(v int32) *QuerySupplementDetailResponseBody {
	s.Status = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetType(v int32) *QuerySupplementDetailResponseBody {
	s.Type = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetSbjDeadTime(v int64) *QuerySupplementDetailResponseBody {
	s.SbjDeadTime = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetAcceptDeadTime(v int64) *QuerySupplementDetailResponseBody {
	s.AcceptDeadTime = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetSendTime(v int64) *QuerySupplementDetailResponseBody {
	s.SendTime = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetAcceptTime(v int64) *QuerySupplementDetailResponseBody {
	s.AcceptTime = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetRequestId(v string) *QuerySupplementDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetTmNumber(v string) *QuerySupplementDetailResponseBody {
	s.TmNumber = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetUploadFileTemplateUrl(v string) *QuerySupplementDetailResponseBody {
	s.UploadFileTemplateUrl = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetContent(v string) *QuerySupplementDetailResponseBody {
	s.Content = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetId(v int64) *QuerySupplementDetailResponseBody {
	s.Id = &v
	return s
}

func (s *QuerySupplementDetailResponseBody) SetFileTemplateUrls(v *QuerySupplementDetailResponseBodyFileTemplateUrls) *QuerySupplementDetailResponseBody {
	s.FileTemplateUrls = v
	return s
}

type QuerySupplementDetailResponseBodyFileTemplateUrls struct {
	FileTemplateUrls []*string `json:"FileTemplateUrls,omitempty" xml:"FileTemplateUrls,omitempty" type:"Repeated"`
}

func (s QuerySupplementDetailResponseBodyFileTemplateUrls) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplementDetailResponseBodyFileTemplateUrls) GoString() string {
	return s.String()
}

func (s *QuerySupplementDetailResponseBodyFileTemplateUrls) SetFileTemplateUrls(v []*string) *QuerySupplementDetailResponseBodyFileTemplateUrls {
	s.FileTemplateUrls = v
	return s
}

type QuerySupplementDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySupplementDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySupplementDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySupplementDetailResponse) GoString() string {
	return s.String()
}

func (s *QuerySupplementDetailResponse) SetHeaders(v map[string]*string) *QuerySupplementDetailResponse {
	s.Headers = v
	return s
}

func (s *QuerySupplementDetailResponse) SetBody(v *QuerySupplementDetailResponseBody) *QuerySupplementDetailResponse {
	s.Body = v
	return s
}

type UploadTrademarkOnSaleRequest struct {
	ClassificationCode      *string  `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	TmName                  *string  `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmIcon                  *string  `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	OriginalPrice           *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TmNumber                *string  `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	EndTime                 *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BeginTime               *int64   `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Description             *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Label                   *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	RegAnnDate              *int64   `json:"RegAnnDate,omitempty" xml:"RegAnnDate,omitempty"`
	OwnerName               *string  `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerEnName             *string  `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	SecondaryClassification *string  `json:"SecondaryClassification,omitempty" xml:"SecondaryClassification,omitempty"`
	ThirdClassification     *string  `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty"`
	Type                    *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Reason                  *string  `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Status                  *string  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UploadTrademarkOnSaleRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadTrademarkOnSaleRequest) GoString() string {
	return s.String()
}

func (s *UploadTrademarkOnSaleRequest) SetClassificationCode(v string) *UploadTrademarkOnSaleRequest {
	s.ClassificationCode = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetTmName(v string) *UploadTrademarkOnSaleRequest {
	s.TmName = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetTmIcon(v string) *UploadTrademarkOnSaleRequest {
	s.TmIcon = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetOriginalPrice(v float32) *UploadTrademarkOnSaleRequest {
	s.OriginalPrice = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetTmNumber(v string) *UploadTrademarkOnSaleRequest {
	s.TmNumber = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetEndTime(v int64) *UploadTrademarkOnSaleRequest {
	s.EndTime = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetBeginTime(v int64) *UploadTrademarkOnSaleRequest {
	s.BeginTime = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetDescription(v string) *UploadTrademarkOnSaleRequest {
	s.Description = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetLabel(v string) *UploadTrademarkOnSaleRequest {
	s.Label = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetRegAnnDate(v int64) *UploadTrademarkOnSaleRequest {
	s.RegAnnDate = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetOwnerName(v string) *UploadTrademarkOnSaleRequest {
	s.OwnerName = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetOwnerEnName(v string) *UploadTrademarkOnSaleRequest {
	s.OwnerEnName = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetSecondaryClassification(v string) *UploadTrademarkOnSaleRequest {
	s.SecondaryClassification = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetThirdClassification(v string) *UploadTrademarkOnSaleRequest {
	s.ThirdClassification = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetType(v string) *UploadTrademarkOnSaleRequest {
	s.Type = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetReason(v string) *UploadTrademarkOnSaleRequest {
	s.Reason = &v
	return s
}

func (s *UploadTrademarkOnSaleRequest) SetStatus(v string) *UploadTrademarkOnSaleRequest {
	s.Status = &v
	return s
}

type UploadTrademarkOnSaleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadTrademarkOnSaleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadTrademarkOnSaleResponseBody) GoString() string {
	return s.String()
}

func (s *UploadTrademarkOnSaleResponseBody) SetRequestId(v string) *UploadTrademarkOnSaleResponseBody {
	s.RequestId = &v
	return s
}

type UploadTrademarkOnSaleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadTrademarkOnSaleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadTrademarkOnSaleResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadTrademarkOnSaleResponse) GoString() string {
	return s.String()
}

func (s *UploadTrademarkOnSaleResponse) SetHeaders(v map[string]*string) *UploadTrademarkOnSaleResponse {
	s.Headers = v
	return s
}

func (s *UploadTrademarkOnSaleResponse) SetBody(v *UploadTrademarkOnSaleResponseBody) *UploadTrademarkOnSaleResponse {
	s.Body = v
	return s
}

type ApplyNotaryPostRequest struct {
	NotaryOrderId   *int64  `json:"NotaryOrderId,omitempty" xml:"NotaryOrderId,omitempty"`
	ReceiverName    *string `json:"ReceiverName,omitempty" xml:"ReceiverName,omitempty"`
	ReceiverAddress *string `json:"ReceiverAddress,omitempty" xml:"ReceiverAddress,omitempty"`
	ReceiverPhone   *string `json:"ReceiverPhone,omitempty" xml:"ReceiverPhone,omitempty"`
}

func (s ApplyNotaryPostRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyNotaryPostRequest) GoString() string {
	return s.String()
}

func (s *ApplyNotaryPostRequest) SetNotaryOrderId(v int64) *ApplyNotaryPostRequest {
	s.NotaryOrderId = &v
	return s
}

func (s *ApplyNotaryPostRequest) SetReceiverName(v string) *ApplyNotaryPostRequest {
	s.ReceiverName = &v
	return s
}

func (s *ApplyNotaryPostRequest) SetReceiverAddress(v string) *ApplyNotaryPostRequest {
	s.ReceiverAddress = &v
	return s
}

func (s *ApplyNotaryPostRequest) SetReceiverPhone(v string) *ApplyNotaryPostRequest {
	s.ReceiverPhone = &v
	return s
}

type ApplyNotaryPostResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s ApplyNotaryPostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyNotaryPostResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyNotaryPostResponseBody) SetErrorMsg(v string) *ApplyNotaryPostResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ApplyNotaryPostResponseBody) SetRequestId(v string) *ApplyNotaryPostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyNotaryPostResponseBody) SetSuccess(v bool) *ApplyNotaryPostResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyNotaryPostResponseBody) SetErrorCode(v string) *ApplyNotaryPostResponseBody {
	s.ErrorCode = &v
	return s
}

type ApplyNotaryPostResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyNotaryPostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyNotaryPostResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyNotaryPostResponse) GoString() string {
	return s.String()
}

func (s *ApplyNotaryPostResponse) SetHeaders(v map[string]*string) *ApplyNotaryPostResponse {
	s.Headers = v
	return s
}

func (s *ApplyNotaryPostResponse) SetBody(v *ApplyNotaryPostResponseBody) *ApplyNotaryPostResponse {
	s.Body = v
	return s
}

type QueryTradeMarkApplicationsByIntentionRequest struct {
	IntentionBizId  *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Channel         *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	PageNum         *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TmProduceStatus *string `json:"TmProduceStatus,omitempty" xml:"TmProduceStatus,omitempty"`
}

func (s QueryTradeMarkApplicationsByIntentionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsByIntentionRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsByIntentionRequest) SetIntentionBizId(v string) *QueryTradeMarkApplicationsByIntentionRequest {
	s.IntentionBizId = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionRequest) SetChannel(v string) *QueryTradeMarkApplicationsByIntentionRequest {
	s.Channel = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionRequest) SetPageNum(v int32) *QueryTradeMarkApplicationsByIntentionRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionRequest) SetPageSize(v int32) *QueryTradeMarkApplicationsByIntentionRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionRequest) SetTmProduceStatus(v string) *QueryTradeMarkApplicationsByIntentionRequest {
	s.TmProduceStatus = &v
	return s
}

type QueryTradeMarkApplicationsByIntentionResponseBody struct {
	CurrentPageNum *int32                                                 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                                                 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                                                 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryTradeMarkApplicationsByIntentionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationsByIntentionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsByIntentionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBody) SetCurrentPageNum(v int32) *QueryTradeMarkApplicationsByIntentionResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBody) SetTotalPageNum(v int32) *QueryTradeMarkApplicationsByIntentionResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBody) SetRequestId(v string) *QueryTradeMarkApplicationsByIntentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBody) SetPageSize(v int32) *QueryTradeMarkApplicationsByIntentionResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBody) SetTotalItemNum(v int32) *QueryTradeMarkApplicationsByIntentionResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBody) SetData(v *QueryTradeMarkApplicationsByIntentionResponseBodyData) *QueryTradeMarkApplicationsByIntentionResponseBody {
	s.Data = v
	return s
}

type QueryTradeMarkApplicationsByIntentionResponseBodyData struct {
	TmProduces []*QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces `json:"TmProduces,omitempty" xml:"TmProduces,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyData) SetTmProduces(v []*QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) *QueryTradeMarkApplicationsByIntentionResponseBodyData {
	s.TmProduces = v
	return s
}

type QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces struct {
	PrincipalDescription *string                                                                             `json:"PrincipalDescription,omitempty" xml:"PrincipalDescription,omitempty"`
	Type                 *int32                                                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	Status               *int32                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	OrderPrice           *float32                                                                            `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	UpdateTime           *int64                                                                              `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	MaterialName         *string                                                                             `json:"MaterialName,omitempty" xml:"MaterialName,omitempty"`
	PrincipalValue       *int32                                                                              `json:"PrincipalValue,omitempty" xml:"PrincipalValue,omitempty"`
	CreateTime           *int64                                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	BizId                *string                                                                             `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ServicePrice         *float32                                                                            `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	TmIcon               *string                                                                             `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	TmName               *string                                                                             `json:"TmName,omitempty" xml:"TmName,omitempty"`
	MaterialId           *string                                                                             `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	SupplementId         *int64                                                                              `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	LoaUrl               *string                                                                             `json:"LoaUrl,omitempty" xml:"LoaUrl,omitempty"`
	TmNumber             *string                                                                             `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	Note                 *string                                                                             `json:"Note,omitempty" xml:"Note,omitempty"`
	SupplementStatus     *int32                                                                              `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	TotalPrice           *float32                                                                            `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	ThirdClassification  *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassification `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty" type:"Struct"`
	FirstClassification  *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetPrincipalDescription(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.PrincipalDescription = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetType(v int32) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.Type = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetStatus(v int32) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.Status = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetOrderPrice(v float32) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.OrderPrice = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetUpdateTime(v int64) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.UpdateTime = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetMaterialName(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.MaterialName = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetPrincipalValue(v int32) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.PrincipalValue = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetCreateTime(v int64) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.CreateTime = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetBizId(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.BizId = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetServicePrice(v float32) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.ServicePrice = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetTmIcon(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.TmIcon = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetTmName(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.TmName = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetMaterialId(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.MaterialId = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetSupplementId(v int64) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.SupplementId = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetLoaUrl(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.LoaUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetTmNumber(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.TmNumber = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetNote(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.Note = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetSupplementStatus(v int32) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.SupplementStatus = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetTotalPrice(v float32) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.TotalPrice = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetThirdClassification(v *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassification) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.ThirdClassification = v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces) SetFirstClassification(v *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProduces {
	s.FirstClassification = v
	return s
}

type QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassification struct {
	ThirdClassifications []*QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications `json:"ThirdClassifications,omitempty" xml:"ThirdClassifications,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassification) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassification) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassification) SetThirdClassifications(v []*QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassification {
	s.ThirdClassifications = v
	return s
}

type QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications struct {
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications) SetClassificationName(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications {
	s.ClassificationName = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications) SetClassificationCode(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesThirdClassificationThirdClassifications {
	s.ClassificationCode = &v
	return s
}

type QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification struct {
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification) SetClassificationName(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification {
	s.ClassificationName = &v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification) SetClassificationCode(v string) *QueryTradeMarkApplicationsByIntentionResponseBodyDataTmProducesFirstClassification {
	s.ClassificationCode = &v
	return s
}

type QueryTradeMarkApplicationsByIntentionResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTradeMarkApplicationsByIntentionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTradeMarkApplicationsByIntentionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsByIntentionResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsByIntentionResponse) SetHeaders(v map[string]*string) *QueryTradeMarkApplicationsByIntentionResponse {
	s.Headers = v
	return s
}

func (s *QueryTradeMarkApplicationsByIntentionResponse) SetBody(v *QueryTradeMarkApplicationsByIntentionResponseBody) *QueryTradeMarkApplicationsByIntentionResponse {
	s.Body = v
	return s
}

type SaveExtensionAttributeRequest struct {
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	AttributeKey   *string `json:"AttributeKey,omitempty" xml:"AttributeKey,omitempty"`
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s SaveExtensionAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveExtensionAttributeRequest) GoString() string {
	return s.String()
}

func (s *SaveExtensionAttributeRequest) SetBizId(v string) *SaveExtensionAttributeRequest {
	s.BizId = &v
	return s
}

func (s *SaveExtensionAttributeRequest) SetAttributeKey(v string) *SaveExtensionAttributeRequest {
	s.AttributeKey = &v
	return s
}

func (s *SaveExtensionAttributeRequest) SetAttributeValue(v string) *SaveExtensionAttributeRequest {
	s.AttributeValue = &v
	return s
}

type SaveExtensionAttributeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveExtensionAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveExtensionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveExtensionAttributeResponseBody) SetCode(v string) *SaveExtensionAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *SaveExtensionAttributeResponseBody) SetMessage(v string) *SaveExtensionAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *SaveExtensionAttributeResponseBody) SetRequestId(v string) *SaveExtensionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveExtensionAttributeResponseBody) SetSuccess(v bool) *SaveExtensionAttributeResponseBody {
	s.Success = &v
	return s
}

type SaveExtensionAttributeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveExtensionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveExtensionAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveExtensionAttributeResponse) GoString() string {
	return s.String()
}

func (s *SaveExtensionAttributeResponse) SetHeaders(v map[string]*string) *SaveExtensionAttributeResponse {
	s.Headers = v
	return s
}

func (s *SaveExtensionAttributeResponse) SetBody(v *SaveExtensionAttributeResponseBody) *SaveExtensionAttributeResponse {
	s.Body = v
	return s
}

type AcceptPartnerNotificationRequest struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Material  *string `json:"Material,omitempty" xml:"Material,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AcceptPartnerNotificationRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptPartnerNotificationRequest) GoString() string {
	return s.String()
}

func (s *AcceptPartnerNotificationRequest) SetBizId(v string) *AcceptPartnerNotificationRequest {
	s.BizId = &v
	return s
}

func (s *AcceptPartnerNotificationRequest) SetOperation(v string) *AcceptPartnerNotificationRequest {
	s.Operation = &v
	return s
}

func (s *AcceptPartnerNotificationRequest) SetMaterial(v string) *AcceptPartnerNotificationRequest {
	s.Material = &v
	return s
}

func (s *AcceptPartnerNotificationRequest) SetRemark(v string) *AcceptPartnerNotificationRequest {
	s.Remark = &v
	return s
}

type AcceptPartnerNotificationResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s AcceptPartnerNotificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptPartnerNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptPartnerNotificationResponseBody) SetErrorMsg(v string) *AcceptPartnerNotificationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AcceptPartnerNotificationResponseBody) SetRequestId(v string) *AcceptPartnerNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptPartnerNotificationResponseBody) SetSuccess(v bool) *AcceptPartnerNotificationResponseBody {
	s.Success = &v
	return s
}

func (s *AcceptPartnerNotificationResponseBody) SetErrorCode(v string) *AcceptPartnerNotificationResponseBody {
	s.ErrorCode = &v
	return s
}

type AcceptPartnerNotificationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AcceptPartnerNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AcceptPartnerNotificationResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptPartnerNotificationResponse) GoString() string {
	return s.String()
}

func (s *AcceptPartnerNotificationResponse) SetHeaders(v map[string]*string) *AcceptPartnerNotificationResponse {
	s.Headers = v
	return s
}

func (s *AcceptPartnerNotificationResponse) SetBody(v *AcceptPartnerNotificationResponseBody) *AcceptPartnerNotificationResponse {
	s.Body = v
	return s
}

type SubmitSupplementRequest struct {
	Id               *int64                 `json:"Id,omitempty" xml:"Id,omitempty"`
	UploadOssKeyList map[string]interface{} `json:"UploadOssKeyList,omitempty" xml:"UploadOssKeyList,omitempty"`
	Content          *string                `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s SubmitSupplementRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSupplementRequest) GoString() string {
	return s.String()
}

func (s *SubmitSupplementRequest) SetId(v int64) *SubmitSupplementRequest {
	s.Id = &v
	return s
}

func (s *SubmitSupplementRequest) SetUploadOssKeyList(v map[string]interface{}) *SubmitSupplementRequest {
	s.UploadOssKeyList = v
	return s
}

func (s *SubmitSupplementRequest) SetContent(v string) *SubmitSupplementRequest {
	s.Content = &v
	return s
}

type SubmitSupplementShrinkRequest struct {
	Id                     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	UploadOssKeyListShrink *string `json:"UploadOssKeyList,omitempty" xml:"UploadOssKeyList,omitempty"`
	Content                *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s SubmitSupplementShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSupplementShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSupplementShrinkRequest) SetId(v int64) *SubmitSupplementShrinkRequest {
	s.Id = &v
	return s
}

func (s *SubmitSupplementShrinkRequest) SetUploadOssKeyListShrink(v string) *SubmitSupplementShrinkRequest {
	s.UploadOssKeyListShrink = &v
	return s
}

func (s *SubmitSupplementShrinkRequest) SetContent(v string) *SubmitSupplementShrinkRequest {
	s.Content = &v
	return s
}

type SubmitSupplementResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s SubmitSupplementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSupplementResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSupplementResponseBody) SetErrorMsg(v string) *SubmitSupplementResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitSupplementResponseBody) SetRequestId(v string) *SubmitSupplementResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSupplementResponseBody) SetSuccess(v bool) *SubmitSupplementResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSupplementResponseBody) SetErrorCode(v string) *SubmitSupplementResponseBody {
	s.ErrorCode = &v
	return s
}

type SubmitSupplementResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSupplementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSupplementResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSupplementResponse) GoString() string {
	return s.String()
}

func (s *SubmitSupplementResponse) SetHeaders(v map[string]*string) *SubmitSupplementResponse {
	s.Headers = v
	return s
}

func (s *SubmitSupplementResponse) SetBody(v *SubmitSupplementResponseBody) *SubmitSupplementResponse {
	s.Body = v
	return s
}

type ForceUploadTrademarkOnsaleRequest struct {
	ClassificationCode      *string  `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	TmName                  *string  `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmIcon                  *string  `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	OriginalPrice           *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TmNumber                *string  `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	EndTime                 *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BeginTime               *int64   `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Description             *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Label                   *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	RegAnnDate              *int64   `json:"RegAnnDate,omitempty" xml:"RegAnnDate,omitempty"`
	OwnerName               *string  `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerEnName             *string  `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	SecondaryClassification *string  `json:"SecondaryClassification,omitempty" xml:"SecondaryClassification,omitempty"`
	ThirdClassification     *string  `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty"`
	Type                    *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Reason                  *string  `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ForceUploadTrademarkOnsaleRequest) String() string {
	return tea.Prettify(s)
}

func (s ForceUploadTrademarkOnsaleRequest) GoString() string {
	return s.String()
}

func (s *ForceUploadTrademarkOnsaleRequest) SetClassificationCode(v string) *ForceUploadTrademarkOnsaleRequest {
	s.ClassificationCode = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetTmName(v string) *ForceUploadTrademarkOnsaleRequest {
	s.TmName = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetTmIcon(v string) *ForceUploadTrademarkOnsaleRequest {
	s.TmIcon = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetOriginalPrice(v float32) *ForceUploadTrademarkOnsaleRequest {
	s.OriginalPrice = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetTmNumber(v string) *ForceUploadTrademarkOnsaleRequest {
	s.TmNumber = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetEndTime(v int64) *ForceUploadTrademarkOnsaleRequest {
	s.EndTime = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetBeginTime(v int64) *ForceUploadTrademarkOnsaleRequest {
	s.BeginTime = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetDescription(v string) *ForceUploadTrademarkOnsaleRequest {
	s.Description = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetLabel(v string) *ForceUploadTrademarkOnsaleRequest {
	s.Label = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetRegAnnDate(v int64) *ForceUploadTrademarkOnsaleRequest {
	s.RegAnnDate = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetOwnerName(v string) *ForceUploadTrademarkOnsaleRequest {
	s.OwnerName = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetOwnerEnName(v string) *ForceUploadTrademarkOnsaleRequest {
	s.OwnerEnName = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetSecondaryClassification(v string) *ForceUploadTrademarkOnsaleRequest {
	s.SecondaryClassification = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetThirdClassification(v string) *ForceUploadTrademarkOnsaleRequest {
	s.ThirdClassification = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetType(v string) *ForceUploadTrademarkOnsaleRequest {
	s.Type = &v
	return s
}

func (s *ForceUploadTrademarkOnsaleRequest) SetReason(v string) *ForceUploadTrademarkOnsaleRequest {
	s.Reason = &v
	return s
}

type ForceUploadTrademarkOnsaleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ForceUploadTrademarkOnsaleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ForceUploadTrademarkOnsaleResponseBody) GoString() string {
	return s.String()
}

func (s *ForceUploadTrademarkOnsaleResponseBody) SetRequestId(v string) *ForceUploadTrademarkOnsaleResponseBody {
	s.RequestId = &v
	return s
}

type ForceUploadTrademarkOnsaleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ForceUploadTrademarkOnsaleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ForceUploadTrademarkOnsaleResponse) String() string {
	return tea.Prettify(s)
}

func (s ForceUploadTrademarkOnsaleResponse) GoString() string {
	return s.String()
}

func (s *ForceUploadTrademarkOnsaleResponse) SetHeaders(v map[string]*string) *ForceUploadTrademarkOnsaleResponse {
	s.Headers = v
	return s
}

func (s *ForceUploadTrademarkOnsaleResponse) SetBody(v *ForceUploadTrademarkOnsaleResponseBody) *ForceUploadTrademarkOnsaleResponse {
	s.Body = v
	return s
}

type BindMaterialRequest struct {
	MaterialId     *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LoaOssKey      *string `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	LegalNoticeKey *string `json:"LegalNoticeKey,omitempty" xml:"LegalNoticeKey,omitempty"`
}

func (s BindMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s BindMaterialRequest) GoString() string {
	return s.String()
}

func (s *BindMaterialRequest) SetMaterialId(v string) *BindMaterialRequest {
	s.MaterialId = &v
	return s
}

func (s *BindMaterialRequest) SetBizId(v string) *BindMaterialRequest {
	s.BizId = &v
	return s
}

func (s *BindMaterialRequest) SetLoaOssKey(v string) *BindMaterialRequest {
	s.LoaOssKey = &v
	return s
}

func (s *BindMaterialRequest) SetLegalNoticeKey(v string) *BindMaterialRequest {
	s.LegalNoticeKey = &v
	return s
}

type BindMaterialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *BindMaterialResponseBody) SetRequestId(v string) *BindMaterialResponseBody {
	s.RequestId = &v
	return s
}

type BindMaterialResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s BindMaterialResponse) GoString() string {
	return s.String()
}

func (s *BindMaterialResponse) SetHeaders(v map[string]*string) *BindMaterialResponse {
	s.Headers = v
	return s
}

func (s *BindMaterialResponse) SetBody(v *BindMaterialResponseBody) *BindMaterialResponse {
	s.Body = v
	return s
}

type GetDefaultPrincipalResponseBody struct {
	PrincipalDescription *string `json:"PrincipalDescription,omitempty" xml:"PrincipalDescription,omitempty"`
	PrincipalName        *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrincipalValue       *int32  `json:"PrincipalValue,omitempty" xml:"PrincipalValue,omitempty"`
}

func (s GetDefaultPrincipalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultPrincipalResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultPrincipalResponseBody) SetPrincipalDescription(v string) *GetDefaultPrincipalResponseBody {
	s.PrincipalDescription = &v
	return s
}

func (s *GetDefaultPrincipalResponseBody) SetPrincipalName(v string) *GetDefaultPrincipalResponseBody {
	s.PrincipalName = &v
	return s
}

func (s *GetDefaultPrincipalResponseBody) SetRequestId(v string) *GetDefaultPrincipalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultPrincipalResponseBody) SetPrincipalValue(v int32) *GetDefaultPrincipalResponseBody {
	s.PrincipalValue = &v
	return s
}

type GetDefaultPrincipalResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDefaultPrincipalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDefaultPrincipalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultPrincipalResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultPrincipalResponse) SetHeaders(v map[string]*string) *GetDefaultPrincipalResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultPrincipalResponse) SetBody(v *GetDefaultPrincipalResponseBody) *GetDefaultPrincipalResponse {
	s.Body = v
	return s
}

type QueryCommunicationLogsRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryCommunicationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCommunicationLogsRequest) GoString() string {
	return s.String()
}

func (s *QueryCommunicationLogsRequest) SetBizId(v string) *QueryCommunicationLogsRequest {
	s.BizId = &v
	return s
}

func (s *QueryCommunicationLogsRequest) SetType(v int32) *QueryCommunicationLogsRequest {
	s.Type = &v
	return s
}

func (s *QueryCommunicationLogsRequest) SetPageNum(v int32) *QueryCommunicationLogsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCommunicationLogsRequest) SetPageSize(v int32) *QueryCommunicationLogsRequest {
	s.PageSize = &v
	return s
}

type QueryCommunicationLogsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryCommunicationLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryCommunicationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCommunicationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCommunicationLogsResponseBody) SetRequestId(v string) *QueryCommunicationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCommunicationLogsResponseBody) SetData(v *QueryCommunicationLogsResponseBodyData) *QueryCommunicationLogsResponseBody {
	s.Data = v
	return s
}

type QueryCommunicationLogsResponseBodyData struct {
	TaskList []*QueryCommunicationLogsResponseBodyDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s QueryCommunicationLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCommunicationLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCommunicationLogsResponseBodyData) SetTaskList(v []*QueryCommunicationLogsResponseBodyDataTaskList) *QueryCommunicationLogsResponseBodyData {
	s.TaskList = v
	return s
}

type QueryCommunicationLogsResponseBodyDataTaskList struct {
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	PartnerCode *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s QueryCommunicationLogsResponseBodyDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s QueryCommunicationLogsResponseBodyDataTaskList) GoString() string {
	return s.String()
}

func (s *QueryCommunicationLogsResponseBodyDataTaskList) SetNote(v string) *QueryCommunicationLogsResponseBodyDataTaskList {
	s.Note = &v
	return s
}

func (s *QueryCommunicationLogsResponseBodyDataTaskList) SetBizId(v string) *QueryCommunicationLogsResponseBodyDataTaskList {
	s.BizId = &v
	return s
}

func (s *QueryCommunicationLogsResponseBodyDataTaskList) SetUpdateTime(v int64) *QueryCommunicationLogsResponseBodyDataTaskList {
	s.UpdateTime = &v
	return s
}

func (s *QueryCommunicationLogsResponseBodyDataTaskList) SetPartnerCode(v string) *QueryCommunicationLogsResponseBodyDataTaskList {
	s.PartnerCode = &v
	return s
}

func (s *QueryCommunicationLogsResponseBodyDataTaskList) SetCreateTime(v int64) *QueryCommunicationLogsResponseBodyDataTaskList {
	s.CreateTime = &v
	return s
}

type QueryCommunicationLogsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCommunicationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCommunicationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCommunicationLogsResponse) GoString() string {
	return s.String()
}

func (s *QueryCommunicationLogsResponse) SetHeaders(v map[string]*string) *QueryCommunicationLogsResponse {
	s.Headers = v
	return s
}

func (s *QueryCommunicationLogsResponse) SetBody(v *QueryCommunicationLogsResponseBody) *QueryCommunicationLogsResponse {
	s.Body = v
	return s
}

type GenerateQrCodeRequest struct {
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	OssKey   *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
}

func (s GenerateQrCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateQrCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateQrCodeRequest) SetUuid(v string) *GenerateQrCodeRequest {
	s.Uuid = &v
	return s
}

func (s *GenerateQrCodeRequest) SetOssKey(v string) *GenerateQrCodeRequest {
	s.OssKey = &v
	return s
}

func (s *GenerateQrCodeRequest) SetFieldKey(v string) *GenerateQrCodeRequest {
	s.FieldKey = &v
	return s
}

type GenerateQrCodeResponseBody struct {
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	QrcodeUrl  *string `json:"QrcodeUrl,omitempty" xml:"QrcodeUrl,omitempty"`
	FieldKey   *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
}

func (s GenerateQrCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateQrCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateQrCodeResponseBody) SetUuid(v string) *GenerateQrCodeResponseBody {
	s.Uuid = &v
	return s
}

func (s *GenerateQrCodeResponseBody) SetRequestId(v string) *GenerateQrCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateQrCodeResponseBody) SetExpireTime(v int64) *GenerateQrCodeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GenerateQrCodeResponseBody) SetSuccess(v bool) *GenerateQrCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateQrCodeResponseBody) SetQrcodeUrl(v string) *GenerateQrCodeResponseBody {
	s.QrcodeUrl = &v
	return s
}

func (s *GenerateQrCodeResponseBody) SetFieldKey(v string) *GenerateQrCodeResponseBody {
	s.FieldKey = &v
	return s
}

type GenerateQrCodeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateQrCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateQrCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateQrCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateQrCodeResponse) SetHeaders(v map[string]*string) *GenerateQrCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateQrCodeResponse) SetBody(v *GenerateQrCodeResponseBody) *GenerateQrCodeResponse {
	s.Body = v
	return s
}

type ConfirmDissentOriginalRequest struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ContactName     *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactAddress  *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactNumber   *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactProvince *string `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	ContactCity     *string `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	ContactDistrict *string `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	ContactCounty   *string `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
}

func (s ConfirmDissentOriginalRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDissentOriginalRequest) GoString() string {
	return s.String()
}

func (s *ConfirmDissentOriginalRequest) SetBizId(v string) *ConfirmDissentOriginalRequest {
	s.BizId = &v
	return s
}

func (s *ConfirmDissentOriginalRequest) SetContactName(v string) *ConfirmDissentOriginalRequest {
	s.ContactName = &v
	return s
}

func (s *ConfirmDissentOriginalRequest) SetContactAddress(v string) *ConfirmDissentOriginalRequest {
	s.ContactAddress = &v
	return s
}

func (s *ConfirmDissentOriginalRequest) SetContactNumber(v string) *ConfirmDissentOriginalRequest {
	s.ContactNumber = &v
	return s
}

func (s *ConfirmDissentOriginalRequest) SetContactProvince(v string) *ConfirmDissentOriginalRequest {
	s.ContactProvince = &v
	return s
}

func (s *ConfirmDissentOriginalRequest) SetContactCity(v string) *ConfirmDissentOriginalRequest {
	s.ContactCity = &v
	return s
}

func (s *ConfirmDissentOriginalRequest) SetContactDistrict(v string) *ConfirmDissentOriginalRequest {
	s.ContactDistrict = &v
	return s
}

func (s *ConfirmDissentOriginalRequest) SetContactCounty(v string) *ConfirmDissentOriginalRequest {
	s.ContactCounty = &v
	return s
}

type ConfirmDissentOriginalResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfirmDissentOriginalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDissentOriginalResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmDissentOriginalResponseBody) SetRequestId(v string) *ConfirmDissentOriginalResponseBody {
	s.RequestId = &v
	return s
}

type ConfirmDissentOriginalResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfirmDissentOriginalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmDissentOriginalResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDissentOriginalResponse) GoString() string {
	return s.String()
}

func (s *ConfirmDissentOriginalResponse) SetHeaders(v map[string]*string) *ConfirmDissentOriginalResponse {
	s.Headers = v
	return s
}

func (s *ConfirmDissentOriginalResponse) SetBody(v *ConfirmDissentOriginalResponseBody) *ConfirmDissentOriginalResponse {
	s.Body = v
	return s
}

type ConvertImageToGrayRequest struct {
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
}

func (s ConvertImageToGrayRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertImageToGrayRequest) GoString() string {
	return s.String()
}

func (s *ConvertImageToGrayRequest) SetOssKey(v string) *ConvertImageToGrayRequest {
	s.OssKey = &v
	return s
}

type ConvertImageToGrayResponseBody struct {
	SignatureUrl *string `json:"SignatureUrl,omitempty" xml:"SignatureUrl,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertImageToGrayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertImageToGrayResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertImageToGrayResponseBody) SetSignatureUrl(v string) *ConvertImageToGrayResponseBody {
	s.SignatureUrl = &v
	return s
}

func (s *ConvertImageToGrayResponseBody) SetRequestId(v string) *ConvertImageToGrayResponseBody {
	s.RequestId = &v
	return s
}

type ConvertImageToGrayResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConvertImageToGrayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConvertImageToGrayResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertImageToGrayResponse) GoString() string {
	return s.String()
}

func (s *ConvertImageToGrayResponse) SetHeaders(v map[string]*string) *ConvertImageToGrayResponse {
	s.Headers = v
	return s
}

func (s *ConvertImageToGrayResponse) SetBody(v *ConvertImageToGrayResponseBody) *ConvertImageToGrayResponse {
	s.Body = v
	return s
}

type QueryIntentionListRequest struct {
	Type      *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum   *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	SortFiled *string `json:"SortFiled,omitempty" xml:"SortFiled,omitempty"`
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s QueryIntentionListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionListRequest) GoString() string {
	return s.String()
}

func (s *QueryIntentionListRequest) SetType(v int32) *QueryIntentionListRequest {
	s.Type = &v
	return s
}

func (s *QueryIntentionListRequest) SetStatus(v int32) *QueryIntentionListRequest {
	s.Status = &v
	return s
}

func (s *QueryIntentionListRequest) SetPageSize(v int32) *QueryIntentionListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryIntentionListRequest) SetPageNum(v int32) *QueryIntentionListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryIntentionListRequest) SetSortFiled(v string) *QueryIntentionListRequest {
	s.SortFiled = &v
	return s
}

func (s *QueryIntentionListRequest) SetSortOrder(v string) *QueryIntentionListRequest {
	s.SortOrder = &v
	return s
}

type QueryIntentionListResponseBody struct {
	CurrentPageNum *int32                              `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                              `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                              `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryIntentionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryIntentionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIntentionListResponseBody) SetCurrentPageNum(v int32) *QueryIntentionListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryIntentionListResponseBody) SetTotalPageNum(v int32) *QueryIntentionListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryIntentionListResponseBody) SetRequestId(v string) *QueryIntentionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIntentionListResponseBody) SetPageSize(v int32) *QueryIntentionListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryIntentionListResponseBody) SetTotalItemNum(v int32) *QueryIntentionListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryIntentionListResponseBody) SetData(v *QueryIntentionListResponseBodyData) *QueryIntentionListResponseBody {
	s.Data = v
	return s
}

type QueryIntentionListResponseBodyData struct {
	Intention []*QueryIntentionListResponseBodyDataIntention `json:"Intention,omitempty" xml:"Intention,omitempty" type:"Repeated"`
}

func (s QueryIntentionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryIntentionListResponseBodyData) SetIntention(v []*QueryIntentionListResponseBodyDataIntention) *QueryIntentionListResponseBodyData {
	s.Intention = v
	return s
}

type QueryIntentionListResponseBodyDataIntention struct {
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime     *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegisterNumber *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
}

func (s QueryIntentionListResponseBodyDataIntention) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionListResponseBodyDataIntention) GoString() string {
	return s.String()
}

func (s *QueryIntentionListResponseBodyDataIntention) SetType(v int32) *QueryIntentionListResponseBodyDataIntention {
	s.Type = &v
	return s
}

func (s *QueryIntentionListResponseBodyDataIntention) SetStatus(v int32) *QueryIntentionListResponseBodyDataIntention {
	s.Status = &v
	return s
}

func (s *QueryIntentionListResponseBodyDataIntention) SetUpdateTime(v int64) *QueryIntentionListResponseBodyDataIntention {
	s.UpdateTime = &v
	return s
}

func (s *QueryIntentionListResponseBodyDataIntention) SetDescription(v string) *QueryIntentionListResponseBodyDataIntention {
	s.Description = &v
	return s
}

func (s *QueryIntentionListResponseBodyDataIntention) SetRegisterNumber(v string) *QueryIntentionListResponseBodyDataIntention {
	s.RegisterNumber = &v
	return s
}

func (s *QueryIntentionListResponseBodyDataIntention) SetCreateTime(v int64) *QueryIntentionListResponseBodyDataIntention {
	s.CreateTime = &v
	return s
}

func (s *QueryIntentionListResponseBodyDataIntention) SetUserId(v string) *QueryIntentionListResponseBodyDataIntention {
	s.UserId = &v
	return s
}

func (s *QueryIntentionListResponseBodyDataIntention) SetBizId(v string) *QueryIntentionListResponseBodyDataIntention {
	s.BizId = &v
	return s
}

func (s *QueryIntentionListResponseBodyDataIntention) SetClassification(v string) *QueryIntentionListResponseBodyDataIntention {
	s.Classification = &v
	return s
}

type QueryIntentionListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryIntentionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryIntentionListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIntentionListResponse) GoString() string {
	return s.String()
}

func (s *QueryIntentionListResponse) SetHeaders(v map[string]*string) *QueryIntentionListResponse {
	s.Headers = v
	return s
}

func (s *QueryIntentionListResponse) SetBody(v *QueryIntentionListResponseBody) *QueryIntentionListResponse {
	s.Body = v
	return s
}

type GetAuthorizationLetterVersionRequest struct {
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
}

func (s GetAuthorizationLetterVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationLetterVersionRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationLetterVersionRequest) SetOssKey(v string) *GetAuthorizationLetterVersionRequest {
	s.OssKey = &v
	return s
}

type GetAuthorizationLetterVersionResponseBody struct {
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthorizationLetterVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationLetterVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationLetterVersionResponseBody) SetVersion(v string) *GetAuthorizationLetterVersionResponseBody {
	s.Version = &v
	return s
}

func (s *GetAuthorizationLetterVersionResponseBody) SetRequestId(v string) *GetAuthorizationLetterVersionResponseBody {
	s.RequestId = &v
	return s
}

type GetAuthorizationLetterVersionResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuthorizationLetterVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthorizationLetterVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationLetterVersionResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationLetterVersionResponse) SetHeaders(v map[string]*string) *GetAuthorizationLetterVersionResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationLetterVersionResponse) SetBody(v *GetAuthorizationLetterVersionResponseBody) *GetAuthorizationLetterVersionResponse {
	s.Body = v
	return s
}

type QueryTrademarkPriceRequest struct {
	UserId    *int64                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TmName    *string                `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmIcon    *string                `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	Type      *int32                 `json:"Type,omitempty" xml:"Type,omitempty"`
	OrderData map[string]interface{} `json:"OrderData,omitempty" xml:"OrderData,omitempty"`
}

func (s QueryTrademarkPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryTrademarkPriceRequest) SetUserId(v int64) *QueryTrademarkPriceRequest {
	s.UserId = &v
	return s
}

func (s *QueryTrademarkPriceRequest) SetTmName(v string) *QueryTrademarkPriceRequest {
	s.TmName = &v
	return s
}

func (s *QueryTrademarkPriceRequest) SetTmIcon(v string) *QueryTrademarkPriceRequest {
	s.TmIcon = &v
	return s
}

func (s *QueryTrademarkPriceRequest) SetType(v int32) *QueryTrademarkPriceRequest {
	s.Type = &v
	return s
}

func (s *QueryTrademarkPriceRequest) SetOrderData(v map[string]interface{}) *QueryTrademarkPriceRequest {
	s.OrderData = v
	return s
}

type QueryTrademarkPriceShrinkRequest struct {
	UserId          *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TmName          *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmIcon          *string `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	OrderDataShrink *string `json:"OrderData,omitempty" xml:"OrderData,omitempty"`
}

func (s QueryTrademarkPriceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkPriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryTrademarkPriceShrinkRequest) SetUserId(v int64) *QueryTrademarkPriceShrinkRequest {
	s.UserId = &v
	return s
}

func (s *QueryTrademarkPriceShrinkRequest) SetTmName(v string) *QueryTrademarkPriceShrinkRequest {
	s.TmName = &v
	return s
}

func (s *QueryTrademarkPriceShrinkRequest) SetTmIcon(v string) *QueryTrademarkPriceShrinkRequest {
	s.TmIcon = &v
	return s
}

func (s *QueryTrademarkPriceShrinkRequest) SetType(v int32) *QueryTrademarkPriceShrinkRequest {
	s.Type = &v
	return s
}

func (s *QueryTrademarkPriceShrinkRequest) SetOrderDataShrink(v string) *QueryTrademarkPriceShrinkRequest {
	s.OrderDataShrink = &v
	return s
}

type QueryTrademarkPriceResponseBody struct {
	OriginalPrice *float32                               `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DiscountPrice *float32                               `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	Currency      *string                                `json:"Currency,omitempty" xml:"Currency,omitempty"`
	TradePrice    *float32                               `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	Prices        *QueryTrademarkPriceResponseBodyPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Struct"`
}

func (s QueryTrademarkPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTrademarkPriceResponseBody) SetOriginalPrice(v float32) *QueryTrademarkPriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *QueryTrademarkPriceResponseBody) SetRequestId(v string) *QueryTrademarkPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTrademarkPriceResponseBody) SetDiscountPrice(v float32) *QueryTrademarkPriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *QueryTrademarkPriceResponseBody) SetCurrency(v string) *QueryTrademarkPriceResponseBody {
	s.Currency = &v
	return s
}

func (s *QueryTrademarkPriceResponseBody) SetTradePrice(v float32) *QueryTrademarkPriceResponseBody {
	s.TradePrice = &v
	return s
}

func (s *QueryTrademarkPriceResponseBody) SetPrices(v *QueryTrademarkPriceResponseBodyPrices) *QueryTrademarkPriceResponseBody {
	s.Prices = v
	return s
}

type QueryTrademarkPriceResponseBodyPrices struct {
	Prices []*QueryTrademarkPriceResponseBodyPricesPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Repeated"`
}

func (s QueryTrademarkPriceResponseBodyPrices) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkPriceResponseBodyPrices) GoString() string {
	return s.String()
}

func (s *QueryTrademarkPriceResponseBodyPrices) SetPrices(v []*QueryTrademarkPriceResponseBodyPricesPrices) *QueryTrademarkPriceResponseBodyPrices {
	s.Prices = v
	return s
}

type QueryTrademarkPriceResponseBodyPricesPrices struct {
	OriginalPrice      *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	DiscountPrice      *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	Currency           *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	TradePrice         *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	ClassificationCode *string  `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
}

func (s QueryTrademarkPriceResponseBodyPricesPrices) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkPriceResponseBodyPricesPrices) GoString() string {
	return s.String()
}

func (s *QueryTrademarkPriceResponseBodyPricesPrices) SetOriginalPrice(v float32) *QueryTrademarkPriceResponseBodyPricesPrices {
	s.OriginalPrice = &v
	return s
}

func (s *QueryTrademarkPriceResponseBodyPricesPrices) SetDiscountPrice(v float32) *QueryTrademarkPriceResponseBodyPricesPrices {
	s.DiscountPrice = &v
	return s
}

func (s *QueryTrademarkPriceResponseBodyPricesPrices) SetCurrency(v string) *QueryTrademarkPriceResponseBodyPricesPrices {
	s.Currency = &v
	return s
}

func (s *QueryTrademarkPriceResponseBodyPricesPrices) SetTradePrice(v float32) *QueryTrademarkPriceResponseBodyPricesPrices {
	s.TradePrice = &v
	return s
}

func (s *QueryTrademarkPriceResponseBodyPricesPrices) SetClassificationCode(v string) *QueryTrademarkPriceResponseBodyPricesPrices {
	s.ClassificationCode = &v
	return s
}

type QueryTrademarkPriceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTrademarkPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTrademarkPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryTrademarkPriceResponse) SetHeaders(v map[string]*string) *QueryTrademarkPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryTrademarkPriceResponse) SetBody(v *QueryTrademarkPriceResponseBody) *QueryTrademarkPriceResponse {
	s.Body = v
	return s
}

type InsertTmMonitorRuleRequest struct {
	RuleSource     *string                `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	RuleName       *string                `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType       *int32                 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	RuleKeyword    *string                `json:"RuleKeyword,omitempty" xml:"RuleKeyword,omitempty"`
	StartApplyDate *string                `json:"StartApplyDate,omitempty" xml:"StartApplyDate,omitempty"`
	EndApplyDate   *string                `json:"EndApplyDate,omitempty" xml:"EndApplyDate,omitempty"`
	Classification map[string]interface{} `json:"Classification,omitempty" xml:"Classification,omitempty"`
	NotifyStatus   map[string]interface{} `json:"NotifyStatus,omitempty" xml:"NotifyStatus,omitempty"`
}

func (s InsertTmMonitorRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertTmMonitorRuleRequest) GoString() string {
	return s.String()
}

func (s *InsertTmMonitorRuleRequest) SetRuleSource(v string) *InsertTmMonitorRuleRequest {
	s.RuleSource = &v
	return s
}

func (s *InsertTmMonitorRuleRequest) SetRuleName(v string) *InsertTmMonitorRuleRequest {
	s.RuleName = &v
	return s
}

func (s *InsertTmMonitorRuleRequest) SetRuleType(v int32) *InsertTmMonitorRuleRequest {
	s.RuleType = &v
	return s
}

func (s *InsertTmMonitorRuleRequest) SetRuleKeyword(v string) *InsertTmMonitorRuleRequest {
	s.RuleKeyword = &v
	return s
}

func (s *InsertTmMonitorRuleRequest) SetStartApplyDate(v string) *InsertTmMonitorRuleRequest {
	s.StartApplyDate = &v
	return s
}

func (s *InsertTmMonitorRuleRequest) SetEndApplyDate(v string) *InsertTmMonitorRuleRequest {
	s.EndApplyDate = &v
	return s
}

func (s *InsertTmMonitorRuleRequest) SetClassification(v map[string]interface{}) *InsertTmMonitorRuleRequest {
	s.Classification = v
	return s
}

func (s *InsertTmMonitorRuleRequest) SetNotifyStatus(v map[string]interface{}) *InsertTmMonitorRuleRequest {
	s.NotifyStatus = v
	return s
}

type InsertTmMonitorRuleShrinkRequest struct {
	RuleSource           *string `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	RuleName             *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType             *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	RuleKeyword          *string `json:"RuleKeyword,omitempty" xml:"RuleKeyword,omitempty"`
	StartApplyDate       *string `json:"StartApplyDate,omitempty" xml:"StartApplyDate,omitempty"`
	EndApplyDate         *string `json:"EndApplyDate,omitempty" xml:"EndApplyDate,omitempty"`
	ClassificationShrink *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	NotifyStatusShrink   *string `json:"NotifyStatus,omitempty" xml:"NotifyStatus,omitempty"`
}

func (s InsertTmMonitorRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertTmMonitorRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertTmMonitorRuleShrinkRequest) SetRuleSource(v string) *InsertTmMonitorRuleShrinkRequest {
	s.RuleSource = &v
	return s
}

func (s *InsertTmMonitorRuleShrinkRequest) SetRuleName(v string) *InsertTmMonitorRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *InsertTmMonitorRuleShrinkRequest) SetRuleType(v int32) *InsertTmMonitorRuleShrinkRequest {
	s.RuleType = &v
	return s
}

func (s *InsertTmMonitorRuleShrinkRequest) SetRuleKeyword(v string) *InsertTmMonitorRuleShrinkRequest {
	s.RuleKeyword = &v
	return s
}

func (s *InsertTmMonitorRuleShrinkRequest) SetStartApplyDate(v string) *InsertTmMonitorRuleShrinkRequest {
	s.StartApplyDate = &v
	return s
}

func (s *InsertTmMonitorRuleShrinkRequest) SetEndApplyDate(v string) *InsertTmMonitorRuleShrinkRequest {
	s.EndApplyDate = &v
	return s
}

func (s *InsertTmMonitorRuleShrinkRequest) SetClassificationShrink(v string) *InsertTmMonitorRuleShrinkRequest {
	s.ClassificationShrink = &v
	return s
}

func (s *InsertTmMonitorRuleShrinkRequest) SetNotifyStatusShrink(v string) *InsertTmMonitorRuleShrinkRequest {
	s.NotifyStatusShrink = &v
	return s
}

type InsertTmMonitorRuleResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s InsertTmMonitorRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertTmMonitorRuleResponseBody) GoString() string {
	return s.String()
}

func (s *InsertTmMonitorRuleResponseBody) SetErrorMsg(v string) *InsertTmMonitorRuleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InsertTmMonitorRuleResponseBody) SetRequestId(v string) *InsertTmMonitorRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertTmMonitorRuleResponseBody) SetSuccess(v bool) *InsertTmMonitorRuleResponseBody {
	s.Success = &v
	return s
}

func (s *InsertTmMonitorRuleResponseBody) SetErrorCode(v string) *InsertTmMonitorRuleResponseBody {
	s.ErrorCode = &v
	return s
}

type InsertTmMonitorRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertTmMonitorRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertTmMonitorRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertTmMonitorRuleResponse) GoString() string {
	return s.String()
}

func (s *InsertTmMonitorRuleResponse) SetHeaders(v map[string]*string) *InsertTmMonitorRuleResponse {
	s.Headers = v
	return s
}

func (s *InsertTmMonitorRuleResponse) SetBody(v *InsertTmMonitorRuleResponseBody) *InsertTmMonitorRuleResponse {
	s.Body = v
	return s
}

type QueryTrademarkMonitorRulesRequest struct {
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	NotifyUpdate *int32  `json:"NotifyUpdate,omitempty" xml:"NotifyUpdate,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryTrademarkMonitorRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorRulesRequest) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorRulesRequest) SetId(v string) *QueryTrademarkMonitorRulesRequest {
	s.Id = &v
	return s
}

func (s *QueryTrademarkMonitorRulesRequest) SetRuleName(v string) *QueryTrademarkMonitorRulesRequest {
	s.RuleName = &v
	return s
}

func (s *QueryTrademarkMonitorRulesRequest) SetNotifyUpdate(v int32) *QueryTrademarkMonitorRulesRequest {
	s.NotifyUpdate = &v
	return s
}

func (s *QueryTrademarkMonitorRulesRequest) SetPageNum(v int32) *QueryTrademarkMonitorRulesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTrademarkMonitorRulesRequest) SetPageSize(v int32) *QueryTrademarkMonitorRulesRequest {
	s.PageSize = &v
	return s
}

type QueryTrademarkMonitorRulesResponseBody struct {
	NextPage       *bool                                       `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrePage        *bool                                       `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	TotalItemNum   *int32                                      `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	CurrentPageNum *int32                                      `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                                      `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	PageSize       *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data           *QueryTrademarkMonitorRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTrademarkMonitorRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorRulesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorRulesResponseBody) SetNextPage(v bool) *QueryTrademarkMonitorRulesResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBody) SetRequestId(v string) *QueryTrademarkMonitorRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBody) SetPrePage(v bool) *QueryTrademarkMonitorRulesResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBody) SetTotalItemNum(v int32) *QueryTrademarkMonitorRulesResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBody) SetCurrentPageNum(v int32) *QueryTrademarkMonitorRulesResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBody) SetTotalPageNum(v int32) *QueryTrademarkMonitorRulesResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBody) SetPageSize(v int32) *QueryTrademarkMonitorRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBody) SetData(v *QueryTrademarkMonitorRulesResponseBodyData) *QueryTrademarkMonitorRulesResponseBody {
	s.Data = v
	return s
}

type QueryTrademarkMonitorRulesResponseBodyData struct {
	TmMonitorRule []*QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule `json:"TmMonitorRule,omitempty" xml:"TmMonitorRule,omitempty" type:"Repeated"`
}

func (s QueryTrademarkMonitorRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorRulesResponseBodyData) SetTmMonitorRule(v []*QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) *QueryTrademarkMonitorRulesResponseBodyData {
	s.TmMonitorRule = v
	return s
}

type QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule struct {
	RuleStatus     *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	LastFinishTime *string `json:"LastFinishTime,omitempty" xml:"LastFinishTime,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	RuleType       *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RuleExtend     *string `json:"RuleExtend,omitempty" xml:"RuleExtend,omitempty"`
	RuleName       *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RuleKeyword    *string `json:"RuleKeyword,omitempty" xml:"RuleKeyword,omitempty"`
	LastRunTime    *string `json:"LastRunTime,omitempty" xml:"LastRunTime,omitempty"`
	Version        *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	RuleSource     *string `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	Env            *string `json:"Env,omitempty" xml:"Env,omitempty"`
	NotifyUpdate   *int32  `json:"NotifyUpdate,omitempty" xml:"NotifyUpdate,omitempty"`
	RuleDetail     *string `json:"RuleDetail,omitempty" xml:"RuleDetail,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetRuleStatus(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.RuleStatus = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetLastFinishTime(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.LastFinishTime = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetUpdateTime(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.UpdateTime = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetRuleType(v int32) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.RuleType = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetCreateTime(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.CreateTime = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetUserId(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.UserId = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetRuleExtend(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.RuleExtend = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetRuleName(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.RuleName = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetEndTime(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.EndTime = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetStartTime(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.StartTime = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetRuleKeyword(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.RuleKeyword = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetLastRunTime(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.LastRunTime = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetVersion(v int32) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.Version = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetRuleSource(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.RuleSource = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetLastUpdateTime(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.LastUpdateTime = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetEnv(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.Env = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetNotifyUpdate(v int32) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.NotifyUpdate = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetRuleDetail(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.RuleDetail = &v
	return s
}

func (s *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule) SetId(v string) *QueryTrademarkMonitorRulesResponseBodyDataTmMonitorRule {
	s.Id = &v
	return s
}

type QueryTrademarkMonitorRulesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTrademarkMonitorRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTrademarkMonitorRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTrademarkMonitorRulesResponse) GoString() string {
	return s.String()
}

func (s *QueryTrademarkMonitorRulesResponse) SetHeaders(v map[string]*string) *QueryTrademarkMonitorRulesResponse {
	s.Headers = v
	return s
}

func (s *QueryTrademarkMonitorRulesResponse) SetBody(v *QueryTrademarkMonitorRulesResponseBody) *QueryTrademarkMonitorRulesResponse {
	s.Body = v
	return s
}

type DenySupplementRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DenySupplementRequest) String() string {
	return tea.Prettify(s)
}

func (s DenySupplementRequest) GoString() string {
	return s.String()
}

func (s *DenySupplementRequest) SetId(v int64) *DenySupplementRequest {
	s.Id = &v
	return s
}

type DenySupplementResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s DenySupplementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DenySupplementResponseBody) GoString() string {
	return s.String()
}

func (s *DenySupplementResponseBody) SetErrorMsg(v string) *DenySupplementResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DenySupplementResponseBody) SetRequestId(v string) *DenySupplementResponseBody {
	s.RequestId = &v
	return s
}

func (s *DenySupplementResponseBody) SetSuccess(v bool) *DenySupplementResponseBody {
	s.Success = &v
	return s
}

func (s *DenySupplementResponseBody) SetErrorCode(v string) *DenySupplementResponseBody {
	s.ErrorCode = &v
	return s
}

type DenySupplementResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DenySupplementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DenySupplementResponse) String() string {
	return tea.Prettify(s)
}

func (s DenySupplementResponse) GoString() string {
	return s.String()
}

func (s *DenySupplementResponse) SetHeaders(v map[string]*string) *DenySupplementResponse {
	s.Headers = v
	return s
}

func (s *DenySupplementResponse) SetBody(v *DenySupplementResponseBody) *DenySupplementResponse {
	s.Body = v
	return s
}

type QueryMaterialRequest struct {
	Id                   *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	QueryUnconfirmedInfo *bool  `json:"QueryUnconfirmedInfo,omitempty" xml:"QueryUnconfirmedInfo,omitempty"`
}

func (s QueryMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialRequest) SetId(v int64) *QueryMaterialRequest {
	s.Id = &v
	return s
}

func (s *QueryMaterialRequest) SetQueryUnconfirmedInfo(v bool) *QueryMaterialRequest {
	s.QueryUnconfirmedInfo = &v
	return s
}

type QueryMaterialResponseBody struct {
	Type                  *int32                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	Status                *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	ReviewApplicationFile *string                                         `json:"ReviewApplicationFile,omitempty" xml:"ReviewApplicationFile,omitempty"`
	ContactDistrict       *string                                         `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	BusinessLicenceUrl    *string                                         `json:"BusinessLicenceUrl,omitempty" xml:"BusinessLicenceUrl,omitempty"`
	PassportUrl           *string                                         `json:"PassportUrl,omitempty" xml:"PassportUrl,omitempty"`
	ContactProvince       *string                                         `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	LegalNoticeUrl        *string                                         `json:"LegalNoticeUrl,omitempty" xml:"LegalNoticeUrl,omitempty"`
	City                  *string                                         `json:"City,omitempty" xml:"City,omitempty"`
	EAddress              *string                                         `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	ContactCounty         *string                                         `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
	ContactEmail          *string                                         `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	RequestId             *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContactCity           *string                                         `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	Region                *int32                                          `json:"Region,omitempty" xml:"Region,omitempty"`
	LoaUrl                *string                                         `json:"LoaUrl,omitempty" xml:"LoaUrl,omitempty"`
	Address               *string                                         `json:"Address,omitempty" xml:"Address,omitempty"`
	Note                  *string                                         `json:"Note,omitempty" xml:"Note,omitempty"`
	PrincipalName         *int32                                          `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Name                  *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PrincipalDescription  *string                                         `json:"PrincipalDescription,omitempty" xml:"PrincipalDescription,omitempty"`
	ContactNumber         *string                                         `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactAddress        *string                                         `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactZipcode        *string                                         `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	ContactName           *string                                         `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	EName                 *string                                         `json:"EName,omitempty" xml:"EName,omitempty"`
	ValidDate             *int64                                          `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
	IdCardUrl             *string                                         `json:"IdCardUrl,omitempty" xml:"IdCardUrl,omitempty"`
	ExpirationDate        *int64                                          `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	CardNumber            *string                                         `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Country               *string                                         `json:"Country,omitempty" xml:"Country,omitempty"`
	Town                  *string                                         `json:"Town,omitempty" xml:"Town,omitempty"`
	LoaStatus             *int32                                          `json:"LoaStatus,omitempty" xml:"LoaStatus,omitempty"`
	Reason                *string                                         `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Id                    *int64                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	Province              *string                                         `json:"Province,omitempty" xml:"Province,omitempty"`
	LegalNoticeKey        *string                                         `json:"LegalNoticeKey,omitempty" xml:"LegalNoticeKey,omitempty"`
	ReviewAdditionalFiles *QueryMaterialResponseBodyReviewAdditionalFiles `json:"ReviewAdditionalFiles,omitempty" xml:"ReviewAdditionalFiles,omitempty" type:"Struct"`
}

func (s QueryMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMaterialResponseBody) SetType(v int32) *QueryMaterialResponseBody {
	s.Type = &v
	return s
}

func (s *QueryMaterialResponseBody) SetStatus(v int32) *QueryMaterialResponseBody {
	s.Status = &v
	return s
}

func (s *QueryMaterialResponseBody) SetReviewApplicationFile(v string) *QueryMaterialResponseBody {
	s.ReviewApplicationFile = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactDistrict(v string) *QueryMaterialResponseBody {
	s.ContactDistrict = &v
	return s
}

func (s *QueryMaterialResponseBody) SetBusinessLicenceUrl(v string) *QueryMaterialResponseBody {
	s.BusinessLicenceUrl = &v
	return s
}

func (s *QueryMaterialResponseBody) SetPassportUrl(v string) *QueryMaterialResponseBody {
	s.PassportUrl = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactProvince(v string) *QueryMaterialResponseBody {
	s.ContactProvince = &v
	return s
}

func (s *QueryMaterialResponseBody) SetLegalNoticeUrl(v string) *QueryMaterialResponseBody {
	s.LegalNoticeUrl = &v
	return s
}

func (s *QueryMaterialResponseBody) SetCity(v string) *QueryMaterialResponseBody {
	s.City = &v
	return s
}

func (s *QueryMaterialResponseBody) SetEAddress(v string) *QueryMaterialResponseBody {
	s.EAddress = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactCounty(v string) *QueryMaterialResponseBody {
	s.ContactCounty = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactEmail(v string) *QueryMaterialResponseBody {
	s.ContactEmail = &v
	return s
}

func (s *QueryMaterialResponseBody) SetRequestId(v string) *QueryMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactCity(v string) *QueryMaterialResponseBody {
	s.ContactCity = &v
	return s
}

func (s *QueryMaterialResponseBody) SetRegion(v int32) *QueryMaterialResponseBody {
	s.Region = &v
	return s
}

func (s *QueryMaterialResponseBody) SetLoaUrl(v string) *QueryMaterialResponseBody {
	s.LoaUrl = &v
	return s
}

func (s *QueryMaterialResponseBody) SetAddress(v string) *QueryMaterialResponseBody {
	s.Address = &v
	return s
}

func (s *QueryMaterialResponseBody) SetNote(v string) *QueryMaterialResponseBody {
	s.Note = &v
	return s
}

func (s *QueryMaterialResponseBody) SetPrincipalName(v int32) *QueryMaterialResponseBody {
	s.PrincipalName = &v
	return s
}

func (s *QueryMaterialResponseBody) SetName(v string) *QueryMaterialResponseBody {
	s.Name = &v
	return s
}

func (s *QueryMaterialResponseBody) SetPrincipalDescription(v string) *QueryMaterialResponseBody {
	s.PrincipalDescription = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactNumber(v string) *QueryMaterialResponseBody {
	s.ContactNumber = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactAddress(v string) *QueryMaterialResponseBody {
	s.ContactAddress = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactZipcode(v string) *QueryMaterialResponseBody {
	s.ContactZipcode = &v
	return s
}

func (s *QueryMaterialResponseBody) SetContactName(v string) *QueryMaterialResponseBody {
	s.ContactName = &v
	return s
}

func (s *QueryMaterialResponseBody) SetEName(v string) *QueryMaterialResponseBody {
	s.EName = &v
	return s
}

func (s *QueryMaterialResponseBody) SetValidDate(v int64) *QueryMaterialResponseBody {
	s.ValidDate = &v
	return s
}

func (s *QueryMaterialResponseBody) SetIdCardUrl(v string) *QueryMaterialResponseBody {
	s.IdCardUrl = &v
	return s
}

func (s *QueryMaterialResponseBody) SetExpirationDate(v int64) *QueryMaterialResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryMaterialResponseBody) SetCardNumber(v string) *QueryMaterialResponseBody {
	s.CardNumber = &v
	return s
}

func (s *QueryMaterialResponseBody) SetCountry(v string) *QueryMaterialResponseBody {
	s.Country = &v
	return s
}

func (s *QueryMaterialResponseBody) SetTown(v string) *QueryMaterialResponseBody {
	s.Town = &v
	return s
}

func (s *QueryMaterialResponseBody) SetLoaStatus(v int32) *QueryMaterialResponseBody {
	s.LoaStatus = &v
	return s
}

func (s *QueryMaterialResponseBody) SetReason(v string) *QueryMaterialResponseBody {
	s.Reason = &v
	return s
}

func (s *QueryMaterialResponseBody) SetId(v int64) *QueryMaterialResponseBody {
	s.Id = &v
	return s
}

func (s *QueryMaterialResponseBody) SetProvince(v string) *QueryMaterialResponseBody {
	s.Province = &v
	return s
}

func (s *QueryMaterialResponseBody) SetLegalNoticeKey(v string) *QueryMaterialResponseBody {
	s.LegalNoticeKey = &v
	return s
}

func (s *QueryMaterialResponseBody) SetReviewAdditionalFiles(v *QueryMaterialResponseBodyReviewAdditionalFiles) *QueryMaterialResponseBody {
	s.ReviewAdditionalFiles = v
	return s
}

type QueryMaterialResponseBodyReviewAdditionalFiles struct {
	ReviewAdditionalFile []*string `json:"ReviewAdditionalFile,omitempty" xml:"ReviewAdditionalFile,omitempty" type:"Repeated"`
}

func (s QueryMaterialResponseBodyReviewAdditionalFiles) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialResponseBodyReviewAdditionalFiles) GoString() string {
	return s.String()
}

func (s *QueryMaterialResponseBodyReviewAdditionalFiles) SetReviewAdditionalFile(v []*string) *QueryMaterialResponseBodyReviewAdditionalFiles {
	s.ReviewAdditionalFile = v
	return s
}

type QueryMaterialResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialResponse) GoString() string {
	return s.String()
}

func (s *QueryMaterialResponse) SetHeaders(v map[string]*string) *QueryMaterialResponse {
	s.Headers = v
	return s
}

func (s *QueryMaterialResponse) SetBody(v *QueryMaterialResponseBody) *QueryMaterialResponse {
	s.Body = v
	return s
}

type CreateTrademarkOrderRequest struct {
	UserId          *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TmName          *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmIcon          *string `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	OrderData       *string `json:"OrderData,omitempty" xml:"OrderData,omitempty"`
	MaterialId      *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	LoaOssKey       *string `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	IsBlackIcon     *bool   `json:"IsBlackIcon,omitempty" xml:"IsBlackIcon,omitempty"`
	RenewInfoId     *string `json:"RenewInfoId,omitempty" xml:"RenewInfoId,omitempty"`
	RootCode        *string `json:"RootCode,omitempty" xml:"RootCode,omitempty"`
	Channel         *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	RegisterNumber  *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	TmNameType      *string `json:"TmNameType,omitempty" xml:"TmNameType,omitempty"`
	RegisterName    *string `json:"RegisterName,omitempty" xml:"RegisterName,omitempty"`
	TmComment       *string `json:"TmComment,omitempty" xml:"TmComment,omitempty"`
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Uid             *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	PartnerCode     *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	RealUserName    *string `json:"RealUserName,omitempty" xml:"RealUserName,omitempty"`
	PhoneNum        *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	PrincipalName   *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	BigDipperSource *string `json:"BigDipperSource,omitempty" xml:"BigDipperSource,omitempty"`
	Ua              *string `json:"Ua,omitempty" xml:"Ua,omitempty"`
	LegalNoticeKey  *string `json:"LegalNoticeKey,omitempty" xml:"LegalNoticeKey,omitempty"`
}

func (s CreateTrademarkOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrademarkOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateTrademarkOrderRequest) SetUserId(v int64) *CreateTrademarkOrderRequest {
	s.UserId = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetTmName(v string) *CreateTrademarkOrderRequest {
	s.TmName = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetTmIcon(v string) *CreateTrademarkOrderRequest {
	s.TmIcon = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetType(v int32) *CreateTrademarkOrderRequest {
	s.Type = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetOrderData(v string) *CreateTrademarkOrderRequest {
	s.OrderData = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetMaterialId(v string) *CreateTrademarkOrderRequest {
	s.MaterialId = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetLoaOssKey(v string) *CreateTrademarkOrderRequest {
	s.LoaOssKey = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetIsBlackIcon(v bool) *CreateTrademarkOrderRequest {
	s.IsBlackIcon = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetRenewInfoId(v string) *CreateTrademarkOrderRequest {
	s.RenewInfoId = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetRootCode(v string) *CreateTrademarkOrderRequest {
	s.RootCode = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetChannel(v string) *CreateTrademarkOrderRequest {
	s.Channel = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetRegisterNumber(v string) *CreateTrademarkOrderRequest {
	s.RegisterNumber = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetTmNameType(v string) *CreateTrademarkOrderRequest {
	s.TmNameType = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetRegisterName(v string) *CreateTrademarkOrderRequest {
	s.RegisterName = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetTmComment(v string) *CreateTrademarkOrderRequest {
	s.TmComment = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetBizId(v string) *CreateTrademarkOrderRequest {
	s.BizId = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetUid(v string) *CreateTrademarkOrderRequest {
	s.Uid = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetPartnerCode(v string) *CreateTrademarkOrderRequest {
	s.PartnerCode = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetRealUserName(v string) *CreateTrademarkOrderRequest {
	s.RealUserName = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetPhoneNum(v string) *CreateTrademarkOrderRequest {
	s.PhoneNum = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetPrincipalName(v int32) *CreateTrademarkOrderRequest {
	s.PrincipalName = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetBigDipperSource(v string) *CreateTrademarkOrderRequest {
	s.BigDipperSource = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetUa(v string) *CreateTrademarkOrderRequest {
	s.Ua = &v
	return s
}

func (s *CreateTrademarkOrderRequest) SetLegalNoticeKey(v string) *CreateTrademarkOrderRequest {
	s.LegalNoticeKey = &v
	return s
}

type CreateTrademarkOrderResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateTrademarkOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrademarkOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrademarkOrderResponseBody) SetErrorMsg(v string) *CreateTrademarkOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateTrademarkOrderResponseBody) SetRequestId(v string) *CreateTrademarkOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrademarkOrderResponseBody) SetSuccess(v bool) *CreateTrademarkOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTrademarkOrderResponseBody) SetOrderId(v int64) *CreateTrademarkOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateTrademarkOrderResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTrademarkOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrademarkOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrademarkOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateTrademarkOrderResponse) SetHeaders(v map[string]*string) *CreateTrademarkOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateTrademarkOrderResponse) SetBody(v *CreateTrademarkOrderResponseBody) *CreateTrademarkOrderResponse {
	s.Body = v
	return s
}

type QueryMaterialListRequest struct {
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type          *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Region        *int32  `json:"Region,omitempty" xml:"Region,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNum       *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CardNumber    *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	PrincipalName *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	MaterialId    *int64  `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
}

func (s QueryMaterialListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialListRequest) GoString() string {
	return s.String()
}

func (s *QueryMaterialListRequest) SetName(v string) *QueryMaterialListRequest {
	s.Name = &v
	return s
}

func (s *QueryMaterialListRequest) SetType(v int32) *QueryMaterialListRequest {
	s.Type = &v
	return s
}

func (s *QueryMaterialListRequest) SetRegion(v int32) *QueryMaterialListRequest {
	s.Region = &v
	return s
}

func (s *QueryMaterialListRequest) SetStatus(v int32) *QueryMaterialListRequest {
	s.Status = &v
	return s
}

func (s *QueryMaterialListRequest) SetPageNum(v int32) *QueryMaterialListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMaterialListRequest) SetPageSize(v int32) *QueryMaterialListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialListRequest) SetCardNumber(v string) *QueryMaterialListRequest {
	s.CardNumber = &v
	return s
}

func (s *QueryMaterialListRequest) SetPrincipalName(v int32) *QueryMaterialListRequest {
	s.PrincipalName = &v
	return s
}

func (s *QueryMaterialListRequest) SetMaterialId(v int64) *QueryMaterialListRequest {
	s.MaterialId = &v
	return s
}

type QueryMaterialListResponseBody struct {
	CurrentPageNum *int32                             `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                             `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                             `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryMaterialListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryMaterialListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMaterialListResponseBody) SetCurrentPageNum(v int32) *QueryMaterialListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryMaterialListResponseBody) SetTotalPageNum(v int32) *QueryMaterialListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryMaterialListResponseBody) SetRequestId(v string) *QueryMaterialListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMaterialListResponseBody) SetPageSize(v int32) *QueryMaterialListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMaterialListResponseBody) SetTotalItemNum(v int32) *QueryMaterialListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryMaterialListResponseBody) SetData(v *QueryMaterialListResponseBodyData) *QueryMaterialListResponseBody {
	s.Data = v
	return s
}

type QueryMaterialListResponseBodyData struct {
	Trademark []*QueryMaterialListResponseBodyDataTrademark `json:"Trademark,omitempty" xml:"Trademark,omitempty" type:"Repeated"`
}

func (s QueryMaterialListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMaterialListResponseBodyData) SetTrademark(v []*QueryMaterialListResponseBodyDataTrademark) *QueryMaterialListResponseBodyData {
	s.Trademark = v
	return s
}

type QueryMaterialListResponseBodyDataTrademark struct {
	PrincipalDescription *string `json:"PrincipalDescription,omitempty" xml:"PrincipalDescription,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                 *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	ContactName          *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CardNumber           *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ValidDate            *int64  `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
	Region               *int32  `json:"Region,omitempty" xml:"Region,omitempty"`
	PrincipalName        *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	LoaStatus            *int32  `json:"LoaStatus,omitempty" xml:"LoaStatus,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	LoaKey               *string `json:"LoaKey,omitempty" xml:"LoaKey,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Reason               *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s QueryMaterialListResponseBodyDataTrademark) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialListResponseBodyDataTrademark) GoString() string {
	return s.String()
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetPrincipalDescription(v string) *QueryMaterialListResponseBodyDataTrademark {
	s.PrincipalDescription = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetStatus(v int32) *QueryMaterialListResponseBodyDataTrademark {
	s.Status = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetType(v int32) *QueryMaterialListResponseBodyDataTrademark {
	s.Type = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetContactName(v string) *QueryMaterialListResponseBodyDataTrademark {
	s.ContactName = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetCardNumber(v string) *QueryMaterialListResponseBodyDataTrademark {
	s.CardNumber = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetValidDate(v int64) *QueryMaterialListResponseBodyDataTrademark {
	s.ValidDate = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetRegion(v int32) *QueryMaterialListResponseBodyDataTrademark {
	s.Region = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetPrincipalName(v int32) *QueryMaterialListResponseBodyDataTrademark {
	s.PrincipalName = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetLoaStatus(v int32) *QueryMaterialListResponseBodyDataTrademark {
	s.LoaStatus = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetName(v string) *QueryMaterialListResponseBodyDataTrademark {
	s.Name = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetLoaKey(v string) *QueryMaterialListResponseBodyDataTrademark {
	s.LoaKey = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetId(v int64) *QueryMaterialListResponseBodyDataTrademark {
	s.Id = &v
	return s
}

func (s *QueryMaterialListResponseBodyDataTrademark) SetReason(v string) *QueryMaterialListResponseBodyDataTrademark {
	s.Reason = &v
	return s
}

type QueryMaterialListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMaterialListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMaterialListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMaterialListResponse) GoString() string {
	return s.String()
}

func (s *QueryMaterialListResponse) SetHeaders(v map[string]*string) *QueryMaterialListResponse {
	s.Headers = v
	return s
}

func (s *QueryMaterialListResponse) SetBody(v *QueryMaterialListResponseBody) *QueryMaterialListResponse {
	s.Body = v
	return s
}

type CheckTrademarkOrderRequest struct {
	UserId         *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	TmName         *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmIcon         *string `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	OrderData      *string `json:"OrderData,omitempty" xml:"OrderData,omitempty"`
	MaterialId     *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	LoaOssKey      *string `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	IsBlackIcon    *bool   `json:"IsBlackIcon,omitempty" xml:"IsBlackIcon,omitempty"`
	RenewInfoId    *string `json:"RenewInfoId,omitempty" xml:"RenewInfoId,omitempty"`
	RootCode       *string `json:"RootCode,omitempty" xml:"RootCode,omitempty"`
	Channel        *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	RegisterNumber *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	TmNameType     *string `json:"TmNameType,omitempty" xml:"TmNameType,omitempty"`
	RegisterName   *string `json:"RegisterName,omitempty" xml:"RegisterName,omitempty"`
	TmComment      *string `json:"TmComment,omitempty" xml:"TmComment,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Uid            *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	PartnerCode    *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	RealUserName   *string `json:"RealUserName,omitempty" xml:"RealUserName,omitempty"`
	PhoneNum       *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	LogoGoodsId    *string `json:"LogoGoodsId,omitempty" xml:"LogoGoodsId,omitempty"`
}

func (s CheckTrademarkOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkOrderRequest) GoString() string {
	return s.String()
}

func (s *CheckTrademarkOrderRequest) SetUserId(v int64) *CheckTrademarkOrderRequest {
	s.UserId = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetTmName(v string) *CheckTrademarkOrderRequest {
	s.TmName = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetTmIcon(v string) *CheckTrademarkOrderRequest {
	s.TmIcon = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetType(v int32) *CheckTrademarkOrderRequest {
	s.Type = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetOrderData(v string) *CheckTrademarkOrderRequest {
	s.OrderData = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetMaterialId(v string) *CheckTrademarkOrderRequest {
	s.MaterialId = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetLoaOssKey(v string) *CheckTrademarkOrderRequest {
	s.LoaOssKey = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetIsBlackIcon(v bool) *CheckTrademarkOrderRequest {
	s.IsBlackIcon = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetRenewInfoId(v string) *CheckTrademarkOrderRequest {
	s.RenewInfoId = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetRootCode(v string) *CheckTrademarkOrderRequest {
	s.RootCode = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetChannel(v string) *CheckTrademarkOrderRequest {
	s.Channel = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetRegisterNumber(v string) *CheckTrademarkOrderRequest {
	s.RegisterNumber = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetTmNameType(v string) *CheckTrademarkOrderRequest {
	s.TmNameType = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetRegisterName(v string) *CheckTrademarkOrderRequest {
	s.RegisterName = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetTmComment(v string) *CheckTrademarkOrderRequest {
	s.TmComment = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetBizId(v string) *CheckTrademarkOrderRequest {
	s.BizId = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetUid(v string) *CheckTrademarkOrderRequest {
	s.Uid = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetPartnerCode(v string) *CheckTrademarkOrderRequest {
	s.PartnerCode = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetRealUserName(v string) *CheckTrademarkOrderRequest {
	s.RealUserName = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetPhoneNum(v string) *CheckTrademarkOrderRequest {
	s.PhoneNum = &v
	return s
}

func (s *CheckTrademarkOrderRequest) SetLogoGoodsId(v string) *CheckTrademarkOrderRequest {
	s.LogoGoodsId = &v
	return s
}

type CheckTrademarkOrderResponseBody struct {
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckTrademarkOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTrademarkOrderResponseBody) SetData(v map[string]interface{}) *CheckTrademarkOrderResponseBody {
	s.Data = v
	return s
}

func (s *CheckTrademarkOrderResponseBody) SetErrorMsg(v string) *CheckTrademarkOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CheckTrademarkOrderResponseBody) SetRequestId(v string) *CheckTrademarkOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTrademarkOrderResponseBody) SetSuccess(v bool) *CheckTrademarkOrderResponseBody {
	s.Success = &v
	return s
}

type CheckTrademarkOrderResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckTrademarkOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckTrademarkOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkOrderResponse) GoString() string {
	return s.String()
}

func (s *CheckTrademarkOrderResponse) SetHeaders(v map[string]*string) *CheckTrademarkOrderResponse {
	s.Headers = v
	return s
}

func (s *CheckTrademarkOrderResponse) SetBody(v *CheckTrademarkOrderResponseBody) *CheckTrademarkOrderResponse {
	s.Body = v
	return s
}

type QueryTradeMarkApplicationsRequest struct {
	TmName             *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	PageNum            *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	MaterialName       *string `json:"MaterialName,omitempty" xml:"MaterialName,omitempty"`
	TmNumber           *string `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplementStatus   *int32  `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	SortOrder          *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId              *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IntentionBizId     *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Hidden             *int32  `json:"Hidden,omitempty" xml:"Hidden,omitempty"`
	ProductType        *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	LogisticsNo        *string `json:"LogisticsNo,omitempty" xml:"LogisticsNo,omitempty"`
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	SortFiled          *string `json:"SortFiled,omitempty" xml:"SortFiled,omitempty"`
}

func (s QueryTradeMarkApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsRequest) SetTmName(v string) *QueryTradeMarkApplicationsRequest {
	s.TmName = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetPageNum(v int32) *QueryTradeMarkApplicationsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetPageSize(v int32) *QueryTradeMarkApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetMaterialName(v string) *QueryTradeMarkApplicationsRequest {
	s.MaterialName = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetTmNumber(v string) *QueryTradeMarkApplicationsRequest {
	s.TmNumber = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetOrderId(v string) *QueryTradeMarkApplicationsRequest {
	s.OrderId = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetStatus(v int32) *QueryTradeMarkApplicationsRequest {
	s.Status = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetSupplementStatus(v int32) *QueryTradeMarkApplicationsRequest {
	s.SupplementStatus = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetSortOrder(v string) *QueryTradeMarkApplicationsRequest {
	s.SortOrder = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetType(v string) *QueryTradeMarkApplicationsRequest {
	s.Type = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetBizId(v string) *QueryTradeMarkApplicationsRequest {
	s.BizId = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetIntentionBizId(v string) *QueryTradeMarkApplicationsRequest {
	s.IntentionBizId = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetHidden(v int32) *QueryTradeMarkApplicationsRequest {
	s.Hidden = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetProductType(v int32) *QueryTradeMarkApplicationsRequest {
	s.ProductType = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetLogisticsNo(v string) *QueryTradeMarkApplicationsRequest {
	s.LogisticsNo = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetClassificationCode(v string) *QueryTradeMarkApplicationsRequest {
	s.ClassificationCode = &v
	return s
}

func (s *QueryTradeMarkApplicationsRequest) SetSortFiled(v string) *QueryTradeMarkApplicationsRequest {
	s.SortFiled = &v
	return s
}

type QueryTradeMarkApplicationsResponseBody struct {
	CurrentPageNum *int32                                      `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                                      `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                                      `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryTradeMarkApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponseBody) SetCurrentPageNum(v int32) *QueryTradeMarkApplicationsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBody) SetTotalPageNum(v int32) *QueryTradeMarkApplicationsResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBody) SetRequestId(v string) *QueryTradeMarkApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBody) SetPageSize(v int32) *QueryTradeMarkApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBody) SetTotalItemNum(v int32) *QueryTradeMarkApplicationsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBody) SetData(v *QueryTradeMarkApplicationsResponseBodyData) *QueryTradeMarkApplicationsResponseBody {
	s.Data = v
	return s
}

type QueryTradeMarkApplicationsResponseBodyData struct {
	TmProduces []*QueryTradeMarkApplicationsResponseBodyDataTmProduces `json:"TmProduces,omitempty" xml:"TmProduces,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponseBodyData) SetTmProduces(v []*QueryTradeMarkApplicationsResponseBodyDataTmProduces) *QueryTradeMarkApplicationsResponseBodyData {
	s.TmProduces = v
	return s
}

type QueryTradeMarkApplicationsResponseBodyDataTmProduces struct {
	Type                *int32                                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Status              *int32                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	OrderPrice          *float32                                                                 `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	SubmitAuditTime     *int64                                                                   `json:"SubmitAuditTime,omitempty" xml:"SubmitAuditTime,omitempty"`
	UpdateTime          *int64                                                                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	MaterialName        *string                                                                  `json:"MaterialName,omitempty" xml:"MaterialName,omitempty"`
	Remark              *string                                                                  `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime          *int64                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId              *string                                                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizId               *string                                                                  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ServicePrice        *float32                                                                 `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	TmIcon              *string                                                                  `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	TmName              *string                                                                  `json:"TmName,omitempty" xml:"TmName,omitempty"`
	MaterialId          *int64                                                                   `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	SupplementId        *int64                                                                   `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	LoaUrl              *string                                                                  `json:"LoaUrl,omitempty" xml:"LoaUrl,omitempty"`
	TmNumber            *string                                                                  `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	Note                *string                                                                  `json:"Note,omitempty" xml:"Note,omitempty"`
	SupplementStatus    *int32                                                                   `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	PrincipalName       *int32                                                                   `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	TotalPrice          *float32                                                                 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	SubmitTime          *int64                                                                   `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	OrderId             *string                                                                  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ThirdClassification *QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassification `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty" type:"Struct"`
	Flags               *QueryTradeMarkApplicationsResponseBodyDataTmProducesFlags               `json:"Flags,omitempty" xml:"Flags,omitempty" type:"Struct"`
	FirstClassification *QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
	RenewResponse       *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse       `json:"RenewResponse,omitempty" xml:"RenewResponse,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProduces) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProduces) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetType(v int32) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.Type = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetStatus(v int32) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.Status = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetOrderPrice(v float32) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.OrderPrice = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetSubmitAuditTime(v int64) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.SubmitAuditTime = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetUpdateTime(v int64) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.UpdateTime = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetMaterialName(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.MaterialName = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetRemark(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.Remark = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetCreateTime(v int64) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.CreateTime = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetUserId(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.UserId = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetBizId(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.BizId = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetServicePrice(v float32) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.ServicePrice = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetTmIcon(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.TmIcon = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetTmName(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.TmName = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetMaterialId(v int64) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.MaterialId = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetSupplementId(v int64) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.SupplementId = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetLoaUrl(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.LoaUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetTmNumber(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.TmNumber = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetNote(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.Note = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetSupplementStatus(v int32) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.SupplementStatus = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetPrincipalName(v int32) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.PrincipalName = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetTotalPrice(v float32) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.TotalPrice = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetSubmitTime(v int64) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.SubmitTime = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetOrderId(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.OrderId = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetThirdClassification(v *QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassification) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.ThirdClassification = v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetFlags(v *QueryTradeMarkApplicationsResponseBodyDataTmProducesFlags) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.Flags = v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetFirstClassification(v *QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.FirstClassification = v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProduces) SetRenewResponse(v *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) *QueryTradeMarkApplicationsResponseBodyDataTmProduces {
	s.RenewResponse = v
	return s
}

type QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassification struct {
	ThirdClassifications []*QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications `json:"ThirdClassifications,omitempty" xml:"ThirdClassifications,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassification) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassification) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassification) SetThirdClassifications(v []*QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications) *QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassification {
	s.ThirdClassifications = v
	return s
}

type QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications struct {
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications) SetClassificationName(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications {
	s.ClassificationName = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications) SetClassificationCode(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesThirdClassificationThirdClassifications {
	s.ClassificationCode = &v
	return s
}

type QueryTradeMarkApplicationsResponseBodyDataTmProducesFlags struct {
	Flags []*string `json:"Flags,omitempty" xml:"Flags,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesFlags) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesFlags) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesFlags) SetFlags(v []*string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesFlags {
	s.Flags = v
	return s
}

type QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification struct {
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification) SetClassificationName(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification {
	s.ClassificationName = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification) SetClassificationCode(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesFirstClassification {
	s.ClassificationCode = &v
	return s
}

type QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse struct {
	EngName       *string `json:"EngName,omitempty" xml:"EngName,omitempty"`
	RegisterTime  *int64  `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	EngAddress    *string `json:"EngAddress,omitempty" xml:"EngAddress,omitempty"`
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SubmitSbjtime *int64  `json:"SubmitSbjtime,omitempty" xml:"SubmitSbjtime,omitempty"`
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) SetEngName(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse {
	s.EngName = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) SetRegisterTime(v int64) *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse {
	s.RegisterTime = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) SetEngAddress(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse {
	s.EngAddress = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) SetAddress(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse {
	s.Address = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) SetName(v string) *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse {
	s.Name = &v
	return s
}

func (s *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse) SetSubmitSbjtime(v int64) *QueryTradeMarkApplicationsResponseBodyDataTmProducesRenewResponse {
	s.SubmitSbjtime = &v
	return s
}

type QueryTradeMarkApplicationsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTradeMarkApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTradeMarkApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationsResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationsResponse) SetHeaders(v map[string]*string) *QueryTradeMarkApplicationsResponse {
	s.Headers = v
	return s
}

func (s *QueryTradeMarkApplicationsResponse) SetBody(v *QueryTradeMarkApplicationsResponseBody) *QueryTradeMarkApplicationsResponse {
	s.Body = v
	return s
}

type UpdateApplicantContacterRequest struct {
	ContactAddress *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactName    *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber  *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactEmail   *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ApplicantId    *int64  `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ContactZipCode *string `json:"ContactZipCode,omitempty" xml:"ContactZipCode,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s UpdateApplicantContacterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicantContacterRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicantContacterRequest) SetContactAddress(v string) *UpdateApplicantContacterRequest {
	s.ContactAddress = &v
	return s
}

func (s *UpdateApplicantContacterRequest) SetContactName(v string) *UpdateApplicantContacterRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateApplicantContacterRequest) SetContactNumber(v string) *UpdateApplicantContacterRequest {
	s.ContactNumber = &v
	return s
}

func (s *UpdateApplicantContacterRequest) SetContactEmail(v string) *UpdateApplicantContacterRequest {
	s.ContactEmail = &v
	return s
}

func (s *UpdateApplicantContacterRequest) SetApplicantId(v int64) *UpdateApplicantContacterRequest {
	s.ApplicantId = &v
	return s
}

func (s *UpdateApplicantContacterRequest) SetContactZipCode(v string) *UpdateApplicantContacterRequest {
	s.ContactZipCode = &v
	return s
}

func (s *UpdateApplicantContacterRequest) SetBizId(v string) *UpdateApplicantContacterRequest {
	s.BizId = &v
	return s
}

type UpdateApplicantContacterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicantContacterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicantContacterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicantContacterResponseBody) SetRequestId(v string) *UpdateApplicantContacterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApplicantContacterResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApplicantContacterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicantContacterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicantContacterResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicantContacterResponse) SetHeaders(v map[string]*string) *UpdateApplicantContacterResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicantContacterResponse) SetBody(v *UpdateApplicantContacterResponseBody) *UpdateApplicantContacterResponse {
	s.Body = v
	return s
}

type SaveTaskRequest struct {
	Request *string `json:"Request,omitempty" xml:"Request,omitempty"`
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s SaveTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskRequest) SetRequest(v string) *SaveTaskRequest {
	s.Request = &v
	return s
}

func (s *SaveTaskRequest) SetBizType(v string) *SaveTaskRequest {
	s.BizType = &v
	return s
}

type SaveTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskResponseBody) SetRequestId(v string) *SaveTaskResponseBody {
	s.RequestId = &v
	return s
}

type SaveTaskResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskResponse) SetHeaders(v map[string]*string) *SaveTaskResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskResponse) SetBody(v *SaveTaskResponseBody) *SaveTaskResponse {
	s.Body = v
	return s
}

type SubmitTrademarkApplicationComplaintRequest struct {
	BizId   *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Files   map[string]interface{} `json:"Files,omitempty" xml:"Files,omitempty"`
	Content *string                `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s SubmitTrademarkApplicationComplaintRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTrademarkApplicationComplaintRequest) GoString() string {
	return s.String()
}

func (s *SubmitTrademarkApplicationComplaintRequest) SetBizId(v string) *SubmitTrademarkApplicationComplaintRequest {
	s.BizId = &v
	return s
}

func (s *SubmitTrademarkApplicationComplaintRequest) SetFiles(v map[string]interface{}) *SubmitTrademarkApplicationComplaintRequest {
	s.Files = v
	return s
}

func (s *SubmitTrademarkApplicationComplaintRequest) SetContent(v string) *SubmitTrademarkApplicationComplaintRequest {
	s.Content = &v
	return s
}

type SubmitTrademarkApplicationComplaintShrinkRequest struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s SubmitTrademarkApplicationComplaintShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTrademarkApplicationComplaintShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTrademarkApplicationComplaintShrinkRequest) SetBizId(v string) *SubmitTrademarkApplicationComplaintShrinkRequest {
	s.BizId = &v
	return s
}

func (s *SubmitTrademarkApplicationComplaintShrinkRequest) SetFilesShrink(v string) *SubmitTrademarkApplicationComplaintShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *SubmitTrademarkApplicationComplaintShrinkRequest) SetContent(v string) *SubmitTrademarkApplicationComplaintShrinkRequest {
	s.Content = &v
	return s
}

type SubmitTrademarkApplicationComplaintResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitTrademarkApplicationComplaintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTrademarkApplicationComplaintResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTrademarkApplicationComplaintResponseBody) SetRequestId(v string) *SubmitTrademarkApplicationComplaintResponseBody {
	s.RequestId = &v
	return s
}

type SubmitTrademarkApplicationComplaintResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitTrademarkApplicationComplaintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTrademarkApplicationComplaintResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTrademarkApplicationComplaintResponse) GoString() string {
	return s.String()
}

func (s *SubmitTrademarkApplicationComplaintResponse) SetHeaders(v map[string]*string) *SubmitTrademarkApplicationComplaintResponse {
	s.Headers = v
	return s
}

func (s *SubmitTrademarkApplicationComplaintResponse) SetBody(v *SubmitTrademarkApplicationComplaintResponseBody) *SubmitTrademarkApplicationComplaintResponse {
	s.Body = v
	return s
}

type WriteIntentionCommunicationLogRequest struct {
	BizId  *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note   *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Reject *bool   `json:"Reject,omitempty" xml:"Reject,omitempty"`
}

func (s WriteIntentionCommunicationLogRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteIntentionCommunicationLogRequest) GoString() string {
	return s.String()
}

func (s *WriteIntentionCommunicationLogRequest) SetBizId(v string) *WriteIntentionCommunicationLogRequest {
	s.BizId = &v
	return s
}

func (s *WriteIntentionCommunicationLogRequest) SetNote(v string) *WriteIntentionCommunicationLogRequest {
	s.Note = &v
	return s
}

func (s *WriteIntentionCommunicationLogRequest) SetReject(v bool) *WriteIntentionCommunicationLogRequest {
	s.Reject = &v
	return s
}

type WriteIntentionCommunicationLogResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s WriteIntentionCommunicationLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteIntentionCommunicationLogResponseBody) GoString() string {
	return s.String()
}

func (s *WriteIntentionCommunicationLogResponseBody) SetErrorMsg(v string) *WriteIntentionCommunicationLogResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *WriteIntentionCommunicationLogResponseBody) SetRequestId(v string) *WriteIntentionCommunicationLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *WriteIntentionCommunicationLogResponseBody) SetSuccess(v bool) *WriteIntentionCommunicationLogResponseBody {
	s.Success = &v
	return s
}

func (s *WriteIntentionCommunicationLogResponseBody) SetErrorCode(v string) *WriteIntentionCommunicationLogResponseBody {
	s.ErrorCode = &v
	return s
}

type WriteIntentionCommunicationLogResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WriteIntentionCommunicationLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WriteIntentionCommunicationLogResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteIntentionCommunicationLogResponse) GoString() string {
	return s.String()
}

func (s *WriteIntentionCommunicationLogResponse) SetHeaders(v map[string]*string) *WriteIntentionCommunicationLogResponse {
	s.Headers = v
	return s
}

func (s *WriteIntentionCommunicationLogResponse) SetBody(v *WriteIntentionCommunicationLogResponseBody) *WriteIntentionCommunicationLogResponse {
	s.Body = v
	return s
}

type SaveTaskForOfficialFileCustomRequest struct {
	EndAcceptTime   *int64 `json:"EndAcceptTime,omitempty" xml:"EndAcceptTime,omitempty"`
	StartAcceptTime *int64 `json:"StartAcceptTime,omitempty" xml:"StartAcceptTime,omitempty"`
}

func (s SaveTaskForOfficialFileCustomRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForOfficialFileCustomRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForOfficialFileCustomRequest) SetEndAcceptTime(v int64) *SaveTaskForOfficialFileCustomRequest {
	s.EndAcceptTime = &v
	return s
}

func (s *SaveTaskForOfficialFileCustomRequest) SetStartAcceptTime(v int64) *SaveTaskForOfficialFileCustomRequest {
	s.StartAcceptTime = &v
	return s
}

type SaveTaskForOfficialFileCustomResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveTaskForOfficialFileCustomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForOfficialFileCustomResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForOfficialFileCustomResponseBody) SetSuccess(v bool) *SaveTaskForOfficialFileCustomResponseBody {
	s.Success = &v
	return s
}

func (s *SaveTaskForOfficialFileCustomResponseBody) SetRequestId(v string) *SaveTaskForOfficialFileCustomResponseBody {
	s.RequestId = &v
	return s
}

type SaveTaskForOfficialFileCustomResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveTaskForOfficialFileCustomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTaskForOfficialFileCustomResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTaskForOfficialFileCustomResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForOfficialFileCustomResponse) SetHeaders(v map[string]*string) *SaveTaskForOfficialFileCustomResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForOfficialFileCustomResponse) SetBody(v *SaveTaskForOfficialFileCustomResponseBody) *SaveTaskForOfficialFileCustomResponse {
	s.Body = v
	return s
}

type DescirbeCombineTrademarkRequest struct {
	RegistrationNumber *string `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName          *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	Products           *string `json:"Products,omitempty" xml:"Products,omitempty"`
	AccurateMatch      *bool   `json:"AccurateMatch,omitempty" xml:"AccurateMatch,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Classification     *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	SimilarGroups      *string `json:"SimilarGroups,omitempty" xml:"SimilarGroups,omitempty"`
}

func (s DescirbeCombineTrademarkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescirbeCombineTrademarkRequest) GoString() string {
	return s.String()
}

func (s *DescirbeCombineTrademarkRequest) SetRegistrationNumber(v string) *DescirbeCombineTrademarkRequest {
	s.RegistrationNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkRequest) SetName(v string) *DescirbeCombineTrademarkRequest {
	s.Name = &v
	return s
}

func (s *DescirbeCombineTrademarkRequest) SetOwnerName(v string) *DescirbeCombineTrademarkRequest {
	s.OwnerName = &v
	return s
}

func (s *DescirbeCombineTrademarkRequest) SetProducts(v string) *DescirbeCombineTrademarkRequest {
	s.Products = &v
	return s
}

func (s *DescirbeCombineTrademarkRequest) SetAccurateMatch(v bool) *DescirbeCombineTrademarkRequest {
	s.AccurateMatch = &v
	return s
}

func (s *DescirbeCombineTrademarkRequest) SetPageNumber(v int32) *DescirbeCombineTrademarkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkRequest) SetPageSize(v int32) *DescirbeCombineTrademarkRequest {
	s.PageSize = &v
	return s
}

func (s *DescirbeCombineTrademarkRequest) SetClassification(v string) *DescirbeCombineTrademarkRequest {
	s.Classification = &v
	return s
}

func (s *DescirbeCombineTrademarkRequest) SetSimilarGroups(v string) *DescirbeCombineTrademarkRequest {
	s.SimilarGroups = &v
	return s
}

type DescirbeCombineTrademarkResponseBody struct {
	NextPage          *bool                                       `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	RequestId         *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalPageNumber   *int32                                      `json:"TotalPageNumber,omitempty" xml:"TotalPageNumber,omitempty"`
	PrePage           *bool                                       `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	CurrentPageNumber *int32                                      `json:"CurrentPageNumber,omitempty" xml:"CurrentPageNumber,omitempty"`
	TotalItemNumber   *int32                                      `json:"TotalItemNumber,omitempty" xml:"TotalItemNumber,omitempty"`
	PageSize          *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data              []*DescirbeCombineTrademarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescirbeCombineTrademarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescirbeCombineTrademarkResponseBody) GoString() string {
	return s.String()
}

func (s *DescirbeCombineTrademarkResponseBody) SetNextPage(v bool) *DescirbeCombineTrademarkResponseBody {
	s.NextPage = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBody) SetRequestId(v string) *DescirbeCombineTrademarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBody) SetTotalPageNumber(v int32) *DescirbeCombineTrademarkResponseBody {
	s.TotalPageNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBody) SetPrePage(v bool) *DescirbeCombineTrademarkResponseBody {
	s.PrePage = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBody) SetCurrentPageNumber(v int32) *DescirbeCombineTrademarkResponseBody {
	s.CurrentPageNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBody) SetTotalItemNumber(v int32) *DescirbeCombineTrademarkResponseBody {
	s.TotalItemNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBody) SetPageSize(v int32) *DescirbeCombineTrademarkResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBody) SetData(v []*DescirbeCombineTrademarkResponseBodyData) *DescirbeCombineTrademarkResponseBody {
	s.Data = v
	return s
}

type DescirbeCombineTrademarkResponseBodyData struct {
	Status                    *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	OwnerAddress              *string                                                     `json:"OwnerAddress,omitempty" xml:"OwnerAddress,omitempty"`
	PreAnnDate                *string                                                     `json:"PreAnnDate,omitempty" xml:"PreAnnDate,omitempty"`
	PreAnnNumber              *string                                                     `json:"PreAnnNumber,omitempty" xml:"PreAnnNumber,omitempty"`
	IntlRegDate               *string                                                     `json:"IntlRegDate,omitempty" xml:"IntlRegDate,omitempty"`
	Share                     *string                                                     `json:"Share,omitempty" xml:"Share,omitempty"`
	OwnerEnName               *string                                                     `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	SubsequentDesignationDate *string                                                     `json:"SubsequentDesignationDate,omitempty" xml:"SubsequentDesignationDate,omitempty"`
	IndexId                   *string                                                     `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	RegAnnNumber              *string                                                     `json:"RegAnnNumber,omitempty" xml:"RegAnnNumber,omitempty"`
	RegistrationNumber        *string                                                     `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	SecondAnnoType            *string                                                     `json:"SecondAnnoType,omitempty" xml:"SecondAnnoType,omitempty"`
	Agency                    *string                                                     `json:"Agency,omitempty" xml:"Agency,omitempty"`
	OwnerEnAddress            *string                                                     `json:"OwnerEnAddress,omitempty" xml:"OwnerEnAddress,omitempty"`
	Classification            *string                                                     `json:"Classification,omitempty" xml:"Classification,omitempty"`
	Name                      *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	ApplyDate                 *string                                                     `json:"ApplyDate,omitempty" xml:"ApplyDate,omitempty"`
	PriorityDate              *string                                                     `json:"PriorityDate,omitempty" xml:"PriorityDate,omitempty"`
	ProductDescription        *string                                                     `json:"ProductDescription,omitempty" xml:"ProductDescription,omitempty"`
	Image                     *string                                                     `json:"Image,omitempty" xml:"Image,omitempty"`
	SecondAnnoNumber          *string                                                     `json:"SecondAnnoNumber,omitempty" xml:"SecondAnnoNumber,omitempty"`
	RegistrationType          *string                                                     `json:"RegistrationType,omitempty" xml:"RegistrationType,omitempty"`
	FirstAnnoNumber           *string                                                     `json:"FirstAnnoNumber,omitempty" xml:"FirstAnnoNumber,omitempty"`
	OwnerName                 *string                                                     `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	RegAnnDate                *string                                                     `json:"RegAnnDate,omitempty" xml:"RegAnnDate,omitempty"`
	SimilarGroup              *string                                                     `json:"SimilarGroup,omitempty" xml:"SimilarGroup,omitempty"`
	OnSale                    *int32                                                      `json:"OnSale,omitempty" xml:"OnSale,omitempty"`
	ExclusiveDateLimit        *string                                                     `json:"ExclusiveDateLimit,omitempty" xml:"ExclusiveDateLimit,omitempty"`
	FirstAnnoType             *string                                                     `json:"FirstAnnoType,omitempty" xml:"FirstAnnoType,omitempty"`
	LastProcedureStatus       *string                                                     `json:"LastProcedureStatus,omitempty" xml:"LastProcedureStatus,omitempty"`
	LawFinalStatus            *string                                                     `json:"LawFinalStatus,omitempty" xml:"LawFinalStatus,omitempty"`
	AnnouncementList          []*DescirbeCombineTrademarkResponseBodyDataAnnouncementList `json:"AnnouncementList,omitempty" xml:"AnnouncementList,omitempty" type:"Repeated"`
	Procedures                []*DescirbeCombineTrademarkResponseBodyDataProcedures       `json:"Procedures,omitempty" xml:"Procedures,omitempty" type:"Repeated"`
}

func (s DescirbeCombineTrademarkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescirbeCombineTrademarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetStatus(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetOwnerAddress(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.OwnerAddress = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetPreAnnDate(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.PreAnnDate = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetPreAnnNumber(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.PreAnnNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetIntlRegDate(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.IntlRegDate = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetShare(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.Share = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetOwnerEnName(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.OwnerEnName = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetSubsequentDesignationDate(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.SubsequentDesignationDate = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetIndexId(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.IndexId = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetRegAnnNumber(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.RegAnnNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetRegistrationNumber(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.RegistrationNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetSecondAnnoType(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.SecondAnnoType = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetAgency(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.Agency = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetOwnerEnAddress(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.OwnerEnAddress = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetClassification(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.Classification = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetName(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetApplyDate(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.ApplyDate = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetPriorityDate(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.PriorityDate = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetProductDescription(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.ProductDescription = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetImage(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.Image = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetSecondAnnoNumber(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.SecondAnnoNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetRegistrationType(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.RegistrationType = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetFirstAnnoNumber(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.FirstAnnoNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetOwnerName(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.OwnerName = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetRegAnnDate(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.RegAnnDate = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetSimilarGroup(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.SimilarGroup = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetOnSale(v int32) *DescirbeCombineTrademarkResponseBodyData {
	s.OnSale = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetExclusiveDateLimit(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.ExclusiveDateLimit = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetFirstAnnoType(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.FirstAnnoType = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetLastProcedureStatus(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.LastProcedureStatus = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetLawFinalStatus(v string) *DescirbeCombineTrademarkResponseBodyData {
	s.LawFinalStatus = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetAnnouncementList(v []*DescirbeCombineTrademarkResponseBodyDataAnnouncementList) *DescirbeCombineTrademarkResponseBodyData {
	s.AnnouncementList = v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyData) SetProcedures(v []*DescirbeCombineTrademarkResponseBodyDataProcedures) *DescirbeCombineTrademarkResponseBodyData {
	s.Procedures = v
	return s
}

type DescirbeCombineTrademarkResponseBodyDataAnnouncementList struct {
	ImageUrl         *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	AnnDate          *string `json:"AnnDate,omitempty" xml:"AnnDate,omitempty"`
	OriginalImageUrl *string `json:"OriginalImageUrl,omitempty" xml:"OriginalImageUrl,omitempty"`
	AnnTypeName      *string `json:"AnnTypeName,omitempty" xml:"AnnTypeName,omitempty"`
	AnnNumber        *string `json:"AnnNumber,omitempty" xml:"AnnNumber,omitempty"`
	AnnTypeCode      *string `json:"AnnTypeCode,omitempty" xml:"AnnTypeCode,omitempty"`
}

func (s DescirbeCombineTrademarkResponseBodyDataAnnouncementList) String() string {
	return tea.Prettify(s)
}

func (s DescirbeCombineTrademarkResponseBodyDataAnnouncementList) GoString() string {
	return s.String()
}

func (s *DescirbeCombineTrademarkResponseBodyDataAnnouncementList) SetImageUrl(v string) *DescirbeCombineTrademarkResponseBodyDataAnnouncementList {
	s.ImageUrl = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataAnnouncementList) SetAnnDate(v string) *DescirbeCombineTrademarkResponseBodyDataAnnouncementList {
	s.AnnDate = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataAnnouncementList) SetOriginalImageUrl(v string) *DescirbeCombineTrademarkResponseBodyDataAnnouncementList {
	s.OriginalImageUrl = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataAnnouncementList) SetAnnTypeName(v string) *DescirbeCombineTrademarkResponseBodyDataAnnouncementList {
	s.AnnTypeName = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataAnnouncementList) SetAnnNumber(v string) *DescirbeCombineTrademarkResponseBodyDataAnnouncementList {
	s.AnnNumber = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataAnnouncementList) SetAnnTypeCode(v string) *DescirbeCombineTrademarkResponseBodyDataAnnouncementList {
	s.AnnTypeCode = &v
	return s
}

type DescirbeCombineTrademarkResponseBodyDataProcedures struct {
	ProcedureStep   *string `json:"ProcedureStep,omitempty" xml:"ProcedureStep,omitempty"`
	ProcedureResult *string `json:"ProcedureResult,omitempty" xml:"ProcedureResult,omitempty"`
	ProcedureCode   *string `json:"ProcedureCode,omitempty" xml:"ProcedureCode,omitempty"`
	ProcedureDate   *string `json:"ProcedureDate,omitempty" xml:"ProcedureDate,omitempty"`
	ProcedureName   *string `json:"ProcedureName,omitempty" xml:"ProcedureName,omitempty"`
}

func (s DescirbeCombineTrademarkResponseBodyDataProcedures) String() string {
	return tea.Prettify(s)
}

func (s DescirbeCombineTrademarkResponseBodyDataProcedures) GoString() string {
	return s.String()
}

func (s *DescirbeCombineTrademarkResponseBodyDataProcedures) SetProcedureStep(v string) *DescirbeCombineTrademarkResponseBodyDataProcedures {
	s.ProcedureStep = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataProcedures) SetProcedureResult(v string) *DescirbeCombineTrademarkResponseBodyDataProcedures {
	s.ProcedureResult = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataProcedures) SetProcedureCode(v string) *DescirbeCombineTrademarkResponseBodyDataProcedures {
	s.ProcedureCode = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataProcedures) SetProcedureDate(v string) *DescirbeCombineTrademarkResponseBodyDataProcedures {
	s.ProcedureDate = &v
	return s
}

func (s *DescirbeCombineTrademarkResponseBodyDataProcedures) SetProcedureName(v string) *DescirbeCombineTrademarkResponseBodyDataProcedures {
	s.ProcedureName = &v
	return s
}

type DescirbeCombineTrademarkResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescirbeCombineTrademarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescirbeCombineTrademarkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescirbeCombineTrademarkResponse) GoString() string {
	return s.String()
}

func (s *DescirbeCombineTrademarkResponse) SetHeaders(v map[string]*string) *DescirbeCombineTrademarkResponse {
	s.Headers = v
	return s
}

func (s *DescirbeCombineTrademarkResponse) SetBody(v *DescirbeCombineTrademarkResponseBody) *DescirbeCombineTrademarkResponse {
	s.Body = v
	return s
}

type GetNotaryOrderRequest struct {
	NotaryOrderId *int64 `json:"NotaryOrderId,omitempty" xml:"NotaryOrderId,omitempty"`
}

func (s GetNotaryOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNotaryOrderRequest) GoString() string {
	return s.String()
}

func (s *GetNotaryOrderRequest) SetNotaryOrderId(v int64) *GetNotaryOrderRequest {
	s.NotaryOrderId = &v
	return s
}

type GetNotaryOrderResponseBody struct {
	Type                        *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	OrderPrice                  *float32 `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	LegalPersonIdCard           *string  `json:"LegalPersonIdCard,omitempty" xml:"LegalPersonIdCard,omitempty"`
	BusinessLicenseId           *string  `json:"BusinessLicenseId,omitempty" xml:"BusinessLicenseId,omitempty"`
	NotaryPostReceipt           *string  `json:"NotaryPostReceipt,omitempty" xml:"NotaryPostReceipt,omitempty"`
	CompanyContactName          *string  `json:"CompanyContactName,omitempty" xml:"CompanyContactName,omitempty"`
	NotaryStatus                *int32   `json:"NotaryStatus,omitempty" xml:"NotaryStatus,omitempty"`
	SellerBackOfIdCard          *string  `json:"SellerBackOfIdCard,omitempty" xml:"SellerBackOfIdCard,omitempty"`
	TmRegisterChangeCertificate *string  `json:"TmRegisterChangeCertificate,omitempty" xml:"TmRegisterChangeCertificate,omitempty"`
	RequestId                   *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LegalPersonName             *string  `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	TmImage                     *string  `json:"TmImage,omitempty" xml:"TmImage,omitempty"`
	NotaryAcceptDate            *int64   `json:"NotaryAcceptDate,omitempty" xml:"NotaryAcceptDate,omitempty"`
	ErrorCode                   *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	AliyunOrderId               *string  `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	NotarySucceedDate           *int64   `json:"NotarySucceedDate,omitempty" xml:"NotarySucceedDate,omitempty"`
	ApplyPostStatus             *int32   `json:"ApplyPostStatus,omitempty" xml:"ApplyPostStatus,omitempty"`
	ErrorMsg                    *string  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Name                        *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	BusinessLicense             *string  `json:"BusinessLicense,omitempty" xml:"BusinessLicense,omitempty"`
	ReceiverName                *string  `json:"ReceiverName,omitempty" xml:"ReceiverName,omitempty"`
	OrderDate                   *int64   `json:"OrderDate,omitempty" xml:"OrderDate,omitempty"`
	CompanyContactPhone         *string  `json:"CompanyContactPhone,omitempty" xml:"CompanyContactPhone,omitempty"`
	NotaryType                  *int32   `json:"NotaryType,omitempty" xml:"NotaryType,omitempty"`
	NotaryFailedDate            *int64   `json:"NotaryFailedDate,omitempty" xml:"NotaryFailedDate,omitempty"`
	TmClassification            *string  `json:"TmClassification,omitempty" xml:"TmClassification,omitempty"`
	Success                     *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	BizId                       *string  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	NotaryOrderId               *int64   `json:"NotaryOrderId,omitempty" xml:"NotaryOrderId,omitempty"`
	Phone                       *string  `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ReceiverPhone               *string  `json:"ReceiverPhone,omitempty" xml:"ReceiverPhone,omitempty"`
	TmRegisterCertificate       *string  `json:"TmRegisterCertificate,omitempty" xml:"TmRegisterCertificate,omitempty"`
	TmName                      *string  `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmRegisterNo                *string  `json:"TmRegisterNo,omitempty" xml:"TmRegisterNo,omitempty"`
	SellerCompanyName           *string  `json:"SellerCompanyName,omitempty" xml:"SellerCompanyName,omitempty"`
	TmAcceptCertificate         *string  `json:"TmAcceptCertificate,omitempty" xml:"TmAcceptCertificate,omitempty"`
	ReceiverPostalCode          *string  `json:"ReceiverPostalCode,omitempty" xml:"ReceiverPostalCode,omitempty"`
	NotaryCertificate           *string  `json:"NotaryCertificate,omitempty" xml:"NotaryCertificate,omitempty"`
	LegalPersonPhone            *string  `json:"LegalPersonPhone,omitempty" xml:"LegalPersonPhone,omitempty"`
	NotaryFailedReason          *string  `json:"NotaryFailedReason,omitempty" xml:"NotaryFailedReason,omitempty"`
	SellerFrontOfIdCard         *string  `json:"SellerFrontOfIdCard,omitempty" xml:"SellerFrontOfIdCard,omitempty"`
	ReceiverAddress             *string  `json:"ReceiverAddress,omitempty" xml:"ReceiverAddress,omitempty"`
	NotaryPlatformName          *string  `json:"NotaryPlatformName,omitempty" xml:"NotaryPlatformName,omitempty"`
}

func (s GetNotaryOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNotaryOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotaryOrderResponseBody) SetType(v string) *GetNotaryOrderResponseBody {
	s.Type = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetOrderPrice(v float32) *GetNotaryOrderResponseBody {
	s.OrderPrice = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetLegalPersonIdCard(v string) *GetNotaryOrderResponseBody {
	s.LegalPersonIdCard = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetBusinessLicenseId(v string) *GetNotaryOrderResponseBody {
	s.BusinessLicenseId = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryPostReceipt(v string) *GetNotaryOrderResponseBody {
	s.NotaryPostReceipt = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetCompanyContactName(v string) *GetNotaryOrderResponseBody {
	s.CompanyContactName = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryStatus(v int32) *GetNotaryOrderResponseBody {
	s.NotaryStatus = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetSellerBackOfIdCard(v string) *GetNotaryOrderResponseBody {
	s.SellerBackOfIdCard = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetTmRegisterChangeCertificate(v string) *GetNotaryOrderResponseBody {
	s.TmRegisterChangeCertificate = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetRequestId(v string) *GetNotaryOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetLegalPersonName(v string) *GetNotaryOrderResponseBody {
	s.LegalPersonName = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetTmImage(v string) *GetNotaryOrderResponseBody {
	s.TmImage = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryAcceptDate(v int64) *GetNotaryOrderResponseBody {
	s.NotaryAcceptDate = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetErrorCode(v string) *GetNotaryOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetAliyunOrderId(v string) *GetNotaryOrderResponseBody {
	s.AliyunOrderId = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotarySucceedDate(v int64) *GetNotaryOrderResponseBody {
	s.NotarySucceedDate = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetApplyPostStatus(v int32) *GetNotaryOrderResponseBody {
	s.ApplyPostStatus = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetErrorMsg(v string) *GetNotaryOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetName(v string) *GetNotaryOrderResponseBody {
	s.Name = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetBusinessLicense(v string) *GetNotaryOrderResponseBody {
	s.BusinessLicense = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetReceiverName(v string) *GetNotaryOrderResponseBody {
	s.ReceiverName = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetOrderDate(v int64) *GetNotaryOrderResponseBody {
	s.OrderDate = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetCompanyContactPhone(v string) *GetNotaryOrderResponseBody {
	s.CompanyContactPhone = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryType(v int32) *GetNotaryOrderResponseBody {
	s.NotaryType = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryFailedDate(v int64) *GetNotaryOrderResponseBody {
	s.NotaryFailedDate = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetTmClassification(v string) *GetNotaryOrderResponseBody {
	s.TmClassification = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetSuccess(v bool) *GetNotaryOrderResponseBody {
	s.Success = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetBizId(v string) *GetNotaryOrderResponseBody {
	s.BizId = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryOrderId(v int64) *GetNotaryOrderResponseBody {
	s.NotaryOrderId = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetPhone(v string) *GetNotaryOrderResponseBody {
	s.Phone = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetReceiverPhone(v string) *GetNotaryOrderResponseBody {
	s.ReceiverPhone = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetTmRegisterCertificate(v string) *GetNotaryOrderResponseBody {
	s.TmRegisterCertificate = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetTmName(v string) *GetNotaryOrderResponseBody {
	s.TmName = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetTmRegisterNo(v string) *GetNotaryOrderResponseBody {
	s.TmRegisterNo = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetSellerCompanyName(v string) *GetNotaryOrderResponseBody {
	s.SellerCompanyName = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetTmAcceptCertificate(v string) *GetNotaryOrderResponseBody {
	s.TmAcceptCertificate = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetReceiverPostalCode(v string) *GetNotaryOrderResponseBody {
	s.ReceiverPostalCode = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryCertificate(v string) *GetNotaryOrderResponseBody {
	s.NotaryCertificate = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetLegalPersonPhone(v string) *GetNotaryOrderResponseBody {
	s.LegalPersonPhone = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryFailedReason(v string) *GetNotaryOrderResponseBody {
	s.NotaryFailedReason = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetSellerFrontOfIdCard(v string) *GetNotaryOrderResponseBody {
	s.SellerFrontOfIdCard = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetReceiverAddress(v string) *GetNotaryOrderResponseBody {
	s.ReceiverAddress = &v
	return s
}

func (s *GetNotaryOrderResponseBody) SetNotaryPlatformName(v string) *GetNotaryOrderResponseBody {
	s.NotaryPlatformName = &v
	return s
}

type GetNotaryOrderResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNotaryOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNotaryOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNotaryOrderResponse) GoString() string {
	return s.String()
}

func (s *GetNotaryOrderResponse) SetHeaders(v map[string]*string) *GetNotaryOrderResponse {
	s.Headers = v
	return s
}

func (s *GetNotaryOrderResponse) SetBody(v *GetNotaryOrderResponseBody) *GetNotaryOrderResponse {
	s.Body = v
	return s
}

type ConfirmAdditionalMaterialRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note  *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s ConfirmAdditionalMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmAdditionalMaterialRequest) GoString() string {
	return s.String()
}

func (s *ConfirmAdditionalMaterialRequest) SetBizId(v string) *ConfirmAdditionalMaterialRequest {
	s.BizId = &v
	return s
}

func (s *ConfirmAdditionalMaterialRequest) SetNote(v string) *ConfirmAdditionalMaterialRequest {
	s.Note = &v
	return s
}

type ConfirmAdditionalMaterialResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s ConfirmAdditionalMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmAdditionalMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmAdditionalMaterialResponseBody) SetErrorMsg(v string) *ConfirmAdditionalMaterialResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ConfirmAdditionalMaterialResponseBody) SetRequestId(v string) *ConfirmAdditionalMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmAdditionalMaterialResponseBody) SetSuccess(v bool) *ConfirmAdditionalMaterialResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmAdditionalMaterialResponseBody) SetErrorCode(v string) *ConfirmAdditionalMaterialResponseBody {
	s.ErrorCode = &v
	return s
}

type ConfirmAdditionalMaterialResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfirmAdditionalMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmAdditionalMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmAdditionalMaterialResponse) GoString() string {
	return s.String()
}

func (s *ConfirmAdditionalMaterialResponse) SetHeaders(v map[string]*string) *ConfirmAdditionalMaterialResponse {
	s.Headers = v
	return s
}

func (s *ConfirmAdditionalMaterialResponse) SetBody(v *ConfirmAdditionalMaterialResponseBody) *ConfirmAdditionalMaterialResponse {
	s.Body = v
	return s
}

type InsertRenewInfoRequest struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	EngName      *string `json:"EngName,omitempty" xml:"EngName,omitempty"`
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EngAddress   *string `json:"EngAddress,omitempty" xml:"EngAddress,omitempty"`
	RegisterTime *int64  `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
}

func (s InsertRenewInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertRenewInfoRequest) GoString() string {
	return s.String()
}

func (s *InsertRenewInfoRequest) SetName(v string) *InsertRenewInfoRequest {
	s.Name = &v
	return s
}

func (s *InsertRenewInfoRequest) SetEngName(v string) *InsertRenewInfoRequest {
	s.EngName = &v
	return s
}

func (s *InsertRenewInfoRequest) SetAddress(v string) *InsertRenewInfoRequest {
	s.Address = &v
	return s
}

func (s *InsertRenewInfoRequest) SetEngAddress(v string) *InsertRenewInfoRequest {
	s.EngAddress = &v
	return s
}

func (s *InsertRenewInfoRequest) SetRegisterTime(v int64) *InsertRenewInfoRequest {
	s.RegisterTime = &v
	return s
}

type InsertRenewInfoResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s InsertRenewInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertRenewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *InsertRenewInfoResponseBody) SetErrorMsg(v string) *InsertRenewInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InsertRenewInfoResponseBody) SetRequestId(v string) *InsertRenewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertRenewInfoResponseBody) SetSuccess(v bool) *InsertRenewInfoResponseBody {
	s.Success = &v
	return s
}

func (s *InsertRenewInfoResponseBody) SetErrorCode(v string) *InsertRenewInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InsertRenewInfoResponseBody) SetId(v int64) *InsertRenewInfoResponseBody {
	s.Id = &v
	return s
}

type InsertRenewInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertRenewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertRenewInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertRenewInfoResponse) GoString() string {
	return s.String()
}

func (s *InsertRenewInfoResponse) SetHeaders(v map[string]*string) *InsertRenewInfoResponse {
	s.Headers = v
	return s
}

func (s *InsertRenewInfoResponse) SetBody(v *InsertRenewInfoResponseBody) *InsertRenewInfoResponse {
	s.Body = v
	return s
}

type QueryCredentialsInfoRequest struct {
	OssKey       *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	CompanyName  *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
}

func (s QueryCredentialsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCredentialsInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCredentialsInfoRequest) SetOssKey(v string) *QueryCredentialsInfoRequest {
	s.OssKey = &v
	return s
}

func (s *QueryCredentialsInfoRequest) SetMaterialType(v string) *QueryCredentialsInfoRequest {
	s.MaterialType = &v
	return s
}

func (s *QueryCredentialsInfoRequest) SetCompanyName(v string) *QueryCredentialsInfoRequest {
	s.CompanyName = &v
	return s
}

type QueryCredentialsInfoResponseBody struct {
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CredentialsInfo *QueryCredentialsInfoResponseBodyCredentialsInfo `json:"CredentialsInfo,omitempty" xml:"CredentialsInfo,omitempty" type:"Struct"`
}

func (s QueryCredentialsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCredentialsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCredentialsInfoResponseBody) SetRequestId(v string) *QueryCredentialsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCredentialsInfoResponseBody) SetCredentialsInfo(v *QueryCredentialsInfoResponseBodyCredentialsInfo) *QueryCredentialsInfoResponseBody {
	s.CredentialsInfo = v
	return s
}

type QueryCredentialsInfoResponseBodyCredentialsInfo struct {
	CardNumber  *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	PersonName  *string `json:"PersonName,omitempty" xml:"PersonName,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
}

func (s QueryCredentialsInfoResponseBodyCredentialsInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryCredentialsInfoResponseBodyCredentialsInfo) GoString() string {
	return s.String()
}

func (s *QueryCredentialsInfoResponseBodyCredentialsInfo) SetCardNumber(v string) *QueryCredentialsInfoResponseBodyCredentialsInfo {
	s.CardNumber = &v
	return s
}

func (s *QueryCredentialsInfoResponseBodyCredentialsInfo) SetAddress(v string) *QueryCredentialsInfoResponseBodyCredentialsInfo {
	s.Address = &v
	return s
}

func (s *QueryCredentialsInfoResponseBodyCredentialsInfo) SetPersonName(v string) *QueryCredentialsInfoResponseBodyCredentialsInfo {
	s.PersonName = &v
	return s
}

func (s *QueryCredentialsInfoResponseBodyCredentialsInfo) SetProvince(v string) *QueryCredentialsInfoResponseBodyCredentialsInfo {
	s.Province = &v
	return s
}

func (s *QueryCredentialsInfoResponseBodyCredentialsInfo) SetCompanyName(v string) *QueryCredentialsInfoResponseBodyCredentialsInfo {
	s.CompanyName = &v
	return s
}

type QueryCredentialsInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCredentialsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCredentialsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCredentialsInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCredentialsInfoResponse) SetHeaders(v map[string]*string) *QueryCredentialsInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCredentialsInfoResponse) SetBody(v *QueryCredentialsInfoResponseBody) *QueryCredentialsInfoResponse {
	s.Body = v
	return s
}

type SearchTmOnsalesRequest struct {
	Keyword         *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Classification  *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PageNum         *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegisterNumber  *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	TmName          *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TopSearch       *string `json:"TopSearch,omitempty" xml:"TopSearch,omitempty"`
	Tag             *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	OrderPriceLeft  *int64  `json:"OrderPriceLeft,omitempty" xml:"OrderPriceLeft,omitempty"`
	OrderPriceRight *int64  `json:"OrderPriceRight,omitempty" xml:"OrderPriceRight,omitempty"`
	RegLeft         *int32  `json:"RegLeft,omitempty" xml:"RegLeft,omitempty"`
	RegRight        *int32  `json:"RegRight,omitempty" xml:"RegRight,omitempty"`
	SortName        *string `json:"SortName,omitempty" xml:"SortName,omitempty"`
	SortOrder       *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	QueryAll        *bool   `json:"QueryAll,omitempty" xml:"QueryAll,omitempty"`
}

func (s SearchTmOnsalesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTmOnsalesRequest) GoString() string {
	return s.String()
}

func (s *SearchTmOnsalesRequest) SetKeyword(v string) *SearchTmOnsalesRequest {
	s.Keyword = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetClassification(v string) *SearchTmOnsalesRequest {
	s.Classification = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetProductCode(v string) *SearchTmOnsalesRequest {
	s.ProductCode = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetPageNum(v int32) *SearchTmOnsalesRequest {
	s.PageNum = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetPageSize(v int32) *SearchTmOnsalesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetRegisterNumber(v string) *SearchTmOnsalesRequest {
	s.RegisterNumber = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetTmName(v string) *SearchTmOnsalesRequest {
	s.TmName = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetTopSearch(v string) *SearchTmOnsalesRequest {
	s.TopSearch = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetTag(v string) *SearchTmOnsalesRequest {
	s.Tag = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetOrderPriceLeft(v int64) *SearchTmOnsalesRequest {
	s.OrderPriceLeft = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetOrderPriceRight(v int64) *SearchTmOnsalesRequest {
	s.OrderPriceRight = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetRegLeft(v int32) *SearchTmOnsalesRequest {
	s.RegLeft = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetRegRight(v int32) *SearchTmOnsalesRequest {
	s.RegRight = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetSortName(v string) *SearchTmOnsalesRequest {
	s.SortName = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetSortOrder(v string) *SearchTmOnsalesRequest {
	s.SortOrder = &v
	return s
}

func (s *SearchTmOnsalesRequest) SetQueryAll(v bool) *SearchTmOnsalesRequest {
	s.QueryAll = &v
	return s
}

type SearchTmOnsalesResponseBody struct {
	RequestId       *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize        *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalPageNumber *int32                                   `json:"TotalPageNumber,omitempty" xml:"TotalPageNumber,omitempty"`
	TotalCount      *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Trademarks      []*SearchTmOnsalesResponseBodyTrademarks `json:"Trademarks,omitempty" xml:"Trademarks,omitempty" type:"Repeated"`
}

func (s SearchTmOnsalesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTmOnsalesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTmOnsalesResponseBody) SetRequestId(v string) *SearchTmOnsalesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTmOnsalesResponseBody) SetPageSize(v int32) *SearchTmOnsalesResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchTmOnsalesResponseBody) SetPageNumber(v int32) *SearchTmOnsalesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchTmOnsalesResponseBody) SetTotalPageNumber(v int32) *SearchTmOnsalesResponseBody {
	s.TotalPageNumber = &v
	return s
}

func (s *SearchTmOnsalesResponseBody) SetTotalCount(v int32) *SearchTmOnsalesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchTmOnsalesResponseBody) SetTrademarks(v []*SearchTmOnsalesResponseBodyTrademarks) *SearchTmOnsalesResponseBody {
	s.Trademarks = v
	return s
}

type SearchTmOnsalesResponseBodyTrademarks struct {
	TrademarkName      *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	Status             *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	ProductDesc        *string `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
	RegistrationNumber *string `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	Icon               *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	PartnerCode        *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	Classification     *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	Uid                *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	ProductCode        *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	OrderPrice         *string `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
}

func (s SearchTmOnsalesResponseBodyTrademarks) String() string {
	return tea.Prettify(s)
}

func (s SearchTmOnsalesResponseBodyTrademarks) GoString() string {
	return s.String()
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetTrademarkName(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.TrademarkName = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetStatus(v int64) *SearchTmOnsalesResponseBodyTrademarks {
	s.Status = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetProductDesc(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.ProductDesc = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetRegistrationNumber(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.RegistrationNumber = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetIcon(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.Icon = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetPartnerCode(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.PartnerCode = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetClassification(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.Classification = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetUid(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.Uid = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetProductCode(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.ProductCode = &v
	return s
}

func (s *SearchTmOnsalesResponseBodyTrademarks) SetOrderPrice(v string) *SearchTmOnsalesResponseBodyTrademarks {
	s.OrderPrice = &v
	return s
}

type SearchTmOnsalesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTmOnsalesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTmOnsalesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTmOnsalesResponse) GoString() string {
	return s.String()
}

func (s *SearchTmOnsalesResponse) SetHeaders(v map[string]*string) *SearchTmOnsalesResponse {
	s.Headers = v
	return s
}

func (s *SearchTmOnsalesResponse) SetBody(v *SearchTmOnsalesResponseBody) *SearchTmOnsalesResponse {
	s.Body = v
	return s
}

type GenerateUploadFilePolicyRequest struct {
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GenerateUploadFilePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyRequest) SetFileType(v string) *GenerateUploadFilePolicyRequest {
	s.FileType = &v
	return s
}

func (s *GenerateUploadFilePolicyRequest) SetBizId(v string) *GenerateUploadFilePolicyRequest {
	s.BizId = &v
	return s
}

type GenerateUploadFilePolicyResponseBody struct {
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExpireTime    *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	FileDir       *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	AccessId      *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
}

func (s GenerateUploadFilePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyResponseBody) SetSignature(v string) *GenerateUploadFilePolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetHost(v string) *GenerateUploadFilePolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetRequestId(v string) *GenerateUploadFilePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetExpireTime(v int64) *GenerateUploadFilePolicyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetEncodedPolicy(v string) *GenerateUploadFilePolicyResponseBody {
	s.EncodedPolicy = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetFileDir(v string) *GenerateUploadFilePolicyResponseBody {
	s.FileDir = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetAccessId(v string) *GenerateUploadFilePolicyResponseBody {
	s.AccessId = &v
	return s
}

type GenerateUploadFilePolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateUploadFilePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateUploadFilePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyResponse) SetHeaders(v map[string]*string) *GenerateUploadFilePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadFilePolicyResponse) SetBody(v *GenerateUploadFilePolicyResponseBody) *GenerateUploadFilePolicyResponse {
	s.Body = v
	return s
}

type DeleteMaterialRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaterialRequest) SetId(v int64) *DeleteMaterialRequest {
	s.Id = &v
	return s
}

type DeleteMaterialResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s DeleteMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaterialResponseBody) SetErrorMsg(v string) *DeleteMaterialResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteMaterialResponseBody) SetRequestId(v string) *DeleteMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaterialResponseBody) SetSuccess(v bool) *DeleteMaterialResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMaterialResponseBody) SetErrorCode(v string) *DeleteMaterialResponseBody {
	s.ErrorCode = &v
	return s
}

type DeleteMaterialResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaterialResponse) SetHeaders(v map[string]*string) *DeleteMaterialResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaterialResponse) SetBody(v *DeleteMaterialResponseBody) *DeleteMaterialResponse {
	s.Body = v
	return s
}

type WriteCommunicationLogRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note     *string `json:"Note,omitempty" xml:"Note,omitempty"`
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s WriteCommunicationLogRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteCommunicationLogRequest) GoString() string {
	return s.String()
}

func (s *WriteCommunicationLogRequest) SetBizId(v string) *WriteCommunicationLogRequest {
	s.BizId = &v
	return s
}

func (s *WriteCommunicationLogRequest) SetNote(v string) *WriteCommunicationLogRequest {
	s.Note = &v
	return s
}

func (s *WriteCommunicationLogRequest) SetTargetId(v string) *WriteCommunicationLogRequest {
	s.TargetId = &v
	return s
}

type WriteCommunicationLogResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s WriteCommunicationLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteCommunicationLogResponseBody) GoString() string {
	return s.String()
}

func (s *WriteCommunicationLogResponseBody) SetErrorMsg(v string) *WriteCommunicationLogResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *WriteCommunicationLogResponseBody) SetRequestId(v string) *WriteCommunicationLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *WriteCommunicationLogResponseBody) SetSuccess(v bool) *WriteCommunicationLogResponseBody {
	s.Success = &v
	return s
}

func (s *WriteCommunicationLogResponseBody) SetErrorCode(v string) *WriteCommunicationLogResponseBody {
	s.ErrorCode = &v
	return s
}

type WriteCommunicationLogResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WriteCommunicationLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WriteCommunicationLogResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteCommunicationLogResponse) GoString() string {
	return s.String()
}

func (s *WriteCommunicationLogResponse) SetHeaders(v map[string]*string) *WriteCommunicationLogResponse {
	s.Headers = v
	return s
}

func (s *WriteCommunicationLogResponse) SetBody(v *WriteCommunicationLogResponseBody) *WriteCommunicationLogResponse {
	s.Body = v
	return s
}

type InsertTradeIntentionUserRequest struct {
	RegisterNumber *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Mobile         *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Vcode          *string `json:"Vcode,omitempty" xml:"Vcode,omitempty"`
	PartnerCode    *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Channel        *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Ua             *string `json:"Ua,omitempty" xml:"Ua,omitempty"`
}

func (s InsertTradeIntentionUserRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertTradeIntentionUserRequest) GoString() string {
	return s.String()
}

func (s *InsertTradeIntentionUserRequest) SetRegisterNumber(v string) *InsertTradeIntentionUserRequest {
	s.RegisterNumber = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetClassification(v string) *InsertTradeIntentionUserRequest {
	s.Classification = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetType(v int32) *InsertTradeIntentionUserRequest {
	s.Type = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetMobile(v string) *InsertTradeIntentionUserRequest {
	s.Mobile = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetVcode(v string) *InsertTradeIntentionUserRequest {
	s.Vcode = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetPartnerCode(v string) *InsertTradeIntentionUserRequest {
	s.PartnerCode = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetUserName(v string) *InsertTradeIntentionUserRequest {
	s.UserName = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetDescription(v string) *InsertTradeIntentionUserRequest {
	s.Description = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetChannel(v string) *InsertTradeIntentionUserRequest {
	s.Channel = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetToken(v string) *InsertTradeIntentionUserRequest {
	s.Token = &v
	return s
}

func (s *InsertTradeIntentionUserRequest) SetUa(v string) *InsertTradeIntentionUserRequest {
	s.Ua = &v
	return s
}

type InsertTradeIntentionUserResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertTradeIntentionUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertTradeIntentionUserResponseBody) GoString() string {
	return s.String()
}

func (s *InsertTradeIntentionUserResponseBody) SetErrorMsg(v string) *InsertTradeIntentionUserResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InsertTradeIntentionUserResponseBody) SetRequestId(v string) *InsertTradeIntentionUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertTradeIntentionUserResponseBody) SetSuccess(v bool) *InsertTradeIntentionUserResponseBody {
	s.Success = &v
	return s
}

type InsertTradeIntentionUserResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertTradeIntentionUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertTradeIntentionUserResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertTradeIntentionUserResponse) GoString() string {
	return s.String()
}

func (s *InsertTradeIntentionUserResponse) SetHeaders(v map[string]*string) *InsertTradeIntentionUserResponse {
	s.Headers = v
	return s
}

func (s *InsertTradeIntentionUserResponse) SetBody(v *InsertTradeIntentionUserResponseBody) *InsertTradeIntentionUserResponse {
	s.Body = v
	return s
}

type QueryExtensionAttributeRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	AttributeKey *string `json:"AttributeKey,omitempty" xml:"AttributeKey,omitempty"`
}

func (s QueryExtensionAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryExtensionAttributeRequest) GoString() string {
	return s.String()
}

func (s *QueryExtensionAttributeRequest) SetBizId(v string) *QueryExtensionAttributeRequest {
	s.BizId = &v
	return s
}

func (s *QueryExtensionAttributeRequest) SetAttributeKey(v string) *QueryExtensionAttributeRequest {
	s.AttributeKey = &v
	return s
}

type QueryExtensionAttributeResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s QueryExtensionAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryExtensionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExtensionAttributeResponseBody) SetCode(v string) *QueryExtensionAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryExtensionAttributeResponseBody) SetMessage(v string) *QueryExtensionAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryExtensionAttributeResponseBody) SetRequestId(v string) *QueryExtensionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryExtensionAttributeResponseBody) SetSuccess(v bool) *QueryExtensionAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryExtensionAttributeResponseBody) SetAttributeValue(v string) *QueryExtensionAttributeResponseBody {
	s.AttributeValue = &v
	return s
}

type QueryExtensionAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryExtensionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryExtensionAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryExtensionAttributeResponse) GoString() string {
	return s.String()
}

func (s *QueryExtensionAttributeResponse) SetHeaders(v map[string]*string) *QueryExtensionAttributeResponse {
	s.Headers = v
	return s
}

func (s *QueryExtensionAttributeResponse) SetBody(v *QueryExtensionAttributeResponseBody) *QueryExtensionAttributeResponse {
	s.Body = v
	return s
}

type UpdateTrademarkOnsaleRequest struct {
	ClassificationCode      *string  `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	TmName                  *string  `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmIcon                  *string  `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	OriginalPrice           *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TmNumber                *string  `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	EndTime                 *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BeginTime               *int64   `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Description             *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Label                   *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	RegAnnDate              *int64   `json:"RegAnnDate,omitempty" xml:"RegAnnDate,omitempty"`
	OwnerName               *string  `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerEnName             *string  `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	SecondaryClassification *string  `json:"SecondaryClassification,omitempty" xml:"SecondaryClassification,omitempty"`
	ThirdClassification     *string  `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty"`
	Type                    *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Reason                  *string  `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s UpdateTrademarkOnsaleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrademarkOnsaleRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrademarkOnsaleRequest) SetClassificationCode(v string) *UpdateTrademarkOnsaleRequest {
	s.ClassificationCode = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetTmName(v string) *UpdateTrademarkOnsaleRequest {
	s.TmName = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetTmIcon(v string) *UpdateTrademarkOnsaleRequest {
	s.TmIcon = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetOriginalPrice(v float32) *UpdateTrademarkOnsaleRequest {
	s.OriginalPrice = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetTmNumber(v string) *UpdateTrademarkOnsaleRequest {
	s.TmNumber = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetEndTime(v int64) *UpdateTrademarkOnsaleRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetBeginTime(v int64) *UpdateTrademarkOnsaleRequest {
	s.BeginTime = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetDescription(v string) *UpdateTrademarkOnsaleRequest {
	s.Description = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetLabel(v string) *UpdateTrademarkOnsaleRequest {
	s.Label = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetRegAnnDate(v int64) *UpdateTrademarkOnsaleRequest {
	s.RegAnnDate = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetOwnerName(v string) *UpdateTrademarkOnsaleRequest {
	s.OwnerName = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetOwnerEnName(v string) *UpdateTrademarkOnsaleRequest {
	s.OwnerEnName = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetSecondaryClassification(v string) *UpdateTrademarkOnsaleRequest {
	s.SecondaryClassification = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetThirdClassification(v string) *UpdateTrademarkOnsaleRequest {
	s.ThirdClassification = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetType(v string) *UpdateTrademarkOnsaleRequest {
	s.Type = &v
	return s
}

func (s *UpdateTrademarkOnsaleRequest) SetReason(v string) *UpdateTrademarkOnsaleRequest {
	s.Reason = &v
	return s
}

type UpdateTrademarkOnsaleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrademarkOnsaleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrademarkOnsaleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrademarkOnsaleResponseBody) SetRequestId(v string) *UpdateTrademarkOnsaleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTrademarkOnsaleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTrademarkOnsaleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTrademarkOnsaleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrademarkOnsaleResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrademarkOnsaleResponse) SetHeaders(v map[string]*string) *UpdateTrademarkOnsaleResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrademarkOnsaleResponse) SetBody(v *UpdateTrademarkOnsaleResponseBody) *UpdateTrademarkOnsaleResponse {
	s.Body = v
	return s
}

type QueryTradeProduceDetailRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s QueryTradeProduceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceDetailRequest) SetBizId(v string) *QueryTradeProduceDetailRequest {
	s.BizId = &v
	return s
}

type QueryTradeProduceDetailResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryTradeProduceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTradeProduceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceDetailResponseBody) SetRequestId(v string) *QueryTradeProduceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBody) SetData(v *QueryTradeProduceDetailResponseBodyData) *QueryTradeProduceDetailResponseBody {
	s.Data = v
	return s
}

type QueryTradeProduceDetailResponseBodyData struct {
	UpdateTime         *int64                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ThirdCode          *string                `json:"ThirdCode,omitempty" xml:"ThirdCode,omitempty"`
	Share              *string                `json:"Share,omitempty" xml:"Share,omitempty"`
	PreAmount          *float32               `json:"PreAmount,omitempty" xml:"PreAmount,omitempty"`
	CreateTime         *int64                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId             *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RefundAmount       *float32               `json:"RefundAmount,omitempty" xml:"RefundAmount,omitempty"`
	Icon               *string                `json:"Icon,omitempty" xml:"Icon,omitempty"`
	BizId              *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BuyerStatus        *int32                 `json:"BuyerStatus,omitempty" xml:"BuyerStatus,omitempty"`
	Source             *int32                 `json:"Source,omitempty" xml:"Source,omitempty"`
	ConfiscateAmount   *float32               `json:"ConfiscateAmount,omitempty" xml:"ConfiscateAmount,omitempty"`
	OperateNote        *string                `json:"OperateNote,omitempty" xml:"OperateNote,omitempty"`
	PreOrderId         *string                `json:"PreOrderId,omitempty" xml:"PreOrderId,omitempty"`
	Extend             map[string]interface{} `json:"Extend,omitempty" xml:"Extend,omitempty"`
	TmName             *string                `json:"TmName,omitempty" xml:"TmName,omitempty"`
	ExclusiveDateLimit *string                `json:"ExclusiveDateLimit,omitempty" xml:"ExclusiveDateLimit,omitempty"`
	AllowCancel        *bool                  `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	RegisterNumber     *string                `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	FinalAmount        *float32               `json:"FinalAmount,omitempty" xml:"FinalAmount,omitempty"`
	Classification     *string                `json:"Classification,omitempty" xml:"Classification,omitempty"`
	PaidAmount         *float32               `json:"PaidAmount,omitempty" xml:"PaidAmount,omitempty"`
}

func (s QueryTradeProduceDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceDetailResponseBodyData) SetUpdateTime(v int64) *QueryTradeProduceDetailResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetThirdCode(v string) *QueryTradeProduceDetailResponseBodyData {
	s.ThirdCode = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetShare(v string) *QueryTradeProduceDetailResponseBodyData {
	s.Share = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetPreAmount(v float32) *QueryTradeProduceDetailResponseBodyData {
	s.PreAmount = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetCreateTime(v int64) *QueryTradeProduceDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetUserId(v string) *QueryTradeProduceDetailResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetRefundAmount(v float32) *QueryTradeProduceDetailResponseBodyData {
	s.RefundAmount = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetIcon(v string) *QueryTradeProduceDetailResponseBodyData {
	s.Icon = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetBizId(v string) *QueryTradeProduceDetailResponseBodyData {
	s.BizId = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetBuyerStatus(v int32) *QueryTradeProduceDetailResponseBodyData {
	s.BuyerStatus = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetSource(v int32) *QueryTradeProduceDetailResponseBodyData {
	s.Source = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetConfiscateAmount(v float32) *QueryTradeProduceDetailResponseBodyData {
	s.ConfiscateAmount = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetOperateNote(v string) *QueryTradeProduceDetailResponseBodyData {
	s.OperateNote = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetPreOrderId(v string) *QueryTradeProduceDetailResponseBodyData {
	s.PreOrderId = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetExtend(v map[string]interface{}) *QueryTradeProduceDetailResponseBodyData {
	s.Extend = v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetTmName(v string) *QueryTradeProduceDetailResponseBodyData {
	s.TmName = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetExclusiveDateLimit(v string) *QueryTradeProduceDetailResponseBodyData {
	s.ExclusiveDateLimit = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetAllowCancel(v bool) *QueryTradeProduceDetailResponseBodyData {
	s.AllowCancel = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetRegisterNumber(v string) *QueryTradeProduceDetailResponseBodyData {
	s.RegisterNumber = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetFinalAmount(v float32) *QueryTradeProduceDetailResponseBodyData {
	s.FinalAmount = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetClassification(v string) *QueryTradeProduceDetailResponseBodyData {
	s.Classification = &v
	return s
}

func (s *QueryTradeProduceDetailResponseBodyData) SetPaidAmount(v float32) *QueryTradeProduceDetailResponseBodyData {
	s.PaidAmount = &v
	return s
}

type QueryTradeProduceDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTradeProduceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTradeProduceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeProduceDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeProduceDetailResponse) SetHeaders(v map[string]*string) *QueryTradeProduceDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTradeProduceDetailResponse) SetBody(v *QueryTradeProduceDetailResponseBody) *QueryTradeProduceDetailResponse {
	s.Body = v
	return s
}

type QueryQrCodeUploadStatusRequest struct {
	OssKey   *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryQrCodeUploadStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryQrCodeUploadStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryQrCodeUploadStatusRequest) SetOssKey(v string) *QueryQrCodeUploadStatusRequest {
	s.OssKey = &v
	return s
}

func (s *QueryQrCodeUploadStatusRequest) SetFieldKey(v string) *QueryQrCodeUploadStatusRequest {
	s.FieldKey = &v
	return s
}

func (s *QueryQrCodeUploadStatusRequest) SetUuid(v string) *QueryQrCodeUploadStatusRequest {
	s.Uuid = &v
	return s
}

type QueryQrCodeUploadStatusResponseBody struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OssUrl    *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	OssKey    *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryQrCodeUploadStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryQrCodeUploadStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQrCodeUploadStatusResponseBody) SetStatus(v int32) *QueryQrCodeUploadStatusResponseBody {
	s.Status = &v
	return s
}

func (s *QueryQrCodeUploadStatusResponseBody) SetRequestId(v string) *QueryQrCodeUploadStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryQrCodeUploadStatusResponseBody) SetOssUrl(v string) *QueryQrCodeUploadStatusResponseBody {
	s.OssUrl = &v
	return s
}

func (s *QueryQrCodeUploadStatusResponseBody) SetOssKey(v string) *QueryQrCodeUploadStatusResponseBody {
	s.OssKey = &v
	return s
}

func (s *QueryQrCodeUploadStatusResponseBody) SetSuccess(v bool) *QueryQrCodeUploadStatusResponseBody {
	s.Success = &v
	return s
}

type QueryQrCodeUploadStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryQrCodeUploadStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryQrCodeUploadStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryQrCodeUploadStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryQrCodeUploadStatusResponse) SetHeaders(v map[string]*string) *QueryQrCodeUploadStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryQrCodeUploadStatusResponse) SetBody(v *QueryQrCodeUploadStatusResponseBody) *QueryQrCodeUploadStatusResponse {
	s.Body = v
	return s
}

type RejectApplicantRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Note       *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s RejectApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectApplicantRequest) GoString() string {
	return s.String()
}

func (s *RejectApplicantRequest) SetInstanceId(v string) *RejectApplicantRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectApplicantRequest) SetNote(v string) *RejectApplicantRequest {
	s.Note = &v
	return s
}

type RejectApplicantResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *RejectApplicantResponseBody) SetRequestId(v string) *RejectApplicantResponseBody {
	s.RequestId = &v
	return s
}

type RejectApplicantResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RejectApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RejectApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectApplicantResponse) GoString() string {
	return s.String()
}

func (s *RejectApplicantResponse) SetHeaders(v map[string]*string) *RejectApplicantResponse {
	s.Headers = v
	return s
}

func (s *RejectApplicantResponse) SetBody(v *RejectApplicantResponseBody) *RejectApplicantResponse {
	s.Body = v
	return s
}

type QueryTradeIntentionUserListRequest struct {
	Begin    *int64  `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End      *int64  `json:"End,omitempty" xml:"End,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryTradeIntentionUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeIntentionUserListRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeIntentionUserListRequest) SetBegin(v int64) *QueryTradeIntentionUserListRequest {
	s.Begin = &v
	return s
}

func (s *QueryTradeIntentionUserListRequest) SetEnd(v int64) *QueryTradeIntentionUserListRequest {
	s.End = &v
	return s
}

func (s *QueryTradeIntentionUserListRequest) SetPageSize(v int32) *QueryTradeIntentionUserListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTradeIntentionUserListRequest) SetPageNum(v int32) *QueryTradeIntentionUserListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTradeIntentionUserListRequest) SetBizId(v string) *QueryTradeIntentionUserListRequest {
	s.BizId = &v
	return s
}

func (s *QueryTradeIntentionUserListRequest) SetStatus(v int32) *QueryTradeIntentionUserListRequest {
	s.Status = &v
	return s
}

type QueryTradeIntentionUserListResponseBody struct {
	CurrentPageNum *int32                                       `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                                       `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                                       `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryTradeIntentionUserListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTradeIntentionUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeIntentionUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeIntentionUserListResponseBody) SetCurrentPageNum(v int32) *QueryTradeIntentionUserListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBody) SetTotalPageNum(v int32) *QueryTradeIntentionUserListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBody) SetRequestId(v string) *QueryTradeIntentionUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBody) SetPageSize(v int32) *QueryTradeIntentionUserListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBody) SetTotalItemNum(v int32) *QueryTradeIntentionUserListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBody) SetData(v *QueryTradeIntentionUserListResponseBodyData) *QueryTradeIntentionUserListResponseBody {
	s.Data = v
	return s
}

type QueryTradeIntentionUserListResponseBodyData struct {
	Trademark []*QueryTradeIntentionUserListResponseBodyDataTrademark `json:"Trademark,omitempty" xml:"Trademark,omitempty" type:"Repeated"`
}

func (s QueryTradeIntentionUserListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeIntentionUserListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTradeIntentionUserListResponseBodyData) SetTrademark(v []*QueryTradeIntentionUserListResponseBodyDataTrademark) *QueryTradeIntentionUserListResponseBodyData {
	s.Trademark = v
	return s
}

type QueryTradeIntentionUserListResponseBodyDataTrademark struct {
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Mobile         *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	RegisterNumber *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryTradeIntentionUserListResponseBodyDataTrademark) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeIntentionUserListResponseBodyDataTrademark) GoString() string {
	return s.String()
}

func (s *QueryTradeIntentionUserListResponseBodyDataTrademark) SetType(v int32) *QueryTradeIntentionUserListResponseBodyDataTrademark {
	s.Type = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBodyDataTrademark) SetStatus(v int32) *QueryTradeIntentionUserListResponseBodyDataTrademark {
	s.Status = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBodyDataTrademark) SetDescription(v string) *QueryTradeIntentionUserListResponseBodyDataTrademark {
	s.Description = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBodyDataTrademark) SetMobile(v string) *QueryTradeIntentionUserListResponseBodyDataTrademark {
	s.Mobile = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBodyDataTrademark) SetRegisterNumber(v string) *QueryTradeIntentionUserListResponseBodyDataTrademark {
	s.RegisterNumber = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBodyDataTrademark) SetBizId(v string) *QueryTradeIntentionUserListResponseBodyDataTrademark {
	s.BizId = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBodyDataTrademark) SetClassification(v string) *QueryTradeIntentionUserListResponseBodyDataTrademark {
	s.Classification = &v
	return s
}

func (s *QueryTradeIntentionUserListResponseBodyDataTrademark) SetUserName(v string) *QueryTradeIntentionUserListResponseBodyDataTrademark {
	s.UserName = &v
	return s
}

type QueryTradeIntentionUserListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTradeIntentionUserListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTradeIntentionUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeIntentionUserListResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeIntentionUserListResponse) SetHeaders(v map[string]*string) *QueryTradeIntentionUserListResponse {
	s.Headers = v
	return s
}

func (s *QueryTradeIntentionUserListResponse) SetBody(v *QueryTradeIntentionUserListResponseBody) *QueryTradeIntentionUserListResponse {
	s.Body = v
	return s
}

type StoreMaterialTemporarilyRequest struct {
	ContactZipcode        *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	Type                  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Region                *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ContactName           *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber         *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactEmail          *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactAddress        *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	LoaOssKey             *string `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CardNumber            *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Province              *string `json:"Province,omitempty" xml:"Province,omitempty"`
	City                  *string `json:"City,omitempty" xml:"City,omitempty"`
	Town                  *string `json:"Town,omitempty" xml:"Town,omitempty"`
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EName                 *string `json:"EName,omitempty" xml:"EName,omitempty"`
	EAddress              *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	Country               *string `json:"Country,omitempty" xml:"Country,omitempty"`
	IdCardOssKey          *string `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	BusinessLicenceOssKey *string `json:"BusinessLicenceOssKey,omitempty" xml:"BusinessLicenceOssKey,omitempty"`
	PassportOssKey        *string `json:"PassportOssKey,omitempty" xml:"PassportOssKey,omitempty"`
	LegalNoticeOssKey     *string `json:"LegalNoticeOssKey,omitempty" xml:"LegalNoticeOssKey,omitempty"`
	PrincipalName         *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	ContactProvince       *string `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	ContactCity           *string `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	ContactDistrict       *string `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	ContactCounty         *string `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
}

func (s StoreMaterialTemporarilyRequest) String() string {
	return tea.Prettify(s)
}

func (s StoreMaterialTemporarilyRequest) GoString() string {
	return s.String()
}

func (s *StoreMaterialTemporarilyRequest) SetContactZipcode(v string) *StoreMaterialTemporarilyRequest {
	s.ContactZipcode = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetType(v string) *StoreMaterialTemporarilyRequest {
	s.Type = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetRegion(v string) *StoreMaterialTemporarilyRequest {
	s.Region = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetContactName(v string) *StoreMaterialTemporarilyRequest {
	s.ContactName = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetContactNumber(v string) *StoreMaterialTemporarilyRequest {
	s.ContactNumber = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetContactEmail(v string) *StoreMaterialTemporarilyRequest {
	s.ContactEmail = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetContactAddress(v string) *StoreMaterialTemporarilyRequest {
	s.ContactAddress = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetLoaOssKey(v string) *StoreMaterialTemporarilyRequest {
	s.LoaOssKey = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetName(v string) *StoreMaterialTemporarilyRequest {
	s.Name = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetCardNumber(v string) *StoreMaterialTemporarilyRequest {
	s.CardNumber = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetProvince(v string) *StoreMaterialTemporarilyRequest {
	s.Province = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetCity(v string) *StoreMaterialTemporarilyRequest {
	s.City = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetTown(v string) *StoreMaterialTemporarilyRequest {
	s.Town = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetAddress(v string) *StoreMaterialTemporarilyRequest {
	s.Address = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetEName(v string) *StoreMaterialTemporarilyRequest {
	s.EName = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetEAddress(v string) *StoreMaterialTemporarilyRequest {
	s.EAddress = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetCountry(v string) *StoreMaterialTemporarilyRequest {
	s.Country = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetIdCardOssKey(v string) *StoreMaterialTemporarilyRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetBusinessLicenceOssKey(v string) *StoreMaterialTemporarilyRequest {
	s.BusinessLicenceOssKey = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetPassportOssKey(v string) *StoreMaterialTemporarilyRequest {
	s.PassportOssKey = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetLegalNoticeOssKey(v string) *StoreMaterialTemporarilyRequest {
	s.LegalNoticeOssKey = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetPrincipalName(v int32) *StoreMaterialTemporarilyRequest {
	s.PrincipalName = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetContactProvince(v string) *StoreMaterialTemporarilyRequest {
	s.ContactProvince = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetContactCity(v string) *StoreMaterialTemporarilyRequest {
	s.ContactCity = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetContactDistrict(v string) *StoreMaterialTemporarilyRequest {
	s.ContactDistrict = &v
	return s
}

func (s *StoreMaterialTemporarilyRequest) SetContactCounty(v string) *StoreMaterialTemporarilyRequest {
	s.ContactCounty = &v
	return s
}

type StoreMaterialTemporarilyResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StoreMaterialTemporarilyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StoreMaterialTemporarilyResponseBody) GoString() string {
	return s.String()
}

func (s *StoreMaterialTemporarilyResponseBody) SetSuccess(v bool) *StoreMaterialTemporarilyResponseBody {
	s.Success = &v
	return s
}

func (s *StoreMaterialTemporarilyResponseBody) SetRequestId(v string) *StoreMaterialTemporarilyResponseBody {
	s.RequestId = &v
	return s
}

type StoreMaterialTemporarilyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StoreMaterialTemporarilyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StoreMaterialTemporarilyResponse) String() string {
	return tea.Prettify(s)
}

func (s StoreMaterialTemporarilyResponse) GoString() string {
	return s.String()
}

func (s *StoreMaterialTemporarilyResponse) SetHeaders(v map[string]*string) *StoreMaterialTemporarilyResponse {
	s.Headers = v
	return s
}

func (s *StoreMaterialTemporarilyResponse) SetBody(v *StoreMaterialTemporarilyResponseBody) *StoreMaterialTemporarilyResponse {
	s.Body = v
	return s
}

type RefuseAdditionalMaterialRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note  *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s RefuseAdditionalMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s RefuseAdditionalMaterialRequest) GoString() string {
	return s.String()
}

func (s *RefuseAdditionalMaterialRequest) SetBizId(v string) *RefuseAdditionalMaterialRequest {
	s.BizId = &v
	return s
}

func (s *RefuseAdditionalMaterialRequest) SetNote(v string) *RefuseAdditionalMaterialRequest {
	s.Note = &v
	return s
}

type RefuseAdditionalMaterialResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s RefuseAdditionalMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefuseAdditionalMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *RefuseAdditionalMaterialResponseBody) SetErrorMsg(v string) *RefuseAdditionalMaterialResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefuseAdditionalMaterialResponseBody) SetRequestId(v string) *RefuseAdditionalMaterialResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefuseAdditionalMaterialResponseBody) SetSuccess(v bool) *RefuseAdditionalMaterialResponseBody {
	s.Success = &v
	return s
}

func (s *RefuseAdditionalMaterialResponseBody) SetErrorCode(v string) *RefuseAdditionalMaterialResponseBody {
	s.ErrorCode = &v
	return s
}

type RefuseAdditionalMaterialResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefuseAdditionalMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefuseAdditionalMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s RefuseAdditionalMaterialResponse) GoString() string {
	return s.String()
}

func (s *RefuseAdditionalMaterialResponse) SetHeaders(v map[string]*string) *RefuseAdditionalMaterialResponse {
	s.Headers = v
	return s
}

func (s *RefuseAdditionalMaterialResponse) SetBody(v *RefuseAdditionalMaterialResponseBody) *RefuseAdditionalMaterialResponse {
	s.Body = v
	return s
}

type ListNotaryInfosRequest struct {
	NotaryType *int32  `json:"NotaryType,omitempty" xml:"NotaryType,omitempty"`
	BizOrderNo *string `json:"BizOrderNo,omitempty" xml:"BizOrderNo,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	PageNum    *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNotaryInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryInfosRequest) GoString() string {
	return s.String()
}

func (s *ListNotaryInfosRequest) SetNotaryType(v int32) *ListNotaryInfosRequest {
	s.NotaryType = &v
	return s
}

func (s *ListNotaryInfosRequest) SetBizOrderNo(v string) *ListNotaryInfosRequest {
	s.BizOrderNo = &v
	return s
}

func (s *ListNotaryInfosRequest) SetToken(v string) *ListNotaryInfosRequest {
	s.Token = &v
	return s
}

func (s *ListNotaryInfosRequest) SetPageNum(v int32) *ListNotaryInfosRequest {
	s.PageNum = &v
	return s
}

func (s *ListNotaryInfosRequest) SetPageSize(v int32) *ListNotaryInfosRequest {
	s.PageSize = &v
	return s
}

type ListNotaryInfosResponseBody struct {
	NextPage       *bool                            `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode      *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	TotalItemNum   *int32                           `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	PrePage        *bool                            `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	CurrentPageNum *int32                           `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	ErrorMsg       *string                          `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	TotalPageNum   *int32                           `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	PageSize       *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data           *ListNotaryInfosResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListNotaryInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListNotaryInfosResponseBody) SetNextPage(v bool) *ListNotaryInfosResponseBody {
	s.NextPage = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetRequestId(v string) *ListNotaryInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetSuccess(v bool) *ListNotaryInfosResponseBody {
	s.Success = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetErrorCode(v string) *ListNotaryInfosResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetTotalItemNum(v int32) *ListNotaryInfosResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetPrePage(v bool) *ListNotaryInfosResponseBody {
	s.PrePage = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetCurrentPageNum(v int32) *ListNotaryInfosResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetErrorMsg(v string) *ListNotaryInfosResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetTotalPageNum(v int32) *ListNotaryInfosResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetPageSize(v int32) *ListNotaryInfosResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNotaryInfosResponseBody) SetData(v *ListNotaryInfosResponseBodyData) *ListNotaryInfosResponseBody {
	s.Data = v
	return s
}

type ListNotaryInfosResponseBodyData struct {
	NotaryInfo []*ListNotaryInfosResponseBodyDataNotaryInfo `json:"NotaryInfo,omitempty" xml:"NotaryInfo,omitempty" type:"Repeated"`
}

func (s ListNotaryInfosResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryInfosResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNotaryInfosResponseBodyData) SetNotaryInfo(v []*ListNotaryInfosResponseBodyDataNotaryInfo) *ListNotaryInfosResponseBodyData {
	s.NotaryInfo = v
	return s
}

type ListNotaryInfosResponseBodyDataNotaryInfo struct {
	Token              *string `json:"Token,omitempty" xml:"Token,omitempty"`
	TmRegisterNo       *string `json:"TmRegisterNo,omitempty" xml:"TmRegisterNo,omitempty"`
	TmClassification   *string `json:"TmClassification,omitempty" xml:"TmClassification,omitempty"`
	NotaryFailedReason *string `json:"NotaryFailedReason,omitempty" xml:"NotaryFailedReason,omitempty"`
	GmtModified        *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NotaryStatus       *int32  `json:"NotaryStatus,omitempty" xml:"NotaryStatus,omitempty"`
	BizOrderNo         *string `json:"BizOrderNo,omitempty" xml:"BizOrderNo,omitempty"`
}

func (s ListNotaryInfosResponseBodyDataNotaryInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryInfosResponseBodyDataNotaryInfo) GoString() string {
	return s.String()
}

func (s *ListNotaryInfosResponseBodyDataNotaryInfo) SetToken(v string) *ListNotaryInfosResponseBodyDataNotaryInfo {
	s.Token = &v
	return s
}

func (s *ListNotaryInfosResponseBodyDataNotaryInfo) SetTmRegisterNo(v string) *ListNotaryInfosResponseBodyDataNotaryInfo {
	s.TmRegisterNo = &v
	return s
}

func (s *ListNotaryInfosResponseBodyDataNotaryInfo) SetTmClassification(v string) *ListNotaryInfosResponseBodyDataNotaryInfo {
	s.TmClassification = &v
	return s
}

func (s *ListNotaryInfosResponseBodyDataNotaryInfo) SetNotaryFailedReason(v string) *ListNotaryInfosResponseBodyDataNotaryInfo {
	s.NotaryFailedReason = &v
	return s
}

func (s *ListNotaryInfosResponseBodyDataNotaryInfo) SetGmtModified(v int64) *ListNotaryInfosResponseBodyDataNotaryInfo {
	s.GmtModified = &v
	return s
}

func (s *ListNotaryInfosResponseBodyDataNotaryInfo) SetNotaryStatus(v int32) *ListNotaryInfosResponseBodyDataNotaryInfo {
	s.NotaryStatus = &v
	return s
}

func (s *ListNotaryInfosResponseBodyDataNotaryInfo) SetBizOrderNo(v string) *ListNotaryInfosResponseBodyDataNotaryInfo {
	s.BizOrderNo = &v
	return s
}

type ListNotaryInfosResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNotaryInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNotaryInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryInfosResponse) GoString() string {
	return s.String()
}

func (s *ListNotaryInfosResponse) SetHeaders(v map[string]*string) *ListNotaryInfosResponse {
	s.Headers = v
	return s
}

func (s *ListNotaryInfosResponse) SetBody(v *ListNotaryInfosResponseBody) *ListNotaryInfosResponse {
	s.Body = v
	return s
}

type GetDefaultPrincipalNameRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s GetDefaultPrincipalNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultPrincipalNameRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultPrincipalNameRequest) SetBizType(v string) *GetDefaultPrincipalNameRequest {
	s.BizType = &v
	return s
}

type GetDefaultPrincipalNameResponseBody struct {
	PrincipalName *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDefaultPrincipalNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultPrincipalNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultPrincipalNameResponseBody) SetPrincipalName(v int32) *GetDefaultPrincipalNameResponseBody {
	s.PrincipalName = &v
	return s
}

func (s *GetDefaultPrincipalNameResponseBody) SetRequestId(v string) *GetDefaultPrincipalNameResponseBody {
	s.RequestId = &v
	return s
}

type GetDefaultPrincipalNameResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDefaultPrincipalNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDefaultPrincipalNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultPrincipalNameResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultPrincipalNameResponse) SetHeaders(v map[string]*string) *GetDefaultPrincipalNameResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultPrincipalNameResponse) SetBody(v *GetDefaultPrincipalNameResponseBody) *GetDefaultPrincipalNameResponse {
	s.Body = v
	return s
}

type QueryTradeMarkApplicationDetailRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s QueryTradeMarkApplicationDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailRequest) SetBizId(v string) *QueryTradeMarkApplicationDetailRequest {
	s.BizId = &v
	return s
}

type QueryTradeMarkApplicationDetailResponseBody struct {
	Type                    *int32                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	Status                  *int32                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	AcceptUrl               *string                                                         `json:"AcceptUrl,omitempty" xml:"AcceptUrl,omitempty"`
	OrderPrice              *float32                                                        `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	SubmitAuditTime         *int64                                                          `json:"SubmitAuditTime,omitempty" xml:"SubmitAuditTime,omitempty"`
	UpdateTime              *int64                                                          `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime              *int64                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NotAcceptUrl            *string                                                         `json:"NotAcceptUrl,omitempty" xml:"NotAcceptUrl,omitempty"`
	SendTime                *string                                                         `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	ServicePrice            *float32                                                        `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	PartnerMobile           *string                                                         `json:"PartnerMobile,omitempty" xml:"PartnerMobile,omitempty"`
	RecvUserLogistics       *string                                                         `json:"RecvUserLogistics,omitempty" xml:"RecvUserLogistics,omitempty"`
	RequestId               *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GrayIconUrl             *string                                                         `json:"GrayIconUrl,omitempty" xml:"GrayIconUrl,omitempty"`
	MaterialId              *int64                                                          `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	SendSbjLogistics        *string                                                         `json:"SendSbjLogistics,omitempty" xml:"SendSbjLogistics,omitempty"`
	SendUserLogistics       *string                                                         `json:"SendUserLogistics,omitempty" xml:"SendUserLogistics,omitempty"`
	LoaUrl                  *string                                                         `json:"LoaUrl,omitempty" xml:"LoaUrl,omitempty"`
	TmNumber                *string                                                         `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	Note                    *string                                                         `json:"Note,omitempty" xml:"Note,omitempty"`
	PrincipalName           *int32                                                          `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	PartnerName             *string                                                         `json:"PartnerName,omitempty" xml:"PartnerName,omitempty"`
	LogisticsCertificateUrl *string                                                         `json:"LogisticsCertificateUrl,omitempty" xml:"LogisticsCertificateUrl,omitempty"`
	BizId                   *string                                                         `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PartnerCode             *string                                                         `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	TmNameType              *int32                                                          `json:"TmNameType,omitempty" xml:"TmNameType,omitempty"`
	ExtendInfo              map[string]interface{}                                          `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	TmIcon                  *string                                                         `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	TmName                  *string                                                         `json:"TmName,omitempty" xml:"TmName,omitempty"`
	LogisticsNo             *string                                                         `json:"LogisticsNo,omitempty" xml:"LogisticsNo,omitempty"`
	TotalPrice              *float32                                                        `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	SubmitTime              *int64                                                          `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	OrderId                 *string                                                         `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ReceiptUrl              *QueryTradeMarkApplicationDetailResponseBodyReceiptUrl          `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty" type:"Struct"`
	JudgeResultUrl          *QueryTradeMarkApplicationDetailResponseBodyJudgeResultUrl      `json:"JudgeResultUrl,omitempty" xml:"JudgeResultUrl,omitempty" type:"Struct"`
	Flags                   *QueryTradeMarkApplicationDetailResponseBodyFlags               `json:"Flags,omitempty" xml:"Flags,omitempty" type:"Struct"`
	AdminUploads            *QueryTradeMarkApplicationDetailResponseBodyAdminUploads        `json:"AdminUploads,omitempty" xml:"AdminUploads,omitempty" type:"Struct"`
	FirstClassification     *QueryTradeMarkApplicationDetailResponseBodyFirstClassification `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
	MaterialDetail          *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail      `json:"MaterialDetail,omitempty" xml:"MaterialDetail,omitempty" type:"Struct"`
	RenewResponse           *QueryTradeMarkApplicationDetailResponseBodyRenewResponse       `json:"RenewResponse,omitempty" xml:"RenewResponse,omitempty" type:"Struct"`
	ReviewOfficialFiles     *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles `json:"ReviewOfficialFiles,omitempty" xml:"ReviewOfficialFiles,omitempty" type:"Struct"`
	Supplements             *QueryTradeMarkApplicationDetailResponseBodySupplements         `json:"Supplements,omitempty" xml:"Supplements,omitempty" type:"Struct"`
	ThirdClassification     *QueryTradeMarkApplicationDetailResponseBodyThirdClassification `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetType(v int32) *QueryTradeMarkApplicationDetailResponseBody {
	s.Type = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetStatus(v int32) *QueryTradeMarkApplicationDetailResponseBody {
	s.Status = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetAcceptUrl(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.AcceptUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetOrderPrice(v float32) *QueryTradeMarkApplicationDetailResponseBody {
	s.OrderPrice = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetSubmitAuditTime(v int64) *QueryTradeMarkApplicationDetailResponseBody {
	s.SubmitAuditTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetUpdateTime(v int64) *QueryTradeMarkApplicationDetailResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetCreateTime(v int64) *QueryTradeMarkApplicationDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetNotAcceptUrl(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.NotAcceptUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetSendTime(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.SendTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetServicePrice(v float32) *QueryTradeMarkApplicationDetailResponseBody {
	s.ServicePrice = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetPartnerMobile(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.PartnerMobile = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetRecvUserLogistics(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.RecvUserLogistics = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetRequestId(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetGrayIconUrl(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.GrayIconUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetMaterialId(v int64) *QueryTradeMarkApplicationDetailResponseBody {
	s.MaterialId = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetSendSbjLogistics(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.SendSbjLogistics = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetSendUserLogistics(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.SendUserLogistics = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetLoaUrl(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.LoaUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetTmNumber(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.TmNumber = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetNote(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.Note = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetPrincipalName(v int32) *QueryTradeMarkApplicationDetailResponseBody {
	s.PrincipalName = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetPartnerName(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.PartnerName = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetLogisticsCertificateUrl(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.LogisticsCertificateUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetBizId(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.BizId = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetPartnerCode(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.PartnerCode = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetTmNameType(v int32) *QueryTradeMarkApplicationDetailResponseBody {
	s.TmNameType = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetExtendInfo(v map[string]interface{}) *QueryTradeMarkApplicationDetailResponseBody {
	s.ExtendInfo = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetTmIcon(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.TmIcon = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetTmName(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.TmName = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetLogisticsNo(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.LogisticsNo = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetTotalPrice(v float32) *QueryTradeMarkApplicationDetailResponseBody {
	s.TotalPrice = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetSubmitTime(v int64) *QueryTradeMarkApplicationDetailResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetOrderId(v string) *QueryTradeMarkApplicationDetailResponseBody {
	s.OrderId = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetReceiptUrl(v *QueryTradeMarkApplicationDetailResponseBodyReceiptUrl) *QueryTradeMarkApplicationDetailResponseBody {
	s.ReceiptUrl = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetJudgeResultUrl(v *QueryTradeMarkApplicationDetailResponseBodyJudgeResultUrl) *QueryTradeMarkApplicationDetailResponseBody {
	s.JudgeResultUrl = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetFlags(v *QueryTradeMarkApplicationDetailResponseBodyFlags) *QueryTradeMarkApplicationDetailResponseBody {
	s.Flags = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetAdminUploads(v *QueryTradeMarkApplicationDetailResponseBodyAdminUploads) *QueryTradeMarkApplicationDetailResponseBody {
	s.AdminUploads = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetFirstClassification(v *QueryTradeMarkApplicationDetailResponseBodyFirstClassification) *QueryTradeMarkApplicationDetailResponseBody {
	s.FirstClassification = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetMaterialDetail(v *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) *QueryTradeMarkApplicationDetailResponseBody {
	s.MaterialDetail = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetRenewResponse(v *QueryTradeMarkApplicationDetailResponseBodyRenewResponse) *QueryTradeMarkApplicationDetailResponseBody {
	s.RenewResponse = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetReviewOfficialFiles(v *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles) *QueryTradeMarkApplicationDetailResponseBody {
	s.ReviewOfficialFiles = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetSupplements(v *QueryTradeMarkApplicationDetailResponseBodySupplements) *QueryTradeMarkApplicationDetailResponseBody {
	s.Supplements = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBody) SetThirdClassification(v *QueryTradeMarkApplicationDetailResponseBodyThirdClassification) *QueryTradeMarkApplicationDetailResponseBody {
	s.ThirdClassification = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyReceiptUrl struct {
	ReceiptUrl []*string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyReceiptUrl) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyReceiptUrl) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyReceiptUrl) SetReceiptUrl(v []*string) *QueryTradeMarkApplicationDetailResponseBodyReceiptUrl {
	s.ReceiptUrl = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyJudgeResultUrl struct {
	JudgeResultUrl []*string `json:"JudgeResultUrl,omitempty" xml:"JudgeResultUrl,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyJudgeResultUrl) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyJudgeResultUrl) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyJudgeResultUrl) SetJudgeResultUrl(v []*string) *QueryTradeMarkApplicationDetailResponseBodyJudgeResultUrl {
	s.JudgeResultUrl = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyFlags struct {
	Flag []*int32 `json:"Flag,omitempty" xml:"Flag,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyFlags) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyFlags) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyFlags) SetFlag(v []*int32) *QueryTradeMarkApplicationDetailResponseBodyFlags {
	s.Flag = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyAdminUploads struct {
	LoaPicUrl     *string `json:"LoaPicUrl,omitempty" xml:"LoaPicUrl,omitempty"`
	LicensePicUrl *string `json:"LicensePicUrl,omitempty" xml:"LicensePicUrl,omitempty"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyAdminUploads) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyAdminUploads) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyAdminUploads) SetLoaPicUrl(v string) *QueryTradeMarkApplicationDetailResponseBodyAdminUploads {
	s.LoaPicUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyAdminUploads) SetLicensePicUrl(v string) *QueryTradeMarkApplicationDetailResponseBodyAdminUploads {
	s.LicensePicUrl = &v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyFirstClassification struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyFirstClassification) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyFirstClassification) SetName(v string) *QueryTradeMarkApplicationDetailResponseBodyFirstClassification {
	s.Name = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyFirstClassification) SetCode(v string) *QueryTradeMarkApplicationDetailResponseBodyFirstClassification {
	s.Code = &v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyMaterialDetail struct {
	Type                  *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	ReviewApplicationFile *string `json:"ReviewApplicationFile,omitempty" xml:"ReviewApplicationFile,omitempty"`
	Status                *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	BusinessLicenceUrl    *string `json:"BusinessLicenceUrl,omitempty" xml:"BusinessLicenceUrl,omitempty"`
	PassportUrl           *string `json:"PassportUrl,omitempty" xml:"PassportUrl,omitempty"`
	City                  *string `json:"City,omitempty" xml:"City,omitempty"`
	LegalNoticeUrl        *string `json:"LegalNoticeUrl,omitempty" xml:"LegalNoticeUrl,omitempty"`
	EAddress              *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	ContactEmail          *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	Region                *int32  `json:"Region,omitempty" xml:"Region,omitempty"`
	LoaUrl                *string `json:"LoaUrl,omitempty" xml:"LoaUrl,omitempty"`
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	PrincipalName         *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ContactNumber         *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactAddress        *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactZipcode        *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	ContactName           *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	EName                 *string `json:"EName,omitempty" xml:"EName,omitempty"`
	CardNumber            *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ExpirationDate        *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	IdCardUrl             *string `json:"IdCardUrl,omitempty" xml:"IdCardUrl,omitempty"`
	Country               *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Town                  *string `json:"Town,omitempty" xml:"Town,omitempty"`
	Province              *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 详细收件地址
	DetailedContactAddress *string                                                                         `json:"DetailedContactAddress,omitempty" xml:"DetailedContactAddress,omitempty"`
	ReviewAdditionalFiles  *QueryTradeMarkApplicationDetailResponseBodyMaterialDetailReviewAdditionalFiles `json:"ReviewAdditionalFiles,omitempty" xml:"ReviewAdditionalFiles,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetType(v int32) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.Type = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetReviewApplicationFile(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.ReviewApplicationFile = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetStatus(v int32) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.Status = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetBusinessLicenceUrl(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.BusinessLicenceUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetPassportUrl(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.PassportUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetCity(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.City = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetLegalNoticeUrl(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.LegalNoticeUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetEAddress(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.EAddress = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetContactEmail(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.ContactEmail = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetRegion(v int32) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.Region = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetLoaUrl(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.LoaUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetAddress(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.Address = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetPrincipalName(v int32) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.PrincipalName = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetName(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.Name = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetContactNumber(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.ContactNumber = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetContactAddress(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.ContactAddress = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetContactZipcode(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.ContactZipcode = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetContactName(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.ContactName = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetEName(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.EName = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetCardNumber(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.CardNumber = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetExpirationDate(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.ExpirationDate = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetIdCardUrl(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.IdCardUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetCountry(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.Country = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetTown(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.Town = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetProvince(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.Province = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetDetailedContactAddress(v string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.DetailedContactAddress = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail) SetReviewAdditionalFiles(v *QueryTradeMarkApplicationDetailResponseBodyMaterialDetailReviewAdditionalFiles) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetail {
	s.ReviewAdditionalFiles = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyMaterialDetailReviewAdditionalFiles struct {
	ReviewAdditionalFile []*string `json:"ReviewAdditionalFile,omitempty" xml:"ReviewAdditionalFile,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyMaterialDetailReviewAdditionalFiles) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyMaterialDetailReviewAdditionalFiles) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyMaterialDetailReviewAdditionalFiles) SetReviewAdditionalFile(v []*string) *QueryTradeMarkApplicationDetailResponseBodyMaterialDetailReviewAdditionalFiles {
	s.ReviewAdditionalFile = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyRenewResponse struct {
	EngName       *string `json:"EngName,omitempty" xml:"EngName,omitempty"`
	RegisterTime  *int64  `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	EngAddress    *string `json:"EngAddress,omitempty" xml:"EngAddress,omitempty"`
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SubmitSbjtime *int64  `json:"SubmitSbjtime,omitempty" xml:"SubmitSbjtime,omitempty"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyRenewResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyRenewResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyRenewResponse) SetEngName(v string) *QueryTradeMarkApplicationDetailResponseBodyRenewResponse {
	s.EngName = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyRenewResponse) SetRegisterTime(v int64) *QueryTradeMarkApplicationDetailResponseBodyRenewResponse {
	s.RegisterTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyRenewResponse) SetEngAddress(v string) *QueryTradeMarkApplicationDetailResponseBodyRenewResponse {
	s.EngAddress = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyRenewResponse) SetAddress(v string) *QueryTradeMarkApplicationDetailResponseBodyRenewResponse {
	s.Address = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyRenewResponse) SetName(v string) *QueryTradeMarkApplicationDetailResponseBodyRenewResponse {
	s.Name = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyRenewResponse) SetSubmitSbjtime(v int64) *QueryTradeMarkApplicationDetailResponseBodyRenewResponse {
	s.SubmitSbjtime = &v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles struct {
	ReviewKeep        *string                                                                          `json:"ReviewKeep,omitempty" xml:"ReviewKeep,omitempty"`
	ReviewAudit       *string                                                                          `json:"ReviewAudit,omitempty" xml:"ReviewAudit,omitempty"`
	ReviewPart        *string                                                                          `json:"ReviewPart,omitempty" xml:"ReviewPart,omitempty"`
	ReviewPass        *string                                                                          `json:"ReviewPass,omitempty" xml:"ReviewPass,omitempty"`
	ReviewSupplements *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFilesReviewSupplements `json:"ReviewSupplements,omitempty" xml:"ReviewSupplements,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles) SetReviewKeep(v string) *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles {
	s.ReviewKeep = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles) SetReviewAudit(v string) *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles {
	s.ReviewAudit = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles) SetReviewPart(v string) *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles {
	s.ReviewPart = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles) SetReviewPass(v string) *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles {
	s.ReviewPass = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles) SetReviewSupplements(v *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFilesReviewSupplements) *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFiles {
	s.ReviewSupplements = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFilesReviewSupplements struct {
	ReviewSupplement []*string `json:"ReviewSupplement,omitempty" xml:"ReviewSupplement,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFilesReviewSupplements) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFilesReviewSupplements) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFilesReviewSupplements) SetReviewSupplement(v []*string) *QueryTradeMarkApplicationDetailResponseBodyReviewOfficialFilesReviewSupplements {
	s.ReviewSupplement = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodySupplements struct {
	Supplements []*QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements `json:"Supplements,omitempty" xml:"Supplements,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationDetailResponseBodySupplements) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodySupplements) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplements) SetSupplements(v []*QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) *QueryTradeMarkApplicationDetailResponseBodySupplements {
	s.Supplements = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements struct {
	Type                  *int32                                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	OperateTime           *int64                                                                             `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	SerialNumber          *string                                                                            `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status                *int32                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	SbjDeadTime           *int64                                                                             `json:"SbjDeadTime,omitempty" xml:"SbjDeadTime,omitempty"`
	AcceptDeadTime        *int64                                                                             `json:"AcceptDeadTime,omitempty" xml:"AcceptDeadTime,omitempty"`
	SendTime              *int64                                                                             `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	BatchNum              *string                                                                            `json:"BatchNum,omitempty" xml:"BatchNum,omitempty"`
	AcceptTime            *int64                                                                             `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	TmNumber              *string                                                                            `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	UploadFileTemplateUrl *string                                                                            `json:"UploadFileTemplateUrl,omitempty" xml:"UploadFileTemplateUrl,omitempty"`
	Content               *string                                                                            `json:"Content,omitempty" xml:"Content,omitempty"`
	Id                    *int64                                                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	OrderId               *string                                                                            `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	FileTemplateUrls      *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplementsFileTemplateUrls `json:"FileTemplateUrls,omitempty" xml:"FileTemplateUrls,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetType(v int32) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.Type = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetOperateTime(v int64) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.OperateTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetSerialNumber(v string) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.SerialNumber = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetStatus(v int32) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.Status = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetSbjDeadTime(v int64) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.SbjDeadTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetAcceptDeadTime(v int64) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.AcceptDeadTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetSendTime(v int64) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.SendTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetBatchNum(v string) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.BatchNum = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetAcceptTime(v int64) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.AcceptTime = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetTmNumber(v string) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.TmNumber = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetUploadFileTemplateUrl(v string) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.UploadFileTemplateUrl = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetContent(v string) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.Content = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetId(v int64) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.Id = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetOrderId(v string) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.OrderId = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements) SetFileTemplateUrls(v *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplementsFileTemplateUrls) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplements {
	s.FileTemplateUrls = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodySupplementsSupplementsFileTemplateUrls struct {
	FileTemplateUrls []*string `json:"FileTemplateUrls,omitempty" xml:"FileTemplateUrls,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationDetailResponseBodySupplementsSupplementsFileTemplateUrls) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodySupplementsSupplementsFileTemplateUrls) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplementsFileTemplateUrls) SetFileTemplateUrls(v []*string) *QueryTradeMarkApplicationDetailResponseBodySupplementsSupplementsFileTemplateUrls {
	s.FileTemplateUrls = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyThirdClassification struct {
	ThirdClassifications []*QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications `json:"ThirdClassifications,omitempty" xml:"ThirdClassifications,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyThirdClassification) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyThirdClassification) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyThirdClassification) SetThirdClassifications(v []*QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications) *QueryTradeMarkApplicationDetailResponseBodyThirdClassification {
	s.ThirdClassifications = v
	return s
}

type QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications) SetName(v string) *QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications {
	s.Name = &v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications) SetCode(v string) *QueryTradeMarkApplicationDetailResponseBodyThirdClassificationThirdClassifications {
	s.Code = &v
	return s
}

type QueryTradeMarkApplicationDetailResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTradeMarkApplicationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTradeMarkApplicationDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationDetailResponse) SetHeaders(v map[string]*string) *QueryTradeMarkApplicationDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTradeMarkApplicationDetailResponse) SetBody(v *QueryTradeMarkApplicationDetailResponseBody) *QueryTradeMarkApplicationDetailResponse {
	s.Body = v
	return s
}

type SaveClassificationConditionsRequest struct {
	Type      *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
}

func (s SaveClassificationConditionsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveClassificationConditionsRequest) GoString() string {
	return s.String()
}

func (s *SaveClassificationConditionsRequest) SetType(v int32) *SaveClassificationConditionsRequest {
	s.Type = &v
	return s
}

func (s *SaveClassificationConditionsRequest) SetBizId(v string) *SaveClassificationConditionsRequest {
	s.BizId = &v
	return s
}

func (s *SaveClassificationConditionsRequest) SetCondition(v string) *SaveClassificationConditionsRequest {
	s.Condition = &v
	return s
}

type SaveClassificationConditionsResponseBody struct {
	ErrorMsg    *string                                                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId   *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagName     *string                                                `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Success     *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	InvalidList []*SaveClassificationConditionsResponseBodyInvalidList `json:"InvalidList,omitempty" xml:"InvalidList,omitempty" type:"Repeated"`
}

func (s SaveClassificationConditionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveClassificationConditionsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveClassificationConditionsResponseBody) SetErrorMsg(v string) *SaveClassificationConditionsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SaveClassificationConditionsResponseBody) SetRequestId(v string) *SaveClassificationConditionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveClassificationConditionsResponseBody) SetTagName(v string) *SaveClassificationConditionsResponseBody {
	s.TagName = &v
	return s
}

func (s *SaveClassificationConditionsResponseBody) SetSuccess(v bool) *SaveClassificationConditionsResponseBody {
	s.Success = &v
	return s
}

func (s *SaveClassificationConditionsResponseBody) SetInvalidList(v []*SaveClassificationConditionsResponseBodyInvalidList) *SaveClassificationConditionsResponseBody {
	s.InvalidList = v
	return s
}

type SaveClassificationConditionsResponseBodyInvalidList struct {
	ParentCode         *string `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
	OfficialCode       *string `json:"OfficialCode,omitempty" xml:"OfficialCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
}

func (s SaveClassificationConditionsResponseBodyInvalidList) String() string {
	return tea.Prettify(s)
}

func (s SaveClassificationConditionsResponseBodyInvalidList) GoString() string {
	return s.String()
}

func (s *SaveClassificationConditionsResponseBodyInvalidList) SetParentCode(v string) *SaveClassificationConditionsResponseBodyInvalidList {
	s.ParentCode = &v
	return s
}

func (s *SaveClassificationConditionsResponseBodyInvalidList) SetOfficialCode(v string) *SaveClassificationConditionsResponseBodyInvalidList {
	s.OfficialCode = &v
	return s
}

func (s *SaveClassificationConditionsResponseBodyInvalidList) SetClassificationName(v string) *SaveClassificationConditionsResponseBodyInvalidList {
	s.ClassificationName = &v
	return s
}

func (s *SaveClassificationConditionsResponseBodyInvalidList) SetClassificationCode(v string) *SaveClassificationConditionsResponseBodyInvalidList {
	s.ClassificationCode = &v
	return s
}

type SaveClassificationConditionsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveClassificationConditionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveClassificationConditionsResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveClassificationConditionsResponse) GoString() string {
	return s.String()
}

func (s *SaveClassificationConditionsResponse) SetHeaders(v map[string]*string) *SaveClassificationConditionsResponse {
	s.Headers = v
	return s
}

func (s *SaveClassificationConditionsResponse) SetBody(v *SaveClassificationConditionsResponseBody) *SaveClassificationConditionsResponse {
	s.Body = v
	return s
}

type FillLogisticsRequest struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Logistics *string `json:"Logistics,omitempty" xml:"Logistics,omitempty"`
}

func (s FillLogisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s FillLogisticsRequest) GoString() string {
	return s.String()
}

func (s *FillLogisticsRequest) SetBizId(v string) *FillLogisticsRequest {
	s.BizId = &v
	return s
}

func (s *FillLogisticsRequest) SetLogistics(v string) *FillLogisticsRequest {
	s.Logistics = &v
	return s
}

type FillLogisticsResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s FillLogisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FillLogisticsResponseBody) GoString() string {
	return s.String()
}

func (s *FillLogisticsResponseBody) SetErrorMsg(v string) *FillLogisticsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *FillLogisticsResponseBody) SetRequestId(v string) *FillLogisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *FillLogisticsResponseBody) SetSuccess(v bool) *FillLogisticsResponseBody {
	s.Success = &v
	return s
}

func (s *FillLogisticsResponseBody) SetErrorCode(v string) *FillLogisticsResponseBody {
	s.ErrorCode = &v
	return s
}

type FillLogisticsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FillLogisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FillLogisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s FillLogisticsResponse) GoString() string {
	return s.String()
}

func (s *FillLogisticsResponse) SetHeaders(v map[string]*string) *FillLogisticsResponse {
	s.Headers = v
	return s
}

func (s *FillLogisticsResponse) SetBody(v *FillLogisticsResponseBody) *FillLogisticsResponse {
	s.Body = v
	return s
}

type UpdateMaterialRequest struct {
	LoaId                 *int64  `json:"LoaId,omitempty" xml:"LoaId,omitempty"`
	ContactAddress        *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactName           *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber         *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactEmail          *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ContactZipcode        *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CardNumber            *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Province              *string `json:"Province,omitempty" xml:"Province,omitempty"`
	City                  *string `json:"City,omitempty" xml:"City,omitempty"`
	Town                  *string `json:"Town,omitempty" xml:"Town,omitempty"`
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EName                 *string `json:"EName,omitempty" xml:"EName,omitempty"`
	EAddress              *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	IdCardOssKey          *string `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	BusinessLicenceOssKey *string `json:"BusinessLicenceOssKey,omitempty" xml:"BusinessLicenceOssKey,omitempty"`
	PassportOssKey        *string `json:"PassportOssKey,omitempty" xml:"PassportOssKey,omitempty"`
	LoaOssKey             *string `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	LegalNoticeOssKey     *string `json:"LegalNoticeOssKey,omitempty" xml:"LegalNoticeOssKey,omitempty"`
	ContactProvince       *string `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	ContactCity           *string `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	ContactDistrict       *string `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	ContactCounty         *string `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
}

func (s UpdateMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaterialRequest) SetLoaId(v int64) *UpdateMaterialRequest {
	s.LoaId = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactAddress(v string) *UpdateMaterialRequest {
	s.ContactAddress = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactName(v string) *UpdateMaterialRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactNumber(v string) *UpdateMaterialRequest {
	s.ContactNumber = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactEmail(v string) *UpdateMaterialRequest {
	s.ContactEmail = &v
	return s
}

func (s *UpdateMaterialRequest) SetId(v int64) *UpdateMaterialRequest {
	s.Id = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactZipcode(v string) *UpdateMaterialRequest {
	s.ContactZipcode = &v
	return s
}

func (s *UpdateMaterialRequest) SetName(v string) *UpdateMaterialRequest {
	s.Name = &v
	return s
}

func (s *UpdateMaterialRequest) SetCardNumber(v string) *UpdateMaterialRequest {
	s.CardNumber = &v
	return s
}

func (s *UpdateMaterialRequest) SetProvince(v string) *UpdateMaterialRequest {
	s.Province = &v
	return s
}

func (s *UpdateMaterialRequest) SetCity(v string) *UpdateMaterialRequest {
	s.City = &v
	return s
}

func (s *UpdateMaterialRequest) SetTown(v string) *UpdateMaterialRequest {
	s.Town = &v
	return s
}

func (s *UpdateMaterialRequest) SetAddress(v string) *UpdateMaterialRequest {
	s.Address = &v
	return s
}

func (s *UpdateMaterialRequest) SetEName(v string) *UpdateMaterialRequest {
	s.EName = &v
	return s
}

func (s *UpdateMaterialRequest) SetEAddress(v string) *UpdateMaterialRequest {
	s.EAddress = &v
	return s
}

func (s *UpdateMaterialRequest) SetIdCardOssKey(v string) *UpdateMaterialRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *UpdateMaterialRequest) SetBusinessLicenceOssKey(v string) *UpdateMaterialRequest {
	s.BusinessLicenceOssKey = &v
	return s
}

func (s *UpdateMaterialRequest) SetPassportOssKey(v string) *UpdateMaterialRequest {
	s.PassportOssKey = &v
	return s
}

func (s *UpdateMaterialRequest) SetLoaOssKey(v string) *UpdateMaterialRequest {
	s.LoaOssKey = &v
	return s
}

func (s *UpdateMaterialRequest) SetLegalNoticeOssKey(v string) *UpdateMaterialRequest {
	s.LegalNoticeOssKey = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactProvince(v string) *UpdateMaterialRequest {
	s.ContactProvince = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactCity(v string) *UpdateMaterialRequest {
	s.ContactCity = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactDistrict(v string) *UpdateMaterialRequest {
	s.ContactDistrict = &v
	return s
}

func (s *UpdateMaterialRequest) SetContactCounty(v string) *UpdateMaterialRequest {
	s.ContactCounty = &v
	return s
}

type UpdateMaterialResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMaterialResponseBody) SetSuccess(v bool) *UpdateMaterialResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMaterialResponseBody) SetRequestId(v string) *UpdateMaterialResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMaterialResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialResponse) GoString() string {
	return s.String()
}

func (s *UpdateMaterialResponse) SetHeaders(v map[string]*string) *UpdateMaterialResponse {
	s.Headers = v
	return s
}

func (s *UpdateMaterialResponse) SetBody(v *UpdateMaterialResponseBody) *UpdateMaterialResponse {
	s.Body = v
	return s
}

type QueryTradeMarkApplicationLogsRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s QueryTradeMarkApplicationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationLogsRequest) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationLogsRequest) SetBizId(v string) *QueryTradeMarkApplicationLogsRequest {
	s.BizId = &v
	return s
}

type QueryTradeMarkApplicationLogsResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryTradeMarkApplicationLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTradeMarkApplicationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationLogsResponseBody) SetRequestId(v string) *QueryTradeMarkApplicationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTradeMarkApplicationLogsResponseBody) SetData(v *QueryTradeMarkApplicationLogsResponseBodyData) *QueryTradeMarkApplicationLogsResponseBody {
	s.Data = v
	return s
}

type QueryTradeMarkApplicationLogsResponseBodyData struct {
	Data []*QueryTradeMarkApplicationLogsResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QueryTradeMarkApplicationLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationLogsResponseBodyData) SetData(v []*QueryTradeMarkApplicationLogsResponseBodyDataData) *QueryTradeMarkApplicationLogsResponseBodyData {
	s.Data = v
	return s
}

type QueryTradeMarkApplicationLogsResponseBodyDataData struct {
	OperateTime   *int64  `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OperateType   *int32  `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	BizStatus     *int32  `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
}

func (s QueryTradeMarkApplicationLogsResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationLogsResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationLogsResponseBodyDataData) SetOperateTime(v int64) *QueryTradeMarkApplicationLogsResponseBodyDataData {
	s.OperateTime = &v
	return s
}

func (s *QueryTradeMarkApplicationLogsResponseBodyDataData) SetOperateType(v int32) *QueryTradeMarkApplicationLogsResponseBodyDataData {
	s.OperateType = &v
	return s
}

func (s *QueryTradeMarkApplicationLogsResponseBodyDataData) SetExtendContent(v string) *QueryTradeMarkApplicationLogsResponseBodyDataData {
	s.ExtendContent = &v
	return s
}

func (s *QueryTradeMarkApplicationLogsResponseBodyDataData) SetBizId(v string) *QueryTradeMarkApplicationLogsResponseBodyDataData {
	s.BizId = &v
	return s
}

func (s *QueryTradeMarkApplicationLogsResponseBodyDataData) SetNote(v string) *QueryTradeMarkApplicationLogsResponseBodyDataData {
	s.Note = &v
	return s
}

func (s *QueryTradeMarkApplicationLogsResponseBodyDataData) SetBizStatus(v int32) *QueryTradeMarkApplicationLogsResponseBodyDataData {
	s.BizStatus = &v
	return s
}

type QueryTradeMarkApplicationLogsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTradeMarkApplicationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTradeMarkApplicationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTradeMarkApplicationLogsResponse) GoString() string {
	return s.String()
}

func (s *QueryTradeMarkApplicationLogsResponse) SetHeaders(v map[string]*string) *QueryTradeMarkApplicationLogsResponse {
	s.Headers = v
	return s
}

func (s *QueryTradeMarkApplicationLogsResponse) SetBody(v *QueryTradeMarkApplicationLogsResponseBody) *QueryTradeMarkApplicationLogsResponse {
	s.Body = v
	return s
}

type RefundProduceRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s RefundProduceRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundProduceRequest) GoString() string {
	return s.String()
}

func (s *RefundProduceRequest) SetBizId(v string) *RefundProduceRequest {
	s.BizId = &v
	return s
}

type RefundProduceResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s RefundProduceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundProduceResponseBody) GoString() string {
	return s.String()
}

func (s *RefundProduceResponseBody) SetErrorMsg(v string) *RefundProduceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefundProduceResponseBody) SetRequestId(v string) *RefundProduceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundProduceResponseBody) SetSuccess(v bool) *RefundProduceResponseBody {
	s.Success = &v
	return s
}

func (s *RefundProduceResponseBody) SetErrorCode(v string) *RefundProduceResponseBody {
	s.ErrorCode = &v
	return s
}

type RefundProduceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefundProduceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefundProduceResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundProduceResponse) GoString() string {
	return s.String()
}

func (s *RefundProduceResponse) SetHeaders(v map[string]*string) *RefundProduceResponse {
	s.Headers = v
	return s
}

func (s *RefundProduceResponse) SetBody(v *RefundProduceResponseBody) *RefundProduceResponse {
	s.Body = v
	return s
}

type SyncTrademarkRequest struct {
	ClassificationCode      *string  `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	TmName                  *string  `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmIcon                  *string  `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	OriginalPrice           *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TmNumber                *string  `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	Status                  *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	EndTime                 *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	BeginTime               *int64   `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Description             *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Label                   *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	RegAnnDate              *int64   `json:"RegAnnDate,omitempty" xml:"RegAnnDate,omitempty"`
	OwnerName               *string  `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerEnName             *string  `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	SecondaryClassification *string  `json:"SecondaryClassification,omitempty" xml:"SecondaryClassification,omitempty"`
	ThirdClassification     *string  `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty"`
	Type                    *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Reason                  *string  `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s SyncTrademarkRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncTrademarkRequest) GoString() string {
	return s.String()
}

func (s *SyncTrademarkRequest) SetClassificationCode(v string) *SyncTrademarkRequest {
	s.ClassificationCode = &v
	return s
}

func (s *SyncTrademarkRequest) SetTmName(v string) *SyncTrademarkRequest {
	s.TmName = &v
	return s
}

func (s *SyncTrademarkRequest) SetTmIcon(v string) *SyncTrademarkRequest {
	s.TmIcon = &v
	return s
}

func (s *SyncTrademarkRequest) SetOriginalPrice(v float32) *SyncTrademarkRequest {
	s.OriginalPrice = &v
	return s
}

func (s *SyncTrademarkRequest) SetTmNumber(v string) *SyncTrademarkRequest {
	s.TmNumber = &v
	return s
}

func (s *SyncTrademarkRequest) SetStatus(v string) *SyncTrademarkRequest {
	s.Status = &v
	return s
}

func (s *SyncTrademarkRequest) SetEndTime(v int64) *SyncTrademarkRequest {
	s.EndTime = &v
	return s
}

func (s *SyncTrademarkRequest) SetBeginTime(v int64) *SyncTrademarkRequest {
	s.BeginTime = &v
	return s
}

func (s *SyncTrademarkRequest) SetDescription(v string) *SyncTrademarkRequest {
	s.Description = &v
	return s
}

func (s *SyncTrademarkRequest) SetLabel(v string) *SyncTrademarkRequest {
	s.Label = &v
	return s
}

func (s *SyncTrademarkRequest) SetRegAnnDate(v int64) *SyncTrademarkRequest {
	s.RegAnnDate = &v
	return s
}

func (s *SyncTrademarkRequest) SetOwnerName(v string) *SyncTrademarkRequest {
	s.OwnerName = &v
	return s
}

func (s *SyncTrademarkRequest) SetOwnerEnName(v string) *SyncTrademarkRequest {
	s.OwnerEnName = &v
	return s
}

func (s *SyncTrademarkRequest) SetSecondaryClassification(v string) *SyncTrademarkRequest {
	s.SecondaryClassification = &v
	return s
}

func (s *SyncTrademarkRequest) SetThirdClassification(v string) *SyncTrademarkRequest {
	s.ThirdClassification = &v
	return s
}

func (s *SyncTrademarkRequest) SetType(v string) *SyncTrademarkRequest {
	s.Type = &v
	return s
}

func (s *SyncTrademarkRequest) SetReason(v string) *SyncTrademarkRequest {
	s.Reason = &v
	return s
}

type SyncTrademarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncTrademarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncTrademarkResponseBody) GoString() string {
	return s.String()
}

func (s *SyncTrademarkResponseBody) SetRequestId(v string) *SyncTrademarkResponseBody {
	s.RequestId = &v
	return s
}

type SyncTrademarkResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncTrademarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncTrademarkResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncTrademarkResponse) GoString() string {
	return s.String()
}

func (s *SyncTrademarkResponse) SetHeaders(v map[string]*string) *SyncTrademarkResponse {
	s.Headers = v
	return s
}

func (s *SyncTrademarkResponse) SetBody(v *SyncTrademarkResponseBody) *SyncTrademarkResponse {
	s.Body = v
	return s
}

type CombineLoaRequest struct {
	MaterialId    *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	TrademarkName *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	MaterialName  *string `json:"MaterialName,omitempty" xml:"MaterialName,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	TmProduceType *string `json:"TmProduceType,omitempty" xml:"TmProduceType,omitempty"`
	PrincipalName *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
}

func (s CombineLoaRequest) String() string {
	return tea.Prettify(s)
}

func (s CombineLoaRequest) GoString() string {
	return s.String()
}

func (s *CombineLoaRequest) SetMaterialId(v string) *CombineLoaRequest {
	s.MaterialId = &v
	return s
}

func (s *CombineLoaRequest) SetTrademarkName(v string) *CombineLoaRequest {
	s.TrademarkName = &v
	return s
}

func (s *CombineLoaRequest) SetMaterialName(v string) *CombineLoaRequest {
	s.MaterialName = &v
	return s
}

func (s *CombineLoaRequest) SetNationality(v string) *CombineLoaRequest {
	s.Nationality = &v
	return s
}

func (s *CombineLoaRequest) SetAddress(v string) *CombineLoaRequest {
	s.Address = &v
	return s
}

func (s *CombineLoaRequest) SetTmProduceType(v string) *CombineLoaRequest {
	s.TmProduceType = &v
	return s
}

func (s *CombineLoaRequest) SetPrincipalName(v int32) *CombineLoaRequest {
	s.PrincipalName = &v
	return s
}

type CombineLoaResponseBody struct {
	TemplateCombineUrl *string `json:"TemplateCombineUrl,omitempty" xml:"TemplateCombineUrl,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CombineLoaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CombineLoaResponseBody) GoString() string {
	return s.String()
}

func (s *CombineLoaResponseBody) SetTemplateCombineUrl(v string) *CombineLoaResponseBody {
	s.TemplateCombineUrl = &v
	return s
}

func (s *CombineLoaResponseBody) SetRequestId(v string) *CombineLoaResponseBody {
	s.RequestId = &v
	return s
}

type CombineLoaResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CombineLoaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CombineLoaResponse) String() string {
	return tea.Prettify(s)
}

func (s CombineLoaResponse) GoString() string {
	return s.String()
}

func (s *CombineLoaResponse) SetHeaders(v map[string]*string) *CombineLoaResponse {
	s.Headers = v
	return s
}

func (s *CombineLoaResponse) SetBody(v *CombineLoaResponseBody) *CombineLoaResponse {
	s.Body = v
	return s
}

type FilterUnavailableCodesRequest struct {
	Codes map[string]interface{} `json:"Codes,omitempty" xml:"Codes,omitempty"`
}

func (s FilterUnavailableCodesRequest) String() string {
	return tea.Prettify(s)
}

func (s FilterUnavailableCodesRequest) GoString() string {
	return s.String()
}

func (s *FilterUnavailableCodesRequest) SetCodes(v map[string]interface{}) *FilterUnavailableCodesRequest {
	s.Codes = v
	return s
}

type FilterUnavailableCodesShrinkRequest struct {
	CodesShrink *string `json:"Codes,omitempty" xml:"Codes,omitempty"`
}

func (s FilterUnavailableCodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FilterUnavailableCodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *FilterUnavailableCodesShrinkRequest) SetCodesShrink(v string) *FilterUnavailableCodesShrinkRequest {
	s.CodesShrink = &v
	return s
}

type FilterUnavailableCodesResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *FilterUnavailableCodesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s FilterUnavailableCodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FilterUnavailableCodesResponseBody) GoString() string {
	return s.String()
}

func (s *FilterUnavailableCodesResponseBody) SetRequestId(v string) *FilterUnavailableCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FilterUnavailableCodesResponseBody) SetData(v *FilterUnavailableCodesResponseBodyData) *FilterUnavailableCodesResponseBody {
	s.Data = v
	return s
}

type FilterUnavailableCodesResponseBodyData struct {
	Codes []*string `json:"Codes,omitempty" xml:"Codes,omitempty" type:"Repeated"`
}

func (s FilterUnavailableCodesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FilterUnavailableCodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *FilterUnavailableCodesResponseBodyData) SetCodes(v []*string) *FilterUnavailableCodesResponseBodyData {
	s.Codes = v
	return s
}

type FilterUnavailableCodesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FilterUnavailableCodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FilterUnavailableCodesResponse) String() string {
	return tea.Prettify(s)
}

func (s FilterUnavailableCodesResponse) GoString() string {
	return s.String()
}

func (s *FilterUnavailableCodesResponse) SetHeaders(v map[string]*string) *FilterUnavailableCodesResponse {
	s.Headers = v
	return s
}

func (s *FilterUnavailableCodesResponse) SetBody(v *FilterUnavailableCodesResponseBody) *FilterUnavailableCodesResponse {
	s.Body = v
	return s
}

type InsertMaterialRequest struct {
	ContactZipcode        *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	Type                  *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Region                *int32  `json:"Region,omitempty" xml:"Region,omitempty"`
	ContactName           *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber         *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactEmail          *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactAddress        *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	LoaOssKey             *string `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CardNumber            *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Province              *string `json:"Province,omitempty" xml:"Province,omitempty"`
	City                  *string `json:"City,omitempty" xml:"City,omitempty"`
	Town                  *string `json:"Town,omitempty" xml:"Town,omitempty"`
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EName                 *string `json:"EName,omitempty" xml:"EName,omitempty"`
	EAddress              *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	Country               *string `json:"Country,omitempty" xml:"Country,omitempty"`
	IdCardOssKey          *string `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	BusinessLicenceOssKey *string `json:"BusinessLicenceOssKey,omitempty" xml:"BusinessLicenceOssKey,omitempty"`
	PassportOssKey        *string `json:"PassportOssKey,omitempty" xml:"PassportOssKey,omitempty"`
	LegalNoticeOssKey     *string `json:"LegalNoticeOssKey,omitempty" xml:"LegalNoticeOssKey,omitempty"`
	PrincipalName         *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	ContactProvince       *string `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	ContactCity           *string `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	ContactDistrict       *string `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	ContactCounty         *string `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
}

func (s InsertMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertMaterialRequest) GoString() string {
	return s.String()
}

func (s *InsertMaterialRequest) SetContactZipcode(v string) *InsertMaterialRequest {
	s.ContactZipcode = &v
	return s
}

func (s *InsertMaterialRequest) SetType(v int32) *InsertMaterialRequest {
	s.Type = &v
	return s
}

func (s *InsertMaterialRequest) SetRegion(v int32) *InsertMaterialRequest {
	s.Region = &v
	return s
}

func (s *InsertMaterialRequest) SetContactName(v string) *InsertMaterialRequest {
	s.ContactName = &v
	return s
}

func (s *InsertMaterialRequest) SetContactNumber(v string) *InsertMaterialRequest {
	s.ContactNumber = &v
	return s
}

func (s *InsertMaterialRequest) SetContactEmail(v string) *InsertMaterialRequest {
	s.ContactEmail = &v
	return s
}

func (s *InsertMaterialRequest) SetContactAddress(v string) *InsertMaterialRequest {
	s.ContactAddress = &v
	return s
}

func (s *InsertMaterialRequest) SetLoaOssKey(v string) *InsertMaterialRequest {
	s.LoaOssKey = &v
	return s
}

func (s *InsertMaterialRequest) SetName(v string) *InsertMaterialRequest {
	s.Name = &v
	return s
}

func (s *InsertMaterialRequest) SetCardNumber(v string) *InsertMaterialRequest {
	s.CardNumber = &v
	return s
}

func (s *InsertMaterialRequest) SetProvince(v string) *InsertMaterialRequest {
	s.Province = &v
	return s
}

func (s *InsertMaterialRequest) SetCity(v string) *InsertMaterialRequest {
	s.City = &v
	return s
}

func (s *InsertMaterialRequest) SetTown(v string) *InsertMaterialRequest {
	s.Town = &v
	return s
}

func (s *InsertMaterialRequest) SetAddress(v string) *InsertMaterialRequest {
	s.Address = &v
	return s
}

func (s *InsertMaterialRequest) SetEName(v string) *InsertMaterialRequest {
	s.EName = &v
	return s
}

func (s *InsertMaterialRequest) SetEAddress(v string) *InsertMaterialRequest {
	s.EAddress = &v
	return s
}

func (s *InsertMaterialRequest) SetCountry(v string) *InsertMaterialRequest {
	s.Country = &v
	return s
}

func (s *InsertMaterialRequest) SetIdCardOssKey(v string) *InsertMaterialRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *InsertMaterialRequest) SetBusinessLicenceOssKey(v string) *InsertMaterialRequest {
	s.BusinessLicenceOssKey = &v
	return s
}

func (s *InsertMaterialRequest) SetPassportOssKey(v string) *InsertMaterialRequest {
	s.PassportOssKey = &v
	return s
}

func (s *InsertMaterialRequest) SetLegalNoticeOssKey(v string) *InsertMaterialRequest {
	s.LegalNoticeOssKey = &v
	return s
}

func (s *InsertMaterialRequest) SetPrincipalName(v int32) *InsertMaterialRequest {
	s.PrincipalName = &v
	return s
}

func (s *InsertMaterialRequest) SetContactProvince(v string) *InsertMaterialRequest {
	s.ContactProvince = &v
	return s
}

func (s *InsertMaterialRequest) SetContactCity(v string) *InsertMaterialRequest {
	s.ContactCity = &v
	return s
}

func (s *InsertMaterialRequest) SetContactDistrict(v string) *InsertMaterialRequest {
	s.ContactDistrict = &v
	return s
}

func (s *InsertMaterialRequest) SetContactCounty(v string) *InsertMaterialRequest {
	s.ContactCounty = &v
	return s
}

type InsertMaterialResponseBody struct {
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *InsertMaterialResponseBody) SetSuccess(v bool) *InsertMaterialResponseBody {
	s.Success = &v
	return s
}

func (s *InsertMaterialResponseBody) SetRequestId(v string) *InsertMaterialResponseBody {
	s.RequestId = &v
	return s
}

type InsertMaterialResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertMaterialResponse) GoString() string {
	return s.String()
}

func (s *InsertMaterialResponse) SetHeaders(v map[string]*string) *InsertMaterialResponse {
	s.Headers = v
	return s
}

func (s *InsertMaterialResponse) SetBody(v *InsertMaterialResponseBody) *InsertMaterialResponse {
	s.Body = v
	return s
}

type SaveTradeMarkReviewMaterialDetailRequest struct {
	BizId                 *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Type                  *int32                 `json:"Type,omitempty" xml:"Type,omitempty"`
	Region                *int32                 `json:"Region,omitempty" xml:"Region,omitempty"`
	Country               *string                `json:"Country,omitempty" xml:"Country,omitempty"`
	ContactName           *string                `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber         *string                `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactEmail          *string                `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactAddress        *string                `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	LoaOssKey             *string                `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	Name                  *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	CardNumber            *string                `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Province              *string                `json:"Province,omitempty" xml:"Province,omitempty"`
	Address               *string                `json:"Address,omitempty" xml:"Address,omitempty"`
	EngName               *string                `json:"EngName,omitempty" xml:"EngName,omitempty"`
	EngAddress            *string                `json:"EngAddress,omitempty" xml:"EngAddress,omitempty"`
	IdCardOssKey          *string                `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	BusinessLicenceOssKey *string                `json:"BusinessLicenceOssKey,omitempty" xml:"BusinessLicenceOssKey,omitempty"`
	PassportOssKey        *string                `json:"PassportOssKey,omitempty" xml:"PassportOssKey,omitempty"`
	LegalNoticeOssKey     *string                `json:"LegalNoticeOssKey,omitempty" xml:"LegalNoticeOssKey,omitempty"`
	ApplicationOssKey     *string                `json:"ApplicationOssKey,omitempty" xml:"ApplicationOssKey,omitempty"`
	AdditionalOssKeyList  map[string]interface{} `json:"AdditionalOssKeyList,omitempty" xml:"AdditionalOssKeyList,omitempty"`
	SubmitType            *int32                 `json:"SubmitType,omitempty" xml:"SubmitType,omitempty"`
}

func (s SaveTradeMarkReviewMaterialDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTradeMarkReviewMaterialDetailRequest) GoString() string {
	return s.String()
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetBizId(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.BizId = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetType(v int32) *SaveTradeMarkReviewMaterialDetailRequest {
	s.Type = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetRegion(v int32) *SaveTradeMarkReviewMaterialDetailRequest {
	s.Region = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetCountry(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.Country = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetContactName(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.ContactName = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetContactNumber(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.ContactNumber = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetContactEmail(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.ContactEmail = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetContactAddress(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.ContactAddress = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetLoaOssKey(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.LoaOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetName(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.Name = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetCardNumber(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.CardNumber = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetProvince(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.Province = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetAddress(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.Address = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetEngName(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.EngName = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetEngAddress(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.EngAddress = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetIdCardOssKey(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetBusinessLicenceOssKey(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.BusinessLicenceOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetPassportOssKey(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.PassportOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetLegalNoticeOssKey(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.LegalNoticeOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetApplicationOssKey(v string) *SaveTradeMarkReviewMaterialDetailRequest {
	s.ApplicationOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetAdditionalOssKeyList(v map[string]interface{}) *SaveTradeMarkReviewMaterialDetailRequest {
	s.AdditionalOssKeyList = v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailRequest) SetSubmitType(v int32) *SaveTradeMarkReviewMaterialDetailRequest {
	s.SubmitType = &v
	return s
}

type SaveTradeMarkReviewMaterialDetailShrinkRequest struct {
	BizId                      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Type                       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Region                     *int32  `json:"Region,omitempty" xml:"Region,omitempty"`
	Country                    *string `json:"Country,omitempty" xml:"Country,omitempty"`
	ContactName                *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber              *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactEmail               *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactAddress             *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	LoaOssKey                  *string `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	Name                       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CardNumber                 *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Province                   *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Address                    *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EngName                    *string `json:"EngName,omitempty" xml:"EngName,omitempty"`
	EngAddress                 *string `json:"EngAddress,omitempty" xml:"EngAddress,omitempty"`
	IdCardOssKey               *string `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	BusinessLicenceOssKey      *string `json:"BusinessLicenceOssKey,omitempty" xml:"BusinessLicenceOssKey,omitempty"`
	PassportOssKey             *string `json:"PassportOssKey,omitempty" xml:"PassportOssKey,omitempty"`
	LegalNoticeOssKey          *string `json:"LegalNoticeOssKey,omitempty" xml:"LegalNoticeOssKey,omitempty"`
	ApplicationOssKey          *string `json:"ApplicationOssKey,omitempty" xml:"ApplicationOssKey,omitempty"`
	AdditionalOssKeyListShrink *string `json:"AdditionalOssKeyList,omitempty" xml:"AdditionalOssKeyList,omitempty"`
	SubmitType                 *int32  `json:"SubmitType,omitempty" xml:"SubmitType,omitempty"`
}

func (s SaveTradeMarkReviewMaterialDetailShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTradeMarkReviewMaterialDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetBizId(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.BizId = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetType(v int32) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.Type = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetRegion(v int32) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.Region = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetCountry(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.Country = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetContactName(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.ContactName = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetContactNumber(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.ContactNumber = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetContactEmail(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.ContactEmail = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetContactAddress(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.ContactAddress = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetLoaOssKey(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.LoaOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetName(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.Name = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetCardNumber(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.CardNumber = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetProvince(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.Province = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetAddress(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.Address = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetEngName(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.EngName = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetEngAddress(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.EngAddress = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetIdCardOssKey(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetBusinessLicenceOssKey(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.BusinessLicenceOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetPassportOssKey(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.PassportOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetLegalNoticeOssKey(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.LegalNoticeOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetApplicationOssKey(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.ApplicationOssKey = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetAdditionalOssKeyListShrink(v string) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.AdditionalOssKeyListShrink = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailShrinkRequest) SetSubmitType(v int32) *SaveTradeMarkReviewMaterialDetailShrinkRequest {
	s.SubmitType = &v
	return s
}

type SaveTradeMarkReviewMaterialDetailResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveTradeMarkReviewMaterialDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTradeMarkReviewMaterialDetailResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTradeMarkReviewMaterialDetailResponseBody) SetErrorMsg(v string) *SaveTradeMarkReviewMaterialDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailResponseBody) SetRequestId(v string) *SaveTradeMarkReviewMaterialDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailResponseBody) SetSuccess(v bool) *SaveTradeMarkReviewMaterialDetailResponseBody {
	s.Success = &v
	return s
}

type SaveTradeMarkReviewMaterialDetailResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveTradeMarkReviewMaterialDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTradeMarkReviewMaterialDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTradeMarkReviewMaterialDetailResponse) GoString() string {
	return s.String()
}

func (s *SaveTradeMarkReviewMaterialDetailResponse) SetHeaders(v map[string]*string) *SaveTradeMarkReviewMaterialDetailResponse {
	s.Headers = v
	return s
}

func (s *SaveTradeMarkReviewMaterialDetailResponse) SetBody(v *SaveTradeMarkReviewMaterialDetailResponseBody) *SaveTradeMarkReviewMaterialDetailResponse {
	s.Body = v
	return s
}

type QueryMonitorKeywordsRequest struct {
	RuleType *int32    `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
}

func (s QueryMonitorKeywordsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorKeywordsRequest) GoString() string {
	return s.String()
}

func (s *QueryMonitorKeywordsRequest) SetRuleType(v int32) *QueryMonitorKeywordsRequest {
	s.RuleType = &v
	return s
}

func (s *QueryMonitorKeywordsRequest) SetKeywords(v []*string) *QueryMonitorKeywordsRequest {
	s.Keywords = v
	return s
}

type QueryMonitorKeywordsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryMonitorKeywordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryMonitorKeywordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonitorKeywordsResponseBody) SetRequestId(v string) *QueryMonitorKeywordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonitorKeywordsResponseBody) SetData(v *QueryMonitorKeywordsResponseBodyData) *QueryMonitorKeywordsResponseBody {
	s.Data = v
	return s
}

type QueryMonitorKeywordsResponseBodyData struct {
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
}

func (s QueryMonitorKeywordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorKeywordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMonitorKeywordsResponseBodyData) SetKeywords(v []*string) *QueryMonitorKeywordsResponseBodyData {
	s.Keywords = v
	return s
}

type QueryMonitorKeywordsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMonitorKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMonitorKeywordsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorKeywordsResponse) GoString() string {
	return s.String()
}

func (s *QueryMonitorKeywordsResponse) SetHeaders(v map[string]*string) *QueryMonitorKeywordsResponse {
	s.Headers = v
	return s
}

func (s *QueryMonitorKeywordsResponse) SetBody(v *QueryMonitorKeywordsResponseBody) *QueryMonitorKeywordsResponse {
	s.Body = v
	return s
}

type QueryTaskListRequest struct {
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
}

func (s QueryTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskListRequest) SetBizType(v string) *QueryTaskListRequest {
	s.BizType = &v
	return s
}

func (s *QueryTaskListRequest) SetPageSize(v int32) *QueryTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskListRequest) SetPageNum(v int32) *QueryTaskListRequest {
	s.PageNum = &v
	return s
}

type QueryTaskListResponseBody struct {
	CurrentPageNum *int32                         `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	TotalPageNum   *int32                         `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                         `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	Data           *QueryTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBody) SetCurrentPageNum(v int32) *QueryTaskListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTaskListResponseBody) SetTotalPageNum(v int32) *QueryTaskListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTaskListResponseBody) SetRequestId(v string) *QueryTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskListResponseBody) SetPageSize(v int32) *QueryTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskListResponseBody) SetTotalItemNum(v int32) *QueryTaskListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTaskListResponseBody) SetData(v *QueryTaskListResponseBodyData) *QueryTaskListResponseBody {
	s.Data = v
	return s
}

type QueryTaskListResponseBodyData struct {
	TaskList []*QueryTaskListResponseBodyDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s QueryTaskListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBodyData) SetTaskList(v []*QueryTaskListResponseBodyDataTaskList) *QueryTaskListResponseBodyData {
	s.TaskList = v
	return s
}

type QueryTaskListResponseBodyDataTaskList struct {
	TaskType     *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	CompleteTime *int64  `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErrMsg       *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s QueryTaskListResponseBodyDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListResponseBodyDataTaskList) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBodyDataTaskList) SetTaskType(v string) *QueryTaskListResponseBodyDataTaskList {
	s.TaskType = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskList) SetResult(v string) *QueryTaskListResponseBodyDataTaskList {
	s.Result = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskList) SetTaskStatus(v string) *QueryTaskListResponseBodyDataTaskList {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskList) SetCompleteTime(v int64) *QueryTaskListResponseBodyDataTaskList {
	s.CompleteTime = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskList) SetCreateTime(v int64) *QueryTaskListResponseBodyDataTaskList {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskList) SetErrMsg(v string) *QueryTaskListResponseBodyDataTaskList {
	s.ErrMsg = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskList) SetFileName(v string) *QueryTaskListResponseBodyDataTaskList {
	s.FileName = &v
	return s
}

type QueryTaskListResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponse) SetHeaders(v map[string]*string) *QueryTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskListResponse) SetBody(v *QueryTaskListResponseBody) *QueryTaskListResponse {
	s.Body = v
	return s
}

type UpdateTrademarkNameRequest struct {
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 业务id
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// 商标名称
	TmName *string `json:"TmName,omitempty" xml:"TmName,omitempty"`
	// 商标图片
	TmIcon    *string `json:"TmIcon,omitempty" xml:"TmIcon,omitempty"`
	TmComment *string `json:"TmComment,omitempty" xml:"TmComment,omitempty"`
	// 商标类型
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTrademarkNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrademarkNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrademarkNameRequest) SetClientToken(v string) *UpdateTrademarkNameRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTrademarkNameRequest) SetBizId(v string) *UpdateTrademarkNameRequest {
	s.BizId = &v
	return s
}

func (s *UpdateTrademarkNameRequest) SetTmName(v string) *UpdateTrademarkNameRequest {
	s.TmName = &v
	return s
}

func (s *UpdateTrademarkNameRequest) SetTmIcon(v string) *UpdateTrademarkNameRequest {
	s.TmIcon = &v
	return s
}

func (s *UpdateTrademarkNameRequest) SetTmComment(v string) *UpdateTrademarkNameRequest {
	s.TmComment = &v
	return s
}

func (s *UpdateTrademarkNameRequest) SetType(v int64) *UpdateTrademarkNameRequest {
	s.Type = &v
	return s
}

type UpdateTrademarkNameResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTrademarkNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrademarkNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrademarkNameResponseBody) SetRequestId(v string) *UpdateTrademarkNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrademarkNameResponseBody) SetErrorCode(v string) *UpdateTrademarkNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTrademarkNameResponseBody) SetErrorMsg(v string) *UpdateTrademarkNameResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateTrademarkNameResponseBody) SetSuccess(v bool) *UpdateTrademarkNameResponseBody {
	s.Success = &v
	return s
}

type UpdateTrademarkNameResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTrademarkNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTrademarkNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrademarkNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrademarkNameResponse) SetHeaders(v map[string]*string) *UpdateTrademarkNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrademarkNameResponse) SetBody(v *UpdateTrademarkNameResponseBody) *UpdateTrademarkNameResponse {
	s.Body = v
	return s
}

type CheckLoaFillRequest struct {
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckLoaFillRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckLoaFillRequest) GoString() string {
	return s.String()
}

func (s *CheckLoaFillRequest) SetOssKey(v string) *CheckLoaFillRequest {
	s.OssKey = &v
	return s
}

func (s *CheckLoaFillRequest) SetType(v string) *CheckLoaFillRequest {
	s.Type = &v
	return s
}

type CheckLoaFillResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CheckLoaFillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CheckLoaFillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckLoaFillResponseBody) GoString() string {
	return s.String()
}

func (s *CheckLoaFillResponseBody) SetRequestId(v string) *CheckLoaFillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckLoaFillResponseBody) SetData(v *CheckLoaFillResponseBodyData) *CheckLoaFillResponseBody {
	s.Data = v
	return s
}

type CheckLoaFillResponseBodyData struct {
	AddressFill       *bool                                  `json:"AddressFill,omitempty" xml:"AddressFill,omitempty"`
	TemplateUrl       *string                                `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
	CountryFill       *bool                                  `json:"CountryFill,omitempty" xml:"CountryFill,omitempty"`
	NationalityFill   *bool                                  `json:"NationalityFill,omitempty" xml:"NationalityFill,omitempty"`
	StampFill         *bool                                  `json:"StampFill,omitempty" xml:"StampFill,omitempty"`
	TradeMarkNameFill *bool                                  `json:"TradeMarkNameFill,omitempty" xml:"TradeMarkNameFill,omitempty"`
	MaterialNameFill  *bool                                  `json:"MaterialNameFill,omitempty" xml:"MaterialNameFill,omitempty"`
	ErrorMsgs         *CheckLoaFillResponseBodyDataErrorMsgs `json:"ErrorMsgs,omitempty" xml:"ErrorMsgs,omitempty" type:"Struct"`
}

func (s CheckLoaFillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckLoaFillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckLoaFillResponseBodyData) SetAddressFill(v bool) *CheckLoaFillResponseBodyData {
	s.AddressFill = &v
	return s
}

func (s *CheckLoaFillResponseBodyData) SetTemplateUrl(v string) *CheckLoaFillResponseBodyData {
	s.TemplateUrl = &v
	return s
}

func (s *CheckLoaFillResponseBodyData) SetCountryFill(v bool) *CheckLoaFillResponseBodyData {
	s.CountryFill = &v
	return s
}

func (s *CheckLoaFillResponseBodyData) SetNationalityFill(v bool) *CheckLoaFillResponseBodyData {
	s.NationalityFill = &v
	return s
}

func (s *CheckLoaFillResponseBodyData) SetStampFill(v bool) *CheckLoaFillResponseBodyData {
	s.StampFill = &v
	return s
}

func (s *CheckLoaFillResponseBodyData) SetTradeMarkNameFill(v bool) *CheckLoaFillResponseBodyData {
	s.TradeMarkNameFill = &v
	return s
}

func (s *CheckLoaFillResponseBodyData) SetMaterialNameFill(v bool) *CheckLoaFillResponseBodyData {
	s.MaterialNameFill = &v
	return s
}

func (s *CheckLoaFillResponseBodyData) SetErrorMsgs(v *CheckLoaFillResponseBodyDataErrorMsgs) *CheckLoaFillResponseBodyData {
	s.ErrorMsgs = v
	return s
}

type CheckLoaFillResponseBodyDataErrorMsgs struct {
	ErrorMsg []*string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty" type:"Repeated"`
}

func (s CheckLoaFillResponseBodyDataErrorMsgs) String() string {
	return tea.Prettify(s)
}

func (s CheckLoaFillResponseBodyDataErrorMsgs) GoString() string {
	return s.String()
}

func (s *CheckLoaFillResponseBodyDataErrorMsgs) SetErrorMsg(v []*string) *CheckLoaFillResponseBodyDataErrorMsgs {
	s.ErrorMsg = v
	return s
}

type CheckLoaFillResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckLoaFillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckLoaFillResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckLoaFillResponse) GoString() string {
	return s.String()
}

func (s *CheckLoaFillResponse) SetHeaders(v map[string]*string) *CheckLoaFillResponse {
	s.Headers = v
	return s
}

func (s *CheckLoaFillResponse) SetBody(v *CheckLoaFillResponseBody) *CheckLoaFillResponse {
	s.Body = v
	return s
}

type ConfirmApplicantRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note  *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s ConfirmApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmApplicantRequest) GoString() string {
	return s.String()
}

func (s *ConfirmApplicantRequest) SetBizId(v string) *ConfirmApplicantRequest {
	s.BizId = &v
	return s
}

func (s *ConfirmApplicantRequest) SetNote(v string) *ConfirmApplicantRequest {
	s.Note = &v
	return s
}

type ConfirmApplicantResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s ConfirmApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmApplicantResponseBody) SetErrorMsg(v string) *ConfirmApplicantResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ConfirmApplicantResponseBody) SetRequestId(v string) *ConfirmApplicantResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmApplicantResponseBody) SetSuccess(v bool) *ConfirmApplicantResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmApplicantResponseBody) SetErrorCode(v string) *ConfirmApplicantResponseBody {
	s.ErrorCode = &v
	return s
}

type ConfirmApplicantResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfirmApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfirmApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmApplicantResponse) GoString() string {
	return s.String()
}

func (s *ConfirmApplicantResponse) SetHeaders(v map[string]*string) *ConfirmApplicantResponse {
	s.Headers = v
	return s
}

func (s *ConfirmApplicantResponse) SetBody(v *ConfirmApplicantResponseBody) *ConfirmApplicantResponse {
	s.Body = v
	return s
}

type CreateIntentionOrderRequest struct {
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Channel        *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s CreateIntentionOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentionOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentionOrderRequest) SetIntentionBizId(v string) *CreateIntentionOrderRequest {
	s.IntentionBizId = &v
	return s
}

func (s *CreateIntentionOrderRequest) SetChannel(v string) *CreateIntentionOrderRequest {
	s.Channel = &v
	return s
}

type CreateIntentionOrderResponseBody struct {
	ErrorMsg  *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *CreateIntentionOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateIntentionOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentionOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntentionOrderResponseBody) SetErrorMsg(v string) *CreateIntentionOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateIntentionOrderResponseBody) SetRequestId(v string) *CreateIntentionOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntentionOrderResponseBody) SetSuccess(v bool) *CreateIntentionOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIntentionOrderResponseBody) SetData(v *CreateIntentionOrderResponseBodyData) *CreateIntentionOrderResponseBody {
	s.Data = v
	return s
}

type CreateIntentionOrderResponseBodyData struct {
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateIntentionOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentionOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIntentionOrderResponseBodyData) SetOrderIds(v []*string) *CreateIntentionOrderResponseBodyData {
	s.OrderIds = v
	return s
}

type CreateIntentionOrderResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIntentionOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIntentionOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentionOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateIntentionOrderResponse) SetHeaders(v map[string]*string) *CreateIntentionOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateIntentionOrderResponse) SetBody(v *CreateIntentionOrderResponseBody) *CreateIntentionOrderResponse {
	s.Body = v
	return s
}

type QueryOssResourcesRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s QueryOssResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOssResourcesRequest) GoString() string {
	return s.String()
}

func (s *QueryOssResourcesRequest) SetBizId(v string) *QueryOssResourcesRequest {
	s.BizId = &v
	return s
}

type QueryOssResourcesResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryOssResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryOssResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOssResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOssResourcesResponseBody) SetRequestId(v string) *QueryOssResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOssResourcesResponseBody) SetData(v *QueryOssResourcesResponseBodyData) *QueryOssResourcesResponseBody {
	s.Data = v
	return s
}

type QueryOssResourcesResponseBodyData struct {
	TaskList []*QueryOssResourcesResponseBodyDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s QueryOssResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOssResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOssResourcesResponseBodyData) SetTaskList(v []*QueryOssResourcesResponseBodyDataTaskList) *QueryOssResourcesResponseBodyData {
	s.TaskList = v
	return s
}

type QueryOssResourcesResponseBodyDataTaskList struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	OssUrl     *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s QueryOssResourcesResponseBodyDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s QueryOssResourcesResponseBodyDataTaskList) GoString() string {
	return s.String()
}

func (s *QueryOssResourcesResponseBodyDataTaskList) SetBizId(v string) *QueryOssResourcesResponseBodyDataTaskList {
	s.BizId = &v
	return s
}

func (s *QueryOssResourcesResponseBodyDataTaskList) SetUpdateTime(v int64) *QueryOssResourcesResponseBodyDataTaskList {
	s.UpdateTime = &v
	return s
}

func (s *QueryOssResourcesResponseBodyDataTaskList) SetOssUrl(v string) *QueryOssResourcesResponseBodyDataTaskList {
	s.OssUrl = &v
	return s
}

func (s *QueryOssResourcesResponseBodyDataTaskList) SetName(v string) *QueryOssResourcesResponseBodyDataTaskList {
	s.Name = &v
	return s
}

func (s *QueryOssResourcesResponseBodyDataTaskList) SetCreateTime(v int64) *QueryOssResourcesResponseBodyDataTaskList {
	s.CreateTime = &v
	return s
}

type QueryOssResourcesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOssResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOssResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOssResourcesResponse) GoString() string {
	return s.String()
}

func (s *QueryOssResourcesResponse) SetHeaders(v map[string]*string) *QueryOssResourcesResponse {
	s.Headers = v
	return s
}

func (s *QueryOssResourcesResponse) SetBody(v *QueryOssResourcesResponseBody) *QueryOssResourcesResponse {
	s.Body = v
	return s
}

type RefuseApplicantRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note  *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s RefuseApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s RefuseApplicantRequest) GoString() string {
	return s.String()
}

func (s *RefuseApplicantRequest) SetBizId(v string) *RefuseApplicantRequest {
	s.BizId = &v
	return s
}

func (s *RefuseApplicantRequest) SetNote(v string) *RefuseApplicantRequest {
	s.Note = &v
	return s
}

type RefuseApplicantResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s RefuseApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefuseApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *RefuseApplicantResponseBody) SetErrorMsg(v string) *RefuseApplicantResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefuseApplicantResponseBody) SetRequestId(v string) *RefuseApplicantResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefuseApplicantResponseBody) SetSuccess(v bool) *RefuseApplicantResponseBody {
	s.Success = &v
	return s
}

func (s *RefuseApplicantResponseBody) SetErrorCode(v string) *RefuseApplicantResponseBody {
	s.ErrorCode = &v
	return s
}

type RefuseApplicantResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefuseApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefuseApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s RefuseApplicantResponse) GoString() string {
	return s.String()
}

func (s *RefuseApplicantResponse) SetHeaders(v map[string]*string) *RefuseApplicantResponse {
	s.Headers = v
	return s
}

func (s *RefuseApplicantResponse) SetBody(v *RefuseApplicantResponseBody) *RefuseApplicantResponse {
	s.Body = v
	return s
}

type CreateIntentionOrderGeneratingPayRequest struct {
	IntentionBizId  *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	PaymentCallback *string `json:"PaymentCallback,omitempty" xml:"PaymentCallback,omitempty"`
	Channel         *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s CreateIntentionOrderGeneratingPayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentionOrderGeneratingPayRequest) GoString() string {
	return s.String()
}

func (s *CreateIntentionOrderGeneratingPayRequest) SetIntentionBizId(v string) *CreateIntentionOrderGeneratingPayRequest {
	s.IntentionBizId = &v
	return s
}

func (s *CreateIntentionOrderGeneratingPayRequest) SetPaymentCallback(v string) *CreateIntentionOrderGeneratingPayRequest {
	s.PaymentCallback = &v
	return s
}

func (s *CreateIntentionOrderGeneratingPayRequest) SetChannel(v string) *CreateIntentionOrderGeneratingPayRequest {
	s.Channel = &v
	return s
}

type CreateIntentionOrderGeneratingPayResponseBody struct {
	ErrorMsg  *string  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PayUrl    *string  `json:"PayUrl,omitempty" xml:"PayUrl,omitempty"`
	Success   *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	OrderIds  []*int64 `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateIntentionOrderGeneratingPayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentionOrderGeneratingPayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntentionOrderGeneratingPayResponseBody) SetErrorMsg(v string) *CreateIntentionOrderGeneratingPayResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateIntentionOrderGeneratingPayResponseBody) SetRequestId(v string) *CreateIntentionOrderGeneratingPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntentionOrderGeneratingPayResponseBody) SetPayUrl(v string) *CreateIntentionOrderGeneratingPayResponseBody {
	s.PayUrl = &v
	return s
}

func (s *CreateIntentionOrderGeneratingPayResponseBody) SetSuccess(v bool) *CreateIntentionOrderGeneratingPayResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIntentionOrderGeneratingPayResponseBody) SetOrderIds(v []*int64) *CreateIntentionOrderGeneratingPayResponseBody {
	s.OrderIds = v
	return s
}

type CreateIntentionOrderGeneratingPayResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIntentionOrderGeneratingPayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIntentionOrderGeneratingPayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntentionOrderGeneratingPayResponse) GoString() string {
	return s.String()
}

func (s *CreateIntentionOrderGeneratingPayResponse) SetHeaders(v map[string]*string) *CreateIntentionOrderGeneratingPayResponse {
	s.Headers = v
	return s
}

func (s *CreateIntentionOrderGeneratingPayResponse) SetBody(v *CreateIntentionOrderGeneratingPayResponseBody) *CreateIntentionOrderGeneratingPayResponse {
	s.Body = v
	return s
}

type ListNotaryOrdersRequest struct {
	StartOrderDate *int64  `json:"StartOrderDate,omitempty" xml:"StartOrderDate,omitempty"`
	EndOrderDate   *int64  `json:"EndOrderDate,omitempty" xml:"EndOrderDate,omitempty"`
	NotaryStatus   *int32  `json:"NotaryStatus,omitempty" xml:"NotaryStatus,omitempty"`
	AliyunOrderId  *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	SortByType     *string `json:"SortByType,omitempty" xml:"SortByType,omitempty"`
	SortKeyType    *int32  `json:"SortKeyType,omitempty" xml:"SortKeyType,omitempty"`
	PageNum        *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	NotaryType     *int32  `json:"NotaryType,omitempty" xml:"NotaryType,omitempty"`
}

func (s ListNotaryOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListNotaryOrdersRequest) SetStartOrderDate(v int64) *ListNotaryOrdersRequest {
	s.StartOrderDate = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetEndOrderDate(v int64) *ListNotaryOrdersRequest {
	s.EndOrderDate = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetNotaryStatus(v int32) *ListNotaryOrdersRequest {
	s.NotaryStatus = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetAliyunOrderId(v string) *ListNotaryOrdersRequest {
	s.AliyunOrderId = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetSortByType(v string) *ListNotaryOrdersRequest {
	s.SortByType = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetSortKeyType(v int32) *ListNotaryOrdersRequest {
	s.SortKeyType = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetPageNum(v int32) *ListNotaryOrdersRequest {
	s.PageNum = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetPageSize(v int32) *ListNotaryOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetBizId(v string) *ListNotaryOrdersRequest {
	s.BizId = &v
	return s
}

func (s *ListNotaryOrdersRequest) SetNotaryType(v int32) *ListNotaryOrdersRequest {
	s.NotaryType = &v
	return s
}

type ListNotaryOrdersResponseBody struct {
	NextPage       *bool                             `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode      *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	TotalItemNum   *int32                            `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	PrePage        *bool                             `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	CurrentPageNum *int32                            `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	ErrorMsg       *string                           `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	TotalPageNum   *int32                            `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
	PageSize       *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data           *ListNotaryOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListNotaryOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListNotaryOrdersResponseBody) SetNextPage(v bool) *ListNotaryOrdersResponseBody {
	s.NextPage = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetRequestId(v string) *ListNotaryOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetSuccess(v bool) *ListNotaryOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetErrorCode(v string) *ListNotaryOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetTotalItemNum(v int32) *ListNotaryOrdersResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetPrePage(v bool) *ListNotaryOrdersResponseBody {
	s.PrePage = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetCurrentPageNum(v int32) *ListNotaryOrdersResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetErrorMsg(v string) *ListNotaryOrdersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetTotalPageNum(v int32) *ListNotaryOrdersResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetPageSize(v int32) *ListNotaryOrdersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNotaryOrdersResponseBody) SetData(v *ListNotaryOrdersResponseBodyData) *ListNotaryOrdersResponseBody {
	s.Data = v
	return s
}

type ListNotaryOrdersResponseBodyData struct {
	NotaryOrder []*ListNotaryOrdersResponseBodyDataNotaryOrder `json:"NotaryOrder,omitempty" xml:"NotaryOrder,omitempty" type:"Repeated"`
}

func (s ListNotaryOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNotaryOrdersResponseBodyData) SetNotaryOrder(v []*ListNotaryOrdersResponseBodyDataNotaryOrder) *ListNotaryOrdersResponseBodyData {
	s.NotaryOrder = v
	return s
}

type ListNotaryOrdersResponseBodyDataNotaryOrder struct {
	OrderDate          *int64   `json:"OrderDate,omitempty" xml:"OrderDate,omitempty"`
	OrderPrice         *float32 `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	NotaryType         *int32   `json:"NotaryType,omitempty" xml:"NotaryType,omitempty"`
	TmClassification   *string  `json:"TmClassification,omitempty" xml:"TmClassification,omitempty"`
	BizId              *string  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GmtModified        *int64   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	NotaryStatus       *int32   `json:"NotaryStatus,omitempty" xml:"NotaryStatus,omitempty"`
	NotaryOrderId      *int64   `json:"NotaryOrderId,omitempty" xml:"NotaryOrderId,omitempty"`
	TmName             *string  `json:"TmName,omitempty" xml:"TmName,omitempty"`
	TmRegisterNo       *string  `json:"TmRegisterNo,omitempty" xml:"TmRegisterNo,omitempty"`
	TmImage            *string  `json:"TmImage,omitempty" xml:"TmImage,omitempty"`
	AliyunOrderId      *string  `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	ApplyPostStatus    *string  `json:"ApplyPostStatus,omitempty" xml:"ApplyPostStatus,omitempty"`
	NotaryCertificate  *string  `json:"NotaryCertificate,omitempty" xml:"NotaryCertificate,omitempty"`
	NotaryPlatformName *string  `json:"NotaryPlatformName,omitempty" xml:"NotaryPlatformName,omitempty"`
}

func (s ListNotaryOrdersResponseBodyDataNotaryOrder) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryOrdersResponseBodyDataNotaryOrder) GoString() string {
	return s.String()
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetOrderDate(v int64) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.OrderDate = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetOrderPrice(v float32) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.OrderPrice = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetNotaryType(v int32) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.NotaryType = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetTmClassification(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.TmClassification = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetBizId(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.BizId = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetGmtModified(v int64) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.GmtModified = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetNotaryStatus(v int32) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.NotaryStatus = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetNotaryOrderId(v int64) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.NotaryOrderId = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetTmName(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.TmName = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetTmRegisterNo(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.TmRegisterNo = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetTmImage(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.TmImage = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetAliyunOrderId(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.AliyunOrderId = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetApplyPostStatus(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.ApplyPostStatus = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetNotaryCertificate(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.NotaryCertificate = &v
	return s
}

func (s *ListNotaryOrdersResponseBodyDataNotaryOrder) SetNotaryPlatformName(v string) *ListNotaryOrdersResponseBodyDataNotaryOrder {
	s.NotaryPlatformName = &v
	return s
}

type ListNotaryOrdersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNotaryOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNotaryOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNotaryOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListNotaryOrdersResponse) SetHeaders(v map[string]*string) *ListNotaryOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListNotaryOrdersResponse) SetBody(v *ListNotaryOrdersResponseBody) *ListNotaryOrdersResponse {
	s.Body = v
	return s
}

type GetSupportPrincipalNameResponseBody struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Principals []*GetSupportPrincipalNameResponseBodyPrincipals `json:"Principals,omitempty" xml:"Principals,omitempty" type:"Repeated"`
}

func (s GetSupportPrincipalNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSupportPrincipalNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupportPrincipalNameResponseBody) SetRequestId(v string) *GetSupportPrincipalNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupportPrincipalNameResponseBody) SetPrincipals(v []*GetSupportPrincipalNameResponseBodyPrincipals) *GetSupportPrincipalNameResponseBody {
	s.Principals = v
	return s
}

type GetSupportPrincipalNameResponseBodyPrincipals struct {
	PrincipalDescription *string `json:"PrincipalDescription,omitempty" xml:"PrincipalDescription,omitempty"`
	DefaultPrincipal     *bool   `json:"DefaultPrincipal,omitempty" xml:"DefaultPrincipal,omitempty"`
	PrincipalValue       *int32  `json:"PrincipalValue,omitempty" xml:"PrincipalValue,omitempty"`
}

func (s GetSupportPrincipalNameResponseBodyPrincipals) String() string {
	return tea.Prettify(s)
}

func (s GetSupportPrincipalNameResponseBodyPrincipals) GoString() string {
	return s.String()
}

func (s *GetSupportPrincipalNameResponseBodyPrincipals) SetPrincipalDescription(v string) *GetSupportPrincipalNameResponseBodyPrincipals {
	s.PrincipalDescription = &v
	return s
}

func (s *GetSupportPrincipalNameResponseBodyPrincipals) SetDefaultPrincipal(v bool) *GetSupportPrincipalNameResponseBodyPrincipals {
	s.DefaultPrincipal = &v
	return s
}

func (s *GetSupportPrincipalNameResponseBodyPrincipals) SetPrincipalValue(v int32) *GetSupportPrincipalNameResponseBodyPrincipals {
	s.PrincipalValue = &v
	return s
}

type GetSupportPrincipalNameResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSupportPrincipalNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSupportPrincipalNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSupportPrincipalNameResponse) GoString() string {
	return s.String()
}

func (s *GetSupportPrincipalNameResponse) SetHeaders(v map[string]*string) *GetSupportPrincipalNameResponse {
	s.Headers = v
	return s
}

func (s *GetSupportPrincipalNameResponse) SetBody(v *GetSupportPrincipalNameResponseBody) *GetSupportPrincipalNameResponse {
	s.Body = v
	return s
}

type StartNotaryRequest struct {
	NotaryOrderId *int64 `json:"NotaryOrderId,omitempty" xml:"NotaryOrderId,omitempty"`
}

func (s StartNotaryRequest) String() string {
	return tea.Prettify(s)
}

func (s StartNotaryRequest) GoString() string {
	return s.String()
}

func (s *StartNotaryRequest) SetNotaryOrderId(v int64) *StartNotaryRequest {
	s.NotaryOrderId = &v
	return s
}

type StartNotaryResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	NotaryUrl *string `json:"NotaryUrl,omitempty" xml:"NotaryUrl,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s StartNotaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartNotaryResponseBody) GoString() string {
	return s.String()
}

func (s *StartNotaryResponseBody) SetErrorMsg(v string) *StartNotaryResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StartNotaryResponseBody) SetRequestId(v string) *StartNotaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartNotaryResponseBody) SetSuccess(v bool) *StartNotaryResponseBody {
	s.Success = &v
	return s
}

func (s *StartNotaryResponseBody) SetNotaryUrl(v string) *StartNotaryResponseBody {
	s.NotaryUrl = &v
	return s
}

func (s *StartNotaryResponseBody) SetErrorCode(v string) *StartNotaryResponseBody {
	s.ErrorCode = &v
	return s
}

type StartNotaryResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartNotaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartNotaryResponse) String() string {
	return tea.Prettify(s)
}

func (s StartNotaryResponse) GoString() string {
	return s.String()
}

func (s *StartNotaryResponse) SetHeaders(v map[string]*string) *StartNotaryResponse {
	s.Headers = v
	return s
}

func (s *StartNotaryResponse) SetBody(v *StartNotaryResponseBody) *StartNotaryResponse {
	s.Body = v
	return s
}

type UpdateSendMaterialNumRequest struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Num         *string `json:"Num,omitempty" xml:"Num,omitempty"`
	OperateType *int32  `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s UpdateSendMaterialNumRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSendMaterialNumRequest) GoString() string {
	return s.String()
}

func (s *UpdateSendMaterialNumRequest) SetBizId(v string) *UpdateSendMaterialNumRequest {
	s.BizId = &v
	return s
}

func (s *UpdateSendMaterialNumRequest) SetNum(v string) *UpdateSendMaterialNumRequest {
	s.Num = &v
	return s
}

func (s *UpdateSendMaterialNumRequest) SetOperateType(v int32) *UpdateSendMaterialNumRequest {
	s.OperateType = &v
	return s
}

type UpdateSendMaterialNumResponseBody struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s UpdateSendMaterialNumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSendMaterialNumResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSendMaterialNumResponseBody) SetErrorMsg(v string) *UpdateSendMaterialNumResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateSendMaterialNumResponseBody) SetRequestId(v string) *UpdateSendMaterialNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSendMaterialNumResponseBody) SetSuccess(v bool) *UpdateSendMaterialNumResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSendMaterialNumResponseBody) SetErrorCode(v string) *UpdateSendMaterialNumResponseBody {
	s.ErrorCode = &v
	return s
}

type UpdateSendMaterialNumResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSendMaterialNumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSendMaterialNumResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSendMaterialNumResponse) GoString() string {
	return s.String()
}

func (s *UpdateSendMaterialNumResponse) SetHeaders(v map[string]*string) *UpdateSendMaterialNumResponse {
	s.Headers = v
	return s
}

func (s *UpdateSendMaterialNumResponse) SetBody(v *UpdateSendMaterialNumResponseBody) *UpdateSendMaterialNumResponse {
	s.Body = v
	return s
}

type DeleteTrademarkApplicationRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DeleteTrademarkApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrademarkApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrademarkApplicationRequest) SetBizId(v string) *DeleteTrademarkApplicationRequest {
	s.BizId = &v
	return s
}

type DeleteTrademarkApplicationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTrademarkApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrademarkApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrademarkApplicationResponseBody) SetCode(v string) *DeleteTrademarkApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTrademarkApplicationResponseBody) SetMessage(v string) *DeleteTrademarkApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTrademarkApplicationResponseBody) SetRequestId(v string) *DeleteTrademarkApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrademarkApplicationResponseBody) SetSuccess(v bool) *DeleteTrademarkApplicationResponseBody {
	s.Success = &v
	return s
}

type DeleteTrademarkApplicationResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTrademarkApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTrademarkApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrademarkApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrademarkApplicationResponse) SetHeaders(v map[string]*string) *DeleteTrademarkApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrademarkApplicationResponse) SetBody(v *DeleteTrademarkApplicationResponseBody) *DeleteTrademarkApplicationResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("trademark"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) QueryTradeProduceListWithOptions(request *QueryTradeProduceListRequest, runtime *util.RuntimeOptions) (_result *QueryTradeProduceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTradeProduceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTradeProduceList"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTradeProduceList(request *QueryTradeProduceListRequest) (_result *QueryTradeProduceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTradeProduceListResponse{}
	_body, _err := client.QueryTradeProduceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTrademarkMonitorResultsWithOptions(request *QueryTrademarkMonitorResultsRequest, runtime *util.RuntimeOptions) (_result *QueryTrademarkMonitorResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTrademarkMonitorResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTrademarkMonitorResults"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTrademarkMonitorResults(request *QueryTrademarkMonitorResultsRequest) (_result *QueryTrademarkMonitorResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTrademarkMonitorResultsResponse{}
	_body, _err := client.QueryTrademarkMonitorResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelTradeOrderWithOptions(request *CancelTradeOrderRequest, runtime *util.RuntimeOptions) (_result *CancelTradeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelTradeOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelTradeOrder"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelTradeOrder(request *CancelTradeOrderRequest) (_result *CancelTradeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelTradeOrderResponse{}
	_body, _err := client.CancelTradeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTmMonitorRuleWithOptions(request *DeleteTmMonitorRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteTmMonitorRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTmMonitorRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTmMonitorRule"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTmMonitorRule(request *DeleteTmMonitorRuleRequest) (_result *DeleteTmMonitorRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTmMonitorRuleResponse{}
	_body, _err := client.DeleteTmMonitorRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadNotaryDataWithOptions(request *UploadNotaryDataRequest, runtime *util.RuntimeOptions) (_result *UploadNotaryDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadNotaryDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadNotaryData"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadNotaryData(request *UploadNotaryDataRequest) (_result *UploadNotaryDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadNotaryDataResponse{}
	_body, _err := client.UploadNotaryDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyApplicantWithOptions(request *CopyApplicantRequest, runtime *util.RuntimeOptions) (_result *CopyApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CopyApplicantResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CopyApplicant"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyApplicant(request *CopyApplicantRequest) (_result *CopyApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyApplicantResponse{}
	_body, _err := client.CopyApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PartnerUpdateTrademarkNameWithOptions(request *PartnerUpdateTrademarkNameRequest, runtime *util.RuntimeOptions) (_result *PartnerUpdateTrademarkNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PartnerUpdateTrademarkNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PartnerUpdateTrademarkName"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PartnerUpdateTrademarkName(request *PartnerUpdateTrademarkNameRequest) (_result *PartnerUpdateTrademarkNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PartnerUpdateTrademarkNameResponse{}
	_body, _err := client.PartnerUpdateTrademarkNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryIntentionDetailWithOptions(request *QueryIntentionDetailRequest, runtime *util.RuntimeOptions) (_result *QueryIntentionDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryIntentionDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryIntentionDetail"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryIntentionDetail(request *QueryIntentionDetailRequest) (_result *QueryIntentionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryIntentionDetailResponse{}
	_body, _err := client.QueryIntentionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryIntentionPriceWithOptions(request *QueryIntentionPriceRequest, runtime *util.RuntimeOptions) (_result *QueryIntentionPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryIntentionPriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryIntentionPrice"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryIntentionPrice(request *QueryIntentionPriceRequest) (_result *QueryIntentionPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryIntentionPriceResponse{}
	_body, _err := client.QueryIntentionPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOfficialFileCustomListWithOptions(request *QueryOfficialFileCustomListRequest, runtime *util.RuntimeOptions) (_result *QueryOfficialFileCustomListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOfficialFileCustomListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOfficialFileCustomList"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOfficialFileCustomList(request *QueryOfficialFileCustomListRequest) (_result *QueryOfficialFileCustomListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOfficialFileCustomListResponse{}
	_body, _err := client.QueryOfficialFileCustomListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckTrademarkIconWithOptions(request *CheckTrademarkIconRequest, runtime *util.RuntimeOptions) (_result *CheckTrademarkIconResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckTrademarkIconResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckTrademarkIcon"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckTrademarkIcon(request *CheckTrademarkIconRequest) (_result *CheckTrademarkIconResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckTrademarkIconResponse{}
	_body, _err := client.CheckTrademarkIconWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySupplementDetailWithOptions(request *QuerySupplementDetailRequest, runtime *util.RuntimeOptions) (_result *QuerySupplementDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySupplementDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySupplementDetail"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySupplementDetail(request *QuerySupplementDetailRequest) (_result *QuerySupplementDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySupplementDetailResponse{}
	_body, _err := client.QuerySupplementDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadTrademarkOnSaleWithOptions(request *UploadTrademarkOnSaleRequest, runtime *util.RuntimeOptions) (_result *UploadTrademarkOnSaleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadTrademarkOnSaleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadTrademarkOnSale"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadTrademarkOnSale(request *UploadTrademarkOnSaleRequest) (_result *UploadTrademarkOnSaleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadTrademarkOnSaleResponse{}
	_body, _err := client.UploadTrademarkOnSaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyNotaryPostWithOptions(request *ApplyNotaryPostRequest, runtime *util.RuntimeOptions) (_result *ApplyNotaryPostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyNotaryPostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyNotaryPost"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyNotaryPost(request *ApplyNotaryPostRequest) (_result *ApplyNotaryPostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyNotaryPostResponse{}
	_body, _err := client.ApplyNotaryPostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTradeMarkApplicationsByIntentionWithOptions(request *QueryTradeMarkApplicationsByIntentionRequest, runtime *util.RuntimeOptions) (_result *QueryTradeMarkApplicationsByIntentionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTradeMarkApplicationsByIntentionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTradeMarkApplicationsByIntention"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTradeMarkApplicationsByIntention(request *QueryTradeMarkApplicationsByIntentionRequest) (_result *QueryTradeMarkApplicationsByIntentionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTradeMarkApplicationsByIntentionResponse{}
	_body, _err := client.QueryTradeMarkApplicationsByIntentionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveExtensionAttributeWithOptions(request *SaveExtensionAttributeRequest, runtime *util.RuntimeOptions) (_result *SaveExtensionAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveExtensionAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveExtensionAttribute"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveExtensionAttribute(request *SaveExtensionAttributeRequest) (_result *SaveExtensionAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveExtensionAttributeResponse{}
	_body, _err := client.SaveExtensionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AcceptPartnerNotificationWithOptions(request *AcceptPartnerNotificationRequest, runtime *util.RuntimeOptions) (_result *AcceptPartnerNotificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AcceptPartnerNotificationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AcceptPartnerNotification"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AcceptPartnerNotification(request *AcceptPartnerNotificationRequest) (_result *AcceptPartnerNotificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptPartnerNotificationResponse{}
	_body, _err := client.AcceptPartnerNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSupplementWithOptions(tmpReq *SubmitSupplementRequest, runtime *util.RuntimeOptions) (_result *SubmitSupplementResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitSupplementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UploadOssKeyList)) {
		request.UploadOssKeyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UploadOssKeyList, tea.String("UploadOssKeyList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSupplementResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSupplement"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSupplement(request *SubmitSupplementRequest) (_result *SubmitSupplementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSupplementResponse{}
	_body, _err := client.SubmitSupplementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ForceUploadTrademarkOnsaleWithOptions(request *ForceUploadTrademarkOnsaleRequest, runtime *util.RuntimeOptions) (_result *ForceUploadTrademarkOnsaleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ForceUploadTrademarkOnsaleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ForceUploadTrademarkOnsale"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ForceUploadTrademarkOnsale(request *ForceUploadTrademarkOnsaleRequest) (_result *ForceUploadTrademarkOnsaleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ForceUploadTrademarkOnsaleResponse{}
	_body, _err := client.ForceUploadTrademarkOnsaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindMaterialWithOptions(request *BindMaterialRequest, runtime *util.RuntimeOptions) (_result *BindMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindMaterial"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindMaterial(request *BindMaterialRequest) (_result *BindMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindMaterialResponse{}
	_body, _err := client.BindMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultPrincipalWithOptions(runtime *util.RuntimeOptions) (_result *GetDefaultPrincipalResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &GetDefaultPrincipalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDefaultPrincipal"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultPrincipal() (_result *GetDefaultPrincipalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDefaultPrincipalResponse{}
	_body, _err := client.GetDefaultPrincipalWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCommunicationLogsWithOptions(request *QueryCommunicationLogsRequest, runtime *util.RuntimeOptions) (_result *QueryCommunicationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCommunicationLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCommunicationLogs"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCommunicationLogs(request *QueryCommunicationLogsRequest) (_result *QueryCommunicationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCommunicationLogsResponse{}
	_body, _err := client.QueryCommunicationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateQrCodeWithOptions(request *GenerateQrCodeRequest, runtime *util.RuntimeOptions) (_result *GenerateQrCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateQrCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateQrCode"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateQrCode(request *GenerateQrCodeRequest) (_result *GenerateQrCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateQrCodeResponse{}
	_body, _err := client.GenerateQrCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmDissentOriginalWithOptions(request *ConfirmDissentOriginalRequest, runtime *util.RuntimeOptions) (_result *ConfirmDissentOriginalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfirmDissentOriginalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfirmDissentOriginal"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmDissentOriginal(request *ConfirmDissentOriginalRequest) (_result *ConfirmDissentOriginalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmDissentOriginalResponse{}
	_body, _err := client.ConfirmDissentOriginalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertImageToGrayWithOptions(request *ConvertImageToGrayRequest, runtime *util.RuntimeOptions) (_result *ConvertImageToGrayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConvertImageToGrayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConvertImageToGray"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertImageToGray(request *ConvertImageToGrayRequest) (_result *ConvertImageToGrayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertImageToGrayResponse{}
	_body, _err := client.ConvertImageToGrayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryIntentionListWithOptions(request *QueryIntentionListRequest, runtime *util.RuntimeOptions) (_result *QueryIntentionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryIntentionListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryIntentionList"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryIntentionList(request *QueryIntentionListRequest) (_result *QueryIntentionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryIntentionListResponse{}
	_body, _err := client.QueryIntentionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthorizationLetterVersionWithOptions(request *GetAuthorizationLetterVersionRequest, runtime *util.RuntimeOptions) (_result *GetAuthorizationLetterVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAuthorizationLetterVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAuthorizationLetterVersion"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthorizationLetterVersion(request *GetAuthorizationLetterVersionRequest) (_result *GetAuthorizationLetterVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthorizationLetterVersionResponse{}
	_body, _err := client.GetAuthorizationLetterVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTrademarkPriceWithOptions(tmpReq *QueryTrademarkPriceRequest, runtime *util.RuntimeOptions) (_result *QueryTrademarkPriceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryTrademarkPriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OrderData)) {
		request.OrderDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderData, tea.String("OrderData"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTrademarkPriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTrademarkPrice"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTrademarkPrice(request *QueryTrademarkPriceRequest) (_result *QueryTrademarkPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTrademarkPriceResponse{}
	_body, _err := client.QueryTrademarkPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertTmMonitorRuleWithOptions(tmpReq *InsertTmMonitorRuleRequest, runtime *util.RuntimeOptions) (_result *InsertTmMonitorRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InsertTmMonitorRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Classification)) {
		request.ClassificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Classification, tea.String("Classification"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NotifyStatus)) {
		request.NotifyStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyStatus, tea.String("NotifyStatus"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertTmMonitorRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertTmMonitorRule"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertTmMonitorRule(request *InsertTmMonitorRuleRequest) (_result *InsertTmMonitorRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertTmMonitorRuleResponse{}
	_body, _err := client.InsertTmMonitorRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTrademarkMonitorRulesWithOptions(request *QueryTrademarkMonitorRulesRequest, runtime *util.RuntimeOptions) (_result *QueryTrademarkMonitorRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTrademarkMonitorRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTrademarkMonitorRules"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTrademarkMonitorRules(request *QueryTrademarkMonitorRulesRequest) (_result *QueryTrademarkMonitorRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTrademarkMonitorRulesResponse{}
	_body, _err := client.QueryTrademarkMonitorRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DenySupplementWithOptions(request *DenySupplementRequest, runtime *util.RuntimeOptions) (_result *DenySupplementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DenySupplementResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DenySupplement"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DenySupplement(request *DenySupplementRequest) (_result *DenySupplementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DenySupplementResponse{}
	_body, _err := client.DenySupplementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMaterialWithOptions(request *QueryMaterialRequest, runtime *util.RuntimeOptions) (_result *QueryMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMaterial"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMaterial(request *QueryMaterialRequest) (_result *QueryMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMaterialResponse{}
	_body, _err := client.QueryMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrademarkOrderWithOptions(request *CreateTrademarkOrderRequest, runtime *util.RuntimeOptions) (_result *CreateTrademarkOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTrademarkOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTrademarkOrder"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrademarkOrder(request *CreateTrademarkOrderRequest) (_result *CreateTrademarkOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTrademarkOrderResponse{}
	_body, _err := client.CreateTrademarkOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMaterialListWithOptions(request *QueryMaterialListRequest, runtime *util.RuntimeOptions) (_result *QueryMaterialListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMaterialListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMaterialList"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMaterialList(request *QueryMaterialListRequest) (_result *QueryMaterialListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMaterialListResponse{}
	_body, _err := client.QueryMaterialListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckTrademarkOrderWithOptions(request *CheckTrademarkOrderRequest, runtime *util.RuntimeOptions) (_result *CheckTrademarkOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckTrademarkOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckTrademarkOrder"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckTrademarkOrder(request *CheckTrademarkOrderRequest) (_result *CheckTrademarkOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckTrademarkOrderResponse{}
	_body, _err := client.CheckTrademarkOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTradeMarkApplicationsWithOptions(request *QueryTradeMarkApplicationsRequest, runtime *util.RuntimeOptions) (_result *QueryTradeMarkApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTradeMarkApplicationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTradeMarkApplications"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTradeMarkApplications(request *QueryTradeMarkApplicationsRequest) (_result *QueryTradeMarkApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTradeMarkApplicationsResponse{}
	_body, _err := client.QueryTradeMarkApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicantContacterWithOptions(request *UpdateApplicantContacterRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicantContacterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateApplicantContacterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateApplicantContacter"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicantContacter(request *UpdateApplicantContacterRequest) (_result *UpdateApplicantContacterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicantContacterResponse{}
	_body, _err := client.UpdateApplicantContacterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTaskWithOptions(request *SaveTaskRequest, runtime *util.RuntimeOptions) (_result *SaveTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveTask"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTask(request *SaveTaskRequest) (_result *SaveTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTaskResponse{}
	_body, _err := client.SaveTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTrademarkApplicationComplaintWithOptions(tmpReq *SubmitTrademarkApplicationComplaintRequest, runtime *util.RuntimeOptions) (_result *SubmitTrademarkApplicationComplaintResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTrademarkApplicationComplaintShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Files)) {
		request.FilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Files, tea.String("Files"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitTrademarkApplicationComplaintResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitTrademarkApplicationComplaint"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTrademarkApplicationComplaint(request *SubmitTrademarkApplicationComplaintRequest) (_result *SubmitTrademarkApplicationComplaintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitTrademarkApplicationComplaintResponse{}
	_body, _err := client.SubmitTrademarkApplicationComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WriteIntentionCommunicationLogWithOptions(request *WriteIntentionCommunicationLogRequest, runtime *util.RuntimeOptions) (_result *WriteIntentionCommunicationLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WriteIntentionCommunicationLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("WriteIntentionCommunicationLog"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WriteIntentionCommunicationLog(request *WriteIntentionCommunicationLogRequest) (_result *WriteIntentionCommunicationLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WriteIntentionCommunicationLogResponse{}
	_body, _err := client.WriteIntentionCommunicationLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTaskForOfficialFileCustomWithOptions(request *SaveTaskForOfficialFileCustomRequest, runtime *util.RuntimeOptions) (_result *SaveTaskForOfficialFileCustomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveTaskForOfficialFileCustomResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveTaskForOfficialFileCustom"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTaskForOfficialFileCustom(request *SaveTaskForOfficialFileCustomRequest) (_result *SaveTaskForOfficialFileCustomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTaskForOfficialFileCustomResponse{}
	_body, _err := client.SaveTaskForOfficialFileCustomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescirbeCombineTrademarkWithOptions(request *DescirbeCombineTrademarkRequest, runtime *util.RuntimeOptions) (_result *DescirbeCombineTrademarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescirbeCombineTrademarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescirbeCombineTrademark"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescirbeCombineTrademark(request *DescirbeCombineTrademarkRequest) (_result *DescirbeCombineTrademarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescirbeCombineTrademarkResponse{}
	_body, _err := client.DescirbeCombineTrademarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNotaryOrderWithOptions(request *GetNotaryOrderRequest, runtime *util.RuntimeOptions) (_result *GetNotaryOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetNotaryOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetNotaryOrder"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNotaryOrder(request *GetNotaryOrderRequest) (_result *GetNotaryOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNotaryOrderResponse{}
	_body, _err := client.GetNotaryOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmAdditionalMaterialWithOptions(request *ConfirmAdditionalMaterialRequest, runtime *util.RuntimeOptions) (_result *ConfirmAdditionalMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfirmAdditionalMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfirmAdditionalMaterial"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmAdditionalMaterial(request *ConfirmAdditionalMaterialRequest) (_result *ConfirmAdditionalMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmAdditionalMaterialResponse{}
	_body, _err := client.ConfirmAdditionalMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertRenewInfoWithOptions(request *InsertRenewInfoRequest, runtime *util.RuntimeOptions) (_result *InsertRenewInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertRenewInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertRenewInfo"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertRenewInfo(request *InsertRenewInfoRequest) (_result *InsertRenewInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertRenewInfoResponse{}
	_body, _err := client.InsertRenewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCredentialsInfoWithOptions(request *QueryCredentialsInfoRequest, runtime *util.RuntimeOptions) (_result *QueryCredentialsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCredentialsInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCredentialsInfo"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCredentialsInfo(request *QueryCredentialsInfoRequest) (_result *QueryCredentialsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCredentialsInfoResponse{}
	_body, _err := client.QueryCredentialsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTmOnsalesWithOptions(request *SearchTmOnsalesRequest, runtime *util.RuntimeOptions) (_result *SearchTmOnsalesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchTmOnsalesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTmOnsales"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTmOnsales(request *SearchTmOnsalesRequest) (_result *SearchTmOnsalesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTmOnsalesResponse{}
	_body, _err := client.SearchTmOnsalesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateUploadFilePolicyWithOptions(request *GenerateUploadFilePolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadFilePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateUploadFilePolicy"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateUploadFilePolicy(request *GenerateUploadFilePolicyRequest) (_result *GenerateUploadFilePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.GenerateUploadFilePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMaterialWithOptions(request *DeleteMaterialRequest, runtime *util.RuntimeOptions) (_result *DeleteMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMaterial"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMaterial(request *DeleteMaterialRequest) (_result *DeleteMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMaterialResponse{}
	_body, _err := client.DeleteMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WriteCommunicationLogWithOptions(request *WriteCommunicationLogRequest, runtime *util.RuntimeOptions) (_result *WriteCommunicationLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WriteCommunicationLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("WriteCommunicationLog"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WriteCommunicationLog(request *WriteCommunicationLogRequest) (_result *WriteCommunicationLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WriteCommunicationLogResponse{}
	_body, _err := client.WriteCommunicationLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertTradeIntentionUserWithOptions(request *InsertTradeIntentionUserRequest, runtime *util.RuntimeOptions) (_result *InsertTradeIntentionUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertTradeIntentionUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertTradeIntentionUser"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertTradeIntentionUser(request *InsertTradeIntentionUserRequest) (_result *InsertTradeIntentionUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertTradeIntentionUserResponse{}
	_body, _err := client.InsertTradeIntentionUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryExtensionAttributeWithOptions(request *QueryExtensionAttributeRequest, runtime *util.RuntimeOptions) (_result *QueryExtensionAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryExtensionAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryExtensionAttribute"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryExtensionAttribute(request *QueryExtensionAttributeRequest) (_result *QueryExtensionAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryExtensionAttributeResponse{}
	_body, _err := client.QueryExtensionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTrademarkOnsaleWithOptions(request *UpdateTrademarkOnsaleRequest, runtime *util.RuntimeOptions) (_result *UpdateTrademarkOnsaleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTrademarkOnsaleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTrademarkOnsale"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrademarkOnsale(request *UpdateTrademarkOnsaleRequest) (_result *UpdateTrademarkOnsaleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTrademarkOnsaleResponse{}
	_body, _err := client.UpdateTrademarkOnsaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTradeProduceDetailWithOptions(request *QueryTradeProduceDetailRequest, runtime *util.RuntimeOptions) (_result *QueryTradeProduceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTradeProduceDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTradeProduceDetail"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTradeProduceDetail(request *QueryTradeProduceDetailRequest) (_result *QueryTradeProduceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTradeProduceDetailResponse{}
	_body, _err := client.QueryTradeProduceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryQrCodeUploadStatusWithOptions(request *QueryQrCodeUploadStatusRequest, runtime *util.RuntimeOptions) (_result *QueryQrCodeUploadStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryQrCodeUploadStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryQrCodeUploadStatus"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryQrCodeUploadStatus(request *QueryQrCodeUploadStatusRequest) (_result *QueryQrCodeUploadStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryQrCodeUploadStatusResponse{}
	_body, _err := client.QueryQrCodeUploadStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RejectApplicantWithOptions(request *RejectApplicantRequest, runtime *util.RuntimeOptions) (_result *RejectApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RejectApplicantResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RejectApplicant"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RejectApplicant(request *RejectApplicantRequest) (_result *RejectApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectApplicantResponse{}
	_body, _err := client.RejectApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTradeIntentionUserListWithOptions(request *QueryTradeIntentionUserListRequest, runtime *util.RuntimeOptions) (_result *QueryTradeIntentionUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTradeIntentionUserListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTradeIntentionUserList"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTradeIntentionUserList(request *QueryTradeIntentionUserListRequest) (_result *QueryTradeIntentionUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTradeIntentionUserListResponse{}
	_body, _err := client.QueryTradeIntentionUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StoreMaterialTemporarilyWithOptions(request *StoreMaterialTemporarilyRequest, runtime *util.RuntimeOptions) (_result *StoreMaterialTemporarilyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StoreMaterialTemporarilyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StoreMaterialTemporarily"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StoreMaterialTemporarily(request *StoreMaterialTemporarilyRequest) (_result *StoreMaterialTemporarilyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StoreMaterialTemporarilyResponse{}
	_body, _err := client.StoreMaterialTemporarilyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefuseAdditionalMaterialWithOptions(request *RefuseAdditionalMaterialRequest, runtime *util.RuntimeOptions) (_result *RefuseAdditionalMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefuseAdditionalMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefuseAdditionalMaterial"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefuseAdditionalMaterial(request *RefuseAdditionalMaterialRequest) (_result *RefuseAdditionalMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefuseAdditionalMaterialResponse{}
	_body, _err := client.RefuseAdditionalMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNotaryInfosWithOptions(request *ListNotaryInfosRequest, runtime *util.RuntimeOptions) (_result *ListNotaryInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListNotaryInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNotaryInfos"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNotaryInfos(request *ListNotaryInfosRequest) (_result *ListNotaryInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNotaryInfosResponse{}
	_body, _err := client.ListNotaryInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultPrincipalNameWithOptions(request *GetDefaultPrincipalNameRequest, runtime *util.RuntimeOptions) (_result *GetDefaultPrincipalNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDefaultPrincipalNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDefaultPrincipalName"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultPrincipalName(request *GetDefaultPrincipalNameRequest) (_result *GetDefaultPrincipalNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDefaultPrincipalNameResponse{}
	_body, _err := client.GetDefaultPrincipalNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTradeMarkApplicationDetailWithOptions(request *QueryTradeMarkApplicationDetailRequest, runtime *util.RuntimeOptions) (_result *QueryTradeMarkApplicationDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTradeMarkApplicationDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTradeMarkApplicationDetail"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTradeMarkApplicationDetail(request *QueryTradeMarkApplicationDetailRequest) (_result *QueryTradeMarkApplicationDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTradeMarkApplicationDetailResponse{}
	_body, _err := client.QueryTradeMarkApplicationDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveClassificationConditionsWithOptions(request *SaveClassificationConditionsRequest, runtime *util.RuntimeOptions) (_result *SaveClassificationConditionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveClassificationConditionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveClassificationConditions"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveClassificationConditions(request *SaveClassificationConditionsRequest) (_result *SaveClassificationConditionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveClassificationConditionsResponse{}
	_body, _err := client.SaveClassificationConditionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FillLogisticsWithOptions(request *FillLogisticsRequest, runtime *util.RuntimeOptions) (_result *FillLogisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FillLogisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FillLogistics"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FillLogistics(request *FillLogisticsRequest) (_result *FillLogisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FillLogisticsResponse{}
	_body, _err := client.FillLogisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMaterialWithOptions(request *UpdateMaterialRequest, runtime *util.RuntimeOptions) (_result *UpdateMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMaterial"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMaterial(request *UpdateMaterialRequest) (_result *UpdateMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMaterialResponse{}
	_body, _err := client.UpdateMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTradeMarkApplicationLogsWithOptions(request *QueryTradeMarkApplicationLogsRequest, runtime *util.RuntimeOptions) (_result *QueryTradeMarkApplicationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTradeMarkApplicationLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTradeMarkApplicationLogs"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTradeMarkApplicationLogs(request *QueryTradeMarkApplicationLogsRequest) (_result *QueryTradeMarkApplicationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTradeMarkApplicationLogsResponse{}
	_body, _err := client.QueryTradeMarkApplicationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundProduceWithOptions(request *RefundProduceRequest, runtime *util.RuntimeOptions) (_result *RefundProduceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefundProduceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefundProduce"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundProduce(request *RefundProduceRequest) (_result *RefundProduceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundProduceResponse{}
	_body, _err := client.RefundProduceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncTrademarkWithOptions(request *SyncTrademarkRequest, runtime *util.RuntimeOptions) (_result *SyncTrademarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SyncTrademarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SyncTrademark"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncTrademark(request *SyncTrademarkRequest) (_result *SyncTrademarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncTrademarkResponse{}
	_body, _err := client.SyncTrademarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CombineLoaWithOptions(request *CombineLoaRequest, runtime *util.RuntimeOptions) (_result *CombineLoaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CombineLoaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CombineLoa"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CombineLoa(request *CombineLoaRequest) (_result *CombineLoaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CombineLoaResponse{}
	_body, _err := client.CombineLoaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FilterUnavailableCodesWithOptions(tmpReq *FilterUnavailableCodesRequest, runtime *util.RuntimeOptions) (_result *FilterUnavailableCodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FilterUnavailableCodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Codes)) {
		request.CodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Codes, tea.String("Codes"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FilterUnavailableCodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FilterUnavailableCodes"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FilterUnavailableCodes(request *FilterUnavailableCodesRequest) (_result *FilterUnavailableCodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FilterUnavailableCodesResponse{}
	_body, _err := client.FilterUnavailableCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertMaterialWithOptions(request *InsertMaterialRequest, runtime *util.RuntimeOptions) (_result *InsertMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertMaterialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertMaterial"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertMaterial(request *InsertMaterialRequest) (_result *InsertMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertMaterialResponse{}
	_body, _err := client.InsertMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTradeMarkReviewMaterialDetailWithOptions(tmpReq *SaveTradeMarkReviewMaterialDetailRequest, runtime *util.RuntimeOptions) (_result *SaveTradeMarkReviewMaterialDetailResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SaveTradeMarkReviewMaterialDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AdditionalOssKeyList)) {
		request.AdditionalOssKeyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalOssKeyList, tea.String("AdditionalOssKeyList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveTradeMarkReviewMaterialDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveTradeMarkReviewMaterialDetail"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTradeMarkReviewMaterialDetail(request *SaveTradeMarkReviewMaterialDetailRequest) (_result *SaveTradeMarkReviewMaterialDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTradeMarkReviewMaterialDetailResponse{}
	_body, _err := client.SaveTradeMarkReviewMaterialDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMonitorKeywordsWithOptions(request *QueryMonitorKeywordsRequest, runtime *util.RuntimeOptions) (_result *QueryMonitorKeywordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMonitorKeywordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMonitorKeywords"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMonitorKeywords(request *QueryMonitorKeywordsRequest) (_result *QueryMonitorKeywordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMonitorKeywordsResponse{}
	_body, _err := client.QueryMonitorKeywordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTaskListWithOptions(request *QueryTaskListRequest, runtime *util.RuntimeOptions) (_result *QueryTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTaskListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTaskList"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTaskList(request *QueryTaskListRequest) (_result *QueryTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTaskListResponse{}
	_body, _err := client.QueryTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTrademarkNameWithOptions(request *UpdateTrademarkNameRequest, runtime *util.RuntimeOptions) (_result *UpdateTrademarkNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTrademarkNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTrademarkName"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrademarkName(request *UpdateTrademarkNameRequest) (_result *UpdateTrademarkNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTrademarkNameResponse{}
	_body, _err := client.UpdateTrademarkNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckLoaFillWithOptions(request *CheckLoaFillRequest, runtime *util.RuntimeOptions) (_result *CheckLoaFillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckLoaFillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckLoaFill"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckLoaFill(request *CheckLoaFillRequest) (_result *CheckLoaFillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckLoaFillResponse{}
	_body, _err := client.CheckLoaFillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmApplicantWithOptions(request *ConfirmApplicantRequest, runtime *util.RuntimeOptions) (_result *ConfirmApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfirmApplicantResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfirmApplicant"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmApplicant(request *ConfirmApplicantRequest) (_result *ConfirmApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmApplicantResponse{}
	_body, _err := client.ConfirmApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIntentionOrderWithOptions(request *CreateIntentionOrderRequest, runtime *util.RuntimeOptions) (_result *CreateIntentionOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIntentionOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIntentionOrder"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIntentionOrder(request *CreateIntentionOrderRequest) (_result *CreateIntentionOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIntentionOrderResponse{}
	_body, _err := client.CreateIntentionOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOssResourcesWithOptions(request *QueryOssResourcesRequest, runtime *util.RuntimeOptions) (_result *QueryOssResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOssResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOssResources"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOssResources(request *QueryOssResourcesRequest) (_result *QueryOssResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOssResourcesResponse{}
	_body, _err := client.QueryOssResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefuseApplicantWithOptions(request *RefuseApplicantRequest, runtime *util.RuntimeOptions) (_result *RefuseApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefuseApplicantResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefuseApplicant"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefuseApplicant(request *RefuseApplicantRequest) (_result *RefuseApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefuseApplicantResponse{}
	_body, _err := client.RefuseApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIntentionOrderGeneratingPayWithOptions(request *CreateIntentionOrderGeneratingPayRequest, runtime *util.RuntimeOptions) (_result *CreateIntentionOrderGeneratingPayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIntentionOrderGeneratingPayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIntentionOrderGeneratingPay"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIntentionOrderGeneratingPay(request *CreateIntentionOrderGeneratingPayRequest) (_result *CreateIntentionOrderGeneratingPayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIntentionOrderGeneratingPayResponse{}
	_body, _err := client.CreateIntentionOrderGeneratingPayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNotaryOrdersWithOptions(request *ListNotaryOrdersRequest, runtime *util.RuntimeOptions) (_result *ListNotaryOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListNotaryOrdersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNotaryOrders"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNotaryOrders(request *ListNotaryOrdersRequest) (_result *ListNotaryOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNotaryOrdersResponse{}
	_body, _err := client.ListNotaryOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSupportPrincipalNameWithOptions(runtime *util.RuntimeOptions) (_result *GetSupportPrincipalNameResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &GetSupportPrincipalNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSupportPrincipalName"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSupportPrincipalName() (_result *GetSupportPrincipalNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSupportPrincipalNameResponse{}
	_body, _err := client.GetSupportPrincipalNameWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartNotaryWithOptions(request *StartNotaryRequest, runtime *util.RuntimeOptions) (_result *StartNotaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartNotaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartNotary"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartNotary(request *StartNotaryRequest) (_result *StartNotaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartNotaryResponse{}
	_body, _err := client.StartNotaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSendMaterialNumWithOptions(request *UpdateSendMaterialNumRequest, runtime *util.RuntimeOptions) (_result *UpdateSendMaterialNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSendMaterialNumResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSendMaterialNum"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSendMaterialNum(request *UpdateSendMaterialNumRequest) (_result *UpdateSendMaterialNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSendMaterialNumResponse{}
	_body, _err := client.UpdateSendMaterialNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTrademarkApplicationWithOptions(request *DeleteTrademarkApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteTrademarkApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTrademarkApplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTrademarkApplication"), tea.String("2018-07-24"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrademarkApplication(request *DeleteTrademarkApplicationRequest) (_result *DeleteTrademarkApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTrademarkApplicationResponse{}
	_body, _err := client.DeleteTrademarkApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
