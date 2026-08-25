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
		"cn-hangzhou": dara.String("qualitycheck.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("qualitycheck"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a business category.
//
// @param request - AddBusinessCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBusinessCategoryResponse
func (client *Client) AddBusinessCategoryWithOptions(request *AddBusinessCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddBusinessCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBusinessCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBusinessCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a business category.
//
// @param request - AddBusinessCategoryRequest
//
// @return AddBusinessCategoryResponse
func (client *Client) AddBusinessCategory(request *AddBusinessCategoryRequest) (_result *AddBusinessCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBusinessCategoryResponse{}
	_body, _err := client.AddBusinessCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a rule category.
//
// @param request - AddRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRuleCategoryResponse
func (client *Client) AddRuleCategoryWithOptions(request *AddRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRuleCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a rule category.
//
// @param request - AddRuleCategoryRequest
//
// @return AddRuleCategoryResponse
func (client *Client) AddRuleCategory(request *AddRuleCategoryRequest) (_result *AddRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRuleCategoryResponse{}
	_body, _err := client.AddRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation creates a rule on the Quality Inspection Rule Configuration page. For Apsara Stack, the URL is ip:port/api/client/UpdateRuleById.json.
//
// @param request - AddRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRuleV4Response
func (client *Client) AddRuleV4WithOptions(request *AddRuleV4Request, runtime *dara.RuntimeOptions) (_result *AddRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IsCopy) {
		body["IsCopy"] = request.IsCopy
	}

	if !dara.IsNil(request.JsonStrForRule) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation creates a rule on the Quality Inspection Rule Configuration page. For Apsara Stack, the URL is ip:port/api/client/UpdateRuleById.json.
//
// @param request - AddRuleV4Request
//
// @return AddRuleV4Response
func (client *Client) AddRuleV4(request *AddRuleV4Request) (_result *AddRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRuleV4Response{}
	_body, _err := client.AddRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs tag categorization.
//
// @param request - AnalyzeLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeLabelResponse
func (client *Client) AnalyzeLabelWithOptions(request *AnalyzeLabelRequest, runtime *dara.RuntimeOptions) (_result *AnalyzeLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeLabel"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AnalyzeLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs tag categorization.
//
// @param request - AnalyzeLabelRequest
//
// @return AnalyzeLabelResponse
func (client *Client) AnalyzeLabel(request *AnalyzeLabelRequest) (_result *AnalyzeLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AnalyzeLabelResponse{}
	_body, _err := client.AnalyzeLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Apply for the token required for real-time speech processing.
//
// @param request - ApplyWsTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyWsTokenResponse
func (client *Client) ApplyWsTokenWithOptions(request *ApplyWsTokenRequest, runtime *dara.RuntimeOptions) (_result *ApplyWsTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyWsToken"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyWsTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Apply for the token required for real-time speech processing.
//
// @param request - ApplyWsTokenRequest
//
// @return ApplyWsTokenResponse
func (client *Client) ApplyWsToken(request *ApplyWsTokenRequest) (_result *ApplyWsTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyWsTokenResponse{}
	_body, _err := client.ApplyWsTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually assign quality reviewers.
//
// Description:
//
// You can manually assign files that have completed quality inspection to reviewers. Assignments can be made one file at a time or in batches:
//
// Single-file assignment: Assign a specific file to a specified reviewer.
//
// Batch assignment: Assign multiple filtered files to one or more reviewers. You can specify how many files each reviewer receives, or let the system distribute the files evenly among reviewers.
//
// @param request - AssignReviewerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignReviewerResponse
func (client *Client) AssignReviewerWithOptions(request *AssignReviewerRequest, runtime *dara.RuntimeOptions) (_result *AssignReviewerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignReviewer"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignReviewerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually assign quality reviewers.
//
// Description:
//
// You can manually assign files that have completed quality inspection to reviewers. Assignments can be made one file at a time or in batches:
//
// Single-file assignment: Assign a specific file to a specified reviewer.
//
// Batch assignment: Assign multiple filtered files to one or more reviewers. You can specify how many files each reviewer receives, or let the system distribute the files evenly among reviewers.
//
// @param request - AssignReviewerRequest
//
// @return AssignReviewerResponse
func (client *Client) AssignReviewer(request *AssignReviewerRequest) (_result *AssignReviewerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssignReviewerResponse{}
	_body, _err := client.AssignReviewerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI AssignReviewerBySessionGroup is deprecated
//
// Summary:
//
// Frontend location: Quality Check Plan Management > Task Results > Session Groups > Batch Assign. Apsara Stack URL: ip:port/api/job/AssignReviewerBySessionGroup.json.
//
// @param request - AssignReviewerBySessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignReviewerBySessionGroupResponse
func (client *Client) AssignReviewerBySessionGroupWithOptions(request *AssignReviewerBySessionGroupRequest, runtime *dara.RuntimeOptions) (_result *AssignReviewerBySessionGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignReviewerBySessionGroup"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignReviewerBySessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AssignReviewerBySessionGroup is deprecated
//
// Summary:
//
// Frontend location: Quality Check Plan Management > Task Results > Session Groups > Batch Assign. Apsara Stack URL: ip:port/api/job/AssignReviewerBySessionGroup.json.
//
// @param request - AssignReviewerBySessionGroupRequest
//
// @return AssignReviewerBySessionGroupResponse
// Deprecated
func (client *Client) AssignReviewerBySessionGroup(request *AssignReviewerBySessionGroupRequest) (_result *AssignReviewerBySessionGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssignReviewerBySessionGroupResponse{}
	_body, _err := client.AssignReviewerBySessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation implements the Batch Review feature, which is available in the frontend under Task Management > Task Result.
//
// For private cloud deployments, use the URL: ip:port/api/qcsBatchSubmitReviewInfo.json.
//
// You can use this operation to perform a batch review on all filtered data.
//
// Note: This operation updates a large volume of data. The changes may take some time to appear.
//
// @param request - BatchSubmitReviewInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSubmitReviewInfoResponse
func (client *Client) BatchSubmitReviewInfoWithOptions(request *BatchSubmitReviewInfoRequest, runtime *dara.RuntimeOptions) (_result *BatchSubmitReviewInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSubmitReviewInfo"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSubmitReviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation implements the Batch Review feature, which is available in the frontend under Task Management > Task Result.
//
// For private cloud deployments, use the URL: ip:port/api/qcsBatchSubmitReviewInfo.json.
//
// You can use this operation to perform a batch review on all filtered data.
//
// Note: This operation updates a large volume of data. The changes may take some time to appear.
//
// @param request - BatchSubmitReviewInfoRequest
//
// @return BatchSubmitReviewInfoResponse
func (client *Client) BatchSubmitReviewInfo(request *BatchSubmitReviewInfoRequest) (_result *BatchSubmitReviewInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSubmitReviewInfoResponse{}
	_body, _err := client.BatchSubmitReviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an agent.
//
// @param request - CreateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentResponse
func (client *Client) CreateAgentWithOptions(request *CreateAgentRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgent"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an agent.
//
// @param request - CreateAgentRequest
//
// @return CreateAgentResponse
func (client *Client) CreateAgent(request *CreateAgentRequest) (_result *CreateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAgentResponse{}
	_body, _err := client.CreateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Agent batch task for conversation analysis. The application call supports HTTP calls to complete the customer response.
//
// @param request - CreateAgentTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentTaskResponse
func (client *Client) CreateAgentTaskWithOptions(request *CreateAgentTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Agent batch task for conversation analysis. The application call supports HTTP calls to complete the customer response.
//
// @param request - CreateAgentTaskRequest
//
// @return CreateAgentTaskResponse
func (client *Client) CreateAgentTask(request *CreateAgentTaskRequest) (_result *CreateAgentTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAgentTaskResponse{}
	_body, _err := client.CreateAgentTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a set of speech hotwords to the server and obtains the hotword ID in the response.
//
// Description:
//
// > Hotwords help improve recognition accuracy for specific terms, such as names, place names, or technical terms. [Learn more](https://help.aliyun.com/document_detail/213249.html).
//
// @param request - CreateAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsrVocabResponse
func (client *Client) CreateAsrVocabWithOptions(request *CreateAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *CreateAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a set of speech hotwords to the server and obtains the hotword ID in the response.
//
// Description:
//
// > Hotwords help improve recognition accuracy for specific terms, such as names, place names, or technical terms. [Learn more](https://help.aliyun.com/document_detail/213249.html).
//
// @param request - CreateAsrVocabRequest
//
// @return CreateAsrVocabResponse
func (client *Client) CreateAsrVocab(request *CreateAsrVocabRequest) (_result *CreateAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAsrVocabResponse{}
	_body, _err := client.CreateAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// UI path: Quality Check Plan Management > Add or Edit Quality Check Dimension > Add Quality Check Dimension. Apsara Stack API endpoint: ip:port/api/qcs/CreateCheckTypeToScheme.json.
//
// @param request - CreateCheckTypeToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCheckTypeToSchemeResponse
func (client *Client) CreateCheckTypeToSchemeWithOptions(request *CreateCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *CreateCheckTypeToSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCheckTypeToScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCheckTypeToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// UI path: Quality Check Plan Management > Add or Edit Quality Check Dimension > Add Quality Check Dimension. Apsara Stack API endpoint: ip:port/api/qcs/CreateCheckTypeToScheme.json.
//
// @param request - CreateCheckTypeToSchemeRequest
//
// @return CreateCheckTypeToSchemeResponse
func (client *Client) CreateCheckTypeToScheme(request *CreateCheckTypeToSchemeRequest) (_result *CreateCheckTypeToSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCheckTypeToSchemeResponse{}
	_body, _err := client.CreateCheckTypeToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a label mining task.
//
// @param request - CreateMiningTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMiningTaskResponse
func (client *Client) CreateMiningTaskWithOptions(request *CreateMiningTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMiningTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackUrl) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.FilePath) {
		body["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.Param) {
		body["Param"] = request.Param
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMiningTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMiningTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a label mining task.
//
// @param request - CreateMiningTaskRequest
//
// @return CreateMiningTaskResponse
func (client *Client) CreateMiningTask(request *CreateMiningTaskRequest) (_result *CreateMiningTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMiningTaskResponse{}
	_body, _err := client.CreateMiningTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can access this operation from the Quality Check Plan Management page in the console. The Apsara Stack endpoint is ip:port/api/qcs/CreateQualityCheckScheme.json.
//
// @param request - CreateQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityCheckSchemeResponse
func (client *Client) CreateQualityCheckSchemeWithOptions(request *CreateQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can access this operation from the Quality Check Plan Management page in the console. The Apsara Stack endpoint is ip:port/api/qcs/CreateQualityCheckScheme.json.
//
// @param request - CreateQualityCheckSchemeRequest
//
// @return CreateQualityCheckSchemeResponse
func (client *Client) CreateQualityCheckScheme(request *CreateQualityCheckSchemeRequest) (_result *CreateQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQualityCheckSchemeResponse{}
	_body, _err := client.CreateQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Corresponding frontend feature location: Plan Management > Create Quality Inspection Job. Apsara Stack URL: ip:port/api/task/CreateSchemeTaskConfig.json.
//
// @param request - CreateSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchemeTaskConfigResponse
func (client *Client) CreateSchemeTaskConfigWithOptions(request *CreateSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Corresponding frontend feature location: Plan Management > Create Quality Inspection Job. Apsara Stack URL: ip:port/api/task/CreateSchemeTaskConfig.json.
//
// @param request - CreateSchemeTaskConfigRequest
//
// @return CreateSchemeTaskConfigResponse
func (client *Client) CreateSchemeTaskConfig(request *CreateSchemeTaskConfigRequest) (_result *CreateSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSchemeTaskConfigResponse{}
	_body, _err := client.CreateSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateSkillGroupConfig is deprecated
//
// Summary:
//
// Create a configuration.
//
// @param request - CreateSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillGroupConfigResponse
func (client *Client) CreateSkillGroupConfigWithOptions(request *CreateSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateSkillGroupConfig is deprecated
//
// Summary:
//
// Create a configuration.
//
// @param request - CreateSkillGroupConfigRequest
//
// @return CreateSkillGroupConfigResponse
// Deprecated
func (client *Client) CreateSkillGroupConfig(request *CreateSkillGroupConfigRequest) (_result *CreateSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillGroupConfigResponse{}
	_body, _err := client.CreateSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a label node.
//
// @param request - CreateTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTagResponse
func (client *Client) CreateTagWithOptions(request *CreateTagRequest, runtime *dara.RuntimeOptions) (_result *CreateTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTag"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a label node.
//
// @param request - CreateTagRequest
//
// @return CreateTagResponse
func (client *Client) CreateTag(request *CreateTagRequest) (_result *CreateTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTagResponse{}
	_body, _err := client.CreateTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an automatic allocation rule for quality review tasks.
//
// @param request - CreateTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskAssignRuleResponse
func (client *Client) CreateTaskAssignRuleWithOptions(request *CreateTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateTaskAssignRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTaskAssignRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an automatic allocation rule for quality review tasks.
//
// @param request - CreateTaskAssignRuleRequest
//
// @return CreateTaskAssignRuleResponse
func (client *Client) CreateTaskAssignRule(request *CreateTaskAssignRuleRequest) (_result *CreateTaskAssignRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTaskAssignRuleResponse{}
	_body, _err := client.CreateTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a user
//
// Description:
//
// Alibaba Cloud uses Resource Access Management (RAM) for unified account management. Before you create a user in Smart Conversation Analysis, first create the user in [RAM](https://ram.console.aliyun.com). Then, obtain the user’s UID, username, and display name. Finally, add the RAM user to Smart Conversation Analysis to grant them access to the Smart Conversation Analysis service.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a user
//
// Description:
//
// Alibaba Cloud uses Resource Access Management (RAM) for unified account management. Before you create a user in Smart Conversation Analysis, first create the user in [RAM](https://ram.console.aliyun.com). Then, obtain the user’s UID, username, and display name. Finally, add the RAM user to Smart Conversation Analysis to grant them access to the Smart Conversation Analysis service.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a warning configuration.
//
// @param request - CreateWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarningConfigResponse
func (client *Client) CreateWarningConfigWithOptions(request *CreateWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateWarningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWarningConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a warning configuration.
//
// @param request - CreateWarningConfigRequest
//
// @return CreateWarningConfigResponse
func (client *Client) CreateWarningConfig(request *CreateWarningConfigRequest) (_result *CreateWarningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWarningConfigResponse{}
	_body, _err := client.CreateWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-新增
//
// @param request - CreateWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarningStrategyConfigResponse
func (client *Client) CreateWarningStrategyConfigWithOptions(request *CreateWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-新增
//
// @param request - CreateWarningStrategyConfigRequest
//
// @return CreateWarningStrategyConfigResponse
func (client *Client) CreateWarningStrategyConfig(request *CreateWarningStrategyConfigRequest) (_result *CreateWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWarningStrategyConfigResponse{}
	_body, _err := client.CreateWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a rule category.
//
// @param request - DelRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelRuleCategoryResponse
func (client *Client) DelRuleCategoryWithOptions(request *DelRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *DelRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelRuleCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a rule category.
//
// @param request - DelRuleCategoryRequest
//
// @return DelRuleCategoryResponse
func (client *Client) DelRuleCategory(request *DelRuleCategoryRequest) (_result *DelRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelRuleCategoryResponse{}
	_body, _err := client.DelRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an agent.
//
// @param request - DeleteAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithOptions(request *DeleteAgentRequest, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgent"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an agent.
//
// @param request - DeleteAgentRequest
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgent(request *DeleteAgentRequest) (_result *DeleteAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a hotword group.
//
// @param request - DeleteAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAsrVocabResponse
func (client *Client) DeleteAsrVocabWithOptions(request *DeleteAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *DeleteAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a hotword group.
//
// @param request - DeleteAsrVocabRequest
//
// @return DeleteAsrVocabResponse
func (client *Client) DeleteAsrVocab(request *DeleteAsrVocabRequest) (_result *DeleteAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAsrVocabResponse{}
	_body, _err := client.DeleteAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a business category.
//
// @param request - DeleteBusinessCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBusinessCategoryResponse
func (client *Client) DeleteBusinessCategoryWithOptions(request *DeleteBusinessCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteBusinessCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBusinessCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBusinessCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a business category.
//
// @param request - DeleteBusinessCategoryRequest
//
// @return DeleteBusinessCategoryResponse
func (client *Client) DeleteBusinessCategory(request *DeleteBusinessCategoryRequest) (_result *DeleteBusinessCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBusinessCategoryResponse{}
	_body, _err := client.DeleteBusinessCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a dimension from a quality inspection scheme.
//
// @param request - DeleteCheckTypeToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCheckTypeToSchemeResponse
func (client *Client) DeleteCheckTypeToSchemeWithOptions(request *DeleteCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteCheckTypeToSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCheckTypeToScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCheckTypeToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dimension from a quality inspection scheme.
//
// @param request - DeleteCheckTypeToSchemeRequest
//
// @return DeleteCheckTypeToSchemeResponse
func (client *Client) DeleteCheckTypeToScheme(request *DeleteCheckTypeToSchemeRequest) (_result *DeleteCheckTypeToSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCheckTypeToSchemeResponse{}
	_body, _err := client.DeleteCheckTypeToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a language model.
//
// @param request - DeleteCustomizationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizationConfigResponse
func (client *Client) DeleteCustomizationConfigWithOptions(request *DeleteCustomizationConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomizationConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomizationConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomizationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a language model.
//
// @param request - DeleteCustomizationConfigRequest
//
// @return DeleteCustomizationConfigResponse
func (client *Client) DeleteCustomizationConfig(request *DeleteCustomizationConfigRequest) (_result *DeleteCustomizationConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomizationConfigResponse{}
	_body, _err := client.DeleteCustomizationConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteDataSet is deprecated
//
// Summary:
//
// Deletes a dataset.
//
// @param request - DeleteDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSetResponse
func (client *Client) DeleteDataSetWithOptions(request *DeleteDataSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSet"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteDataSet is deprecated
//
// Summary:
//
// Deletes a dataset.
//
// @param request - DeleteDataSetRequest
//
// @return DeleteDataSetResponse
// Deprecated
func (client *Client) DeleteDataSet(request *DeleteDataSetRequest) (_result *DeleteDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.DeleteDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can delete a quality check plan from the Quality Check Plan Management page by clicking the Delete button on the right side of the plan. The Apsara Stack API endpoint is ip:port/api/qcs/DeleteQualityCheckScheme.json.
//
// @param request - DeleteQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityCheckSchemeResponse
func (client *Client) DeleteQualityCheckSchemeWithOptions(request *DeleteQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can delete a quality check plan from the Quality Check Plan Management page by clicking the Delete button on the right side of the plan. The Apsara Stack API endpoint is ip:port/api/qcs/DeleteQualityCheckScheme.json.
//
// @param request - DeleteQualityCheckSchemeRequest
//
// @return DeleteQualityCheckSchemeResponse
func (client *Client) DeleteQualityCheckScheme(request *DeleteQualityCheckSchemeRequest) (_result *DeleteQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityCheckSchemeResponse{}
	_body, _err := client.DeleteQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// Summary:
//
// This operation deletes a quality check rule. You can access it from the Quality Check Rule Configuration page in the Apsara Stack console. The API endpoint is ip:port/api/client/DeleteRule.json.
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ForceDelete) {
		body["ForceDelete"] = request.ForceDelete
	}

	if !dara.IsNil(request.IsSchemeData) {
		body["IsSchemeData"] = request.IsSchemeData
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRule"),
		Version:     dara.String("2019-01-15"),
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

// Deprecated: OpenAPI DeleteRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// Summary:
//
// This operation deletes a quality check rule. You can access it from the Quality Check Rule Configuration page in the Apsara Stack console. The API endpoint is ip:port/api/client/DeleteRule.json.
//
// @param request - DeleteRuleRequest
//
// @return DeleteRuleResponse
// Deprecated
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
// Frontend feature location: Quality Inspection Rule Configuration — Delete. Apsara Stack URL: ip:port/api/client/DeleteRule.json.
//
// @param request - DeleteRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleV4Response
func (client *Client) DeleteRuleV4WithOptions(request *DeleteRuleV4Request, runtime *dara.RuntimeOptions) (_result *DeleteRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ForceDelete) {
		body["ForceDelete"] = request.ForceDelete
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Frontend feature location: Quality Inspection Rule Configuration — Delete. Apsara Stack URL: ip:port/api/client/DeleteRule.json.
//
// @param request - DeleteRuleV4Request
//
// @return DeleteRuleV4Response
func (client *Client) DeleteRuleV4(request *DeleteRuleV4Request) (_result *DeleteRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRuleV4Response{}
	_body, _err := client.DeleteRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This feature is not available on the frontend. The Apsara Stack API endpoint is ip:port/api/task/DeleteSchemeTaskConfig.json.
//
// @param request - DeleteSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchemeTaskConfigResponse
func (client *Client) DeleteSchemeTaskConfigWithOptions(request *DeleteSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature is not available on the frontend. The Apsara Stack API endpoint is ip:port/api/task/DeleteSchemeTaskConfig.json.
//
// @param request - DeleteSchemeTaskConfigRequest
//
// @return DeleteSchemeTaskConfigResponse
func (client *Client) DeleteSchemeTaskConfig(request *DeleteSchemeTaskConfigRequest) (_result *DeleteSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSchemeTaskConfigResponse{}
	_body, _err := client.DeleteSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteSkillGroupConfig is deprecated
//
// Summary:
//
// Delete a configuration.
//
// @param request - DeleteSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillGroupConfigResponse
func (client *Client) DeleteSkillGroupConfigWithOptions(request *DeleteSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteSkillGroupConfig is deprecated
//
// Summary:
//
// Delete a configuration.
//
// @param request - DeleteSkillGroupConfigRequest
//
// @return DeleteSkillGroupConfigResponse
// Deprecated
func (client *Client) DeleteSkillGroupConfig(request *DeleteSkillGroupConfigRequest) (_result *DeleteSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSkillGroupConfigResponse{}
	_body, _err := client.DeleteSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a label node.
//
// @param request - DeleteTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTagResponse
func (client *Client) DeleteTagWithOptions(request *DeleteTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTag"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a label node.
//
// @param request - DeleteTagRequest
//
// @return DeleteTagResponse
func (client *Client) DeleteTag(request *DeleteTagRequest) (_result *DeleteTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTagResponse{}
	_body, _err := client.DeleteTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an automatic allocation rule for review tasks.
//
// @param request - DeleteTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskAssignRuleResponse
func (client *Client) DeleteTaskAssignRuleWithOptions(request *DeleteTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteTaskAssignRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTaskAssignRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an automatic allocation rule for review tasks.
//
// @param request - DeleteTaskAssignRuleRequest
//
// @return DeleteTaskAssignRuleResponse
func (client *Client) DeleteTaskAssignRule(request *DeleteTaskAssignRuleRequest) (_result *DeleteTaskAssignRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTaskAssignRuleResponse{}
	_body, _err := client.DeleteTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a warning configuration.
//
// @param request - DeleteWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarningConfigResponse
func (client *Client) DeleteWarningConfigWithOptions(request *DeleteWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteWarningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWarningConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a warning configuration.
//
// @param request - DeleteWarningConfigRequest
//
// @return DeleteWarningConfigResponse
func (client *Client) DeleteWarningConfig(request *DeleteWarningConfigRequest) (_result *DeleteWarningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWarningConfigResponse{}
	_body, _err := client.DeleteWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-删除
//
// @param request - DeleteWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarningStrategyConfigResponse
func (client *Client) DeleteWarningStrategyConfigWithOptions(request *DeleteWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-删除
//
// @param request - DeleteWarningStrategyConfigRequest
//
// @return DeleteWarningStrategyConfigResponse
func (client *Client) DeleteWarningStrategyConfig(request *DeleteWarningStrategyConfigRequest) (_result *DeleteWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWarningStrategyConfigResponse{}
	_body, _err := client.DeleteWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs an agent.
//
// @param request - ExecuteAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAgentResponse
func (client *Client) ExecuteAgentWithSSE(request *ExecuteAgentRequest, runtime *dara.RuntimeOptions, _yield chan *ExecuteAgentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.executeAgentWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// Runs an agent.
//
// @param request - ExecuteAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAgentResponse
func (client *Client) ExecuteAgentWithOptions(request *ExecuteAgentRequest, runtime *dara.RuntimeOptions) (_result *ExecuteAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAgent"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs an agent.
//
// @param request - ExecuteAgentRequest
//
// @return ExecuteAgentResponse
func (client *Client) ExecuteAgent(request *ExecuteAgentRequest) (_result *ExecuteAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteAgentResponse{}
	_body, _err := client.ExecuteAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates labels.
//
// @param request - GenerateLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateLabelResponse
func (client *Client) GenerateLabelWithOptions(request *GenerateLabelRequest, runtime *dara.RuntimeOptions) (_result *GenerateLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateLabel"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates labels.
//
// @param request - GenerateLabelRequest
//
// @return GenerateLabelResponse
func (client *Client) GenerateLabel(request *GenerateLabelRequest) (_result *GenerateLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateLabelResponse{}
	_body, _err := client.GenerateLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an agent.
//
// @param request - GetAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithOptions(request *GetAgentRequest, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgent"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an agent.
//
// @param request - GetAgentRequest
//
// @return GetAgentResponse
func (client *Client) GetAgent(request *GetAgentRequest) (_result *GetAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the task result of an agent node.
//
// Description:
//
// 可以查询通过[UploadAudioData](https://help.aliyun.com/document_detail/139399.html)、[UploadData](https://help.aliyun.com/document_detail/111394.html)上传的数据，也可以查询数据集质检任务[SubmitQualityCheckTask](https://help.aliyun.com/document_detail/158890.html)的数据。可以根据任务ID（taskId）查询，也可以根据时间范围查询。
//
// 此接⼝返回结果中默认只返回部分参数，可通过请求参数中的requiredFields来⾃定义设置返回参数中需要返回哪些字段。
//
// @param request - GetAgentTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentTaskResultResponse
func (client *Client) GetAgentTaskResultWithOptions(request *GetAgentTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetAgentTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentTaskResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the task result of an agent node.
//
// Description:
//
// 可以查询通过[UploadAudioData](https://help.aliyun.com/document_detail/139399.html)、[UploadData](https://help.aliyun.com/document_detail/111394.html)上传的数据，也可以查询数据集质检任务[SubmitQualityCheckTask](https://help.aliyun.com/document_detail/158890.html)的数据。可以根据任务ID（taskId）查询，也可以根据时间范围查询。
//
// 此接⼝返回结果中默认只返回部分参数，可通过请求参数中的requiredFields来⾃定义设置返回参数中需要返回哪些字段。
//
// @param request - GetAgentTaskResultRequest
//
// @return GetAgentTaskResultResponse
func (client *Client) GetAgentTaskResult(request *GetAgentTaskResultRequest) (_result *GetAgentTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAgentTaskResultResponse{}
	_body, _err := client.GetAgentTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves details for a specified hotword group.
//
// @param request - GetAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsrVocabResponse
func (client *Client) GetAsrVocabWithOptions(request *GetAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *GetAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves details for a specified hotword group.
//
// @param request - GetAsrVocabRequest
//
// @return GetAsrVocabResponse
func (client *Client) GetAsrVocab(request *GetAsrVocabRequest) (_result *GetAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsrVocabResponse{}
	_body, _err := client.GetAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the list of applicable businesses.
//
// @param request - GetBusinessCategoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusinessCategoryListResponse
func (client *Client) GetBusinessCategoryListWithOptions(request *GetBusinessCategoryListRequest, runtime *dara.RuntimeOptions) (_result *GetBusinessCategoryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBusinessCategoryList"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBusinessCategoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the list of applicable businesses.
//
// @param request - GetBusinessCategoryListRequest
//
// @return GetBusinessCategoryListResponse
func (client *Client) GetBusinessCategoryList(request *GetBusinessCategoryListRequest) (_result *GetBusinessCategoryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBusinessCategoryListResponse{}
	_body, _err := client.GetBusinessCategoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of language models.
//
// @param request - GetCustomizationConfigListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomizationConfigListResponse
func (client *Client) GetCustomizationConfigListWithOptions(request *GetCustomizationConfigListRequest, runtime *dara.RuntimeOptions) (_result *GetCustomizationConfigListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomizationConfigList"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomizationConfigListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of language models.
//
// @param request - GetCustomizationConfigListRequest
//
// @return GetCustomizationConfigListResponse
func (client *Client) GetCustomizationConfigList(request *GetCustomizationConfigListRequest) (_result *GetCustomizationConfigListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomizationConfigListResponse{}
	_body, _err := client.GetCustomizationConfigListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the results of tag categorization analysis.
//
// @param request - GetLabelAnalysisResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLabelAnalysisResultResponse
func (client *Client) GetLabelAnalysisResultWithOptions(request *GetLabelAnalysisResultRequest, runtime *dara.RuntimeOptions) (_result *GetLabelAnalysisResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLabelAnalysisResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLabelAnalysisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of tag categorization analysis.
//
// @param request - GetLabelAnalysisResultRequest
//
// @return GetLabelAnalysisResultResponse
func (client *Client) GetLabelAnalysisResult(request *GetLabelAnalysisResultRequest) (_result *GetLabelAnalysisResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLabelAnalysisResultResponse{}
	_body, _err := client.GetLabelAnalysisResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of a label generation task.
//
// @param request - GetLabelGeneratedResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLabelGeneratedResultResponse
func (client *Client) GetLabelGeneratedResultWithOptions(request *GetLabelGeneratedResultRequest, runtime *dara.RuntimeOptions) (_result *GetLabelGeneratedResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLabelGeneratedResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLabelGeneratedResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of a label generation task.
//
// @param request - GetLabelGeneratedResultRequest
//
// @return GetLabelGeneratedResultResponse
func (client *Client) GetLabelGeneratedResult(request *GetLabelGeneratedResultRequest) (_result *GetLabelGeneratedResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLabelGeneratedResultResponse{}
	_body, _err := client.GetLabelGeneratedResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the result of a tag mining task.
//
// @param request - GetMiningTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMiningTaskResultResponse
func (client *Client) GetMiningTaskResultWithOptions(request *GetMiningTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetMiningTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMiningTaskResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMiningTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the result of a tag mining task.
//
// @param request - GetMiningTaskResultRequest
//
// @return GetMiningTaskResultResponse
func (client *Client) GetMiningTaskResult(request *GetMiningTaskResultRequest) (_result *GetMiningTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMiningTaskResultResponse{}
	_body, _err := client.GetMiningTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the next file details for manual verification.
//
// @param request - GetNextResultToVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNextResultToVerifyResponse
func (client *Client) GetNextResultToVerifyWithOptions(request *GetNextResultToVerifyRequest, runtime *dara.RuntimeOptions) (_result *GetNextResultToVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNextResultToVerify"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNextResultToVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the next file details for manual verification.
//
// @param request - GetNextResultToVerifyRequest
//
// @return GetNextResultToVerifyResponse
func (client *Client) GetNextResultToVerify(request *GetNextResultToVerifyRequest) (_result *GetNextResultToVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNextResultToVerifyResponse{}
	_body, _err := client.GetNextResultToVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation implements the query feature in quality check plan management on the frontend. The Apsara Stack URL is ip:port/api/qcs/GetQualityCheckScheme.json.
//
// @param request - GetQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityCheckSchemeResponse
func (client *Client) GetQualityCheckSchemeWithOptions(request *GetQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *GetQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation implements the query feature in quality check plan management on the frontend. The Apsara Stack URL is ip:port/api/qcs/GetQualityCheckScheme.json.
//
// @param request - GetQualityCheckSchemeRequest
//
// @return GetQualityCheckSchemeResponse
func (client *Client) GetQualityCheckScheme(request *GetQualityCheckSchemeRequest) (_result *GetQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityCheckSchemeResponse{}
	_body, _err := client.GetQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query quality inspection results. Some fields require the requiredFields parameter to be explicitly specified in the request. Set the service endpoint (Region) to Hangzhou (cn-hangzhou).
//
// Description:
//
// You can query data uploaded using [UploadAudioData](https://help.aliyun.com/document_detail/139399.html) or [UploadData](https://help.aliyun.com/document_detail/111394.html). You can also query data from dataset-based quality inspection tasks created with [SubmitQualityCheckTask](https://help.aliyun.com/document_detail/158890.html). You can search by task ID (taskId) or by time range.
//
// @param request - GetResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResultResponse
func (client *Client) GetResultWithOptions(request *GetResultRequest, runtime *dara.RuntimeOptions) (_result *GetResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query quality inspection results. Some fields require the requiredFields parameter to be explicitly specified in the request. Set the service endpoint (Region) to Hangzhou (cn-hangzhou).
//
// Description:
//
// You can query data uploaded using [UploadAudioData](https://help.aliyun.com/document_detail/139399.html) or [UploadData](https://help.aliyun.com/document_detail/111394.html). You can also query data from dataset-based quality inspection tasks created with [SubmitQualityCheckTask](https://help.aliyun.com/document_detail/158890.html). You can search by task ID (taskId) or by time range.
//
// @param request - GetResultRequest
//
// @return GetResultResponse
func (client *Client) GetResult(request *GetResultRequest) (_result *GetResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResultResponse{}
	_body, _err := client.GetResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the quality inspection results for a specified file. The response includes the transcript, audio URL, and details of detected rule hits. You can use this information to review the file by listening to the audio, reading the transcript, and locating where rules were triggered.
//
// @param request - GetResultToReviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResultToReviewResponse
func (client *Client) GetResultToReviewWithOptions(request *GetResultToReviewRequest, runtime *dara.RuntimeOptions) (_result *GetResultToReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResultToReview"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResultToReviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the quality inspection results for a specified file. The response includes the transcript, audio URL, and details of detected rule hits. You can use this information to review the file by listening to the audio, reading the transcript, and locating where rules were triggered.
//
// @param request - GetResultToReviewRequest
//
// @return GetResultToReviewResponse
func (client *Client) GetResultToReview(request *GetResultToReviewRequest) (_result *GetResultToReviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResultToReviewResponse{}
	_body, _err := client.GetResultToReviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRule is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// Retrieves basic information about rules.
//
// Description:
//
// > This operation returns basic rule information such as the **id*	- and **name**. You can use this information with [GetRuleDetails](https://help.aliyun.com/document_detail/142310.html).
//
// @param request - GetRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleResponse
func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRule is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// Retrieves basic information about rules.
//
// Description:
//
// > This operation returns basic rule information such as the **id*	- and **name**. You can use this information with [GetRuleDetails](https://help.aliyun.com/document_detail/142310.html).
//
// @param request - GetRuleRequest
//
// @return GetRuleResponse
// Deprecated
func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRuleById is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// Retrieves the details of a specific quality inspection rule. It corresponds to the **Edit*	- action on the **Quality Inspection Rule Configuration*	- page. The URL for this operation in a private cloud is `ip:port/api/client/GetRuleById.json`.
//
// @param request - GetRuleByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleByIdResponse
func (client *Client) GetRuleByIdWithOptions(request *GetRuleByIdRequest, runtime *dara.RuntimeOptions) (_result *GetRuleByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleById"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRuleById is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// Retrieves the details of a specific quality inspection rule. It corresponds to the **Edit*	- action on the **Quality Inspection Rule Configuration*	- page. The URL for this operation in a private cloud is `ip:port/api/client/GetRuleById.json`.
//
// @param request - GetRuleByIdRequest
//
// @return GetRuleByIdResponse
// Deprecated
func (client *Client) GetRuleById(request *GetRuleByIdRequest) (_result *GetRuleByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleByIdResponse{}
	_body, _err := client.GetRuleByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of rule types.
//
// @param request - GetRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleCategoryResponse
func (client *Client) GetRuleCategoryWithOptions(request *GetRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *GetRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleCategory"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of rule types.
//
// @param request - GetRuleCategoryRequest
//
// @return GetRuleCategoryResponse
func (client *Client) GetRuleCategory(request *GetRuleCategoryRequest) (_result *GetRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleCategoryResponse{}
	_body, _err := client.GetRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetRuleDetail is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// Retrieves the detailed information of a rule.
//
// Description:
//
// > This operation is used in conjunction with [Get basic rule information](https://help.aliyun.com/document_detail/142333.html). First, call the GetRule operation to obtain the rule ID. Then, use the rule ID as a parameter to call the **GetRuleDetail*	- operation.
//
// @param request - GetRuleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleDetailResponse
func (client *Client) GetRuleDetailWithOptions(request *GetRuleDetailRequest, runtime *dara.RuntimeOptions) (_result *GetRuleDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleDetail"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetRuleDetail is deprecated, please use Qualitycheck::2019-01-15::GetRuleV4 instead.
//
// Summary:
//
// Retrieves the detailed information of a rule.
//
// Description:
//
// > This operation is used in conjunction with [Get basic rule information](https://help.aliyun.com/document_detail/142333.html). First, call the GetRule operation to obtain the rule ID. Then, use the rule ID as a parameter to call the **GetRuleDetail*	- operation.
//
// @param request - GetRuleDetailRequest
//
// @return GetRuleDetailResponse
// Deprecated
func (client *Client) GetRuleDetail(request *GetRuleDetailRequest) (_result *GetRuleDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleDetailResponse{}
	_body, _err := client.GetRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API is located in the frontend at Quality Check Rule Configuration > Query. The Apsara Stack URL is ip:port/api/client/GetRuleById.json.
//
// @param request - GetRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleV4Response
func (client *Client) GetRuleV4WithOptions(request *GetRuleV4Request, runtime *dara.RuntimeOptions) (_result *GetRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is located in the frontend at Quality Check Rule Configuration > Query. The Apsara Stack URL is ip:port/api/client/GetRuleById.json.
//
// @param request - GetRuleV4Request
//
// @return GetRuleV4Response
func (client *Client) GetRuleV4(request *GetRuleV4Request) (_result *GetRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleV4Response{}
	_body, _err := client.GetRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Corresponds to the frontend feature location: Quality Inspection Rule Configuration > List. Apsara Stack URL: ip:port/api/rule/GetRulesCountList.json.
//
// @param request - GetRulesCountListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRulesCountListResponse
func (client *Client) GetRulesCountListWithOptions(request *GetRulesCountListRequest, runtime *dara.RuntimeOptions) (_result *GetRulesCountListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.BusinessRange) {
		body["BusinessRange"] = request.BusinessRange
	}

	if !dara.IsNil(request.CategoryName) {
		body["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.CountTotal) {
		body["CountTotal"] = request.CountTotal
	}

	if !dara.IsNil(request.CreateEmpid) {
		body["CreateEmpid"] = request.CreateEmpid
	}

	if !dara.IsNil(request.CreateUserId) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LastUpdateEmpid) {
		body["LastUpdateEmpid"] = request.LastUpdateEmpid
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequireInfos) {
		body["RequireInfos"] = request.RequireInfos
	}

	if !dara.IsNil(request.Rid) {
		body["Rid"] = request.Rid
	}

	if !dara.IsNil(request.RuleIdOrRuleName) {
		body["RuleIdOrRuleName"] = request.RuleIdOrRuleName
	}

	if !dara.IsNil(request.RuleScoreSingleType) {
		body["RuleScoreSingleType"] = request.RuleScoreSingleType
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SchemeId) {
		body["SchemeId"] = request.SchemeId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.TypeName) {
		body["TypeName"] = request.TypeName
	}

	if !dara.IsNil(request.UpdateEndTime) {
		body["UpdateEndTime"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		body["UpdateStartTime"] = request.UpdateStartTime
	}

	if !dara.IsNil(request.UpdateUserId) {
		body["UpdateUserId"] = request.UpdateUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRulesCountList"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRulesCountListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Corresponds to the frontend feature location: Quality Inspection Rule Configuration > List. Apsara Stack URL: ip:port/api/rule/GetRulesCountList.json.
//
// @param request - GetRulesCountListRequest
//
// @return GetRulesCountListResponse
func (client *Client) GetRulesCountList(request *GetRulesCountListRequest) (_result *GetRulesCountListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRulesCountListResponse{}
	_body, _err := client.GetRulesCountListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the configuration details of a quality inspection task.
//
// @param request - GetSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemeTaskConfigResponse
func (client *Client) GetSchemeTaskConfigWithOptions(request *GetSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the configuration details of a quality inspection task.
//
// @param request - GetSchemeTaskConfigRequest
//
// @return GetSchemeTaskConfigResponse
func (client *Client) GetSchemeTaskConfig(request *GetSchemeTaskConfigRequest) (_result *GetSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSchemeTaskConfigResponse{}
	_body, _err := client.GetSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetSkillGroupConfig is deprecated
//
// Summary:
//
// Retrieves the configuration that is specified by its ID.
//
// @param request - GetSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupConfigResponse
func (client *Client) GetSkillGroupConfigWithOptions(request *GetSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetSkillGroupConfig is deprecated
//
// Summary:
//
// Retrieves the configuration that is specified by its ID.
//
// @param request - GetSkillGroupConfigRequest
//
// @return GetSkillGroupConfigResponse
// Deprecated
func (client *Client) GetSkillGroupConfig(request *GetSkillGroupConfigRequest) (_result *GetSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillGroupConfigResponse{}
	_body, _err := client.GetSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetSyncResult is deprecated, please use Qualitycheck::2019-01-15::GetResult instead.
//
// Summary:
//
// You can obtain the real-time quality inspection result of the hotline.
//
// @param request - GetSyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSyncResultResponse
func (client *Client) GetSyncResultWithOptions(request *GetSyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetSyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSyncResult"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetSyncResult is deprecated, please use Qualitycheck::2019-01-15::GetResult instead.
//
// Summary:
//
// You can obtain the real-time quality inspection result of the hotline.
//
// @param request - GetSyncResultRequest
//
// @return GetSyncResultResponse
// Deprecated
func (client *Client) GetSyncResult(request *GetSyncResultRequest) (_result *GetSyncResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSyncResultResponse{}
	_body, _err := client.GetSyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a label node.
//
// @param request - GetTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTagResponse
func (client *Client) GetTagWithOptions(request *GetTagRequest, runtime *dara.RuntimeOptions) (_result *GetTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTag"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a label node.
//
// @param request - GetTagRequest
//
// @return GetTagResponse
func (client *Client) GetTag(request *GetTagRequest) (_result *GetTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTagResponse{}
	_body, _err := client.GetTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-详情
//
// @param request - GetWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWarningStrategyConfigResponse
func (client *Client) GetWarningStrategyConfigWithOptions(request *GetWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *GetWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-详情
//
// @param request - GetWarningStrategyConfigRequest
//
// @return GetWarningStrategyConfigResponse
func (client *Client) GetWarningStrategyConfig(request *GetWarningStrategyConfigRequest) (_result *GetWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWarningStrategyConfigResponse{}
	_body, _err := client.GetWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Handles a complaint.
//
// Description:
//
// Only quality checkers or administrators can call this operation.
//
// @param request - HandleComplaintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandleComplaintResponse
func (client *Client) HandleComplaintWithOptions(request *HandleComplaintRequest, runtime *dara.RuntimeOptions) (_result *HandleComplaintResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HandleComplaint"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HandleComplaintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Handles a complaint.
//
// Description:
//
// Only quality checkers or administrators can call this operation.
//
// @param request - HandleComplaintRequest
//
// @return HandleComplaintResponse
func (client *Client) HandleComplaint(request *HandleComplaintRequest) (_result *HandleComplaintResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &HandleComplaintResponse{}
	_body, _err := client.HandleComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI InvalidRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// Summary:
//
// Delete rules.
//
// @param request - InvalidRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidRuleResponse
func (client *Client) InvalidRuleWithOptions(request *InvalidRuleRequest, runtime *dara.RuntimeOptions) (_result *InvalidRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvalidRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvalidRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI InvalidRule is deprecated, please use Qualitycheck::2019-01-15::DeleteRuleV4 instead.
//
// Summary:
//
// Delete rules.
//
// @param request - InvalidRuleRequest
//
// @return InvalidRuleResponse
// Deprecated
func (client *Client) InvalidRule(request *InvalidRuleRequest) (_result *InvalidRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InvalidRuleResponse{}
	_body, _err := client.InvalidRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of vocabulary groups without their specific content.
//
// @param request - ListAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsrVocabResponse
func (client *Client) ListAsrVocabWithOptions(request *ListAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *ListAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of vocabulary groups without their specific content.
//
// @param request - ListAsrVocabRequest
//
// @return ListAsrVocabResponse
func (client *Client) ListAsrVocab(request *ListAsrVocabRequest) (_result *ListAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAsrVocabResponse{}
	_body, _err := client.ListAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListDataSet is deprecated
//
// Summary:
//
// This feature is located in the Dataset management section of the frontend. The Apsara Stack URL is ip:port/api/dataset/ListDataSet.json.
//
// @param request - ListDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSetResponse
func (client *Client) ListDataSetWithOptions(request *ListDataSetRequest, runtime *dara.RuntimeOptions) (_result *ListDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSet"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListDataSet is deprecated
//
// Summary:
//
// This feature is located in the Dataset management section of the frontend. The Apsara Stack URL is ip:port/api/dataset/ListDataSet.json.
//
// @param request - ListDataSetRequest
//
// @return ListDataSetResponse
// Deprecated
func (client *Client) ListDataSet(request *ListDataSetRequest) (_result *ListDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSetResponse{}
	_body, _err := client.ListDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of quality check schemes.
//
// @param request - ListQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityCheckSchemeResponse
func (client *Client) ListQualityCheckSchemeWithOptions(request *ListQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *ListQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of quality check schemes.
//
// @param request - ListQualityCheckSchemeRequest
//
// @return ListQualityCheckSchemeResponse
func (client *Client) ListQualityCheckScheme(request *ListQualityCheckSchemeRequest) (_result *ListQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityCheckSchemeResponse{}
	_body, _err := client.ListQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListRules is deprecated, please use Qualitycheck::2019-01-15::ListRulesV4 instead.
//
// Summary:
//
// Lists rules.
//
// @param request - ListRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRules"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListRules is deprecated, please use Qualitycheck::2019-01-15::ListRulesV4 instead.
//
// Summary:
//
// Lists rules.
//
// @param request - ListRulesRequest
//
// @return ListRulesResponse
// Deprecated
func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Frontend location: Quality inspection rule configuration — List. Apsara Stack URL: ip:port/api/rule/GetRulesCountList.json.
//
// @param request - ListRulesV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesV4Response
func (client *Client) ListRulesV4WithOptions(request *ListRulesV4Request, runtime *dara.RuntimeOptions) (_result *ListRulesV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.BusinessRange) {
		body["BusinessRange"] = request.BusinessRange
	}

	if !dara.IsNil(request.CategoryName) {
		body["CategoryName"] = request.CategoryName
	}

	if !dara.IsNil(request.CountTotal) {
		body["CountTotal"] = request.CountTotal
	}

	if !dara.IsNil(request.CreateEmpid) {
		body["CreateEmpid"] = request.CreateEmpid
	}

	if !dara.IsNil(request.CreateUserId) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LastUpdateEmpid) {
		body["LastUpdateEmpid"] = request.LastUpdateEmpid
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequireInfos) {
		body["RequireInfos"] = request.RequireInfos
	}

	if !dara.IsNil(request.Rid) {
		body["Rid"] = request.Rid
	}

	if !dara.IsNil(request.RuleIdOrRuleName) {
		body["RuleIdOrRuleName"] = request.RuleIdOrRuleName
	}

	if !dara.IsNil(request.RuleScoreSingleType) {
		body["RuleScoreSingleType"] = request.RuleScoreSingleType
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SchemeId) {
		body["SchemeId"] = request.SchemeId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TargetType) {
		body["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.TypeName) {
		body["TypeName"] = request.TypeName
	}

	if !dara.IsNil(request.UpdateEndTime) {
		body["UpdateEndTime"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		body["UpdateStartTime"] = request.UpdateStartTime
	}

	if !dara.IsNil(request.UpdateUserId) {
		body["UpdateUserId"] = request.UpdateUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRulesV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRulesV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Frontend location: Quality inspection rule configuration — List. Apsara Stack URL: ip:port/api/rule/GetRulesCountList.json.
//
// @param request - ListRulesV4Request
//
// @return ListRulesV4Response
func (client *Client) ListRulesV4(request *ListRulesV4Request) (_result *ListRulesV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRulesV4Response{}
	_body, _err := client.ListRulesV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation is used for the Task Management feature on the frontend. The Apsara Stack URL is ip:port/api/task/ListSchemeTaskInfo.json.
//
// @param request - ListSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemeTaskConfigResponse
func (client *Client) ListSchemeTaskConfigWithOptions(request *ListSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *ListSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation is used for the Task Management feature on the frontend. The Apsara Stack URL is ip:port/api/task/ListSchemeTaskInfo.json.
//
// @param request - ListSchemeTaskConfigRequest
//
// @return ListSchemeTaskConfigResponse
func (client *Client) ListSchemeTaskConfig(request *ListSchemeTaskConfigRequest) (_result *ListSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSchemeTaskConfigResponse{}
	_body, _err := client.ListSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListSessionGroup is deprecated
//
// Summary:
//
// This API corresponds to the frontend location: Task Management > View Results > Task Result > Session Group Results tab. The Apsara Stack URL is: ip:port/api/session/group/ListSessionGroup.json. It aggregates multi-turn sessions by their session group ID for unified management. You must pass the `sessionGroupId` field. For more information, see the UploadData and UploadAudioData API documentation.
//
// @param request - ListSessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionGroupResponse
func (client *Client) ListSessionGroupWithOptions(request *ListSessionGroupRequest, runtime *dara.RuntimeOptions) (_result *ListSessionGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessionGroup"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListSessionGroup is deprecated
//
// Summary:
//
// This API corresponds to the frontend location: Task Management > View Results > Task Result > Session Group Results tab. The Apsara Stack URL is: ip:port/api/session/group/ListSessionGroup.json. It aggregates multi-turn sessions by their session group ID for unified management. You must pass the `sessionGroupId` field. For more information, see the UploadData and UploadAudioData API documentation.
//
// @param request - ListSessionGroupRequest
//
// @return ListSessionGroupResponse
// Deprecated
func (client *Client) ListSessionGroup(request *ListSessionGroupRequest) (_result *ListSessionGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSessionGroupResponse{}
	_body, _err := client.ListSessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListSkillGroupConfig is deprecated
//
// Summary:
//
// You can call ListSkillGroupConfig to obtain the configuration list.
//
// @param request - ListSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillGroupConfigResponse
func (client *Client) ListSkillGroupConfigWithOptions(request *ListSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *ListSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListSkillGroupConfig is deprecated
//
// Summary:
//
// You can call ListSkillGroupConfig to obtain the configuration list.
//
// @param request - ListSkillGroupConfigRequest
//
// @return ListSkillGroupConfigResponse
// Deprecated
func (client *Client) ListSkillGroupConfig(request *ListSkillGroupConfigRequest) (_result *ListSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillGroupConfigResponse{}
	_body, _err := client.ListSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of label nodes.
//
// @param request - ListTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResponse
func (client *Client) ListTagWithOptions(request *ListTagRequest, runtime *dara.RuntimeOptions) (_result *ListTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTag"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of label nodes.
//
// @param request - ListTagRequest
//
// @return ListTagResponse
func (client *Client) ListTag(request *ListTagRequest) (_result *ListTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResponse{}
	_body, _err := client.ListTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists automatic allocation rules for review tasks.
//
// @param request - ListTaskAssignRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskAssignRulesResponse
func (client *Client) ListTaskAssignRulesWithOptions(request *ListTaskAssignRulesRequest, runtime *dara.RuntimeOptions) (_result *ListTaskAssignRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskAssignRules"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskAssignRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists automatic allocation rules for review tasks.
//
// @param request - ListTaskAssignRulesRequest
//
// @return ListTaskAssignRulesResponse
func (client *Client) ListTaskAssignRules(request *ListTaskAssignRulesRequest) (_result *ListTaskAssignRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskAssignRulesResponse{}
	_body, _err := client.ListTaskAssignRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of users. Set the service endpoint to China (Hangzhou) (cn-hangzhou).
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of users. Set the service endpoint to China (Hangzhou) (cn-hangzhou).
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists warning configurations.
//
// @param request - ListWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarningConfigResponse
func (client *Client) ListWarningConfigWithOptions(request *ListWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *ListWarningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWarningConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists warning configurations.
//
// @param request - ListWarningConfigRequest
//
// @return ListWarningConfigResponse
func (client *Client) ListWarningConfig(request *ListWarningConfigRequest) (_result *ListWarningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWarningConfigResponse{}
	_body, _err := client.ListWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-列表
//
// @param request - ListWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarningStrategyConfigResponse
func (client *Client) ListWarningStrategyConfigWithOptions(request *ListWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *ListWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-列表
//
// @param request - ListWarningStrategyConfigRequest
//
// @return ListWarningStrategyConfigResponse
func (client *Client) ListWarningStrategyConfig(request *ListWarningStrategyConfigRequest) (_result *ListWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWarningStrategyConfigResponse{}
	_body, _err := client.ListWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Frontend location: Task Management > Task Results > Batch Revoke. Apsara Stack URL: ip:port/api/job/RevertAssignedSession.json.
//
// @param request - RevertAssignedSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAssignedSessionResponse
func (client *Client) RevertAssignedSessionWithOptions(request *RevertAssignedSessionRequest, runtime *dara.RuntimeOptions) (_result *RevertAssignedSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevertAssignedSession"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevertAssignedSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Frontend location: Task Management > Task Results > Batch Revoke. Apsara Stack URL: ip:port/api/job/RevertAssignedSession.json.
//
// @param request - RevertAssignedSessionRequest
//
// @return RevertAssignedSessionResponse
func (client *Client) RevertAssignedSession(request *RevertAssignedSessionRequest) (_result *RevertAssignedSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevertAssignedSessionResponse{}
	_body, _err := client.RevertAssignedSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI RevertAssignedSessionGroup is deprecated
//
// Summary:
//
// You can use this feature in the frontend console under Plan Management > Task Result > Session Group > Batch Revoke. The Apsara Stack URL for this operation is `ip:port/api/job/RevertAssignedSessionGroup.json`.
//
// @param request - RevertAssignedSessionGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertAssignedSessionGroupResponse
func (client *Client) RevertAssignedSessionGroupWithOptions(request *RevertAssignedSessionGroupRequest, runtime *dara.RuntimeOptions) (_result *RevertAssignedSessionGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevertAssignedSessionGroup"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevertAssignedSessionGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RevertAssignedSessionGroup is deprecated
//
// Summary:
//
// You can use this feature in the frontend console under Plan Management > Task Result > Session Group > Batch Revoke. The Apsara Stack URL for this operation is `ip:port/api/job/RevertAssignedSessionGroup.json`.
//
// @param request - RevertAssignedSessionGroupRequest
//
// @return RevertAssignedSessionGroupResponse
// Deprecated
func (client *Client) RevertAssignedSessionGroup(request *RevertAssignedSessionGroupRequest) (_result *RevertAssignedSessionGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevertAssignedSessionGroupResponse{}
	_body, _err := client.RevertAssignedSessionGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation calls a large model using the message protocol to generate a response. You can make calls using standard HTTP for a complete response or use Server-Sent Events (SSE) for a streaming response.
//
// @param tmpReq - RunCompletionMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCompletionMessageResponse
func (client *Client) RunCompletionMessageWithSSE(tmpReq *RunCompletionMessageRequest, runtime *dara.RuntimeOptions, _yield chan *RunCompletionMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCompletionMessageWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
	return
}

// Summary:
//
// This operation calls a large model using the message protocol to generate a response. You can make calls using standard HTTP for a complete response or use Server-Sent Events (SSE) for a streaming response.
//
// @param tmpReq - RunCompletionMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCompletionMessageResponse
func (client *Client) RunCompletionMessageWithOptions(tmpReq *RunCompletionMessageRequest, runtime *dara.RuntimeOptions) (_result *RunCompletionMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunCompletionMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Messages) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, dara.String("Messages"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MessagesShrink) {
		body["Messages"] = request.MessagesShrink
	}

	if !dara.IsNil(request.ModelCode) {
		body["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCompletionMessage"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCompletionMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation calls a large model using the message protocol to generate a response. You can make calls using standard HTTP for a complete response or use Server-Sent Events (SSE) for a streaming response.
//
// @param request - RunCompletionMessageRequest
//
// @return RunCompletionMessageResponse
func (client *Client) RunCompletionMessage(request *RunCompletionMessageRequest) (_result *RunCompletionMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunCompletionMessageResponse{}
	_body, _err := client.RunCompletionMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI SaveConfigDataSet is deprecated
//
// Summary:
//
// Saves the speaker role configuration for a dataset.
//
// @param request - SaveConfigDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveConfigDataSetResponse
func (client *Client) SaveConfigDataSetWithOptions(request *SaveConfigDataSetRequest, runtime *dara.RuntimeOptions) (_result *SaveConfigDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveConfigDataSet"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveConfigDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI SaveConfigDataSet is deprecated
//
// Summary:
//
// Saves the speaker role configuration for a dataset.
//
// @param request - SaveConfigDataSetRequest
//
// @return SaveConfigDataSetResponse
// Deprecated
func (client *Client) SaveConfigDataSet(request *SaveConfigDataSetRequest) (_result *SaveConfigDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveConfigDataSetResponse{}
	_body, _err := client.SaveConfigDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a complaint.
//
// @param request - SubmitComplaintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitComplaintResponse
func (client *Client) SubmitComplaintWithOptions(request *SubmitComplaintRequest, runtime *dara.RuntimeOptions) (_result *SubmitComplaintResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitComplaint"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitComplaintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a complaint.
//
// @param request - SubmitComplaintRequest
//
// @return SubmitComplaintResponse
func (client *Client) SubmitComplaint(request *SubmitComplaintRequest) (_result *SubmitComplaintResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitComplaintResponse{}
	_body, _err := client.SubmitComplaintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a new dataset quality check task.
//
// @param request - SubmitQualityCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitQualityCheckTaskResponse
func (client *Client) SubmitQualityCheckTaskWithOptions(request *SubmitQualityCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitQualityCheckTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitQualityCheckTask"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitQualityCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a new dataset quality check task.
//
// @param request - SubmitQualityCheckTaskRequest
//
// @return SubmitQualityCheckTaskResponse
func (client *Client) SubmitQualityCheckTask(request *SubmitQualityCheckTaskRequest) (_result *SubmitQualityCheckTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitQualityCheckTaskResponse{}
	_body, _err := client.SubmitQualityCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Save review results. This is only supported by the legacy Smart Conversation Analysis.
//
// Description:
//
// You can manually review files after quality inspection. After completing the review, call this API to save the review results. This involves manually reviewing rules identified by the system as hits to determine if they are true hits or false positives. Refer to the file review feature on the console page. For more information, see [File Review](https://help.aliyun.com/document_detail/139653.html#h2-u6587u4EF6u590Du68385).
//
// @param request - SubmitReviewInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitReviewInfoResponse
func (client *Client) SubmitReviewInfoWithOptions(request *SubmitReviewInfoRequest, runtime *dara.RuntimeOptions) (_result *SubmitReviewInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitReviewInfo"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitReviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Save review results. This is only supported by the legacy Smart Conversation Analysis.
//
// Description:
//
// You can manually review files after quality inspection. After completing the review, call this API to save the review results. This involves manually reviewing rules identified by the system as hits to determine if they are true hits or false positives. Refer to the file review feature on the console page. For more information, see [File Review](https://help.aliyun.com/document_detail/139653.html#h2-u6587u4EF6u590Du68385).
//
// @param request - SubmitReviewInfoRequest
//
// @return SubmitReviewInfoResponse
func (client *Client) SubmitReviewInfo(request *SubmitReviewInfoRequest) (_result *SubmitReviewInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitReviewInfoResponse{}
	_body, _err := client.SubmitReviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Review quality inspection results
//
// @param request - SubmitReviewInfoV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitReviewInfoV4Response
func (client *Client) SubmitReviewInfoV4WithOptions(request *SubmitReviewInfoV4Request, runtime *dara.RuntimeOptions) (_result *SubmitReviewInfoV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitReviewInfoV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitReviewInfoV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Review quality inspection results
//
// @param request - SubmitReviewInfoV4Request
//
// @return SubmitReviewInfoV4Response
func (client *Client) SubmitReviewInfoV4(request *SubmitReviewInfoV4Request) (_result *SubmitReviewInfoV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitReviewInfoV4Response{}
	_body, _err := client.SubmitReviewInfoV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Perform real-time quality inspection for hotlines.
//
// Description:
//
// Real-time hotline quality inspection transcribes spoken dialogue into text during a call. It sends the text to the Smart Conversation Analysis system for real-time quality inspection to detect potential issues or risks. You can display the dialogue text and inspection results in real time on the customer service representative\\"s workbench (a third-party system). This differs from offline quality inspection, which uses UploadAudioData for inspection or dataset inspection. For details, see the developer guide. Offline quality inspection occurs after the call ends and the recording file is generated.
//
// **Usage Flow**
//
// You can implement real-time transcription of audio streams to text during calls, or use Alibaba Cloud Call Center (CC) directly. CC integrates deeply with Smart Conversation Analysis, enabling real-time quality inspection during calls without API integration.
//
// If you implement audio-to-text conversion yourself, invoke the SyncQualityCheck API for real-time quality inspection after a speaker finishes a sentence and generates dialogue text. This returns the inspection result for that sentence synchronously.
//
// You should include skill group information when uploading data. Then, you can use the Call Center Quality Inspection - Configuration Management feature to configure different quality inspection rules for calls from different skill groups.
//
// After the call ends, you can store the recording file on a storage server accessible over the public network. You can invoke the recording information maintenance API: UpdateSyncQualityCheckData. You can submit the recording name, recording file URL, and other details to the Smart Conversation Analysis service. This lets quality inspectors play back the recording during review.
//
// After the call ends, you can view the quality inspection results in Call Center Quality Inspection - Result Display - Real-time Quality Inspection Results. You can also invoke the real-time quality inspection result query API: GetSyncResult to retrieve the results. You can use Score Dashboard - Real-time Dashboard to view data charts for customer service representatives, skill groups, and scoring items.
//
// **Full-Text Quality Inspection**
//
// Quality inspection rules include dozens of operators. Some operators require dialogue context (multi-turn conversations between customer service representatives and customers) for analysis. However, real-time quality inspection occurs during a call and typically uses text from only one sentence spoken by a single speaker. Some operators are not suitable for real-time quality inspection. Therefore, quality inspection rules are divided into real-time quality inspection rules and full-text quality inspection rules:
//
// **Real-time quality inspection rules**: Rules used for real-time quality inspection. They support a limited number of operator types. They do not support specifying the detection range for operators.
//
// **Full-text quality inspection rules**: Rules used for offline quality inspection. They support all operator types. They support custom detection ranges for operators.
//
// For calls that underwent real-time quality inspection, you can apply full-text quality inspection rules to the complete dialogue text after the call ends. To enable full-text quality inspection after real-time inspection, see the full-text quality inspection description in Call Center Quality Inspection - Configuration Management.
//
// @param request - SyncQualityCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncQualityCheckResponse
func (client *Client) SyncQualityCheckWithOptions(request *SyncQualityCheckRequest, runtime *dara.RuntimeOptions) (_result *SyncQualityCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncQualityCheck"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncQualityCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Perform real-time quality inspection for hotlines.
//
// Description:
//
// Real-time hotline quality inspection transcribes spoken dialogue into text during a call. It sends the text to the Smart Conversation Analysis system for real-time quality inspection to detect potential issues or risks. You can display the dialogue text and inspection results in real time on the customer service representative\\"s workbench (a third-party system). This differs from offline quality inspection, which uses UploadAudioData for inspection or dataset inspection. For details, see the developer guide. Offline quality inspection occurs after the call ends and the recording file is generated.
//
// **Usage Flow**
//
// You can implement real-time transcription of audio streams to text during calls, or use Alibaba Cloud Call Center (CC) directly. CC integrates deeply with Smart Conversation Analysis, enabling real-time quality inspection during calls without API integration.
//
// If you implement audio-to-text conversion yourself, invoke the SyncQualityCheck API for real-time quality inspection after a speaker finishes a sentence and generates dialogue text. This returns the inspection result for that sentence synchronously.
//
// You should include skill group information when uploading data. Then, you can use the Call Center Quality Inspection - Configuration Management feature to configure different quality inspection rules for calls from different skill groups.
//
// After the call ends, you can store the recording file on a storage server accessible over the public network. You can invoke the recording information maintenance API: UpdateSyncQualityCheckData. You can submit the recording name, recording file URL, and other details to the Smart Conversation Analysis service. This lets quality inspectors play back the recording during review.
//
// After the call ends, you can view the quality inspection results in Call Center Quality Inspection - Result Display - Real-time Quality Inspection Results. You can also invoke the real-time quality inspection result query API: GetSyncResult to retrieve the results. You can use Score Dashboard - Real-time Dashboard to view data charts for customer service representatives, skill groups, and scoring items.
//
// **Full-Text Quality Inspection**
//
// Quality inspection rules include dozens of operators. Some operators require dialogue context (multi-turn conversations between customer service representatives and customers) for analysis. However, real-time quality inspection occurs during a call and typically uses text from only one sentence spoken by a single speaker. Some operators are not suitable for real-time quality inspection. Therefore, quality inspection rules are divided into real-time quality inspection rules and full-text quality inspection rules:
//
// **Real-time quality inspection rules**: Rules used for real-time quality inspection. They support a limited number of operator types. They do not support specifying the detection range for operators.
//
// **Full-text quality inspection rules**: Rules used for offline quality inspection. They support all operator types. They support custom detection ranges for operators.
//
// For calls that underwent real-time quality inspection, you can apply full-text quality inspection rules to the complete dialogue text after the call ends. To enable full-text quality inspection after real-time inspection, see the full-text quality inspection description in Call Center Quality Inspection - Configuration Management.
//
// @param request - SyncQualityCheckRequest
//
// @return SyncQualityCheckResponse
func (client *Client) SyncQualityCheck(request *SyncQualityCheckRequest) (_result *SyncQualityCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncQualityCheckResponse{}
	_body, _err := client.SyncQualityCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Frontend location: Quality Check Rule Configuration > Test. Apsara Stack URL: http://<ip>:<port>/api/client/TestRule.json.
//
// @param request - TestRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestRuleV4Response
func (client *Client) TestRuleV4WithOptions(request *TestRuleV4Request, runtime *dara.RuntimeOptions) (_result *TestRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsSchemeData) {
		body["IsSchemeData"] = request.IsSchemeData
	}

	if !dara.IsNil(request.TestJson) {
		body["TestJson"] = request.TestJson
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Frontend location: Quality Check Rule Configuration > Test. Apsara Stack URL: http://<ip>:<port>/api/client/TestRule.json.
//
// @param request - TestRuleV4Request
//
// @return TestRuleV4Response
func (client *Client) TestRuleV4(request *TestRuleV4Request) (_result *TestRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestRuleV4Response{}
	_body, _err := client.TestRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an agent.
//
// @param request - UpdateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgentWithOptions(request *UpdateAgentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgent"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an agent.
//
// @param request - UpdateAgentRequest
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgent(request *UpdateAgentRequest) (_result *UpdateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAgentResponse{}
	_body, _err := client.UpdateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the hotword vocabulary.
//
// @param request - UpdateAsrVocabRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAsrVocabResponse
func (client *Client) UpdateAsrVocabWithOptions(request *UpdateAsrVocabRequest, runtime *dara.RuntimeOptions) (_result *UpdateAsrVocabResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAsrVocab"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAsrVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the hotword vocabulary.
//
// @param request - UpdateAsrVocabRequest
//
// @return UpdateAsrVocabResponse
func (client *Client) UpdateAsrVocab(request *UpdateAsrVocabRequest) (_result *UpdateAsrVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAsrVocabResponse{}
	_body, _err := client.UpdateAsrVocabWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can access this operation from the frontend by navigating to Plan Management > Create Quality Inspection Task or Edit > Edit icon next to the quality inspection dimension name. The Apsara Stack endpoint is ip:port/api/qcs/UpdateCheckTypeToScheme.json.
//
// @param request - UpdateCheckTypeToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCheckTypeToSchemeResponse
func (client *Client) UpdateCheckTypeToSchemeWithOptions(request *UpdateCheckTypeToSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCheckTypeToSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCheckTypeToScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCheckTypeToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can access this operation from the frontend by navigating to Plan Management > Create Quality Inspection Task or Edit > Edit icon next to the quality inspection dimension name. The Apsara Stack endpoint is ip:port/api/qcs/UpdateCheckTypeToScheme.json.
//
// @param request - UpdateCheckTypeToSchemeRequest
//
// @return UpdateCheckTypeToSchemeResponse
func (client *Client) UpdateCheckTypeToScheme(request *UpdateCheckTypeToSchemeRequest) (_result *UpdateCheckTypeToSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCheckTypeToSchemeResponse{}
	_body, _err := client.UpdateCheckTypeToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update session recording data (third-party business fields) to facilitate statistics and queries across more business dimensions.
//
// @param request - UpdateQualityCheckDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityCheckDataResponse
func (client *Client) UpdateQualityCheckDataWithOptions(request *UpdateQualityCheckDataRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityCheckDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQualityCheckData"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update session recording data (third-party business fields) to facilitate statistics and queries across more business dimensions.
//
// @param request - UpdateQualityCheckDataRequest
//
// @return UpdateQualityCheckDataResponse
func (client *Client) UpdateQualityCheckData(request *UpdateQualityCheckDataRequest) (_result *UpdateQualityCheckDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQualityCheckDataResponse{}
	_body, _err := client.UpdateQualityCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a quality check scheme.
//
// @param request - UpdateQualityCheckSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityCheckSchemeResponse
func (client *Client) UpdateQualityCheckSchemeWithOptions(request *UpdateQualityCheckSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityCheckSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQualityCheckScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityCheckSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a quality check scheme.
//
// @param request - UpdateQualityCheckSchemeRequest
//
// @return UpdateQualityCheckSchemeResponse
func (client *Client) UpdateQualityCheckScheme(request *UpdateQualityCheckSchemeRequest) (_result *UpdateQualityCheckSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQualityCheckSchemeResponse{}
	_body, _err := client.UpdateQualityCheckSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateRule is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// Summary:
//
// Update rule information.
//
// Description:
//
// > Update an existing rule. You can modify its conditions and operators as needed. The rule ID (rid) remains unchanged, but condition IDs and operator IDs may change.
//
// @param request - UpdateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleResponse
func (client *Client) UpdateRuleWithOptions(request *UpdateRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRule"),
		Version:     dara.String("2019-01-15"),
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

// Deprecated: OpenAPI UpdateRule is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// Summary:
//
// Update rule information.
//
// Description:
//
// > Update an existing rule. You can modify its conditions and operators as needed. The rule ID (rid) remains unchanged, but condition IDs and operator IDs may change.
//
// @param request - UpdateRuleRequest
//
// @return UpdateRuleResponse
// Deprecated
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

// Deprecated: OpenAPI UpdateRuleById is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// Summary:
//
// This API maps to the frontend function location: Quality Inspection Rule Configuration - Create & Update. The Apsara Stack URL is: ip:port/api/client/UpdateRuleById.json.
//
// @param request - UpdateRuleByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleByIdResponse
func (client *Client) UpdateRuleByIdWithOptions(request *UpdateRuleByIdRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IsCopy) {
		body["IsCopy"] = request.IsCopy
	}

	if !dara.IsNil(request.JsonStrForRule) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	if !dara.IsNil(request.ReturnRelatedSchemes) {
		body["ReturnRelatedSchemes"] = request.ReturnRelatedSchemes
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRuleById"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateRuleById is deprecated, please use Qualitycheck::2019-01-15::UpdateRuleV4 instead.
//
// Summary:
//
// This API maps to the frontend function location: Quality Inspection Rule Configuration - Create & Update. The Apsara Stack URL is: ip:port/api/client/UpdateRuleById.json.
//
// @param request - UpdateRuleByIdRequest
//
// @return UpdateRuleByIdResponse
// Deprecated
func (client *Client) UpdateRuleById(request *UpdateRuleByIdRequest) (_result *UpdateRuleByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleByIdResponse{}
	_body, _err := client.UpdateRuleByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Frontend location: Quality Check Plan Management > Create or edit a quality check task > Associate quality check rules. Apsara Stack URL: ip:port/api/qcs/UpdateRuleToScheme.json.
//
// @param request - UpdateRuleToSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleToSchemeResponse
func (client *Client) UpdateRuleToSchemeWithOptions(request *UpdateRuleToSchemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRuleToSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRuleToScheme"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleToSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Frontend location: Quality Check Plan Management > Create or edit a quality check task > Associate quality check rules. Apsara Stack URL: ip:port/api/qcs/UpdateRuleToScheme.json.
//
// @param request - UpdateRuleToSchemeRequest
//
// @return UpdateRuleToSchemeResponse
func (client *Client) UpdateRuleToScheme(request *UpdateRuleToSchemeRequest) (_result *UpdateRuleToSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleToSchemeResponse{}
	_body, _err := client.UpdateRuleToSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Corresponding frontend feature location: Rule Configuration - Update. Apsara Stack URL: ip:port/api/client/UpdateRuleById.json.
//
// @param request - UpdateRuleV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRuleV4Response
func (client *Client) UpdateRuleV4WithOptions(request *UpdateRuleV4Request, runtime *dara.RuntimeOptions) (_result *UpdateRuleV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JsonStrForRule) {
		body["JsonStrForRule"] = request.JsonStrForRule
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRuleV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRuleV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Corresponding frontend feature location: Rule Configuration - Update. Apsara Stack URL: ip:port/api/client/UpdateRuleById.json.
//
// @param request - UpdateRuleV4Request
//
// @return UpdateRuleV4Response
func (client *Client) UpdateRuleV4(request *UpdateRuleV4Request) (_result *UpdateRuleV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRuleV4Response{}
	_body, _err := client.UpdateRuleV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Frontend path: Task Management > Edit any data on the right. Apsara Stack URL: ip:port/api/task/UpdateSchemeTaskConfig.json.
//
// Description:
//
// Updates quality inspection task information.
//
// @param request - UpdateSchemeTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSchemeTaskConfigResponse
func (client *Client) UpdateSchemeTaskConfigWithOptions(request *UpdateSchemeTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSchemeTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["jsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSchemeTaskConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSchemeTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Frontend path: Task Management > Edit any data on the right. Apsara Stack URL: ip:port/api/task/UpdateSchemeTaskConfig.json.
//
// Description:
//
// Updates quality inspection task information.
//
// @param request - UpdateSchemeTaskConfigRequest
//
// @return UpdateSchemeTaskConfigResponse
func (client *Client) UpdateSchemeTaskConfig(request *UpdateSchemeTaskConfigRequest) (_result *UpdateSchemeTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSchemeTaskConfigResponse{}
	_body, _err := client.UpdateSchemeTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateSkillGroupConfig is deprecated
//
// Summary:
//
// You can call UpdateSkillGroupConfig to update a configuration.
//
// @param request - UpdateSkillGroupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillGroupConfigResponse
func (client *Client) UpdateSkillGroupConfigWithOptions(request *UpdateSkillGroupConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillGroupConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillGroupConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillGroupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateSkillGroupConfig is deprecated
//
// Summary:
//
// You can call UpdateSkillGroupConfig to update a configuration.
//
// @param request - UpdateSkillGroupConfigRequest
//
// @return UpdateSkillGroupConfigResponse
// Deprecated
func (client *Client) UpdateSkillGroupConfig(request *UpdateSkillGroupConfigRequest) (_result *UpdateSkillGroupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSkillGroupConfigResponse{}
	_body, _err := client.UpdateSkillGroupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Maintain the recording information after real-time quality inspection is completed, which is used to play back the recording during review. After the recording information is maintained, the task status will change to Succeeded.
//
// @param request - UpdateSyncQualityCheckDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSyncQualityCheckDataResponse
func (client *Client) UpdateSyncQualityCheckDataWithOptions(request *UpdateSyncQualityCheckDataRequest, runtime *dara.RuntimeOptions) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSyncQualityCheckData"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSyncQualityCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Maintain the recording information after real-time quality inspection is completed, which is used to play back the recording during review. After the recording information is maintained, the task status will change to Succeeded.
//
// @param request - UpdateSyncQualityCheckDataRequest
//
// @return UpdateSyncQualityCheckDataResponse
func (client *Client) UpdateSyncQualityCheckData(request *UpdateSyncQualityCheckDataRequest) (_result *UpdateSyncQualityCheckDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSyncQualityCheckDataResponse{}
	_body, _err := client.UpdateSyncQualityCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a label node.
//
// @param request - UpdateTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTagResponse
func (client *Client) UpdateTagWithOptions(request *UpdateTagRequest, runtime *dara.RuntimeOptions) (_result *UpdateTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTag"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a label node.
//
// @param request - UpdateTagRequest
//
// @return UpdateTagResponse
func (client *Client) UpdateTag(request *UpdateTagRequest) (_result *UpdateTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTagResponse{}
	_body, _err := client.UpdateTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the automatic allocation rule for quality review tasks.
//
// @param request - UpdateTaskAssignRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskAssignRuleResponse
func (client *Client) UpdateTaskAssignRuleWithOptions(request *UpdateTaskAssignRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskAssignRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskAssignRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskAssignRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the automatic allocation rule for quality review tasks.
//
// @param request - UpdateTaskAssignRuleRequest
//
// @return UpdateTaskAssignRuleResponse
func (client *Client) UpdateTaskAssignRule(request *UpdateTaskAssignRuleRequest) (_result *UpdateTaskAssignRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskAssignRuleResponse{}
	_body, _err := client.UpdateTaskAssignRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update users by modifying their roles in batches.
//
// Description:
//
// When you update users, you can modify only their roles. You cannot modify other account information because all Alibaba Cloud products use a unified account management system. Smart Conversation Analysis uses these accounts. To modify account information, go to [Resource Access Management (RAM)](https://ram.console.aliyun.com/).
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update users by modifying their roles in batches.
//
// Description:
//
// When you update users, you can modify only their roles. You cannot modify other account information because all Alibaba Cloud products use a unified account management system. Smart Conversation Analysis uses these accounts. To modify account information, go to [Resource Access Management (RAM)](https://ram.console.aliyun.com/).
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the warning configuration.
//
// @param request - UpdateWarningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarningConfigResponse
func (client *Client) UpdateWarningConfigWithOptions(request *UpdateWarningConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateWarningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWarningConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWarningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the warning configuration.
//
// @param request - UpdateWarningConfigRequest
//
// @return UpdateWarningConfigResponse
func (client *Client) UpdateWarningConfig(request *UpdateWarningConfigRequest) (_result *UpdateWarningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWarningConfigResponse{}
	_body, _err := client.UpdateWarningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预警策略-更新
//
// @param request - UpdateWarningStrategyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarningStrategyConfigResponse
func (client *Client) UpdateWarningStrategyConfigWithOptions(request *UpdateWarningStrategyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateWarningStrategyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWarningStrategyConfig"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWarningStrategyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警策略-更新
//
// @param request - UpdateWarningStrategyConfigRequest
//
// @return UpdateWarningStrategyConfigResponse
func (client *Client) UpdateWarningStrategyConfig(request *UpdateWarningStrategyConfigRequest) (_result *UpdateWarningStrategyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWarningStrategyConfigResponse{}
	_body, _err := client.UpdateWarningStrategyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads offline voice quality inspection data (recording session files). This operation is applicable to hotline agent scenarios. Scenario 1: Natively integrates with Alibaba Cloud Call Center (CCC), requiring no development. You can enable one-click push of call data to SCA. Scenario 2: Integrates with your own call center system. Each time the call center generates a recording, it pushes the recording to SCA for analysis.
//
// Description:
//
// ### Process description
//
// Call the API to upload audio quality inspection => Convert the recording file to text => Separate roles in the text based on the specified channel splitting method (distinguish between agent and customer) => Analyze using quality inspection rules => Quality inspection complete.
//
// ### Task execution efficiency
//
// The speed of task execution depends on the speed of converting the recording file to text. Ideally, a 5-minute recording file can be transcribed within 2 minutes. However, when the transcription service has many queued tasks, there will be a queuing wait time. Generally, transcription completes within 6 hours, except for bulk uploads of large-scale data (more than 500 hours of recordings uploaded within 30 minutes). After transcription is complete, quality inspection analysis takes only milliseconds.
//
// ### Recording file URL requirements
//
// - Supports single-channel/dual-channel WAV and MP3 format recording files. The file size must be less than 512 MB.
//
// - The URL must be an HTTP-accessible URL address. Local file submission is not supported. The recording file access permissions must be set to public.
//
// - The URL can only use domain names, not IP addresses. The URL cannot contain spaces. Avoid using Chinese characters.
//
// - After converting the recording to text, the system deletes the downloaded recording file and does not retain a copy.
//
// - If your recording URL has an access expiration period (for example, the recording is stored in Alibaba Cloud OSS and you specified an expiration period when generating the recording URL through OSS), set the expiration period to at least 12 hours, or 24 hours if possible. This is because file transcription takes time and occasional queuing may occur. If the queuing time is long, the recording is downloaded only when transcription begins. This prevents the recording URL from expiring before the download.
//
// - After quality inspection analysis is complete, the recording is still played using the URL you provided when reviewing files in the console. Ensure that the URL remains active long-term. Otherwise, the recording cannot be played.
//
// ### Role separation description
//
// After the recording is converted to text, the system automatically separates the text into two conversation roles. However, the system cannot determine which role is the agent and which is the customer. You need to perform role separation based on certain rules. The accuracy of role separation is critical because the rules used for quality inspection analysis often have role detection restrictions (a rule only checks the agent or the customer). If role separation is incorrect, the accuracy of quality inspection results is significantly affected.
//
// Recording files are typically divided into two types: single-channel (mono) and dual-channel (stereo):
//
// - Single-channel recording: The voices of both the agent and customer are stored on one channel. After the recording file is converted to text, the system uses a built-in algorithm to distinguish between two roles. By setting a list of keywords that the agent is likely to say, the system analyzes the transcribed text sentence by sentence from top to bottom. When a sentence matches a keyword, the role of that sentence is determined to be the agent, and the other role is the customer. For details, see recognizeRoleDataSetId and serviceChannelKeywords in the request parameters. Due to the unpredictability of conversation content (for example, cross-talk between two roles or both people speaking simultaneously), role separation for single-channel recordings cannot be guaranteed to be 100% accurate. Save recording files as dual-channel recordings whenever possible.
//
// - Dual-channel recording: The voices of the agent and customer are stored on two separate channels. Even if the conversation overlaps, the recording-to-text conversion can accurately distinguish between the two. Specify the agent and customer by using the serviceChannel and clientChannel request parameters.
//
// ### Retrieve quality inspection analysis results
//
// Because recording file recognition is not real-time, you need to asynchronously retrieve quality inspection analysis results. The following three methods are available:
//
// - Message notification: For details, see [MSMQ](https://help.aliyun.com/document_detail/213237.html). After receiving a message, invoke the GetResult operation to retrieve detailed results. (Recommended)
//
// - Callback: Specify a callbackUrl in the request parameters. The system initiates a callback after the task is complete. After receiving the callback, invoke the GetResult operation to retrieve detailed results.
//
// - Polling: The operation returns a task ID (taskId). Use the taskId to poll the `getResult` operation to asynchronously retrieve results. Check whether the `status` in the response parameters indicates completion. Do not set the polling interval too short. Analysis normally completes within a few minutes. Set the polling interval to 30 seconds or more. (Not recommended)
//
// @param request - UploadAudioDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadAudioDataResponse
func (client *Client) UploadAudioDataWithOptions(request *UploadAudioDataRequest, runtime *dara.RuntimeOptions) (_result *UploadAudioDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadAudioData"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadAudioDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads offline voice quality inspection data (recording session files). This operation is applicable to hotline agent scenarios. Scenario 1: Natively integrates with Alibaba Cloud Call Center (CCC), requiring no development. You can enable one-click push of call data to SCA. Scenario 2: Integrates with your own call center system. Each time the call center generates a recording, it pushes the recording to SCA for analysis.
//
// Description:
//
// ### Process description
//
// Call the API to upload audio quality inspection => Convert the recording file to text => Separate roles in the text based on the specified channel splitting method (distinguish between agent and customer) => Analyze using quality inspection rules => Quality inspection complete.
//
// ### Task execution efficiency
//
// The speed of task execution depends on the speed of converting the recording file to text. Ideally, a 5-minute recording file can be transcribed within 2 minutes. However, when the transcription service has many queued tasks, there will be a queuing wait time. Generally, transcription completes within 6 hours, except for bulk uploads of large-scale data (more than 500 hours of recordings uploaded within 30 minutes). After transcription is complete, quality inspection analysis takes only milliseconds.
//
// ### Recording file URL requirements
//
// - Supports single-channel/dual-channel WAV and MP3 format recording files. The file size must be less than 512 MB.
//
// - The URL must be an HTTP-accessible URL address. Local file submission is not supported. The recording file access permissions must be set to public.
//
// - The URL can only use domain names, not IP addresses. The URL cannot contain spaces. Avoid using Chinese characters.
//
// - After converting the recording to text, the system deletes the downloaded recording file and does not retain a copy.
//
// - If your recording URL has an access expiration period (for example, the recording is stored in Alibaba Cloud OSS and you specified an expiration period when generating the recording URL through OSS), set the expiration period to at least 12 hours, or 24 hours if possible. This is because file transcription takes time and occasional queuing may occur. If the queuing time is long, the recording is downloaded only when transcription begins. This prevents the recording URL from expiring before the download.
//
// - After quality inspection analysis is complete, the recording is still played using the URL you provided when reviewing files in the console. Ensure that the URL remains active long-term. Otherwise, the recording cannot be played.
//
// ### Role separation description
//
// After the recording is converted to text, the system automatically separates the text into two conversation roles. However, the system cannot determine which role is the agent and which is the customer. You need to perform role separation based on certain rules. The accuracy of role separation is critical because the rules used for quality inspection analysis often have role detection restrictions (a rule only checks the agent or the customer). If role separation is incorrect, the accuracy of quality inspection results is significantly affected.
//
// Recording files are typically divided into two types: single-channel (mono) and dual-channel (stereo):
//
// - Single-channel recording: The voices of both the agent and customer are stored on one channel. After the recording file is converted to text, the system uses a built-in algorithm to distinguish between two roles. By setting a list of keywords that the agent is likely to say, the system analyzes the transcribed text sentence by sentence from top to bottom. When a sentence matches a keyword, the role of that sentence is determined to be the agent, and the other role is the customer. For details, see recognizeRoleDataSetId and serviceChannelKeywords in the request parameters. Due to the unpredictability of conversation content (for example, cross-talk between two roles or both people speaking simultaneously), role separation for single-channel recordings cannot be guaranteed to be 100% accurate. Save recording files as dual-channel recordings whenever possible.
//
// - Dual-channel recording: The voices of the agent and customer are stored on two separate channels. Even if the conversation overlaps, the recording-to-text conversion can accurately distinguish between the two. Specify the agent and customer by using the serviceChannel and clientChannel request parameters.
//
// ### Retrieve quality inspection analysis results
//
// Because recording file recognition is not real-time, you need to asynchronously retrieve quality inspection analysis results. The following three methods are available:
//
// - Message notification: For details, see [MSMQ](https://help.aliyun.com/document_detail/213237.html). After receiving a message, invoke the GetResult operation to retrieve detailed results. (Recommended)
//
// - Callback: Specify a callbackUrl in the request parameters. The system initiates a callback after the task is complete. After receiving the callback, invoke the GetResult operation to retrieve detailed results.
//
// - Polling: The operation returns a task ID (taskId). Use the taskId to poll the `getResult` operation to asynchronously retrieve results. Check whether the `status` in the response parameters indicates completion. Do not set the polling interval too short. Analysis normally completes within a few minutes. Set the polling interval to 30 seconds or more. (Not recommended)
//
// @param request - UploadAudioDataRequest
//
// @return UploadAudioDataResponse
func (client *Client) UploadAudioData(request *UploadAudioDataRequest) (_result *UploadAudioDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadAudioDataResponse{}
	_body, _err := client.UploadAudioDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UploadData is deprecated, please use Qualitycheck::2019-01-15::UploadDataV4 instead.
//
// Summary:
//
// Upload offline plain text quality inspection data (plain text sessions). This applies to online agent scenarios. Use the UploadDataV4 API. Differences between UploadDataV4 and UploadData include the following: UploadDataV4 supports only POST requests, and it supports longer JsonStr values.
//
// Description:
//
// You can call UploadData.json to upload text-based quality inspection data. Text typically originates from online customer service interactions or tickets. The API returns a task ID. You can retrieve results in one of three ways:
//
// - Message notification: For details, see [message queues](https://help.aliyun.com/document_detail/213237.html). After you receive a message, call the GetResult API to retrieve detailed results. (Recommended)
//
// - Callback: Specify a callback URL in your request parameters. After the task completes, the system sends a callback to that URL. Then call the GetResult API to retrieve detailed results.
//
// - Polling: Use the returned task ID to poll the GetResult API asynchronously. Check whether the status field in the response indicates completion. (Not recommended)
//
// @param request - UploadDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataResponse
func (client *Client) UploadDataWithOptions(request *UploadDataRequest, runtime *dara.RuntimeOptions) (_result *UploadDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadData"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UploadData is deprecated, please use Qualitycheck::2019-01-15::UploadDataV4 instead.
//
// Summary:
//
// Upload offline plain text quality inspection data (plain text sessions). This applies to online agent scenarios. Use the UploadDataV4 API. Differences between UploadDataV4 and UploadData include the following: UploadDataV4 supports only POST requests, and it supports longer JsonStr values.
//
// Description:
//
// You can call UploadData.json to upload text-based quality inspection data. Text typically originates from online customer service interactions or tickets. The API returns a task ID. You can retrieve results in one of three ways:
//
// - Message notification: For details, see [message queues](https://help.aliyun.com/document_detail/213237.html). After you receive a message, call the GetResult API to retrieve detailed results. (Recommended)
//
// - Callback: Specify a callback URL in your request parameters. After the task completes, the system sends a callback to that URL. Then call the GetResult API to retrieve detailed results.
//
// - Polling: Use the returned task ID to poll the GetResult API asynchronously. Check whether the status field in the response indicates completion. (Not recommended)
//
// @param request - UploadDataRequest
//
// @return UploadDataResponse
// Deprecated
func (client *Client) UploadData(request *UploadDataRequest) (_result *UploadDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDataResponse{}
	_body, _err := client.UploadDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs real-time text-based quality inspection.
//
// Description:
//
// Pushes text data in a specific format to SCA for real-time quality inspection analysis based on user-specified rules, and synchronously returns the analysis results. Compared with uploaded text quality inspection, which typically uploads the complete conversation text after a conversation ends, real-time text quality inspection allows you to push text to SCA for analysis after one role finishes one or more sentences, providing higher real-time performance. Notes:
//
// - If the pushed text is a single sentence from one role, some operators in the rules become ineffective due to the lack of conversation context, such as context repetition check, interruption check, and call silence check.
//
// - Real-time quality inspection synchronously returns analysis results. SCA does not save call records, so you cannot query quality inspection results through APIs.
//
// @param request - UploadDataSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataSyncResponse
func (client *Client) UploadDataSyncWithOptions(request *UploadDataSyncRequest, runtime *dara.RuntimeOptions) (_result *UploadDataSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDataSync"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDataSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs real-time text-based quality inspection.
//
// Description:
//
// Pushes text data in a specific format to SCA for real-time quality inspection analysis based on user-specified rules, and synchronously returns the analysis results. Compared with uploaded text quality inspection, which typically uploads the complete conversation text after a conversation ends, real-time text quality inspection allows you to push text to SCA for analysis after one role finishes one or more sentences, providing higher real-time performance. Notes:
//
// - If the pushed text is a single sentence from one role, some operators in the rules become ineffective due to the lack of conversation context, such as context repetition check, interruption check, and call silence check.
//
// - Real-time quality inspection synchronously returns analysis results. SCA does not save call records, so you cannot query quality inspection results through APIs.
//
// @param request - UploadDataSyncRequest
//
// @return UploadDataSyncResponse
func (client *Client) UploadDataSync(request *UploadDataSyncRequest) (_result *UploadDataSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDataSyncResponse{}
	_body, _err := client.UploadDataSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs real-time text quality inspection using a large language model.
//
// Description:
//
// Pushes text data in a specific format to Smart Conversation Analysis (SCA) for real-time quality inspection based on user-specified rules, and synchronously returns the analysis results. Compared with uploaded text quality inspection, which typically uploads the complete conversation text after a conversation ends, real-time text quality inspection allows you to push text to SCA for analysis after one or more sentences are spoken by a role, providing higher real-time performance. Special notes:
//
// If the pushed text is a single sentence from one role, some operators in the rules may not work due to the lack of conversation context, such as context repetition check, interruption check, and call silence check.
//
// Real-time quality inspection synchronously returns analysis results. SCA does not save call records, so you cannot query quality inspection results through APIs.
//
// @param request - UploadDataSyncForLLMRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataSyncForLLMResponse
func (client *Client) UploadDataSyncForLLMWithOptions(request *UploadDataSyncForLLMRequest, runtime *dara.RuntimeOptions) (_result *UploadDataSyncForLLMResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDataSyncForLLM"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDataSyncForLLMResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs real-time text quality inspection using a large language model.
//
// Description:
//
// Pushes text data in a specific format to Smart Conversation Analysis (SCA) for real-time quality inspection based on user-specified rules, and synchronously returns the analysis results. Compared with uploaded text quality inspection, which typically uploads the complete conversation text after a conversation ends, real-time text quality inspection allows you to push text to SCA for analysis after one or more sentences are spoken by a role, providing higher real-time performance. Special notes:
//
// If the pushed text is a single sentence from one role, some operators in the rules may not work due to the lack of conversation context, such as context repetition check, interruption check, and call silence check.
//
// Real-time quality inspection synchronously returns analysis results. SCA does not save call records, so you cannot query quality inspection results through APIs.
//
// @param request - UploadDataSyncForLLMRequest
//
// @return UploadDataSyncForLLMResponse
func (client *Client) UploadDataSyncForLLM(request *UploadDataSyncForLLMRequest) (_result *UploadDataSyncForLLMResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDataSyncForLLMResponse{}
	_body, _err := client.UploadDataSyncForLLMWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads offline text quality inspection data (plain text conversations) applicable to online agent scenarios. We recommend using the UploadDataV4 operation. Differences between UploadDataV4 and UploadData: 1. V4 supports only POST requests. 2. V4 supports longer JsonStr values.
//
// Description:
//
// Calls UploadData.json to upload text quality inspection data. The text typically comes from online customer service or tickets. The operation returns a taskId. You can obtain results in three ways:
//
// - Message notification: For details, see [Message channel](https://help.aliyun.com/document_detail/213237.html). After receiving a message, call the GetResult operation to obtain detailed results. (Recommended)
//
// - Callback: Specify a callbackUrl in the request parameters. The system initiates a callback after the task is completed. After receiving the callback, call the GetResult operation to obtain detailed results.
//
// - Polling: Use the task ID returned by this operation to poll the GetResult operation asynchronously. Check whether the status in the response indicates completion. (Not recommended)
//
// @param request - UploadDataV4Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDataV4Response
func (client *Client) UploadDataV4WithOptions(request *UploadDataV4Request, runtime *dara.RuntimeOptions) (_result *UploadDataV4Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDataV4"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDataV4Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads offline text quality inspection data (plain text conversations) applicable to online agent scenarios. We recommend using the UploadDataV4 operation. Differences between UploadDataV4 and UploadData: 1. V4 supports only POST requests. 2. V4 supports longer JsonStr values.
//
// Description:
//
// Calls UploadData.json to upload text quality inspection data. The text typically comes from online customer service or tickets. The operation returns a taskId. You can obtain results in three ways:
//
// - Message notification: For details, see [Message channel](https://help.aliyun.com/document_detail/213237.html). After receiving a message, call the GetResult operation to obtain detailed results. (Recommended)
//
// - Callback: Specify a callbackUrl in the request parameters. The system initiates a callback after the task is completed. After receiving the callback, call the GetResult operation to obtain detailed results.
//
// - Polling: Use the task ID returned by this operation to poll the GetResult operation asynchronously. Check whether the status in the response indicates completion. (Not recommended)
//
// @param request - UploadDataV4Request
//
// @return UploadDataV4Response
func (client *Client) UploadDataV4(request *UploadDataV4Request) (_result *UploadDataV4Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadDataV4Response{}
	_body, _err := client.UploadDataV4WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a rule. You can use this operation to provide your own rule editing interface.
//
// Description:
//
// > For more information, see [Rule configuration](https://help.aliyun.com/document_detail/213225.html).
//
// @param request - UploadRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRuleResponse
func (client *Client) UploadRuleWithOptions(request *UploadRuleRequest, runtime *dara.RuntimeOptions) (_result *UploadRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		query["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		query["JsonStr"] = request.JsonStr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadRule"),
		Version:     dara.String("2019-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a rule. You can use this operation to provide your own rule editing interface.
//
// Description:
//
// > For more information, see [Rule configuration](https://help.aliyun.com/document_detail/213225.html).
//
// @param request - UploadRuleRequest
//
// @return UploadRuleResponse
func (client *Client) UploadRule(request *UploadRuleRequest) (_result *UploadRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadRuleResponse{}
	_body, _err := client.UploadRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) executeAgentWithSSE_opYieldFunc(_yield chan *ExecuteAgentResponse, _yieldErr chan error, request *ExecuteAgentRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseMeAgentId) {
		body["BaseMeAgentId"] = request.BaseMeAgentId
	}

	if !dara.IsNil(request.JsonStr) {
		body["JsonStr"] = request.JsonStr
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAgent"),
		Version:     dara.String("2019-01-15"),
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

func (client *Client) runCompletionMessageWithSSE_opYieldFunc(_yield chan *RunCompletionMessageResponse, _yieldErr chan error, tmpReq *RunCompletionMessageRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &RunCompletionMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Messages) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, dara.String("Messages"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MessagesShrink) {
		body["Messages"] = request.MessagesShrink
	}

	if !dara.IsNil(request.ModelCode) {
		body["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCompletionMessage"),
		Version:     dara.String("2019-01-15"),
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
