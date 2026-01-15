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
	client.Endpoint, _err = client.GetEndpoint(dara.String("safconsole"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Create basic model information
//
// @param request - CreateCustomerModuleBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerModuleBasicInfoResponse
func (client *Client) CreateCustomerModuleBasicInfoWithOptions(request *CreateCustomerModuleBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomerModuleBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleName) {
		query["CustomerModuleName"] = request.CustomerModuleName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ModuleType) {
		query["ModuleType"] = request.ModuleType
	}

	if !dara.IsNil(request.StorePath) {
		query["StorePath"] = request.StorePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomerModuleBasicInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomerModuleBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create basic model information
//
// @param request - CreateCustomerModuleBasicInfoRequest
//
// @return CreateCustomerModuleBasicInfoResponse
func (client *Client) CreateCustomerModuleBasicInfo(request *CreateCustomerModuleBasicInfoRequest) (_result *CreateCustomerModuleBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomerModuleBasicInfoResponse{}
	_body, _err := client.CreateCustomerModuleBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Save Model Feature Configuration
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each option including a label (label) and value (value).
//
// @param request - CreateCustomerModuleMetaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerModuleMetaInfoResponse
func (client *Client) CreateCustomerModuleMetaInfoWithOptions(request *CreateCustomerModuleMetaInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomerModuleMetaInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.FeatureString) {
		query["FeatureString"] = request.FeatureString
	}

	if !dara.IsNil(request.FeatureTemplate) {
		query["FeatureTemplate"] = request.FeatureTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomerModuleMetaInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomerModuleMetaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Save Model Feature Configuration
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each option including a label (label) and value (value).
//
// @param request - CreateCustomerModuleMetaInfoRequest
//
// @return CreateCustomerModuleMetaInfoResponse
func (client *Client) CreateCustomerModuleMetaInfo(request *CreateCustomerModuleMetaInfoRequest) (_result *CreateCustomerModuleMetaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomerModuleMetaInfoResponse{}
	_body, _err := client.CreateCustomerModuleMetaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Model Output Information
//
// @param request - CreateCustomerModuleOutputInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerModuleOutputInfoResponse
func (client *Client) CreateCustomerModuleOutputInfoWithOptions(request *CreateCustomerModuleOutputInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomerModuleOutputInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.FinalScoreFormat) {
		query["FinalScoreFormat"] = request.FinalScoreFormat
	}

	if !dara.IsNil(request.ProcessExpression) {
		query["ProcessExpression"] = request.ProcessExpression
	}

	if !dara.IsNil(request.TestFileUrl) {
		query["TestFileUrl"] = request.TestFileUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomerModuleOutputInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomerModuleOutputInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Model Output Information
//
// @param request - CreateCustomerModuleOutputInfoRequest
//
// @return CreateCustomerModuleOutputInfoResponse
func (client *Client) CreateCustomerModuleOutputInfo(request *CreateCustomerModuleOutputInfoRequest) (_result *CreateCustomerModuleOutputInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomerModuleOutputInfoResponse{}
	_body, _err := client.CreateCustomerModuleOutputInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified customer model based on the provided customer model ID.
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: Deletion is irreversible, please use with caution.
//
// @param request - DeleteModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
func (client *Client) DeleteModelWithOptions(request *DeleteModelRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified customer model based on the provided customer model ID.
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: Deletion is irreversible, please use with caution.
//
// @param request - DeleteModelRequest
//
// @return DeleteModelResponse
func (client *Client) DeleteModel(request *DeleteModelRequest) (_result *DeleteModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query basic model information based on the customer model ID
//
// Description:
//
// ## Request Description
//
// By providing the `customerModuleId` parameter, you can query the current status of a specified customer model. The status values may include but are not limited to "EDIT", "ONLINE", etc.
//
// @param request - DescribeCustomerModuleBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomerModuleBasicInfoResponse
func (client *Client) DescribeCustomerModuleBasicInfoWithOptions(request *DescribeCustomerModuleBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomerModuleBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomerModuleBasicInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomerModuleBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query basic model information based on the customer model ID
//
// Description:
//
// ## Request Description
//
// By providing the `customerModuleId` parameter, you can query the current status of a specified customer model. The status values may include but are not limited to "EDIT", "ONLINE", etc.
//
// @param request - DescribeCustomerModuleBasicInfoRequest
//
// @return DescribeCustomerModuleBasicInfoResponse
func (client *Client) DescribeCustomerModuleBasicInfo(request *DescribeCustomerModuleBasicInfoRequest) (_result *DescribeCustomerModuleBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomerModuleBasicInfoResponse{}
	_body, _err := client.DescribeCustomerModuleBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据客户模型ID查询模型特征信息
//
// Description:
//
// ## 请求说明
//
// 通过提供`customerModuleId`参数，可以查询指定客户模型的当前状态。状态值可能包括但不限于"EDIT"、"ONLINE"等。
//
// @param request - DescribeCustomerModuleMetaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomerModuleMetaInfoResponse
func (client *Client) DescribeCustomerModuleMetaInfoWithOptions(request *DescribeCustomerModuleMetaInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomerModuleMetaInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomerModuleMetaInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomerModuleMetaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据客户模型ID查询模型特征信息
//
// Description:
//
// ## 请求说明
//
// 通过提供`customerModuleId`参数，可以查询指定客户模型的当前状态。状态值可能包括但不限于"EDIT"、"ONLINE"等。
//
// @param request - DescribeCustomerModuleMetaInfoRequest
//
// @return DescribeCustomerModuleMetaInfoResponse
func (client *Client) DescribeCustomerModuleMetaInfo(request *DescribeCustomerModuleMetaInfoRequest) (_result *DescribeCustomerModuleMetaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomerModuleMetaInfoResponse{}
	_body, _err := client.DescribeCustomerModuleMetaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据客户模型ID查询模型出参信息
//
// Description:
//
// ## 请求说明
//
// 通过提供`customerModuleId`参数，可以查询指定客户模型的当前状态。状态值可能包括但不限于"EDIT"、"ONLINE"等。
//
// @param request - DescribeCustomerModuleOutputInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomerModuleOutputInfoResponse
func (client *Client) DescribeCustomerModuleOutputInfoWithOptions(request *DescribeCustomerModuleOutputInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomerModuleOutputInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomerModuleOutputInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomerModuleOutputInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据客户模型ID查询模型出参信息
//
// Description:
//
// ## 请求说明
//
// 通过提供`customerModuleId`参数，可以查询指定客户模型的当前状态。状态值可能包括但不限于"EDIT"、"ONLINE"等。
//
// @param request - DescribeCustomerModuleOutputInfoRequest
//
// @return DescribeCustomerModuleOutputInfoResponse
func (client *Client) DescribeCustomerModuleOutputInfo(request *DescribeCustomerModuleOutputInfoRequest) (_result *DescribeCustomerModuleOutputInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomerModuleOutputInfoResponse{}
	_body, _err := client.DescribeCustomerModuleOutputInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据特征模板返回特征模板具体特征选项
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - DescribeFeatureOptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFeatureOptionResponse
func (client *Client) DescribeFeatureOptionWithOptions(request *DescribeFeatureOptionRequest, runtime *dara.RuntimeOptions) (_result *DescribeFeatureOptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureTemplate) {
		query["FeatureTemplate"] = request.FeatureTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFeatureOption"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFeatureOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据特征模板返回特征模板具体特征选项
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - DescribeFeatureOptionRequest
//
// @return DescribeFeatureOptionResponse
func (client *Client) DescribeFeatureOption(request *DescribeFeatureOptionRequest) (_result *DescribeFeatureOptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFeatureOptionResponse{}
	_body, _err := client.DescribeFeatureOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取可用的特征模板列表，用于模型配置。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - DescribeFeatureTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFeatureTemplateListResponse
func (client *Client) DescribeFeatureTemplateListWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeFeatureTemplateListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFeatureTemplateList"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFeatureTemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取可用的特征模板列表，用于模型配置。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @return DescribeFeatureTemplateListResponse
func (client *Client) DescribeFeatureTemplateList() (_result *DescribeFeatureTemplateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFeatureTemplateListResponse{}
	_body, _err := client.DescribeFeatureTemplateListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 渲染模型的特征配置
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - DescribeModelFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelFeatureResponse
func (client *Client) DescribeModelFeatureWithOptions(request *DescribeModelFeatureRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelFeatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.FeatureTemplate) {
		query["FeatureTemplate"] = request.FeatureTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelFeature"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 渲染模型的特征配置
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - DescribeModelFeatureRequest
//
// @return DescribeModelFeatureResponse
func (client *Client) DescribeModelFeature(request *DescribeModelFeatureRequest) (_result *DescribeModelFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModelFeatureResponse{}
	_body, _err := client.DescribeModelFeatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取上传oss鉴权数据
//
// @param request - DescribeModelOssTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelOssTokenResponse
func (client *Client) DescribeModelOssTokenWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeModelOssTokenResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelOssToken"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelOssTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取上传oss鉴权数据
//
// @return DescribeModelOssTokenResponse
func (client *Client) DescribeModelOssToken() (_result *DescribeModelOssTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModelOssTokenResponse{}
	_body, _err := client.DescribeModelOssTokenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据客户模型ID查询指定模型对应服务是否存在
//
// Description:
//
// ## 请求说明
//
// 该接口用于通过提供的`customerModuleId`来检查特定的模型服务是否已经存在。如果存在，则返回`true`；反之则返回`false`。
//
// ### 注意事项
//
// - `customerModuleId`是必须提供的参数，且为字符串类型。
//
// - 此API主要用于前端页面展示或逻辑判断时使用，以确认用户所选模型是否有对应的服务被创建。
//
// @param request - DescribeModuleServiceExistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModuleServiceExistResponse
func (client *Client) DescribeModuleServiceExistWithOptions(request *DescribeModuleServiceExistRequest, runtime *dara.RuntimeOptions) (_result *DescribeModuleServiceExistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModuleServiceExist"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModuleServiceExistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据客户模型ID查询指定模型对应服务是否存在
//
// Description:
//
// ## 请求说明
//
// 该接口用于通过提供的`customerModuleId`来检查特定的模型服务是否已经存在。如果存在，则返回`true`；反之则返回`false`。
//
// ### 注意事项
//
// - `customerModuleId`是必须提供的参数，且为字符串类型。
//
// - 此API主要用于前端页面展示或逻辑判断时使用，以确认用户所选模型是否有对应的服务被创建。
//
// @param request - DescribeModuleServiceExistRequest
//
// @return DescribeModuleServiceExistResponse
func (client *Client) DescribeModuleServiceExist(request *DescribeModuleServiceExistRequest) (_result *DescribeModuleServiceExistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModuleServiceExistResponse{}
	_body, _err := client.DescribeModuleServiceExistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据客户模型ID查询模型当前状态
//
// Description:
//
// ## 请求说明
//
// 通过提供`customerModuleId`参数，可以查询指定客户模型的当前状态。状态值可能包括但不限于"EDIT"、"ONLINE"等。
//
// @param request - DescribeModuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModuleStatusResponse
func (client *Client) DescribeModuleStatusWithOptions(request *DescribeModuleStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeModuleStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModuleStatus"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据客户模型ID查询模型当前状态
//
// Description:
//
// ## 请求说明
//
// 通过提供`customerModuleId`参数，可以查询指定客户模型的当前状态。状态值可能包括但不限于"EDIT"、"ONLINE"等。
//
// @param request - DescribeModuleStatusRequest
//
// @return DescribeModuleStatusResponse
func (client *Client) DescribeModuleStatus(request *DescribeModuleStatusRequest) (_result *DescribeModuleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModuleStatusResponse{}
	_body, _err := client.DescribeModuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模型托管订单
//
// @param request - DescribeSafRmmpOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSafRmmpOrderResponse
func (client *Client) DescribeSafRmmpOrderWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeSafRmmpOrderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSafRmmpOrder"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSafRmmpOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型托管订单
//
// @return DescribeSafRmmpOrderResponse
func (client *Client) DescribeSafRmmpOrder() (_result *DescribeSafRmmpOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSafRmmpOrderResponse{}
	_body, _err := client.DescribeSafRmmpOrderWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询场景和服务
//
// @param request - DescribeServiceAndSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceAndSceneResponse
func (client *Client) DescribeServiceAndSceneWithOptions(request *DescribeServiceAndSceneRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceAndSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceAndScene"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceAndSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询场景和服务
//
// @param request - DescribeServiceAndSceneRequest
//
// @return DescribeServiceAndSceneResponse
func (client *Client) DescribeServiceAndScene(request *DescribeServiceAndSceneRequest) (_result *DescribeServiceAndSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServiceAndSceneResponse{}
	_body, _err := client.DescribeServiceAndSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定用户下的所有模型信息，支持分页查询
//
// Description:
//
// ## 请求说明
//
// 该 API 用于查询特定用户下的所有模型信息，并支持通过分页参数进行分页查询。可以通过 `name` 参数进行模糊搜索。
//
// - `regId`: 地域标识，必填。
//
// - `pageSize`: 每页显示的条目数，必填。
//
// - `currentPage`: 当前页码，从1开始计数，必填。
//
// - `userId`: 用户ID，必填。
//
// @param request - DescribeUserModelListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserModelListResponse
func (client *Client) DescribeUserModelListWithOptions(request *DescribeUserModelListRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserModelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserModelList"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserModelListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定用户下的所有模型信息，支持分页查询
//
// Description:
//
// ## 请求说明
//
// 该 API 用于查询特定用户下的所有模型信息，并支持通过分页参数进行分页查询。可以通过 `name` 参数进行模糊搜索。
//
// - `regId`: 地域标识，必填。
//
// - `pageSize`: 每页显示的条目数，必填。
//
// - `currentPage`: 当前页码，从1开始计数，必填。
//
// - `userId`: 用户ID，必填。
//
// @param request - DescribeUserModelListRequest
//
// @return DescribeUserModelListResponse
func (client *Client) DescribeUserModelList(request *DescribeUserModelListRequest) (_result *DescribeUserModelListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserModelListResponse{}
	_body, _err := client.DescribeUserModelListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 迭代模型
//
// Description:
//
// ## 请求说明
//
// 本API用于从系统中删除指定的客户模型。调用时必须提供`customerModuleId`参数，该参数标识了要删除的具体模型。
//
// - **注意**：删除操作不可逆，请谨慎使用。
//
// @param request - DuplicateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DuplicateModelResponse
func (client *Client) DuplicateModelWithOptions(request *DuplicateModelRequest, runtime *dara.RuntimeOptions) (_result *DuplicateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DuplicateModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DuplicateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 迭代模型
//
// Description:
//
// ## 请求说明
//
// 本API用于从系统中删除指定的客户模型。调用时必须提供`customerModuleId`参数，该参数标识了要删除的具体模型。
//
// - **注意**：删除操作不可逆，请谨慎使用。
//
// @param request - DuplicateModelRequest
//
// @return DuplicateModelResponse
func (client *Client) DuplicateModel(request *DuplicateModelRequest) (_result *DuplicateModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DuplicateModelResponse{}
	_body, _err := client.DuplicateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑模型
//
// @param request - EditModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditModelResponse
func (client *Client) EditModelWithOptions(request *EditModelRequest, runtime *dara.RuntimeOptions) (_result *EditModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑模型
//
// @param request - EditModelRequest
//
// @return EditModelResponse
func (client *Client) EditModel(request *EditModelRequest) (_result *EditModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditModelResponse{}
	_body, _err := client.EditModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 迭代模型
//
// Description:
//
// ## 请求说明
//
// 本API用于从系统中删除指定的客户模型。调用时必须提供`customerModuleId`参数，该参数标识了要删除的具体模型。
//
// - **注意**：删除操作不可逆，请谨慎使用。
//
// @param request - IterateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IterateModelResponse
func (client *Client) IterateModelWithOptions(request *IterateModelRequest, runtime *dara.RuntimeOptions) (_result *IterateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IterateModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IterateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 迭代模型
//
// Description:
//
// ## 请求说明
//
// 本API用于从系统中删除指定的客户模型。调用时必须提供`customerModuleId`参数，该参数标识了要删除的具体模型。
//
// - **注意**：删除操作不可逆，请谨慎使用。
//
// @param request - IterateModelRequest
//
// @return IterateModelResponse
func (client *Client) IterateModel(request *IterateModelRequest) (_result *IterateModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IterateModelResponse{}
	_body, _err := client.IterateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下线模型
//
// @param request - OfflineModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineModelResponse
func (client *Client) OfflineModelWithOptions(request *OfflineModelRequest, runtime *dara.RuntimeOptions) (_result *OfflineModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下线模型
//
// @param request - OfflineModelRequest
//
// @return OfflineModelResponse
func (client *Client) OfflineModel(request *OfflineModelRequest) (_result *OfflineModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineModelResponse{}
	_body, _err := client.OfflineModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上线模型
//
// Description:
//
// ## 请求说明
//
// 本API用于从系统中删除指定的客户模型。调用时必须提供`customerModuleId`参数，该参数标识了要删除的具体模型。
//
// - **注意**：删除操作不可逆，请谨慎使用。
//
// @param request - OnlineModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineModelResponse
func (client *Client) OnlineModelWithOptions(request *OnlineModelRequest, runtime *dara.RuntimeOptions) (_result *OnlineModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线模型
//
// Description:
//
// ## 请求说明
//
// 本API用于从系统中删除指定的客户模型。调用时必须提供`customerModuleId`参数，该参数标识了要删除的具体模型。
//
// - **注意**：删除操作不可逆，请谨慎使用。
//
// @param request - OnlineModelRequest
//
// @return OnlineModelResponse
func (client *Client) OnlineModel(request *OnlineModelRequest) (_result *OnlineModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OnlineModelResponse{}
	_body, _err := client.OnlineModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解析表达式参数
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - ParseExpressionParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseExpressionParametersResponse
func (client *Client) ParseExpressionParametersWithOptions(request *ParseExpressionParametersRequest, runtime *dara.RuntimeOptions) (_result *ParseExpressionParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ParseExpressionParameters"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ParseExpressionParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解析表达式参数
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - ParseExpressionParametersRequest
//
// @return ParseExpressionParametersResponse
func (client *Client) ParseExpressionParameters(request *ParseExpressionParametersRequest) (_result *ParseExpressionParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ParseExpressionParametersResponse{}
	_body, _err := client.ParseExpressionParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预发布模型
//
// @param request - PrepublishModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrepublishModelResponse
func (client *Client) PrepublishModelWithOptions(request *PrepublishModelRequest, runtime *dara.RuntimeOptions) (_result *PrepublishModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrepublishModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PrepublishModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预发布模型
//
// @param request - PrepublishModelRequest
//
// @return PrepublishModelResponse
func (client *Client) PrepublishModel(request *PrepublishModelRequest) (_result *PrepublishModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PrepublishModelResponse{}
	_body, _err := client.PrepublishModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Roll back the specified customer model based on the provided customer model ID.
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: The deletion operation is irreversible, please use with caution.
//
// @param request - RollbackModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackModelResponse
func (client *Client) RollbackModelWithOptions(request *RollbackModelRequest, runtime *dara.RuntimeOptions) (_result *RollbackModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Roll back the specified customer model based on the provided customer model ID.
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: The deletion operation is irreversible, please use with caution.
//
// @param request - RollbackModelRequest
//
// @return RollbackModelResponse
func (client *Client) RollbackModel(request *RollbackModelRequest) (_result *RollbackModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackModelResponse{}
	_body, _err := client.RollbackModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预发布测试模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - TestModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestModelResponse
func (client *Client) TestModelWithOptions(request *TestModelRequest, runtime *dara.RuntimeOptions) (_result *TestModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预发布测试模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - TestModelRequest
//
// @return TestModelResponse
func (client *Client) TestModel(request *TestModelRequest) (_result *TestModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestModelResponse{}
	_body, _err := client.TestModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预发布测试模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - TestPreModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestPreModelResponse
func (client *Client) TestPreModelWithOptions(request *TestPreModelRequest, runtime *dara.RuntimeOptions) (_result *TestPreModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestPreModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestPreModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预发布测试模型
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - TestPreModelRequest
//
// @return TestPreModelResponse
func (client *Client) TestPreModel(request *TestPreModelRequest) (_result *TestPreModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestPreModelResponse{}
	_body, _err := client.TestPreModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 测试表达式
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - TestProcessExpressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestProcessExpressionResponse
func (client *Client) TestProcessExpressionWithOptions(request *TestProcessExpressionRequest, runtime *dara.RuntimeOptions) (_result *TestProcessExpressionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestProcessExpression"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestProcessExpressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 测试表达式
//
// Description:
//
// ## 请求说明
//
// - 该接口用于查询系统中所有可用的特征模板。
//
// - 请求方式为 GET，无需提供额外参数。
//
// - 返回结果包含多个特征模板选项，每个选项包括标签（label）和值（value）。
//
// @param request - TestProcessExpressionRequest
//
// @return TestProcessExpressionResponse
func (client *Client) TestProcessExpression(request *TestProcessExpressionRequest) (_result *TestProcessExpressionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestProcessExpressionResponse{}
	_body, _err := client.TestProcessExpressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update basic model information
//
// @param request - UpdateModuleBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModuleBasicInfoResponse
func (client *Client) UpdateModuleBasicInfoWithOptions(request *UpdateModuleBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateModuleBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.CustomerModuleName) {
		query["CustomerModuleName"] = request.CustomerModuleName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ModuleType) {
		query["ModuleType"] = request.ModuleType
	}

	if !dara.IsNil(request.StorePath) {
		query["StorePath"] = request.StorePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModuleBasicInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModuleBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update basic model information
//
// @param request - UpdateModuleBasicInfoRequest
//
// @return UpdateModuleBasicInfoResponse
func (client *Client) UpdateModuleBasicInfo(request *UpdateModuleBasicInfoRequest) (_result *UpdateModuleBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateModuleBasicInfoResponse{}
	_body, _err := client.UpdateModuleBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Validate model file
//
// @param request - ValidateModelFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateModelFileResponse
func (client *Client) ValidateModelFileWithOptions(request *ValidateModelFileRequest, runtime *dara.RuntimeOptions) (_result *ValidateModelFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateModelFile"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateModelFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Validate model file
//
// @param request - ValidateModelFileRequest
//
// @return ValidateModelFileResponse
func (client *Client) ValidateModelFile(request *ValidateModelFileRequest) (_result *ValidateModelFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateModelFileResponse{}
	_body, _err := client.ValidateModelFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Validate Test File
//
// @param request - ValidateTestFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateTestFileResponse
func (client *Client) ValidateTestFileWithOptions(request *ValidateTestFileRequest, runtime *dara.RuntimeOptions) (_result *ValidateTestFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateTestFile"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateTestFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Validate Test File
//
// @param request - ValidateTestFileRequest
//
// @return ValidateTestFileResponse
func (client *Client) ValidateTestFile(request *ValidateTestFileRequest) (_result *ValidateTestFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateTestFileResponse{}
	_body, _err := client.ValidateTestFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
