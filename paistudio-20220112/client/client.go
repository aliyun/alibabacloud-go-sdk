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
		"cn-beijing":            dara.String("pai.cn-beijing.aliyuncs.com"),
		"cn-hangzhou":           dara.String("pai.cn-hangzhou.data.aliyun.com"),
		"cn-shanghai":           dara.String("pai.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           dara.String("pai.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":           dara.String("pai.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1":        dara.String("pai.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        dara.String("pai.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        dara.String("pai.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        dara.String("pai.ap-southeast-5.aliyuncs.com"),
		"us-east-1":             dara.String("pai.us-east-1.aliyuncs.com"),
		"us-west-1":             dara.String("pai.us-west-1.aliyuncs.com"),
		"eu-central-1":          dara.String("pai.eu-central-1.aliyuncs.com"),
		"ap-south-1":            dara.String("pai.ap-south-1.aliyuncs.com"),
		"me-east-1":             dara.String("pai.me-east-1.aliyuncs.com"),
		"ap-northeast-1":        dara.String("pai.ap-northeast-1.aliyuncs.com"),
		"cn-qingdao":            dara.String("pai.cn-qingdao.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("pai.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("pai.cn-wulanchabu.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("pai.cn-zhangjiakou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("paistudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 检查WebTerminal
//
// @param request - CheckInstanceWebTerminalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceWebTerminalResponse
func (client *Client) CheckInstanceWebTerminalWithOptions(TrainingJobId *string, InstanceId *string, request *CheckInstanceWebTerminalRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInstanceWebTerminalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckInfo) {
		body["CheckInfo"] = request.CheckInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckInstanceWebTerminal"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/webterminals/action/check"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInstanceWebTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查WebTerminal
//
// @param request - CheckInstanceWebTerminalRequest
//
// @return CheckInstanceWebTerminalResponse
func (client *Client) CheckInstanceWebTerminal(TrainingJobId *string, InstanceId *string, request *CheckInstanceWebTerminalRequest) (_result *CheckInstanceWebTerminalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInstanceWebTerminalResponse{}
	_body, _err := client.CheckInstanceWebTerminalWithOptions(TrainingJobId, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建新的算法
//
// @param request - CreateAlgorithmRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlgorithmResponse
func (client *Client) CreateAlgorithmWithOptions(request *CreateAlgorithmRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlgorithmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlgorithmDescription) {
		body["AlgorithmDescription"] = request.AlgorithmDescription
	}

	if !dara.IsNil(request.AlgorithmName) {
		body["AlgorithmName"] = request.AlgorithmName
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlgorithm"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建新的算法
//
// @param request - CreateAlgorithmRequest
//
// @return CreateAlgorithmResponse
func (client *Client) CreateAlgorithm(request *CreateAlgorithmRequest) (_result *CreateAlgorithmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAlgorithmResponse{}
	_body, _err := client.CreateAlgorithmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个新的算法版本
//
// @param tmpReq - CreateAlgorithmVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlgorithmVersionResponse
func (client *Client) CreateAlgorithmVersionWithOptions(AlgorithmId *string, AlgorithmVersion *string, tmpReq *CreateAlgorithmVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlgorithmVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAlgorithmVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlgorithmSpec) {
		request.AlgorithmSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlgorithmSpec, dara.String("AlgorithmSpec"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlgorithmSpecShrink) {
		body["AlgorithmSpec"] = request.AlgorithmSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlgorithmVersion"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms/" + dara.PercentEncode(dara.StringValue(AlgorithmId)) + "/versions/" + dara.PercentEncode(dara.StringValue(AlgorithmVersion))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlgorithmVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个新的算法版本
//
// @param request - CreateAlgorithmVersionRequest
//
// @return CreateAlgorithmVersionResponse
func (client *Client) CreateAlgorithmVersion(AlgorithmId *string, AlgorithmVersion *string, request *CreateAlgorithmVersionRequest) (_result *CreateAlgorithmVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAlgorithmVersionResponse{}
	_body, _err := client.CreateAlgorithmVersionWithOptions(AlgorithmId, AlgorithmVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建WebTerminal
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceWebTerminalResponse
func (client *Client) CreateInstanceWebTerminalWithOptions(TrainingJobId *string, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceWebTerminalResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceWebTerminal"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/webterminals"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceWebTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建WebTerminal
//
// @return CreateInstanceWebTerminalResponse
func (client *Client) CreateInstanceWebTerminal(TrainingJobId *string, InstanceId *string) (_result *CreateInstanceWebTerminalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceWebTerminalResponse{}
	_body, _err := client.CreateInstanceWebTerminalWithOptions(TrainingJobId, InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Quota
//
// @param request - CreateQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaResponse
func (client *Client) CreateQuotaWithOptions(request *CreateQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllocateStrategy) {
		body["AllocateStrategy"] = request.AllocateStrategy
	}

	if !dara.IsNil(request.ClusterSpec) {
		body["ClusterSpec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Min) {
		body["Min"] = request.Min
	}

	if !dara.IsNil(request.ParentQuotaId) {
		body["ParentQuotaId"] = request.ParentQuotaId
	}

	if !dara.IsNil(request.QueueStrategy) {
		body["QueueStrategy"] = request.QueueStrategy
	}

	if !dara.IsNil(request.QuotaConfig) {
		body["QuotaConfig"] = request.QuotaConfig
	}

	if !dara.IsNil(request.QuotaName) {
		body["QuotaName"] = request.QuotaName
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		body["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQuota"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Quota
//
// @param request - CreateQuotaRequest
//
// @return CreateQuotaResponse
func (client *Client) CreateQuota(request *CreateQuotaRequest) (_result *CreateQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQuotaResponse{}
	_body, _err := client.CreateQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源组
//
// @param request - CreateResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComputingResourceProvider) {
		body["ComputingResourceProvider"] = request.ComputingResourceProvider
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserVpc) {
		body["UserVpc"] = request.UserVpc
	}

	if !dara.IsNil(request.Version) {
		body["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceGroup"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源组
//
// @param request - CreateResourceGroupRequest
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建TrainingJob
//
// @param request - CreateTrainingJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrainingJobResponse
func (client *Client) CreateTrainingJobWithOptions(request *CreateTrainingJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrainingJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlgorithmName) {
		body["AlgorithmName"] = request.AlgorithmName
	}

	if !dara.IsNil(request.AlgorithmProvider) {
		body["AlgorithmProvider"] = request.AlgorithmProvider
	}

	if !dara.IsNil(request.AlgorithmSpec) {
		body["AlgorithmSpec"] = request.AlgorithmSpec
	}

	if !dara.IsNil(request.AlgorithmVersion) {
		body["AlgorithmVersion"] = request.AlgorithmVersion
	}

	if !dara.IsNil(request.AssignNodeSpec) {
		body["AssignNodeSpec"] = request.AssignNodeSpec
	}

	if !dara.IsNil(request.CodeDir) {
		body["CodeDir"] = request.CodeDir
	}

	if !dara.IsNil(request.ComputeResource) {
		body["ComputeResource"] = request.ComputeResource
	}

	if !dara.IsNil(request.Environments) {
		body["Environments"] = request.Environments
	}

	if !dara.IsNil(request.ExperimentConfig) {
		body["ExperimentConfig"] = request.ExperimentConfig
	}

	if !dara.IsNil(request.HyperParameters) {
		body["HyperParameters"] = request.HyperParameters
	}

	if !dara.IsNil(request.InputChannels) {
		body["InputChannels"] = request.InputChannels
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.OutputChannels) {
		body["OutputChannels"] = request.OutputChannels
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.PythonRequirements) {
		body["PythonRequirements"] = request.PythonRequirements
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.Scheduler) {
		body["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.Settings) {
		body["Settings"] = request.Settings
	}

	if !dara.IsNil(request.TrainingJobDescription) {
		body["TrainingJobDescription"] = request.TrainingJobDescription
	}

	if !dara.IsNil(request.TrainingJobName) {
		body["TrainingJobName"] = request.TrainingJobName
	}

	if !dara.IsNil(request.UserVpc) {
		body["UserVpc"] = request.UserVpc
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrainingJob"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建TrainingJob
//
// @param request - CreateTrainingJobRequest
//
// @return CreateTrainingJobResponse
func (client *Client) CreateTrainingJob(request *CreateTrainingJobRequest) (_result *CreateTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrainingJobResponse{}
	_body, _err := client.CreateTrainingJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除算法
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlgorithmResponse
func (client *Client) DeleteAlgorithmWithOptions(AlgorithmId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlgorithmResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlgorithm"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms/" + dara.PercentEncode(dara.StringValue(AlgorithmId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除算法
//
// @return DeleteAlgorithmResponse
func (client *Client) DeleteAlgorithm(AlgorithmId *string) (_result *DeleteAlgorithmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlgorithmResponse{}
	_body, _err := client.DeleteAlgorithmWithOptions(AlgorithmId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除算法版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlgorithmVersionResponse
func (client *Client) DeleteAlgorithmVersionWithOptions(AlgorithmId *string, AlgorithmVersion *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlgorithmVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlgorithmVersion"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms/" + dara.PercentEncode(dara.StringValue(AlgorithmId)) + "/versions/" + dara.PercentEncode(dara.StringValue(AlgorithmVersion))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlgorithmVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除算法版本
//
// @return DeleteAlgorithmVersionResponse
func (client *Client) DeleteAlgorithmVersion(AlgorithmId *string, AlgorithmVersion *string) (_result *DeleteAlgorithmVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAlgorithmVersionResponse{}
	_body, _err := client.DeleteAlgorithmVersionWithOptions(AlgorithmId, AlgorithmVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteMachineGroup is deprecated
//
// Summary:
//
// delete machine group
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMachineGroupResponse
func (client *Client) DeleteMachineGroupWithOptions(MachineGroupID *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMachineGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMachineGroup"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/machinegroups/" + dara.PercentEncode(dara.StringValue(MachineGroupID))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMachineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteMachineGroup is deprecated
//
// Summary:
//
// delete machine group
//
// @return DeleteMachineGroupResponse
// Deprecated
func (client *Client) DeleteMachineGroup(MachineGroupID *string) (_result *DeleteMachineGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMachineGroupResponse{}
	_body, _err := client.DeleteMachineGroupWithOptions(MachineGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Quota
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaResponse
func (client *Client) DeleteQuotaWithOptions(QuotaId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteQuotaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQuota"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(QuotaId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Quota
//
// @return DeleteQuotaResponse
func (client *Client) DeleteQuota(QuotaId *string) (_result *DeleteQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteQuotaResponse{}
	_body, _err := client.DeleteQuotaWithOptions(QuotaId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithOptions(ResourceGroupID *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceGroup"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/" + dara.PercentEncode(dara.StringValue(ResourceGroupID))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源组
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroup(ResourceGroupID *string) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(ResourceGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteResourceGroupMachineGroup is deprecated
//
// Summary:
//
// delete machine group
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupMachineGroupResponse
func (client *Client) DeleteResourceGroupMachineGroupWithOptions(MachineGroupID *string, ResourceGroupID *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupMachineGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceGroupMachineGroup"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/" + dara.PercentEncode(dara.StringValue(ResourceGroupID)) + "/machinegroups/" + dara.PercentEncode(dara.StringValue(MachineGroupID))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceGroupMachineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteResourceGroupMachineGroup is deprecated
//
// Summary:
//
// delete machine group
//
// @return DeleteResourceGroupMachineGroupResponse
// Deprecated
func (client *Client) DeleteResourceGroupMachineGroup(MachineGroupID *string, ResourceGroupID *string) (_result *DeleteResourceGroupMachineGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceGroupMachineGroupResponse{}
	_body, _err := client.DeleteResourceGroupMachineGroupWithOptions(MachineGroupID, ResourceGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除一个TrainingJob
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrainingJobResponse
func (client *Client) DeleteTrainingJobWithOptions(TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTrainingJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrainingJob"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个TrainingJob
//
// @return DeleteTrainingJobResponse
func (client *Client) DeleteTrainingJob(TrainingJobId *string) (_result *DeleteTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrainingJobResponse{}
	_body, _err := client.DeleteTrainingJobWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除TrainingJob的Labels
//
// @param request - DeleteTrainingJobLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrainingJobLabelsResponse
func (client *Client) DeleteTrainingJobLabelsWithOptions(TrainingJobId *string, request *DeleteTrainingJobLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTrainingJobLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keys) {
		query["Keys"] = request.Keys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrainingJobLabels"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/labels"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrainingJobLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除TrainingJob的Labels
//
// @param request - DeleteTrainingJobLabelsRequest
//
// @return DeleteTrainingJobLabelsResponse
func (client *Client) DeleteTrainingJobLabels(TrainingJobId *string, request *DeleteTrainingJobLabelsRequest) (_result *DeleteTrainingJobLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTrainingJobLabelsResponse{}
	_body, _err := client.DeleteTrainingJobLabelsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一个算法信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmResponse
func (client *Client) GetAlgorithmWithOptions(AlgorithmId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlgorithmResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlgorithm"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms/" + dara.PercentEncode(dara.StringValue(AlgorithmId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一个算法信息
//
// @return GetAlgorithmResponse
func (client *Client) GetAlgorithm(AlgorithmId *string) (_result *GetAlgorithmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgorithmResponse{}
	_body, _err := client.GetAlgorithmWithOptions(AlgorithmId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个新的算法版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmVersionResponse
func (client *Client) GetAlgorithmVersionWithOptions(AlgorithmId *string, AlgorithmVersion *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlgorithmVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlgorithmVersion"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms/" + dara.PercentEncode(dara.StringValue(AlgorithmId)) + "/versions/" + dara.PercentEncode(dara.StringValue(AlgorithmVersion))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlgorithmVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个新的算法版本
//
// @return GetAlgorithmVersionResponse
func (client *Client) GetAlgorithmVersion(AlgorithmId *string, AlgorithmVersion *string) (_result *GetAlgorithmVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAlgorithmVersionResponse{}
	_body, _err := client.GetAlgorithmVersionWithOptions(AlgorithmId, AlgorithmVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetMachineGroup is deprecated
//
// Summary:
//
// get machine group
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMachineGroupResponse
func (client *Client) GetMachineGroupWithOptions(MachineGroupID *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMachineGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMachineGroup"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/machinegroups/" + dara.PercentEncode(dara.StringValue(MachineGroupID))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMachineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetMachineGroup is deprecated
//
// Summary:
//
// get machine group
//
// @return GetMachineGroupResponse
// Deprecated
func (client *Client) GetMachineGroup(MachineGroupID *string) (_result *GetMachineGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMachineGroupResponse{}
	_body, _err := client.GetMachineGroupWithOptions(MachineGroupID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetNodeMetrics is deprecated
//
// Summary:
//
// get resource group node metrics
//
// @param request - GetNodeMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeMetricsResponse
func (client *Client) GetNodeMetricsWithOptions(ResourceGroupID *string, MetricType *string, request *GetNodeMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetNodeMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GPUType) {
		query["GPUType"] = request.GPUType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeStep) {
		query["TimeStep"] = request.TimeStep
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeMetrics"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/" + dara.PercentEncode(dara.StringValue(ResourceGroupID)) + "/nodemetrics/" + dara.PercentEncode(dara.StringValue(MetricType))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetNodeMetrics is deprecated
//
// Summary:
//
// get resource group node metrics
//
// @param request - GetNodeMetricsRequest
//
// @return GetNodeMetricsResponse
// Deprecated
func (client *Client) GetNodeMetrics(ResourceGroupID *string, MetricType *string, request *GetNodeMetricsRequest) (_result *GetNodeMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNodeMetricsResponse{}
	_body, _err := client.GetNodeMetricsWithOptions(ResourceGroupID, MetricType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Quota
//
// @param request - GetQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaResponse
func (client *Client) GetQuotaWithOptions(QuotaId *string, request *GetQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WithNodeMeta) {
		query["WithNodeMeta"] = request.WithNodeMeta
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuota"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(QuotaId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Quota
//
// @param request - GetQuotaRequest
//
// @return GetQuotaResponse
func (client *Client) GetQuota(QuotaId *string, request *GetQuotaRequest) (_result *GetQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get resource group by group id
//
// @param tmpReq - GetResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroupWithOptions(ResourceGroupID *string, tmpReq *GetResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IsAIWorkspaceDataEnabled) {
		query["IsAIWorkspaceDataEnabled"] = request.IsAIWorkspaceDataEnabled
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroup"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/" + dara.PercentEncode(dara.StringValue(ResourceGroupID))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get resource group by group id
//
// @param request - GetResourceGroupRequest
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroup(ResourceGroupID *string, request *GetResourceGroupRequest) (_result *GetResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupResponse{}
	_body, _err := client.GetResourceGroupWithOptions(ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get machine group
//
// @param tmpReq - GetResourceGroupMachineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupMachineGroupResponse
func (client *Client) GetResourceGroupMachineGroupWithOptions(MachineGroupID *string, ResourceGroupID *string, tmpReq *GetResourceGroupMachineGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupMachineGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetResourceGroupMachineGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroupMachineGroup"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/" + dara.PercentEncode(dara.StringValue(ResourceGroupID)) + "/machinegroups/" + dara.PercentEncode(dara.StringValue(MachineGroupID))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupMachineGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get machine group
//
// @param request - GetResourceGroupMachineGroupRequest
//
// @return GetResourceGroupMachineGroupResponse
func (client *Client) GetResourceGroupMachineGroup(MachineGroupID *string, ResourceGroupID *string, request *GetResourceGroupMachineGroupRequest) (_result *GetResourceGroupMachineGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupMachineGroupResponse{}
	_body, _err := client.GetResourceGroupMachineGroupWithOptions(MachineGroupID, ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetResourceGroupRequest is deprecated
//
// Summary:
//
// get resource group requested resource by resource group id
//
// @param request - GetResourceGroupRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupRequestResponse
func (client *Client) GetResourceGroupRequestWithOptions(request *GetResourceGroupRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PodStatus) {
		query["PodStatus"] = request.PodStatus
	}

	if !dara.IsNil(request.ResourceGroupID) {
		query["ResourceGroupID"] = request.ResourceGroupID
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroupRequest"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/data/request"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetResourceGroupRequest is deprecated
//
// Summary:
//
// get resource group requested resource by resource group id
//
// @param request - GetResourceGroupRequestRequest
//
// @return GetResourceGroupRequestResponse
// Deprecated
func (client *Client) GetResourceGroupRequest(request *GetResourceGroupRequestRequest) (_result *GetResourceGroupRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupRequestResponse{}
	_body, _err := client.GetResourceGroupRequestWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get resource group total resource by group id
//
// @param request - GetResourceGroupTotalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupTotalResponse
func (client *Client) GetResourceGroupTotalWithOptions(request *GetResourceGroupTotalRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupID) {
		query["ResourceGroupID"] = request.ResourceGroupID
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroupTotal"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/data/total"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get resource group total resource by group id
//
// @param request - GetResourceGroupTotalRequest
//
// @return GetResourceGroupTotalResponse
func (client *Client) GetResourceGroupTotal(request *GetResourceGroupTotalRequest) (_result *GetResourceGroupTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupTotalResponse{}
	_body, _err := client.GetResourceGroupTotalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取抢占式实例历史价格
//
// @param request - GetSpotPriceHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpotPriceHistoryResponse
func (client *Client) GetSpotPriceHistoryWithOptions(InstanceType *string, request *GetSpotPriceHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSpotPriceHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SpotDuration) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSpotPriceHistory"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/spots/" + dara.PercentEncode(dara.StringValue(InstanceType)) + "/pricehistory"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSpotPriceHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取抢占式实例历史价格
//
// @param request - GetSpotPriceHistoryRequest
//
// @return GetSpotPriceHistoryResponse
func (client *Client) GetSpotPriceHistory(InstanceType *string, request *GetSpotPriceHistoryRequest) (_result *GetSpotPriceHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSpotPriceHistoryResponse{}
	_body, _err := client.GetSpotPriceHistoryWithOptions(InstanceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用GetToken获取临时鉴权信息
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.TrainingJobId) {
		query["TrainingJobId"] = request.TrainingJobId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tokens"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用GetToken获取临时鉴权信息
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取TrainingJob的详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobResponse
func (client *Client) GetTrainingJobWithOptions(TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrainingJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrainingJob"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取TrainingJob的详情
//
// @return GetTrainingJobResponse
func (client *Client) GetTrainingJob(TrainingJobId *string) (_result *GetTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrainingJobResponse{}
	_body, _err := client.GetTrainingJobWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job的算法错误信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobErrorInfoResponse
func (client *Client) GetTrainingJobErrorInfoWithOptions(TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrainingJobErrorInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrainingJobErrorInfo"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/errorinfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrainingJobErrorInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Training Job的算法错误信息
//
// @return GetTrainingJobErrorInfoResponse
func (client *Client) GetTrainingJobErrorInfo(TrainingJobId *string) (_result *GetTrainingJobErrorInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrainingJobErrorInfoResponse{}
	_body, _err := client.GetTrainingJobErrorInfoWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取TrainingJob最近的Metrics
//
// @param request - GetTrainingJobLatestMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobLatestMetricsResponse
func (client *Client) GetTrainingJobLatestMetricsWithOptions(TrainingJobId *string, request *GetTrainingJobLatestMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrainingJobLatestMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Names) {
		query["Names"] = request.Names
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrainingJobLatestMetrics"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/latestmetrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrainingJobLatestMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取TrainingJob最近的Metrics
//
// @param request - GetTrainingJobLatestMetricsRequest
//
// @return GetTrainingJobLatestMetricsResponse
func (client *Client) GetTrainingJobLatestMetrics(TrainingJobId *string, request *GetTrainingJobLatestMetricsRequest) (_result *GetTrainingJobLatestMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrainingJobLatestMetricsResponse{}
	_body, _err := client.GetTrainingJobLatestMetricsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetUserViewMetrics is deprecated
//
// Summary:
//
// get user view  metrics
//
// @param request - GetUserViewMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserViewMetricsResponse
func (client *Client) GetUserViewMetricsWithOptions(ResourceGroupID *string, request *GetUserViewMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserViewMetricsResponse, _err error) {
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TimeStep) {
		query["TimeStep"] = request.TimeStep
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserViewMetrics"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/" + dara.PercentEncode(dara.StringValue(ResourceGroupID)) + "/usermetrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserViewMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetUserViewMetrics is deprecated
//
// Summary:
//
// get user view  metrics
//
// @param request - GetUserViewMetricsRequest
//
// @return GetUserViewMetricsResponse
// Deprecated
func (client *Client) GetUserViewMetrics(ResourceGroupID *string, request *GetUserViewMetricsRequest) (_result *GetUserViewMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserViewMetricsResponse{}
	_body, _err := client.GetUserViewMetricsWithOptions(ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取算法的所有版本信息
//
// @param request - ListAlgorithmVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlgorithmVersionsResponse
func (client *Client) ListAlgorithmVersionsWithOptions(AlgorithmId *string, request *ListAlgorithmVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlgorithmVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlgorithmVersions"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms/" + dara.PercentEncode(dara.StringValue(AlgorithmId)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlgorithmVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取算法的所有版本信息
//
// @param request - ListAlgorithmVersionsRequest
//
// @return ListAlgorithmVersionsResponse
func (client *Client) ListAlgorithmVersions(AlgorithmId *string, request *ListAlgorithmVersionsRequest) (_result *ListAlgorithmVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlgorithmVersionsResponse{}
	_body, _err := client.ListAlgorithmVersionsWithOptions(AlgorithmId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取算法列表
//
// @param request - ListAlgorithmsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlgorithmsResponse
func (client *Client) ListAlgorithmsWithOptions(request *ListAlgorithmsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlgorithmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlgorithmId) {
		query["AlgorithmId"] = request.AlgorithmId
	}

	if !dara.IsNil(request.AlgorithmName) {
		query["AlgorithmName"] = request.AlgorithmName
	}

	if !dara.IsNil(request.AlgorithmProvider) {
		query["AlgorithmProvider"] = request.AlgorithmProvider
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlgorithms"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlgorithmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取算法列表
//
// @param request - ListAlgorithmsRequest
//
// @return ListAlgorithmsResponse
func (client *Client) ListAlgorithms(request *ListAlgorithmsRequest) (_result *ListAlgorithmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlgorithmsResponse{}
	_body, _err := client.ListAlgorithmsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源节点列表
//
// @param request - ListNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorType) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !dara.IsNil(request.AvailabilityZone) {
		query["AvailabilityZone"] = request.AvailabilityZone
	}

	if !dara.IsNil(request.CliqueID) {
		query["CliqueID"] = request.CliqueID
	}

	if !dara.IsNil(request.FilterByQuotaId) {
		query["FilterByQuotaId"] = request.FilterByQuotaId
	}

	if !dara.IsNil(request.FilterByResourceGroupIds) {
		query["FilterByResourceGroupIds"] = request.FilterByResourceGroupIds
	}

	if !dara.IsNil(request.GPUType) {
		query["GPUType"] = request.GPUType
	}

	if !dara.IsNil(request.HyperNode) {
		query["HyperNode"] = request.HyperNode
	}

	if !dara.IsNil(request.HyperZone) {
		query["HyperZone"] = request.HyperZone
	}

	if !dara.IsNil(request.MachineGroupIds) {
		query["MachineGroupIds"] = request.MachineGroupIds
	}

	if !dara.IsNil(request.NodeNames) {
		query["NodeNames"] = request.NodeNames
	}

	if !dara.IsNil(request.NodeStatuses) {
		query["NodeStatuses"] = request.NodeStatuses
	}

	if !dara.IsNil(request.NodeTypes) {
		query["NodeTypes"] = request.NodeTypes
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderInstanceIds) {
		query["OrderInstanceIds"] = request.OrderInstanceIds
	}

	if !dara.IsNil(request.OrderStatuses) {
		query["OrderStatuses"] = request.OrderStatuses
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.QuotaId) {
		query["QuotaId"] = request.QuotaId
	}

	if !dara.IsNil(request.ReasonCodes) {
		query["ReasonCodes"] = request.ReasonCodes
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/nodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源节点列表
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 您可以通过此API获取Quota上的任务信息列表
//
// @param request - ListQuotaWorkloadsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaWorkloadsResponse
func (client *Client) ListQuotaWorkloadsWithOptions(QuotaId *string, request *ListQuotaWorkloadsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotaWorkloadsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeforeWorkloadId) {
		query["BeforeWorkloadId"] = request.BeforeWorkloadId
	}

	if !dara.IsNil(request.GmtDequeuedTimeRange) {
		query["GmtDequeuedTimeRange"] = request.GmtDequeuedTimeRange
	}

	if !dara.IsNil(request.GmtEnqueuedTimeRange) {
		query["GmtEnqueuedTimeRange"] = request.GmtEnqueuedTimeRange
	}

	if !dara.IsNil(request.GmtPositionModifiedTimeRange) {
		query["GmtPositionModifiedTimeRange"] = request.GmtPositionModifiedTimeRange
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
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

	if !dara.IsNil(request.ShowOwn) {
		query["ShowOwn"] = request.ShowOwn
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SubQuotaIds) {
		query["SubQuotaIds"] = request.SubQuotaIds
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	if !dara.IsNil(request.WithHistoricalData) {
		query["WithHistoricalData"] = request.WithHistoricalData
	}

	if !dara.IsNil(request.WorkloadCreatedTimeRange) {
		query["WorkloadCreatedTimeRange"] = request.WorkloadCreatedTimeRange
	}

	if !dara.IsNil(request.WorkloadIds) {
		query["WorkloadIds"] = request.WorkloadIds
	}

	if !dara.IsNil(request.WorkloadStatuses) {
		query["WorkloadStatuses"] = request.WorkloadStatuses
	}

	if !dara.IsNil(request.WorkloadType) {
		query["WorkloadType"] = request.WorkloadType
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotaWorkloads"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(QuotaId)) + "/workloads"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaWorkloadsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 您可以通过此API获取Quota上的任务信息列表
//
// @param request - ListQuotaWorkloadsRequest
//
// @return ListQuotaWorkloadsResponse
func (client *Client) ListQuotaWorkloads(QuotaId *string, request *ListQuotaWorkloadsRequest) (_result *ListQuotaWorkloadsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotaWorkloadsResponse{}
	_body, _err := client.ListQuotaWorkloadsWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Quota列表
//
// @param request - ListQuotasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasResponse
func (client *Client) ListQuotasWithOptions(request *ListQuotasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.HasResource) {
		query["HasResource"] = request.HasResource
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.LayoutMode) {
		query["LayoutMode"] = request.LayoutMode
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

	if !dara.IsNil(request.ParentQuotaId) {
		query["ParentQuotaId"] = request.ParentQuotaId
	}

	if !dara.IsNil(request.QuotaIds) {
		query["QuotaIds"] = request.QuotaIds
	}

	if !dara.IsNil(request.QuotaName) {
		query["QuotaName"] = request.QuotaName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Statuses) {
		query["Statuses"] = request.Statuses
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.Versions) {
		query["Versions"] = request.Versions
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotas"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Quota列表
//
// @param request - ListQuotasRequest
//
// @return ListQuotasResponse
func (client *Client) ListQuotas(request *ListQuotasRequest) (_result *ListQuotasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasResponse{}
	_body, _err := client.ListQuotasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// list machine groups
//
// @param request - ListResourceGroupMachineGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupMachineGroupsResponse
func (client *Client) ListResourceGroupMachineGroupsWithOptions(ResourceGroupID *string, request *ListResourceGroupMachineGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceGroupMachineGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreatorID) {
		query["CreatorID"] = request.CreatorID
	}

	if !dara.IsNil(request.EcsSpec) {
		query["EcsSpec"] = request.EcsSpec
	}

	if !dara.IsNil(request.MachineGroupIDs) {
		query["MachineGroupIDs"] = request.MachineGroupIDs
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderInstanceId) {
		query["OrderInstanceId"] = request.OrderInstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PaymentDuration) {
		query["PaymentDuration"] = request.PaymentDuration
	}

	if !dara.IsNil(request.PaymentDurationUnit) {
		query["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroupMachineGroups"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/" + dara.PercentEncode(dara.StringValue(ResourceGroupID)) + "/machinegroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupMachineGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// list machine groups
//
// @param request - ListResourceGroupMachineGroupsRequest
//
// @return ListResourceGroupMachineGroupsResponse
func (client *Client) ListResourceGroupMachineGroups(ResourceGroupID *string, request *ListResourceGroupMachineGroupsRequest) (_result *ListResourceGroupMachineGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceGroupMachineGroupsResponse{}
	_body, _err := client.ListResourceGroupMachineGroupsWithOptions(ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// list resource group
//
// @param request - ListResourceGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithOptions(request *ListResourceGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingResourceProvider) {
		query["ComputingResourceProvider"] = request.ComputingResourceProvider
	}

	if !dara.IsNil(request.HasResource) {
		query["HasResource"] = request.HasResource
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.ResourceGroupIDs) {
		query["ResourceGroupIDs"] = request.ResourceGroupIDs
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ShowAll) {
		query["ShowAll"] = request.ShowAll
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Versions) {
		query["Versions"] = request.Versions
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroups"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// list resource group
//
// @param request - ListResourceGroupsRequest
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查标签接口
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("ResourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["ResourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查标签接口
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定TrainingJob的事件。
//
// @param request - ListTrainingJobEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobEventsResponse
func (client *Client) ListTrainingJobEventsWithOptions(TrainingJobId *string, request *ListTrainingJobEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrainingJobEvents"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrainingJobEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定TrainingJob的事件。
//
// @param request - ListTrainingJobEventsRequest
//
// @return ListTrainingJobEventsResponse
func (client *Client) ListTrainingJobEvents(TrainingJobId *string, request *ListTrainingJobEventsRequest) (_result *ListTrainingJobEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobEventsResponse{}
	_body, _err := client.ListTrainingJobEventsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定Instance（TrainingJob的运行单元）的日志。
//
// @param request - ListTrainingJobInstanceEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobInstanceEventsResponse
func (client *Client) ListTrainingJobInstanceEventsWithOptions(TrainingJobId *string, InstanceId *string, request *ListTrainingJobInstanceEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobInstanceEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrainingJobInstanceEvents"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrainingJobInstanceEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定Instance（TrainingJob的运行单元）的日志。
//
// @param request - ListTrainingJobInstanceEventsRequest
//
// @return ListTrainingJobInstanceEventsResponse
func (client *Client) ListTrainingJobInstanceEvents(TrainingJobId *string, InstanceId *string, request *ListTrainingJobInstanceEventsRequest) (_result *ListTrainingJobInstanceEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobInstanceEventsResponse{}
	_body, _err := client.ListTrainingJobInstanceEventsWithOptions(TrainingJobId, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job实例的Metrics
//
// @param request - ListTrainingJobInstanceMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobInstanceMetricsResponse
func (client *Client) ListTrainingJobInstanceMetricsWithOptions(TrainingJobId *string, request *ListTrainingJobInstanceMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobInstanceMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeStep) {
		query["TimeStep"] = request.TimeStep
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrainingJobInstanceMetrics"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/instancemetrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrainingJobInstanceMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Training Job实例的Metrics
//
// @param request - ListTrainingJobInstanceMetricsRequest
//
// @return ListTrainingJobInstanceMetricsResponse
func (client *Client) ListTrainingJobInstanceMetrics(TrainingJobId *string, request *ListTrainingJobInstanceMetricsRequest) (_result *ListTrainingJobInstanceMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobInstanceMetricsResponse{}
	_body, _err := client.ListTrainingJobInstanceMetricsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job的日志
//
// @param request - ListTrainingJobLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobLogsResponse
func (client *Client) ListTrainingJobLogsWithOptions(TrainingJobId *string, request *ListTrainingJobLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.WorkerId) {
		query["WorkerId"] = request.WorkerId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrainingJobLogs"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrainingJobLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Training Job的日志
//
// @param request - ListTrainingJobLogsRequest
//
// @return ListTrainingJobLogsResponse
func (client *Client) ListTrainingJobLogs(TrainingJobId *string, request *ListTrainingJobLogsRequest) (_result *ListTrainingJobLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobLogsResponse{}
	_body, _err := client.ListTrainingJobLogsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job的Metrics
//
// @param request - ListTrainingJobMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobMetricsResponse
func (client *Client) ListTrainingJobMetricsWithOptions(TrainingJobId *string, request *ListTrainingJobMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrainingJobMetrics"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrainingJobMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Training Job的Metrics
//
// @param request - ListTrainingJobMetricsRequest
//
// @return ListTrainingJobMetricsResponse
func (client *Client) ListTrainingJobMetrics(TrainingJobId *string, request *ListTrainingJobMetricsRequest) (_result *ListTrainingJobMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobMetricsResponse{}
	_body, _err := client.ListTrainingJobMetricsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Training Job 产出的所有模型信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobOutputModelsResponse
func (client *Client) ListTrainingJobOutputModelsWithOptions(TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobOutputModelsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrainingJobOutputModels"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/outputmodels"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrainingJobOutputModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Training Job 产出的所有模型信息
//
// @return ListTrainingJobOutputModelsResponse
func (client *Client) ListTrainingJobOutputModels(TrainingJobId *string) (_result *ListTrainingJobOutputModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobOutputModelsResponse{}
	_body, _err := client.ListTrainingJobOutputModelsWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取TrainingJob的列表
//
// @param tmpReq - ListTrainingJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobsResponse
func (client *Client) ListTrainingJobsWithOptions(tmpReq *ListTrainingJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTrainingJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Labels) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, dara.String("Labels"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlgorithmName) {
		query["AlgorithmName"] = request.AlgorithmName
	}

	if !dara.IsNil(request.AlgorithmProvider) {
		query["AlgorithmProvider"] = request.AlgorithmProvider
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IsTempAlgo) {
		query["IsTempAlgo"] = request.IsTempAlgo
	}

	if !dara.IsNil(request.LabelsShrink) {
		query["Labels"] = request.LabelsShrink
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

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TrainingJobId) {
		query["TrainingJobId"] = request.TrainingJobId
	}

	if !dara.IsNil(request.TrainingJobName) {
		query["TrainingJobName"] = request.TrainingJobName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrainingJobs"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrainingJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取TrainingJob的列表
//
// @param request - ListTrainingJobsRequest
//
// @return ListTrainingJobsResponse
func (client *Client) ListTrainingJobs(request *ListTrainingJobsRequest) (_result *ListTrainingJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTrainingJobsResponse{}
	_body, _err := client.ListTrainingJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 扩缩容Quota
//
// @param request - ScaleQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleQuotaResponse
func (client *Client) ScaleQuotaWithOptions(QuotaId *string, request *ScaleQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Min) {
		body["Min"] = request.Min
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		body["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleQuota"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(QuotaId)) + "/action/scale"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 扩缩容Quota
//
// @param request - ScaleQuotaRequest
//
// @return ScaleQuotaResponse
func (client *Client) ScaleQuota(QuotaId *string, request *ScaleQuotaRequest) (_result *ScaleQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleQuotaResponse{}
	_body, _err := client.ScaleQuotaWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止一个TrainingJob
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrainingJobResponse
func (client *Client) StopTrainingJobWithOptions(TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTrainingJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTrainingJob"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止一个TrainingJob
//
// @return StopTrainingJobResponse
func (client *Client) StopTrainingJob(TrainingJobId *string) (_result *StopTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTrainingJobResponse{}
	_body, _err := client.StopTrainingJobWithOptions(TrainingJobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删标签接口
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("ResourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKey) {
		request.TagKeyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKey, dara.String("TagKey"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["ResourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeyShrink) {
		query["TagKey"] = request.TagKeyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tags"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删标签接口
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新算法
//
// @param request - UpdateAlgorithmRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlgorithmResponse
func (client *Client) UpdateAlgorithmWithOptions(AlgorithmId *string, request *UpdateAlgorithmRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlgorithmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlgorithmDescription) {
		body["AlgorithmDescription"] = request.AlgorithmDescription
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlgorithm"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms/" + dara.PercentEncode(dara.StringValue(AlgorithmId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新算法
//
// @param request - UpdateAlgorithmRequest
//
// @return UpdateAlgorithmResponse
func (client *Client) UpdateAlgorithm(AlgorithmId *string, request *UpdateAlgorithmRequest) (_result *UpdateAlgorithmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlgorithmResponse{}
	_body, _err := client.UpdateAlgorithmWithOptions(AlgorithmId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新算法
//
// @param tmpReq - UpdateAlgorithmVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlgorithmVersionResponse
func (client *Client) UpdateAlgorithmVersionWithOptions(AlgorithmId *string, AlgorithmVersion *string, tmpReq *UpdateAlgorithmVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlgorithmVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAlgorithmVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlgorithmSpec) {
		request.AlgorithmSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlgorithmSpec, dara.String("AlgorithmSpec"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlgorithmSpecShrink) {
		body["AlgorithmSpec"] = request.AlgorithmSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlgorithmVersion"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/algorithms/" + dara.PercentEncode(dara.StringValue(AlgorithmId)) + "/versions/" + dara.PercentEncode(dara.StringValue(AlgorithmVersion))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlgorithmVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新算法
//
// @param request - UpdateAlgorithmVersionRequest
//
// @return UpdateAlgorithmVersionResponse
func (client *Client) UpdateAlgorithmVersion(AlgorithmId *string, AlgorithmVersion *string, request *UpdateAlgorithmVersionRequest) (_result *UpdateAlgorithmVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAlgorithmVersionResponse{}
	_body, _err := client.UpdateAlgorithmVersionWithOptions(AlgorithmId, AlgorithmVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Quota
//
// @param request - UpdateQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuotaWithOptions(QuotaId *string, request *UpdateQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.QueueStrategy) {
		body["QueueStrategy"] = request.QueueStrategy
	}

	if !dara.IsNil(request.QuotaConfig) {
		body["QuotaConfig"] = request.QuotaConfig
	}

	if !dara.IsNil(request.QuotaName) {
		body["QuotaName"] = request.QuotaName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQuota"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(QuotaId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Quota
//
// @param request - UpdateQuotaRequest
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuota(QuotaId *string, request *UpdateQuotaRequest) (_result *UpdateQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaResponse{}
	_body, _err := client.UpdateQuotaWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新Resource Group
//
// @param request - UpdateResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroupWithOptions(ResourceGroupID *string, request *UpdateResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Unbind) {
		body["Unbind"] = request.Unbind
	}

	if !dara.IsNil(request.UserVpc) {
		body["UserVpc"] = request.UserVpc
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceGroup"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/resources/" + dara.PercentEncode(dara.StringValue(ResourceGroupID))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Resource Group
//
// @param request - UpdateResourceGroupRequest
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroup(ResourceGroupID *string, request *UpdateResourceGroupRequest) (_result *UpdateResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.UpdateResourceGroupWithOptions(ResourceGroupID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新一个TrainingJob的Labels
//
// @param request - UpdateTrainingJobLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrainingJobLabelsResponse
func (client *Client) UpdateTrainingJobLabelsWithOptions(TrainingJobId *string, request *UpdateTrainingJobLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTrainingJobLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrainingJobLabels"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/trainingjobs/" + dara.PercentEncode(dara.StringValue(TrainingJobId)) + "/labels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrainingJobLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个TrainingJob的Labels
//
// @param request - UpdateTrainingJobLabelsRequest
//
// @return UpdateTrainingJobLabelsResponse
func (client *Client) UpdateTrainingJobLabels(TrainingJobId *string, request *UpdateTrainingJobLabelsRequest) (_result *UpdateTrainingJobLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTrainingJobLabelsResponse{}
	_body, _err := client.UpdateTrainingJobLabelsWithOptions(TrainingJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
