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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        dara.String("ccc.aliyuncs.com"),
		"ap-south-1":            dara.String("ccc.aliyuncs.com"),
		"ap-southeast-1":        dara.String("ccc.aliyuncs.com"),
		"ap-southeast-2":        dara.String("ccc.aliyuncs.com"),
		"ap-southeast-3":        dara.String("ccc.aliyuncs.com"),
		"ap-southeast-5":        dara.String("ccc.aliyuncs.com"),
		"cn-beijing":            dara.String("ccc.aliyuncs.com"),
		"cn-chengdu":            dara.String("ccc.aliyuncs.com"),
		"cn-hongkong":           dara.String("ccc.aliyuncs.com"),
		"cn-huhehaote":          dara.String("ccc.aliyuncs.com"),
		"cn-qingdao":            dara.String("ccc.aliyuncs.com"),
		"cn-shenzhen":           dara.String("ccc.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("ccc.aliyuncs.com"),
		"eu-central-1":          dara.String("ccc.aliyuncs.com"),
		"eu-west-1":             dara.String("ccc.aliyuncs.com"),
		"me-east-1":             dara.String("ccc.aliyuncs.com"),
		"us-east-1":             dara.String("ccc.aliyuncs.com"),
		"us-west-1":             dara.String("ccc.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("ccc.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("ccc.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("ccc.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("ccc.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ccc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 废弃预测式外呼活动
//
// @param request - AbortCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortCampaignResponse
func (client *Client) AbortCampaignWithOptions(request *AbortCampaignRequest, runtime *dara.RuntimeOptions) (_result *AbortCampaignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 废弃预测式外呼活动
//
// @param request - AbortCampaignRequest
//
// @return AbortCampaignResponse
func (client *Client) AbortCampaign(request *AbortCampaignRequest) (_result *AbortCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AbortCampaignResponse{}
	_body, _err := client.AbortCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AcceptChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptChatResponse
func (client *Client) AcceptChatWithOptions(request *AcceptChatRequest, runtime *dara.RuntimeOptions) (_result *AcceptChatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptChat"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AcceptChatRequest
//
// @return AcceptChatResponse
func (client *Client) AcceptChat(request *AcceptChatRequest) (_result *AcceptChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptChatResponse{}
	_body, _err := client.AcceptChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑呼入控制号码
//
// @param request - AddBlacklistCallTaggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBlacklistCallTaggingResponse
func (client *Client) AddBlacklistCallTaggingWithOptions(request *AddBlacklistCallTaggingRequest, runtime *dara.RuntimeOptions) (_result *AddBlacklistCallTaggingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBlacklistCallTagging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBlacklistCallTaggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑呼入控制号码
//
// @param request - AddBlacklistCallTaggingRequest
//
// @return AddBlacklistCallTaggingResponse
func (client *Client) AddBlacklistCallTagging(request *AddBlacklistCallTaggingRequest) (_result *AddBlacklistCallTaggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBlacklistCallTaggingResponse{}
	_body, _err := client.AddBlacklistCallTaggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 追加联系人
//
// @param tmpReq - AddCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCasesResponse
func (client *Client) AddCasesWithOptions(tmpReq *AddCasesRequest, runtime *dara.RuntimeOptions) (_result *AddCasesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddCasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CaseList) {
		request.CaseListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseList, dara.String("CaseList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.CaseListShrink) {
		query["CaseList"] = request.CaseListShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCases"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 追加联系人
//
// @param request - AddCasesRequest
//
// @return AddCasesResponse
func (client *Client) AddCases(request *AddCasesRequest) (_result *AddCasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCasesResponse{}
	_body, _err := client.AddCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddFeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFeedbackResponse
func (client *Client) AddFeedbackWithOptions(request *AddFeedbackRequest, runtime *dara.RuntimeOptions) (_result *AddFeedbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Feedback) {
		query["Feedback"] = request.Feedback
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Rating) {
		query["Rating"] = request.Rating
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFeedback"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddFeedbackRequest
//
// @return AddFeedbackResponse
func (client *Client) AddFeedback(request *AddFeedbackRequest) (_result *AddFeedbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFeedbackResponse{}
	_body, _err := client.AddFeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddNumbersToSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddNumbersToSkillGroupResponse
func (client *Client) AddNumbersToSkillGroupWithOptions(request *AddNumbersToSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *AddNumbersToSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstNumberGroupIdList) {
		query["InstNumberGroupIdList"] = request.InstNumberGroupIdList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddNumbersToSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddNumbersToSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddNumbersToSkillGroupRequest
//
// @return AddNumbersToSkillGroupResponse
func (client *Client) AddNumbersToSkillGroup(request *AddNumbersToSkillGroupRequest) (_result *AddNumbersToSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddNumbersToSkillGroupResponse{}
	_body, _err := client.AddNumbersToSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddPersonalNumbersToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPersonalNumbersToUserResponse
func (client *Client) AddPersonalNumbersToUserWithOptions(request *AddPersonalNumbersToUserRequest, runtime *dara.RuntimeOptions) (_result *AddPersonalNumbersToUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPersonalNumbersToUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPersonalNumbersToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddPersonalNumbersToUserRequest
//
// @return AddPersonalNumbersToUserResponse
func (client *Client) AddPersonalNumbersToUser(request *AddPersonalNumbersToUserRequest) (_result *AddPersonalNumbersToUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPersonalNumbersToUserResponse{}
	_body, _err := client.AddPersonalNumbersToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddPhoneNumberToSkillGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPhoneNumberToSkillGroupsResponse
func (client *Client) AddPhoneNumberToSkillGroupsWithOptions(request *AddPhoneNumberToSkillGroupsRequest, runtime *dara.RuntimeOptions) (_result *AddPhoneNumberToSkillGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPhoneNumberToSkillGroups"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPhoneNumberToSkillGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddPhoneNumberToSkillGroupsRequest
//
// @return AddPhoneNumberToSkillGroupsResponse
func (client *Client) AddPhoneNumberToSkillGroups(request *AddPhoneNumberToSkillGroupsRequest) (_result *AddPhoneNumberToSkillGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPhoneNumberToSkillGroupsResponse{}
	_body, _err := client.AddPhoneNumberToSkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddPhoneNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPhoneNumbersResponse
func (client *Client) AddPhoneNumbersWithOptions(request *AddPhoneNumbersRequest, runtime *dara.RuntimeOptions) (_result *AddPhoneNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberGroupId) {
		query["NumberGroupId"] = request.NumberGroupId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.Usage) {
		query["Usage"] = request.Usage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPhoneNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPhoneNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddPhoneNumbersRequest
//
// @return AddPhoneNumbersResponse
func (client *Client) AddPhoneNumbers(request *AddPhoneNumbersRequest) (_result *AddPhoneNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPhoneNumbersResponse{}
	_body, _err := client.AddPhoneNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - AddSchemaPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSchemaPropertyResponse
func (client *Client) AddSchemaPropertyWithOptions(tmpReq *AddSchemaPropertyRequest, runtime *dara.RuntimeOptions) (_result *AddSchemaPropertyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddSchemaPropertyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Property) {
		request.PropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Property, dara.String("Property"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PropertyShrink) {
		body["Property"] = request.PropertyShrink
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSchemaProperty"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSchemaPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddSchemaPropertyRequest
//
// @return AddSchemaPropertyResponse
func (client *Client) AddSchemaProperty(request *AddSchemaPropertyRequest) (_result *AddSchemaPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSchemaPropertyResponse{}
	_body, _err := client.AddSchemaPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddSkillGroupsToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSkillGroupsToUserResponse
func (client *Client) AddSkillGroupsToUserWithOptions(request *AddSkillGroupsToUserRequest, runtime *dara.RuntimeOptions) (_result *AddSkillGroupsToUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillLevelList) {
		query["SkillLevelList"] = request.SkillLevelList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSkillGroupsToUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSkillGroupsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddSkillGroupsToUserRequest
//
// @return AddSkillGroupsToUserResponse
func (client *Client) AddSkillGroupsToUser(request *AddSkillGroupsToUserRequest) (_result *AddSkillGroupsToUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSkillGroupsToUserResponse{}
	_body, _err := client.AddSkillGroupsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddTicketTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTicketTaskResponse
func (client *Client) AddTicketTaskWithOptions(request *AddTicketTaskRequest, runtime *dara.RuntimeOptions) (_result *AddTicketTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Assignee) {
		query["Assignee"] = request.Assignee
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Position) {
		query["Position"] = request.Position
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTicketTask"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTicketTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddTicketTaskRequest
//
// @return AddTicketTaskResponse
func (client *Client) AddTicketTask(request *AddTicketTaskRequest) (_result *AddTicketTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTicketTaskResponse{}
	_body, _err := client.AddTicketTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddUsersToSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUsersToSkillGroupResponse
func (client *Client) AddUsersToSkillGroupWithOptions(request *AddUsersToSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUsersToSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.UserSkillLevelList) {
		query["UserSkillLevelList"] = request.UserSkillLevelList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUsersToSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUsersToSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddUsersToSkillGroupRequest
//
// @return AddUsersToSkillGroupResponse
func (client *Client) AddUsersToSkillGroup(request *AddUsersToSkillGroupRequest) (_result *AddUsersToSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUsersToSkillGroupResponse{}
	_body, _err := client.AddUsersToSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AnalyzeConversationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeConversationResponse
func (client *Client) AnalyzeConversationWithOptions(request *AnalyzeConversationRequest, runtime *dara.RuntimeOptions) (_result *AnalyzeConversationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.FieldListJson) {
		query["FieldListJson"] = request.FieldListJson
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskListJson) {
		query["TaskListJson"] = request.TaskListJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeConversation"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AnalyzeConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AnalyzeConversationRequest
//
// @return AnalyzeConversationResponse
func (client *Client) AnalyzeConversation(request *AnalyzeConversationRequest) (_result *AnalyzeConversationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AnalyzeConversationResponse{}
	_body, _err := client.AnalyzeConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AnswerCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnswerCallResponse
func (client *Client) AnswerCallWithOptions(request *AnswerCallRequest, runtime *dara.RuntimeOptions) (_result *AnswerCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnswerCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AnswerCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AnswerCallRequest
//
// @return AnswerCallResponse
func (client *Client) AnswerCall(request *AnswerCallRequest) (_result *AnswerCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AnswerCallResponse{}
	_body, _err := client.AnswerCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 追加联系人
//
// @param tmpReq - AppendCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppendCasesResponse
func (client *Client) AppendCasesWithOptions(tmpReq *AppendCasesRequest, runtime *dara.RuntimeOptions) (_result *AppendCasesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AppendCasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AppendCases"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AppendCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 追加联系人
//
// @param request - AppendCasesRequest
//
// @return AppendCasesResponse
func (client *Client) AppendCases(request *AppendCasesRequest) (_result *AppendCasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AppendCasesResponse{}
	_body, _err := client.AppendCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI AssignUsers is deprecated, please use CCC::2020-07-01::ImportRamUsers instead.
//
// @param request - AssignUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignUsersResponse
// Deprecated
func (client *Client) AssignUsersWithOptions(request *AssignUsersRequest, runtime *dara.RuntimeOptions) (_result *AssignUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RamIdList) {
		query["RamIdList"] = request.RamIdList
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.SkillLevelList) {
		query["SkillLevelList"] = request.SkillLevelList
	}

	if !dara.IsNil(request.WorkMode) {
		query["WorkMode"] = request.WorkMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignUsers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AssignUsers is deprecated, please use CCC::2020-07-01::ImportRamUsers instead.
//
// @param request - AssignUsersRequest
//
// @return AssignUsersResponse
// Deprecated
func (client *Client) AssignUsers(request *AssignUsersRequest) (_result *AssignUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssignUsersResponse{}
	_body, _err := client.AssignUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BargeInCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BargeInCallResponse
func (client *Client) BargeInCallWithOptions(request *BargeInCallRequest, runtime *dara.RuntimeOptions) (_result *BargeInCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BargedUserId) {
		query["BargedUserId"] = request.BargedUserId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BargeInCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BargeInCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BargeInCallRequest
//
// @return BargeInCallResponse
func (client *Client) BargeInCall(request *BargeInCallRequest) (_result *BargeInCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BargeInCallResponse{}
	_body, _err := client.BargeInCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BlindTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BlindTransferResponse
func (client *Client) BlindTransferWithOptions(request *BlindTransferRequest, runtime *dara.RuntimeOptions) (_result *BlindTransferResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallPriority) {
		query["CallPriority"] = request.CallPriority
	}

	if !dara.IsNil(request.ContactFlowVariables) {
		query["ContactFlowVariables"] = request.ContactFlowVariables
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.QueuingOverflowThreshold) {
		query["QueuingOverflowThreshold"] = request.QueuingOverflowThreshold
	}

	if !dara.IsNil(request.QueuingTimeoutSeconds) {
		query["QueuingTimeoutSeconds"] = request.QueuingTimeoutSeconds
	}

	if !dara.IsNil(request.RoutingType) {
		query["RoutingType"] = request.RoutingType
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	if !dara.IsNil(request.StrategyParams) {
		query["StrategyParams"] = request.StrategyParams
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.Transferee) {
		query["Transferee"] = request.Transferee
	}

	if !dara.IsNil(request.TransfereeType) {
		query["TransfereeType"] = request.TransfereeType
	}

	if !dara.IsNil(request.Transferor) {
		query["Transferor"] = request.Transferor
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BlindTransfer"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BlindTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BlindTransferRequest
//
// @return BlindTransferResponse
func (client *Client) BlindTransfer(request *BlindTransferRequest) (_result *BlindTransferResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BlindTransferResponse{}
	_body, _err := client.BlindTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BridgeRtcCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BridgeRtcCallResponse
func (client *Client) BridgeRtcCallWithOptions(request *BridgeRtcCallRequest, runtime *dara.RuntimeOptions) (_result *BridgeRtcCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ServiceProvider) {
		query["ServiceProvider"] = request.ServiceProvider
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.VideoEnabled) {
		query["VideoEnabled"] = request.VideoEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BridgeRtcCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BridgeRtcCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BridgeRtcCallRequest
//
// @return BridgeRtcCallResponse
func (client *Client) BridgeRtcCall(request *BridgeRtcCallRequest) (_result *BridgeRtcCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BridgeRtcCallResponse{}
	_body, _err := client.BridgeRtcCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CancelAttendedTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAttendedTransferResponse
func (client *Client) CancelAttendedTransferWithOptions(request *CancelAttendedTransferRequest, runtime *dara.RuntimeOptions) (_result *CancelAttendedTransferResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAttendedTransfer"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAttendedTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CancelAttendedTransferRequest
//
// @return CancelAttendedTransferResponse
func (client *Client) CancelAttendedTransfer(request *CancelAttendedTransferRequest) (_result *CancelAttendedTransferResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelAttendedTransferResponse{}
	_body, _err := client.CancelAttendedTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ChangeVisibilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeVisibilityResponse
func (client *Client) ChangeVisibilityWithOptions(request *ChangeVisibilityRequest, runtime *dara.RuntimeOptions) (_result *ChangeVisibilityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Invisible) {
		query["Invisible"] = request.Invisible
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeVisibility"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeVisibilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeVisibilityRequest
//
// @return ChangeVisibilityResponse
func (client *Client) ChangeVisibility(request *ChangeVisibilityRequest) (_result *ChangeVisibilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeVisibilityResponse{}
	_body, _err := client.ChangeVisibilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ChangeWorkModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeWorkModeResponse
func (client *Client) ChangeWorkModeWithOptions(request *ChangeWorkModeRequest, runtime *dara.RuntimeOptions) (_result *ChangeWorkModeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.SignedSkillGroupIdList) {
		query["SignedSkillGroupIdList"] = request.SignedSkillGroupIdList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkMode) {
		query["WorkMode"] = request.WorkMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeWorkMode"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeWorkModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeWorkModeRequest
//
// @return ChangeWorkModeResponse
func (client *Client) ChangeWorkMode(request *ChangeWorkModeRequest) (_result *ChangeWorkModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeWorkModeResponse{}
	_body, _err := client.ChangeWorkModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ClaimChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClaimChatResponse
func (client *Client) ClaimChatWithOptions(request *ClaimChatRequest, runtime *dara.RuntimeOptions) (_result *ClaimChatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClaimChat"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClaimChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ClaimChatRequest
//
// @return ClaimChatResponse
func (client *Client) ClaimChat(request *ClaimChatRequest) (_result *ClaimChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClaimChatResponse{}
	_body, _err := client.ClaimChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CoachCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CoachCallResponse
func (client *Client) CoachCallWithOptions(request *CoachCallRequest, runtime *dara.RuntimeOptions) (_result *CoachCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoachedUserId) {
		query["CoachedUserId"] = request.CoachedUserId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CoachCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CoachCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CoachCallRequest
//
// @return CoachCallResponse
func (client *Client) CoachCall(request *CoachCallRequest) (_result *CoachCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CoachCallResponse{}
	_body, _err := client.CoachCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CommitContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitContactFlowResponse
func (client *Client) CommitContactFlowWithOptions(request *CommitContactFlowRequest, runtime *dara.RuntimeOptions) (_result *CommitContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.Definition) {
		query["Definition"] = request.Definition
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DraftId) {
		query["DraftId"] = request.DraftId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommitContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CommitContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CommitContactFlowRequest
//
// @return CommitContactFlowResponse
func (client *Client) CommitContactFlow(request *CommitContactFlowRequest) (_result *CommitContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CommitContactFlowResponse{}
	_body, _err := client.CommitContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CompleteAttendedTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteAttendedTransferResponse
func (client *Client) CompleteAttendedTransferWithOptions(request *CompleteAttendedTransferRequest, runtime *dara.RuntimeOptions) (_result *CompleteAttendedTransferResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteAttendedTransfer"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteAttendedTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CompleteAttendedTransferRequest
//
// @return CompleteAttendedTransferResponse
func (client *Client) CompleteAttendedTransfer(request *CompleteAttendedTransferRequest) (_result *CompleteAttendedTransferResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompleteAttendedTransferResponse{}
	_body, _err := client.CompleteAttendedTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAudioFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAudioFileResponse
func (client *Client) CreateAudioFileWithOptions(request *CreateAudioFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAudioFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioFileName) {
		query["AudioFileName"] = request.AudioFileName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssFileKey) {
		query["OssFileKey"] = request.OssFileKey
	}

	if !dara.IsNil(request.Usage) {
		query["Usage"] = request.Usage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAudioFile"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAudioFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAudioFileRequest
//
// @return CreateAudioFileResponse
func (client *Client) CreateAudioFile(request *CreateAudioFileRequest) (_result *CreateAudioFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAudioFileResponse{}
	_body, _err := client.CreateAudioFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateCallSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCallSummaryResponse
func (client *Client) CreateCallSummaryWithOptions(request *CreateCallSummaryRequest, runtime *dara.RuntimeOptions) (_result *CreateCallSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCallSummary"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCallSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateCallSummaryRequest
//
// @return CreateCallSummaryResponse
func (client *Client) CreateCallSummary(request *CreateCallSummaryRequest) (_result *CreateCallSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCallSummaryResponse{}
	_body, _err := client.CreateCallSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量创建号码标签
//
// @param request - CreateCallTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCallTagsResponse
func (client *Client) CreateCallTagsWithOptions(request *CreateCallTagsRequest, runtime *dara.RuntimeOptions) (_result *CreateCallTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallTagNameList) {
		query["CallTagNameList"] = request.CallTagNameList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCallTags"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCallTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建号码标签
//
// @param request - CreateCallTagsRequest
//
// @return CreateCallTagsResponse
func (client *Client) CreateCallTags(request *CreateCallTagsRequest) (_result *CreateCallTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCallTagsResponse{}
	_body, _err := client.CreateCallTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建预测式外呼活动
//
// @param tmpReq - CreateCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCampaignResponse
func (client *Client) CreateCampaignWithOptions(tmpReq *CreateCampaignRequest, runtime *dara.RuntimeOptions) (_result *CreateCampaignResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateCampaignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CaseList) {
		request.CaseListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseList, dara.String("CaseList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallableTime) {
		query["CallableTime"] = request.CallableTime
	}

	if !dara.IsNil(request.CaseFileKey) {
		query["CaseFileKey"] = request.CaseFileKey
	}

	if !dara.IsNil(request.CaseListShrink) {
		query["CaseList"] = request.CaseListShrink
	}

	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecutingUntilTimeout) {
		query["ExecutingUntilTimeout"] = request.ExecutingUntilTimeout
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxAttemptCount) {
		query["MaxAttemptCount"] = request.MaxAttemptCount
	}

	if !dara.IsNil(request.MinAttemptInterval) {
		query["MinAttemptInterval"] = request.MinAttemptInterval
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.QueueId) {
		query["QueueId"] = request.QueueId
	}

	if !dara.IsNil(request.Simulation) {
		query["Simulation"] = request.Simulation
	}

	if !dara.IsNil(request.SimulationParameters) {
		query["SimulationParameters"] = request.SimulationParameters
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StrategyParameters) {
		query["StrategyParameters"] = request.StrategyParameters
	}

	if !dara.IsNil(request.StrategyType) {
		query["StrategyType"] = request.StrategyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预测式外呼活动
//
// @param request - CreateCampaignRequest
//
// @return CreateCampaignResponse
func (client *Client) CreateCampaign(request *CreateCampaignRequest) (_result *CreateCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCampaignResponse{}
	_body, _err := client.CreateCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateChatMediaUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatMediaUrlResponse
func (client *Client) CreateChatMediaUrlWithOptions(request *CreateChatMediaUrlRequest, runtime *dara.RuntimeOptions) (_result *CreateChatMediaUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MimeType) {
		body["MimeType"] = request.MimeType
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatMediaUrl"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatMediaUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateChatMediaUrlRequest
//
// @return CreateChatMediaUrlResponse
func (client *Client) CreateChatMediaUrl(request *CreateChatMediaUrlRequest) (_result *CreateChatMediaUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChatMediaUrlResponse{}
	_body, _err := client.CreateChatMediaUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContactFlowResponse
func (client *Client) CreateContactFlowWithOptions(request *CreateContactFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Definition) {
		query["Definition"] = request.Definition
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateContactFlowRequest
//
// @return CreateContactFlowResponse
func (client *Client) CreateContactFlow(request *CreateContactFlowRequest) (_result *CreateContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateContactFlowResponse{}
	_body, _err := client.CreateContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateCustomCallTagging is deprecated, please use CCC::2020-07-01::CreateCustomCallTaggings instead.
//
// Summary:
//
// 创建呼入控制号码
//
// @param request - CreateCustomCallTaggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomCallTaggingResponse
// Deprecated
func (client *Client) CreateCustomCallTaggingWithOptions(request *CreateCustomCallTaggingRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomCallTaggingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomNumberList) {
		query["CustomNumberList"] = request.CustomNumberList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomCallTagging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomCallTaggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateCustomCallTagging is deprecated, please use CCC::2020-07-01::CreateCustomCallTaggings instead.
//
// Summary:
//
// 创建呼入控制号码
//
// @param request - CreateCustomCallTaggingRequest
//
// @return CreateCustomCallTaggingResponse
// Deprecated
func (client *Client) CreateCustomCallTagging(request *CreateCustomCallTaggingRequest) (_result *CreateCustomCallTaggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomCallTaggingResponse{}
	_body, _err := client.CreateCustomCallTaggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminRamIdList) {
		query["AdminRamIdList"] = request.AdminRamIdList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - CreateSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchemaResponse
func (client *Client) CreateSchemaWithOptions(tmpReq *CreateSchemaRequest, runtime *dara.RuntimeOptions) (_result *CreateSchemaResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateSchemaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Properties) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, dara.String("Properties"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PropertiesShrink) {
		body["Properties"] = request.PropertiesShrink
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSchema"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateSchemaRequest
//
// @return CreateSchemaResponse
func (client *Client) CreateSchema(request *CreateSchemaRequest) (_result *CreateSchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSchemaResponse{}
	_body, _err := client.CreateSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillGroupResponse
func (client *Client) CreateSkillGroupWithOptions(request *CreateSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateSkillGroupRequest
//
// @return CreateSkillGroupResponse
func (client *Client) CreateSkillGroup(request *CreateSkillGroupRequest) (_result *CreateSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.CreateSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTicket"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarUrl) {
		query["AvatarUrl"] = request.AvatarUrl
	}

	if !dara.IsNil(request.DisplayId) {
		query["DisplayId"] = request.DisplayId
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LoginName) {
		query["LoginName"] = request.LoginName
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.Nickname) {
		query["Nickname"] = request.Nickname
	}

	if !dara.IsNil(request.ResetPassword) {
		query["ResetPassword"] = request.ResetPassword
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.SkillLevelList) {
		query["SkillLevelList"] = request.SkillLevelList
	}

	if !dara.IsNil(request.WorkMode) {
		query["WorkMode"] = request.WorkMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2020-07-01"),
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

// @param request - DeleteAudioFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAudioFileResponse
func (client *Client) DeleteAudioFileWithOptions(request *DeleteAudioFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteAudioFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioResourceId) {
		query["AudioResourceId"] = request.AudioResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAudioFile"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAudioFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAudioFileRequest
//
// @return DeleteAudioFileResponse
func (client *Client) DeleteAudioFile(request *DeleteAudioFileRequest) (_result *DeleteAudioFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAudioFileResponse{}
	_body, _err := client.DeleteAudioFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除号码标签
//
// @param request - DeleteCallTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCallTagResponse
func (client *Client) DeleteCallTagWithOptions(request *DeleteCallTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteCallTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCallTag"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCallTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除号码标签
//
// @param request - DeleteCallTagRequest
//
// @return DeleteCallTagResponse
func (client *Client) DeleteCallTag(request *DeleteCallTagRequest) (_result *DeleteCallTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCallTagResponse{}
	_body, _err := client.DeleteCallTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// -
//
// @param request - DeleteContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactFlowResponse
func (client *Client) DeleteContactFlowWithOptions(request *DeleteContactFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// -
//
// @param request - DeleteContactFlowRequest
//
// @return DeleteContactFlowResponse
func (client *Client) DeleteContactFlow(request *DeleteContactFlowRequest) (_result *DeleteContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContactFlowResponse{}
	_body, _err := client.DeleteContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除呼入控制号码
//
// @param request - DeleteCustomCallTaggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomCallTaggingResponse
func (client *Client) DeleteCustomCallTaggingWithOptions(request *DeleteCustomCallTaggingRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomCallTaggingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomCallTagging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomCallTaggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除呼入控制号码
//
// @param request - DeleteCustomCallTaggingRequest
//
// @return DeleteCustomCallTaggingResponse
func (client *Client) DeleteCustomCallTagging(request *DeleteCustomCallTaggingRequest) (_result *DeleteCustomCallTaggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomCallTaggingResponse{}
	_body, _err := client.DeleteCustomCallTaggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocumentWithOptions(request *DeleteDocumentRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		body["DocumentId"] = request.DocumentId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocument"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteDocumentRequest
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocument(request *DeleteDocumentRequest) (_result *DeleteDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDocumentResponse{}
	_body, _err := client.DeleteDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - DeleteDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentsResponse
func (client *Client) DeleteDocumentsWithOptions(tmpReq *DeleteDocumentsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocumentsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteDocumentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocumentIds) {
		request.DocumentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentIds, dara.String("DocumentIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIdsShrink) {
		body["DocumentIds"] = request.DocumentIdsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocuments"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteDocumentsRequest
//
// @return DeleteDocumentsResponse
func (client *Client) DeleteDocuments(request *DeleteDocumentsRequest) (_result *DeleteDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDocumentsResponse{}
	_body, _err := client.DeleteDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchemaResponse
func (client *Client) DeleteSchemaWithOptions(request *DeleteSchemaRequest, runtime *dara.RuntimeOptions) (_result *DeleteSchemaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSchema"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteSchemaRequest
//
// @return DeleteSchemaResponse
func (client *Client) DeleteSchema(request *DeleteSchemaRequest) (_result *DeleteSchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSchemaResponse{}
	_body, _err := client.DeleteSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteSchemaPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchemaPropertyResponse
func (client *Client) DeleteSchemaPropertyWithOptions(request *DeleteSchemaPropertyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSchemaPropertyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PropertyName) {
		body["PropertyName"] = request.PropertyName
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSchemaProperty"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSchemaPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteSchemaPropertyRequest
//
// @return DeleteSchemaPropertyResponse
func (client *Client) DeleteSchemaProperty(request *DeleteSchemaPropertyRequest) (_result *DeleteSchemaPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSchemaPropertyResponse{}
	_body, _err := client.DeleteSchemaPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillGroupResponse
func (client *Client) DeleteSkillGroupWithOptions(request *DeleteSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteSkillGroupRequest
//
// @return DeleteSkillGroupResponse
func (client *Client) DeleteSkillGroup(request *DeleteSkillGroupRequest) (_result *DeleteSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSkillGroupResponse{}
	_body, _err := client.DeleteSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTicketResponse
func (client *Client) DeleteTicketWithOptions(request *DeleteTicketRequest, runtime *dara.RuntimeOptions) (_result *DeleteTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTicket"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteTicketRequest
//
// @return DeleteTicketResponse
func (client *Client) DeleteTicket(request *DeleteTicketRequest) (_result *DeleteTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTicketResponse{}
	_body, _err := client.DeleteTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteTicketTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTicketTemplateResponse
func (client *Client) DeleteTicketTemplateWithOptions(request *DeleteTicketTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTicketTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTicketTemplate"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTicketTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteTicketTemplateRequest
//
// @return DeleteTicketTemplateResponse
func (client *Client) DeleteTicketTemplate(request *DeleteTicketTemplateRequest) (_result *DeleteTicketTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTicketTemplateResponse{}
	_body, _err := client.DeleteTicketTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableSchemaPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSchemaPropertyResponse
func (client *Client) DisableSchemaPropertyWithOptions(request *DisableSchemaPropertyRequest, runtime *dara.RuntimeOptions) (_result *DisableSchemaPropertyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PropertyName) {
		body["PropertyName"] = request.PropertyName
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSchemaProperty"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSchemaPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableSchemaPropertyRequest
//
// @return DisableSchemaPropertyResponse
func (client *Client) DisableSchemaProperty(request *DisableSchemaPropertyRequest) (_result *DisableSchemaPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableSchemaPropertyResponse{}
	_body, _err := client.DisableSchemaPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableTicketTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableTicketTemplateResponse
func (client *Client) DisableTicketTemplateWithOptions(request *DisableTicketTemplateRequest, runtime *dara.RuntimeOptions) (_result *DisableTicketTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableTicketTemplate"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableTicketTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableTicketTemplateRequest
//
// @return DisableTicketTemplateResponse
func (client *Client) DisableTicketTemplate(request *DisableTicketTemplateRequest) (_result *DisableTicketTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableTicketTemplateResponse{}
	_body, _err := client.DisableTicketTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DiscardEditingContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DiscardEditingContactFlowResponse
func (client *Client) DiscardEditingContactFlowWithOptions(request *DiscardEditingContactFlowRequest, runtime *dara.RuntimeOptions) (_result *DiscardEditingContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.DraftId) {
		query["DraftId"] = request.DraftId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DiscardEditingContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DiscardEditingContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DiscardEditingContactFlowRequest
//
// @return DiscardEditingContactFlowResponse
func (client *Client) DiscardEditingContactFlow(request *DiscardEditingContactFlowRequest) (_result *DiscardEditingContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DiscardEditingContactFlowResponse{}
	_body, _err := client.DiscardEditingContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableSchemaPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSchemaPropertyResponse
func (client *Client) EnableSchemaPropertyWithOptions(request *EnableSchemaPropertyRequest, runtime *dara.RuntimeOptions) (_result *EnableSchemaPropertyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PropertyName) {
		body["PropertyName"] = request.PropertyName
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSchemaProperty"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSchemaPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableSchemaPropertyRequest
//
// @return EnableSchemaPropertyResponse
func (client *Client) EnableSchemaProperty(request *EnableSchemaPropertyRequest) (_result *EnableSchemaPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableSchemaPropertyResponse{}
	_body, _err := client.EnableSchemaPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableTicketTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableTicketTemplateResponse
func (client *Client) EnableTicketTemplateWithOptions(request *EnableTicketTemplateRequest, runtime *dara.RuntimeOptions) (_result *EnableTicketTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableTicketTemplate"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableTicketTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableTicketTemplateRequest
//
// @return EnableTicketTemplateResponse
func (client *Client) EnableTicketTemplate(request *EnableTicketTemplateRequest) (_result *EnableTicketTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableTicketTemplateResponse{}
	_body, _err := client.EnableTicketTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EndConferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndConferenceResponse
func (client *Client) EndConferenceWithOptions(request *EndConferenceRequest, runtime *dara.RuntimeOptions) (_result *EndConferenceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndConference"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EndConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EndConferenceRequest
//
// @return EndConferenceResponse
func (client *Client) EndConference(request *EndConferenceRequest) (_result *EndConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EndConferenceResponse{}
	_body, _err := client.EndConferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ExportContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportContactFlowResponse
func (client *Client) ExportContactFlowWithOptions(request *ExportContactFlowRequest, runtime *dara.RuntimeOptions) (_result *ExportContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		body["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExportContactFlowRequest
//
// @return ExportContactFlowResponse
func (client *Client) ExportContactFlow(request *ExportContactFlowRequest) (_result *ExportContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportContactFlowResponse{}
	_body, _err := client.ExportContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ExportCustomCallTagging is deprecated, please use CCC::2020-07-01::ExportCustomCallTaggings instead.
//
// Summary:
//
// 导出全部呼入号码标签
//
// @param request - ExportCustomCallTaggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportCustomCallTaggingResponse
// Deprecated
func (client *Client) ExportCustomCallTaggingWithOptions(request *ExportCustomCallTaggingRequest, runtime *dara.RuntimeOptions) (_result *ExportCustomCallTaggingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportCustomCallTagging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportCustomCallTaggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ExportCustomCallTagging is deprecated, please use CCC::2020-07-01::ExportCustomCallTaggings instead.
//
// Summary:
//
// 导出全部呼入号码标签
//
// @param request - ExportCustomCallTaggingRequest
//
// @return ExportCustomCallTaggingResponse
// Deprecated
func (client *Client) ExportCustomCallTagging(request *ExportCustomCallTaggingRequest) (_result *ExportCustomCallTaggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportCustomCallTaggingResponse{}
	_body, _err := client.ExportCustomCallTaggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出黑名单号码
//
// @param request - ExportDoNotCallNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportDoNotCallNumbersResponse
func (client *Client) ExportDoNotCallNumbersWithOptions(request *ExportDoNotCallNumbersRequest, runtime *dara.RuntimeOptions) (_result *ExportDoNotCallNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportDoNotCallNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportDoNotCallNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出黑名单号码
//
// @param request - ExportDoNotCallNumbersRequest
//
// @return ExportDoNotCallNumbersResponse
func (client *Client) ExportDoNotCallNumbers(request *ExportDoNotCallNumbersRequest) (_result *ExportDoNotCallNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportDoNotCallNumbersResponse{}
	_body, _err := client.ExportDoNotCallNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FinishTicketTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishTicketTaskResponse
func (client *Client) FinishTicketTaskWithOptions(request *FinishTicketTaskRequest, runtime *dara.RuntimeOptions) (_result *FinishTicketTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FinishTicketTask"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FinishTicketTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FinishTicketTaskRequest
//
// @return FinishTicketTaskResponse
func (client *Client) FinishTicketTask(request *FinishTicketTaskRequest) (_result *FinishTicketTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FinishTicketTaskResponse{}
	_body, _err := client.FinishTicketTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetAccessChannelOfStaging
//
// @param request - GetAccessChannelOfStagingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessChannelOfStagingResponse
func (client *Client) GetAccessChannelOfStagingWithOptions(request *GetAccessChannelOfStagingRequest, runtime *dara.RuntimeOptions) (_result *GetAccessChannelOfStagingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessChannelOfStaging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessChannelOfStagingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetAccessChannelOfStaging
//
// @param request - GetAccessChannelOfStagingRequest
//
// @return GetAccessChannelOfStagingResponse
func (client *Client) GetAccessChannelOfStaging(request *GetAccessChannelOfStagingRequest) (_result *GetAccessChannelOfStagingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessChannelOfStagingResponse{}
	_body, _err := client.GetAccessChannelOfStagingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取音频文件
//
// @param request - GetAudioFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAudioFileResponse
func (client *Client) GetAudioFileWithOptions(request *GetAudioFileRequest, runtime *dara.RuntimeOptions) (_result *GetAudioFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioResourceId) {
		query["AudioResourceId"] = request.AudioResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAudioFile"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAudioFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取音频文件
//
// @param request - GetAudioFileRequest
//
// @return GetAudioFileResponse
func (client *Client) GetAudioFile(request *GetAudioFileRequest) (_result *GetAudioFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAudioFileResponse{}
	_body, _err := client.GetAudioFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAudioFileDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAudioFileDownloadUrlResponse
func (client *Client) GetAudioFileDownloadUrlWithOptions(request *GetAudioFileDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAudioFileDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioResourceId) {
		query["AudioResourceId"] = request.AudioResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAudioFileDownloadUrl"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAudioFileDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAudioFileDownloadUrlRequest
//
// @return GetAudioFileDownloadUrlResponse
func (client *Client) GetAudioFileDownloadUrl(request *GetAudioFileDownloadUrlRequest) (_result *GetAudioFileDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAudioFileDownloadUrlResponse{}
	_body, _err := client.GetAudioFileDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAudioFileUploadParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAudioFileUploadParametersResponse
func (client *Client) GetAudioFileUploadParametersWithOptions(request *GetAudioFileUploadParametersRequest, runtime *dara.RuntimeOptions) (_result *GetAudioFileUploadParametersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioFileName) {
		query["AudioFileName"] = request.AudioFileName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAudioFileUploadParameters"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAudioFileUploadParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAudioFileUploadParametersRequest
//
// @return GetAudioFileUploadParametersResponse
func (client *Client) GetAudioFileUploadParameters(request *GetAudioFileUploadParametersRequest) (_result *GetAudioFileUploadParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAudioFileUploadParametersResponse{}
	_body, _err := client.GetAudioFileUploadParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通话记录详情
//
// @param request - GetCallDetailRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCallDetailRecordResponse
func (client *Client) GetCallDetailRecordWithOptions(request *GetCallDetailRecordRequest, runtime *dara.RuntimeOptions) (_result *GetCallDetailRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCallDetailRecord"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCallDetailRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通话记录详情
//
// @param request - GetCallDetailRecordRequest
//
// @return GetCallDetailRecordResponse
func (client *Client) GetCallDetailRecord(request *GetCallDetailRecordRequest) (_result *GetCallDetailRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCallDetailRecordResponse{}
	_body, _err := client.GetCallDetailRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动信息
//
// @param request - GetCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCampaignResponse
func (client *Client) GetCampaignWithOptions(request *GetCampaignRequest, runtime *dara.RuntimeOptions) (_result *GetCampaignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动信息
//
// @param request - GetCampaignRequest
//
// @return GetCampaignResponse
func (client *Client) GetCampaign(request *GetCampaignRequest) (_result *GetCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCampaignResponse{}
	_body, _err := client.GetCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetCaseFileUploadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCaseFileUploadUrlResponse
func (client *Client) GetCaseFileUploadUrlWithOptions(request *GetCaseFileUploadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetCaseFileUploadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCaseFileUploadUrl"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCaseFileUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetCaseFileUploadUrlRequest
//
// @return GetCaseFileUploadUrlResponse
func (client *Client) GetCaseFileUploadUrl(request *GetCaseFileUploadUrlRequest) (_result *GetCaseFileUploadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCaseFileUploadUrlResponse{}
	_body, _err := client.GetCaseFileUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetChatMediaUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatMediaUrlResponse
func (client *Client) GetChatMediaUrlWithOptions(request *GetChatMediaUrlRequest, runtime *dara.RuntimeOptions) (_result *GetChatMediaUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaId) {
		body["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatMediaUrl"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatMediaUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetChatMediaUrlRequest
//
// @return GetChatMediaUrlResponse
func (client *Client) GetChatMediaUrl(request *GetChatMediaUrlRequest) (_result *GetChatMediaUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatMediaUrlResponse{}
	_body, _err := client.GetChatMediaUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetChatRoutingProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatRoutingProfileResponse
func (client *Client) GetChatRoutingProfileWithOptions(request *GetChatRoutingProfileRequest, runtime *dara.RuntimeOptions) (_result *GetChatRoutingProfileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatRoutingProfile"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatRoutingProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetChatRoutingProfileRequest
//
// @return GetChatRoutingProfileResponse
func (client *Client) GetChatRoutingProfile(request *GetChatRoutingProfileRequest) (_result *GetChatRoutingProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatRoutingProfileResponse{}
	_body, _err := client.GetChatRoutingProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContactFlowResponse
func (client *Client) GetContactFlowWithOptions(request *GetContactFlowRequest, runtime *dara.RuntimeOptions) (_result *GetContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.DraftId) {
		query["DraftId"] = request.DraftId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetContactFlowRequest
//
// @return GetContactFlowResponse
func (client *Client) GetContactFlow(request *GetContactFlowRequest) (_result *GetContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContactFlowResponse{}
	_body, _err := client.GetContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取通话文本信息
//
// @param request - GetConversationDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversationDetailResponse
func (client *Client) GetConversationDetailWithOptions(request *GetConversationDetailRequest, runtime *dara.RuntimeOptions) (_result *GetConversationDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConversationDetail"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConversationDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取通话文本信息
//
// @param request - GetConversationDetailRequest
//
// @return GetConversationDetailResponse
func (client *Client) GetConversationDetail(request *GetConversationDetailRequest) (_result *GetConversationDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConversationDetailResponse{}
	_body, _err := client.GetConversationDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetDataChannelCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataChannelCredentialsResponse
func (client *Client) GetDataChannelCredentialsWithOptions(request *GetDataChannelCredentialsRequest, runtime *dara.RuntimeOptions) (_result *GetDataChannelCredentialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataChannelCredentials"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataChannelCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDataChannelCredentialsRequest
//
// @return GetDataChannelCredentialsResponse
func (client *Client) GetDataChannelCredentials(request *GetDataChannelCredentialsRequest) (_result *GetDataChannelCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataChannelCredentialsResponse{}
	_body, _err := client.GetDataChannelCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取黑名单文件上传地址
//
// @param request - GetDoNotCallFileUploadParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoNotCallFileUploadParametersResponse
func (client *Client) GetDoNotCallFileUploadParametersWithOptions(request *GetDoNotCallFileUploadParametersRequest, runtime *dara.RuntimeOptions) (_result *GetDoNotCallFileUploadParametersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoNotCallFileUploadParameters"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoNotCallFileUploadParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取黑名单文件上传地址
//
// @param request - GetDoNotCallFileUploadParametersRequest
//
// @return GetDoNotCallFileUploadParametersResponse
func (client *Client) GetDoNotCallFileUploadParameters(request *GetDoNotCallFileUploadParametersRequest) (_result *GetDoNotCallFileUploadParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoNotCallFileUploadParametersResponse{}
	_body, _err := client.GetDoNotCallFileUploadParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetDocumentUploadParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentUploadParametersResponse
func (client *Client) GetDocumentUploadParametersWithOptions(request *GetDocumentUploadParametersRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentUploadParametersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentUploadParameters"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentUploadParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDocumentUploadParametersRequest
//
// @return GetDocumentUploadParametersResponse
func (client *Client) GetDocumentUploadParameters(request *GetDocumentUploadParametersRequest) (_result *GetDocumentUploadParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocumentUploadParametersResponse{}
	_body, _err := client.GetDocumentUploadParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取早媒体音频
//
// @param request - GetEarlyMediaRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEarlyMediaRecordingResponse
func (client *Client) GetEarlyMediaRecordingWithOptions(request *GetEarlyMediaRecordingRequest, runtime *dara.RuntimeOptions) (_result *GetEarlyMediaRecordingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEarlyMediaRecording"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEarlyMediaRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取早媒体音频
//
// @param request - GetEarlyMediaRecordingRequest
//
// @return GetEarlyMediaRecordingResponse
func (client *Client) GetEarlyMediaRecording(request *GetEarlyMediaRecordingRequest) (_result *GetEarlyMediaRecordingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEarlyMediaRecordingResponse{}
	_body, _err := client.GetEarlyMediaRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetHistoricalCallerReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoricalCallerReportResponse
func (client *Client) GetHistoricalCallerReportWithOptions(request *GetHistoricalCallerReportRequest, runtime *dara.RuntimeOptions) (_result *GetHistoricalCallerReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StopTime) {
		query["StopTime"] = request.StopTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHistoricalCallerReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHistoricalCallerReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetHistoricalCallerReportRequest
//
// @return GetHistoricalCallerReportResponse
func (client *Client) GetHistoricalCallerReport(request *GetHistoricalCallerReportRequest) (_result *GetHistoricalCallerReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHistoricalCallerReportResponse{}
	_body, _err := client.GetHistoricalCallerReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动历史报表
//
// @param request - GetHistoricalCampaignReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoricalCampaignReportResponse
func (client *Client) GetHistoricalCampaignReportWithOptions(request *GetHistoricalCampaignReportRequest, runtime *dara.RuntimeOptions) (_result *GetHistoricalCampaignReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHistoricalCampaignReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHistoricalCampaignReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动历史报表
//
// @param request - GetHistoricalCampaignReportRequest
//
// @return GetHistoricalCampaignReportResponse
func (client *Client) GetHistoricalCampaignReport(request *GetHistoricalCampaignReportRequest) (_result *GetHistoricalCampaignReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHistoricalCampaignReportResponse{}
	_body, _err := client.GetHistoricalCampaignReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetHistoricalInstanceReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoricalInstanceReportResponse
func (client *Client) GetHistoricalInstanceReportWithOptions(request *GetHistoricalInstanceReportRequest, runtime *dara.RuntimeOptions) (_result *GetHistoricalInstanceReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHistoricalInstanceReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHistoricalInstanceReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetHistoricalInstanceReportRequest
//
// @return GetHistoricalInstanceReportResponse
func (client *Client) GetHistoricalInstanceReport(request *GetHistoricalInstanceReportRequest) (_result *GetHistoricalInstanceReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHistoricalInstanceReportResponse{}
	_body, _err := client.GetHistoricalInstanceReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetInstanceTrendingReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceTrendingReportResponse
func (client *Client) GetInstanceTrendingReportWithOptions(request *GetInstanceTrendingReportRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceTrendingReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceTrendingReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceTrendingReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetInstanceTrendingReportRequest
//
// @return GetInstanceTrendingReportResponse
func (client *Client) GetInstanceTrendingReport(request *GetInstanceTrendingReportRequest) (_result *GetInstanceTrendingReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceTrendingReportResponse{}
	_body, _err := client.GetInstanceTrendingReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetIvrTrackingSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIvrTrackingSummaryResponse
func (client *Client) GetIvrTrackingSummaryWithOptions(request *GetIvrTrackingSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetIvrTrackingSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIvrTrackingSummary"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIvrTrackingSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetIvrTrackingSummaryRequest
//
// @return GetIvrTrackingSummaryResponse
func (client *Client) GetIvrTrackingSummary(request *GetIvrTrackingSummaryRequest) (_result *GetIvrTrackingSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIvrTrackingSummaryResponse{}
	_body, _err := client.GetIvrTrackingSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLoginDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginDetailsResponse
func (client *Client) GetLoginDetailsWithOptions(request *GetLoginDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetLoginDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatDeviceId) {
		query["ChatDeviceId"] = request.ChatDeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoginDetails"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoginDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetLoginDetailsRequest
//
// @return GetLoginDetailsResponse
func (client *Client) GetLoginDetails(request *GetLoginDetailsRequest) (_result *GetLoginDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLoginDetailsResponse{}
	_body, _err := client.GetLoginDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetMonoRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonoRecordingResponse
func (client *Client) GetMonoRecordingWithOptions(request *GetMonoRecordingRequest, runtime *dara.RuntimeOptions) (_result *GetMonoRecordingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ExpireSeconds) {
		query["ExpireSeconds"] = request.ExpireSeconds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMonoRecording"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMonoRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMonoRecordingRequest
//
// @return GetMonoRecordingResponse
func (client *Client) GetMonoRecording(request *GetMonoRecordingRequest) (_result *GetMonoRecordingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMonoRecordingResponse{}
	_body, _err := client.GetMonoRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetMultiChannelRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiChannelRecordingResponse
func (client *Client) GetMultiChannelRecordingWithOptions(request *GetMultiChannelRecordingRequest, runtime *dara.RuntimeOptions) (_result *GetMultiChannelRecordingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiChannelRecording"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiChannelRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMultiChannelRecordingRequest
//
// @return GetMultiChannelRecordingResponse
func (client *Client) GetMultiChannelRecording(request *GetMultiChannelRecordingRequest) (_result *GetMultiChannelRecordingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMultiChannelRecordingResponse{}
	_body, _err := client.GetMultiChannelRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetNumberLocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNumberLocationResponse
func (client *Client) GetNumberLocationWithOptions(request *GetNumberLocationRequest, runtime *dara.RuntimeOptions) (_result *GetNumberLocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNumberLocation"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNumberLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetNumberLocationRequest
//
// @return GetNumberLocationResponse
func (client *Client) GetNumberLocation(request *GetNumberLocationRequest) (_result *GetNumberLocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNumberLocationResponse{}
	_body, _err := client.GetNumberLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取预测式外呼实时状态
//
// @param request - GetRealtimeCampaignStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealtimeCampaignStatsResponse
func (client *Client) GetRealtimeCampaignStatsWithOptions(request *GetRealtimeCampaignStatsRequest, runtime *dara.RuntimeOptions) (_result *GetRealtimeCampaignStatsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRealtimeCampaignStats"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRealtimeCampaignStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼实时状态
//
// @param request - GetRealtimeCampaignStatsRequest
//
// @return GetRealtimeCampaignStatsResponse
func (client *Client) GetRealtimeCampaignStats(request *GetRealtimeCampaignStatsRequest) (_result *GetRealtimeCampaignStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRealtimeCampaignStatsResponse{}
	_body, _err := client.GetRealtimeCampaignStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRealtimeInstanceStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealtimeInstanceStatesResponse
func (client *Client) GetRealtimeInstanceStatesWithOptions(request *GetRealtimeInstanceStatesRequest, runtime *dara.RuntimeOptions) (_result *GetRealtimeInstanceStatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRealtimeInstanceStates"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRealtimeInstanceStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRealtimeInstanceStatesRequest
//
// @return GetRealtimeInstanceStatesResponse
func (client *Client) GetRealtimeInstanceStates(request *GetRealtimeInstanceStatesRequest) (_result *GetRealtimeInstanceStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRealtimeInstanceStatesResponse{}
	_body, _err := client.GetRealtimeInstanceStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemaResponse
func (client *Client) GetSchemaWithOptions(request *GetSchemaRequest, runtime *dara.RuntimeOptions) (_result *GetSchemaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSchema"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetSchemaRequest
//
// @return GetSchemaResponse
func (client *Client) GetSchema(request *GetSchemaRequest) (_result *GetSchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSchemaResponse{}
	_body, _err := client.GetSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询技能组
//
// @param request - GetSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillGroupResponse
func (client *Client) GetSkillGroupWithOptions(request *GetSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *GetSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询技能组
//
// @param request - GetSkillGroupRequest
//
// @return GetSkillGroupResponse
func (client *Client) GetSkillGroup(request *GetSkillGroupRequest) (_result *GetSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillGroupResponse{}
	_body, _err := client.GetSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetSummaryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSummaryTemplateResponse
func (client *Client) GetSummaryTemplateWithOptions(request *GetSummaryTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetSummaryTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSummaryTemplate"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSummaryTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetSummaryTemplateRequest
//
// @return GetSummaryTemplateResponse
func (client *Client) GetSummaryTemplate(request *GetSummaryTemplateRequest) (_result *GetSummaryTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSummaryTemplateResponse{}
	_body, _err := client.GetSummaryTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketResponse
func (client *Client) GetTicketWithOptions(request *GetTicketRequest, runtime *dara.RuntimeOptions) (_result *GetTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTicket"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetTicketRequest
//
// @return GetTicketResponse
func (client *Client) GetTicket(request *GetTicketRequest) (_result *GetTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTicketResponse{}
	_body, _err := client.GetTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetTicketSummaryReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketSummaryReportResponse
func (client *Client) GetTicketSummaryReportWithOptions(request *GetTicketSummaryReportRequest, runtime *dara.RuntimeOptions) (_result *GetTicketSummaryReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Assignee) {
		query["Assignee"] = request.Assignee
	}

	if !dara.IsNil(request.AssigneeType) {
		query["AssigneeType"] = request.AssigneeType
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Participant) {
		query["Participant"] = request.Participant
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTicketSummaryReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTicketSummaryReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetTicketSummaryReportRequest
//
// @return GetTicketSummaryReportResponse
func (client *Client) GetTicketSummaryReport(request *GetTicketSummaryReportRequest) (_result *GetTicketSummaryReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTicketSummaryReportResponse{}
	_body, _err := client.GetTicketSummaryReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetTicketTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketTemplateResponse
func (client *Client) GetTicketTemplateWithOptions(request *GetTicketTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetTicketTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTicketTemplate"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTicketTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetTicketTemplateRequest
//
// @return GetTicketTemplateResponse
func (client *Client) GetTicketTemplate(request *GetTicketTemplateRequest) (_result *GetTicketTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTicketTemplateResponse{}
	_body, _err := client.GetTicketTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetTurnCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTurnCredentialsResponse
func (client *Client) GetTurnCredentialsWithOptions(request *GetTurnCredentialsRequest, runtime *dara.RuntimeOptions) (_result *GetTurnCredentialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTurnCredentials"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTurnCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetTurnCredentialsRequest
//
// @return GetTurnCredentialsResponse
func (client *Client) GetTurnCredentials(request *GetTurnCredentialsRequest) (_result *GetTurnCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTurnCredentialsResponse{}
	_body, _err := client.GetTurnCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetTurnServerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTurnServerListResponse
func (client *Client) GetTurnServerListWithOptions(request *GetTurnServerListRequest, runtime *dara.RuntimeOptions) (_result *GetTurnServerListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTurnServerList"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTurnServerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetTurnServerListRequest
//
// @return GetTurnServerListResponse
func (client *Client) GetTurnServerList(request *GetTurnServerListRequest) (_result *GetTurnServerListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTurnServerListResponse{}
	_body, _err := client.GetTurnServerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质检参数
//
// @param request - GetUploadAudioDataParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadAudioDataParamsResponse
func (client *Client) GetUploadAudioDataParamsWithOptions(request *GetUploadAudioDataParamsRequest, runtime *dara.RuntimeOptions) (_result *GetUploadAudioDataParamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadAudioDataParams"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadAudioDataParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质检参数
//
// @param request - GetUploadAudioDataParamsRequest
//
// @return GetUploadAudioDataParamsResponse
func (client *Client) GetUploadAudioDataParams(request *GetUploadAudioDataParamsRequest) (_result *GetUploadAudioDataParamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUploadAudioDataParamsResponse{}
	_body, _err := client.GetUploadAudioDataParamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Extension) {
		query["Extension"] = request.Extension
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取视频
//
// @param request - GetVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoResponse
func (client *Client) GetVideoWithOptions(request *GetVideoRequest, runtime *dara.RuntimeOptions) (_result *GetVideoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideo"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频
//
// @param request - GetVideoRequest
//
// @return GetVideoResponse
func (client *Client) GetVideo(request *GetVideoRequest) (_result *GetVideoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoResponse{}
	_body, _err := client.GetVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetVisitorLoginDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVisitorLoginDetailsResponse
func (client *Client) GetVisitorLoginDetailsWithOptions(request *GetVisitorLoginDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetVisitorLoginDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatDeviceId) {
		query["ChatDeviceId"] = request.ChatDeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VisitorId) {
		query["VisitorId"] = request.VisitorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVisitorLoginDetails"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVisitorLoginDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetVisitorLoginDetailsRequest
//
// @return GetVisitorLoginDetailsResponse
func (client *Client) GetVisitorLoginDetails(request *GetVisitorLoginDetailsRequest) (_result *GetVisitorLoginDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVisitorLoginDetailsResponse{}
	_body, _err := client.GetVisitorLoginDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetVoicemailRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVoicemailRecordingResponse
func (client *Client) GetVoicemailRecordingWithOptions(request *GetVoicemailRecordingRequest, runtime *dara.RuntimeOptions) (_result *GetVoicemailRecordingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVoicemailRecording"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVoicemailRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetVoicemailRecordingRequest
//
// @return GetVoicemailRecordingResponse
func (client *Client) GetVoicemailRecording(request *GetVoicemailRecordingRequest) (_result *GetVoicemailRecordingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVoicemailRecordingResponse{}
	_body, _err := client.GetVoicemailRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - HoldCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HoldCallResponse
func (client *Client) HoldCallWithOptions(request *HoldCallRequest, runtime *dara.RuntimeOptions) (_result *HoldCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Music) {
		query["Music"] = request.Music
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HoldCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HoldCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - HoldCallRequest
//
// @return HoldCallResponse
func (client *Client) HoldCall(request *HoldCallRequest) (_result *HoldCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &HoldCallResponse{}
	_body, _err := client.HoldCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ImportAdminsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportAdminsResponse
func (client *Client) ImportAdminsWithOptions(request *ImportAdminsRequest, runtime *dara.RuntimeOptions) (_result *ImportAdminsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RamIdList) {
		query["RamIdList"] = request.RamIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportAdmins"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportAdminsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ImportAdminsRequest
//
// @return ImportAdminsResponse
func (client *Client) ImportAdmins(request *ImportAdminsRequest) (_result *ImportAdminsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportAdminsResponse{}
	_body, _err := client.ImportAdminsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ImportContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportContactFlowResponse
func (client *Client) ImportContactFlowWithOptions(request *ImportContactFlowRequest, runtime *dara.RuntimeOptions) (_result *ImportContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FlowPackageData) {
		body["FlowPackageData"] = request.FlowPackageData
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ImportContactFlowRequest
//
// @return ImportContactFlowResponse
func (client *Client) ImportContactFlow(request *ImportContactFlowRequest) (_result *ImportContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportContactFlowResponse{}
	_body, _err := client.ImportContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ImportCorpNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportCorpNumbersResponse
func (client *Client) ImportCorpNumbersWithOptions(request *ImportCorpNumbersRequest, runtime *dara.RuntimeOptions) (_result *ImportCorpNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.CorpName) {
		query["CorpName"] = request.CorpName
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.Provider) {
		query["Provider"] = request.Provider
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.TagList) {
		query["TagList"] = request.TagList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportCorpNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportCorpNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ImportCorpNumbersRequest
//
// @return ImportCorpNumbersResponse
func (client *Client) ImportCorpNumbers(request *ImportCorpNumbersRequest) (_result *ImportCorpNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportCorpNumbersResponse{}
	_body, _err := client.ImportCorpNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ImportCustomCallTagging is deprecated, please use CCC::2020-07-01::ImportCustomCallTaggings instead.
//
// Summary:
//
// 文件导入呼入控制号码
//
// @param request - ImportCustomCallTaggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportCustomCallTaggingResponse
// Deprecated
func (client *Client) ImportCustomCallTaggingWithOptions(request *ImportCustomCallTaggingRequest, runtime *dara.RuntimeOptions) (_result *ImportCustomCallTaggingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportCustomCallTagging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportCustomCallTaggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ImportCustomCallTagging is deprecated, please use CCC::2020-07-01::ImportCustomCallTaggings instead.
//
// Summary:
//
// 文件导入呼入控制号码
//
// @param request - ImportCustomCallTaggingRequest
//
// @return ImportCustomCallTaggingResponse
// Deprecated
func (client *Client) ImportCustomCallTagging(request *ImportCustomCallTaggingRequest) (_result *ImportCustomCallTaggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportCustomCallTaggingResponse{}
	_body, _err := client.ImportCustomCallTaggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加黑名单号码
//
// @param request - ImportDoNotCallNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportDoNotCallNumbersResponse
func (client *Client) ImportDoNotCallNumbersWithOptions(request *ImportDoNotCallNumbersRequest, runtime *dara.RuntimeOptions) (_result *ImportDoNotCallNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportDoNotCallNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportDoNotCallNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加黑名单号码
//
// @param request - ImportDoNotCallNumbersRequest
//
// @return ImportDoNotCallNumbersResponse
func (client *Client) ImportDoNotCallNumbers(request *ImportDoNotCallNumbersRequest) (_result *ImportDoNotCallNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportDoNotCallNumbersResponse{}
	_body, _err := client.ImportDoNotCallNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ImportDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportDocumentsResponse
func (client *Client) ImportDocumentsWithOptions(request *ImportDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ImportDocumentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OssFileKey) {
		body["OssFileKey"] = request.OssFileKey
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportDocuments"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ImportDocumentsRequest
//
// @return ImportDocumentsResponse
func (client *Client) ImportDocuments(request *ImportDocumentsRequest) (_result *ImportDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportDocumentsResponse{}
	_body, _err := client.ImportDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - InitiateAttendedTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitiateAttendedTransferResponse
func (client *Client) InitiateAttendedTransferWithOptions(request *InitiateAttendedTransferRequest, runtime *dara.RuntimeOptions) (_result *InitiateAttendedTransferResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallPriority) {
		query["CallPriority"] = request.CallPriority
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.QueuingOverflowThreshold) {
		query["QueuingOverflowThreshold"] = request.QueuingOverflowThreshold
	}

	if !dara.IsNil(request.QueuingTimeoutSeconds) {
		query["QueuingTimeoutSeconds"] = request.QueuingTimeoutSeconds
	}

	if !dara.IsNil(request.RoutingType) {
		query["RoutingType"] = request.RoutingType
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	if !dara.IsNil(request.StrategyParams) {
		query["StrategyParams"] = request.StrategyParams
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.Transferee) {
		query["Transferee"] = request.Transferee
	}

	if !dara.IsNil(request.TransfereeType) {
		query["TransfereeType"] = request.TransfereeType
	}

	if !dara.IsNil(request.Transferor) {
		query["Transferor"] = request.Transferor
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitiateAttendedTransfer"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitiateAttendedTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - InitiateAttendedTransferRequest
//
// @return InitiateAttendedTransferResponse
func (client *Client) InitiateAttendedTransfer(request *InitiateAttendedTransferRequest) (_result *InitiateAttendedTransferResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitiateAttendedTransferResponse{}
	_body, _err := client.InitiateAttendedTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - InterceptCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InterceptCallResponse
func (client *Client) InterceptCallWithOptions(request *InterceptCallRequest, runtime *dara.RuntimeOptions) (_result *InterceptCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InterceptedUserId) {
		query["InterceptedUserId"] = request.InterceptedUserId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InterceptCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InterceptCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - InterceptCallRequest
//
// @return InterceptCallResponse
func (client *Client) InterceptCall(request *InterceptCallRequest) (_result *InterceptCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InterceptCallResponse{}
	_body, _err := client.InterceptCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - LaunchAuthenticationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LaunchAuthenticationResponse
func (client *Client) LaunchAuthenticationWithOptions(request *LaunchAuthenticationRequest, runtime *dara.RuntimeOptions) (_result *LaunchAuthenticationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.ContactFlowVariables) {
		query["ContactFlowVariables"] = request.ContactFlowVariables
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LaunchAuthentication"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LaunchAuthenticationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - LaunchAuthenticationRequest
//
// @return LaunchAuthenticationResponse
func (client *Client) LaunchAuthentication(request *LaunchAuthenticationRequest) (_result *LaunchAuthenticationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LaunchAuthenticationResponse{}
	_body, _err := client.LaunchAuthenticationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - LaunchSurveyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LaunchSurveyResponse
func (client *Client) LaunchSurveyWithOptions(request *LaunchSurveyRequest, runtime *dara.RuntimeOptions) (_result *LaunchSurveyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.ContactFlowVariables) {
		query["ContactFlowVariables"] = request.ContactFlowVariables
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.SmsMetadataId) {
		query["SmsMetadataId"] = request.SmsMetadataId
	}

	if !dara.IsNil(request.SurveyChannel) {
		query["SurveyChannel"] = request.SurveyChannel
	}

	if !dara.IsNil(request.SurveyTemplateId) {
		query["SurveyTemplateId"] = request.SurveyTemplateId
	}

	if !dara.IsNil(request.SurveyTemplateVariables) {
		query["SurveyTemplateVariables"] = request.SurveyTemplateVariables
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LaunchSurvey"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LaunchSurveyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - LaunchSurveyRequest
//
// @return LaunchSurveyResponse
func (client *Client) LaunchSurvey(request *LaunchSurveyRequest) (_result *LaunchSurveyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LaunchSurveyResponse{}
	_body, _err := client.LaunchSurveyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAgentStateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentStateLogsResponse
func (client *Client) ListAgentStateLogsWithOptions(request *ListAgentStateLogsRequest, runtime *dara.RuntimeOptions) (_result *ListAgentStateLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentStateLogs"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentStateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAgentStateLogsRequest
//
// @return ListAgentStateLogsResponse
func (client *Client) ListAgentStateLogs(request *ListAgentStateLogsRequest) (_result *ListAgentStateLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAgentStateLogsResponse{}
	_body, _err := client.ListAgentStateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListAgentStates is deprecated, please use CCC::2020-07-01::ListRealtimeAgentStates instead.
//
// Summary:
//
// # ListAgentStates for ACC
//
// @param request - ListAgentStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentStatesResponse
// Deprecated
func (client *Client) ListAgentStatesWithOptions(request *ListAgentStatesRequest, runtime *dara.RuntimeOptions) (_result *ListAgentStatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentIds) {
		query["AgentIds"] = request.AgentIds
	}

	if !dara.IsNil(request.ExcludeOfflineUsers) {
		query["ExcludeOfflineUsers"] = request.ExcludeOfflineUsers
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentStates"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListAgentStates is deprecated, please use CCC::2020-07-01::ListRealtimeAgentStates instead.
//
// Summary:
//
// # ListAgentStates for ACC
//
// @param request - ListAgentStatesRequest
//
// @return ListAgentStatesResponse
// Deprecated
func (client *Client) ListAgentStates(request *ListAgentStatesRequest) (_result *ListAgentStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAgentStatesResponse{}
	_body, _err := client.ListAgentStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListAgentSummaryReportsSinceMidnight is deprecated, please use CCC::2020-07-01::ListHistoricalAgentReport instead.
//
// Summary:
//
// # ListAgentSummaryReportsSinceMidnight for acc
//
// @param request - ListAgentSummaryReportsSinceMidnightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentSummaryReportsSinceMidnightResponse
// Deprecated
func (client *Client) ListAgentSummaryReportsSinceMidnightWithOptions(request *ListAgentSummaryReportsSinceMidnightRequest, runtime *dara.RuntimeOptions) (_result *ListAgentSummaryReportsSinceMidnightResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentSummaryReportsSinceMidnight"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentSummaryReportsSinceMidnightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListAgentSummaryReportsSinceMidnight is deprecated, please use CCC::2020-07-01::ListHistoricalAgentReport instead.
//
// Summary:
//
// # ListAgentSummaryReportsSinceMidnight for acc
//
// @param request - ListAgentSummaryReportsSinceMidnightRequest
//
// @return ListAgentSummaryReportsSinceMidnightResponse
// Deprecated
func (client *Client) ListAgentSummaryReportsSinceMidnight(request *ListAgentSummaryReportsSinceMidnightRequest) (_result *ListAgentSummaryReportsSinceMidnightResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAgentSummaryReportsSinceMidnightResponse{}
	_body, _err := client.ListAgentSummaryReportsSinceMidnightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取预测式外呼呼叫记录
//
// @param request - ListAttemptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAttemptsResponse
func (client *Client) ListAttemptsWithOptions(request *ListAttemptsRequest, runtime *dara.RuntimeOptions) (_result *ListAttemptsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAttempts"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAttemptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼呼叫记录
//
// @param request - ListAttemptsRequest
//
// @return ListAttemptsResponse
func (client *Client) ListAttempts(request *ListAttemptsRequest) (_result *ListAttemptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAttemptsResponse{}
	_body, _err := client.ListAttemptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取音频文件列表
//
// @param request - ListAudioFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAudioFilesResponse
func (client *Client) ListAudioFilesWithOptions(request *ListAudioFilesRequest, runtime *dara.RuntimeOptions) (_result *ListAudioFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Usage) {
		query["Usage"] = request.Usage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAudioFiles"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAudioFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取音频文件列表
//
// @param request - ListAudioFilesRequest
//
// @return ListAudioFilesResponse
func (client *Client) ListAudioFiles(request *ListAudioFilesRequest) (_result *ListAudioFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAudioFilesResponse{}
	_body, _err := client.ListAudioFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑呼入控制号码
//
// @param request - ListBlacklistCallTaggingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBlacklistCallTaggingsResponse
func (client *Client) ListBlacklistCallTaggingsWithOptions(request *ListBlacklistCallTaggingsRequest, runtime *dara.RuntimeOptions) (_result *ListBlacklistCallTaggingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBlacklistCallTaggings"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBlacklistCallTaggingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑呼入控制号码
//
// @param request - ListBlacklistCallTaggingsRequest
//
// @return ListBlacklistCallTaggingsResponse
func (client *Client) ListBlacklistCallTaggings(request *ListBlacklistCallTaggingsRequest) (_result *ListBlacklistCallTaggingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBlacklistCallTaggingsResponse{}
	_body, _err := client.ListBlacklistCallTaggingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBriefSkillGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBriefSkillGroupsResponse
func (client *Client) ListBriefSkillGroupsWithOptions(request *ListBriefSkillGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListBriefSkillGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBriefSkillGroups"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBriefSkillGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBriefSkillGroupsRequest
//
// @return ListBriefSkillGroupsResponse
func (client *Client) ListBriefSkillGroups(request *ListBriefSkillGroupsRequest) (_result *ListBriefSkillGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBriefSkillGroupsResponse{}
	_body, _err := client.ListBriefSkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListCallDetailRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallDetailRecordsResponse
func (client *Client) ListCallDetailRecordsWithOptions(request *ListCallDetailRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListCallDetailRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.ContactDisposition) {
		query["ContactDisposition"] = request.ContactDisposition
	}

	if !dara.IsNil(request.ContactDispositionList) {
		query["ContactDispositionList"] = request.ContactDispositionList
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactType) {
		query["ContactType"] = request.ContactType
	}

	if !dara.IsNil(request.ContactTypeList) {
		query["ContactTypeList"] = request.ContactTypeList
	}

	if !dara.IsNil(request.Criteria) {
		query["Criteria"] = request.Criteria
	}

	if !dara.IsNil(request.EarlyMediaStateList) {
		query["EarlyMediaStateList"] = request.EarlyMediaStateList
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderByField) {
		query["OrderByField"] = request.OrderByField
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SatisfactionDescriptionList) {
		query["SatisfactionDescriptionList"] = request.SatisfactionDescriptionList
	}

	if !dara.IsNil(request.SatisfactionList) {
		query["SatisfactionList"] = request.SatisfactionList
	}

	if !dara.IsNil(request.SatisfactionSurveyChannel) {
		query["SatisfactionSurveyChannel"] = request.SatisfactionSurveyChannel
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCallDetailRecords"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCallDetailRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCallDetailRecordsRequest
//
// @return ListCallDetailRecordsResponse
func (client *Client) ListCallDetailRecords(request *ListCallDetailRecordsRequest) (_result *ListCallDetailRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCallDetailRecordsResponse{}
	_body, _err := client.ListCallDetailRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通话记录列表
//
// @param request - ListCallDetailRecordsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallDetailRecordsV2Response
func (client *Client) ListCallDetailRecordsV2WithOptions(request *ListCallDetailRecordsV2Request, runtime *dara.RuntimeOptions) (_result *ListCallDetailRecordsV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessChannelTypeList) {
		query["AccessChannelTypeList"] = request.AccessChannelTypeList
	}

	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AnalyticsReportReady) {
		query["AnalyticsReportReady"] = request.AnalyticsReportReady
	}

	if !dara.IsNil(request.Broker) {
		query["Broker"] = request.Broker
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.ContactDispositionList) {
		query["ContactDispositionList"] = request.ContactDispositionList
	}

	if !dara.IsNil(request.ContactIdList) {
		query["ContactIdList"] = request.ContactIdList
	}

	if !dara.IsNil(request.ContactTypeList) {
		query["ContactTypeList"] = request.ContactTypeList
	}

	if !dara.IsNil(request.EarlyMediaStateList) {
		query["EarlyMediaStateList"] = request.EarlyMediaStateList
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FirstAgentId) {
		query["FirstAgentId"] = request.FirstAgentId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.OrderByField) {
		query["OrderByField"] = request.OrderByField
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReleaseInitiatorList) {
		query["ReleaseInitiatorList"] = request.ReleaseInitiatorList
	}

	if !dara.IsNil(request.ReleaseReasonList) {
		query["ReleaseReasonList"] = request.ReleaseReasonList
	}

	if !dara.IsNil(request.SatisfactionDescriptionList) {
		query["SatisfactionDescriptionList"] = request.SatisfactionDescriptionList
	}

	if !dara.IsNil(request.SatisfactionRateList) {
		query["SatisfactionRateList"] = request.SatisfactionRateList
	}

	if !dara.IsNil(request.SatisfactionSurveyChannel) {
		query["SatisfactionSurveyChannel"] = request.SatisfactionSurveyChannel
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCallDetailRecordsV2"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCallDetailRecordsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通话记录列表
//
// @param request - ListCallDetailRecordsV2Request
//
// @return ListCallDetailRecordsV2Response
func (client *Client) ListCallDetailRecordsV2(request *ListCallDetailRecordsV2Request) (_result *ListCallDetailRecordsV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCallDetailRecordsV2Response{}
	_body, _err := client.ListCallDetailRecordsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - ListCallSummariesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallSummariesResponse
func (client *Client) ListCallSummariesWithOptions(tmpReq *ListCallSummariesRequest, runtime *dara.RuntimeOptions) (_result *ListCallSummariesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListCallSummariesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactIdList) {
		request.ContactIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactIdList, dara.String("ContactIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactIdListShrink) {
		query["ContactIdList"] = request.ContactIdListShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCallSummaries"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCallSummariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCallSummariesRequest
//
// @return ListCallSummariesResponse
func (client *Client) ListCallSummaries(request *ListCallSummariesRequest) (_result *ListCallSummariesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCallSummariesResponse{}
	_body, _err := client.ListCallSummariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出号码标签
//
// @param request - ListCallTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallTagsResponse
func (client *Client) ListCallTagsWithOptions(request *ListCallTagsRequest, runtime *dara.RuntimeOptions) (_result *ListCallTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCallTags"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCallTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出号码标签
//
// @param request - ListCallTagsRequest
//
// @return ListCallTagsResponse
func (client *Client) ListCallTags(request *ListCallTagsRequest) (_result *ListCallTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCallTagsResponse{}
	_body, _err := client.ListCallTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动趋势报表
//
// @param request - ListCampaignTrendingReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCampaignTrendingReportResponse
func (client *Client) ListCampaignTrendingReportWithOptions(request *ListCampaignTrendingReportRequest, runtime *dara.RuntimeOptions) (_result *ListCampaignTrendingReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCampaignTrendingReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCampaignTrendingReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动趋势报表
//
// @param request - ListCampaignTrendingReportRequest
//
// @return ListCampaignTrendingReportResponse
func (client *Client) ListCampaignTrendingReport(request *ListCampaignTrendingReportRequest) (_result *ListCampaignTrendingReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCampaignTrendingReportResponse{}
	_body, _err := client.ListCampaignTrendingReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动列表
//
// @param request - ListCampaignsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCampaignsResponse
func (client *Client) ListCampaignsWithOptions(request *ListCampaignsRequest, runtime *dara.RuntimeOptions) (_result *ListCampaignsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActualStartTimeFrom) {
		query["ActualStartTimeFrom"] = request.ActualStartTimeFrom
	}

	if !dara.IsNil(request.ActualStartTimeTo) {
		query["ActualStartTimeTo"] = request.ActualStartTimeTo
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlanedStartTimeFrom) {
		query["PlanedStartTimeFrom"] = request.PlanedStartTimeFrom
	}

	if !dara.IsNil(request.PlanedStartTimeTo) {
		query["PlanedStartTimeTo"] = request.PlanedStartTimeTo
	}

	if !dara.IsNil(request.QueueId) {
		query["QueueId"] = request.QueueId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCampaigns"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCampaignsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动列表
//
// @param request - ListCampaignsRequest
//
// @return ListCampaignsResponse
func (client *Client) ListCampaigns(request *ListCampaignsRequest) (_result *ListCampaignsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCampaignsResponse{}
	_body, _err := client.ListCampaignsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动的联系人呼叫详情
//
// @param request - ListCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCasesResponse
func (client *Client) ListCasesWithOptions(request *ListCasesRequest, runtime *dara.RuntimeOptions) (_result *ListCasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCases"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动的联系人呼叫详情
//
// @param request - ListCasesRequest
//
// @return ListCasesResponse
func (client *Client) ListCases(request *ListCasesRequest) (_result *ListCasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCasesResponse{}
	_body, _err := client.ListCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoriesResponse
func (client *Client) ListCategoriesWithOptions(request *ListCategoriesRequest, runtime *dara.RuntimeOptions) (_result *ListCategoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCategories"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCategoriesRequest
//
// @return ListCategoriesResponse
func (client *Client) ListCategories(request *ListCategoriesRequest) (_result *ListCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCategoriesResponse{}
	_body, _err := client.ListCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListCommonTicketFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommonTicketFieldsResponse
func (client *Client) ListCommonTicketFieldsWithOptions(request *ListCommonTicketFieldsRequest, runtime *dara.RuntimeOptions) (_result *ListCommonTicketFieldsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCommonTicketFields"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCommonTicketFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCommonTicketFieldsRequest
//
// @return ListCommonTicketFieldsResponse
func (client *Client) ListCommonTicketFields(request *ListCommonTicketFieldsRequest) (_result *ListCommonTicketFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCommonTicketFieldsResponse{}
	_body, _err := client.ListCommonTicketFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListConfigItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigItemsResponse
func (client *Client) ListConfigItemsWithOptions(request *ListConfigItemsRequest, runtime *dara.RuntimeOptions) (_result *ListConfigItemsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigItems"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListConfigItemsRequest
//
// @return ListConfigItemsResponse
func (client *Client) ListConfigItems(request *ListConfigItemsRequest) (_result *ListConfigItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConfigItemsResponse{}
	_body, _err := client.ListConfigItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// -
//
// @param request - ListContactFlowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactFlowsResponse
func (client *Client) ListContactFlowsWithOptions(request *ListContactFlowsRequest, runtime *dara.RuntimeOptions) (_result *ListContactFlowsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderByField) {
		query["OrderByField"] = request.OrderByField
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContactFlows"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContactFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// -
//
// @param request - ListContactFlowsRequest
//
// @return ListContactFlowsResponse
func (client *Client) ListContactFlows(request *ListContactFlowsRequest) (_result *ListContactFlowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContactFlowsResponse{}
	_body, _err := client.ListContactFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListCustomCallTagging is deprecated, please use CCC::2020-07-01::ListCustomCallTaggings instead.
//
// Summary:
//
// 列出呼入控制号码
//
// @param request - ListCustomCallTaggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomCallTaggingResponse
// Deprecated
func (client *Client) ListCustomCallTaggingWithOptions(request *ListCustomCallTaggingRequest, runtime *dara.RuntimeOptions) (_result *ListCustomCallTaggingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallTagNameList) {
		query["CallTagNameList"] = request.CallTagNameList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomCallTagging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomCallTaggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListCustomCallTagging is deprecated, please use CCC::2020-07-01::ListCustomCallTaggings instead.
//
// Summary:
//
// 列出呼入控制号码
//
// @param request - ListCustomCallTaggingRequest
//
// @return ListCustomCallTaggingResponse
// Deprecated
func (client *Client) ListCustomCallTagging(request *ListCustomCallTaggingRequest) (_result *ListCustomCallTaggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomCallTaggingResponse{}
	_body, _err := client.ListCustomCallTaggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDevicesResponse
func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *dara.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDevices"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDevicesRequest
//
// @return ListDevicesResponse
func (client *Client) ListDevices(request *ListDevicesRequest) (_result *ListDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDevicesResponse{}
	_body, _err := client.ListDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询黑名单号码
//
// @param request - ListDoNotCallNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoNotCallNumbersResponse
func (client *Client) ListDoNotCallNumbersWithOptions(request *ListDoNotCallNumbersRequest, runtime *dara.RuntimeOptions) (_result *ListDoNotCallNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoNotCallNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoNotCallNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询黑名单号码
//
// @param request - ListDoNotCallNumbersRequest
//
// @return ListDoNotCallNumbersResponse
func (client *Client) ListDoNotCallNumbers(request *ListDoNotCallNumbersRequest) (_result *ListDoNotCallNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoNotCallNumbersResponse{}
	_body, _err := client.ListDoNotCallNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - ListDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDocumentsResponse
func (client *Client) ListDocumentsWithOptions(tmpReq *ListDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ListDocumentsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDocumentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Sorts) {
		request.SortsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sorts, dara.String("Sorts"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NextPageToken) {
		body["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	if !dara.IsNil(request.SearchPattern) {
		body["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.SortsShrink) {
		body["Sorts"] = request.SortsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDocuments"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDocumentsRequest
//
// @return ListDocumentsResponse
func (client *Client) ListDocuments(request *ListDocumentsRequest) (_result *ListDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDocumentsResponse{}
	_body, _err := client.ListDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListFlashSmsApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsApplicationsResponse
func (client *Client) ListFlashSmsApplicationsWithOptions(request *ListFlashSmsApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProviderId) {
		query["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlashSmsApplications"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFlashSmsApplicationsRequest
//
// @return ListFlashSmsApplicationsResponse
func (client *Client) ListFlashSmsApplications(request *ListFlashSmsApplicationsRequest) (_result *ListFlashSmsApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlashSmsApplicationsResponse{}
	_body, _err := client.ListFlashSmsApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取闪信模板列表
//
// @param request - ListFlashSmsTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsTemplatesResponse
func (client *Client) ListFlashSmsTemplatesWithOptions(request *ListFlashSmsTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProviderId) {
		query["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlashSmsTemplates"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取闪信模板列表
//
// @param request - ListFlashSmsTemplatesRequest
//
// @return ListFlashSmsTemplatesResponse
func (client *Client) ListFlashSmsTemplates(request *ListFlashSmsTemplatesRequest) (_result *ListFlashSmsTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlashSmsTemplatesResponse{}
	_body, _err := client.ListFlashSmsTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListGroupChatMessages
//
// @param request - ListGroupChatMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupChatMessagesResponse
func (client *Client) ListGroupChatMessagesWithOptions(request *ListGroupChatMessagesRequest, runtime *dara.RuntimeOptions) (_result *ListGroupChatMessagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupChatMessages"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupChatMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListGroupChatMessages
//
// @param request - ListGroupChatMessagesRequest
//
// @return ListGroupChatMessagesResponse
func (client *Client) ListGroupChatMessages(request *ListGroupChatMessagesRequest) (_result *ListGroupChatMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGroupChatMessagesResponse{}
	_body, _err := client.ListGroupChatMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListHistoricalAgentReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHistoricalAgentReportResponse
func (client *Client) ListHistoricalAgentReportWithOptions(request *ListHistoricalAgentReportRequest, runtime *dara.RuntimeOptions) (_result *ListHistoricalAgentReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StopTime) {
		query["StopTime"] = request.StopTime
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentIdList) {
		body["AgentIdList"] = request.AgentIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHistoricalAgentReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHistoricalAgentReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListHistoricalAgentReportRequest
//
// @return ListHistoricalAgentReportResponse
func (client *Client) ListHistoricalAgentReport(request *ListHistoricalAgentReportRequest) (_result *ListHistoricalAgentReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHistoricalAgentReportResponse{}
	_body, _err := client.ListHistoricalAgentReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListHistoricalAgentSkillGroupReport
//
// @param request - ListHistoricalAgentSkillGroupReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHistoricalAgentSkillGroupReportResponse
func (client *Client) ListHistoricalAgentSkillGroupReportWithOptions(request *ListHistoricalAgentSkillGroupReportRequest, runtime *dara.RuntimeOptions) (_result *ListHistoricalAgentSkillGroupReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentIdList) {
		body["AgentIdList"] = request.AgentIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHistoricalAgentSkillGroupReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHistoricalAgentSkillGroupReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListHistoricalAgentSkillGroupReport
//
// @param request - ListHistoricalAgentSkillGroupReportRequest
//
// @return ListHistoricalAgentSkillGroupReportResponse
func (client *Client) ListHistoricalAgentSkillGroupReport(request *ListHistoricalAgentSkillGroupReportRequest) (_result *ListHistoricalAgentSkillGroupReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHistoricalAgentSkillGroupReportResponse{}
	_body, _err := client.ListHistoricalAgentSkillGroupReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListHistoricalSkillGroupReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHistoricalSkillGroupReportResponse
func (client *Client) ListHistoricalSkillGroupReportWithOptions(request *ListHistoricalSkillGroupReportRequest, runtime *dara.RuntimeOptions) (_result *ListHistoricalSkillGroupReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SkillGroupIdList) {
		body["SkillGroupIdList"] = request.SkillGroupIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHistoricalSkillGroupReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHistoricalSkillGroupReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListHistoricalSkillGroupReportRequest
//
// @return ListHistoricalSkillGroupReportResponse
func (client *Client) ListHistoricalSkillGroupReport(request *ListHistoricalSkillGroupReportRequest) (_result *ListHistoricalSkillGroupReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHistoricalSkillGroupReportResponse{}
	_body, _err := client.ListHistoricalSkillGroupReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListInstancesOfUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesOfUserResponse
func (client *Client) ListInstancesOfUserWithOptions(request *ListInstancesOfUserRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesOfUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstancesOfUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesOfUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListInstancesOfUserRequest
//
// @return ListInstancesOfUserResponse
func (client *Client) ListInstancesOfUser(request *ListInstancesOfUserRequest) (_result *ListInstancesOfUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancesOfUserResponse{}
	_body, _err := client.ListInstancesOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListIntervalAgentReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntervalAgentReportResponse
func (client *Client) ListIntervalAgentReportWithOptions(request *ListIntervalAgentReportRequest, runtime *dara.RuntimeOptions) (_result *ListIntervalAgentReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntervalAgentReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntervalAgentReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListIntervalAgentReportRequest
//
// @return ListIntervalAgentReportResponse
func (client *Client) ListIntervalAgentReport(request *ListIntervalAgentReportRequest) (_result *ListIntervalAgentReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIntervalAgentReportResponse{}
	_body, _err := client.ListIntervalAgentReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListIntervalAgentSkillGroupReport
//
// @param request - ListIntervalAgentSkillGroupReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntervalAgentSkillGroupReportResponse
func (client *Client) ListIntervalAgentSkillGroupReportWithOptions(request *ListIntervalAgentSkillGroupReportRequest, runtime *dara.RuntimeOptions) (_result *ListIntervalAgentSkillGroupReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntervalAgentSkillGroupReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntervalAgentSkillGroupReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListIntervalAgentSkillGroupReport
//
// @param request - ListIntervalAgentSkillGroupReportRequest
//
// @return ListIntervalAgentSkillGroupReportResponse
func (client *Client) ListIntervalAgentSkillGroupReport(request *ListIntervalAgentSkillGroupReportRequest) (_result *ListIntervalAgentSkillGroupReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIntervalAgentSkillGroupReportResponse{}
	_body, _err := client.ListIntervalAgentSkillGroupReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListIntervalInstanceReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntervalInstanceReportResponse
func (client *Client) ListIntervalInstanceReportWithOptions(request *ListIntervalInstanceReportRequest, runtime *dara.RuntimeOptions) (_result *ListIntervalInstanceReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntervalInstanceReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntervalInstanceReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListIntervalInstanceReportRequest
//
// @return ListIntervalInstanceReportResponse
func (client *Client) ListIntervalInstanceReport(request *ListIntervalInstanceReportRequest) (_result *ListIntervalInstanceReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIntervalInstanceReportResponse{}
	_body, _err := client.ListIntervalInstanceReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListIntervalSkillGroupReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntervalSkillGroupReportResponse
func (client *Client) ListIntervalSkillGroupReportWithOptions(request *ListIntervalSkillGroupReportRequest, runtime *dara.RuntimeOptions) (_result *ListIntervalSkillGroupReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntervalSkillGroupReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntervalSkillGroupReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListIntervalSkillGroupReportRequest
//
// @return ListIntervalSkillGroupReportResponse
func (client *Client) ListIntervalSkillGroupReport(request *ListIntervalSkillGroupReportRequest) (_result *ListIntervalSkillGroupReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIntervalSkillGroupReportResponse{}
	_body, _err := client.ListIntervalSkillGroupReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListIvrTrackingDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIvrTrackingDetailsResponse
func (client *Client) ListIvrTrackingDetailsWithOptions(request *ListIvrTrackingDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListIvrTrackingDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIvrTrackingDetails"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIvrTrackingDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListIvrTrackingDetailsRequest
//
// @return ListIvrTrackingDetailsResponse
func (client *Client) ListIvrTrackingDetails(request *ListIvrTrackingDetailsRequest) (_result *ListIvrTrackingDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIvrTrackingDetailsResponse{}
	_body, _err := client.ListIvrTrackingDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListLegacyAgentEventLogs is deprecated, please use CCC::2020-07-01::ListAgentStateLogs instead.
//
// Summary:
//
// # ListLegacyAgentEventLogs
//
// @param request - ListLegacyAgentEventLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLegacyAgentEventLogsResponse
// Deprecated
func (client *Client) ListLegacyAgentEventLogsWithOptions(request *ListLegacyAgentEventLogsRequest, runtime *dara.RuntimeOptions) (_result *ListLegacyAgentEventLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLegacyAgentEventLogs"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLegacyAgentEventLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListLegacyAgentEventLogs is deprecated, please use CCC::2020-07-01::ListAgentStateLogs instead.
//
// Summary:
//
// # ListLegacyAgentEventLogs
//
// @param request - ListLegacyAgentEventLogsRequest
//
// @return ListLegacyAgentEventLogsResponse
// Deprecated
func (client *Client) ListLegacyAgentEventLogs(request *ListLegacyAgentEventLogsRequest) (_result *ListLegacyAgentEventLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLegacyAgentEventLogsResponse{}
	_body, _err := client.ListLegacyAgentEventLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListLegacyAgentStatusLogs is deprecated, please use CCC::2020-07-01::ListAgentStateLogs instead.
//
// Summary:
//
// # ListLegacyAgentStatusLogs
//
// @param request - ListLegacyAgentStatusLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLegacyAgentStatusLogsResponse
// Deprecated
func (client *Client) ListLegacyAgentStatusLogsWithOptions(request *ListLegacyAgentStatusLogsRequest, runtime *dara.RuntimeOptions) (_result *ListLegacyAgentStatusLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLegacyAgentStatusLogs"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLegacyAgentStatusLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListLegacyAgentStatusLogs is deprecated, please use CCC::2020-07-01::ListAgentStateLogs instead.
//
// Summary:
//
// # ListLegacyAgentStatusLogs
//
// @param request - ListLegacyAgentStatusLogsRequest
//
// @return ListLegacyAgentStatusLogsResponse
// Deprecated
func (client *Client) ListLegacyAgentStatusLogs(request *ListLegacyAgentStatusLogsRequest) (_result *ListLegacyAgentStatusLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLegacyAgentStatusLogsResponse{}
	_body, _err := client.ListLegacyAgentStatusLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListLegacyAppraiseLogs
//
// @param request - ListLegacyAppraiseLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLegacyAppraiseLogsResponse
func (client *Client) ListLegacyAppraiseLogsWithOptions(request *ListLegacyAppraiseLogsRequest, runtime *dara.RuntimeOptions) (_result *ListLegacyAppraiseLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLegacyAppraiseLogs"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLegacyAppraiseLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListLegacyAppraiseLogs
//
// @param request - ListLegacyAppraiseLogsRequest
//
// @return ListLegacyAppraiseLogsResponse
func (client *Client) ListLegacyAppraiseLogs(request *ListLegacyAppraiseLogsRequest) (_result *ListLegacyAppraiseLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLegacyAppraiseLogsResponse{}
	_body, _err := client.ListLegacyAppraiseLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListLegacyQueueEventLogs
//
// @param request - ListLegacyQueueEventLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLegacyQueueEventLogsResponse
func (client *Client) ListLegacyQueueEventLogsWithOptions(request *ListLegacyQueueEventLogsRequest, runtime *dara.RuntimeOptions) (_result *ListLegacyQueueEventLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLegacyQueueEventLogs"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLegacyQueueEventLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListLegacyQueueEventLogs
//
// @param request - ListLegacyQueueEventLogsRequest
//
// @return ListLegacyQueueEventLogsResponse
func (client *Client) ListLegacyQueueEventLogs(request *ListLegacyQueueEventLogsRequest) (_result *ListLegacyQueueEventLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLegacyQueueEventLogsResponse{}
	_body, _err := client.ListLegacyQueueEventLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMonoRecordingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMonoRecordingsResponse
func (client *Client) ListMonoRecordingsWithOptions(request *ListMonoRecordingsRequest, runtime *dara.RuntimeOptions) (_result *ListMonoRecordingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMonoRecordings"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMonoRecordingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMonoRecordingsRequest
//
// @return ListMonoRecordingsResponse
func (client *Client) ListMonoRecordings(request *ListMonoRecordingsRequest) (_result *ListMonoRecordingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMonoRecordingsResponse{}
	_body, _err := client.ListMonoRecordingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMultiChannelRecordingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiChannelRecordingsResponse
func (client *Client) ListMultiChannelRecordingsWithOptions(request *ListMultiChannelRecordingsRequest, runtime *dara.RuntimeOptions) (_result *ListMultiChannelRecordingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultiChannelRecordings"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultiChannelRecordingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMultiChannelRecordingsRequest
//
// @return ListMultiChannelRecordingsResponse
func (client *Client) ListMultiChannelRecordings(request *ListMultiChannelRecordingsRequest) (_result *ListMultiChannelRecordingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultiChannelRecordingsResponse{}
	_body, _err := client.ListMultiChannelRecordingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 该坐席可用的外呼号码列表
//
// @param request - ListOutboundNumbersOfUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOutboundNumbersOfUserResponse
func (client *Client) ListOutboundNumbersOfUserWithOptions(request *ListOutboundNumbersOfUserRequest, runtime *dara.RuntimeOptions) (_result *ListOutboundNumbersOfUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOutboundNumbersOfUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOutboundNumbersOfUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 该坐席可用的外呼号码列表
//
// @param request - ListOutboundNumbersOfUserRequest
//
// @return ListOutboundNumbersOfUserResponse
func (client *Client) ListOutboundNumbersOfUser(request *ListOutboundNumbersOfUserRequest) (_result *ListOutboundNumbersOfUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOutboundNumbersOfUserResponse{}
	_body, _err := client.ListOutboundNumbersOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPersonalNumbersOfUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPersonalNumbersOfUserResponse
func (client *Client) ListPersonalNumbersOfUserWithOptions(request *ListPersonalNumbersOfUserRequest, runtime *dara.RuntimeOptions) (_result *ListPersonalNumbersOfUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsMember) {
		query["IsMember"] = request.IsMember
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPersonalNumbersOfUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPersonalNumbersOfUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPersonalNumbersOfUserRequest
//
// @return ListPersonalNumbersOfUserResponse
func (client *Client) ListPersonalNumbersOfUser(request *ListPersonalNumbersOfUserRequest) (_result *ListPersonalNumbersOfUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPersonalNumbersOfUserResponse{}
	_body, _err := client.ListPersonalNumbersOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取号码列表
//
// @param request - ListPhoneNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPhoneNumbersResponse
func (client *Client) ListPhoneNumbersWithOptions(request *ListPhoneNumbersRequest, runtime *dara.RuntimeOptions) (_result *ListPhoneNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.Usage) {
		query["Usage"] = request.Usage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPhoneNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPhoneNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取号码列表
//
// @param request - ListPhoneNumbersRequest
//
// @return ListPhoneNumbersResponse
func (client *Client) ListPhoneNumbers(request *ListPhoneNumbersRequest) (_result *ListPhoneNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPhoneNumbersResponse{}
	_body, _err := client.ListPhoneNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPhoneNumbersOfSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPhoneNumbersOfSkillGroupResponse
func (client *Client) ListPhoneNumbersOfSkillGroupWithOptions(request *ListPhoneNumbersOfSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *ListPhoneNumbersOfSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsMember) {
		query["IsMember"] = request.IsMember
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPhoneNumbersOfSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPhoneNumbersOfSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPhoneNumbersOfSkillGroupRequest
//
// @return ListPhoneNumbersOfSkillGroupResponse
func (client *Client) ListPhoneNumbersOfSkillGroup(request *ListPhoneNumbersOfSkillGroupRequest) (_result *ListPhoneNumbersOfSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPhoneNumbersOfSkillGroupResponse{}
	_body, _err := client.ListPhoneNumbersOfSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPrivilegesOfUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivilegesOfUserResponse
func (client *Client) ListPrivilegesOfUserWithOptions(request *ListPrivilegesOfUserRequest, runtime *dara.RuntimeOptions) (_result *ListPrivilegesOfUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivilegesOfUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivilegesOfUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPrivilegesOfUserRequest
//
// @return ListPrivilegesOfUserResponse
func (client *Client) ListPrivilegesOfUser(request *ListPrivilegesOfUserRequest) (_result *ListPrivilegesOfUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivilegesOfUserResponse{}
	_body, _err := client.ListPrivilegesOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListRamUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRamUsersResponse
func (client *Client) ListRamUsersWithOptions(request *ListRamUsersRequest, runtime *dara.RuntimeOptions) (_result *ListRamUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRamUsers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRamUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListRamUsersRequest
//
// @return ListRamUsersResponse
func (client *Client) ListRamUsers(request *ListRamUsersRequest) (_result *ListRamUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRamUsersResponse{}
	_body, _err := client.ListRamUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListRealtimeAgentStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRealtimeAgentStatesResponse
func (client *Client) ListRealtimeAgentStatesWithOptions(request *ListRealtimeAgentStatesRequest, runtime *dara.RuntimeOptions) (_result *ListRealtimeAgentStatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.CallTypeList) {
		query["CallTypeList"] = request.CallTypeList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.OutboundScenario) {
		query["OutboundScenario"] = request.OutboundScenario
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

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.WorkModeList) {
		query["WorkModeList"] = request.WorkModeList
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentIdList) {
		body["AgentIdList"] = request.AgentIdList
	}

	if !dara.IsNil(request.StateList) {
		body["StateList"] = request.StateList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRealtimeAgentStates"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRealtimeAgentStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListRealtimeAgentStatesRequest
//
// @return ListRealtimeAgentStatesResponse
func (client *Client) ListRealtimeAgentStates(request *ListRealtimeAgentStatesRequest) (_result *ListRealtimeAgentStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRealtimeAgentStatesResponse{}
	_body, _err := client.ListRealtimeAgentStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListRealtimeSkillGroupStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRealtimeSkillGroupStatesResponse
func (client *Client) ListRealtimeSkillGroupStatesWithOptions(request *ListRealtimeSkillGroupStatesRequest, runtime *dara.RuntimeOptions) (_result *ListRealtimeSkillGroupStatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SkillGroupIdList) {
		body["SkillGroupIdList"] = request.SkillGroupIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRealtimeSkillGroupStates"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRealtimeSkillGroupStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListRealtimeSkillGroupStatesRequest
//
// @return ListRealtimeSkillGroupStatesResponse
func (client *Client) ListRealtimeSkillGroupStates(request *ListRealtimeSkillGroupStatesRequest) (_result *ListRealtimeSkillGroupStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRealtimeSkillGroupStatesResponse{}
	_body, _err := client.ListRealtimeSkillGroupStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListRecentCallDetailRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentCallDetailRecordsResponse
func (client *Client) ListRecentCallDetailRecordsWithOptions(request *ListRecentCallDetailRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListRecentCallDetailRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Criteria) {
		query["Criteria"] = request.Criteria
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecentCallDetailRecords"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecentCallDetailRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListRecentCallDetailRecordsRequest
//
// @return ListRecentCallDetailRecordsResponse
func (client *Client) ListRecentCallDetailRecords(request *ListRecentCallDetailRecordsRequest) (_result *ListRecentCallDetailRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecentCallDetailRecordsResponse{}
	_body, _err := client.ListRecentCallDetailRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithOptions(request *ListRolesRequest, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoles"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListRolesRequest
//
// @return ListRolesResponse
func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListSkillGroupStates is deprecated, please use CCC::2020-07-01::ListRealtimeSkillGroupStates instead.
//
// Summary:
//
// # ListSkillGroupStates for acc
//
// @param request - ListSkillGroupStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillGroupStatesResponse
// Deprecated
func (client *Client) ListSkillGroupStatesWithOptions(request *ListSkillGroupStatesRequest, runtime *dara.RuntimeOptions) (_result *ListSkillGroupStatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillGroupStates"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillGroupStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListSkillGroupStates is deprecated, please use CCC::2020-07-01::ListRealtimeSkillGroupStates instead.
//
// Summary:
//
// # ListSkillGroupStates for acc
//
// @param request - ListSkillGroupStatesRequest
//
// @return ListSkillGroupStatesResponse
// Deprecated
func (client *Client) ListSkillGroupStates(request *ListSkillGroupStatesRequest) (_result *ListSkillGroupStatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillGroupStatesResponse{}
	_body, _err := client.ListSkillGroupStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListSkillGroupSummaryReportsSinceMidnight is deprecated, please use CCC::2020-07-01::ListHistoricalSkillGroupReport instead.
//
// Summary:
//
// # ListSkillGroupSummaryReportsSinceMidnight for acc
//
// @param request - ListSkillGroupSummaryReportsSinceMidnightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillGroupSummaryReportsSinceMidnightResponse
// Deprecated
func (client *Client) ListSkillGroupSummaryReportsSinceMidnightWithOptions(request *ListSkillGroupSummaryReportsSinceMidnightRequest, runtime *dara.RuntimeOptions) (_result *ListSkillGroupSummaryReportsSinceMidnightResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillGroupSummaryReportsSinceMidnight"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillGroupSummaryReportsSinceMidnightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListSkillGroupSummaryReportsSinceMidnight is deprecated, please use CCC::2020-07-01::ListHistoricalSkillGroupReport instead.
//
// Summary:
//
// # ListSkillGroupSummaryReportsSinceMidnight for acc
//
// @param request - ListSkillGroupSummaryReportsSinceMidnightRequest
//
// @return ListSkillGroupSummaryReportsSinceMidnightResponse
// Deprecated
func (client *Client) ListSkillGroupSummaryReportsSinceMidnight(request *ListSkillGroupSummaryReportsSinceMidnightRequest) (_result *ListSkillGroupSummaryReportsSinceMidnightResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillGroupSummaryReportsSinceMidnightResponse{}
	_body, _err := client.ListSkillGroupSummaryReportsSinceMidnightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListSkillGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillGroupsResponse
func (client *Client) ListSkillGroupsWithOptions(request *ListSkillGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListSkillGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillGroups"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListSkillGroupsRequest
//
// @return ListSkillGroupsResponse
func (client *Client) ListSkillGroups(request *ListSkillGroupsRequest) (_result *ListSkillGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillGroupsResponse{}
	_body, _err := client.ListSkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListSkillLevelsOfUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillLevelsOfUserResponse
func (client *Client) ListSkillLevelsOfUserWithOptions(request *ListSkillLevelsOfUserRequest, runtime *dara.RuntimeOptions) (_result *ListSkillLevelsOfUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsMember) {
		query["IsMember"] = request.IsMember
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillLevelsOfUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillLevelsOfUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListSkillLevelsOfUserRequest
//
// @return ListSkillLevelsOfUserResponse
func (client *Client) ListSkillLevelsOfUser(request *ListSkillLevelsOfUserRequest) (_result *ListSkillLevelsOfUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillLevelsOfUserResponse{}
	_body, _err := client.ListSkillLevelsOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListSmsMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSmsMetadataResponse
func (client *Client) ListSmsMetadataWithOptions(request *ListSmsMetadataRequest, runtime *dara.RuntimeOptions) (_result *ListSmsMetadataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScenarioListJson) {
		query["ScenarioListJson"] = request.ScenarioListJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSmsMetadata"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSmsMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListSmsMetadataRequest
//
// @return ListSmsMetadataResponse
func (client *Client) ListSmsMetadata(request *ListSmsMetadataRequest) (_result *ListSmsMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSmsMetadataResponse{}
	_body, _err := client.ListSmsMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListTicketTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketTasksResponse
func (client *Client) ListTicketTasksWithOptions(request *ListTicketTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTicketTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTicketTasks"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTicketTasksRequest
//
// @return ListTicketTasksResponse
func (client *Client) ListTicketTasks(request *ListTicketTasksRequest) (_result *ListTicketTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTicketTasksResponse{}
	_body, _err := client.ListTicketTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListTicketTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketTemplatesResponse
func (client *Client) ListTicketTemplatesWithOptions(request *ListTicketTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListTicketTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTicketTemplates"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTicketTemplatesRequest
//
// @return ListTicketTemplatesResponse
func (client *Client) ListTicketTemplates(request *ListTicketTemplatesRequest) (_result *ListTicketTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTicketTemplatesResponse{}
	_body, _err := client.ListTicketTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListTicketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketsResponse
func (client *Client) ListTicketsWithOptions(request *ListTicketsRequest, runtime *dara.RuntimeOptions) (_result *ListTicketsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Assignee) {
		query["Assignee"] = request.Assignee
	}

	if !dara.IsNil(request.AssigneeType) {
		query["AssigneeType"] = request.AssigneeType
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobIdList) {
		query["JobIdList"] = request.JobIdList
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Participant) {
		query["Participant"] = request.Participant
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTickets"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListTicketsRequest
//
// @return ListTicketsResponse
func (client *Client) ListTickets(request *ListTicketsRequest) (_result *ListTicketsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTicketsResponse{}
	_body, _err := client.ListTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取未分配号码列表
//
// @param request - ListUnassignedNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUnassignedNumbersResponse
func (client *Client) ListUnassignedNumbersWithOptions(request *ListUnassignedNumbersRequest, runtime *dara.RuntimeOptions) (_result *ListUnassignedNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUnassignedNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUnassignedNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取未分配号码列表
//
// @param request - ListUnassignedNumbersRequest
//
// @return ListUnassignedNumbersResponse
func (client *Client) ListUnassignedNumbers(request *ListUnassignedNumbersRequest) (_result *ListUnassignedNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUnassignedNumbersResponse{}
	_body, _err := client.ListUnassignedNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserLevelsOfSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserLevelsOfSkillGroupResponse
func (client *Client) ListUserLevelsOfSkillGroupWithOptions(request *ListUserLevelsOfSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *ListUserLevelsOfSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsMember) {
		query["IsMember"] = request.IsMember
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserLevelsOfSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserLevelsOfSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserLevelsOfSkillGroupRequest
//
// @return ListUserLevelsOfSkillGroupResponse
func (client *Client) ListUserLevelsOfSkillGroup(request *ListUserLevelsOfSkillGroupRequest) (_result *ListUserLevelsOfSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserLevelsOfSkillGroupResponse{}
	_body, _err := client.ListUserLevelsOfSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		query["SearchPattern"] = request.SearchPattern
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2020-07-01"),
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
// # ListVisitorChatMessages
//
// @param request - ListVisitorChatMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVisitorChatMessagesResponse
func (client *Client) ListVisitorChatMessagesWithOptions(request *ListVisitorChatMessagesRequest, runtime *dara.RuntimeOptions) (_result *ListVisitorChatMessagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessChannelId) {
		query["AccessChannelId"] = request.AccessChannelId
	}

	if !dara.IsNil(request.AccessToken) {
		query["AccessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VisitorId) {
		query["VisitorId"] = request.VisitorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVisitorChatMessages"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVisitorChatMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListVisitorChatMessages
//
// @param request - ListVisitorChatMessagesRequest
//
// @return ListVisitorChatMessagesResponse
func (client *Client) ListVisitorChatMessages(request *ListVisitorChatMessagesRequest) (_result *ListVisitorChatMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVisitorChatMessagesResponse{}
	_body, _err := client.ListVisitorChatMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListVoicemailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVoicemailsResponse
func (client *Client) ListVoicemailsWithOptions(request *ListVoicemailsRequest, runtime *dara.RuntimeOptions) (_result *ListVoicemailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVoicemails"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVoicemailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListVoicemailsRequest
//
// @return ListVoicemailsResponse
func (client *Client) ListVoicemails(request *ListVoicemailsRequest) (_result *ListVoicemailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVoicemailsResponse{}
	_body, _err := client.ListVoicemailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListWaitingChatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWaitingChatsResponse
func (client *Client) ListWaitingChatsWithOptions(request *ListWaitingChatsRequest, runtime *dara.RuntimeOptions) (_result *ListWaitingChatsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWaitingChats"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWaitingChatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListWaitingChatsRequest
//
// @return ListWaitingChatsResponse
func (client *Client) ListWaitingChats(request *ListWaitingChatsRequest) (_result *ListWaitingChatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWaitingChatsResponse{}
	_body, _err := client.ListWaitingChatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MakeCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MakeCallResponse
func (client *Client) MakeCallWithOptions(request *MakeCallRequest, runtime *dara.RuntimeOptions) (_result *MakeCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.FlashSmsVariables) {
		query["FlashSmsVariables"] = request.FlashSmsVariables
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaskedCallee) {
		query["MaskedCallee"] = request.MaskedCallee
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MakeCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MakeCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MakeCallRequest
//
// @return MakeCallResponse
func (client *Client) MakeCall(request *MakeCallRequest) (_result *MakeCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MakeCallResponse{}
	_body, _err := client.MakeCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAudioFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAudioFileResponse
func (client *Client) ModifyAudioFileWithOptions(request *ModifyAudioFileRequest, runtime *dara.RuntimeOptions) (_result *ModifyAudioFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioFileName) {
		query["AudioFileName"] = request.AudioFileName
	}

	if !dara.IsNil(request.AudioResourceId) {
		query["AudioResourceId"] = request.AudioResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssFileKey) {
		query["OssFileKey"] = request.OssFileKey
	}

	if !dara.IsNil(request.Usage) {
		query["Usage"] = request.Usage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAudioFile"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAudioFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAudioFileRequest
//
// @return ModifyAudioFileResponse
func (client *Client) ModifyAudioFile(request *ModifyAudioFileRequest) (_result *ModifyAudioFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAudioFileResponse{}
	_body, _err := client.ModifyAudioFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑呼入控制号码
//
// @param request - ModifyCustomCallTaggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomCallTaggingResponse
func (client *Client) ModifyCustomCallTaggingWithOptions(request *ModifyCustomCallTaggingRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomCallTaggingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallTagNameList) {
		query["CallTagNameList"] = request.CallTagNameList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomCallTagging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomCallTaggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑呼入控制号码
//
// @param request - ModifyCustomCallTaggingRequest
//
// @return ModifyCustomCallTaggingResponse
func (client *Client) ModifyCustomCallTagging(request *ModifyCustomCallTaggingRequest) (_result *ModifyCustomCallTaggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCustomCallTaggingResponse{}
	_body, _err := client.ModifyCustomCallTaggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstanceWithOptions(request *ModifyInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstance"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyInstanceRequest
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstance(request *ModifyInstanceRequest) (_result *ModifyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.ModifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyPhoneNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPhoneNumberResponse
func (client *Client) ModifyPhoneNumberWithOptions(request *ModifyPhoneNumberRequest, runtime *dara.RuntimeOptions) (_result *ModifyPhoneNumberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.Usage) {
		query["Usage"] = request.Usage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPhoneNumber"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyPhoneNumberRequest
//
// @return ModifyPhoneNumberResponse
func (client *Client) ModifyPhoneNumber(request *ModifyPhoneNumberRequest) (_result *ModifyPhoneNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPhoneNumberResponse{}
	_body, _err := client.ModifyPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifySkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySkillGroupResponse
func (client *Client) ModifySkillGroupWithOptions(request *ModifySkillGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifySkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifySkillGroupRequest
//
// @return ModifySkillGroupResponse
func (client *Client) ModifySkillGroup(request *ModifySkillGroupRequest) (_result *ModifySkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySkillGroupResponse{}
	_body, _err := client.ModifySkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifySkillLevelsOfUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySkillLevelsOfUserResponse
func (client *Client) ModifySkillLevelsOfUserWithOptions(request *ModifySkillLevelsOfUserRequest, runtime *dara.RuntimeOptions) (_result *ModifySkillLevelsOfUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillLevelList) {
		query["SkillLevelList"] = request.SkillLevelList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySkillLevelsOfUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySkillLevelsOfUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifySkillLevelsOfUserRequest
//
// @return ModifySkillLevelsOfUserResponse
func (client *Client) ModifySkillLevelsOfUser(request *ModifySkillLevelsOfUserRequest) (_result *ModifySkillLevelsOfUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySkillLevelsOfUserResponse{}
	_body, _err := client.ModifySkillLevelsOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserResponse
func (client *Client) ModifyUserWithOptions(request *ModifyUserRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarUrl) {
		query["AvatarUrl"] = request.AvatarUrl
	}

	if !dara.IsNil(request.DisplayId) {
		query["DisplayId"] = request.DisplayId
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.Nickname) {
		query["Nickname"] = request.Nickname
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkMode) {
		query["WorkMode"] = request.WorkMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyUserRequest
//
// @return ModifyUserResponse
func (client *Client) ModifyUser(request *ModifyUserRequest) (_result *ModifyUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserResponse{}
	_body, _err := client.ModifyUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyUserLevelsOfSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserLevelsOfSkillGroupResponse
func (client *Client) ModifyUserLevelsOfSkillGroupWithOptions(request *ModifyUserLevelsOfSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserLevelsOfSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.UserLevelList) {
		query["UserLevelList"] = request.UserLevelList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserLevelsOfSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserLevelsOfSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyUserLevelsOfSkillGroupRequest
//
// @return ModifyUserLevelsOfSkillGroupResponse
func (client *Client) ModifyUserLevelsOfSkillGroup(request *ModifyUserLevelsOfSkillGroupRequest) (_result *ModifyUserLevelsOfSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserLevelsOfSkillGroupResponse{}
	_body, _err := client.ModifyUserLevelsOfSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MonitorCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MonitorCallResponse
func (client *Client) MonitorCallWithOptions(request *MonitorCallRequest, runtime *dara.RuntimeOptions) (_result *MonitorCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MonitoredUserId) {
		query["MonitoredUserId"] = request.MonitoredUserId
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MonitorCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MonitorCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MonitorCallRequest
//
// @return MonitorCallResponse
func (client *Client) MonitorCall(request *MonitorCallRequest) (_result *MonitorCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MonitorCallResponse{}
	_body, _err := client.MonitorCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MuteCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MuteCallResponse
func (client *Client) MuteCallWithOptions(request *MuteCallRequest, runtime *dara.RuntimeOptions) (_result *MuteCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MuteCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MuteCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MuteCallRequest
//
// @return MuteCallResponse
func (client *Client) MuteCall(request *MuteCallRequest) (_result *MuteCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MuteCallResponse{}
	_body, _err := client.MuteCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂停预测式外呼活动
//
// @param request - PauseCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseCampaignResponse
func (client *Client) PauseCampaignWithOptions(request *PauseCampaignRequest, runtime *dara.RuntimeOptions) (_result *PauseCampaignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停预测式外呼活动
//
// @param request - PauseCampaignRequest
//
// @return PauseCampaignResponse
func (client *Client) PauseCampaign(request *PauseCampaignRequest) (_result *PauseCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseCampaignResponse{}
	_body, _err := client.PauseCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PickOutboundNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PickOutboundNumbersResponse
func (client *Client) PickOutboundNumbersWithOptions(request *PickOutboundNumbersRequest, runtime *dara.RuntimeOptions) (_result *PickOutboundNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PickOutboundNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PickOutboundNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PickOutboundNumbersRequest
//
// @return PickOutboundNumbersResponse
func (client *Client) PickOutboundNumbers(request *PickOutboundNumbersRequest) (_result *PickOutboundNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PickOutboundNumbersResponse{}
	_body, _err := client.PickOutboundNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PollUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PollUserStatusResponse
func (client *Client) PollUserStatusWithOptions(request *PollUserStatusRequest, runtime *dara.RuntimeOptions) (_result *PollUserStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PollUserStatus"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PollUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PollUserStatusRequest
//
// @return PollUserStatusResponse
func (client *Client) PollUserStatus(request *PollUserStatusRequest) (_result *PollUserStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PollUserStatusResponse{}
	_body, _err := client.PollUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ProcessAliMeCallbackOfStaging
//
// @param request - ProcessAliMeCallbackOfStagingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProcessAliMeCallbackOfStagingResponse
func (client *Client) ProcessAliMeCallbackOfStagingWithOptions(request *ProcessAliMeCallbackOfStagingRequest, runtime *dara.RuntimeOptions) (_result *ProcessAliMeCallbackOfStagingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProcessAliMeCallbackOfStaging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProcessAliMeCallbackOfStagingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ProcessAliMeCallbackOfStaging
//
// @param request - ProcessAliMeCallbackOfStagingRequest
//
// @return ProcessAliMeCallbackOfStagingResponse
func (client *Client) ProcessAliMeCallbackOfStaging(request *ProcessAliMeCallbackOfStagingRequest) (_result *ProcessAliMeCallbackOfStagingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ProcessAliMeCallbackOfStagingResponse{}
	_body, _err := client.ProcessAliMeCallbackOfStagingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ProcessCustomIMCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProcessCustomIMCallbackResponse
func (client *Client) ProcessCustomIMCallbackWithOptions(request *ProcessCustomIMCallbackRequest, runtime *dara.RuntimeOptions) (_result *ProcessCustomIMCallbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessChannelId) {
		body["AccessChannelId"] = request.AccessChannelId
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MessageContent) {
		body["MessageContent"] = request.MessageContent
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SenderAvatarMediaId) {
		body["SenderAvatarMediaId"] = request.SenderAvatarMediaId
	}

	if !dara.IsNil(request.SenderId) {
		body["SenderId"] = request.SenderId
	}

	if !dara.IsNil(request.SenderName) {
		body["SenderName"] = request.SenderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProcessCustomIMCallback"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProcessCustomIMCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ProcessCustomIMCallbackRequest
//
// @return ProcessCustomIMCallbackResponse
func (client *Client) ProcessCustomIMCallback(request *ProcessCustomIMCallbackRequest) (_result *ProcessCustomIMCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ProcessCustomIMCallbackResponse{}
	_body, _err := client.ProcessCustomIMCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PublishContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishContactFlowResponse
func (client *Client) PublishContactFlowWithOptions(request *PublishContactFlowRequest, runtime *dara.RuntimeOptions) (_result *PublishContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.DraftId) {
		query["DraftId"] = request.DraftId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PublishContactFlowRequest
//
// @return PublishContactFlowResponse
func (client *Client) PublishContactFlow(request *PublishContactFlowRequest) (_result *PublishContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishContactFlowResponse{}
	_body, _err := client.PublishContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReadyForServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadyForServiceResponse
func (client *Client) ReadyForServiceWithOptions(request *ReadyForServiceRequest, runtime *dara.RuntimeOptions) (_result *ReadyForServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundScenario) {
		query["OutboundScenario"] = request.OutboundScenario
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadyForService"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadyForServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReadyForServiceRequest
//
// @return ReadyForServiceResponse
func (client *Client) ReadyForService(request *ReadyForServiceRequest) (_result *ReadyForServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadyForServiceResponse{}
	_body, _err := client.ReadyForServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RedialCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RedialCallResponse
func (client *Client) RedialCallWithOptions(request *RedialCallRequest, runtime *dara.RuntimeOptions) (_result *RedialCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RedialCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RedialCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RedialCallRequest
//
// @return RedialCallResponse
func (client *Client) RedialCall(request *RedialCallRequest) (_result *RedialCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RedialCallResponse{}
	_body, _err := client.RedialCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RegisterDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDeviceResponse
func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, runtime *dara.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterDevice"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RegisterDeviceRequest
//
// @return RegisterDeviceResponse
func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RegisterDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDevicesResponse
func (client *Client) RegisterDevicesWithOptions(request *RegisterDevicesRequest, runtime *dara.RuntimeOptions) (_result *RegisterDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.UserIdListJson) {
		query["UserIdListJson"] = request.UserIdListJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterDevices"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RegisterDevicesRequest
//
// @return RegisterDevicesResponse
func (client *Client) RegisterDevices(request *RegisterDevicesRequest) (_result *RegisterDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterDevicesResponse{}
	_body, _err := client.RegisterDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RejectChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectChatResponse
func (client *Client) RejectChatWithOptions(request *RejectChatRequest, runtime *dara.RuntimeOptions) (_result *RejectChatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectChat"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RejectChatRequest
//
// @return RejectChatResponse
func (client *Client) RejectChat(request *RejectChatRequest) (_result *RejectChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectChatResponse{}
	_body, _err := client.RejectChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RejectTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectTicketResponse
func (client *Client) RejectTicketWithOptions(request *RejectTicketRequest, runtime *dara.RuntimeOptions) (_result *RejectTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectTicket"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RejectTicketRequest
//
// @return RejectTicketResponse
func (client *Client) RejectTicket(request *RejectTicketRequest) (_result *RejectTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectTicketResponse{}
	_body, _err := client.RejectTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReleaseCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseCallResponse
func (client *Client) ReleaseCallWithOptions(request *ReleaseCallRequest, runtime *dara.RuntimeOptions) (_result *ReleaseCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReleaseCallRequest
//
// @return ReleaseCallResponse
func (client *Client) ReleaseCall(request *ReleaseCallRequest) (_result *ReleaseCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseCallResponse{}
	_body, _err := client.ReleaseCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReleaseChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseChatResponse
func (client *Client) ReleaseChatWithOptions(request *ReleaseChatRequest, runtime *dara.RuntimeOptions) (_result *ReleaseChatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseChat"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReleaseChatRequest
//
// @return ReleaseChatResponse
func (client *Client) ReleaseChat(request *ReleaseChatRequest) (_result *ReleaseChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseChatResponse{}
	_body, _err := client.ReleaseChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑呼入控制号码
//
// @param request - RemoveBlacklistCallTaggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveBlacklistCallTaggingResponse
func (client *Client) RemoveBlacklistCallTaggingWithOptions(request *RemoveBlacklistCallTaggingRequest, runtime *dara.RuntimeOptions) (_result *RemoveBlacklistCallTaggingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveBlacklistCallTagging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveBlacklistCallTaggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑呼入控制号码
//
// @param request - RemoveBlacklistCallTaggingRequest
//
// @return RemoveBlacklistCallTaggingResponse
func (client *Client) RemoveBlacklistCallTagging(request *RemoveBlacklistCallTaggingRequest) (_result *RemoveBlacklistCallTaggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveBlacklistCallTaggingResponse{}
	_body, _err := client.RemoveBlacklistCallTaggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除黑名单号码
//
// @param request - RemoveDoNotCallNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDoNotCallNumbersResponse
func (client *Client) RemoveDoNotCallNumbersWithOptions(request *RemoveDoNotCallNumbersRequest, runtime *dara.RuntimeOptions) (_result *RemoveDoNotCallNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDoNotCallNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDoNotCallNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除黑名单号码
//
// @param request - RemoveDoNotCallNumbersRequest
//
// @return RemoveDoNotCallNumbersResponse
func (client *Client) RemoveDoNotCallNumbers(request *RemoveDoNotCallNumbersRequest) (_result *RemoveDoNotCallNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveDoNotCallNumbersResponse{}
	_body, _err := client.RemoveDoNotCallNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemovePersonalNumbersFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePersonalNumbersFromUserResponse
func (client *Client) RemovePersonalNumbersFromUserWithOptions(request *RemovePersonalNumbersFromUserRequest, runtime *dara.RuntimeOptions) (_result *RemovePersonalNumbersFromUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePersonalNumbersFromUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePersonalNumbersFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemovePersonalNumbersFromUserRequest
//
// @return RemovePersonalNumbersFromUserResponse
func (client *Client) RemovePersonalNumbersFromUser(request *RemovePersonalNumbersFromUserRequest) (_result *RemovePersonalNumbersFromUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePersonalNumbersFromUserResponse{}
	_body, _err := client.RemovePersonalNumbersFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemovePhoneNumberFromSkillGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePhoneNumberFromSkillGroupsResponse
func (client *Client) RemovePhoneNumberFromSkillGroupsWithOptions(request *RemovePhoneNumberFromSkillGroupsRequest, runtime *dara.RuntimeOptions) (_result *RemovePhoneNumberFromSkillGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePhoneNumberFromSkillGroups"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePhoneNumberFromSkillGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemovePhoneNumberFromSkillGroupsRequest
//
// @return RemovePhoneNumberFromSkillGroupsResponse
func (client *Client) RemovePhoneNumberFromSkillGroups(request *RemovePhoneNumberFromSkillGroupsRequest) (_result *RemovePhoneNumberFromSkillGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePhoneNumberFromSkillGroupsResponse{}
	_body, _err := client.RemovePhoneNumberFromSkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemovePhoneNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePhoneNumbersResponse
func (client *Client) RemovePhoneNumbersWithOptions(request *RemovePhoneNumbersRequest, runtime *dara.RuntimeOptions) (_result *RemovePhoneNumbersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePhoneNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePhoneNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemovePhoneNumbersRequest
//
// @return RemovePhoneNumbersResponse
func (client *Client) RemovePhoneNumbers(request *RemovePhoneNumbersRequest) (_result *RemovePhoneNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePhoneNumbersResponse{}
	_body, _err := client.RemovePhoneNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemovePhoneNumbersFromSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePhoneNumbersFromSkillGroupResponse
func (client *Client) RemovePhoneNumbersFromSkillGroupWithOptions(request *RemovePhoneNumbersFromSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *RemovePhoneNumbersFromSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePhoneNumbersFromSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePhoneNumbersFromSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemovePhoneNumbersFromSkillGroupRequest
//
// @return RemovePhoneNumbersFromSkillGroupResponse
func (client *Client) RemovePhoneNumbersFromSkillGroup(request *RemovePhoneNumbersFromSkillGroupRequest) (_result *RemovePhoneNumbersFromSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePhoneNumbersFromSkillGroupResponse{}
	_body, _err := client.RemovePhoneNumbersFromSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemoveSkillGroupsFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSkillGroupsFromUserResponse
func (client *Client) RemoveSkillGroupsFromUserWithOptions(request *RemoveSkillGroupsFromUserRequest, runtime *dara.RuntimeOptions) (_result *RemoveSkillGroupsFromUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSkillGroupsFromUser"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSkillGroupsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveSkillGroupsFromUserRequest
//
// @return RemoveSkillGroupsFromUserResponse
func (client *Client) RemoveSkillGroupsFromUser(request *RemoveSkillGroupsFromUserRequest) (_result *RemoveSkillGroupsFromUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveSkillGroupsFromUserResponse{}
	_body, _err := client.RemoveSkillGroupsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemoveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsersWithOptions(request *RemoveUsersRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NotificationEmail) {
		query["NotificationEmail"] = request.NotificationEmail
	}

	if !dara.IsNil(request.UserIdList) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveUsersRequest
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsers(request *RemoveUsersRequest) (_result *RemoveUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUsersResponse{}
	_body, _err := client.RemoveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemoveUsersFromSkillGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersFromSkillGroupResponse
func (client *Client) RemoveUsersFromSkillGroupWithOptions(request *RemoveUsersFromSkillGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersFromSkillGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.UserIdList) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsersFromSkillGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUsersFromSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveUsersFromSkillGroupRequest
//
// @return RemoveUsersFromSkillGroupResponse
func (client *Client) RemoveUsersFromSkillGroup(request *RemoveUsersFromSkillGroupRequest) (_result *RemoveUsersFromSkillGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUsersFromSkillGroupResponse{}
	_body, _err := client.RemoveUsersFromSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResetAgentStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAgentStateResponse
func (client *Client) ResetAgentStateWithOptions(request *ResetAgentStateRequest, runtime *dara.RuntimeOptions) (_result *ResetAgentStateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAgentState"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAgentStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetAgentStateRequest
//
// @return ResetAgentStateResponse
func (client *Client) ResetAgentState(request *ResetAgentStateRequest) (_result *ResetAgentStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAgentStateResponse{}
	_body, _err := client.ResetAgentStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResetUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPasswordWithOptions(request *ResetUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetUserPassword"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetUserPasswordRequest
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPassword(request *ResetUserPasswordRequest) (_result *ResetUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.ResetUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 录音解冻
//
// @param request - RestoreArchivedRecordingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreArchivedRecordingsResponse
func (client *Client) RestoreArchivedRecordingsWithOptions(request *RestoreArchivedRecordingsRequest, runtime *dara.RuntimeOptions) (_result *RestoreArchivedRecordingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactIds) {
		query["ContactIds"] = request.ContactIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreArchivedRecordings"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreArchivedRecordingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 录音解冻
//
// @param request - RestoreArchivedRecordingsRequest
//
// @return RestoreArchivedRecordingsResponse
func (client *Client) RestoreArchivedRecordings(request *RestoreArchivedRecordingsRequest) (_result *RestoreArchivedRecordingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreArchivedRecordingsResponse{}
	_body, _err := client.RestoreArchivedRecordingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResubmitTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResubmitTicketResponse
func (client *Client) ResubmitTicketWithOptions(request *ResubmitTicketRequest, runtime *dara.RuntimeOptions) (_result *ResubmitTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResubmitTicket"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResubmitTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResubmitTicketRequest
//
// @return ResubmitTicketResponse
func (client *Client) ResubmitTicket(request *ResubmitTicketRequest) (_result *ResubmitTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResubmitTicketResponse{}
	_body, _err := client.ResubmitTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复预测式外呼活动
//
// @param request - ResumeCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeCampaignResponse
func (client *Client) ResumeCampaignWithOptions(request *ResumeCampaignRequest, runtime *dara.RuntimeOptions) (_result *ResumeCampaignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复预测式外呼活动
//
// @param request - ResumeCampaignRequest
//
// @return ResumeCampaignResponse
func (client *Client) ResumeCampaign(request *ResumeCampaignRequest) (_result *ResumeCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeCampaignResponse{}
	_body, _err := client.ResumeCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RetrieveCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveCallResponse
func (client *Client) RetrieveCallWithOptions(request *RetrieveCallRequest, runtime *dara.RuntimeOptions) (_result *RetrieveCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetrieveCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RetrieveCallRequest
//
// @return RetrieveCallResponse
func (client *Client) RetrieveCall(request *RetrieveCallRequest) (_result *RetrieveCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetrieveCallResponse{}
	_body, _err := client.RetrieveCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveDocumentResponse
func (client *Client) SaveDocumentWithOptions(request *SaveDocumentRequest, runtime *dara.RuntimeOptions) (_result *SaveDocumentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		body["DocumentId"] = request.DocumentId
	}

	if !dara.IsNil(request.DocumentJson) {
		body["DocumentJson"] = request.DocumentJson
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveDocument"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveDocumentRequest
//
// @return SaveDocumentResponse
func (client *Client) SaveDocument(request *SaveDocumentRequest) (_result *SaveDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveDocumentResponse{}
	_body, _err := client.SaveDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveRTCStatsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveRTCStatsV2Response
func (client *Client) SaveRTCStatsV2WithOptions(request *SaveRTCStatsV2Request, runtime *dara.RuntimeOptions) (_result *SaveRTCStatsV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.GeneralInfo) {
		query["GeneralInfo"] = request.GeneralInfo
	}

	if !dara.IsNil(request.GoogAddress) {
		query["GoogAddress"] = request.GoogAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ReceiverReport) {
		query["ReceiverReport"] = request.ReceiverReport
	}

	if !dara.IsNil(request.SenderReport) {
		query["SenderReport"] = request.SenderReport
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveRTCStatsV2"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveRTCStatsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveRTCStatsV2Request
//
// @return SaveRTCStatsV2Response
func (client *Client) SaveRTCStatsV2(request *SaveRTCStatsV2Request) (_result *SaveRTCStatsV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveRTCStatsV2Response{}
	_body, _err := client.SaveRTCStatsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveTerminalLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTerminalLogResponse
func (client *Client) SaveTerminalLogWithOptions(request *SaveTerminalLogRequest, runtime *dara.RuntimeOptions) (_result *SaveTerminalLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MethodName) {
		query["MethodName"] = request.MethodName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UniqueRequestId) {
		query["UniqueRequestId"] = request.UniqueRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTerminalLog"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTerminalLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveTerminalLogRequest
//
// @return SaveTerminalLogResponse
func (client *Client) SaveTerminalLog(request *SaveTerminalLogRequest) (_result *SaveTerminalLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveTerminalLogResponse{}
	_body, _err := client.SaveTerminalLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveWebRTCStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveWebRTCStatsResponse
func (client *Client) SaveWebRTCStatsWithOptions(request *SaveWebRTCStatsRequest, runtime *dara.RuntimeOptions) (_result *SaveWebRTCStatsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.GeneralInfo) {
		query["GeneralInfo"] = request.GeneralInfo
	}

	if !dara.IsNil(request.GoogAddress) {
		query["GoogAddress"] = request.GoogAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ReceiverReport) {
		query["ReceiverReport"] = request.ReceiverReport
	}

	if !dara.IsNil(request.SenderReport) {
		query["SenderReport"] = request.SenderReport
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveWebRTCStats"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveWebRTCStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveWebRTCStatsRequest
//
// @return SaveWebRTCStatsResponse
func (client *Client) SaveWebRTCStats(request *SaveWebRTCStatsRequest) (_result *SaveWebRTCStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveWebRTCStatsResponse{}
	_body, _err := client.SaveWebRTCStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveWebRtcInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveWebRtcInfoResponse
func (client *Client) SaveWebRtcInfoWithOptions(request *SaveWebRtcInfoRequest, runtime *dara.RuntimeOptions) (_result *SaveWebRtcInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveWebRtcInfo"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveWebRtcInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveWebRtcInfoRequest
//
// @return SaveWebRtcInfoResponse
func (client *Client) SaveWebRtcInfo(request *SaveWebRtcInfoRequest) (_result *SaveWebRtcInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveWebRtcInfoResponse{}
	_body, _err := client.SaveWebRtcInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SendDtmfSignalingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendDtmfSignalingResponse
func (client *Client) SendDtmfSignalingWithOptions(request *SendDtmfSignalingRequest, runtime *dara.RuntimeOptions) (_result *SendDtmfSignalingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Dtmf) {
		query["Dtmf"] = request.Dtmf
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendDtmfSignaling"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendDtmfSignalingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SendDtmfSignalingRequest
//
// @return SendDtmfSignalingResponse
func (client *Client) SendDtmfSignaling(request *SendDtmfSignalingRequest) (_result *SendDtmfSignalingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendDtmfSignalingResponse{}
	_body, _err := client.SendDtmfSignalingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SignInGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignInGroupResponse
func (client *Client) SignInGroupWithOptions(request *SignInGroupRequest, runtime *dara.RuntimeOptions) (_result *SignInGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Additivity) {
		query["Additivity"] = request.Additivity
	}

	if !dara.IsNil(request.ChatDeviceId) {
		query["ChatDeviceId"] = request.ChatDeviceId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SignedSkillGroupIdList) {
		query["SignedSkillGroupIdList"] = request.SignedSkillGroupIdList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SignInGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SignInGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SignInGroupRequest
//
// @return SignInGroupResponse
func (client *Client) SignInGroup(request *SignInGroupRequest) (_result *SignInGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SignInGroupResponse{}
	_body, _err := client.SignInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SignOutGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignOutGroupResponse
func (client *Client) SignOutGroupWithOptions(request *SignOutGroupRequest, runtime *dara.RuntimeOptions) (_result *SignOutGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SignOutGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SignOutGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SignOutGroupRequest
//
// @return SignOutGroupResponse
func (client *Client) SignOutGroup(request *SignOutGroupRequest) (_result *SignOutGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SignOutGroupResponse{}
	_body, _err := client.SignOutGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartBack2BackCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBack2BackCallResponse
func (client *Client) StartBack2BackCallWithOptions(request *StartBack2BackCallRequest, runtime *dara.RuntimeOptions) (_result *StartBack2BackCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalBroker) {
		query["AdditionalBroker"] = request.AdditionalBroker
	}

	if !dara.IsNil(request.Broker) {
		query["Broker"] = request.Broker
	}

	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartBack2BackCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartBack2BackCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartBack2BackCallRequest
//
// @return StartBack2BackCallResponse
func (client *Client) StartBack2BackCall(request *StartBack2BackCallRequest) (_result *StartBack2BackCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartBack2BackCallResponse{}
	_body, _err := client.StartBack2BackCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - StartChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartChatResponse
func (client *Client) StartChatWithOptions(tmpReq *StartChatRequest, runtime *dara.RuntimeOptions) (_result *StartChatResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserList) {
		request.UserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserList, dara.String("UserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessChannelId) {
		query["AccessChannelId"] = request.AccessChannelId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserListShrink) {
		query["UserList"] = request.UserListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartChat"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartChatRequest
//
// @return StartChatResponse
func (client *Client) StartChat(request *StartChatRequest) (_result *StartChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartChatResponse{}
	_body, _err := client.StartChatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartConferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartConferenceResponse
func (client *Client) StartConferenceWithOptions(request *StartConferenceRequest, runtime *dara.RuntimeOptions) (_result *StartConferenceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ParticipantListJson) {
		query["ParticipantListJson"] = request.ParticipantListJson
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartConference"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartConferenceRequest
//
// @return StartConferenceResponse
func (client *Client) StartConference(request *StartConferenceRequest) (_result *StartConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartConferenceResponse{}
	_body, _err := client.StartConferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartEditContactFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartEditContactFlowResponse
func (client *Client) StartEditContactFlowWithOptions(request *StartEditContactFlowRequest, runtime *dara.RuntimeOptions) (_result *StartEditContactFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartEditContactFlow"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartEditContactFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartEditContactFlowRequest
//
// @return StartEditContactFlowResponse
func (client *Client) StartEditContactFlow(request *StartEditContactFlowRequest) (_result *StartEditContactFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartEditContactFlowResponse{}
	_body, _err := client.StartEditContactFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartPredictiveCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPredictiveCallResponse
func (client *Client) StartPredictiveCallWithOptions(request *StartPredictiveCallRequest, runtime *dara.RuntimeOptions) (_result *StartPredictiveCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.ContactFlowVariables) {
		query["ContactFlowVariables"] = request.ContactFlowVariables
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaskedCallee) {
		query["MaskedCallee"] = request.MaskedCallee
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartPredictiveCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartPredictiveCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartPredictiveCallRequest
//
// @return StartPredictiveCallResponse
func (client *Client) StartPredictiveCall(request *StartPredictiveCallRequest) (_result *StartPredictiveCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartPredictiveCallResponse{}
	_body, _err := client.StartPredictiveCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发起隐私呼叫
//
// @param request - StartPrivacyCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPrivacyCallResponse
func (client *Client) StartPrivacyCallWithOptions(request *StartPrivacyCallRequest, runtime *dara.RuntimeOptions) (_result *StartPrivacyCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartPrivacyCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartPrivacyCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起隐私呼叫
//
// @param request - StartPrivacyCallRequest
//
// @return StartPrivacyCallResponse
func (client *Client) StartPrivacyCall(request *StartPrivacyCallRequest) (_result *StartPrivacyCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartPrivacyCallResponse{}
	_body, _err := client.StartPrivacyCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交预测式外呼活动
//
// @param request - SubmitCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCampaignResponse
func (client *Client) SubmitCampaignWithOptions(request *SubmitCampaignRequest, runtime *dara.RuntimeOptions) (_result *SubmitCampaignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交预测式外呼活动
//
// @param request - SubmitCampaignRequest
//
// @return SubmitCampaignResponse
func (client *Client) SubmitCampaign(request *SubmitCampaignRequest) (_result *SubmitCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitCampaignResponse{}
	_body, _err := client.SubmitCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SwitchToConferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchToConferenceResponse
func (client *Client) SwitchToConferenceWithOptions(request *SwitchToConferenceRequest, runtime *dara.RuntimeOptions) (_result *SwitchToConferenceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchToConference"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchToConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SwitchToConferenceRequest
//
// @return SwitchToConferenceResponse
func (client *Client) SwitchToConference(request *SwitchToConferenceRequest) (_result *SwitchToConferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchToConferenceResponse{}
	_body, _err := client.SwitchToConferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - TakeBreakRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TakeBreakResponse
func (client *Client) TakeBreakWithOptions(request *TakeBreakRequest, runtime *dara.RuntimeOptions) (_result *TakeBreakResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TakeBreak"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TakeBreakResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TakeBreakRequest
//
// @return TakeBreakResponse
func (client *Client) TakeBreak(request *TakeBreakRequest) (_result *TakeBreakResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TakeBreakResponse{}
	_body, _err := client.TakeBreakWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - TerminateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateTicketResponse
func (client *Client) TerminateTicketWithOptions(request *TerminateTicketRequest, runtime *dara.RuntimeOptions) (_result *TerminateTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminateTicket"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TerminateTicketRequest
//
// @return TerminateTicketResponse
func (client *Client) TerminateTicket(request *TerminateTicketRequest) (_result *TerminateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TerminateTicketResponse{}
	_body, _err := client.TerminateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - TransferTicketTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferTicketTaskResponse
func (client *Client) TransferTicketTaskWithOptions(request *TransferTicketTaskRequest, runtime *dara.RuntimeOptions) (_result *TransferTicketTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Assignee) {
		query["Assignee"] = request.Assignee
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferTicketTask"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferTicketTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TransferTicketTaskRequest
//
// @return TransferTicketTaskResponse
func (client *Client) TransferTicketTask(request *TransferTicketTaskRequest) (_result *TransferTicketTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferTicketTaskResponse{}
	_body, _err := client.TransferTicketTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UnmuteCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnmuteCallResponse
func (client *Client) UnmuteCallWithOptions(request *UnmuteCallRequest, runtime *dara.RuntimeOptions) (_result *UnmuteCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnmuteCall"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnmuteCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UnmuteCallRequest
//
// @return UnmuteCallResponse
func (client *Client) UnmuteCall(request *UnmuteCallRequest) (_result *UnmuteCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnmuteCallResponse{}
	_body, _err := client.UnmuteCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除注册设备
//
// @param request - UnregisterDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnregisterDeviceResponse
func (client *Client) UnregisterDeviceWithOptions(request *UnregisterDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnregisterDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnregisterDevice"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnregisterDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除注册设备
//
// @param request - UnregisterDeviceRequest
//
// @return UnregisterDeviceResponse
func (client *Client) UnregisterDevice(request *UnregisterDeviceRequest) (_result *UnregisterDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnregisterDeviceResponse{}
	_body, _err := client.UnregisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCallSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCallSummaryResponse
func (client *Client) UpdateCallSummaryWithOptions(request *UpdateCallSummaryRequest, runtime *dara.RuntimeOptions) (_result *UpdateCallSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCallSummary"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCallSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCallSummaryRequest
//
// @return UpdateCallSummaryResponse
func (client *Client) UpdateCallSummary(request *UpdateCallSummaryRequest) (_result *UpdateCallSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCallSummaryResponse{}
	_body, _err := client.UpdateCallSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update campaign
//
// @param request - UpdateCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCampaignResponse
func (client *Client) UpdateCampaignWithOptions(request *UpdateCampaignRequest, runtime *dara.RuntimeOptions) (_result *UpdateCampaignResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallableTime) {
		query["CallableTime"] = request.CallableTime
	}

	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StrategyParameters) {
		query["StrategyParameters"] = request.StrategyParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update campaign
//
// @param request - UpdateCampaignRequest
//
// @return UpdateCampaignResponse
func (client *Client) UpdateCampaign(request *UpdateCampaignRequest) (_result *UpdateCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCampaignResponse{}
	_body, _err := client.UpdateCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateChatRoutingProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChatRoutingProfileResponse
func (client *Client) UpdateChatRoutingProfileWithOptions(request *UpdateChatRoutingProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateChatRoutingProfileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RoutingProfiles) {
		query["RoutingProfiles"] = request.RoutingProfiles
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChatRoutingProfile"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChatRoutingProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateChatRoutingProfileRequest
//
// @return UpdateChatRoutingProfileResponse
func (client *Client) UpdateChatRoutingProfile(request *UpdateChatRoutingProfileRequest) (_result *UpdateChatRoutingProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateChatRoutingProfileResponse{}
	_body, _err := client.UpdateChatRoutingProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateConfigItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigItemsResponse
func (client *Client) UpdateConfigItemsWithOptions(request *UpdateConfigItemsRequest, runtime *dara.RuntimeOptions) (_result *UpdateConfigItemsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigItems) {
		query["ConfigItems"] = request.ConfigItems
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfigItems"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateConfigItemsRequest
//
// @return UpdateConfigItemsResponse
func (client *Client) UpdateConfigItems(request *UpdateConfigItemsRequest) (_result *UpdateConfigItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConfigItemsResponse{}
	_body, _err := client.UpdateConfigItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateSchemaPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSchemaPropertyResponse
func (client *Client) UpdateSchemaPropertyWithOptions(tmpReq *UpdateSchemaPropertyRequest, runtime *dara.RuntimeOptions) (_result *UpdateSchemaPropertyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSchemaPropertyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Property) {
		request.PropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Property, dara.String("Property"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PropertyShrink) {
		body["Property"] = request.PropertyShrink
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSchemaProperty"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSchemaPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSchemaPropertyRequest
//
// @return UpdateSchemaPropertyResponse
func (client *Client) UpdateSchemaProperty(request *UpdateSchemaPropertyRequest) (_result *UpdateSchemaPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSchemaPropertyResponse{}
	_body, _err := client.UpdateSchemaPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketResponse
func (client *Client) UpdateTicketWithOptions(request *UpdateTicketRequest, runtime *dara.RuntimeOptions) (_result *UpdateTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		query["Context"] = request.Context
	}

	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTicket"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateTicketRequest
//
// @return UpdateTicketResponse
func (client *Client) UpdateTicket(request *UpdateTicketRequest) (_result *UpdateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTicketResponse{}
	_body, _err := client.UpdateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - WithdrawTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawTicketResponse
func (client *Client) WithdrawTicketWithOptions(request *WithdrawTicketRequest, runtime *dara.RuntimeOptions) (_result *WithdrawTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WithdrawTicket"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WithdrawTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - WithdrawTicketRequest
//
// @return WithdrawTicketResponse
func (client *Client) WithdrawTicket(request *WithdrawTicketRequest) (_result *WithdrawTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &WithdrawTicketResponse{}
	_body, _err := client.WithdrawTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
