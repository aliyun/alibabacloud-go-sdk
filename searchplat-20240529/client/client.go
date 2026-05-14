// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("Searchplat")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("")
	return nil
}

// Summary:
//
// 创建语音转录异步任务
//
// @param request - CreateAudioAsrTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAudioAsrTaskResponse
func (client *Client) CreateAudioAsrTaskWithOptions(workspaceName *string, serviceId *string, request *CreateAudioAsrTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAudioAsrTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAudioAsrTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/audio-asr/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAudioAsrTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建语音转录异步任务
//
// @param request - CreateAudioAsrTaskRequest
//
// @return CreateAudioAsrTaskResponse
func (client *Client) CreateAudioAsrTask(workspaceName *string, serviceId *string, request *CreateAudioAsrTaskRequest) (_result *CreateAudioAsrTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAudioAsrTaskResponse{}
	_body, _err := client.CreateAudioAsrTaskWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文档解析异步提取任务
//
// @param request - CreateDocumentAnalyzeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocumentAnalyzeTaskResponse
func (client *Client) CreateDocumentAnalyzeTaskWithOptions(workspaceName *string, serviceId *string, request *CreateDocumentAnalyzeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDocumentAnalyzeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Document) {
		body["document"] = request.Document
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocumentAnalyzeTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/document-analyze/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocumentAnalyzeTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文档解析异步提取任务
//
// @param request - CreateDocumentAnalyzeTaskRequest
//
// @return CreateDocumentAnalyzeTaskResponse
func (client *Client) CreateDocumentAnalyzeTask(workspaceName *string, serviceId *string, request *CreateDocumentAnalyzeTaskRequest) (_result *CreateDocumentAnalyzeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDocumentAnalyzeTaskResponse{}
	_body, _err := client.CreateDocumentAnalyzeTaskWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建图片解析异步提取任务
//
// @param request - CreateImageAnalyzeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageAnalyzeTaskResponse
func (client *Client) CreateImageAnalyzeTaskWithOptions(workspaceName *string, serviceId *string, request *CreateImageAnalyzeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateImageAnalyzeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Document) {
		body["document"] = request.Document
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageAnalyzeTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/image-analyze/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageAnalyzeTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图片解析异步提取任务
//
// @param request - CreateImageAnalyzeTaskRequest
//
// @return CreateImageAnalyzeTaskResponse
func (client *Client) CreateImageAnalyzeTask(workspaceName *string, serviceId *string, request *CreateImageAnalyzeTaskRequest) (_result *CreateImageAnalyzeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateImageAnalyzeTaskResponse{}
	_body, _err := client.CreateImageAnalyzeTaskWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 存储 Memory 内容
//
// @param request - CreateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryResponse
func (client *Client) CreateMemoryWithOptions(workspaceName *string, serviceId *string, request *CreateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.Enhancements) {
		body["enhancements"] = request.Enhancements
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.RunId) {
		body["run_id"] = request.RunId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemory"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/memories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 存储 Memory 内容
//
// @param request - CreateMemoryRequest
//
// @return CreateMemoryResponse
func (client *Client) CreateMemory(workspaceName *string, serviceId *string, request *CreateMemoryRequest) (_result *CreateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryResponse{}
	_body, _err := client.CreateMemoryWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传skill
//
// @param request - CreateMemorySkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemorySkillResponse
func (client *Client) CreateMemorySkillWithOptions(workspaceName *string, serviceId *string, request *CreateMemorySkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemorySkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.ZipBase64) {
		body["zip_base64"] = request.ZipBase64
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemorySkill"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/skills"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemorySkillResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传skill
//
// @param request - CreateMemorySkillRequest
//
// @return CreateMemorySkillResponse
func (client *Client) CreateMemorySkill(workspaceName *string, serviceId *string, request *CreateMemorySkillRequest) (_result *CreateMemorySkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemorySkillResponse{}
	_body, _err := client.CreateMemorySkillWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建视频切割异步任务
//
// @param request - CreateVideoSegmentationTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoSegmentationTaskResponse
func (client *Client) CreateVideoSegmentationTaskWithOptions(workspaceName *string, serviceId *string, request *CreateVideoSegmentationTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVideoSegmentationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoSegmentationTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/video-segmentation/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoSegmentationTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建视频切割异步任务
//
// @param request - CreateVideoSegmentationTaskRequest
//
// @return CreateVideoSegmentationTaskResponse
func (client *Client) CreateVideoSegmentationTask(workspaceName *string, serviceId *string, request *CreateVideoSegmentationTaskRequest) (_result *CreateVideoSegmentationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVideoSegmentationTaskResponse{}
	_body, _err := client.CreateVideoSegmentationTaskWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建语音转录异步任务
//
// @param request - CreateVideoSnapshotTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoSnapshotTaskResponse
func (client *Client) CreateVideoSnapshotTaskWithOptions(workspaceName *string, serviceId *string, request *CreateVideoSnapshotTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVideoSnapshotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoSnapshotTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/video-snapshot/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoSnapshotTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建语音转录异步任务
//
// @param request - CreateVideoSnapshotTaskRequest
//
// @return CreateVideoSnapshotTaskResponse
func (client *Client) CreateVideoSnapshotTask(workspaceName *string, serviceId *string, request *CreateVideoSnapshotTaskRequest) (_result *CreateVideoSnapshotTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVideoSnapshotTaskResponse{}
	_body, _err := client.CreateVideoSnapshotTaskWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建视频总结异步任务
//
// @param request - CreateVideoSummarizationTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoSummarizationTaskResponse
func (client *Client) CreateVideoSummarizationTaskWithOptions(workspaceName *string, serviceId *string, request *CreateVideoSummarizationTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVideoSummarizationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoSummarizationTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/video-summarization/" + dara.StringValue(serviceId) + "/async"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoSummarizationTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建视频总结异步任务
//
// @param request - CreateVideoSummarizationTaskRequest
//
// @return CreateVideoSummarizationTaskResponse
func (client *Client) CreateVideoSummarizationTask(workspaceName *string, serviceId *string, request *CreateVideoSummarizationTaskRequest) (_result *CreateVideoSummarizationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVideoSummarizationTaskResponse{}
	_body, _err := client.CreateVideoSummarizationTaskWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除一条 Memory
//
// @param request - DeleteMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithOptions(workspaceName *string, serviceId *string, memoryId *string, request *DeleteMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemory"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/memories/" + dara.StringValue(memoryId)),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一条 Memory
//
// @param request - DeleteMemoryRequest
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemory(workspaceName *string, serviceId *string, memoryId *string, request *DeleteMemoryRequest) (_result *DeleteMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryResponse{}
	_body, _err := client.DeleteMemoryWithOptions(workspaceName, serviceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除skill
//
// @param request - DeleteMemorySkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemorySkillResponse
func (client *Client) DeleteMemorySkillWithOptions(workspaceName *string, serviceId *string, skillId *string, request *DeleteMemorySkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemorySkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemorySkill"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/skills/" + dara.StringValue(skillId)),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemorySkillResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除skill
//
// @param request - DeleteMemorySkillRequest
//
// @return DeleteMemorySkillResponse
func (client *Client) DeleteMemorySkill(workspaceName *string, serviceId *string, skillId *string, request *DeleteMemorySkillRequest) (_result *DeleteMemorySkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemorySkillResponse{}
	_body, _err := client.DeleteMemorySkillWithOptions(workspaceName, serviceId, skillId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取视频截帧异步提取任务状态
//
// @param request - GetAudioAsrTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAudioAsrTaskStatusResponse
func (client *Client) GetAudioAsrTaskStatusWithOptions(workspaceName *string, serviceId *string, request *GetAudioAsrTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAudioAsrTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAudioAsrTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/audio-asr/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAudioAsrTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频截帧异步提取任务状态
//
// @param request - GetAudioAsrTaskStatusRequest
//
// @return GetAudioAsrTaskStatusResponse
func (client *Client) GetAudioAsrTaskStatus(workspaceName *string, serviceId *string, request *GetAudioAsrTaskStatusRequest) (_result *GetAudioAsrTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAudioAsrTaskStatusResponse{}
	_body, _err := client.GetAudioAsrTaskStatusWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档解析异步提取任务状态
//
// @param request - GetDocumentAnalyzeTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentAnalyzeTaskStatusResponse
func (client *Client) GetDocumentAnalyzeTaskStatusWithOptions(workspaceName *string, serviceId *string, request *GetDocumentAnalyzeTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentAnalyzeTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentAnalyzeTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/document-analyze/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentAnalyzeTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档解析异步提取任务状态
//
// @param request - GetDocumentAnalyzeTaskStatusRequest
//
// @return GetDocumentAnalyzeTaskStatusResponse
func (client *Client) GetDocumentAnalyzeTaskStatus(workspaceName *string, serviceId *string, request *GetDocumentAnalyzeTaskStatusRequest) (_result *GetDocumentAnalyzeTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentAnalyzeTaskStatusResponse{}
	_body, _err := client.GetDocumentAnalyzeTaskStatusWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档相关性打分
//
// @param request - GetDocumentRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentRankResponse
func (client *Client) GetDocumentRankWithOptions(workspaceName *string, serviceId *string, request *GetDocumentRankRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentRankResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Docs) {
		body["docs"] = request.Docs
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentRank"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/ranker/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentRankResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档相关性打分
//
// @param request - GetDocumentRankRequest
//
// @return GetDocumentRankResponse
func (client *Client) GetDocumentRank(workspaceName *string, serviceId *string, request *GetDocumentRankRequest) (_result *GetDocumentRankResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentRankResponse{}
	_body, _err := client.GetDocumentRankWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档切片
//
// @param request - GetDocumentSplitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentSplitResponse
func (client *Client) GetDocumentSplitWithOptions(workspaceName *string, serviceId *string, request *GetDocumentSplitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentSplitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Document) {
		body["document"] = request.Document
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentSplit"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/document-split/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentSplitResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档切片
//
// @param request - GetDocumentSplitRequest
//
// @return GetDocumentSplitResponse
func (client *Client) GetDocumentSplit(workspaceName *string, serviceId *string, request *GetDocumentSplitRequest) (_result *GetDocumentSplitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentSplitResponse{}
	_body, _err := client.GetDocumentSplitWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 向量微调
//
// @param request - GetEmbeddingTuningRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmbeddingTuningResponse
func (client *Client) GetEmbeddingTuningWithOptions(workspaceName *string, serviceId *string, request *GetEmbeddingTuningRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmbeddingTuningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmbeddingTuning"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/embedding-tuning/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmbeddingTuningResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向量微调
//
// @param request - GetEmbeddingTuningRequest
//
// @return GetEmbeddingTuningResponse
func (client *Client) GetEmbeddingTuning(workspaceName *string, serviceId *string, request *GetEmbeddingTuningRequest) (_result *GetEmbeddingTuningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmbeddingTuningResponse{}
	_body, _err := client.GetEmbeddingTuningWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图片解析异步提取任务状态
//
// @param request - GetImageAnalyzeTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageAnalyzeTaskStatusResponse
func (client *Client) GetImageAnalyzeTaskStatusWithOptions(workspaceName *string, serviceId *string, request *GetImageAnalyzeTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetImageAnalyzeTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageAnalyzeTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/image-analyze/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageAnalyzeTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片解析异步提取任务状态
//
// @param request - GetImageAnalyzeTaskStatusRequest
//
// @return GetImageAnalyzeTaskStatusResponse
func (client *Client) GetImageAnalyzeTaskStatus(workspaceName *string, serviceId *string, request *GetImageAnalyzeTaskStatusRequest) (_result *GetImageAnalyzeTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImageAnalyzeTaskStatusResponse{}
	_body, _err := client.GetImageAnalyzeTaskStatusWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片主体检测
//
// @param request - GetImageObjectDetectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageObjectDetectionResponse
func (client *Client) GetImageObjectDetectionWithOptions(workspaceName *string, serviceId *string, request *GetImageObjectDetectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetImageObjectDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Image) {
		body["image"] = request.Image
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageObjectDetection"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/image-object-detection/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageObjectDetectionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片主体检测
//
// @param request - GetImageObjectDetectionRequest
//
// @return GetImageObjectDetectionResponse
func (client *Client) GetImageObjectDetection(workspaceName *string, serviceId *string, request *GetImageObjectDetectionRequest) (_result *GetImageObjectDetectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImageObjectDetectionResponse{}
	_body, _err := client.GetImageObjectDetectionWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取JinaAiReader解析结果
//
// @param request - GetJinaAiReaderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJinaAiReaderResponse
func (client *Client) GetJinaAiReaderWithOptions(workspaceName *string, serviceId *string, request *GetJinaAiReaderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJinaAiReaderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJinaAiReader"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/jina-ai-reader/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJinaAiReaderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取JinaAiReader解析结果
//
// @param request - GetJinaAiReaderRequest
//
// @return GetJinaAiReaderResponse
func (client *Client) GetJinaAiReader(workspaceName *string, serviceId *string, request *GetJinaAiReaderRequest) (_result *GetJinaAiReaderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJinaAiReaderResponse{}
	_body, _err := client.GetJinaAiReaderWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看memory详情
//
// @param request - GetMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryResponse
func (client *Client) GetMemoryWithOptions(workspaceName *string, serviceId *string, memoryId *string, request *GetMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemory"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/memories/" + dara.StringValue(memoryId)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看memory详情
//
// @param request - GetMemoryRequest
//
// @return GetMemoryResponse
func (client *Client) GetMemory(workspaceName *string, serviceId *string, memoryId *string, request *GetMemoryRequest) (_result *GetMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryResponse{}
	_body, _err := client.GetMemoryWithOptions(workspaceName, serviceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查 Memory 服务健康状态
//
// @param request - GetMemoryHealthRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryHealthResponse
func (client *Client) GetMemoryHealthWithOptions(workspaceName *string, serviceId *string, request *GetMemoryHealthRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryHealthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemoryHealth"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/health"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryHealthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查 Memory 服务健康状态
//
// @param request - GetMemoryHealthRequest
//
// @return GetMemoryHealthResponse
func (client *Client) GetMemoryHealth(workspaceName *string, serviceId *string, request *GetMemoryHealthRequest) (_result *GetMemoryHealthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryHealthResponse{}
	_body, _err := client.GetMemoryHealthWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看skill详情
//
// @param request - GetMemorySkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemorySkillResponse
func (client *Client) GetMemorySkillWithOptions(workspaceName *string, serviceId *string, skillId *string, request *GetMemorySkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemorySkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemorySkill"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/skills/" + dara.StringValue(skillId)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemorySkillResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看skill详情
//
// @param request - GetMemorySkillRequest
//
// @return GetMemorySkillResponse
func (client *Client) GetMemorySkill(workspaceName *string, serviceId *string, skillId *string, request *GetMemorySkillRequest) (_result *GetMemorySkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemorySkillResponse{}
	_body, _err := client.GetMemorySkillWithOptions(workspaceName, serviceId, skillId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询memory异步任务的处理状态
//
// @param request - GetMemoryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryTaskResponse
func (client *Client) GetMemoryTaskWithOptions(workspaceName *string, serviceId *string, taskId *string, request *GetMemoryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemoryTask"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/tasks/" + dara.StringValue(taskId)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询memory异步任务的处理状态
//
// @param request - GetMemoryTaskRequest
//
// @return GetMemoryTaskResponse
func (client *Client) GetMemoryTask(workspaceName *string, serviceId *string, taskId *string, request *GetMemoryTaskRequest) (_result *GetMemoryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryTaskResponse{}
	_body, _err := client.GetMemoryTaskWithOptions(workspaceName, serviceId, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多模态向量化
//
// @param request - GetMultiModalEmbeddingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiModalEmbeddingResponse
func (client *Client) GetMultiModalEmbeddingWithOptions(workspaceName *string, serviceId *string, request *GetMultiModalEmbeddingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMultiModalEmbeddingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiModalEmbedding"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/multi-modal-embedding/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiModalEmbeddingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态向量化
//
// @param request - GetMultiModalEmbeddingRequest
//
// @return GetMultiModalEmbeddingResponse
func (client *Client) GetMultiModalEmbedding(workspaceName *string, serviceId *string, request *GetMultiModalEmbeddingRequest) (_result *GetMultiModalEmbeddingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMultiModalEmbeddingResponse{}
	_body, _err := client.GetMultiModalEmbeddingWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 多模态 Reranker
//
// @param request - GetMultiModalRerankerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiModalRerankerResponse
func (client *Client) GetMultiModalRerankerWithOptions(workspaceName *string, serviceId *string, request *GetMultiModalRerankerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMultiModalRerankerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Docs) {
		body["docs"] = request.Docs
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultiModalReranker"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/multi-modal-reranker/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultiModalRerankerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 多模态 Reranker
//
// @param request - GetMultiModalRerankerRequest
//
// @return GetMultiModalRerankerResponse
func (client *Client) GetMultiModalReranker(workspaceName *string, serviceId *string, request *GetMultiModalRerankerRequest) (_result *GetMultiModalRerankerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMultiModalRerankerResponse{}
	_body, _err := client.GetMultiModalRerankerWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取推理结果
//
// @param request - GetPredictionRequest
//
// @param headers - GetPredictionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPredictionResponse
func (client *Client) GetPredictionWithOptions(deploymentId *string, request *GetPredictionRequest, headers *GetPredictionHeaders, runtime *dara.RuntimeOptions) (_result *GetPredictionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Token) {
		realHeaders["Token"] = dara.String(dara.ToString(dara.StringValue(headers.Token)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    request.Body,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrediction"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/deployments/" + dara.StringValue(deploymentId) + "/predict"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &GetPredictionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取推理结果
//
// @param request - GetPredictionRequest
//
// @return GetPredictionResponse
func (client *Client) GetPrediction(deploymentId *string, request *GetPredictionRequest) (_result *GetPredictionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetPredictionHeaders{}
	_result = &GetPredictionResponse{}
	_body, _err := client.GetPredictionWithOptions(deploymentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取query分析结果
//
// @param request - GetQueryAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryAnalysisResponse
func (client *Client) GetQueryAnalysisWithOptions(workspaceName *string, serviceId *string, request *GetQueryAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQueryAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Functions) {
		body["functions"] = request.Functions
	}

	if !dara.IsNil(request.History) {
		body["history"] = request.History
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueryAnalysis"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/query-analyze/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryAnalysisResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取query分析结果
//
// @param request - GetQueryAnalysisRequest
//
// @return GetQueryAnalysisResponse
func (client *Client) GetQueryAnalysis(workspaceName *string, serviceId *string, request *GetQueryAnalysisRequest) (_result *GetQueryAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQueryAnalysisResponse{}
	_body, _err := client.GetQueryAnalysisWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文本向量化
//
// @param request - GetTextEmbeddingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextEmbeddingResponse
func (client *Client) GetTextEmbeddingWithOptions(workspaceName *string, serviceId *string, request *GetTextEmbeddingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextEmbeddingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.InputType) {
		body["input_type"] = request.InputType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextEmbedding"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/text-embedding/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextEmbeddingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本向量化
//
// @param request - GetTextEmbeddingRequest
//
// @return GetTextEmbeddingResponse
func (client *Client) GetTextEmbedding(workspaceName *string, serviceId *string, request *GetTextEmbeddingRequest) (_result *GetTextEmbeddingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextEmbeddingResponse{}
	_body, _err := client.GetTextEmbeddingWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 大模型问答
//
// @param request - GetTextGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextGenerationResponse
func (client *Client) GetTextGenerationWithOptions(workspaceName *string, serviceId *string, request *GetTextGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CsiLevel) {
		body["csi_level"] = request.CsiLevel
	}

	if !dara.IsNil(request.EnableSearch) {
		body["enable_search"] = request.EnableSearch
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextGeneration"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/text-generation/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextGenerationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 大模型问答
//
// @param request - GetTextGenerationRequest
//
// @return GetTextGenerationResponse
func (client *Client) GetTextGeneration(workspaceName *string, serviceId *string, request *GetTextGenerationRequest) (_result *GetTextGenerationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextGenerationResponse{}
	_body, _err := client.GetTextGenerationWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文本稀疏向量化
//
// @param request - GetTextSparseEmbeddingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextSparseEmbeddingResponse
func (client *Client) GetTextSparseEmbeddingWithOptions(workspaceName *string, serviceId *string, request *GetTextSparseEmbeddingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextSparseEmbeddingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Input) {
		body["input"] = request.Input
	}

	if !dara.IsNil(request.InputType) {
		body["input_type"] = request.InputType
	}

	if !dara.IsNil(request.ReturnToken) {
		body["return_token"] = request.ReturnToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextSparseEmbedding"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/text-sparse-embedding/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextSparseEmbeddingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本稀疏向量化
//
// @param request - GetTextSparseEmbeddingRequest
//
// @return GetTextSparseEmbeddingResponse
func (client *Client) GetTextSparseEmbedding(workspaceName *string, serviceId *string, request *GetTextSparseEmbeddingRequest) (_result *GetTextSparseEmbeddingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextSparseEmbeddingResponse{}
	_body, _err := client.GetTextSparseEmbeddingWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取视频切割异步任务状态
//
// @param request - GetVideoSegmentationTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoSegmentationTaskStatusResponse
func (client *Client) GetVideoSegmentationTaskStatusWithOptions(workspaceName *string, serviceId *string, request *GetVideoSegmentationTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoSegmentationTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoSegmentationTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/video-segmentation/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoSegmentationTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频切割异步任务状态
//
// @param request - GetVideoSegmentationTaskStatusRequest
//
// @return GetVideoSegmentationTaskStatusResponse
func (client *Client) GetVideoSegmentationTaskStatus(workspaceName *string, serviceId *string, request *GetVideoSegmentationTaskStatusRequest) (_result *GetVideoSegmentationTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoSegmentationTaskStatusResponse{}
	_body, _err := client.GetVideoSegmentationTaskStatusWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取视频截帧异步提取任务状态
//
// @param request - GetVideoSnapshotTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoSnapshotTaskStatusResponse
func (client *Client) GetVideoSnapshotTaskStatusWithOptions(workspaceName *string, serviceId *string, request *GetVideoSnapshotTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoSnapshotTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoSnapshotTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/video-snapshot/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoSnapshotTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频截帧异步提取任务状态
//
// @param request - GetVideoSnapshotTaskStatusRequest
//
// @return GetVideoSnapshotTaskStatusResponse
func (client *Client) GetVideoSnapshotTaskStatus(workspaceName *string, serviceId *string, request *GetVideoSnapshotTaskStatusRequest) (_result *GetVideoSnapshotTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoSnapshotTaskStatusResponse{}
	_body, _err := client.GetVideoSnapshotTaskStatusWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取视频总结异步任务状态
//
// @param request - GetVideoSummarizationTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoSummarizationTaskStatusResponse
func (client *Client) GetVideoSummarizationTaskStatusWithOptions(workspaceName *string, serviceId *string, request *GetVideoSummarizationTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoSummarizationTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoSummarizationTaskStatus"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/video-summarization/" + dara.StringValue(serviceId) + "/async/task-status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoSummarizationTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频总结异步任务状态
//
// @param request - GetVideoSummarizationTaskStatusRequest
//
// @return GetVideoSummarizationTaskStatusResponse
func (client *Client) GetVideoSummarizationTaskStatus(workspaceName *string, serviceId *string, request *GetVideoSummarizationTaskStatusRequest) (_result *GetVideoSummarizationTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoSummarizationTaskStatusResponse{}
	_body, _err := client.GetVideoSummarizationTaskStatusWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 联网搜索
//
// @param request - GetWebSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWebSearchResponse
func (client *Client) GetWebSearchWithOptions(workspaceName *string, serviceId *string, request *GetWebSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWebSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["content_type"] = request.ContentType
	}

	if !dara.IsNil(request.History) {
		body["history"] = request.History
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryRewrite) {
		body["query_rewrite"] = request.QueryRewrite
	}

	if !dara.IsNil(request.TopK) {
		body["top_k"] = request.TopK
	}

	if !dara.IsNil(request.Way) {
		body["way"] = request.Way
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWebSearch"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/web-search/" + dara.StringValue(serviceId)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWebSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 联网搜索
//
// @param request - GetWebSearchRequest
//
// @return GetWebSearchResponse
func (client *Client) GetWebSearch(workspaceName *string, serviceId *string, request *GetWebSearchRequest) (_result *GetWebSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWebSearchResponse{}
	_body, _err := client.GetWebSearchWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据查询条件搜索 Memory
//
// @param request - SearchMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMemoryResponse
func (client *Client) SearchMemoryWithOptions(workspaceName *string, serviceId *string, request *SearchMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.Enhancements) {
		body["enhancements"] = request.Enhancements
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.RunId) {
		body["run_id"] = request.RunId
	}

	if !dara.IsNil(request.Size) {
		body["size"] = request.Size
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMemory"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMemoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据查询条件搜索 Memory
//
// @param request - SearchMemoryRequest
//
// @return SearchMemoryResponse
func (client *Client) SearchMemory(workspaceName *string, serviceId *string, request *SearchMemoryRequest) (_result *SearchMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchMemoryResponse{}
	_body, _err := client.SearchMemoryWithOptions(workspaceName, serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Memory
//
// @param request - UpdateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithOptions(workspaceName *string, serviceId *string, memoryId *string, request *UpdateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Memory) {
		body["memory"] = request.Memory
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemory"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/memories/" + dara.StringValue(memoryId)),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Memory
//
// @param request - UpdateMemoryRequest
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemory(workspaceName *string, serviceId *string, memoryId *string, request *UpdateMemoryRequest) (_result *UpdateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryResponse{}
	_body, _err := client.UpdateMemoryWithOptions(workspaceName, serviceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Skill
//
// @param request - UpdateMemorySkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemorySkillResponse
func (client *Client) UpdateMemorySkillWithOptions(workspaceName *string, serviceId *string, skillId *string, request *UpdateMemorySkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemorySkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemorySkill"),
		Version:     dara.String("2024-05-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/memory/" + dara.StringValue(serviceId) + "/skills/" + dara.StringValue(skillId)),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemorySkillResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Skill
//
// @param request - UpdateMemorySkillRequest
//
// @return UpdateMemorySkillResponse
func (client *Client) UpdateMemorySkill(workspaceName *string, serviceId *string, skillId *string, request *UpdateMemorySkillRequest) (_result *UpdateMemorySkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemorySkillResponse{}
	_body, _err := client.UpdateMemorySkillWithOptions(workspaceName, serviceId, skillId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
