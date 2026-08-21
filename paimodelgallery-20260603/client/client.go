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
		"cn-shenzhen":    dara.String("paimodelgallery.cn-shenzhen.aliyuncs.com"),
		"cn-wulanchabu":  dara.String("paimodelgallery.cn-wulanchabu.aliyuncs.com"),
		"cn-beijing":     dara.String("paimodelgallery.cn-beijing.aliyuncs.com"),
		"ap-northeast-2": dara.String("paimodelgallery.ap-northeast-2.aliyuncs.com"),
		"ap-northeast-1": dara.String("paimodelgallery.ap-northeast-1.aliyuncs.com"),
		"cn-shanghai":    dara.String("paimodelgallery.cn-shanghai.aliyuncs.com"),
		"cn-guangzhou":   dara.String("paimodelgallery.cn-guangzhou.aliyuncs.com"),
		"cn-hongkong":    dara.String("paimodelgallery.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1": dara.String("paimodelgallery.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3": dara.String("paimodelgallery.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5": dara.String("paimodelgallery.ap-southeast-5.aliyuncs.com"),
		"cn-hangzhou":    dara.String("paimodelgallery.cn-hangzhou.aliyuncs.com"),
		"us-west-1":      dara.String("paimodelgallery.us-west-1.aliyuncs.com"),
		"us-east-1":      dara.String("paimodelgallery.us-east-1.aliyuncs.com"),
		"eu-central-1":   dara.String("paimodelgallery.eu-central-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("paimodelgallery"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 获取模型部署方案匹配资源
//
// @param request - GetModelDeploymentResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelDeploymentResourcesResponse
func (client *Client) GetModelDeploymentResourcesWithOptions(ModelId *string, request *GetModelDeploymentResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelDeploymentResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizKey) {
		query["BizKey"] = request.BizKey
	}

	if !dara.IsNil(request.ModelVersion) {
		query["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.ProfileId) {
		query["ProfileId"] = request.ProfileId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelDeploymentResources"),
		Version:     dara.String("2026-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/modelgallery/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/deployment-resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelDeploymentResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型部署方案匹配资源
//
// @param request - GetModelDeploymentResourcesRequest
//
// @return GetModelDeploymentResourcesResponse
func (client *Client) GetModelDeploymentResources(ModelId *string, request *GetModelDeploymentResourcesRequest) (_result *GetModelDeploymentResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelDeploymentResourcesResponse{}
	_body, _err := client.GetModelDeploymentResourcesWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成模型部署配置
//
// @param request - GetModelDeploymentSpecRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelDeploymentSpecResponse
func (client *Client) GetModelDeploymentSpecWithOptions(ModelId *string, request *GetModelDeploymentSpecRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelDeploymentSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizKey) {
		query["BizKey"] = request.BizKey
	}

	if !dara.IsNil(request.ModelVersion) {
		query["ModelVersion"] = request.ModelVersion
	}

	if !dara.IsNil(request.ProfileId) {
		query["ProfileId"] = request.ProfileId
	}

	if !dara.IsNil(request.ResourceSelections) {
		query["ResourceSelections"] = request.ResourceSelections
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelDeploymentSpec"),
		Version:     dara.String("2026-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/modelgallery/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/deployment-spec"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelDeploymentSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成模型部署配置
//
// @param request - GetModelDeploymentSpecRequest
//
// @return GetModelDeploymentSpecResponse
func (client *Client) GetModelDeploymentSpec(ModelId *string, request *GetModelDeploymentSpecRequest) (_result *GetModelDeploymentSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelDeploymentSpecResponse{}
	_body, _err := client.GetModelDeploymentSpecWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模型部署方案列表
//
// @param request - ListModelDeploymentProfilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelDeploymentProfilesResponse
func (client *Client) ListModelDeploymentProfilesWithOptions(ModelId *string, request *ListModelDeploymentProfilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelDeploymentProfilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizKey) {
		query["BizKey"] = request.BizKey
	}

	if !dara.IsNil(request.ModelVersion) {
		query["ModelVersion"] = request.ModelVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelDeploymentProfiles"),
		Version:     dara.String("2026-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/modelgallery/models/" + dara.PercentEncode(dara.StringValue(ModelId)) + "/deployment-profiles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelDeploymentProfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模型部署方案列表
//
// @param request - ListModelDeploymentProfilesRequest
//
// @return ListModelDeploymentProfilesResponse
func (client *Client) ListModelDeploymentProfiles(ModelId *string, request *ListModelDeploymentProfilesRequest) (_result *ListModelDeploymentProfilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelDeploymentProfilesResponse{}
	_body, _err := client.ListModelDeploymentProfilesWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取ModelGallery模型列表
//
// @param tmpReq - ListModelGalleryModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelGalleryModelsResponse
func (client *Client) ListModelGalleryModelsWithOptions(tmpReq *ListModelGalleryModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelGalleryModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListModelGalleryModelsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Conditions) {
		request.ConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Conditions, dara.String("Conditions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Collections) {
		query["Collections"] = request.Collections
	}

	if !dara.IsNil(request.Compressible) {
		query["Compressible"] = request.Compressible
	}

	if !dara.IsNil(request.ConditionsShrink) {
		query["Conditions"] = request.ConditionsShrink
	}

	if !dara.IsNil(request.DeepThink) {
		query["DeepThink"] = request.DeepThink
	}

	if !dara.IsNil(request.Demonstrable) {
		query["Demonstrable"] = request.Demonstrable
	}

	if !dara.IsNil(request.Deployable) {
		query["Deployable"] = request.Deployable
	}

	if !dara.IsNil(request.Distillable) {
		query["Distillable"] = request.Distillable
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Evaluable) {
		query["Evaluable"] = request.Evaluable
	}

	if !dara.IsNil(request.FunctionCall) {
		query["FunctionCall"] = request.FunctionCall
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelSeries) {
		query["ModelSeries"] = request.ModelSeries
	}

	if !dara.IsNil(request.ModelType) {
		query["ModelType"] = request.ModelType
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SupportedCompressionResource) {
		query["SupportedCompressionResource"] = request.SupportedCompressionResource
	}

	if !dara.IsNil(request.SupportedDistillationResource) {
		query["SupportedDistillationResource"] = request.SupportedDistillationResource
	}

	if !dara.IsNil(request.SupportedEvaluationResource) {
		query["SupportedEvaluationResource"] = request.SupportedEvaluationResource
	}

	if !dara.IsNil(request.SupportedInferenceResource) {
		query["SupportedInferenceResource"] = request.SupportedInferenceResource
	}

	if !dara.IsNil(request.SupportedTrainingResource) {
		query["SupportedTrainingResource"] = request.SupportedTrainingResource
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.Task) {
		query["Task"] = request.Task
	}

	if !dara.IsNil(request.Trainable) {
		query["Trainable"] = request.Trainable
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelGalleryModels"),
		Version:     dara.String("2026-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/modelgallery/models"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelGalleryModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ModelGallery模型列表
//
// @param request - ListModelGalleryModelsRequest
//
// @return ListModelGalleryModelsResponse
func (client *Client) ListModelGalleryModels(request *ListModelGalleryModelsRequest) (_result *ListModelGalleryModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelGalleryModelsResponse{}
	_body, _err := client.ListModelGalleryModelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
