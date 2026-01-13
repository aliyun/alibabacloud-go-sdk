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
