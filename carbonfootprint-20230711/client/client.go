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

type AllowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllowResponseBody) GoString() string {
	return s.String()
}

func (s *AllowResponseBody) SetData(v bool) *AllowResponseBody {
	s.Data = &v
	return s
}

func (s *AllowResponseBody) SetRequestId(v string) *AllowResponseBody {
	s.RequestId = &v
	return s
}

type AllowResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllowResponse) String() string {
	return tea.Prettify(s)
}

func (s AllowResponse) GoString() string {
	return s.String()
}

func (s *AllowResponse) SetHeaders(v map[string]*string) *AllowResponse {
	s.Headers = v
	return s
}

func (s *AllowResponse) SetStatusCode(v int32) *AllowResponse {
	s.StatusCode = &v
	return s
}

func (s *AllowResponse) SetBody(v *AllowResponseBody) *AllowResponse {
	s.Body = v
	return s
}

type GetSummaryDataRequest struct {
	EndTime   *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Group     *string   `json:"Group,omitempty" xml:"Group,omitempty"`
	StartTime *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Uids      []*string `json:"Uids,omitempty" xml:"Uids,omitempty" type:"Repeated"`
}

func (s GetSummaryDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryDataRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryDataRequest) SetEndTime(v string) *GetSummaryDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetSummaryDataRequest) SetGroup(v string) *GetSummaryDataRequest {
	s.Group = &v
	return s
}

func (s *GetSummaryDataRequest) SetStartTime(v string) *GetSummaryDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetSummaryDataRequest) SetUids(v []*string) *GetSummaryDataRequest {
	s.Uids = v
	return s
}

type GetSummaryDataShrinkRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Group      *string `json:"Group,omitempty" xml:"Group,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UidsShrink *string `json:"Uids,omitempty" xml:"Uids,omitempty"`
}

func (s GetSummaryDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryDataShrinkRequest) SetEndTime(v string) *GetSummaryDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetSummaryDataShrinkRequest) SetGroup(v string) *GetSummaryDataShrinkRequest {
	s.Group = &v
	return s
}

func (s *GetSummaryDataShrinkRequest) SetStartTime(v string) *GetSummaryDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetSummaryDataShrinkRequest) SetUidsShrink(v string) *GetSummaryDataShrinkRequest {
	s.UidsShrink = &v
	return s
}

type GetSummaryDataResponseBody struct {
	Data *GetSummaryDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSummaryDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetSummaryDataResponseBody) SetData(v *GetSummaryDataResponseBodyData) *GetSummaryDataResponseBody {
	s.Data = v
	return s
}

func (s *GetSummaryDataResponseBody) SetRequestId(v string) *GetSummaryDataResponseBody {
	s.RequestId = &v
	return s
}

type GetSummaryDataResponseBodyData struct {
	AircraftConsumptionConversion    *string `json:"AircraftConsumptionConversion,omitempty" xml:"AircraftConsumptionConversion,omitempty"`
	CarConsumptionConversion         *string `json:"CarConsumptionConversion,omitempty" xml:"CarConsumptionConversion,omitempty"`
	LastMonthConsumptionConversion   *string `json:"LastMonthConsumptionConversion,omitempty" xml:"LastMonthConsumptionConversion,omitempty"`
	LastYearConsumptionConversion    *string `json:"LastYearConsumptionConversion,omitempty" xml:"LastYearConsumptionConversion,omitempty"`
	LastYearConsumptionConversionSum *string `json:"LastYearConsumptionConversionSum,omitempty" xml:"LastYearConsumptionConversionSum,omitempty"`
	LatestDataTime                   *string `json:"LatestDataTime,omitempty" xml:"LatestDataTime,omitempty"`
	ThisMonthConsumptionConversion   *string `json:"ThisMonthConsumptionConversion,omitempty" xml:"ThisMonthConsumptionConversion,omitempty"`
	ThisYearConsumptionConversion    *string `json:"ThisYearConsumptionConversion,omitempty" xml:"ThisYearConsumptionConversion,omitempty"`
	TotalCarbonConsumptionConversion *string `json:"TotalCarbonConsumptionConversion,omitempty" xml:"TotalCarbonConsumptionConversion,omitempty"`
	TreeConsumptionConversion        *string `json:"TreeConsumptionConversion,omitempty" xml:"TreeConsumptionConversion,omitempty"`
}

func (s GetSummaryDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSummaryDataResponseBodyData) SetAircraftConsumptionConversion(v string) *GetSummaryDataResponseBodyData {
	s.AircraftConsumptionConversion = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetCarConsumptionConversion(v string) *GetSummaryDataResponseBodyData {
	s.CarConsumptionConversion = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetLastMonthConsumptionConversion(v string) *GetSummaryDataResponseBodyData {
	s.LastMonthConsumptionConversion = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetLastYearConsumptionConversion(v string) *GetSummaryDataResponseBodyData {
	s.LastYearConsumptionConversion = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetLastYearConsumptionConversionSum(v string) *GetSummaryDataResponseBodyData {
	s.LastYearConsumptionConversionSum = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetLatestDataTime(v string) *GetSummaryDataResponseBodyData {
	s.LatestDataTime = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetThisMonthConsumptionConversion(v string) *GetSummaryDataResponseBodyData {
	s.ThisMonthConsumptionConversion = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetThisYearConsumptionConversion(v string) *GetSummaryDataResponseBodyData {
	s.ThisYearConsumptionConversion = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetTotalCarbonConsumptionConversion(v string) *GetSummaryDataResponseBodyData {
	s.TotalCarbonConsumptionConversion = &v
	return s
}

func (s *GetSummaryDataResponseBodyData) SetTreeConsumptionConversion(v string) *GetSummaryDataResponseBodyData {
	s.TreeConsumptionConversion = &v
	return s
}

type GetSummaryDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSummaryDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSummaryDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryDataResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryDataResponse) SetHeaders(v map[string]*string) *GetSummaryDataResponse {
	s.Headers = v
	return s
}

func (s *GetSummaryDataResponse) SetStatusCode(v int32) *GetSummaryDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSummaryDataResponse) SetBody(v *GetSummaryDataResponseBody) *GetSummaryDataResponse {
	s.Body = v
	return s
}

type QueryCarbonTrackRequest struct {
	EndTime         *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterRDAccount *int32    `json:"FilterRDAccount,omitempty" xml:"FilterRDAccount,omitempty"`
	Group           *string   `json:"Group,omitempty" xml:"Group,omitempty"`
	StartTime       *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopNum          *int32    `json:"TopNum,omitempty" xml:"TopNum,omitempty"`
	Uids            []*string `json:"Uids,omitempty" xml:"Uids,omitempty" type:"Repeated"`
}

func (s QueryCarbonTrackRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCarbonTrackRequest) GoString() string {
	return s.String()
}

func (s *QueryCarbonTrackRequest) SetEndTime(v string) *QueryCarbonTrackRequest {
	s.EndTime = &v
	return s
}

func (s *QueryCarbonTrackRequest) SetFilterRDAccount(v int32) *QueryCarbonTrackRequest {
	s.FilterRDAccount = &v
	return s
}

func (s *QueryCarbonTrackRequest) SetGroup(v string) *QueryCarbonTrackRequest {
	s.Group = &v
	return s
}

func (s *QueryCarbonTrackRequest) SetStartTime(v string) *QueryCarbonTrackRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCarbonTrackRequest) SetTopNum(v int32) *QueryCarbonTrackRequest {
	s.TopNum = &v
	return s
}

func (s *QueryCarbonTrackRequest) SetUids(v []*string) *QueryCarbonTrackRequest {
	s.Uids = v
	return s
}

type QueryCarbonTrackShrinkRequest struct {
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterRDAccount *int32  `json:"FilterRDAccount,omitempty" xml:"FilterRDAccount,omitempty"`
	Group           *string `json:"Group,omitempty" xml:"Group,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopNum          *int32  `json:"TopNum,omitempty" xml:"TopNum,omitempty"`
	UidsShrink      *string `json:"Uids,omitempty" xml:"Uids,omitempty"`
}

func (s QueryCarbonTrackShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCarbonTrackShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCarbonTrackShrinkRequest) SetEndTime(v string) *QueryCarbonTrackShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *QueryCarbonTrackShrinkRequest) SetFilterRDAccount(v int32) *QueryCarbonTrackShrinkRequest {
	s.FilterRDAccount = &v
	return s
}

func (s *QueryCarbonTrackShrinkRequest) SetGroup(v string) *QueryCarbonTrackShrinkRequest {
	s.Group = &v
	return s
}

func (s *QueryCarbonTrackShrinkRequest) SetStartTime(v string) *QueryCarbonTrackShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCarbonTrackShrinkRequest) SetTopNum(v int32) *QueryCarbonTrackShrinkRequest {
	s.TopNum = &v
	return s
}

func (s *QueryCarbonTrackShrinkRequest) SetUidsShrink(v string) *QueryCarbonTrackShrinkRequest {
	s.UidsShrink = &v
	return s
}

type QueryCarbonTrackResponseBody struct {
	Data      []*QueryCarbonTrackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCarbonTrackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCarbonTrackResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCarbonTrackResponseBody) SetData(v []*QueryCarbonTrackResponseBodyData) *QueryCarbonTrackResponseBody {
	s.Data = v
	return s
}

func (s *QueryCarbonTrackResponseBody) SetRequestId(v string) *QueryCarbonTrackResponseBody {
	s.RequestId = &v
	return s
}

type QueryCarbonTrackResponseBodyData struct {
	ProductCode    *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaValue     *float64 `json:"QuotaValue,omitempty" xml:"QuotaValue,omitempty"`
	Region         *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	StatisticsDate *int64   `json:"StatisticsDate,omitempty" xml:"StatisticsDate,omitempty"`
	SubUid         *string  `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
	Uid            *string  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryCarbonTrackResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCarbonTrackResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCarbonTrackResponseBodyData) SetProductCode(v string) *QueryCarbonTrackResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *QueryCarbonTrackResponseBodyData) SetQuotaValue(v float64) *QueryCarbonTrackResponseBodyData {
	s.QuotaValue = &v
	return s
}

func (s *QueryCarbonTrackResponseBodyData) SetRegion(v string) *QueryCarbonTrackResponseBodyData {
	s.Region = &v
	return s
}

func (s *QueryCarbonTrackResponseBodyData) SetStatisticsDate(v int64) *QueryCarbonTrackResponseBodyData {
	s.StatisticsDate = &v
	return s
}

func (s *QueryCarbonTrackResponseBodyData) SetSubUid(v string) *QueryCarbonTrackResponseBodyData {
	s.SubUid = &v
	return s
}

func (s *QueryCarbonTrackResponseBodyData) SetUid(v string) *QueryCarbonTrackResponseBodyData {
	s.Uid = &v
	return s
}

type QueryCarbonTrackResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCarbonTrackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCarbonTrackResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCarbonTrackResponse) GoString() string {
	return s.String()
}

func (s *QueryCarbonTrackResponse) SetHeaders(v map[string]*string) *QueryCarbonTrackResponse {
	s.Headers = v
	return s
}

func (s *QueryCarbonTrackResponse) SetStatusCode(v int32) *QueryCarbonTrackResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCarbonTrackResponse) SetBody(v *QueryCarbonTrackResponseBody) *QueryCarbonTrackResponse {
	s.Body = v
	return s
}

type QueryMultiAccountCarbonTrackRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMultiAccountCarbonTrackRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMultiAccountCarbonTrackRequest) GoString() string {
	return s.String()
}

func (s *QueryMultiAccountCarbonTrackRequest) SetEndTime(v string) *QueryMultiAccountCarbonTrackRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMultiAccountCarbonTrackRequest) SetStartTime(v string) *QueryMultiAccountCarbonTrackRequest {
	s.StartTime = &v
	return s
}

type QueryMultiAccountCarbonTrackResponseBody struct {
	Data []*QueryMultiAccountCarbonTrackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMultiAccountCarbonTrackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMultiAccountCarbonTrackResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMultiAccountCarbonTrackResponseBody) SetData(v []*QueryMultiAccountCarbonTrackResponseBodyData) *QueryMultiAccountCarbonTrackResponseBody {
	s.Data = v
	return s
}

func (s *QueryMultiAccountCarbonTrackResponseBody) SetRequestId(v string) *QueryMultiAccountCarbonTrackResponseBody {
	s.RequestId = &v
	return s
}

type QueryMultiAccountCarbonTrackResponseBodyData struct {
	CarbonActualEmission *string `json:"CarbonActualEmission,omitempty" xml:"CarbonActualEmission,omitempty"`
	Month                *string `json:"Month,omitempty" xml:"Month,omitempty"`
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Uid                  *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryMultiAccountCarbonTrackResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMultiAccountCarbonTrackResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMultiAccountCarbonTrackResponseBodyData) SetCarbonActualEmission(v string) *QueryMultiAccountCarbonTrackResponseBodyData {
	s.CarbonActualEmission = &v
	return s
}

func (s *QueryMultiAccountCarbonTrackResponseBodyData) SetMonth(v string) *QueryMultiAccountCarbonTrackResponseBodyData {
	s.Month = &v
	return s
}

func (s *QueryMultiAccountCarbonTrackResponseBodyData) SetProductCode(v string) *QueryMultiAccountCarbonTrackResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *QueryMultiAccountCarbonTrackResponseBodyData) SetRegion(v string) *QueryMultiAccountCarbonTrackResponseBodyData {
	s.Region = &v
	return s
}

func (s *QueryMultiAccountCarbonTrackResponseBodyData) SetUid(v string) *QueryMultiAccountCarbonTrackResponseBodyData {
	s.Uid = &v
	return s
}

type QueryMultiAccountCarbonTrackResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMultiAccountCarbonTrackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMultiAccountCarbonTrackResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMultiAccountCarbonTrackResponse) GoString() string {
	return s.String()
}

func (s *QueryMultiAccountCarbonTrackResponse) SetHeaders(v map[string]*string) *QueryMultiAccountCarbonTrackResponse {
	s.Headers = v
	return s
}

func (s *QueryMultiAccountCarbonTrackResponse) SetStatusCode(v int32) *QueryMultiAccountCarbonTrackResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMultiAccountCarbonTrackResponse) SetBody(v *QueryMultiAccountCarbonTrackResponseBody) *QueryMultiAccountCarbonTrackResponse {
	s.Body = v
	return s
}

type VerifyResponseBody struct {
	Data      *VerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyResponseBody) SetData(v *VerifyResponseBodyData) *VerifyResponseBody {
	s.Data = v
	return s
}

func (s *VerifyResponseBody) SetRequestId(v string) *VerifyResponseBody {
	s.RequestId = &v
	return s
}

type VerifyResponseBodyData struct {
	AllowedUids         []*string                                    `json:"AllowedUids,omitempty" xml:"AllowedUids,omitempty" type:"Repeated"`
	AccountType         *int32                                       `json:"accountType,omitempty" xml:"accountType,omitempty"`
	AllMultiAccountUids []*VerifyResponseBodyDataAllMultiAccountUids `json:"allMultiAccountUids,omitempty" xml:"allMultiAccountUids,omitempty" type:"Repeated"`
	Code                *string                                      `json:"code,omitempty" xml:"code,omitempty"`
	Message             *string                                      `json:"message,omitempty" xml:"message,omitempty"`
	MultiAccountsAllow  *int32                                       `json:"multiAccountsAllow,omitempty" xml:"multiAccountsAllow,omitempty"`
}

func (s VerifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyResponseBodyData) SetAllowedUids(v []*string) *VerifyResponseBodyData {
	s.AllowedUids = v
	return s
}

func (s *VerifyResponseBodyData) SetAccountType(v int32) *VerifyResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *VerifyResponseBodyData) SetAllMultiAccountUids(v []*VerifyResponseBodyDataAllMultiAccountUids) *VerifyResponseBodyData {
	s.AllMultiAccountUids = v
	return s
}

func (s *VerifyResponseBodyData) SetCode(v string) *VerifyResponseBodyData {
	s.Code = &v
	return s
}

func (s *VerifyResponseBodyData) SetMessage(v string) *VerifyResponseBodyData {
	s.Message = &v
	return s
}

func (s *VerifyResponseBodyData) SetMultiAccountsAllow(v int32) *VerifyResponseBodyData {
	s.MultiAccountsAllow = &v
	return s
}

type VerifyResponseBodyDataAllMultiAccountUids struct {
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s VerifyResponseBodyDataAllMultiAccountUids) String() string {
	return tea.Prettify(s)
}

func (s VerifyResponseBodyDataAllMultiAccountUids) GoString() string {
	return s.String()
}

func (s *VerifyResponseBodyDataAllMultiAccountUids) SetAccountId(v string) *VerifyResponseBodyDataAllMultiAccountUids {
	s.AccountId = &v
	return s
}

func (s *VerifyResponseBodyDataAllMultiAccountUids) SetDisplayName(v string) *VerifyResponseBodyDataAllMultiAccountUids {
	s.DisplayName = &v
	return s
}

type VerifyResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyResponse) GoString() string {
	return s.String()
}

func (s *VerifyResponse) SetHeaders(v map[string]*string) *VerifyResponse {
	s.Headers = v
	return s
}

func (s *VerifyResponse) SetStatusCode(v int32) *VerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyResponse) SetBody(v *VerifyResponseBody) *VerifyResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("carbonfootprint"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllowWithOptions(runtime *util.RuntimeOptions) (_result *AllowResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("Allow"),
		Version:     tea.String("2023-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Allow() (_result *AllowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllowResponse{}
	_body, _err := client.AllowWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSummaryDataWithOptions(tmpReq *GetSummaryDataRequest, runtime *util.RuntimeOptions) (_result *GetSummaryDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSummaryDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Uids)) {
		request.UidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Uids, tea.String("Uids"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UidsShrink)) {
		query["Uids"] = request.UidsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSummaryData"),
		Version:     tea.String("2023-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSummaryDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSummaryData(request *GetSummaryDataRequest) (_result *GetSummaryDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSummaryDataResponse{}
	_body, _err := client.GetSummaryDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCarbonTrackWithOptions(tmpReq *QueryCarbonTrackRequest, runtime *util.RuntimeOptions) (_result *QueryCarbonTrackResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryCarbonTrackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Uids)) {
		request.UidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Uids, tea.String("Uids"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterRDAccount)) {
		query["FilterRDAccount"] = request.FilterRDAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TopNum)) {
		query["TopNum"] = request.TopNum
	}

	if !tea.BoolValue(util.IsUnset(request.UidsShrink)) {
		query["Uids"] = request.UidsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCarbonTrack"),
		Version:     tea.String("2023-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCarbonTrackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCarbonTrack(request *QueryCarbonTrackRequest) (_result *QueryCarbonTrackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCarbonTrackResponse{}
	_body, _err := client.QueryCarbonTrackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMultiAccountCarbonTrackWithOptions(request *QueryMultiAccountCarbonTrackRequest, runtime *util.RuntimeOptions) (_result *QueryMultiAccountCarbonTrackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMultiAccountCarbonTrack"),
		Version:     tea.String("2023-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMultiAccountCarbonTrackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMultiAccountCarbonTrack(request *QueryMultiAccountCarbonTrackRequest) (_result *QueryMultiAccountCarbonTrackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMultiAccountCarbonTrackResponse{}
	_body, _err := client.QueryMultiAccountCarbonTrackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyWithOptions(runtime *util.RuntimeOptions) (_result *VerifyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("Verify"),
		Version:     tea.String("2023-07-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Verify() (_result *VerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyResponse{}
	_body, _err := client.VerifyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
