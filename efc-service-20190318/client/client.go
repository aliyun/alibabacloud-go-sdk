// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckAllMasterTrusteeshipRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserId    *int64  `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
}

func (s CheckAllMasterTrusteeshipRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAllMasterTrusteeshipRequest) GoString() string {
	return s.String()
}

func (s *CheckAllMasterTrusteeshipRequest) SetRequestId(v string) *CheckAllMasterTrusteeshipRequest {
	s.RequestId = &v
	return s
}

func (s *CheckAllMasterTrusteeshipRequest) SetUserId(v int64) *CheckAllMasterTrusteeshipRequest {
	s.UserId = &v
	return s
}

type CheckAllMasterTrusteeshipResponse struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *CheckAllMasterTrusteeshipResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CheckAllMasterTrusteeshipResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAllMasterTrusteeshipResponse) GoString() string {
	return s.String()
}

func (s *CheckAllMasterTrusteeshipResponse) SetRequestId(v string) *CheckAllMasterTrusteeshipResponse {
	s.RequestId = &v
	return s
}

func (s *CheckAllMasterTrusteeshipResponse) SetSuccess(v bool) *CheckAllMasterTrusteeshipResponse {
	s.Success = &v
	return s
}

func (s *CheckAllMasterTrusteeshipResponse) SetCode(v string) *CheckAllMasterTrusteeshipResponse {
	s.Code = &v
	return s
}

func (s *CheckAllMasterTrusteeshipResponse) SetMessage(v string) *CheckAllMasterTrusteeshipResponse {
	s.Message = &v
	return s
}

func (s *CheckAllMasterTrusteeshipResponse) SetData(v *CheckAllMasterTrusteeshipResponseData) *CheckAllMasterTrusteeshipResponse {
	s.Data = v
	return s
}

type CheckAllMasterTrusteeshipResponseData struct {
	CheckResult   *bool                                                 `json:"CheckResult,omitempty" xml:"CheckResult,omitempty" require:"true"`
	CheckDataList []*CheckAllMasterTrusteeshipResponseDataCheckDataList `json:"CheckDataList,omitempty" xml:"CheckDataList,omitempty" require:"true" type:"Repeated"`
}

func (s CheckAllMasterTrusteeshipResponseData) String() string {
	return tea.Prettify(s)
}

func (s CheckAllMasterTrusteeshipResponseData) GoString() string {
	return s.String()
}

func (s *CheckAllMasterTrusteeshipResponseData) SetCheckResult(v bool) *CheckAllMasterTrusteeshipResponseData {
	s.CheckResult = &v
	return s
}

func (s *CheckAllMasterTrusteeshipResponseData) SetCheckDataList(v []*CheckAllMasterTrusteeshipResponseDataCheckDataList) *CheckAllMasterTrusteeshipResponseData {
	s.CheckDataList = v
	return s
}

type CheckAllMasterTrusteeshipResponseDataCheckDataList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s CheckAllMasterTrusteeshipResponseDataCheckDataList) String() string {
	return tea.Prettify(s)
}

func (s CheckAllMasterTrusteeshipResponseDataCheckDataList) GoString() string {
	return s.String()
}

func (s *CheckAllMasterTrusteeshipResponseDataCheckDataList) SetCode(v string) *CheckAllMasterTrusteeshipResponseDataCheckDataList {
	s.Code = &v
	return s
}

func (s *CheckAllMasterTrusteeshipResponseDataCheckDataList) SetMessage(v string) *CheckAllMasterTrusteeshipResponseDataCheckDataList {
	s.Message = &v
	return s
}

type QueryReservedInstanceSharedInfosRequest struct {
	Uid         *int64                                            `json:"Uid,omitempty" xml:"Uid,omitempty" require:"true"`
	RiInfos     []*QueryReservedInstanceSharedInfosRequestRiInfos `json:"RiInfos,omitempty" xml:"RiInfos,omitempty" type:"Repeated"`
	PageNo      *int                                              `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	AccountType *string                                           `json:"AccountType,omitempty" xml:"AccountType,omitempty" require:"true"`
	PageSize    *int                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Region      *string                                           `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryReservedInstanceSharedInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReservedInstanceSharedInfosRequest) GoString() string {
	return s.String()
}

func (s *QueryReservedInstanceSharedInfosRequest) SetUid(v int64) *QueryReservedInstanceSharedInfosRequest {
	s.Uid = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosRequest) SetRiInfos(v []*QueryReservedInstanceSharedInfosRequestRiInfos) *QueryReservedInstanceSharedInfosRequest {
	s.RiInfos = v
	return s
}

func (s *QueryReservedInstanceSharedInfosRequest) SetPageNo(v int) *QueryReservedInstanceSharedInfosRequest {
	s.PageNo = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosRequest) SetAccountType(v string) *QueryReservedInstanceSharedInfosRequest {
	s.AccountType = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosRequest) SetPageSize(v int) *QueryReservedInstanceSharedInfosRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosRequest) SetRegion(v string) *QueryReservedInstanceSharedInfosRequest {
	s.Region = &v
	return s
}

type QueryReservedInstanceSharedInfosRequestRiInfos struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty" require:"true"`
	RiInstanceId  *string `json:"RiInstanceId,omitempty" xml:"RiInstanceId,omitempty" require:"true"`
}

func (s QueryReservedInstanceSharedInfosRequestRiInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryReservedInstanceSharedInfosRequestRiInfos) GoString() string {
	return s.String()
}

func (s *QueryReservedInstanceSharedInfosRequestRiInfos) SetCommodityCode(v string) *QueryReservedInstanceSharedInfosRequestRiInfos {
	s.CommodityCode = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosRequestRiInfos) SetRiInstanceId(v string) *QueryReservedInstanceSharedInfosRequestRiInfos {
	s.RiInstanceId = &v
	return s
}

type QueryReservedInstanceSharedInfosResponse struct {
	Code        *string                                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Count       *int                                            `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	PageSize    *int                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage *int                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	TotalCount  *int                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Message     *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Region      *string                                         `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data        []*QueryReservedInstanceSharedInfosResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryReservedInstanceSharedInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReservedInstanceSharedInfosResponse) GoString() string {
	return s.String()
}

func (s *QueryReservedInstanceSharedInfosResponse) SetCode(v string) *QueryReservedInstanceSharedInfosResponse {
	s.Code = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetSuccess(v bool) *QueryReservedInstanceSharedInfosResponse {
	s.Success = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetCount(v int) *QueryReservedInstanceSharedInfosResponse {
	s.Count = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetPageSize(v int) *QueryReservedInstanceSharedInfosResponse {
	s.PageSize = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetCurrentPage(v int) *QueryReservedInstanceSharedInfosResponse {
	s.CurrentPage = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetTotalCount(v int) *QueryReservedInstanceSharedInfosResponse {
	s.TotalCount = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetMessage(v string) *QueryReservedInstanceSharedInfosResponse {
	s.Message = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetRegion(v string) *QueryReservedInstanceSharedInfosResponse {
	s.Region = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetRequestId(v string) *QueryReservedInstanceSharedInfosResponse {
	s.RequestId = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponse) SetData(v []*QueryReservedInstanceSharedInfosResponseData) *QueryReservedInstanceSharedInfosResponse {
	s.Data = v
	return s
}

type QueryReservedInstanceSharedInfosResponseData struct {
	MainAccountPk *int64  `json:"MainAccountPk,omitempty" xml:"MainAccountPk,omitempty" require:"true"`
	SubAccountPk  *int64  `json:"SubAccountPk,omitempty" xml:"SubAccountPk,omitempty" require:"true"`
	RelationType  *string `json:"RelationType,omitempty" xml:"RelationType,omitempty" require:"true"`
	RiInstanceId  *string `json:"RiInstanceId,omitempty" xml:"RiInstanceId,omitempty" require:"true"`
	EffectTime    *int64  `json:"EffectTime,omitempty" xml:"EffectTime,omitempty" require:"true"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
}

func (s QueryReservedInstanceSharedInfosResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryReservedInstanceSharedInfosResponseData) GoString() string {
	return s.String()
}

func (s *QueryReservedInstanceSharedInfosResponseData) SetMainAccountPk(v int64) *QueryReservedInstanceSharedInfosResponseData {
	s.MainAccountPk = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponseData) SetSubAccountPk(v int64) *QueryReservedInstanceSharedInfosResponseData {
	s.SubAccountPk = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponseData) SetRelationType(v string) *QueryReservedInstanceSharedInfosResponseData {
	s.RelationType = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponseData) SetRiInstanceId(v string) *QueryReservedInstanceSharedInfosResponseData {
	s.RiInstanceId = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponseData) SetEffectTime(v int64) *QueryReservedInstanceSharedInfosResponseData {
	s.EffectTime = &v
	return s
}

func (s *QueryReservedInstanceSharedInfosResponseData) SetRegion(v string) *QueryReservedInstanceSharedInfosResponseData {
	s.Region = &v
	return s
}

type QueryRdTrusteeshiperRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserId    *int64  `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
}

func (s QueryRdTrusteeshiperRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRdTrusteeshiperRequest) GoString() string {
	return s.String()
}

func (s *QueryRdTrusteeshiperRequest) SetRequestId(v string) *QueryRdTrusteeshiperRequest {
	s.RequestId = &v
	return s
}

func (s *QueryRdTrusteeshiperRequest) SetUserId(v int64) *QueryRdTrusteeshiperRequest {
	s.UserId = &v
	return s
}

type QueryRdTrusteeshiperResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *QueryRdTrusteeshiperResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s QueryRdTrusteeshiperResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRdTrusteeshiperResponse) GoString() string {
	return s.String()
}

func (s *QueryRdTrusteeshiperResponse) SetRequestId(v string) *QueryRdTrusteeshiperResponse {
	s.RequestId = &v
	return s
}

func (s *QueryRdTrusteeshiperResponse) SetSuccess(v bool) *QueryRdTrusteeshiperResponse {
	s.Success = &v
	return s
}

func (s *QueryRdTrusteeshiperResponse) SetCode(v string) *QueryRdTrusteeshiperResponse {
	s.Code = &v
	return s
}

func (s *QueryRdTrusteeshiperResponse) SetMessage(v string) *QueryRdTrusteeshiperResponse {
	s.Message = &v
	return s
}

func (s *QueryRdTrusteeshiperResponse) SetData(v *QueryRdTrusteeshiperResponseData) *QueryRdTrusteeshiperResponse {
	s.Data = v
	return s
}

type QueryRdTrusteeshiperResponseData struct {
	TrusteeshipUserId   *int64  `json:"TrusteeshipUserId,omitempty" xml:"TrusteeshipUserId,omitempty" require:"true"`
	TrusteeshipUserName *string `json:"TrusteeshipUserName,omitempty" xml:"TrusteeshipUserName,omitempty" require:"true"`
}

func (s QueryRdTrusteeshiperResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryRdTrusteeshiperResponseData) GoString() string {
	return s.String()
}

func (s *QueryRdTrusteeshiperResponseData) SetTrusteeshipUserId(v int64) *QueryRdTrusteeshiperResponseData {
	s.TrusteeshipUserId = &v
	return s
}

func (s *QueryRdTrusteeshiperResponseData) SetTrusteeshipUserName(v string) *QueryRdTrusteeshiperResponseData {
	s.TrusteeshipUserName = &v
	return s
}

type CheckMasterTrusteeshipRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserId    *int64  `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
}

func (s CheckMasterTrusteeshipRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMasterTrusteeshipRequest) GoString() string {
	return s.String()
}

func (s *CheckMasterTrusteeshipRequest) SetRequestId(v string) *CheckMasterTrusteeshipRequest {
	s.RequestId = &v
	return s
}

func (s *CheckMasterTrusteeshipRequest) SetUserId(v int64) *CheckMasterTrusteeshipRequest {
	s.UserId = &v
	return s
}

type CheckMasterTrusteeshipResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *CheckMasterTrusteeshipResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CheckMasterTrusteeshipResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMasterTrusteeshipResponse) GoString() string {
	return s.String()
}

func (s *CheckMasterTrusteeshipResponse) SetRequestId(v string) *CheckMasterTrusteeshipResponse {
	s.RequestId = &v
	return s
}

func (s *CheckMasterTrusteeshipResponse) SetSuccess(v bool) *CheckMasterTrusteeshipResponse {
	s.Success = &v
	return s
}

func (s *CheckMasterTrusteeshipResponse) SetCode(v string) *CheckMasterTrusteeshipResponse {
	s.Code = &v
	return s
}

func (s *CheckMasterTrusteeshipResponse) SetMessage(v string) *CheckMasterTrusteeshipResponse {
	s.Message = &v
	return s
}

func (s *CheckMasterTrusteeshipResponse) SetData(v *CheckMasterTrusteeshipResponseData) *CheckMasterTrusteeshipResponse {
	s.Data = v
	return s
}

type CheckMasterTrusteeshipResponseData struct {
	CheckResult *bool   `json:"CheckResult,omitempty" xml:"CheckResult,omitempty" require:"true"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s CheckMasterTrusteeshipResponseData) String() string {
	return tea.Prettify(s)
}

func (s CheckMasterTrusteeshipResponseData) GoString() string {
	return s.String()
}

func (s *CheckMasterTrusteeshipResponseData) SetCheckResult(v bool) *CheckMasterTrusteeshipResponseData {
	s.CheckResult = &v
	return s
}

func (s *CheckMasterTrusteeshipResponseData) SetCode(v string) *CheckMasterTrusteeshipResponseData {
	s.Code = &v
	return s
}

func (s *CheckMasterTrusteeshipResponseData) SetMessage(v string) *CheckMasterTrusteeshipResponseData {
	s.Message = &v
	return s
}

type VerifyRdOperationRequest struct {
	RdOperation *string `json:"RdOperation,omitempty" xml:"RdOperation,omitempty"`
}

func (s VerifyRdOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyRdOperationRequest) GoString() string {
	return s.String()
}

func (s *VerifyRdOperationRequest) SetRdOperation(v string) *VerifyRdOperationRequest {
	s.RdOperation = &v
	return s
}

type VerifyRdOperationResponse struct {
	RequestId             *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RdOprationCheckResult *VerifyRdOperationResponseRdOprationCheckResult `json:"RdOprationCheckResult,omitempty" xml:"RdOprationCheckResult,omitempty" require:"true" type:"Struct"`
}

func (s VerifyRdOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyRdOperationResponse) GoString() string {
	return s.String()
}

func (s *VerifyRdOperationResponse) SetRequestId(v string) *VerifyRdOperationResponse {
	s.RequestId = &v
	return s
}

func (s *VerifyRdOperationResponse) SetRdOprationCheckResult(v *VerifyRdOperationResponseRdOprationCheckResult) *VerifyRdOperationResponse {
	s.RdOprationCheckResult = v
	return s
}

type VerifyRdOperationResponseRdOprationCheckResult struct {
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty" require:"true"`
	FailCode    *string `json:"FailCode,omitempty" xml:"FailCode,omitempty" require:"true"`
	FailMessage *string `json:"FailMessage,omitempty" xml:"FailMessage,omitempty" require:"true"`
	ExtraData   *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty" require:"true"`
}

func (s VerifyRdOperationResponseRdOprationCheckResult) String() string {
	return tea.Prettify(s)
}

func (s VerifyRdOperationResponseRdOprationCheckResult) GoString() string {
	return s.String()
}

func (s *VerifyRdOperationResponseRdOprationCheckResult) SetSuccess(v bool) *VerifyRdOperationResponseRdOprationCheckResult {
	s.Success = &v
	return s
}

func (s *VerifyRdOperationResponseRdOprationCheckResult) SetBizType(v string) *VerifyRdOperationResponseRdOprationCheckResult {
	s.BizType = &v
	return s
}

func (s *VerifyRdOperationResponseRdOprationCheckResult) SetFailCode(v string) *VerifyRdOperationResponseRdOprationCheckResult {
	s.FailCode = &v
	return s
}

func (s *VerifyRdOperationResponseRdOprationCheckResult) SetFailMessage(v string) *VerifyRdOperationResponseRdOprationCheckResult {
	s.FailMessage = &v
	return s
}

func (s *VerifyRdOperationResponseRdOprationCheckResult) SetExtraData(v string) *VerifyRdOperationResponseRdOprationCheckResult {
	s.ExtraData = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("efc-service"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) CheckAllMasterTrusteeshipWithOptions(request *CheckAllMasterTrusteeshipRequest, runtime *util.RuntimeOptions) (_result *CheckAllMasterTrusteeshipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckAllMasterTrusteeshipResponse{}
	_body, _err := client.DoRequest(tea.String("CheckAllMasterTrusteeship"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAllMasterTrusteeship(request *CheckAllMasterTrusteeshipRequest) (_result *CheckAllMasterTrusteeshipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAllMasterTrusteeshipResponse{}
	_body, _err := client.CheckAllMasterTrusteeshipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryReservedInstanceSharedInfosWithOptions(request *QueryReservedInstanceSharedInfosRequest, runtime *util.RuntimeOptions) (_result *QueryReservedInstanceSharedInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryReservedInstanceSharedInfosResponse{}
	_body, _err := client.DoRequest(tea.String("QueryReservedInstanceSharedInfos"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-03-18"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryReservedInstanceSharedInfos(request *QueryReservedInstanceSharedInfosRequest) (_result *QueryReservedInstanceSharedInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryReservedInstanceSharedInfosResponse{}
	_body, _err := client.QueryReservedInstanceSharedInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRdTrusteeshiperWithOptions(request *QueryRdTrusteeshiperRequest, runtime *util.RuntimeOptions) (_result *QueryRdTrusteeshiperResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryRdTrusteeshiperResponse{}
	_body, _err := client.DoRequest(tea.String("QueryRdTrusteeshiper"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRdTrusteeshiper(request *QueryRdTrusteeshiperRequest) (_result *QueryRdTrusteeshiperResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRdTrusteeshiperResponse{}
	_body, _err := client.QueryRdTrusteeshiperWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckMasterTrusteeshipWithOptions(request *CheckMasterTrusteeshipRequest, runtime *util.RuntimeOptions) (_result *CheckMasterTrusteeshipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckMasterTrusteeshipResponse{}
	_body, _err := client.DoRequest(tea.String("CheckMasterTrusteeship"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMasterTrusteeship(request *CheckMasterTrusteeshipRequest) (_result *CheckMasterTrusteeshipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMasterTrusteeshipResponse{}
	_body, _err := client.CheckMasterTrusteeshipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyRdOperationWithOptions(request *VerifyRdOperationRequest, runtime *util.RuntimeOptions) (_result *VerifyRdOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &VerifyRdOperationResponse{}
	_body, _err := client.DoRequest(tea.String("VerifyRdOperation"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-03-18"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyRdOperation(request *VerifyRdOperationRequest) (_result *VerifyRdOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyRdOperationResponse{}
	_body, _err := client.VerifyRdOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
