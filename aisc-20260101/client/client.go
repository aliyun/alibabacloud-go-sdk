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
		"cn-zhangjiakou":        dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":            dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-nanjing":            dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-huhehaote":          dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou":           dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-guangzhou":          dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-beijing":            dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"ap-southeast-7":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-6":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-1":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":          dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":  dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-heyuan-acdr-1":      dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("aisc.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("aisc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates an AI Red Teaming scan target (agent or model) and saves its connection configuration for subsequent connectivity tests and scan tasks.
//
// Description:
//
// Creates an attack target (agent or model) and saves its connection configuration for subsequent connectivity tests and scan tasks.
//
// This is a synchronous call. Upon success, the system-generated TargetId is returned in Data. You can use this TargetId as a parameter in subsequent calls such as TestConnectivity and scan task creation.
//
// Metric description:
//
// - When ConnectionMethod is set to enterprise_relay (access through a corporate internal network agent), the values of Endpoint and ModelName are ignored. The platform uses fixed internal network values. The actual target endpoint and credentials are held by the corporate internal network agent.
//
// - After ApiKey is submitted, it is encrypted and stored. Subsequent queries do not return the plaintext value.
//
// - ConnectionConfig is a JSON character string in JSON format that specifies advanced connection settings. For common provider templates, refer to the metric description of this parameter.
//
// After the target is created, its initial connectivity status is verified. You can call TestConnectivity at any time to re-verify.
//
// Internal network access (enterprise_relay) workflow:
//
// - After the target is created, invoke GenerateRelayPollerScript to obtain an installation script (Linux only) and run it on a machine within the corporate internal network. The actual target endpoint and credentials are entered interactively during installation. The platform does not retain them.
//
// - After installation, the poller automatically registers and enters a polling loop. No manual registration or polling invocations are required.
//
// - After invoking TestConnectivity to authenticate end-to-end connectivity, you can use CreateTargetScanTask to initiate a scan.
//
// @param request - CreateAttackTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAttackTargetResponse
func (client *Client) CreateAttackTargetWithOptions(request *CreateAttackTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateAttackTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ConnectionConfig) {
		query["ConnectionConfig"] = request.ConnectionConfig
	}

	if !dara.IsNil(request.ConnectionMethod) {
		query["ConnectionMethod"] = request.ConnectionMethod
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Provider) {
		query["Provider"] = request.Provider
	}

	if !dara.IsNil(request.TargetName) {
		query["TargetName"] = request.TargetName
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAttackTarget"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAttackTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AI Red Teaming scan target (agent or model) and saves its connection configuration for subsequent connectivity tests and scan tasks.
//
// Description:
//
// Creates an attack target (agent or model) and saves its connection configuration for subsequent connectivity tests and scan tasks.
//
// This is a synchronous call. Upon success, the system-generated TargetId is returned in Data. You can use this TargetId as a parameter in subsequent calls such as TestConnectivity and scan task creation.
//
// Metric description:
//
// - When ConnectionMethod is set to enterprise_relay (access through a corporate internal network agent), the values of Endpoint and ModelName are ignored. The platform uses fixed internal network values. The actual target endpoint and credentials are held by the corporate internal network agent.
//
// - After ApiKey is submitted, it is encrypted and stored. Subsequent queries do not return the plaintext value.
//
// - ConnectionConfig is a JSON character string in JSON format that specifies advanced connection settings. For common provider templates, refer to the metric description of this parameter.
//
// After the target is created, its initial connectivity status is verified. You can call TestConnectivity at any time to re-verify.
//
// Internal network access (enterprise_relay) workflow:
//
// - After the target is created, invoke GenerateRelayPollerScript to obtain an installation script (Linux only) and run it on a machine within the corporate internal network. The actual target endpoint and credentials are entered interactively during installation. The platform does not retain them.
//
// - After installation, the poller automatically registers and enters a polling loop. No manual registration or polling invocations are required.
//
// - After invoking TestConnectivity to authenticate end-to-end connectivity, you can use CreateTargetScanTask to initiate a scan.
//
// @param request - CreateAttackTargetRequest
//
// @return CreateAttackTargetResponse
func (client *Client) CreateAttackTarget(request *CreateAttackTargetRequest) (_result *CreateAttackTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAttackTargetResponse{}
	_body, _err := client.CreateAttackTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates batch detection for user-defined skills.
//
// @param request - CreateSkillFileCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillFileCheckResponse
func (client *Client) CreateSkillFileCheckWithOptions(request *CreateSkillFileCheckRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillFileCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Files) {
		query["Files"] = request.Files
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillFileCheck"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillFileCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates batch detection for user-defined skills.
//
// @param request - CreateSkillFileCheckRequest
//
// @return CreateSkillFileCheckResponse
func (client *Client) CreateSkillFileCheck(request *CreateSkillFileCheckRequest) (_result *CreateSkillFileCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillFileCheckResponse{}
	_body, _err := client.CreateSkillFileCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a security scan task against a scan target that has passed connectivity verification. After the task is created, it is asynchronously prepared and executed.
//
// Description:
//
// *Before you use this operation, make sure that you fully understand the billing method and [pricing](https://www.alibabacloud.com/help/en/asc/user-guide/ai-red-teaming#aefbf9b5b4noh) of AI Red Teaming.**
//
// Initiates a security scan (AI Red Teaming detection) task against a specified scan target.
//
// Before you begin:
//
// - The account must have a normal subscription status. Otherwise, a 403 error is returned.
//
// - The TargetId must exist and belong to the current tenant.
//
// - The connectivity verification status of the target must be verified. You can call TestConnectivity to verify the target first. Otherwise, a 400 error is returned.
//
// - Available attack samples must exist within the current scan scope. Otherwise, a 400 error is returned.
//
// Execution mode:
//
// - The call synchronously returns a TaskId. The initial task status is PREPARING. Sample preparation and scan execution are performed asynchronously. You can call ListScanTasksByTarget to query the task status and progress.
//
// Sample selection:
//
// - The sample scope is determined based on the target type (agent/model) plus general-purpose samples. SampleLevel determines the detection intensity and derives the technique level. Lang is used to filter samples by language.
//
// - If no sample intent is specified, the system automatically derives all available intents based on the scope described above.
//
// @param request - CreateTargetScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTargetScanTaskResponse
func (client *Client) CreateTargetScanTaskWithOptions(request *CreateTargetScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateTargetScanTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SampleLevel) {
		query["SampleLevel"] = request.SampleLevel
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTargetScanTask"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTargetScanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a security scan task against a scan target that has passed connectivity verification. After the task is created, it is asynchronously prepared and executed.
//
// Description:
//
// *Before you use this operation, make sure that you fully understand the billing method and [pricing](https://www.alibabacloud.com/help/en/asc/user-guide/ai-red-teaming#aefbf9b5b4noh) of AI Red Teaming.**
//
// Initiates a security scan (AI Red Teaming detection) task against a specified scan target.
//
// Before you begin:
//
// - The account must have a normal subscription status. Otherwise, a 403 error is returned.
//
// - The TargetId must exist and belong to the current tenant.
//
// - The connectivity verification status of the target must be verified. You can call TestConnectivity to verify the target first. Otherwise, a 400 error is returned.
//
// - Available attack samples must exist within the current scan scope. Otherwise, a 400 error is returned.
//
// Execution mode:
//
// - The call synchronously returns a TaskId. The initial task status is PREPARING. Sample preparation and scan execution are performed asynchronously. You can call ListScanTasksByTarget to query the task status and progress.
//
// Sample selection:
//
// - The sample scope is determined based on the target type (agent/model) plus general-purpose samples. SampleLevel determines the detection intensity and derives the technique level. Lang is used to filter samples by language.
//
// - If no sample intent is specified, the system automatically derives all available intents based on the scope described above.
//
// @param request - CreateTargetScanTaskRequest
//
// @return CreateTargetScanTaskResponse
func (client *Client) CreateTargetScanTask(request *CreateTargetScanTaskRequest) (_result *CreateTargetScanTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTargetScanTaskResponse{}
	_body, _err := client.CreateTargetScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified scan target by TargetId in AI Red Teaming. The target is physically deleted and cannot be recovered. This operation does not stop scan tasks that are in progress.
//
// Description:
//
// Deletes a specified scan target by TargetId. The target is physically deleted.
//
// - You can delete only targets that belong to the current tenant. If the target does not exist or belongs to another tenant, a 400 error is returned. This prevents exposing whether the resource exists.
//
// - Physical deletion: The target cannot be recovered after deletion. Confirm before you proceed.
//
// - This operation deletes only the target record. It does not stop scan tasks that are in progress for the target or delete historical scan task records. To stop or clean up tasks, call StopScannerTask or DeleteScannerTask first.
//
// - After deletion, the connection configurations of the target, including encrypted credentials and connectivity verification results, are also removed.
//
// @param request - DeleteAttackTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAttackTargetResponse
func (client *Client) DeleteAttackTargetWithOptions(request *DeleteAttackTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteAttackTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAttackTarget"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAttackTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified scan target by TargetId in AI Red Teaming. The target is physically deleted and cannot be recovered. This operation does not stop scan tasks that are in progress.
//
// Description:
//
// Deletes a specified scan target by TargetId. The target is physically deleted.
//
// - You can delete only targets that belong to the current tenant. If the target does not exist or belongs to another tenant, a 400 error is returned. This prevents exposing whether the resource exists.
//
// - Physical deletion: The target cannot be recovered after deletion. Confirm before you proceed.
//
// - This operation deletes only the target record. It does not stop scan tasks that are in progress for the target or delete historical scan task records. To stop or clean up tasks, call StopScannerTask or DeleteScannerTask first.
//
// - After deletion, the connection configurations of the target, including encrypted credentials and connectivity verification results, are also removed.
//
// @param request - DeleteAttackTargetRequest
//
// @return DeleteAttackTargetResponse
func (client *Client) DeleteAttackTarget(request *DeleteAttackTargetRequest) (_result *DeleteAttackTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAttackTargetResponse{}
	_body, _err := client.DeleteAttackTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Physically deletes a specified scan task by ScannerTaskId in AI Red Teaming.
//
// Description:
//
// Physically deletes a specified scan task by ScannerTaskId.
//
// - Only tasks that belong to the current tenant can be deleted. If the task does not exist or belongs to another tenant, a 400 error is returned without exposing whether the resource exists.
//
// - If the task is in progress (sample preparation, waiting, processing, or report generation), the task is automatically canceled before deletion. A cancellation failure does not block the deletion.
//
// - Physical deletion: After deletion, the task record and its status and progress information cannot be queried or recovered. Confirm before you delete.
//
// - Deleting a task record does not affect the scan target itself.
//
// @param request - DeleteScannerTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScannerTaskResponse
func (client *Client) DeleteScannerTaskWithOptions(request *DeleteScannerTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteScannerTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScannerTaskId) {
		query["ScannerTaskId"] = request.ScannerTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScannerTask"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScannerTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Physically deletes a specified scan task by ScannerTaskId in AI Red Teaming.
//
// Description:
//
// Physically deletes a specified scan task by ScannerTaskId.
//
// - Only tasks that belong to the current tenant can be deleted. If the task does not exist or belongs to another tenant, a 400 error is returned without exposing whether the resource exists.
//
// - If the task is in progress (sample preparation, waiting, processing, or report generation), the task is automatically canceled before deletion. A cancellation failure does not block the deletion.
//
// - Physical deletion: After deletion, the task record and its status and progress information cannot be queried or recovered. Confirm before you delete.
//
// - Deleting a task record does not affect the scan target itself.
//
// @param request - DeleteScannerTaskRequest
//
// @return DeleteScannerTaskResponse
func (client *Client) DeleteScannerTask(request *DeleteScannerTaskRequest) (_result *DeleteScannerTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScannerTaskResponse{}
	_body, _err := client.DeleteScannerTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a one-time installation script for a scan target of the enterprise_relay type in AI Red Teaming. Only Linux platforms are supported. The script contains a one-time access code.
//
// Description:
//
// Generates an installation script for the internal network agent (relay-poller) for a specified scan target.
//
// - The target must use the enterprise_relay connection method (see CreateAttackTarget). Otherwise, HTTP status code 400 is returned. If the target does not exist or belongs to another tenant, HTTP status code 400 is returned without exposing whether the resource exists.
//
// - Only Linux is supported for the platform. The Platform parameter uses the "operating system-architecture" format and accepts only linux-amd64 and linux-arm64. Compatible architecture values include amd64, x86_64, x86, arm64, and aarch64. If only the architecture is specified, the operating system defaults to linux. Other operating systems such as macOS and Windows return HTTP status code 400. If this parameter is not specified, the default value is linux-amd64.
//
// - The script contains a one-time access code. Each call issues a new access code, and the previous code automatically expires. Re-downloading the script generates a new access code. Use the latest generated script for installation.
//
// - The script contains a temporary download link (a signed link valid for 1 hour) and a checksum for the poller binary. The binary is available only for Linux in both architectures.
//
// - The script does not contain the actual endpoint or credentials of the target. The installer interactively enters these values when running the script. The platform does not store them.
//
// - After installation, the poller automatically completes registration and enters a polling cycle. The registration and polling operations are automatically called by the script and do not require manual invocation. You can call TestConnectivity to verify end-to-end connectivity.
//
// @param request - GenerateRelayPollerScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateRelayPollerScriptResponse
func (client *Client) GenerateRelayPollerScriptWithOptions(request *GenerateRelayPollerScriptRequest, runtime *dara.RuntimeOptions) (_result *GenerateRelayPollerScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateRelayPollerScript"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateRelayPollerScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a one-time installation script for a scan target of the enterprise_relay type in AI Red Teaming. Only Linux platforms are supported. The script contains a one-time access code.
//
// Description:
//
// Generates an installation script for the internal network agent (relay-poller) for a specified scan target.
//
// - The target must use the enterprise_relay connection method (see CreateAttackTarget). Otherwise, HTTP status code 400 is returned. If the target does not exist or belongs to another tenant, HTTP status code 400 is returned without exposing whether the resource exists.
//
// - Only Linux is supported for the platform. The Platform parameter uses the "operating system-architecture" format and accepts only linux-amd64 and linux-arm64. Compatible architecture values include amd64, x86_64, x86, arm64, and aarch64. If only the architecture is specified, the operating system defaults to linux. Other operating systems such as macOS and Windows return HTTP status code 400. If this parameter is not specified, the default value is linux-amd64.
//
// - The script contains a one-time access code. Each call issues a new access code, and the previous code automatically expires. Re-downloading the script generates a new access code. Use the latest generated script for installation.
//
// - The script contains a temporary download link (a signed link valid for 1 hour) and a checksum for the poller binary. The binary is available only for Linux in both architectures.
//
// - The script does not contain the actual endpoint or credentials of the target. The installer interactively enters these values when running the script. The platform does not store them.
//
// - After installation, the poller automatically completes registration and enters a polling cycle. The registration and polling operations are automatically called by the script and do not require manual invocation. You can call TestConnectivity to verify end-to-end connectivity.
//
// @param request - GenerateRelayPollerScriptRequest
//
// @return GenerateRelayPollerScriptResponse
func (client *Client) GenerateRelayPollerScript(request *GenerateRelayPollerScriptRequest) (_result *GenerateRelayPollerScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateRelayPollerScriptResponse{}
	_body, _err := client.GenerateRelayPollerScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a scan target by TargetId for AI Red Teaming, including connection configuration and scan task configuration.
//
// Description:
//
// Queries the details of a scan target by TargetId.
//
// - Only targets that belong to the current tenant can be queried. If the target does not exist or belongs to another tenant, a 400 error is returned to avoid exposing whether the resource exists.
//
// - The response includes basic target information, advanced connection configuration (ConnectionConfig), and scan task configuration (ScanTaskConfig).
//
// - The following six aggregate fields are not populated by this operation and return empty values: cumulative scan count (ScanCount), last scan status (LastScanStatus), risk level (RiskLevel), first scan time (FirstScanTime), last scan time (LastScanTime), and last scan failure reason (LastScanFailMessage). Query these fields by calling ListAttackTargets or ListScanTasksByTarget.
//
// - The response does not include sensitive credentials such as ApiKey in plaintext.
//
// @param request - GetAttackTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttackTargetResponse
func (client *Client) GetAttackTargetWithOptions(request *GetAttackTargetRequest, runtime *dara.RuntimeOptions) (_result *GetAttackTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAttackTarget"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAttackTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a scan target by TargetId for AI Red Teaming, including connection configuration and scan task configuration.
//
// Description:
//
// Queries the details of a scan target by TargetId.
//
// - Only targets that belong to the current tenant can be queried. If the target does not exist or belongs to another tenant, a 400 error is returned to avoid exposing whether the resource exists.
//
// - The response includes basic target information, advanced connection configuration (ConnectionConfig), and scan task configuration (ScanTaskConfig).
//
// - The following six aggregate fields are not populated by this operation and return empty values: cumulative scan count (ScanCount), last scan status (LastScanStatus), risk level (RiskLevel), first scan time (FirstScanTime), last scan time (LastScanTime), and last scan failure reason (LastScanFailMessage). Query these fields by calling ListAttackTargets or ListScanTasksByTarget.
//
// - The response does not include sensitive credentials such as ApiKey in plaintext.
//
// @param request - GetAttackTargetRequest
//
// @return GetAttackTargetResponse
func (client *Client) GetAttackTarget(request *GetAttackTargetRequest) (_result *GetAttackTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAttackTargetResponse{}
	_body, _err := client.GetAttackTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a temporary download URL for the attack hit data (hits.csv) of a specified scan task in AI Red Teaming. The URL is valid for 2 hours. An empty string is returned if the data has not been generated.
//
// Description:
//
// Retrieves a temporary download URL for the attack hit data (hits.csv) of a specified scan task.
//
// - ScannerTaskId is required in practice. An empty value returns HTTP status code 400.
//
// - You can query only tasks that belong to the current tenant. If the task does not exist or belongs to another tenant, HTTP status code 400 is returned uniformly to avoid exposing whether the resource exists.
//
// - The download URL is a signed temporary URL of Object Storage Service (OSS) that is valid for 2 hours (7,200 seconds). After the URL expires, call this operation again to obtain a new URL.
//
// - If the attack hit data has not been generated (for existing tasks or when the agent execution mode does not produce hit data), the download URL in the response is an empty string. No error is returned.
//
// - This operation is a read-only action (with the Get prefix). A RAM user with read-only permissions can call this operation. The behavior is consistent with the deprecated GenerateScannerTaskHitDataUrl operation.
//
// @param request - GetScannerTaskHitDataUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScannerTaskHitDataUrlResponse
func (client *Client) GetScannerTaskHitDataUrlWithOptions(request *GetScannerTaskHitDataUrlRequest, runtime *dara.RuntimeOptions) (_result *GetScannerTaskHitDataUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScannerTaskId) {
		query["ScannerTaskId"] = request.ScannerTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScannerTaskHitDataUrl"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScannerTaskHitDataUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a temporary download URL for the attack hit data (hits.csv) of a specified scan task in AI Red Teaming. The URL is valid for 2 hours. An empty string is returned if the data has not been generated.
//
// Description:
//
// Retrieves a temporary download URL for the attack hit data (hits.csv) of a specified scan task.
//
// - ScannerTaskId is required in practice. An empty value returns HTTP status code 400.
//
// - You can query only tasks that belong to the current tenant. If the task does not exist or belongs to another tenant, HTTP status code 400 is returned uniformly to avoid exposing whether the resource exists.
//
// - The download URL is a signed temporary URL of Object Storage Service (OSS) that is valid for 2 hours (7,200 seconds). After the URL expires, call this operation again to obtain a new URL.
//
// - If the attack hit data has not been generated (for existing tasks or when the agent execution mode does not produce hit data), the download URL in the response is an empty string. No error is returned.
//
// - This operation is a read-only action (with the Get prefix). A RAM user with read-only permissions can call this operation. The behavior is consistent with the deprecated GenerateScannerTaskHitDataUrl operation.
//
// @param request - GetScannerTaskHitDataUrlRequest
//
// @return GetScannerTaskHitDataUrlResponse
func (client *Client) GetScannerTaskHitDataUrl(request *GetScannerTaskHitDataUrlRequest) (_result *GetScannerTaskHitDataUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScannerTaskHitDataUrlResponse{}
	_body, _err := client.GetScannerTaskHitDataUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a temporary download URL for the HTML result report of an AI Red Teaming scan task. The URL is valid for 2 hours. An empty string is returned if the report has not been generated.
//
// Description:
//
// Retrieves a temporary download URL for the HTML result report of a specified scan task.
//
// - You can only query tasks that belong to the current tenant. If the task does not exist or belongs to another tenant, a 400 error is returned to avoid exposing whether the resource exists.
//
// - The download URL is a signed temporary URL from object storage, valid for 2 hours (7,200 seconds). After the URL expires, call this operation again to obtain a new URL.
//
// - If the task result report has not been generated (the task is not complete or the report has not been produced), the download URL in the response is an empty string and no error is returned. Call this operation after the task status changes to completed.
//
// @param request - GetScannerTaskResultHtmlUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScannerTaskResultHtmlUrlResponse
func (client *Client) GetScannerTaskResultHtmlUrlWithOptions(request *GetScannerTaskResultHtmlUrlRequest, runtime *dara.RuntimeOptions) (_result *GetScannerTaskResultHtmlUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScannerTaskId) {
		query["ScannerTaskId"] = request.ScannerTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScannerTaskResultHtmlUrl"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScannerTaskResultHtmlUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a temporary download URL for the HTML result report of an AI Red Teaming scan task. The URL is valid for 2 hours. An empty string is returned if the report has not been generated.
//
// Description:
//
// Retrieves a temporary download URL for the HTML result report of a specified scan task.
//
// - You can only query tasks that belong to the current tenant. If the task does not exist or belongs to another tenant, a 400 error is returned to avoid exposing whether the resource exists.
//
// - The download URL is a signed temporary URL from object storage, valid for 2 hours (7,200 seconds). After the URL expires, call this operation again to obtain a new URL.
//
// - If the task result report has not been generated (the task is not complete or the report has not been produced), the download URL in the response is an empty string and no error is returned. Call this operation after the task status changes to completed.
//
// @param request - GetScannerTaskResultHtmlUrlRequest
//
// @return GetScannerTaskResultHtmlUrlResponse
func (client *Client) GetScannerTaskResultHtmlUrl(request *GetScannerTaskResultHtmlUrlRequest) (_result *GetScannerTaskResultHtmlUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScannerTaskResultHtmlUrlResponse{}
	_body, _err := client.GetScannerTaskResultHtmlUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of agent risk events.
//
// @param request - ListAIAgentEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIAgentEventResponse
func (client *Client) ListAIAgentEventWithOptions(request *ListAIAgentEventRequest, runtime *dara.RuntimeOptions) (_result *ListAIAgentEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AssetName) {
		query["AssetName"] = request.AssetName
	}

	if !dara.IsNil(request.AssetType) {
		query["AssetType"] = request.AssetType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InfraInstanceId) {
		query["InfraInstanceId"] = request.InfraInstanceId
	}

	if !dara.IsNil(request.InfraName) {
		query["InfraName"] = request.InfraName
	}

	if !dara.IsNil(request.InfraRegionId) {
		query["InfraRegionId"] = request.InfraRegionId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceDirectoryAccountId) {
		query["ResourceDirectoryAccountId"] = request.ResourceDirectoryAccountId
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.RiskName) {
		query["RiskName"] = request.RiskName
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	if !dara.IsNil(request.Vendor) {
		query["Vendor"] = request.Vendor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAIAgentEvent"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIAgentEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of agent risk events.
//
// @param request - ListAIAgentEventRequest
//
// @return ListAIAgentEventResponse
func (client *Client) ListAIAgentEvent(request *ListAIAgentEventRequest) (_result *ListAIAgentEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAIAgentEventResponse{}
	_body, _err := client.ListAIAgentEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of scan targets for AI Red Teaming. This operation supports multi-dimensional filtering and sorting.
//
// @param request - ListAttackTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAttackTargetsResponse
func (client *Client) ListAttackTargetsWithOptions(request *ListAttackTargetsRequest, runtime *dara.RuntimeOptions) (_result *ListAttackTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirstScanTimeEnd) {
		query["FirstScanTimeEnd"] = request.FirstScanTimeEnd
	}

	if !dara.IsNil(request.FirstScanTimeStart) {
		query["FirstScanTimeStart"] = request.FirstScanTimeStart
	}

	if !dara.IsNil(request.LastScanStatus) {
		query["LastScanStatus"] = request.LastScanStatus
	}

	if !dara.IsNil(request.LastScanTimeEnd) {
		query["LastScanTimeEnd"] = request.LastScanTimeEnd
	}

	if !dara.IsNil(request.LastScanTimeStart) {
		query["LastScanTimeStart"] = request.LastScanTimeStart
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Provider) {
		query["Provider"] = request.Provider
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.SortField) {
		query["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.TargetName) {
		query["TargetName"] = request.TargetName
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAttackTargets"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAttackTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of scan targets for AI Red Teaming. This operation supports multi-dimensional filtering and sorting.
//
// @param request - ListAttackTargetsRequest
//
// @return ListAttackTargetsResponse
func (client *Client) ListAttackTargets(request *ListAttackTargetsRequest) (_result *ListAttackTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAttackTargetsResponse{}
	_body, _err := client.ListAttackTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paged query of scan tasks under a specified scan target for AI Red Teaming. Supports filtering by detection intensity, task status, and scan mode.
//
// Description:
//
// Performs a paged query of the scan task list under a specified scan target. Only tasks belonging to targets owned by the current tenant are returned.
//
// Query scope and sorting:
//
// - Only tasks created within the last 366 days are returned.
//
// - Results are sorted by creation time in descending order.
//
// - TaskStatus filters by task status. ScanType filters by scan mode. The scan mode is stored in the task execute parameters. Historical tasks without a recorded scan mode are treated as attack.
//
// Paged query rules:
//
// - PageNumber starts from 1. Values less than 1 are normalized to 1.
//
// - PageSize defaults to 10, with a maximum of 100 per page. Values greater than 100 are clamped to 100. Values less than 1 return HTTP status code 400.
//
// - The PageNumber and PageSize values in the response are the normalization values that actually take effect.
//
// @param request - ListScanTasksByTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScanTasksByTargetResponse
func (client *Client) ListScanTasksByTargetWithOptions(request *ListScanTasksByTargetRequest, runtime *dara.RuntimeOptions) (_result *ListScanTasksByTargetResponse, _err error) {
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

	if !dara.IsNil(request.SampleLevel) {
		query["SampleLevel"] = request.SampleLevel
	}

	if !dara.IsNil(request.ScanType) {
		query["ScanType"] = request.ScanType
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TaskStatus) {
		query["TaskStatus"] = request.TaskStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScanTasksByTarget"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScanTasksByTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of scan tasks under a specified scan target for AI Red Teaming. Supports filtering by detection intensity, task status, and scan mode.
//
// Description:
//
// Performs a paged query of the scan task list under a specified scan target. Only tasks belonging to targets owned by the current tenant are returned.
//
// Query scope and sorting:
//
// - Only tasks created within the last 366 days are returned.
//
// - Results are sorted by creation time in descending order.
//
// - TaskStatus filters by task status. ScanType filters by scan mode. The scan mode is stored in the task execute parameters. Historical tasks without a recorded scan mode are treated as attack.
//
// Paged query rules:
//
// - PageNumber starts from 1. Values less than 1 are normalized to 1.
//
// - PageSize defaults to 10, with a maximum of 100 per page. Values greater than 100 are clamped to 100. Values less than 1 return HTTP status code 400.
//
// - The PageNumber and PageSize values in the response are the normalization values that actually take effect.
//
// @param request - ListScanTasksByTargetRequest
//
// @return ListScanTasksByTargetResponse
func (client *Client) ListScanTasksByTarget(request *ListScanTasksByTargetRequest) (_result *ListScanTasksByTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScanTasksByTargetResponse{}
	_body, _err := client.ListScanTasksByTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get subtask information.
//
// @param request - ListSubTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubTasksResponse
func (client *Client) ListSubTasksWithOptions(request *ListSubTasksRequest, runtime *dara.RuntimeOptions) (_result *ListSubTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RootTaskId) {
		query["RootTaskId"] = request.RootTaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubTasks"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get subtask information.
//
// @param request - ListSubTasksRequest
//
// @return ListSubTasksResponse
func (client *Client) ListSubTasks(request *ListSubTasksRequest) (_result *ListSubTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSubTasksResponse{}
	_body, _err := client.ListSubTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops (cancels) an in-progress scan task for AI Red Teaming. Calling this operation on a task that is already in a desired state is idempotent and performs no action.
//
// Description:
//
// Stops (cancels) the scan task specified by ScannerTaskId.
//
// - You can only operate on tasks that belong to the current tenant. If the task does not exist or belongs to another tenant, a 400 error is returned without exposing whether the resource exists.
//
// - Only tasks in an in-progress state (sample preparation, waiting, processing, or report generation) are actually canceled. The task status is set to canceled, the end time is recorded, and the underlying execution job is stopped asynchronously.
//
// - Idempotent: If the task is already in a desired state (completed, failed, timed out, or canceled), the call returns success without modifying the task.
//
// - The underlying execution job is stopped asynchronously. A failure to stop the job does not affect the cancellation result of the task itself.
//
// @param request - StopScannerTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopScannerTaskResponse
func (client *Client) StopScannerTaskWithOptions(request *StopScannerTaskRequest, runtime *dara.RuntimeOptions) (_result *StopScannerTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScannerTaskId) {
		query["ScannerTaskId"] = request.ScannerTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopScannerTask"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopScannerTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops (cancels) an in-progress scan task for AI Red Teaming. Calling this operation on a task that is already in a desired state is idempotent and performs no action.
//
// Description:
//
// Stops (cancels) the scan task specified by ScannerTaskId.
//
// - You can only operate on tasks that belong to the current tenant. If the task does not exist or belongs to another tenant, a 400 error is returned without exposing whether the resource exists.
//
// - Only tasks in an in-progress state (sample preparation, waiting, processing, or report generation) are actually canceled. The task status is set to canceled, the end time is recorded, and the underlying execution job is stopped asynchronously.
//
// - Idempotent: If the task is already in a desired state (completed, failed, timed out, or canceled), the call returns success without modifying the task.
//
// - The underlying execution job is stopped asynchronously. A failure to stop the job does not affect the cancellation result of the task itself.
//
// @param request - StopScannerTaskRequest
//
// @return StopScannerTaskResponse
func (client *Client) StopScannerTask(request *StopScannerTaskRequest) (_result *StopScannerTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopScannerTaskResponse{}
	_body, _err := client.StopScannerTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests the network connectivity and authentication validity of an AI Red Teaming target.
//
// Description:
//
// Tests the network connectivity and authentication validity of a specified attack target.
//
// This operation uses an asynchronous polling model:
//
// - First call (without CheckId): Immediately returns a CheckId with VerifyStatus=checking. The actual test runs asynchronously in the background for up to 60 seconds.
//
// - Subsequent calls (with the CheckId returned from the first call): Queries the latest status of the corresponding CheckId, which may be checking, verified, or failed.
//
// - Poll at 2-second intervals for up to 60 seconds. After the CheckId expires, the operation returns failed with VerifyMessage set to "check expired, please retry".
//
// Use one of the following two approaches for parameters:
//
// - Approach A: Specify only TargetId. The system reads Endpoint, ApiKey, ModelName, ConnectionMethod, and ConnectionConfig from the saved target configuration and ignores any parameters with the same names in the request.
//
// - Approach B: Do not specify TargetId. Instead, provide the five connection parameters directly in the request.
//
// @param request - TestConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestConnectivityResponse
func (client *Client) TestConnectivityWithOptions(request *TestConnectivityRequest, runtime *dara.RuntimeOptions) (_result *TestConnectivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.CheckId) {
		query["CheckId"] = request.CheckId
	}

	if !dara.IsNil(request.ConnectionConfig) {
		query["ConnectionConfig"] = request.ConnectionConfig
	}

	if !dara.IsNil(request.ConnectionMethod) {
		query["ConnectionMethod"] = request.ConnectionMethod
	}

	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestConnectivity"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestConnectivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the network connectivity and authentication validity of an AI Red Teaming target.
//
// Description:
//
// Tests the network connectivity and authentication validity of a specified attack target.
//
// This operation uses an asynchronous polling model:
//
// - First call (without CheckId): Immediately returns a CheckId with VerifyStatus=checking. The actual test runs asynchronously in the background for up to 60 seconds.
//
// - Subsequent calls (with the CheckId returned from the first call): Queries the latest status of the corresponding CheckId, which may be checking, verified, or failed.
//
// - Poll at 2-second intervals for up to 60 seconds. After the CheckId expires, the operation returns failed with VerifyMessage set to "check expired, please retry".
//
// Use one of the following two approaches for parameters:
//
// - Approach A: Specify only TargetId. The system reads Endpoint, ApiKey, ModelName, ConnectionMethod, and ConnectionConfig from the saved target configuration and ignores any parameters with the same names in the request.
//
// - Approach B: Do not specify TargetId. Instead, provide the five connection parameters directly in the request.
//
// @param request - TestConnectivityRequest
//
// @return TestConnectivityResponse
func (client *Client) TestConnectivity(request *TestConnectivityRequest) (_result *TestConnectivityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestConnectivityResponse{}
	_body, _err := client.TestConnectivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
