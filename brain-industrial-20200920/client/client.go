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
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("brain-industrial"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 激活License
//
// @param request - ActivateLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicenseWithOptions(request *ActivateLicenseRequest, runtime *dara.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Fingerprint) {
		body["Fingerprint"] = request.Fingerprint
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderId) {
		body["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateLicense"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 激活License
//
// @param request - ActivateLicenseRequest
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicense(request *ActivateLicenseRequest) (_result *ActivateLicenseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.ActivateLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用aics openapi
//
// @param tmpReq - AicsOpenApiInvokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AicsOpenApiInvokeResponse
func (client *Client) AicsOpenApiInvokeWithOptions(tmpReq *AicsOpenApiInvokeRequest, runtime *dara.RuntimeOptions) (_result *AicsOpenApiInvokeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AicsOpenApiInvokeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ParamShrink) {
		body["Param"] = request.ParamShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AicsOpenApiInvoke"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AicsOpenApiInvokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用aics openapi
//
// @param request - AicsOpenApiInvokeRequest
//
// @return AicsOpenApiInvokeResponse
func (client *Client) AicsOpenApiInvoke(request *AicsOpenApiInvokeRequest) (_result *AicsOpenApiInvokeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AicsOpenApiInvokeResponse{}
	_body, _err := client.AicsOpenApiInvokeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建储能运行优化任务
//
// @param tmpReq - CreateEssOptJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEssOptJobResponse
func (client *Client) CreateEssOptJobWithOptions(tmpReq *CreateEssOptJobRequest, runtime *dara.RuntimeOptions) (_result *CreateEssOptJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEssOptJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ElecPrice) {
		request.ElecPriceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ElecPrice, dara.String("ElecPrice"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GenPrice) {
		request.GenPriceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenPrice, dara.String("GenPrice"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Location) {
		request.LocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Location, dara.String("Location"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SystemData) {
		request.SystemDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SystemData, dara.String("SystemData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessKey) {
		body["BusinessKey"] = request.BusinessKey
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.ElecPriceShrink) {
		body["ElecPrice"] = request.ElecPriceShrink
	}

	if !dara.IsNil(request.Freq) {
		body["Freq"] = request.Freq
	}

	if !dara.IsNil(request.GenPriceShrink) {
		body["GenPrice"] = request.GenPriceShrink
	}

	if !dara.IsNil(request.LocationShrink) {
		body["Location"] = request.LocationShrink
	}

	if !dara.IsNil(request.ModelVersion) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.RunDate) {
		body["RunDate"] = request.RunDate
	}

	if !dara.IsNil(request.SystemDataShrink) {
		body["SystemData"] = request.SystemDataShrink
	}

	if !dara.IsNil(request.TimeZone) {
		body["TimeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.TopoType) {
		body["TopoType"] = request.TopoType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEssOptJob"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEssOptJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建储能运行优化任务
//
// @param request - CreateEssOptJobRequest
//
// @return CreateEssOptJobResponse
func (client *Client) CreateEssOptJob(request *CreateEssOptJobRequest) (_result *CreateEssOptJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEssOptJobResponse{}
	_body, _err := client.CreateEssOptJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用电负荷预测任务，历史数据来自文件url
//
// @param request - CreateLoadForecastByFileUrlJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadForecastByFileUrlJobResponse
func (client *Client) CreateLoadForecastByFileUrlJobWithOptions(request *CreateLoadForecastByFileUrlJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadForecastByFileUrlJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessKey) {
		body["BusinessKey"] = request.BusinessKey
	}

	if !dara.IsNil(request.DataMode) {
		body["DataMode"] = request.DataMode
	}

	if !dara.IsNil(request.DeviceType) {
		body["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.ForecastHorizon) {
		body["ForecastHorizon"] = request.ForecastHorizon
	}

	if !dara.IsNil(request.Freq) {
		body["Freq"] = request.Freq
	}

	if !dara.IsNil(request.HistoryUrl) {
		body["HistoryUrl"] = request.HistoryUrl
	}

	if !dara.IsNil(request.ModelVersion) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.RunDate) {
		body["RunDate"] = request.RunDate
	}

	if !dara.IsNil(request.SystemType) {
		body["SystemType"] = request.SystemType
	}

	if !dara.IsNil(request.TimeColumn) {
		body["TimeColumn"] = request.TimeColumn
	}

	if !dara.IsNil(request.TimeZone) {
		body["TimeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.ValueColumn) {
		body["ValueColumn"] = request.ValueColumn
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadForecastByFileUrlJob"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadForecastByFileUrlJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用电负荷预测任务，历史数据来自文件url
//
// @param request - CreateLoadForecastByFileUrlJobRequest
//
// @return CreateLoadForecastByFileUrlJobResponse
func (client *Client) CreateLoadForecastByFileUrlJob(request *CreateLoadForecastByFileUrlJobRequest) (_result *CreateLoadForecastByFileUrlJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoadForecastByFileUrlJobResponse{}
	_body, _err := client.CreateLoadForecastByFileUrlJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用电负荷预测任务
//
// @param tmpReq - CreateLoadForecastJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadForecastJobResponse
func (client *Client) CreateLoadForecastJobWithOptions(tmpReq *CreateLoadForecastJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadForecastJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLoadForecastJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HistoryData) {
		request.HistoryDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HistoryData, dara.String("HistoryData"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessKey) {
		body["BusinessKey"] = request.BusinessKey
	}

	if !dara.IsNil(request.DataMode) {
		body["DataMode"] = request.DataMode
	}

	if !dara.IsNil(request.DeviceType) {
		body["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.ForecastHorizon) {
		body["ForecastHorizon"] = request.ForecastHorizon
	}

	if !dara.IsNil(request.Freq) {
		body["Freq"] = request.Freq
	}

	if !dara.IsNil(request.HistoryDataShrink) {
		body["HistoryData"] = request.HistoryDataShrink
	}

	if !dara.IsNil(request.ModelVersion) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.RunDate) {
		body["RunDate"] = request.RunDate
	}

	if !dara.IsNil(request.SystemType) {
		body["SystemType"] = request.SystemType
	}

	if !dara.IsNil(request.TimeZone) {
		body["TimeZone"] = request.TimeZone
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadForecastJob"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadForecastJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用电负荷预测任务
//
// @param request - CreateLoadForecastJobRequest
//
// @return CreateLoadForecastJobResponse
func (client *Client) CreateLoadForecastJob(request *CreateLoadForecastJobRequest) (_result *CreateLoadForecastJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoadForecastJobResponse{}
	_body, _err := client.CreateLoadForecastJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建发电功率预测任务，历史数据来自文件url
//
// @param tmpReq - CreatePowerForecastByFileUrlJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePowerForecastByFileUrlJobResponse
func (client *Client) CreatePowerForecastByFileUrlJobWithOptions(tmpReq *CreatePowerForecastByFileUrlJobRequest, runtime *dara.RuntimeOptions) (_result *CreatePowerForecastByFileUrlJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePowerForecastByFileUrlJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Location) {
		request.LocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Location, dara.String("Location"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessKey) {
		body["BusinessKey"] = request.BusinessKey
	}

	if !dara.IsNil(request.DataMode) {
		body["DataMode"] = request.DataMode
	}

	if !dara.IsNil(request.DeviceType) {
		body["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.ForecastHorizon) {
		body["ForecastHorizon"] = request.ForecastHorizon
	}

	if !dara.IsNil(request.Freq) {
		body["Freq"] = request.Freq
	}

	if !dara.IsNil(request.HistoryUrl) {
		body["HistoryUrl"] = request.HistoryUrl
	}

	if !dara.IsNil(request.LocationShrink) {
		body["Location"] = request.LocationShrink
	}

	if !dara.IsNil(request.ModelVersion) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.RunDate) {
		body["RunDate"] = request.RunDate
	}

	if !dara.IsNil(request.SystemType) {
		body["SystemType"] = request.SystemType
	}

	if !dara.IsNil(request.TimeColumn) {
		body["TimeColumn"] = request.TimeColumn
	}

	if !dara.IsNil(request.TimeZone) {
		body["TimeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.ValueColumn) {
		body["ValueColumn"] = request.ValueColumn
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePowerForecastByFileUrlJob"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePowerForecastByFileUrlJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建发电功率预测任务，历史数据来自文件url
//
// @param request - CreatePowerForecastByFileUrlJobRequest
//
// @return CreatePowerForecastByFileUrlJobResponse
func (client *Client) CreatePowerForecastByFileUrlJob(request *CreatePowerForecastByFileUrlJobRequest) (_result *CreatePowerForecastByFileUrlJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePowerForecastByFileUrlJobResponse{}
	_body, _err := client.CreatePowerForecastByFileUrlJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建发电功率预测任务
//
// @param tmpReq - CreatePowerForecastJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePowerForecastJobResponse
func (client *Client) CreatePowerForecastJobWithOptions(tmpReq *CreatePowerForecastJobRequest, runtime *dara.RuntimeOptions) (_result *CreatePowerForecastJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePowerForecastJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HistoryData) {
		request.HistoryDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HistoryData, dara.String("HistoryData"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Location) {
		request.LocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Location, dara.String("Location"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessKey) {
		body["BusinessKey"] = request.BusinessKey
	}

	if !dara.IsNil(request.DataMode) {
		body["DataMode"] = request.DataMode
	}

	if !dara.IsNil(request.DeviceType) {
		body["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.ForecastHorizon) {
		body["ForecastHorizon"] = request.ForecastHorizon
	}

	if !dara.IsNil(request.Freq) {
		body["Freq"] = request.Freq
	}

	if !dara.IsNil(request.HistoryDataShrink) {
		body["HistoryData"] = request.HistoryDataShrink
	}

	if !dara.IsNil(request.LocationShrink) {
		body["Location"] = request.LocationShrink
	}

	if !dara.IsNil(request.ModelVersion) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.RunDate) {
		body["RunDate"] = request.RunDate
	}

	if !dara.IsNil(request.SystemType) {
		body["SystemType"] = request.SystemType
	}

	if !dara.IsNil(request.TimeZone) {
		body["TimeZone"] = request.TimeZone
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePowerForecastJob"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePowerForecastJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建发电功率预测任务
//
// @param request - CreatePowerForecastJobRequest
//
// @return CreatePowerForecastJobResponse
func (client *Client) CreatePowerForecastJob(request *CreatePowerForecastJobRequest) (_result *CreatePowerForecastJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePowerForecastJobResponse{}
	_body, _err := client.CreatePowerForecastJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询aivpp算法job
//
// @param request - GetAivppAlgoJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAivppAlgoJobResponse
func (client *Client) GetAivppAlgoJobWithOptions(request *GetAivppAlgoJobRequest, runtime *dara.RuntimeOptions) (_result *GetAivppAlgoJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAivppAlgoJob"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAivppAlgoJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询aivpp算法job
//
// @param request - GetAivppAlgoJobRequest
//
// @return GetAivppAlgoJobResponse
func (client *Client) GetAivppAlgoJob(request *GetAivppAlgoJobRequest) (_result *GetAivppAlgoJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAivppAlgoJobResponse{}
	_body, _err := client.GetAivppAlgoJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # License详情
//
// @param request - GetLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLicenseResponse
func (client *Client) GetLicenseWithOptions(request *GetLicenseRequest, runtime *dara.RuntimeOptions) (_result *GetLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLicense"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # License详情
//
// @param request - GetLicenseRequest
//
// @return GetLicenseResponse
func (client *Client) GetLicense(request *GetLicenseRequest) (_result *GetLicenseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLicenseResponse{}
	_body, _err := client.GetLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户AIVPP资源列表
//
// @param request - ListAivppResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAivppResourcesResponse
func (client *Client) ListAivppResourcesWithOptions(request *ListAivppResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListAivppResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAivppResources"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAivppResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户AIVPP资源列表
//
// @param request - ListAivppResourcesRequest
//
// @return ListAivppResourcesResponse
func (client *Client) ListAivppResources(request *ListAivppResourcesRequest) (_result *ListAivppResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAivppResourcesResponse{}
	_body, _err := client.ListAivppResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # License列表
//
// @param request - ListLicensesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLicensesResponse
func (client *Client) ListLicensesWithOptions(request *ListLicensesRequest, runtime *dara.RuntimeOptions) (_result *ListLicensesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryStr) {
		body["QueryStr"] = request.QueryStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLicenses"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLicensesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # License列表
//
// @param request - ListLicensesRequest
//
// @return ListLicensesResponse
func (client *Client) ListLicenses(request *ListLicensesRequest) (_result *ListLicensesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLicensesResponse{}
	_body, _err := client.ListLicensesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户资源列表
//
// @param request - ListUserResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserResourcesResponse
func (client *Client) ListUserResourcesWithOptions(request *ListUserResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListUserResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserResources"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户资源列表
//
// @param request - ListUserResourcesRequest
//
// @return ListUserResourcesResponse
func (client *Client) ListUserResources(request *ListUserResourcesRequest) (_result *ListUserResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserResourcesResponse{}
	_body, _err := client.ListUserResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新license描述
//
// @param request - UpdateLicenseDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLicenseDescriptionResponse
func (client *Client) UpdateLicenseDescriptionWithOptions(request *UpdateLicenseDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateLicenseDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLicenseDescription"),
		Version:     dara.String("2020-09-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLicenseDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新license描述
//
// @param request - UpdateLicenseDescriptionRequest
//
// @return UpdateLicenseDescriptionResponse
func (client *Client) UpdateLicenseDescription(request *UpdateLicenseDescriptionRequest) (_result *UpdateLicenseDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLicenseDescriptionResponse{}
	_body, _err := client.UpdateLicenseDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
