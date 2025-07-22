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
	client.Endpoint, _err = client.GetEndpoint(dara.String("pai-dsw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates an automatic stop policy for a specific Data Science Workshop (DSW) instance. When the conditions are met, the instance is automatically stopped. You can create only one automatic stop policy for an idle DSW instance. If the specific instance has an automatic stop policy, call DeleteIdleInstanceCuller to delete it first.
//
// @param request - CreateIdleInstanceCullerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdleInstanceCullerResponse
func (client *Client) CreateIdleInstanceCullerWithOptions(InstanceId *string, request *CreateIdleInstanceCullerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIdleInstanceCullerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CpuPercentThreshold) {
		body["CpuPercentThreshold"] = request.CpuPercentThreshold
	}

	if !dara.IsNil(request.GpuPercentThreshold) {
		body["GpuPercentThreshold"] = request.GpuPercentThreshold
	}

	if !dara.IsNil(request.MaxIdleTimeInMinutes) {
		body["MaxIdleTimeInMinutes"] = request.MaxIdleTimeInMinutes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdleInstanceCuller"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/idleinstanceculler"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdleInstanceCullerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an automatic stop policy for a specific Data Science Workshop (DSW) instance. When the conditions are met, the instance is automatically stopped. You can create only one automatic stop policy for an idle DSW instance. If the specific instance has an automatic stop policy, call DeleteIdleInstanceCuller to delete it first.
//
// @param request - CreateIdleInstanceCullerRequest
//
// @return CreateIdleInstanceCullerResponse
func (client *Client) CreateIdleInstanceCuller(InstanceId *string, request *CreateIdleInstanceCullerRequest) (_result *CreateIdleInstanceCullerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIdleInstanceCullerResponse{}
	_body, _err := client.CreateIdleInstanceCullerWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Data Science Workshop (DSW) instance.
//
// @param request - CreateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Affinity) {
		body["Affinity"] = request.Affinity
	}

	if !dara.IsNil(request.CloudDisks) {
		body["CloudDisks"] = request.CloudDisks
	}

	if !dara.IsNil(request.CredentialConfig) {
		body["CredentialConfig"] = request.CredentialConfig
	}

	if !dara.IsNil(request.Datasets) {
		body["Datasets"] = request.Datasets
	}

	if !dara.IsNil(request.Driver) {
		body["Driver"] = request.Driver
	}

	if !dara.IsNil(request.DynamicMount) {
		body["DynamicMount"] = request.DynamicMount
	}

	if !dara.IsNil(request.EcsSpec) {
		body["EcsSpec"] = request.EcsSpec
	}

	if !dara.IsNil(request.EnvironmentVariables) {
		body["EnvironmentVariables"] = request.EnvironmentVariables
	}

	if !dara.IsNil(request.ImageAuth) {
		body["ImageAuth"] = request.ImageAuth
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.OversoldType) {
		body["OversoldType"] = request.OversoldType
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RequestedResource) {
		body["RequestedResource"] = request.RequestedResource
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserCommand) {
		body["UserCommand"] = request.UserCommand
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserVpc) {
		body["UserVpc"] = request.UserVpc
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceSource) {
		body["WorkspaceSource"] = request.WorkspaceSource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

// Summary:
//
// Creates a Data Science Workshop (DSW) instance.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scheduled stop task for an instance.
//
// @param request - CreateInstanceShutdownTimerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceShutdownTimerResponse
func (client *Client) CreateInstanceShutdownTimerWithOptions(InstanceId *string, request *CreateInstanceShutdownTimerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceShutdownTimerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DueTime) {
		body["DueTime"] = request.DueTime
	}

	if !dara.IsNil(request.RemainingTimeInMs) {
		body["RemainingTimeInMs"] = request.RemainingTimeInMs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceShutdownTimer"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/shutdowntimer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled stop task for an instance.
//
// @param request - CreateInstanceShutdownTimerRequest
//
// @return CreateInstanceShutdownTimerResponse
func (client *Client) CreateInstanceShutdownTimer(InstanceId *string, request *CreateInstanceShutdownTimerRequest) (_result *CreateInstanceShutdownTimerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceShutdownTimerResponse{}
	_body, _err := client.CreateInstanceShutdownTimerWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实例快照
//
// @param request - CreateInstanceSnapshotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceSnapshotResponse
func (client *Client) CreateInstanceSnapshotWithOptions(InstanceId *string, request *CreateInstanceSnapshotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceSnapshotResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExcludePaths) {
		body["ExcludePaths"] = request.ExcludePaths
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Overwrite) {
		body["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.SnapshotDescription) {
		body["SnapshotDescription"] = request.SnapshotDescription
	}

	if !dara.IsNil(request.SnapshotName) {
		body["SnapshotName"] = request.SnapshotName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceSnapshot"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshots"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例快照
//
// @param request - CreateInstanceSnapshotRequest
//
// @return CreateInstanceSnapshotResponse
func (client *Client) CreateInstanceSnapshot(InstanceId *string, request *CreateInstanceSnapshotRequest) (_result *CreateInstanceSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceSnapshotResponse{}
	_body, _err := client.CreateInstanceSnapshotWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the automatic stop policy of an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdleInstanceCullerResponse
func (client *Client) DeleteIdleInstanceCullerWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIdleInstanceCullerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIdleInstanceCuller"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/idleinstanceculler"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIdleInstanceCullerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the automatic stop policy of an instance.
//
// @return DeleteIdleInstanceCullerResponse
func (client *Client) DeleteIdleInstanceCuller(InstanceId *string) (_result *DeleteIdleInstanceCullerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIdleInstanceCullerResponse{}
	_body, _err := client.DeleteIdleInstanceCullerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

// Summary:
//
// Deletes a specific instance.
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(InstanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete tags from a Data Science Workshop (DSW) instance.
//
// @param request - DeleteInstanceLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceLabelsResponse
func (client *Client) DeleteInstanceLabelsWithOptions(InstanceId *string, request *DeleteInstanceLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceLabelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelKeys) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceLabels"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/labels"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete tags from a Data Science Workshop (DSW) instance.
//
// @param request - DeleteInstanceLabelsRequest
//
// @return DeleteInstanceLabelsResponse
func (client *Client) DeleteInstanceLabels(InstanceId *string, request *DeleteInstanceLabelsRequest) (_result *DeleteInstanceLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceLabelsResponse{}
	_body, _err := client.DeleteInstanceLabelsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scheduled stop task of an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceShutdownTimerResponse
func (client *Client) DeleteInstanceShutdownTimerWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceShutdownTimerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceShutdownTimer"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/shutdowntimer"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scheduled stop task of an instance.
//
// @return DeleteInstanceShutdownTimerResponse
func (client *Client) DeleteInstanceShutdownTimer(InstanceId *string) (_result *DeleteInstanceShutdownTimerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceShutdownTimerResponse{}
	_body, _err := client.DeleteInstanceShutdownTimerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例快照详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceSnapshotResponse
func (client *Client) DeleteInstanceSnapshotWithOptions(InstanceId *string, SnapshotId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceSnapshotResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceSnapshot"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshots/" + dara.PercentEncode(dara.StringValue(SnapshotId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例快照详情
//
// @return DeleteInstanceSnapshotResponse
func (client *Client) DeleteInstanceSnapshot(InstanceId *string, SnapshotId *string) (_result *DeleteInstanceSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceSnapshotResponse{}
	_body, _err := client.DeleteInstanceSnapshotWithOptions(InstanceId, SnapshotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an auto stop policy for a specific idle instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdleInstanceCullerResponse
func (client *Client) GetIdleInstanceCullerWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIdleInstanceCullerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdleInstanceCuller"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/idleinstanceculler"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdleInstanceCullerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an auto stop policy for a specific idle instance.
//
// @return GetIdleInstanceCullerResponse
func (client *Client) GetIdleInstanceCuller(InstanceId *string) (_result *GetIdleInstanceCullerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIdleInstanceCullerResponse{}
	_body, _err := client.GetIdleInstanceCullerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a DSW instance.
//
// @param request - GetInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(InstanceId *string, request *GetInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

// Summary:
//
// Queries the details of a DSW instance.
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(InstanceId *string, request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of system events for a Data Science Workshop (DSW) instance.
//
// @param request - GetInstanceEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceEventsResponse
func (client *Client) GetInstanceEventsWithOptions(InstanceId *string, request *GetInstanceEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxEventsNum) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceEvents"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of system events for a Data Science Workshop (DSW) instance.
//
// @param request - GetInstanceEventsRequest
//
// @return GetInstanceEventsResponse
func (client *Client) GetInstanceEvents(InstanceId *string, request *GetInstanceEventsRequest) (_result *GetInstanceEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceEventsResponse{}
	_body, _err := client.GetInstanceEventsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource metrics of an instance.
//
// @param request - GetInstanceMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceMetricsResponse
func (client *Client) GetInstanceMetricsWithOptions(InstanceId *string, request *GetInstanceMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceMetricsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("GetInstanceMetrics"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instance/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource metrics of an instance.
//
// @param request - GetInstanceMetricsRequest
//
// @return GetInstanceMetricsResponse
func (client *Client) GetInstanceMetrics(InstanceId *string, request *GetInstanceMetricsRequest) (_result *GetInstanceMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceMetricsResponse{}
	_body, _err := client.GetInstanceMetricsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取定时关机任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceShutdownTimerResponse
func (client *Client) GetInstanceShutdownTimerWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceShutdownTimerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceShutdownTimer"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/shutdowntimer"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceShutdownTimerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取定时关机任务
//
// @return GetInstanceShutdownTimerResponse
func (client *Client) GetInstanceShutdownTimer(InstanceId *string) (_result *GetInstanceShutdownTimerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceShutdownTimerResponse{}
	_body, _err := client.GetInstanceShutdownTimerWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例快照详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceSnapshotResponse
func (client *Client) GetInstanceSnapshotWithOptions(InstanceId *string, SnapshotId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceSnapshotResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceSnapshot"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshots/" + dara.PercentEncode(dara.StringValue(SnapshotId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例快照详情
//
// @return GetInstanceSnapshotResponse
func (client *Client) GetInstanceSnapshot(InstanceId *string, SnapshotId *string) (_result *GetInstanceSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceSnapshotResponse{}
	_body, _err := client.GetInstanceSnapshotWithOptions(InstanceId, SnapshotId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtains the lifecycle of an instance
//
// Description:
//
// Obtains the lifecycle transition information for an instance, including details on the status an instance transitions to at a specific point in time.
//
// @param request - GetLifecycleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLifecycleResponse
func (client *Client) GetLifecycleWithOptions(InstanceId *string, request *GetLifecycleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLifecycleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.SessionNumber) {
		query["SessionNumber"] = request.SessionNumber
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLifecycle"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/lifecycle"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLifecycleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtains the lifecycle of an instance
//
// Description:
//
// Obtains the lifecycle transition information for an instance, including details on the status an instance transitions to at a specific point in time.
//
// @param request - GetLifecycleRequest
//
// @return GetLifecycleResponse
func (client *Client) GetLifecycle(InstanceId *string, request *GetLifecycleRequest) (_result *GetLifecycleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLifecycleResponse{}
	_body, _err := client.GetLifecycleWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取metrics数据
//
// @param request - GetMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetricsResponse
func (client *Client) GetMetricsWithOptions(InstanceId *string, request *GetMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMetricsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetrics"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instance/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/cms/metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取metrics数据
//
// @param request - GetMetricsRequest
//
// @return GetMetricsResponse
func (client *Client) GetMetrics(InstanceId *string, request *GetMetricsRequest) (_result *GetMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMetricsResponse{}
	_body, _err := client.GetMetricsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetResourceGroupStatisticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupStatisticsResponse
func (client *Client) GetResourceGroupStatisticsWithOptions(request *GetResourceGroupStatisticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceGroupStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceGroupStatistics"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/resourcegroupstatistics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetResourceGroupStatisticsRequest
//
// @return GetResourceGroupStatisticsResponse
func (client *Client) GetResourceGroupStatistics(request *GetResourceGroupStatisticsRequest) (_result *GetResourceGroupStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceGroupStatisticsResponse{}
	_body, _err := client.GetResourceGroupStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the temporary authentication information of a DSW instance.
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/tokens"),
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
// Obtains the temporary authentication information of a DSW instance.
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
// 获取自定义用户命令
//
// @param request - GetUserCommandRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserCommandResponse
func (client *Client) GetUserCommandWithOptions(UserCommandId *string, request *GetUserCommandRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserCommandResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserCommand"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/usercommands/" + dara.PercentEncode(dara.StringValue(UserCommandId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义用户命令
//
// @param request - GetUserCommandRequest
//
// @return GetUserCommandResponse
func (client *Client) GetUserCommand(UserCommandId *string, request *GetUserCommandRequest) (_result *GetUserCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserCommandResponse{}
	_body, _err := client.GetUserCommandWithOptions(UserCommandId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserConfigResponse
func (client *Client) GetUserConfigWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/userconfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户配置
//
// @return GetUserConfigResponse
func (client *Client) GetUserConfig() (_result *GetUserConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserConfigResponse{}
	_body, _err := client.GetUserConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of specifications of ECS instances.
//
// @param request - ListEcsSpecsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEcsSpecsResponse
func (client *Client) ListEcsSpecsWithOptions(request *ListEcsSpecsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEcsSpecsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorType) {
		query["AcceleratorType"] = request.AcceleratorType
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

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEcsSpecs"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/ecsspecs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of specifications of ECS instances.
//
// @param request - ListEcsSpecsRequest
//
// @return ListEcsSpecsResponse
func (client *Client) ListEcsSpecs(request *ListEcsSpecsRequest) (_result *ListEcsSpecsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.ListEcsSpecsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例快照列表
//
// @param request - ListInstanceSnapshotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceSnapshotResponse
func (client *Client) ListInstanceSnapshotWithOptions(InstanceId *string, request *ListInstanceSnapshotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceSnapshotResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceSnapshot"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/snapshots"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例快照列表
//
// @param request - ListInstanceSnapshotRequest
//
// @return ListInstanceSnapshotResponse
func (client *Client) ListInstanceSnapshot(InstanceId *string, request *ListInstanceSnapshotRequest) (_result *ListInstanceSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceSnapshotResponse{}
	_body, _err := client.ListInstanceSnapshotWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例统计信息
//
// @param request - ListInstanceStatisticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceStatisticsResponse
func (client *Client) ListInstanceStatisticsWithOptions(request *ListInstanceStatisticsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceIds) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceStatistics"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instancestatistics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例统计信息
//
// @param request - ListInstanceStatisticsRequest
//
// @return ListInstanceStatisticsResponse
func (client *Client) ListInstanceStatistics(request *ListInstanceStatisticsRequest) (_result *ListInstanceStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceStatisticsResponse{}
	_body, _err := client.ListInstanceStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Data Science Workshop (DSW) instances.
//
// @param tmpReq - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Labels) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, dara.String("Labels"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorType) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.CreateUserId) {
		query["CreateUserId"] = request.CreateUserId
	}

	if !dara.IsNil(request.GpuType) {
		query["GpuType"] = request.GpuType
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.LabelsShrink) {
		query["Labels"] = request.LabelsShrink
	}

	if !dara.IsNil(request.MaxCpu) {
		query["MaxCpu"] = request.MaxCpu
	}

	if !dara.IsNil(request.MaxGpu) {
		query["MaxGpu"] = request.MaxGpu
	}

	if !dara.IsNil(request.MaxGpuMemory) {
		query["MaxGpuMemory"] = request.MaxGpuMemory
	}

	if !dara.IsNil(request.MaxMemory) {
		query["MaxMemory"] = request.MaxMemory
	}

	if !dara.IsNil(request.MinCpu) {
		query["MinCpu"] = request.MinCpu
	}

	if !dara.IsNil(request.MinGpu) {
		query["MinGpu"] = request.MinGpu
	}

	if !dara.IsNil(request.MinGpuMemory) {
		query["MinGpuMemory"] = request.MinGpuMemory
	}

	if !dara.IsNil(request.MinMemory) {
		query["MinMemory"] = request.MinMemory
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OversoldInfo) {
		query["OversoldInfo"] = request.OversoldInfo
	}

	if !dara.IsNil(request.OversoldType) {
		query["OversoldType"] = request.OversoldType
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

// Summary:
//
// Queries a list of Data Science Workshop (DSW) instances.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取系统日志
//
// @param request - ListSystemLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSystemLogsResponse
func (client *Client) ListSystemLogsWithOptions(request *ListSystemLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSystemLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GmtEndTime) {
		query["GmtEndTime"] = request.GmtEndTime
	}

	if !dara.IsNil(request.GmtStartTime) {
		query["GmtStartTime"] = request.GmtStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogLevel) {
		query["LogLevel"] = request.LogLevel
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

	if !dara.IsNil(request.ProblemCategory) {
		query["ProblemCategory"] = request.ProblemCategory
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SourceRequestId) {
		query["SourceRequestId"] = request.SourceRequestId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSystemLogs"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/systemlogs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSystemLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统日志
//
// @param request - ListSystemLogsRequest
//
// @return ListSystemLogsResponse
func (client *Client) ListSystemLogs(request *ListSystemLogsRequest) (_result *ListSystemLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSystemLogsResponse{}
	_body, _err := client.ListSystemLogsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动实例
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithOptions(InstanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartInstance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/start"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动实例
//
// @return StartInstanceResponse
func (client *Client) StartInstance(InstanceId *string) (_result *StartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(InstanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止实例
//
// @param request - StopInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithOptions(InstanceId *string, request *StopInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SaveImage) {
		query["SaveImage"] = request.SaveImage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInstance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止实例
//
// @param request - StopInstanceRequest
//
// @return StopInstanceResponse
func (client *Client) StopInstance(InstanceId *string, request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the properties of a DSW instance.
//
// @param request - UpdateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(InstanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Affinity) {
		body["Affinity"] = request.Affinity
	}

	if !dara.IsNil(request.CloudDisks) {
		body["CloudDisks"] = request.CloudDisks
	}

	if !dara.IsNil(request.CredentialConfig) {
		body["CredentialConfig"] = request.CredentialConfig
	}

	if !dara.IsNil(request.Datasets) {
		body["Datasets"] = request.Datasets
	}

	if !dara.IsNil(request.DisassociateCredential) {
		body["DisassociateCredential"] = request.DisassociateCredential
	}

	if !dara.IsNil(request.DisassociateDatasets) {
		body["DisassociateDatasets"] = request.DisassociateDatasets
	}

	if !dara.IsNil(request.DisassociateDriver) {
		body["DisassociateDriver"] = request.DisassociateDriver
	}

	if !dara.IsNil(request.DisassociateEnvironmentVariables) {
		body["DisassociateEnvironmentVariables"] = request.DisassociateEnvironmentVariables
	}

	if !dara.IsNil(request.DisassociateForwardInfos) {
		body["DisassociateForwardInfos"] = request.DisassociateForwardInfos
	}

	if !dara.IsNil(request.DisassociateUserCommand) {
		body["DisassociateUserCommand"] = request.DisassociateUserCommand
	}

	if !dara.IsNil(request.DisassociateVpc) {
		body["DisassociateVpc"] = request.DisassociateVpc
	}

	if !dara.IsNil(request.Driver) {
		body["Driver"] = request.Driver
	}

	if !dara.IsNil(request.DynamicMount) {
		body["DynamicMount"] = request.DynamicMount
	}

	if !dara.IsNil(request.EcsSpec) {
		body["EcsSpec"] = request.EcsSpec
	}

	if !dara.IsNil(request.EnvironmentVariables) {
		body["EnvironmentVariables"] = request.EnvironmentVariables
	}

	if !dara.IsNil(request.ImageAuth) {
		body["ImageAuth"] = request.ImageAuth
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageUrl) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.OversoldType) {
		body["OversoldType"] = request.OversoldType
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RequestedResource) {
		body["RequestedResource"] = request.RequestedResource
	}

	if !dara.IsNil(request.UserCommand) {
		body["UserCommand"] = request.UserCommand
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserVpc) {
		body["UserVpc"] = request.UserVpc
	}

	if !dara.IsNil(request.WorkspaceSource) {
		body["WorkspaceSource"] = request.WorkspaceSource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the properties of a DSW instance.
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(InstanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the tags of a Data Science Workshop (DSW) instance. If the tags do not exist, the system adds tags.
//
// @param request - UpdateInstanceLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceLabelsResponse
func (client *Client) UpdateInstanceLabelsWithOptions(InstanceId *string, request *UpdateInstanceLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceLabelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("UpdateInstanceLabels"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/instances/" + dara.PercentEncode(dara.StringValue(InstanceId)) + "/labels"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the tags of a Data Science Workshop (DSW) instance. If the tags do not exist, the system adds tags.
//
// @param request - UpdateInstanceLabelsRequest
//
// @return UpdateInstanceLabelsResponse
func (client *Client) UpdateInstanceLabels(InstanceId *string, request *UpdateInstanceLabelsRequest) (_result *UpdateInstanceLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceLabelsResponse{}
	_body, _err := client.UpdateInstanceLabelsWithOptions(InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
