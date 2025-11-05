// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHoloWarehouseResponse
func (client *Client) CreateHoloWarehouseWithContext(ctx context.Context, instanceId *string, request *CreateHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHoloWarehouseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHoloWarehouseResponse
func (client *Client) DeleteHoloWarehouseWithContext(ctx context.Context, instanceId *string, request *DeleteHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHoloWarehouseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, instanceId *string, request *DeleteInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableHiveAccessResponse
func (client *Client) DisableHiveAccessWithContext(ctx context.Context, instanceId *string, request *DisableHiveAccessRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableHiveAccessResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSSLResponse
func (client *Client) DisableSSLWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableSSLResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHiveAccessResponse
func (client *Client) EnableHiveAccessWithContext(ctx context.Context, instanceId *string, request *EnableHiveAccessRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableHiveAccessResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSSLResponse
func (client *Client) EnableSSLWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableSSLResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertificateAttributeResponse
func (client *Client) GetCertificateAttributeWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCertificateAttributeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRootCertificateResponse
func (client *Client) GetRootCertificateWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRootCertificateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWarehouseDetailResponse
func (client *Client) GetWarehouseDetailWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWarehouseDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantDatabasePermissionResponse
func (client *Client) GrantDatabasePermissionWithContext(ctx context.Context, instanceId *string, request *GrantDatabasePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantDatabasePermissionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantSchemaPermissionResponse
func (client *Client) GrantSchemaPermissionWithContext(ctx context.Context, instanceId *string, request *GrantSchemaPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantSchemaPermissionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantTablePermissionResponse
func (client *Client) GrantTablePermissionWithContext(ctx context.Context, instanceId *string, request *GrantTablePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantTablePermissionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBackupDataResponse
func (client *Client) ListBackupDataWithContext(ctx context.Context, request *ListBackupDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBackupDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithContext(ctx context.Context, instanceId *string, request *ListDatabasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarehousesResponse
func (client *Client) ListWarehousesWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWarehousesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebalanceHoloWarehouseResponse
func (client *Client) RebalanceHoloWarehouseWithContext(ctx context.Context, instanceId *string, request *RebalanceHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RebalanceHoloWarehouseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameHoloWarehouseResponse
func (client *Client) RenameHoloWarehouseWithContext(ctx context.Context, instanceId *string, request *RenameHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenameHoloWarehouseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithContext(ctx context.Context, instanceId *string, request *RenewInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewSSLCertificateResponse
func (client *Client) RenewSSLCertificateWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewSSLCertificateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartHoloWarehouseResponse
func (client *Client) RestartHoloWarehouseWithContext(ctx context.Context, instanceId *string, request *RestartHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartHoloWarehouseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeHoloWarehouseResponse
func (client *Client) ResumeHoloWarehouseWithContext(ctx context.Context, instanceId *string, request *ResumeHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeHoloWarehouseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstanceWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeDatabasePermissionResponse
func (client *Client) RevokeDatabasePermissionWithContext(ctx context.Context, instanceId *string, request *RevokeDatabasePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeDatabasePermissionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeSchemaPermissionResponse
func (client *Client) RevokeSchemaPermissionWithContext(ctx context.Context, instanceId *string, request *RevokeSchemaPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeSchemaPermissionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTablePermissionResponse
func (client *Client) RevokeTablePermissionWithContext(ctx context.Context, instanceId *string, request *RevokeTablePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeTablePermissionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleHoloWarehouseResponse
func (client *Client) ScaleHoloWarehouseWithContext(ctx context.Context, instanceId *string, request *ScaleHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleHoloWarehouseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleInstanceResponse
func (client *Client) ScaleInstanceWithContext(ctx context.Context, instanceId *string, request *ScaleInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendHoloWarehouseResponse
func (client *Client) SuspendHoloWarehouseWithContext(ctx context.Context, instanceId *string, request *SuspendHoloWarehouseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SuspendHoloWarehouseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceNameWithContext(ctx context.Context, instanceId *string, request *UpdateInstanceNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNetworkTypeResponse
func (client *Client) UpdateInstanceNetworkTypeWithContext(ctx context.Context, instanceId *string, request *UpdateInstanceNetworkTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceNetworkTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
