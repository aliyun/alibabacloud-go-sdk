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
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        dara.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            dara.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        dara.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        dara.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        dara.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            dara.String("green.aliyuncs.com"),
		"cn-hongkong":           dara.String("green.aliyuncs.com"),
		"cn-huhehaote":          dara.String("green.aliyuncs.com"),
		"cn-qingdao":            dara.String("green.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("green.aliyuncs.com"),
		"eu-central-1":          dara.String("green.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             dara.String("green.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             dara.String("green.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             dara.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("green.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("green.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("green.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("green.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("green"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加代答样本
//
// @param request - AddAnswerSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAnswerSampleResponse
func (client *Client) AddAnswerSampleWithOptions(request *AddAnswerSampleRequest, runtime *dara.RuntimeOptions) (_result *AddAnswerSampleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SampleObject) {
		query["SampleObject"] = request.SampleObject
	}

	if !dara.IsNil(request.Samples) {
		query["Samples"] = request.Samples
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAnswerSample"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAnswerSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加代答样本
//
// @param request - AddAnswerSampleRequest
//
// @return AddAnswerSampleResponse
func (client *Client) AddAnswerSample(request *AddAnswerSampleRequest) (_result *AddAnswerSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAnswerSampleResponse{}
	_body, _err := client.AddAnswerSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Image Library
//
// @param request - AddImageLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageLibResponse
func (client *Client) AddImageLibWithOptions(request *AddImageLibRequest, runtime *dara.RuntimeOptions) (_result *AddImageLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.LibName) {
		body["LibName"] = request.LibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddImageLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Image Library
//
// @param request - AddImageLibRequest
//
// @return AddImageLibResponse
func (client *Client) AddImageLib(request *AddImageLibRequest) (_result *AddImageLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddImageLibResponse{}
	_body, _err := client.AddImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add image to image lib
//
// @param request - AddImages2LibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImages2LibResponse
func (client *Client) AddImages2LibWithOptions(request *AddImages2LibRequest, runtime *dara.RuntimeOptions) (_result *AddImages2LibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImgUrl) {
		body["ImgUrl"] = request.ImgUrl
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddImages2Lib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddImages2LibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add image to image lib
//
// @param request - AddImages2LibRequest
//
// @return AddImages2LibResponse
func (client *Client) AddImages2Lib(request *AddImages2LibRequest) (_result *AddImages2LibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddImages2LibResponse{}
	_body, _err := client.AddImages2LibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create keyword library
//
// @param request - AddKeywordLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddKeywordLibResponse
func (client *Client) AddKeywordLibWithOptions(request *AddKeywordLibRequest, runtime *dara.RuntimeOptions) (_result *AddKeywordLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Keywords) {
		body["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.KeywordsObject) {
		body["KeywordsObject"] = request.KeywordsObject
	}

	if !dara.IsNil(request.LibName) {
		body["LibName"] = request.LibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddKeywordLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create keyword library
//
// @param request - AddKeywordLibRequest
//
// @return AddKeywordLibResponse
func (client *Client) AddKeywordLib(request *AddKeywordLibRequest) (_result *AddKeywordLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddKeywordLibResponse{}
	_body, _err := client.AddKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add keywords
//
// @param request - AddKeywordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddKeywordsResponse
func (client *Client) AddKeywordsWithOptions(request *AddKeywordsRequest, runtime *dara.RuntimeOptions) (_result *AddKeywordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Keywords) {
		body["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.KeywordsObject) {
		body["KeywordsObject"] = request.KeywordsObject
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddKeywords"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddKeywordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add keywords
//
// @param request - AddKeywordsRequest
//
// @return AddKeywordsResponse
func (client *Client) AddKeywords(request *AddKeywordsRequest) (_result *AddKeywordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddKeywordsResponse{}
	_body, _err := client.AddKeywordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add keywords to keyword library.
//
// @param request - AddKeywordsToLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddKeywordsToLibResponse
func (client *Client) AddKeywordsToLibWithOptions(request *AddKeywordsToLibRequest, runtime *dara.RuntimeOptions) (_result *AddKeywordsToLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Keywords) {
		body["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.KeywordsObject) {
		body["KeywordsObject"] = request.KeywordsObject
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddKeywordsToLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddKeywordsToLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add keywords to keyword library.
//
// @param request - AddKeywordsToLibRequest
//
// @return AddKeywordsToLibResponse
func (client *Client) AddKeywordsToLib(request *AddKeywordsToLibRequest) (_result *AddKeywordsToLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddKeywordsToLibResponse{}
	_body, _err := client.AddKeywordsToLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Cancel OSS detection task
//
// @param request - CancelStockOssCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelStockOssCheckTaskResponse
func (client *Client) CancelStockOssCheckTaskWithOptions(request *CancelStockOssCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelStockOssCheckTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelStockOssCheckTask"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelStockOssCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Cancel OSS detection task
//
// @param request - CancelStockOssCheckTaskRequest
//
// @return CancelStockOssCheckTaskResponse
func (client *Client) CancelStockOssCheckTask(request *CancelStockOssCheckTaskRequest) (_result *CancelStockOssCheckTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelStockOssCheckTaskResponse{}
	_body, _err := client.CancelStockOssCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// copy service config
//
// @param request - CopyServiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyServiceConfigResponse
func (client *Client) CopyServiceConfigWithOptions(request *CopyServiceConfigRequest, runtime *dara.RuntimeOptions) (_result *CopyServiceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.ServiceDesc) {
		body["ServiceDesc"] = request.ServiceDesc
	}

	if !dara.IsNil(request.ServiceName) {
		body["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyServiceConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// copy service config
//
// @param request - CopyServiceConfigRequest
//
// @return CopyServiceConfigResponse
func (client *Client) CopyServiceConfig(request *CopyServiceConfigRequest) (_result *CopyServiceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyServiceConfigResponse{}
	_body, _err := client.CopyServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create stock oss check task
//
// @param request - CreatStockOssCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatStockOssCheckTaskResponse
func (client *Client) CreatStockOssCheckTaskWithOptions(request *CreatStockOssCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *CreatStockOssCheckTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketPrefixFilterConfig) {
		query["BucketPrefixFilterConfig"] = request.BucketPrefixFilterConfig
	}

	if !dara.IsNil(request.Buckets) {
		query["Buckets"] = request.Buckets
	}

	if !dara.IsNil(request.CallbackId) {
		query["CallbackId"] = request.CallbackId
	}

	if !dara.IsNil(request.DistinctHistoryTasks) {
		query["DistinctHistoryTasks"] = request.DistinctHistoryTasks
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecuteDate) {
		query["ExecuteDate"] = request.ExecuteDate
	}

	if !dara.IsNil(request.ExecuteTime) {
		query["ExecuteTime"] = request.ExecuteTime
	}

	if !dara.IsNil(request.Freeze) {
		query["Freeze"] = request.Freeze
	}

	if !dara.IsNil(request.FreezeHighRisk1) {
		query["FreezeHighRisk1"] = request.FreezeHighRisk1
	}

	if !dara.IsNil(request.FreezeHighRisk2) {
		query["FreezeHighRisk2"] = request.FreezeHighRisk2
	}

	if !dara.IsNil(request.FreezeMediumRisk1) {
		query["FreezeMediumRisk1"] = request.FreezeMediumRisk1
	}

	if !dara.IsNil(request.FreezeMediumRisk2) {
		query["FreezeMediumRisk2"] = request.FreezeMediumRisk2
	}

	if !dara.IsNil(request.FreezeRestorePath) {
		query["FreezeRestorePath"] = request.FreezeRestorePath
	}

	if !dara.IsNil(request.FreezeType) {
		query["FreezeType"] = request.FreezeType
	}

	if !dara.IsNil(request.IsInc) {
		query["IsInc"] = request.IsInc
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PrefixFilterType) {
		query["PrefixFilterType"] = request.PrefixFilterType
	}

	if !dara.IsNil(request.PrefixFilters) {
		query["PrefixFilters"] = request.PrefixFilters
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Referer) {
		query["Referer"] = request.Referer
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScanLimit) {
		query["ScanLimit"] = request.ScanLimit
	}

	if !dara.IsNil(request.ScanNoFileType) {
		query["ScanNoFileType"] = request.ScanNoFileType
	}

	if !dara.IsNil(request.ScanResourceType) {
		query["ScanResourceType"] = request.ScanResourceType
	}

	if !dara.IsNil(request.ScanService) {
		query["ScanService"] = request.ScanService
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskCycle) {
		query["TaskCycle"] = request.TaskCycle
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatStockOssCheckTask"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatStockOssCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create stock oss check task
//
// @param request - CreatStockOssCheckTaskRequest
//
// @return CreatStockOssCheckTaskResponse
func (client *Client) CreatStockOssCheckTask(request *CreatStockOssCheckTaskRequest) (_result *CreatStockOssCheckTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatStockOssCheckTaskResponse{}
	_body, _err := client.CreatStockOssCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建代答库
//
// @param request - CreateAnswerLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnswerLibResponse
func (client *Client) CreateAnswerLibWithOptions(request *CreateAnswerLibRequest, runtime *dara.RuntimeOptions) (_result *CreateAnswerLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LibName) {
		body["LibName"] = request.LibName
	}

	if !dara.IsNil(request.SampleBucket) {
		body["SampleBucket"] = request.SampleBucket
	}

	if !dara.IsNil(request.SampleObject) {
		body["SampleObject"] = request.SampleObject
	}

	if !dara.IsNil(request.Samples) {
		body["Samples"] = request.Samples
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAnswerLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAnswerLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建代答库
//
// @param request - CreateAnswerLibRequest
//
// @return CreateAnswerLibResponse
func (client *Client) CreateAnswerLib(request *CreateAnswerLibRequest) (_result *CreateAnswerLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAnswerLibResponse{}
	_body, _err := client.CreateAnswerLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a new message notification
//
// @param request - CreateCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCallbackResponse
func (client *Client) CreateCallbackWithOptions(request *CreateCallbackRequest, runtime *dara.RuntimeOptions) (_result *CreateCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CryptType) {
		body["CryptType"] = request.CryptType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Scope) {
		body["Scope"] = request.Scope
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCallback"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a new message notification
//
// @param request - CreateCallbackRequest
//
// @return CreateCallbackResponse
func (client *Client) CreateCallback(request *CreateCallbackRequest) (_result *CreateCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCallbackResponse{}
	_body, _err := client.CreateCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Online Test
//
// @param request - CreateOnlineTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOnlineTestResponse
func (client *Client) CreateOnlineTestWithOptions(request *CreateOnlineTestRequest, runtime *dara.RuntimeOptions) (_result *CreateOnlineTestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOnlineTest"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOnlineTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Online Test
//
// @param request - CreateOnlineTestRequest
//
// @return CreateOnlineTestResponse
func (client *Client) CreateOnlineTest(request *CreateOnlineTestRequest) (_result *CreateOnlineTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOnlineTestResponse{}
	_body, _err := client.CreateOnlineTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check before creating an OSS scan task
//
// @param request - CreatePreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePreCheckResponse
func (client *Client) CreatePreCheckWithOptions(request *CreatePreCheckRequest, runtime *dara.RuntimeOptions) (_result *CreatePreCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BucketPrefixFilterConfig) {
		body["BucketPrefixFilterConfig"] = request.BucketPrefixFilterConfig
	}

	if !dara.IsNil(request.Buckets) {
		body["Buckets"] = request.Buckets
	}

	if !dara.IsNil(request.DistinctHistoryTasks) {
		body["DistinctHistoryTasks"] = request.DistinctHistoryTasks
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IsInc) {
		body["IsInc"] = request.IsInc
	}

	if !dara.IsNil(request.MediaType) {
		body["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PrefixFilterType) {
		body["PrefixFilterType"] = request.PrefixFilterType
	}

	if !dara.IsNil(request.PrefixFilters) {
		body["PrefixFilters"] = request.PrefixFilters
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ScanLimit) {
		body["ScanLimit"] = request.ScanLimit
	}

	if !dara.IsNil(request.ScanNoFileType) {
		body["ScanNoFileType"] = request.ScanNoFileType
	}

	if !dara.IsNil(request.ScanService) {
		body["ScanService"] = request.ScanService
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePreCheck"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check before creating an OSS scan task
//
// @param request - CreatePreCheckRequest
//
// @return CreatePreCheckResponse
func (client *Client) CreatePreCheck(request *CreatePreCheckRequest) (_result *CreatePreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePreCheckResponse{}
	_body, _err := client.CreatePreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除代答库
//
// @param request - DeleteAnswerLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAnswerLibResponse
func (client *Client) DeleteAnswerLibWithOptions(request *DeleteAnswerLibRequest, runtime *dara.RuntimeOptions) (_result *DeleteAnswerLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAnswerLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAnswerLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代答库
//
// @param request - DeleteAnswerLibRequest
//
// @return DeleteAnswerLibResponse
func (client *Client) DeleteAnswerLib(request *DeleteAnswerLibRequest) (_result *DeleteAnswerLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAnswerLibResponse{}
	_body, _err := client.DeleteAnswerLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除代答答案
//
// @param request - DeleteAnswerSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAnswerSampleResponse
func (client *Client) DeleteAnswerSampleWithOptions(request *DeleteAnswerSampleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAnswerSampleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		body["Ids"] = request.Ids
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAnswerSample"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAnswerSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代答答案
//
// @param request - DeleteAnswerSampleRequest
//
// @return DeleteAnswerSampleResponse
func (client *Client) DeleteAnswerSample(request *DeleteAnswerSampleRequest) (_result *DeleteAnswerSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAnswerSampleResponse{}
	_body, _err := client.DeleteAnswerSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// delete callback
//
// @param request - DeleteCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCallbackResponse
func (client *Client) DeleteCallbackWithOptions(request *DeleteCallbackRequest, runtime *dara.RuntimeOptions) (_result *DeleteCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCallback"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// delete callback
//
// @param request - DeleteCallbackRequest
//
// @return DeleteCallbackResponse
func (client *Client) DeleteCallback(request *DeleteCallbackRequest) (_result *DeleteCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCallbackResponse{}
	_body, _err := client.DeleteCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete feature configuration
//
// @param request - DeleteFeatureConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFeatureConfigResponse
func (client *Client) DeleteFeatureConfigWithOptions(request *DeleteFeatureConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteFeatureConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Field) {
		body["Field"] = request.Field
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFeatureConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFeatureConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete feature configuration
//
// @param request - DeleteFeatureConfigRequest
//
// @return DeleteFeatureConfigResponse
func (client *Client) DeleteFeatureConfig(request *DeleteFeatureConfigRequest) (_result *DeleteFeatureConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFeatureConfigResponse{}
	_body, _err := client.DeleteFeatureConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete images from library.
//
// @param request - DeleteImagesFromLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImagesFromLibResponse
func (client *Client) DeleteImagesFromLibWithOptions(request *DeleteImagesFromLibRequest, runtime *dara.RuntimeOptions) (_result *DeleteImagesFromLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageIds) {
		body["ImageIds"] = request.ImageIds
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImagesFromLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImagesFromLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete images from library.
//
// @param request - DeleteImagesFromLibRequest
//
// @return DeleteImagesFromLibResponse
func (client *Client) DeleteImagesFromLib(request *DeleteImagesFromLibRequest) (_result *DeleteImagesFromLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteImagesFromLibResponse{}
	_body, _err := client.DeleteImagesFromLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete keyword
//
// @param request - DeleteKeywordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeywordResponse
func (client *Client) DeleteKeywordWithOptions(request *DeleteKeywordRequest, runtime *dara.RuntimeOptions) (_result *DeleteKeywordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.KeywordIdList) {
		body["KeywordIdList"] = request.KeywordIdList
	}

	if !dara.IsNil(request.KeywordIds) {
		body["KeywordIds"] = request.KeywordIds
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKeyword"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKeywordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete keyword
//
// @param request - DeleteKeywordRequest
//
// @return DeleteKeywordResponse
func (client *Client) DeleteKeyword(request *DeleteKeywordRequest) (_result *DeleteKeywordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteKeywordResponse{}
	_body, _err := client.DeleteKeywordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Keyword Library
//
// @param request - DeleteKeywordLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeywordLibResponse
func (client *Client) DeleteKeywordLibWithOptions(request *DeleteKeywordLibRequest, runtime *dara.RuntimeOptions) (_result *DeleteKeywordLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKeywordLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Keyword Library
//
// @param request - DeleteKeywordLibRequest
//
// @return DeleteKeywordLibResponse
func (client *Client) DeleteKeywordLib(request *DeleteKeywordLibRequest) (_result *DeleteKeywordLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteKeywordLibResponse{}
	_body, _err := client.DeleteKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete online test
//
// @param request - DeleteOnlineTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOnlineTestResponse
func (client *Client) DeleteOnlineTestWithOptions(request *DeleteOnlineTestRequest, runtime *dara.RuntimeOptions) (_result *DeleteOnlineTestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOnlineTest"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOnlineTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete online test
//
// @param request - DeleteOnlineTestRequest
//
// @return DeleteOnlineTestResponse
func (client *Client) DeleteOnlineTest(request *DeleteOnlineTestRequest) (_result *DeleteOnlineTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOnlineTestResponse{}
	_body, _err := client.DeleteOnlineTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询在线测试结果
//
// @param request - DescribeOnlineTestResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOnlineTestResultResponse
func (client *Client) DescribeOnlineTestResultWithOptions(request *DescribeOnlineTestResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeOnlineTestResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOnlineTestResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOnlineTestResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询在线测试结果
//
// @param request - DescribeOnlineTestResultRequest
//
// @return DescribeOnlineTestResultResponse
func (client *Client) DescribeOnlineTestResult(request *DescribeOnlineTestResultRequest) (_result *DescribeOnlineTestResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOnlineTestResultResponse{}
	_body, _err := client.DescribeOnlineTestResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出代答答案
//
// @param request - ExportAnswerSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportAnswerSampleResponse
func (client *Client) ExportAnswerSampleWithOptions(request *ExportAnswerSampleRequest, runtime *dara.RuntimeOptions) (_result *ExportAnswerSampleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportAnswerSample"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportAnswerSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出代答答案
//
// @param request - ExportAnswerSampleRequest
//
// @return ExportAnswerSampleResponse
func (client *Client) ExportAnswerSample(request *ExportAnswerSampleRequest) (_result *ExportAnswerSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportAnswerSampleResponse{}
	_body, _err := client.ExportAnswerSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Export Call Volume
//
// @param request - ExportCipStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportCipStatsResponse
func (client *Client) ExportCipStatsWithOptions(request *ExportCipStatsRequest, runtime *dara.RuntimeOptions) (_result *ExportCipStatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ByMonth) {
		body["ByMonth"] = request.ByMonth
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ExportType) {
		body["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubUid) {
		body["SubUid"] = request.SubUid
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportCipStats"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportCipStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Export Call Volume
//
// @param request - ExportCipStatsRequest
//
// @return ExportCipStatsResponse
func (client *Client) ExportCipStats(request *ExportCipStatsRequest) (_result *ExportCipStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportCipStatsResponse{}
	_body, _err := client.ExportCipStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Export Keywords
//
// @param request - ExportKeywordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportKeywordResponse
func (client *Client) ExportKeywordWithOptions(request *ExportKeywordRequest, runtime *dara.RuntimeOptions) (_result *ExportKeywordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportKeyword"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportKeywordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Export Keywords
//
// @param request - ExportKeywordRequest
//
// @return ExportKeywordResponse
func (client *Client) ExportKeyword(request *ExportKeywordRequest) (_result *ExportKeywordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportKeywordResponse{}
	_body, _err := client.ExportKeywordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OSS Usage Statistics Export
//
// @param request - ExportOssCheckStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportOssCheckStatResponse
func (client *Client) ExportOssCheckStatWithOptions(request *ExportOssCheckStatRequest, runtime *dara.RuntimeOptions) (_result *ExportOssCheckStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ByMonth) {
		body["ByMonth"] = request.ByMonth
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ParentTaskId) {
		body["ParentTaskId"] = request.ParentTaskId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportOssCheckStat"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportOssCheckStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OSS Usage Statistics Export
//
// @param request - ExportOssCheckStatRequest
//
// @return ExportOssCheckStatResponse
func (client *Client) ExportOssCheckStat(request *ExportOssCheckStatRequest) (_result *ExportOssCheckStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportOssCheckStatResponse{}
	_body, _err := client.ExportOssCheckStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Export OSS scan results
//
// @param tmpReq - ExportResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportResultResponse
func (client *Client) ExportResultWithOptions(tmpReq *ExportResultRequest, runtime *dara.RuntimeOptions) (_result *ExportResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.SortShrink) {
		body["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.Source) {
		body["Source"] = request.Source
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Export OSS scan results
//
// @param request - ExportResultRequest
//
// @return ExportResultResponse
func (client *Client) ExportResult(request *ExportResultRequest) (_result *ExportResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportResultResponse{}
	_body, _err := client.ExportResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Export scan results, Excel file
//
// @param tmpReq - ExportScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportScanResultResponse
func (client *Client) ExportScanResultWithOptions(tmpReq *ExportScanResultRequest, runtime *dara.RuntimeOptions) (_result *ExportScanResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportScanResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryShrink) {
		body["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SortShrink) {
		body["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportScanResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Export scan results, Excel file
//
// @param request - ExportScanResultRequest
//
// @return ExportScanResultResponse
func (client *Client) ExportScanResult(request *ExportScanResultRequest) (_result *ExportScanResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportScanResultResponse{}
	_body, _err := client.ExportScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Export text scan results, Excel file
//
// @param tmpReq - ExportTextScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportTextScanResultResponse
func (client *Client) ExportTextScanResultWithOptions(tmpReq *ExportTextScanResultRequest, runtime *dara.RuntimeOptions) (_result *ExportTextScanResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportTextScanResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.QueryShrink) {
		body["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportTextScanResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportTextScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Export text scan results, Excel file
//
// @param request - ExportTextScanResultRequest
//
// @return ExportTextScanResultResponse
func (client *Client) ExportTextScanResult(request *ExportTextScanResultRequest) (_result *ExportTextScanResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportTextScanResultResponse{}
	_body, _err := client.ExportTextScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取代答样本导入进度
//
// @param request - GetAnswerImportProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnswerImportProgressResponse
func (client *Client) GetAnswerImportProgressWithOptions(request *GetAnswerImportProgressRequest, runtime *dara.RuntimeOptions) (_result *GetAnswerImportProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAnswerImportProgress"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAnswerImportProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取代答样本导入进度
//
// @param request - GetAnswerImportProgressRequest
//
// @return GetAnswerImportProgressResponse
func (client *Client) GetAnswerImportProgress(request *GetAnswerImportProgressRequest) (_result *GetAnswerImportProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAnswerImportProgressResponse{}
	_body, _err := client.GetAnswerImportProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Evidence Transfer to Get User\\"s Bucket List
//
// @param request - GetBackupBucketsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupBucketsListResponse
func (client *Client) GetBackupBucketsListWithOptions(request *GetBackupBucketsListRequest, runtime *dara.RuntimeOptions) (_result *GetBackupBucketsListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBackupBucketsList"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBackupBucketsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Evidence Transfer to Get User\\"s Bucket List
//
// @param request - GetBackupBucketsListRequest
//
// @return GetBackupBucketsListResponse
func (client *Client) GetBackupBucketsList(request *GetBackupBucketsListRequest) (_result *GetBackupBucketsListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBackupBucketsListResponse{}
	_body, _err := client.GetBackupBucketsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Evidence Backup Configuration
//
// @param request - GetBackupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupConfigResponse
func (client *Client) GetBackupConfigWithOptions(request *GetBackupConfigRequest, runtime *dara.RuntimeOptions) (_result *GetBackupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBackupConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBackupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Evidence Backup Configuration
//
// @param request - GetBackupConfigRequest
//
// @return GetBackupConfigResponse
func (client *Client) GetBackupConfig(request *GetBackupConfigRequest) (_result *GetBackupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBackupConfigResponse{}
	_body, _err := client.GetBackupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # User Backup Authorization Verification
//
// @param request - GetBackupStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupStatusResponse
func (client *Client) GetBackupStatusWithOptions(request *GetBackupStatusRequest, runtime *dara.RuntimeOptions) (_result *GetBackupStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBackupStatus"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBackupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # User Backup Authorization Verification
//
// @param request - GetBackupStatusRequest
//
// @return GetBackupStatusResponse
func (client *Client) GetBackupStatus(request *GetBackupStatusRequest) (_result *GetBackupStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBackupStatusResponse{}
	_body, _err := client.GetBackupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get User OSS Scan Bucket List
//
// @param request - GetBucketsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBucketsListResponse
func (client *Client) GetBucketsListWithOptions(request *GetBucketsListRequest, runtime *dara.RuntimeOptions) (_result *GetBucketsListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBucketsList"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBucketsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get User OSS Scan Bucket List
//
// @param request - GetBucketsListRequest
//
// @return GetBucketsListResponse
func (client *Client) GetBucketsList(request *GetBucketsListRequest) (_result *GetBucketsListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBucketsListResponse{}
	_body, _err := client.GetBucketsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Call Volume
//
// @param request - GetCipStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCipStatsResponse
func (client *Client) GetCipStatsWithOptions(request *GetCipStatsRequest, runtime *dara.RuntimeOptions) (_result *GetCipStatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ByMonth) {
		body["ByMonth"] = request.ByMonth
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubUid) {
		body["SubUid"] = request.SubUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCipStats"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCipStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Call Volume
//
// @param request - GetCipStatsRequest
//
// @return GetCipStatsResponse
func (client *Client) GetCipStats(request *GetCipStatsRequest) (_result *GetCipStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCipStatsResponse{}
	_body, _err := client.GetCipStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Scheduled  OSS Scan  Task Estimated Execution Time
//
// @param request - GetExecuteTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecuteTimeResponse
func (client *Client) GetExecuteTimeWithOptions(request *GetExecuteTimeRequest, runtime *dara.RuntimeOptions) (_result *GetExecuteTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExecuteTime"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExecuteTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Scheduled  OSS Scan  Task Estimated Execution Time
//
// @param request - GetExecuteTimeRequest
//
// @return GetExecuteTimeResponse
func (client *Client) GetExecuteTime(request *GetExecuteTimeRequest) (_result *GetExecuteTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExecuteTimeResponse{}
	_body, _err := client.GetExecuteTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Feature Configuration
//
// @param request - GetFeatureConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureConfigResponse
func (client *Client) GetFeatureConfigWithOptions(request *GetFeatureConfigRequest, runtime *dara.RuntimeOptions) (_result *GetFeatureConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFeatureConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFeatureConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Feature Configuration
//
// @param request - GetFeatureConfigRequest
//
// @return GetFeatureConfigResponse
func (client *Client) GetFeatureConfig(request *GetFeatureConfigRequest) (_result *GetFeatureConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFeatureConfigResponse{}
	_body, _err := client.GetFeatureConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Image Rule Label Information
//
// @param request - GetImageSceneLabelConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageSceneLabelConfResponse
func (client *Client) GetImageSceneLabelConfWithOptions(request *GetImageSceneLabelConfRequest, runtime *dara.RuntimeOptions) (_result *GetImageSceneLabelConfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageSceneLabelConf"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageSceneLabelConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Image Rule Label Information
//
// @param request - GetImageSceneLabelConfRequest
//
// @return GetImageSceneLabelConfResponse
func (client *Client) GetImageSceneLabelConf(request *GetImageSceneLabelConfRequest) (_result *GetImageSceneLabelConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageSceneLabelConfResponse{}
	_body, _err := client.GetImageSceneLabelConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Image Rule Label Information
//
// @param request - GetImageSceneLabelListConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageSceneLabelListConfResponse
func (client *Client) GetImageSceneLabelListConfWithOptions(request *GetImageSceneLabelListConfRequest, runtime *dara.RuntimeOptions) (_result *GetImageSceneLabelListConfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageServiceCode) {
		query["ImageServiceCode"] = request.ImageServiceCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageSceneLabelListConf"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageSceneLabelListConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Image Rule Label Information
//
// @param request - GetImageSceneLabelListConfRequest
//
// @return GetImageSceneLabelListConfResponse
func (client *Client) GetImageSceneLabelListConf(request *GetImageSceneLabelListConfRequest) (_result *GetImageSceneLabelListConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageSceneLabelListConfResponse{}
	_body, _err := client.GetImageSceneLabelListConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OSS scheduled scan detection cycle query
//
// @param tmpReq - GetJobNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobNameListResponse
func (client *Client) GetJobNameListWithOptions(tmpReq *GetJobNameListRequest, runtime *dara.RuntimeOptions) (_result *GetJobNameListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetJobNameListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortShrink) {
		query["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobNameList"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OSS scheduled scan detection cycle query
//
// @param request - GetJobNameListRequest
//
// @return GetJobNameListResponse
func (client *Client) GetJobNameList(request *GetJobNameListRequest) (_result *GetJobNameListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobNameListResponse{}
	_body, _err := client.GetJobNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the result of keyword import
//
// @param request - GetKeywordImportResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKeywordImportResultResponse
func (client *Client) GetKeywordImportResultWithOptions(request *GetKeywordImportResultRequest, runtime *dara.RuntimeOptions) (_result *GetKeywordImportResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKeywordImportResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKeywordImportResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the result of keyword import
//
// @param request - GetKeywordImportResultRequest
//
// @return GetKeywordImportResultResponse
func (client *Client) GetKeywordImportResult(request *GetKeywordImportResultRequest) (_result *GetKeywordImportResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetKeywordImportResultResponse{}
	_body, _err := client.GetKeywordImportResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Keyword Library Information
//
// @param request - GetKeywordLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKeywordLibResponse
func (client *Client) GetKeywordLibWithOptions(request *GetKeywordLibRequest, runtime *dara.RuntimeOptions) (_result *GetKeywordLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKeywordLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Keyword Library Information
//
// @param request - GetKeywordLibRequest
//
// @return GetKeywordLibResponse
func (client *Client) GetKeywordLib(request *GetKeywordLibRequest) (_result *GetKeywordLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetKeywordLibResponse{}
	_body, _err := client.GetKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query OSS freeze result
//
// @param tmpReq - GetOssCheckFreezeResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssCheckFreezeResultResponse
func (client *Client) GetOssCheckFreezeResultWithOptions(tmpReq *GetOssCheckFreezeResultRequest, runtime *dara.RuntimeOptions) (_result *GetOssCheckFreezeResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetOssCheckFreezeResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.FinishNum) {
		query["FinishNum"] = request.FinishNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortShrink) {
		query["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssCheckFreezeResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssCheckFreezeResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query OSS freeze result
//
// @param request - GetOssCheckFreezeResultRequest
//
// @return GetOssCheckFreezeResultResponse
func (client *Client) GetOssCheckFreezeResult(request *GetOssCheckFreezeResultRequest) (_result *GetOssCheckFreezeResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssCheckFreezeResultResponse{}
	_body, _err := client.GetOssCheckFreezeResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OSS result details
//
// @param request - GetOssCheckResultDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssCheckResultDetailResponse
func (client *Client) GetOssCheckResultDetailWithOptions(request *GetOssCheckResultDetailRequest, runtime *dara.RuntimeOptions) (_result *GetOssCheckResultDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Object) {
		query["Object"] = request.Object
	}

	if !dara.IsNil(request.ParentTaskId) {
		query["ParentTaskId"] = request.ParentTaskId
	}

	if !dara.IsNil(request.QueryRequestId) {
		query["QueryRequestId"] = request.QueryRequestId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssCheckResultDetail"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssCheckResultDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OSS result details
//
// @param request - GetOssCheckResultDetailRequest
//
// @return GetOssCheckResultDetailResponse
func (client *Client) GetOssCheckResultDetail(request *GetOssCheckResultDetailRequest) (_result *GetOssCheckResultDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssCheckResultDetailResponse{}
	_body, _err := client.GetOssCheckResultDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OSS Check Usage Statistics
//
// @param request - GetOssCheckStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssCheckStatResponse
func (client *Client) GetOssCheckStatWithOptions(request *GetOssCheckStatRequest, runtime *dara.RuntimeOptions) (_result *GetOssCheckStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ByMonth) {
		body["ByMonth"] = request.ByMonth
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ParentTaskId) {
		body["ParentTaskId"] = request.ParentTaskId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssCheckStat"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssCheckStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OSS Check Usage Statistics
//
// @param request - GetOssCheckStatRequest
//
// @return GetOssCheckStatResponse
func (client *Client) GetOssCheckStat(request *GetOssCheckStatRequest) (_result *GetOssCheckStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssCheckStatResponse{}
	_body, _err := client.GetOssCheckStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get User OSS check user status
//
// @param request - GetOssCheckStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssCheckStatusResponse
func (client *Client) GetOssCheckStatusWithOptions(request *GetOssCheckStatusRequest, runtime *dara.RuntimeOptions) (_result *GetOssCheckStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssCheckStatus"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssCheckStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get User OSS check user status
//
// @param request - GetOssCheckStatusRequest
//
// @return GetOssCheckStatusResponse
func (client *Client) GetOssCheckStatus(request *GetOssCheckStatusRequest) (_result *GetOssCheckStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssCheckStatusResponse{}
	_body, _err := client.GetOssCheckStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询oss扫描任务详情
//
// @param request - GetOssCheckTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssCheckTaskInfoResponse
func (client *Client) GetOssCheckTaskInfoWithOptions(request *GetOssCheckTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetOssCheckTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParentTaskId) {
		query["ParentTaskId"] = request.ParentTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssCheckTaskInfo"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssCheckTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询oss扫描任务详情
//
// @param request - GetOssCheckTaskInfoRequest
//
// @return GetOssCheckTaskInfoResponse
func (client *Client) GetOssCheckTaskInfo(request *GetOssCheckTaskInfoRequest) (_result *GetOssCheckTaskInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssCheckTaskInfoResponse{}
	_body, _err := client.GetOssCheckTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # User OSS Check Task Pending Inspection Information
//
// @param request - GetScanNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScanNumResponse
func (client *Client) GetScanNumWithOptions(request *GetScanNumRequest, runtime *dara.RuntimeOptions) (_result *GetScanNumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Buckets) {
		query["Buckets"] = request.Buckets
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScanNum"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScanNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # User OSS Check Task Pending Inspection Information
//
// @param request - GetScanNumRequest
//
// @return GetScanNumResponse
func (client *Client) GetScanNum(request *GetScanNumRequest) (_result *GetScanNumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScanNumResponse{}
	_body, _err := client.GetScanNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the Scan results
//
// @param tmpReq - GetScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScanResultResponse
func (client *Client) GetScanResultWithOptions(tmpReq *GetScanResultRequest, runtime *dara.RuntimeOptions) (_result *GetScanResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetScanResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryShrink) {
		body["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SortShrink) {
		body["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScanResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the Scan results
//
// @param request - GetScanResultRequest
//
// @return GetScanResultResponse
func (client *Client) GetScanResult(request *GetScanResultRequest) (_result *GetScanResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScanResultResponse{}
	_body, _err := client.GetScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get a Single Service Configuration
//
// @param request - GetServiceConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceConfResponse
func (client *Client) GetServiceConfWithOptions(request *GetServiceConfRequest, runtime *dara.RuntimeOptions) (_result *GetServiceConfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ByDefault) {
		body["ByDefault"] = request.ByDefault
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceConf"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get a Single Service Configuration
//
// @param request - GetServiceConfRequest
//
// @return GetServiceConfResponse
func (client *Client) GetServiceConf(request *GetServiceConfRequest) (_result *GetServiceConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceConfResponse{}
	_body, _err := client.GetServiceConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get a Single Service Configuration
//
// @param request - GetServiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceConfigResponse
func (client *Client) GetServiceConfigWithOptions(request *GetServiceConfigRequest, runtime *dara.RuntimeOptions) (_result *GetServiceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// # Get a Single Service Configuration
//
// @param request - GetServiceConfigRequest
//
// @return GetServiceConfigResponse
func (client *Client) GetServiceConfig(request *GetServiceConfigRequest) (_result *GetServiceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceConfigResponse{}
	_body, _err := client.GetServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get the label configuration of a single service
//
// @param request - GetServiceLabelConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceLabelConfigResponse
func (client *Client) GetServiceLabelConfigWithOptions(request *GetServiceLabelConfigRequest, runtime *dara.RuntimeOptions) (_result *GetServiceLabelConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceLabelConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceLabelConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the label configuration of a single service
//
// @param request - GetServiceLabelConfigRequest
//
// @return GetServiceLabelConfigResponse
func (client *Client) GetServiceLabelConfig(request *GetServiceLabelConfigRequest) (_result *GetServiceLabelConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceLabelConfigResponse{}
	_body, _err := client.GetServiceLabelConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query OSS Scan Task List
//
// @param tmpReq - GetStockOssCheckTasksListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStockOssCheckTasksListResponse
func (client *Client) GetStockOssCheckTasksListWithOptions(tmpReq *GetStockOssCheckTasksListRequest, runtime *dara.RuntimeOptions) (_result *GetStockOssCheckTasksListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetStockOssCheckTasksListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IsInc) {
		query["IsInc"] = request.IsInc
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MediaType) {
		body["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortShrink) {
		body["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStockOssCheckTasksList"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStockOssCheckTasksListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query OSS Scan Task List
//
// @param request - GetStockOssCheckTasksListRequest
//
// @return GetStockOssCheckTasksListResponse
func (client *Client) GetStockOssCheckTasksList(request *GetStockOssCheckTasksListRequest) (_result *GetStockOssCheckTasksListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStockOssCheckTasksListResponse{}
	_body, _err := client.GetStockOssCheckTasksListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the invocation result
//
// @param tmpReq - GetTextScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextScanResultResponse
func (client *Client) GetTextScanResultWithOptions(tmpReq *GetTextScanResultRequest, runtime *dara.RuntimeOptions) (_result *GetTextScanResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTextScanResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryShrink) {
		body["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.SortShrink) {
		body["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextScanResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the invocation result
//
// @param request - GetTextScanResultRequest
//
// @return GetTextScanResultResponse
func (client *Client) GetTextScanResult(request *GetTextScanResultRequest) (_result *GetTextScanResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTextScanResultResponse{}
	_body, _err := client.GetTextScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get the corresponding information for file upload
//
// @param request - GetUploadInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadInfoResponse
func (client *Client) GetUploadInfoWithOptions(request *GetUploadInfoRequest, runtime *dara.RuntimeOptions) (_result *GetUploadInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadInfo"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the corresponding information for file upload
//
// @param request - GetUploadInfoRequest
//
// @return GetUploadInfoResponse
func (client *Client) GetUploadInfo(request *GetUploadInfoRequest) (_result *GetUploadInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUploadInfoResponse{}
	_body, _err := client.GetUploadInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取上传链接
//
// @param request - GetUploadLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadLinkResponse
func (client *Client) GetUploadLinkWithOptions(request *GetUploadLinkRequest, runtime *dara.RuntimeOptions) (_result *GetUploadLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UploadUrl) {
		query["UploadUrl"] = request.UploadUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadLink"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取上传链接
//
// @param request - GetUploadLinkRequest
//
// @return GetUploadLinkResponse
func (client *Client) GetUploadLink(request *GetUploadLinkRequest) (_result *GetUploadLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUploadLinkResponse{}
	_body, _err := client.GetUploadLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get User Purchase Status
//
// @param request - GetUserBuyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBuyStatusResponse
func (client *Client) GetUserBuyStatusWithOptions(request *GetUserBuyStatusRequest, runtime *dara.RuntimeOptions) (_result *GetUserBuyStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserBuyStatus"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserBuyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get User Purchase Status
//
// @param request - GetUserBuyStatusRequest
//
// @return GetUserBuyStatusResponse
func (client *Client) GetUserBuyStatus(request *GetUserBuyStatusRequest) (_result *GetUserBuyStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserBuyStatusResponse{}
	_body, _err := client.GetUserBuyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 代答库列表
//
// @param request - ListAnswerLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnswerLibResponse
func (client *Client) ListAnswerLibWithOptions(request *ListAnswerLibRequest, runtime *dara.RuntimeOptions) (_result *ListAnswerLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnswerLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnswerLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 代答库列表
//
// @param request - ListAnswerLibRequest
//
// @return ListAnswerLibResponse
func (client *Client) ListAnswerLib(request *ListAnswerLibRequest) (_result *ListAnswerLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAnswerLibResponse{}
	_body, _err := client.ListAnswerLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Callback List
//
// @param request - ListCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallbackResponse
func (client *Client) ListCallbackWithOptions(request *ListCallbackRequest, runtime *dara.RuntimeOptions) (_result *ListCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCallback"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Callback List
//
// @param request - ListCallbackRequest
//
// @return ListCallbackResponse
func (client *Client) ListCallback(request *ListCallbackRequest) (_result *ListCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCallbackResponse{}
	_body, _err := client.ListCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Image Library List
//
// @param request - ListImageLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageLibResponse
func (client *Client) ListImageLibWithOptions(request *ListImageLibRequest, runtime *dara.RuntimeOptions) (_result *ListImageLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImageLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Image Library List
//
// @param request - ListImageLibRequest
//
// @return ListImageLibResponse
func (client *Client) ListImageLib(request *ListImageLibRequest) (_result *ListImageLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListImageLibResponse{}
	_body, _err := client.ListImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Paged Image List
//
// @param tmpReq - ListImagesFromLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesFromLibResponse
func (client *Client) ListImagesFromLibWithOptions(tmpReq *ListImagesFromLibRequest, runtime *dara.RuntimeOptions) (_result *ListImagesFromLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListImagesFromLibShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ImgId) {
		body["ImgId"] = request.ImgId
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortShrink) {
		body["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImagesFromLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImagesFromLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Paged Image List
//
// @param request - ListImagesFromLibRequest
//
// @return ListImagesFromLibResponse
func (client *Client) ListImagesFromLib(request *ListImagesFromLibRequest) (_result *ListImagesFromLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListImagesFromLibResponse{}
	_body, _err := client.ListImagesFromLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Keyword Library List
//
// @param request - ListKeywordLibsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeywordLibsResponse
func (client *Client) ListKeywordLibsWithOptions(request *ListKeywordLibsRequest, runtime *dara.RuntimeOptions) (_result *ListKeywordLibsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKeywordLibs"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKeywordLibsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Keyword Library List
//
// @param request - ListKeywordLibsRequest
//
// @return ListKeywordLibsResponse
func (client *Client) ListKeywordLibs(request *ListKeywordLibsRequest) (_result *ListKeywordLibsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListKeywordLibsResponse{}
	_body, _err := client.ListKeywordLibsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Keyword List
//
// @param tmpReq - ListKeywordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeywordsResponse
func (client *Client) ListKeywordsWithOptions(tmpReq *ListKeywordsRequest, runtime *dara.RuntimeOptions) (_result *ListKeywordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListKeywordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortShrink) {
		body["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.Word) {
		body["Word"] = request.Word
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKeywords"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKeywordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Keyword List
//
// @param request - ListKeywordsRequest
//
// @return ListKeywordsResponse
func (client *Client) ListKeywords(request *ListKeywordsRequest) (_result *ListKeywordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListKeywordsResponse{}
	_body, _err := client.ListKeywordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// query OSS scan result list
//
// @param tmpReq - ListOssCheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOssCheckResultResponse
func (client *Client) ListOssCheckResultWithOptions(tmpReq *ListOssCheckResultRequest, runtime *dara.RuntimeOptions) (_result *ListOssCheckResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListOssCheckResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.FinishNum) {
		query["FinishNum"] = request.FinishNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortShrink) {
		query["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOssCheckResult"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOssCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// query OSS scan result list
//
// @param request - ListOssCheckResultRequest
//
// @return ListOssCheckResultResponse
func (client *Client) ListOssCheckResult(request *ListOssCheckResultRequest) (_result *ListOssCheckResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOssCheckResultResponse{}
	_body, _err := client.ListOssCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Service List
//
// @param request - ListServiceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceConfigsResponse
func (client *Client) ListServiceConfigsWithOptions(request *ListServiceConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListServiceConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Classify) {
		query["Classify"] = request.Classify
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UseStatus) {
		query["UseStatus"] = request.UseStatus
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceConfigs"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Service List
//
// @param request - ListServiceConfigsRequest
//
// @return ListServiceConfigsResponse
func (client *Client) ListServiceConfigs(request *ListServiceConfigsRequest) (_result *ListServiceConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceConfigsResponse{}
	_body, _err := client.ListServiceConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Use SSE interface to stream large model calls
//
// @param request - LlmStreamChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LlmStreamChatResponse
func (client *Client) LlmStreamChatWithSSE(request *LlmStreamChatRequest, runtime *dara.RuntimeOptions, _yield chan *LlmStreamChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.llmStreamChatWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// # Use SSE interface to stream large model calls
//
// @param request - LlmStreamChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LlmStreamChatResponse
func (client *Client) LlmStreamChatWithOptions(request *LlmStreamChatRequest, runtime *dara.RuntimeOptions) (_result *LlmStreamChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.Messages) {
		body["Messages"] = request.Messages
	}

	if !dara.IsNil(request.Temperature) {
		body["Temperature"] = request.Temperature
	}

	if !dara.IsNil(request.TopP) {
		body["TopP"] = request.TopP
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LlmStreamChat"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LlmStreamChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Use SSE interface to stream large model calls
//
// @param request - LlmStreamChatRequest
//
// @return LlmStreamChatResponse
func (client *Client) LlmStreamChat(request *LlmStreamChatRequest) (_result *LlmStreamChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LlmStreamChatResponse{}
	_body, _err := client.LlmStreamChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新代答库
//
// @param request - ModifyAnswerLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAnswerLibResponse
func (client *Client) ModifyAnswerLibWithOptions(request *ModifyAnswerLibRequest, runtime *dara.RuntimeOptions) (_result *ModifyAnswerLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.LibName) {
		query["LibName"] = request.LibName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAnswerLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAnswerLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代答库
//
// @param request - ModifyAnswerLibRequest
//
// @return ModifyAnswerLibResponse
func (client *Client) ModifyAnswerLib(request *ModifyAnswerLibRequest) (_result *ModifyAnswerLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAnswerLibResponse{}
	_body, _err := client.ModifyAnswerLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Message Notification
//
// @param request - ModifyCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCallbackResponse
func (client *Client) ModifyCallbackWithOptions(request *ModifyCallbackRequest, runtime *dara.RuntimeOptions) (_result *ModifyCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CryptType) {
		body["CryptType"] = request.CryptType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Scope) {
		body["Scope"] = request.Scope
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCallback"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Message Notification
//
// @param request - ModifyCallbackRequest
//
// @return ModifyCallbackResponse
func (client *Client) ModifyCallback(request *ModifyCallbackRequest) (_result *ModifyCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCallbackResponse{}
	_body, _err := client.ModifyCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Save Feature Configuration
//
// @param request - ModifyFeatureConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFeatureConfigResponse
func (client *Client) ModifyFeatureConfigWithOptions(request *ModifyFeatureConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyFeatureConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Field) {
		body["Field"] = request.Field
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFeatureConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFeatureConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Save Feature Configuration
//
// @param request - ModifyFeatureConfigRequest
//
// @return ModifyFeatureConfigResponse
func (client *Client) ModifyFeatureConfig(request *ModifyFeatureConfigRequest) (_result *ModifyFeatureConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFeatureConfigResponse{}
	_body, _err := client.ModifyFeatureConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit Service
//
// @param request - ModifyServiceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyServiceInfoResponse
func (client *Client) ModifyServiceInfoWithOptions(request *ModifyServiceInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyServiceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.ServiceDesc) {
		body["ServiceDesc"] = request.ServiceDesc
	}

	if !dara.IsNil(request.ServiceName) {
		body["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyServiceInfo"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyServiceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Service
//
// @param request - ModifyServiceInfoRequest
//
// @return ModifyServiceInfoResponse
func (client *Client) ModifyServiceInfo(request *ModifyServiceInfoRequest) (_result *ModifyServiceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyServiceInfoResponse{}
	_body, _err := client.ModifyServiceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # OSS scan result query
//
// @param tmpReq - OssCheckResultListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OssCheckResultListResponse
func (client *Client) OssCheckResultListWithOptions(tmpReq *OssCheckResultListRequest, runtime *dara.RuntimeOptions) (_result *OssCheckResultListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OssCheckResultListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.FinishNum) {
		query["FinishNum"] = request.FinishNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortShrink) {
		query["Sort"] = request.SortShrink
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OssCheckResultList"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OssCheckResultListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OSS scan result query
//
// @param request - OssCheckResultListRequest
//
// @return OssCheckResultListResponse
func (client *Client) OssCheckResultList(request *OssCheckResultListRequest) (_result *OssCheckResultListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OssCheckResultListResponse{}
	_body, _err := client.OssCheckResultListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询代答样本
//
// @param tmpReq - QueryAnswerSampleByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAnswerSampleByPageResponse
func (client *Client) QueryAnswerSampleByPageWithOptions(tmpReq *QueryAnswerSampleByPageRequest, runtime *dara.RuntimeOptions) (_result *QueryAnswerSampleByPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryAnswerSampleByPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sort) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, dara.String("Sort"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Answer) {
		query["Answer"] = request.Answer
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortShrink) {
		query["Sort"] = request.SortShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAnswerSampleByPage"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAnswerSampleByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询代答样本
//
// @param request - QueryAnswerSampleByPageRequest
//
// @return QueryAnswerSampleByPageResponse
func (client *Client) QueryAnswerSampleByPage(request *QueryAnswerSampleByPageRequest) (_result *QueryAnswerSampleByPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAnswerSampleByPageResponse{}
	_body, _err := client.QueryAnswerSampleByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query a Single Callback Configuration
//
// @param request - QueryCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallbackResponse
func (client *Client) QueryCallbackWithOptions(request *QueryCallbackRequest, runtime *dara.RuntimeOptions) (_result *QueryCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckForOss) {
		body["CheckForOss"] = request.CheckForOss
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCallback"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a Single Callback Configuration
//
// @param request - QueryCallbackRequest
//
// @return QueryCallbackResponse
func (client *Client) QueryCallback(request *QueryCallbackRequest) (_result *QueryCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCallbackResponse{}
	_body, _err := client.QueryCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Paginated Query of Message Notification List
//
// @param request - QueryCallbackByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallbackByPageResponse
func (client *Client) QueryCallbackByPageWithOptions(request *QueryCallbackByPageRequest, runtime *dara.RuntimeOptions) (_result *QueryCallbackByPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCallbackByPage"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCallbackByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Paginated Query of Message Notification List
//
// @param request - QueryCallbackByPageRequest
//
// @return QueryCallbackByPageResponse
func (client *Client) QueryCallbackByPage(request *QueryCallbackByPageRequest) (_result *QueryCallbackByPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCallbackByPageResponse{}
	_body, _err := client.QueryCallbackByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止在线测试
//
// @param request - StopOnlineTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopOnlineTestResponse
func (client *Client) StopOnlineTestWithOptions(request *StopOnlineTestRequest, runtime *dara.RuntimeOptions) (_result *StopOnlineTestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopOnlineTest"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopOnlineTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止在线测试
//
// @param request - StopOnlineTestRequest
//
// @return StopOnlineTestResponse
func (client *Client) StopOnlineTest(request *StopOnlineTestRequest) (_result *StopOnlineTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopOnlineTestResponse{}
	_body, _err := client.StopOnlineTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Evidence Backup Configuration
//
// @param request - UpdateBackupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupConfigResponse
func (client *Client) UpdateBackupConfigWithOptions(request *UpdateBackupConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateBackupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupConfig) {
		query["BackupConfig"] = request.BackupConfig
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBackupConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBackupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Evidence Backup Configuration
//
// @param request - UpdateBackupConfigRequest
//
// @return UpdateBackupConfigResponse
func (client *Client) UpdateBackupConfig(request *UpdateBackupConfigRequest) (_result *UpdateBackupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBackupConfigResponse{}
	_body, _err := client.UpdateBackupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit Image Library
//
// @param request - UpdateImageLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageLibResponse
func (client *Client) UpdateImageLibWithOptions(request *UpdateImageLibRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.FreeInspection) {
		body["FreeInspection"] = request.FreeInspection
	}

	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	if !dara.IsNil(request.LibName) {
		body["LibName"] = request.LibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImageLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Image Library
//
// @param request - UpdateImageLibRequest
//
// @return UpdateImageLibResponse
func (client *Client) UpdateImageLib(request *UpdateImageLibRequest) (_result *UpdateImageLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateImageLibResponse{}
	_body, _err := client.UpdateImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit Image Library Free Inspection Configuration
//
// @param tmpReq - UpdateImageLibFreeInspectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageLibFreeInspectionResponse
func (client *Client) UpdateImageLibFreeInspectionWithOptions(tmpReq *UpdateImageLibFreeInspectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageLibFreeInspectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateImageLibFreeInspectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigShrink) {
		body["Config"] = request.ConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImageLibFreeInspection"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageLibFreeInspectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Image Library Free Inspection Configuration
//
// @param request - UpdateImageLibFreeInspectionRequest
//
// @return UpdateImageLibFreeInspectionResponse
func (client *Client) UpdateImageLibFreeInspection(request *UpdateImageLibFreeInspectionRequest) (_result *UpdateImageLibFreeInspectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateImageLibFreeInspectionResponse{}
	_body, _err := client.UpdateImageLibFreeInspectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit Keyword Library
//
// @param request - UpdateKeywordLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKeywordLibResponse
func (client *Client) UpdateKeywordLibWithOptions(request *UpdateKeywordLibRequest, runtime *dara.RuntimeOptions) (_result *UpdateKeywordLibResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LibId) {
		body["LibId"] = request.LibId
	}

	if !dara.IsNil(request.LibName) {
		body["LibName"] = request.LibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKeywordLib"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Keyword Library
//
// @param request - UpdateKeywordLibRequest
//
// @return UpdateKeywordLibResponse
func (client *Client) UpdateKeywordLib(request *UpdateKeywordLibRequest) (_result *UpdateKeywordLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateKeywordLibResponse{}
	_body, _err := client.UpdateKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量反馈任务
//
// @param request - UpdateOssCheckResultsBatchFeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOssCheckResultsBatchFeedbackResponse
func (client *Client) UpdateOssCheckResultsBatchFeedbackWithOptions(request *UpdateOssCheckResultsBatchFeedbackRequest, runtime *dara.RuntimeOptions) (_result *UpdateOssCheckResultsBatchFeedbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Feedback) {
		query["Feedback"] = request.Feedback
	}

	if !dara.IsNil(request.Items) {
		query["Items"] = request.Items
	}

	if !dara.IsNil(request.ParentTaskId) {
		query["ParentTaskId"] = request.ParentTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOssCheckResultsBatchFeedback"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOssCheckResultsBatchFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量反馈任务
//
// @param request - UpdateOssCheckResultsBatchFeedbackRequest
//
// @return UpdateOssCheckResultsBatchFeedbackResponse
func (client *Client) UpdateOssCheckResultsBatchFeedback(request *UpdateOssCheckResultsBatchFeedbackRequest) (_result *UpdateOssCheckResultsBatchFeedbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOssCheckResultsBatchFeedbackResponse{}
	_body, _err := client.UpdateOssCheckResultsBatchFeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// oss结果反馈
//
// @param request - UpdateOssCheckResultsFeedBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOssCheckResultsFeedBackResponse
func (client *Client) UpdateOssCheckResultsFeedBackWithOptions(request *UpdateOssCheckResultsFeedBackRequest, runtime *dara.RuntimeOptions) (_result *UpdateOssCheckResultsFeedBackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Feedback) {
		query["Feedback"] = request.Feedback
	}

	if !dara.IsNil(request.QueryRequestId) {
		query["QueryRequestId"] = request.QueryRequestId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOssCheckResultsFeedBack"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOssCheckResultsFeedBackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// oss结果反馈
//
// @param request - UpdateOssCheckResultsFeedBackRequest
//
// @return UpdateOssCheckResultsFeedBackResponse
func (client *Client) UpdateOssCheckResultsFeedBack(request *UpdateOssCheckResultsFeedBackRequest) (_result *UpdateOssCheckResultsFeedBackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOssCheckResultsFeedBackResponse{}
	_body, _err := client.UpdateOssCheckResultsFeedBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量冻结任务
//
// @param request - UpdateOssCheckResultsFreezeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOssCheckResultsFreezeResponse
func (client *Client) UpdateOssCheckResultsFreezeWithOptions(request *UpdateOssCheckResultsFreezeRequest, runtime *dara.RuntimeOptions) (_result *UpdateOssCheckResultsFreezeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.FreezeItems) {
		query["FreezeItems"] = request.FreezeItems
	}

	if !dara.IsNil(request.FreezeRestorePath) {
		query["FreezeRestorePath"] = request.FreezeRestorePath
	}

	if !dara.IsNil(request.FreezeType) {
		query["FreezeType"] = request.FreezeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOssCheckResultsFreeze"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOssCheckResultsFreezeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量冻结任务
//
// @param request - UpdateOssCheckResultsFreezeRequest
//
// @return UpdateOssCheckResultsFreezeResponse
func (client *Client) UpdateOssCheckResultsFreeze(request *UpdateOssCheckResultsFreezeRequest) (_result *UpdateOssCheckResultsFreezeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOssCheckResultsFreezeResponse{}
	_body, _err := client.UpdateOssCheckResultsFreezeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量解冻任务
//
// @param request - UpdateOssCheckResultsUnfreezeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOssCheckResultsUnfreezeResponse
func (client *Client) UpdateOssCheckResultsUnfreezeWithOptions(request *UpdateOssCheckResultsUnfreezeRequest, runtime *dara.RuntimeOptions) (_result *UpdateOssCheckResultsUnfreezeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.FreezeItems) {
		query["FreezeItems"] = request.FreezeItems
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOssCheckResultsUnfreeze"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOssCheckResultsUnfreezeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量解冻任务
//
// @param request - UpdateOssCheckResultsUnfreezeRequest
//
// @return UpdateOssCheckResultsUnfreezeResponse
func (client *Client) UpdateOssCheckResultsUnfreeze(request *UpdateOssCheckResultsUnfreezeRequest) (_result *UpdateOssCheckResultsUnfreezeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOssCheckResultsUnfreezeResponse{}
	_body, _err := client.UpdateOssCheckResultsUnfreezeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Feedback on Scan Results
//
// @param request - UpdateScanResultFeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScanResultFeedbackResponse
func (client *Client) UpdateScanResultFeedbackWithOptions(request *UpdateScanResultFeedbackRequest, runtime *dara.RuntimeOptions) (_result *UpdateScanResultFeedbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Feedback) {
		body["Feedback"] = request.Feedback
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.QueryRequestId) {
		body["QueryRequestId"] = request.QueryRequestId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RiskLevel) {
		body["RiskLevel"] = request.RiskLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScanResultFeedback"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScanResultFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Feedback on Scan Results
//
// @param request - UpdateScanResultFeedbackRequest
//
// @return UpdateScanResultFeedbackResponse
func (client *Client) UpdateScanResultFeedback(request *UpdateScanResultFeedbackRequest) (_result *UpdateScanResultFeedbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateScanResultFeedbackResponse{}
	_body, _err := client.UpdateScanResultFeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新服务
//
// @param request - UpdateServiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceConfigResponse
func (client *Client) UpdateServiceConfigWithOptions(request *UpdateServiceConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileConfig) {
		body["FileConfig"] = request.FileConfig
	}

	if !dara.IsNil(request.KeywordFilterLibs) {
		body["KeywordFilterLibs"] = request.KeywordFilterLibs
	}

	if !dara.IsNil(request.KeywordHitLibs) {
		body["KeywordHitLibs"] = request.KeywordHitLibs
	}

	if !dara.IsNil(request.ManualMachineConfig) {
		body["ManualMachineConfig"] = request.ManualMachineConfig
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SceneConfig) {
		body["SceneConfig"] = request.SceneConfig
	}

	if !dara.IsNil(request.ServiceCode) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.ServiceConfig) {
		body["ServiceConfig"] = request.ServiceConfig
	}

	if !dara.IsNil(request.VideoConfig) {
		body["VideoConfig"] = request.VideoConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceConfig"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务
//
// @param request - UpdateServiceConfigRequest
//
// @return UpdateServiceConfigResponse
func (client *Client) UpdateServiceConfig(request *UpdateServiceConfigRequest) (_result *UpdateServiceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceConfigResponse{}
	_body, _err := client.UpdateServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) llmStreamChatWithSSE_opYieldFunc(_yield chan *LlmStreamChatResponse, _yieldErr chan error, request *LlmStreamChatRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.Messages) {
		body["Messages"] = request.Messages
	}

	if !dara.IsNil(request.Temperature) {
		body["Temperature"] = request.Temperature
	}

	if !dara.IsNil(request.TopP) {
		body["TopP"] = request.TopP
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LlmStreamChat"),
		Version:     dara.String("2022-09-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
