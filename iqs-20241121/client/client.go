// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("iqs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申请正式开通
//
// @param request - ApplyFormalServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyFormalServiceResponse
func (client *Client) ApplyFormalServiceWithOptions(request *ApplyFormalServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ApplyFormalServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyFormalService"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/services/commands/applyFormalService"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyFormalServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请正式开通
//
// @param request - ApplyFormalServiceRequest
//
// @return ApplyFormalServiceResponse
func (client *Client) ApplyFormalService(request *ApplyFormalServiceRequest) (_result *ApplyFormalServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyFormalServiceResponse{}
	_body, _err := client.ApplyFormalServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验账号类型
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountTypeResponse
func (client *Client) CheckAccountTypeWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckAccountTypeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAccountType"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/openService/v1/account/commands/checkAccountType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAccountTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验账号类型
//
// @return CheckAccountTypeResponse
func (client *Client) CheckAccountType() (_result *CheckAccountTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckAccountTypeResponse{}
	_body, _err := client.CheckAccountTypeWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载接口计量日明细
//
// @param request - DownloadApiCallDailyDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadApiCallDailyDetailResponse
func (client *Client) DownloadApiCallDailyDetailWithOptions(request *DownloadApiCallDailyDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DownloadApiCallDailyDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadApiCallDailyDetail"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/monitors/commands/downloadApiCallDailyDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadApiCallDailyDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载接口计量日明细
//
// @param request - DownloadApiCallDailyDetailRequest
//
// @return DownloadApiCallDailyDetailResponse
func (client *Client) DownloadApiCallDailyDetail(request *DownloadApiCallDailyDetailRequest) (_result *DownloadApiCallDailyDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DownloadApiCallDailyDetailResponse{}
	_body, _err := client.DownloadApiCallDailyDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载接口计量日明细
//
// @param request - DownloadMeteringDailyDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadMeteringDailyDetailResponse
func (client *Client) DownloadMeteringDailyDetailWithOptions(request *DownloadMeteringDailyDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DownloadMeteringDailyDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadMeteringDailyDetail"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/monitors/commands/downloadMeteringDailyDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadMeteringDailyDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载接口计量日明细
//
// @param request - DownloadMeteringDailyDetailRequest
//
// @return DownloadMeteringDailyDetailResponse
func (client *Client) DownloadMeteringDailyDetail(request *DownloadMeteringDailyDetailRequest) (_result *DownloadMeteringDailyDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DownloadMeteringDailyDetailResponse{}
	_body, _err := client.DownloadMeteringDailyDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 延长测试时间
//
// @param request - ExpandSearchExpiredTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpandSearchExpiredTimeResponse
func (client *Client) ExpandSearchExpiredTimeWithOptions(request *ExpandSearchExpiredTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExpandSearchExpiredTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExpandSearchExpiredTime"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/openService/v1/account/commands/expendExpiredTime"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExpandSearchExpiredTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 延长测试时间
//
// @param request - ExpandSearchExpiredTimeRequest
//
// @return ExpandSearchExpiredTimeResponse
func (client *Client) ExpandSearchExpiredTime(request *ExpandSearchExpiredTimeRequest) (_result *ExpandSearchExpiredTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExpandSearchExpiredTimeResponse{}
	_body, _err := client.ExpandSearchExpiredTimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取账号配置信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountConfigInfoResponse
func (client *Client) GetAccountConfigInfoWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAccountConfigInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountConfigInfo"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/openService/v1/account/command/accountConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountConfigInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取账号配置信息
//
// @return GetAccountConfigInfoResponse
func (client *Client) GetAccountConfigInfo() (_result *GetAccountConfigInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAccountConfigInfoResponse{}
	_body, _err := client.GetAccountConfigInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户申请记录
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountReviewRecordResponse
func (client *Client) GetAccountReviewRecordWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAccountReviewRecordResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountReviewRecord"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/services/commands/reviewRecord"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountReviewRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户申请记录
//
// @return GetAccountReviewRecordResponse
func (client *Client) GetAccountReviewRecord() (_result *GetAccountReviewRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAccountReviewRecordResponse{}
	_body, _err := client.GetAccountReviewRecordWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询计量汇总信息
//
// @param request - GetMeteringSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMeteringSummaryResponse
func (client *Client) GetMeteringSummaryWithOptions(request *GetMeteringSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMeteringSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillPeriod) {
		query["billPeriod"] = request.BillPeriod
	}

	if !dara.IsNil(request.BillingItem) {
		query["billingItem"] = request.BillingItem
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubAccountId) {
		query["subAccountId"] = request.SubAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMeteringSummary"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/monitors/commands/metering/summary"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMeteringSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询计量汇总信息
//
// @param request - GetMeteringSummaryRequest
//
// @return GetMeteringSummaryResponse
func (client *Client) GetMeteringSummary(request *GetMeteringSummaryRequest) (_result *GetMeteringSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMeteringSummaryResponse{}
	_body, _err := client.GetMeteringSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询转正用户调用量信息
//
// @param request - GetNormalServiceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNormalServiceConfigResponse
func (client *Client) GetNormalServiceConfigWithOptions(request *GetNormalServiceConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetNormalServiceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNormalServiceConfig"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/services/commands/normalServiceConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNormalServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询转正用户调用量信息
//
// @param request - GetNormalServiceConfigRequest
//
// @return GetNormalServiceConfigResponse
func (client *Client) GetNormalServiceConfig(request *GetNormalServiceConfigRequest) (_result *GetNormalServiceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNormalServiceConfigResponse{}
	_body, _err := client.GetNormalServiceConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务额度信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceConfigResponse
func (client *Client) GetServiceConfigWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceConfig"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/services/commands/serviceConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务额度信息
//
// @return GetServiceConfigResponse
func (client *Client) GetServiceConfig() (_result *GetServiceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceConfigResponse{}
	_body, _err := client.GetServiceConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Api调用量（性能）日明细
//
// @param request - ListApiCallDailyDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiCallDailyDetailResponse
func (client *Client) ListApiCallDailyDetailWithOptions(request *ListApiCallDailyDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiCallDailyDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["apiName"] = request.ApiName
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EngineType) {
		query["engineType"] = request.EngineType
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiCallDailyDetail"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/monitors/commands/apiCall/dailyDetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiCallDailyDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Api调用量（性能）日明细
//
// @param request - ListApiCallDailyDetailRequest
//
// @return ListApiCallDailyDetailResponse
func (client *Client) ListApiCallDailyDetail(request *ListApiCallDailyDetailRequest) (_result *ListApiCallDailyDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApiCallDailyDetailResponse{}
	_body, _err := client.ListApiCallDailyDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// api基础信息
//
// @param request - ListApiInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiInfosResponse
func (client *Client) ListApiInfosWithOptions(request *ListApiInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiInfos"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/monitors/commands/apiInfos"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// api基础信息
//
// @param request - ListApiInfosRequest
//
// @return ListApiInfosResponse
func (client *Client) ListApiInfos(request *ListApiInfosRequest) (_result *ListApiInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApiInfosResponse{}
	_body, _err := client.ListApiInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Api名称列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiNamesResponse
func (client *Client) ListApiNamesWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiNamesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiNames"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/monitors/commands/apiNames"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Api名称列表
//
// @return ListApiNamesResponse
func (client *Client) ListApiNames() (_result *ListApiNamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApiNamesResponse{}
	_body, _err := client.ListApiNamesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// api基础信息
//
// @param request - ListLimitationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLimitationsResponse
func (client *Client) ListLimitationsWithOptions(request *ListLimitationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLimitationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLimitations"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/openService/v1/console/limitation/commands/list/account"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLimitationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// api基础信息
//
// @param request - ListLimitationsRequest
//
// @return ListLimitationsResponse
func (client *Client) ListLimitations(request *ListLimitationsRequest) (_result *ListLimitationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLimitationsResponse{}
	_body, _err := client.ListLimitationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询计量日明细信息
//
// @param request - ListMeteringDailyDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMeteringDailyDetailResponse
func (client *Client) ListMeteringDailyDetailWithOptions(request *ListMeteringDailyDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMeteringDailyDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillPeriod) {
		query["billPeriod"] = request.BillPeriod
	}

	if !dara.IsNil(request.BillingItem) {
		query["billingItem"] = request.BillingItem
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.SubAccountId) {
		query["subAccountId"] = request.SubAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMeteringDailyDetail"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/monitors/commands/metering/dailyDetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMeteringDailyDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询计量日明细信息
//
// @param request - ListMeteringDailyDetailRequest
//
// @return ListMeteringDailyDetailResponse
func (client *Client) ListMeteringDailyDetail(request *ListMeteringDailyDetailRequest) (_result *ListMeteringDailyDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMeteringDailyDetailResponse{}
	_body, _err := client.ListMeteringDailyDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询主账号的所有子账号信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubAccountInfoResponse
func (client *Client) ListSubAccountInfoWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSubAccountInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubAccountInfo"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/console/v1/monitors/commands/subAccountInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询主账号的所有子账号信息
//
// @return ListSubAccountInfoResponse
func (client *Client) ListSubAccountInfo() (_result *ListSubAccountInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSubAccountInfoResponse{}
	_body, _err := client.ListSubAccountInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 管理智搜用户
//
// @param request - ManageSearchAccountInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageSearchAccountInfoResponse
func (client *Client) ManageSearchAccountInfoWithOptions(request *ManageSearchAccountInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ManageSearchAccountInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManageSearchAccountInfo"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/openService/v1/account/commands/manage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ManageSearchAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 管理智搜用户
//
// @param request - ManageSearchAccountInfoRequest
//
// @return ManageSearchAccountInfoResponse
func (client *Client) ManageSearchAccountInfo(request *ManageSearchAccountInfoRequest) (_result *ManageSearchAccountInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ManageSearchAccountInfoResponse{}
	_body, _err := client.ManageSearchAccountInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止自动转正
//
// @param request - OpenAutoNormalReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenAutoNormalReviewResponse
func (client *Client) OpenAutoNormalReviewWithOptions(request *OpenAutoNormalReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenAutoNormalReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenAutoNormalReview"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/openService/v1/account/commands/openAutoNormalReview"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenAutoNormalReviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止自动转正
//
// @param request - OpenAutoNormalReviewRequest
//
// @return OpenAutoNormalReviewResponse
func (client *Client) OpenAutoNormalReview(request *OpenAutoNormalReviewRequest) (_result *OpenAutoNormalReviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenAutoNormalReviewResponse{}
	_body, _err := client.OpenAutoNormalReviewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止自动转正
//
// @param request - StopAutoNormalReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAutoNormalReviewResponse
func (client *Client) StopAutoNormalReviewWithOptions(request *StopAutoNormalReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopAutoNormalReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAutoNormalReview"),
		Version:     dara.String("2024-11-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/linked-retrieval/linked-retrieval-admin/openService/v1/account/commands/stopAutoNormalReview"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAutoNormalReviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止自动转正
//
// @param request - StopAutoNormalReviewRequest
//
// @return StopAutoNormalReviewResponse
func (client *Client) StopAutoNormalReview(request *StopAutoNormalReviewRequest) (_result *StopAutoNormalReviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopAutoNormalReviewResponse{}
	_body, _err := client.StopAutoNormalReviewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
