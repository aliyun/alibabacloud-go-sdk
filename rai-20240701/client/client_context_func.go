// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # BatchContentAsyncDetect
//
// @param request - BatchContentAsyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchContentAsyncDetectResponse
func (client *Client) BatchContentAsyncDetectWithContext(ctx context.Context, request *BatchContentAsyncDetectRequest, runtime *dara.RuntimeOptions) (_result *BatchContentAsyncDetectResponse, _err error) {
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

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceParameterList) {
		body["serviceParameterList"] = request.ServiceParameterList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchContentAsyncDetect"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchContentAsyncDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # BatchContentSyncDetect
//
// @param request - BatchContentSyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchContentSyncDetectResponse
func (client *Client) BatchContentSyncDetectWithContext(ctx context.Context, request *BatchContentSyncDetectRequest, runtime *dara.RuntimeOptions) (_result *BatchContentSyncDetectResponse, _err error) {
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

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceParameterList) {
		body["serviceParameterList"] = request.ServiceParameterList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchContentSyncDetect"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchContentSyncDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查用户是否开通RAI服务
//
// @param request - CheckAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountResponse
func (client *Client) CheckAccountWithContext(ctx context.Context, request *CheckAccountRequest, runtime *dara.RuntimeOptions) (_result *CheckAccountResponse, _err error) {
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
		Action:      dara.String("CheckAccount"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ContentAsyncDetect
//
// @param request - ContentAsyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContentAsyncDetectResponse
func (client *Client) ContentAsyncDetectWithContext(ctx context.Context, request *ContentAsyncDetectRequest, runtime *dara.RuntimeOptions) (_result *ContentAsyncDetectResponse, _err error) {
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

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceParameter) {
		body["serviceParameter"] = request.ServiceParameter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContentAsyncDetect"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ContentAsyncDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ContentSyncDetect
//
// @param request - ContentSyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContentSyncDetectResponse
func (client *Client) ContentSyncDetectWithContext(ctx context.Context, request *ContentSyncDetectRequest, runtime *dara.RuntimeOptions) (_result *ContentSyncDetectResponse, _err error) {
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

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceParameter) {
		body["serviceParameter"] = request.ServiceParameter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContentSyncDetect"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ContentSyncDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateModelInstance
//
// @param request - CreateModelInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelInstanceResponse
func (client *Client) CreateModelInstanceWithContext(ctx context.Context, request *CreateModelInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateModelInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EasServiceId) {
		query["EasServiceId"] = request.EasServiceId
	}

	if !dara.IsNil(request.EasServiceName) {
		query["EasServiceName"] = request.EasServiceName
	}

	if !dara.IsNil(request.ModelCallName) {
		query["ModelCallName"] = request.ModelCallName
	}

	if !dara.IsNil(request.ModelCategoryId) {
		query["ModelCategoryId"] = request.ModelCategoryId
	}

	if !dara.IsNil(request.ModelToken) {
		query["ModelToken"] = request.ModelToken
	}

	if !dara.IsNil(request.ModelUrl) {
		query["ModelUrl"] = request.ModelUrl
	}

	if !dara.IsNil(request.ModelVpcUrl) {
		query["ModelVpcUrl"] = request.ModelVpcUrl
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelInstance"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreatePolicy
//
// @param tmpReq - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithContext(ctx context.Context, tmpReq *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HarmfulCategoryConfigInfoList) {
		request.HarmfulCategoryConfigInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmfulCategoryConfigInfoList, dara.String("HarmfulCategoryConfigInfoList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PromptAttackInfo) {
		request.PromptAttackInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PromptAttackInfo, dara.String("PromptAttackInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PromptAttackInfoList) {
		request.PromptAttackInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PromptAttackInfoList, dara.String("PromptAttackInfoList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegularExpressList) {
		request.RegularExpressListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegularExpressList, dara.String("RegularExpressList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SensitiveConfigList) {
		request.SensitiveConfigListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SensitiveConfigList, dara.String("SensitiveConfigList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SensitiveTopicList) {
		request.SensitiveTopicListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SensitiveTopicList, dara.String("SensitiveTopicList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SensitiveWordList) {
		request.SensitiveWordListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SensitiveWordList, dara.String("SensitiveWordList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TopicConfigInfoList) {
		request.TopicConfigInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TopicConfigInfoList, dara.String("TopicConfigInfoList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WordGroupInfoList) {
		request.WordGroupInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WordGroupInfoList, dara.String("WordGroupInfoList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentSafeModelInstanceId) {
		query["ContentSafeModelInstanceId"] = request.ContentSafeModelInstanceId
	}

	if !dara.IsNil(request.EnableSensitiveInputCheck) {
		query["EnableSensitiveInputCheck"] = request.EnableSensitiveInputCheck
	}

	if !dara.IsNil(request.EnableSensitiveOutputCheck) {
		query["EnableSensitiveOutputCheck"] = request.EnableSensitiveOutputCheck
	}

	if !dara.IsNil(request.HarmfulCategoryConfigInfoListShrink) {
		query["HarmfulCategoryConfigInfoList"] = request.HarmfulCategoryConfigInfoListShrink
	}

	if !dara.IsNil(request.InputSafeAnswer) {
		query["InputSafeAnswer"] = request.InputSafeAnswer
	}

	if !dara.IsNil(request.InputSafeAnswerSwitch) {
		query["InputSafeAnswerSwitch"] = request.InputSafeAnswerSwitch
	}

	if !dara.IsNil(request.IsSidecarPolicy) {
		query["IsSidecarPolicy"] = request.IsSidecarPolicy
	}

	if !dara.IsNil(request.OutputSafeAnswer) {
		query["OutputSafeAnswer"] = request.OutputSafeAnswer
	}

	if !dara.IsNil(request.OutputSafeAnswerSwitch) {
		query["OutputSafeAnswerSwitch"] = request.OutputSafeAnswerSwitch
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PromptAttackInfoShrink) {
		query["PromptAttackInfo"] = request.PromptAttackInfoShrink
	}

	if !dara.IsNil(request.PromptAttackInfoListShrink) {
		query["PromptAttackInfoList"] = request.PromptAttackInfoListShrink
	}

	if !dara.IsNil(request.PromptAttackModelInstanceId) {
		query["PromptAttackModelInstanceId"] = request.PromptAttackModelInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegularExpressListShrink) {
		query["RegularExpressList"] = request.RegularExpressListShrink
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	if !dara.IsNil(request.SensitiveConfigListShrink) {
		query["SensitiveConfigList"] = request.SensitiveConfigListShrink
	}

	if !dara.IsNil(request.SensitiveTopicListShrink) {
		query["SensitiveTopicList"] = request.SensitiveTopicListShrink
	}

	if !dara.IsNil(request.SensitiveTopicModelInstanceId) {
		query["SensitiveTopicModelInstanceId"] = request.SensitiveTopicModelInstanceId
	}

	if !dara.IsNil(request.SensitiveWordListShrink) {
		query["SensitiveWordList"] = request.SensitiveWordListShrink
	}

	if !dara.IsNil(request.TopicConfigInfoListShrink) {
		query["TopicConfigInfoList"] = request.TopicConfigInfoListShrink
	}

	if !dara.IsNil(request.WordGroupInfoListShrink) {
		query["WordGroupInfoList"] = request.WordGroupInfoListShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateTopic
//
// @param tmpReq - CreateTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTopicResponse
func (client *Client) CreateTopicWithContext(ctx context.Context, tmpReq *CreateTopicRequest, runtime *dara.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTopicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BodyData) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, dara.String("BodyData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TopicDefinition) {
		query["TopicDefinition"] = request.TopicDefinition
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyDataShrink) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTopic"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateWordGroup
//
// @param tmpReq - CreateWordGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWordGroupResponse
func (client *Client) CreateWordGroupWithContext(ctx context.Context, tmpReq *CreateWordGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateWordGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWordGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BodyData) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, dara.String("BodyData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Commit) {
		query["Commit"] = request.Commit
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyDataShrink) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWordGroup"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWordGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteModelInstance
//
// @param tmpReq - DeleteModelInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelInstanceResponse
func (client *Client) DeleteModelInstanceWithContext(ctx context.Context, tmpReq *DeleteModelInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteModelInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ModelInstanceIdList) {
		request.ModelInstanceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModelInstanceIdList, dara.String("ModelInstanceIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelInstanceIdListShrink) {
		query["ModelInstanceIdList"] = request.ModelInstanceIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelInstance"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeletePolicy
//
// @param tmpReq - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithContext(ctx context.Context, tmpReq *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeletePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PolicyIdList) {
		request.PolicyIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PolicyIdList, dara.String("PolicyIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIdListShrink) {
		query["PolicyIdList"] = request.PolicyIdListShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteTopic
//
// @param tmpReq - DeleteTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopicWithContext(ctx context.Context, tmpReq *DeleteTopicRequest, runtime *dara.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteTopicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TopicIdList) {
		request.TopicIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TopicIdList, dara.String("TopicIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TopicIdListShrink) {
		query["TopicIdList"] = request.TopicIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTopic"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteWordGroup
//
// @param tmpReq - DeleteWordGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWordGroupResponse
func (client *Client) DeleteWordGroupWithContext(ctx context.Context, tmpReq *DeleteWordGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteWordGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteWordGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GroupIdList) {
		request.GroupIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupIdList, dara.String("GroupIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupIdListShrink) {
		query["GroupIdList"] = request.GroupIdListShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWordGroup"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWordGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetContentDetectResult
//
// @param request - GetContentDetectResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContentDetectResultResponse
func (client *Client) GetContentDetectResultWithContext(ctx context.Context, request *GetContentDetectResultRequest, runtime *dara.RuntimeOptions) (_result *GetContentDetectResultResponse, _err error) {
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
		Action:      dara.String("GetContentDetectResult"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContentDetectResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetModelInputContentDetectResult
//
// @param request - GetModelInputContentDetectResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelInputContentDetectResultResponse
func (client *Client) GetModelInputContentDetectResultWithContext(ctx context.Context, request *GetModelInputContentDetectResultRequest, runtime *dara.RuntimeOptions) (_result *GetModelInputContentDetectResultResponse, _err error) {
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
		Action:      dara.String("GetModelInputContentDetectResult"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelInputContentDetectResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetModelInstanceInfo
//
// @param request - GetModelInstanceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelInstanceInfoResponse
func (client *Client) GetModelInstanceInfoWithContext(ctx context.Context, request *GetModelInstanceInfoRequest, runtime *dara.RuntimeOptions) (_result *GetModelInstanceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelInstanceId) {
		query["ModelInstanceId"] = request.ModelInstanceId
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelInstanceInfo"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelInstanceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetModelOutputContentDetectResult
//
// @param request - GetModelOutputContentDetectResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelOutputContentDetectResultResponse
func (client *Client) GetModelOutputContentDetectResultWithContext(ctx context.Context, request *GetModelOutputContentDetectResultRequest, runtime *dara.RuntimeOptions) (_result *GetModelOutputContentDetectResultResponse, _err error) {
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
		Action:      dara.String("GetModelOutputContentDetectResult"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelOutputContentDetectResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetPolicyDefaultOptions
//
// @param request - GetPolicyDefaultOptionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyDefaultOptionsResponse
func (client *Client) GetPolicyDefaultOptionsWithContext(ctx context.Context, request *GetPolicyDefaultOptionsRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyDefaultOptionsResponse, _err error) {
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
		Action:      dara.String("GetPolicyDefaultOptions"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyDefaultOptionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetPolicyInfo
//
// @param request - GetPolicyInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyInfoResponse
func (client *Client) GetPolicyInfoWithContext(ctx context.Context, request *GetPolicyInfoRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicyInfo"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetTopic
//
// @param request - GetTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicResponse
func (client *Client) GetTopicWithContext(ctx context.Context, request *GetTopicRequest, runtime *dara.RuntimeOptions) (_result *GetTopicResponse, _err error) {
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

	if !dara.IsNil(request.TopicId) {
		query["TopicId"] = request.TopicId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopic"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetWordGroup
//
// @param request - GetWordGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWordGroupResponse
func (client *Client) GetWordGroupWithContext(ctx context.Context, request *GetWordGroupRequest, runtime *dara.RuntimeOptions) (_result *GetWordGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWordGroup"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWordGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListModelCategory
//
// @param request - ListModelCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelCategoryResponse
func (client *Client) ListModelCategoryWithContext(ctx context.Context, request *ListModelCategoryRequest, runtime *dara.RuntimeOptions) (_result *ListModelCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentSafeImageSupported) {
		query["ContentSafeImageSupported"] = request.ContentSafeImageSupported
	}

	if !dara.IsNil(request.ContentSafeTextSupported) {
		query["ContentSafeTextSupported"] = request.ContentSafeTextSupported
	}

	if !dara.IsNil(request.ModelCategoryName) {
		query["ModelCategoryName"] = request.ModelCategoryName
	}

	if !dara.IsNil(request.ModelSource) {
		query["ModelSource"] = request.ModelSource
	}

	if !dara.IsNil(request.PromptAttackTextSupported) {
		query["PromptAttackTextSupported"] = request.PromptAttackTextSupported
	}

	if !dara.IsNil(request.SensitiveTopicTextSupported) {
		query["SensitiveTopicTextSupported"] = request.SensitiveTopicTextSupported
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelCategory"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListModelInstance
//
// @param request - ListModelInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelInstanceResponse
func (client *Client) ListModelInstanceWithContext(ctx context.Context, request *ListModelInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListModelInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EasServiceName) {
		query["EasServiceName"] = request.EasServiceName
	}

	if !dara.IsNil(request.IsSidecarPolicy) {
		query["IsSidecarPolicy"] = request.IsSidecarPolicy
	}

	if !dara.IsNil(request.IsSupportContentSafe) {
		query["IsSupportContentSafe"] = request.IsSupportContentSafe
	}

	if !dara.IsNil(request.IsSupportPromptAttack) {
		query["IsSupportPromptAttack"] = request.IsSupportPromptAttack
	}

	if !dara.IsNil(request.IsSupportSensitiveTopic) {
		query["IsSupportSensitiveTopic"] = request.IsSupportSensitiveTopic
	}

	if !dara.IsNil(request.ModelSource) {
		query["ModelSource"] = request.ModelSource
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelInstance"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListPolicy
//
// @param request - ListPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyResponse
func (client *Client) ListPolicyWithContext(ctx context.Context, request *ListPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsSidecarPolicy) {
		query["IsSidecarPolicy"] = request.IsSidecarPolicy
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyIdentifier) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicy"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListTopic
//
// @param request - ListTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicResponse
func (client *Client) ListTopicWithContext(ctx context.Context, request *ListTopicRequest, runtime *dara.RuntimeOptions) (_result *ListTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTopic"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListWordGroup
//
// @param request - ListWordGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWordGroupResponse
func (client *Client) ListWordGroupWithContext(ctx context.Context, request *ListWordGroupRequest, runtime *dara.RuntimeOptions) (_result *ListWordGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWordGroup"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWordGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModelInputContentAsyncDetect
//
// @param tmpReq - ModelInputContentAsyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelInputContentAsyncDetectResponse
func (client *Client) ModelInputContentAsyncDetectWithContext(ctx context.Context, tmpReq *ModelInputContentAsyncDetectRequest, runtime *dara.RuntimeOptions) (_result *ModelInputContentAsyncDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModelInputContentAsyncDetectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BodyData) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, dara.String("BodyData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIdentifier) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyDataShrink) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelInputContentAsyncDetect"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelInputContentAsyncDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModelInputContentSyncDetect
//
// @param tmpReq - ModelInputContentSyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelInputContentSyncDetectResponse
func (client *Client) ModelInputContentSyncDetectWithContext(ctx context.Context, tmpReq *ModelInputContentSyncDetectRequest, runtime *dara.RuntimeOptions) (_result *ModelInputContentSyncDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModelInputContentSyncDetectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BodyData) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, dara.String("BodyData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIdentifier) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyDataShrink) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelInputContentSyncDetect"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelInputContentSyncDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModelOutputContentAsyncDetect
//
// @param tmpReq - ModelOutputContentAsyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelOutputContentAsyncDetectResponse
func (client *Client) ModelOutputContentAsyncDetectWithContext(ctx context.Context, tmpReq *ModelOutputContentAsyncDetectRequest, runtime *dara.RuntimeOptions) (_result *ModelOutputContentAsyncDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModelOutputContentAsyncDetectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BodyData) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, dara.String("BodyData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIdentifier) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyDataShrink) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelOutputContentAsyncDetect"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelOutputContentAsyncDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModelOutputContentSyncDetect
//
// @param tmpReq - ModelOutputContentSyncDetectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModelOutputContentSyncDetectResponse
func (client *Client) ModelOutputContentSyncDetectWithContext(ctx context.Context, tmpReq *ModelOutputContentSyncDetectRequest, runtime *dara.RuntimeOptions) (_result *ModelOutputContentSyncDetectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModelOutputContentSyncDetectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BodyData) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, dara.String("BodyData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIdentifier) {
		query["PolicyIdentifier"] = request.PolicyIdentifier
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SceneName) {
		query["SceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyDataShrink) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModelOutputContentSyncDetect"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModelOutputContentSyncDetectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册RAI账号
//
// @param request - RegisterAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterAccountResponse
func (client *Client) RegisterAccountWithContext(ctx context.Context, request *RegisterAccountRequest, runtime *dara.RuntimeOptions) (_result *RegisterAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Memo) {
		query["Memo"] = request.Memo
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterAccount"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateModelInstance
//
// @param request - UpdateModelInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelInstanceResponse
func (client *Client) UpdateModelInstanceWithContext(ctx context.Context, request *UpdateModelInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateModelInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EasServiceId) {
		query["EasServiceId"] = request.EasServiceId
	}

	if !dara.IsNil(request.EasServiceName) {
		query["EasServiceName"] = request.EasServiceName
	}

	if !dara.IsNil(request.ModelCallName) {
		query["ModelCallName"] = request.ModelCallName
	}

	if !dara.IsNil(request.ModelCategoryId) {
		query["ModelCategoryId"] = request.ModelCategoryId
	}

	if !dara.IsNil(request.ModelInstanceId) {
		query["ModelInstanceId"] = request.ModelInstanceId
	}

	if !dara.IsNil(request.ModelToken) {
		query["ModelToken"] = request.ModelToken
	}

	if !dara.IsNil(request.ModelUrl) {
		query["ModelUrl"] = request.ModelUrl
	}

	if !dara.IsNil(request.ModelVpcUrl) {
		query["ModelVpcUrl"] = request.ModelVpcUrl
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelInstance"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdatePolicy
//
// @param tmpReq - UpdatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicyWithContext(ctx context.Context, tmpReq *UpdatePolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HarmfulCategoryConfigInfoList) {
		request.HarmfulCategoryConfigInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmfulCategoryConfigInfoList, dara.String("HarmfulCategoryConfigInfoList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PromptAttackInfo) {
		request.PromptAttackInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PromptAttackInfo, dara.String("PromptAttackInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PromptAttackInfoList) {
		request.PromptAttackInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PromptAttackInfoList, dara.String("PromptAttackInfoList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegularExpressList) {
		request.RegularExpressListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegularExpressList, dara.String("RegularExpressList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SensitiveConfigList) {
		request.SensitiveConfigListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SensitiveConfigList, dara.String("SensitiveConfigList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SensitiveTopicList) {
		request.SensitiveTopicListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SensitiveTopicList, dara.String("SensitiveTopicList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SensitiveWordList) {
		request.SensitiveWordListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SensitiveWordList, dara.String("SensitiveWordList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TopicConfigInfoList) {
		request.TopicConfigInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TopicConfigInfoList, dara.String("TopicConfigInfoList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WordGroupInfoList) {
		request.WordGroupInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WordGroupInfoList, dara.String("WordGroupInfoList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentSafeModelInstanceId) {
		query["ContentSafeModelInstanceId"] = request.ContentSafeModelInstanceId
	}

	if !dara.IsNil(request.EnableSensitiveInputCheck) {
		query["EnableSensitiveInputCheck"] = request.EnableSensitiveInputCheck
	}

	if !dara.IsNil(request.EnableSensitiveOutputCheck) {
		query["EnableSensitiveOutputCheck"] = request.EnableSensitiveOutputCheck
	}

	if !dara.IsNil(request.HarmfulCategoryConfigInfoListShrink) {
		query["HarmfulCategoryConfigInfoList"] = request.HarmfulCategoryConfigInfoListShrink
	}

	if !dara.IsNil(request.InputSafeAnswer) {
		query["InputSafeAnswer"] = request.InputSafeAnswer
	}

	if !dara.IsNil(request.InputSafeAnswerSwitch) {
		query["InputSafeAnswerSwitch"] = request.InputSafeAnswerSwitch
	}

	if !dara.IsNil(request.IsSidecarPolicy) {
		query["IsSidecarPolicy"] = request.IsSidecarPolicy
	}

	if !dara.IsNil(request.OutputSafeAnswer) {
		query["OutputSafeAnswer"] = request.OutputSafeAnswer
	}

	if !dara.IsNil(request.OutputSafeAnswerSwitch) {
		query["OutputSafeAnswerSwitch"] = request.OutputSafeAnswerSwitch
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PromptAttackInfoShrink) {
		query["PromptAttackInfo"] = request.PromptAttackInfoShrink
	}

	if !dara.IsNil(request.PromptAttackInfoListShrink) {
		query["PromptAttackInfoList"] = request.PromptAttackInfoListShrink
	}

	if !dara.IsNil(request.PromptAttackModelInstanceId) {
		query["PromptAttackModelInstanceId"] = request.PromptAttackModelInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegularExpressListShrink) {
		query["RegularExpressList"] = request.RegularExpressListShrink
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	if !dara.IsNil(request.SensitiveConfigListShrink) {
		query["SensitiveConfigList"] = request.SensitiveConfigListShrink
	}

	if !dara.IsNil(request.SensitiveTopicListShrink) {
		query["SensitiveTopicList"] = request.SensitiveTopicListShrink
	}

	if !dara.IsNil(request.SensitiveTopicModelInstanceId) {
		query["SensitiveTopicModelInstanceId"] = request.SensitiveTopicModelInstanceId
	}

	if !dara.IsNil(request.SensitiveWordListShrink) {
		query["SensitiveWordList"] = request.SensitiveWordListShrink
	}

	if !dara.IsNil(request.TopicConfigInfoListShrink) {
		query["TopicConfigInfoList"] = request.TopicConfigInfoListShrink
	}

	if !dara.IsNil(request.WordGroupInfoListShrink) {
		query["WordGroupInfoList"] = request.WordGroupInfoListShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolicy"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateTopic
//
// @param tmpReq - UpdateTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTopicResponse
func (client *Client) UpdateTopicWithContext(ctx context.Context, tmpReq *UpdateTopicRequest, runtime *dara.RuntimeOptions) (_result *UpdateTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTopicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BodyData) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, dara.String("BodyData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TopicDefinition) {
		query["TopicDefinition"] = request.TopicDefinition
	}

	if !dara.IsNil(request.TopicId) {
		query["TopicId"] = request.TopicId
	}

	if !dara.IsNil(request.TopicName) {
		query["TopicName"] = request.TopicName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyDataShrink) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTopic"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateWordGroup
//
// @param tmpReq - UpdateWordGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWordGroupResponse
func (client *Client) UpdateWordGroupWithContext(ctx context.Context, tmpReq *UpdateWordGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateWordGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWordGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BodyData) {
		request.BodyDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BodyData, dara.String("BodyData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Commit) {
		query["Commit"] = request.Commit
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WordInfoListModified) {
		query["WordInfoListModified"] = request.WordInfoListModified
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyDataShrink) {
		body["BodyData"] = request.BodyDataShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWordGroup"),
		Version:     dara.String("2024-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWordGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
