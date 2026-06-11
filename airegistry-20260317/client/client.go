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
	client.Endpoint, _err = client.GetEndpoint(dara.String("airegistry"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建 AI Registry 命名空间
//
// @param request - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ScanPolicy) {
		query["ScanPolicy"] = request.ScanPolicy
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNamespace"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 AI Registry 命名空间
//
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建 Prompt
//
// @param request - CreatePromptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePromptResponse
func (client *Client) CreatePromptWithOptions(request *CreatePromptRequest, runtime *dara.RuntimeOptions) (_result *CreatePromptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizTags) {
		query["BizTags"] = request.BizTags
	}

	if !dara.IsNil(request.CommitMsg) {
		query["CommitMsg"] = request.CommitMsg
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	if !dara.IsNil(request.TargetVersion) {
		query["TargetVersion"] = request.TargetVersion
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.Variables) {
		query["Variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrompt"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePromptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 Prompt
//
// @param request - CreatePromptRequest
//
// @return CreatePromptResponse
func (client *Client) CreatePrompt(request *CreatePromptRequest) (_result *CreatePromptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePromptResponse{}
	_body, _err := client.CreatePromptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建 Prompt 草稿版本。Prompt 必须已存在，且当前没有正在编辑的草稿。只对草稿版本生效。
//
// @param request - CreatePromptVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePromptVersionResponse
func (client *Client) CreatePromptVersionWithOptions(request *CreatePromptVersionRequest, runtime *dara.RuntimeOptions) (_result *CreatePromptVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasedOnVersion) {
		query["BasedOnVersion"] = request.BasedOnVersion
	}

	if !dara.IsNil(request.CommitMsg) {
		query["CommitMsg"] = request.CommitMsg
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	if !dara.IsNil(request.TargetVersion) {
		query["TargetVersion"] = request.TargetVersion
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.Variables) {
		query["Variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePromptVersion"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePromptVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 Prompt 草稿版本。Prompt 必须已存在，且当前没有正在编辑的草稿。只对草稿版本生效。
//
// @param request - CreatePromptVersionRequest
//
// @return CreatePromptVersionResponse
func (client *Client) CreatePromptVersion(request *CreatePromptVersionRequest) (_result *CreatePromptVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePromptVersionResponse{}
	_body, _err := client.CreatePromptVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建 Skill 草稿版本
//
// @param request - CreateSkillDraftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillDraftResponse
func (client *Client) CreateSkillDraftWithOptions(request *CreateSkillDraftRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillDraftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasedOnVersion) {
		query["BasedOnVersion"] = request.BasedOnVersion
	}

	if !dara.IsNil(request.CommitMsg) {
		query["CommitMsg"] = request.CommitMsg
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillCard) {
		query["SkillCard"] = request.SkillCard
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.TargetVersion) {
		query["TargetVersion"] = request.TargetVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillDraft"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 Skill 草稿版本
//
// @param request - CreateSkillDraftRequest
//
// @return CreateSkillDraftResponse
func (client *Client) CreateSkillDraft(request *CreateSkillDraftRequest) (_result *CreateSkillDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillDraftResponse{}
	_body, _err := client.CreateSkillDraftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除命名空间
//
// @param request - DeleteNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNamespace"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除命名空间
//
// @param request - DeleteNamespaceRequest
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 Prompt
//
// @param request - DeletePromptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePromptResponse
func (client *Client) DeletePromptWithOptions(request *DeletePromptRequest, runtime *dara.RuntimeOptions) (_result *DeletePromptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrompt"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePromptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 Prompt
//
// @param request - DeletePromptRequest
//
// @return DeletePromptResponse
func (client *Client) DeletePrompt(request *DeletePromptRequest) (_result *DeletePromptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePromptResponse{}
	_body, _err := client.DeletePromptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 Skill
//
// @param request - DeleteSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkillWithOptions(request *DeleteSkillRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkill"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 Skill
//
// @param request - DeleteSkillRequest
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkill(request *DeleteSkillRequest) (_result *DeleteSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSkillResponse{}
	_body, _err := client.DeleteSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过 OSS 下载 Skill 版本 - 返回 OSS 下载 URL
//
// @param request - DownloadSkillVersionViaOssRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadSkillVersionViaOssResponse
func (client *Client) DownloadSkillVersionViaOssWithOptions(request *DownloadSkillVersionViaOssRequest, runtime *dara.RuntimeOptions) (_result *DownloadSkillVersionViaOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadSkillVersionViaOss"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadSkillVersionViaOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过 OSS 下载 Skill 版本 - 返回 OSS 下载 URL
//
// @param request - DownloadSkillVersionViaOssRequest
//
// @return DownloadSkillVersionViaOssResponse
func (client *Client) DownloadSkillVersionViaOss(request *DownloadSkillVersionViaOssRequest) (_result *DownloadSkillVersionViaOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadSkillVersionViaOssResponse{}
	_body, _err := client.DownloadSkillVersionViaOssWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 强制发布版本
//
// @param request - ForcePublishSkillVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForcePublishSkillVersionResponse
func (client *Client) ForcePublishSkillVersionWithOptions(request *ForcePublishSkillVersionRequest, runtime *dara.RuntimeOptions) (_result *ForcePublishSkillVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	if !dara.IsNil(request.UpdateLatestLabel) {
		query["UpdateLatestLabel"] = request.UpdateLatestLabel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForcePublishSkillVersion"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForcePublishSkillVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 强制发布版本
//
// @param request - ForcePublishSkillVersionRequest
//
// @return ForcePublishSkillVersionResponse
func (client *Client) ForcePublishSkillVersion(request *ForcePublishSkillVersionRequest) (_result *ForcePublishSkillVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ForcePublishSkillVersionResponse{}
	_body, _err := client.ForcePublishSkillVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取命名空间详细信息
//
// @param request - GetNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNamespaceResponse
func (client *Client) GetNamespaceWithOptions(request *GetNamespaceRequest, runtime *dara.RuntimeOptions) (_result *GetNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNamespace"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取命名空间详细信息
//
// @param request - GetNamespaceRequest
//
// @return GetNamespaceResponse
func (client *Client) GetNamespace(request *GetNamespaceRequest) (_result *GetNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNamespaceResponse{}
	_body, _err := client.GetNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 Prompt 详情信息
//
// @param request - GetPromptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPromptResponse
func (client *Client) GetPromptWithOptions(request *GetPromptRequest, runtime *dara.RuntimeOptions) (_result *GetPromptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrompt"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPromptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 Prompt 详情信息
//
// @param request - GetPromptRequest
//
// @return GetPromptResponse
func (client *Client) GetPrompt(request *GetPromptRequest) (_result *GetPromptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPromptResponse{}
	_body, _err := client.GetPromptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 Prompt 某个版本的信息
//
// @param request - GetPromptVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPromptVersionResponse
func (client *Client) GetPromptVersionWithOptions(request *GetPromptVersionRequest, runtime *dara.RuntimeOptions) (_result *GetPromptVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	if !dara.IsNil(request.PromptVersion) {
		query["PromptVersion"] = request.PromptVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPromptVersion"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPromptVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 Prompt 某个版本的信息
//
// @param request - GetPromptVersionRequest
//
// @return GetPromptVersionResponse
func (client *Client) GetPromptVersion(request *GetPromptVersionRequest) (_result *GetPromptVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPromptVersionResponse{}
	_body, _err := client.GetPromptVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 Skill 详情
//
// @param request - GetSkillDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillDetailResponse
func (client *Client) GetSkillDetailWithOptions(request *GetSkillDetailRequest, runtime *dara.RuntimeOptions) (_result *GetSkillDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillDetail"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 Skill 详情
//
// @param request - GetSkillDetailRequest
//
// @return GetSkillDetailResponse
func (client *Client) GetSkillDetail(request *GetSkillDetailRequest) (_result *GetSkillDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillDetailResponse{}
	_body, _err := client.GetSkillDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 Skill 导入用 OSS 上传 URL。客户端使用返回的 uploadUrl 执行 PUT 上传后，
//
// @param request - GetSkillImportFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillImportFileUrlResponse
func (client *Client) GetSkillImportFileUrlWithOptions(request *GetSkillImportFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetSkillImportFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillImportFileUrl"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillImportFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 Skill 导入用 OSS 上传 URL。客户端使用返回的 uploadUrl 执行 PUT 上传后，
//
// @param request - GetSkillImportFileUrlRequest
//
// @return GetSkillImportFileUrlResponse
func (client *Client) GetSkillImportFileUrl(request *GetSkillImportFileUrlRequest) (_result *GetSkillImportFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillImportFileUrlResponse{}
	_body, _err := client.GetSkillImportFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定版本详情
//
// @param request - GetSkillVersionDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillVersionDetailResponse
func (client *Client) GetSkillVersionDetailWithOptions(request *GetSkillVersionDetailRequest, runtime *dara.RuntimeOptions) (_result *GetSkillVersionDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkillVersionDetail"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillVersionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定版本详情
//
// @param request - GetSkillVersionDetailRequest
//
// @return GetSkillVersionDetailResponse
func (client *Client) GetSkillVersionDetail(request *GetSkillVersionDetailRequest) (_result *GetSkillVersionDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSkillVersionDetailResponse{}
	_body, _err := client.GetSkillVersionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取命名空间列表
//
// @param request - ListNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespacesResponse
func (client *Client) ListNamespacesWithOptions(request *ListNamespacesRequest, runtime *dara.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNamespaces"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取命名空间列表
//
// @param request - ListNamespacesRequest
//
// @return ListNamespacesResponse
func (client *Client) ListNamespaces(request *ListNamespacesRequest) (_result *ListNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNamespacesResponse{}
	_body, _err := client.ListNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出Prompt版本列表
//
// @param request - ListPromptVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPromptVersionsResponse
func (client *Client) ListPromptVersionsWithOptions(request *ListPromptVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListPromptVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPromptVersions"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPromptVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出Prompt版本列表
//
// @param request - ListPromptVersionsRequest
//
// @return ListPromptVersionsResponse
func (client *Client) ListPromptVersions(request *ListPromptVersionsRequest) (_result *ListPromptVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPromptVersionsResponse{}
	_body, _err := client.ListPromptVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Prompt列表
//
// @param request - ListPromptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPromptsResponse
func (client *Client) ListPromptsWithOptions(request *ListPromptsRequest, runtime *dara.RuntimeOptions) (_result *ListPromptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizTags) {
		query["BizTags"] = request.BizTags
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrompts"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPromptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Prompt列表
//
// @param request - ListPromptsRequest
//
// @return ListPromptsResponse
func (client *Client) ListPrompts(request *ListPromptsRequest) (_result *ListPromptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPromptsResponse{}
	_body, _err := client.ListPromptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出 Skills
//
// @param request - ListSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithOptions(request *ListSkillsRequest, runtime *dara.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkills"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出 Skills
//
// @param request - ListSkillsRequest
//
// @return ListSkillsResponse
func (client *Client) ListSkills(request *ListSkillsRequest) (_result *ListSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSkillsResponse{}
	_body, _err := client.ListSkillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下线版本
//
// @param request - OfflineSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineSkillResponse
func (client *Client) OfflineSkillWithOptions(request *OfflineSkillRequest, runtime *dara.RuntimeOptions) (_result *OfflineSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineSkill"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下线版本
//
// @param request - OfflineSkillRequest
//
// @return OfflineSkillResponse
func (client *Client) OfflineSkill(request *OfflineSkillRequest) (_result *OfflineSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineSkillResponse{}
	_body, _err := client.OfflineSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上线 Skill
//
// @param request - OnlineSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineSkillResponse
func (client *Client) OnlineSkillWithOptions(request *OnlineSkillRequest, runtime *dara.RuntimeOptions) (_result *OnlineSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineSkill"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线 Skill
//
// @param request - OnlineSkillRequest
//
// @return OnlineSkillResponse
func (client *Client) OnlineSkill(request *OnlineSkillRequest) (_result *OnlineSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OnlineSkillResponse{}
	_body, _err := client.OnlineSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布版本
//
// @param request - PublishSkillVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishSkillVersionResponse
func (client *Client) PublishSkillVersionWithOptions(request *PublishSkillVersionRequest, runtime *dara.RuntimeOptions) (_result *PublishSkillVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	if !dara.IsNil(request.UpdateLatestLabel) {
		query["UpdateLatestLabel"] = request.UpdateLatestLabel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishSkillVersion"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishSkillVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布版本
//
// @param request - PublishSkillVersionRequest
//
// @return PublishSkillVersionResponse
func (client *Client) PublishSkillVersion(request *PublishSkillVersionRequest) (_result *PublishSkillVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishSkillVersionResponse{}
	_body, _err := client.PublishSkillVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交 Prompt 版本, 将 Prompt 的草稿版本转化为正式版本
//
// @param request - SubmitPromptVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitPromptVersionResponse
func (client *Client) SubmitPromptVersionWithOptions(request *SubmitPromptVersionRequest, runtime *dara.RuntimeOptions) (_result *SubmitPromptVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	if !dara.IsNil(request.PromptVersion) {
		query["PromptVersion"] = request.PromptVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitPromptVersion"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitPromptVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交 Prompt 版本, 将 Prompt 的草稿版本转化为正式版本
//
// @param request - SubmitPromptVersionRequest
//
// @return SubmitPromptVersionResponse
func (client *Client) SubmitPromptVersion(request *SubmitPromptVersionRequest) (_result *SubmitPromptVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitPromptVersionResponse{}
	_body, _err := client.SubmitPromptVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交 Skill Draft 审核
//
// @param request - SubmitSkillVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSkillVersionResponse
func (client *Client) SubmitSkillVersionWithOptions(request *SubmitSkillVersionRequest, runtime *dara.RuntimeOptions) (_result *SubmitSkillVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSkillVersion"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSkillVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交 Skill Draft 审核
//
// @param request - SubmitSkillVersionRequest
//
// @return SubmitSkillVersionResponse
func (client *Client) SubmitSkillVersion(request *SubmitSkillVersionRequest) (_result *SubmitSkillVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSkillVersionResponse{}
	_body, _err := client.SubmitSkillVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新命名空间信息
//
// @param request - UpdateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespaceWithOptions(request *UpdateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.ScanPolicy) {
		query["ScanPolicy"] = request.ScanPolicy
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNamespace"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新命名空间信息
//
// @param request - UpdateNamespaceRequest
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespace(request *UpdateNamespaceRequest) (_result *UpdateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.UpdateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Prompt 元数据，支持同时更新 description、bizTags、labels。
//
// @param tmpReq - UpdatePromptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePromptResponse
func (client *Client) UpdatePromptWithOptions(tmpReq *UpdatePromptRequest, runtime *dara.RuntimeOptions) (_result *UpdatePromptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePromptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizTags) {
		request.BizTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizTags, dara.String("BizTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Labels) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, dara.String("Labels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizTagsShrink) {
		query["BizTags"] = request.BizTagsShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.LabelsShrink) {
		query["Labels"] = request.LabelsShrink
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrompt"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePromptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Prompt 元数据，支持同时更新 description、bizTags、labels。
//
// @param request - UpdatePromptRequest
//
// @return UpdatePromptResponse
func (client *Client) UpdatePrompt(request *UpdatePromptRequest) (_result *UpdatePromptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePromptResponse{}
	_body, _err := client.UpdatePromptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Prompt 草稿版本内容。只对草稿版本生效，已发布的版本不可修改。
//
// @param request - UpdatePromptVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePromptVersionResponse
func (client *Client) UpdatePromptVersionWithOptions(request *UpdatePromptVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdatePromptVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommitMsg) {
		query["CommitMsg"] = request.CommitMsg
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PromptKey) {
		query["PromptKey"] = request.PromptKey
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.Variables) {
		query["Variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePromptVersion"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePromptVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Prompt 草稿版本内容。只对草稿版本生效，已发布的版本不可修改。
//
// @param request - UpdatePromptVersionRequest
//
// @return UpdatePromptVersionResponse
func (client *Client) UpdatePromptVersion(request *UpdatePromptVersionRequest) (_result *UpdatePromptVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePromptVersionResponse{}
	_body, _err := client.UpdatePromptVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新业务标签
//
// @param request - UpdateSkillBizTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillBizTagsResponse
func (client *Client) UpdateSkillBizTagsWithOptions(request *UpdateSkillBizTagsRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillBizTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizTags) {
		query["BizTags"] = request.BizTags
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillBizTags"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillBizTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新业务标签
//
// @param request - UpdateSkillBizTagsRequest
//
// @return UpdateSkillBizTagsResponse
func (client *Client) UpdateSkillBizTags(request *UpdateSkillBizTagsRequest) (_result *UpdateSkillBizTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSkillBizTagsResponse{}
	_body, _err := client.UpdateSkillBizTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Draft
//
// @param request - UpdateSkillDraftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillDraftResponse
func (client *Client) UpdateSkillDraftWithOptions(request *UpdateSkillDraftRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillDraftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommitMsg) {
		query["CommitMsg"] = request.CommitMsg
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillCard) {
		query["SkillCard"] = request.SkillCard
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillDraft"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Draft
//
// @param request - UpdateSkillDraftRequest
//
// @return UpdateSkillDraftResponse
func (client *Client) UpdateSkillDraft(request *UpdateSkillDraftRequest) (_result *UpdateSkillDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSkillDraftResponse{}
	_body, _err := client.UpdateSkillDraftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新版本标签
//
// @param request - UpdateSkillLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillLabelsResponse
func (client *Client) UpdateSkillLabelsWithOptions(request *UpdateSkillLabelsRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillLabels"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新版本标签
//
// @param request - UpdateSkillLabelsRequest
//
// @return UpdateSkillLabelsResponse
func (client *Client) UpdateSkillLabels(request *UpdateSkillLabelsRequest) (_result *UpdateSkillLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSkillLabelsResponse{}
	_body, _err := client.UpdateSkillLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新可见性
//
// @param request - UpdateSkillScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillScopeResponse
func (client *Client) UpdateSkillScopeWithOptions(request *UpdateSkillScopeRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SkillName) {
		query["SkillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkillScope"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新可见性
//
// @param request - UpdateSkillScopeRequest
//
// @return UpdateSkillScopeResponse
func (client *Client) UpdateSkillScope(request *UpdateSkillScopeRequest) (_result *UpdateSkillScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSkillScopeResponse{}
	_body, _err := client.UpdateSkillScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过 OSS 上传 Skill (ZIP) - 从 OSS 拉取文件内容后上传到 Nacos
//
// @param request - UploadSkillViaOssRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadSkillViaOssResponse
func (client *Client) UploadSkillViaOssWithOptions(request *UploadSkillViaOssRequest, runtime *dara.RuntimeOptions) (_result *UploadSkillViaOssResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommitMsg) {
		query["CommitMsg"] = request.CommitMsg
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.OssObjectName) {
		query["OssObjectName"] = request.OssObjectName
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadSkillViaOss"),
		Version:     dara.String("2026-03-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadSkillViaOssResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过 OSS 上传 Skill (ZIP) - 从 OSS 拉取文件内容后上传到 Nacos
//
// @param request - UploadSkillViaOssRequest
//
// @return UploadSkillViaOssResponse
func (client *Client) UploadSkillViaOss(request *UploadSkillViaOssRequest) (_result *UploadSkillViaOssResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadSkillViaOssResponse{}
	_body, _err := client.UploadSkillViaOssWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
