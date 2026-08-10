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
		"ap-northeast-2-pop":          dara.String("clickhouse.aliyuncs.com"),
		"ap-southeast-1":              dara.String("clickhouse.aliyuncs.com"),
		"cn-beijing":                  dara.String("clickhouse.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("clickhouse.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("clickhouse.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("clickhouse.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("clickhouse.aliyuncs.com"),
		"cn-edge-1":                   dara.String("clickhouse.aliyuncs.com"),
		"cn-fujian":                   dara.String("clickhouse.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("clickhouse.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("clickhouse.aliyuncs.com"),
		"cn-hongkong":                 dara.String("clickhouse.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("clickhouse.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("clickhouse.aliyuncs.com"),
		"cn-qingdao":                  dara.String("clickhouse.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("clickhouse.aliyuncs.com"),
		"cn-shanghai":                 dara.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("clickhouse.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("clickhouse.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("clickhouse.aliyuncs.com"),
		"cn-wuhan":                    dara.String("clickhouse.aliyuncs.com"),
		"cn-yushanfang":               dara.String("clickhouse.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("clickhouse.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("clickhouse.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("clickhouse.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("clickhouse.aliyuncs.com"),
		"me-east-1":                   dara.String("clickhouse.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("clickhouse.aliyuncs.com"),
		"us-east-1":                   dara.String("clickhouse.aliyuncs.com"),
		"us-west-1":                   dara.String("clickhouse.aliyuncs.com"),
		"us-southeast-1":              dara.String("clickhouse.us-southeast-1.aliyuncs.com"),
		"na-south-1":                  dara.String("clickhouse.na-south-1.aliyuncs.com"),
		"me-central-1":                dara.String("clickhouse.me-central-1.aliyuncs.com"),
		"eu-west-1":                   dara.String("clickhouse.eu-west-1.aliyuncs.com"),
		"eu-central-1":                dara.String("clickhouse.eu-central-1.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("clickhouse.cn-zhangjiakou.aliyuncs.com"),
		"cn-wulanchabu-gic-1":         dara.String("clickhouse.cn-wulanchabu-gic-1.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("clickhouse.cn-wulanchabu.aliyuncs.com"),
		"cn-huhehaote":                dara.String("clickhouse.cn-huhehaote.aliyuncs.com"),
		"cn-guangzhou":                dara.String("clickhouse.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":                  dara.String("clickhouse.cn-chengdu.aliyuncs.com"),
		"ap-southeast-8":              dara.String("clickhouse.ap-southeast-8.aliyuncs.com"),
		"ap-southeast-6":              dara.String("clickhouse.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-5":              dara.String("clickhouse.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-3":              dara.String("clickhouse.aliyuncs.com"),
		"ap-northeast-1":              dara.String("clickhouse.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("clickhouse"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Attaches a whitelist template to instances.
//
// @param request - AttachWhitelistTemplateToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachWhitelistTemplateToInstanceResponse
func (client *Client) AttachWhitelistTemplateToInstanceWithOptions(request *AttachWhitelistTemplateToInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachWhitelistTemplateToInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachWhitelistTemplateToInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachWhitelistTemplateToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a whitelist template to instances.
//
// @param request - AttachWhitelistTemplateToInstanceRequest
//
// @return AttachWhitelistTemplateToInstanceResponse
func (client *Client) AttachWhitelistTemplateToInstance(request *AttachWhitelistTemplateToInstanceRequest) (_result *AttachWhitelistTemplateToInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachWhitelistTemplateToInstanceResponse{}
	_body, _err := client.AttachWhitelistTemplateToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group of a ClickHouse Enterprise instance.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Changes the resource group of a ClickHouse Enterprise instance.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a database account for an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param tmpReq - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithOptions(tmpReq *CreateAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAccountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DmlAuthSetting) {
		request.DmlAuthSettingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DmlAuthSetting, dara.String("DmlAuthSetting"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DmlAuthSettingShrink) {
		query["DmlAuthSetting"] = request.DmlAuthSettingShrink
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database account for an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - CreateAccountRequest
//
// @return CreateAccountResponse
func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backup policy for a specified ClickHouse Enterprise Edition cluster.
//
// @param request - CreateBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupPolicyResponse
func (client *Client) CreateBackupPolicyWithOptions(request *CreateBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.PreferredBackupPeriod) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackupPolicy"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup policy for a specified ClickHouse Enterprise Edition cluster.
//
// @param request - CreateBackupPolicyRequest
//
// @return CreateBackupPolicyResponse
func (client *Client) CreateBackupPolicy(request *CreateBackupPolicyRequest) (_result *CreateBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackupPolicyResponse{}
	_body, _err := client.CreateBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a database.
//
// @param request - CreateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBResponse
func (client *Client) CreateDBWithOptions(request *CreateDBRequest, runtime *dara.RuntimeOptions) (_result *CreateDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDB"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database.
//
// @param request - CreateDBRequest
//
// @return CreateDBResponse
func (client *Client) CreateDB(request *CreateDBRequest) (_result *CreateDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBResponse{}
	_body, _err := client.CreateDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// To create a ClickHouse Enterprise Edition cluster, call the `CreateDBInstance` API.
//
// @param tmpReq - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithOptions(tmpReq *CreateDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDBInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MultiZone) {
		request.MultiZoneShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiZone, dara.String("MultiZone"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBTimeZone) {
		query["DBTimeZone"] = request.DBTimeZone
	}

	if !dara.IsNil(request.DeploySchema) {
		query["DeploySchema"] = request.DeploySchema
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.MultiZoneShrink) {
		query["MultiZone"] = request.MultiZoneShrink
	}

	if !dara.IsNil(request.NodeCount) {
		query["NodeCount"] = request.NodeCount
	}

	if !dara.IsNil(request.NodeScaleMax) {
		query["NodeScaleMax"] = request.NodeScaleMax
	}

	if !dara.IsNil(request.NodeScaleMin) {
		query["NodeScaleMin"] = request.NodeScaleMin
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.SourceDBInstanceId) {
		query["SourceDBInstanceId"] = request.SourceDBInstanceId
	}

	if !dara.IsNil(request.StorageQuota) {
		query["StorageQuota"] = request.StorageQuota
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// To create a ClickHouse Enterprise Edition cluster, call the `CreateDBInstance` API.
//
// @param request - CreateDBInstanceRequest
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (_result *CreateDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a public endpoint that provides a public address to access a cluster.
//
// @param request - CreateEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEndpointResponse
func (client *Client) CreateEndpointWithOptions(request *CreateEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.ConnectionPrefix) {
		query["ConnectionPrefix"] = request.ConnectionPrefix
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceNetType) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEndpoint"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public endpoint that provides a public address to access a cluster.
//
// @param request - CreateEndpointRequest
//
// @return CreateEndpointResponse
func (client *Client) CreateEndpoint(request *CreateEndpointRequest) (_result *CreateEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEndpointResponse{}
	_body, _err := client.CreateEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Langfuse organization.
//
// @param request - CreateLangfuseOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLangfuseOrgResponse
func (client *Client) CreateLangfuseOrgWithOptions(request *CreateLangfuseOrgRequest, runtime *dara.RuntimeOptions) (_result *CreateLangfuseOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerEmail) {
		query["OwnerEmail"] = request.OwnerEmail
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLangfuseOrg"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLangfuseOrgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Langfuse organization.
//
// @param request - CreateLangfuseOrgRequest
//
// @return CreateLangfuseOrgResponse
func (client *Client) CreateLangfuseOrg(request *CreateLangfuseOrgRequest) (_result *CreateLangfuseOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLangfuseOrgResponse{}
	_body, _err := client.CreateLangfuseOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a user to a Langfuse organization.
//
// @param request - CreateLangfuseOrgMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLangfuseOrgMemberResponse
func (client *Client) CreateLangfuseOrgMemberWithOptions(request *CreateLangfuseOrgMemberRequest, runtime *dara.RuntimeOptions) (_result *CreateLangfuseOrgMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLangfuseOrgMember"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLangfuseOrgMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a Langfuse organization.
//
// @param request - CreateLangfuseOrgMemberRequest
//
// @return CreateLangfuseOrgMemberResponse
func (client *Client) CreateLangfuseOrgMember(request *CreateLangfuseOrgMemberRequest) (_result *CreateLangfuseOrgMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLangfuseOrgMemberResponse{}
	_body, _err := client.CreateLangfuseOrgMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Langfuse project.
//
// @param request - CreateLangfuseProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLangfuseProjectResponse
func (client *Client) CreateLangfuseProjectWithOptions(request *CreateLangfuseProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateLangfuseProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLangfuseProject"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLangfuseProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Langfuse project.
//
// @param request - CreateLangfuseProjectRequest
//
// @return CreateLangfuseProjectResponse
func (client *Client) CreateLangfuseProject(request *CreateLangfuseProjectRequest) (_result *CreateLangfuseProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLangfuseProjectResponse{}
	_body, _err := client.CreateLangfuseProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Langfuse user.
//
// @param request - CreateLangfuseUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLangfuseUserResponse
func (client *Client) CreateLangfuseUserWithOptions(request *CreateLangfuseUserRequest, runtime *dara.RuntimeOptions) (_result *CreateLangfuseUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OrgRole) {
		query["OrgRole"] = request.OrgRole
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLangfuseUser"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLangfuseUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Langfuse user.
//
// @param request - CreateLangfuseUserRequest
//
// @return CreateLangfuseUserResponse
func (client *Client) CreateLangfuseUser(request *CreateLangfuseUserRequest) (_result *CreateLangfuseUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLangfuseUserResponse{}
	_body, _err := client.CreateLangfuseUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a whitelist template.
//
// @param request - CreateWhitelistTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWhitelistTemplateResponse
func (client *Client) CreateWhitelistTemplateWithOptions(request *CreateWhitelistTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateWhitelistTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SecurityIPList) {
		body["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWhitelistTemplate"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWhitelistTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a whitelist template.
//
// @param request - CreateWhitelistTemplateRequest
//
// @return CreateWhitelistTemplateResponse
func (client *Client) CreateWhitelistTemplate(request *CreateWhitelistTemplateRequest) (_result *CreateWhitelistTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWhitelistTemplateResponse{}
	_body, _err := client.CreateWhitelistTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database account from an ApsaraDB for ClickHouse cluster.
//
// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database account from an ApsaraDB for ClickHouse cluster.
//
// @param request - DeleteAccountRequest
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the backup policy for a ClickHouse Enterprise Edition cluster.
//
// @param request - DeleteBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupPolicyResponse
func (client *Client) DeleteBackupPolicyWithOptions(request *DeleteBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackupPolicy"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the backup policy for a ClickHouse Enterprise Edition cluster.
//
// @param request - DeleteBackupPolicyRequest
//
// @return DeleteBackupPolicyResponse
func (client *Client) DeleteBackupPolicy(request *DeleteBackupPolicyRequest) (_result *DeleteBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupPolicyResponse{}
	_body, _err := client.DeleteBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call this operation to delete a database.
//
// @param request - DeleteDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBResponse
func (client *Client) DeleteDBWithOptions(request *DeleteDBRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDB"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call this operation to delete a database.
//
// @param request - DeleteDBRequest
//
// @return DeleteDBResponse
func (client *Client) DeleteDB(request *DeleteDBRequest) (_result *DeleteDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBResponse{}
	_body, _err := client.DeleteDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DeleteDBInstance operation to release an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - DeleteDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceResponse
func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteDBInstance operation to release an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - DeleteDBInstanceRequest
//
// @return DeleteDBInstanceResponse
func (client *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (_result *DeleteDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DeleteDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a public endpoint.
//
// @param request - DeleteEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEndpointResponse
func (client *Client) DeleteEndpointWithOptions(request *DeleteEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceNetType) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEndpoint"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a public endpoint.
//
// @param request - DeleteEndpointRequest
//
// @return DeleteEndpointResponse
func (client *Client) DeleteEndpoint(request *DeleteEndpointRequest) (_result *DeleteEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEndpointResponse{}
	_body, _err := client.DeleteEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a Langfuse instance.
//
// @param request - DeleteLangfuseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLangfuseInstanceResponse
func (client *Client) DeleteLangfuseInstanceWithOptions(request *DeleteLangfuseInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteLangfuseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLangfuseInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLangfuseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a Langfuse instance.
//
// @param request - DeleteLangfuseInstanceRequest
//
// @return DeleteLangfuseInstanceResponse
func (client *Client) DeleteLangfuseInstance(request *DeleteLangfuseInstanceRequest) (_result *DeleteLangfuseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLangfuseInstanceResponse{}
	_body, _err := client.DeleteLangfuseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Langfuse organization.
//
// @param request - DeleteLangfuseOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLangfuseOrgResponse
func (client *Client) DeleteLangfuseOrgWithOptions(request *DeleteLangfuseOrgRequest, runtime *dara.RuntimeOptions) (_result *DeleteLangfuseOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLangfuseOrg"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLangfuseOrgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Langfuse organization.
//
// @param request - DeleteLangfuseOrgRequest
//
// @return DeleteLangfuseOrgResponse
func (client *Client) DeleteLangfuseOrg(request *DeleteLangfuseOrgRequest) (_result *DeleteLangfuseOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLangfuseOrgResponse{}
	_body, _err := client.DeleteLangfuseOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a user from a Langfuse organization.
//
// @param request - DeleteLangfuseOrgMembershipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLangfuseOrgMembershipResponse
func (client *Client) DeleteLangfuseOrgMembershipWithOptions(request *DeleteLangfuseOrgMembershipRequest, runtime *dara.RuntimeOptions) (_result *DeleteLangfuseOrgMembershipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLangfuseOrgMembership"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLangfuseOrgMembershipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a user from a Langfuse organization.
//
// @param request - DeleteLangfuseOrgMembershipRequest
//
// @return DeleteLangfuseOrgMembershipResponse
func (client *Client) DeleteLangfuseOrgMembership(request *DeleteLangfuseOrgMembershipRequest) (_result *DeleteLangfuseOrgMembershipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLangfuseOrgMembershipResponse{}
	_body, _err := client.DeleteLangfuseOrgMembershipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Langfuse project.
//
// @param request - DeleteLangfuseProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLangfuseProjectResponse
func (client *Client) DeleteLangfuseProjectWithOptions(request *DeleteLangfuseProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteLangfuseProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLangfuseProject"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLangfuseProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Langfuse project.
//
// @param request - DeleteLangfuseProjectRequest
//
// @return DeleteLangfuseProjectResponse
func (client *Client) DeleteLangfuseProject(request *DeleteLangfuseProjectRequest) (_result *DeleteLangfuseProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLangfuseProjectResponse{}
	_body, _err := client.DeleteLangfuseProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Langfuse user.
//
// @param request - DeleteLangfuseUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLangfuseUserResponse
func (client *Client) DeleteLangfuseUserWithOptions(request *DeleteLangfuseUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteLangfuseUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLangfuseUser"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLangfuseUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Langfuse user.
//
// @param request - DeleteLangfuseUserRequest
//
// @return DeleteLangfuseUserResponse
func (client *Client) DeleteLangfuseUser(request *DeleteLangfuseUserRequest) (_result *DeleteLangfuseUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLangfuseUserResponse{}
	_body, _err := client.DeleteLangfuseUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a whitelist template.
//
// @param request - DeleteWhitelistTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWhitelistTemplateResponse
func (client *Client) DeleteWhitelistTemplateWithOptions(request *DeleteWhitelistTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteWhitelistTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWhitelistTemplate"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWhitelistTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a whitelist template.
//
// @param request - DeleteWhitelistTemplateRequest
//
// @return DeleteWhitelistTemplateResponse
func (client *Client) DeleteWhitelistTemplate(request *DeleteWhitelistTemplateRequest) (_result *DeleteWhitelistTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWhitelistTemplateResponse{}
	_body, _err := client.DeleteWhitelistTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DescribeAccountAuthority operation to query the permissions of an account.
//
// @param request - DescribeAccountAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountAuthorityResponse
func (client *Client) DescribeAccountAuthorityWithOptions(request *DescribeAccountAuthorityRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountAuthorityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountAuthority"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeAccountAuthority operation to query the permissions of an account.
//
// @param request - DescribeAccountAuthorityRequest
//
// @return DescribeAccountAuthorityResponse
func (client *Client) DescribeAccountAuthority(request *DescribeAccountAuthorityRequest) (_result *DescribeAccountAuthorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountAuthorityResponse{}
	_body, _err := client.DescribeAccountAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of database accounts for a specified ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccounts"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of database accounts for a specified ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeAccountsRequest
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup policy of a ClickHouse Enterprise Edition cluster.
//
// @param request - DescribeBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup policy of a ClickHouse Enterprise Edition cluster.
//
// @param request - DescribeBackupPolicyRequest
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup sets of a ClickHouse Enterprise Edition cluster.
//
// @param request - DescribeBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeBackups"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup sets of a ClickHouse Enterprise Edition cluster.
//
// @param request - DescribeBackupsRequest
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceAttributeResponse
func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceAttribute"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an ApsaraDB for ClickHouse cluster.
//
// @param request - DescribeDBInstanceAttributeRequest
//
// @return DescribeDBInstanceAttributeResponse
func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DescribeDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries the parameter configuration of an instance
//
// @param request - DescribeDBInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceConfigResponse
func (client *Client) DescribeDBInstanceConfigWithOptions(request *DescribeDBInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceConfig"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the parameter configuration of an instance
//
// @param request - DescribeDBInstanceConfigRequest
//
// @return DescribeDBInstanceConfigResponse
func (client *Client) DescribeDBInstanceConfig(request *DescribeDBInstanceConfigRequest) (_result *DescribeDBInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceConfigResponse{}
	_body, _err := client.DescribeDBInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries the parameter configuration records of an instance
//
// @param request - DescribeDBInstanceConfigChangeLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceConfigChangeLogResponse
func (client *Client) DescribeDBInstanceConfigChangeLogWithOptions(request *DescribeDBInstanceConfigChangeLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceConfigChangeLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceConfigChangeLog"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceConfigChangeLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the parameter configuration records of an instance
//
// @param request - DescribeDBInstanceConfigChangeLogRequest
//
// @return DescribeDBInstanceConfigChangeLogResponse
func (client *Client) DescribeDBInstanceConfigChangeLog(request *DescribeDBInstanceConfigChangeLogRequest) (_result *DescribeDBInstanceConfigChangeLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceConfigChangeLogResponse{}
	_body, _err := client.DescribeDBInstanceConfigChangeLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data structure of a database or table by calling the DescribeDBInstanceDataSources operation.
//
// @param request - DescribeDBInstanceDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceDataSourcesResponse
func (client *Client) DescribeDBInstanceDataSourcesWithOptions(request *DescribeDBInstanceDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceDataSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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
		Action:      dara.String("DescribeDBInstanceDataSources"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data structure of a database or table by calling the DescribeDBInstanceDataSources operation.
//
// @param request - DescribeDBInstanceDataSourcesRequest
//
// @return DescribeDBInstanceDataSourcesResponse
func (client *Client) DescribeDBInstanceDataSources(request *DescribeDBInstanceDataSourcesRequest) (_result *DescribeDBInstanceDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceDataSourcesResponse{}
	_body, _err := client.DescribeDBInstanceDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of clusters.
//
// @param request - DescribeDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesResponse
func (client *Client) DescribeDBInstancesWithOptions(request *DescribeDBInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceIds) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !dara.IsNil(request.DBInstanceStatus) {
		query["DBInstanceStatus"] = request.DBInstanceStatus
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.VpcIds) {
		query["VpcIds"] = request.VpcIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstances"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of clusters.
//
// @param request - DescribeDBInstancesRequest
//
// @return DescribeDBInstancesResponse
func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (_result *DescribeDBInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DescribeDBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DescribeEndpoints operation to query the endpoints of a cluster.
//
// @param request - DescribeEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndpointsResponse
func (client *Client) DescribeEndpointsWithOptions(request *DescribeEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEndpoints"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeEndpoints operation to query the endpoints of a cluster.
//
// @param request - DescribeEndpointsRequest
//
// @return DescribeEndpointsResponse
func (client *Client) DescribeEndpoints(request *DescribeEndpointsRequest) (_result *DescribeEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEndpointsResponse{}
	_body, _err := client.DescribeEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access endpoint of a Langfuse instance.
//
// @param request - DescribeLangfuseEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLangfuseEndpointsResponse
func (client *Client) DescribeLangfuseEndpointsWithOptions(request *DescribeLangfuseEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLangfuseEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLangfuseEndpoints"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLangfuseEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access endpoint of a Langfuse instance.
//
// @param request - DescribeLangfuseEndpointsRequest
//
// @return DescribeLangfuseEndpointsResponse
func (client *Client) DescribeLangfuseEndpoints(request *DescribeLangfuseEndpointsRequest) (_result *DescribeLangfuseEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLangfuseEndpointsResponse{}
	_body, _err := client.DescribeLangfuseEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the member information of a Langfuse organization.
//
// @param request - DescribeLangfuseOrgMembershipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLangfuseOrgMembershipsResponse
func (client *Client) DescribeLangfuseOrgMembershipsWithOptions(request *DescribeLangfuseOrgMembershipsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLangfuseOrgMembershipsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
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
		Action:      dara.String("DescribeLangfuseOrgMemberships"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLangfuseOrgMembershipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the member information of a Langfuse organization.
//
// @param request - DescribeLangfuseOrgMembershipsRequest
//
// @return DescribeLangfuseOrgMembershipsResponse
func (client *Client) DescribeLangfuseOrgMemberships(request *DescribeLangfuseOrgMembershipsRequest) (_result *DescribeLangfuseOrgMembershipsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLangfuseOrgMembershipsResponse{}
	_body, _err := client.DescribeLangfuseOrgMembershipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of organizations for a Langfuse instance.
//
// @param request - DescribeLangfuseOrgsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLangfuseOrgsResponse
func (client *Client) DescribeLangfuseOrgsWithOptions(request *DescribeLangfuseOrgsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLangfuseOrgsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("DescribeLangfuseOrgs"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLangfuseOrgsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of organizations for a Langfuse instance.
//
// @param request - DescribeLangfuseOrgsRequest
//
// @return DescribeLangfuseOrgsResponse
func (client *Client) DescribeLangfuseOrgs(request *DescribeLangfuseOrgsRequest) (_result *DescribeLangfuseOrgsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLangfuseOrgsResponse{}
	_body, _err := client.DescribeLangfuseOrgsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user roles in a Langfuse project.
//
// @param request - DescribeLangfuseProjectMembershipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLangfuseProjectMembershipsResponse
func (client *Client) DescribeLangfuseProjectMembershipsWithOptions(request *DescribeLangfuseProjectMembershipsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLangfuseProjectMembershipsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLangfuseProjectMemberships"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLangfuseProjectMembershipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user roles in a Langfuse project.
//
// @param request - DescribeLangfuseProjectMembershipsRequest
//
// @return DescribeLangfuseProjectMembershipsResponse
func (client *Client) DescribeLangfuseProjectMemberships(request *DescribeLangfuseProjectMembershipsRequest) (_result *DescribeLangfuseProjectMembershipsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLangfuseProjectMembershipsResponse{}
	_body, _err := client.DescribeLangfuseProjectMembershipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Langfuse projects.
//
// @param request - DescribeLangfuseProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLangfuseProjectsResponse
func (client *Client) DescribeLangfuseProjectsWithOptions(request *DescribeLangfuseProjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLangfuseProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLangfuseProjects"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLangfuseProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Langfuse projects.
//
// @param request - DescribeLangfuseProjectsRequest
//
// @return DescribeLangfuseProjectsResponse
func (client *Client) DescribeLangfuseProjects(request *DescribeLangfuseProjectsRequest) (_result *DescribeLangfuseProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLangfuseProjectsResponse{}
	_body, _err := client.DescribeLangfuseProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Langfuse whitelist.
//
// @param request - DescribeLangfuseSecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLangfuseSecurityIPListResponse
func (client *Client) DescribeLangfuseSecurityIPListWithOptions(request *DescribeLangfuseSecurityIPListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLangfuseSecurityIPListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLangfuseSecurityIPList"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLangfuseSecurityIPListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Langfuse whitelist.
//
// @param request - DescribeLangfuseSecurityIPListRequest
//
// @return DescribeLangfuseSecurityIPListResponse
func (client *Client) DescribeLangfuseSecurityIPList(request *DescribeLangfuseSecurityIPListRequest) (_result *DescribeLangfuseSecurityIPListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLangfuseSecurityIPListResponse{}
	_body, _err := client.DescribeLangfuseSecurityIPListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Langfuse user.
//
// @param request - DescribeLangfuseUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLangfuseUserResponse
func (client *Client) DescribeLangfuseUserWithOptions(request *DescribeLangfuseUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeLangfuseUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLangfuseUser"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLangfuseUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Langfuse user.
//
// @param request - DescribeLangfuseUserRequest
//
// @return DescribeLangfuseUserResponse
func (client *Client) DescribeLangfuseUser(request *DescribeLangfuseUserRequest) (_result *DescribeLangfuseUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLangfuseUserResponse{}
	_body, _err := client.DescribeLangfuseUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Langfuse users.
//
// @param request - DescribeLangfuseUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLangfuseUsersResponse
func (client *Client) DescribeLangfuseUsersWithOptions(request *DescribeLangfuseUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeLangfuseUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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
		Action:      dara.String("DescribeLangfuseUsers"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLangfuseUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Langfuse users.
//
// @param request - DescribeLangfuseUsersRequest
//
// @return DescribeLangfuseUsersResponse
func (client *Client) DescribeLangfuseUsers(request *DescribeLangfuseUsersRequest) (_result *DescribeLangfuseUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLangfuseUsersResponse{}
	_body, _err := client.DescribeLangfuseUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DescribeProcessList operation to view queries that are currently running.
//
// Description:
//
// > This operation supports only community-compatible edition clusters created after December 01, 2021.
//
// @param request - DescribeProcessListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProcessListResponse
func (client *Client) DescribeProcessListWithOptions(request *DescribeProcessListRequest, runtime *dara.RuntimeOptions) (_result *DescribeProcessListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.InitialQueryId) {
		query["InitialQueryId"] = request.InitialQueryId
	}

	if !dara.IsNil(request.InitialUser) {
		query["InitialUser"] = request.InitialUser
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryDurationMs) {
		query["QueryDurationMs"] = request.QueryDurationMs
	}

	if !dara.IsNil(request.QueryOrder) {
		query["QueryOrder"] = request.QueryOrder
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProcessList"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProcessListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeProcessList operation to view queries that are currently running.
//
// Description:
//
// > This operation supports only community-compatible edition clusters created after December 01, 2021.
//
// @param request - DescribeProcessListRequest
//
// @return DescribeProcessListResponse
func (client *Client) DescribeProcessList(request *DescribeProcessListRequest) (_result *DescribeProcessListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProcessListResponse{}
	_body, _err := client.DescribeProcessListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns information about the available regions and zones for ApsaraDB for ClickHouse.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns information about the available regions and zones for ApsaraDB for ClickHouse.
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelist of a cluster by calling the DescribeSecurityIPList operation.
//
// @param request - DescribeSecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityIPListResponse
func (client *Client) DescribeSecurityIPListWithOptions(request *DescribeSecurityIPListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityIPListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityIPList"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityIPListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelist of a cluster by calling the DescribeSecurityIPList operation.
//
// @param request - DescribeSecurityIPListRequest
//
// @return DescribeSecurityIPListResponse
func (client *Client) DescribeSecurityIPList(request *DescribeSecurityIPListRequest) (_result *DescribeSecurityIPListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityIPListResponse{}
	_body, _err := client.DescribeSecurityIPListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries slow query log details by calling the DescribeSlowLogRecords operation.
//
// @param request - DescribeSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryDurationMs) {
		query["QueryDurationMs"] = request.QueryDurationMs
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
		Action:      dara.String("DescribeSlowLogRecords"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries slow query log details by calling the DescribeSlowLogRecords operation.
//
// @param request - DescribeSlowLogRecordsRequest
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries slow log trends by calling the DescribeSlowLogTrend operation.
//
// @param request - DescribeSlowLogTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogTrendResponse
func (client *Client) DescribeSlowLogTrendWithOptions(request *DescribeSlowLogTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.QueryDurationMs) {
		query["QueryDurationMs"] = request.QueryDurationMs
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
		Action:      dara.String("DescribeSlowLogTrend"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries slow log trends by calling the DescribeSlowLogTrend operation.
//
// @param request - DescribeSlowLogTrendRequest
//
// @return DescribeSlowLogTrendResponse
func (client *Client) DescribeSlowLogTrend(request *DescribeSlowLogTrendRequest) (_result *DescribeSlowLogTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlowLogTrendResponse{}
	_body, _err := client.DescribeSlowLogTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociate a whitelist template from an instance.
//
// @param request - DetachWhitelistTemplateToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachWhitelistTemplateToInstanceResponse
func (client *Client) DetachWhitelistTemplateToInstanceWithOptions(request *DetachWhitelistTemplateToInstanceRequest, runtime *dara.RuntimeOptions) (_result *DetachWhitelistTemplateToInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachWhitelistTemplateToInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachWhitelistTemplateToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociate a whitelist template from an instance.
//
// @param request - DetachWhitelistTemplateToInstanceRequest
//
// @return DetachWhitelistTemplateToInstanceResponse
func (client *Client) DetachWhitelistTemplateToInstance(request *DetachWhitelistTemplateToInstanceRequest) (_result *DetachWhitelistTemplateToInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachWhitelistTemplateToInstanceResponse{}
	_body, _err := client.DetachWhitelistTemplateToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a whitelist template.
//
// @param request - GetWhitelistTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWhitelistTemplateResponse
func (client *Client) GetWhitelistTemplateWithOptions(request *GetWhitelistTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetWhitelistTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWhitelistTemplate"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWhitelistTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a whitelist template.
//
// @param request - GetWhitelistTemplateRequest
//
// @return GetWhitelistTemplateResponse
func (client *Client) GetWhitelistTemplate(request *GetWhitelistTemplateRequest) (_result *GetWhitelistTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWhitelistTemplateResponse{}
	_body, _err := client.GetWhitelistTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an ongoing task.
//
// Description:
//
// Make sure that you are familiar with the billing methods and [pricing](https://help.aliyun.com/document_detail/167450.html) of ApsaraDB for ClickHouse before you call this operation.
//
// @param request - KillProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillProcessResponse
func (client *Client) KillProcessWithOptions(request *KillProcessRequest, runtime *dara.RuntimeOptions) (_result *KillProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.InitialQueryId) {
		query["InitialQueryId"] = request.InitialQueryId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KillProcess"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KillProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an ongoing task.
//
// Description:
//
// Make sure that you are familiar with the billing methods and [pricing](https://help.aliyun.com/document_detail/167450.html) of ApsaraDB for ClickHouse before you call this operation.
//
// @param request - KillProcessRequest
//
// @return KillProcessResponse
func (client *Client) KillProcess(request *KillProcessRequest) (_result *KillProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KillProcessResponse{}
	_body, _err := client.KillProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the available time zones for ClickHouse.
//
// @param request - ListClickHouseDBTimezonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClickHouseDBTimezonesResponse
func (client *Client) ListClickHouseDBTimezonesWithOptions(request *ListClickHouseDBTimezonesRequest, runtime *dara.RuntimeOptions) (_result *ListClickHouseDBTimezonesResponse, _err error) {
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
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClickHouseDBTimezones"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClickHouseDBTimezonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the available time zones for ClickHouse.
//
// @param request - ListClickHouseDBTimezonesRequest
//
// @return ListClickHouseDBTimezonesResponse
func (client *Client) ListClickHouseDBTimezones(request *ListClickHouseDBTimezonesRequest) (_result *ListClickHouseDBTimezonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClickHouseDBTimezonesResponse{}
	_body, _err := client.ListClickHouseDBTimezonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the IP address whitelist templates linked to an instance.
//
// @param request - ListInstanceLinkedWhitelistTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceLinkedWhitelistTemplatesResponse
func (client *Client) ListInstanceLinkedWhitelistTemplatesWithOptions(request *ListInstanceLinkedWhitelistTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceLinkedWhitelistTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceLinkedWhitelistTemplates"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceLinkedWhitelistTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the IP address whitelist templates linked to an instance.
//
// @param request - ListInstanceLinkedWhitelistTemplatesRequest
//
// @return ListInstanceLinkedWhitelistTemplatesResponse
func (client *Client) ListInstanceLinkedWhitelistTemplates(request *ListInstanceLinkedWhitelistTemplatesRequest) (_result *ListInstanceLinkedWhitelistTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceLinkedWhitelistTemplatesResponse{}
	_body, _err := client.ListInstanceLinkedWhitelistTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the whitelist templates associated with an instance.
//
// @param request - ListWhitelistTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWhitelistTemplatesResponse
func (client *Client) ListWhitelistTemplatesWithOptions(request *ListWhitelistTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListWhitelistTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWhitelistTemplates"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWhitelistTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the whitelist templates associated with an instance.
//
// @param request - ListWhitelistTemplatesRequest
//
// @return ListWhitelistTemplatesResponse
func (client *Client) ListWhitelistTemplates(request *ListWhitelistTemplatesRequest) (_result *ListWhitelistTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWhitelistTemplatesResponse{}
	_body, _err := client.ListWhitelistTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a database account.
//
// @param tmpReq - ModifyAccountAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountAuthorityResponse
func (client *Client) ModifyAccountAuthorityWithOptions(tmpReq *ModifyAccountAuthorityRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountAuthorityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAccountAuthorityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DmlAuthSetting) {
		request.DmlAuthSettingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DmlAuthSetting, dara.String("DmlAuthSetting"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DmlAuthSettingShrink) {
		query["DmlAuthSetting"] = request.DmlAuthSettingShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountAuthority"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a database account.
//
// @param request - ModifyAccountAuthorityRequest
//
// @return ModifyAccountAuthorityResponse
func (client *Client) ModifyAccountAuthority(request *ModifyAccountAuthorityRequest) (_result *ModifyAccountAuthorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountAuthorityResponse{}
	_body, _err := client.ModifyAccountAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a database account.
//
// @param request - ModifyAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a database account.
//
// @param request - ModifyAccountDescriptionRequest
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the backup policy of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >ApsaraDB for ClickHouse clusters of versions 20.3, 20.8, and 21.8 support data backup.
//
// @param request - ModifyBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPolicyResponse
func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.PreferredBackupPeriod) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the backup policy of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// >ApsaraDB for ClickHouse clusters of versions 20.3, 20.8, and 21.8 support data backup.
//
// @param request - ModifyBackupPolicyRequest
//
// @return ModifyBackupPolicyResponse
func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ModifyDBInstanceAttribute operation to modify the configuration of a cluster.
//
// @param request - ModifyDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceAttributeResponse
func (client *Client) ModifyDBInstanceAttributeWithOptions(request *ModifyDBInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttributeType) {
		query["AttributeType"] = request.AttributeType
	}

	if !dara.IsNil(request.AttributeValue) {
		query["AttributeValue"] = request.AttributeValue
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceAttribute"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyDBInstanceAttribute operation to modify the configuration of a cluster.
//
// @param request - ModifyDBInstanceAttributeRequest
//
// @return ModifyDBInstanceAttributeResponse
func (client *Client) ModifyDBInstanceAttribute(request *ModifyDBInstanceAttributeRequest) (_result *ModifyDBInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceAttributeResponse{}
	_body, _err := client.ModifyDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the elastic configuration of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/167450.html) of ApsaraDB for ClickHouse.
//
// @param tmpReq - ModifyDBInstanceClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceClassResponse
func (client *Client) ModifyDBInstanceClassWithOptions(tmpReq *ModifyDBInstanceClassRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDBInstanceClassShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AutoScaleConfig) {
		request.AutoScaleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AutoScaleConfig, dara.String("AutoScaleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoScaleConfigShrink) {
		query["AutoScaleConfig"] = request.AutoScaleConfigShrink
	}

	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeCount) {
		query["NodeCount"] = request.NodeCount
	}

	if !dara.IsNil(request.NodeScaleMax) {
		query["NodeScaleMax"] = request.NodeScaleMax
	}

	if !dara.IsNil(request.NodeScaleMin) {
		query["NodeScaleMin"] = request.NodeScaleMin
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.StorageQuota) {
		query["StorageQuota"] = request.StorageQuota
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceClass"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the elastic configuration of an ApsaraDB for ClickHouse cluster.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/167450.html) of ApsaraDB for ClickHouse.
//
// @param request - ModifyDBInstanceClassRequest
//
// @return ModifyDBInstanceClassResponse
func (client *Client) ModifyDBInstanceClass(request *ModifyDBInstanceClassRequest) (_result *ModifyDBInstanceClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceClassResponse{}
	_body, _err := client.ModifyDBInstanceClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the parameter settings for a DB instance.
//
// @param request - ModifyDBInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConfigResponse
func (client *Client) ModifyDBInstanceConfigWithOptions(request *ModifyDBInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConfig"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameter settings for a DB instance.
//
// @param request - ModifyDBInstanceConfigRequest
//
// @return ModifyDBInstanceConfigResponse
func (client *Client) ModifyDBInstanceConfig(request *ModifyDBInstanceConfigRequest) (_result *ModifyDBInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceConfigResponse{}
	_body, _err := client.ModifyDBInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the ModifyDBInstanceConnectionString operation to modify the endpoint of a cluster.
//
// @param request - ModifyDBInstanceConnectionStringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConnectionStringResponse
func (client *Client) ModifyDBInstanceConnectionStringWithOptions(request *ModifyDBInstanceConnectionStringRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComputingGroupId) {
		query["ComputingGroupId"] = request.ComputingGroupId
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceNetType) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !dara.IsNil(request.DisablePorts) {
		query["DisablePorts"] = request.DisablePorts
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConnectionString"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyDBInstanceConnectionString operation to modify the endpoint of a cluster.
//
// @param request - ModifyDBInstanceConnectionStringRequest
//
// @return ModifyDBInstanceConnectionStringResponse
func (client *Client) ModifyDBInstanceConnectionString(request *ModifyDBInstanceConnectionStringRequest) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.ModifyDBInstanceConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a user in a Langfuse organization.
//
// @param request - ModifyLangfuseOrgMembershipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLangfuseOrgMembershipResponse
func (client *Client) ModifyLangfuseOrgMembershipWithOptions(request *ModifyLangfuseOrgMembershipRequest, runtime *dara.RuntimeOptions) (_result *ModifyLangfuseOrgMembershipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLangfuseOrgMembership"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLangfuseOrgMembershipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a user in a Langfuse organization.
//
// @param request - ModifyLangfuseOrgMembershipRequest
//
// @return ModifyLangfuseOrgMembershipResponse
func (client *Client) ModifyLangfuseOrgMembership(request *ModifyLangfuseOrgMembershipRequest) (_result *ModifyLangfuseOrgMembershipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLangfuseOrgMembershipResponse{}
	_body, _err := client.ModifyLangfuseOrgMembershipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a user in a Langfuse project.
//
// @param request - ModifyLangfuseProjectMembershipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLangfuseProjectMembershipResponse
func (client *Client) ModifyLangfuseProjectMembershipWithOptions(request *ModifyLangfuseProjectMembershipRequest, runtime *dara.RuntimeOptions) (_result *ModifyLangfuseProjectMembershipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLangfuseProjectMembership"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLangfuseProjectMembershipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a user in a Langfuse project.
//
// @param request - ModifyLangfuseProjectMembershipRequest
//
// @return ModifyLangfuseProjectMembershipResponse
func (client *Client) ModifyLangfuseProjectMembership(request *ModifyLangfuseProjectMembershipRequest) (_result *ModifyLangfuseProjectMembershipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLangfuseProjectMembershipResponse{}
	_body, _err := client.ModifyLangfuseProjectMembershipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the whitelist of a Langfuse instance.
//
// @param request - ModifyLangfuseSecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLangfuseSecurityIPListResponse
func (client *Client) ModifyLangfuseSecurityIPListWithOptions(request *ModifyLangfuseSecurityIPListRequest, runtime *dara.RuntimeOptions) (_result *ModifyLangfuseSecurityIPListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLangfuseSecurityIPList"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLangfuseSecurityIPListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the whitelist of a Langfuse instance.
//
// @param request - ModifyLangfuseSecurityIPListRequest
//
// @return ModifyLangfuseSecurityIPListResponse
func (client *Client) ModifyLangfuseSecurityIPList(request *ModifyLangfuseSecurityIPListRequest) (_result *ModifyLangfuseSecurityIPListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLangfuseSecurityIPListResponse{}
	_body, _err := client.ModifyLangfuseSecurityIPListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the whitelist settings of a cluster.
//
// @param request - ModifySecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityIPListResponse
func (client *Client) ModifySecurityIPListWithOptions(request *ModifySecurityIPListRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityIPListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityIPList"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityIPListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the whitelist settings of a cluster.
//
// @param request - ModifySecurityIPListRequest
//
// @return ModifySecurityIPListResponse
func (client *Client) ModifySecurityIPList(request *ModifySecurityIPListRequest) (_result *ModifySecurityIPListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySecurityIPListResponse{}
	_body, _err := client.ModifySecurityIPListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of a database account for an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - ResetAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of a database account for an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - ResetAccountPasswordRequest
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of a Langfuse user.
//
// @param request - ResetLangfuseUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetLangfuseUserPasswordResponse
func (client *Client) ResetLangfuseUserPasswordWithOptions(request *ResetLangfuseUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetLangfuseUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.NewPassword) {
		query["NewPassword"] = request.NewPassword
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetLangfuseUserPassword"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetLangfuseUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of a Langfuse user.
//
// @param request - ResetLangfuseUserPasswordRequest
//
// @return ResetLangfuseUserPasswordResponse
func (client *Client) ResetLangfuseUserPassword(request *ResetLangfuseUserPasswordRequest) (_result *ResetLangfuseUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetLangfuseUserPasswordResponse{}
	_body, _err := client.ResetLangfuseUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an ApsaraDB for ClickHouse cluster.
//
// @param request - RestartDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBInstanceResponse
func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDBInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an ApsaraDB for ClickHouse cluster.
//
// @param request - RestartDBInstanceRequest
//
// @return RestartDBInstanceResponse
func (client *Client) RestartDBInstance(request *RestartDBInstanceRequest) (_result *RestartDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.RestartDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the StartDBInstance operation to start an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - StartDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDBInstanceResponse
func (client *Client) StartDBInstanceWithOptions(request *StartDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDBInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the StartDBInstance operation to start an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - StartDBInstanceRequest
//
// @return StartDBInstanceResponse
func (client *Client) StartDBInstance(request *StartDBInstanceRequest) (_result *StartDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDBInstanceResponse{}
	_body, _err := client.StartDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the StopDBInstance operation to pause an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - StopDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDBInstanceResponse
func (client *Client) StopDBInstanceWithOptions(request *StopDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDBInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the StopDBInstance operation to pause an ApsaraDB for ClickHouse Enterprise Edition cluster.
//
// @param request - StopDBInstanceRequest
//
// @return StopDBInstanceResponse
func (client *Client) StopDBInstance(request *StopDBInstanceRequest) (_result *StopDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDBInstanceResponse{}
	_body, _err := client.StopDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a whitelist template.
//
// @param request - UpdateWhitelistTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWhitelistTemplateResponse
func (client *Client) UpdateWhitelistTemplateWithOptions(request *UpdateWhitelistTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateWhitelistTemplateResponse, _err error) {
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

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWhitelistTemplate"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWhitelistTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a whitelist template.
//
// @param request - UpdateWhitelistTemplateRequest
//
// @return UpdateWhitelistTemplateResponse
func (client *Client) UpdateWhitelistTemplate(request *UpdateWhitelistTemplateRequest) (_result *UpdateWhitelistTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWhitelistTemplateResponse{}
	_body, _err := client.UpdateWhitelistTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the minor engine version of a specified ApsaraDB for ClickHouse cluster.
//
// @param request - UpgradeMinorVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeMinorVersionResponse
func (client *Client) UpgradeMinorVersionWithOptions(request *UpgradeMinorVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeMinorVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !dara.IsNil(request.TargetMinorVersion) {
		query["TargetMinorVersion"] = request.TargetMinorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeMinorVersion"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeMinorVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the minor engine version of a specified ApsaraDB for ClickHouse cluster.
//
// @param request - UpgradeMinorVersionRequest
//
// @return UpgradeMinorVersionResponse
func (client *Client) UpgradeMinorVersion(request *UpgradeMinorVersionRequest) (_result *UpgradeMinorVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeMinorVersionResponse{}
	_body, _err := client.UpgradeMinorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
