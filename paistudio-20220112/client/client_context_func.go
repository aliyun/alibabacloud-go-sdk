// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Checks the connection to the instance container.
//
// @param request - CheckInstanceWebTerminalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceWebTerminalResponse
func (client *Client) CheckInstanceWebTerminalWithContext(ctx context.Context, TrainingJobId *string, InstanceId *string, request *CheckInstanceWebTerminalRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInstanceWebTerminalResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an algorithm.
//
// @param request - CreateAlgorithmRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlgorithmResponse
func (client *Client) CreateAlgorithmWithContext(ctx context.Context, request *CreateAlgorithmRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlgorithmResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a version of an algorithm.
//
// @param tmpReq - CreateAlgorithmVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlgorithmVersionResponse
func (client *Client) CreateAlgorithmVersionWithContext(ctx context.Context, AlgorithmId *string, AlgorithmVersion *string, tmpReq *CreateAlgorithmVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAlgorithmVersionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can create a link to access the container of the instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceWebTerminalResponse
func (client *Client) CreateInstanceWebTerminalWithContext(ctx context.Context, TrainingJobId *string, InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceWebTerminalResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the `CreateQuota` operation to create a quota.
//
// @param request - CreateQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaResponse
func (client *Client) CreateQuotaWithContext(ctx context.Context, request *CreateQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use the CreateResourceGroup operation to create a resource group.
//
// @param request - CreateResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithContext(ctx context.Context, request *CreateResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a training job.
//
// @param request - CreateTrainingJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrainingJobResponse
func (client *Client) CreateTrainingJobWithContext(ctx context.Context, request *CreateTrainingJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrainingJobResponse, _err error) {
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

	if !dara.IsNil(request.CredentialConfig) {
		body["CredentialConfig"] = request.CredentialConfig
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an algorithm.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlgorithmResponse
func (client *Client) DeleteAlgorithmWithContext(ctx context.Context, AlgorithmId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlgorithmResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an algorithm version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlgorithmVersionResponse
func (client *Client) DeleteAlgorithmVersionWithContext(ctx context.Context, AlgorithmId *string, AlgorithmVersion *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAlgorithmVersionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Deletes a machine group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMachineGroupResponse
func (client *Client) DeleteMachineGroupWithContext(ctx context.Context, MachineGroupID *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMachineGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can delete a resource quota by using DeleteQuota.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaResponse
func (client *Client) DeleteQuotaWithContext(ctx context.Context, QuotaId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithContext(ctx context.Context, ResourceGroupID *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Deletes the specified machine group from a resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupMachineGroupResponse
func (client *Client) DeleteResourceGroupMachineGroupWithContext(ctx context.Context, MachineGroupID *string, ResourceGroupID *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupMachineGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a training task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrainingJobResponse
func (client *Client) DeleteTrainingJobWithContext(ctx context.Context, TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTrainingJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the labels of a training task.
//
// @param request - DeleteTrainingJobLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrainingJobLabelsResponse
func (client *Client) DeleteTrainingJobLabelsWithContext(ctx context.Context, TrainingJobId *string, request *DeleteTrainingJobLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTrainingJobLabelsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves algorithm details.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmResponse
func (client *Client) GetAlgorithmWithContext(ctx context.Context, AlgorithmId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlgorithmResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve details about an algorithm version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlgorithmVersionResponse
func (client *Client) GetAlgorithmVersionWithContext(ctx context.Context, AlgorithmId *string, AlgorithmVersion *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAlgorithmVersionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Retrieves the details of a specific machine group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMachineGroupResponse
func (client *Client) GetMachineGroupWithContext(ctx context.Context, MachineGroupID *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMachineGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// You can call the GetNodeMetrics operation to obtain the metric information of resource group nodes.
//
// @param request - GetNodeMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeMetricsResponse
func (client *Client) GetNodeMetricsWithContext(ctx context.Context, ResourceGroupID *string, MetricType *string, request *GetNodeMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetNodeMetricsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call GetQuota to retrieve the details of a resource quota. This helps you monitor current resource usage for effective resource management and planning.
//
// @param request - GetQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaResponse
func (client *Client) GetQuotaWithContext(ctx context.Context, QuotaId *string, request *GetQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// `GetResourceGroup` retrieves details for a resource group.
//
// @param tmpReq - GetResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroupWithContext(ctx context.Context, ResourceGroupID *string, tmpReq *GetResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call GetResourceGroupMachineGroup to retrieve the details of a machine group.
//
// @param tmpReq - GetResourceGroupMachineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupMachineGroupResponse
func (client *Client) GetResourceGroupMachineGroupWithContext(ctx context.Context, MachineGroupID *string, ResourceGroupID *string, tmpReq *GetResourceGroupMachineGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupMachineGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// You can call the GetResourceGroupRequest operation to obtain the number of resources that have been allocated to a resource group.
//
// @param request - GetResourceGroupRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupRequestResponse
func (client *Client) GetResourceGroupRequestWithContext(ctx context.Context, request *GetResourceGroupRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupRequestResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call GetResourceGroupTotal to get the total resources of a Resource Group.
//
// @param request - GetResourceGroupTotalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupTotalResponse
func (client *Client) GetResourceGroupTotalWithContext(ctx context.Context, request *GetResourceGroupTotalRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupTotalResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the historical prices of spot instances.
//
// @param request - GetSpotPriceHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpotPriceHistoryResponse
func (client *Client) GetSpotPriceHistoryWithContext(ctx context.Context, InstanceType *string, request *GetSpotPriceHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSpotPriceHistoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a share token for a training task.
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithContext(ctx context.Context, request *GetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a training job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobResponse
func (client *Client) GetTrainingJobWithContext(ctx context.Context, TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrainingJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the runtime fault error message of a training node.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobErrorInfoResponse
func (client *Client) GetTrainingJobErrorInfoWithContext(ctx context.Context, TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrainingJobErrorInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the latest metrics for a training task.
//
// @param request - GetTrainingJobLatestMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrainingJobLatestMetricsResponse
func (client *Client) GetTrainingJobLatestMetricsWithContext(ctx context.Context, TrainingJobId *string, request *GetTrainingJobLatestMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTrainingJobLatestMetricsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Retrieves metric data at the user dimension.
//
// @param request - GetUserViewMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserViewMetricsResponse
func (client *Client) GetUserViewMetricsWithContext(ctx context.Context, ResourceGroupID *string, request *GetUserViewMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserViewMetricsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve a list of algorithm version details.
//
// @param request - ListAlgorithmVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlgorithmVersionsResponse
func (client *Client) ListAlgorithmVersionsWithContext(ctx context.Context, AlgorithmId *string, request *ListAlgorithmVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlgorithmVersionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists algorithm details.
//
// @param request - ListAlgorithmsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlgorithmsResponse
func (client *Client) ListAlgorithmsWithContext(ctx context.Context, request *ListAlgorithmsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAlgorithmsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves detailed information about all pods on a specified node.
//
// Description:
//
// ## Operation description
//
// - This feature is currently available only on the Lingjun AI Computing Service platform and is accessible to customers through a whitelist.
//
// - `NodeName` and `ResourceGroupId` are required parameters that represent the node name and the ID of the resource group to which the node belongs.
//
// - Optional parameters include `OversoldTypes` and `GPUIndexes`, which are used to further filter or specify conditions.
//
// @param request - ListNodePodsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodePodsResponse
func (client *Client) ListNodePodsWithContext(ctx context.Context, NodeId *string, request *ListNodePodsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodePodsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GPUIndexes) {
		query["GPUIndexes"] = request.GPUIndexes
	}

	if !dara.IsNil(request.OversoldTypes) {
		query["OversoldTypes"] = request.OversoldTypes
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodePods"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/nodes/" + dara.PercentEncode(dara.StringValue(NodeId)) + "/Pods"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodePodsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the available node types.
//
// Description:
//
// ## Description
//
// - Lists the node types that match the specified filter criteria.
//
// - You must specify either the `ResourceGroupIds` or `QuotaId` parameter. The specified resource group IDs or quota ID must belong to the requesting user.
//
// - You can filter the results by using parameters such as accelerator type and GPU type.
//
// - The response includes details and statistics for each node type, indicating the availability of your resources.
//
// @param request - ListNodeTypesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeTypesResponse
func (client *Client) ListNodeTypesWithContext(ctx context.Context, request *ListNodeTypesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodeTypesResponse, _err error) {
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

	if !dara.IsNil(request.GPUType) {
		query["GPUType"] = request.GPUType
	}

	if !dara.IsNil(request.NodeTypes) {
		query["NodeTypes"] = request.NodeTypes
	}

	if !dara.IsNil(request.QuotaId) {
		query["QuotaId"] = request.QuotaId
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeTypes"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/nodetypes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a list of resource nodes.
//
// Description:
//
// ListNodes retrieves resource nodes that match specified filter conditions.
//
// @param tmpReq - ListNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithContext(ctx context.Context, tmpReq *ListNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HealthCount) {
		request.HealthCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HealthCount, dara.String("HealthCount"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HealthRate) {
		request.HealthRateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HealthRate, dara.String("HealthRate"), dara.String("json"))
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

	if !dara.IsNil(request.DiskPL) {
		query["DiskPL"] = request.DiskPL
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

	if !dara.IsNil(request.HealthCountShrink) {
		query["HealthCount"] = request.HealthCountShrink
	}

	if !dara.IsNil(request.HealthRateShrink) {
		query["HealthRate"] = request.HealthRateShrink
	}

	if !dara.IsNil(request.HyperNode) {
		query["HyperNode"] = request.HyperNode
	}

	if !dara.IsNil(request.HyperZone) {
		query["HyperZone"] = request.HyperZone
	}

	if !dara.IsNil(request.LayoutMode) {
		query["LayoutMode"] = request.LayoutMode
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

	if !dara.IsNil(request.PodNum) {
		query["PodNum"] = request.PodNum
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

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WorkloadNum) {
		query["WorkloadNum"] = request.WorkloadNum
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists users within a quota and their resource usage.
//
// Description:
//
// Retrieves the resource usage of users in a specified quota.
//
// @param request - ListQuotaActiveUserUsagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaActiveUserUsagesResponse
func (client *Client) ListQuotaActiveUserUsagesWithContext(ctx context.Context, QuotaId *string, request *ListQuotaActiveUserUsagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotaActiveUserUsagesResponse, _err error) {
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

	if !dara.IsNil(request.SelfOnly) {
		query["SelfOnly"] = request.SelfOnly
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.WorkloadCount) {
		query["WorkloadCount"] = request.WorkloadCount
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotaActiveUserUsages"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/quotas/" + dara.PercentEncode(dara.StringValue(QuotaId)) + "/activeuserusages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaActiveUserUsagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of workloads associated with a specific quota.
//
// @param request - ListQuotaWorkloadsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaWorkloadsResponse
func (client *Client) ListQuotaWorkloadsWithContext(ctx context.Context, QuotaId *string, request *ListQuotaWorkloadsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotaWorkloadsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.Position) {
		query["Position"] = request.Position
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
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

	if !dara.IsNil(request.UseOversoldResource) {
		query["UseOversoldResource"] = request.UseOversoldResource
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

	if !dara.IsNil(request.WorkloadNames) {
		query["WorkloadNames"] = request.WorkloadNames
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists quotas.
//
// @param request - ListQuotasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasResponse
func (client *Client) ListQuotasWithContext(ctx context.Context, request *ListQuotasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
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

	if !dara.IsNil(request.GPUType) {
		query["GPUType"] = request.GPUType
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of machine groups in a specified resource group.
//
// @param request - ListResourceGroupMachineGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupMachineGroupsResponse
func (client *Client) ListResourceGroupMachineGroupsWithContext(ctx context.Context, ResourceGroupID *string, request *ListResourceGroupMachineGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceGroupMachineGroupsResponse, _err error) {
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

	if !dara.IsNil(request.DiskPL) {
		query["DiskPL"] = request.DiskPL
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call this operation to list resource groups for ECS or Lingjun intelligent computing resources.
//
// Description:
//
// Currently, AI computing resource groups (including ECS and Lingjun intelligent computing resources) are available only in specific regions. For a list of supported regions, see [AI computing resource management](https://help.aliyun.com/document_detail/143986.html).
//
// @param request - ListResourceGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithContext(ctx context.Context, request *ListResourceGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tag query API.
//
// Description:
//
// Specify at least one of the following parameters to query for resources.
//
// - ResourceId.N
//
// - Tag.N (Tag.N.Key and Tag.N.Value)
//
// If you specify both parameters, the operation returns only resources that match both criteria.
//
// - Tag.N and ResourceId.N
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve events for a specific training task to gain detailed insight into the training process.
//
// @param request - ListTrainingJobEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobEventsResponse
func (client *Client) ListTrainingJobEventsWithContext(ctx context.Context, TrainingJobId *string, request *ListTrainingJobEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobEventsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve events for a training job instance.
//
// @param request - ListTrainingJobInstanceEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobInstanceEventsResponse
func (client *Client) ListTrainingJobInstanceEventsWithContext(ctx context.Context, TrainingJobId *string, InstanceId *string, request *ListTrainingJobInstanceEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobInstanceEventsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve monitoring metrics for training job instances.
//
// @param request - ListTrainingJobInstanceMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobInstanceMetricsResponse
func (client *Client) ListTrainingJobInstanceMetricsWithContext(ctx context.Context, TrainingJobId *string, request *ListTrainingJobInstanceMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobInstanceMetricsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the logs for a training task.
//
// @param request - ListTrainingJobLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobLogsResponse
func (client *Client) ListTrainingJobLogsWithContext(ctx context.Context, TrainingJobId *string, request *ListTrainingJobLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobLogsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve monitoring data for a training task.
//
// @param request - ListTrainingJobMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobMetricsResponse
func (client *Client) ListTrainingJobMetricsWithContext(ctx context.Context, TrainingJobId *string, request *ListTrainingJobMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobMetricsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about the output models of a training job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobOutputModelsResponse
func (client *Client) ListTrainingJobOutputModelsWithContext(ctx context.Context, TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobOutputModelsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of training jobs.
//
// @param tmpReq - ListTrainingJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrainingJobsResponse
func (client *Client) ListTrainingJobsWithContext(ctx context.Context, tmpReq *ListTrainingJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTrainingJobsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use OperateNode to manage nodes.
//
// Description:
//
// Use OperateNode to manage a specified resource node. Supported operations include disabling or enabling scheduling.
//
// @param request - OperateNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateNodeResponse
func (client *Client) OperateNodeWithContext(ctx context.Context, NodeId *string, request *OperateNodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OperateNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Operation) {
		body["Operation"] = request.Operation
	}

	if !dara.IsNil(request.OperationParameters) {
		body["OperationParameters"] = request.OperationParameters
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateNode"),
		Version:     dara.String("2022-01-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/nodes/" + dara.PercentEncode(dara.StringValue(NodeId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use ScaleQuota to scale resource quotas.
//
// @param request - ScaleQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleQuotaResponse
func (client *Client) ScaleQuotaWithContext(ctx context.Context, QuotaId *string, request *ScaleQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a training task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTrainingJobResponse
func (client *Client) StopTrainingJobWithContext(ctx context.Context, TrainingJobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTrainingJobResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches tags to resources.
//
// Description:
//
// Alibaba Cloud checks the number of existing tags on the resource before attaching a new tag. If the limit is exceeded, an error message is returned.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach tags from a list of specified resources. If a tag is no longer attached to any resource after detachment, Alibaba Cloud deletes it automatically.
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an algorithm.
//
// @param request - UpdateAlgorithmRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlgorithmResponse
func (client *Client) UpdateAlgorithmWithContext(ctx context.Context, AlgorithmId *string, request *UpdateAlgorithmRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlgorithmResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an algorithm version.
//
// @param tmpReq - UpdateAlgorithmVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlgorithmVersionResponse
func (client *Client) UpdateAlgorithmVersionWithContext(ctx context.Context, AlgorithmId *string, AlgorithmVersion *string, tmpReq *UpdateAlgorithmVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAlgorithmVersionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a resource quota.
//
// @param request - UpdateQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuotaWithContext(ctx context.Context, QuotaId *string, request *UpdateQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateQuotaResponse, _err error) {
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

	if !dara.IsNil(request.PropagateDefaultGPUDriver) {
		body["PropagateDefaultGPUDriver"] = request.PropagateDefaultGPUDriver
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use UpdateResourceGroup to update resource group information.
//
// @param request - UpdateResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroupWithContext(ctx context.Context, ResourceGroupID *string, request *UpdateResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the labels of a training task.
//
// @param request - UpdateTrainingJobLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrainingJobLabelsResponse
func (client *Client) UpdateTrainingJobLabelsWithContext(ctx context.Context, TrainingJobId *string, request *UpdateTrainingJobLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTrainingJobLabelsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
