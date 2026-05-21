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
	client.Endpoint, _err = client.GetEndpoint(dara.String("hologram"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 绑定主实例
//
// @param request - BindLeaderInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindLeaderInstanceResponse
func (client *Client) BindLeaderInstanceWithOptions(instanceId *string, request *BindLeaderInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BindLeaderInstanceResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.LeaderInstanceId) {
		body["leaderInstanceId"] = request.LeaderInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindLeaderInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/bindReadOnly"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BindLeaderInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定主实例
//
// @param request - BindLeaderInstanceRequest
//
// @return BindLeaderInstanceResponse
func (client *Client) BindLeaderInstance(instanceId *string, request *BindLeaderInstanceRequest) (_result *BindLeaderInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindLeaderInstanceResponse{}
	_body, _err := client.BindLeaderInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		body["newResourceGroupId"] = request.NewResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tag/changeResourceGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建db
//
// @param request - CreateDatabaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabaseWithOptions(instanceId *string, request *CreateDatabaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		body["databaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.PermissionModel) {
		body["permissionModel"] = request.PermissionModel
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabase"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/createDatabase"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建db
//
// @param request - CreateDatabaseRequest
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabase(instanceId *string, request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建外部db
//
// @param request - CreateExternalDatabaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExternalDatabaseResponse
func (client *Client) CreateExternalDatabaseWithOptions(instanceId *string, request *CreateExternalDatabaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExternalDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.DatabaseName) {
		body["databaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DefaultUserMapping) {
		body["defaultUserMapping"] = request.DefaultUserMapping
	}

	if !dara.IsNil(request.ExternalConfig) {
		body["externalConfig"] = request.ExternalConfig
	}

	if !dara.IsNil(request.MetastoreType) {
		body["metastoreType"] = request.MetastoreType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExternalDatabase"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/createExternalDatabase"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExternalDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建外部db
//
// @param request - CreateExternalDatabaseRequest
//
// @return CreateExternalDatabaseResponse
func (client *Client) CreateExternalDatabase(instanceId *string, request *CreateExternalDatabaseRequest) (_result *CreateExternalDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExternalDatabaseResponse{}
	_body, _err := client.CreateExternalDatabaseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a virtual warehouse.
//
// @param request - CreateHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHoloWarehouseResponse
func (client *Client) CreateHoloWarehouseWithOptions(instanceId *string, request *CreateHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHoloWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterCount) {
		body["clusterCount"] = request.ClusterCount
	}

	if !dara.IsNil(request.Cpu) {
		body["cpu"] = request.Cpu
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHoloWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/createHoloWarehouse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHoloWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual warehouse.
//
// @param request - CreateHoloWarehouseRequest
//
// @return CreateHoloWarehouseResponse
func (client *Client) CreateHoloWarehouse(instanceId *string, request *CreateHoloWarehouseRequest) (_result *CreateHoloWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHoloWarehouseResponse{}
	_body, _err := client.CreateHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Hologres instance.
//
// Description:
//
// > Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
//   - For more information about the billing details of Hologres, see [Pricing](https://www.alibabacloud.com/help/en/hologres/developer-reference/api-hologram-2022-06-01-createinstance).
//
//   - When you purchase a Hologres instance, you must specify the region and zone in which the Hologres instance resides. A region may correspond to multiple zones. Example:
//
// <!---->
//
//	cn-hangzhou: cn-hangzhou-h, cn-hangzhou-j
//
//	   cn-shanghai: cn-shanghai-e, cn-shanghai-f
//
//	   cn-beijing: cn-beijing-i, cn-beijing-g
//
//	   cn-zhangjiakou: cn-zhangjiakou-b
//
//	   cn-shenzhen: cn-shenzhen-e
//
//	   cn-hongkong: cn-hongkong-b
//
//	   cn-shanghai-finance-1: cn-shanghai-finance-1z
//
//	   ap-northeast-1: ap-northeast-1a
//
//	   ap-southeast-1: ap-southeast-1c
//
//	   ap-southeast-3: ap-southeast-3b
//
//	   ap-southeast-5: ap-southeast-5b
//
//	   ap-south-1: ap-south-1b
//
//	   eu-central-1: eu-central-1a
//
//	   us-east-1: us-east-1a
//
//	   us-west-1: us-west-1b
//
// @param request - CreateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		body["autoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		body["autoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ChargeType) {
		body["chargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ColdStorageSize) {
		body["coldStorageSize"] = request.ColdStorageSize
	}

	if !dara.IsNil(request.Cpu) {
		body["cpu"] = request.Cpu
	}

	if !dara.IsNil(request.Duration) {
		body["duration"] = request.Duration
	}

	if !dara.IsNil(request.EnableServerlessComputing) {
		body["enableServerlessComputing"] = request.EnableServerlessComputing
	}

	if !dara.IsNil(request.GatewayCount) {
		body["gatewayCount"] = request.GatewayCount
	}

	if !dara.IsNil(request.InitialDatabases) {
		body["initialDatabases"] = request.InitialDatabases
	}

	if !dara.IsNil(request.InstanceName) {
		body["instanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		body["instanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.LeaderInstanceId) {
		body["leaderInstanceId"] = request.LeaderInstanceId
	}

	if !dara.IsNil(request.PricingCycle) {
		body["pricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.RegionId) {
		body["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StorageSize) {
		body["storageSize"] = request.StorageSize
	}

	if !dara.IsNil(request.StorageType) {
		body["storageType"] = request.StorageType
	}

	if !dara.IsNil(request.VSwitchId) {
		body["vSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		body["vpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		body["zoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/create"),
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
// Creates a Hologres instance.
//
// Description:
//
// > Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
//   - For more information about the billing details of Hologres, see [Pricing](https://www.alibabacloud.com/help/en/hologres/developer-reference/api-hologram-2022-06-01-createinstance).
//
//   - When you purchase a Hologres instance, you must specify the region and zone in which the Hologres instance resides. A region may correspond to multiple zones. Example:
//
// <!---->
//
//	cn-hangzhou: cn-hangzhou-h, cn-hangzhou-j
//
//	   cn-shanghai: cn-shanghai-e, cn-shanghai-f
//
//	   cn-beijing: cn-beijing-i, cn-beijing-g
//
//	   cn-zhangjiakou: cn-zhangjiakou-b
//
//	   cn-shenzhen: cn-shenzhen-e
//
//	   cn-hongkong: cn-hongkong-b
//
//	   cn-shanghai-finance-1: cn-shanghai-finance-1z
//
//	   ap-northeast-1: ap-northeast-1a
//
//	   ap-southeast-1: ap-southeast-1c
//
//	   ap-southeast-3: ap-southeast-3b
//
//	   ap-southeast-5: ap-southeast-5b
//
//	   ap-south-1: ap-south-1b
//
//	   eu-central-1: eu-central-1a
//
//	   us-east-1: us-east-1a
//
//	   us-west-1: us-west-1b
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
// 创建手动备份
//
// @param request - CreateManualBackupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateManualBackupResponse
func (client *Client) CreateManualBackupWithOptions(request *CreateManualBackupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateManualBackupResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateManualBackup"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/backups/manual"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateManualBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建手动备份
//
// @param request - CreateManualBackupRequest
//
// @return CreateManualBackupResponse
func (client *Client) CreateManualBackup(request *CreateManualBackupRequest) (_result *CreateManualBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateManualBackupResponse{}
	_body, _err := client.CreateManualBackupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模型服务
//
// @param request - CreateModelServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelServiceResponse
func (client *Client) CreateModelServiceWithOptions(instanceId *string, request *CreateModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		body["apiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.Cpu) {
		body["cpu"] = request.Cpu
	}

	if !dara.IsNil(request.Gpu) {
		body["gpu"] = request.Gpu
	}

	if !dara.IsNil(request.GpuMemory) {
		body["gpuMemory"] = request.GpuMemory
	}

	if !dara.IsNil(request.Memory) {
		body["memory"] = request.Memory
	}

	if !dara.IsNil(request.ModelParams) {
		body["modelParams"] = request.ModelParams
	}

	if !dara.IsNil(request.ModelServiceName) {
		body["modelServiceName"] = request.ModelServiceName
	}

	if !dara.IsNil(request.ModelType) {
		body["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.Provider) {
		body["provider"] = request.Provider
	}

	if !dara.IsNil(request.ServiceCount) {
		body["serviceCount"] = request.ServiceCount
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelService"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/createModelService"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型服务
//
// @param request - CreateModelServiceRequest
//
// @return CreateModelServiceResponse
func (client *Client) CreateModelService(instanceId *string, request *CreateModelServiceRequest) (_result *CreateModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelServiceResponse{}
	_body, _err := client.CreateModelServiceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加用户
//
// @param request - CreateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(instanceId *string, request *CreateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SuperUser) {
		body["superUser"] = request.SuperUser
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/createUser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

// Summary:
//
// 添加用户
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(instanceId *string, request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建warehouse
//
// @param request - CreateWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarehouseResponse
func (client *Client) CreateWarehouseWithOptions(instanceId *string, request *CreateWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWarehouseResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Cpu) {
		body["cpu"] = request.Cpu
	}

	if !dara.IsNil(request.WarehouseName) {
		body["warehouseName"] = request.WarehouseName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/warehouse/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建warehouse
//
// @param request - CreateWarehouseRequest
//
// @return CreateWarehouseResponse
func (client *Client) CreateWarehouse(instanceId *string, request *CreateWarehouseRequest) (_result *CreateWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWarehouseResponse{}
	_body, _err := client.CreateWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建分时弹性计划
//
// @param request - CreateWarehouseScheduleTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWarehouseScheduleTaskResponse
func (client *Client) CreateWarehouseScheduleTaskWithOptions(instanceId *string, request *CreateWarehouseScheduleTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWarehouseScheduleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ElasticCu) {
		body["elasticCu"] = request.ElasticCu
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.WarehouseId) {
		body["warehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWarehouseScheduleTask"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/createWarehouseScheduleTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWarehouseScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建分时弹性计划
//
// @param request - CreateWarehouseScheduleTaskRequest
//
// @return CreateWarehouseScheduleTaskResponse
func (client *Client) CreateWarehouseScheduleTask(instanceId *string, request *CreateWarehouseScheduleTaskRequest) (_result *CreateWarehouseScheduleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWarehouseScheduleTaskResponse{}
	_body, _err := client.CreateWarehouseScheduleTaskWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除备份
//
// @param request - DeleteBackupDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupDataResponse
func (client *Client) DeleteBackupDataWithOptions(id *string, request *DeleteBackupDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBackupDataResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackupData"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/backups/" + dara.PercentEncode(dara.StringValue(id)) + "/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除备份
//
// @param request - DeleteBackupDataRequest
//
// @return DeleteBackupDataResponse
func (client *Client) DeleteBackupData(id *string, request *DeleteBackupDataRequest) (_result *DeleteBackupDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBackupDataResponse{}
	_body, _err := client.DeleteBackupDataWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virtual warehouse.
//
// @param request - DeleteHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHoloWarehouseResponse
func (client *Client) DeleteHoloWarehouseWithOptions(instanceId *string, request *DeleteHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHoloWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHoloWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/deleteHoloWarehouse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHoloWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual warehouse.
//
// @param request - DeleteHoloWarehouseRequest
//
// @return DeleteHoloWarehouseResponse
func (client *Client) DeleteHoloWarehouse(instanceId *string, request *DeleteHoloWarehouseRequest) (_result *DeleteHoloWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHoloWarehouseResponse{}
	_body, _err := client.DeleteHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Hologres instance.
//
// Description:
//
// > Before you call this operation, read the documentation and make sure that you understand the prerequisites and impacts of this operation.
//
//   - After you delete a Hologres instance, data and objects in the instance cannot be restored. Proceed with caution. For more information, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview?spm=a2c63.p38356.0.0.efc33b87i5pDl7).
//
//   - You can delete only pay-as-you-go instances.
//
// @param request - DeleteInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(instanceId *string, request *DeleteInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/delete"),
		Method:      dara.String("POST"),
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
// Deletes a Hologres instance.
//
// Description:
//
// > Before you call this operation, read the documentation and make sure that you understand the prerequisites and impacts of this operation.
//
//   - After you delete a Hologres instance, data and objects in the instance cannot be restored. Proceed with caution. For more information, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview?spm=a2c63.p38356.0.0.efc33b87i5pDl7).
//
//   - You can delete only pay-as-you-go instances.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(instanceId *string, request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除模型资源
//
// @param request - DeleteModelResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResourceResponse
func (client *Client) DeleteModelResourceWithOptions(instanceId *string, request *DeleteModelResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AiInstanceId) {
		query["aiInstanceId"] = request.AiInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelResource"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/deleteModelResource"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模型资源
//
// @param request - DeleteModelResourceRequest
//
// @return DeleteModelResourceResponse
func (client *Client) DeleteModelResource(instanceId *string, request *DeleteModelResourceRequest) (_result *DeleteModelResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelResourceResponse{}
	_body, _err := client.DeleteModelResourceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模型服务
//
// @param request - DeleteModelServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelServiceWithOptions(instanceId *string, request *DeleteModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelServiceName) {
		query["modelServiceName"] = request.ModelServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelService"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/deleteModelService"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型服务
//
// @param request - DeleteModelServiceRequest
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelService(instanceId *string, request *DeleteModelServiceRequest) (_result *DeleteModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelServiceResponse{}
	_body, _err := client.DeleteModelServiceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除计算组弹性计划
//
// @param request - DeleteWarehouseScheduleTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWarehouseScheduleTaskResponse
func (client *Client) DeleteWarehouseScheduleTaskWithOptions(instanceId *string, request *DeleteWarehouseScheduleTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWarehouseScheduleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.WarehouseId) {
		body["warehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWarehouseScheduleTask"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/deleteWarehouseScheduleTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWarehouseScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除计算组弹性计划
//
// @param request - DeleteWarehouseScheduleTaskRequest
//
// @return DeleteWarehouseScheduleTaskResponse
func (client *Client) DeleteWarehouseScheduleTask(instanceId *string, request *DeleteWarehouseScheduleTaskRequest) (_result *DeleteWarehouseScheduleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWarehouseScheduleTaskResponse{}
	_body, _err := client.DeleteWarehouseScheduleTaskWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂停实例
//
// @param request - DescribeTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
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

	if !dara.IsNil(request.TagsOnly) {
		query["tagsOnly"] = request.TagsOnly
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tag"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停实例
//
// @param request - DescribeTagsRequest
//
// @return DescribeTagsResponse
func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭OpenAPI执行SQL功能
//
// @param request - DisableExecuteStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableExecuteStatementResponse
func (client *Client) DisableExecuteStatementWithOptions(instanceId *string, request *DisableExecuteStatementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableExecuteStatementResponse, _err error) {
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
		Action:      dara.String("DisableExecuteStatement"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/disableExecuteStatement"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableExecuteStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭OpenAPI执行SQL功能
//
// @param request - DisableExecuteStatementRequest
//
// @return DisableExecuteStatementResponse
func (client *Client) DisableExecuteStatement(instanceId *string, request *DisableExecuteStatementRequest) (_result *DisableExecuteStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableExecuteStatementResponse{}
	_body, _err := client.DisableExecuteStatementWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables data lake acceleration.
//
// @param request - DisableHiveAccessRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableHiveAccessResponse
func (client *Client) DisableHiveAccessWithOptions(instanceId *string, request *DisableHiveAccessRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableHiveAccessResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableHiveAccess"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/disableHiveAccess"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableHiveAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables data lake acceleration.
//
// @param request - DisableHiveAccessRequest
//
// @return DisableHiveAccessResponse
func (client *Client) DisableHiveAccess(instanceId *string, request *DisableHiveAccessRequest) (_result *DisableHiveAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableHiveAccessResponse{}
	_body, _err := client.DisableHiveAccessWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消执行计划
//
// @param request - DisableOperationEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableOperationEventResponse
func (client *Client) DisableOperationEventWithOptions(instanceId *string, request *DisableOperationEventRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableOperationEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableOperationEvent"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/disableOperationEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableOperationEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消执行计划
//
// @param request - DisableOperationEventRequest
//
// @return DisableOperationEventResponse
func (client *Client) DisableOperationEvent(instanceId *string, request *DisableOperationEventRequest) (_result *DisableOperationEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableOperationEventResponse{}
	_body, _err := client.DisableOperationEventWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭SSL
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSSLResponse
func (client *Client) DisableSSLWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableSSLResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSSL"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/disableSSL"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭SSL
//
// @return DisableSSLResponse
func (client *Client) DisableSSL(instanceId *string) (_result *DisableSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableSSLResponse{}
	_body, _err := client.DisableSSLWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭服务账号
//
// @param request - DisableSupportAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSupportAccountResponse
func (client *Client) DisableSupportAccountWithOptions(instanceId *string, request *DisableSupportAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableSupportAccountResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSupportAccount"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/supportAccount/disable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSupportAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭服务账号
//
// @param request - DisableSupportAccountRequest
//
// @return DisableSupportAccountResponse
func (client *Client) DisableSupportAccount(instanceId *string, request *DisableSupportAccountRequest) (_result *DisableSupportAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableSupportAccountResponse{}
	_body, _err := client.DisableSupportAccountWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭自动弹性
//
// @param request - DisableWarehouseAutoScaleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableWarehouseAutoScaleResponse
func (client *Client) DisableWarehouseAutoScaleWithOptions(instanceId *string, request *DisableWarehouseAutoScaleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableWarehouseAutoScaleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableWarehouseAutoScale"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/disableWarehouseAutoScale"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableWarehouseAutoScaleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭自动弹性
//
// @param request - DisableWarehouseAutoScaleRequest
//
// @return DisableWarehouseAutoScaleResponse
func (client *Client) DisableWarehouseAutoScale(instanceId *string, request *DisableWarehouseAutoScaleRequest) (_result *DisableWarehouseAutoScaleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableWarehouseAutoScaleResponse{}
	_body, _err := client.DisableWarehouseAutoScaleWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消升级
//
// @param request - DiscardUpgradeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DiscardUpgradeResponse
func (client *Client) DiscardUpgradeWithOptions(instanceId *string, request *DiscardUpgradeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DiscardUpgradeResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DiscardUpgrade"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/upgrade/discard"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DiscardUpgradeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消升级
//
// @param request - DiscardUpgradeRequest
//
// @return DiscardUpgradeResponse
func (client *Client) DiscardUpgrade(instanceId *string, request *DiscardUpgradeRequest) (_result *DiscardUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DiscardUpgradeResponse{}
	_body, _err := client.DiscardUpgradeWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户
//
// @param request - DropUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropUserResponse
func (client *Client) DropUserWithOptions(instanceId *string, request *DropUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SuperUser) {
		body["superUser"] = request.SuperUser
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DropUser"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/dropUser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DropUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户
//
// @param request - DropUserRequest
//
// @return DropUserResponse
func (client *Client) DropUser(instanceId *string, request *DropUserRequest) (_result *DropUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DropUserResponse{}
	_body, _err := client.DropUserWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启或关闭OpenAPI执行SQL功能
//
// @param request - EnableExecuteStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableExecuteStatementResponse
func (client *Client) EnableExecuteStatementWithOptions(instanceId *string, request *EnableExecuteStatementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableExecuteStatementResponse, _err error) {
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
		Action:      dara.String("EnableExecuteStatement"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/enableExecuteStatement"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableExecuteStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启或关闭OpenAPI执行SQL功能
//
// @param request - EnableExecuteStatementRequest
//
// @return EnableExecuteStatementResponse
func (client *Client) EnableExecuteStatement(instanceId *string, request *EnableExecuteStatementRequest) (_result *EnableExecuteStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableExecuteStatementResponse{}
	_body, _err := client.EnableExecuteStatementWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables data lake acceleration.
//
// @param request - EnableHiveAccessRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHiveAccessResponse
func (client *Client) EnableHiveAccessWithOptions(instanceId *string, request *EnableHiveAccessRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableHiveAccessResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableHiveAccess"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/enableHiveAccess"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableHiveAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables data lake acceleration.
//
// @param request - EnableHiveAccessRequest
//
// @return EnableHiveAccessResponse
func (client *Client) EnableHiveAccess(instanceId *string, request *EnableHiveAccessRequest) (_result *EnableHiveAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableHiveAccessResponse{}
	_body, _err := client.EnableHiveAccessWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 打开SSL
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSSLResponse
func (client *Client) EnableSSLWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableSSLResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSSL"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/enableSSL"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 打开SSL
//
// @return EnableSSLResponse
func (client *Client) EnableSSL(instanceId *string) (_result *EnableSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableSSLResponse{}
	_body, _err := client.EnableSSLWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 打开服务账号
//
// @param request - EnableSupportAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSupportAccountResponse
func (client *Client) EnableSupportAccountWithOptions(instanceId *string, request *EnableSupportAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableSupportAccountResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ExpireTime) {
		body["expireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSupportAccount"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/supportAccount/enable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSupportAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 打开服务账号
//
// @param request - EnableSupportAccountRequest
//
// @return EnableSupportAccountResponse
func (client *Client) EnableSupportAccount(instanceId *string, request *EnableSupportAccountRequest) (_result *EnableSupportAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableSupportAccountResponse{}
	_body, _err := client.EnableSupportAccountWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启自动弹性
//
// @param request - EnableWarehouseAutoScaleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableWarehouseAutoScaleResponse
func (client *Client) EnableWarehouseAutoScaleWithOptions(instanceId *string, request *EnableWarehouseAutoScaleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableWarehouseAutoScaleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxClusterCount) {
		body["maxClusterCount"] = request.MaxClusterCount
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableWarehouseAutoScale"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/enableWarehouseAutoScale"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableWarehouseAutoScaleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启自动弹性
//
// @param request - EnableWarehouseAutoScaleRequest
//
// @return EnableWarehouseAutoScaleResponse
func (client *Client) EnableWarehouseAutoScale(instanceId *string, request *EnableWarehouseAutoScaleRequest) (_result *EnableWarehouseAutoScaleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableWarehouseAutoScaleResponse{}
	_body, _err := client.EnableWarehouseAutoScaleWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SQL执行
//
// @param request - ExecuteStatementRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteStatementResponse
func (client *Client) ExecuteStatementWithOptions(instanceId *string, request *ExecuteStatementRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteStatementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["dbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxBytes) {
		body["maxBytes"] = request.MaxBytes
	}

	if !dara.IsNil(request.MaxRows) {
		body["maxRows"] = request.MaxRows
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	if !dara.IsNil(request.QueryTimeout) {
		body["queryTimeout"] = request.QueryTimeout
	}

	if !dara.IsNil(request.Sql) {
		body["sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteStatement"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/executeStatement"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SQL执行
//
// @param request - ExecuteStatementRequest
//
// @return ExecuteStatementResponse
func (client *Client) ExecuteStatement(instanceId *string, request *ExecuteStatementRequest) (_result *ExecuteStatementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteStatementResponse{}
	_body, _err := client.ExecuteStatementWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得证书信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertificateAttributeResponse
func (client *Client) GetCertificateAttributeWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCertificateAttributeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCertificateAttribute"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/certificateAttribute"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCertificateAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得证书信息
//
// @return GetCertificateAttributeResponse
func (client *Client) GetCertificateAttribute(instanceId *string) (_result *GetCertificateAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCertificateAttributeResponse{}
	_body, _err := client.GetCertificateAttributeWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例是否已开启OpenAPI执行SQL功能
//
// @param request - GetExecuteStatementEnabledRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecuteStatementEnabledResponse
func (client *Client) GetExecuteStatementEnabledWithOptions(instanceId *string, request *GetExecuteStatementEnabledRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExecuteStatementEnabledResponse, _err error) {
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
		Action:      dara.String("GetExecuteStatementEnabled"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/executeStatementEnabled"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExecuteStatementEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例是否已开启OpenAPI执行SQL功能
//
// @param request - GetExecuteStatementEnabledRequest
//
// @return GetExecuteStatementEnabledResponse
func (client *Client) GetExecuteStatementEnabled(instanceId *string, request *GetExecuteStatementEnabledRequest) (_result *GetExecuteStatementEnabledResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExecuteStatementEnabledResponse{}
	_body, _err := client.GetExecuteStatementEnabledWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取holoweb登陆权限
//
// @param request - GetHoloWebLoginSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHoloWebLoginSettingResponse
func (client *Client) GetHoloWebLoginSettingWithOptions(instanceId *string, request *GetHoloWebLoginSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHoloWebLoginSettingResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHoloWebLoginSetting"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/getHoloWebLoginSetting"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHoloWebLoginSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取holoweb登陆权限
//
// @param request - GetHoloWebLoginSettingRequest
//
// @return GetHoloWebLoginSettingResponse
func (client *Client) GetHoloWebLoginSetting(instanceId *string, request *GetHoloWebLoginSettingRequest) (_result *GetHoloWebLoginSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHoloWebLoginSettingResponse{}
	_body, _err := client.GetHoloWebLoginSettingWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId))),
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
// Obtains the details of an instance.
//
// @return GetInstanceResponse
func (client *Client) GetInstance(instanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模型信息
//
// @param request - GetInstanceModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceModelResponse
func (client *Client) GetInstanceModelWithOptions(instanceId *string, request *GetInstanceModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceModelResponse, _err error) {
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
		Action:      dara.String("GetInstanceModel"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/getInstanceModel"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模型信息
//
// @param request - GetInstanceModelRequest
//
// @return GetInstanceModelResponse
func (client *Client) GetInstanceModel(instanceId *string, request *GetInstanceModelRequest) (_result *GetInstanceModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceModelResponse{}
	_body, _err := client.GetInstanceModelWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取上次升级历史
//
// @param request - GetLastUpgradeRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLastUpgradeRecordResponse
func (client *Client) GetLastUpgradeRecordWithOptions(instanceId *string, request *GetLastUpgradeRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLastUpgradeRecordResponse, _err error) {
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
		Action:      dara.String("GetLastUpgradeRecord"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/lastUpgradeRecord"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLastUpgradeRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取上次升级历史
//
// @param request - GetLastUpgradeRecordRequest
//
// @return GetLastUpgradeRecordResponse
func (client *Client) GetLastUpgradeRecord(instanceId *string, request *GetLastUpgradeRecordRequest) (_result *GetLastUpgradeRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLastUpgradeRecordResponse{}
	_body, _err := client.GetLastUpgradeRecordWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得根证书
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRootCertificateResponse
func (client *Client) GetRootCertificateWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRootCertificateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRootCertificate"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/rootCertificate"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRootCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得根证书
//
// @return GetRootCertificateResponse
func (client *Client) GetRootCertificate(instanceId *string) (_result *GetRootCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRootCertificateResponse{}
	_body, _err := client.GetRootCertificateWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取周期备份配置
//
// @param request - GetScheduledBackupConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduledBackupConfigResponse
func (client *Client) GetScheduledBackupConfigWithOptions(request *GetScheduledBackupConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScheduledBackupConfigResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScheduleType) {
		query["scheduleType"] = request.ScheduleType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheduledBackupConfig"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/backups/scheduledConfig"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduledBackupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取周期备份配置
//
// @param request - GetScheduledBackupConfigRequest
//
// @return GetScheduledBackupConfigResponse
func (client *Client) GetScheduledBackupConfig(request *GetScheduledBackupConfigRequest) (_result *GetScheduledBackupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScheduledBackupConfigResponse{}
	_body, _err := client.GetScheduledBackupConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取升级状态
//
// @param request - GetUpgradeStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUpgradeStatusResponse
func (client *Client) GetUpgradeStatusWithOptions(instanceId *string, request *GetUpgradeStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUpgradeStatusResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUpgradeStatus"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/upgrade/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUpgradeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取升级状态
//
// @param request - GetUpgradeStatusRequest
//
// @return GetUpgradeStatusResponse
func (client *Client) GetUpgradeStatus(instanceId *string, request *GetUpgradeStatusRequest) (_result *GetUpgradeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUpgradeStatusResponse{}
	_body, _err := client.GetUpgradeStatusWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 是否可升级
//
// @param request - GetUpgradeableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUpgradeableResponse
func (client *Client) GetUpgradeableWithOptions(instanceId *string, request *GetUpgradeableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUpgradeableResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUpgradeable"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/upgradeable"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUpgradeableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 是否可升级
//
// @param request - GetUpgradeableRequest
//
// @return GetUpgradeableResponse
func (client *Client) GetUpgradeable(instanceId *string, request *GetUpgradeableRequest) (_result *GetUpgradeableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUpgradeableResponse{}
	_body, _err := client.GetUpgradeableWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details of a virtual warehouse instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWarehouseDetailResponse
func (client *Client) GetWarehouseDetailWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWarehouseDetailResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWarehouseDetail"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/getWarehouseDetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWarehouseDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details of a virtual warehouse instance.
//
// @return GetWarehouseDetailResponse
func (client *Client) GetWarehouseDetail(instanceId *string) (_result *GetWarehouseDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWarehouseDetailResponse{}
	_body, _err := client.GetWarehouseDetailWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DB授权
//
// @param request - GrantDatabasePermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantDatabasePermissionResponse
func (client *Client) GrantDatabasePermissionWithOptions(instanceId *string, request *GrantDatabasePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantDatabasePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		body["databaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Privileges) {
		body["privileges"] = request.Privileges
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantDatabasePermission"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/grantDatabasePermission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantDatabasePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DB授权
//
// @param request - GrantDatabasePermissionRequest
//
// @return GrantDatabasePermissionResponse
func (client *Client) GrantDatabasePermission(instanceId *string, request *GrantDatabasePermissionRequest) (_result *GrantDatabasePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantDatabasePermissionResponse{}
	_body, _err := client.GrantDatabasePermissionWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DB授权
//
// @param request - GrantSchemaPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantSchemaPermissionResponse
func (client *Client) GrantSchemaPermissionWithOptions(instanceId *string, request *GrantSchemaPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantSchemaPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		body["databaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Privileges) {
		body["privileges"] = request.Privileges
	}

	if !dara.IsNil(request.SchemaName) {
		body["schemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantSchemaPermission"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/grantSchemaPermission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantSchemaPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DB授权
//
// @param request - GrantSchemaPermissionRequest
//
// @return GrantSchemaPermissionResponse
func (client *Client) GrantSchemaPermission(instanceId *string, request *GrantSchemaPermissionRequest) (_result *GrantSchemaPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantSchemaPermissionResponse{}
	_body, _err := client.GrantSchemaPermissionWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DB授权
//
// @param request - GrantTablePermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantTablePermissionResponse
func (client *Client) GrantTablePermissionWithOptions(instanceId *string, request *GrantTablePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantTablePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllTable) {
		body["allTable"] = request.AllTable
	}

	if !dara.IsNil(request.ColumnNames) {
		body["columnNames"] = request.ColumnNames
	}

	if !dara.IsNil(request.DatabaseName) {
		body["databaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Privileges) {
		body["privileges"] = request.Privileges
	}

	if !dara.IsNil(request.SchemaName) {
		body["schemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		body["tableName"] = request.TableName
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantTablePermission"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/grantTablePermission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantTablePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DB授权
//
// @param request - GrantTablePermissionRequest
//
// @return GrantTablePermissionResponse
func (client *Client) GrantTablePermission(instanceId *string, request *GrantTablePermissionRequest) (_result *GrantTablePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantTablePermissionResponse{}
	_body, _err := client.GrantTablePermissionWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of backups. A backup is a full data snapshot of an instance at the end of the snapshot time. You can purchase another instance to completely restore the original data.
//
// @param request - ListBackupDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBackupDataResponse
func (client *Client) ListBackupDataWithOptions(request *ListBackupDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBackupDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupType) {
		query["backupType"] = request.BackupType
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBackupData"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/backups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBackupDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of backups. A backup is a full data snapshot of an instance at the end of the snapshot time. You can purchase another instance to completely restore the original data.
//
// @param request - ListBackupDataRequest
//
// @return ListBackupDataResponse
func (client *Client) ListBackupData(request *ListBackupDataRequest) (_result *ListBackupDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBackupDataResponse{}
	_body, _err := client.ListBackupDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取DB列表
//
// @param request - ListDatabasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithOptions(instanceId *string, request *ListDatabasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.External) {
		query["external"] = request.External
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/listDatabases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取DB列表
//
// @param request - ListDatabasesRequest
//
// @return ListDatabasesResponse
func (client *Client) ListDatabases(instanceId *string, request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取只读从实例
//
// @param request - ListFollowerInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFollowerInstancesResponse
func (client *Client) ListFollowerInstancesWithOptions(instanceId *string, request *ListFollowerInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFollowerInstancesResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFollowerInstances"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/listFollowerInstances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFollowerInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取只读从实例
//
// @param request - ListFollowerInstancesRequest
//
// @return ListFollowerInstancesResponse
func (client *Client) ListFollowerInstances(instanceId *string, request *ListFollowerInstancesRequest) (_result *ListFollowerInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFollowerInstancesResponse{}
	_body, _err := client.ListFollowerInstancesWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AI资源列表
//
// @param request - ListInstanceModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceModelResponse
func (client *Client) ListInstanceModelWithOptions(request *ListInstanceModelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceModelResponse, _err error) {
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
		Action:      dara.String("ListInstanceModel"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/rpc/listInstanceModel"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI资源列表
//
// @param request - ListInstanceModelRequest
//
// @return ListInstanceModelResponse
func (client *Client) ListInstanceModel(request *ListInstanceModelRequest) (_result *ListInstanceModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceModelResponse{}
	_body, _err := client.ListInstanceModelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CmsInstanceType) {
		body["cmsInstanceType"] = request.CmsInstanceType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances"),
		Method:      dara.String("POST"),
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
// Queries a list of instances.
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
// 获取主实例
//
// @param request - ListLeaderInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLeaderInstancesResponse
func (client *Client) ListLeaderInstancesWithOptions(instanceId *string, request *ListLeaderInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLeaderInstancesResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLeaderInstances"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/primaryInstances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLeaderInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主实例
//
// @param request - ListLeaderInstancesRequest
//
// @return ListLeaderInstancesResponse
func (client *Client) ListLeaderInstances(instanceId *string, request *ListLeaderInstancesRequest) (_result *ListLeaderInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLeaderInstancesResponse{}
	_body, _err := client.ListLeaderInstancesWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例可迁移可用区列表
//
// @param request - ListMigrationZonesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMigrationZonesResponse
func (client *Client) ListMigrationZonesWithOptions(instanceId *string, request *ListMigrationZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMigrationZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMigrationZones"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/listMigrationZones"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMigrationZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例可迁移可用区列表
//
// @param request - ListMigrationZonesRequest
//
// @return ListMigrationZonesResponse
func (client *Client) ListMigrationZones(instanceId *string, request *ListMigrationZonesRequest) (_result *ListMigrationZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMigrationZonesResponse{}
	_body, _err := client.ListMigrationZonesWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出模型列表
//
// @param request - ListModelCatalogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelCatalogResponse
func (client *Client) ListModelCatalogWithOptions(instanceId *string, request *ListModelCatalogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelCatalog"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/listModelCatalog"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出模型列表
//
// @param request - ListModelCatalogRequest
//
// @return ListModelCatalogResponse
func (client *Client) ListModelCatalog(instanceId *string, request *ListModelCatalogRequest) (_result *ListModelCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelCatalogResponse{}
	_body, _err := client.ListModelCatalogWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运维事件列表
//
// @param request - ListOperationEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationEventsResponse
func (client *Client) ListOperationEventsWithOptions(request *ListOperationEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOperationEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EventName) {
		body["eventName"] = request.EventName
	}

	if !dara.IsNil(request.EventNameDesc) {
		body["eventNameDesc"] = request.EventNameDesc
	}

	if !dara.IsNil(request.EventType) {
		body["eventType"] = request.EventType
	}

	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScheduleTimeDesc) {
		body["scheduleTimeDesc"] = request.ScheduleTimeDesc
	}

	if !dara.IsNil(request.State) {
		body["state"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationEvents"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/listOperationEvents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运维事件列表
//
// @param request - ListOperationEventsRequest
//
// @return ListOperationEventsResponse
func (client *Client) ListOperationEvents(request *ListOperationEventsRequest) (_result *ListOperationEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOperationEventsResponse{}
	_body, _err := client.ListOperationEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// holoapp共有云所有开服的region
//
// @param request - ListRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
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
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/regions/listRegions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// holoapp共有云所有开服的region
//
// @param request - ListRegionsRequest
//
// @return ListRegionsResponse
func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取版本列表
//
// @param request - ListUpgradeReleaseVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUpgradeReleaseVersionsResponse
func (client *Client) ListUpgradeReleaseVersionsWithOptions(instanceId *string, request *ListUpgradeReleaseVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUpgradeReleaseVersionsResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUpgradeReleaseVersions"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/upgrade/listReleaseVersions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUpgradeReleaseVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取版本列表
//
// @param request - ListUpgradeReleaseVersionsRequest
//
// @return ListUpgradeReleaseVersionsResponse
func (client *Client) ListUpgradeReleaseVersions(instanceId *string, request *ListUpgradeReleaseVersionsRequest) (_result *ListUpgradeReleaseVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUpgradeReleaseVersionsResponse{}
	_body, _err := client.ListUpgradeReleaseVersionsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分时弹性日志
//
// @param request - ListWarehouseScheduleEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarehouseScheduleEventResponse
func (client *Client) ListWarehouseScheduleEventWithOptions(instanceId *string, request *ListWarehouseScheduleEventRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWarehouseScheduleEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticType) {
		body["elasticType"] = request.ElasticType
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		body["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWarehouseScheduleEvent"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/listWarehouseScheduleEvent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWarehouseScheduleEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分时弹性日志
//
// @param request - ListWarehouseScheduleEventRequest
//
// @return ListWarehouseScheduleEventResponse
func (client *Client) ListWarehouseScheduleEvent(instanceId *string, request *ListWarehouseScheduleEventRequest) (_result *ListWarehouseScheduleEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWarehouseScheduleEventResponse{}
	_body, _err := client.ListWarehouseScheduleEventWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 计算分时弹性计划列表
//
// @param request - ListWarehouseScheduleTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarehouseScheduleTaskResponse
func (client *Client) ListWarehouseScheduleTaskWithOptions(instanceId *string, request *ListWarehouseScheduleTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWarehouseScheduleTaskResponse, _err error) {
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
		Action:      dara.String("ListWarehouseScheduleTask"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/listWarehouseScheduleTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWarehouseScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 计算分时弹性计划列表
//
// @param request - ListWarehouseScheduleTaskRequest
//
// @return ListWarehouseScheduleTaskResponse
func (client *Client) ListWarehouseScheduleTask(instanceId *string, request *ListWarehouseScheduleTaskRequest) (_result *ListWarehouseScheduleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWarehouseScheduleTaskResponse{}
	_body, _err := client.ListWarehouseScheduleTaskWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of virtual warehouse instances.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarehousesResponse
func (client *Client) ListWarehousesWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWarehousesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWarehouses"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/listWarehouses"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWarehousesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of virtual warehouse instances.
//
// @return ListWarehousesResponse
func (client *Client) ListWarehouses(instanceId *string) (_result *ListWarehousesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWarehousesResponse{}
	_body, _err := client.ListWarehousesWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例迁移
//
// @param request - MigrateInstanceZoneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateInstanceZoneResponse
func (client *Client) MigrateInstanceZoneWithOptions(instanceId *string, request *MigrateInstanceZoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateInstanceZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ZoneId) {
		body["zoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateInstanceZone"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/migrateInstanceZone"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateInstanceZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例迁移
//
// @param request - MigrateInstanceZoneRequest
//
// @return MigrateInstanceZoneResponse
func (client *Client) MigrateInstanceZone(instanceId *string, request *MigrateInstanceZoneRequest) (_result *MigrateInstanceZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MigrateInstanceZoneResponse{}
	_body, _err := client.MigrateInstanceZoneWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 准备升级
//
// @param request - PrepareUpgradeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrepareUpgradeResponse
func (client *Client) PrepareUpgradeWithOptions(instanceId *string, request *PrepareUpgradeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PrepareUpgradeResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrepareUpgrade"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/upgrade/prepare"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PrepareUpgradeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 准备升级
//
// @param request - PrepareUpgradeRequest
//
// @return PrepareUpgradeResponse
func (client *Client) PrepareUpgrade(instanceId *string, request *PrepareUpgradeRequest) (_result *PrepareUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PrepareUpgradeResponse{}
	_body, _err := client.PrepareUpgradeWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers shard rebalancing for a virtual warehouse.
//
// @param request - RebalanceHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebalanceHoloWarehouseResponse
func (client *Client) RebalanceHoloWarehouseWithOptions(instanceId *string, request *RebalanceHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RebalanceHoloWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebalanceHoloWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/rebalanceHoloWarehouse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RebalanceHoloWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers shard rebalancing for a virtual warehouse.
//
// @param request - RebalanceHoloWarehouseRequest
//
// @return RebalanceHoloWarehouseResponse
func (client *Client) RebalanceHoloWarehouse(instanceId *string, request *RebalanceHoloWarehouseRequest) (_result *RebalanceHoloWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RebalanceHoloWarehouseResponse{}
	_body, _err := client.RebalanceHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a virtual warehouse.
//
// @param request - RenameHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameHoloWarehouseResponse
func (client *Client) RenameHoloWarehouseWithOptions(instanceId *string, request *RenameHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenameHoloWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NewWarehouseName) {
		body["newWarehouseName"] = request.NewWarehouseName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameHoloWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/renameHoloWarehouse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameHoloWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a virtual warehouse.
//
// @param request - RenameHoloWarehouseRequest
//
// @return RenameHoloWarehouseResponse
func (client *Client) RenameHoloWarehouse(instanceId *string, request *RenameHoloWarehouseRequest) (_result *RenameHoloWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenameHoloWarehouseResponse{}
	_body, _err := client.RenameHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually renews a Hologres instance. You can enable monthly auto-renewal when you renew a Hologres instance.
//
// Description:
//
// >  Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
//   - For more information about the billing of Hologres, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview).
//
//   - For more information about how to renew a Hologres instance, see [Manage renewals](https://www.alibabacloud.com/help/zh/hologres/product-overview/manage-renewals?spm=a2c63.p38356.0.0.38e731c9VAwtDP).
//
//   - You can renew only subscription instances.
//
// @param request - RenewInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(instanceId *string, request *RenewInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		body["autoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.Duration) {
		body["duration"] = request.Duration
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/renew"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually renews a Hologres instance. You can enable monthly auto-renewal when you renew a Hologres instance.
//
// Description:
//
// >  Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
//   - For more information about the billing of Hologres, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview).
//
//   - For more information about how to renew a Hologres instance, see [Manage renewals](https://www.alibabacloud.com/help/zh/hologres/product-overview/manage-renewals?spm=a2c63.p38356.0.0.38e731c9VAwtDP).
//
//   - You can renew only subscription instances.
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(instanceId *string, request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新证书
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewSSLCertificateResponse
func (client *Client) RenewSSLCertificateWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewSSLCertificateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewSSLCertificate"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/renewSSLCertificate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewSSLCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新证书
//
// @return RenewSSLCertificateResponse
func (client *Client) RenewSSLCertificate(instanceId *string) (_result *RenewSSLCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewSSLCertificateResponse{}
	_body, _err := client.RenewSSLCertificateWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a virtual warehouse.
//
// @param request - RestartHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartHoloWarehouseResponse
func (client *Client) RestartHoloWarehouseWithOptions(instanceId *string, request *RestartHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartHoloWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartHoloWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/restartHoloWarehouse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartHoloWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a virtual warehouse.
//
// @param request - RestartHoloWarehouseRequest
//
// @return RestartHoloWarehouseResponse
func (client *Client) RestartHoloWarehouse(instanceId *string, request *RestartHoloWarehouseRequest) (_result *RestartHoloWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartHoloWarehouseResponse{}
	_body, _err := client.RestartHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an instance.
//
// @return RestartInstanceResponse
func (client *Client) RestartInstance(instanceId *string) (_result *RestartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a virtual warehouse.
//
// @param request - ResumeHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeHoloWarehouseResponse
func (client *Client) ResumeHoloWarehouseWithOptions(instanceId *string, request *ResumeHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeHoloWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeHoloWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/resumeHoloWarehouse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeHoloWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a virtual warehouse.
//
// @param request - ResumeHoloWarehouseRequest
//
// @return ResumeHoloWarehouseResponse
func (client *Client) ResumeHoloWarehouse(instanceId *string, request *ResumeHoloWarehouseRequest) (_result *ResumeHoloWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeHoloWarehouseResponse{}
	_body, _err := client.ResumeHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/resume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes an instance.
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstance(instanceId *string) (_result *ResumeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消DB授权
//
// @param request - RevokeDatabasePermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeDatabasePermissionResponse
func (client *Client) RevokeDatabasePermissionWithOptions(instanceId *string, request *RevokeDatabasePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeDatabasePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		body["databaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Privileges) {
		body["privileges"] = request.Privileges
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeDatabasePermission"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/revokeDatabasePermission"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeDatabasePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消DB授权
//
// @param request - RevokeDatabasePermissionRequest
//
// @return RevokeDatabasePermissionResponse
func (client *Client) RevokeDatabasePermission(instanceId *string, request *RevokeDatabasePermissionRequest) (_result *RevokeDatabasePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeDatabasePermissionResponse{}
	_body, _err := client.RevokeDatabasePermissionWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消Schema授权
//
// @param request - RevokeSchemaPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeSchemaPermissionResponse
func (client *Client) RevokeSchemaPermissionWithOptions(instanceId *string, request *RevokeSchemaPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeSchemaPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		body["databaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Privileges) {
		body["privileges"] = request.Privileges
	}

	if !dara.IsNil(request.SchemaName) {
		body["schemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeSchemaPermission"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/revokeSchemaPermission"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeSchemaPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消Schema授权
//
// @param request - RevokeSchemaPermissionRequest
//
// @return RevokeSchemaPermissionResponse
func (client *Client) RevokeSchemaPermission(instanceId *string, request *RevokeSchemaPermissionRequest) (_result *RevokeSchemaPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeSchemaPermissionResponse{}
	_body, _err := client.RevokeSchemaPermissionWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消表授权
//
// @param request - RevokeTablePermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTablePermissionResponse
func (client *Client) RevokeTablePermissionWithOptions(instanceId *string, request *RevokeTablePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeTablePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllTable) {
		body["allTable"] = request.AllTable
	}

	if !dara.IsNil(request.ColumnNames) {
		body["columnNames"] = request.ColumnNames
	}

	if !dara.IsNil(request.DatabaseName) {
		body["databaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Privileges) {
		body["privileges"] = request.Privileges
	}

	if !dara.IsNil(request.SchemaName) {
		body["schemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		body["tableName"] = request.TableName
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeTablePermission"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/revokeTablePermission"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeTablePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消表授权
//
// @param request - RevokeTablePermissionRequest
//
// @return RevokeTablePermissionResponse
func (client *Client) RevokeTablePermission(instanceId *string, request *RevokeTablePermissionRequest) (_result *RevokeTablePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeTablePermissionResponse{}
	_body, _err := client.RevokeTablePermissionWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales in or out a virtual warehouse.
//
// @param request - ScaleHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleHoloWarehouseResponse
func (client *Client) ScaleHoloWarehouseWithOptions(instanceId *string, request *ScaleHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleHoloWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterCount) {
		body["clusterCount"] = request.ClusterCount
	}

	if !dara.IsNil(request.Cpu) {
		body["cpu"] = request.Cpu
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleHoloWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/scaleHoloWarehouse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleHoloWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales in or out a virtual warehouse.
//
// @param request - ScaleHoloWarehouseRequest
//
// @return ScaleHoloWarehouseResponse
func (client *Client) ScaleHoloWarehouse(instanceId *string, request *ScaleHoloWarehouseRequest) (_result *ScaleHoloWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleHoloWarehouseResponse{}
	_body, _err := client.ScaleHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the specifications and storage space of a Hologres instance.
//
// Description:
//
// > Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
//   - For more information about the billing of Hologres, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview).
//
//   - During the change of computing resource specifications of a Hologres instance, the instance is unavailable. During the change of storage resource specifications of a Hologres instance, the instance can work normally. Do not frequently change instance specifications. For more information, see [Upgrade or downgrade instance specifications](https://www.alibabacloud.com/help/en/hologres/product-overview/upgrade-or-downgrade-instance-specifications).
//
// @param request - ScaleInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleInstanceResponse
func (client *Client) ScaleInstanceWithOptions(instanceId *string, request *ScaleInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ColdStorageSize) {
		body["coldStorageSize"] = request.ColdStorageSize
	}

	if !dara.IsNil(request.Cpu) {
		body["cpu"] = request.Cpu
	}

	if !dara.IsNil(request.EnableServerlessComputing) {
		body["enableServerlessComputing"] = request.EnableServerlessComputing
	}

	if !dara.IsNil(request.GatewayCount) {
		body["gatewayCount"] = request.GatewayCount
	}

	if !dara.IsNil(request.ScaleType) {
		body["scaleType"] = request.ScaleType
	}

	if !dara.IsNil(request.StorageSize) {
		body["storageSize"] = request.StorageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/scale"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the specifications and storage space of a Hologres instance.
//
// Description:
//
// > Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
//   - For more information about the billing of Hologres, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview).
//
//   - During the change of computing resource specifications of a Hologres instance, the instance is unavailable. During the change of storage resource specifications of a Hologres instance, the instance can work normally. Do not frequently change instance specifications. For more information, see [Upgrade or downgrade instance specifications](https://www.alibabacloud.com/help/en/hologres/product-overview/upgrade-or-downgrade-instance-specifications).
//
// @param request - ScaleInstanceRequest
//
// @return ScaleInstanceResponse
func (client *Client) ScaleInstance(instanceId *string, request *ScaleInstanceRequest) (_result *ScaleInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleInstanceResponse{}
	_body, _err := client.ScaleInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/stop"),
		Method:      dara.String("POST"),
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
// Stops an instance.
//
// @return StopInstanceResponse
func (client *Client) StopInstance(instanceId *string) (_result *StopInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends a virtual warehouse.
//
// @param request - SuspendHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendHoloWarehouseResponse
func (client *Client) SuspendHoloWarehouseWithOptions(instanceId *string, request *SuspendHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SuspendHoloWarehouseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendHoloWarehouse"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/suspendHoloWarehouse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendHoloWarehouseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends a virtual warehouse.
//
// @param request - SuspendHoloWarehouseRequest
//
// @return SuspendHoloWarehouseResponse
func (client *Client) SuspendHoloWarehouse(instanceId *string, request *SuspendHoloWarehouseRequest) (_result *SuspendHoloWarehouseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SuspendHoloWarehouseResponse{}
	_body, _err := client.SuspendHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增tag
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.ResourceIds) {
		body["resourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.TagOwnerUid) {
		body["tagOwnerUid"] = request.TagOwnerUid
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tag"),
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
// 新增tag
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
// 解除绑定主实例
//
// @param request - UnBindLeaderInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnBindLeaderInstanceResponse
func (client *Client) UnBindLeaderInstanceWithOptions(instanceId *string, request *UnBindLeaderInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnBindLeaderInstanceResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnBindLeaderInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/unBindReadOnly"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnBindLeaderInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除绑定主实例
//
// @param request - UnBindLeaderInstanceRequest
//
// @return UnBindLeaderInstanceResponse
func (client *Client) UnBindLeaderInstance(instanceId *string, request *UnBindLeaderInstanceRequest) (_result *UnBindLeaderInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnBindLeaderInstanceResponse{}
	_body, _err := client.UnBindLeaderInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除tag
//
// @param request - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.ResourceIds) {
		body["resourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeys) {
		body["tagKeys"] = request.TagKeys
	}

	if !dara.IsNil(request.TagOwnerUid) {
		body["tagOwnerUid"] = request.TagOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tag/unTag"),
		Method:      dara.String("POST"),
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
// 删除tag
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
// 修改小版本自动升级开关
//
// @param request - UpdateAutoUpgradeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoUpgradeResponse
func (client *Client) UpdateAutoUpgradeWithOptions(instanceId *string, request *UpdateAutoUpgradeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAutoUpgradeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoUpgrade) {
		body["autoUpgrade"] = request.AutoUpgrade
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutoUpgrade"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/autoUpgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutoUpgradeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改小版本自动升级开关
//
// @param request - UpdateAutoUpgradeRequest
//
// @return UpdateAutoUpgradeResponse
func (client *Client) UpdateAutoUpgrade(instanceId *string, request *UpdateAutoUpgradeRequest) (_result *UpdateAutoUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAutoUpgradeResponse{}
	_body, _err := client.UpdateAutoUpgradeWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新备份描述
//
// @param request - UpdateBackupDataDescRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupDataDescResponse
func (client *Client) UpdateBackupDataDescWithOptions(id *string, request *UpdateBackupDataDescRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBackupDataDescResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Desc) {
		body["desc"] = request.Desc
	}

	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBackupDataDesc"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/backups/" + dara.PercentEncode(dara.StringValue(id)) + "/desc"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBackupDataDescResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新备份描述
//
// @param request - UpdateBackupDataDescRequest
//
// @return UpdateBackupDataDescResponse
func (client *Client) UpdateBackupDataDesc(id *string, request *UpdateBackupDataDescRequest) (_result *UpdateBackupDataDescResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBackupDataDescResponse{}
	_body, _err := client.UpdateBackupDataDescWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改holoweb登陆权限
//
// @param request - UpdateHoloWebLoginSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHoloWebLoginSettingResponse
func (client *Client) UpdateHoloWebLoginSettingWithOptions(instanceId *string, request *UpdateHoloWebLoginSettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHoloWebLoginSettingResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowExternalAccountsLogin) {
		body["allowExternalAccountsLogin"] = request.AllowExternalAccountsLogin
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHoloWebLoginSetting"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/updateHoloWebLoginSetting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHoloWebLoginSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改holoweb登陆权限
//
// @param request - UpdateHoloWebLoginSettingRequest
//
// @return UpdateHoloWebLoginSettingResponse
func (client *Client) UpdateHoloWebLoginSetting(instanceId *string, request *UpdateHoloWebLoginSettingRequest) (_result *UpdateHoloWebLoginSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHoloWebLoginSettingResponse{}
	_body, _err := client.UpdateHoloWebLoginSettingWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of an instance.
//
// @param request - UpdateInstanceNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceNameWithOptions(instanceId *string, request *UpdateInstanceNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		body["instanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceName"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/instanceName"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of an instance.
//
// @param request - UpdateInstanceNameRequest
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceName(instanceId *string, request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the network configuration of an instance.
//
// @param request - UpdateInstanceNetworkTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNetworkTypeResponse
func (client *Client) UpdateInstanceNetworkTypeWithOptions(instanceId *string, request *UpdateInstanceNetworkTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceNetworkTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnyTunnelToSingleTunnel) {
		body["anyTunnelToSingleTunnel"] = request.AnyTunnelToSingleTunnel
	}

	if !dara.IsNil(request.NetworkTypes) {
		body["networkTypes"] = request.NetworkTypes
	}

	if !dara.IsNil(request.VSwitchId) {
		body["vSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		body["vpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcOwnerId) {
		body["vpcOwnerId"] = request.VpcOwnerId
	}

	if !dara.IsNil(request.VpcRegionId) {
		body["vpcRegionId"] = request.VpcRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceNetworkType"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/network"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the network configuration of an instance.
//
// @param request - UpdateInstanceNetworkTypeRequest
//
// @return UpdateInstanceNetworkTypeResponse
func (client *Client) UpdateInstanceNetworkType(instanceId *string, request *UpdateInstanceNetworkTypeRequest) (_result *UpdateInstanceNetworkTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNetworkTypeResponse{}
	_body, _err := client.UpdateInstanceNetworkTypeWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改端口号
//
// @param request - UpdateInstancePortRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstancePortResponse
func (client *Client) UpdateInstancePortWithOptions(instanceId *string, request *UpdateInstancePortRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstancePortResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Port) {
		body["port"] = request.Port
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstancePort"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/updatePort"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstancePortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改端口号
//
// @param request - UpdateInstancePortRequest
//
// @return UpdateInstancePortResponse
func (client *Client) UpdateInstancePort(instanceId *string, request *UpdateInstancePortRequest) (_result *UpdateInstancePortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstancePortResponse{}
	_body, _err := client.UpdateInstancePortWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改可维护时间窗口
//
// @param request - UpdateMaintenanceWindowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMaintenanceWindowResponse
func (client *Client) UpdateMaintenanceWindowWithOptions(instanceId *string, request *UpdateMaintenanceWindowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMaintenanceWindowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMaintenanceWindow"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/maintenanceWindow"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMaintenanceWindowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改可维护时间窗口
//
// @param request - UpdateMaintenanceWindowRequest
//
// @return UpdateMaintenanceWindowResponse
func (client *Client) UpdateMaintenanceWindow(instanceId *string, request *UpdateMaintenanceWindowRequest) (_result *UpdateMaintenanceWindowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMaintenanceWindowResponse{}
	_body, _err := client.UpdateMaintenanceWindowWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模型服务
//
// @param request - UpdateModelServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelServiceResponse
func (client *Client) UpdateModelServiceWithOptions(instanceId *string, request *UpdateModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cpu) {
		body["cpu"] = request.Cpu
	}

	if !dara.IsNil(request.Gpu) {
		body["gpu"] = request.Gpu
	}

	if !dara.IsNil(request.Memory) {
		body["memory"] = request.Memory
	}

	if !dara.IsNil(request.ModelServiceName) {
		body["modelServiceName"] = request.ModelServiceName
	}

	if !dara.IsNil(request.ModelType) {
		body["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.ServiceCount) {
		body["serviceCount"] = request.ServiceCount
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModelService"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/updateModelService"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模型服务
//
// @param request - UpdateModelServiceRequest
//
// @return UpdateModelServiceResponse
func (client *Client) UpdateModelService(instanceId *string, request *UpdateModelServiceRequest) (_result *UpdateModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelServiceResponse{}
	_body, _err := client.UpdateModelServiceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改运维事件执行时间
//
// @param request - UpdateOperationEventScheduleTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOperationEventScheduleTimeResponse
func (client *Client) UpdateOperationEventScheduleTimeWithOptions(instanceId *string, request *UpdateOperationEventScheduleTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateOperationEventScheduleTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.ScheduleTime) {
		body["scheduleTime"] = request.ScheduleTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOperationEventScheduleTime"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/updateOperationEventScheduleTime"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOperationEventScheduleTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改运维事件执行时间
//
// @param request - UpdateOperationEventScheduleTimeRequest
//
// @return UpdateOperationEventScheduleTimeResponse
func (client *Client) UpdateOperationEventScheduleTime(instanceId *string, request *UpdateOperationEventScheduleTimeRequest) (_result *UpdateOperationEventScheduleTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateOperationEventScheduleTimeResponse{}
	_body, _err := client.UpdateOperationEventScheduleTimeWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置周期备份
//
// @param request - UpdateScheduledBackupConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduledBackupConfigResponse
func (client *Client) UpdateScheduledBackupConfigWithOptions(request *UpdateScheduledBackupConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateScheduledBackupConfigResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataKeepQuantity) {
		body["dataKeepQuantity"] = request.DataKeepQuantity
	}

	if !dara.IsNil(request.DstRegion) {
		body["dstRegion"] = request.DstRegion
	}

	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Hour) {
		body["hour"] = request.Hour
	}

	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ManualDataKeepQuantity) {
		body["manualDataKeepQuantity"] = request.ManualDataKeepQuantity
	}

	if !dara.IsNil(request.ScheduleType) {
		body["scheduleType"] = request.ScheduleType
	}

	if !dara.IsNil(request.Week) {
		body["week"] = request.Week
	}

	if !dara.IsNil(request.ZoneId) {
		body["zoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScheduledBackupConfig"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/backups/scheduledConfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScheduledBackupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置周期备份
//
// @param request - UpdateScheduledBackupConfigRequest
//
// @return UpdateScheduledBackupConfigResponse
func (client *Client) UpdateScheduledBackupConfig(request *UpdateScheduledBackupConfigRequest) (_result *UpdateScheduledBackupConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateScheduledBackupConfigResponse{}
	_body, _err := client.UpdateScheduledBackupConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建分时弹性计划
//
// @param request - UpdateWarehouseScheduleTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWarehouseScheduleTaskResponse
func (client *Client) UpdateWarehouseScheduleTaskWithOptions(instanceId *string, request *UpdateWarehouseScheduleTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWarehouseScheduleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticCu) {
		body["elasticCu"] = request.ElasticCu
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.WarehouseId) {
		body["warehouseId"] = request.WarehouseId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWarehouseScheduleTask"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/updateWarehouseScheduleTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWarehouseScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建分时弹性计划
//
// @param request - UpdateWarehouseScheduleTaskRequest
//
// @return UpdateWarehouseScheduleTaskResponse
func (client *Client) UpdateWarehouseScheduleTask(instanceId *string, request *UpdateWarehouseScheduleTaskRequest) (_result *UpdateWarehouseScheduleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWarehouseScheduleTaskResponse{}
	_body, _err := client.UpdateWarehouseScheduleTaskWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开始升级
//
// @param request - UpgradeInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeInstanceResponse
func (client *Client) UpgradeInstanceWithOptions(instanceId *string, request *UpgradeInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeInstanceResponse, _err error) {
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

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.UpgradeTime) {
		query["upgradeTime"] = request.UpgradeTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeInstance"),
		Version:     dara.String("2022-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/upgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开始升级
//
// @param request - UpgradeInstanceRequest
//
// @return UpgradeInstanceResponse
func (client *Client) UpgradeInstance(instanceId *string, request *UpgradeInstanceRequest) (_result *UpgradeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.UpgradeInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
