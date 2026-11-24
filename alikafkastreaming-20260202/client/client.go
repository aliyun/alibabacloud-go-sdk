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
	client.ProductId = dara.String("AlikafkaStreaming")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("alikafkastreaming"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Validates whether a specified Flink SQL statement complies with the specifications.
//
// Description:
//
// ## Operation description
//
// - You can invoke this operation by using the GET or POST method.
//
// - You must specify the `InstanceId`, `JobName`, and `SqlContent` parameters. The `SqlContent` parameter specifies the Flink SQL statement to be validated.
//
// - In the response, the `Data.Valid` field indicates whether the SQL statement passed the validation. If the validation fails, the fault details are returned in the `Data.ErrorList` field.
//
// - The current version requires both the instance ID (`InstanceId`) and the job name (`JobName`) to build the job context.
//
// - A successful response only indicates that the system starts to execute the validation flow. It does not directly indicate whether the SQL statement is valid. Check the `Data.Valid` field to determine whether the SQL statement is valid.
//
// - For error codes and exception handling, refer to the Error codes section in the References.
//
// @param request - CheckSqlContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSqlContentResponse
func (client *Client) CheckSqlContentWithOptions(request *CheckSqlContentRequest, runtime *dara.RuntimeOptions) (_result *CheckSqlContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SqlContent) {
		query["SqlContent"] = request.SqlContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckSqlContent"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSqlContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates whether a specified Flink SQL statement complies with the specifications.
//
// Description:
//
// ## Operation description
//
// - You can invoke this operation by using the GET or POST method.
//
// - You must specify the `InstanceId`, `JobName`, and `SqlContent` parameters. The `SqlContent` parameter specifies the Flink SQL statement to be validated.
//
// - In the response, the `Data.Valid` field indicates whether the SQL statement passed the validation. If the validation fails, the fault details are returned in the `Data.ErrorList` field.
//
// - The current version requires both the instance ID (`InstanceId`) and the job name (`JobName`) to build the job context.
//
// - A successful response only indicates that the system starts to execute the validation flow. It does not directly indicate whether the SQL statement is valid. Check the `Data.Valid` field to determine whether the SQL statement is valid.
//
// - For error codes and exception handling, refer to the Error codes section in the References.
//
// @param request - CheckSqlContentRequest
//
// @return CheckSqlContentResponse
func (client *Client) CheckSqlContent(request *CheckSqlContentRequest) (_result *CheckSqlContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckSqlContentResponse{}
	_body, _err := client.CheckSqlContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a compute instance. This operation only completes the purchase phase. After the instance is created, call StartComputeInstance to complete network configuration and deployment.
//
// Description:
//
// Creates a compute instance. This operation only completes the purchase phase. After the instance is created, call StartComputeInstance to complete network configuration and deployment.
//
// - API version: 2026-02-02
//
// - Action: CreateComputeInstance
//
// @param request - CreateComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeInstanceResponse
func (client *Client) CreateComputeInstanceWithOptions(request *CreateComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PaidType) {
		query["PaidType"] = request.PaidType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a compute instance. This operation only completes the purchase phase. After the instance is created, call StartComputeInstance to complete network configuration and deployment.
//
// Description:
//
// Creates a compute instance. This operation only completes the purchase phase. After the instance is created, call StartComputeInstance to complete network configuration and deployment.
//
// - API version: 2026-02-02
//
// - Action: CreateComputeInstance
//
// @param request - CreateComputeInstanceRequest
//
// @return CreateComputeInstanceResponse
func (client *Client) CreateComputeInstance(request *CreateComputeInstanceRequest) (_result *CreateComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComputeInstanceResponse{}
	_body, _err := client.CreateComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Flink SQL job on a running compute instance.
//
// Description:
//
// ## Request description
//
// - This API creates a new Flink SQL job on a specified running compute instance.
//
// - The created job is in the `INIT` state.
//
// - You can control the resource usage of the job by setting `CuLimit` and `CuReserved`.
//
// - The `Remark` field allows you to add remarks to the job for easier management and identification.
//
// - Make sure that the `RegionId`, `InstanceId`, and `JobName` parameters are accurate. Otherwise, the request may fail.
//
// - If you attempt to create a job with a duplicate name, an error is returned.
//
// - The compute instance must be in the running state to successfully create a job.
//
// @param request - CreateComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeJobResponse
func (client *Client) CreateComputeJobWithOptions(request *CreateComputeJobRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CuLimit) {
		query["CuLimit"] = request.CuLimit
	}

	if !dara.IsNil(request.CuReserved) {
		query["CuReserved"] = request.CuReserved
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Flink SQL job on a running compute instance.
//
// Description:
//
// ## Request description
//
// - This API creates a new Flink SQL job on a specified running compute instance.
//
// - The created job is in the `INIT` state.
//
// - You can control the resource usage of the job by setting `CuLimit` and `CuReserved`.
//
// - The `Remark` field allows you to add remarks to the job for easier management and identification.
//
// - Make sure that the `RegionId`, `InstanceId`, and `JobName` parameters are accurate. Otherwise, the request may fail.
//
// - If you attempt to create a job with a duplicate name, an error is returned.
//
// - The compute instance must be in the running state to successfully create a job.
//
// @param request - CreateComputeJobRequest
//
// @return CreateComputeJobResponse
func (client *Client) CreateComputeJob(request *CreateComputeJobRequest) (_result *CreateComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComputeJobResponse{}
	_body, _err := client.CreateComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a compute instance that is in the pending deployment, stopped, or released state.
//
// Description:
//
// Deletes a compute instance that is in the pending deployment, stopped, or released state.
//
// - API version: 2026-02-02
//
// - Action: DeleteComputeInstance
//
// @param request - DeleteComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeInstanceResponse
func (client *Client) DeleteComputeInstanceWithOptions(request *DeleteComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a compute instance that is in the pending deployment, stopped, or released state.
//
// Description:
//
// Deletes a compute instance that is in the pending deployment, stopped, or released state.
//
// - API version: 2026-02-02
//
// - Action: DeleteComputeInstance
//
// @param request - DeleteComputeInstanceRequest
//
// @return DeleteComputeInstanceResponse
func (client *Client) DeleteComputeInstance(request *DeleteComputeInstanceRequest) (_result *DeleteComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComputeInstanceResponse{}
	_body, _err := client.DeleteComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified compute job. A successful request indicates that the deletion request has been accepted.
//
// Description:
//
// ## Request description
//
// - This operation deletes a specific compute job.
//
// - A successful call to this operation only indicates that the deletion request has been accepted by the system. The deletion is not completed immediately.
//
// - Make sure that the `RegionId`, `InstanceId`, and `JobName` parameters are accurate. Otherwise, the request may fail.
//
// - If the compute instance or job is in a state that does not allow deletion, such as a non-running state, an error message is returned.
//
// - The deletion operation is irreversible. Proceed with caution.
//
// @param request - DeleteComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeJobResponse
func (client *Client) DeleteComputeJobWithOptions(request *DeleteComputeJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified compute job. A successful request indicates that the deletion request has been accepted.
//
// Description:
//
// ## Request description
//
// - This operation deletes a specific compute job.
//
// - A successful call to this operation only indicates that the deletion request has been accepted by the system. The deletion is not completed immediately.
//
// - Make sure that the `RegionId`, `InstanceId`, and `JobName` parameters are accurate. Otherwise, the request may fail.
//
// - If the compute instance or job is in a state that does not allow deletion, such as a non-running state, an error message is returned.
//
// - The deletion operation is irreversible. Proceed with caution.
//
// @param request - DeleteComputeJobRequest
//
// @return DeleteComputeJobResponse
func (client *Client) DeleteComputeJob(request *DeleteComputeJobRequest) (_result *DeleteComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComputeJobResponse{}
	_body, _err := client.DeleteComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified compute instance.
//
// @param request - GetComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeInstanceResponse
func (client *Client) GetComputeInstanceWithOptions(request *GetComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified compute instance.
//
// @param request - GetComputeInstanceRequest
//
// @return GetComputeInstanceResponse
func (client *Client) GetComputeInstance(request *GetComputeInstanceRequest) (_result *GetComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetComputeInstanceResponse{}
	_body, _err := client.GetComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a compute job in Alikafka Streaming by using a parameter query.
//
// Description:
//
// ## Request description
//
// - This operation queries the details of a specified compute job.
//
// - You can use the GET or POST method to send requests.
//
// - The `RegionId`, `InstanceId`, and `JobName` parameters are required.
//
// - The authorization action is `alikafkastreaming:GetComputeJob`, and the access level is Read.
//
// @param request - GetComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeJobResponse
func (client *Client) GetComputeJobWithOptions(request *GetComputeJobRequest, runtime *dara.RuntimeOptions) (_result *GetComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a compute job in Alikafka Streaming by using a parameter query.
//
// Description:
//
// ## Request description
//
// - This operation queries the details of a specified compute job.
//
// - You can use the GET or POST method to send requests.
//
// - The `RegionId`, `InstanceId`, and `JobName` parameters are required.
//
// - The authorization action is `alikafkastreaming:GetComputeJob`, and the access level is Read.
//
// @param request - GetComputeJobRequest
//
// @return GetComputeJobResponse
func (client *Client) GetComputeJob(request *GetComputeJobRequest) (_result *GetComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetComputeJobResponse{}
	_body, _err := client.GetComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取debug信息
//
// @param request - GetJobDebugDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDebugDataResponse
func (client *Client) GetJobDebugDataWithOptions(request *GetJobDebugDataRequest, runtime *dara.RuntimeOptions) (_result *GetJobDebugDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobDebugData"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobDebugDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取debug信息
//
// @param request - GetJobDebugDataRequest
//
// @return GetJobDebugDataResponse
func (client *Client) GetJobDebugData(request *GetJobDebugDataRequest) (_result *GetJobDebugDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobDebugDataResponse{}
	_body, _err := client.GetJobDebugDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例列表（分页）
//
// @param tmpReq - ListComputeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeInstancesResponse
func (client *Client) ListComputeInstancesWithOptions(tmpReq *ListComputeInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListComputeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListComputeInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeInstances"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例列表（分页）
//
// @param request - ListComputeInstancesRequest
//
// @return ListComputeInstancesResponse
func (client *Client) ListComputeInstances(request *ListComputeInstancesRequest) (_result *ListComputeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComputeInstancesResponse{}
	_body, _err := client.ListComputeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries compute instances under the current account by using paging.
//
// @param tmpReq - ListComputeInstancesInPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeInstancesInPageResponse
func (client *Client) ListComputeInstancesInPageWithOptions(tmpReq *ListComputeInstancesInPageRequest, runtime *dara.RuntimeOptions) (_result *ListComputeInstancesInPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListComputeInstancesInPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeInstancesInPage"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeInstancesInPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries compute instances under the current account by using paging.
//
// @param request - ListComputeInstancesInPageRequest
//
// @return ListComputeInstancesInPageResponse
func (client *Client) ListComputeInstancesInPage(request *ListComputeInstancesInPageRequest) (_result *ListComputeInstancesInPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComputeInstancesInPageResponse{}
	_body, _err := client.ListComputeInstancesInPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all job information under a specified compute instance by using paging.
//
// Description:
//
// ## Request description
//
// - This operation supports a cursor-based paged query by using the `MaxResults` and `NextToken` parameters.
//
// - You do not need to specify `NextToken` in the first request. For subsequent requests, use the `NextToken` value returned in the previous response.
//
// - You can search by job name or remark, and specify different sorting fields and directions.
//
// - The authorization action is `alikafkastreaming:ListComputeJobs`. The access level is List, which applies to all resources.
//
// @param request - ListComputeJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeJobsResponse
func (client *Client) ListComputeJobsWithOptions(request *ListComputeJobsRequest, runtime *dara.RuntimeOptions) (_result *ListComputeJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.SortField) {
		query["SortField"] = request.SortField
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeJobs"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all job information under a specified compute instance by using paging.
//
// Description:
//
// ## Request description
//
// - This operation supports a cursor-based paged query by using the `MaxResults` and `NextToken` parameters.
//
// - You do not need to specify `NextToken` in the first request. For subsequent requests, use the `NextToken` value returned in the previous response.
//
// - You can search by job name or remark, and specify different sorting fields and directions.
//
// - The authorization action is `alikafkastreaming:ListComputeJobs`. The access level is List, which applies to all resources.
//
// @param request - ListComputeJobsRequest
//
// @return ListComputeJobsResponse
func (client *Client) ListComputeJobs(request *ListComputeJobsRequest) (_result *ListComputeJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComputeJobsResponse{}
	_body, _err := client.ListComputeJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询创建 SQL 任务时支持的连接器列表
//
// @param request - ListSupportedConnectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupportedConnectorsResponse
func (client *Client) ListSupportedConnectorsWithOptions(request *ListSupportedConnectorsRequest, runtime *dara.RuntimeOptions) (_result *ListSupportedConnectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSupportedConnectors"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSupportedConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询创建 SQL 任务时支持的连接器列表
//
// @param request - ListSupportedConnectorsRequest
//
// @return ListSupportedConnectorsResponse
func (client *Client) ListSupportedConnectors(request *ListSupportedConnectorsRequest) (_result *ListSupportedConnectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSupportedConnectorsResponse{}
	_body, _err := client.ListSupportedConnectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Re-enables a stopped pay-as-you-go compute instance. A successful response indicates that the re-enablement request has been accepted.
//
// Description:
//
// Re-enables a stopped pay-as-you-go compute instance. A successful response indicates that the re-enablement request has been accepted.
//
// - API version: 2026-02-02
//
// - Action: ReopenComputeInstance
//
// @param request - ReopenComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReopenComputeInstanceResponse
func (client *Client) ReopenComputeInstanceWithOptions(request *ReopenComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReopenComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReopenComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReopenComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-enables a stopped pay-as-you-go compute instance. A successful response indicates that the re-enablement request has been accepted.
//
// Description:
//
// Re-enables a stopped pay-as-you-go compute instance. A successful response indicates that the re-enablement request has been accepted.
//
// - API version: 2026-02-02
//
// - Action: ReopenComputeInstance
//
// @param request - ReopenComputeInstanceRequest
//
// @return ReopenComputeInstanceResponse
func (client *Client) ReopenComputeInstance(request *ReopenComputeInstanceRequest) (_result *ReopenComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReopenComputeInstanceResponse{}
	_body, _err := client.ReopenComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启 JOB
//
// @param request - RestartComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartComputeJobResponse
func (client *Client) RestartComputeJobWithOptions(request *RestartComputeJobRequest, runtime *dara.RuntimeOptions) (_result *RestartComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启 JOB
//
// @param request - RestartComputeJobRequest
//
// @return RestartComputeJobResponse
func (client *Client) RestartComputeJob(request *RestartComputeJobRequest) (_result *RestartComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartComputeJobResponse{}
	_body, _err := client.RestartComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the network and deploys a compute instance in the pending deployment state.
//
// Description:
//
// Configures the network and deploys a compute instance in the pending deployment state.
//
// - API version: 2026-02-02
//
// - Action: StartComputeInstance
//
// @param tmpReq - StartComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartComputeInstanceResponse
func (client *Client) StartComputeInstanceWithOptions(tmpReq *StartComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartComputeInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VSwitchIds) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, dara.String("VSwitchIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchIdsShrink) {
		query["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the network and deploys a compute instance in the pending deployment state.
//
// Description:
//
// Configures the network and deploys a compute instance in the pending deployment state.
//
// - API version: 2026-02-02
//
// - Action: StartComputeInstance
//
// @param request - StartComputeInstanceRequest
//
// @return StartComputeInstanceResponse
func (client *Client) StartComputeInstance(request *StartComputeInstanceRequest) (_result *StartComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartComputeInstanceResponse{}
	_body, _err := client.StartComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts, resumes, or debugs a compute job.
//
// Description:
//
// ## Request description
//
// - `RecoveryMode` supports two modes: `savepoint` and `stateless`. If you select the `savepoint` mode but no savepoint is available, an error is returned.
//
// - The `CuLimit` and `CuReserved` parameters are used to set the CU limit and the number of reserved CUs for the job, respectively. Integers and decimals are supported.
//
// - Make sure that the values of the `RegionId`, `InstanceId`, and `JobName` parameters are correct and exist. Otherwise, the request fails.
//
// @param request - StartComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartComputeJobResponse
func (client *Client) StartComputeJobWithOptions(request *StartComputeJobRequest, runtime *dara.RuntimeOptions) (_result *StartComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CuLimit) {
		query["CuLimit"] = request.CuLimit
	}

	if !dara.IsNil(request.CuReserved) {
		query["CuReserved"] = request.CuReserved
	}

	if !dara.IsNil(request.DraftSql) {
		query["DraftSql"] = request.DraftSql
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RecoveryMode) {
		query["RecoveryMode"] = request.RecoveryMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts, resumes, or debugs a compute job.
//
// Description:
//
// ## Request description
//
// - `RecoveryMode` supports two modes: `savepoint` and `stateless`. If you select the `savepoint` mode but no savepoint is available, an error is returned.
//
// - The `CuLimit` and `CuReserved` parameters are used to set the CU limit and the number of reserved CUs for the job, respectively. Integers and decimals are supported.
//
// - Make sure that the values of the `RegionId`, `InstanceId`, and `JobName` parameters are correct and exist. Otherwise, the request fails.
//
// @param request - StartComputeJobRequest
//
// @return StartComputeJobResponse
func (client *Client) StartComputeJob(request *StartComputeJobRequest) (_result *StartComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartComputeJobResponse{}
	_body, _err := client.StartComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a running pay-as-you-go compute instance. A successful response indicates that the stop request has been accepted.
//
// Description:
//
// Stops a running pay-as-you-go compute instance. A successful response indicates that the stop request has been accepted.
//
// - API version: 2026-02-02
//
// - Action: StopComputeInstance
//
// @param request - StopComputeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopComputeInstanceResponse
func (client *Client) StopComputeInstanceWithOptions(request *StopComputeInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopComputeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopComputeInstance"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopComputeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a running pay-as-you-go compute instance. A successful response indicates that the stop request has been accepted.
//
// Description:
//
// Stops a running pay-as-you-go compute instance. A successful response indicates that the stop request has been accepted.
//
// - API version: 2026-02-02
//
// - Action: StopComputeInstance
//
// @param request - StopComputeInstanceRequest
//
// @return StopComputeInstanceResponse
func (client *Client) StopComputeInstance(request *StopComputeInstanceRequest) (_result *StopComputeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopComputeInstanceResponse{}
	_body, _err := client.StopComputeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a compute job in a specified region and instance.
//
// Description:
//
// ## Request description
//
// - This operation stops the production or debug runtime instance of a specified compute job.
//
// - A successful response indicates that the stop request has been accepted, but it does not mean that the job stops immediately.
//
// @param request - StopComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopComputeJobResponse
func (client *Client) StopComputeJobWithOptions(request *StopComputeJobRequest, runtime *dara.RuntimeOptions) (_result *StopComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a compute job in a specified region and instance.
//
// Description:
//
// ## Request description
//
// - This operation stops the production or debug runtime instance of a specified compute job.
//
// - A successful response indicates that the stop request has been accepted, but it does not mean that the job stops immediately.
//
// @param request - StopComputeJobRequest
//
// @return StopComputeJobResponse
func (client *Client) StopComputeJob(request *StopComputeJobRequest) (_result *StopComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopComputeJobResponse{}
	_body, _err := client.StopComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name of a compute instance. The instance must be in the deployment preparation or running state.
//
// Description:
//
// Updates the name of a compute instance. The instance must be in the deployment preparation or running state.
//
// - API version: 2026-02-02
//
// - Action: UpdateComputeInstanceName
//
// @param request - UpdateComputeInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeInstanceNameResponse
func (client *Client) UpdateComputeInstanceNameWithOptions(request *UpdateComputeInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeInstanceName"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name of a compute instance. The instance must be in the deployment preparation or running state.
//
// Description:
//
// Updates the name of a compute instance. The instance must be in the deployment preparation or running state.
//
// - API version: 2026-02-02
//
// - Action: UpdateComputeInstanceName
//
// @param request - UpdateComputeInstanceNameRequest
//
// @return UpdateComputeInstanceNameResponse
func (client *Client) UpdateComputeInstanceName(request *UpdateComputeInstanceNameRequest) (_result *UpdateComputeInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeInstanceNameResponse{}
	_body, _err := client.UpdateComputeInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the remarks of a specified compute job.
//
// Description:
//
// ## Request description
//
// - Ensure that the provided `InstanceId` and `JobName` are valid. Otherwise, an error is returned.
//
// - You cannot execute this operation if the instance is not in the running state.
//
// - Modifications are not supported if the current job is running as a debug node or undergoing changes.
//
// @param request - UpdateComputeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobResponse
func (client *Client) UpdateComputeJobWithOptions(request *UpdateComputeJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeJob"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the remarks of a specified compute job.
//
// Description:
//
// ## Request description
//
// - Ensure that the provided `InstanceId` and `JobName` are valid. Otherwise, an error is returned.
//
// - You cannot execute this operation if the instance is not in the running state.
//
// - Modifications are not supported if the current job is running as a debug node or undergoing changes.
//
// @param request - UpdateComputeJobRequest
//
// @return UpdateComputeJobResponse
func (client *Client) UpdateComputeJob(request *UpdateComputeJobRequest) (_result *UpdateComputeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeJobResponse{}
	_body, _err := client.UpdateComputeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the compute unit (CU) limit and the number of reserved CUs for a specified compute job.
//
// Description:
//
// ## Request description
//
// This operation allows you to modify the compute unit (CU) limit and the number of reserved CUs for a specific compute job. Before calling this operation, ensure that the specified `InstanceId` and `JobName` are correct and that the instance is in the running state. Additionally, verify the validity and reasonableness of the `CuLimit` and `CuReserved` parameters to prevent request failures caused by exceeding limits or violating business logic.
//
// @param request - UpdateComputeJobCuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobCuResponse
func (client *Client) UpdateComputeJobCuWithOptions(request *UpdateComputeJobCuRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobCuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CuLimit) {
		query["CuLimit"] = request.CuLimit
	}

	if !dara.IsNil(request.CuReserved) {
		query["CuReserved"] = request.CuReserved
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeJobCu"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeJobCuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the compute unit (CU) limit and the number of reserved CUs for a specified compute job.
//
// Description:
//
// ## Request description
//
// This operation allows you to modify the compute unit (CU) limit and the number of reserved CUs for a specific compute job. Before calling this operation, ensure that the specified `InstanceId` and `JobName` are correct and that the instance is in the running state. Additionally, verify the validity and reasonableness of the `CuLimit` and `CuReserved` parameters to prevent request failures caused by exceeding limits or violating business logic.
//
// @param request - UpdateComputeJobCuRequest
//
// @return UpdateComputeJobCuResponse
func (client *Client) UpdateComputeJobCu(request *UpdateComputeJobCuRequest) (_result *UpdateComputeJobCuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeJobCuResponse{}
	_body, _err := client.UpdateComputeJobCuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the Flink SQL draft saved for a specified compute job.
//
// Description:
//
// ## Request description
//
// This operation updates the Flink SQL draft saved for a specific compute job under a given compute instance. Make sure that the provided `InstanceId` and `JobName` are accurate, and the current status of the job supports SQL modification.
//
// - **Precautions**:
//
//   - Make sure that the target instance is in the running state.
//
//   - The current job status must allow SQL modification. The job must not be undergoing debugging or modification.
//
//   - The `DraftSql` parameter must contain a complete and correctly formatted Flink SQL statement.
//
// @param request - UpdateComputeJobDraftSqlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeJobDraftSqlResponse
func (client *Client) UpdateComputeJobDraftSqlWithOptions(request *UpdateComputeJobDraftSqlRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeJobDraftSqlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DraftSql) {
		query["DraftSql"] = request.DraftSql
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeJobDraftSql"),
		Version:     dara.String("2026-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeJobDraftSqlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Flink SQL draft saved for a specified compute job.
//
// Description:
//
// ## Request description
//
// This operation updates the Flink SQL draft saved for a specific compute job under a given compute instance. Make sure that the provided `InstanceId` and `JobName` are accurate, and the current status of the job supports SQL modification.
//
// - **Precautions**:
//
//   - Make sure that the target instance is in the running state.
//
//   - The current job status must allow SQL modification. The job must not be undergoing debugging or modification.
//
//   - The `DraftSql` parameter must contain a complete and correctly formatted Flink SQL statement.
//
// @param request - UpdateComputeJobDraftSqlRequest
//
// @return UpdateComputeJobDraftSqlResponse
func (client *Client) UpdateComputeJobDraftSql(request *UpdateComputeJobDraftSqlRequest) (_result *UpdateComputeJobDraftSqlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComputeJobDraftSqlResponse{}
	_body, _err := client.UpdateComputeJobDraftSqlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
