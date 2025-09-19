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
	client.Endpoint, _err = client.GetEndpoint(dara.String("xtee"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Add Sample Data via CSV
//
// @param request - AddSampleDataByCsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSampleDataByCsvResponse
func (client *Client) AddSampleDataByCsvWithOptions(request *AddSampleDataByCsvRequest, runtime *dara.RuntimeOptions) (_result *AddSampleDataByCsvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OssFileName) {
		query["ossFileName"] = request.OssFileName
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleBatchUuid) {
		query["sampleBatchUuid"] = request.SampleBatchUuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSampleDataByCsv"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSampleDataByCsvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add Sample Data via CSV
//
// @param request - AddSampleDataByCsvRequest
//
// @return AddSampleDataByCsvResponse
func (client *Client) AddSampleDataByCsv(request *AddSampleDataByCsvRequest) (_result *AddSampleDataByCsvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSampleDataByCsvResponse{}
	_body, _err := client.AddSampleDataByCsvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add list data through a text box for samples
//
// @param request - AddSampleDataByTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSampleDataByTextResponse
func (client *Client) AddSampleDataByTextWithOptions(request *AddSampleDataByTextRequest, runtime *dara.RuntimeOptions) (_result *AddSampleDataByTextResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataValue) {
		query["dataValue"] = request.DataValue
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleBatchUuid) {
		query["sampleBatchUuid"] = request.SampleBatchUuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSampleDataByText"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSampleDataByTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add list data through a text box for samples
//
// @param request - AddSampleDataByTextRequest
//
// @return AddSampleDataByTextResponse
func (client *Client) AddSampleDataByText(request *AddSampleDataByTextRequest) (_result *AddSampleDataByTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSampleDataByTextResponse{}
	_body, _err := client.AddSampleDataByTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch Delete Sample List Data
//
// @param request - BatchDeleteSampleDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteSampleDataResponse
func (client *Client) BatchDeleteSampleDataWithOptions(request *BatchDeleteSampleDataRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteSampleDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Uuids) {
		query["uuids"] = request.Uuids
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteSampleData"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Delete Sample List Data
//
// @param request - BatchDeleteSampleDataRequest
//
// @return BatchDeleteSampleDataResponse
func (client *Client) BatchDeleteSampleData(request *BatchDeleteSampleDataRequest) (_result *BatchDeleteSampleDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteSampleDataResponse{}
	_body, _err := client.BatchDeleteSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Variable binding operation
//
// @param request - BindVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindVariableResponse
func (client *Client) BindVariableWithOptions(request *BindVariableRequest, runtime *dara.RuntimeOptions) (_result *BindVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ApiRegionId) {
		query["apiRegionId"] = request.ApiRegionId
	}

	if !dara.IsNil(request.ApiType) {
		query["apiType"] = request.ApiType
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.DefineId) {
		query["defineId"] = request.DefineId
	}

	if !dara.IsNil(request.DefineIds) {
		query["defineIds"] = request.DefineIds
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.ExceptionValue) {
		query["exceptionValue"] = request.ExceptionValue
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.OutputField) {
		query["outputField"] = request.OutputField
	}

	if !dara.IsNil(request.OutputType) {
		query["outputType"] = request.OutputType
	}

	if !dara.IsNil(request.Params) {
		query["params"] = request.Params
	}

	if !dara.IsNil(request.ParamsList) {
		query["paramsList"] = request.ParamsList
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Variable binding operation
//
// @param request - BindVariableRequest
//
// @return BindVariableResponse
func (client *Client) BindVariable(request *BindVariableRequest) (_result *BindVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindVariableResponse{}
	_body, _err := client.BindVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Policy Replication Lineage Check
//
// @param request - CheckCopyRuleVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCopyRuleVariableResponse
func (client *Client) CheckCopyRuleVariableWithOptions(request *CheckCopyRuleVariableRequest, runtime *dara.RuntimeOptions) (_result *CheckCopyRuleVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateType) {
		query["CreateType"] = request.CreateType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.SourceRuleId) {
		query["SourceRuleId"] = request.SourceRuleId
	}

	if !dara.IsNil(request.SourceRuleIds) {
		query["SourceRuleIds"] = request.SourceRuleIds
	}

	if !dara.IsNil(request.TargetEventCode) {
		query["TargetEventCode"] = request.TargetEventCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCopyRuleVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCopyRuleVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Policy Replication Lineage Check
//
// @param request - CheckCopyRuleVariableRequest
//
// @return CheckCopyRuleVariableResponse
func (client *Client) CheckCopyRuleVariable(request *CheckCopyRuleVariableRequest) (_result *CheckCopyRuleVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckCopyRuleVariableResponse{}
	_body, _err := client.CheckCopyRuleVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check if the cumulative number of variables exceeds the limit
//
// @param request - CheckCustVariableLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCustVariableLimitResponse
func (client *Client) CheckCustVariableLimitWithOptions(request *CheckCustVariableLimitRequest, runtime *dara.RuntimeOptions) (_result *CheckCustVariableLimitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCustVariableLimit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCustVariableLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check if the cumulative number of variables exceeds the limit
//
// @param request - CheckCustVariableLimitRequest
//
// @return CheckCustVariableLimitResponse
func (client *Client) CheckCustVariableLimit(request *CheckCustVariableLimitRequest) (_result *CheckCustVariableLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckCustVariableLimitResponse{}
	_body, _err := client.CheckCustVariableLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check if Creating Variables Exceeds the Limit
//
// @param request - CheckExpressionVariableLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckExpressionVariableLimitResponse
func (client *Client) CheckExpressionVariableLimitWithOptions(request *CheckExpressionVariableLimitRequest, runtime *dara.RuntimeOptions) (_result *CheckExpressionVariableLimitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckExpressionVariableLimit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckExpressionVariableLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check if Creating Variables Exceeds the Limit
//
// @param request - CheckExpressionVariableLimitRequest
//
// @return CheckExpressionVariableLimitResponse
func (client *Client) CheckExpressionVariableLimit(request *CheckExpressionVariableLimitRequest) (_result *CheckExpressionVariableLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckExpressionVariableLimitResponse{}
	_body, _err := client.CheckExpressionVariableLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check if the number of fields exceeds the limit
//
// @param request - CheckFieldLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckFieldLimitResponse
func (client *Client) CheckFieldLimitWithOptions(request *CheckFieldLimitRequest, runtime *dara.RuntimeOptions) (_result *CheckFieldLimitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckFieldLimit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckFieldLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check if the number of fields exceeds the limit
//
// @param request - CheckFieldLimitRequest
//
// @return CheckFieldLimitResponse
func (client *Client) CheckFieldLimit(request *CheckFieldLimitRequest) (_result *CheckFieldLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckFieldLimitResponse{}
	_body, _err := client.CheckFieldLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Validate Variable Reference
//
// @param request - CheckUsageVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUsageVariableResponse
func (client *Client) CheckUsageVariableWithOptions(request *CheckUsageVariableRequest, runtime *dara.RuntimeOptions) (_result *CheckUsageVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckUsageVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckUsageVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Validate Variable Reference
//
// @param request - CheckUsageVariableRequest
//
// @return CheckUsageVariableResponse
func (client *Client) CheckUsageVariable(request *CheckUsageVariableRequest) (_result *CheckUsageVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckUsageVariableResponse{}
	_body, _err := client.CheckUsageVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Policy Comparison
//
// @param request - CompareCopyRuleVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareCopyRuleVariableResponse
func (client *Client) CompareCopyRuleVariableWithOptions(request *CompareCopyRuleVariableRequest, runtime *dara.RuntimeOptions) (_result *CompareCopyRuleVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateType) {
		query["CreateType"] = request.CreateType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.SourceRuleId) {
		query["SourceRuleId"] = request.SourceRuleId
	}

	if !dara.IsNil(request.SourceRuleIds) {
		query["SourceRuleIds"] = request.SourceRuleIds
	}

	if !dara.IsNil(request.TargetEventCode) {
		query["TargetEventCode"] = request.TargetEventCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompareCopyRuleVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompareCopyRuleVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Policy Comparison
//
// @param request - CompareCopyRuleVariableRequest
//
// @return CompareCopyRuleVariableResponse
func (client *Client) CompareCopyRuleVariable(request *CompareCopyRuleVariableRequest) (_result *CompareCopyRuleVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompareCopyRuleVariableResponse{}
	_body, _err := client.CompareCopyRuleVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add Query Conditions
//
// @param request - CreateAnalysisConditionFavoriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnalysisConditionFavoriteResponse
func (client *Client) CreateAnalysisConditionFavoriteWithOptions(request *CreateAnalysisConditionFavoriteRequest, runtime *dara.RuntimeOptions) (_result *CreateAnalysisConditionFavoriteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Condition) {
		query["condition"] = request.Condition
	}

	if !dara.IsNil(request.EventBeginTime) {
		query["eventBeginTime"] = request.EventBeginTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.EventEndTime) {
		query["eventEndTime"] = request.EventEndTime
	}

	if !dara.IsNil(request.FieldName) {
		query["fieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FieldValue) {
		query["fieldValue"] = request.FieldValue
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAnalysisConditionFavorite"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAnalysisConditionFavoriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add Query Conditions
//
// @param request - CreateAnalysisConditionFavoriteRequest
//
// @return CreateAnalysisConditionFavoriteResponse
func (client *Client) CreateAnalysisConditionFavorite(request *CreateAnalysisConditionFavoriteRequest) (_result *CreateAnalysisConditionFavoriteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAnalysisConditionFavoriteResponse{}
	_body, _err := client.CreateAnalysisConditionFavoriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Export Task
//
// @param request - CreateAnalysisExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnalysisExportTaskResponse
func (client *Client) CreateAnalysisExportTaskWithOptions(request *CreateAnalysisExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateAnalysisExportTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Columns) {
		query["columns"] = request.Columns
	}

	if !dara.IsNil(request.Conditions) {
		query["conditions"] = request.Conditions
	}

	if !dara.IsNil(request.EventBeginTime) {
		query["eventBeginTime"] = request.EventBeginTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.EventEndTime) {
		query["eventEndTime"] = request.EventEndTime
	}

	if !dara.IsNil(request.FieldName) {
		query["fieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FieldValue) {
		query["fieldValue"] = request.FieldValue
	}

	if !dara.IsNil(request.FileFormat) {
		query["fileFormat"] = request.FileFormat
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Scope) {
		query["scope"] = request.Scope
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAnalysisExportTask"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAnalysisExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Export Task
//
// @param request - CreateAnalysisExportTaskRequest
//
// @return CreateAnalysisExportTaskResponse
func (client *Client) CreateAnalysisExportTask(request *CreateAnalysisExportTaskRequest) (_result *CreateAnalysisExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAnalysisExportTaskResponse{}
	_body, _err := client.CreateAnalysisExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create AppKey
//
// @param request - CreateAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppKeyResponse
func (client *Client) CreateAppKeyWithOptions(request *CreateAppKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateAppKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppKey"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create AppKey
//
// @param request - CreateAppKeyRequest
//
// @return CreateAppKeyResponse
func (client *Client) CreateAppKey(request *CreateAppKeyRequest) (_result *CreateAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppKeyResponse{}
	_body, _err := client.CreateAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Accumulative Variable
//
// @param request - CreateCustVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustVariableResponse
func (client *Client) CreateCustVariableWithOptions(request *CreateCustVariableRequest, runtime *dara.RuntimeOptions) (_result *CreateCustVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Condition) {
		query["condition"] = request.Condition
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.HistoryValueType) {
		query["historyValueType"] = request.HistoryValueType
	}

	if !dara.IsNil(request.Object) {
		query["object"] = request.Object
	}

	if !dara.IsNil(request.Outputs) {
		query["outputs"] = request.Outputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Subject) {
		query["subject"] = request.Subject
	}

	if !dara.IsNil(request.TimeType) {
		query["timeType"] = request.TimeType
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	if !dara.IsNil(request.TwCount) {
		query["twCount"] = request.TwCount
	}

	if !dara.IsNil(request.VelocityFC) {
		query["velocityFC"] = request.VelocityFC
	}

	if !dara.IsNil(request.VelocityTW) {
		query["velocityTW"] = request.VelocityTW
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Accumulative Variable
//
// @param request - CreateCustVariableRequest
//
// @return CreateCustVariableResponse
func (client *Client) CreateCustVariable(request *CreateCustVariableRequest) (_result *CreateCustVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustVariableResponse{}
	_body, _err := client.CreateCustVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add Data Source
//
// @param request - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(request *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OssKey) {
		query["ossKey"] = request.OssKey
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSource"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add Data Source
//
// @param request - CreateDataSourceRequest
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Event
//
// @param request - CreateEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventResponse
func (client *Client) CreateEventWithOptions(request *CreateEventRequest, runtime *dara.RuntimeOptions) (_result *CreateEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.EventName) {
		query["eventName"] = request.EventName
	}

	if !dara.IsNil(request.InputFieldsStr) {
		query["inputFieldsStr"] = request.InputFieldsStr
	}

	if !dara.IsNil(request.Memo) {
		query["memo"] = request.Memo
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["templateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		query["templateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["templateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Event
//
// @param request - CreateEventRequest
//
// @return CreateEventResponse
func (client *Client) CreateEvent(request *CreateEventRequest) (_result *CreateEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEventResponse{}
	_body, _err := client.CreateEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Custom Variable
//
// @param request - CreateExpressionVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressionVariableResponse
func (client *Client) CreateExpressionVariableWithOptions(request *CreateExpressionVariableRequest, runtime *dara.RuntimeOptions) (_result *CreateExpressionVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Expression) {
		query["expression"] = request.Expression
	}

	if !dara.IsNil(request.ExpressionTitle) {
		query["expressionTitle"] = request.ExpressionTitle
	}

	if !dara.IsNil(request.ExpressionVariable) {
		query["expressionVariable"] = request.ExpressionVariable
	}

	if !dara.IsNil(request.Outlier) {
		query["outlier"] = request.Outlier
	}

	if !dara.IsNil(request.Outputs) {
		query["outputs"] = request.Outputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExpressionVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExpressionVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Custom Variable
//
// @param request - CreateExpressionVariableRequest
//
// @return CreateExpressionVariableResponse
func (client *Client) CreateExpressionVariable(request *CreateExpressionVariableRequest) (_result *CreateExpressionVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExpressionVariableResponse{}
	_body, _err := client.CreateExpressionVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add New Field
//
// @param request - CreateFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFieldResponse
func (client *Client) CreateFieldWithOptions(request *CreateFieldRequest, runtime *dara.RuntimeOptions) (_result *CreateFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Classify) {
		query["classify"] = request.Classify
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EnumData) {
		query["enumData"] = request.EnumData
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateField"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add New Field
//
// @param request - CreateFieldRequest
//
// @return CreateFieldResponse
func (client *Client) CreateField(request *CreateFieldRequest) (_result *CreateFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFieldResponse{}
	_body, _err := client.CreateFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Submit Task
//
// @param request - CreateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelResponse
func (client *Client) CreateModelWithOptions(request *CreateModelRequest, runtime *dara.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucId) {
		query["BucId"] = request.BucId
	}

	if !dara.IsNil(request.Counts) {
		query["Counts"] = request.Counts
	}

	if !dara.IsNil(request.FileMD5) {
		query["FileMD5"] = request.FileMD5
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelScene) {
		query["ModelScene"] = request.ModelScene
	}

	if !dara.IsNil(request.ParameterNum) {
		query["ParameterNum"] = request.ParameterNum
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.UserLocalFileName) {
		query["UserLocalFileName"] = request.UserLocalFileName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModel"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Submit Task
//
// @param request - CreateModelRequest
//
// @return CreateModelResponse
func (client *Client) CreateModel(request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create POC
//
// @param request - CreatePocEvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePocEvResponse
func (client *Client) CreatePocEvWithOptions(request *CreatePocEvRequest, runtime *dara.RuntimeOptions) (_result *CreatePocEvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DateFormat) {
		query["DateFormat"] = request.DateFormat
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.Tab) {
		query["Tab"] = request.Tab
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePocEv"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePocEvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create POC
//
// @param request - CreatePocEvRequest
//
// @return CreatePocEvResponse
func (client *Client) CreatePocEv(request *CreatePocEvRequest) (_result *CreatePocEvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePocEvResponse{}
	_body, _err := client.CreatePocEvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add New Custom Query Variable
//
// @param request - CreateQueryVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueryVariableResponse
func (client *Client) CreateQueryVariableWithOptions(request *CreateQueryVariableRequest, runtime *dara.RuntimeOptions) (_result *CreateQueryVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataSourceCode) {
		query["dataSourceCode"] = request.DataSourceCode
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Expression) {
		query["expression"] = request.Expression
	}

	if !dara.IsNil(request.ExpressionTitle) {
		query["expressionTitle"] = request.ExpressionTitle
	}

	if !dara.IsNil(request.ExpressionVariable) {
		query["expressionVariable"] = request.ExpressionVariable
	}

	if !dara.IsNil(request.Outlier) {
		query["outlier"] = request.Outlier
	}

	if !dara.IsNil(request.Outputs) {
		query["outputs"] = request.Outputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQueryVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQueryVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add New Custom Query Variable
//
// @param request - CreateQueryVariableRequest
//
// @return CreateQueryVariableResponse
func (client *Client) CreateQueryVariable(request *CreateQueryVariableRequest) (_result *CreateQueryVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQueryVariableResponse{}
	_body, _err := client.CreateQueryVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Recommended Event Strategy
//
// @param request - CreateRecommendEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecommendEventRuleResponse
func (client *Client) CreateRecommendEventRuleWithOptions(request *CreateRecommendEventRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRecommendEventRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.EventName) {
		query["eventName"] = request.EventName
	}

	if !dara.IsNil(request.RecommendRuleIdsStr) {
		query["recommendRuleIdsStr"] = request.RecommendRuleIdsStr
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecommendEventRule"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecommendEventRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Recommended Event Strategy
//
// @param request - CreateRecommendEventRuleRequest
//
// @return CreateRecommendEventRuleResponse
func (client *Client) CreateRecommendEventRule(request *CreateRecommendEventRuleRequest) (_result *CreateRecommendEventRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecommendEventRuleResponse{}
	_body, _err := client.CreateRecommendEventRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Recommendation Task
//
// @param request - CreateRecommendTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecommendTaskResponse
func (client *Client) CreateRecommendTaskWithOptions(request *CreateRecommendTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRecommendTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleId) {
		query["sampleId"] = request.SampleId
	}

	if !dara.IsNil(request.VariablesStr) {
		query["variablesStr"] = request.VariablesStr
	}

	if !dara.IsNil(request.VelocitiesStr) {
		query["velocitiesStr"] = request.VelocitiesStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecommendTask"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecommendTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Recommendation Task
//
// @param request - CreateRecommendTaskRequest
//
// @return CreateRecommendTaskResponse
func (client *Client) CreateRecommendTask(request *CreateRecommendTaskRequest) (_result *CreateRecommendTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecommendTaskResponse{}
	_body, _err := client.CreateRecommendTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create Policy & Version
//
// @param request - CreateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.EventName) {
		query["eventName"] = request.EventName
	}

	if !dara.IsNil(request.LogicExpression) {
		query["logicExpression"] = request.LogicExpression
	}

	if !dara.IsNil(request.Memo) {
		query["memo"] = request.Memo
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleActions) {
		query["ruleActions"] = request.RuleActions
	}

	if !dara.IsNil(request.RuleBody) {
		query["ruleBody"] = request.RuleBody
	}

	if !dara.IsNil(request.RuleExpressions) {
		query["ruleExpressions"] = request.RuleExpressions
	}

	if !dara.IsNil(request.RuleName) {
		query["ruleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleStatus) {
		query["ruleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.RuleType) {
		query["ruleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRule"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create Policy & Version
//
// @param request - CreateRuleRequest
//
// @return CreateRuleResponse
func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Add Sample
//
// @param request - CreateSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSampleResponse
func (client *Client) CreateSampleWithOptions(request *CreateSampleRequest, runtime *dara.RuntimeOptions) (_result *CreateSampleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ClientFileName) {
		query["clientFileName"] = request.ClientFileName
	}

	if !dara.IsNil(request.ClientPath) {
		query["clientPath"] = request.ClientPath
	}

	if !dara.IsNil(request.FileType) {
		query["fileType"] = request.FileType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleTag) {
		query["sampleTag"] = request.SampleTag
	}

	if !dara.IsNil(request.SampleType) {
		query["sampleType"] = request.SampleType
	}

	if !dara.IsNil(request.SampleValues) {
		query["sampleValues"] = request.SampleValues
	}

	if !dara.IsNil(request.UploadType) {
		query["uploadType"] = request.UploadType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSample"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add Sample
//
// @param request - CreateSampleRequest
//
// @return CreateSampleResponse
func (client *Client) CreateSample(request *CreateSampleRequest) (_result *CreateSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSampleResponse{}
	_body, _err := client.CreateSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # User-level Single API to Create Sample Batches
//
// @param request - CreateSampleApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSampleApiResponse
func (client *Client) CreateSampleApiWithOptions(request *CreateSampleApiRequest, runtime *dara.RuntimeOptions) (_result *CreateSampleApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.DataValue) {
		query["DataValue"] = request.DataValue
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.SampleBatchType) {
		query["SampleBatchType"] = request.SampleBatchType
	}

	if !dara.IsNil(request.ServiceList) {
		query["ServiceList"] = request.ServiceList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSampleApi"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSampleApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # User-level Single API to Create Sample Batches
//
// @param request - CreateSampleApiRequest
//
// @return CreateSampleApiResponse
func (client *Client) CreateSampleApi(request *CreateSampleApiRequest) (_result *CreateSampleApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSampleApiResponse{}
	_body, _err := client.CreateSampleApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Sample Batch
//
// @param request - CreateSampleBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSampleBatchResponse
func (client *Client) CreateSampleBatchWithOptions(request *CreateSampleBatchRequest, runtime *dara.RuntimeOptions) (_result *CreateSampleBatchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BatchName) {
		query["batchName"] = request.BatchName
	}

	if !dara.IsNil(request.DataType) {
		query["dataType"] = request.DataType
	}

	if !dara.IsNil(request.DataValue) {
		query["dataValue"] = request.DataValue
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.OssFileName) {
		query["ossFileName"] = request.OssFileName
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleBatchType) {
		query["sampleBatchType"] = request.SampleBatchType
	}

	if !dara.IsNil(request.ServiceList) {
		query["serviceList"] = request.ServiceList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSampleBatch"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSampleBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Sample Batch
//
// @param request - CreateSampleBatchRequest
//
// @return CreateSampleBatchResponse
func (client *Client) CreateSampleBatch(request *CreateSampleBatchRequest) (_result *CreateSampleBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSampleBatchResponse{}
	_body, _err := client.CreateSampleBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Sample Data
//
// @param request - CreateSampleDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSampleDataResponse
func (client *Client) CreateSampleDataWithOptions(request *CreateSampleDataRequest, runtime *dara.RuntimeOptions) (_result *CreateSampleDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EncryptType) {
		query["encryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RiskValue) {
		query["riskValue"] = request.RiskValue
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	if !dara.IsNil(request.StorePath) {
		query["storePath"] = request.StorePath
	}

	if !dara.IsNil(request.StoreType) {
		query["storeType"] = request.StoreType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSampleData"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Sample Data
//
// @param request - CreateSampleDataRequest
//
// @return CreateSampleDataResponse
func (client *Client) CreateSampleData(request *CreateSampleDataRequest) (_result *CreateSampleDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSampleDataResponse{}
	_body, _err := client.CreateSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Task
//
// @param request - CreateSimulationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimulationTaskResponse
func (client *Client) CreateSimulationTaskWithOptions(request *CreateSimulationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSimulationTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataSourceConfig) {
		query["dataSourceConfig"] = request.DataSourceConfig
	}

	if !dara.IsNil(request.DataSourceType) {
		query["dataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.FiltersStr) {
		query["filtersStr"] = request.FiltersStr
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RulesStr) {
		query["rulesStr"] = request.RulesStr
	}

	if !dara.IsNil(request.RunTask) {
		query["runTask"] = request.RunTask
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskName) {
		query["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSimulationTask"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSimulationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Task
//
// @param request - CreateSimulationTaskRequest
//
// @return CreateSimulationTaskResponse
func (client *Client) CreateSimulationTask(request *CreateSimulationTaskRequest) (_result *CreateSimulationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSimulationTaskResponse{}
	_body, _err := client.CreateSimulationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Policy Replication
//
// @param request - DeepCopyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepCopyRuleResponse
func (client *Client) DeepCopyRuleWithOptions(request *DeepCopyRuleRequest, runtime *dara.RuntimeOptions) (_result *DeepCopyRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateType) {
		query["CreateType"] = request.CreateType
	}

	if !dara.IsNil(request.CustInsertInfo) {
		query["CustInsertInfo"] = request.CustInsertInfo
	}

	if !dara.IsNil(request.CustWriteInfo) {
		query["CustWriteInfo"] = request.CustWriteInfo
	}

	if !dara.IsNil(request.ExpressionVariableInfo) {
		query["ExpressionVariableInfo"] = request.ExpressionVariableInfo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.QueryExpressionVariableInfo) {
		query["QueryExpressionVariableInfo"] = request.QueryExpressionVariableInfo
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.SourceRuleId) {
		query["SourceRuleId"] = request.SourceRuleId
	}

	if !dara.IsNil(request.SourceRuleIds) {
		query["SourceRuleIds"] = request.SourceRuleIds
	}

	if !dara.IsNil(request.TargetEventCode) {
		query["TargetEventCode"] = request.TargetEventCode
	}

	if !dara.IsNil(request.TargetEventName) {
		query["TargetEventName"] = request.TargetEventName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeepCopyRule"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeepCopyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Policy Replication
//
// @param request - DeepCopyRuleRequest
//
// @return DeepCopyRuleResponse
func (client *Client) DeepCopyRule(request *DeepCopyRuleRequest) (_result *DeepCopyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeepCopyRuleResponse{}
	_body, _err := client.DeepCopyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Query Condition
//
// @param request - DeleteAnalysisConditionFavoriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAnalysisConditionFavoriteResponse
func (client *Client) DeleteAnalysisConditionFavoriteWithOptions(request *DeleteAnalysisConditionFavoriteRequest, runtime *dara.RuntimeOptions) (_result *DeleteAnalysisConditionFavoriteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAnalysisConditionFavorite"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAnalysisConditionFavoriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Query Condition
//
// @param request - DeleteAnalysisConditionFavoriteRequest
//
// @return DeleteAnalysisConditionFavoriteResponse
func (client *Client) DeleteAnalysisConditionFavorite(request *DeleteAnalysisConditionFavoriteRequest) (_result *DeleteAnalysisConditionFavoriteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAnalysisConditionFavoriteResponse{}
	_body, _err := client.DeleteAnalysisConditionFavoriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Bypass Event
//
// @param request - DeleteByPassShuntEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteByPassShuntEventResponse
func (client *Client) DeleteByPassShuntEventWithOptions(request *DeleteByPassShuntEventRequest, runtime *dara.RuntimeOptions) (_result *DeleteByPassShuntEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventId) {
		query["eventId"] = request.EventId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteByPassShuntEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteByPassShuntEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Bypass Event
//
// @param request - DeleteByPassShuntEventRequest
//
// @return DeleteByPassShuntEventResponse
func (client *Client) DeleteByPassShuntEvent(request *DeleteByPassShuntEventRequest) (_result *DeleteByPassShuntEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteByPassShuntEventResponse{}
	_body, _err := client.DeleteByPassShuntEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Accumulated Variable
//
// @param request - DeleteCustVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustVariableResponse
func (client *Client) DeleteCustVariableWithOptions(request *DeleteCustVariableRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.DataVersion) {
		query["dataVersion"] = request.DataVersion
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.VariableId) {
		query["variableId"] = request.VariableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Accumulated Variable
//
// @param request - DeleteCustVariableRequest
//
// @return DeleteCustVariableResponse
func (client *Client) DeleteCustVariable(request *DeleteCustVariableRequest) (_result *DeleteCustVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustVariableResponse{}
	_body, _err := client.DeleteCustVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Data Source
//
// @param request - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Data Source
//
// @param request - DeleteDataSourceRequest
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Event Field
//
// @param request - DeleteEventFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventFieldResponse
func (client *Client) DeleteEventFieldWithOptions(request *DeleteEventFieldRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.FieldName) {
		query["fieldName"] = request.FieldName
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventField"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Event Field
//
// @param request - DeleteEventFieldRequest
//
// @return DeleteEventFieldResponse
func (client *Client) DeleteEventField(request *DeleteEventFieldRequest) (_result *DeleteEventFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEventFieldResponse{}
	_body, _err := client.DeleteEventFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Custom Variable
//
// @param request - DeleteExpressionVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressionVariableResponse
func (client *Client) DeleteExpressionVariableWithOptions(request *DeleteExpressionVariableRequest, runtime *dara.RuntimeOptions) (_result *DeleteExpressionVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataVersion) {
		query["dataVersion"] = request.DataVersion
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExpressionVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExpressionVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Custom Variable
//
// @param request - DeleteExpressionVariableRequest
//
// @return DeleteExpressionVariableResponse
func (client *Client) DeleteExpressionVariable(request *DeleteExpressionVariableRequest) (_result *DeleteExpressionVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExpressionVariableResponse{}
	_body, _err := client.DeleteExpressionVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除字段
//
// @param request - DeleteFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFieldResponse
func (client *Client) DeleteFieldWithOptions(request *DeleteFieldRequest, runtime *dara.RuntimeOptions) (_result *DeleteFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteField"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除字段
//
// @param request - DeleteFieldRequest
//
// @return DeleteFieldResponse
func (client *Client) DeleteField(request *DeleteFieldRequest) (_result *DeleteFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFieldResponse{}
	_body, _err := client.DeleteFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Name List
//
// @param request - DeleteNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNameListResponse
func (client *Client) DeleteNameListWithOptions(request *DeleteNameListRequest, runtime *dara.RuntimeOptions) (_result *DeleteNameListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Ids) {
		query["ids"] = request.Ids
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNameList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Name List
//
// @param request - DeleteNameListRequest
//
// @return DeleteNameListResponse
func (client *Client) DeleteNameList(request *DeleteNameListRequest) (_result *DeleteNameListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNameListResponse{}
	_body, _err := client.DeleteNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete (pseudo) name list variable data
//
// @param request - DeleteNameListDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNameListDataResponse
func (client *Client) DeleteNameListDataWithOptions(request *DeleteNameListDataRequest, runtime *dara.RuntimeOptions) (_result *DeleteNameListDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.VariableId) {
		query["variableId"] = request.VariableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNameListData"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNameListDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete (pseudo) name list variable data
//
// @param request - DeleteNameListDataRequest
//
// @return DeleteNameListDataResponse
func (client *Client) DeleteNameListData(request *DeleteNameListDataRequest) (_result *DeleteNameListDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNameListDataResponse{}
	_body, _err := client.DeleteNameListDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Query Variable
//
// @param request - DeleteQueryVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueryVariableResponse
func (client *Client) DeleteQueryVariableWithOptions(request *DeleteQueryVariableRequest, runtime *dara.RuntimeOptions) (_result *DeleteQueryVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQueryVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQueryVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Query Variable
//
// @param request - DeleteQueryVariableRequest
//
// @return DeleteQueryVariableResponse
func (client *Client) DeleteQueryVariable(request *DeleteQueryVariableRequest) (_result *DeleteQueryVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQueryVariableResponse{}
	_body, _err := client.DeleteQueryVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Policy Version
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleVersionId) {
		query["ruleVersionId"] = request.RuleVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRule"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Policy Version
//
// @param request - DeleteRuleRequest
//
// @return DeleteRuleResponse
func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch Delete Samples
//
// @param request - DeleteSampleBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSampleBatchResponse
func (client *Client) DeleteSampleBatchWithOptions(request *DeleteSampleBatchRequest, runtime *dara.RuntimeOptions) (_result *DeleteSampleBatchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Ids) {
		query["ids"] = request.Ids
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Versions) {
		query["versions"] = request.Versions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSampleBatch"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSampleBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Delete Samples
//
// @param request - DeleteSampleBatchRequest
//
// @return DeleteSampleBatchResponse
func (client *Client) DeleteSampleBatch(request *DeleteSampleBatchRequest) (_result *DeleteSampleBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSampleBatchResponse{}
	_body, _err := client.DeleteSampleBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Sample Deletion
//
// @param request - DeleteSampleBatchMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSampleBatchMetaResponse
func (client *Client) DeleteSampleBatchMetaWithOptions(request *DeleteSampleBatchMetaRequest, runtime *dara.RuntimeOptions) (_result *DeleteSampleBatchMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BatchUuid) {
		query["batchUuid"] = request.BatchUuid
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSampleBatchMeta"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSampleBatchMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Sample Deletion
//
// @param request - DeleteSampleBatchMetaRequest
//
// @return DeleteSampleBatchMetaResponse
func (client *Client) DeleteSampleBatchMeta(request *DeleteSampleBatchMetaRequest) (_result *DeleteSampleBatchMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSampleBatchMetaResponse{}
	_body, _err := client.DeleteSampleBatchMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Sample Data
//
// @param request - DeleteSampleDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSampleDataResponse
func (client *Client) DeleteSampleDataWithOptions(request *DeleteSampleDataRequest, runtime *dara.RuntimeOptions) (_result *DeleteSampleDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSampleData"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSampleDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Sample Data
//
// @param request - DeleteSampleDataRequest
//
// @return DeleteSampleDataResponse
func (client *Client) DeleteSampleData(request *DeleteSampleDataRequest) (_result *DeleteSampleDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSampleDataResponse{}
	_body, _err := client.DeleteSampleDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Custom System Variable
//
// @param request - DeleteSelfBindVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSelfBindVariableResponse
func (client *Client) DeleteSelfBindVariableWithOptions(request *DeleteSelfBindVariableRequest, runtime *dara.RuntimeOptions) (_result *DeleteSelfBindVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSelfBindVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSelfBindVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Custom System Variable
//
// @param request - DeleteSelfBindVariableRequest
//
// @return DeleteSelfBindVariableResponse
func (client *Client) DeleteSelfBindVariable(request *DeleteSelfBindVariableRequest) (_result *DeleteSelfBindVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSelfBindVariableResponse{}
	_body, _err := client.DeleteSelfBindVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 高级查询获取左变量接口
//
// @param request - DescribeAdvanceSearchLeftVariableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvanceSearchLeftVariableListResponse
func (client *Client) DescribeAdvanceSearchLeftVariableListWithOptions(request *DescribeAdvanceSearchLeftVariableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvanceSearchLeftVariableListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvanceSearchLeftVariableList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvanceSearchLeftVariableListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 高级查询获取左变量接口
//
// @param request - DescribeAdvanceSearchLeftVariableListRequest
//
// @return DescribeAdvanceSearchLeftVariableListResponse
func (client *Client) DescribeAdvanceSearchLeftVariableList(request *DescribeAdvanceSearchLeftVariableListRequest) (_result *DescribeAdvanceSearchLeftVariableListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdvanceSearchLeftVariableListResponse{}
	_body, _err := client.DescribeAdvanceSearchLeftVariableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Advanced Query
//
// @param request - DescribeAdvanceSearchPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvanceSearchPageListResponse
func (client *Client) DescribeAdvanceSearchPageListWithOptions(request *DescribeAdvanceSearchPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvanceSearchPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Condition) {
		query["condition"] = request.Condition
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventBeginTime) {
		query["eventBeginTime"] = request.EventBeginTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.EventEndTime) {
		query["eventEndTime"] = request.EventEndTime
	}

	if !dara.IsNil(request.FieldName) {
		query["fieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FieldValue) {
		query["fieldValue"] = request.FieldValue
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvanceSearchPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvanceSearchPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Advanced Query
//
// @param request - DescribeAdvanceSearchPageListRequest
//
// @return DescribeAdvanceSearchPageListResponse
func (client *Client) DescribeAdvanceSearchPageList(request *DescribeAdvanceSearchPageListRequest) (_result *DescribeAdvanceSearchPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdvanceSearchPageListResponse{}
	_body, _err := client.DescribeAdvanceSearchPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据源列表
//
// @param request - DescribeAllDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllDataSourceResponse
func (client *Client) DescribeAllDataSourceWithOptions(request *DescribeAllDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAllDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAllDataSource"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据源列表
//
// @param request - DescribeAllDataSourceRequest
//
// @return DescribeAllDataSourceResponse
func (client *Client) DescribeAllDataSource(request *DescribeAllDataSourceRequest) (_result *DescribeAllDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAllDataSourceResponse{}
	_body, _err := client.DescribeAllDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Event List Query
//
// @param request - DescribeAllEventNameAndCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllEventNameAndCodeResponse
func (client *Client) DescribeAllEventNameAndCodeWithOptions(request *DescribeAllEventNameAndCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAllEventNameAndCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAllEventNameAndCode"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAllEventNameAndCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Event List Query
//
// @param request - DescribeAllEventNameAndCodeRequest
//
// @return DescribeAllEventNameAndCodeResponse
func (client *Client) DescribeAllEventNameAndCode(request *DescribeAllEventNameAndCodeRequest) (_result *DescribeAllEventNameAndCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAllEventNameAndCodeResponse{}
	_body, _err := client.DescribeAllEventNameAndCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Display all root variables when testing custom expressions
//
// @param request - DescribeAllRootVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllRootVariableResponse
func (client *Client) DescribeAllRootVariableWithOptions(request *DescribeAllRootVariableRequest, runtime *dara.RuntimeOptions) (_result *DescribeAllRootVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.DeviceVariableIds) {
		query["deviceVariableIds"] = request.DeviceVariableIds
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.ExpressionVariableIds) {
		query["expressionVariableIds"] = request.ExpressionVariableIds
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.NativeVariableIds) {
		query["nativeVariableIds"] = request.NativeVariableIds
	}

	if !dara.IsNil(request.QueryVariableIds) {
		query["queryVariableIds"] = request.QueryVariableIds
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.VelocityVariableIds) {
		query["velocityVariableIds"] = request.VelocityVariableIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAllRootVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAllRootVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Display all root variables when testing custom expressions
//
// @param request - DescribeAllRootVariableRequest
//
// @return DescribeAllRootVariableResponse
func (client *Client) DescribeAllRootVariable(request *DescribeAllRootVariableRequest) (_result *DescribeAllRootVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAllRootVariableResponse{}
	_body, _err := client.DescribeAllRootVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Display All Fields
//
// @param request - DescribeAnalysisColumnFieldListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAnalysisColumnFieldListResponse
func (client *Client) DescribeAnalysisColumnFieldListWithOptions(request *DescribeAnalysisColumnFieldListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAnalysisColumnFieldListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAnalysisColumnFieldList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAnalysisColumnFieldListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Display All Fields
//
// @param request - DescribeAnalysisColumnFieldListRequest
//
// @return DescribeAnalysisColumnFieldListResponse
func (client *Client) DescribeAnalysisColumnFieldList(request *DescribeAnalysisColumnFieldListRequest) (_result *DescribeAnalysisColumnFieldListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAnalysisColumnFieldListResponse{}
	_body, _err := client.DescribeAnalysisColumnFieldListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Custom Columns
//
// @param request - DescribeAnalysisColumnListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAnalysisColumnListResponse
func (client *Client) DescribeAnalysisColumnListWithOptions(request *DescribeAnalysisColumnListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAnalysisColumnListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAnalysisColumnList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAnalysisColumnListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Custom Columns
//
// @param request - DescribeAnalysisColumnListRequest
//
// @return DescribeAnalysisColumnListResponse
func (client *Client) DescribeAnalysisColumnList(request *DescribeAnalysisColumnListRequest) (_result *DescribeAnalysisColumnListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAnalysisColumnListResponse{}
	_body, _err := client.DescribeAnalysisColumnListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Condition List
//
// @param request - DescribeAnalysisConditionFavoriteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAnalysisConditionFavoriteListResponse
func (client *Client) DescribeAnalysisConditionFavoriteListWithOptions(request *DescribeAnalysisConditionFavoriteListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAnalysisConditionFavoriteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAnalysisConditionFavoriteList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAnalysisConditionFavoriteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Condition List
//
// @param request - DescribeAnalysisConditionFavoriteListRequest
//
// @return DescribeAnalysisConditionFavoriteListResponse
func (client *Client) DescribeAnalysisConditionFavoriteList(request *DescribeAnalysisConditionFavoriteListRequest) (_result *DescribeAnalysisConditionFavoriteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAnalysisConditionFavoriteListResponse{}
	_body, _err := client.DescribeAnalysisConditionFavoriteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Download Query Results
//
// @param request - DescribeAnalysisExportTaskDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAnalysisExportTaskDownloadUrlResponse
func (client *Client) DescribeAnalysisExportTaskDownloadUrlWithOptions(request *DescribeAnalysisExportTaskDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeAnalysisExportTaskDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAnalysisExportTaskDownloadUrl"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAnalysisExportTaskDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Download Query Results
//
// @param request - DescribeAnalysisExportTaskDownloadUrlRequest
//
// @return DescribeAnalysisExportTaskDownloadUrlResponse
func (client *Client) DescribeAnalysisExportTaskDownloadUrl(request *DescribeAnalysisExportTaskDownloadUrlRequest) (_result *DescribeAnalysisExportTaskDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAnalysisExportTaskDownloadUrlResponse{}
	_body, _err := client.DescribeAnalysisExportTaskDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get API Details
//
// @param request - DescribeApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiResponse
func (client *Client) DescribeApiWithOptions(request *DescribeApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ApiId) {
		query["apiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiRegionId) {
		query["apiRegionId"] = request.ApiRegionId
	}

	if !dara.IsNil(request.ApiType) {
		query["apiType"] = request.ApiType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApi"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get API Details
//
// @param request - DescribeApiRequest
//
// @return DescribeApiResponse
func (client *Client) DescribeApi(request *DescribeApiRequest) (_result *DescribeApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiResponse{}
	_body, _err := client.DescribeApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get API groups including those purchased by the user and custom ones
//
// @param request - DescribeApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiGroupsResponse
func (client *Client) DescribeApiGroupsWithOptions(request *DescribeApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ApiRegionId) {
		query["apiRegionId"] = request.ApiRegionId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiGroups"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get API groups including those purchased by the user and custom ones
//
// @param request - DescribeApiGroupsRequest
//
// @return DescribeApiGroupsResponse
func (client *Client) DescribeApiGroups(request *DescribeApiGroupsRequest) (_result *DescribeApiGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiGroupsResponse{}
	_body, _err := client.DescribeApiGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the limit information for creating API tasks
//
// @param request - DescribeApiLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiLimitResponse
func (client *Client) DescribeApiLimitWithOptions(request *DescribeApiLimitRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiLimitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiLimit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the limit information for creating API tasks
//
// @param request - DescribeApiLimitRequest
//
// @return DescribeApiLimitResponse
func (client *Client) DescribeApiLimit(request *DescribeApiLimitRequest) (_result *DescribeApiLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiLimitResponse{}
	_body, _err := client.DescribeApiLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get API Service Name
//
// @param request - DescribeApiNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiNameListResponse
func (client *Client) DescribeApiNameListWithOptions(request *DescribeApiNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiNameListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiNameList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get API Service Name
//
// @param request - DescribeApiNameListRequest
//
// @return DescribeApiNameListResponse
func (client *Client) DescribeApiNameList(request *DescribeApiNameListRequest) (_result *DescribeApiNameListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiNameListResponse{}
	_body, _err := client.DescribeApiNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Variable Details
//
// @param request - DescribeApiVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiVariableResponse
func (client *Client) DescribeApiVariableWithOptions(request *DescribeApiVariableRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Variable Details
//
// @param request - DescribeApiVariableRequest
//
// @return DescribeApiVariableResponse
func (client *Client) DescribeApiVariable(request *DescribeApiVariableRequest) (_result *DescribeApiVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiVariableResponse{}
	_body, _err := client.DescribeApiVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get API list including purchased and customized APIs
//
// @param request - DescribeApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisResponse
func (client *Client) DescribeApisWithOptions(request *DescribeApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ApiGroupId) {
		query["apiGroupId"] = request.ApiGroupId
	}

	if !dara.IsNil(request.ApiRegionId) {
		query["apiRegionId"] = request.ApiRegionId
	}

	if !dara.IsNil(request.ApiType) {
		query["apiType"] = request.ApiType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApis"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get API list including purchased and customized APIs
//
// @param request - DescribeApisRequest
//
// @return DescribeApisResponse
func (client *Client) DescribeApis(request *DescribeApisRequest) (_result *DescribeApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisResponse{}
	_body, _err := client.DescribeApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query appKey List
//
// @param request - DescribeAppKeyPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppKeyPageResponse
func (client *Client) DescribeAppKeyPageWithOptions(request *DescribeAppKeyPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppKeyPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppKeyPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppKeyPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query appKey List
//
// @param request - DescribeAppKeyPageRequest
//
// @return DescribeAppKeyPageResponse
func (client *Client) DescribeAppKeyPage(request *DescribeAppKeyPageRequest) (_result *DescribeAppKeyPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppKeyPageResponse{}
	_body, _err := client.DescribeAppKeyPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Approval Switch
//
// @param request - DescribeAuditConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditConfigResponse
func (client *Client) DescribeAuditConfigWithOptions(request *DescribeAuditConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.AuditRelationType) {
		query["auditRelationType"] = request.AuditRelationType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditConfig"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Approval Switch
//
// @param request - DescribeAuditConfigRequest
//
// @return DescribeAuditConfigResponse
func (client *Client) DescribeAuditConfig(request *DescribeAuditConfigRequest) (_result *DescribeAuditConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuditConfigResponse{}
	_body, _err := client.DescribeAuditConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Approval Details
//
// @param request - DescribeAuditDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditDetailsResponse
func (client *Client) DescribeAuditDetailsWithOptions(request *DescribeAuditDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditDetails"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Approval Details
//
// @param request - DescribeAuditDetailsRequest
//
// @return DescribeAuditDetailsResponse
func (client *Client) DescribeAuditDetails(request *DescribeAuditDetailsRequest) (_result *DescribeAuditDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuditDetailsResponse{}
	_body, _err := client.DescribeAuditDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Display and Query of Audit List
//
// @param request - DescribeAuditPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditPageListResponse
func (client *Client) DescribeAuditPageListWithOptions(request *DescribeAuditPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.AuditStatus) {
		query["auditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleName) {
		query["ruleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Display and Query of Audit List
//
// @param request - DescribeAuditPageListRequest
//
// @return DescribeAuditPageListResponse
func (client *Client) DescribeAuditPageList(request *DescribeAuditPageListRequest) (_result *DescribeAuditPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuditPageListResponse{}
	_body, _err := client.DescribeAuditPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of event names for the current user
//
// @param request - DescribeAuthEventNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthEventNameListResponse
func (client *Client) DescribeAuthEventNameListWithOptions(request *DescribeAuthEventNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthEventNameListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthEventNameList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthEventNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of event names for the current user
//
// @param request - DescribeAuthEventNameListRequest
//
// @return DescribeAuthEventNameListResponse
func (client *Client) DescribeAuthEventNameList(request *DescribeAuthEventNameListRequest) (_result *DescribeAuthEventNameListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthEventNameListResponse{}
	_body, _err := client.DescribeAuthEventNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 策略列表
//
// @param request - DescribeAuthRulePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthRulePageListResponse
func (client *Client) DescribeAuthRulePageListWithOptions(request *DescribeAuthRulePageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthRulePageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleName) {
		query["ruleName"] = request.RuleName
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthRulePageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthRulePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 策略列表
//
// @param request - DescribeAuthRulePageListRequest
//
// @return DescribeAuthRulePageListResponse
func (client *Client) DescribeAuthRulePageList(request *DescribeAuthRulePageListRequest) (_result *DescribeAuthRulePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthRulePageListResponse{}
	_body, _err := client.DescribeAuthRulePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 场景列表
//
// @param request - DescribeAuthSceneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthSceneListResponse
func (client *Client) DescribeAuthSceneListWithOptions(request *DescribeAuthSceneListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthSceneListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthSceneList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthSceneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 场景列表
//
// @param request - DescribeAuthSceneListRequest
//
// @return DescribeAuthSceneListResponse
func (client *Client) DescribeAuthSceneList(request *DescribeAuthSceneListRequest) (_result *DescribeAuthSceneListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthSceneListResponse{}
	_body, _err := client.DescribeAuthSceneListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Scene List
//
// @param request - DescribeAuthScenePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthScenePageListResponse
func (client *Client) DescribeAuthScenePageListWithOptions(request *DescribeAuthScenePageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthScenePageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SceneName) {
		query["sceneName"] = request.SceneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthScenePageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthScenePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Scene List
//
// @param request - DescribeAuthScenePageListRequest
//
// @return DescribeAuthScenePageListResponse
func (client *Client) DescribeAuthScenePageList(request *DescribeAuthScenePageListRequest) (_result *DescribeAuthScenePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthScenePageListResponse{}
	_body, _err := client.DescribeAuthScenePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check Authorization
//
// @param request - DescribeAuthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthStatusResponse
func (client *Client) DescribeAuthStatusWithOptions(request *DescribeAuthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthStatus"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check Authorization
//
// @param request - DescribeAuthStatusRequest
//
// @return DescribeAuthStatusResponse
func (client *Client) DescribeAuthStatus(request *DescribeAuthStatusRequest) (_result *DescribeAuthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthStatusResponse{}
	_body, _err := client.DescribeAuthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Average Execution Time
//
// @param request - DescribeAvgExecuteCostReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvgExecuteCostReportResponse
func (client *Client) DescribeAvgExecuteCostReportWithOptions(request *DescribeAvgExecuteCostReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvgExecuteCostReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvgExecuteCostReport"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvgExecuteCostReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Average Execution Time
//
// @param request - DescribeAvgExecuteCostReportRequest
//
// @return DescribeAvgExecuteCostReportResponse
func (client *Client) DescribeAvgExecuteCostReport(request *DescribeAvgExecuteCostReportRequest) (_result *DescribeAvgExecuteCostReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAvgExecuteCostReportResponse{}
	_body, _err := client.DescribeAvgExecuteCostReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Basic Query
//
// @param request - DescribeBasicSearchPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBasicSearchPageListResponse
func (client *Client) DescribeBasicSearchPageListWithOptions(request *DescribeBasicSearchPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBasicSearchPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventBeginTime) {
		query["eventBeginTime"] = request.EventBeginTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.EventEndTime) {
		query["eventEndTime"] = request.EventEndTime
	}

	if !dara.IsNil(request.FieldName) {
		query["fieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FieldValue) {
		query["fieldValue"] = request.FieldValue
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBasicSearchPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBasicSearchPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Basic Query
//
// @param request - DescribeBasicSearchPageListRequest
//
// @return DescribeBasicSearchPageListResponse
func (client *Client) DescribeBasicSearchPageList(request *DescribeBasicSearchPageListRequest) (_result *DescribeBasicSearchPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBasicSearchPageListResponse{}
	_body, _err := client.DescribeBasicSearchPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 基础统计
//
// @param request - DescribeBasicStartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBasicStartResponse
func (client *Client) DescribeBasicStartWithOptions(request *DescribeBasicStartRequest, runtime *dara.RuntimeOptions) (_result *DescribeBasicStartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["appKey"] = request.AppKey
	}

	if !dara.IsNil(request.EndDs) {
		query["endDs"] = request.EndDs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Service) {
		query["service"] = request.Service
	}

	if !dara.IsNil(request.StartDs) {
		query["startDs"] = request.StartDs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBasicStart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBasicStartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 基础统计
//
// @param request - DescribeBasicStartRequest
//
// @return DescribeBasicStartResponse
func (client *Client) DescribeBasicStart(request *DescribeBasicStartRequest) (_result *DescribeBasicStartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBasicStartResponse{}
	_body, _err := client.DescribeBasicStartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # View Bypass Event
//
// @param request - DescribeByPassShuntEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeByPassShuntEventResponse
func (client *Client) DescribeByPassShuntEventWithOptions(request *DescribeByPassShuntEventRequest, runtime *dara.RuntimeOptions) (_result *DescribeByPassShuntEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventId) {
		query["eventId"] = request.EventId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeByPassShuntEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeByPassShuntEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View Bypass Event
//
// @param request - DescribeByPassShuntEventRequest
//
// @return DescribeByPassShuntEventResponse
func (client *Client) DescribeByPassShuntEvent(request *DescribeByPassShuntEventRequest) (_result *DescribeByPassShuntEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeByPassShuntEventResponse{}
	_body, _err := client.DescribeByPassShuntEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the type configuration of custom accumulated variables
//
// @param request - DescribeCustVariableConfigListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustVariableConfigListResponse
func (client *Client) DescribeCustVariableConfigListWithOptions(request *DescribeCustVariableConfigListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustVariableConfigListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BizType) {
		query["bizType"] = request.BizType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TimeType) {
		query["timeType"] = request.TimeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustVariableConfigList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustVariableConfigListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the type configuration of custom accumulated variables
//
// @param request - DescribeCustVariableConfigListRequest
//
// @return DescribeCustVariableConfigListResponse
func (client *Client) DescribeCustVariableConfigList(request *DescribeCustVariableConfigListRequest) (_result *DescribeCustVariableConfigListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustVariableConfigListResponse{}
	_body, _err := client.DescribeCustVariableConfigListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Cumulative Variable Details
//
// @param request - DescribeCustVariableDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustVariableDetailResponse
func (client *Client) DescribeCustVariableDetailWithOptions(request *DescribeCustVariableDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustVariableDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustVariableDetail"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustVariableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Cumulative Variable Details
//
// @param request - DescribeCustVariableDetailRequest
//
// @return DescribeCustVariableDetailResponse
func (client *Client) DescribeCustVariableDetail(request *DescribeCustVariableDetailRequest) (_result *DescribeCustVariableDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustVariableDetailResponse{}
	_body, _err := client.DescribeCustVariableDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Custom Accumulated Variable List
//
// Description:
//
// # List Query of Accumulated Variables
//
// @param request - DescribeCustVariablePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustVariablePageResponse
func (client *Client) DescribeCustVariablePageWithOptions(request *DescribeCustVariablePageRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustVariablePageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustVariablePage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustVariablePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Custom Accumulated Variable List
//
// Description:
//
// # List Query of Accumulated Variables
//
// @param request - DescribeCustVariablePageRequest
//
// @return DescribeCustVariablePageResponse
func (client *Client) DescribeCustVariablePage(request *DescribeCustVariablePageRequest) (_result *DescribeCustVariablePageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustVariablePageResponse{}
	_body, _err := client.DescribeCustVariablePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Data Source Data Download Link
//
// @param request - DescribeDataSourceDataDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataSourceDataDownloadUrlResponse
func (client *Client) DescribeDataSourceDataDownloadUrlWithOptions(request *DescribeDataSourceDataDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataSourceDataDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataSourceId) {
		query["dataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataSourceDataDownloadUrl"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataSourceDataDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Data Source Data Download Link
//
// @param request - DescribeDataSourceDataDownloadUrlRequest
//
// @return DescribeDataSourceDataDownloadUrlResponse
func (client *Client) DescribeDataSourceDataDownloadUrl(request *DescribeDataSourceDataDownloadUrlRequest) (_result *DescribeDataSourceDataDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataSourceDataDownloadUrlResponse{}
	_body, _err := client.DescribeDataSourceDataDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve all fields of a data source
//
// @param request - DescribeDataSourceFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataSourceFieldsResponse
func (client *Client) DescribeDataSourceFieldsWithOptions(request *DescribeDataSourceFieldsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataSourceFieldsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataSourceCode) {
		query["dataSourceCode"] = request.DataSourceCode
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataSourceFields"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataSourceFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve all fields of a data source
//
// @param request - DescribeDataSourceFieldsRequest
//
// @return DescribeDataSourceFieldsResponse
func (client *Client) DescribeDataSourceFields(request *DescribeDataSourceFieldsRequest) (_result *DescribeDataSourceFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataSourceFieldsResponse{}
	_body, _err := client.DescribeDataSourceFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Data Source List Interface
//
// @param request - DescribeDataSourcePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataSourcePageListResponse
func (client *Client) DescribeDataSourcePageListWithOptions(request *DescribeDataSourcePageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataSourcePageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataSourcePageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataSourcePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Data Source List Interface
//
// @param request - DescribeDataSourcePageListRequest
//
// @return DescribeDataSourcePageListResponse
func (client *Client) DescribeDataSourcePageList(request *DescribeDataSourcePageListRequest) (_result *DescribeDataSourcePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataSourcePageListResponse{}
	_body, _err := client.DescribeDataSourcePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Decision Result Fluctuation Detection
//
// @param request - DescribeDecisionResultFluctuationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDecisionResultFluctuationResponse
func (client *Client) DescribeDecisionResultFluctuationWithOptions(request *DescribeDecisionResultFluctuationRequest, runtime *dara.RuntimeOptions) (_result *DescribeDecisionResultFluctuationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDecisionResultFluctuation"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDecisionResultFluctuationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Decision Result Fluctuation Detection
//
// @param request - DescribeDecisionResultFluctuationRequest
//
// @return DescribeDecisionResultFluctuationResponse
func (client *Client) DescribeDecisionResultFluctuation(request *DescribeDecisionResultFluctuationRequest) (_result *DescribeDecisionResultFluctuationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDecisionResultFluctuationResponse{}
	_body, _err := client.DescribeDecisionResultFluctuationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Decision Result Fluctuation Trend
//
// @param request - DescribeDecisionResultTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDecisionResultTrendResponse
func (client *Client) DescribeDecisionResultTrendWithOptions(request *DescribeDecisionResultTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeDecisionResultTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDecisionResultTrend"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDecisionResultTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Decision Result Fluctuation Trend
//
// @param request - DescribeDecisionResultTrendRequest
//
// @return DescribeDecisionResultTrendResponse
func (client *Client) DescribeDecisionResultTrend(request *DescribeDecisionResultTrendRequest) (_result *DescribeDecisionResultTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDecisionResultTrendResponse{}
	_body, _err := client.DescribeDecisionResultTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Detailed Statistics
//
// @param request - DescribeDetailStartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDetailStartResponse
func (client *Client) DescribeDetailStartWithOptions(request *DescribeDetailStartRequest, runtime *dara.RuntimeOptions) (_result *DescribeDetailStartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["appKey"] = request.AppKey
	}

	if !dara.IsNil(request.EndDs) {
		query["endDs"] = request.EndDs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Service) {
		query["service"] = request.Service
	}

	if !dara.IsNil(request.StartDs) {
		query["startDs"] = request.StartDs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDetailStart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDetailStartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Detailed Statistics
//
// @param request - DescribeDetailStartRequest
//
// @return DescribeDetailStartResponse
func (client *Client) DescribeDetailStart(request *DescribeDetailStartRequest) (_result *DescribeDetailStartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDetailStartResponse{}
	_body, _err := client.DescribeDetailStartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Download
//
// @param request - DescribeDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadUrlResponse
func (client *Client) DescribeDownloadUrlWithOptions(request *DescribeDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadUrl"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Download
//
// @param request - DescribeDownloadUrlRequest
//
// @return DescribeDownloadUrlResponse
func (client *Client) DescribeDownloadUrl(request *DescribeDownloadUrlRequest) (_result *DescribeDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDownloadUrlResponse{}
	_body, _err := client.DescribeDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Event Details
//
// @param request - DescribeEventBaseInfoByEventCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventBaseInfoByEventCodeResponse
func (client *Client) DescribeEventBaseInfoByEventCodeWithOptions(request *DescribeEventBaseInfoByEventCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventBaseInfoByEventCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventBaseInfoByEventCode"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventBaseInfoByEventCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Event Details
//
// @param request - DescribeEventBaseInfoByEventCodeRequest
//
// @return DescribeEventBaseInfoByEventCodeResponse
func (client *Client) DescribeEventBaseInfoByEventCode(request *DescribeEventBaseInfoByEventCodeRequest) (_result *DescribeEventBaseInfoByEventCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventBaseInfoByEventCodeResponse{}
	_body, _err := client.DescribeEventBaseInfoByEventCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Total Event Count
//
// @param request - DescribeEventCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventCountResponse
func (client *Client) DescribeEventCountWithOptions(request *DescribeEventCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventCount"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Total Event Count
//
// @param request - DescribeEventCountRequest
//
// @return DescribeEventCountResponse
func (client *Client) DescribeEventCount(request *DescribeEventCountRequest) (_result *DescribeEventCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventCountResponse{}
	_body, _err := client.DescribeEventCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query event details based on the request ID
//
// @param request - DescribeEventDetailByRequestIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventDetailByRequestIdResponse
func (client *Client) DescribeEventDetailByRequestIdWithOptions(request *DescribeEventDetailByRequestIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventDetailByRequestIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.EventTime) {
		query["eventTime"] = request.EventTime
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SRequestId) {
		query["sRequestId"] = request.SRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventDetailByRequestId"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventDetailByRequestIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query event details based on the request ID
//
// @param request - DescribeEventDetailByRequestIdRequest
//
// @return DescribeEventDetailByRequestIdResponse
func (client *Client) DescribeEventDetailByRequestId(request *DescribeEventDetailByRequestIdRequest) (_result *DescribeEventDetailByRequestIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventDetailByRequestIdResponse{}
	_body, _err := client.DescribeEventDetailByRequestIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Event History Details
//
// @param request - DescribeEventLogDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventLogDetailResponse
func (client *Client) DescribeEventLogDetailWithOptions(request *DescribeEventLogDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventLogDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.ReqIdByLog) {
		query["reqIdByLog"] = request.ReqIdByLog
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventLogDetail"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventLogDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Event History Details
//
// @param request - DescribeEventLogDetailRequest
//
// @return DescribeEventLogDetailResponse
func (client *Client) DescribeEventLogDetail(request *DescribeEventLogDetailRequest) (_result *DescribeEventLogDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventLogDetailResponse{}
	_body, _err := client.DescribeEventLogDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Event History List
//
// @param request - DescribeEventLogPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventLogPageResponse
func (client *Client) DescribeEventLogPageWithOptions(request *DescribeEventLogPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventLogPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.AccountIdPRP) {
		query["accountIdPRP"] = request.AccountIdPRP
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.Condition1AL) {
		query["condition1AL"] = request.Condition1AL
	}

	if !dara.IsNil(request.Condition2AL) {
		query["condition2AL"] = request.Condition2AL
	}

	if !dara.IsNil(request.Condition3AL) {
		query["condition3AL"] = request.Condition3AL
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceTypeLRP) {
		query["deviceTypeLRP"] = request.DeviceTypeLRP
	}

	if !dara.IsNil(request.EmailPRP) {
		query["emailPRP"] = request.EmailPRP
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.FailReasonLRP) {
		query["failReasonLRP"] = request.FailReasonLRP
	}

	if !dara.IsNil(request.IpPRP) {
		query["ipPRP"] = request.IpPRP
	}

	if !dara.IsNil(request.LoginResultARP) {
		query["loginResultARP"] = request.LoginResultARP
	}

	if !dara.IsNil(request.LoginTypeLRP) {
		query["loginTypeLRP"] = request.LoginTypeLRP
	}

	if !dara.IsNil(request.MacPRP) {
		query["macPRP"] = request.MacPRP
	}

	if !dara.IsNil(request.MobilePRP) {
		query["mobilePRP"] = request.MobilePRP
	}

	if !dara.IsNil(request.NickNamePRP) {
		query["nickNamePRP"] = request.NickNamePRP
	}

	if !dara.IsNil(request.OperateSourceLRP) {
		query["operateSourceLRP"] = request.OperateSourceLRP
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReferPRP) {
		query["referPRP"] = request.ReferPRP
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RegisterIpPRP) {
		query["registerIpPRP"] = request.RegisterIpPRP
	}

	if !dara.IsNil(request.ReqIdPBS) {
		query["reqIdPBS"] = request.ReqIdPBS
	}

	if !dara.IsNil(request.ScoreEBS) {
		query["scoreEBS"] = request.ScoreEBS
	}

	if !dara.IsNil(request.ScoreSBS) {
		query["scoreSBS"] = request.ScoreSBS
	}

	if !dara.IsNil(request.ServiceABS) {
		query["serviceABS"] = request.ServiceABS
	}

	if !dara.IsNil(request.TagsLBS) {
		query["tagsLBS"] = request.TagsLBS
	}

	if !dara.IsNil(request.UmidPDI) {
		query["umidPDI"] = request.UmidPDI
	}

	if !dara.IsNil(request.UserAgentPRP) {
		query["userAgentPRP"] = request.UserAgentPRP
	}

	if !dara.IsNil(request.UserNameTypeLRP) {
		query["userNameTypeLRP"] = request.UserNameTypeLRP
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventLogPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventLogPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Event History List
//
// @param request - DescribeEventLogPageRequest
//
// @return DescribeEventLogPageResponse
func (client *Client) DescribeEventLogPage(request *DescribeEventLogPageRequest) (_result *DescribeEventLogPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventLogPageResponse{}
	_body, _err := client.DescribeEventLogPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Paged Query for Events
//
// @param request - DescribeEventPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventPageListResponse
func (client *Client) DescribeEventPageListWithOptions(request *DescribeEventPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.EventName) {
		query["eventName"] = request.EventName
	}

	if !dara.IsNil(request.EventStatus) {
		query["eventStatus"] = request.EventStatus
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Paged Query for Events
//
// @param request - DescribeEventPageListRequest
//
// @return DescribeEventPageListResponse
func (client *Client) DescribeEventPageList(request *DescribeEventPageListRequest) (_result *DescribeEventPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventPageListResponse{}
	_body, _err := client.DescribeEventPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Risk Dashboard
//
// @param request - DescribeEventResultBarChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventResultBarChartResponse
func (client *Client) DescribeEventResultBarChartWithOptions(request *DescribeEventResultBarChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventResultBarChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventResultBarChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventResultBarChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Risk Dashboard
//
// @param request - DescribeEventResultBarChartRequest
//
// @return DescribeEventResultBarChartResponse
func (client *Client) DescribeEventResultBarChart(request *DescribeEventResultBarChartRequest) (_result *DescribeEventResultBarChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventResultBarChartResponse{}
	_body, _err := client.DescribeEventResultBarChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Event Overview List
//
// @param request - DescribeEventResultListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventResultListResponse
func (client *Client) DescribeEventResultListWithOptions(request *DescribeEventResultListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventResultListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventResultList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventResultListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Event Overview List
//
// @param request - DescribeEventResultListRequest
//
// @return DescribeEventResultListResponse
func (client *Client) DescribeEventResultList(request *DescribeEventResultListRequest) (_result *DescribeEventResultListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventResultListResponse{}
	_body, _err := client.DescribeEventResultListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Policy Download List
//
// @param request - DescribeEventTaskHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventTaskHistoryResponse
func (client *Client) DescribeEventTaskHistoryWithOptions(request *DescribeEventTaskHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventTaskHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventTaskHistory"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventTaskHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Policy Download List
//
// @param request - DescribeEventTaskHistoryRequest
//
// @return DescribeEventTaskHistoryResponse
func (client *Client) DescribeEventTaskHistory(request *DescribeEventTaskHistoryRequest) (_result *DescribeEventTaskHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventTaskHistoryResponse{}
	_body, _err := client.DescribeEventTaskHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Event Invocation Count
//
// @param request - DescribeEventTotalCountReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventTotalCountReportResponse
func (client *Client) DescribeEventTotalCountReportWithOptions(request *DescribeEventTotalCountReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventTotalCountReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventTotalCountReport"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventTotalCountReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Event Invocation Count
//
// @param request - DescribeEventTotalCountReportRequest
//
// @return DescribeEventTotalCountReportResponse
func (client *Client) DescribeEventTotalCountReport(request *DescribeEventTotalCountReportRequest) (_result *DescribeEventTotalCountReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventTotalCountReportResponse{}
	_body, _err := client.DescribeEventTotalCountReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch Import Policy
//
// @param request - DescribeEventUploadPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventUploadPolicyResponse
func (client *Client) DescribeEventUploadPolicyWithOptions(request *DescribeEventUploadPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventUploadPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventUploadPolicy"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Import Policy
//
// @param request - DescribeEventUploadPolicyRequest
//
// @return DescribeEventUploadPolicyResponse
func (client *Client) DescribeEventUploadPolicy(request *DescribeEventUploadPolicyRequest) (_result *DescribeEventUploadPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventUploadPolicyResponse{}
	_body, _err := client.DescribeEventUploadPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query event variables
//
// Description:
//
// # Cumulative Variable List Query
//
// @param request - DescribeEventVariableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventVariableListResponse
func (client *Client) DescribeEventVariableListWithOptions(request *DescribeEventVariableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventVariableListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.FilterDTO) {
		query["filterDTO"] = request.FilterDTO
	}

	if !dara.IsNil(request.RefObjId) {
		query["refObjId"] = request.RefObjId
	}

	if !dara.IsNil(request.RefObjType) {
		query["refObjType"] = request.RefObjType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventVariableList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventVariableListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query event variables
//
// Description:
//
// # Cumulative Variable List Query
//
// @param request - DescribeEventVariableListRequest
//
// @return DescribeEventVariableListResponse
func (client *Client) DescribeEventVariableList(request *DescribeEventVariableListRequest) (_result *DescribeEventVariableListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventVariableListResponse{}
	_body, _err := client.DescribeEventVariableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Event Template
//
// @param request - DescribeEventVariableTemplateBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventVariableTemplateBindResponse
func (client *Client) DescribeEventVariableTemplateBindWithOptions(request *DescribeEventVariableTemplateBindRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventVariableTemplateBindResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Inputs) {
		query["inputs"] = request.Inputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["templateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventVariableTemplateBind"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventVariableTemplateBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Event Template
//
// @param request - DescribeEventVariableTemplateBindRequest
//
// @return DescribeEventVariableTemplateBindResponse
func (client *Client) DescribeEventVariableTemplateBind(request *DescribeEventVariableTemplateBindRequest) (_result *DescribeEventVariableTemplateBindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventVariableTemplateBindResponse{}
	_body, _err := client.DescribeEventVariableTemplateBindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Event Template
//
// @param request - DescribeEventVariableTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventVariableTemplateListResponse
func (client *Client) DescribeEventVariableTemplateListWithOptions(request *DescribeEventVariableTemplateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventVariableTemplateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Inputs) {
		query["inputs"] = request.Inputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["templateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventVariableTemplateList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventVariableTemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Event Template
//
// @param request - DescribeEventVariableTemplateListRequest
//
// @return DescribeEventVariableTemplateListResponse
func (client *Client) DescribeEventVariableTemplateList(request *DescribeEventVariableTemplateListRequest) (_result *DescribeEventVariableTemplateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventVariableTemplateListResponse{}
	_body, _err := client.DescribeEventVariableTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Event Variables
//
// @param request - DescribeEventsVariableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsVariableListResponse
func (client *Client) DescribeEventsVariableListWithOptions(request *DescribeEventsVariableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsVariableListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.FilterDTO) {
		query["filterDTO"] = request.FilterDTO
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventsVariableList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventsVariableListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Event Variables
//
// @param request - DescribeEventsVariableListRequest
//
// @return DescribeEventsVariableListResponse
func (client *Client) DescribeEventsVariableList(request *DescribeEventsVariableListRequest) (_result *DescribeEventsVariableListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventsVariableListResponse{}
	_body, _err := client.DescribeEventsVariableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Self-service call list.
//
// @param request - DescribeExcuteNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExcuteNumResponse
func (client *Client) DescribeExcuteNumWithOptions(request *DescribeExcuteNumRequest, runtime *dara.RuntimeOptions) (_result *DescribeExcuteNumResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.Degree) {
		query["Degree"] = request.Degree
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExcuteNum"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExcuteNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Self-service call list.
//
// @param request - DescribeExcuteNumRequest
//
// @return DescribeExcuteNumResponse
func (client *Client) DescribeExcuteNum(request *DescribeExcuteNumRequest) (_result *DescribeExcuteNumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExcuteNumResponse{}
	_body, _err := client.DescribeExcuteNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validate if the field name is duplicated (based on user\\"s organization)
//
// @param request - DescribeExistNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExistNameResponse
func (client *Client) DescribeExistNameWithOptions(request *DescribeExistNameRequest, runtime *dara.RuntimeOptions) (_result *DescribeExistNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExistName"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExistNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validate if the field name is duplicated (based on user\\"s organization)
//
// @param request - DescribeExistNameRequest
//
// @return DescribeExistNameResponse
func (client *Client) DescribeExistName(request *DescribeExistNameRequest) (_result *DescribeExistNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExistNameResponse{}
	_body, _err := client.DescribeExistNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check if Scene Exists
//
// @param request - DescribeExistSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExistSceneResponse
func (client *Client) DescribeExistSceneWithOptions(request *DescribeExistSceneRequest, runtime *dara.RuntimeOptions) (_result *DescribeExistSceneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExistScene"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExistSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check if Scene Exists
//
// @param request - DescribeExistSceneRequest
//
// @return DescribeExistSceneResponse
func (client *Client) DescribeExistScene(request *DescribeExistSceneRequest) (_result *DescribeExistSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExistSceneResponse{}
	_body, _err := client.DescribeExistSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Custom Variable Details
//
// @param request - DescribeExpressionVariableDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressionVariableDetailResponse
func (client *Client) DescribeExpressionVariableDetailWithOptions(request *DescribeExpressionVariableDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressionVariableDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressionVariableDetail"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressionVariableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Custom Variable Details
//
// @param request - DescribeExpressionVariableDetailRequest
//
// @return DescribeExpressionVariableDetailResponse
func (client *Client) DescribeExpressionVariableDetail(request *DescribeExpressionVariableDetailRequest) (_result *DescribeExpressionVariableDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressionVariableDetailResponse{}
	_body, _err := client.DescribeExpressionVariableDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Function List
//
// @param request - DescribeExpressionVariableFunctionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressionVariableFunctionListResponse
func (client *Client) DescribeExpressionVariableFunctionListWithOptions(request *DescribeExpressionVariableFunctionListRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressionVariableFunctionListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressionVariableFunctionList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressionVariableFunctionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Function List
//
// @param request - DescribeExpressionVariableFunctionListRequest
//
// @return DescribeExpressionVariableFunctionListResponse
func (client *Client) DescribeExpressionVariableFunctionList(request *DescribeExpressionVariableFunctionListRequest) (_result *DescribeExpressionVariableFunctionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressionVariableFunctionListResponse{}
	_body, _err := client.DescribeExpressionVariableFunctionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Paged Query for Custom Variables.
//
// @param request - DescribeExpressionVariablePageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressionVariablePageResponse
func (client *Client) DescribeExpressionVariablePageWithOptions(request *DescribeExpressionVariablePageRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressionVariablePageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Outputs) {
		query["outputs"] = request.Outputs
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Value) {
		query["value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressionVariablePage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressionVariablePageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Paged Query for Custom Variables.
//
// @param request - DescribeExpressionVariablePageRequest
//
// @return DescribeExpressionVariablePageResponse
func (client *Client) DescribeExpressionVariablePage(request *DescribeExpressionVariablePageRequest) (_result *DescribeExpressionVariablePageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressionVariablePageResponse{}
	_body, _err := client.DescribeExpressionVariablePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Field Details
//
// @param request - DescribeFieldByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFieldByIdResponse
func (client *Client) DescribeFieldByIdWithOptions(request *DescribeFieldByIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeFieldByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFieldById"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFieldByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Field Details
//
// @param request - DescribeFieldByIdRequest
//
// @return DescribeFieldByIdResponse
func (client *Client) DescribeFieldById(request *DescribeFieldByIdRequest) (_result *DescribeFieldByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFieldByIdResponse{}
	_body, _err := client.DescribeFieldByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Field List
//
// @param request - DescribeFieldListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFieldListResponse
func (client *Client) DescribeFieldListWithOptions(request *DescribeFieldListRequest, runtime *dara.RuntimeOptions) (_result *DescribeFieldListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Condition) {
		query["condition"] = request.Condition
	}

	if !dara.IsNil(request.Inputs) {
		query["inputs"] = request.Inputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFieldList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFieldListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Field List
//
// @param request - DescribeFieldListRequest
//
// @return DescribeFieldListResponse
func (client *Client) DescribeFieldList(request *DescribeFieldListRequest) (_result *DescribeFieldListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFieldListResponse{}
	_body, _err := client.DescribeFieldListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query paged list of fields
//
// @param request - DescribeFieldPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFieldPageResponse
func (client *Client) DescribeFieldPageWithOptions(request *DescribeFieldPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeFieldPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Classify) {
		query["classify"] = request.Classify
	}

	if !dara.IsNil(request.Condition) {
		query["condition"] = request.Condition
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFieldPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFieldPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query paged list of fields
//
// @param request - DescribeFieldPageRequest
//
// @return DescribeFieldPageResponse
func (client *Client) DescribeFieldPage(request *DescribeFieldPageRequest) (_result *DescribeFieldPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFieldPageResponse{}
	_body, _err := client.DescribeFieldPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Community Account List
//
// @param request - DescribeGroupAccountPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupAccountPageResponse
func (client *Client) DescribeGroupAccountPageWithOptions(request *DescribeGroupAccountPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupAccountPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CommunityNo) {
		query["communityNo"] = request.CommunityNo
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.FieldKey) {
		query["fieldKey"] = request.FieldKey
	}

	if !dara.IsNil(request.FieldVal) {
		query["fieldVal"] = request.FieldVal
	}

	if !dara.IsNil(request.IsPage) {
		query["isPage"] = request.IsPage
	}

	if !dara.IsNil(request.Order) {
		query["order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupAccountPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupAccountPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Community Account List
//
// @param request - DescribeGroupAccountPageRequest
//
// @return DescribeGroupAccountPageResponse
func (client *Client) DescribeGroupAccountPage(request *DescribeGroupAccountPageRequest) (_result *DescribeGroupAccountPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupAccountPageResponse{}
	_body, _err := client.DescribeGroupAccountPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Community List Query Conditions
//
// @param request - DescribeGroupConditionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupConditionListResponse
func (client *Client) DescribeGroupConditionListWithOptions(request *DescribeGroupConditionListRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupConditionListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupConditionList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupConditionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Community List Query Conditions
//
// @param request - DescribeGroupConditionListRequest
//
// @return DescribeGroupConditionListResponse
func (client *Client) DescribeGroupConditionList(request *DescribeGroupConditionListRequest) (_result *DescribeGroupConditionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupConditionListResponse{}
	_body, _err := client.DescribeGroupConditionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Community List
//
// @param request - DescribeGroupPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupPageResponse
func (client *Client) DescribeGroupPageWithOptions(request *DescribeGroupPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.Order) {
		query["order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.TimeType) {
		query["timeType"] = request.TimeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Community List
//
// @param request - DescribeGroupPageRequest
//
// @return DescribeGroupPageResponse
func (client *Client) DescribeGroupPage(request *DescribeGroupPageRequest) (_result *DescribeGroupPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupPageResponse{}
	_body, _err := client.DescribeGroupPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Risk Communities Discovered Today
//
// @param request - DescribeGroupStatisticsByTodayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupStatisticsByTodayResponse
func (client *Client) DescribeGroupStatisticsByTodayWithOptions(request *DescribeGroupStatisticsByTodayRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupStatisticsByTodayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupStatisticsByToday"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupStatisticsByTodayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Risk Communities Discovered Today
//
// @param request - DescribeGroupStatisticsByTodayRequest
//
// @return DescribeGroupStatisticsByTodayResponse
func (client *Client) DescribeGroupStatisticsByToday(request *DescribeGroupStatisticsByTodayRequest) (_result *DescribeGroupStatisticsByTodayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupStatisticsByTodayResponse{}
	_body, _err := client.DescribeGroupStatisticsByTodayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Recent Trends in Risk Communities
//
// @param request - DescribeGroupTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupTrendResponse
func (client *Client) DescribeGroupTrendWithOptions(request *DescribeGroupTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Day) {
		query["day"] = request.Day
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupTrend"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Recent Trends in Risk Communities
//
// @param request - DescribeGroupTrendRequest
//
// @return DescribeGroupTrendResponse
func (client *Client) DescribeGroupTrend(request *DescribeGroupTrendRequest) (_result *DescribeGroupTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupTrendResponse{}
	_body, _err := client.DescribeGroupTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check if the policy name under the event name exists
//
// @param request - DescribeHasRuleNameByEventCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHasRuleNameByEventCodeResponse
func (client *Client) DescribeHasRuleNameByEventCodeWithOptions(request *DescribeHasRuleNameByEventCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeHasRuleNameByEventCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.ExcludeRuleId) {
		query["excludeRuleId"] = request.ExcludeRuleId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleName) {
		query["ruleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHasRuleNameByEventCode"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHasRuleNameByEventCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check if the policy name under the event name exists
//
// @param request - DescribeHasRuleNameByEventCodeRequest
//
// @return DescribeHasRuleNameByEventCodeResponse
func (client *Client) DescribeHasRuleNameByEventCode(request *DescribeHasRuleNameByEventCodeRequest) (_result *DescribeHasRuleNameByEventCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHasRuleNameByEventCodeResponse{}
	_body, _err := client.DescribeHasRuleNameByEventCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Risk Map Overview Chart (Pie Chart)
//
// @param request - DescribeHighRiskPieChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHighRiskPieChartResponse
func (client *Client) DescribeHighRiskPieChartWithOptions(request *DescribeHighRiskPieChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeHighRiskPieChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHighRiskPieChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHighRiskPieChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Risk Map Overview Chart (Pie Chart)
//
// @param request - DescribeHighRiskPieChartRequest
//
// @return DescribeHighRiskPieChartResponse
func (client *Client) DescribeHighRiskPieChart(request *DescribeHighRiskPieChartRequest) (_result *DescribeHighRiskPieChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHighRiskPieChartResponse{}
	_body, _err := client.DescribeHighRiskPieChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Policy Hit Fluctuation Detection
//
// @param request - DescribeHitRuleFluctuationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHitRuleFluctuationResponse
func (client *Client) DescribeHitRuleFluctuationWithOptions(request *DescribeHitRuleFluctuationRequest, runtime *dara.RuntimeOptions) (_result *DescribeHitRuleFluctuationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleStatus) {
		query["ruleStatus"] = request.RuleStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHitRuleFluctuation"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHitRuleFluctuationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Policy Hit Fluctuation Detection
//
// @param request - DescribeHitRuleFluctuationRequest
//
// @return DescribeHitRuleFluctuationResponse
func (client *Client) DescribeHitRuleFluctuation(request *DescribeHitRuleFluctuationRequest) (_result *DescribeHitRuleFluctuationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHitRuleFluctuationResponse{}
	_body, _err := client.DescribeHitRuleFluctuationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Top 20 Hits for Main Events/Bypass/Diversion Strategies
//
// @param request - DescribeHitRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHitRuleListResponse
func (client *Client) DescribeHitRuleListWithOptions(request *DescribeHitRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHitRuleListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.EventType) {
		query["eventType"] = request.EventType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHitRuleList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHitRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Top 20 Hits for Main Events/Bypass/Diversion Strategies
//
// @param request - DescribeHitRuleListRequest
//
// @return DescribeHitRuleListResponse
func (client *Client) DescribeHitRuleList(request *DescribeHitRuleListRequest) (_result *DescribeHitRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHitRuleListResponse{}
	_body, _err := client.DescribeHitRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Hit Rule Trend
//
// @param request - DescribeHitRuleTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHitRuleTrendResponse
func (client *Client) DescribeHitRuleTrendWithOptions(request *DescribeHitRuleTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeHitRuleTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleStatus) {
		query["ruleStatus"] = request.RuleStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHitRuleTrend"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHitRuleTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Hit Rule Trend
//
// @param request - DescribeHitRuleTrendRequest
//
// @return DescribeHitRuleTrendResponse
func (client *Client) DescribeHitRuleTrend(request *DescribeHitRuleTrendRequest) (_result *DescribeHitRuleTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHitRuleTrendResponse{}
	_body, _err := client.DescribeHitRuleTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Initialization Popup Information
//
// Description:
//
// # Add prompt information in BOPS, POC new page initialization popup prompts this information
//
// @param request - DescribeInitDigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInitDigResponse
func (client *Client) DescribeInitDigWithOptions(request *DescribeInitDigRequest, runtime *dara.RuntimeOptions) (_result *DescribeInitDigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInitDig"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInitDigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Initialization Popup Information
//
// Description:
//
// # Add prompt information in BOPS, POC new page initialization popup prompts this information
//
// @param request - DescribeInitDigRequest
//
// @return DescribeInitDigResponse
func (client *Client) DescribeInitDig(request *DescribeInitDigRequest) (_result *DescribeInitDigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInitDigResponse{}
	_body, _err := client.DescribeInitDigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Total Number of Events
//
// @param request - DescribeInputFeildCountByEventCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInputFeildCountByEventCodeResponse
func (client *Client) DescribeInputFeildCountByEventCodeWithOptions(request *DescribeInputFeildCountByEventCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInputFeildCountByEventCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInputFeildCountByEventCode"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInputFeildCountByEventCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Total Number of Events
//
// @param request - DescribeInputFeildCountByEventCodeRequest
//
// @return DescribeInputFeildCountByEventCodeResponse
func (client *Client) DescribeInputFeildCountByEventCode(request *DescribeInputFeildCountByEventCodeRequest) (_result *DescribeInputFeildCountByEventCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInputFeildCountByEventCodeResponse{}
	_body, _err := client.DescribeInputFeildCountByEventCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Display of Model List
//
// @param request - DescribeListModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListModelResponse
func (client *Client) DescribeListModelWithOptions(request *DescribeListModelRequest, runtime *dara.RuntimeOptions) (_result *DescribeListModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListModel"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Display of Model List
//
// @param request - DescribeListModelRequest
//
// @return DescribeListModelResponse
func (client *Client) DescribeListModel(request *DescribeListModelRequest) (_result *DescribeListModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeListModelResponse{}
	_body, _err := client.DescribeListModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Task List
//
// @param request - DescribeListPocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListPocResponse
func (client *Client) DescribeListPocWithOptions(request *DescribeListPocRequest, runtime *dara.RuntimeOptions) (_result *DescribeListPocResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListPoc"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListPocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Task List
//
// @param request - DescribeListPocRequest
//
// @return DescribeListPocResponse
func (client *Client) DescribeListPoc(request *DescribeListPocRequest) (_result *DescribeListPocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeListPocResponse{}
	_body, _err := client.DescribeListPocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Monitoring Object List
//
// @param request - DescribeLoanExecListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoanExecListResponse
func (client *Client) DescribeLoanExecListWithOptions(request *DescribeLoanExecListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoanExecListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BatchNo) {
		query["batchNo"] = request.BatchNo
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.MonitorObj) {
		query["monitorObj"] = request.MonitorObj
	}

	if !dara.IsNil(request.MonitorStatus) {
		query["monitorStatus"] = request.MonitorStatus
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoanExecList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoanExecListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Monitoring Object List
//
// @param request - DescribeLoanExecListRequest
//
// @return DescribeLoanExecListResponse
func (client *Client) DescribeLoanExecList(request *DescribeLoanExecListRequest) (_result *DescribeLoanExecListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoanExecListResponse{}
	_body, _err := client.DescribeLoanExecListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Loan Monitoring Task List
//
// @param request - DescribeLoanTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoanTaskListResponse
func (client *Client) DescribeLoanTaskListWithOptions(request *DescribeLoanTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoanTaskListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BatchNo) {
		query["batchNo"] = request.BatchNo
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.MonitorStatus) {
		query["monitorStatus"] = request.MonitorStatus
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoanTaskList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoanTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Loan Monitoring Task List
//
// @param request - DescribeLoanTaskListRequest
//
// @return DescribeLoanTaskListResponse
func (client *Client) DescribeLoanTaskList(request *DescribeLoanTaskListRequest) (_result *DescribeLoanTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoanTaskListResponse{}
	_body, _err := client.DescribeLoanTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Mark List
//
// @param request - DescribeMarkPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMarkPageResponse
func (client *Client) DescribeMarkPageWithOptions(request *DescribeMarkPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeMarkPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.IsPage) {
		query["isPage"] = request.IsPage
	}

	if !dara.IsNil(request.Order) {
		query["order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TaskLogId) {
		query["taskLogId"] = request.TaskLogId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMarkPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMarkPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Mark List
//
// @param request - DescribeMarkPageRequest
//
// @return DescribeMarkPageResponse
func (client *Client) DescribeMarkPage(request *DescribeMarkPageRequest) (_result *DescribeMarkPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMarkPageResponse{}
	_body, _err := client.DescribeMarkPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check Permission
//
// @param request - DescribeMenuPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMenuPermissionResponse
func (client *Client) DescribeMenuPermissionWithOptions(request *DescribeMenuPermissionRequest, runtime *dara.RuntimeOptions) (_result *DescribeMenuPermissionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PermissionType) {
		query["permissionType"] = request.PermissionType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMenuPermission"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMenuPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check Permission
//
// @param request - DescribeMenuPermissionRequest
//
// @return DescribeMenuPermissionResponse
func (client *Client) DescribeMenuPermission(request *DescribeMenuPermissionRequest) (_result *DescribeMenuPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMenuPermissionResponse{}
	_body, _err := client.DescribeMenuPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # View Result
//
// @param request - DescribeModelDetailsByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelDetailsByIdResponse
func (client *Client) DescribeModelDetailsByIdWithOptions(request *DescribeModelDetailsByIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelDetailsByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelDetailsById"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelDetailsByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View Result
//
// @param request - DescribeModelDetailsByIdRequest
//
// @return DescribeModelDetailsByIdResponse
func (client *Client) DescribeModelDetailsById(request *DescribeModelDetailsByIdRequest) (_result *DescribeModelDetailsByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModelDetailsByIdResponse{}
	_body, _err := client.DescribeModelDetailsByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get File Upload Credentials
//
// @param request - DescribeModelOssPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelOssPolicyResponse
func (client *Client) DescribeModelOssPolicyWithOptions(request *DescribeModelOssPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelOssPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelOssPolicy"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelOssPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get File Upload Credentials
//
// @param request - DescribeModelOssPolicyRequest
//
// @return DescribeModelOssPolicyResponse
func (client *Client) DescribeModelOssPolicy(request *DescribeModelOssPolicyRequest) (_result *DescribeModelOssPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModelOssPolicyResponse{}
	_body, _err := client.DescribeModelOssPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Task Limit
//
// @param request - DescribeMonitorTaskLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorTaskLimitResponse
func (client *Client) DescribeMonitorTaskLimitWithOptions(request *DescribeMonitorTaskLimitRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorTaskLimitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorTaskLimit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorTaskLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Task Limit
//
// @param request - DescribeMonitorTaskLimitRequest
//
// @return DescribeMonitorTaskLimitResponse
func (client *Client) DescribeMonitorTaskLimit(request *DescribeMonitorTaskLimitRequest) (_result *DescribeMonitorTaskLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorTaskLimitResponse{}
	_body, _err := client.DescribeMonitorTaskLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Name List Pagination
//
// @param request - DescribeNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNameListResponse
func (client *Client) DescribeNameListWithOptions(request *DescribeNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNameListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Value) {
		query["value"] = request.Value
	}

	if !dara.IsNil(request.VariableId) {
		query["variableId"] = request.VariableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNameList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Name List Pagination
//
// @param request - DescribeNameListRequest
//
// @return DescribeNameListResponse
func (client *Client) DescribeNameList(request *DescribeNameListRequest) (_result *DescribeNameListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNameListResponse{}
	_body, _err := client.DescribeNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Download Name List
//
// @param request - DescribeNameListDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNameListDownloadUrlResponse
func (client *Client) DescribeNameListDownloadUrlWithOptions(request *DescribeNameListDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeNameListDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.VariableId) {
		query["variableId"] = request.VariableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNameListDownloadUrl"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNameListDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Download Name List
//
// @param request - DescribeNameListDownloadUrlRequest
//
// @return DescribeNameListDownloadUrlResponse
func (client *Client) DescribeNameListDownloadUrl(request *DescribeNameListDownloadUrlRequest) (_result *DescribeNameListDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNameListDownloadUrlResponse{}
	_body, _err := client.DescribeNameListDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Name List Limit
//
// @param request - DescribeNameListLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNameListLimitResponse
func (client *Client) DescribeNameListLimitWithOptions(request *DescribeNameListLimitRequest, runtime *dara.RuntimeOptions) (_result *DescribeNameListLimitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNameListLimit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNameListLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Name List Limit
//
// @param request - DescribeNameListLimitRequest
//
// @return DescribeNameListLimitResponse
func (client *Client) DescribeNameListLimit(request *DescribeNameListLimitRequest) (_result *DescribeNameListLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNameListLimitResponse{}
	_body, _err := client.DescribeNameListLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the content of the list
//
// @param request - DescribeNameListPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNameListPageListResponse
func (client *Client) DescribeNameListPageListWithOptions(request *DescribeNameListPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNameListPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.UpdateBeginTime) {
		query["updateBeginTime"] = request.UpdateBeginTime
	}

	if !dara.IsNil(request.UpdateEndTime) {
		query["updateEndTime"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.Value) {
		query["value"] = request.Value
	}

	if !dara.IsNil(request.VariableId) {
		query["variableId"] = request.VariableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNameListPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNameListPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the content of the list
//
// @param request - DescribeNameListPageListRequest
//
// @return DescribeNameListPageListResponse
func (client *Client) DescribeNameListPageList(request *DescribeNameListPageListRequest) (_result *DescribeNameListPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNameListPageListResponse{}
	_body, _err := client.DescribeNameListPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of Name Types
//
// @param request - DescribeNameListTypeListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNameListTypeListResponse
func (client *Client) DescribeNameListTypeListWithOptions(request *DescribeNameListTypeListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNameListTypeListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNameListTypeList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNameListTypeListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Name Types
//
// @param request - DescribeNameListTypeListRequest
//
// @return DescribeNameListTypeListResponse
func (client *Client) DescribeNameListTypeList(request *DescribeNameListTypeListRequest) (_result *DescribeNameListTypeListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNameListTypeListResponse{}
	_body, _err := client.DescribeNameListTypeListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Name List
//
// @param request - DescribeNameListVariablePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNameListVariablePageListResponse
func (client *Client) DescribeNameListVariablePageListWithOptions(request *DescribeNameListVariablePageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNameListVariablePageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NameListType) {
		query["nameListType"] = request.NameListType
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Value) {
		query["value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNameListVariablePageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNameListVariablePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Name List
//
// @param request - DescribeNameListVariablePageListRequest
//
// @return DescribeNameListVariablePageListResponse
func (client *Client) DescribeNameListVariablePageList(request *DescribeNameListVariablePageListRequest) (_result *DescribeNameListVariablePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNameListVariablePageListResponse{}
	_body, _err := client.DescribeNameListVariablePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Operation Log Monitoring Statistics
//
// @param request - DescribeOperationLogMonitoringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperationLogMonitoringResponse
func (client *Client) DescribeOperationLogMonitoringWithOptions(request *DescribeOperationLogMonitoringRequest, runtime *dara.RuntimeOptions) (_result *DescribeOperationLogMonitoringResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.StartDate) {
		query["startDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserNameSearch) {
		query["userNameSearch"] = request.UserNameSearch
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOperationLogMonitoring"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOperationLogMonitoringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Operation Log Monitoring Statistics
//
// @param request - DescribeOperationLogMonitoringRequest
//
// @return DescribeOperationLogMonitoringResponse
func (client *Client) DescribeOperationLogMonitoring(request *DescribeOperationLogMonitoringRequest) (_result *DescribeOperationLogMonitoringResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOperationLogMonitoringResponse{}
	_body, _err := client.DescribeOperationLogMonitoringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query event list by event name
//
// @param request - DescribeOperationLogPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperationLogPageListResponse
func (client *Client) DescribeOperationLogPageListWithOptions(request *DescribeOperationLogPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeOperationLogPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.OperationSummary) {
		query["operationSummary"] = request.OperationSummary
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.StartDate) {
		query["startDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserNameSearch) {
		query["userNameSearch"] = request.UserNameSearch
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOperationLogPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOperationLogPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query event list by event name
//
// @param request - DescribeOperationLogPageListRequest
//
// @return DescribeOperationLogPageListResponse
func (client *Client) DescribeOperationLogPageList(request *DescribeOperationLogPageListRequest) (_result *DescribeOperationLogPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOperationLogPageListResponse{}
	_body, _err := client.DescribeOperationLogPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the operator mapping list based on customer ID
//
// @param request - DescribeOperatorListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperatorListResponse
func (client *Client) DescribeOperatorListWithOptions(request *DescribeOperatorListRequest, runtime *dara.RuntimeOptions) (_result *DescribeOperatorListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOperatorList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOperatorListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the operator mapping list based on customer ID
//
// @param request - DescribeOperatorListRequest
//
// @return DescribeOperatorListResponse
func (client *Client) DescribeOperatorList(request *DescribeOperatorListRequest) (_result *DescribeOperatorListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOperatorListResponse{}
	_body, _err := client.DescribeOperatorListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Operator Mapping List
//
// @param request - DescribeOperatorListBySceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperatorListBySceneResponse
func (client *Client) DescribeOperatorListBySceneWithOptions(request *DescribeOperatorListBySceneRequest, runtime *dara.RuntimeOptions) (_result *DescribeOperatorListBySceneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOperatorListByScene"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOperatorListBySceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Operator Mapping List
//
// @param request - DescribeOperatorListBySceneRequest
//
// @return DescribeOperatorListBySceneResponse
func (client *Client) DescribeOperatorListByScene(request *DescribeOperatorListBySceneRequest) (_result *DescribeOperatorListBySceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOperatorListBySceneResponse{}
	_body, _err := client.DescribeOperatorListBySceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Operator Mapping List
//
// @param request - DescribeOperatorListByTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOperatorListByTypeResponse
func (client *Client) DescribeOperatorListByTypeWithOptions(request *DescribeOperatorListByTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeOperatorListByTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOperatorListByType"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOperatorListByTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Operator Mapping List
//
// @param request - DescribeOperatorListByTypeRequest
//
// @return DescribeOperatorListByTypeResponse
func (client *Client) DescribeOperatorListByType(request *DescribeOperatorListByTypeRequest) (_result *DescribeOperatorListByTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOperatorListByTypeResponse{}
	_body, _err := client.DescribeOperatorListByTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check Oss Authorization
//
// @param request - DescribeOssAuthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssAuthStatusResponse
func (client *Client) DescribeOssAuthStatusWithOptions(request *DescribeOssAuthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssAuthStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssAuthStatus"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssAuthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check Oss Authorization
//
// @param request - DescribeOssAuthStatusRequest
//
// @return DescribeOssAuthStatusResponse
func (client *Client) DescribeOssAuthStatus(request *DescribeOssAuthStatusRequest) (_result *DescribeOssAuthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssAuthStatusResponse{}
	_body, _err := client.DescribeOssAuthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get OSS Policy
//
// @param request - DescribeOssPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssPolicyResponse
func (client *Client) DescribeOssPolicyWithOptions(request *DescribeOssPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssPolicy"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get OSS Policy
//
// @param request - DescribeOssPolicyRequest
//
// @return DescribeOssPolicyResponse
func (client *Client) DescribeOssPolicy(request *DescribeOssPolicyRequest) (_result *DescribeOssPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssPolicyResponse{}
	_body, _err := client.DescribeOssPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get File Upload Credentials
//
// @param request - DescribeOssTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssTokenResponse
func (client *Client) DescribeOssTokenWithOptions(request *DescribeOssTokenRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.UploadType) {
		query["uploadType"] = request.UploadType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssToken"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get File Upload Credentials
//
// @param request - DescribeOssTokenRequest
//
// @return DescribeOssTokenResponse
func (client *Client) DescribeOssToken(request *DescribeOssTokenRequest) (_result *DescribeOssTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssTokenResponse{}
	_body, _err := client.DescribeOssTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Event Property List
//
// @param request - DescribeParamByEventCodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParamByEventCodesResponse
func (client *Client) DescribeParamByEventCodesWithOptions(request *DescribeParamByEventCodesRequest, runtime *dara.RuntimeOptions) (_result *DescribeParamByEventCodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.Parma) {
		query["parma"] = request.Parma
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParamByEventCodes"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParamByEventCodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Event Property List
//
// @param request - DescribeParamByEventCodesRequest
//
// @return DescribeParamByEventCodesResponse
func (client *Client) DescribeParamByEventCodes(request *DescribeParamByEventCodesRequest) (_result *DescribeParamByEventCodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParamByEventCodesResponse{}
	_body, _err := client.DescribeParamByEventCodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get File Upload Credentials
//
// @param request - DescribePocOssTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePocOssTokenResponse
func (client *Client) DescribePocOssTokenWithOptions(request *DescribePocOssTokenRequest, runtime *dara.RuntimeOptions) (_result *DescribePocOssTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePocOssToken"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePocOssTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get File Upload Credentials
//
// @param request - DescribePocOssTokenRequest
//
// @return DescribePocOssTokenResponse
func (client *Client) DescribePocOssToken(request *DescribePocOssTokenRequest) (_result *DescribePocOssTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePocOssTokenResponse{}
	_body, _err := client.DescribePocOssTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get POC Task List
//
// @param request - DescribePocTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePocTaskListResponse
func (client *Client) DescribePocTaskListWithOptions(request *DescribePocTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribePocTaskListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePocTaskList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePocTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get POC Task List
//
// @param request - DescribePocTaskListRequest
//
// @return DescribePocTaskListResponse
func (client *Client) DescribePocTaskList(request *DescribePocTaskListRequest) (_result *DescribePocTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePocTaskListResponse{}
	_body, _err := client.DescribePocTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Determine if Stack Private Domain Mode is Enabled
//
// @param request - DescribePrivateStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrivateStackResponse
func (client *Client) DescribePrivateStackWithOptions(request *DescribePrivateStackRequest, runtime *dara.RuntimeOptions) (_result *DescribePrivateStackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrivateStack"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrivateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Determine if Stack Private Domain Mode is Enabled
//
// @param request - DescribePrivateStackRequest
//
// @return DescribePrivateStackResponse
func (client *Client) DescribePrivateStack(request *DescribePrivateStackRequest) (_result *DescribePrivateStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePrivateStackResponse{}
	_body, _err := client.DescribePrivateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Variable Detail Query
//
// @param request - DescribeQueryVariableDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQueryVariableDetailResponse
func (client *Client) DescribeQueryVariableDetailWithOptions(request *DescribeQueryVariableDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeQueryVariableDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQueryVariableDetail"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQueryVariableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Variable Detail Query
//
// @param request - DescribeQueryVariableDetailRequest
//
// @return DescribeQueryVariableDetailResponse
func (client *Client) DescribeQueryVariableDetail(request *DescribeQueryVariableDetailRequest) (_result *DescribeQueryVariableDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeQueryVariableDetailResponse{}
	_body, _err := client.DescribeQueryVariableDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询变量列表查询
//
// @param request - DescribeQueryVariablePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQueryVariablePageListResponse
func (client *Client) DescribeQueryVariablePageListWithOptions(request *DescribeQueryVariablePageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeQueryVariablePageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataSourceCode) {
		query["dataSourceCode"] = request.DataSourceCode
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQueryVariablePageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQueryVariablePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询变量列表查询
//
// @param request - DescribeQueryVariablePageListRequest
//
// @return DescribeQueryVariablePageListResponse
func (client *Client) DescribeQueryVariablePageList(request *DescribeQueryVariablePageListRequest) (_result *DescribeQueryVariablePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeQueryVariablePageListResponse{}
	_body, _err := client.DescribeQueryVariablePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Variable List under Sample & Scenario
//
// @param request - DescribeRecommendSceneVariablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecommendSceneVariablesResponse
func (client *Client) DescribeRecommendSceneVariablesWithOptions(request *DescribeRecommendSceneVariablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecommendSceneVariablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleId) {
		query["sampleId"] = request.SampleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecommendSceneVariables"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecommendSceneVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Variable List under Sample & Scenario
//
// @param request - DescribeRecommendSceneVariablesRequest
//
// @return DescribeRecommendSceneVariablesResponse
func (client *Client) DescribeRecommendSceneVariables(request *DescribeRecommendSceneVariablesRequest) (_result *DescribeRecommendSceneVariablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecommendSceneVariablesResponse{}
	_body, _err := client.DescribeRecommendSceneVariablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Variable Recommendation Details Query Interface
//
// @param request - DescribeRecommendTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecommendTaskDetailResponse
func (client *Client) DescribeRecommendTaskDetailWithOptions(request *DescribeRecommendTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecommendTaskDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecommendTaskDetail"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecommendTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Variable Recommendation Details Query Interface
//
// @param request - DescribeRecommendTaskDetailRequest
//
// @return DescribeRecommendTaskDetailResponse
func (client *Client) DescribeRecommendTaskDetail(request *DescribeRecommendTaskDetailRequest) (_result *DescribeRecommendTaskDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecommendTaskDetailResponse{}
	_body, _err := client.DescribeRecommendTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Variable Recommendation List Query Interface
//
// @param request - DescribeRecommendTaskPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecommendTaskPageListResponse
func (client *Client) DescribeRecommendTaskPageListWithOptions(request *DescribeRecommendTaskPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecommendTaskPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TaskName) {
		query["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecommendTaskPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecommendTaskPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Variable Recommendation List Query Interface
//
// @param request - DescribeRecommendTaskPageListRequest
//
// @return DescribeRecommendTaskPageListResponse
func (client *Client) DescribeRecommendTaskPageList(request *DescribeRecommendTaskPageListRequest) (_result *DescribeRecommendTaskPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecommendTaskPageListResponse{}
	_body, _err := client.DescribeRecommendTaskPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Indicators Information under Variables
//
// @param request - DescribeRecommendVariablesVelocityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecommendVariablesVelocityResponse
func (client *Client) DescribeRecommendVariablesVelocityWithOptions(request *DescribeRecommendVariablesVelocityRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecommendVariablesVelocityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	if !dara.IsNil(request.VariableIdsStr) {
		query["variableIdsStr"] = request.VariableIdsStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecommendVariablesVelocity"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecommendVariablesVelocityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Indicators Information under Variables
//
// @param request - DescribeRecommendVariablesVelocityRequest
//
// @return DescribeRecommendVariablesVelocityResponse
func (client *Client) DescribeRecommendVariablesVelocity(request *DescribeRecommendVariablesVelocityRequest) (_result *DescribeRecommendVariablesVelocityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecommendVariablesVelocityResponse{}
	_body, _err := client.DescribeRecommendVariablesVelocityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Supported Metrics List
//
// @param request - DescribeRecommendVelocitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecommendVelocitiesResponse
func (client *Client) DescribeRecommendVelocitiesWithOptions(request *DescribeRecommendVelocitiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecommendVelocitiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Code) {
		query["code"] = request.Code
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecommendVelocities"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecommendVelocitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Supported Metrics List
//
// @param request - DescribeRecommendVelocitiesRequest
//
// @return DescribeRecommendVelocitiesResponse
func (client *Client) DescribeRecommendVelocities(request *DescribeRecommendVelocitiesRequest) (_result *DescribeRecommendVelocitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecommendVelocitiesResponse{}
	_body, _err := client.DescribeRecommendVelocitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of regions supported by ApiGateway
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of regions supported by ApiGateway
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Request Hit Details
//
// @param request - DescribeRequestHitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRequestHitResponse
func (client *Client) DescribeRequestHitWithOptions(request *DescribeRequestHitRequest, runtime *dara.RuntimeOptions) (_result *DescribeRequestHitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SRequestId) {
		query["sRequestId"] = request.SRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRequestHit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRequestHitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Request Hit Details
//
// @param request - DescribeRequestHitRequest
//
// @return DescribeRequestHitResponse
func (client *Client) DescribeRequestHit(request *DescribeRequestHitRequest) (_result *DescribeRequestHitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRequestHitResponse{}
	_body, _err := client.DescribeRequestHitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Request Peak
//
// @param request - DescribeRequestPeakReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRequestPeakReportResponse
func (client *Client) DescribeRequestPeakReportWithOptions(request *DescribeRequestPeakReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeRequestPeakReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRequestPeakReport"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRequestPeakReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Request Peak
//
// @param request - DescribeRequestPeakReportRequest
//
// @return DescribeRequestPeakReportResponse
func (client *Client) DescribeRequestPeakReport(request *DescribeRequestPeakReportRequest) (_result *DescribeRequestPeakReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRequestPeakReportResponse{}
	_body, _err := client.DescribeRequestPeakReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Drill-down Analysis
//
// @param request - DescribeResultCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResultCountResponse
func (client *Client) DescribeResultCountWithOptions(request *DescribeResultCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeResultCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResultCount"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResultCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Drill-down Analysis
//
// @param request - DescribeResultCountRequest
//
// @return DescribeResultCountResponse
func (client *Client) DescribeResultCount(request *DescribeResultCountRequest) (_result *DescribeResultCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResultCountResponse{}
	_body, _err := client.DescribeResultCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Risk map overview chart (line chart)
//
// @param request - DescribeRiskLineChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskLineChartResponse
func (client *Client) DescribeRiskLineChartWithOptions(request *DescribeRiskLineChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskLineChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRiskLineChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskLineChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Risk map overview chart (line chart)
//
// @param request - DescribeRiskLineChartRequest
//
// @return DescribeRiskLineChartResponse
func (client *Client) DescribeRiskLineChart(request *DescribeRiskLineChartRequest) (_result *DescribeRiskLineChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRiskLineChartResponse{}
	_body, _err := client.DescribeRiskLineChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag Hit Rate Tag Hit Dimension
//
// @param request - DescribeRiskTagsLineChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskTagsLineChartResponse
func (client *Client) DescribeRiskTagsLineChartWithOptions(request *DescribeRiskTagsLineChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskTagsLineChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["EventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRiskTagsLineChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskTagsLineChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag Hit Rate Tag Hit Dimension
//
// @param request - DescribeRiskTagsLineChartRequest
//
// @return DescribeRiskTagsLineChartResponse
func (client *Client) DescribeRiskTagsLineChart(request *DescribeRiskTagsLineChartRequest) (_result *DescribeRiskTagsLineChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRiskTagsLineChartResponse{}
	_body, _err := client.DescribeRiskTagsLineChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Policy Overview List
//
// @param request - DescribeRuleBarChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleBarChartResponse
func (client *Client) DescribeRuleBarChartWithOptions(request *DescribeRuleBarChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleBarChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleBarChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleBarChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Policy Overview List
//
// @param request - DescribeRuleBarChartRequest
//
// @return DescribeRuleBarChartResponse
func (client *Client) DescribeRuleBarChart(request *DescribeRuleBarChartRequest) (_result *DescribeRuleBarChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleBarChartResponse{}
	_body, _err := client.DescribeRuleBarChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Policy Count by User ID
//
// @param request - DescribeRuleCountByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleCountByUserIdResponse
func (client *Client) DescribeRuleCountByUserIdWithOptions(request *DescribeRuleCountByUserIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleCountByUserIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleCountByUserId"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleCountByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Policy Count by User ID
//
// @param request - DescribeRuleCountByUserIdRequest
//
// @return DescribeRuleCountByUserIdResponse
func (client *Client) DescribeRuleCountByUserId(request *DescribeRuleCountByUserIdRequest) (_result *DescribeRuleCountByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleCountByUserIdResponse{}
	_body, _err := client.DescribeRuleCountByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query policy/version details
//
// @param request - DescribeRuleDetailByRuleIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleDetailByRuleIdResponse
func (client *Client) DescribeRuleDetailByRuleIdWithOptions(request *DescribeRuleDetailByRuleIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleDetailByRuleIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleVersionId) {
		query["ruleVersionId"] = request.RuleVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleDetailByRuleId"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleDetailByRuleIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query policy/version details
//
// @param request - DescribeRuleDetailByRuleIdRequest
//
// @return DescribeRuleDetailByRuleIdResponse
func (client *Client) DescribeRuleDetailByRuleId(request *DescribeRuleDetailByRuleIdRequest) (_result *DescribeRuleDetailByRuleIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleDetailByRuleIdResponse{}
	_body, _err := client.DescribeRuleDetailByRuleIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query rule hit details
//
// @param request - DescribeRuleHitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleHitResponse
func (client *Client) DescribeRuleHitWithOptions(request *DescribeRuleHitRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleHitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RequestTime) {
		query["requestTime"] = request.RequestTime
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleSnapshotId) {
		query["ruleSnapshotId"] = request.RuleSnapshotId
	}

	if !dara.IsNil(request.SRequestId) {
		query["sRequestId"] = request.SRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleHit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleHitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query rule hit details
//
// @param request - DescribeRuleHitRequest
//
// @return DescribeRuleHitResponse
func (client *Client) DescribeRuleHit(request *DescribeRuleHitRequest) (_result *DescribeRuleHitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleHitResponse{}
	_body, _err := client.DescribeRuleHitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query policy list
//
// @param request - DescribeRuleListByEventCodesListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleListByEventCodesListResponse
func (client *Client) DescribeRuleListByEventCodesListWithOptions(request *DescribeRuleListByEventCodesListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleListByEventCodesListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleListByEventCodesList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleListByEventCodesListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query policy list
//
// @param request - DescribeRuleListByEventCodesListRequest
//
// @return DescribeRuleListByEventCodesListResponse
func (client *Client) DescribeRuleListByEventCodesList(request *DescribeRuleListByEventCodesListRequest) (_result *DescribeRuleListByEventCodesListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleListByEventCodesListResponse{}
	_body, _err := client.DescribeRuleListByEventCodesListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of policies
//
// @param request - DescribeRulePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRulePageListResponse
func (client *Client) DescribeRulePageListWithOptions(request *DescribeRulePageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRulePageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleAuthType) {
		query["ruleAuthType"] = request.RuleAuthType
	}

	if !dara.IsNil(request.RuleName) {
		query["ruleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleStatus) {
		query["ruleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRulePageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRulePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of policies
//
// @param request - DescribeRulePageListRequest
//
// @return DescribeRulePageListResponse
func (client *Client) DescribeRulePageList(request *DescribeRulePageListRequest) (_result *DescribeRulePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRulePageListResponse{}
	_body, _err := client.DescribeRulePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query historical snapshots based on ruleId and version
//
// @param request - DescribeRuleSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleSnapshotResponse
func (client *Client) DescribeRuleSnapshotWithOptions(request *DescribeRuleSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleSnapshotResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.SnapshotVersion) {
		query["snapshotVersion"] = request.SnapshotVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleSnapshot"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query historical snapshots based on ruleId and version
//
// @param request - DescribeRuleSnapshotRequest
//
// @return DescribeRuleSnapshotResponse
func (client *Client) DescribeRuleSnapshot(request *DescribeRuleSnapshotRequest) (_result *DescribeRuleSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleSnapshotResponse{}
	_body, _err := client.DescribeRuleSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Policy Version List
//
// @param request - DescribeRuleVersionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleVersionListResponse
func (client *Client) DescribeRuleVersionListWithOptions(request *DescribeRuleVersionListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleVersionListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleVersionList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleVersionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Policy Version List
//
// @param request - DescribeRuleVersionListRequest
//
// @return DescribeRuleVersionListResponse
func (client *Client) DescribeRuleVersionList(request *DescribeRuleVersionListRequest) (_result *DescribeRuleVersionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleVersionListResponse{}
	_body, _err := client.DescribeRuleVersionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SDK Download List
//
// @param request - DescribeSDKDownloadListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSDKDownloadListResponse
func (client *Client) DescribeSDKDownloadListWithOptions(request *DescribeSDKDownloadListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSDKDownloadListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DeviceType) {
		query["deviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.ListType) {
		query["listType"] = request.ListType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSDKDownloadList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSDKDownloadListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SDK Download List
//
// @param request - DescribeSDKDownloadListRequest
//
// @return DescribeSDKDownloadListResponse
func (client *Client) DescribeSDKDownloadList(request *DescribeSDKDownloadListRequest) (_result *DescribeSDKDownloadListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSDKDownloadListResponse{}
	_body, _err := client.DescribeSDKDownloadListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query SAF permissions.
//
// @param request - DescribeSafConsoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSafConsoleResponse
func (client *Client) DescribeSafConsoleWithOptions(request *DescribeSafConsoleRequest, runtime *dara.RuntimeOptions) (_result *DescribeSafConsoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	if !dara.IsNil(request.Service) {
		query["service"] = request.Service
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSafConsole"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSafConsoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query SAF permissions.
//
// @param request - DescribeSafConsoleRequest
//
// @return DescribeSafConsoleResponse
func (client *Client) DescribeSafConsole(request *DescribeSafConsoleRequest) (_result *DescribeSafConsoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSafConsoleResponse{}
	_body, _err := client.DescribeSafConsoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query saf_de Order
//
// @param request - DescribeSafDeOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSafDeOrderResponse
func (client *Client) DescribeSafDeOrderWithOptions(request *DescribeSafDeOrderRequest, runtime *dara.RuntimeOptions) (_result *DescribeSafDeOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.StartDate) {
		query["startDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSafDeOrder"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSafDeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query saf_de Order
//
// @param request - DescribeSafDeOrderRequest
//
// @return DescribeSafDeOrderResponse
func (client *Client) DescribeSafDeOrder(request *DescribeSafDeOrderRequest) (_result *DescribeSafDeOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSafDeOrderResponse{}
	_body, _err := client.DescribeSafDeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Order Information
//
// @param request - DescribeSafOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSafOrderResponse
func (client *Client) DescribeSafOrderWithOptions(request *DescribeSafOrderRequest, runtime *dara.RuntimeOptions) (_result *DescribeSafOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.ExactProductCode) {
		query["exactProductCode"] = request.ExactProductCode
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.StartDate) {
		query["startDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSafOrder"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSafOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Order Information
//
// @param request - DescribeSafOrderRequest
//
// @return DescribeSafOrderResponse
func (client *Client) DescribeSafOrder(request *DescribeSafOrderRequest) (_result *DescribeSafOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSafOrderResponse{}
	_body, _err := client.DescribeSafOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Access Configuration
//
// @param request - DescribeSafStartConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSafStartConfigResponse
func (client *Client) DescribeSafStartConfigWithOptions(request *DescribeSafStartConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeSafStartConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSafStartConfig"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSafStartConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Access Configuration
//
// @param request - DescribeSafStartConfigRequest
//
// @return DescribeSafStartConfigResponse
func (client *Client) DescribeSafStartConfig(request *DescribeSafStartConfigRequest) (_result *DescribeSafStartConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSafStartConfigResponse{}
	_body, _err := client.DescribeSafStartConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Access Configuration
//
// @param request - DescribeSafStartStepsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSafStartStepsResponse
func (client *Client) DescribeSafStartStepsWithOptions(request *DescribeSafStartStepsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSafStartStepsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.AliyunServer) {
		query["aliyunServer"] = request.AliyunServer
	}

	if !dara.IsNil(request.DeviceTypesStr) {
		query["deviceTypesStr"] = request.DeviceTypesStr
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.ServerRegion) {
		query["serverRegion"] = request.ServerRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSafStartSteps"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSafStartStepsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Access Configuration
//
// @param request - DescribeSafStartStepsRequest
//
// @return DescribeSafStartStepsResponse
func (client *Client) DescribeSafStartSteps(request *DescribeSafStartStepsRequest) (_result *DescribeSafStartStepsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSafStartStepsResponse{}
	_body, _err := client.DescribeSafStartStepsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Risk Tag List
//
// @param request - DescribeSafTagListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSafTagListResponse
func (client *Client) DescribeSafTagListWithOptions(request *DescribeSafTagListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSafTagListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	if !dara.IsNil(request.ApiId) {
		query["apiId"] = request.ApiId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSafTagList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSafTagListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Risk Tag List
//
// @param request - DescribeSafTagListRequest
//
// @return DescribeSafTagListResponse
func (client *Client) DescribeSafTagList(request *DescribeSafTagListRequest) (_result *DescribeSafTagListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSafTagListResponse{}
	_body, _err := client.DescribeSafTagListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get File Upload Credentials
//
// @param request - DescribeSampleBatchOssPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleBatchOssPolicyResponse
func (client *Client) DescribeSampleBatchOssPolicyWithOptions(request *DescribeSampleBatchOssPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleBatchOssPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BatchName) {
		query["batchName"] = request.BatchName
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleBatchOssPolicy"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleBatchOssPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get File Upload Credentials
//
// @param request - DescribeSampleBatchOssPolicyRequest
//
// @return DescribeSampleBatchOssPolicyResponse
func (client *Client) DescribeSampleBatchOssPolicy(request *DescribeSampleBatchOssPolicyRequest) (_result *DescribeSampleBatchOssPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleBatchOssPolicyResponse{}
	_body, _err := client.DescribeSampleBatchOssPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Sample List
//
// @param request - DescribeSampleDataByBatchUUidPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleDataByBatchUUidPageResponse
func (client *Client) DescribeSampleDataByBatchUUidPageWithOptions(request *DescribeSampleDataByBatchUUidPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleDataByBatchUUidPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BatchUuid) {
		query["batchUuid"] = request.BatchUuid
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataValue) {
		query["dataValue"] = request.DataValue
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.UpdateBeginTime) {
		query["updateBeginTime"] = request.UpdateBeginTime
	}

	if !dara.IsNil(request.UpdateEndTime) {
		query["updateEndTime"] = request.UpdateEndTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleDataByBatchUUidPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleDataByBatchUUidPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Sample List
//
// @param request - DescribeSampleDataByBatchUUidPageRequest
//
// @return DescribeSampleDataByBatchUUidPageResponse
func (client *Client) DescribeSampleDataByBatchUUidPage(request *DescribeSampleDataByBatchUUidPageRequest) (_result *DescribeSampleDataByBatchUUidPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleDataByBatchUUidPageResponse{}
	_body, _err := client.DescribeSampleDataByBatchUUidPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Paged Query of Sample List
//
// @param request - DescribeSampleDataListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleDataListResponse
func (client *Client) DescribeSampleDataListWithOptions(request *DescribeSampleDataListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleDataListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeleteTag) {
		query["deleteTag"] = request.DeleteTag
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryContent) {
		query["queryContent"] = request.QueryContent
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleId) {
		query["sampleId"] = request.SampleId
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleDataList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleDataListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Paged Query of Sample List
//
// @param request - DescribeSampleDataListRequest
//
// @return DescribeSampleDataListResponse
func (client *Client) DescribeSampleDataList(request *DescribeSampleDataListRequest) (_result *DescribeSampleDataListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleDataListResponse{}
	_body, _err := client.DescribeSampleDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Sample List
//
// @param request - DescribeSampleDataPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleDataPageResponse
func (client *Client) DescribeSampleDataPageWithOptions(request *DescribeSampleDataPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleDataPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataValue) {
		query["dataValue"] = request.DataValue
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.UpdateBeginTime) {
		query["updateBeginTime"] = request.UpdateBeginTime
	}

	if !dara.IsNil(request.UpdateEndTime) {
		query["updateEndTime"] = request.UpdateEndTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleDataPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleDataPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Sample List
//
// @param request - DescribeSampleDataPageRequest
//
// @return DescribeSampleDataPageResponse
func (client *Client) DescribeSampleDataPage(request *DescribeSampleDataPageRequest) (_result *DescribeSampleDataPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleDataPageResponse{}
	_body, _err := client.DescribeSampleDataPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Sample Example Authorization
//
// @param request - DescribeSampleDemoDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleDemoDownloadUrlResponse
func (client *Client) DescribeSampleDemoDownloadUrlWithOptions(request *DescribeSampleDemoDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleDemoDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleDemoDownloadUrl"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleDemoDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Sample Example Authorization
//
// @param request - DescribeSampleDemoDownloadUrlRequest
//
// @return DescribeSampleDemoDownloadUrlResponse
func (client *Client) DescribeSampleDemoDownloadUrl(request *DescribeSampleDemoDownloadUrlRequest) (_result *DescribeSampleDemoDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleDemoDownloadUrlResponse{}
	_body, _err := client.DescribeSampleDemoDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Sample Download Authorization Information
//
// @param request - DescribeSampleDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleDownloadUrlResponse
func (client *Client) DescribeSampleDownloadUrlWithOptions(request *DescribeSampleDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleId) {
		query["sampleId"] = request.SampleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleDownloadUrl"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Sample Download Authorization Information
//
// @param request - DescribeSampleDownloadUrlRequest
//
// @return DescribeSampleDownloadUrlResponse
func (client *Client) DescribeSampleDownloadUrl(request *DescribeSampleDownloadUrlRequest) (_result *DescribeSampleDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleDownloadUrlResponse{}
	_body, _err := client.DescribeSampleDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Sample Details
//
// @param request - DescribeSampleInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleInfoResponse
func (client *Client) DescribeSampleInfoWithOptions(request *DescribeSampleInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Versions) {
		query["versions"] = request.Versions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleInfo"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Sample Details
//
// @param request - DescribeSampleInfoRequest
//
// @return DescribeSampleInfoResponse
func (client *Client) DescribeSampleInfo(request *DescribeSampleInfoRequest) (_result *DescribeSampleInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleInfoResponse{}
	_body, _err := client.DescribeSampleInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Sample List
//
// @param request - DescribeSampleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleListResponse
func (client *Client) DescribeSampleListWithOptions(request *DescribeSampleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleType) {
		query["sampleType"] = request.SampleType
	}

	if !dara.IsNil(request.SampleValue) {
		query["sampleValue"] = request.SampleValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Sample List
//
// @param request - DescribeSampleListRequest
//
// @return DescribeSampleListResponse
func (client *Client) DescribeSampleList(request *DescribeSampleListRequest) (_result *DescribeSampleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleListResponse{}
	_body, _err := client.DescribeSampleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Scene List
//
// @param request - DescribeSampleSceneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleSceneListResponse
func (client *Client) DescribeSampleSceneListWithOptions(request *DescribeSampleSceneListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleSceneListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleSceneList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleSceneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Scene List
//
// @param request - DescribeSampleSceneListRequest
//
// @return DescribeSampleSceneListResponse
func (client *Client) DescribeSampleSceneList(request *DescribeSampleSceneListRequest) (_result *DescribeSampleSceneListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleSceneListResponse{}
	_body, _err := client.DescribeSampleSceneListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Tag List
//
// @param request - DescribeSampleTagListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleTagListResponse
func (client *Client) DescribeSampleTagListWithOptions(request *DescribeSampleTagListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleTagListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleTagList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleTagListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Tag List
//
// @param request - DescribeSampleTagListRequest
//
// @return DescribeSampleTagListResponse
func (client *Client) DescribeSampleTagList(request *DescribeSampleTagListRequest) (_result *DescribeSampleTagListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleTagListResponse{}
	_body, _err := client.DescribeSampleTagListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Sample Upload Authorization Information
//
// @param request - DescribeSampleUploadPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSampleUploadPolicyResponse
func (client *Client) DescribeSampleUploadPolicyWithOptions(request *DescribeSampleUploadPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeSampleUploadPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSampleUploadPolicy"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSampleUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Sample Upload Authorization Information
//
// @param request - DescribeSampleUploadPolicyRequest
//
// @return DescribeSampleUploadPolicyResponse
func (client *Client) DescribeSampleUploadPolicy(request *DescribeSampleUploadPolicyRequest) (_result *DescribeSampleUploadPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSampleUploadPolicyResponse{}
	_body, _err := client.DescribeSampleUploadPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Sample Batch List
//
// @param request - DescribeSamplebatchPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSamplebatchPageResponse
func (client *Client) DescribeSamplebatchPageWithOptions(request *DescribeSamplebatchPageRequest, runtime *dara.RuntimeOptions) (_result *DescribeSamplebatchPageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataValue) {
		query["dataValue"] = request.DataValue
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSamplebatchPage"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSamplebatchPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Sample Batch List
//
// @param request - DescribeSamplebatchPageRequest
//
// @return DescribeSamplebatchPageResponse
func (client *Client) DescribeSamplebatchPage(request *DescribeSamplebatchPageRequest) (_result *DescribeSamplebatchPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSamplebatchPageResponse{}
	_body, _err := client.DescribeSamplebatchPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Dropdown list for scenario-based service events
//
// Description:
//
// # Dropdown list for scenario-based risk control events
//
// @param request - DescribeSceneAllEventNameCodeListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSceneAllEventNameCodeListResponse
func (client *Client) DescribeSceneAllEventNameCodeListWithOptions(request *DescribeSceneAllEventNameCodeListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSceneAllEventNameCodeListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSceneAllEventNameCodeList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSceneAllEventNameCodeListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Dropdown list for scenario-based service events
//
// Description:
//
// # Dropdown list for scenario-based risk control events
//
// @param request - DescribeSceneAllEventNameCodeListRequest
//
// @return DescribeSceneAllEventNameCodeListResponse
func (client *Client) DescribeSceneAllEventNameCodeList(request *DescribeSceneAllEventNameCodeListRequest) (_result *DescribeSceneAllEventNameCodeListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSceneAllEventNameCodeListResponse{}
	_body, _err := client.DescribeSceneAllEventNameCodeListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of Scenario-based Risk Control Events
//
// @param request - DescribeSceneEventPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSceneEventPageListResponse
func (client *Client) DescribeSceneEventPageListWithOptions(request *DescribeSceneEventPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSceneEventPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.NameOrCode) {
		query["nameOrCode"] = request.NameOrCode
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSceneEventPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSceneEventPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Scenario-based Risk Control Events
//
// @param request - DescribeSceneEventPageListRequest
//
// @return DescribeSceneEventPageListResponse
func (client *Client) DescribeSceneEventPageList(request *DescribeSceneEventPageListRequest) (_result *DescribeSceneEventPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSceneEventPageListResponse{}
	_body, _err := client.DescribeSceneEventPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # White-boxed strategy list for risk control services
//
// Description:
//
// # Query the list of scenarized risk control event strategies
//
// @param request - DescribeSceneRulePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSceneRulePageListResponse
func (client *Client) DescribeSceneRulePageListWithOptions(request *DescribeSceneRulePageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSceneRulePageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleAuthType) {
		query["ruleAuthType"] = request.RuleAuthType
	}

	if !dara.IsNil(request.RuleName) {
		query["ruleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleStatus) {
		query["ruleStatus"] = request.RuleStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSceneRulePageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSceneRulePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # White-boxed strategy list for risk control services
//
// Description:
//
// # Query the list of scenarized risk control event strategies
//
// @param request - DescribeSceneRulePageListRequest
//
// @return DescribeSceneRulePageListResponse
func (client *Client) DescribeSceneRulePageList(request *DescribeSceneRulePageListRequest) (_result *DescribeSceneRulePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSceneRulePageListResponse{}
	_body, _err := client.DescribeSceneRulePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Score Distribution
//
// @param request - DescribeScoreListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScoreListResponse
func (client *Client) DescribeScoreListWithOptions(request *DescribeScoreListRequest, runtime *dara.RuntimeOptions) (_result *DescribeScoreListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScoreList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScoreListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Score Distribution
//
// @param request - DescribeScoreListRequest
//
// @return DescribeScoreListResponse
func (client *Client) DescribeScoreList(request *DescribeScoreListRequest) (_result *DescribeScoreListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScoreListResponse{}
	_body, _err := client.DescribeScoreListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Score Range Quantity Analysis
//
// @param request - DescribeScoreSectionNumLineChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScoreSectionNumLineChartResponse
func (client *Client) DescribeScoreSectionNumLineChartWithOptions(request *DescribeScoreSectionNumLineChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeScoreSectionNumLineChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.ByPassEventCodes) {
		query["byPassEventCodes"] = request.ByPassEventCodes
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MainEventCodes) {
		query["mainEventCodes"] = request.MainEventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.ShuntEventCodes) {
		query["shuntEventCodes"] = request.ShuntEventCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScoreSectionNumLineChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScoreSectionNumLineChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Score Range Quantity Analysis
//
// @param request - DescribeScoreSectionNumLineChartRequest
//
// @return DescribeScoreSectionNumLineChartResponse
func (client *Client) DescribeScoreSectionNumLineChart(request *DescribeScoreSectionNumLineChartRequest) (_result *DescribeScoreSectionNumLineChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScoreSectionNumLineChartResponse{}
	_body, _err := client.DescribeScoreSectionNumLineChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Proportion of Score Ranges for Main Events/Bypass Events/Diversion Events
//
// @param request - DescribeScoreSectionPieChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScoreSectionPieChartResponse
func (client *Client) DescribeScoreSectionPieChartWithOptions(request *DescribeScoreSectionPieChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeScoreSectionPieChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.EventType) {
		query["eventType"] = request.EventType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScoreSectionPieChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScoreSectionPieChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Proportion of Score Ranges for Main Events/Bypass Events/Diversion Events
//
// @param request - DescribeScoreSectionPieChartRequest
//
// @return DescribeScoreSectionPieChartResponse
func (client *Client) DescribeScoreSectionPieChart(request *DescribeScoreSectionPieChartRequest) (_result *DescribeScoreSectionPieChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScoreSectionPieChartResponse{}
	_body, _err := client.DescribeScoreSectionPieChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Score Section Ratio Analysis
//
// @param request - DescribeScoreSectionRatioLineChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScoreSectionRatioLineChartResponse
func (client *Client) DescribeScoreSectionRatioLineChartWithOptions(request *DescribeScoreSectionRatioLineChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeScoreSectionRatioLineChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.ByPassEventCodes) {
		query["byPassEventCodes"] = request.ByPassEventCodes
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MainEventCodes) {
		query["mainEventCodes"] = request.MainEventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.ShuntEventCodes) {
		query["shuntEventCodes"] = request.ShuntEventCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScoreSectionRatioLineChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScoreSectionRatioLineChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Score Section Ratio Analysis
//
// @param request - DescribeScoreSectionRatioLineChartRequest
//
// @return DescribeScoreSectionRatioLineChartResponse
func (client *Client) DescribeScoreSectionRatioLineChart(request *DescribeScoreSectionRatioLineChartRequest) (_result *DescribeScoreSectionRatioLineChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScoreSectionRatioLineChartResponse{}
	_body, _err := client.DescribeScoreSectionRatioLineChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Task ID List
//
// @param request - DescribeSelectItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSelectItemResponse
func (client *Client) DescribeSelectItemWithOptions(request *DescribeSelectItemRequest, runtime *dara.RuntimeOptions) (_result *DescribeSelectItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSelectItem"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSelectItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Task ID List
//
// @param request - DescribeSelectItemRequest
//
// @return DescribeSelectItemResponse
func (client *Client) DescribeSelectItem(request *DescribeSelectItemRequest) (_result *DescribeSelectItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSelectItemResponse{}
	_body, _err := client.DescribeSelectItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ServiceAppkey dropdown
//
// @param request - DescribeServiceAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceAppKeyResponse
func (client *Client) DescribeServiceAppKeyWithOptions(request *DescribeServiceAppKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceAppKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceAppKey"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ServiceAppkey dropdown
//
// @param request - DescribeServiceAppKeyRequest
//
// @return DescribeServiceAppKeyResponse
func (client *Client) DescribeServiceAppKey(request *DescribeServiceAppKeyRequest) (_result *DescribeServiceAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServiceAppKeyResponse{}
	_body, _err := client.DescribeServiceAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ServiceCodeName Information
//
// @param request - DescribeServiceCodeNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceCodeNameResponse
func (client *Client) DescribeServiceCodeNameWithOptions(request *DescribeServiceCodeNameRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceCodeNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Tab) {
		query["Tab"] = request.Tab
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceCodeName"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceCodeNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ServiceCodeName Information
//
// @param request - DescribeServiceCodeNameRequest
//
// @return DescribeServiceCodeNameResponse
func (client *Client) DescribeServiceCodeName(request *DescribeServiceCodeNameRequest) (_result *DescribeServiceCodeNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServiceCodeNameResponse{}
	_body, _err := client.DescribeServiceCodeNameWithOptions(request, runtime)
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
// @param request - DescribeServiceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceListResponse
func (client *Client) DescribeServiceListWithOptions(request *DescribeServiceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceListResponse{}
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
// @param request - DescribeServiceListRequest
//
// @return DescribeServiceListResponse
func (client *Client) DescribeServiceList(request *DescribeServiceListRequest) (_result *DescribeServiceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServiceListResponse{}
	_body, _err := client.DescribeServiceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Estimate Call Information
//
// @param request - DescribeSimulationPreditInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSimulationPreditInfoResponse
func (client *Client) DescribeSimulationPreditInfoWithOptions(request *DescribeSimulationPreditInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeSimulationPreditInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RulesStr) {
		query["rulesStr"] = request.RulesStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSimulationPreditInfo"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSimulationPreditInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Estimate Call Information
//
// @param request - DescribeSimulationPreditInfoRequest
//
// @return DescribeSimulationPreditInfoResponse
func (client *Client) DescribeSimulationPreditInfo(request *DescribeSimulationPreditInfoRequest) (_result *DescribeSimulationPreditInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSimulationPreditInfoResponse{}
	_body, _err := client.DescribeSimulationPreditInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Task Record Count
//
// @param request - DescribeSimulationTaskCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSimulationTaskCountResponse
func (client *Client) DescribeSimulationTaskCountWithOptions(request *DescribeSimulationTaskCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeSimulationTaskCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataSourceConfig) {
		query["dataSourceConfig"] = request.DataSourceConfig
	}

	if !dara.IsNil(request.DataSourceType) {
		query["dataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.FiltersStr) {
		query["filtersStr"] = request.FiltersStr
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSimulationTaskCount"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSimulationTaskCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Task Record Count
//
// @param request - DescribeSimulationTaskCountRequest
//
// @return DescribeSimulationTaskCountResponse
func (client *Client) DescribeSimulationTaskCount(request *DescribeSimulationTaskCountRequest) (_result *DescribeSimulationTaskCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSimulationTaskCountResponse{}
	_body, _err := client.DescribeSimulationTaskCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Task List
//
// @param request - DescribeSimulationTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSimulationTaskListResponse
func (client *Client) DescribeSimulationTaskListWithOptions(request *DescribeSimulationTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSimulationTaskListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSimulationTaskList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSimulationTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Task List
//
// @param request - DescribeSimulationTaskListRequest
//
// @return DescribeSimulationTaskListResponse
func (client *Client) DescribeSimulationTaskList(request *DescribeSimulationTaskListRequest) (_result *DescribeSimulationTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSimulationTaskListResponse{}
	_body, _err := client.DescribeSimulationTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Project Configuration
//
// @param request - DescribeSlsUrlConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsUrlConfigResponse
func (client *Client) DescribeSlsUrlConfigWithOptions(request *DescribeSlsUrlConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsUrlConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlsUrlConfig"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsUrlConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Project Configuration
//
// @param request - DescribeSlsUrlConfigRequest
//
// @return DescribeSlsUrlConfigResponse
func (client *Client) DescribeSlsUrlConfig(request *DescribeSlsUrlConfigRequest) (_result *DescribeSlsUrlConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlsUrlConfigResponse{}
	_body, _err := client.DescribeSlsUrlConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query List of Policies Supporting Simulation
//
// @param request - DescribeSupportRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportRuleListResponse
func (client *Client) DescribeSupportRuleListWithOptions(request *DescribeSupportRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupportRuleListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupportRuleList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupportRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query List of Policies Supporting Simulation
//
// @param request - DescribeSupportRuleListRequest
//
// @return DescribeSupportRuleListResponse
func (client *Client) DescribeSupportRuleList(request *DescribeSupportRuleListRequest) (_result *DescribeSupportRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSupportRuleListResponse{}
	_body, _err := client.DescribeSupportRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag List
//
// @param request - DescribeTagListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagListResponse
func (client *Client) DescribeTagListWithOptions(request *DescribeTagListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag List
//
// @param request - DescribeTagListRequest
//
// @return DescribeTagListResponse
func (client *Client) DescribeTagList(request *DescribeTagListRequest) (_result *DescribeTagListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagListResponse{}
	_body, _err := client.DescribeTagListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag Overview List
//
// @param request - DescribeTagsBarChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsBarChartResponse
func (client *Client) DescribeTagsBarChartWithOptions(request *DescribeTagsBarChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsBarChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Result) {
		query["result"] = request.Result
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagsBarChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsBarChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag Overview List
//
// @param request - DescribeTagsBarChartRequest
//
// @return DescribeTagsBarChartResponse
func (client *Client) DescribeTagsBarChart(request *DescribeTagsBarChartRequest) (_result *DescribeTagsBarChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsBarChartResponse{}
	_body, _err := client.DescribeTagsBarChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag Fluctuation Detection
//
// @param request - DescribeTagsFluctuationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsFluctuationResponse
func (client *Client) DescribeTagsFluctuationWithOptions(request *DescribeTagsFluctuationRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsFluctuationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagsFluctuation"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsFluctuationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag Fluctuation Detection
//
// @param request - DescribeTagsFluctuationRequest
//
// @return DescribeTagsFluctuationResponse
func (client *Client) DescribeTagsFluctuation(request *DescribeTagsFluctuationRequest) (_result *DescribeTagsFluctuationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsFluctuationResponse{}
	_body, _err := client.DescribeTagsFluctuationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Tag List
//
// @param request - DescribeTagsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsListResponse
func (client *Client) DescribeTagsListWithOptions(request *DescribeTagsListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagsList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Tag List
//
// @param request - DescribeTagsListRequest
//
// @return DescribeTagsListResponse
func (client *Client) DescribeTagsList(request *DescribeTagsListRequest) (_result *DescribeTagsListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsListResponse{}
	_body, _err := client.DescribeTagsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag Hit Count Analysis
//
// @param request - DescribeTagsNumLineChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsNumLineChartResponse
func (client *Client) DescribeTagsNumLineChartWithOptions(request *DescribeTagsNumLineChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsNumLineChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.ByPassEventCodes) {
		query["byPassEventCodes"] = request.ByPassEventCodes
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MainEventCodes) {
		query["mainEventCodes"] = request.MainEventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.ShuntEventCodes) {
		query["shuntEventCodes"] = request.ShuntEventCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagsNumLineChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsNumLineChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag Hit Count Analysis
//
// @param request - DescribeTagsNumLineChartRequest
//
// @return DescribeTagsNumLineChartResponse
func (client *Client) DescribeTagsNumLineChart(request *DescribeTagsNumLineChartRequest) (_result *DescribeTagsNumLineChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsNumLineChartResponse{}
	_body, _err := client.DescribeTagsNumLineChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag Hit Ratio Analysis
//
// @param request - DescribeTagsRatioLineChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsRatioLineChartResponse
func (client *Client) DescribeTagsRatioLineChartWithOptions(request *DescribeTagsRatioLineChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsRatioLineChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.ByPassEventCodes) {
		query["byPassEventCodes"] = request.ByPassEventCodes
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MainEventCodes) {
		query["mainEventCodes"] = request.MainEventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.ShuntEventCodes) {
		query["shuntEventCodes"] = request.ShuntEventCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagsRatioLineChart"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsRatioLineChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag Hit Ratio Analysis
//
// @param request - DescribeTagsRatioLineChartRequest
//
// @return DescribeTagsRatioLineChartResponse
func (client *Client) DescribeTagsRatioLineChart(request *DescribeTagsRatioLineChartRequest) (_result *DescribeTagsRatioLineChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsRatioLineChartResponse{}
	_body, _err := client.DescribeTagsRatioLineChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag Hit Trend
//
// @param request - DescribeTagsTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsTrendResponse
func (client *Client) DescribeTagsTrendWithOptions(request *DescribeTagsTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BeginTime) {
		query["beginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Result) {
		query["result"] = request.Result
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagsTrend"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag Hit Trend
//
// @param request - DescribeTagsTrendRequest
//
// @return DescribeTagsTrendResponse
func (client *Client) DescribeTagsTrend(request *DescribeTagsTrendRequest) (_result *DescribeTagsTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsTrendResponse{}
	_body, _err := client.DescribeTagsTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Task List
//
// @param request - DescribeTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskListResponse
func (client *Client) DescribeTaskListWithOptions(request *DescribeTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTaskListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.IsPage) {
		query["IsPage"] = request.IsPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTaskList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Task List
//
// @param request - DescribeTaskListRequest
//
// @return DescribeTaskListResponse
func (client *Client) DescribeTaskList(request *DescribeTaskListRequest) (_result *DescribeTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTaskListResponse{}
	_body, _err := client.DescribeTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Task Log List
//
// @param request - DescribeTaskLogListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskLogListResponse
func (client *Client) DescribeTaskLogListWithOptions(request *DescribeTaskLogListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTaskLogListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.IsPage) {
		query["IsPage"] = request.IsPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskLogId) {
		query["TaskLogId"] = request.TaskLogId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTaskLogList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskLogListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Task Log List
//
// @param request - DescribeTaskLogListRequest
//
// @return DescribeTaskLogListResponse
func (client *Client) DescribeTaskLogList(request *DescribeTaskLogListRequest) (_result *DescribeTaskLogListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTaskLogListResponse{}
	_body, _err := client.DescribeTaskLogListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Total Event Count
//
// @param request - DescribeTemplateCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateCountResponse
func (client *Client) DescribeTemplateCountWithOptions(request *DescribeTemplateCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplateCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplateCount"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplateCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Total Event Count
//
// @param request - DescribeTemplateCountRequest
//
// @return DescribeTemplateCountResponse
func (client *Client) DescribeTemplateCount(request *DescribeTemplateCountRequest) (_result *DescribeTemplateCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplateCountResponse{}
	_body, _err := client.DescribeTemplateCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Template Download
//
// @param request - DescribeTemplateDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateDownloadResponse
func (client *Client) DescribeTemplateDownloadWithOptions(request *DescribeTemplateDownloadRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplateDownloadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplateDownload"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplateDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Template Download
//
// @param request - DescribeTemplateDownloadRequest
//
// @return DescribeTemplateDownloadResponse
func (client *Client) DescribeTemplateDownload(request *DescribeTemplateDownloadRequest) (_result *DescribeTemplateDownloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplateDownloadResponse{}
	_body, _err := client.DescribeTemplateDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Event List by Event Name
//
// @param request - DescribeTemplatePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplatePageListResponse
func (client *Client) DescribeTemplatePageListWithOptions(request *DescribeTemplatePageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplatePageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TemplateName) {
		query["templateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateSearchItem) {
		query["templateSearchItem"] = request.TemplateSearchItem
	}

	if !dara.IsNil(request.TemplateStatus) {
		query["templateStatus"] = request.TemplateStatus
	}

	if !dara.IsNil(request.TemplateType) {
		query["templateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplatePageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplatePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Event List by Event Name
//
// @param request - DescribeTemplatePageListRequest
//
// @return DescribeTemplatePageListResponse
func (client *Client) DescribeTemplatePageList(request *DescribeTemplatePageListRequest) (_result *DescribeTemplatePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplatePageListResponse{}
	_body, _err := client.DescribeTemplatePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get List of Services Used by User
//
// @param request - DescribeUsedServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsedServiceResponse
func (client *Client) DescribeUsedServiceWithOptions(request *DescribeUsedServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsedServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsedService"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsedServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get List of Services Used by User
//
// @param request - DescribeUsedServiceRequest
//
// @return DescribeUsedServiceResponse
func (client *Client) DescribeUsedService(request *DescribeUsedServiceRequest) (_result *DescribeUsedServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUsedServiceResponse{}
	_body, _err := client.DescribeUsedServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Current Logged-in User Information
//
// @param request - DescribeUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserInfoResponse
func (client *Client) DescribeUserInfoWithOptions(request *DescribeUserInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserInfo"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Current Logged-in User Information
//
// @param request - DescribeUserInfoRequest
//
// @return DescribeUserInfoResponse
func (client *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (_result *DescribeUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserInfoResponse{}
	_body, _err := client.DescribeUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Variable Binding Information
//
// @param request - DescribeVariableBindDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVariableBindDetailResponse
func (client *Client) DescribeVariableBindDetailWithOptions(request *DescribeVariableBindDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVariableBindDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DefineId) {
		query["defineId"] = request.DefineId
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVariableBindDetail"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVariableBindDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Variable Binding Information
//
// @param request - DescribeVariableBindDetailRequest
//
// @return DescribeVariableBindDetailResponse
func (client *Client) DescribeVariableBindDetail(request *DescribeVariableBindDetailRequest) (_result *DescribeVariableBindDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVariableBindDetailResponse{}
	_body, _err := client.DescribeVariableBindDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query variable details
//
// @param request - DescribeVariableDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVariableDetailResponse
func (client *Client) DescribeVariableDetailWithOptions(request *DescribeVariableDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVariableDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVariableDetail"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVariableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query variable details
//
// @param request - DescribeVariableDetailRequest
//
// @return DescribeVariableDetailResponse
func (client *Client) DescribeVariableDetail(request *DescribeVariableDetailRequest) (_result *DescribeVariableDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVariableDetailResponse{}
	_body, _err := client.DescribeVariableDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Variable Fee Information
//
// @param request - DescribeVariableFeeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVariableFeeResponse
func (client *Client) DescribeVariableFeeWithOptions(request *DescribeVariableFeeRequest, runtime *dara.RuntimeOptions) (_result *DescribeVariableFeeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Ids) {
		query["ids"] = request.Ids
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVariableFee"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVariableFeeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Variable Fee Information
//
// @param request - DescribeVariableFeeRequest
//
// @return DescribeVariableFeeResponse
func (client *Client) DescribeVariableFee(request *DescribeVariableFeeRequest) (_result *DescribeVariableFeeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVariableFeeResponse{}
	_body, _err := client.DescribeVariableFeeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Variable Details
//
// @param request - DescribeVariableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVariableListResponse
func (client *Client) DescribeVariableListWithOptions(request *DescribeVariableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVariableListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RefObjId) {
		query["refObjId"] = request.RefObjId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.TypesStr) {
		query["typesStr"] = request.TypesStr
	}

	if !dara.IsNil(request.Value) {
		query["value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVariableList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVariableListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Variable Details
//
// @param request - DescribeVariableListRequest
//
// @return DescribeVariableListResponse
func (client *Client) DescribeVariableList(request *DescribeVariableListRequest) (_result *DescribeVariableListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVariableListResponse{}
	_body, _err := client.DescribeVariableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Variable Definitions
//
// @param request - DescribeVariableMarketListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVariableMarketListResponse
func (client *Client) DescribeVariableMarketListWithOptions(request *DescribeVariableMarketListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVariableMarketListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ChargingMode) {
		query["chargingMode"] = request.ChargingMode
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Paging) {
		query["paging"] = request.Paging
	}

	if !dara.IsNil(request.QueryContent) {
		query["queryContent"] = request.QueryContent
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.ScenesStr) {
		query["scenesStr"] = request.ScenesStr
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVariableMarketList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVariableMarketListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Variable Definitions
//
// @param request - DescribeVariableMarketListRequest
//
// @return DescribeVariableMarketListResponse
func (client *Client) DescribeVariableMarketList(request *DescribeVariableMarketListRequest) (_result *DescribeVariableMarketListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVariableMarketListResponse{}
	_body, _err := client.DescribeVariableMarketListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Configuration Information
//
// @param request - DescribeVariableSceneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVariableSceneListResponse
func (client *Client) DescribeVariableSceneListWithOptions(request *DescribeVariableSceneListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVariableSceneListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BizType) {
		query["bizType"] = request.BizType
	}

	if !dara.IsNil(request.ConfigKey) {
		query["configKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Paging) {
		query["paging"] = request.Paging
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVariableSceneList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVariableSceneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Configuration Information
//
// @param request - DescribeVariableSceneListRequest
//
// @return DescribeVariableSceneListResponse
func (client *Client) DescribeVariableSceneList(request *DescribeVariableSceneListRequest) (_result *DescribeVariableSceneListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVariableSceneListResponse{}
	_body, _err := client.DescribeVariableSceneListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Variable Version List Query
//
// @param request - DescribeVersionPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVersionPageListResponse
func (client *Client) DescribeVersionPageListWithOptions(request *DescribeVersionPageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVersionPageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.ObjectCode) {
		query["objectCode"] = request.ObjectCode
	}

	if !dara.IsNil(request.ObjectId) {
		query["objectId"] = request.ObjectId
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Paging) {
		query["paging"] = request.Paging
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVersionPageList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVersionPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Variable Version List Query
//
// @param request - DescribeVersionPageListRequest
//
// @return DescribeVersionPageListResponse
func (client *Client) DescribeVersionPageList(request *DescribeVersionPageListRequest) (_result *DescribeVersionPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVersionPageListResponse{}
	_body, _err := client.DescribeVersionPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Sample List Data Download
//
// @param request - DownloadSmapleBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadSmapleBatchResponse
func (client *Client) DownloadSmapleBatchWithOptions(request *DownloadSmapleBatchRequest, runtime *dara.RuntimeOptions) (_result *DownloadSmapleBatchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BatchUuid) {
		query["batchUuid"] = request.BatchUuid
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadSmapleBatch"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadSmapleBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Sample List Data Download
//
// @param request - DownloadSmapleBatchRequest
//
// @return DownloadSmapleBatchResponse
func (client *Client) DownloadSmapleBatch(request *DownloadSmapleBatchRequest) (_result *DownloadSmapleBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadSmapleBatchResponse{}
	_body, _err := client.DownloadSmapleBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Custom Variable Test
//
// @param request - ExpressionTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpressionTestResponse
func (client *Client) ExpressionTestWithOptions(request *ExpressionTestRequest, runtime *dara.RuntimeOptions) (_result *ExpressionTestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Expression) {
		query["expression"] = request.Expression
	}

	if !dara.IsNil(request.ExpressionVariable) {
		query["expressionVariable"] = request.ExpressionVariable
	}

	if !dara.IsNil(request.ExpressionVariableIds) {
		query["expressionVariableIds"] = request.ExpressionVariableIds
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExpressionTest"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExpressionTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Custom Variable Test
//
// @param request - ExpressionTestRequest
//
// @return ExpressionTestResponse
func (client *Client) ExpressionTest(request *ExpressionTestRequest) (_result *ExpressionTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExpressionTestResponse{}
	_body, _err := client.ExpressionTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # File Upload
//
// @param request - FileUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileUploadResponse
func (client *Client) FileUploadWithOptions(request *FileUploadRequest, runtime *dara.RuntimeOptions) (_result *FileUploadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Tab) {
		query["Tab"] = request.Tab
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FileUpload"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FileUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # File Upload
//
// @param request - FileUploadRequest
//
// @return FileUploadResponse
func (client *Client) FileUpload(request *FileUploadRequest) (_result *FileUploadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FileUploadResponse{}
	_body, _err := client.FileUploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch Import
//
// @param request - ImportFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportFieldResponse
func (client *Client) ImportFieldWithOptions(request *ImportFieldRequest, runtime *dara.RuntimeOptions) (_result *ImportFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportField"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Import
//
// @param request - ImportFieldRequest
//
// @return ImportFieldResponse
func (client *Client) ImportField(request *ImportFieldRequest) (_result *ImportFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportFieldResponse{}
	_body, _err := client.ImportFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create or Import Name List
//
// @param request - ImportNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportNameListResponse
func (client *Client) ImportNameListWithOptions(request *ImportNameListRequest, runtime *dara.RuntimeOptions) (_result *ImportNameListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.Data) {
		query["data"] = request.Data
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.ImportType) {
		query["importType"] = request.ImportType
	}

	if !dara.IsNil(request.Memo) {
		query["memo"] = request.Memo
	}

	if !dara.IsNil(request.NameListType) {
		query["nameListType"] = request.NameListType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	if !dara.IsNil(request.VariableId) {
		query["variableId"] = request.VariableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportNameList"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create or Import Name List
//
// @param request - ImportNameListRequest
//
// @return ImportNameListResponse
func (client *Client) ImportNameList(request *ImportNameListRequest) (_result *ImportNameListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportNameListResponse{}
	_body, _err := client.ImportNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Import Template Event
//
// @param request - ImportTemplateEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportTemplateEventResponse
func (client *Client) ImportTemplateEventWithOptions(request *ImportTemplateEventRequest, runtime *dara.RuntimeOptions) (_result *ImportTemplateEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventTemplateIds) {
		query["eventTemplateIds"] = request.EventTemplateIds
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportTemplateEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportTemplateEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Import Template Event
//
// @param request - ImportTemplateEventRequest
//
// @return ImportTemplateEventResponse
func (client *Client) ImportTemplateEvent(request *ImportTemplateEventRequest) (_result *ImportTemplateEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportTemplateEventResponse{}
	_body, _err := client.ImportTemplateEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Variable Definition
//
// @param request - ListVariableDefineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVariableDefineResponse
func (client *Client) ListVariableDefineWithOptions(request *ListVariableDefineRequest, runtime *dara.RuntimeOptions) (_result *ListVariableDefineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.AllowBind) {
		query["allowBind"] = request.AllowBind
	}

	if !dara.IsNil(request.ChargingMode) {
		query["chargingMode"] = request.ChargingMode
	}

	if !dara.IsNil(request.CurrentPage) {
		query["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Paging) {
		query["paging"] = request.Paging
	}

	if !dara.IsNil(request.QueryContent) {
		query["queryContent"] = request.QueryContent
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RoleType) {
		query["roleType"] = request.RoleType
	}

	if !dara.IsNil(request.ScenesStr) {
		query["scenesStr"] = request.ScenesStr
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	if !dara.IsNil(request.TypesStr) {
		query["typesStr"] = request.TypesStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVariableDefine"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVariableDefineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Variable Definition
//
// @param request - ListVariableDefineRequest
//
// @return ListVariableDefineResponse
func (client *Client) ListVariableDefine(request *ListVariableDefineRequest) (_result *ListVariableDefineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVariableDefineResponse{}
	_body, _err := client.ListVariableDefineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete
//
// @param request - ModelDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelDeleteResponse
func (client *Client) ModelDeleteWithOptions(request *ModelDeleteRequest, runtime *dara.RuntimeOptions) (_result *ModelDeleteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelDelete"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete
//
// @param request - ModelDeleteRequest
//
// @return ModelDeleteResponse
func (client *Client) ModelDelete(request *ModelDeleteRequest) (_result *ModelDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModelDeleteResponse{}
	_body, _err := client.ModelDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # File Upload
//
// @param request - ModelFileUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelFileUploadResponse
func (client *Client) ModelFileUploadWithOptions(request *ModelFileUploadRequest, runtime *dara.RuntimeOptions) (_result *ModelFileUploadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectName) {
		query["ObjectName"] = request.ObjectName
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelFileUpload"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelFileUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # File Upload
//
// @param request - ModelFileUploadRequest
//
// @return ModelFileUploadResponse
func (client *Client) ModelFileUpload(request *ModelFileUploadRequest) (_result *ModelFileUploadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModelFileUploadResponse{}
	_body, _err := client.ModelFileUploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Enable, Disable
//
// @param request - ModelIsUsingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelIsUsingResponse
func (client *Client) ModelIsUsingWithOptions(request *ModelIsUsingRequest, runtime *dara.RuntimeOptions) (_result *ModelIsUsingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelCode) {
		query["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelIsUsing"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelIsUsingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enable, Disable
//
// @param request - ModelIsUsingRequest
//
// @return ModelIsUsingResponse
func (client *Client) ModelIsUsing(request *ModelIsUsingRequest) (_result *ModelIsUsingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModelIsUsingResponse{}
	_body, _err := client.ModelIsUsingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Is Model Name Duplicated
//
// @param request - ModelNameIsDuplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelNameIsDuplicationResponse
func (client *Client) ModelNameIsDuplicationWithOptions(request *ModelNameIsDuplicationRequest, runtime *dara.RuntimeOptions) (_result *ModelNameIsDuplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelNameIsDuplication"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelNameIsDuplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Is Model Name Duplicated
//
// @param request - ModelNameIsDuplicationRequest
//
// @return ModelNameIsDuplicationResponse
func (client *Client) ModelNameIsDuplication(request *ModelNameIsDuplicationRequest) (_result *ModelNameIsDuplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModelNameIsDuplicationResponse{}
	_body, _err := client.ModelNameIsDuplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Template Download
//
// @param request - ModelSampleDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelSampleDownloadResponse
func (client *Client) ModelSampleDownloadWithOptions(request *ModelSampleDownloadRequest, runtime *dara.RuntimeOptions) (_result *ModelSampleDownloadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelSampleDownload"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelSampleDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Template Download
//
// @param request - ModelSampleDownloadRequest
//
// @return ModelSampleDownloadResponse
func (client *Client) ModelSampleDownload(request *ModelSampleDownloadRequest) (_result *ModelSampleDownloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModelSampleDownloadResponse{}
	_body, _err := client.ModelSampleDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Memo
//
// @param request - ModifyAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppKeyResponse
func (client *Client) ModifyAppKeyWithOptions(request *ModifyAppKeyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.AppKey) {
		query["appKey"] = request.AppKey
	}

	if !dara.IsNil(request.Memo) {
		query["memo"] = request.Memo
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppKey"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Memo
//
// @param request - ModifyAppKeyRequest
//
// @return ModifyAppKeyResponse
func (client *Client) ModifyAppKey(request *ModifyAppKeyRequest) (_result *ModifyAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppKeyResponse{}
	_body, _err := client.ModifyAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit Accumulated Variable
//
// @param request - ModifyCustVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustVariableResponse
func (client *Client) ModifyCustVariableWithOptions(request *ModifyCustVariableRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Condition) {
		query["condition"] = request.Condition
	}

	if !dara.IsNil(request.DataVersion) {
		query["dataVersion"] = request.DataVersion
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EventCodes) {
		query["eventCodes"] = request.EventCodes
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Outputs) {
		query["outputs"] = request.Outputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Accumulated Variable
//
// @param request - ModifyCustVariableRequest
//
// @return ModifyCustVariableResponse
func (client *Client) ModifyCustVariable(request *ModifyCustVariableRequest) (_result *ModifyCustVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCustVariableResponse{}
	_body, _err := client.ModifyCustVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit Event
//
// @param request - ModifyEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEventResponse
func (client *Client) ModifyEventWithOptions(request *ModifyEventRequest, runtime *dara.RuntimeOptions) (_result *ModifyEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BizVersion) {
		query["bizVersion"] = request.BizVersion
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.EventName) {
		query["eventName"] = request.EventName
	}

	if !dara.IsNil(request.InputFieldsStr) {
		query["inputFieldsStr"] = request.InputFieldsStr
	}

	if !dara.IsNil(request.Memo) {
		query["memo"] = request.Memo
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TemplateType) {
		query["templateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Event
//
// @param request - ModifyEventRequest
//
// @return ModifyEventResponse
func (client *Client) ModifyEvent(request *ModifyEventRequest) (_result *ModifyEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEventResponse{}
	_body, _err := client.ModifyEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Event Status
//
// @param request - ModifyEventStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEventStatusResponse
func (client *Client) ModifyEventStatusWithOptions(request *ModifyEventStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyEventStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.FromEventSatus) {
		query["fromEventSatus"] = request.FromEventSatus
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.ToEventSatus) {
		query["toEventSatus"] = request.ToEventSatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEventStatus"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEventStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Event Status
//
// @param request - ModifyEventStatusRequest
//
// @return ModifyEventStatusResponse
func (client *Client) ModifyEventStatus(request *ModifyEventStatusRequest) (_result *ModifyEventStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEventStatusResponse{}
	_body, _err := client.ModifyEventStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit Custom Variable
//
// @param request - ModifyExpressionVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressionVariableResponse
func (client *Client) ModifyExpressionVariableWithOptions(request *ModifyExpressionVariableRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressionVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataVersion) {
		query["dataVersion"] = request.DataVersion
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Expression) {
		query["expression"] = request.Expression
	}

	if !dara.IsNil(request.ExpressionTitle) {
		query["expressionTitle"] = request.ExpressionTitle
	}

	if !dara.IsNil(request.ExpressionVariable) {
		query["expressionVariable"] = request.ExpressionVariable
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Outlier) {
		query["outlier"] = request.Outlier
	}

	if !dara.IsNil(request.Outputs) {
		query["outputs"] = request.Outputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressionVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressionVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Custom Variable
//
// @param request - ModifyExpressionVariableRequest
//
// @return ModifyExpressionVariableResponse
func (client *Client) ModifyExpressionVariable(request *ModifyExpressionVariableRequest) (_result *ModifyExpressionVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressionVariableResponse{}
	_body, _err := client.ModifyExpressionVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Field
//
// @param request - ModifyFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFieldResponse
func (client *Client) ModifyFieldWithOptions(request *ModifyFieldRequest, runtime *dara.RuntimeOptions) (_result *ModifyFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Classify) {
		query["classify"] = request.Classify
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EnumData) {
		query["enumData"] = request.EnumData
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyField"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Field
//
// @param request - ModifyFieldRequest
//
// @return ModifyFieldResponse
func (client *Client) ModifyField(request *ModifyFieldRequest) (_result *ModifyFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFieldResponse{}
	_body, _err := client.ModifyFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Policy Priority
//
// @param request - ModifyRulePriorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRulePriorityResponse
func (client *Client) ModifyRulePriorityWithOptions(request *ModifyRulePriorityRequest, runtime *dara.RuntimeOptions) (_result *ModifyRulePriorityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.Priority) {
		query["priority"] = request.Priority
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRulePriority"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRulePriorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Policy Priority
//
// @param request - ModifyRulePriorityRequest
//
// @return ModifyRulePriorityResponse
func (client *Client) ModifyRulePriority(request *ModifyRulePriorityRequest) (_result *ModifyRulePriorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRulePriorityResponse{}
	_body, _err := client.ModifyRulePriorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Change the status of a policy version application
//
// @param request - ModifyRuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleStatusResponse
func (client *Client) ModifyRuleStatusWithOptions(request *ModifyRuleStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyRuleStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ApplyUserId) {
		query["applyUserId"] = request.ApplyUserId
	}

	if !dara.IsNil(request.ApplyUserName) {
		query["applyUserName"] = request.ApplyUserName
	}

	if !dara.IsNil(request.AuditRemark) {
		query["auditRemark"] = request.AuditRemark
	}

	if !dara.IsNil(request.AuditUserId) {
		query["auditUserId"] = request.AuditUserId
	}

	if !dara.IsNil(request.AuditUserName) {
		query["auditUserName"] = request.AuditUserName
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.EventType) {
		query["eventType"] = request.EventType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleAuditType) {
		query["ruleAuditType"] = request.RuleAuditType
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleVersionId) {
		query["ruleVersionId"] = request.RuleVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRuleStatus"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Change the status of a policy version application
//
// @param request - ModifyRuleStatusRequest
//
// @return ModifyRuleStatusResponse
func (client *Client) ModifyRuleStatus(request *ModifyRuleStatusRequest) (_result *ModifyRuleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRuleStatusResponse{}
	_body, _err := client.ModifyRuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Activate Service
//
// @param request - OpenConsoleSlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenConsoleSlsResponse
func (client *Client) OpenConsoleSlsWithOptions(request *OpenConsoleSlsRequest, runtime *dara.RuntimeOptions) (_result *OpenConsoleSlsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Scene) {
		query["scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenConsoleSls"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenConsoleSlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Activate Service
//
// @param request - OpenConsoleSlsRequest
//
// @return OpenConsoleSlsResponse
func (client *Client) OpenConsoleSls(request *OpenConsoleSlsRequest) (_result *OpenConsoleSlsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenConsoleSlsResponse{}
	_body, _err := client.OpenConsoleSlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Operate Favorites
//
// @param request - OperateFavoriteVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateFavoriteVariableResponse
func (client *Client) OperateFavoriteVariableWithOptions(request *OperateFavoriteVariableRequest, runtime *dara.RuntimeOptions) (_result *OperateFavoriteVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Operate) {
		query["operate"] = request.Operate
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateFavoriteVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateFavoriteVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Operate Favorites
//
// @param request - OperateFavoriteVariableRequest
//
// @return OperateFavoriteVariableResponse
func (client *Client) OperateFavoriteVariable(request *OperateFavoriteVariableRequest) (_result *OperateFavoriteVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateFavoriteVariableResponse{}
	_body, _err := client.OperateFavoriteVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Enterprise Verification
//
// @param request - PermissionCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PermissionCheckResponse
func (client *Client) PermissionCheckWithOptions(request *PermissionCheckRequest, runtime *dara.RuntimeOptions) (_result *PermissionCheckResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PermissionCheck"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PermissionCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enterprise Verification
//
// @param request - PermissionCheckRequest
//
// @return PermissionCheckResponse
func (client *Client) PermissionCheck(request *PermissionCheckRequest) (_result *PermissionCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PermissionCheckResponse{}
	_body, _err := client.PermissionCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// createTask
//
// @param request - PocCreateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PocCreateTaskResponse
func (client *Client) PocCreateTaskWithOptions(request *PocCreateTaskRequest, runtime *dara.RuntimeOptions) (_result *PocCreateTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DateFormat) {
		query["DateFormat"] = request.DateFormat
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PocCreateTask"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PocCreateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// createTask
//
// @param request - PocCreateTaskRequest
//
// @return PocCreateTaskResponse
func (client *Client) PocCreateTask(request *PocCreateTaskRequest) (_result *PocCreateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PocCreateTaskResponse{}
	_body, _err := client.PocCreateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # PocGetDownloadUrl
//
// @param request - PocGetDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PocGetDownloadUrlResponse
func (client *Client) PocGetDownloadUrlWithOptions(request *PocGetDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *PocGetDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PocGetDownloadUrl"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PocGetDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # PocGetDownloadUrl
//
// @param request - PocGetDownloadUrlRequest
//
// @return PocGetDownloadUrlResponse
func (client *Client) PocGetDownloadUrl(request *PocGetDownloadUrlRequest) (_result *PocGetDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PocGetDownloadUrlResponse{}
	_body, _err := client.PocGetDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// getToken
//
// @param request - PocGetTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PocGetTokenResponse
func (client *Client) PocGetTokenWithOptions(request *PocGetTokenRequest, runtime *dara.RuntimeOptions) (_result *PocGetTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.ServiceCode) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PocGetToken"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PocGetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// getToken
//
// @param request - PocGetTokenRequest
//
// @return PocGetTokenResponse
func (client *Client) PocGetToken(request *PocGetTokenRequest) (_result *PocGetTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PocGetTokenResponse{}
	_body, _err := client.PocGetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// sendData
//
// @param request - PocSendDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PocSendDataResponse
func (client *Client) PocSendDataWithOptions(request *PocSendDataRequest, runtime *dara.RuntimeOptions) (_result *PocSendDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchNo) {
		query["BatchNo"] = request.BatchNo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ParamsList) {
		query["ParamsList"] = request.ParamsList
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PocSendData"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PocSendDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// sendData
//
// @param request - PocSendDataRequest
//
// @return PocSendDataResponse
func (client *Client) PocSendData(request *PocSendDataRequest) (_result *PocSendDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PocSendDataResponse{}
	_body, _err := client.PocSendDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query White-box Strategy Details
//
// @param request - QueryAuthRuleDetailByRuleIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAuthRuleDetailByRuleIdResponse
func (client *Client) QueryAuthRuleDetailByRuleIdWithOptions(request *QueryAuthRuleDetailByRuleIdRequest, runtime *dara.RuntimeOptions) (_result *QueryAuthRuleDetailByRuleIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleVersionId) {
		query["ruleVersionId"] = request.RuleVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAuthRuleDetailByRuleId"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAuthRuleDetailByRuleIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query White-box Strategy Details
//
// @param request - QueryAuthRuleDetailByRuleIdRequest
//
// @return QueryAuthRuleDetailByRuleIdResponse
func (client *Client) QueryAuthRuleDetailByRuleId(request *QueryAuthRuleDetailByRuleIdRequest) (_result *QueryAuthRuleDetailByRuleIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAuthRuleDetailByRuleIdResponse{}
	_body, _err := client.QueryAuthRuleDetailByRuleIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Recall.
//
// @param request - RecallRuleAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallRuleAuditResponse
func (client *Client) RecallRuleAuditWithOptions(request *RecallRuleAuditRequest, runtime *dara.RuntimeOptions) (_result *RecallRuleAuditResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecallRuleAudit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecallRuleAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Recall.
//
// @param request - RecallRuleAuditRequest
//
// @return RecallRuleAuditResponse
func (client *Client) RecallRuleAudit(request *RecallRuleAuditRequest) (_result *RecallRuleAuditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecallRuleAuditResponse{}
	_body, _err := client.RecallRuleAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Event
//
// @param request - RemoveEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveEventResponse
func (client *Client) RemoveEventWithOptions(request *RemoveEventRequest, runtime *dara.RuntimeOptions) (_result *RemoveEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.CreateType) {
		query["createType"] = request.CreateType
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["templateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Event
//
// @param request - RemoveEventRequest
//
// @return RemoveEventResponse
func (client *Client) RemoveEvent(request *RemoveEventRequest) (_result *RemoveEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveEventResponse{}
	_body, _err := client.RemoveEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Template Download
//
// @param request - SampleFileDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SampleFileDownloadResponse
func (client *Client) SampleFileDownloadWithOptions(request *SampleFileDownloadRequest, runtime *dara.RuntimeOptions) (_result *SampleFileDownloadResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.Tab) {
		query["Tab"] = request.Tab
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SampleFileDownload"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SampleFileDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Template Download
//
// @param request - SampleFileDownloadRequest
//
// @return SampleFileDownloadResponse
func (client *Client) SampleFileDownload(request *SampleFileDownloadRequest) (_result *SampleFileDownloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SampleFileDownloadResponse{}
	_body, _err := client.SampleFileDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Save Custom Columns
//
// @param request - SaveAnalysisColumnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAnalysisColumnResponse
func (client *Client) SaveAnalysisColumnWithOptions(request *SaveAnalysisColumnRequest, runtime *dara.RuntimeOptions) (_result *SaveAnalysisColumnResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Columns) {
		query["columns"] = request.Columns
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveAnalysisColumn"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveAnalysisColumnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Save Custom Columns
//
// @param request - SaveAnalysisColumnRequest
//
// @return SaveAnalysisColumnResponse
func (client *Client) SaveAnalysisColumn(request *SaveAnalysisColumnRequest) (_result *SaveAnalysisColumnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveAnalysisColumnResponse{}
	_body, _err := client.SaveAnalysisColumnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Bypass/Shunt Configuration
//
// @param request - SaveByPassOrShuntEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveByPassOrShuntEventResponse
func (client *Client) SaveByPassOrShuntEventWithOptions(request *SaveByPassOrShuntEventRequest, runtime *dara.RuntimeOptions) (_result *SaveByPassOrShuntEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventId) {
		query["eventId"] = request.EventId
	}

	if !dara.IsNil(request.EventName) {
		query["eventName"] = request.EventName
	}

	if !dara.IsNil(request.EventType) {
		query["eventType"] = request.EventType
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveByPassOrShuntEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveByPassOrShuntEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Bypass/Shunt Configuration
//
// @param request - SaveByPassOrShuntEventRequest
//
// @return SaveByPassOrShuntEventResponse
func (client *Client) SaveByPassOrShuntEvent(request *SaveByPassOrShuntEventRequest) (_result *SaveByPassOrShuntEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveByPassOrShuntEventResponse{}
	_body, _err := client.SaveByPassOrShuntEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Start/Stop Bypass Event
//
// @param request - StartOrStopByPassShuntEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartOrStopByPassShuntEventResponse
func (client *Client) StartOrStopByPassShuntEventWithOptions(request *StartOrStopByPassShuntEventRequest, runtime *dara.RuntimeOptions) (_result *StartOrStopByPassShuntEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventId) {
		query["eventId"] = request.EventId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartOrStopByPassShuntEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartOrStopByPassShuntEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Start/Stop Bypass Event
//
// @param request - StartOrStopByPassShuntEventRequest
//
// @return StartOrStopByPassShuntEventResponse
func (client *Client) StartOrStopByPassShuntEvent(request *StartOrStopByPassShuntEventRequest) (_result *StartOrStopByPassShuntEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartOrStopByPassShuntEventResponse{}
	_body, _err := client.StartOrStopByPassShuntEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Start Task Execution
//
// @param request - StartSimulationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSimulationTaskResponse
func (client *Client) StartSimulationTaskWithOptions(request *StartSimulationTaskRequest, runtime *dara.RuntimeOptions) (_result *StartSimulationTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSimulationTask"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSimulationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Start Task Execution
//
// @param request - StartSimulationTaskRequest
//
// @return StartSimulationTaskResponse
func (client *Client) StartSimulationTask(request *StartSimulationTaskRequest) (_result *StartSimulationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartSimulationTaskResponse{}
	_body, _err := client.StartSimulationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Stop Task
//
// @param request - StopSimulationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSimulationTaskResponse
func (client *Client) StopSimulationTaskWithOptions(request *StopSimulationTaskRequest, runtime *dara.RuntimeOptions) (_result *StopSimulationTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopSimulationTask"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopSimulationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Stop Task
//
// @param request - StopSimulationTaskRequest
//
// @return StopSimulationTaskResponse
func (client *Client) StopSimulationTask(request *StopSimulationTaskRequest) (_result *StopSimulationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopSimulationTaskResponse{}
	_body, _err := client.StopSimulationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Custom Variable Switch
//
// @param request - SwitchExpressionVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchExpressionVariableResponse
func (client *Client) SwitchExpressionVariableWithOptions(request *SwitchExpressionVariableRequest, runtime *dara.RuntimeOptions) (_result *SwitchExpressionVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataVersion) {
		query["dataVersion"] = request.DataVersion
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchExpressionVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchExpressionVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Custom Variable Switch
//
// @param request - SwitchExpressionVariableRequest
//
// @return SwitchExpressionVariableResponse
func (client *Client) SwitchExpressionVariable(request *SwitchExpressionVariableRequest) (_result *SwitchExpressionVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchExpressionVariableResponse{}
	_body, _err := client.SwitchExpressionVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Field Switch
//
// @param request - SwitchFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchFieldResponse
func (client *Client) SwitchFieldWithOptions(request *SwitchFieldRequest, runtime *dara.RuntimeOptions) (_result *SwitchFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchField"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Field Switch
//
// @param request - SwitchFieldRequest
//
// @return SwitchFieldResponse
func (client *Client) SwitchField(request *SwitchFieldRequest) (_result *SwitchFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchFieldResponse{}
	_body, _err := client.SwitchFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Variable Enable/Disable
//
// @param request - SwitchQueryVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchQueryVariableResponse
func (client *Client) SwitchQueryVariableWithOptions(request *SwitchQueryVariableRequest, runtime *dara.RuntimeOptions) (_result *SwitchQueryVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchQueryVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchQueryVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Variable Enable/Disable
//
// @param request - SwitchQueryVariableRequest
//
// @return SwitchQueryVariableResponse
func (client *Client) SwitchQueryVariable(request *SwitchQueryVariableRequest) (_result *SwitchQueryVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchQueryVariableResponse{}
	_body, _err := client.SwitchQueryVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # One-click switch online
//
// @param request - SwitchToOnlineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchToOnlineResponse
func (client *Client) SwitchToOnlineWithOptions(request *SwitchToOnlineRequest, runtime *dara.RuntimeOptions) (_result *SwitchToOnlineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventId) {
		query["eventId"] = request.EventId
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchToOnline"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchToOnlineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # One-click switch online
//
// @param request - SwitchToOnlineRequest
//
// @return SwitchToOnlineResponse
func (client *Client) SwitchToOnline(request *SwitchToOnlineRequest) (_result *SwitchToOnlineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchToOnlineResponse{}
	_body, _err := client.SwitchToOnlineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Cumulative Variable Switch
//
// @param request - SwitchVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchVariableResponse
func (client *Client) SwitchVariableWithOptions(request *SwitchVariableRequest, runtime *dara.RuntimeOptions) (_result *SwitchVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataVersion) {
		query["dataVersion"] = request.DataVersion
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Cumulative Variable Switch
//
// @param request - SwitchVariableRequest
//
// @return SwitchVariableResponse
func (client *Client) SwitchVariable(request *SwitchVariableRequest) (_result *SwitchVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchVariableResponse{}
	_body, _err := client.SwitchVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Determine if the task name is duplicated
//
// @param request - TaskNameByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskNameByUserIdResponse
func (client *Client) TaskNameByUserIdWithOptions(request *TaskNameByUserIdRequest, runtime *dara.RuntimeOptions) (_result *TaskNameByUserIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegId) {
		query["RegId"] = request.RegId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TaskNameByUserId"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TaskNameByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Determine if the task name is duplicated
//
// @param request - TaskNameByUserIdRequest
//
// @return TaskNameByUserIdResponse
func (client *Client) TaskNameByUserId(request *TaskNameByUserIdRequest) (_result *TaskNameByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TaskNameByUserIdResponse{}
	_body, _err := client.TaskNameByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Query Conditions
//
// @param request - UpdateAnalysisConditionFavoriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAnalysisConditionFavoriteResponse
func (client *Client) UpdateAnalysisConditionFavoriteWithOptions(request *UpdateAnalysisConditionFavoriteRequest, runtime *dara.RuntimeOptions) (_result *UpdateAnalysisConditionFavoriteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Condition) {
		query["condition"] = request.Condition
	}

	if !dara.IsNil(request.EventBeginTime) {
		query["eventBeginTime"] = request.EventBeginTime
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.EventEndTime) {
		query["eventEndTime"] = request.EventEndTime
	}

	if !dara.IsNil(request.FieldName) {
		query["fieldName"] = request.FieldName
	}

	if !dara.IsNil(request.FieldValue) {
		query["fieldValue"] = request.FieldValue
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAnalysisConditionFavorite"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAnalysisConditionFavoriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Query Conditions
//
// @param request - UpdateAnalysisConditionFavoriteRequest
//
// @return UpdateAnalysisConditionFavoriteResponse
func (client *Client) UpdateAnalysisConditionFavorite(request *UpdateAnalysisConditionFavoriteRequest) (_result *UpdateAnalysisConditionFavoriteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAnalysisConditionFavoriteResponse{}
	_body, _err := client.UpdateAnalysisConditionFavoriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Approval
//
// @param request - UpdateAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuditResponse
func (client *Client) UpdateAuditWithOptions(request *UpdateAuditRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuditResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.AuditMsg) {
		query["auditMsg"] = request.AuditMsg
	}

	if !dara.IsNil(request.AuditRelationType) {
		query["auditRelationType"] = request.AuditRelationType
	}

	if !dara.IsNil(request.AuditStatus) {
		query["auditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAudit"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Approval
//
// @param request - UpdateAuditRequest
//
// @return UpdateAuditResponse
func (client *Client) UpdateAudit(request *UpdateAuditRequest) (_result *UpdateAuditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAuditResponse{}
	_body, _err := client.UpdateAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Authorization Policy
//
// @param request - UpdateAuthRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthRuleResponse
func (client *Client) UpdateAuthRuleWithOptions(request *UpdateAuthRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleActions) {
		query["ruleActions"] = request.RuleActions
	}

	if !dara.IsNil(request.RuleExpressions) {
		query["ruleExpressions"] = request.RuleExpressions
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleVersionId) {
		query["ruleVersionId"] = request.RuleVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthRule"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Authorization Policy
//
// @param request - UpdateAuthRuleRequest
//
// @return UpdateAuthRuleResponse
func (client *Client) UpdateAuthRule(request *UpdateAuthRuleRequest) (_result *UpdateAuthRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAuthRuleResponse{}
	_body, _err := client.UpdateAuthRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Edit Bypass Event
//
// @param request - UpdateByPassShuntEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateByPassShuntEventResponse
func (client *Client) UpdateByPassShuntEventWithOptions(request *UpdateByPassShuntEventRequest, runtime *dara.RuntimeOptions) (_result *UpdateByPassShuntEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.EventId) {
		query["eventId"] = request.EventId
	}

	if !dara.IsNil(request.EventName) {
		query["eventName"] = request.EventName
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateByPassShuntEvent"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateByPassShuntEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Bypass Event
//
// @param request - UpdateByPassShuntEventRequest
//
// @return UpdateByPassShuntEventResponse
func (client *Client) UpdateByPassShuntEvent(request *UpdateByPassShuntEventRequest) (_result *UpdateByPassShuntEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateByPassShuntEventResponse{}
	_body, _err := client.UpdateByPassShuntEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Data Source
//
// @param request - UpdateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSourceWithOptions(request *UpdateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OssKey) {
		query["ossKey"] = request.OssKey
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSource"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Data Source
//
// @param request - UpdateDataSourceRequest
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSource(request *UpdateDataSourceRequest) (_result *UpdateDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.UpdateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Custom Query Variable
//
// @param request - UpdateQueryVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQueryVariableResponse
func (client *Client) UpdateQueryVariableWithOptions(request *UpdateQueryVariableRequest, runtime *dara.RuntimeOptions) (_result *UpdateQueryVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.DataSourceCode) {
		query["dataSourceCode"] = request.DataSourceCode
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Expression) {
		query["expression"] = request.Expression
	}

	if !dara.IsNil(request.ExpressionTitle) {
		query["expressionTitle"] = request.ExpressionTitle
	}

	if !dara.IsNil(request.ExpressionVariable) {
		query["expressionVariable"] = request.ExpressionVariable
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Outlier) {
		query["outlier"] = request.Outlier
	}

	if !dara.IsNil(request.Outputs) {
		query["outputs"] = request.Outputs
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Title) {
		query["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQueryVariable"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQueryVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Custom Query Variable
//
// @param request - UpdateQueryVariableRequest
//
// @return UpdateQueryVariableResponse
func (client *Client) UpdateQueryVariable(request *UpdateQueryVariableRequest) (_result *UpdateQueryVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQueryVariableResponse{}
	_body, _err := client.UpdateQueryVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Policy
//
// @param request - UpdateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleResponse
func (client *Client) UpdateRuleWithOptions(request *UpdateRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.LogicExpression) {
		query["logicExpression"] = request.LogicExpression
	}

	if !dara.IsNil(request.Memo) {
		query["memo"] = request.Memo
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleActions) {
		query["ruleActions"] = request.RuleActions
	}

	if !dara.IsNil(request.RuleBody) {
		query["ruleBody"] = request.RuleBody
	}

	if !dara.IsNil(request.RuleExpressions) {
		query["ruleExpressions"] = request.RuleExpressions
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["ruleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleStatus) {
		query["ruleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.RuleType) {
		query["ruleType"] = request.RuleType
	}

	if !dara.IsNil(request.RuleVersionId) {
		query["ruleVersionId"] = request.RuleVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRule"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Policy
//
// @param request - UpdateRuleRequest
//
// @return UpdateRuleResponse
func (client *Client) UpdateRule(request *UpdateRuleRequest) (_result *UpdateRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleResponse{}
	_body, _err := client.UpdateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Basic Policy Information
//
// @param request - UpdateRuleBaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleBaseResponse
func (client *Client) UpdateRuleBaseWithOptions(request *UpdateRuleBaseRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleBaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ConsoleRuleId) {
		query["consoleRuleId"] = request.ConsoleRuleId
	}

	if !dara.IsNil(request.EventCode) {
		query["eventCode"] = request.EventCode
	}

	if !dara.IsNil(request.Memo) {
		query["memo"] = request.Memo
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.RuleId) {
		query["ruleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["ruleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRuleBase"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleBaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Basic Policy Information
//
// @param request - UpdateRuleBaseRequest
//
// @return UpdateRuleBaseResponse
func (client *Client) UpdateRuleBase(request *UpdateRuleBaseRequest) (_result *UpdateRuleBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleBaseResponse{}
	_body, _err := client.UpdateRuleBaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch Update Samples
//
// @param request - UpdateSampleBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSampleBatchResponse
func (client *Client) UpdateSampleBatchWithOptions(request *UpdateSampleBatchRequest, runtime *dara.RuntimeOptions) (_result *UpdateSampleBatchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Ids) {
		query["ids"] = request.Ids
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.Tags) {
		query["tags"] = request.Tags
	}

	if !dara.IsNil(request.Versions) {
		query["versions"] = request.Versions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSampleBatch"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSampleBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Update Samples
//
// @param request - UpdateSampleBatchRequest
//
// @return UpdateSampleBatchResponse
func (client *Client) UpdateSampleBatch(request *UpdateSampleBatchRequest) (_result *UpdateSampleBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSampleBatchResponse{}
	_body, _err := client.UpdateSampleBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Sample Inspection
//
// @param request - UploadFileCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadFileCheckResponse
func (client *Client) UploadFileCheckWithOptions(request *UploadFileCheckRequest, runtime *dara.RuntimeOptions) (_result *UploadFileCheckResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.BatchName) {
		query["batchName"] = request.BatchName
	}

	if !dara.IsNil(request.DataType) {
		query["dataType"] = request.DataType
	}

	if !dara.IsNil(request.OssFileName) {
		query["ossFileName"] = request.OssFileName
	}

	if !dara.IsNil(request.RegId) {
		query["regId"] = request.RegId
	}

	if !dara.IsNil(request.SampleTagType) {
		query["sampleTagType"] = request.SampleTagType
	}

	if !dara.IsNil(request.ServiceList) {
		query["serviceList"] = request.ServiceList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadFileCheck"),
		Version:     dara.String("2021-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadFileCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Sample Inspection
//
// @param request - UploadFileCheckRequest
//
// @return UploadFileCheckResponse
func (client *Client) UploadFileCheck(request *UploadFileCheckRequest) (_result *UploadFileCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadFileCheckResponse{}
	_body, _err := client.UploadFileCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
