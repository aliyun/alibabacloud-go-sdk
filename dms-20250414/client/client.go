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
	client.Endpoint, _err = client.GetEndpoint(dara.String("dms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 批量新建湖仓表分区
//
// @param tmpReq - BatchCreateDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateDataLakePartitionsResponse
func (client *Client) BatchCreateDataLakePartitionsWithOptions(tmpReq *BatchCreateDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCreateDataLakePartitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInputs) {
		request.PartitionInputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInputs, dara.String("PartitionInputs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfNotExists) {
		query["IfNotExists"] = request.IfNotExists
	}

	if !dara.IsNil(request.NeedResult) {
		query["NeedResult"] = request.NeedResult
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputsShrink) {
		body["PartitionInputs"] = request.PartitionInputsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateDataLakePartitions"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateDataLakePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新建湖仓表分区
//
// @param request - BatchCreateDataLakePartitionsRequest
//
// @return BatchCreateDataLakePartitionsResponse
func (client *Client) BatchCreateDataLakePartitions(request *BatchCreateDataLakePartitionsRequest) (_result *BatchCreateDataLakePartitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateDataLakePartitionsResponse{}
	_body, _err := client.BatchCreateDataLakePartitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除湖仓表分区
//
// @param request - BatchDeleteDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDataLakePartitionsResponse
func (client *Client) BatchDeleteDataLakePartitionsWithOptions(request *BatchDeleteDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfExists) {
		query["IfExists"] = request.IfExists
	}

	if !dara.IsNil(request.PartitionValuesList) {
		query["PartitionValuesList"] = request.PartitionValuesList
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDataLakePartitions"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDataLakePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除湖仓表分区
//
// @param request - BatchDeleteDataLakePartitionsRequest
//
// @return BatchDeleteDataLakePartitionsResponse
func (client *Client) BatchDeleteDataLakePartitions(request *BatchDeleteDataLakePartitionsRequest) (_result *BatchDeleteDataLakePartitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteDataLakePartitionsResponse{}
	_body, _err := client.BatchDeleteDataLakePartitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量更新湖仓表分区
//
// @param tmpReq - BatchUpdateDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateDataLakePartitionsResponse
func (client *Client) BatchUpdateDataLakePartitionsWithOptions(tmpReq *BatchUpdateDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateDataLakePartitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInputs) {
		request.PartitionInputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInputs, dara.String("PartitionInputs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputsShrink) {
		body["PartitionInputs"] = request.PartitionInputsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateDataLakePartitions"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateDataLakePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新湖仓表分区
//
// @param request - BatchUpdateDataLakePartitionsRequest
//
// @return BatchUpdateDataLakePartitionsResponse
func (client *Client) BatchUpdateDataLakePartitions(request *BatchUpdateDataLakePartitionsRequest) (_result *BatchUpdateDataLakePartitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUpdateDataLakePartitionsResponse{}
	_body, _err := client.BatchUpdateDataLakePartitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Airflow
//
// @param request - CreateAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAirflowResponse
func (client *Client) CreateAirflowWithOptions(request *CreateAirflowRequest, runtime *dara.RuntimeOptions) (_result *CreateAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowName) {
		query["AirflowName"] = request.AirflowName
	}

	if !dara.IsNil(request.AppSpec) {
		query["AppSpec"] = request.AppSpec
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DagsDir) {
		query["DagsDir"] = request.DagsDir
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssPath) {
		query["OssPath"] = request.OssPath
	}

	if !dara.IsNil(request.PluginsDir) {
		query["PluginsDir"] = request.PluginsDir
	}

	if !dara.IsNil(request.RequirementFile) {
		query["RequirementFile"] = request.RequirementFile
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.StartupFile) {
		query["StartupFile"] = request.StartupFile
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkerServerlessReplicas) {
		query["WorkerServerlessReplicas"] = request.WorkerServerlessReplicas
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAirflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Airflow
//
// @param request - CreateAirflowRequest
//
// @return CreateAirflowResponse
func (client *Client) CreateAirflow(request *CreateAirflowRequest) (_result *CreateAirflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAirflowResponse{}
	_body, _err := client.CreateAirflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Airflow logon credential. You can use this credential to log on to the DMS-managed Airflow instance.
//
// @param request - CreateAirflowLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAirflowLoginTokenResponse
func (client *Client) CreateAirflowLoginTokenWithOptions(request *CreateAirflowLoginTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateAirflowLoginTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAirflowLoginToken"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAirflowLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Airflow logon credential. You can use this credential to log on to the DMS-managed Airflow instance.
//
// @param request - CreateAirflowLoginTokenRequest
//
// @return CreateAirflowLoginTokenResponse
func (client *Client) CreateAirflowLoginToken(request *CreateAirflowLoginTokenRequest) (_result *CreateAirflowLoginTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAirflowLoginTokenResponse{}
	_body, _err := client.CreateAirflowLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CreateDataAgentSession
//
// @param tmpReq - CreateDataAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataAgentSessionResponse
func (client *Client) CreateDataAgentSessionWithOptions(tmpReq *CreateDataAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SessionConfig) {
		request.SessionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionConfig, dara.String("SessionConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.File) {
		query["File"] = request.File
	}

	if !dara.IsNil(request.SessionConfigShrink) {
		query["SessionConfig"] = request.SessionConfigShrink
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataAgentSession"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataAgentSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateDataAgentSession
//
// @param request - CreateDataAgentSessionRequest
//
// @return CreateDataAgentSessionResponse
func (client *Client) CreateDataAgentSession(request *CreateDataAgentSessionRequest) (_result *CreateDataAgentSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataAgentSessionResponse{}
	_body, _err := client.CreateDataAgentSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建湖仓数据库
//
// @param tmpReq - CreateDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeDatabaseResponse
func (client *Client) CreateDataLakeDatabaseWithOptions(tmpReq *CreateDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeDatabaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建湖仓数据库
//
// @param request - CreateDataLakeDatabaseRequest
//
// @return CreateDataLakeDatabaseResponse
func (client *Client) CreateDataLakeDatabase(request *CreateDataLakeDatabaseRequest) (_result *CreateDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLakeDatabaseResponse{}
	_body, _err := client.CreateDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建湖仓自定义函数
//
// @param tmpReq - CreateDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeFunctionResponse
func (client *Client) CreateDataLakeFunctionWithOptions(tmpReq *CreateDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeFunctionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FunctionInput) {
		request.FunctionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FunctionInput, dara.String("FunctionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionInputShrink) {
		body["FunctionInput"] = request.FunctionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建湖仓自定义函数
//
// @param request - CreateDataLakeFunctionRequest
//
// @return CreateDataLakeFunctionResponse
func (client *Client) CreateDataLakeFunction(request *CreateDataLakeFunctionRequest) (_result *CreateDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLakeFunctionResponse{}
	_body, _err := client.CreateDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建湖仓表分区
//
// @param tmpReq - CreateDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakePartitionResponse
func (client *Client) CreateDataLakePartitionWithOptions(tmpReq *CreateDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInput) {
		request.PartitionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInput, dara.String("PartitionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfNotExists) {
		query["IfNotExists"] = request.IfNotExists
	}

	if !dara.IsNil(request.NeedResult) {
		query["NeedResult"] = request.NeedResult
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputShrink) {
		body["PartitionInput"] = request.PartitionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建湖仓表分区
//
// @param request - CreateDataLakePartitionRequest
//
// @return CreateDataLakePartitionResponse
func (client *Client) CreateDataLakePartition(request *CreateDataLakePartitionRequest) (_result *CreateDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLakePartitionResponse{}
	_body, _err := client.CreateDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建湖仓表
//
// @param tmpReq - CreateDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeTableResponse
func (client *Client) CreateDataLakeTableWithOptions(tmpReq *CreateDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableInput) {
		request.TableInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableInput, dara.String("TableInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableInputShrink) {
		body["TableInput"] = request.TableInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建湖仓表
//
// @param request - CreateDataLakeTableRequest
//
// @return CreateDataLakeTableResponse
func (client *Client) CreateDataLakeTable(request *CreateDataLakeTableRequest) (_result *CreateDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLakeTableResponse{}
	_body, _err := client.CreateDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Airflow
//
// @param request - DeleteAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAirflowResponse
func (client *Client) DeleteAirflowWithOptions(request *DeleteAirflowRequest, runtime *dara.RuntimeOptions) (_result *DeleteAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAirflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Airflow
//
// @param request - DeleteAirflowRequest
//
// @return DeleteAirflowResponse
func (client *Client) DeleteAirflow(request *DeleteAirflowRequest) (_result *DeleteAirflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAirflowResponse{}
	_body, _err := client.DeleteAirflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除湖仓数据库
//
// @param request - DeleteDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeDatabaseResponse
func (client *Client) DeleteDataLakeDatabaseWithOptions(request *DeleteDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除湖仓数据库
//
// @param request - DeleteDataLakeDatabaseRequest
//
// @return DeleteDataLakeDatabaseResponse
func (client *Client) DeleteDataLakeDatabase(request *DeleteDataLakeDatabaseRequest) (_result *DeleteDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLakeDatabaseResponse{}
	_body, _err := client.DeleteDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除湖仓自定义函数
//
// @param request - DeleteDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeFunctionResponse
func (client *Client) DeleteDataLakeFunctionWithOptions(request *DeleteDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除湖仓自定义函数
//
// @param request - DeleteDataLakeFunctionRequest
//
// @return DeleteDataLakeFunctionResponse
func (client *Client) DeleteDataLakeFunction(request *DeleteDataLakeFunctionRequest) (_result *DeleteDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLakeFunctionResponse{}
	_body, _err := client.DeleteDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除湖仓表分区
//
// @param tmpReq - DeleteDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakePartitionResponse
func (client *Client) DeleteDataLakePartitionWithOptions(tmpReq *DeleteDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionValues) {
		request.PartitionValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionValues, dara.String("PartitionValues"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfExists) {
		query["IfExists"] = request.IfExists
	}

	if !dara.IsNil(request.PartitionValuesShrink) {
		query["PartitionValues"] = request.PartitionValuesShrink
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除湖仓表分区
//
// @param request - DeleteDataLakePartitionRequest
//
// @return DeleteDataLakePartitionResponse
func (client *Client) DeleteDataLakePartition(request *DeleteDataLakePartitionRequest) (_result *DeleteDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLakePartitionResponse{}
	_body, _err := client.DeleteDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除湖仓表
//
// @param request - DeleteDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeTableResponse
func (client *Client) DeleteDataLakeTableWithOptions(request *DeleteDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除湖仓表
//
// @param request - DeleteDataLakeTableRequest
//
// @return DeleteDataLakeTableResponse
func (client *Client) DeleteDataLakeTable(request *DeleteDataLakeTableRequest) (_result *DeleteDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLakeTableResponse{}
	_body, _err := client.DeleteDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeDataAgentSession
//
// @param request - DescribeDataAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataAgentSessionResponse
func (client *Client) DescribeDataAgentSessionWithOptions(request *DescribeDataAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataAgentSession"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataAgentSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeDataAgentSession
//
// @param request - DescribeDataAgentSessionRequest
//
// @return DescribeDataAgentSessionResponse
func (client *Client) DescribeDataAgentSession(request *DescribeDataAgentSessionRequest) (_result *DescribeDataAgentSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataAgentSessionResponse{}
	_body, _err := client.DescribeDataAgentSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 Airflow
//
// @param request - GetAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAirflowResponse
func (client *Client) GetAirflowWithOptions(request *GetAirflowRequest, runtime *dara.RuntimeOptions) (_result *GetAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAirflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 Airflow
//
// @param request - GetAirflowRequest
//
// @return GetAirflowResponse
func (client *Client) GetAirflow(request *GetAirflowRequest) (_result *GetAirflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAirflowResponse{}
	_body, _err := client.GetAirflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetChatContent
//
// @param request - GetChatContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatContentResponse
func (client *Client) GetChatContentWithSSE(request *GetChatContentRequest, runtime *dara.RuntimeOptions, _yield chan *GetChatContentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.getChatContentWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// # GetChatContent
//
// @param request - GetChatContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatContentResponse
func (client *Client) GetChatContentWithOptions(request *GetChatContentRequest, runtime *dara.RuntimeOptions) (_result *GetChatContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.Checkpoint) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatContent"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetChatContent
//
// @param request - GetChatContentRequest
//
// @return GetChatContentResponse
func (client *Client) GetChatContent(request *GetChatContentRequest) (_result *GetChatContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChatContentResponse{}
	_body, _err := client.GetChatContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取uc的数据库目录
//
// @param request - GetDataLakeCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeCatalogResponse
func (client *Client) GetDataLakeCatalogWithOptions(request *GetDataLakeCatalogRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeCatalog"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取uc的数据库目录
//
// @param request - GetDataLakeCatalogRequest
//
// @return GetDataLakeCatalogResponse
func (client *Client) GetDataLakeCatalog(request *GetDataLakeCatalogRequest) (_result *GetDataLakeCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakeCatalogResponse{}
	_body, _err := client.GetDataLakeCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取UC的数据库
//
// @param request - GetDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeDatabaseResponse
func (client *Client) GetDataLakeDatabaseWithOptions(request *GetDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取UC的数据库
//
// @param request - GetDataLakeDatabaseRequest
//
// @return GetDataLakeDatabaseResponse
func (client *Client) GetDataLakeDatabase(request *GetDataLakeDatabaseRequest) (_result *GetDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakeDatabaseResponse{}
	_body, _err := client.GetDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取湖仓自定义函数详细信息
//
// @param request - GetDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeFunctionResponse
func (client *Client) GetDataLakeFunctionWithOptions(request *GetDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取湖仓自定义函数详细信息
//
// @param request - GetDataLakeFunctionRequest
//
// @return GetDataLakeFunctionResponse
func (client *Client) GetDataLakeFunction(request *GetDataLakeFunctionRequest) (_result *GetDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakeFunctionResponse{}
	_body, _err := client.GetDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取湖仓表分区详情
//
// @param tmpReq - GetDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakePartitionResponse
func (client *Client) GetDataLakePartitionWithOptions(tmpReq *GetDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionValues) {
		request.PartitionValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionValues, dara.String("PartitionValues"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.PartitionValuesShrink) {
		query["PartitionValues"] = request.PartitionValuesShrink
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取湖仓表分区详情
//
// @param request - GetDataLakePartitionRequest
//
// @return GetDataLakePartitionResponse
func (client *Client) GetDataLakePartition(request *GetDataLakePartitionRequest) (_result *GetDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakePartitionResponse{}
	_body, _err := client.GetDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取表信息
//
// @param request - GetDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeTableResponse
func (client *Client) GetDataLakeTableWithOptions(request *GetDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表信息
//
// @param request - GetDataLakeTableRequest
//
// @return GetDataLakeTableResponse
func (client *Client) GetDataLakeTable(request *GetDataLakeTableRequest) (_result *GetDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakeTableResponse{}
	_body, _err := client.GetDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调度运行Notebook文件
//
// @param request - GetNotebookAndSubmitTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotebookAndSubmitTaskResponse
func (client *Client) GetNotebookAndSubmitTaskWithOptions(request *GetNotebookAndSubmitTaskRequest, runtime *dara.RuntimeOptions) (_result *GetNotebookAndSubmitTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.Retry) {
		body["Retry"] = request.Retry
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotebookAndSubmitTask"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotebookAndSubmitTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调度运行Notebook文件
//
// @param request - GetNotebookAndSubmitTaskRequest
//
// @return GetNotebookAndSubmitTaskResponse
func (client *Client) GetNotebookAndSubmitTask(request *GetNotebookAndSubmitTaskRequest) (_result *GetNotebookAndSubmitTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNotebookAndSubmitTaskResponse{}
	_body, _err := client.GetNotebookAndSubmitTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看Notebook任务运行结果
//
// @param request - GetNotebookTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotebookTaskStatusResponse
func (client *Client) GetNotebookTaskStatusWithOptions(request *GetNotebookTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetNotebookTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotebookTaskStatus"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotebookTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看Notebook任务运行结果
//
// @param request - GetNotebookTaskStatusRequest
//
// @return GetNotebookTaskStatusResponse
func (client *Client) GetNotebookTaskStatus(request *GetNotebookTaskStatusRequest) (_result *GetNotebookTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNotebookTaskStatusResponse{}
	_body, _err := client.GetNotebookTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出资源Airflow
//
// @param request - ListAirflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAirflowsResponse
func (client *Client) ListAirflowsWithOptions(request *ListAirflowsRequest, runtime *dara.RuntimeOptions) (_result *ListAirflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAirflows"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAirflowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出资源Airflow
//
// @param request - ListAirflowsRequest
//
// @return ListAirflowsResponse
func (client *Client) ListAirflows(request *ListAirflowsRequest) (_result *ListAirflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAirflowsResponse{}
	_body, _err := client.ListAirflowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取uc的数据库目录列表
//
// @param request - ListDataLakeCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeCatalogResponse
func (client *Client) ListDataLakeCatalogWithOptions(request *ListDataLakeCatalogRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeCatalog"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取uc的数据库目录列表
//
// @param request - ListDataLakeCatalogRequest
//
// @return ListDataLakeCatalogResponse
func (client *Client) ListDataLakeCatalog(request *ListDataLakeCatalogRequest) (_result *ListDataLakeCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeCatalogResponse{}
	_body, _err := client.ListDataLakeCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据库列表
//
// @param request - ListDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeDatabaseResponse
func (client *Client) ListDataLakeDatabaseWithOptions(request *ListDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据库列表
//
// @param request - ListDataLakeDatabaseRequest
//
// @return ListDataLakeDatabaseResponse
func (client *Client) ListDataLakeDatabase(request *ListDataLakeDatabaseRequest) (_result *ListDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeDatabaseResponse{}
	_body, _err := client.ListDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据湖函数列表
//
// @param request - ListDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeFunctionResponse
func (client *Client) ListDataLakeFunctionWithOptions(request *ListDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionNamePattern) {
		query["FunctionNamePattern"] = request.FunctionNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据湖函数列表
//
// @param request - ListDataLakeFunctionRequest
//
// @return ListDataLakeFunctionResponse
func (client *Client) ListDataLakeFunction(request *ListDataLakeFunctionRequest) (_result *ListDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeFunctionResponse{}
	_body, _err := client.ListDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据湖函数名列表
//
// @param request - ListDataLakeFunctionNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeFunctionNameResponse
func (client *Client) ListDataLakeFunctionNameWithOptions(request *ListDataLakeFunctionNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeFunctionNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionNamePattern) {
		query["FunctionNamePattern"] = request.FunctionNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeFunctionName"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeFunctionNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据湖函数名列表
//
// @param request - ListDataLakeFunctionNameRequest
//
// @return ListDataLakeFunctionNameResponse
func (client *Client) ListDataLakeFunctionName(request *ListDataLakeFunctionNameRequest) (_result *ListDataLakeFunctionNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeFunctionNameResponse{}
	_body, _err := client.ListDataLakeFunctionNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据湖表分区列表
//
// @param tmpReq - ListDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionResponse
func (client *Client) ListDataLakePartitionWithOptions(tmpReq *ListDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartNames) {
		request.PartNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartNames, dara.String("PartNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartNamesShrink) {
		body["PartNames"] = request.PartNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据湖表分区列表
//
// @param request - ListDataLakePartitionRequest
//
// @return ListDataLakePartitionResponse
func (client *Client) ListDataLakePartition(request *ListDataLakePartitionRequest) (_result *ListDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakePartitionResponse{}
	_body, _err := client.ListDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据筛选条件获取数据湖表分区列表
//
// @param request - ListDataLakePartitionByFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionByFilterResponse
func (client *Client) ListDataLakePartitionByFilterWithOptions(request *ListDataLakePartitionByFilterRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionByFilterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartitionByFilter"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionByFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据筛选条件获取数据湖表分区列表
//
// @param request - ListDataLakePartitionByFilterRequest
//
// @return ListDataLakePartitionByFilterResponse
func (client *Client) ListDataLakePartitionByFilter(request *ListDataLakePartitionByFilterRequest) (_result *ListDataLakePartitionByFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakePartitionByFilterResponse{}
	_body, _err := client.ListDataLakePartitionByFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据湖表分区名列表
//
// @param request - ListDataLakePartitionNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionNameResponse
func (client *Client) ListDataLakePartitionNameWithOptions(request *ListDataLakePartitionNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartitionName"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据湖表分区名列表
//
// @param request - ListDataLakePartitionNameRequest
//
// @return ListDataLakePartitionNameResponse
func (client *Client) ListDataLakePartitionName(request *ListDataLakePartitionNameRequest) (_result *ListDataLakePartitionNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakePartitionNameResponse{}
	_body, _err := client.ListDataLakePartitionNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据湖表列表
//
// @param request - ListDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTableResponse
func (client *Client) ListDataLakeTableWithOptions(request *ListDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["TableNamePattern"] = request.TableNamePattern
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据湖表列表
//
// @param request - ListDataLakeTableRequest
//
// @return ListDataLakeTableResponse
func (client *Client) ListDataLakeTable(request *ListDataLakeTableRequest) (_result *ListDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeTableResponse{}
	_body, _err := client.ListDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据湖表名列表
//
// @param request - ListDataLakeTableNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTableNameResponse
func (client *Client) ListDataLakeTableNameWithOptions(request *ListDataLakeTableNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTableNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["TableNamePattern"] = request.TableNamePattern
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTableName"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTableNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据湖表名列表
//
// @param request - ListDataLakeTableNameRequest
//
// @return ListDataLakeTableNameResponse
func (client *Client) ListDataLakeTableName(request *ListDataLakeTableNameRequest) (_result *ListDataLakeTableNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeTableNameResponse{}
	_body, _err := client.ListDataLakeTableNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取表信息
//
// @param request - ListDataLakeTablebaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTablebaseInfoResponse
func (client *Client) ListDataLakeTablebaseInfoWithOptions(request *ListDataLakeTablebaseInfoRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTablebaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.Rows) {
		query["Rows"] = request.Rows
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTablebaseInfo"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTablebaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表信息
//
// @param request - ListDataLakeTablebaseInfoRequest
//
// @return ListDataLakeTablebaseInfoResponse
func (client *Client) ListDataLakeTablebaseInfo(request *ListDataLakeTablebaseInfoRequest) (_result *ListDataLakeTablebaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeTablebaseInfoResponse{}
	_body, _err := client.ListDataLakeTablebaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SendChatMessage
//
// @param tmpReq - SendChatMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendChatMessageResponse
func (client *Client) SendChatMessageWithOptions(tmpReq *SendChatMessageRequest, runtime *dara.RuntimeOptions) (_result *SendChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendChatMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSource) {
		request.DataSourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSource, dara.String("DataSource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SessionConfig) {
		request.SessionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionConfig, dara.String("SessionConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.DataSourceShrink) {
		query["DataSource"] = request.DataSourceShrink
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.Question) {
		query["Question"] = request.Question
	}

	if !dara.IsNil(request.QuotedMessage) {
		query["QuotedMessage"] = request.QuotedMessage
	}

	if !dara.IsNil(request.ReplyTo) {
		query["ReplyTo"] = request.ReplyTo
	}

	if !dara.IsNil(request.SessionConfigShrink) {
		query["SessionConfig"] = request.SessionConfigShrink
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendChatMessage"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendChatMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SendChatMessage
//
// @param request - SendChatMessageRequest
//
// @return SendChatMessageResponse
func (client *Client) SendChatMessage(request *SendChatMessageRequest) (_result *SendChatMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendChatMessageResponse{}
	_body, _err := client.SendChatMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新UpdateAirflow
//
// @param request - UpdateAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAirflowResponse
func (client *Client) UpdateAirflowWithOptions(request *UpdateAirflowRequest, runtime *dara.RuntimeOptions) (_result *UpdateAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	if !dara.IsNil(request.AirflowName) {
		query["AirflowName"] = request.AirflowName
	}

	if !dara.IsNil(request.AppSpec) {
		query["AppSpec"] = request.AppSpec
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DagsDir) {
		query["DagsDir"] = request.DagsDir
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.PluginsDir) {
		query["PluginsDir"] = request.PluginsDir
	}

	if !dara.IsNil(request.RequirementFile) {
		query["RequirementFile"] = request.RequirementFile
	}

	if !dara.IsNil(request.StartupFile) {
		query["StartupFile"] = request.StartupFile
	}

	if !dara.IsNil(request.WorkerServerlessReplicas) {
		query["WorkerServerlessReplicas"] = request.WorkerServerlessReplicas
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAirflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新UpdateAirflow
//
// @param request - UpdateAirflowRequest
//
// @return UpdateAirflowResponse
func (client *Client) UpdateAirflow(request *UpdateAirflowRequest) (_result *UpdateAirflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAirflowResponse{}
	_body, _err := client.UpdateAirflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新湖仓数据库
//
// @param tmpReq - UpdateDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeDatabaseResponse
func (client *Client) UpdateDataLakeDatabaseWithOptions(tmpReq *UpdateDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeDatabaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新湖仓数据库
//
// @param request - UpdateDataLakeDatabaseRequest
//
// @return UpdateDataLakeDatabaseResponse
func (client *Client) UpdateDataLakeDatabase(request *UpdateDataLakeDatabaseRequest) (_result *UpdateDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLakeDatabaseResponse{}
	_body, _err := client.UpdateDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新湖仓自定义函数
//
// @param tmpReq - UpdateDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeFunctionResponse
func (client *Client) UpdateDataLakeFunctionWithOptions(tmpReq *UpdateDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeFunctionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FunctionInput) {
		request.FunctionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FunctionInput, dara.String("FunctionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionInputShrink) {
		body["FunctionInput"] = request.FunctionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新湖仓自定义函数
//
// @param request - UpdateDataLakeFunctionRequest
//
// @return UpdateDataLakeFunctionResponse
func (client *Client) UpdateDataLakeFunction(request *UpdateDataLakeFunctionRequest) (_result *UpdateDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLakeFunctionResponse{}
	_body, _err := client.UpdateDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新湖仓表分区
//
// @param tmpReq - UpdateDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakePartitionResponse
func (client *Client) UpdateDataLakePartitionWithOptions(tmpReq *UpdateDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInput) {
		request.PartitionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInput, dara.String("PartitionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputShrink) {
		body["PartitionInput"] = request.PartitionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新湖仓表分区
//
// @param request - UpdateDataLakePartitionRequest
//
// @return UpdateDataLakePartitionResponse
func (client *Client) UpdateDataLakePartition(request *UpdateDataLakePartitionRequest) (_result *UpdateDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLakePartitionResponse{}
	_body, _err := client.UpdateDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新湖仓表信息
//
// @param tmpReq - UpdateDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeTableResponse
func (client *Client) UpdateDataLakeTableWithOptions(tmpReq *UpdateDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableInput) {
		request.TableInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableInput, dara.String("TableInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableInputShrink) {
		body["TableInput"] = request.TableInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新湖仓表信息
//
// @param request - UpdateDataLakeTableRequest
//
// @return UpdateDataLakeTableResponse
func (client *Client) UpdateDataLakeTable(request *UpdateDataLakeTableRequest) (_result *UpdateDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLakeTableResponse{}
	_body, _err := client.UpdateDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) getChatContentWithSSE_opYieldFunc(_yield chan *GetChatContentResponse, _yieldErr chan error, request *GetChatContentRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.Checkpoint) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatContent"),
		Version:     dara.String("2025-04-14"),
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
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
