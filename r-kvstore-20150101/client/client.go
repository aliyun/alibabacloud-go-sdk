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
		"cn-qingdao":                  dara.String("r-kvstore.aliyuncs.com"),
		"cn-beijing":                  dara.String("r-kvstore.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai":                 dara.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("r-kvstore.aliyuncs.com"),
		"cn-heyuan":                   dara.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("r-kvstore.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("r-kvstore.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("r-kvstore.aliyuncs.com"),
		"cn-edge-1":                   dara.String("r-kvstore.aliyuncs.com"),
		"cn-fujian":                   dara.String("r-kvstore.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("r-kvstore.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("r-kvstore.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("r-kvstore.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("r-kvstore.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("r-kvstore.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("r-kvstore.aliyuncs.com"),
		"cn-wuhan":                    dara.String("r-kvstore.aliyuncs.com"),
		"cn-yushanfang":               dara.String("r-kvstore.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("r-kvstore.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("r-kvstore.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("r-kvstore.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("r-kvstore.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("r-kvstore.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("r-kvstore.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("r-kvstore"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds one or more data shards to a Tair cluster instance.
//
// Description:
//
// This operation is available only for cluster instances that use cloud disks.
//
// @param request - AddShardingNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShardingNodeResponse
func (client *Client) AddShardingNodeWithOptions(request *AddShardingNodeRequest, runtime *dara.RuntimeOptions) (_result *AddShardingNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.ForceTrans) {
		query["ForceTrans"] = request.ForceTrans
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ShardCount) {
		query["ShardCount"] = request.ShardCount
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddShardingNode"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddShardingNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more data shards to a Tair cluster instance.
//
// Description:
//
// This operation is available only for cluster instances that use cloud disks.
//
// @param request - AddShardingNodeRequest
//
// @return AddShardingNodeResponse
func (client *Client) AddShardingNode(request *AddShardingNodeRequest) (_result *AddShardingNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddShardingNodeResponse{}
	_body, _err := client.AddShardingNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a private endpoint for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// Clients can bypass proxy nodes and use private endpoints to connect to cluster instances. This is similar to the connection to native Redis clusters. The direct connection mode can reduce communication overheads and the response time of Tair (Redis OSS-compatible).
//
// To call this operation, make sure that the instance meets the following requirements:
//
//   - The instance is a cluster instance.
//
//   - The instance is deployed in classic mode.
//
//   - The instance is deployed in a virtual private cloud (VPC). If the instance is deployed in the classic network, you can call the [SwitchNetwork](https://help.aliyun.com/document_detail/473797.html) operation to change the network type to VPC.
//
//   - SSL encryption is disabled for the instance. If SSL encryption is enabled, you can call the [ModifyInstanceSSL](https://help.aliyun.com/document_detail/473838.html) operation to disable SSL encryption.
//
// @param request - AllocateDirectConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateDirectConnectionResponse
func (client *Client) AllocateDirectConnectionWithOptions(request *AllocateDirectConnectionRequest, runtime *dara.RuntimeOptions) (_result *AllocateDirectConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateDirectConnection"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateDirectConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a private endpoint for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// Clients can bypass proxy nodes and use private endpoints to connect to cluster instances. This is similar to the connection to native Redis clusters. The direct connection mode can reduce communication overheads and the response time of Tair (Redis OSS-compatible).
//
// To call this operation, make sure that the instance meets the following requirements:
//
//   - The instance is a cluster instance.
//
//   - The instance is deployed in classic mode.
//
//   - The instance is deployed in a virtual private cloud (VPC). If the instance is deployed in the classic network, you can call the [SwitchNetwork](https://help.aliyun.com/document_detail/473797.html) operation to change the network type to VPC.
//
//   - SSL encryption is disabled for the instance. If SSL encryption is enabled, you can call the [ModifyInstanceSSL](https://help.aliyun.com/document_detail/473838.html) operation to disable SSL encryption.
//
// @param request - AllocateDirectConnectionRequest
//
// @return AllocateDirectConnectionResponse
func (client *Client) AllocateDirectConnection(request *AllocateDirectConnectionRequest) (_result *AllocateDirectConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateDirectConnectionResponse{}
	_body, _err := client.AllocateDirectConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a public endpoint for an ApsaraDB for Redis instance.
//
// Description:
//
// You can also apply for public endpoints in the ApsaraDB for Redis console. For more information, see [Use a public endpoint to connect to an ApsaraDB for Redis instance](https://help.aliyun.com/document_detail/43850.html).
//
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateInstancePublicConnection"),
		Version:     dara.String("2015-01-01"),
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

// Summary:
//
// Applies for a public endpoint for an ApsaraDB for Redis instance.
//
// Description:
//
// You can also apply for public endpoints in the ApsaraDB for Redis console. For more information, see [Use a public endpoint to connect to an ApsaraDB for Redis instance](https://help.aliyun.com/document_detail/43850.html).
//
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
// Cancels O\\\\\\&M events at a time.
//
// Description:
//
// O\\&M events cannot be canceled in the following scenarios:
//
//   - The allowCancel parameter is set to 0.
//
//   - The current time is later than the start time of the O\\&M event.
//
//   - The state value of the O\\&M event is not 3.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelActiveOperationTasks"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Cancels O\\\\\\&M events at a time.
//
// Description:
//
// O\\&M events cannot be canceled in the following scenarios:
//
//   - The allowCancel parameter is set to 0.
//
//   - The current time is later than the start time of the O\\&M event.
//
//   - The state value of the O\\&M event is not 3.
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
// Queries whether a Tair (Redis OSS-compatible) instance has the permissions to use Key Management Service (KMS).
//
// Description:
//
//	  For information about Transparent Data Encryption (TDE) and the usage notes of TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
//		- If the Tair (Redis OSS-compatible) instance is authorized to use KMS, you can call the [ModifyInstanceTDE](https://help.aliyun.com/document_detail/473859.html) operation to enable TDE.
//
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCloudResourceAuthorized"),
		Version:     dara.String("2015-01-01"),
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

// Summary:
//
// Queries whether a Tair (Redis OSS-compatible) instance has the permissions to use Key Management Service (KMS).
//
// Description:
//
//	  For information about Transparent Data Encryption (TDE) and the usage notes of TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
//		- If the Tair (Redis OSS-compatible) instance is authorized to use KMS, you can call the [ModifyInstanceTDE](https://help.aliyun.com/document_detail/473859.html) operation to enable TDE.
//
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

// Summary:
//
// Creates an account that has specific permissions for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
//	  This operation is supported only for instances that are compatible with Redis 4.0 or later.
//
//		- The instance must be in the running state.
//
//		- You can create up to 18 accounts for an instance.
//
// >  For more information about how to create an account in the console, see [Manage database accounts](https://help.aliyun.com/document_detail/92665.html).
//
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

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2015-01-01"),
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
// Creates an account that has specific permissions for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
//	  This operation is supported only for instances that are compatible with Redis 4.0 or later.
//
//		- The instance must be in the running state.
//
//		- You can create up to 18 accounts for an instance.
//
// >  For more information about how to create an account in the console, see [Manage database accounts](https://help.aliyun.com/document_detail/92665.html).
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
// Backs up a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also back up an instance in the Tair (Redis OSS-compatible) console. For more information, see [Backup and recovery](https://help.aliyun.com/document_detail/43886.html).
//
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
	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2015-01-01"),
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

// Summary:
//
// Backs up a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also back up an instance in the Tair (Redis OSS-compatible) console. For more information, see [Backup and recovery](https://help.aliyun.com/document_detail/43886.html).
//
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

// Summary:
//
// Creates a cache analysis task.
//
// Description:
//
// This operation is no longer available. Use the new operation. For more information, see [Real-time key statistics and offline key analysis](https://help.aliyun.com/document_detail/184226.html).
//
// @param request - CreateCacheAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCacheAnalysisTaskResponse
func (client *Client) CreateCacheAnalysisTaskWithOptions(request *CreateCacheAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateCacheAnalysisTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCacheAnalysisTask"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCacheAnalysisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cache analysis task.
//
// Description:
//
// This operation is no longer available. Use the new operation. For more information, see [Real-time key statistics and offline key analysis](https://help.aliyun.com/document_detail/184226.html).
//
// @param request - CreateCacheAnalysisTaskRequest
//
// @return CreateCacheAnalysisTaskResponse
func (client *Client) CreateCacheAnalysisTask(request *CreateCacheAnalysisTaskRequest) (_result *CreateCacheAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCacheAnalysisTaskResponse{}
	_body, _err := client.CreateCacheAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Converts an existing Tair DRAM-based classic instance to the first child instance of a distributed instance.
//
// Description:
//
// You can call this operation to convert an existing instance to the first child instance of a distributed instance. After the instance is converted, the distributed instance is created. Before you call this operation, make sure that the following requirements are met:
//
//   - The instance that you want to convert must be a Tair [DRAM-based](https://help.aliyun.com/document_detail/126164.html) instance that uses the classic deployment mode.
//
//   - If the existing instance is a cluster instance, the direct connection mode must be disabled for the instance. For more information, see [Release a private endpoint](https://help.aliyun.com/document_detail/150047.html).
//
// >  You can also call the [CreateInstance](https://help.aliyun.com/document_detail/473757.html) operation to create an instance that is specified as the first child instance of a distributed instance. After the child instance is created, the distributed instance to which the child instance belongs is created.
//
// @param request - CreateGlobalDistributeCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalDistributeCacheResponse
func (client *Client) CreateGlobalDistributeCacheWithOptions(request *CreateGlobalDistributeCacheRequest, runtime *dara.RuntimeOptions) (_result *CreateGlobalDistributeCacheResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SeedSubInstanceId) {
		query["SeedSubInstanceId"] = request.SeedSubInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGlobalDistributeCache"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGlobalDistributeCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts an existing Tair DRAM-based classic instance to the first child instance of a distributed instance.
//
// Description:
//
// You can call this operation to convert an existing instance to the first child instance of a distributed instance. After the instance is converted, the distributed instance is created. Before you call this operation, make sure that the following requirements are met:
//
//   - The instance that you want to convert must be a Tair [DRAM-based](https://help.aliyun.com/document_detail/126164.html) instance that uses the classic deployment mode.
//
//   - If the existing instance is a cluster instance, the direct connection mode must be disabled for the instance. For more information, see [Release a private endpoint](https://help.aliyun.com/document_detail/150047.html).
//
// >  You can also call the [CreateInstance](https://help.aliyun.com/document_detail/473757.html) operation to create an instance that is specified as the first child instance of a distributed instance. After the child instance is created, the distributed instance to which the child instance belongs is created.
//
// @param request - CreateGlobalDistributeCacheRequest
//
// @return CreateGlobalDistributeCacheResponse
func (client *Client) CreateGlobalDistributeCache(request *CreateGlobalDistributeCacheRequest) (_result *CreateGlobalDistributeCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGlobalDistributeCacheResponse{}
	_body, _err := client.CreateGlobalDistributeCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a global IP whitelist template.
//
// @param request - CreateGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalSecurityIPGroupResponse
func (client *Client) CreateGlobalSecurityIPGroupWithOptions(request *CreateGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGlobalSecurityIPGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GIpList) {
		query["GIpList"] = request.GIpList
	}

	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGlobalSecurityIPGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a global IP whitelist template.
//
// @param request - CreateGlobalSecurityIPGroupRequest
//
// @return CreateGlobalSecurityIPGroupResponse
func (client *Client) CreateGlobalSecurityIPGroup(request *CreateGlobalSecurityIPGroupRequest) (_result *CreateGlobalSecurityIPGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGlobalSecurityIPGroupResponse{}
	_body, _err := client.CreateGlobalSecurityIPGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Tair (Redis OSS-compatible) instance. If you want to create a Tair (Enterprise Edition) cloud-native instance, you can call the CreateTairInstance operation.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/54532.html) of Tair (Redis OSS-compatible).
//
// You can call this operation to create a Tair (Redis OSS-compatible) instance or a classic Tair DRAM-based instance. To create a cloud-native Tair instance, call the [CreateTairInstance](https://help.aliyun.com/document_detail/473770.html) operation.
//
// > For more information about how to create an instance that meets your requirements in the Tair (Redis OSS-compatible) console, see [Step 1: Create an instance](https://help.aliyun.com/document_detail/26351.html).
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Appendonly) {
		query["Appendonly"] = request.Appendonly
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClusterBackupId) {
		query["ClusterBackupId"] = request.ClusterBackupId
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.GlobalInstance) {
		query["GlobalInstance"] = request.GlobalInstance
	}

	if !dara.IsNil(request.GlobalInstanceId) {
		query["GlobalInstanceId"] = request.GlobalInstanceId
	}

	if !dara.IsNil(request.GlobalSecurityGroupIds) {
		query["GlobalSecurityGroupIds"] = request.GlobalSecurityGroupIds
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParamGroupId) {
		query["ParamGroupId"] = request.ParamGroupId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ReadOnlyCount) {
		query["ReadOnlyCount"] = request.ReadOnlyCount
	}

	if !dara.IsNil(request.RecoverConfigMode) {
		query["RecoverConfigMode"] = request.RecoverConfigMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaCount) {
		query["ReplicaCount"] = request.ReplicaCount
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ShardCount) {
		query["ShardCount"] = request.ShardCount
	}

	if !dara.IsNil(request.SlaveReadOnlyCount) {
		query["SlaveReadOnlyCount"] = request.SlaveReadOnlyCount
	}

	if !dara.IsNil(request.SlaveReplicaCount) {
		query["SlaveReplicaCount"] = request.SlaveReplicaCount
	}

	if !dara.IsNil(request.SrcDBInstanceId) {
		query["SrcDBInstanceId"] = request.SrcDBInstanceId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Creates a Tair (Redis OSS-compatible) instance. If you want to create a Tair (Enterprise Edition) cloud-native instance, you can call the CreateTairInstance operation.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/54532.html) of Tair (Redis OSS-compatible).
//
// You can call this operation to create a Tair (Redis OSS-compatible) instance or a classic Tair DRAM-based instance. To create a cloud-native Tair instance, call the [CreateTairInstance](https://help.aliyun.com/document_detail/473770.html) operation.
//
// > For more information about how to create an instance that meets your requirements in the Tair (Redis OSS-compatible) console, see [Step 1: Create an instance](https://help.aliyun.com/document_detail/26351.html).
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates multiple Tair (Redis OSS-compatible) instances at a time.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/54532.html) of Tair (Redis OSS-compatible).
//
// >  You can call this operation to create classic Redis Open-Source Edition instances or classic Tair DRAM-based instances. We recommend that you use an API operation for creating a single instance:
//
//   - [CreateInstance](https://help.aliyun.com/document_detail/473757.html): creates a Redis Open-Source instance or a classic Tair DRAM-based instance.
//
//   - [CreateTairInstance](https://help.aliyun.com/document_detail/473770.html): creates a Tair (Enterprise Edition) instance. The instance can be a DRAM-based, persistent memory-optimized, or ESSD/SSD-based instance.
//
// @param request - CreateInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstancesResponse
func (client *Client) CreateInstancesWithOptions(request *CreateInstancesRequest, runtime *dara.RuntimeOptions) (_result *CreateInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RebuildInstance) {
		query["RebuildInstance"] = request.RebuildInstance
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstances"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple Tair (Redis OSS-compatible) instances at a time.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/54532.html) of Tair (Redis OSS-compatible).
//
// >  You can call this operation to create classic Redis Open-Source Edition instances or classic Tair DRAM-based instances. We recommend that you use an API operation for creating a single instance:
//
//   - [CreateInstance](https://help.aliyun.com/document_detail/473757.html): creates a Redis Open-Source instance or a classic Tair DRAM-based instance.
//
//   - [CreateTairInstance](https://help.aliyun.com/document_detail/473770.html): creates a Tair (Enterprise Edition) instance. The instance can be a DRAM-based, persistent memory-optimized, or ESSD/SSD-based instance.
//
// @param request - CreateInstancesRequest
//
// @return CreateInstancesResponse
func (client *Client) CreateInstances(request *CreateInstancesRequest) (_result *CreateInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstancesResponse{}
	_body, _err := client.CreateInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a parameter template.
//
// @param request - CreateParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParameterGroupResponse
func (client *Client) CreateParameterGroupWithOptions(request *CreateParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateParameterGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupDesc) {
		query["ParameterGroupDesc"] = request.ParameterGroupDesc
	}

	if !dara.IsNil(request.ParameterGroupName) {
		query["ParameterGroupName"] = request.ParameterGroupName
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateParameterGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParameterGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a parameter template.
//
// @param request - CreateParameterGroupRequest
//
// @return CreateParameterGroupResponse
func (client *Client) CreateParameterGroup(request *CreateParameterGroupRequest) (_result *CreateParameterGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateParameterGroupResponse{}
	_body, _err := client.CreateParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建TairCustom实例
//
// @param request - CreateTCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTCInstanceResponse
func (client *Client) CreateTCInstanceWithOptions(request *CreateTCInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateTCInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DataDisk) {
		query["DataDisk"] = request.DataDisk
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.NeedEni) {
		query["NeedEni"] = request.NeedEni
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTCInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTCInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建TairCustom实例
//
// @param request - CreateTCInstanceRequest
//
// @return CreateTCInstanceResponse
func (client *Client) CreateTCInstance(request *CreateTCInstanceRequest) (_result *CreateTCInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTCInstanceResponse{}
	_body, _err := client.CreateTCInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Tair (Enterprise Edition) cloud-native instance.
//
// Description:
//
// For information about instance selection, see [Instructions for selecting an appropriate Tair (Redis OSS-compatible) instance](https://help.aliyun.com/document_detail/223808.html).
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/54532.html) of Tair (Redis OSS-compatible).
//
// >
//
//   - For information about how to create an instance in the console, see [Step 1: Create an instance](https://help.aliyun.com/document_detail/26351.html).
//
//   - To create other types of instances, such as Redis Open-Source Edition instances or [Tair DRAM-based](https://help.aliyun.com/document_detail/126164.html) instances, you can call the [CreateInstance](https://help.aliyun.com/document_detail/473757.html) operation.
//
// @param request - CreateTairInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTairInstanceResponse
func (client *Client) CreateTairInstanceWithOptions(request *CreateTairInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateTairInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterBackupId) {
		query["ClusterBackupId"] = request.ClusterBackupId
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.GlobalInstanceId) {
		query["GlobalInstanceId"] = request.GlobalInstanceId
	}

	if !dara.IsNil(request.GlobalSecurityGroupIds) {
		query["GlobalSecurityGroupIds"] = request.GlobalSecurityGroupIds
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParamGroupId) {
		query["ParamGroupId"] = request.ParamGroupId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ReadOnlyCount) {
		query["ReadOnlyCount"] = request.ReadOnlyCount
	}

	if !dara.IsNil(request.RecoverConfigMode) {
		query["RecoverConfigMode"] = request.RecoverConfigMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaCount) {
		query["ReplicaCount"] = request.ReplicaCount
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ShardCount) {
		query["ShardCount"] = request.ShardCount
	}

	if !dara.IsNil(request.ShardType) {
		query["ShardType"] = request.ShardType
	}

	if !dara.IsNil(request.SlaveReadOnlyCount) {
		query["SlaveReadOnlyCount"] = request.SlaveReadOnlyCount
	}

	if !dara.IsNil(request.SlaveReplicaCount) {
		query["SlaveReplicaCount"] = request.SlaveReplicaCount
	}

	if !dara.IsNil(request.SrcDBInstanceId) {
		query["SrcDBInstanceId"] = request.SrcDBInstanceId
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTairInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTairInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Tair (Enterprise Edition) cloud-native instance.
//
// Description:
//
// For information about instance selection, see [Instructions for selecting an appropriate Tair (Redis OSS-compatible) instance](https://help.aliyun.com/document_detail/223808.html).
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/54532.html) of Tair (Redis OSS-compatible).
//
// >
//
//   - For information about how to create an instance in the console, see [Step 1: Create an instance](https://help.aliyun.com/document_detail/26351.html).
//
//   - To create other types of instances, such as Redis Open-Source Edition instances or [Tair DRAM-based](https://help.aliyun.com/document_detail/126164.html) instances, you can call the [CreateInstance](https://help.aliyun.com/document_detail/473757.html) operation.
//
// @param request - CreateTairInstanceRequest
//
// @return CreateTairInstanceResponse
func (client *Client) CreateTairInstance(request *CreateTairInstanceRequest) (_result *CreateTairInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTairInstanceResponse{}
	_body, _err := client.CreateTairInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Tair VNode实例
//
// @param request - CreateTairKVCacheVNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTairKVCacheVNodeResponse
func (client *Client) CreateTairKVCacheVNodeWithOptions(request *CreateTairKVCacheVNodeRequest, runtime *dara.RuntimeOptions) (_result *CreateTairKVCacheVNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ComputeUnitNum) {
		query["ComputeUnitNum"] = request.ComputeUnitNum
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VkName) {
		query["VkName"] = request.VkName
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTairKVCacheVNode"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTairKVCacheVNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Tair VNode实例
//
// @param request - CreateTairKVCacheVNodeRequest
//
// @return CreateTairKVCacheVNodeResponse
func (client *Client) CreateTairKVCacheVNode(request *CreateTairKVCacheVNodeRequest) (_result *CreateTairKVCacheVNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTairKVCacheVNodeResponse{}
	_body, _err := client.CreateTairKVCacheVNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an account from a Tair (Redis OSS-compatible) instance.
//
// Description:
//
//	  This operation is supported only for instances that are compatible with Redis 4.0 or later.
//
//		- The instance must be in the Running state.
//
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2015-01-01"),
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
// Deletes an account from a Tair (Redis OSS-compatible) instance.
//
// Description:
//
//	  This operation is supported only for instances that are compatible with Redis 4.0 or later.
//
//		- The instance must be in the Running state.
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
// 删除指定备份集
//
// @param request - DeleteBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackupWithOptions(request *DeleteBackupRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DeleteBackup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定备份集
//
// @param request - DeleteBackupRequest
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackup(request *DeleteBackupRequest) (_result *DeleteBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupResponse{}
	_body, _err := client.DeleteBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a global IP whitelist template.
//
// Description:
//
// Before you delete an IP whitelist template, you must unbind (disassociate) the instances that are currently associated with the template.
//
// @param request - DeleteGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGlobalSecurityIPGroupResponse
func (client *Client) DeleteGlobalSecurityIPGroupWithOptions(request *DeleteGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGlobalSecurityIPGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGlobalSecurityIPGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a global IP whitelist template.
//
// Description:
//
// Before you delete an IP whitelist template, you must unbind (disassociate) the instances that are currently associated with the template.
//
// @param request - DeleteGlobalSecurityIPGroupRequest
//
// @return DeleteGlobalSecurityIPGroupResponse
func (client *Client) DeleteGlobalSecurityIPGroup(request *DeleteGlobalSecurityIPGroupRequest) (_result *DeleteGlobalSecurityIPGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGlobalSecurityIPGroupResponse{}
	_body, _err := client.DeleteGlobalSecurityIPGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Release the Redis instance.
//
// Description:
//
// For more information about how to perform the corresponding operation in the console, see [Release an instance](https://help.aliyun.com/document_detail/43882.html).
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The instance is in the running state.
//
//   - The instance is charged on a pay-as-you-go basis.
//
// >  You cannot call this operation to release a subscription instance, which is automatically released when it expires. To release a subscription instance before it expires, submit a ticket.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalInstanceId) {
		query["GlobalInstanceId"] = request.GlobalInstanceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Release the Redis instance.
//
// Description:
//
// For more information about how to perform the corresponding operation in the console, see [Release an instance](https://help.aliyun.com/document_detail/43882.html).
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The instance is in the running state.
//
//   - The instance is charged on a pay-as-you-go basis.
//
// >  You cannot call this operation to release a subscription instance, which is automatically released when it expires. To release a subscription instance before it expires, submit a ticket.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a parameter template.
//
// @param request - DeleteParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParameterGroupResponse
func (client *Client) DeleteParameterGroupWithOptions(request *DeleteParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteParameterGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteParameterGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParameterGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a parameter template.
//
// @param request - DeleteParameterGroupRequest
//
// @return DeleteParameterGroupResponse
func (client *Client) DeleteParameterGroup(request *DeleteParameterGroupRequest) (_result *DeleteParameterGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteParameterGroupResponse{}
	_body, _err := client.DeleteParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes one or more data shards from a Tair (Redis OSS-compatible) cluster instance.
//
// Description:
//
// You can also remove data shards from an instance in the Tair (Redis OSS-compatible) console. For more information, see [Adjust the number of shards for an instance with cloud disks](https://help.aliyun.com/document_detail/198082.html).\\
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The instance is a persistent memory-optimized instance in the cluster architecture. For more information about persistent memory-optimized instances, see [Persistent memory-optimized instances](https://help.aliyun.com/document_detail/183956.html).
//
//   - The instance has more than one data shard.
//
// @param request - DeleteShardingNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteShardingNodeResponse
func (client *Client) DeleteShardingNodeWithOptions(request *DeleteShardingNodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteShardingNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.ForceTrans) {
		query["ForceTrans"] = request.ForceTrans
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ShardCount) {
		query["ShardCount"] = request.ShardCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteShardingNode"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteShardingNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes one or more data shards from a Tair (Redis OSS-compatible) cluster instance.
//
// Description:
//
// You can also remove data shards from an instance in the Tair (Redis OSS-compatible) console. For more information, see [Adjust the number of shards for an instance with cloud disks](https://help.aliyun.com/document_detail/198082.html).\\
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The instance is a persistent memory-optimized instance in the cluster architecture. For more information about persistent memory-optimized instances, see [Persistent memory-optimized instances](https://help.aliyun.com/document_detail/183956.html).
//
//   - The instance has more than one data shard.
//
// @param request - DeleteShardingNodeRequest
//
// @return DeleteShardingNodeResponse
func (client *Client) DeleteShardingNode(request *DeleteShardingNodeRequest) (_result *DeleteShardingNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteShardingNodeResponse{}
	_body, _err := client.DeleteShardingNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a specified account of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >  Only Tair (Redis OSS-compatible) instances of Redis 4.0 or later are supported.
//
// @param request - DescribeAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccounts"),
		Version:     dara.String("2015-01-01"),
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
// Queries a specified account of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >  Only Tair (Redis OSS-compatible) instances of Redis 4.0 or later are supported.
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
// Queries the details of the O\\&M tasks of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// After you have called this API operation and queried the information about a specific O&M task, you can also call the [ModifyActiveOperationTask](https://help.aliyun.com/document_detail/473864.html) operation to modify the scheduled switchover time of the O&M task.
//
// @param request - DescribeActiveOperationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskResponse
func (client *Client) DescribeActiveOperationTaskWithOptions(request *DescribeActiveOperationTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsHistory) {
		query["IsHistory"] = request.IsHistory
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

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationTask"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the O\\&M tasks of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// After you have called this API operation and queried the information about a specific O&M task, you can also call the [ModifyActiveOperationTask](https://help.aliyun.com/document_detail/473864.html) operation to modify the scheduled switchover time of the O&M task.
//
// @param request - DescribeActiveOperationTaskRequest
//
// @return DescribeActiveOperationTaskResponse
func (client *Client) DescribeActiveOperationTask(request *DescribeActiveOperationTaskRequest) (_result *DescribeActiveOperationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveOperationTaskResponse{}
	_body, _err := client.DescribeActiveOperationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the O\\\\\\\\\\\\&M event details of an instance.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowCancel) {
		query["AllowCancel"] = request.AllowCancel
	}

	if !dara.IsNil(request.AllowChange) {
		query["AllowChange"] = request.AllowChange
	}

	if !dara.IsNil(request.ChangeLevel) {
		query["ChangeLevel"] = request.ChangeLevel
	}

	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.InsName) {
		query["InsName"] = request.InsName
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

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationTasks"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Queries the O\\\\\\\\\\\\&M event details of an instance.
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
// Queries the audit log configurations of a Tair (Redis OSS-compatible) instance. The configurations include whether the audit log feature is enabled and the retention period of audit logs.
//
// Description:
//
// Before you call this operation, you must enable the audit log feature for the instance. For more information, see [ModifyAuditLogConfig](https://help.aliyun.com/document_detail/473829.html) or [Enable the audit log feature](https://help.aliyun.com/document_detail/102015.html).
//
// @param request - DescribeAuditLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditLogConfigResponse
func (client *Client) DescribeAuditLogConfigWithOptions(request *DescribeAuditLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditLogConfig"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audit log configurations of a Tair (Redis OSS-compatible) instance. The configurations include whether the audit log feature is enabled and the retention period of audit logs.
//
// Description:
//
// Before you call this operation, you must enable the audit log feature for the instance. For more information, see [ModifyAuditLogConfig](https://help.aliyun.com/document_detail/473829.html) or [Enable the audit log feature](https://help.aliyun.com/document_detail/102015.html).
//
// @param request - DescribeAuditLogConfigRequest
//
// @return DescribeAuditLogConfigResponse
func (client *Client) DescribeAuditLogConfig(request *DescribeAuditLogConfigRequest) (_result *DescribeAuditLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuditLogConfigResponse{}
	_body, _err := client.DescribeAuditLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the audit logs of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// Before you call this operation, you must enable the audit log feature for the instance. For more information, see [ModifyAuditLogConfig](https://help.aliyun.com/document_detail/473829.html).
//
// @param request - DescribeAuditRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditRecordsResponse
func (client *Client) DescribeAuditRecordsWithOptions(request *DescribeAuditRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
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

	if !dara.IsNil(request.QueryKeywords) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditRecords"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audit logs of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// Before you call this operation, you must enable the audit log feature for the instance. For more information, see [ModifyAuditLogConfig](https://help.aliyun.com/document_detail/473829.html).
//
// @param request - DescribeAuditRecordsRequest
//
// @return DescribeAuditRecordsResponse
func (client *Client) DescribeAuditRecords(request *DescribeAuditRecordsRequest) (_result *DescribeAuditRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuditRecordsResponse{}
	_body, _err := client.DescribeAuditRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the types of Tair (Redis OSS-compatible) instances that can be created in a specified zone.
//
// @param request - DescribeAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableResourceResponse
func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceScene) {
		query["InstanceScene"] = request.InstanceScene
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableResource"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types of Tair (Redis OSS-compatible) instances that can be created in a specified zone.
//
// @param request - DescribeAvailableResourceRequest
//
// @return DescribeAvailableResourceResponse
func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup policy of a Tair (Redis OSS-compatible) instance, including the backup cycle and backup time.
//
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2015-01-01"),
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
// Queries the backup policy of a Tair (Redis OSS-compatible) instance, including the backup cycle and backup time.
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
// Queries the execution status of backup tasks for a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeBackupTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupTasksResponse
func (client *Client) DescribeBackupTasksWithOptions(request *DescribeBackupTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupJobId) {
		query["BackupJobId"] = request.BackupJobId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobMode) {
		query["JobMode"] = request.JobMode
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupTasks"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of backup tasks for a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeBackupTasksRequest
//
// @return DescribeBackupTasksResponse
func (client *Client) DescribeBackupTasks(request *DescribeBackupTasksRequest) (_result *DescribeBackupTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.DescribeBackupTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup files of the Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupJobId) {
		query["BackupJobId"] = request.BackupJobId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NeedAof) {
		query["NeedAof"] = request.NeedAof
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackups"),
		Version:     dara.String("2015-01-01"),
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
// Queries the backup files of the Tair (Redis OSS-compatible) instance.
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
// Queries the cache analysis report of an instance on a specified date.
//
// Description:
//
// > Tair (Redis OSS-compatible) has optimized the cache analytics feature to improve user experience. This API operation is phased out. You can use the new API operation for cache analytics. For more information, see [API operations for cache analytics are upgraded](https://help.aliyun.com/document_detail/186019.html).
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The engine version of the instance is Redis 4.0 or later.
//
//   - The instance uses the latest minor version. For more information about how to check whether to update the minor version of an instance, see [How do I check whether the minor version of a Tair (Redis OSS-compatible) instance is the latest?](https://help.aliyun.com/document_detail/129203.html)
//
// @param request - DescribeCacheAnalysisReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCacheAnalysisReportResponse
func (client *Client) DescribeCacheAnalysisReportWithOptions(request *DescribeCacheAnalysisReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeCacheAnalysisReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisType) {
		query["AnalysisType"] = request.AnalysisType
	}

	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumbers) {
		query["PageNumbers"] = request.PageNumbers
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCacheAnalysisReport"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCacheAnalysisReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cache analysis report of an instance on a specified date.
//
// Description:
//
// > Tair (Redis OSS-compatible) has optimized the cache analytics feature to improve user experience. This API operation is phased out. You can use the new API operation for cache analytics. For more information, see [API operations for cache analytics are upgraded](https://help.aliyun.com/document_detail/186019.html).
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The engine version of the instance is Redis 4.0 or later.
//
//   - The instance uses the latest minor version. For more information about how to check whether to update the minor version of an instance, see [How do I check whether the minor version of a Tair (Redis OSS-compatible) instance is the latest?](https://help.aliyun.com/document_detail/129203.html)
//
// @param request - DescribeCacheAnalysisReportRequest
//
// @return DescribeCacheAnalysisReportResponse
func (client *Client) DescribeCacheAnalysisReport(request *DescribeCacheAnalysisReportRequest) (_result *DescribeCacheAnalysisReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCacheAnalysisReportResponse{}
	_body, _err := client.DescribeCacheAnalysisReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of cache analysis reports for an instance.
//
// Description:
//
// > Tair (Redis OSS-compatible) has optimized the cache analytics feature to improve user experience. This API operation is phased out. You can use the new API operation for cache analytics. For more information, see [API operations for cache analytics are upgraded](https://help.aliyun.com/document_detail/186019.html).
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The engine version of the instance is Redis 4.0 or later.
//
//   - The instance uses the latest minor version. For more information about how to check whether to update the minor version of an instance, see [How do I check whether the minor version of a Tair (Redis OSS-compatible) instance is the latest?](https://help.aliyun.com/document_detail/129203.html)
//
// @param request - DescribeCacheAnalysisReportListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCacheAnalysisReportListResponse
func (client *Client) DescribeCacheAnalysisReportListWithOptions(request *DescribeCacheAnalysisReportListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCacheAnalysisReportListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Days) {
		query["Days"] = request.Days
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumbers) {
		query["PageNumbers"] = request.PageNumbers
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCacheAnalysisReportList"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCacheAnalysisReportListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of cache analysis reports for an instance.
//
// Description:
//
// > Tair (Redis OSS-compatible) has optimized the cache analytics feature to improve user experience. This API operation is phased out. You can use the new API operation for cache analytics. For more information, see [API operations for cache analytics are upgraded](https://help.aliyun.com/document_detail/186019.html).
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The engine version of the instance is Redis 4.0 or later.
//
//   - The instance uses the latest minor version. For more information about how to check whether to update the minor version of an instance, see [How do I check whether the minor version of a Tair (Redis OSS-compatible) instance is the latest?](https://help.aliyun.com/document_detail/129203.html)
//
// @param request - DescribeCacheAnalysisReportListRequest
//
// @return DescribeCacheAnalysisReportListResponse
func (client *Client) DescribeCacheAnalysisReportList(request *DescribeCacheAnalysisReportListRequest) (_result *DescribeCacheAnalysisReportListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCacheAnalysisReportListResponse{}
	_body, _err := client.DescribeCacheAnalysisReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup sets of a Tair (Redis OSS-compatible) cluster instance.
//
// Description:
//
// This operation is applicable only to cloud-native instances.
//
// @param request - DescribeClusterBackupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterBackupListResponse
func (client *Client) DescribeClusterBackupListWithOptions(request *DescribeClusterBackupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterBackupListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterBackupList"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterBackupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup sets of a Tair (Redis OSS-compatible) cluster instance.
//
// Description:
//
// This operation is applicable only to cloud-native instances.
//
// @param request - DescribeClusterBackupListRequest
//
// @return DescribeClusterBackupListResponse
func (client *Client) DescribeClusterBackupList(request *DescribeClusterBackupListRequest) (_result *DescribeClusterBackupListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterBackupListResponse{}
	_body, _err := client.DescribeClusterBackupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration information of nodes in a Tair (Redis OSS-compatible) cluster instance, such as the specifications and the maximum number of connections.
//
// Description:
//
// > This API operation is applicable only to Tair (Redis OSS-compatible) instances that use [cloud disks](https://help.aliyun.com/document_detail/188068.html) and the [cluster architecture](https://help.aliyun.com/document_detail/52228.html).
//
// @param request - DescribeClusterMemberInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterMemberInfoResponse
func (client *Client) DescribeClusterMemberInfoWithOptions(request *DescribeClusterMemberInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterMemberInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterMemberInfo"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterMemberInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration information of nodes in a Tair (Redis OSS-compatible) cluster instance, such as the specifications and the maximum number of connections.
//
// Description:
//
// > This API operation is applicable only to Tair (Redis OSS-compatible) instances that use [cloud disks](https://help.aliyun.com/document_detail/188068.html) and the [cluster architecture](https://help.aliyun.com/document_detail/52228.html).
//
// @param request - DescribeClusterMemberInfoRequest
//
// @return DescribeClusterMemberInfoResponse
func (client *Client) DescribeClusterMemberInfo(request *DescribeClusterMemberInfoRequest) (_result *DescribeClusterMemberInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterMemberInfoResponse{}
	_body, _err := client.DescribeClusterMemberInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network information of an ApsaraDB for Redis instance.
//
// @param request - DescribeDBInstanceNetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceNetInfoResponse
func (client *Client) DescribeDBInstanceNetInfoWithOptions(request *DescribeDBInstanceNetInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceNetInfo"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network information of an ApsaraDB for Redis instance.
//
// @param request - DescribeDBInstanceNetInfoRequest
//
// @return DescribeDBInstanceNetInfoResponse
func (client *Client) DescribeDBInstanceNetInfo(request *DescribeDBInstanceNetInfoRequest) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DescribeDBInstanceNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about virtual IP addresses (VIPs) of child instances of a cluster instance in direct connection mode.
//
// Description:
//
// > Only instances that use cloud disks support this operation.
//
// @param request - DescribeDBNodeDirectVipInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBNodeDirectVipInfoResponse
func (client *Client) DescribeDBNodeDirectVipInfoWithOptions(request *DescribeDBNodeDirectVipInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBNodeDirectVipInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeDBNodeDirectVipInfo"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBNodeDirectVipInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about virtual IP addresses (VIPs) of child instances of a cluster instance in direct connection mode.
//
// Description:
//
// > Only instances that use cloud disks support this operation.
//
// @param request - DescribeDBNodeDirectVipInfoRequest
//
// @return DescribeDBNodeDirectVipInfoResponse
func (client *Client) DescribeDBNodeDirectVipInfo(request *DescribeDBNodeDirectVipInfoRequest) (_result *DescribeDBNodeDirectVipInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBNodeDirectVipInfoResponse{}
	_body, _err := client.DescribeDBNodeDirectVipInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of Tair (Redis OSS-compatible) instances deployed in a dedicated cluster.
//
// Description:
//
// > If you want to query the information about Tair (Redis OSS-compatible) instances that are not deployed in a dedicated cluster, call the [DescribeInstanceAttribute](https://help.aliyun.com/document_detail/473779.html) operation.
//
// @param request - DescribeDedicatedClusterInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedClusterInstanceListResponse
func (client *Client) DescribeDedicatedClusterInstanceListWithOptions(request *DescribeDedicatedClusterInstanceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedClusterInstanceListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DedicatedHostName) {
		query["DedicatedHostName"] = request.DedicatedHostName
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceNetType) {
		query["InstanceNetType"] = request.InstanceNetType
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDedicatedClusterInstanceList"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDedicatedClusterInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of Tair (Redis OSS-compatible) instances deployed in a dedicated cluster.
//
// Description:
//
// > If you want to query the information about Tair (Redis OSS-compatible) instances that are not deployed in a dedicated cluster, call the [DescribeInstanceAttribute](https://help.aliyun.com/document_detail/473779.html) operation.
//
// @param request - DescribeDedicatedClusterInstanceListRequest
//
// @return DescribeDedicatedClusterInstanceListResponse
func (client *Client) DescribeDedicatedClusterInstanceList(request *DescribeDedicatedClusterInstanceListRequest) (_result *DescribeDedicatedClusterInstanceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDedicatedClusterInstanceListResponse{}
	_body, _err := client.DescribeDedicatedClusterInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a custom key for a Tair (Redis OSS-compatible) instance to use transparent data encryption (TDE).
//
// Description:
//
// Before you call this operation, TDE must be enabled for the Tair (Redis OSS-compatible) instance by using a custom key. For more information, see [ModifyInstanceTDE](https://help.aliyun.com/document_detail/473859.html).
//
// > For more information about TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
// @param request - DescribeEncryptionKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEncryptionKeyResponse
func (client *Client) DescribeEncryptionKeyWithOptions(request *DescribeEncryptionKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeEncryptionKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEncryptionKey"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEncryptionKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a custom key for a Tair (Redis OSS-compatible) instance to use transparent data encryption (TDE).
//
// Description:
//
// Before you call this operation, TDE must be enabled for the Tair (Redis OSS-compatible) instance by using a custom key. For more information, see [ModifyInstanceTDE](https://help.aliyun.com/document_detail/473859.html).
//
// > For more information about TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
// @param request - DescribeEncryptionKeyRequest
//
// @return DescribeEncryptionKeyResponse
func (client *Client) DescribeEncryptionKey(request *DescribeEncryptionKeyRequest) (_result *DescribeEncryptionKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEncryptionKeyResponse{}
	_body, _err := client.DescribeEncryptionKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of custom keys used by Tair (Redis OSS-compatible) instances.
//
// Description:
//
//	  You can specify a custom key when you call the [ModifyInstanceTDE](https://help.aliyun.com/document_detail/473859.html) operation to enable Transparent Data Encryption (TDE). You can call the DescribeEncryptionKeyList operation to query the custom keys that are in use. To create a custom key, you can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation of Key Management Service (KMS).
//
//		- For more information about TDE and the usage notes of TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
// @param request - DescribeEncryptionKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEncryptionKeyListResponse
func (client *Client) DescribeEncryptionKeyListWithOptions(request *DescribeEncryptionKeyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEncryptionKeyListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEncryptionKeyList"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEncryptionKeyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of custom keys used by Tair (Redis OSS-compatible) instances.
//
// Description:
//
//	  You can specify a custom key when you call the [ModifyInstanceTDE](https://help.aliyun.com/document_detail/473859.html) operation to enable Transparent Data Encryption (TDE). You can call the DescribeEncryptionKeyList operation to query the custom keys that are in use. To create a custom key, you can call the [CreateKey](https://help.aliyun.com/document_detail/28947.html) operation of Key Management Service (KMS).
//
//		- For more information about TDE and the usage notes of TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
// @param request - DescribeEncryptionKeyListRequest
//
// @return DescribeEncryptionKeyListResponse
func (client *Client) DescribeEncryptionKeyList(request *DescribeEncryptionKeyListRequest) (_result *DescribeEncryptionKeyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEncryptionKeyListResponse{}
	_body, _err := client.DescribeEncryptionKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the major version and minor version of a Tair (Redis OSS-compatible) instance and the release notes for minor versions.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=R-kvstore\\&api=DescribeEngineVersion\\&type=RPC\\&version=2015-01-01)
//
// @param request - DescribeEngineVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEngineVersionResponse
func (client *Client) DescribeEngineVersionWithOptions(request *DescribeEngineVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeEngineVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEngineVersion"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEngineVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the major version and minor version of a Tair (Redis OSS-compatible) instance and the release notes for minor versions.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=R-kvstore\\&api=DescribeEngineVersion\\&type=RPC\\&version=2015-01-01)
//
// @param request - DescribeEngineVersionRequest
//
// @return DescribeEngineVersionResponse
func (client *Client) DescribeEngineVersion(request *DescribeEngineVersionRequest) (_result *DescribeEngineVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEngineVersionResponse{}
	_body, _err := client.DescribeEngineVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a distributed Tair (Redis OSS-compatible) instance.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=R-kvstore\\&api=DescribeGlobalDistributeCache\\&type=RPC\\&version=2015-01-01)
//
// @param request - DescribeGlobalDistributeCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalDistributeCacheResponse
func (client *Client) DescribeGlobalDistributeCacheWithOptions(request *DescribeGlobalDistributeCacheRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalDistributeCacheResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalInstanceId) {
		query["GlobalInstanceId"] = request.GlobalInstanceId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SubInstanceId) {
		query["SubInstanceId"] = request.SubInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalDistributeCache"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalDistributeCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a distributed Tair (Redis OSS-compatible) instance.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=R-kvstore\\&api=DescribeGlobalDistributeCache\\&type=RPC\\&version=2015-01-01)
//
// @param request - DescribeGlobalDistributeCacheRequest
//
// @return DescribeGlobalDistributeCacheResponse
func (client *Client) DescribeGlobalDistributeCache(request *DescribeGlobalDistributeCacheRequest) (_result *DescribeGlobalDistributeCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalDistributeCacheResponse{}
	_body, _err := client.DescribeGlobalDistributeCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries global IP whitelist templates.
//
// @param request - DescribeGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalSecurityIPGroupResponse
func (client *Client) DescribeGlobalSecurityIPGroupWithOptions(request *DescribeGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalSecurityIPGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalSecurityIPGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries global IP whitelist templates.
//
// @param request - DescribeGlobalSecurityIPGroupRequest
//
// @return DescribeGlobalSecurityIPGroupResponse
func (client *Client) DescribeGlobalSecurityIPGroup(request *DescribeGlobalSecurityIPGroupRequest) (_result *DescribeGlobalSecurityIPGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalSecurityIPGroupResponse{}
	_body, _err := client.DescribeGlobalSecurityIPGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the global IP whitelist templates associated with an instance.
//
// @param request - DescribeGlobalSecurityIPGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalSecurityIPGroupRelationResponse
func (client *Client) DescribeGlobalSecurityIPGroupRelationWithOptions(request *DescribeGlobalSecurityIPGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalSecurityIPGroupRelationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalSecurityIPGroupRelation"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the global IP whitelist templates associated with an instance.
//
// @param request - DescribeGlobalSecurityIPGroupRelationRequest
//
// @return DescribeGlobalSecurityIPGroupRelationResponse
func (client *Client) DescribeGlobalSecurityIPGroupRelation(request *DescribeGlobalSecurityIPGroupRelationRequest) (_result *DescribeGlobalSecurityIPGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.DescribeGlobalSecurityIPGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the performance monitoring data of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also query the performance monitoring data of an instance in the Tair console. For more information, see [Metrics](https://help.aliyun.com/document_detail/43887.html).
//
// @param request - DescribeHistoryMonitorValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryMonitorValuesResponse
func (client *Client) DescribeHistoryMonitorValuesWithOptions(request *DescribeHistoryMonitorValuesRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryMonitorValuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntervalForHistory) {
		query["IntervalForHistory"] = request.IntervalForHistory
	}

	if !dara.IsNil(request.MonitorKeys) {
		query["MonitorKeys"] = request.MonitorKeys
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeRole) {
		query["NodeRole"] = request.NodeRole
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryMonitorValues"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryMonitorValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance monitoring data of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also query the performance monitoring data of an instance in the Tair console. For more information, see [Metrics](https://help.aliyun.com/document_detail/43887.html).
//
// @param request - DescribeHistoryMonitorValuesRequest
//
// @return DescribeHistoryMonitorValuesResponse
func (client *Client) DescribeHistoryMonitorValues(request *DescribeHistoryMonitorValuesRequest) (_result *DescribeHistoryMonitorValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHistoryMonitorValuesResponse{}
	_body, _err := client.DescribeHistoryMonitorValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tasks in the task center.
//
// @param request - DescribeHistoryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksResponse
func (client *Client) DescribeHistoryTasksWithOptions(request *DescribeHistoryTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromExecTime) {
		query["FromExecTime"] = request.FromExecTime
	}

	if !dara.IsNil(request.FromStartTime) {
		query["FromStartTime"] = request.FromStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.ToExecTime) {
		query["ToExecTime"] = request.ToExecTime
	}

	if !dara.IsNil(request.ToStartTime) {
		query["ToStartTime"] = request.ToStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryTasks"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tasks in the task center.
//
// @param request - DescribeHistoryTasksRequest
//
// @return DescribeHistoryTasksResponse
func (client *Client) DescribeHistoryTasks(request *DescribeHistoryTasksRequest) (_result *DescribeHistoryTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHistoryTasksResponse{}
	_body, _err := client.DescribeHistoryTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attribute of Tair (Redis OSS-compatible) instances.
//
// @param request - DescribeInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAttributeResponse
func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attribute of Tair (Redis OSS-compatible) instances.
//
// @param request - DescribeInstanceAttributeRequest
//
// @return DescribeInstanceAttributeResponse
func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (_result *DescribeInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DescribeInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether auto-renewal is enabled for a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeInstanceAutoRenewalAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAutoRenewalAttributeResponse
func (client *Client) DescribeInstanceAutoRenewalAttributeWithOptions(request *DescribeInstanceAutoRenewalAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAutoRenewalAttributeResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAutoRenewalAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether auto-renewal is enabled for a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeInstanceAutoRenewalAttributeRequest
//
// @return DescribeInstanceAutoRenewalAttributeResponse
func (client *Client) DescribeInstanceAutoRenewalAttribute(request *DescribeInstanceAutoRenewalAttributeRequest) (_result *DescribeInstanceAutoRenewalAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.DescribeInstanceAutoRenewalAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default parameter configurations of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// This operation is available only for instances that use cloud disks.
//
// > You can call the [DescribeParameters](https://help.aliyun.com/document_detail/473847.html) operation to query the parameter settings of instances that use local disks.
//
// @param request - DescribeInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceConfigResponse
func (client *Client) DescribeInstanceConfigWithOptions(request *DescribeInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceConfig"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default parameter configurations of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// This operation is available only for instances that use cloud disks.
//
// > You can call the [DescribeParameters](https://help.aliyun.com/document_detail/473847.html) operation to query the parameter settings of instances that use local disks.
//
// @param request - DescribeInstanceConfigRequest
//
// @return DescribeInstanceConfigResponse
func (client *Client) DescribeInstanceConfig(request *DescribeInstanceConfigRequest) (_result *DescribeInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceConfigResponse{}
	_body, _err := client.DescribeInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether TLS (SSL) encryption is enabled for an instance.
//
// Description:
//
// SSL encryption is supported for Tair (Redis OSS-compatible) 2.8 standard master-replica instances, Tair (Redis OSS-compatible) 2.8 master-replica cluster instances, and Tair (Redis OSS-compatible) 4.0 master-replica cluster instances. You can enable SSL encryption to enhance data transmission security.
//
// You can use one of the following methods to enable or disable SSL encryption or update the SSL certificate for a Tair (Redis OSS-compatible) instance:
//
//   - Call the [ModifyInstanceSSL](https://help.aliyun.com/document_detail/473838.html) operation.
//
//   - Enable or disable SSL encryption or update the SSL certificate in the Tair (Redis OSS-compatible) console. For more information, see [Configure SSL encryption](https://help.aliyun.com/document_detail/84898.html).
//
// > After SSL encryption is enabled, the instance may respond slower.
//
// @param request - DescribeInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSSLResponse
func (client *Client) DescribeInstanceSSLWithOptions(request *DescribeInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSSLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSSL"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether TLS (SSL) encryption is enabled for an instance.
//
// Description:
//
// SSL encryption is supported for Tair (Redis OSS-compatible) 2.8 standard master-replica instances, Tair (Redis OSS-compatible) 2.8 master-replica cluster instances, and Tair (Redis OSS-compatible) 4.0 master-replica cluster instances. You can enable SSL encryption to enhance data transmission security.
//
// You can use one of the following methods to enable or disable SSL encryption or update the SSL certificate for a Tair (Redis OSS-compatible) instance:
//
//   - Call the [ModifyInstanceSSL](https://help.aliyun.com/document_detail/473838.html) operation.
//
//   - Enable or disable SSL encryption or update the SSL certificate in the Tair (Redis OSS-compatible) console. For more information, see [Configure SSL encryption](https://help.aliyun.com/document_detail/84898.html).
//
// > After SSL encryption is enabled, the instance may respond slower.
//
// @param request - DescribeInstanceSSLRequest
//
// @return DescribeInstanceSSLResponse
func (client *Client) DescribeInstanceSSL(request *DescribeInstanceSSLRequest) (_result *DescribeInstanceSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceSSLResponse{}
	_body, _err := client.DescribeInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether transparent data encryption (TDE) is enabled for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// For more information about TDE and the usage notes of TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
// >  You can call the [ModifyInstanceTDE](https://help.aliyun.com/document_detail/473859.html) to enable or disable TDE.
//
// @param request - DescribeInstanceTDEStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceTDEStatusResponse
func (client *Client) DescribeInstanceTDEStatusWithOptions(request *DescribeInstanceTDEStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceTDEStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceTDEStatus"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceTDEStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether transparent data encryption (TDE) is enabled for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// For more information about TDE and the usage notes of TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
// >  You can call the [ModifyInstanceTDE](https://help.aliyun.com/document_detail/473859.html) to enable or disable TDE.
//
// @param request - DescribeInstanceTDEStatusRequest
//
// @return DescribeInstanceTDEStatusResponse
func (client *Client) DescribeInstanceTDEStatus(request *DescribeInstanceTDEStatusRequest) (_result *DescribeInstanceTDEStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceTDEStatusResponse{}
	_body, _err := client.DescribeInstanceTDEStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more Tair (Redis OSS-compatible) instances.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArchitectureType) {
		query["ArchitectureType"] = request.ArchitectureType
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.EditionType) {
		query["EditionType"] = request.EditionType
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
	}

	if !dara.IsNil(request.GlobalInstance) {
		query["GlobalInstance"] = request.GlobalInstance
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
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

	if !dara.IsNil(request.PrivateIp) {
		query["PrivateIp"] = request.PrivateIp
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more Tair (Redis OSS-compatible) instances.
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the overview information of one or more Tair (Redis OSS-compatible) instances.
//
// Description:
//
// If you do not specify the InstanceIds parameter when you call this operation, the overview information of all instances is returned.
//
// > This operation returns non-paged results.
//
// @param request - DescribeInstancesOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesOverviewResponse
func (client *Client) DescribeInstancesOverviewWithOptions(request *DescribeInstancesOverviewRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesOverviewResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArchitectureType) {
		query["ArchitectureType"] = request.ArchitectureType
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.EditionType) {
		query["EditionType"] = request.EditionType
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PrivateIp) {
		query["PrivateIp"] = request.PrivateIp
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstancesOverview"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the overview information of one or more Tair (Redis OSS-compatible) instances.
//
// Description:
//
// If you do not specify the InstanceIds parameter when you call this operation, the overview information of all instances is returned.
//
// > This operation returns non-paged results.
//
// @param request - DescribeInstancesOverviewRequest
//
// @return DescribeInstancesOverviewResponse
func (client *Client) DescribeInstancesOverview(request *DescribeInstancesOverviewRequest) (_result *DescribeInstancesOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstancesOverviewResponse{}
	_body, _err := client.DescribeInstancesOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the internal bandwidth of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can call the [EnableAdditionalBandwidth](https://help.aliyun.com/document_detail/473771.html) operation to increase the internal bandwidth of an instance.
//
// @param request - DescribeIntranetAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIntranetAttributeResponse
func (client *Client) DescribeIntranetAttributeWithOptions(request *DescribeIntranetAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeIntranetAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIntranetAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIntranetAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the internal bandwidth of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can call the [EnableAdditionalBandwidth](https://help.aliyun.com/document_detail/473771.html) operation to increase the internal bandwidth of an instance.
//
// @param request - DescribeIntranetAttributeRequest
//
// @return DescribeIntranetAttributeResponse
func (client *Client) DescribeIntranetAttribute(request *DescribeIntranetAttributeRequest) (_result *DescribeIntranetAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIntranetAttributeResponse{}
	_body, _err := client.DescribeIntranetAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logical topology of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// This parameter is supported only for cluster and read/write splitting instances.
//
// @param request - DescribeLogicInstanceTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogicInstanceTopologyResponse
func (client *Client) DescribeLogicInstanceTopologyWithOptions(request *DescribeLogicInstanceTopologyRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogicInstanceTopologyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogicInstanceTopology"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogicInstanceTopologyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logical topology of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// This parameter is supported only for cluster and read/write splitting instances.
//
// @param request - DescribeLogicInstanceTopologyRequest
//
// @return DescribeLogicInstanceTopologyResponse
func (client *Client) DescribeLogicInstanceTopology(request *DescribeLogicInstanceTopologyRequest) (_result *DescribeLogicInstanceTopologyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogicInstanceTopologyResponse{}
	_body, _err := client.DescribeLogicInstanceTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metrics of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >  To improve user experience, Tair has upgraded the monitoring metrics. The DescribeMonitorItems operation is phased out. For more information, see [The DescribeMonitorItems operation of Tair (Redis OSS-compatible) is phased out](https://help.aliyun.com/document_detail/189893.html).
//
// After you call this operation to retrieve a list of metrics for a specified instance, you can call the [DescribeHistoryMonitorValues](https://help.aliyun.com/document_detail/473827.html) operation to query the monitoring history of the instance.
//
// @param request - DescribeMonitorItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorItemsResponse
func (client *Client) DescribeMonitorItemsWithOptions(request *DescribeMonitorItemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorItemsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorItems"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metrics of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >  To improve user experience, Tair has upgraded the monitoring metrics. The DescribeMonitorItems operation is phased out. For more information, see [The DescribeMonitorItems operation of Tair (Redis OSS-compatible) is phased out](https://help.aliyun.com/document_detail/189893.html).
//
// After you call this operation to retrieve a list of metrics for a specified instance, you can call the [DescribeHistoryMonitorValues](https://help.aliyun.com/document_detail/473827.html) operation to query the monitoring history of the instance.
//
// @param request - DescribeMonitorItemsRequest
//
// @return DescribeMonitorItemsResponse
func (client *Client) DescribeMonitorItems(request *DescribeMonitorItemsRequest) (_result *DescribeMonitorItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorItemsResponse{}
	_body, _err := client.DescribeMonitorItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information about a parameter template.
//
// @param request - DescribeParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterGroupResponse
func (client *Client) DescribeParameterGroupWithOptions(request *DescribeParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about a parameter template.
//
// @param request - DescribeParameterGroupRequest
//
// @return DescribeParameterGroupResponse
func (client *Client) DescribeParameterGroup(request *DescribeParameterGroupRequest) (_result *DescribeParameterGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterGroupResponse{}
	_body, _err := client.DescribeParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters that can be configured in parameter templates across different database versions.
//
// @param request - DescribeParameterGroupSupportParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterGroupSupportParamResponse
func (client *Client) DescribeParameterGroupSupportParamWithOptions(request *DescribeParameterGroupSupportParamRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterGroupSupportParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterGroupSupportParam"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterGroupSupportParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameters that can be configured in parameter templates across different database versions.
//
// @param request - DescribeParameterGroupSupportParamRequest
//
// @return DescribeParameterGroupSupportParamResponse
func (client *Client) DescribeParameterGroupSupportParam(request *DescribeParameterGroupSupportParamRequest) (_result *DescribeParameterGroupSupportParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterGroupSupportParamResponse{}
	_body, _err := client.DescribeParameterGroupSupportParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the parameters that can be configured in a parameter template, such as the default values, value ranges, and descriptions.
//
// @param request - DescribeParameterGroupTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterGroupTemplateListResponse
func (client *Client) DescribeParameterGroupTemplateListWithOptions(request *DescribeParameterGroupTemplateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterGroupTemplateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterGroupTemplateList"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterGroupTemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the parameters that can be configured in a parameter template, such as the default values, value ranges, and descriptions.
//
// @param request - DescribeParameterGroupTemplateListRequest
//
// @return DescribeParameterGroupTemplateListResponse
func (client *Client) DescribeParameterGroupTemplateList(request *DescribeParameterGroupTemplateListRequest) (_result *DescribeParameterGroupTemplateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterGroupTemplateListResponse{}
	_body, _err := client.DescribeParameterGroupTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of available parameter templates.
//
// @param request - DescribeParameterGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterGroupsResponse
func (client *Client) DescribeParameterGroupsWithOptions(request *DescribeParameterGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterGroups"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of available parameter templates.
//
// @param request - DescribeParameterGroupsRequest
//
// @return DescribeParameterGroupsResponse
func (client *Client) DescribeParameterGroups(request *DescribeParameterGroupsRequest) (_result *DescribeParameterGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterGroupsResponse{}
	_body, _err := client.DescribeParameterGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameter modification history of a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeParameterModificationHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterModificationHistoryResponse
func (client *Client) DescribeParameterModificationHistoryWithOptions(request *DescribeParameterModificationHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterModificationHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterName) {
		query["ParameterName"] = request.ParameterName
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterModificationHistory"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterModificationHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameter modification history of a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeParameterModificationHistoryRequest
//
// @return DescribeParameterModificationHistoryResponse
func (client *Client) DescribeParameterModificationHistory(request *DescribeParameterModificationHistoryRequest) (_result *DescribeParameterModificationHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParameterModificationHistoryResponse{}
	_body, _err := client.DescribeParameterModificationHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters and their default values that are supported by Tair (Redis OSS-compatible) instances of different architectures and major versions.
//
// Description:
//
// After you call this operation to query the parameters and default values of an instance, you can call the [ModifyInstanceConfig](https://help.aliyun.com/document_detail/473844.html) operation to reconfigure the parameters of the instance.
//
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
	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterTemplates"),
		Version:     dara.String("2015-01-01"),
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

// Summary:
//
// Queries the parameters and their default values that are supported by Tair (Redis OSS-compatible) instances of different architectures and major versions.
//
// Description:
//
// After you call this operation to query the parameters and default values of an instance, you can call the [ModifyInstanceConfig](https://help.aliyun.com/document_detail/473844.html) operation to reconfigure the parameters of the instance.
//
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

// Summary:
//
// Queries the configuration parameters and running parameters of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// This operation is applicable only to classic instances.
//
// >  If the instance is deployed in cloud-native mode, you can use the [DescribeInstanceConfig](https://help.aliyun.com/document_detail/473846.html) operation to query the configuration and operational parameters of the instance.
//
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameters"),
		Version:     dara.String("2015-01-01"),
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

// Summary:
//
// Queries the configuration parameters and running parameters of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// This operation is applicable only to classic instances.
//
// >  If the instance is deployed in cloud-native mode, you can use the [DescribeInstanceConfig](https://help.aliyun.com/document_detail/473846.html) operation to query the configuration and operational parameters of the instance.
//
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

// Summary:
//
// Queries the fees that you must pay when you create, upgrade, or renew a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceResponse
func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.ForceUpgrade) {
		query["ForceUpgrade"] = request.ForceUpgrade
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OrderParamOut) {
		query["OrderParamOut"] = request.OrderParamOut
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
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

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ShardCount) {
		query["ShardCount"] = request.ShardCount
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrice"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fees that you must pay when you create, upgrade, or renew a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribePriceRequest
//
// @return DescribePriceResponse
func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions in which ApsaraDB for Redis instances can be created.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2015-01-01"),
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
// Queries the regions in which ApsaraDB for Redis instances can be created.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the role, type, minor version, and zone of each node in a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeRoleZoneInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoleZoneInfoResponse
func (client *Client) DescribeRoleZoneInfoWithOptions(request *DescribeRoleZoneInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoleZoneInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoleZoneInfo"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoleZoneInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the role, type, minor version, and zone of each node in a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeRoleZoneInfoRequest
//
// @return DescribeRoleZoneInfoResponse
func (client *Client) DescribeRoleZoneInfo(request *DescribeRoleZoneInfoRequest) (_result *DescribeRoleZoneInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRoleZoneInfoResponse{}
	_body, _err := client.DescribeRoleZoneInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operational logs of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// For more information about how to view the operational logs of an instance in the Tair (Redis OSS-compatible) console, see [View active logs](https://help.aliyun.com/document_detail/101713.html).
//
// This operation can be called up to 100 times per minute.
//
// @param request - DescribeRunningLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRunningLogRecordsResponse
func (client *Client) DescribeRunningLogRecordsWithOptions(request *DescribeRunningLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRunningLogRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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

	if !dara.IsNil(request.QueryKeyword) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRunningLogRecords"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRunningLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operational logs of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// For more information about how to view the operational logs of an instance in the Tair (Redis OSS-compatible) console, see [View active logs](https://help.aliyun.com/document_detail/101713.html).
//
// This operation can be called up to 100 times per minute.
//
// @param request - DescribeRunningLogRecordsRequest
//
// @return DescribeRunningLogRecordsResponse
func (client *Client) DescribeRunningLogRecords(request *DescribeRunningLogRecordsRequest) (_result *DescribeRunningLogRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRunningLogRecordsResponse{}
	_body, _err := client.DescribeRunningLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security groups that are added to the whitelists of a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeSecurityGroupConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityGroupConfigurationResponse
func (client *Client) DescribeSecurityGroupConfigurationWithOptions(request *DescribeSecurityGroupConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityGroupConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityGroupConfiguration"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityGroupConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security groups that are added to the whitelists of a Tair (Redis OSS-compatible) instance.
//
// @param request - DescribeSecurityGroupConfigurationRequest
//
// @return DescribeSecurityGroupConfigurationResponse
func (client *Client) DescribeSecurityGroupConfiguration(request *DescribeSecurityGroupConfigurationRequest) (_result *DescribeSecurityGroupConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityGroupConfigurationResponse{}
	_body, _err := client.DescribeSecurityGroupConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists of a Tair (Redis OSS-compatible) instance.
//
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityIps"),
		Version:     dara.String("2015-01-01"),
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

// Summary:
//
// Queries the IP address whitelists of a Tair (Redis OSS-compatible) instance.
//
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
// Queries the slow query logs of a Tair (Redis OSS-compatible) instance that are generated within a specified period of time.
//
// Description:
//
// You can also query slow logs in the Tair (Redis OSS-compatible) console. For more information, see [Query slow logs of an instance](https://help.aliyun.com/document_detail/95874.html). This operation can be called up to 100 times per minute.
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
	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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

	if !dara.IsNil(request.QueryKeyword) {
		query["QueryKeyword"] = request.QueryKeyword
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SlowLogRecordType) {
		query["SlowLogRecordType"] = request.SlowLogRecordType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowLogRecords"),
		Version:     dara.String("2015-01-01"),
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
// Queries the slow query logs of a Tair (Redis OSS-compatible) instance that are generated within a specified period of time.
//
// Description:
//
// You can also query slow logs in the Tair (Redis OSS-compatible) console. For more information, see [Query slow logs of an instance](https://help.aliyun.com/document_detail/95874.html). This operation can be called up to 100 times per minute.
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
// 查看TairCustom实例
//
// @param request - DescribeTairKVCacheCustomInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTairKVCacheCustomInstanceAttributeResponse
func (client *Client) DescribeTairKVCacheCustomInstanceAttributeWithOptions(request *DescribeTairKVCacheCustomInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeTairKVCacheCustomInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTairKVCacheCustomInstanceAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTairKVCacheCustomInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看TairCustom实例
//
// @param request - DescribeTairKVCacheCustomInstanceAttributeRequest
//
// @return DescribeTairKVCacheCustomInstanceAttributeResponse
func (client *Client) DescribeTairKVCacheCustomInstanceAttribute(request *DescribeTairKVCacheCustomInstanceAttributeRequest) (_result *DescribeTairKVCacheCustomInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTairKVCacheCustomInstanceAttributeResponse{}
	_body, _err := client.DescribeTairKVCacheCustomInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询TairCustom主机监控
//
// @param request - DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse
func (client *Client) DescribeTairKVCacheCustomInstanceHistoryMonitorValuesWithOptions(request *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTairKVCacheCustomInstanceHistoryMonitorValues"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询TairCustom主机监控
//
// @param request - DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest
//
// @return DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse
func (client *Client) DescribeTairKVCacheCustomInstanceHistoryMonitorValues(request *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesRequest) (_result *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse{}
	_body, _err := client.DescribeTairKVCacheCustomInstanceHistoryMonitorValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看TairCustom实例
//
// @param request - DescribeTairKVCacheCustomInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTairKVCacheCustomInstancesResponse
func (client *Client) DescribeTairKVCacheCustomInstancesWithOptions(request *DescribeTairKVCacheCustomInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTairKVCacheCustomInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTairKVCacheCustomInstances"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTairKVCacheCustomInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看TairCustom实例
//
// @param request - DescribeTairKVCacheCustomInstancesRequest
//
// @return DescribeTairKVCacheCustomInstancesResponse
func (client *Client) DescribeTairKVCacheCustomInstances(request *DescribeTairKVCacheCustomInstancesRequest) (_result *DescribeTairKVCacheCustomInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTairKVCacheCustomInstancesResponse{}
	_body, _err := client.DescribeTairKVCacheCustomInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看TairInfer实例
//
// @param request - DescribeTairKVCacheInferInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTairKVCacheInferInstanceAttributeResponse
func (client *Client) DescribeTairKVCacheInferInstanceAttributeWithOptions(request *DescribeTairKVCacheInferInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeTairKVCacheInferInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTairKVCacheInferInstanceAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTairKVCacheInferInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看TairInfer实例
//
// @param request - DescribeTairKVCacheInferInstanceAttributeRequest
//
// @return DescribeTairKVCacheInferInstanceAttributeResponse
func (client *Client) DescribeTairKVCacheInferInstanceAttribute(request *DescribeTairKVCacheInferInstanceAttributeRequest) (_result *DescribeTairKVCacheInferInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTairKVCacheInferInstanceAttributeResponse{}
	_body, _err := client.DescribeTairKVCacheInferInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看TairInfer实例列表
//
// @param request - DescribeTairKVCacheInferInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTairKVCacheInferInstancesResponse
func (client *Client) DescribeTairKVCacheInferInstancesWithOptions(request *DescribeTairKVCacheInferInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTairKVCacheInferInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTairKVCacheInferInstances"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTairKVCacheInferInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看TairInfer实例列表
//
// @param request - DescribeTairKVCacheInferInstancesRequest
//
// @return DescribeTairKVCacheInferInstancesResponse
func (client *Client) DescribeTairKVCacheInferInstances(request *DescribeTairKVCacheInferInstancesRequest) (_result *DescribeTairKVCacheInferInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTairKVCacheInferInstancesResponse{}
	_body, _err := client.DescribeTairKVCacheInferInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the zones available for Tair (Redis OSS-compatible).
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones available for Tair (Redis OSS-compatible).
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adjusts the bandwidth of a Tair (Redis OSS-compatible) instance. Only the pay-as-you-go billing method is supported for bandwidth adjustment. You need to specify the InstanceId, NodeId (optional), Bandwidth, and ChargeType parameters.
//
// Description:
//
// If you enable the bandwidth auto scaling feature and call this operation at the same time, bandwidth auto scaling takes precedence. During bandwidth scale-back, the instance is scaled back to the default bandwidth of the instance type. For more information about the limits, costs, and FAQ about this feature, see [Adjust the bandwidth of an instance](https://help.aliyun.com/document_detail/102588.html).
//
// >  Before you call this operation, you can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/473782.html) operation to query the current bandwidth of each data node in an instance.
//
// @param request - EnableAdditionalBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAdditionalBandwidthResponse
func (client *Client) EnableAdditionalBandwidthWithOptions(request *EnableAdditionalBandwidthRequest, runtime *dara.RuntimeOptions) (_result *EnableAdditionalBandwidthResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderTimeLength) {
		query["OrderTimeLength"] = request.OrderTimeLength
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAdditionalBandwidth"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAdditionalBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adjusts the bandwidth of a Tair (Redis OSS-compatible) instance. Only the pay-as-you-go billing method is supported for bandwidth adjustment. You need to specify the InstanceId, NodeId (optional), Bandwidth, and ChargeType parameters.
//
// Description:
//
// If you enable the bandwidth auto scaling feature and call this operation at the same time, bandwidth auto scaling takes precedence. During bandwidth scale-back, the instance is scaled back to the default bandwidth of the instance type. For more information about the limits, costs, and FAQ about this feature, see [Adjust the bandwidth of an instance](https://help.aliyun.com/document_detail/102588.html).
//
// >  Before you call this operation, you can call the [DescribeRoleZoneInfo](https://help.aliyun.com/document_detail/473782.html) operation to query the current bandwidth of each data node in an instance.
//
// @param request - EnableAdditionalBandwidthRequest
//
// @return EnableAdditionalBandwidthResponse
func (client *Client) EnableAdditionalBandwidth(request *EnableAdditionalBandwidthRequest) (_result *EnableAdditionalBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableAdditionalBandwidthResponse{}
	_body, _err := client.EnableAdditionalBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears all expired keys in a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// For more information about how to clear the expired keys in the Tair (Redis OSS-compatible) console, see [Clear data](https://help.aliyun.com/document_detail/43881.html).
//
// >  Expired keys cannot be recovered after they are deleted. Exercise caution when you call this operation.
//
// @param request - FlushExpireKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlushExpireKeysResponse
func (client *Client) FlushExpireKeysWithOptions(request *FlushExpireKeysRequest, runtime *dara.RuntimeOptions) (_result *FlushExpireKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlushExpireKeys"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlushExpireKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears all expired keys in a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// For more information about how to clear the expired keys in the Tair (Redis OSS-compatible) console, see [Clear data](https://help.aliyun.com/document_detail/43881.html).
//
// >  Expired keys cannot be recovered after they are deleted. Exercise caution when you call this operation.
//
// @param request - FlushExpireKeysRequest
//
// @return FlushExpireKeysResponse
func (client *Client) FlushExpireKeys(request *FlushExpireKeysRequest) (_result *FlushExpireKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FlushExpireKeysResponse{}
	_body, _err := client.FlushExpireKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears the data of a Tair (Redis OSS-compatible) instance. The cleared data cannot be restored.
//
// @param request - FlushInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlushInstanceResponse
func (client *Client) FlushInstanceWithOptions(request *FlushInstanceRequest, runtime *dara.RuntimeOptions) (_result *FlushInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FlushInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlushInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the data of a Tair (Redis OSS-compatible) instance. The cleared data cannot be restored.
//
// @param request - FlushInstanceRequest
//
// @return FlushInstanceResponse
func (client *Client) FlushInstance(request *FlushInstanceRequest) (_result *FlushInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FlushInstanceResponse{}
	_body, _err := client.FlushInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cleans the data of specified databases in a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// Each Tair (Redis OSS-compatible) instance can contain up to 256 databases named from DB0 to DB255. Each database does not have a separate memory usage limit. The memory capacity that a database can use is subject to the total memory limit of the instance. You can execute the `SELECT` statement to switch between databases. For more information, see [What is the size of each database on a Tair (Redis OSS-compatible) instance, and how can I choose databases?](https://help.aliyun.com/document_detail/38688.html)
//
// >  This operation is available only for cloud-native instances that use cloud disks.
//
// @param request - FlushInstanceForDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlushInstanceForDBResponse
func (client *Client) FlushInstanceForDBWithOptions(request *FlushInstanceForDBRequest, runtime *dara.RuntimeOptions) (_result *FlushInstanceForDBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbIndex) {
		query["DbIndex"] = request.DbIndex
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("FlushInstanceForDB"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FlushInstanceForDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cleans the data of specified databases in a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// Each Tair (Redis OSS-compatible) instance can contain up to 256 databases named from DB0 to DB255. Each database does not have a separate memory usage limit. The memory capacity that a database can use is subject to the total memory limit of the instance. You can execute the `SELECT` statement to switch between databases. For more information, see [What is the size of each database on a Tair (Redis OSS-compatible) instance, and how can I choose databases?](https://help.aliyun.com/document_detail/38688.html)
//
// >  This operation is available only for cloud-native instances that use cloud disks.
//
// @param request - FlushInstanceForDBRequest
//
// @return FlushInstanceForDBResponse
func (client *Client) FlushInstanceForDB(request *FlushInstanceForDBRequest) (_result *FlushInstanceForDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FlushInstanceForDBResponse{}
	_body, _err := client.FlushInstanceForDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the permissions of an account for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >
//
//   - Only Tair (Redis OSS-compatible) instances of Redis 4.0 or later are supported.
//
//   - The Tair (Redis OSS-compatible) instance must be in the running state.
//
// @param request - GrantAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantAccountPrivilegeResponse
func (client *Client) GrantAccountPrivilegeWithOptions(request *GrantAccountPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *GrantAccountPrivilegeResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantAccountPrivilege"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the permissions of an account for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >
//
//   - Only Tair (Redis OSS-compatible) instances of Redis 4.0 or later are supported.
//
//   - The Tair (Redis OSS-compatible) instance must be in the running state.
//
// @param request - GrantAccountPrivilegeRequest
//
// @return GrantAccountPrivilegeResponse
func (client *Client) GrantAccountPrivilege(request *GrantAccountPrivilegeRequest) (_result *GrantAccountPrivilegeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.GrantAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns a service-linked role to Tair (Redis OSS-compatible).
//
// Description:
//
// The log management feature of Tair (Redis OSS-compatible) requires the resources of [Simple Log Service](https://help.aliyun.com/document_detail/48869.html). To use the log management feature, you can call this operation to assign the AliyunServiceRoleForKvstore service-linked role to Tair (Redis OSS-compatible). For more information, see [Service-linked role of Tair (Redis OSS-compatible)](https://help.aliyun.com/document_detail/184337.html).
//
// @param request - InitializeKvstorePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeKvstorePermissionResponse
func (client *Client) InitializeKvstorePermissionWithOptions(request *InitializeKvstorePermissionRequest, runtime *dara.RuntimeOptions) (_result *InitializeKvstorePermissionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitializeKvstorePermission"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeKvstorePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns a service-linked role to Tair (Redis OSS-compatible).
//
// Description:
//
// The log management feature of Tair (Redis OSS-compatible) requires the resources of [Simple Log Service](https://help.aliyun.com/document_detail/48869.html). To use the log management feature, you can call this operation to assign the AliyunServiceRoleForKvstore service-linked role to Tair (Redis OSS-compatible). For more information, see [Service-linked role of Tair (Redis OSS-compatible)](https://help.aliyun.com/document_detail/184337.html).
//
// @param request - InitializeKvstorePermissionRequest
//
// @return InitializeKvstorePermissionResponse
func (client *Client) InitializeKvstorePermission(request *InitializeKvstorePermissionRequest) (_result *InitializeKvstorePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitializeKvstorePermissionResponse{}
	_body, _err := client.InitializeKvstorePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the relationships between Tair (Redis OSS-compatible) instances and tags.
//
// Description:
//
// You can also query the relationships between instances and tags in the Tair (Redis OSS-compatible) console. For more information, see [Filter Tair (Redis OSS-compatible) instances by tag](https://help.aliyun.com/document_detail/119160.html) and [View tags bound to an instance](https://help.aliyun.com/document_detail/134038.html).
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2015-01-01"),
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
// Queries the relationships between Tair (Redis OSS-compatible) instances and tags.
//
// Description:
//
// You can also query the relationships between instances and tags in the Tair (Redis OSS-compatible) console. For more information, see [Filter Tair (Redis OSS-compatible) instances by tag](https://help.aliyun.com/document_detail/119160.html) and [View tags bound to an instance](https://help.aliyun.com/document_detail/134038.html).
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

// Summary:
//
// Places a write lock on an instance. After the instance is locked, it supports only read operations.
//
// @param request - LockDBInstanceWriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockDBInstanceWriteResponse
func (client *Client) LockDBInstanceWriteWithOptions(request *LockDBInstanceWriteRequest, runtime *dara.RuntimeOptions) (_result *LockDBInstanceWriteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.LockReason) {
		query["LockReason"] = request.LockReason
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LockDBInstanceWrite"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LockDBInstanceWriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Places a write lock on an instance. After the instance is locked, it supports only read operations.
//
// @param request - LockDBInstanceWriteRequest
//
// @return LockDBInstanceWriteResponse
func (client *Client) LockDBInstanceWrite(request *LockDBInstanceWriteRequest) (_result *LockDBInstanceWriteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LockDBInstanceWriteResponse{}
	_body, _err := client.LockDBInstanceWriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Simulates database node failures.
//
// @param request - MasterNodeShutDownFailOverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MasterNodeShutDownFailOverResponse
func (client *Client) MasterNodeShutDownFailOverWithOptions(request *MasterNodeShutDownFailOverRequest, runtime *dara.RuntimeOptions) (_result *MasterNodeShutDownFailOverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.DBFaultMode) {
		query["DBFaultMode"] = request.DBFaultMode
	}

	if !dara.IsNil(request.DBNodes) {
		query["DBNodes"] = request.DBNodes
	}

	if !dara.IsNil(request.FailMode) {
		query["FailMode"] = request.FailMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProxyFaultMode) {
		query["ProxyFaultMode"] = request.ProxyFaultMode
	}

	if !dara.IsNil(request.ProxyInstanceIds) {
		query["ProxyInstanceIds"] = request.ProxyInstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MasterNodeShutDownFailOver"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MasterNodeShutDownFailOverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Simulates database node failures.
//
// @param request - MasterNodeShutDownFailOverRequest
//
// @return MasterNodeShutDownFailOverResponse
func (client *Client) MasterNodeShutDownFailOver(request *MasterNodeShutDownFailOverRequest) (_result *MasterNodeShutDownFailOverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MasterNodeShutDownFailOverResponse{}
	_body, _err := client.MasterNodeShutDownFailOverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Migrates a Tair (Redis OSS-compatible) instance to another zone in the same region.
//
// Description:
//
// Before you call this operation, you must release the public endpoint (if any) of the instance. For more information, see [Migrate an instance across zones](https://help.aliyun.com/document_detail/106272.html).
//
// >
//
//   - If the network type of an Tair (Redis OSS-compatible) instance is switched from classic network to Virtual Private Cloud (VPC), and the classic network endpoint is retained, you can migrate the instance across zones only after the classic network endpoint is released upon expiration.
//
//   - After the instance is migrated, the endpoint of the instance remains unchanged. However, the virtual IP address (VIP) is changed. We recommend that you use the endpoint instead of the VIP to connect to the instance.
//
// @param request - MigrateToOtherZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateToOtherZoneResponse
func (client *Client) MigrateToOtherZoneWithOptions(request *MigrateToOtherZoneRequest, runtime *dara.RuntimeOptions) (_result *MigrateToOtherZoneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadOnlyCount) {
		query["ReadOnlyCount"] = request.ReadOnlyCount
	}

	if !dara.IsNil(request.ReplicaCount) {
		query["ReplicaCount"] = request.ReplicaCount
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SlaveReadOnlyCount) {
		query["SlaveReadOnlyCount"] = request.SlaveReadOnlyCount
	}

	if !dara.IsNil(request.SlaveReplicaCount) {
		query["SlaveReplicaCount"] = request.SlaveReplicaCount
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
		Action:      dara.String("MigrateToOtherZone"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates a Tair (Redis OSS-compatible) instance to another zone in the same region.
//
// Description:
//
// Before you call this operation, you must release the public endpoint (if any) of the instance. For more information, see [Migrate an instance across zones](https://help.aliyun.com/document_detail/106272.html).
//
// >
//
//   - If the network type of an Tair (Redis OSS-compatible) instance is switched from classic network to Virtual Private Cloud (VPC), and the classic network endpoint is retained, you can migrate the instance across zones only after the classic network endpoint is released upon expiration.
//
//   - After the instance is migrated, the endpoint of the instance remains unchanged. However, the virtual IP address (VIP) is changed. We recommend that you use the endpoint instead of the VIP to connect to the instance.
//
// @param request - MigrateToOtherZoneRequest
//
// @return MigrateToOtherZoneResponse
func (client *Client) MigrateToOtherZone(request *MigrateToOtherZoneRequest) (_result *MigrateToOtherZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.MigrateToOtherZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an account for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// This operation is supported only for instances that run Redis 4.0 or later.
//
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2015-01-01"),
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
// Modifies the description of an account for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// This operation is supported only for instances that run Redis 4.0 or later.
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
// Changes the password of a specific account for a Tair (Redis OSS-compatible) instance.
//
// @param request - ModifyAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPasswordResponse
func (client *Client) ModifyAccountPasswordWithOptions(request *ModifyAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NewAccountPassword) {
		query["NewAccountPassword"] = request.NewAccountPassword
	}

	if !dara.IsNil(request.OldAccountPassword) {
		query["OldAccountPassword"] = request.OldAccountPassword
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountPassword"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the password of a specific account for a Tair (Redis OSS-compatible) instance.
//
// @param request - ModifyAccountPasswordRequest
//
// @return ModifyAccountPasswordResponse
func (client *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (_result *ModifyAccountPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.ModifyAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the scheduled switchover time of an O&M task.
//
// Description:
//
// You can receive notifications for Tair (Redis OSS-compatible) events such as instance migration and version upgrade by text message, phone call, email, internal message, or by using the console. You can also change the scheduled switchover time of a task by using the console. For more information, see [Query or manage pending events](https://help.aliyun.com/document_detail/187022.html).
//
// @param request - ModifyActiveOperationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationTaskResponse
func (client *Client) ModifyActiveOperationTaskWithOptions(request *ModifyActiveOperationTaskRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationTask"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActiveOperationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the scheduled switchover time of an O&M task.
//
// Description:
//
// You can receive notifications for Tair (Redis OSS-compatible) events such as instance migration and version upgrade by text message, phone call, email, internal message, or by using the console. You can also change the scheduled switchover time of a task by using the console. For more information, see [Query or manage pending events](https://help.aliyun.com/document_detail/187022.html).
//
// @param request - ModifyActiveOperationTaskRequest
//
// @return ModifyActiveOperationTaskResponse
func (client *Client) ModifyActiveOperationTask(request *ModifyActiveOperationTaskRequest) (_result *ModifyActiveOperationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyActiveOperationTaskResponse{}
	_body, _err := client.ModifyActiveOperationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the switching time of scheduled O\\\\\\&M events for an instance.
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationTasks"),
		Version:     dara.String("2015-01-01"),
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
// Modifies the switching time of scheduled O\\\\\\&M events for an instance.
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

// Summary:
//
// Enables the audit log feature or modifies the audit log settings for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/54532.html) of the audit log feature.
//
// Before you call this operation, make sure that the Tair (Redis OSS-compatible) instance meets the following requirements:
//
//   - The instance is a Tair (Redis OSS-compatible) Community Edition instance or Tair [DRAM-based instance](https://help.aliyun.com/document_detail/126164.html).
//
//   - The engine version of the instance is Redis 4.0 or later, and the latest minor version is used. You can call the [DescribeEngineVersion](https://help.aliyun.com/document_detail/473781.html) operation to check whether the instance uses the latest major version and minor version.
//
// @param request - ModifyAuditLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAuditLogConfigResponse
func (client *Client) ModifyAuditLogConfigWithOptions(request *ModifyAuditLogConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyAuditLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbAudit) {
		query["DbAudit"] = request.DbAudit
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAuditLogConfig"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the audit log feature or modifies the audit log settings for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/54532.html) of the audit log feature.
//
// Before you call this operation, make sure that the Tair (Redis OSS-compatible) instance meets the following requirements:
//
//   - The instance is a Tair (Redis OSS-compatible) Community Edition instance or Tair [DRAM-based instance](https://help.aliyun.com/document_detail/126164.html).
//
//   - The engine version of the instance is Redis 4.0 or later, and the latest minor version is used. You can call the [DescribeEngineVersion](https://help.aliyun.com/document_detail/473781.html) operation to check whether the instance uses the latest major version and minor version.
//
// @param request - ModifyAuditLogConfigRequest
//
// @return ModifyAuditLogConfigResponse
func (client *Client) ModifyAuditLogConfig(request *ModifyAuditLogConfigRequest) (_result *ModifyAuditLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAuditLogConfigResponse{}
	_body, _err := client.ModifyAuditLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改备份集过期时间
//
// @param request - ModifyBackupExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupExpireTimeResponse
func (client *Client) ModifyBackupExpireTimeWithOptions(request *ModifyBackupExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupExpireTimeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.ExpectExpireTime) {
		query["ExpectExpireTime"] = request.ExpectExpireTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyBackupExpireTime"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupExpireTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改备份集过期时间
//
// @param request - ModifyBackupExpireTimeRequest
//
// @return ModifyBackupExpireTimeResponse
func (client *Client) ModifyBackupExpireTime(request *ModifyBackupExpireTimeRequest) (_result *ModifyBackupExpireTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupExpireTimeResponse{}
	_body, _err := client.ModifyBackupExpireTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the automatic backup policy of a Tair (Redis OSS-compatible) instance.
//
// @param request - ModifyBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPolicyResponse
func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.EnableBackupLog) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PreferredBackupPeriod) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2015-01-01"),
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
// Modifies the automatic backup policy of a Tair (Redis OSS-compatible) instance.
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
// Modifies the setting related to the automatic update of minor versions for an instance.
//
// @param request - ModifyDBInstanceAutoUpgradeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceAutoUpgradeResponse
func (client *Client) ModifyDBInstanceAutoUpgradeWithOptions(request *ModifyDBInstanceAutoUpgradeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceAutoUpgradeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceAutoUpgrade"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceAutoUpgradeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the setting related to the automatic update of minor versions for an instance.
//
// @param request - ModifyDBInstanceAutoUpgradeRequest
//
// @return ModifyDBInstanceAutoUpgradeResponse
func (client *Client) ModifyDBInstanceAutoUpgrade(request *ModifyDBInstanceAutoUpgradeRequest) (_result *ModifyDBInstanceAutoUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceAutoUpgradeResponse{}
	_body, _err := client.ModifyDBInstanceAutoUpgradeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the endpoint or port number of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also modify the endpoint or port number of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Change the endpoint or port number of an instance](https://help.aliyun.com/document_detail/85683.html).
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
	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.IPType) {
		query["IPType"] = request.IPType
	}

	if !dara.IsNil(request.NewConnectionString) {
		query["NewConnectionString"] = request.NewConnectionString
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConnectionString"),
		Version:     dara.String("2015-01-01"),
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
// Changes the endpoint or port number of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also modify the endpoint or port number of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Change the endpoint or port number of an instance](https://help.aliyun.com/document_detail/85683.html).
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
// Modifies a global IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupResponse
func (client *Client) ModifyGlobalSecurityIPGroupWithOptions(request *ModifyGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GIpList) {
		query["GIpList"] = request.GIpList
	}

	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGlobalSecurityIPGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a global IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupRequest
//
// @return ModifyGlobalSecurityIPGroupResponse
func (client *Client) ModifyGlobalSecurityIPGroup(request *ModifyGlobalSecurityIPGroupRequest) (_result *ModifyGlobalSecurityIPGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGlobalSecurityIPGroupResponse{}
	_body, _err := client.ModifyGlobalSecurityIPGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a global IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupNameResponse
func (client *Client) ModifyGlobalSecurityIPGroupNameWithOptions(request *ModifyGlobalSecurityIPGroupNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGlobalSecurityIPGroupName"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a global IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupNameRequest
//
// @return ModifyGlobalSecurityIPGroupNameResponse
func (client *Client) ModifyGlobalSecurityIPGroupName(request *ModifyGlobalSecurityIPGroupNameRequest) (_result *ModifyGlobalSecurityIPGroupNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGlobalSecurityIPGroupNameResponse{}
	_body, _err := client.ModifyGlobalSecurityIPGroupNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a specified instance to a specified IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupRelationResponse
func (client *Client) ModifyGlobalSecurityIPGroupRelationWithOptions(request *ModifyGlobalSecurityIPGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupRelationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGlobalSecurityIPGroupRelation"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a specified instance to a specified IP whitelist template.
//
// @param request - ModifyGlobalSecurityIPGroupRelationRequest
//
// @return ModifyGlobalSecurityIPGroupRelationResponse
func (client *Client) ModifyGlobalSecurityIPGroupRelation(request *ModifyGlobalSecurityIPGroupRelationRequest) (_result *ModifyGlobalSecurityIPGroupRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.ModifyGlobalSecurityIPGroupRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the specific information of a Tair (Redis OSS-compatible) instance, such as the password and the name.
//
// Description:
//
// You can also modify the information of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Change or reset the password](https://help.aliyun.com/document_detail/43874.html).
//
// @param request - ModifyInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceReleaseProtection) {
		query["InstanceReleaseProtection"] = request.InstanceReleaseProtection
	}

	if !dara.IsNil(request.NewPassword) {
		query["NewPassword"] = request.NewPassword
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the specific information of a Tair (Redis OSS-compatible) instance, such as the password and the name.
//
// Description:
//
// You can also modify the information of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Change or reset the password](https://help.aliyun.com/document_detail/43874.html).
//
// @param request - ModifyInstanceAttributeRequest
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables auto-renewal for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// > Auto-renewal is triggered seven days before the expiration date of the instance.
//
// @param request - ModifyInstanceAutoRenewalAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAutoRenewalAttributeResponse
func (client *Client) ModifyInstanceAutoRenewalAttributeWithOptions(request *ModifyInstanceAutoRenewalAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAutoRenewalAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAutoRenewalAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables auto-renewal for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// > Auto-renewal is triggered seven days before the expiration date of the instance.
//
// @param request - ModifyInstanceAutoRenewalAttributeRequest
//
// @return ModifyInstanceAutoRenewalAttributeResponse
func (client *Client) ModifyInstanceAutoRenewalAttribute(request *ModifyInstanceAutoRenewalAttributeRequest) (_result *ModifyInstanceAutoRenewalAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.ModifyInstanceAutoRenewalAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the intended bandwidth value of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
//	  Before you call this operation, make sure that you understand the billing methods and pricing of instance bandwidth. Tair (Redis OSS-compatible) charges fees per hour based on the amount and usage duration of the extra bandwidth that you purchase. The fees vary based on the region that you select. For more information, see [Billable items](https://help.aliyun.com/document_detail/54532.html).
//
//		- The bandwidth of an instance or a shard can be increased by up to six times the default bandwidth of the instance, but the increase in bandwidth cannot exceed 192 Mbit/s. For example, if the default bandwidth of a Tair DRAM-based master-replica instance equipped with 2 GB of memory is 96 Mbit/s, you can increase the bandwidth of the instance by up to 192 Mbit/s. As a result, the maximum bandwidth of the instance is 288 Mbit/s. If the default bandwidth of a Redis Open-Source Edition master-replica instance equipped with 256 MB of memory is 10 Mbit/s, you can increase the bandwidth of the instance by up to 60 Mbit/s. As a result, the maximum bandwidth of the instance is 70 Mbit/s.
//
// @param request - ModifyInstanceBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceBandwidthResponse
func (client *Client) ModifyInstanceBandwidthWithOptions(request *ModifyInstanceBandwidthRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceBandwidthResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TargetIntranetBandwidth) {
		query["TargetIntranetBandwidth"] = request.TargetIntranetBandwidth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceBandwidth"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the intended bandwidth value of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
//	  Before you call this operation, make sure that you understand the billing methods and pricing of instance bandwidth. Tair (Redis OSS-compatible) charges fees per hour based on the amount and usage duration of the extra bandwidth that you purchase. The fees vary based on the region that you select. For more information, see [Billable items](https://help.aliyun.com/document_detail/54532.html).
//
//		- The bandwidth of an instance or a shard can be increased by up to six times the default bandwidth of the instance, but the increase in bandwidth cannot exceed 192 Mbit/s. For example, if the default bandwidth of a Tair DRAM-based master-replica instance equipped with 2 GB of memory is 96 Mbit/s, you can increase the bandwidth of the instance by up to 192 Mbit/s. As a result, the maximum bandwidth of the instance is 288 Mbit/s. If the default bandwidth of a Redis Open-Source Edition master-replica instance equipped with 256 MB of memory is 10 Mbit/s, you can increase the bandwidth of the instance by up to 60 Mbit/s. As a result, the maximum bandwidth of the instance is 70 Mbit/s.
//
// @param request - ModifyInstanceBandwidthRequest
//
// @return ModifyInstanceBandwidthResponse
func (client *Client) ModifyInstanceBandwidth(request *ModifyInstanceBandwidthRequest) (_result *ModifyInstanceBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceBandwidthResponse{}
	_body, _err := client.ModifyInstanceBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the parameter settings of a Tair (Redis OSS-compatible) instance.
//
// @param request - ModifyInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceConfigResponse
func (client *Client) ModifyInstanceConfigWithOptions(request *ModifyInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParamNoLooseSentinelEnabled) {
		query["ParamNoLooseSentinelEnabled"] = request.ParamNoLooseSentinelEnabled
	}

	if !dara.IsNil(request.ParamNoLooseSentinelPasswordFreeAccess) {
		query["ParamNoLooseSentinelPasswordFreeAccess"] = request.ParamNoLooseSentinelPasswordFreeAccess
	}

	if !dara.IsNil(request.ParamNoLooseSentinelPasswordFreeCommands) {
		query["ParamNoLooseSentinelPasswordFreeCommands"] = request.ParamNoLooseSentinelPasswordFreeCommands
	}

	if !dara.IsNil(request.ParamReplMode) {
		query["ParamReplMode"] = request.ParamReplMode
	}

	if !dara.IsNil(request.ParamSemisyncReplTimeout) {
		query["ParamSemisyncReplTimeout"] = request.ParamSemisyncReplTimeout
	}

	if !dara.IsNil(request.ParamSentinelCompatEnable) {
		query["ParamSentinelCompatEnable"] = request.ParamSentinelCompatEnable
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceConfig"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameter settings of a Tair (Redis OSS-compatible) instance.
//
// @param request - ModifyInstanceConfigRequest
//
// @return ModifyInstanceConfigResponse
func (client *Client) ModifyInstanceConfig(request *ModifyInstanceConfigRequest) (_result *ModifyInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.ModifyInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of an Tair (Redis OSS-compatible) instance. Alibaba Cloud maintains Tair (Redis OSS-compatible) instances during the specified maintenance window.
//
// Description:
//
// You can also modify the maintenance window of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Set a maintenance window](https://help.aliyun.com/document_detail/55252.html).
//
// @param request - ModifyInstanceMaintainTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMaintainTimeResponse
func (client *Client) ModifyInstanceMaintainTimeWithOptions(request *ModifyInstanceMaintainTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaintainEndTime) {
		query["MaintainEndTime"] = request.MaintainEndTime
	}

	if !dara.IsNil(request.MaintainStartTime) {
		query["MaintainStartTime"] = request.MaintainStartTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceMaintainTime"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of an Tair (Redis OSS-compatible) instance. Alibaba Cloud maintains Tair (Redis OSS-compatible) instances during the specified maintenance window.
//
// Description:
//
// You can also modify the maintenance window of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Set a maintenance window](https://help.aliyun.com/document_detail/55252.html).
//
// @param request - ModifyInstanceMaintainTimeRequest
//
// @return ModifyInstanceMaintainTimeResponse
func (client *Client) ModifyInstanceMaintainTime(request *ModifyInstanceMaintainTimeRequest) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the major version of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// For more information about the precautions and impacts of the upgrade, see [Upgrade the major version](https://help.aliyun.com/document_detail/101764.html).
//
// @param request - ModifyInstanceMajorVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMajorVersionResponse
func (client *Client) ModifyInstanceMajorVersionWithOptions(request *ModifyInstanceMajorVersionRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMajorVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MajorVersion) {
		query["MajorVersion"] = request.MajorVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceMajorVersion"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceMajorVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the major version of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// For more information about the precautions and impacts of the upgrade, see [Upgrade the major version](https://help.aliyun.com/document_detail/101764.html).
//
// @param request - ModifyInstanceMajorVersionRequest
//
// @return ModifyInstanceMajorVersionResponse
func (client *Client) ModifyInstanceMajorVersion(request *ModifyInstanceMajorVersionRequest) (_result *ModifyInstanceMajorVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceMajorVersionResponse{}
	_body, _err := client.ModifyInstanceMajorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the minor version of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// The procedure to update the minor version of an instance varies based on types of Tair (Redis OSS-compatible) instances. For more information, see [Upgrade the minor version](https://help.aliyun.com/document_detail/56450.html).
//
// >
//
//   - Before you call this operation, you can call the [DescribeEngineVersion](https://help.aliyun.com/document_detail/473781.html) operation to query the minor version of the current instance.
//
//   - When you switch your workloads over from the original instance to a new instance or from the master node to the replica node in the original instance, you may experience disconnections that last a few seconds. The original instance stays in the read-only state within 60 seconds until all data is synchronized. We recommend that you upgrade the original instance during off-peak hours and make sure that your application is configured to automatically reconnect to the original instance.
//
// @param request - ModifyInstanceMinorVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMinorVersionResponse
func (client *Client) ModifyInstanceMinorVersionWithOptions(request *ModifyInstanceMinorVersionRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMinorVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Minorversion) {
		query["Minorversion"] = request.Minorversion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceMinorVersion"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceMinorVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the minor version of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// The procedure to update the minor version of an instance varies based on types of Tair (Redis OSS-compatible) instances. For more information, see [Upgrade the minor version](https://help.aliyun.com/document_detail/56450.html).
//
// >
//
//   - Before you call this operation, you can call the [DescribeEngineVersion](https://help.aliyun.com/document_detail/473781.html) operation to query the minor version of the current instance.
//
//   - When you switch your workloads over from the original instance to a new instance or from the master node to the replica node in the original instance, you may experience disconnections that last a few seconds. The original instance stays in the read-only state within 60 seconds until all data is synchronized. We recommend that you upgrade the original instance during off-peak hours and make sure that your application is configured to automatically reconnect to the original instance.
//
// @param request - ModifyInstanceMinorVersionRequest
//
// @return ModifyInstanceMinorVersionResponse
func (client *Client) ModifyInstanceMinorVersion(request *ModifyInstanceMinorVersionRequest) (_result *ModifyInstanceMinorVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceMinorVersionResponse{}
	_body, _err := client.ModifyInstanceMinorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Extends the retention period of the classic network endpoint of a Tair (Redis OSS-compatible) instance. You can call this operation after you change the network type of the Tair (Redis OSS-compatible) instance from classic network to Virtual Private Cloud (VPC) with the classic network endpoint retained.
//
// Description:
//
// You can also perform this operation in the Tair (Redis OSS-compatible) console. For more information, see [Change the expiration time for the endpoint of the classic network](https://help.aliyun.com/document_detail/60062.html).
//
// > For more information about how to switch the network type of a Tair (Redis OSS-compatible) instance from classic network to VPC, see [SwitchNetwork](https://help.aliyun.com/document_detail/473797.html).
//
// @param request - ModifyInstanceNetExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceNetExpireTimeResponse
func (client *Client) ModifyInstanceNetExpireTimeWithOptions(request *ModifyInstanceNetExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceNetExpireTimeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassicExpiredDays) {
		query["ClassicExpiredDays"] = request.ClassicExpiredDays
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceNetExpireTime"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceNetExpireTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extends the retention period of the classic network endpoint of a Tair (Redis OSS-compatible) instance. You can call this operation after you change the network type of the Tair (Redis OSS-compatible) instance from classic network to Virtual Private Cloud (VPC) with the classic network endpoint retained.
//
// Description:
//
// You can also perform this operation in the Tair (Redis OSS-compatible) console. For more information, see [Change the expiration time for the endpoint of the classic network](https://help.aliyun.com/document_detail/60062.html).
//
// > For more information about how to switch the network type of a Tair (Redis OSS-compatible) instance from classic network to VPC, see [SwitchNetwork](https://help.aliyun.com/document_detail/473797.html).
//
// @param request - ModifyInstanceNetExpireTimeRequest
//
// @return ModifyInstanceNetExpireTimeResponse
func (client *Client) ModifyInstanceNetExpireTime(request *ModifyInstanceNetExpireTimeRequest) (_result *ModifyInstanceNetExpireTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceNetExpireTimeResponse{}
	_body, _err := client.ModifyInstanceNetExpireTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies a parameter template to specific instances. This indicates that the parameter values in the template take effect on the instances. After you modify a parameter template, you must reapply it to specific instances for the new parameter values to take effect on the instances.
//
// @param request - ModifyInstanceParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceParameterResponse
func (client *Client) ModifyInstanceParameterWithOptions(request *ModifyInstanceParameterRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceParameterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceParameter"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies a parameter template to specific instances. This indicates that the parameter values in the template take effect on the instances. After you modify a parameter template, you must reapply it to specific instances for the new parameter values to take effect on the instances.
//
// @param request - ModifyInstanceParameterRequest
//
// @return ModifyInstanceParameterResponse
func (client *Client) ModifyInstanceParameter(request *ModifyInstanceParameterRequest) (_result *ModifyInstanceParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceParameterResponse{}
	_body, _err := client.ModifyInstanceParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables Transport Layer Security (TLS) for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also configure SSL encryption in the console. For more information, see [Configure SSL encryption](https://help.aliyun.com/document_detail/84898.html).
//
// >  To specify the earliest supported SSL version, you can call the [ModifyInstanceConfig](https://help.aliyun.com/document_detail/473844.html) operation to modify the required parameter.
//
// @param request - ModifyInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceSSLResponse
func (client *Client) ModifyInstanceSSLWithOptions(request *ModifyInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceSSLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SSLEnabled) {
		query["SSLEnabled"] = request.SSLEnabled
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceSSL"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceSSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables Transport Layer Security (TLS) for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also configure SSL encryption in the console. For more information, see [Configure SSL encryption](https://help.aliyun.com/document_detail/84898.html).
//
// >  To specify the earliest supported SSL version, you can call the [ModifyInstanceConfig](https://help.aliyun.com/document_detail/473844.html) operation to modify the required parameter.
//
// @param request - ModifyInstanceSSLRequest
//
// @return ModifyInstanceSSLResponse
func (client *Client) ModifyInstanceSSL(request *ModifyInstanceSSLRequest) (_result *ModifyInstanceSSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceSSLResponse{}
	_body, _err := client.ModifyInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the configurations of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >  For more information about the procedure, impacts, limits, and fees of this operation, see [Change the configurations of an instance](https://help.aliyun.com/document_detail/26353.html).
//
// @param request - ModifyInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceSpecResponse
func (client *Client) ModifyInstanceSpecWithOptions(request *ModifyInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceSpecResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.ForceTrans) {
		query["ForceTrans"] = request.ForceTrans
	}

	if !dara.IsNil(request.ForceUpgrade) {
		query["ForceUpgrade"] = request.ForceUpgrade
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MajorVersion) {
		query["MajorVersion"] = request.MajorVersion
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadOnlyCount) {
		query["ReadOnlyCount"] = request.ReadOnlyCount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaCount) {
		query["ReplicaCount"] = request.ReplicaCount
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ShardCount) {
		query["ShardCount"] = request.ShardCount
	}

	if !dara.IsNil(request.SlaveReadOnlyCount) {
		query["SlaveReadOnlyCount"] = request.SlaveReadOnlyCount
	}

	if !dara.IsNil(request.SlaveReplicaCount) {
		query["SlaveReplicaCount"] = request.SlaveReplicaCount
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceSpec"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the configurations of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >  For more information about the procedure, impacts, limits, and fees of this operation, see [Change the configurations of an instance](https://help.aliyun.com/document_detail/26353.html).
//
// @param request - ModifyInstanceSpecRequest
//
// @return ModifyInstanceSpecResponse
func (client *Client) ModifyInstanceSpec(request *ModifyInstanceSpecRequest) (_result *ModifyInstanceSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.ModifyInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables transparent data encryption (TDE) for a Tair (Redis OSS-compatible) instance. You can use existing custom keys.
//
// Description:
//
// > For more information about TDE and the impact of TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
// @param request - ModifyInstanceTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceTDEResponse
func (client *Client) ModifyInstanceTDEWithOptions(request *ModifyInstanceTDERequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceTDEResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.EncryptionName) {
		query["EncryptionName"] = request.EncryptionName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TDEStatus) {
		query["TDEStatus"] = request.TDEStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceTDE"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceTDEResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables transparent data encryption (TDE) for a Tair (Redis OSS-compatible) instance. You can use existing custom keys.
//
// Description:
//
// > For more information about TDE and the impact of TDE, see [Enable TDE](https://help.aliyun.com/document_detail/265913.html).
//
// @param request - ModifyInstanceTDERequest
//
// @return ModifyInstanceTDEResponse
func (client *Client) ModifyInstanceTDE(request *ModifyInstanceTDERequest) (_result *ModifyInstanceTDEResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceTDEResponse{}
	_body, _err := client.ModifyInstanceTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables password-free access for a Tair (Redis OSS-compatible) instance. This way, you can connect to a database in a convenient and secure manner.
//
// Description:
//
// When the password-free access feature is enabled, Elastic Compute Service (ECS) instances in the same virtual private cloud (VPC) can connect to the Tair instance without a password. You can also use the username and password to connect to the Tair instance.
//
// > The Tair instance is deployed in a VPC. For more information, see [Enable password-free access](https://help.aliyun.com/document_detail/85168.html).
//
// @param request - ModifyInstanceVpcAuthModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceVpcAuthModeResponse
func (client *Client) ModifyInstanceVpcAuthModeWithOptions(request *ModifyInstanceVpcAuthModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceVpcAuthModeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcAuthMode) {
		query["VpcAuthMode"] = request.VpcAuthMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceVpcAuthMode"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceVpcAuthModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables password-free access for a Tair (Redis OSS-compatible) instance. This way, you can connect to a database in a convenient and secure manner.
//
// Description:
//
// When the password-free access feature is enabled, Elastic Compute Service (ECS) instances in the same virtual private cloud (VPC) can connect to the Tair instance without a password. You can also use the username and password to connect to the Tair instance.
//
// > The Tair instance is deployed in a VPC. For more information, see [Enable password-free access](https://help.aliyun.com/document_detail/85168.html).
//
// @param request - ModifyInstanceVpcAuthModeRequest
//
// @return ModifyInstanceVpcAuthModeResponse
func (client *Client) ModifyInstanceVpcAuthMode(request *ModifyInstanceVpcAuthModeRequest) (_result *ModifyInstanceVpcAuthModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceVpcAuthModeResponse{}
	_body, _err := client.ModifyInstanceVpcAuthModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Temporarily adjusts the internal bandwidth of a Tair (Redis OSS-compatible) instance that is deployed in a dedicated cluster.
//
// Description:
//
// >
//
//   - This operation is applicable only to an instance that is deployed in a dedicated cluster. To adjust the bandwidth of a standard instance, call the [EnableAdditionalBandwidth](https://help.aliyun.com/document_detail/473771.html) operation.
//
// @param request - ModifyIntranetAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIntranetAttributeResponse
func (client *Client) ModifyIntranetAttributeWithOptions(request *ModifyIntranetAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyIntranetAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BandWidth) {
		query["BandWidth"] = request.BandWidth
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIntranetAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIntranetAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Temporarily adjusts the internal bandwidth of a Tair (Redis OSS-compatible) instance that is deployed in a dedicated cluster.
//
// Description:
//
// >
//
//   - This operation is applicable only to an instance that is deployed in a dedicated cluster. To adjust the bandwidth of a standard instance, call the [EnableAdditionalBandwidth](https://help.aliyun.com/document_detail/473771.html) operation.
//
// @param request - ModifyIntranetAttributeRequest
//
// @return ModifyIntranetAttributeResponse
func (client *Client) ModifyIntranetAttribute(request *ModifyIntranetAttributeRequest) (_result *ModifyIntranetAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyIntranetAttributeResponse{}
	_body, _err := client.ModifyIntranetAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the settings of a parameter template.
//
// @param request - ModifyParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParameterGroupResponse
func (client *Client) ModifyParameterGroupWithOptions(request *ModifyParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyParameterGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupDesc) {
		query["ParameterGroupDesc"] = request.ParameterGroupDesc
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.ParameterGroupName) {
		query["ParameterGroupName"] = request.ParameterGroupName
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyParameterGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyParameterGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of a parameter template.
//
// @param request - ModifyParameterGroupRequest
//
// @return ModifyParameterGroupResponse
func (client *Client) ModifyParameterGroup(request *ModifyParameterGroupRequest) (_result *ModifyParameterGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyParameterGroupResponse{}
	_body, _err := client.ModifyParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a Tair (Redis OSS-compatible) instance belongs.
//
// Description:
//
// You can also perform this operation in the [Resource Management](https://resourcemanager.console.aliyun.com/resource-center) console. For more information, see [Transfer resources across resource groups](https://help.aliyun.com/document_detail/94487.html).
//
// >  Resource Group allows you to sort resources owned by your Alibaba Cloud account into groups. This simplifies the resource and permission management within your Alibaba Cloud account. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html)
//
// @param request - ModifyResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyResourceGroupResponse
func (client *Client) ModifyResourceGroupWithOptions(request *ModifyResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("ModifyResourceGroup"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a Tair (Redis OSS-compatible) instance belongs.
//
// Description:
//
// You can also perform this operation in the [Resource Management](https://resourcemanager.console.aliyun.com/resource-center) console. For more information, see [Transfer resources across resource groups](https://help.aliyun.com/document_detail/94487.html).
//
// >  Resource Group allows you to sort resources owned by your Alibaba Cloud account into groups. This simplifies the resource and permission management within your Alibaba Cloud account. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html)
//
// @param request - ModifyResourceGroupRequest
//
// @return ModifyResourceGroupResponse
func (client *Client) ModifyResourceGroup(request *ModifyResourceGroupRequest) (_result *ModifyResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyResourceGroupResponse{}
	_body, _err := client.ModifyResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the security groups that are added to the whitelists of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// > After you call this operation, the security groups that are added to the whitelists of the Tair instance are deleted, and the security group specified by the **SecurityGroupId*	- parameter is added to the whitelists. For more information about how to reset security groups in the Tair console, see [Add security groups](https://help.aliyun.com/document_detail/148267.html).
//
// @param request - ModifySecurityGroupConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityGroupConfigurationResponse
func (client *Client) ModifySecurityGroupConfigurationWithOptions(request *ModifySecurityGroupConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityGroupConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityGroupConfiguration"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityGroupConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the security groups that are added to the whitelists of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// > After you call this operation, the security groups that are added to the whitelists of the Tair instance are deleted, and the security group specified by the **SecurityGroupId*	- parameter is added to the whitelists. For more information about how to reset security groups in the Tair console, see [Add security groups](https://help.aliyun.com/document_detail/148267.html).
//
// @param request - ModifySecurityGroupConfigurationRequest
//
// @return ModifySecurityGroupConfigurationResponse
func (client *Client) ModifySecurityGroupConfiguration(request *ModifySecurityGroupConfigurationRequest) (_result *ModifySecurityGroupConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySecurityGroupConfigurationResponse{}
	_body, _err := client.ModifySecurityGroupConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelists of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also modify the whitelists of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Configure a whitelist for an instance](https://help.aliyun.com/document_detail/56464.html).
//
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityIpGroupAttribute) {
		query["SecurityIpGroupAttribute"] = request.SecurityIpGroupAttribute
	}

	if !dara.IsNil(request.SecurityIpGroupName) {
		query["SecurityIpGroupName"] = request.SecurityIpGroupName
	}

	if !dara.IsNil(request.SecurityIps) {
		query["SecurityIps"] = request.SecurityIps
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityIps"),
		Version:     dara.String("2015-01-01"),
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

// Summary:
//
// Modifies the IP address whitelists of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also modify the whitelists of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Configure a whitelist for an instance](https://help.aliyun.com/document_detail/56464.html).
//
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
// 修改TairCustom实例基本参数
//
// @param request - ModifyTairKVCacheCustomInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTairKVCacheCustomInstanceAttributeResponse
func (client *Client) ModifyTairKVCacheCustomInstanceAttributeWithOptions(request *ModifyTairKVCacheCustomInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyTairKVCacheCustomInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTairKVCacheCustomInstanceAttribute"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTairKVCacheCustomInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改TairCustom实例基本参数
//
// @param request - ModifyTairKVCacheCustomInstanceAttributeRequest
//
// @return ModifyTairKVCacheCustomInstanceAttributeResponse
func (client *Client) ModifyTairKVCacheCustomInstanceAttribute(request *ModifyTairKVCacheCustomInstanceAttributeRequest) (_result *ModifyTairKVCacheCustomInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTairKVCacheCustomInstanceAttributeResponse{}
	_body, _err := client.ModifyTairKVCacheCustomInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the task information, such as the task execution time.
//
// @param request - ModifyTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTaskInfoResponse
func (client *Client) ModifyTaskInfoWithOptions(request *ModifyTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyTaskInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionParams) {
		query["ActionParams"] = request.ActionParams
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StepName) {
		query["StepName"] = request.StepName
	}

	if !dara.IsNil(request.TaskAction) {
		query["TaskAction"] = request.TaskAction
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTaskInfo"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the task information, such as the task execution time.
//
// @param request - ModifyTaskInfoRequest
//
// @return ModifyTaskInfoResponse
func (client *Client) ModifyTaskInfo(request *ModifyTaskInfoRequest) (_result *ModifyTaskInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTaskInfoResponse{}
	_body, _err := client.ModifyTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the private endpoint of an ApsaraDB for Redis cluster instance.
//
// Description:
//
// In direct connection mode, clients can bypass proxy nodes and use private endpoints to connect to ApsaraDB for Redis instances. This is similar to the connection to a native Redis cluster. The direct connection mode can reduce communication overheads and the response time of ApsaraDB for Redis. For more information, see [Enable the direct connection mode](https://help.aliyun.com/document_detail/146901.html).
//
// @param request - ReleaseDirectConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseDirectConnectionResponse
func (client *Client) ReleaseDirectConnectionWithOptions(request *ReleaseDirectConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseDirectConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseDirectConnection"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseDirectConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the private endpoint of an ApsaraDB for Redis cluster instance.
//
// Description:
//
// In direct connection mode, clients can bypass proxy nodes and use private endpoints to connect to ApsaraDB for Redis instances. This is similar to the connection to a native Redis cluster. The direct connection mode can reduce communication overheads and the response time of ApsaraDB for Redis. For more information, see [Enable the direct connection mode](https://help.aliyun.com/document_detail/146901.html).
//
// @param request - ReleaseDirectConnectionRequest
//
// @return ReleaseDirectConnectionResponse
func (client *Client) ReleaseDirectConnection(request *ReleaseDirectConnectionRequest) (_result *ReleaseDirectConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseDirectConnectionResponse{}
	_body, _err := client.ReleaseDirectConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also release the public endpoint for an instance in the Tair (Redis OSS-compatible) console. For more information, see [Release public endpoints](https://help.aliyun.com/document_detail/125424.html).
//
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstancePublicConnection"),
		Version:     dara.String("2015-01-01"),
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

// Summary:
//
// Releases the public endpoint of a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// You can also release the public endpoint for an instance in the Tair (Redis OSS-compatible) console. For more information, see [Release public endpoints](https://help.aliyun.com/document_detail/125424.html).
//
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

// Summary:
//
// Removes a child instance from a distributed instance.
//
// Description:
//
// The operation that you want to perform. Set the value to **RemoveSubInstance**.
//
// @param request - RemoveSubInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSubInstanceResponse
func (client *Client) RemoveSubInstanceWithOptions(request *RemoveSubInstanceRequest, runtime *dara.RuntimeOptions) (_result *RemoveSubInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSubInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSubInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a child instance from a distributed instance.
//
// Description:
//
// The operation that you want to perform. Set the value to **RemoveSubInstance**.
//
// @param request - RemoveSubInstanceRequest
//
// @return RemoveSubInstanceResponse
func (client *Client) RemoveSubInstance(request *RemoveSubInstanceRequest) (_result *RemoveSubInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveSubInstanceResponse{}
	_body, _err := client.RemoveSubInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation is not recommended now. The billing method for bandwidth of a Tair (Redis OSS-compatible) instance is changed to pay-as-you-go.
//
// Description:
//
// You can adjust the bandwidth of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Adjust the bandwidth of an instance](https://help.aliyun.com/document_detail/102588.html). You can also call the [EnableAdditionalBandwidth](https://help.aliyun.com/document_detail/473771.html) operation to purchase bandwidth for an instance.
//
// @param request - RenewAdditionalBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAdditionalBandwidthResponse
func (client *Client) RenewAdditionalBandwidthWithOptions(request *RenewAdditionalBandwidthRequest, runtime *dara.RuntimeOptions) (_result *RenewAdditionalBandwidthResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderTimeLength) {
		query["OrderTimeLength"] = request.OrderTimeLength
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewAdditionalBandwidth"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewAdditionalBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation is not recommended now. The billing method for bandwidth of a Tair (Redis OSS-compatible) instance is changed to pay-as-you-go.
//
// Description:
//
// You can adjust the bandwidth of an instance in the Tair (Redis OSS-compatible) console. For more information, see [Adjust the bandwidth of an instance](https://help.aliyun.com/document_detail/102588.html). You can also call the [EnableAdditionalBandwidth](https://help.aliyun.com/document_detail/473771.html) operation to purchase bandwidth for an instance.
//
// @param request - RenewAdditionalBandwidthRequest
//
// @return RenewAdditionalBandwidthResponse
func (client *Client) RenewAdditionalBandwidth(request *RenewAdditionalBandwidthRequest) (_result *RenewAdditionalBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewAdditionalBandwidthResponse{}
	_body, _err := client.RenewAdditionalBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews an ApsaraDB for Redis instance.
//
// Description:
//
// This operation is applicable only to subscription instances.
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.FromApp) {
		query["FromApp"] = request.FromApp
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Renews an ApsaraDB for Redis instance.
//
// Description:
//
// This operation is applicable only to subscription instances.
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of an account for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >  Only Tair (Redis OSS-compatible) instances of Redis 4.0 or later are supported.
//
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2015-01-01"),
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
// Resets the password of an account for a Tair (Redis OSS-compatible) instance.
//
// Description:
//
// >  Only Tair (Redis OSS-compatible) instances of Redis 4.0 or later are supported.
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
// 重置TairCustom上主机密码
//
// @param request - ResetTairKVCacheCustomInstancePasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetTairKVCacheCustomInstancePasswordResponse
func (client *Client) ResetTairKVCacheCustomInstancePasswordWithOptions(request *ResetTairKVCacheCustomInstancePasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetTairKVCacheCustomInstancePasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetTairKVCacheCustomInstancePassword"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetTairKVCacheCustomInstancePasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置TairCustom上主机密码
//
// @param request - ResetTairKVCacheCustomInstancePasswordRequest
//
// @return ResetTairKVCacheCustomInstancePasswordResponse
func (client *Client) ResetTairKVCacheCustomInstancePassword(request *ResetTairKVCacheCustomInstancePasswordRequest) (_result *ResetTairKVCacheCustomInstancePasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetTairKVCacheCustomInstancePasswordResponse{}
	_body, _err := client.ResetTairKVCacheCustomInstancePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变配TairCustom的主机的磁盘
//
// @param request - ResizeTairKVCacheCustomInstanceDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeTairKVCacheCustomInstanceDiskResponse
func (client *Client) ResizeTairKVCacheCustomInstanceDiskWithOptions(request *ResizeTairKVCacheCustomInstanceDiskRequest, runtime *dara.RuntimeOptions) (_result *ResizeTairKVCacheCustomInstanceDiskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResizeTairKVCacheCustomInstanceDisk"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeTairKVCacheCustomInstanceDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变配TairCustom的主机的磁盘
//
// @param request - ResizeTairKVCacheCustomInstanceDiskRequest
//
// @return ResizeTairKVCacheCustomInstanceDiskResponse
func (client *Client) ResizeTairKVCacheCustomInstanceDisk(request *ResizeTairKVCacheCustomInstanceDiskRequest) (_result *ResizeTairKVCacheCustomInstanceDiskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResizeTairKVCacheCustomInstanceDiskResponse{}
	_body, _err := client.ResizeTairKVCacheCustomInstanceDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a running ApsaraDB for Redis instance.
//
// @param request - RestartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.UpgradeMinorVersion) {
		query["UpgradeMinorVersion"] = request.UpgradeMinorVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Restarts a running ApsaraDB for Redis instance.
//
// @param request - RestartInstanceRequest
//
// @return RestartInstanceResponse
func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启TairCustom的主机
//
// @param request - RestartTairKVCacheCustomInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartTairKVCacheCustomInstanceResponse
func (client *Client) RestartTairKVCacheCustomInstanceWithOptions(request *RestartTairKVCacheCustomInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartTairKVCacheCustomInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartTairKVCacheCustomInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartTairKVCacheCustomInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启TairCustom的主机
//
// @param request - RestartTairKVCacheCustomInstanceRequest
//
// @return RestartTairKVCacheCustomInstanceResponse
func (client *Client) RestartTairKVCacheCustomInstance(request *RestartTairKVCacheCustomInstanceRequest) (_result *RestartTairKVCacheCustomInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartTairKVCacheCustomInstanceResponse{}
	_body, _err := client.RestartTairKVCacheCustomInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores the data in a backup file to a specified Tair (Redis OSS-compatible) instance.
//
// Description:
//
//	  If your instance is a [DRAM-based instance](https://help.aliyun.com/document_detail/126164.html) or a [persistent memory-optimized instance](https://help.aliyun.com/document_detail/183956.html) and has the [data flashback](https://help.aliyun.com/document_detail/148479.html) feature enabled, you can call this operation to restore the entire instance or specific keys to a specific point in time accurate to the second. This way, you can achieve more fine-grained data restoration.
//
//		- For other types of instances, we recommend that you call the [CreateInstance](https://help.aliyun.com/document_detail/473757.html) or [CreateTairInstance](https://help.aliyun.com/document_detail/473770.html) operation to restore the backup data to a new instance.
//
// @param request - RestoreInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreInstanceResponse
func (client *Client) RestoreInstanceWithOptions(request *RestoreInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestoreInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.FilterKey) {
		query["FilterKey"] = request.FilterKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TimeShift) {
		query["TimeShift"] = request.TimeShift
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores the data in a backup file to a specified Tair (Redis OSS-compatible) instance.
//
// Description:
//
//	  If your instance is a [DRAM-based instance](https://help.aliyun.com/document_detail/126164.html) or a [persistent memory-optimized instance](https://help.aliyun.com/document_detail/183956.html) and has the [data flashback](https://help.aliyun.com/document_detail/148479.html) feature enabled, you can call this operation to restore the entire instance or specific keys to a specific point in time accurate to the second. This way, you can achieve more fine-grained data restoration.
//
//		- For other types of instances, we recommend that you call the [CreateInstance](https://help.aliyun.com/document_detail/473757.html) or [CreateTairInstance](https://help.aliyun.com/document_detail/473770.html) operation to restore the backup data to a new instance.
//
// @param request - RestoreInstanceRequest
//
// @return RestoreInstanceResponse
func (client *Client) RestoreInstance(request *RestoreInstanceRequest) (_result *RestoreInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.RestoreInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动TairCustom的主机
//
// @param request - StartTairKVCacheCustomInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTairKVCacheCustomInstanceResponse
func (client *Client) StartTairKVCacheCustomInstanceWithOptions(request *StartTairKVCacheCustomInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartTairKVCacheCustomInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTairKVCacheCustomInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTairKVCacheCustomInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动TairCustom的主机
//
// @param request - StartTairKVCacheCustomInstanceRequest
//
// @return StartTairKVCacheCustomInstanceResponse
func (client *Client) StartTairKVCacheCustomInstance(request *StartTairKVCacheCustomInstanceRequest) (_result *StartTairKVCacheCustomInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartTairKVCacheCustomInstanceResponse{}
	_body, _err := client.StartTairKVCacheCustomInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止TairCustom的主机
//
// @param request - StopTairKVCacheCustomInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTairKVCacheCustomInstanceResponse
func (client *Client) StopTairKVCacheCustomInstanceWithOptions(request *StopTairKVCacheCustomInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopTairKVCacheCustomInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTairKVCacheCustomInstance"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTairKVCacheCustomInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止TairCustom的主机
//
// @param request - StopTairKVCacheCustomInstanceRequest
//
// @return StopTairKVCacheCustomInstanceResponse
func (client *Client) StopTairKVCacheCustomInstance(request *StopTairKVCacheCustomInstanceRequest) (_result *StopTairKVCacheCustomInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopTairKVCacheCustomInstanceResponse{}
	_body, _err := client.StopTairKVCacheCustomInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a master-replica switchover to switch node roles. This operation is applicable to disaster recovery drills and nearby access to applications that are deployed across zones.
//
// Description:
//
// > For more information about nearby access to applications that are deployed across zones, see [Switch node roles](https://help.aliyun.com/document_detail/164222.html).
//
// The instance must be a Redis Open-Source Edition instance or Tair (Enterprise Edition) [DRAM-based](https://help.aliyun.com/document_detail/126164.html) instance that uses local disks.
//
// A call to this operation has the following impacts on your instance:
//
//   - The data shards in the instance may change to the read-only state and experience transient connections within seconds. Make sure that your application is configured to automatically reconnect to the instance.
//
//   - If the instance enters the switching state, you cannot manage this instance. For example, you cannot modify the instance configurations or migrate the instance to another zone.
//
// @param request - SwitchInstanceHARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchInstanceHAResponse
func (client *Client) SwitchInstanceHAWithOptions(request *SwitchInstanceHARequest, runtime *dara.RuntimeOptions) (_result *SwitchInstanceHAResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	if !dara.IsNil(request.SwitchType) {
		query["SwitchType"] = request.SwitchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchInstanceHA"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchInstanceHAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a master-replica switchover to switch node roles. This operation is applicable to disaster recovery drills and nearby access to applications that are deployed across zones.
//
// Description:
//
// > For more information about nearby access to applications that are deployed across zones, see [Switch node roles](https://help.aliyun.com/document_detail/164222.html).
//
// The instance must be a Redis Open-Source Edition instance or Tair (Enterprise Edition) [DRAM-based](https://help.aliyun.com/document_detail/126164.html) instance that uses local disks.
//
// A call to this operation has the following impacts on your instance:
//
//   - The data shards in the instance may change to the read-only state and experience transient connections within seconds. Make sure that your application is configured to automatically reconnect to the instance.
//
//   - If the instance enters the switching state, you cannot manage this instance. For example, you cannot modify the instance configurations or migrate the instance to another zone.
//
// @param request - SwitchInstanceHARequest
//
// @return SwitchInstanceHAResponse
func (client *Client) SwitchInstanceHA(request *SwitchInstanceHARequest) (_result *SwitchInstanceHAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchInstanceHAResponse{}
	_body, _err := client.SwitchInstanceHAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the proxy mode for a Tair (Redis OSS-compatible) cluster instance in a dedicated cluster.
//
// Description:
//
// For more information about the proxy mode, see [Features of proxy servers](https://help.aliyun.com/document_detail/142959.html). Before you call this operation, make sure that the following requirements are met:
//
//   - The instance is created by using a dedicated cluster. For more information, see [What is ApsaraDB for MyBase?](https://help.aliyun.com/document_detail/141455.html)
//
//   - The instance uses the [cluster architecture](https://help.aliyun.com/document_detail/52228.html).
//
// >  Before you call the SwitchInstanceProxy operation, you must call the [DescribeDedicatedClusterInstanceList](https://help.aliyun.com/document_detail/473867.html) operation and view the value of the **ProxyCount*	- response parameter to check whether the proxy mode is enabled. A value of 0 indicates that the proxy mode is disabled. A value that is greater than 0 indicates that the proxy mode is enabled.
//
// @param request - SwitchInstanceProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchInstanceProxyResponse
func (client *Client) SwitchInstanceProxyWithOptions(request *SwitchInstanceProxyRequest, runtime *dara.RuntimeOptions) (_result *SwitchInstanceProxyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchInstanceProxy"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchInstanceProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the proxy mode for a Tair (Redis OSS-compatible) cluster instance in a dedicated cluster.
//
// Description:
//
// For more information about the proxy mode, see [Features of proxy servers](https://help.aliyun.com/document_detail/142959.html). Before you call this operation, make sure that the following requirements are met:
//
//   - The instance is created by using a dedicated cluster. For more information, see [What is ApsaraDB for MyBase?](https://help.aliyun.com/document_detail/141455.html)
//
//   - The instance uses the [cluster architecture](https://help.aliyun.com/document_detail/52228.html).
//
// >  Before you call the SwitchInstanceProxy operation, you must call the [DescribeDedicatedClusterInstanceList](https://help.aliyun.com/document_detail/473867.html) operation and view the value of the **ProxyCount*	- response parameter to check whether the proxy mode is enabled. A value of 0 indicates that the proxy mode is disabled. A value that is greater than 0 indicates that the proxy mode is enabled.
//
// @param request - SwitchInstanceProxyRequest
//
// @return SwitchInstanceProxyResponse
func (client *Client) SwitchInstanceProxy(request *SwitchInstanceProxyRequest) (_result *SwitchInstanceProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchInstanceProxyResponse{}
	_body, _err := client.SwitchInstanceProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches an instance from the current zone to the specified zone in the event of a fault.
//
// @param request - SwitchInstanceZoneFailOverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchInstanceZoneFailOverResponse
func (client *Client) SwitchInstanceZoneFailOverWithOptions(request *SwitchInstanceZoneFailOverRequest, runtime *dara.RuntimeOptions) (_result *SwitchInstanceZoneFailOverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SiteFaultTime) {
		query["SiteFaultTime"] = request.SiteFaultTime
	}

	if !dara.IsNil(request.TargetZoneId) {
		query["TargetZoneId"] = request.TargetZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchInstanceZoneFailOver"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchInstanceZoneFailOverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches an instance from the current zone to the specified zone in the event of a fault.
//
// @param request - SwitchInstanceZoneFailOverRequest
//
// @return SwitchInstanceZoneFailOverResponse
func (client *Client) SwitchInstanceZoneFailOver(request *SwitchInstanceZoneFailOverRequest) (_result *SwitchInstanceZoneFailOverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchInstanceZoneFailOverResponse{}
	_body, _err := client.SwitchInstanceZoneFailOverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the VPC or vSwitch of a Tair (Redis OSS-compatible) instance. If the instance is deployed in the classic network, the network type of the instance is changed from the classic network to VPC.
//
// @param request - SwitchNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchNetworkResponse
func (client *Client) SwitchNetworkWithOptions(request *SwitchNetworkRequest, runtime *dara.RuntimeOptions) (_result *SwitchNetworkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassicExpiredDays) {
		query["ClassicExpiredDays"] = request.ClassicExpiredDays
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RetainClassic) {
		query["RetainClassic"] = request.RetainClassic
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TargetNetworkType) {
		query["TargetNetworkType"] = request.TargetNetworkType
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchNetwork"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the VPC or vSwitch of a Tair (Redis OSS-compatible) instance. If the instance is deployed in the classic network, the network type of the instance is changed from the classic network to VPC.
//
// @param request - SwitchNetworkRequest
//
// @return SwitchNetworkResponse
func (client *Client) SwitchNetwork(request *SwitchNetworkRequest) (_result *SwitchNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchNetworkResponse{}
	_body, _err := client.SwitchNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables configuration changes for a Tair (Redis OSS-compatible) instance before you use Data Transmission Service (DTS) to migrate or synchronize data of the instance. This prevents migration and synchronization task failures due to configuration changes.
//
// @param request - SyncDtsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDtsStatusResponse
func (client *Client) SyncDtsStatusWithOptions(request *SyncDtsStatusRequest, runtime *dara.RuntimeOptions) (_result *SyncDtsStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDtsStatus"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDtsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables configuration changes for a Tair (Redis OSS-compatible) instance before you use Data Transmission Service (DTS) to migrate or synchronize data of the instance. This prevents migration and synchronization task failures due to configuration changes.
//
// @param request - SyncDtsStatusRequest
//
// @return SyncDtsStatusResponse
func (client *Client) SyncDtsStatus(request *SyncDtsStatusRequest) (_result *SyncDtsStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncDtsStatusResponse{}
	_body, _err := client.SyncDtsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to Tair (Redis OSS-compatible) instances.
//
// Description:
//
// If you have a large number of instances, you can create multiple tags and add these tags to the instances. Then, you can filter instances by tag.
//
//   - A tag consists of a key and a value. Each key must be unique in a region for an Alibaba Cloud account. Different keys can be mapped to the same value.
//
//   - If the tag that you specify does not exist, this tag is automatically created and added to the specified instance.
//
//   - If the key of the specified tag is the same as that of an existing tag, the specified tag overwrites the existing tag.
//
//   - You can add up to 20 tags to each instance.
//
//   - You can add tags to up to 50 instances in each request.
//
// You can also add tags to instances in the Tair (Redis OSS-compatible) console. For more information, see [Create a tag](https://help.aliyun.com/document_detail/118779.html).
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
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2015-01-01"),
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
// Adds tags to Tair (Redis OSS-compatible) instances.
//
// Description:
//
// If you have a large number of instances, you can create multiple tags and add these tags to the instances. Then, you can filter instances by tag.
//
//   - A tag consists of a key and a value. Each key must be unique in a region for an Alibaba Cloud account. Different keys can be mapped to the same value.
//
//   - If the tag that you specify does not exist, this tag is automatically created and added to the specified instance.
//
//   - If the key of the specified tag is the same as that of an existing tag, the specified tag overwrites the existing tag.
//
//   - You can add up to 20 tags to each instance.
//
//   - You can add tags to up to 50 instances in each request.
//
// You can also add tags to instances in the Tair (Redis OSS-compatible) console. For more information, see [Create a tag](https://help.aliyun.com/document_detail/118779.html).
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
// Changes the billing method of a Tair (Redis OSS-compatible) instance from subscription to pay-as-you-go or from pay-as-you-go to subscription.
//
// Description:
//
// Before you call this operation, make sure that you understand relevant precautions and billing rules. For more information, see the following topics:
//
//   - [Change the billing method to subscription](https://help.aliyun.com/document_detail/54542.html).
//
//   - [Change the billing method to pay-as-you-go](https://help.aliyun.com/document_detail/211549.html).
//
// @param request - TransformInstanceChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformInstanceChargeTypeResponse
func (client *Client) TransformInstanceChargeTypeWithOptions(request *TransformInstanceChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *TransformInstanceChargeTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransformInstanceChargeType"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransformInstanceChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of a Tair (Redis OSS-compatible) instance from subscription to pay-as-you-go or from pay-as-you-go to subscription.
//
// Description:
//
// Before you call this operation, make sure that you understand relevant precautions and billing rules. For more information, see the following topics:
//
//   - [Change the billing method to subscription](https://help.aliyun.com/document_detail/54542.html).
//
//   - [Change the billing method to pay-as-you-go](https://help.aliyun.com/document_detail/211549.html).
//
// @param request - TransformInstanceChargeTypeRequest
//
// @return TransformInstanceChargeTypeResponse
func (client *Client) TransformInstanceChargeType(request *TransformInstanceChargeTypeRequest) (_result *TransformInstanceChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransformInstanceChargeTypeResponse{}
	_body, _err := client.TransformInstanceChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 转换本地盘到云原生
//
// @param request - TransformToEcsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformToEcsResponse
func (client *Client) TransformToEcsWithOptions(request *TransformToEcsRequest, runtime *dara.RuntimeOptions) (_result *TransformToEcsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ShardCount) {
		query["ShardCount"] = request.ShardCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransformToEcs"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransformToEcsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转换本地盘到云原生
//
// @param request - TransformToEcsRequest
//
// @return TransformToEcsResponse
func (client *Client) TransformToEcs(request *TransformToEcsRequest) (_result *TransformToEcsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransformToEcsResponse{}
	_body, _err := client.TransformToEcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes a pay-as-you-go Tair (Redis OSS-compatible) instance to a subscription instance.
//
// Description:
//
// For more information about how to change the billing method in the Tair (Redis OSS-compatible) console, see [Switch to subscription](https://help.aliyun.com/document_detail/54542.html).
//
// >  You cannot change the billing method of a Tair (Redis OSS-compatible) instance from subscription to pay-as-you-go.
//
// @param request - TransformToPrePaidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformToPrePaidResponse
func (client *Client) TransformToPrePaidWithOptions(request *TransformToPrePaidRequest, runtime *dara.RuntimeOptions) (_result *TransformToPrePaidResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransformToPrePaid"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransformToPrePaidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes a pay-as-you-go Tair (Redis OSS-compatible) instance to a subscription instance.
//
// Description:
//
// For more information about how to change the billing method in the Tair (Redis OSS-compatible) console, see [Switch to subscription](https://help.aliyun.com/document_detail/54542.html).
//
// >  You cannot change the billing method of a Tair (Redis OSS-compatible) instance from subscription to pay-as-you-go.
//
// @param request - TransformToPrePaidRequest
//
// @return TransformToPrePaidResponse
func (client *Client) TransformToPrePaid(request *TransformToPrePaidRequest) (_result *TransformToPrePaidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransformToPrePaidResponse{}
	_body, _err := client.TransformToPrePaidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the write lock from an instance. After the instance is unlocked, it supports both read and write operations.
//
// @param request - UnlockDBInstanceWriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockDBInstanceWriteResponse
func (client *Client) UnlockDBInstanceWriteWithOptions(request *UnlockDBInstanceWriteRequest, runtime *dara.RuntimeOptions) (_result *UnlockDBInstanceWriteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockDBInstanceWrite"),
		Version:     dara.String("2015-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockDBInstanceWriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the write lock from an instance. After the instance is unlocked, it supports both read and write operations.
//
// @param request - UnlockDBInstanceWriteRequest
//
// @return UnlockDBInstanceWriteResponse
func (client *Client) UnlockDBInstanceWrite(request *UnlockDBInstanceWriteRequest) (_result *UnlockDBInstanceWriteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockDBInstanceWriteResponse{}
	_body, _err := client.UnlockDBInstanceWriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from Tair (Redis OSS-compatible) instances.
//
// Description:
//
//	  You can remove up to 20 tags at a time.
//
//		- If a tag is removed from an instance and is not added to other instances, the tag is deleted.
//
// You can also remove tags from instances in the Tair (Redis OSS-compatible) console. For more information, see [Remove a tag](https://help.aliyun.com/document_detail/119157.html).
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2015-01-01"),
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
// Removes tags from Tair (Redis OSS-compatible) instances.
//
// Description:
//
//	  You can remove up to 20 tags at a time.
//
//		- If a tag is removed from an instance and is not added to other instances, the tag is deleted.
//
// You can also remove tags from instances in the Tair (Redis OSS-compatible) console. For more information, see [Remove a tag](https://help.aliyun.com/document_detail/119157.html).
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
