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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              dara.String("polardbx.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("polardbx.aliyuncs.com"),
		"ap-south-1":                  dara.String("polardbx.aliyuncs.com"),
		"ap-southeast-2":              dara.String("polardbx.aliyuncs.com"),
		"ap-southeast-3":              dara.String("polardbx.aliyuncs.com"),
		"ap-southeast-5":              dara.String("polardbx.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("polardbx.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("polardbx.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("polardbx.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("polardbx.aliyuncs.com"),
		"cn-edge-1":                   dara.String("polardbx.aliyuncs.com"),
		"cn-fujian":                   dara.String("polardbx.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("polardbx.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("polardbx.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("polardbx.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("polardbx.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("polardbx.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("polardbx.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("polardbx.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("polardbx.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("polardbx.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("polardbx.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("polardbx.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("polardbx.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("polardbx.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("polardbx.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("polardbx.aliyuncs.com"),
		"cn-wuhan":                    dara.String("polardbx.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("polardbx.aliyuncs.com"),
		"cn-yushanfang":               dara.String("polardbx.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("polardbx.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("polardbx.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("polardbx.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("polardbx.aliyuncs.com"),
		"eu-central-1":                dara.String("polardbx.aliyuncs.com"),
		"eu-west-1":                   dara.String("polardbx.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("polardbx.aliyuncs.com"),
		"me-east-1":                   dara.String("polardbx.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("polardbx.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("polardbx"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AlignStoragePrimaryAzoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlignStoragePrimaryAzoneResponse
func (client *Client) AlignStoragePrimaryAzoneWithOptions(request *AlignStoragePrimaryAzoneRequest, runtime *dara.RuntimeOptions) (_result *AlignStoragePrimaryAzoneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StorageInstanceName) {
		query["StorageInstanceName"] = request.StorageInstanceName
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlignStoragePrimaryAzone"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AlignStoragePrimaryAzoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AlignStoragePrimaryAzoneRequest
//
// @return AlignStoragePrimaryAzoneResponse
func (client *Client) AlignStoragePrimaryAzone(request *AlignStoragePrimaryAzoneRequest) (_result *AlignStoragePrimaryAzoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AlignStoragePrimaryAzoneResponse{}
	_body, _err := client.AlignStoragePrimaryAzoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通冷存储
//
// @param request - AllocateColdDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateColdDataVolumeResponse
func (client *Client) AllocateColdDataVolumeWithOptions(request *AllocateColdDataVolumeRequest, runtime *dara.RuntimeOptions) (_result *AllocateColdDataVolumeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateColdDataVolume"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateColdDataVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通冷存储
//
// @param request - AllocateColdDataVolumeRequest
//
// @return AllocateColdDataVolumeResponse
func (client *Client) AllocateColdDataVolume(request *AllocateColdDataVolumeRequest) (_result *AllocateColdDataVolumeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateColdDataVolumeResponse{}
	_body, _err := client.AllocateColdDataVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AllocateInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateInstancePublicConnectionResponse
func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateInstancePublicConnection"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AllocateInstancePublicConnectionRequest
//
// @return AllocateInstancePublicConnectionResponse
func (client *Client) AllocateInstancePublicConnection(request *AllocateInstancePublicConnectionRequest) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.AllocateInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消主动运维任务
//
// @param request - CancelActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasksWithOptions(request *CancelActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *CancelActiveOperationTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelActiveOperationTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消主动运维任务
//
// @param request - CancelActiveOperationTasksRequest
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasks(request *CancelActiveOperationTasksRequest) (_result *CancelActiveOperationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelActiveOperationTasksResponse{}
	_body, _err := client.CancelActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例所在资源组.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2020-02-02"),
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
// 修改实例所在资源组.
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

// @param request - CheckCloudResourceAuthorizedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCloudResourceAuthorizedResponse
func (client *Client) CheckCloudResourceAuthorizedWithOptions(request *CheckCloudResourceAuthorizedRequest, runtime *dara.RuntimeOptions) (_result *CheckCloudResourceAuthorizedResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCloudResourceAuthorized"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCloudResourceAuthorizedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckCloudResourceAuthorizedRequest
//
// @return CheckCloudResourceAuthorizedResponse
func (client *Client) CheckCloudResourceAuthorized(request *CheckCloudResourceAuthorizedRequest) (_result *CheckCloudResourceAuthorizedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckCloudResourceAuthorizedResponse{}
	_body, _err := client.CheckCloudResourceAuthorizedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2020-02-02"),
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

// @param request - CreateBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupResponse
func (client *Client) CreateBackupWithOptions(request *CreateBackupRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateBackupRequest
//
// @return CreateBackupResponse
func (client *Client) CreateBackup(request *CreateBackupRequest) (_result *CreateBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackupResponse{}
	_body, _err := client.CreateBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBResponse
func (client *Client) CreateDBWithOptions(request *CreateDBRequest, runtime *dara.RuntimeOptions) (_result *CreateDBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.Charset) {
		query["Charset"] = request.Charset
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbDescription) {
		query["DbDescription"] = request.DbDescription
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	if !dara.IsNil(request.StoragePoolName) {
		query["StoragePoolName"] = request.StoragePoolName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDB"),
		Version:     dara.String("2020-02-02"),
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
// 创建实例
//
// @param tmpReq - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithOptions(tmpReq *CreateDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDBInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtraParams) {
		request.ExtraParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtraParams, dara.String("ExtraParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.CNNodeCount) {
		query["CNNodeCount"] = request.CNNodeCount
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CnClass) {
		query["CnClass"] = request.CnClass
	}

	if !dara.IsNil(request.DBNodeClass) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !dara.IsNil(request.DBNodeCount) {
		query["DBNodeCount"] = request.DBNodeCount
	}

	if !dara.IsNil(request.DNNodeCount) {
		query["DNNodeCount"] = request.DNNodeCount
	}

	if !dara.IsNil(request.DnClass) {
		query["DnClass"] = request.DnClass
	}

	if !dara.IsNil(request.DnStorageSpace) {
		query["DnStorageSpace"] = request.DnStorageSpace
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.ExtraParamsShrink) {
		query["ExtraParams"] = request.ExtraParamsShrink
	}

	if !dara.IsNil(request.IsColumnarReadDBInstance) {
		query["IsColumnarReadDBInstance"] = request.IsColumnarReadDBInstance
	}

	if !dara.IsNil(request.IsReadDBInstance) {
		query["IsReadDBInstance"] = request.IsReadDBInstance
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrimaryDBInstanceName) {
		query["PrimaryDBInstanceName"] = request.PrimaryDBInstanceName
	}

	if !dara.IsNil(request.PrimaryZone) {
		query["PrimaryZone"] = request.PrimaryZone
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecondaryZone) {
		query["SecondaryZone"] = request.SecondaryZone
	}

	if !dara.IsNil(request.Series) {
		query["Series"] = request.Series
	}

	if !dara.IsNil(request.TertiaryZone) {
		query["TertiaryZone"] = request.TertiaryZone
	}

	if !dara.IsNil(request.TopologyType) {
		query["TopologyType"] = request.TopologyType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstance"),
		Version:     dara.String("2020-02-02"),
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
// 创建实例
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

// @param request - CreateSuperAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSuperAccountResponse
func (client *Client) CreateSuperAccountWithOptions(request *CreateSuperAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateSuperAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSuperAccount"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSuperAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateSuperAccountRequest
//
// @return CreateSuperAccountResponse
func (client *Client) CreateSuperAccount(request *CreateSuperAccountRequest) (_result *CreateSuperAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSuperAccountResponse{}
	_body, _err := client.CreateSuperAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2020-02-02"),
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

// @param request - DeleteDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBResponse
func (client *Client) DeleteDBWithOptions(request *DeleteDBRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDB"),
		Version:     dara.String("2020-02-02"),
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

// @param request - DeleteDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceResponse
func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstance"),
		Version:     dara.String("2020-02-02"),
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

// @param request - DescribeAccountListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountListResponse
func (client *Client) DescribeAccountListWithOptions(request *DescribeAccountListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAccountListRequest
//
// @return DescribeAccountListResponse
func (client *Client) DescribeAccountList(request *DescribeAccountListRequest) (_result *DescribeAccountListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountListResponse{}
	_body, _err := client.DescribeAccountListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示全局运维窗口配置
//
// @param request - DescribeActiveOperationMaintainConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationMaintainConfResponse
func (client *Client) DescribeActiveOperationMaintainConfWithOptions(request *DescribeActiveOperationMaintainConfRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationMaintainConfResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationMaintainConf"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationMaintainConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示全局运维窗口配置
//
// @param request - DescribeActiveOperationMaintainConfRequest
//
// @return DescribeActiveOperationMaintainConfResponse
func (client *Client) DescribeActiveOperationMaintainConf(request *DescribeActiveOperationMaintainConfRequest) (_result *DescribeActiveOperationMaintainConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveOperationMaintainConfResponse{}
	_body, _err := client.DescribeActiveOperationMaintainConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取主动运维任务数量
//
// @param request - DescribeActiveOperationTaskCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskCountResponse
func (client *Client) DescribeActiveOperationTaskCountWithOptions(request *DescribeActiveOperationTaskCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTaskCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationTaskCount"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTaskCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主动运维任务数量
//
// @param request - DescribeActiveOperationTaskCountRequest
//
// @return DescribeActiveOperationTaskCountResponse
func (client *Client) DescribeActiveOperationTaskCount(request *DescribeActiveOperationTaskCountRequest) (_result *DescribeActiveOperationTaskCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveOperationTaskCountResponse{}
	_body, _err := client.DescribeActiveOperationTaskCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取待执行自动运维任务列表
//
// @param request - DescribeActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasksWithOptions(request *DescribeActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取待执行自动运维任务列表
//
// @param request - DescribeActiveOperationTasksRequest
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasks(request *DescribeActiveOperationTasksRequest) (_result *DescribeActiveOperationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveOperationTasksResponse{}
	_body, _err := client.DescribeActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 冷存储表列表
//
// @param request - DescribeArchiveTableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeArchiveTableListResponse
func (client *Client) DescribeArchiveTableListWithOptions(request *DescribeArchiveTableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeArchiveTableListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeArchiveTableList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeArchiveTableListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 冷存储表列表
//
// @param request - DescribeArchiveTableListRequest
//
// @return DescribeArchiveTableListResponse
func (client *Client) DescribeArchiveTableList(request *DescribeArchiveTableListRequest) (_result *DescribeArchiveTableListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeArchiveTableListResponse{}
	_body, _err := client.DescribeArchiveTableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2020-02-02"),
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
// 备份集详情
//
// @param request - DescribeBackupSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetResponse
func (client *Client) DescribeBackupSetWithOptions(request *DescribeBackupSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DestCrossRegion) {
		query["DestCrossRegion"] = request.DestCrossRegion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupSet"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 备份集详情
//
// @param request - DescribeBackupSetRequest
//
// @return DescribeBackupSetResponse
func (client *Client) DescribeBackupSet(request *DescribeBackupSetRequest) (_result *DescribeBackupSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupSetResponse{}
	_body, _err := client.DescribeBackupSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeBackupSetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetListResponse
func (client *Client) DescribeBackupSetListWithOptions(request *DescribeBackupSetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupSetListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupSetList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupSetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBackupSetListRequest
//
// @return DescribeBackupSetListResponse
func (client *Client) DescribeBackupSetList(request *DescribeBackupSetListRequest) (_result *DescribeBackupSetListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupSetListResponse{}
	_body, _err := client.DescribeBackupSetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeBinaryLogListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBinaryLogListResponse
func (client *Client) DescribeBinaryLogListWithOptions(request *DescribeBinaryLogListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBinaryLogListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
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
		Action:      dara.String("DescribeBinaryLogList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBinaryLogListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBinaryLogListRequest
//
// @return DescribeBinaryLogListResponse
func (client *Client) DescribeBinaryLogList(request *DescribeBinaryLogListRequest) (_result *DescribeBinaryLogListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBinaryLogListResponse{}
	_body, _err := client.DescribeBinaryLogListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询CDC信息
//
// @param request - DescribeCdcInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdcInfoResponse
func (client *Client) DescribeCdcInfoWithOptions(request *DescribeCdcInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdcInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdcInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdcInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询CDC信息
//
// @param request - DescribeCdcInfoRequest
//
// @return DescribeCdcInfoResponse
func (client *Client) DescribeCdcInfo(request *DescribeCdcInfoRequest) (_result *DescribeCdcInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdcInfoResponse{}
	_body, _err := client.DescribeCdcInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeCharacterSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCharacterSetResponse
func (client *Client) DescribeCharacterSetWithOptions(request *DescribeCharacterSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeCharacterSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCharacterSet"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCharacterSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCharacterSetRequest
//
// @return DescribeCharacterSetResponse
func (client *Client) DescribeCharacterSet(request *DescribeCharacterSetRequest) (_result *DescribeCharacterSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCharacterSetResponse{}
	_body, _err := client.DescribeCharacterSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 冷存储基础信息
//
// @param request - DescribeColdDataBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColdDataBasicInfoResponse
func (client *Client) DescribeColdDataBasicInfoWithOptions(request *DescribeColdDataBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeColdDataBasicInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColdDataBasicInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColdDataBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 冷存储基础信息
//
// @param request - DescribeColdDataBasicInfoRequest
//
// @return DescribeColdDataBasicInfoResponse
func (client *Client) DescribeColdDataBasicInfo(request *DescribeColdDataBasicInfoRequest) (_result *DescribeColdDataBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeColdDataBasicInfoResponse{}
	_body, _err := client.DescribeColdDataBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - DescribeDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceAttributeResponse
func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
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
		Action:      dara.String("DescribeDBInstanceAttribute"),
		Version:     dara.String("2020-02-02"),
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
// 获取实例详情
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
// 查询实例配置信息
//
// @param request - DescribeDBInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceConfigResponse
func (client *Client) DescribeDBInstanceConfigWithOptions(request *DescribeDBInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceConfig"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// 查询实例配置信息
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
// 查询实例的HA信息
//
// @param request - DescribeDBInstanceHARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceHAResponse
func (client *Client) DescribeDBInstanceHAWithOptions(request *DescribeDBInstanceHARequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceHAResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceHA"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceHAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例的HA信息
//
// @param request - DescribeDBInstanceHARequest
//
// @return DescribeDBInstanceHAResponse
func (client *Client) DescribeDBInstanceHA(request *DescribeDBInstanceHARequest) (_result *DescribeDBInstanceHAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceHAResponse{}
	_body, _err := client.DescribeDBInstanceHAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDBInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceSSLResponse
func (client *Client) DescribeDBInstanceSSLWithOptions(request *DescribeDBInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceSSLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceSSL"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDBInstanceSSLRequest
//
// @return DescribeDBInstanceSSLResponse
func (client *Client) DescribeDBInstanceSSL(request *DescribeDBInstanceSSLRequest) (_result *DescribeDBInstanceSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.DescribeDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDBInstanceTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceTDEResponse
func (client *Client) DescribeDBInstanceTDEWithOptions(request *DescribeDBInstanceTDERequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceTDEResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceTDE"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceTDEResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDBInstanceTDERequest
//
// @return DescribeDBInstanceTDEResponse
func (client *Client) DescribeDBInstanceTDE(request *DescribeDBInstanceTDERequest) (_result *DescribeDBInstanceTDEResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceTDEResponse{}
	_body, _err := client.DescribeDBInstanceTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例拓扑
//
// @param request - DescribeDBInstanceTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceTopologyResponse
func (client *Client) DescribeDBInstanceTopologyWithOptions(request *DescribeDBInstanceTopologyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceTopologyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MinuteSimple) {
		query["MinuteSimple"] = request.MinuteSimple
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
		Action:      dara.String("DescribeDBInstanceTopology"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceTopologyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例拓扑
//
// @param request - DescribeDBInstanceTopologyRequest
//
// @return DescribeDBInstanceTopologyResponse
func (client *Client) DescribeDBInstanceTopology(request *DescribeDBInstanceTopologyRequest) (_result *DescribeDBInstanceTopologyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceTopologyResponse{}
	_body, _err := client.DescribeDBInstanceTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过Endpoint查询实例
//
// @param request - DescribeDBInstanceViaEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceViaEndpointResponse
func (client *Client) DescribeDBInstanceViaEndpointWithOptions(request *DescribeDBInstanceViaEndpointRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceViaEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceViaEndpoint"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceViaEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过Endpoint查询实例
//
// @param request - DescribeDBInstanceViaEndpointRequest
//
// @return DescribeDBInstanceViaEndpointResponse
func (client *Client) DescribeDBInstanceViaEndpoint(request *DescribeDBInstanceViaEndpointRequest) (_result *DescribeDBInstanceViaEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceViaEndpointResponse{}
	_body, _err := client.DescribeDBInstanceViaEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取PolarDB-X实例列表
//
// @param request - DescribeDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesResponse
func (client *Client) DescribeDBInstancesWithOptions(request *DescribeDBInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbVersion) {
		query["DbVersion"] = request.DbVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MustHasCdc) {
		query["MustHasCdc"] = request.MustHasCdc
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

	if !dara.IsNil(request.Series) {
		query["Series"] = request.Series
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstances"),
		Version:     dara.String("2020-02-02"),
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
// 获取PolarDB-X实例列表
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
// 性能监控数据接口
//
// @param request - DescribeDBNodePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBNodePerformanceResponse
func (client *Client) DescribeDBNodePerformanceWithOptions(request *DescribeDBNodePerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBNodePerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !dara.IsNil(request.DBNodeRole) {
		query["DBNodeRole"] = request.DBNodeRole
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
		Action:      dara.String("DescribeDBNodePerformance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 性能监控数据接口
//
// @param request - DescribeDBNodePerformanceRequest
//
// @return DescribeDBNodePerformanceResponse
func (client *Client) DescribeDBNodePerformance(request *DescribeDBNodePerformanceRequest) (_result *DescribeDBNodePerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.DescribeDBNodePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDbListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDbListResponse
func (client *Client) DescribeDbListWithOptions(request *DescribeDbListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDbListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
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
		Action:      dara.String("DescribeDbList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDbListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDbListRequest
//
// @return DescribeDbListResponse
func (client *Client) DescribeDbList(request *DescribeDbListRequest) (_result *DescribeDbListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDbListResponse{}
	_body, _err := client.DescribeDbListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDistributeTableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistributeTableListResponse
func (client *Client) DescribeDistributeTableListWithOptions(request *DescribeDistributeTableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDistributeTableListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDistributeTableList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDistributeTableListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDistributeTableListRequest
//
// @return DescribeDistributeTableListResponse
func (client *Client) DescribeDistributeTableList(request *DescribeDistributeTableListRequest) (_result *DescribeDistributeTableListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDistributeTableListResponse{}
	_body, _err := client.DescribeDistributeTableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 历史事件
//
// @param request - DescribeEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEvents"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 历史事件
//
// @param request - DescribeEventsRequest
//
// @return DescribeEventsResponse
func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取GDN实例列表
//
// @param request - DescribeGdnInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGdnInstancesResponse
func (client *Client) DescribeGdnInstancesWithOptions(request *DescribeGdnInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGdnInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterType) {
		query["FilterType"] = request.FilterType
	}

	if !dara.IsNil(request.FilterValue) {
		query["FilterValue"] = request.FilterValue
	}

	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
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
		Action:      dara.String("DescribeGdnInstances"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGdnInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取GDN实例列表
//
// @param request - DescribeGdnInstancesRequest
//
// @return DescribeGdnInstancesResponse
func (client *Client) DescribeGdnInstances(request *DescribeGdnInstancesRequest) (_result *DescribeGdnInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGdnInstancesResponse{}
	_body, _err := client.DescribeGdnInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放商业备份集
//
// @param request - DescribeOpenBackupSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenBackupSetResponse
func (client *Client) DescribeOpenBackupSetWithOptions(request *DescribeOpenBackupSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeOpenBackupSetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOpenBackupSet"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOpenBackupSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放商业备份集
//
// @param request - DescribeOpenBackupSetRequest
//
// @return DescribeOpenBackupSetResponse
func (client *Client) DescribeOpenBackupSet(request *DescribeOpenBackupSetRequest) (_result *DescribeOpenBackupSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOpenBackupSetResponse{}
	_body, _err := client.DescribeOpenBackupSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeParameterTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterTemplatesResponse
func (client *Client) DescribeParameterTemplatesWithOptions(request *DescribeParameterTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
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
		Action:      dara.String("DescribeParameterTemplates"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeParameterTemplatesRequest
//
// @return DescribeParameterTemplatesResponse
func (client *Client) DescribeParameterTemplates(request *DescribeParameterTemplatesRequest) (_result *DescribeParameterTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.DescribeParameterTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParametersResponse
func (client *Client) DescribeParametersWithOptions(request *DescribeParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("DescribeParameters"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeParametersRequest
//
// @return DescribeParametersResponse
func (client *Client) DescribeParameters(request *DescribeParametersRequest) (_result *DescribeParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DescribeParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2020-02-02"),
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

// @param request - DescribeScaleOutMigrateTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScaleOutMigrateTaskListResponse
func (client *Client) DescribeScaleOutMigrateTaskListWithOptions(request *DescribeScaleOutMigrateTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeScaleOutMigrateTaskListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScaleOutMigrateTaskList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScaleOutMigrateTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeScaleOutMigrateTaskListRequest
//
// @return DescribeScaleOutMigrateTaskListResponse
func (client *Client) DescribeScaleOutMigrateTaskList(request *DescribeScaleOutMigrateTaskListRequest) (_result *DescribeScaleOutMigrateTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScaleOutMigrateTaskListResponse{}
	_body, _err := client.DescribeScaleOutMigrateTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeSecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityIpsResponse
func (client *Client) DescribeSecurityIpsWithOptions(request *DescribeSecurityIpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityIps"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeSecurityIpsRequest
//
// @return DescribeSecurityIpsResponse
func (client *Client) DescribeSecurityIps(request *DescribeSecurityIpsRequest) (_result *DescribeSecurityIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityIpsResponse{}
	_body, _err := client.DescribeSecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 慢SQL明细
//
// @param request - DescribeSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
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
		Action:      dara.String("DescribeSlowLogRecords"),
		Version:     dara.String("2020-02-02"),
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
// 慢SQL明细
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
// 标签列表查询
//
// @param request - DescribeTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 标签列表查询
//
// @param request - DescribeTagsRequest
//
// @return DescribeTagsResponse
func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeTasksRequest
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeUserEncryptionKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserEncryptionKeyListResponse
func (client *Client) DescribeUserEncryptionKeyListWithOptions(request *DescribeUserEncryptionKeyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserEncryptionKeyList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeUserEncryptionKeyListRequest
//
// @return DescribeUserEncryptionKeyListResponse
func (client *Client) DescribeUserEncryptionKeyList(request *DescribeUserEncryptionKeyListRequest) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DescribeUserEncryptionKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableRightsSeparationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableRightsSeparationResponse
func (client *Client) DisableRightsSeparationWithOptions(request *DisableRightsSeparationRequest, runtime *dara.RuntimeOptions) (_result *DisableRightsSeparationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbaAccountName) {
		query["DbaAccountName"] = request.DbaAccountName
	}

	if !dara.IsNil(request.DbaAccountPassword) {
		query["DbaAccountPassword"] = request.DbaAccountPassword
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableRightsSeparation"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableRightsSeparationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableRightsSeparationRequest
//
// @return DisableRightsSeparationResponse
func (client *Client) DisableRightsSeparation(request *DisableRightsSeparationRequest) (_result *DisableRightsSeparationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableRightsSeparationResponse{}
	_body, _err := client.DisableRightsSeparationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启三权分立
//
// @param request - EnableRightsSeparationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRightsSeparationResponse
func (client *Client) EnableRightsSeparationWithOptions(request *EnableRightsSeparationRequest, runtime *dara.RuntimeOptions) (_result *EnableRightsSeparationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditAccountDescription) {
		query["AuditAccountDescription"] = request.AuditAccountDescription
	}

	if !dara.IsNil(request.AuditAccountName) {
		query["AuditAccountName"] = request.AuditAccountName
	}

	if !dara.IsNil(request.AuditAccountPassword) {
		query["AuditAccountPassword"] = request.AuditAccountPassword
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountDescription) {
		query["SecurityAccountDescription"] = request.SecurityAccountDescription
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableRightsSeparation"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableRightsSeparationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启三权分立
//
// @param request - EnableRightsSeparationRequest
//
// @return EnableRightsSeparationResponse
func (client *Client) EnableRightsSeparation(request *EnableRightsSeparationRequest) (_result *EnableRightsSeparationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableRightsSeparationResponse{}
	_body, _err := client.EnableRightsSeparationWithOptions(request, runtime)
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
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2020-02-02"),
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

// @param request - ModifyAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPrivilegeResponse
func (client *Client) ModifyAccountPrivilegeWithOptions(request *ModifyAccountPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPrivilegeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountPrivilege"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAccountPrivilegeRequest
//
// @return ModifyAccountPrivilegeResponse
func (client *Client) ModifyAccountPrivilege(request *ModifyAccountPrivilegeRequest) (_result *ModifyAccountPrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.ModifyAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改全局运维窗口信息
//
// @param request - ModifyActiveOperationMaintainConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationMaintainConfResponse
func (client *Client) ModifyActiveOperationMaintainConfWithOptions(request *ModifyActiveOperationMaintainConfRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationMaintainConfResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationMaintainConf"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActiveOperationMaintainConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改全局运维窗口信息
//
// @param request - ModifyActiveOperationMaintainConfRequest
//
// @return ModifyActiveOperationMaintainConfResponse
func (client *Client) ModifyActiveOperationMaintainConf(request *ModifyActiveOperationMaintainConfRequest) (_result *ModifyActiveOperationMaintainConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyActiveOperationMaintainConfResponse{}
	_body, _err := client.ModifyActiveOperationMaintainConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改主动运维任务
//
// @param request - ModifyActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasksWithOptions(request *ModifyActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.ImmediateStart) {
		query["ImmediateStart"] = request.ImmediateStart
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActiveOperationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改主动运维任务
//
// @param request - ModifyActiveOperationTasksRequest
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasks(request *ModifyActiveOperationTasksRequest) (_result *ModifyActiveOperationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyActiveOperationTasksResponse{}
	_body, _err := client.ModifyActiveOperationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyDBInstanceClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceClassResponse
func (client *Client) ModifyDBInstanceClassWithOptions(request *ModifyDBInstanceClassRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceClassResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CnClass) {
		query["CnClass"] = request.CnClass
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DnClass) {
		query["DnClass"] = request.DnClass
	}

	if !dara.IsNil(request.DnStorageSpace) {
		query["DnStorageSpace"] = request.DnStorageSpace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SpecifiedDNScale) {
		query["SpecifiedDNScale"] = request.SpecifiedDNScale
	}

	if !dara.IsNil(request.SpecifiedDNSpecMapJson) {
		query["SpecifiedDNSpecMapJson"] = request.SpecifiedDNSpecMapJson
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !dara.IsNil(request.TargetDBInstanceClass) {
		query["TargetDBInstanceClass"] = request.TargetDBInstanceClass
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceClass"),
		Version:     dara.String("2020-02-02"),
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
// 修改实例配置
//
// @param request - ModifyDBInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConfigResponse
func (client *Client) ModifyDBInstanceConfigWithOptions(request *ModifyDBInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.ConfigValue) {
		query["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConfig"),
		Version:     dara.String("2020-02-02"),
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
// 修改实例配置
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
// 修改实例链接串
//
// @param request - ModifyDBInstanceConnectionStringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConnectionStringResponse
func (client *Client) ModifyDBInstanceConnectionStringWithOptions(request *ModifyDBInstanceConnectionStringRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.NewPort) {
		query["NewPort"] = request.NewPort
	}

	if !dara.IsNil(request.NewPrefix) {
		query["NewPrefix"] = request.NewPrefix
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConnectionString"),
		Version:     dara.String("2020-02-02"),
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
// 修改实例链接串
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

// @param request - ModifyDBInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceDescriptionResponse
func (client *Client) ModifyDBInstanceDescriptionWithOptions(request *ModifyDBInstanceDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceDescription"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyDBInstanceDescriptionRequest
//
// @return ModifyDBInstanceDescriptionResponse
func (client *Client) ModifyDBInstanceDescription(request *ModifyDBInstanceDescriptionRequest) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.ModifyDBInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyDatabaseDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseDescriptionResponse
func (client *Client) ModifyDatabaseDescriptionWithOptions(request *ModifyDatabaseDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbDescription) {
		query["DbDescription"] = request.DbDescription
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatabaseDescription"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyDatabaseDescriptionRequest
//
// @return ModifyDatabaseDescriptionResponse
func (client *Client) ModifyDatabaseDescription(request *ModifyDatabaseDescriptionRequest) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.ModifyDatabaseDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParameterResponse
func (client *Client) ModifyParameterWithOptions(request *ModifyParameterRequest, runtime *dara.RuntimeOptions) (_result *ModifyParameterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ParamLevel) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
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
		Action:      dara.String("ModifyParameter"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyParameterRequest
//
// @return ModifyParameterResponse
func (client *Client) ModifyParameter(request *ModifyParameterRequest) (_result *ModifyParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyParameterResponse{}
	_body, _err := client.ModifyParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifySecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityIpsResponse
func (client *Client) ModifySecurityIpsWithOptions(request *ModifySecurityIpsRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
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
		Action:      dara.String("ModifySecurityIps"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifySecurityIpsRequest
//
// @return ModifySecurityIpsResponse
func (client *Client) ModifySecurityIps(request *ModifySecurityIpsRequest) (_result *ModifySecurityIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.ModifySecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭冷存储
//
// @param request - ReleaseColdDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseColdDataVolumeResponse
func (client *Client) ReleaseColdDataVolumeWithOptions(request *ReleaseColdDataVolumeRequest, runtime *dara.RuntimeOptions) (_result *ReleaseColdDataVolumeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseColdDataVolume"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseColdDataVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭冷存储
//
// @param request - ReleaseColdDataVolumeRequest
//
// @return ReleaseColdDataVolumeResponse
func (client *Client) ReleaseColdDataVolume(request *ReleaseColdDataVolumeRequest) (_result *ReleaseColdDataVolumeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseColdDataVolumeResponse{}
	_body, _err := client.ReleaseColdDataVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReleaseInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstancePublicConnectionResponse
func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstancePublicConnection"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReleaseInstancePublicConnectionRequest
//
// @return ReleaseInstancePublicConnectionResponse
func (client *Client) ReleaseInstancePublicConnection(request *ReleaseInstancePublicConnectionRequest) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.ReleaseInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResetAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2020-02-02"),
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

// @param request - RestartDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBInstanceResponse
func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDBInstance"),
		Version:     dara.String("2020-02-02"),
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

// @param request - SwitchDBInstanceHARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchDBInstanceHAResponse
func (client *Client) SwitchDBInstanceHAWithOptions(request *SwitchDBInstanceHARequest, runtime *dara.RuntimeOptions) (_result *SwitchDBInstanceHAResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
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

	if !dara.IsNil(request.TargetPrimaryAzoneId) {
		query["TargetPrimaryAzoneId"] = request.TargetPrimaryAzoneId
	}

	if !dara.IsNil(request.TargetPrimaryRegionId) {
		query["TargetPrimaryRegionId"] = request.TargetPrimaryRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchDBInstanceHA"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchDBInstanceHAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SwitchDBInstanceHARequest
//
// @return SwitchDBInstanceHAResponse
func (client *Client) SwitchDBInstanceHA(request *SwitchDBInstanceHARequest) (_result *SwitchDBInstanceHAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchDBInstanceHAResponse{}
	_body, _err := client.SwitchDBInstanceHAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GDN主备切换
//
// @param request - SwitchGdnMemberRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchGdnMemberRoleResponse
func (client *Client) SwitchGdnMemberRoleWithOptions(request *SwitchGdnMemberRoleRequest, runtime *dara.RuntimeOptions) (_result *SwitchGdnMemberRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	if !dara.IsNil(request.TaskTimeout) {
		query["TaskTimeout"] = request.TaskTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchGdnMemberRole"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchGdnMemberRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GDN主备切换
//
// @param request - SwitchGdnMemberRoleRequest
//
// @return SwitchGdnMemberRoleResponse
func (client *Client) SwitchGdnMemberRole(request *SwitchGdnMemberRoleRequest) (_result *SwitchGdnMemberRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchGdnMemberRoleResponse{}
	_body, _err := client.SwitchGdnMemberRoleWithOptions(request, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
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
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupPolicyResponse
func (client *Client) UpdateBackupPolicyWithOptions(request *UpdateBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateBackupPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.BackupPlanBegin) {
		query["BackupPlanBegin"] = request.BackupPlanBegin
	}

	if !dara.IsNil(request.BackupSetRetention) {
		query["BackupSetRetention"] = request.BackupSetRetention
	}

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.BackupWay) {
		query["BackupWay"] = request.BackupWay
	}

	if !dara.IsNil(request.ColdDataBackupInterval) {
		query["ColdDataBackupInterval"] = request.ColdDataBackupInterval
	}

	if !dara.IsNil(request.ColdDataBackupRetention) {
		query["ColdDataBackupRetention"] = request.ColdDataBackupRetention
	}

	if !dara.IsNil(request.CrossRegionDataBackupRetention) {
		query["CrossRegionDataBackupRetention"] = request.CrossRegionDataBackupRetention
	}

	if !dara.IsNil(request.CrossRegionLogBackupRetention) {
		query["CrossRegionLogBackupRetention"] = request.CrossRegionLogBackupRetention
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DestCrossRegion) {
		query["DestCrossRegion"] = request.DestCrossRegion
	}

	if !dara.IsNil(request.ForceCleanOnHighSpaceUsage) {
		query["ForceCleanOnHighSpaceUsage"] = request.ForceCleanOnHighSpaceUsage
	}

	if !dara.IsNil(request.IsCrossRegionDataBackupEnabled) {
		query["IsCrossRegionDataBackupEnabled"] = request.IsCrossRegionDataBackupEnabled
	}

	if !dara.IsNil(request.IsCrossRegionLogBackupEnabled) {
		query["IsCrossRegionLogBackupEnabled"] = request.IsCrossRegionLogBackupEnabled
	}

	if !dara.IsNil(request.IsEnabled) {
		query["IsEnabled"] = request.IsEnabled
	}

	if !dara.IsNil(request.LocalLogRetention) {
		query["LocalLogRetention"] = request.LocalLogRetention
	}

	if !dara.IsNil(request.LocalLogRetentionNumber) {
		query["LocalLogRetentionNumber"] = request.LocalLogRetentionNumber
	}

	if !dara.IsNil(request.LogLocalRetentionSpace) {
		query["LogLocalRetentionSpace"] = request.LogLocalRetentionSpace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoveLogRetention) {
		query["RemoveLogRetention"] = request.RemoveLogRetention
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBackupPolicy"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateBackupPolicyRequest
//
// @return UpdateBackupPolicyResponse
func (client *Client) UpdateBackupPolicy(request *UpdateBackupPolicyRequest) (_result *UpdateBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBackupPolicyResponse{}
	_body, _err := client.UpdateBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateDBInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDBInstanceSSLResponse
func (client *Client) UpdateDBInstanceSSLWithOptions(request *UpdateDBInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *UpdateDBInstanceSSLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertCommonName) {
		query["CertCommonName"] = request.CertCommonName
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.EnableSSL) {
		query["EnableSSL"] = request.EnableSSL
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDBInstanceSSL"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDBInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateDBInstanceSSLRequest
//
// @return UpdateDBInstanceSSLResponse
func (client *Client) UpdateDBInstanceSSL(request *UpdateDBInstanceSSLRequest) (_result *UpdateDBInstanceSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDBInstanceSSLResponse{}
	_body, _err := client.UpdateDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateDBInstanceTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDBInstanceTDEResponse
func (client *Client) UpdateDBInstanceTDEWithOptions(request *UpdateDBInstanceTDERequest, runtime *dara.RuntimeOptions) (_result *UpdateDBInstanceTDEResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.TDEStatus) {
		query["TDEStatus"] = request.TDEStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDBInstanceTDE"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDBInstanceTDEResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateDBInstanceTDERequest
//
// @return UpdateDBInstanceTDEResponse
func (client *Client) UpdateDBInstanceTDE(request *UpdateDBInstanceTDERequest) (_result *UpdateDBInstanceTDEResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDBInstanceTDEResponse{}
	_body, _err := client.UpdateDBInstanceTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 扩缩容实例节点数
//
// @param request - UpdatePolarDBXInstanceNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolarDBXInstanceNodeResponse
func (client *Client) UpdatePolarDBXInstanceNodeWithOptions(request *UpdatePolarDBXInstanceNodeRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolarDBXInstanceNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddDNSpec) {
		query["AddDNSpec"] = request.AddDNSpec
	}

	if !dara.IsNil(request.CNNodeCount) {
		query["CNNodeCount"] = request.CNNodeCount
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DNNodeCount) {
		query["DNNodeCount"] = request.DNNodeCount
	}

	if !dara.IsNil(request.DbInstanceNodeCount) {
		query["DbInstanceNodeCount"] = request.DbInstanceNodeCount
	}

	if !dara.IsNil(request.DeleteDNIds) {
		query["DeleteDNIds"] = request.DeleteDNIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StoragePoolName) {
		query["StoragePoolName"] = request.StoragePoolName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolarDBXInstanceNode"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolarDBXInstanceNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 扩缩容实例节点数
//
// @param request - UpdatePolarDBXInstanceNodeRequest
//
// @return UpdatePolarDBXInstanceNodeResponse
func (client *Client) UpdatePolarDBXInstanceNode(request *UpdatePolarDBXInstanceNodeRequest) (_result *UpdatePolarDBXInstanceNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePolarDBXInstanceNodeResponse{}
	_body, _err := client.UpdatePolarDBXInstanceNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpgradeDBInstanceKernelVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceKernelVersionResponse
func (client *Client) UpgradeDBInstanceKernelVersionWithOptions(request *UpgradeDBInstanceKernelVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBInstanceKernelVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.MinorVersion) {
		query["MinorVersion"] = request.MinorVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceKernelVersion"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBInstanceKernelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpgradeDBInstanceKernelVersionRequest
//
// @return UpgradeDBInstanceKernelVersionResponse
func (client *Client) UpgradeDBInstanceKernelVersion(request *UpgradeDBInstanceKernelVersionRequest) (_result *UpgradeDBInstanceKernelVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeDBInstanceKernelVersionResponse{}
	_body, _err := client.UpgradeDBInstanceKernelVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
