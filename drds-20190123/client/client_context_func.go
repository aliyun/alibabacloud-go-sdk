// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - ChangeAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAccountPasswordResponse
func (client *Client) ChangeAccountPasswordWithContext(ctx context.Context, request *ChangeAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ChangeAccountPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeAccountPassword"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeAccountPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeInstanceAzoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeInstanceAzoneResponse
func (client *Client) ChangeInstanceAzoneWithContext(ctx context.Context, request *ChangeInstanceAzoneRequest, runtime *dara.RuntimeOptions) (_result *ChangeInstanceAzoneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeVSwitch) {
		query["ChangeVSwitch"] = request.ChangeVSwitch
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.DrdsRegionId) {
		query["DrdsRegionId"] = request.DrdsRegionId
	}

	if !dara.IsNil(request.NewVSwitch) {
		query["NewVSwitch"] = request.NewVSwitch
	}

	if !dara.IsNil(request.OriginAzoneId) {
		query["OriginAzoneId"] = request.OriginAzoneId
	}

	if !dara.IsNil(request.TargetAzoneId) {
		query["TargetAzoneId"] = request.TargetAzoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeInstanceAzone"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeInstanceAzoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckDrdsDbNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDrdsDbNameResponse
func (client *Client) CheckDrdsDbNameWithContext(ctx context.Context, request *CheckDrdsDbNameRequest, runtime *dara.RuntimeOptions) (_result *CheckDrdsDbNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDrdsDbName"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDrdsDbNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether scale-out operations such as smooth scale-out can be performed on a PolarDB-X database.
//
// @param request - CheckExpandStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckExpandStatusResponse
func (client *Client) CheckExpandStatusWithContext(ctx context.Context, request *CheckExpandStatusRequest, runtime *dara.RuntimeOptions) (_result *CheckExpandStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckExpandStatus"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckExpandStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the SQL audit feature is enabled for the logical database of a PolarDB-X 1.0 instance.
//
// @param request - CheckSqlAuditEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSqlAuditEnableStatusResponse
func (client *Client) CheckSqlAuditEnableStatusWithContext(ctx context.Context, request *CheckSqlAuditEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *CheckSqlAuditEnableStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckSqlAuditEnableStatus"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSqlAuditEnableStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateDrdsDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDrdsDBResponse
func (client *Client) CreateDrdsDBWithContext(ctx context.Context, request *CreateDrdsDBRequest, runtime *dara.RuntimeOptions) (_result *CreateDrdsDBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DbInstType) {
		query["DbInstType"] = request.DbInstType
	}

	if !dara.IsNil(request.DbInstanceIsCreating) {
		query["DbInstanceIsCreating"] = request.DbInstanceIsCreating
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.Encode) {
		query["Encode"] = request.Encode
	}

	if !dara.IsNil(request.InstDbName) {
		query["InstDbName"] = request.InstDbName
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RdsInstance) {
		query["RdsInstance"] = request.RdsInstance
	}

	if !dara.IsNil(request.RdsSuperAccount) {
		query["RdsSuperAccount"] = request.RdsSuperAccount
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDrdsDB"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDrdsDBResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateDrdsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDrdsInstanceResponse
func (client *Client) CreateDrdsInstanceWithContext(ctx context.Context, request *CreateDrdsInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDrdsInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceSeries) {
		query["InstanceSeries"] = request.InstanceSeries
	}

	if !dara.IsNil(request.IsAutoRenew) {
		query["IsAutoRenew"] = request.IsAutoRenew
	}

	if !dara.IsNil(request.MasterInstId) {
		query["MasterInstId"] = request.MasterInstId
	}

	if !dara.IsNil(request.MySQLVersion) {
		query["MySQLVersion"] = request.MySQLVersion
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.IsHa) {
		query["isHa"] = request.IsHa
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDrdsInstance"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDrdsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateInstanceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceAccountResponse
func (client *Client) CreateInstanceAccountWithContext(ctx context.Context, request *CreateInstanceAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DbPrivilege) {
		query["DbPrivilege"] = request.DbPrivilege
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceAccount"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateInstanceInternetAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceInternetAddressResponse
func (client *Client) CreateInstanceInternetAddressWithContext(ctx context.Context, request *CreateInstanceInternetAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceInternetAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceInternetAddress"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceInternetAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an order to purchase an ApsaraDB RDS for MySQL instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and pricing of PolarDB-X 1.0. For more information, visit the [pricing page](https://www.aliyun.com/price/product#/rds/detail).
//
// @param request - CreateOrderForRdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderForRdsResponse
func (client *Client) CreateOrderForRdsWithContext(ctx context.Context, request *CreateOrderForRdsRequest, runtime *dara.RuntimeOptions) (_result *CreateOrderForRdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrderForRds"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderForRdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateShardTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateShardTaskResponse
func (client *Client) CreateShardTaskWithContext(ctx context.Context, request *CreateShardTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateShardTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceTableName) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !dara.IsNil(request.TargetTableName) {
		query["TargetTableName"] = request.TargetTableName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateShardTask"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateShardTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackMenuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackMenuResponse
func (client *Client) DescribeBackMenuWithContext(ctx context.Context, request *DescribeBackMenuRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackMenuResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackMenu"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackMenuResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackupDbsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupDbsResponse
func (client *Client) DescribeBackupDbsWithContext(ctx context.Context, request *DescribeBackupDbsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupDbsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PreferredRestoreTime) {
		query["PreferredRestoreTime"] = request.PreferredRestoreTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupDbs"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupDbsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup settings of local logs.
//
// @param request - DescribeBackupLocalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupLocalResponse
func (client *Client) DescribeBackupLocalWithContext(ctx context.Context, request *DescribeBackupLocalRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupLocalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupLocal"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupLocalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a backup policy.
//
// @param request - DescribeBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicyWithContext(ctx context.Context, request *DescribeBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackupSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetsResponse
func (client *Client) DescribeBackupSetsWithContext(ctx context.Context, request *DescribeBackupSetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupSetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupSets"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackupTimesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupTimesResponse
func (client *Client) DescribeBackupTimesWithContext(ctx context.Context, request *DescribeBackupTimesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupTimesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupTimes"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupTimesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBroadcastTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBroadcastTablesResponse
func (client *Client) DescribeBroadcastTablesWithContext(ctx context.Context, request *DescribeBroadcastTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBroadcastTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBroadcastTables"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBroadcastTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDbInstanceDbsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDbInstanceDbsResponse
func (client *Client) DescribeDbInstanceDbsWithContext(ctx context.Context, request *DescribeDbInstanceDbsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDbInstanceDbsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DbInstType) {
		query["DbInstType"] = request.DbInstType
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDbInstanceDbs"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDbInstanceDbsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries DescribeDbInstances of the storage layer, such as RDS or PolarDB.
//
// @param request - DescribeDbInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDbInstancesResponse
func (client *Client) DescribeDbInstancesWithContext(ctx context.Context, request *DescribeDbInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDbInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstType) {
		query["DbInstType"] = request.DbInstType
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDbInstances"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDbInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsDBResponse
func (client *Client) DescribeDrdsDBWithContext(ctx context.Context, request *DescribeDrdsDBRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsDBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsDB"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsDBResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the information of the PolarDB cluster in the DRDS logical database.
//
// @param request - DescribeDrdsDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsDBClusterResponse
func (client *Client) DescribeDrdsDBClusterWithContext(ctx context.Context, request *DescribeDrdsDBClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsDBClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsDBCluster"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsDBClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsDBIpWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsDBIpWhiteListResponse
func (client *Client) DescribeDrdsDBIpWhiteListWithContext(ctx context.Context, request *DescribeDrdsDBIpWhiteListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsDBIpWhiteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsDBIpWhiteList"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsDBIpWhiteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsDBsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsDBsResponse
func (client *Client) DescribeDrdsDBsWithContext(ctx context.Context, request *DescribeDrdsDBsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsDBsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsDBs"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsDBsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsDbInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsDbInstanceResponse
func (client *Client) DescribeDrdsDbInstanceWithContext(ctx context.Context, request *DescribeDrdsDbInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsDbInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsDbInstance"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsDbInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries ApsaraDB RDS for MySQL instances that are used to store the data of a database.
//
// @param request - DescribeDrdsDbInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsDbInstancesResponse
func (client *Client) DescribeDrdsDbInstancesWithContext(ctx context.Context, request *DescribeDrdsDbInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsDbInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsDbInstances"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsDbInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询存储实例列表
//
// @param request - DescribeDrdsDbRdsNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsDbRdsNameListResponse
func (client *Client) DescribeDrdsDbRdsNameListWithContext(ctx context.Context, request *DescribeDrdsDbRdsNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsDbRdsNameListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsDbRdsNameList"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsDbRdsNameListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a PolarDB-X 1.0 instance.
//
// @param request - DescribeDrdsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsInstanceResponse
func (client *Client) DescribeDrdsInstanceWithContext(ctx context.Context, request *DescribeDrdsInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsInstance"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsInstanceDbMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsInstanceDbMonitorResponse
func (client *Client) DescribeDrdsInstanceDbMonitorWithContext(ctx context.Context, request *DescribeDrdsInstanceDbMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsInstanceDbMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
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
		Action:      dara.String("DescribeDrdsInstanceDbMonitor"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsInstanceDbMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsInstanceLevelTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsInstanceLevelTasksResponse
func (client *Client) DescribeDrdsInstanceLevelTasksWithContext(ctx context.Context, request *DescribeDrdsInstanceLevelTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsInstanceLevelTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsInstanceLevelTasks"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsInstanceLevelTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询监控数据
//
// @param request - DescribeDrdsInstanceMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsInstanceMonitorResponse
func (client *Client) DescribeDrdsInstanceMonitorWithContext(ctx context.Context, request *DescribeDrdsInstanceMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsInstanceMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.PeriodMultiple) {
		query["PeriodMultiple"] = request.PeriodMultiple
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
		Action:      dara.String("DescribeDrdsInstanceMonitor"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsInstanceMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsInstanceVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsInstanceVersionResponse
func (client *Client) DescribeDrdsInstanceVersionWithContext(ctx context.Context, request *DescribeDrdsInstanceVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsInstanceVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsInstanceVersion"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsInstanceVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instances that meet the specified conditions.
//
// @param request - DescribeDrdsInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsInstancesResponse
func (client *Client) DescribeDrdsInstancesWithContext(ctx context.Context, request *DescribeDrdsInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
	}

	if !dara.IsNil(request.Mix) {
		query["Mix"] = request.Mix
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductVersion) {
		query["ProductVersion"] = request.ProductVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsInstances"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsParamsResponse
func (client *Client) DescribeDrdsParamsWithContext(ctx context.Context, request *DescribeDrdsParamsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsParamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.ParamLevel) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsParams"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all custom ApsaraDB RDS for MySQL instances in a PolarDB-X instance.
//
// @param request - DescribeDrdsRdsInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsRdsInstancesResponse
func (client *Client) DescribeDrdsRdsInstancesWithContext(ctx context.Context, request *DescribeDrdsRdsInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsRdsInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsRdsInstances"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsRdsInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the database shards of an PolarDB-X 1.0 instance.
//
// @param request - DescribeDrdsShardingDbsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsShardingDbsResponse
func (client *Client) DescribeDrdsShardingDbsWithContext(ctx context.Context, request *DescribeDrdsShardingDbsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsShardingDbsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DbNamePattern) {
		query["DbNamePattern"] = request.DbNamePattern
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsShardingDbs"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsShardingDbsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a slow SQL query.
//
// @param request - DescribeDrdsSlowSqlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsSlowSqlsResponse
func (client *Client) DescribeDrdsSlowSqlsWithContext(ctx context.Context, request *DescribeDrdsSlowSqlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsSlowSqlsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExeTime) {
		query["ExeTime"] = request.ExeTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeDrdsSlowSqls"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsSlowSqlsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SQL audit details of a PolarDB-X 1.0 instance.
//
// @param request - DescribeDrdsSqlAuditStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsSqlAuditStatusResponse
func (client *Client) DescribeDrdsSqlAuditStatusWithContext(ctx context.Context, request *DescribeDrdsSqlAuditStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsSqlAuditStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsSqlAuditStatus"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsSqlAuditStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDrdsTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDrdsTasksResponse
func (client *Client) DescribeDrdsTasksWithContext(ctx context.Context, request *DescribeDrdsTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDrdsTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDrdsTasks"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDrdsTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeExpandLogicTableInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpandLogicTableInfoListResponse
func (client *Client) DescribeExpandLogicTableInfoListWithContext(ctx context.Context, request *DescribeExpandLogicTableInfoListRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpandLogicTableInfoListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpandLogicTableInfoList"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpandLogicTableInfoListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about databases on which hots-pot scale-out is performed.
//
// @param request - DescribeHotDbListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHotDbListResponse
func (client *Client) DescribeHotDbListWithContext(ctx context.Context, request *DescribeHotDbListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHotDbListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHotDbList"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHotDbListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeInstDbLogInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstDbLogInfoResponse
func (client *Client) DescribeInstDbLogInfoWithContext(ctx context.Context, request *DescribeInstDbLogInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstDbLogInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstDbLogInfo"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstDbLogInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of the Log Service project and the Logstore used by the SQL audit feature.
//
// @param request - DescribeInstDbSlsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstDbSlsInfoResponse
func (client *Client) DescribeInstDbSlsInfoWithContext(ctx context.Context, request *DescribeInstDbSlsInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstDbSlsInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstDbSlsInfo"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstDbSlsInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an instance account.
//
// @param request - DescribeInstanceAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAccountsResponse
func (client *Client) DescribeInstanceAccountsWithContext(ctx context.Context, request *DescribeInstanceAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAccounts"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check whether zone switching is enabled
//
// @param request - DescribeInstanceSwitchAzoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSwitchAzoneResponse
func (client *Client) DescribeInstanceSwitchAzoneWithContext(ctx context.Context, request *DescribeInstanceSwitchAzoneRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSwitchAzoneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSwitchAzone"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSwitchAzoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether you can change the network type of a PolarDB-X 1.0 instance.
//
// Description:
//
// ***
//
// @param request - DescribeInstanceSwitchNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSwitchNetworkResponse
func (client *Client) DescribeInstanceSwitchNetworkWithContext(ctx context.Context, request *DescribeInstanceSwitchNetworkRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSwitchNetworkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSwitchNetwork"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSwitchNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribePreCheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreCheckResultResponse
func (client *Client) DescribePreCheckResultWithContext(ctx context.Context, request *DescribePreCheckResultRequest, runtime *dara.RuntimeOptions) (_result *DescribePreCheckResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePreCheckResult"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePreCheckResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRDSPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRDSPerformanceResponse
func (client *Client) DescribeRDSPerformanceWithContext(ctx context.Context, request *DescribeRDSPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRDSPerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstType) {
		query["DbInstType"] = request.DbInstType
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keys) {
		query["Keys"] = request.Keys
	}

	if !dara.IsNil(request.RdsInstanceId) {
		query["RdsInstanceId"] = request.RdsInstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRDSPerformance"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRDSPerformanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRdsCommodityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsCommodityResponse
func (client *Client) DescribeRdsCommodityWithContext(ctx context.Context, request *DescribeRdsCommodityRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsCommodityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdsCommodity"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsCommodityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRdsPerformanceSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsPerformanceSummaryResponse
func (client *Client) DescribeRdsPerformanceSummaryWithContext(ctx context.Context, request *DescribeRdsPerformanceSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsPerformanceSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RdsInstanceId) {
		query["RdsInstanceId"] = request.RdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdsPerformanceSummary"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsPerformanceSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRdsSuperAccountInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsSuperAccountInstancesResponse
func (client *Client) DescribeRdsSuperAccountInstancesWithContext(ctx context.Context, request *DescribeRdsSuperAccountInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsSuperAccountInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstType) {
		query["DbInstType"] = request.DbInstType
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RdsInstance) {
		query["RdsInstance"] = request.RdsInstance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdsSuperAccountInstances"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsSuperAccountInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the table recycle bin.
//
// @param request - DescribeRecycleBinStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecycleBinStatusResponse
func (client *Client) DescribeRecycleBinStatusWithContext(ctx context.Context, request *DescribeRecycleBinStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecycleBinStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecycleBinStatus"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecycleBinStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tables that can be restored in the recycle bin.
//
// @param request - DescribeRecycleBinTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecycleBinTablesResponse
func (client *Client) DescribeRecycleBinTablesWithContext(ctx context.Context, request *DescribeRecycleBinTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecycleBinTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecycleBinTables"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecycleBinTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeRestoreOrder operation to view the details of the order.
//
// @param request - DescribeRestoreOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreOrderResponse
func (client *Client) DescribeRestoreOrderWithContext(ctx context.Context, request *DescribeRestoreOrderRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupDbNames) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupLevel) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreOrder"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeShardTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeShardTaskInfoResponse
func (client *Client) DescribeShardTaskInfoWithContext(ctx context.Context, request *DescribeShardTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeShardTaskInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceTableName) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !dara.IsNil(request.TargetTableName) {
		query["TargetTableName"] = request.TargetTableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeShardTaskInfo"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeShardTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the flashback tasks that are performed on a PolarDB-X 1.0 instance.
//
// @param request - DescribeSqlFlashbakTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlFlashbakTaskResponse
func (client *Client) DescribeSqlFlashbakTaskWithContext(ctx context.Context, request *DescribeSqlFlashbakTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlFlashbakTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSqlFlashbakTask"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlFlashbakTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the schema of a table.
//
// @param request - DescribeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableResponse
func (client *Client) DescribeTableWithContext(ctx context.Context, request *DescribeTableRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTable"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeTableListByTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTableListByTypeResponse
func (client *Client) DescribeTableListByTypeWithContext(ctx context.Context, request *DescribeTableListByTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeTableListByTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTableListByType"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTableListByTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeTables文档变更
//
// @param request - DescribeTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTablesResponse
func (client *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTables"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the SQL audit feature for a database.
//
// @param request - DisableSqlAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSqlAuditResponse
func (client *Client) DisableSqlAuditWithContext(ctx context.Context, request *DisableSqlAuditRequest, runtime *dara.RuntimeOptions) (_result *DisableSqlAuditResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSqlAudit"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSqlAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IPv6 address.
//
// @param request - EnableInstanceIpv6AddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInstanceIpv6AddressResponse
func (client *Client) EnableInstanceIpv6AddressWithContext(ctx context.Context, request *EnableInstanceIpv6AddressRequest, runtime *dara.RuntimeOptions) (_result *EnableInstanceIpv6AddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInstanceIpv6Address"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInstanceIpv6AddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the SQL audit feature for a database.
//
// @param request - EnableSqlAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSqlAuditResponse
func (client *Client) EnableSqlAuditWithContext(ctx context.Context, request *EnableSqlAuditRequest, runtime *dara.RuntimeOptions) (_result *EnableSqlAuditResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.IsRecall) {
		query["IsRecall"] = request.IsRecall
	}

	if !dara.IsNil(request.RecallEndTimestamp) {
		query["RecallEndTimestamp"] = request.RecallEndTimestamp
	}

	if !dara.IsNil(request.RecallStartTimestamp) {
		query["RecallStartTimestamp"] = request.RecallStartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSqlAudit"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSqlAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableSqlFlashbackMatchSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSqlFlashbackMatchSwitchResponse
func (client *Client) EnableSqlFlashbackMatchSwitchWithContext(ctx context.Context, request *EnableSqlFlashbackMatchSwitchRequest, runtime *dara.RuntimeOptions) (_result *EnableSqlFlashbackMatchSwitchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSqlFlashbackMatchSwitch"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSqlFlashbackMatchSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a logical table that is deleted.
//
// @param request - FlashbackRecycleBinTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlashbackRecycleBinTableResponse
func (client *Client) FlashbackRecycleBinTableWithContext(ctx context.Context, request *FlashbackRecycleBinTableRequest, runtime *dara.RuntimeOptions) (_result *FlashbackRecycleBinTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlashbackRecycleBinTable"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlashbackRecycleBinTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDrdsDbRdsRelationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDrdsDbRdsRelationInfoResponse
func (client *Client) GetDrdsDbRdsRelationInfoWithContext(ctx context.Context, request *GetDrdsDbRdsRelationInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDrdsDbRdsRelationInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDrdsDbRdsRelationInfo"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDrdsDbRdsRelationInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看Tag标签
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Manages a custom ApsaraDB RDS instance at the storage layer.
//
// @param request - ManagePrivateRdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManagePrivateRdsResponse
func (client *Client) ManagePrivateRdsWithContext(ctx context.Context, request *ManagePrivateRdsRequest, runtime *dara.RuntimeOptions) (_result *ManagePrivateRdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.RdsAction) {
		query["RdsAction"] = request.RdsAction
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManagePrivateRds"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManagePrivateRdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPrivilegeResponse
func (client *Client) ModifyAccountPrivilegeWithContext(ctx context.Context, request *ModifyAccountPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPrivilegeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DbPrivilege) {
		query["DbPrivilege"] = request.DbPrivilege
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountPrivilege"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyDrdsInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDrdsInstanceDescriptionResponse
func (client *Client) ModifyDrdsInstanceDescriptionWithContext(ctx context.Context, request *ModifyDrdsInstanceDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDrdsInstanceDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDrdsInstanceDescription"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDrdsInstanceDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyDrdsIpWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDrdsIpWhiteListResponse
func (client *Client) ModifyDrdsIpWhiteListWithContext(ctx context.Context, request *ModifyDrdsIpWhiteListRequest, runtime *dara.RuntimeOptions) (_result *ModifyDrdsIpWhiteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.GroupAttribute) {
		query["GroupAttribute"] = request.GroupAttribute
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IpWhiteList) {
		query["IpWhiteList"] = request.IpWhiteList
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDrdsIpWhiteList"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDrdsIpWhiteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyPolarDbReadWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolarDbReadWeightResponse
func (client *Client) ModifyPolarDbReadWeightWithContext(ctx context.Context, request *ModifyPolarDbReadWeightRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolarDbReadWeightResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DbNodeIds) {
		query["DbNodeIds"] = request.DbNodeIds
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.Weights) {
		query["Weights"] = request.Weights
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolarDbReadWeight"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolarDbReadWeightResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyRdsReadWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRdsReadWeightResponse
func (client *Client) ModifyRdsReadWeightWithContext(ctx context.Context, request *ModifyRdsReadWeightRequest, runtime *dara.RuntimeOptions) (_result *ModifyRdsReadWeightResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.InstanceNames) {
		query["InstanceNames"] = request.InstanceNames
	}

	if !dara.IsNil(request.Weights) {
		query["Weights"] = request.Weights
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRdsReadWeight"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRdsReadWeightResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutStartBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutStartBackupResponse
func (client *Client) PutStartBackupWithContext(ctx context.Context, request *PutStartBackupRequest, runtime *dara.RuntimeOptions) (_result *PutStartBackupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupDbNames) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !dara.IsNil(request.BackupLevel) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutStartBackup"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutStartBackupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefreshDrdsAtomUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshDrdsAtomUrlResponse
func (client *Client) RefreshDrdsAtomUrlWithContext(ctx context.Context, request *RefreshDrdsAtomUrlRequest, runtime *dara.RuntimeOptions) (_result *RefreshDrdsAtomUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshDrdsAtomUrl"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshDrdsAtomUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReleaseInstanceInternetAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceInternetAddressResponse
func (client *Client) ReleaseInstanceInternetAddressWithContext(ctx context.Context, request *ReleaseInstanceInternetAddressRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceInternetAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstanceInternetAddress"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceInternetAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveBackupsSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveBackupsSetResponse
func (client *Client) RemoveBackupsSetWithContext(ctx context.Context, request *RemoveBackupsSetRequest, runtime *dara.RuntimeOptions) (_result *RemoveBackupsSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveBackupsSet"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveBackupsSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveDrdsDbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDrdsDbResponse
func (client *Client) RemoveDrdsDbWithContext(ctx context.Context, request *RemoveDrdsDbRequest, runtime *dara.RuntimeOptions) (_result *RemoveDrdsDbResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDrdsDb"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDrdsDbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveDrdsDbFailedRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDrdsDbFailedRecordResponse
func (client *Client) RemoveDrdsDbFailedRecordWithContext(ctx context.Context, request *RemoveDrdsDbFailedRecordRequest, runtime *dara.RuntimeOptions) (_result *RemoveDrdsDbFailedRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDrdsDbFailedRecord"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDrdsDbFailedRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases an instance.
//
// Description:
//
// > 	- You can call this operation to release an instance that is charged based on only the pay-as-you-go billing method.
//
// >	- If the specifications of the instance are being changed, or one or more databases exist in the instance, you cannot call this operation to release the instance.
//
// @param request - RemoveDrdsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDrdsInstanceResponse
func (client *Client) RemoveDrdsInstanceWithContext(ctx context.Context, request *RemoveDrdsInstanceRequest, runtime *dara.RuntimeOptions) (_result *RemoveDrdsInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDrdsInstance"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDrdsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveInstanceAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInstanceAccountResponse
func (client *Client) RemoveInstanceAccountWithContext(ctx context.Context, request *RemoveInstanceAccountRequest, runtime *dara.RuntimeOptions) (_result *RemoveInstanceAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveInstanceAccount"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveInstanceAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a table in the recycle bin.
//
// @param request - RemoveRecycleBinTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveRecycleBinTableResponse
func (client *Client) RemoveRecycleBinTableWithContext(ctx context.Context, request *RemoveRecycleBinTableRequest, runtime *dara.RuntimeOptions) (_result *RemoveRecycleBinTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveRecycleBinTable"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveRecycleBinTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RestartDrdsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDrdsInstanceResponse
func (client *Client) RestartDrdsInstanceWithContext(ctx context.Context, request *RestartDrdsInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartDrdsInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDrdsInstance"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDrdsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RollbackInstanceVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackInstanceVersionResponse
func (client *Client) RollbackInstanceVersionWithContext(ctx context.Context, request *RollbackInstanceVersionRequest, runtime *dara.RuntimeOptions) (_result *RollbackInstanceVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackInstanceVersion"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackInstanceVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a backup policy.
//
// @param request - SetBackupLocalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetBackupLocalResponse
func (client *Client) SetBackupLocalWithContext(ctx context.Context, request *SetBackupLocalRequest, runtime *dara.RuntimeOptions) (_result *SetBackupLocalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.HighSpaceUsageProtection) {
		query["HighSpaceUsageProtection"] = request.HighSpaceUsageProtection
	}

	if !dara.IsNil(request.LocalLogRetentionHours) {
		query["LocalLogRetentionHours"] = request.LocalLogRetentionHours
	}

	if !dara.IsNil(request.LocalLogRetentionSpace) {
		query["LocalLogRetentionSpace"] = request.LocalLogRetentionSpace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetBackupLocal"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetBackupLocalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetBackupPolicyResponse
func (client *Client) SetBackupPolicyWithContext(ctx context.Context, request *SetBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetBackupPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupDbNames) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !dara.IsNil(request.BackupLevel) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !dara.IsNil(request.BackupLog) {
		query["BackupLog"] = request.BackupLog
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.DataBackupRetentionPeriod) {
		query["DataBackupRetentionPeriod"] = request.DataBackupRetentionPeriod
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.LogBackupRetentionPeriod) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !dara.IsNil(request.PreferredBackupEndTime) {
		query["PreferredBackupEndTime"] = request.PreferredBackupEndTime
	}

	if !dara.IsNil(request.PreferredBackupPeriod) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !dara.IsNil(request.PreferredBackupStartTime) {
		query["PreferredBackupStartTime"] = request.PreferredBackupStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetBackupPolicy"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetBackupPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a broadcast table for a database.
//
// @param request - SetupBroadcastTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetupBroadcastTablesResponse
func (client *Client) SetupBroadcastTablesWithContext(ctx context.Context, request *SetupBroadcastTablesRequest, runtime *dara.RuntimeOptions) (_result *SetupBroadcastTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetupBroadcastTables"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetupBroadcastTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetupDrdsParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetupDrdsParamsResponse
func (client *Client) SetupDrdsParamsWithContext(ctx context.Context, request *SetupDrdsParamsRequest, runtime *dara.RuntimeOptions) (_result *SetupDrdsParamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.ParamLevel) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetupDrdsParams"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetupDrdsParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the table recycle bin for a database.
//
// @param request - SetupRecycleBinStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetupRecycleBinStatusResponse
func (client *Client) SetupRecycleBinStatusWithContext(ctx context.Context, request *SetupRecycleBinStatusRequest, runtime *dara.RuntimeOptions) (_result *SetupRecycleBinStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StatusAction) {
		query["StatusAction"] = request.StatusAction
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetupRecycleBinStatus"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetupRecycleBinStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetupTableResponse
func (client *Client) SetupTableWithContext(ctx context.Context, request *SetupTableRequest, runtime *dara.RuntimeOptions) (_result *SetupTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowFullTableScan) {
		query["AllowFullTableScan"] = request.AllowFullTableScan
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetupTable"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetupTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartRestoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRestoreResponse
func (client *Client) StartRestoreWithContext(ctx context.Context, request *StartRestoreRequest, runtime *dara.RuntimeOptions) (_result *StartRestoreResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupDbNames) {
		query["BackupDbNames"] = request.BackupDbNames
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupLevel) {
		query["BackupLevel"] = request.BackupLevel
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRestore"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRestoreResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a cleanup task for the scale-out of a PolarDB-X database.
//
// @param request - SubmitCleanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCleanTaskResponse
func (client *Client) SubmitCleanTaskWithContext(ctx context.Context, request *SubmitCleanTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitCleanTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.ExpandType) {
		query["ExpandType"] = request.ExpandType
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.ParentJobId) {
		query["ParentJobId"] = request.ParentJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCleanTask"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCleanTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a precheck task for the hot-spot scale-out of a PolarDB-X database. The task is used to check the table that does not contain the primary key.
//
// @param request - SubmitHotExpandPreCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitHotExpandPreCheckTaskResponse
func (client *Client) SubmitHotExpandPreCheckTaskWithContext(ctx context.Context, request *SubmitHotExpandPreCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitHotExpandPreCheckTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstType) {
		query["DbInstType"] = request.DbInstType
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.TableList) {
		query["TableList"] = request.TableList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitHotExpandPreCheckTask"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitHotExpandPreCheckTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a hot-spot scale-out task for a database.
//
// @param request - SubmitHotExpandTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitHotExpandTaskResponse
func (client *Client) SubmitHotExpandTaskWithContext(ctx context.Context, request *SubmitHotExpandTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitHotExpandTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.ExtendedMapping) {
		query["ExtendedMapping"] = request.ExtendedMapping
	}

	if !dara.IsNil(request.InstanceDbMapping) {
		query["InstanceDbMapping"] = request.InstanceDbMapping
	}

	if !dara.IsNil(request.Mapping) {
		query["Mapping"] = request.Mapping
	}

	if !dara.IsNil(request.SupperAccountMapping) {
		query["SupperAccountMapping"] = request.SupperAccountMapping
	}

	if !dara.IsNil(request.TaskDesc) {
		query["TaskDesc"] = request.TaskDesc
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitHotExpandTask"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitHotExpandTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a precheck task for the smooth scale-out of a PolarDB-X database.
//
// @param request - SubmitSmoothExpandPreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSmoothExpandPreCheckResponse
func (client *Client) SubmitSmoothExpandPreCheckWithContext(ctx context.Context, request *SubmitSmoothExpandPreCheckRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmoothExpandPreCheckResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstType) {
		query["DbInstType"] = request.DbInstType
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSmoothExpandPreCheck"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSmoothExpandPreCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a precheck task for the smooth scale-out of a PolarDB-X 1.0 database.
//
// @param request - SubmitSmoothExpandPreCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSmoothExpandPreCheckTaskResponse
func (client *Client) SubmitSmoothExpandPreCheckTaskWithContext(ctx context.Context, request *SubmitSmoothExpandPreCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmoothExpandPreCheckTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSmoothExpandPreCheckTask"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSmoothExpandPreCheckTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitSqlFlashbackTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSqlFlashbackTaskResponse
func (client *Client) SubmitSqlFlashbackTaskWithContext(ctx context.Context, request *SubmitSqlFlashbackTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitSqlFlashbackTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RecallRestoreType) {
		query["RecallRestoreType"] = request.RecallRestoreType
	}

	if !dara.IsNil(request.RecallType) {
		query["RecallType"] = request.RecallType
	}

	if !dara.IsNil(request.SqlPk) {
		query["SqlPk"] = request.SqlPk
	}

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TraceId) {
		query["TraceId"] = request.TraceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSqlFlashbackTask"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSqlFlashbackTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the mode of broadcast tables from the multi-write mode to the asynchronous link mode.
//
// @param request - SwitchGlobalBroadcastTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchGlobalBroadcastTypeResponse
func (client *Client) SwitchGlobalBroadcastTypeWithContext(ctx context.Context, request *SwitchGlobalBroadcastTypeRequest, runtime *dara.RuntimeOptions) (_result *SwitchGlobalBroadcastTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchGlobalBroadcastType"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchGlobalBroadcastTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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

// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Changes the network type of a PolarDB-X 1.0 instance.
//
// @param request - UpdateInstanceNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNetworkResponse
func (client *Client) UpdateInstanceNetworkWithContext(ctx context.Context, request *UpdateInstanceNetworkRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceNetworkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassicExpiredDays) {
		query["ClassicExpiredDays"] = request.ClassicExpiredDays
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RetainClassic) {
		query["RetainClassic"] = request.RetainClassic
	}

	if !dara.IsNil(request.SrcInstanceNetworkType) {
		query["SrcInstanceNetworkType"] = request.SrcInstanceNetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceNetwork"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the specifications of a custom ApsaraDB RDS instance at the storage layer.
//
// @param request - UpdatePrivateRdsClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateRdsClassResponse
func (client *Client) UpdatePrivateRdsClassWithContext(ctx context.Context, request *UpdatePrivateRdsClassRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivateRdsClassResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.PrePayDuration) {
		query["PrePayDuration"] = request.PrePayDuration
	}

	if !dara.IsNil(request.RdsClass) {
		query["RdsClass"] = request.RdsClass
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateRdsClass"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateRdsClassResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateResourceGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupAttributeResponse
func (client *Client) UpdateResourceGroupAttributeWithContext(ctx context.Context, request *UpdateResourceGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceGroupAttribute"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the version of a column-oriented storage instance of a PolarDB-X 1.0 instance.
//
// @param request - UpgradeHiStoreInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeHiStoreInstanceResponse
func (client *Client) UpgradeHiStoreInstanceWithContext(ctx context.Context, request *UpgradeHiStoreInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpgradeHiStoreInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.HistoreInstanceId) {
		query["HistoreInstanceId"] = request.HistoreInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeHiStoreInstance"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeHiStoreInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpgradeInstanceVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeInstanceVersionResponse
func (client *Client) UpgradeInstanceVersionWithContext(ctx context.Context, request *UpgradeInstanceVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeInstanceVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Rpm) {
		query["Rpm"] = request.Rpm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeInstanceVersion"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ValidateShardTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateShardTaskResponse
func (client *Client) ValidateShardTaskWithContext(ctx context.Context, request *ValidateShardTaskRequest, runtime *dara.RuntimeOptions) (_result *ValidateShardTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DrdsInstanceId) {
		query["DrdsInstanceId"] = request.DrdsInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceTableName) {
		query["SourceTableName"] = request.SourceTableName
	}

	if !dara.IsNil(request.TargetTableName) {
		query["TargetTableName"] = request.TargetTableName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateShardTask"),
		Version:     dara.String("2019-01-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateShardTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
