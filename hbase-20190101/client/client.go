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
		"ap-northeast-2-pop":          dara.String("hbase.aliyuncs.com"),
		"ap-south-1":                  dara.String("hbase.aliyuncs.com"),
		"ap-southeast-2":              dara.String("hbase.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("hbase.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("hbase.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("hbase.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("hbase.aliyuncs.com"),
		"cn-edge-1":                   dara.String("hbase.aliyuncs.com"),
		"cn-fujian":                   dara.String("hbase.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("hbase.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("hbase.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("hbase.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("hbase.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("hbase.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("hbase.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("hbase.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("hbase.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("hbase.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("hbase.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("hbase.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("hbase.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("hbase.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("hbase.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("hbase.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("hbase.aliyuncs.com"),
		"cn-wuhan":                    dara.String("hbase.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("hbase.aliyuncs.com"),
		"cn-yushanfang":               dara.String("hbase.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("hbase.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("hbase.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("hbase.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("hbase.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("hbase.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("hbase.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("hbase"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a self-managed HDFS address to Xpack.
//
// @param request - AddUserHdfsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserHdfsInfoResponse
func (client *Client) AddUserHdfsInfoWithOptions(request *AddUserHdfsInfoRequest, runtime *dara.RuntimeOptions) (_result *AddUserHdfsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserHdfsInfo"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserHdfsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a self-managed HDFS address to Xpack.
//
// @param request - AddUserHdfsInfoRequest
//
// @return AddUserHdfsInfoResponse
func (client *Client) AddUserHdfsInfo(request *AddUserHdfsInfoRequest) (_result *AddUserHdfsInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserHdfsInfoResponse{}
	_body, _err := client.AddUserHdfsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a public endpoint.
//
// @param request - AllocatePublicNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocatePublicNetworkAddressResponse
func (client *Client) AllocatePublicNetworkAddressWithOptions(request *AllocatePublicNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *AllocatePublicNetworkAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocatePublicNetworkAddress"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a public endpoint.
//
// @param request - AllocatePublicNetworkAddressRequest
//
// @return AllocatePublicNetworkAddressResponse
func (client *Client) AllocatePublicNetworkAddress(request *AllocatePublicNetworkAddressRequest) (_result *AllocatePublicNetworkAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.AllocatePublicNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels O&M event tasks by calling the CancelActiveOperationTasks operation.
//
// @param request - CancelActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasksWithOptions(request *CancelActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *CancelActiveOperationTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2019-01-01"),
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
// Cancels O&M event tasks by calling the CancelActiveOperationTasks operation.
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
// Checks whether the versions of components are the latest versions by calling CheckComponentsVersion.
//
// @param request - CheckComponentsVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComponentsVersionResponse
func (client *Client) CheckComponentsVersionWithOptions(request *CheckComponentsVersionRequest, runtime *dara.RuntimeOptions) (_result *CheckComponentsVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Components) {
		query["Components"] = request.Components
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckComponentsVersion"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckComponentsVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the versions of components are the latest versions by calling CheckComponentsVersion.
//
// @param request - CheckComponentsVersionRequest
//
// @return CheckComponentsVersionResponse
func (client *Client) CheckComponentsVersion(request *CheckComponentsVersionRequest) (_result *CheckComponentsVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckComponentsVersionResponse{}
	_body, _err := client.CheckComponentsVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the backup and recovery feature for an HBase cluster by calling CloseBackup.
//
// @param request - CloseBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseBackupResponse
func (client *Client) CloseBackupWithOptions(request *CloseBackupRequest, runtime *dara.RuntimeOptions) (_result *CloseBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseBackup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the backup and recovery feature for an HBase cluster by calling CloseBackup.
//
// @param request - CloseBackupRequest
//
// @return CloseBackupResponse
func (client *Client) CloseBackup(request *CloseBackupRequest) (_result *CloseBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseBackupResponse{}
	_body, _err := client.CloseBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Converts an instance from pay-as-you-go to subscription billing.
//
// Description:
//
// Note: This operation can be called only when payType is set to Postpaid.
//
// @param request - ConvertInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertInstanceResponse
func (client *Client) ConvertInstanceWithOptions(request *ConvertInstanceRequest, runtime *dara.RuntimeOptions) (_result *ConvertInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertInstance"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts an instance from pay-as-you-go to subscription billing.
//
// Description:
//
// Note: This operation can be called only when payType is set to Postpaid.
//
// @param request - ConvertInstanceRequest
//
// @return ConvertInstanceResponse
func (client *Client) ConvertInstance(request *ConvertInstanceRequest) (_result *ConvertInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConvertInstanceResponse{}
	_body, _err := client.ConvertInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a wide table account.
//
// Description:
//
// Only ApsaraDB for HBase Performance-enhanced Edition is supported.
//
// @param request - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2019-01-01"),
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
// Creates a wide table account.
//
// Description:
//
// Only ApsaraDB for HBase Performance-enhanced Edition is supported.
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
// Creates a backup plan. Currently, only HBaseue is supported.
//
// Description:
//
// Currently, this operation supports only HBaseue clusters. The EnableHBaseueBackup operation is compatible with this operation and performs automatic creation of a backup plan after a BDS cluster is created.
//
// @param request - CreateBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupPlanResponse
func (client *Client) CreateBackupPlanWithOptions(request *CreateBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackupPlan"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup plan. Currently, only HBaseue is supported.
//
// Description:
//
// Currently, this operation supports only HBaseue clusters. The EnableHBaseueBackup operation is compatible with this operation and performs automatic creation of a backup plan after a BDS cluster is created.
//
// @param request - CreateBackupPlanRequest
//
// @return CreateBackupPlanResponse
func (client *Client) CreateBackupPlan(request *CreateBackupPlanRequest) (_result *CreateBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CreateBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an ApsaraDB for HBase cluster by calling CreateCluster.
//
// @param request - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ColdStorageSize) {
		query["ColdStorageSize"] = request.ColdStorageSize
	}

	if !dara.IsNil(request.CoreInstanceType) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.MasterInstanceType) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !dara.IsNil(request.NodeCount) {
		query["NodeCount"] = request.NodeCount
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
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
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ApsaraDB for HBase cluster by calling CreateCluster.
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a global resource in a cluster by calling CreateGlobalResource.
//
// @param request - CreateGlobalResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalResourceResponse
func (client *Client) CreateGlobalResourceWithOptions(request *CreateGlobalResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateGlobalResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGlobalResource"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGlobalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a global resource in a cluster by calling CreateGlobalResource.
//
// @param request - CreateGlobalResourceRequest
//
// @return CreateGlobalResourceResponse
func (client *Client) CreateGlobalResource(request *CreateGlobalResourceRequest) (_result *CreateGlobalResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGlobalResourceResponse{}
	_body, _err := client.CreateGlobalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a load balancing service by calling CreateHBaseSlbServer.
//
// @param request - CreateHBaseSlbServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHBaseSlbServerResponse
func (client *Client) CreateHBaseSlbServerWithOptions(request *CreateHBaseSlbServerRequest, runtime *dara.RuntimeOptions) (_result *CreateHBaseSlbServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.SlbServer) {
		query["SlbServer"] = request.SlbServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHBaseSlbServer"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHBaseSlbServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a load balancing service by calling CreateHBaseSlbServer.
//
// @param request - CreateHBaseSlbServerRequest
//
// @return CreateHBaseSlbServerResponse
func (client *Client) CreateHBaseSlbServer(request *CreateHBaseSlbServerRequest) (_result *CreateHBaseSlbServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHBaseSlbServerResponse{}
	_body, _err := client.CreateHBaseSlbServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a high-availability Thrift or Phoenix service for a BDS cluster that already has a high-availability HBase setup.
//
// @param request - CreateHbaseHaSlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHbaseHaSlbResponse
func (client *Client) CreateHbaseHaSlbWithOptions(request *CreateHbaseHaSlbRequest, runtime *dara.RuntimeOptions) (_result *CreateHbaseHaSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BdsId) {
		query["BdsId"] = request.BdsId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.HaId) {
		query["HaId"] = request.HaId
	}

	if !dara.IsNil(request.HaTypes) {
		query["HaTypes"] = request.HaTypes
	}

	if !dara.IsNil(request.HbaseType) {
		query["HbaseType"] = request.HbaseType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHbaseHaSlb"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHbaseHaSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a high-availability Thrift or Phoenix service for a BDS cluster that already has a high-availability HBase setup.
//
// @param request - CreateHbaseHaSlbRequest
//
// @return CreateHbaseHaSlbResponse
func (client *Client) CreateHbaseHaSlb(request *CreateHbaseHaSlbRequest) (_result *CreateHbaseHaSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHbaseHaSlbResponse{}
	_body, _err := client.CreateHbaseHaSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a multi-zone cluster by calling CreateMultiZoneCluster.
//
// Description:
//
// Currently, only version 2.0 of the HBaseue engine is supported.
//
// @param request - CreateMultiZoneClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultiZoneClusterResponse
func (client *Client) CreateMultiZoneClusterWithOptions(request *CreateMultiZoneClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateMultiZoneClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArbiterVSwitchId) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !dara.IsNil(request.ArbiterZoneId) {
		query["ArbiterZoneId"] = request.ArbiterZoneId
	}

	if !dara.IsNil(request.ArchVersion) {
		query["ArchVersion"] = request.ArchVersion
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.CoreDiskSize) {
		query["CoreDiskSize"] = request.CoreDiskSize
	}

	if !dara.IsNil(request.CoreDiskType) {
		query["CoreDiskType"] = request.CoreDiskType
	}

	if !dara.IsNil(request.CoreInstanceType) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !dara.IsNil(request.CoreNodeCount) {
		query["CoreNodeCount"] = request.CoreNodeCount
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.LogDiskSize) {
		query["LogDiskSize"] = request.LogDiskSize
	}

	if !dara.IsNil(request.LogDiskType) {
		query["LogDiskType"] = request.LogDiskType
	}

	if !dara.IsNil(request.LogInstanceType) {
		query["LogInstanceType"] = request.LogInstanceType
	}

	if !dara.IsNil(request.LogNodeCount) {
		query["LogNodeCount"] = request.LogNodeCount
	}

	if !dara.IsNil(request.MasterInstanceType) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !dara.IsNil(request.MultiZoneCombination) {
		query["MultiZoneCombination"] = request.MultiZoneCombination
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PrimaryVSwitchId) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !dara.IsNil(request.PrimaryZoneId) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !dara.IsNil(request.StandbyZoneId) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultiZoneCluster"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultiZoneClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a multi-zone cluster by calling CreateMultiZoneCluster.
//
// Description:
//
// Currently, only version 2.0 of the HBaseue engine is supported.
//
// @param request - CreateMultiZoneClusterRequest
//
// @return CreateMultiZoneClusterResponse
func (client *Client) CreateMultiZoneCluster(request *CreateMultiZoneClusterRequest) (_result *CreateMultiZoneClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMultiZoneClusterResponse{}
	_body, _err := client.CreateMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores backed-up data from a cluster with the backup and restore feature enabled to another ApsaraDB for HBase Performance-enhanced Edition cluster in the same region.
//
// Description:
//
// Before calling this operation, make sure that the backup and restore feature is enabled for the ApsaraDB for HBase Performance-enhanced Edition cluster and that the target ApsaraDB for HBase Performance-enhanced Edition cluster is associated with the corresponding BDS.
//
// @param request - CreateRestorePlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRestorePlanResponse
func (client *Client) CreateRestorePlanWithOptions(request *CreateRestorePlanRequest, runtime *dara.RuntimeOptions) (_result *CreateRestorePlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RestoreAllTable) {
		query["RestoreAllTable"] = request.RestoreAllTable
	}

	if !dara.IsNil(request.RestoreByCopy) {
		query["RestoreByCopy"] = request.RestoreByCopy
	}

	if !dara.IsNil(request.RestoreToDate) {
		query["RestoreToDate"] = request.RestoreToDate
	}

	if !dara.IsNil(request.Tables) {
		query["Tables"] = request.Tables
	}

	if !dara.IsNil(request.TargetClusterId) {
		query["TargetClusterId"] = request.TargetClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRestorePlan"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRestorePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores backed-up data from a cluster with the backup and restore feature enabled to another ApsaraDB for HBase Performance-enhanced Edition cluster in the same region.
//
// Description:
//
// Before calling this operation, make sure that the backup and restore feature is enabled for the ApsaraDB for HBase Performance-enhanced Edition cluster and that the target ApsaraDB for HBase Performance-enhanced Edition cluster is associated with the corresponding BDS.
//
// @param request - CreateRestorePlanRequest
//
// @return CreateRestorePlanResponse
func (client *Client) CreateRestorePlan(request *CreateRestorePlanRequest) (_result *CreateRestorePlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRestorePlanResponse{}
	_body, _err := client.CreateRestorePlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an HBase Serverless cluster.
//
// @param request - CreateServerlessClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerlessClusterResponse
func (client *Client) CreateServerlessClusterWithOptions(request *CreateServerlessClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateServerlessClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServerlessCapability) {
		query["ServerlessCapability"] = request.ServerlessCapability
	}

	if !dara.IsNil(request.ServerlessSpec) {
		query["ServerlessSpec"] = request.ServerlessSpec
	}

	if !dara.IsNil(request.ServerlessStorage) {
		query["ServerlessStorage"] = request.ServerlessStorage
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
		Action:      dara.String("CreateServerlessCluster"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServerlessClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an HBase Serverless cluster.
//
// @param request - CreateServerlessClusterRequest
//
// @return CreateServerlessClusterResponse
func (client *Client) CreateServerlessCluster(request *CreateServerlessClusterRequest) (_result *CreateServerlessClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServerlessClusterResponse{}
	_body, _err := client.CreateServerlessClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Wide Table Database Account
//
// Description:
//
// Only supported for HBase Enhanced Edition (HBaseue).
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
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2019-01-01"),
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
// # Delete Wide Table Database Account
//
// Description:
//
// Only supported for HBase Enhanced Edition (HBaseue).
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
// Deletes a global resource from a cluster by calling DeleteGlobalResource.
//
// @param request - DeleteGlobalResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGlobalResourceResponse
func (client *Client) DeleteGlobalResourceWithOptions(request *DeleteGlobalResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGlobalResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGlobalResource"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGlobalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a global resource from a cluster by calling DeleteGlobalResource.
//
// @param request - DeleteGlobalResourceRequest
//
// @return DeleteGlobalResourceResponse
func (client *Client) DeleteGlobalResource(request *DeleteGlobalResourceRequest) (_result *DeleteGlobalResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGlobalResourceResponse{}
	_body, _err := client.DeleteGlobalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a created HA instance.
//
// @param request - DeleteHBaseHaDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHBaseHaDBResponse
func (client *Client) DeleteHBaseHaDBWithOptions(request *DeleteHBaseHaDBRequest, runtime *dara.RuntimeOptions) (_result *DeleteHBaseHaDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BdsId) {
		query["BdsId"] = request.BdsId
	}

	if !dara.IsNil(request.HaId) {
		query["HaId"] = request.HaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHBaseHaDB"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHBaseHaDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a created HA instance.
//
// @param request - DeleteHBaseHaDBRequest
//
// @return DeleteHBaseHaDBResponse
func (client *Client) DeleteHBaseHaDB(request *DeleteHBaseHaDBRequest) (_result *DeleteHBaseHaDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHBaseHaDBResponse{}
	_body, _err := client.DeleteHBaseHaDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls DeleteHBaseSlbServer to delete an activated load balancing service.
//
// @param request - DeleteHBaseSlbServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHBaseSlbServerResponse
func (client *Client) DeleteHBaseSlbServerWithOptions(request *DeleteHBaseSlbServerRequest, runtime *dara.RuntimeOptions) (_result *DeleteHBaseSlbServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.SlbServer) {
		query["SlbServer"] = request.SlbServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHBaseSlbServer"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHBaseSlbServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DeleteHBaseSlbServer to delete an activated load balancing service.
//
// @param request - DeleteHBaseSlbServerRequest
//
// @return DeleteHBaseSlbServerResponse
func (client *Client) DeleteHBaseSlbServer(request *DeleteHBaseSlbServerRequest) (_result *DeleteHBaseSlbServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHBaseSlbServerResponse{}
	_body, _err := client.DeleteHBaseSlbServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the corresponding high-availability Thrift or high-availability Phoenix configuration. This operation is the counterpart of the CreateHbaseHaSlb operation.
//
// @param request - DeleteHbaseHaSlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHbaseHaSlbResponse
func (client *Client) DeleteHbaseHaSlbWithOptions(request *DeleteHbaseHaSlbRequest, runtime *dara.RuntimeOptions) (_result *DeleteHbaseHaSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BdsId) {
		query["BdsId"] = request.BdsId
	}

	if !dara.IsNil(request.HaId) {
		query["HaId"] = request.HaId
	}

	if !dara.IsNil(request.HaTypes) {
		query["HaTypes"] = request.HaTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHbaseHaSlb"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHbaseHaSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the corresponding high-availability Thrift or high-availability Phoenix configuration. This operation is the counterpart of the CreateHbaseHaSlb operation.
//
// @param request - DeleteHbaseHaSlbRequest
//
// @return DeleteHbaseHaSlbResponse
func (client *Client) DeleteHbaseHaSlb(request *DeleteHbaseHaSlbRequest) (_result *DeleteHbaseHaSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHbaseHaSlbResponse{}
	_body, _err := client.DeleteHbaseHaSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls DeleteInstance to release an HBase instance.
//
// Description:
//
// When you call this operation, the instance must meet the following conditions:
//
// - The instance status is **Running**.
//
// - The billing method of the instance is **pay-as-you-go**.
//
// > Subscription instances cannot be deleted by calling this operation. They are automatically released upon expiration. To release a subscription instance in advance, submit a ticket.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ImmediateDeleteFlag) {
		query["ImmediateDeleteFlag"] = request.ImmediateDeleteFlag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2019-01-01"),
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
// Calls DeleteInstance to release an HBase instance.
//
// Description:
//
// When you call this operation, the instance must meet the following conditions:
//
// - The instance status is **Running**.
//
// - The billing method of the instance is **pay-as-you-go**.
//
// > Subscription instances cannot be deleted by calling this operation. They are automatically released upon expiration. To release a subscription instance in advance, submit a ticket.
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
// Deletes a multi-zone instance by calling the DeleteMultiZoneCluster operation.
//
// @param request - DeleteMultiZoneClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultiZoneClusterResponse
func (client *Client) DeleteMultiZoneClusterWithOptions(request *DeleteMultiZoneClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultiZoneClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ImmediateDeleteFlag) {
		query["ImmediateDeleteFlag"] = request.ImmediateDeleteFlag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMultiZoneCluster"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultiZoneClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a multi-zone instance by calling the DeleteMultiZoneCluster operation.
//
// @param request - DeleteMultiZoneClusterRequest
//
// @return DeleteMultiZoneClusterResponse
func (client *Client) DeleteMultiZoneCluster(request *DeleteMultiZoneClusterRequest) (_result *DeleteMultiZoneClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMultiZoneClusterResponse{}
	_body, _err := client.DeleteMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an HBase Serverless cluster.
//
// @param request - DeleteServerlessClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServerlessClusterResponse
func (client *Client) DeleteServerlessClusterWithOptions(request *DeleteServerlessClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteServerlessClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServerlessCluster"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServerlessClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an HBase Serverless cluster.
//
// @param request - DeleteServerlessClusterRequest
//
// @return DeleteServerlessClusterResponse
func (client *Client) DeleteServerlessCluster(request *DeleteServerlessClusterRequest) (_result *DeleteServerlessClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteServerlessClusterResponse{}
	_body, _err := client.DeleteServerlessClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls DeleteUserHdfsInfo to delete user-created HDFS information in Xpack.
//
// @param request - DeleteUserHdfsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserHdfsInfoResponse
func (client *Client) DeleteUserHdfsInfoWithOptions(request *DeleteUserHdfsInfoRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserHdfsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NameService) {
		query["NameService"] = request.NameService
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserHdfsInfo"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserHdfsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DeleteUserHdfsInfo to delete user-created HDFS information in Xpack.
//
// @param request - DeleteUserHdfsInfoRequest
//
// @return DeleteUserHdfsInfoResponse
func (client *Client) DeleteUserHdfsInfo(request *DeleteUserHdfsInfoRequest) (_result *DeleteUserHdfsInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserHdfsInfoResponse{}
	_body, _err := client.DeleteUserHdfsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of database accounts.
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
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccounts"),
		Version:     dara.String("2019-01-01"),
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
// Queries the list of database accounts.
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
// Calls the DescribeActiveOperationTaskType operation to query the O&M task types, the number of tasks of each type, and the details of an HBase instance.
//
// @param request - DescribeActiveOperationTaskTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskTypeResponse
func (client *Client) DescribeActiveOperationTaskTypeWithOptions(request *DescribeActiveOperationTaskTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTaskTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DescribeActiveOperationTaskType"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTaskTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeActiveOperationTaskType operation to query the O&M task types, the number of tasks of each type, and the details of an HBase instance.
//
// @param request - DescribeActiveOperationTaskTypeRequest
//
// @return DescribeActiveOperationTaskTypeResponse
func (client *Client) DescribeActiveOperationTaskType(request *DescribeActiveOperationTaskTypeRequest) (_result *DescribeActiveOperationTaskTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveOperationTaskTypeResponse{}
	_body, _err := client.DescribeActiveOperationTaskTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DescribeActiveOperationTasks operation to query details of O&M tasks for HBase instances.
//
// @param request - DescribeActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasksWithOptions(request *DescribeActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2019-01-01"),
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
// Calls the DescribeActiveOperationTasks operation to query details of O&M tasks for HBase instances.
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
// Queries available resources by calling DescribeAvailableResource.
//
// @param request - DescribeAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableResourceResponse
func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.CoreInstanceType) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableResource"),
		Version:     dara.String("2019-01-01"),
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
// Queries available resources by calling DescribeAvailableResource.
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
// Queries the configuration of a backup plan by calling the DescribeBackupPlanConfig operation.
//
// Description:
//
// This operation currently supports only HBaseue. Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue instance.
//
// @param request - DescribeBackupPlanConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPlanConfigResponse
func (client *Client) DescribeBackupPlanConfigWithOptions(request *DescribeBackupPlanConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPlanConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPlanConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPlanConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of a backup plan by calling the DescribeBackupPlanConfig operation.
//
// Description:
//
// This operation currently supports only HBaseue. Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue instance.
//
// @param request - DescribeBackupPlanConfigRequest
//
// @return DescribeBackupPlanConfigResponse
func (client *Client) DescribeBackupPlanConfig(request *DescribeBackupPlanConfigRequest) (_result *DescribeBackupPlanConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupPlanConfigResponse{}
	_body, _err := client.DescribeBackupPlanConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup policy of an HBase cluster by calling DescribeBackupPolicy.
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
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2019-01-01"),
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
// Queries the backup policy of an HBase cluster by calling DescribeBackupPolicy.
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
// Describes the status of a backup task by calling the DescribeBackupStatus operation.
//
// @param request - DescribeBackupStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupStatusResponse
func (client *Client) DescribeBackupStatusWithOptions(request *DescribeBackupStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupStatus"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the status of a backup task by calling the DescribeBackupStatus operation.
//
// @param request - DescribeBackupStatusRequest
//
// @return DescribeBackupStatusResponse
func (client *Client) DescribeBackupStatus(request *DescribeBackupStatusRequest) (_result *DescribeBackupStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupStatusResponse{}
	_body, _err := client.DescribeBackupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After HBase Enhanced Edition backup is enabled, you can call the DescribeBackupSummary operation to query backup details.
//
// Description:
//
// The backup and restore feature of the HBase Enhanced Edition instance is enabled. For more information, see [Enable backup and restore](https://help.aliyun.com/document_detail/145767.html).
//
// @param request - DescribeBackupSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSummaryResponse
func (client *Client) DescribeBackupSummaryWithOptions(request *DescribeBackupSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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
		Action:      dara.String("DescribeBackupSummary"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After HBase Enhanced Edition backup is enabled, you can call the DescribeBackupSummary operation to query backup details.
//
// Description:
//
// The backup and restore feature of the HBase Enhanced Edition instance is enabled. For more information, see [Enable backup and restore](https://help.aliyun.com/document_detail/145767.html).
//
// @param request - DescribeBackupSummaryRequest
//
// @return DescribeBackupSummaryResponse
func (client *Client) DescribeBackupSummary(request *DescribeBackupSummaryRequest) (_result *DescribeBackupSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupSummaryResponse{}
	_body, _err := client.DescribeBackupSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup table information of a specific backup record by calling the DescribeBackupTables operation.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that backup records exist. You can call the DescribeBackupSummary operation to obtain backup records.
//
// @param request - DescribeBackupTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupTablesResponse
func (client *Client) DescribeBackupTablesWithOptions(request *DescribeBackupTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupRecordId) {
		query["BackupRecordId"] = request.BackupRecordId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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
		Action:      dara.String("DescribeBackupTables"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup table information of a specific backup record by calling the DescribeBackupTables operation.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that backup records exist. You can call the DescribeBackupSummary operation to obtain backup records.
//
// @param request - DescribeBackupTablesRequest
//
// @return DescribeBackupTablesResponse
func (client *Client) DescribeBackupTables(request *DescribeBackupTablesRequest) (_result *DescribeBackupTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupTablesResponse{}
	_body, _err := client.DescribeBackupTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves backup records of an ApsaraDB for HBase cluster by calling DescribeBackups.
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

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndTimeUTC) {
		query["EndTimeUTC"] = request.EndTimeUTC
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

	if !dara.IsNil(request.StartTimeUTC) {
		query["StartTimeUTC"] = request.StartTimeUTC
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackups"),
		Version:     dara.String("2019-01-01"),
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
// Retrieves backup records of an ApsaraDB for HBase cluster by calling DescribeBackups.
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
// Queries the connection list of a cluster by calling DescribeClusterConnection.
//
// @param request - DescribeClusterConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterConnectionResponse
func (client *Client) DescribeClusterConnectionWithOptions(request *DescribeClusterConnectionRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterConnection"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the connection list of a cluster by calling DescribeClusterConnection.
//
// @param request - DescribeClusterConnectionRequest
//
// @return DescribeClusterConnectionResponse
func (client *Client) DescribeClusterConnection(request *DescribeClusterConnectionRequest) (_result *DescribeClusterConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClusterConnectionResponse{}
	_body, _err := client.DescribeClusterConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls DescribeColdStorage to retrieve the cold storage information of an instance.
//
// @param request - DescribeColdStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColdStorageResponse
func (client *Client) DescribeColdStorageWithOptions(request *DescribeColdStorageRequest, runtime *dara.RuntimeOptions) (_result *DescribeColdStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColdStorage"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColdStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeColdStorage to retrieve the cold storage information of an instance.
//
// @param request - DescribeColdStorageRequest
//
// @return DescribeColdStorageResponse
func (client *Client) DescribeColdStorage(request *DescribeColdStorageRequest) (_result *DescribeColdStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeColdStorageResponse{}
	_body, _err := client.DescribeColdStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls DescribeDBInstanceUsage to retrieve the running watermark status of an HBase cluster by cluster ID.
//
// @param request - DescribeDBInstanceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceUsageResponse
func (client *Client) DescribeDBInstanceUsageWithOptions(request *DescribeDBInstanceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceUsage"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeDBInstanceUsage to retrieve the running watermark status of an HBase cluster by cluster ID.
//
// @param request - DescribeDBInstanceUsageRequest
//
// @return DescribeDBInstanceUsageResponse
func (client *Client) DescribeDBInstanceUsage(request *DescribeDBInstanceUsageRequest) (_result *DescribeDBInstanceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceUsageResponse{}
	_body, _err := client.DescribeDBInstanceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of clusters that have been deleted but not fully released.
//
// Description:
//
// By default, cluster resources are fully cleaned up 7 days after deletion. If you have called the PurgeInstance operation or set ImmediateDeleteFlag to true when calling the delete operation, the resources have already been fully cleaned up.
//
// @param request - DescribeDeletedInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeletedInstancesResponse
func (client *Client) DescribeDeletedInstancesWithOptions(request *DescribeDeletedInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeletedInstancesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeletedInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeletedInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of clusters that have been deleted but not fully released.
//
// Description:
//
// By default, cluster resources are fully cleaned up 7 days after deletion. If you have called the PurgeInstance operation or set ImmediateDeleteFlag to true when calling the delete operation, the resources have already been fully cleaned up.
//
// @param request - DescribeDeletedInstancesRequest
//
// @return DescribeDeletedInstancesResponse
func (client *Client) DescribeDeletedInstances(request *DescribeDeletedInstancesRequest) (_result *DescribeDeletedInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeletedInstancesResponse{}
	_body, _err := client.DescribeDeletedInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Currently, only HBase and HBaseue are supported.
//
// @param request - DescribeDiskWarningLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskWarningLineResponse
func (client *Client) DescribeDiskWarningLineWithOptions(request *DescribeDiskWarningLineRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskWarningLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiskWarningLine"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiskWarningLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Currently, only HBase and HBaseue are supported.
//
// @param request - DescribeDiskWarningLineRequest
//
// @return DescribeDiskWarningLineResponse
func (client *Client) DescribeDiskWarningLine(request *DescribeDiskWarningLineRequest) (_result *DescribeDiskWarningLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiskWarningLineResponse{}
	_body, _err := client.DescribeDiskWarningLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the database connection information of an ApsaraDB for HBase instance by calling DescribeEndpoints.
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
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEndpoints"),
		Version:     dara.String("2019-01-01"),
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
// Queries the database connection information of an ApsaraDB for HBase instance by calling DescribeEndpoints.
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
// Queries the details of an ApsaraDB for HBase instance.
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an ApsaraDB for HBase instance.
//
// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instance specifications (CPU and memory) by calling DescribeInstanceType.
//
// Description:
//
// The complete specification mapping table is based on the list returned by this operation.
//
// If the response returns an "instanceType is not found" error, the specification has been deprecated. If you have a strong dependency on this specification, contact the Alibaba Cloud HBase management team.
//
// @param request - DescribeInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceTypeResponse
func (client *Client) DescribeInstanceTypeWithOptions(request *DescribeInstanceTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceType"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instance specifications (CPU and memory) by calling DescribeInstanceType.
//
// Description:
//
// The complete specification mapping table is based on the list returned by this operation.
//
// If the response returns an "instanceType is not found" error, the specification has been deprecated. If you have a strong dependency on this specification, contact the Alibaba Cloud HBase management team.
//
// @param request - DescribeInstanceTypeRequest
//
// @return DescribeInstanceTypeResponse
func (client *Client) DescribeInstanceType(request *DescribeInstanceTypeRequest) (_result *DescribeInstanceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceTypeResponse{}
	_body, _err := client.DescribeInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of ApsaraDB for HBase instances in a specified region by calling DescribeInstances.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2019-01-01"),
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
// Queries the list of ApsaraDB for HBase instances in a specified region by calling DescribeInstances.
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
// Retrieves the IP whitelist information of a cluster by cluster ID.
//
// @param request - DescribeIpWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpWhitelistResponse
func (client *Client) DescribeIpWhitelistWithOptions(request *DescribeIpWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpWhitelist"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the IP whitelist information of a cluster by cluster ID.
//
// @param request - DescribeIpWhitelistRequest
//
// @return DescribeIpWhitelistResponse
func (client *Client) DescribeIpWhitelist(request *DescribeIpWhitelistRequest) (_result *DescribeIpWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpWhitelistResponse{}
	_body, _err := client.DescribeIpWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available region and zone combinations for multi-zone deployment by calling DescribeMultiZoneAvailableRegions.
//
// @param request - DescribeMultiZoneAvailableRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiZoneAvailableRegionsResponse
func (client *Client) DescribeMultiZoneAvailableRegionsWithOptions(request *DescribeMultiZoneAvailableRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMultiZoneAvailableRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMultiZoneAvailableRegions"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMultiZoneAvailableRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available region and zone combinations for multi-zone deployment by calling DescribeMultiZoneAvailableRegions.
//
// @param request - DescribeMultiZoneAvailableRegionsRequest
//
// @return DescribeMultiZoneAvailableRegionsResponse
func (client *Client) DescribeMultiZoneAvailableRegions(request *DescribeMultiZoneAvailableRegionsRequest) (_result *DescribeMultiZoneAvailableRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMultiZoneAvailableRegionsResponse{}
	_body, _err := client.DescribeMultiZoneAvailableRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries purchasable resources in multiple zones by calling DescribeMultiZoneAvailableResource.
//
// @param request - DescribeMultiZoneAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiZoneAvailableResourceResponse
func (client *Client) DescribeMultiZoneAvailableResourceWithOptions(request *DescribeMultiZoneAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeMultiZoneAvailableResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneCombination) {
		query["ZoneCombination"] = request.ZoneCombination
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMultiZoneAvailableResource"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMultiZoneAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries purchasable resources in multiple zones by calling DescribeMultiZoneAvailableResource.
//
// @param request - DescribeMultiZoneAvailableResourceRequest
//
// @return DescribeMultiZoneAvailableResourceResponse
func (client *Client) DescribeMultiZoneAvailableResource(request *DescribeMultiZoneAvailableResourceRequest) (_result *DescribeMultiZoneAvailableResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMultiZoneAvailableResourceResponse{}
	_body, _err := client.DescribeMultiZoneAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a multi-zone instance.
//
// @param request - DescribeMultiZoneClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMultiZoneClusterResponse
func (client *Client) DescribeMultiZoneClusterWithOptions(request *DescribeMultiZoneClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeMultiZoneClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMultiZoneCluster"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMultiZoneClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a multi-zone instance.
//
// @param request - DescribeMultiZoneClusterRequest
//
// @return DescribeMultiZoneClusterResponse
func (client *Client) DescribeMultiZoneCluster(request *DescribeMultiZoneClusterRequest) (_result *DescribeMultiZoneClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMultiZoneClusterResponse{}
	_body, _err := client.DescribeMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the recoverable time range after backup is enabled for an HBaseue cluster.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster.
//
// @param request - DescribeRecoverableTimeRangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecoverableTimeRangeResponse
func (client *Client) DescribeRecoverableTimeRangeWithOptions(request *DescribeRecoverableTimeRangeRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecoverableTimeRangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecoverableTimeRange"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecoverableTimeRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the recoverable time range after backup is enabled for an HBaseue cluster.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster.
//
// @param request - DescribeRecoverableTimeRangeRequest
//
// @return DescribeRecoverableTimeRangeResponse
func (client *Client) DescribeRecoverableTimeRange(request *DescribeRecoverableTimeRangeRequest) (_result *DescribeRecoverableTimeRangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecoverableTimeRangeResponse{}
	_body, _err := client.DescribeRecoverableTimeRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available regions by calling DescribeRegions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-01-01"),
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
// Queries available regions by calling DescribeRegions.
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
// Queries the details of a specific full restoration for an HBaseue cluster that has backup and recovery enabled.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that restoration records exist. You can call DescribeRestoreSummary to obtain restoration records.
//
// @param request - DescribeRestoreFullDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreFullDetailsResponse
func (client *Client) DescribeRestoreFullDetailsWithOptions(request *DescribeRestoreFullDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreFullDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RestoreRecordId) {
		query["RestoreRecordId"] = request.RestoreRecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreFullDetails"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreFullDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific full restoration for an HBaseue cluster that has backup and recovery enabled.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that restoration records exist. You can call DescribeRestoreSummary to obtain restoration records.
//
// @param request - DescribeRestoreFullDetailsRequest
//
// @return DescribeRestoreFullDetailsResponse
func (client *Client) DescribeRestoreFullDetails(request *DescribeRestoreFullDetailsRequest) (_result *DescribeRestoreFullDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRestoreFullDetailsResponse{}
	_body, _err := client.DescribeRestoreFullDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an incremental restoration by calling the DescribeRestoreIncrDetail operation.
//
// Description:
//
// Before you call this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that restoration records exist. You can call the DescribeRestoreSummary operation to obtain restoration records.
//
// @param request - DescribeRestoreIncrDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreIncrDetailResponse
func (client *Client) DescribeRestoreIncrDetailWithOptions(request *DescribeRestoreIncrDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreIncrDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RestoreRecordId) {
		query["RestoreRecordId"] = request.RestoreRecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreIncrDetail"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreIncrDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an incremental restoration by calling the DescribeRestoreIncrDetail operation.
//
// Description:
//
// Before you call this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that restoration records exist. You can call the DescribeRestoreSummary operation to obtain restoration records.
//
// @param request - DescribeRestoreIncrDetailRequest
//
// @return DescribeRestoreIncrDetailResponse
func (client *Client) DescribeRestoreIncrDetail(request *DescribeRestoreIncrDetailRequest) (_result *DescribeRestoreIncrDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRestoreIncrDetailResponse{}
	_body, _err := client.DescribeRestoreIncrDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of schema restoration for an HBaseue cluster that has backup and recovery enabled.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that restoration records exist.
//
// @param request - DescribeRestoreSchemaDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreSchemaDetailsResponse
func (client *Client) DescribeRestoreSchemaDetailsWithOptions(request *DescribeRestoreSchemaDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreSchemaDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RestoreRecordId) {
		query["RestoreRecordId"] = request.RestoreRecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreSchemaDetails"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreSchemaDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of schema restoration for an HBaseue cluster that has backup and recovery enabled.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that restoration records exist.
//
// @param request - DescribeRestoreSchemaDetailsRequest
//
// @return DescribeRestoreSchemaDetailsResponse
func (client *Client) DescribeRestoreSchemaDetails(request *DescribeRestoreSchemaDetailsRequest) (_result *DescribeRestoreSchemaDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRestoreSchemaDetailsResponse{}
	_body, _err := client.DescribeRestoreSchemaDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the data recovery summary for an HBaseue cluster that has backup and recovery enabled by calling the DescribeRestoreSummary operation.
//
// Description:
//
// Before calling this operation, make sure that backup and recovery is enabled for the HBaseue cluster.
//
// @param request - DescribeRestoreSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreSummaryResponse
func (client *Client) DescribeRestoreSummaryWithOptions(request *DescribeRestoreSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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
		Action:      dara.String("DescribeRestoreSummary"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the data recovery summary for an HBaseue cluster that has backup and recovery enabled by calling the DescribeRestoreSummary operation.
//
// Description:
//
// Before calling this operation, make sure that backup and recovery is enabled for the HBaseue cluster.
//
// @param request - DescribeRestoreSummaryRequest
//
// @return DescribeRestoreSummaryResponse
func (client *Client) DescribeRestoreSummary(request *DescribeRestoreSummaryRequest) (_result *DescribeRestoreSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRestoreSummaryResponse{}
	_body, _err := client.DescribeRestoreSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a restore record for a cluster that has been restored from a backup.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that restore records exist. You can call DescribeRestoreSummary to obtain restore records.
//
// @param request - DescribeRestoreTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreTablesResponse
func (client *Client) DescribeRestoreTablesWithOptions(request *DescribeRestoreTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RestoreRecordId) {
		query["RestoreRecordId"] = request.RestoreRecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreTables"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a restore record for a cluster that has been restored from a backup.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster and that restore records exist. You can call DescribeRestoreSummary to obtain restore records.
//
// @param request - DescribeRestoreTablesRequest
//
// @return DescribeRestoreTablesResponse
func (client *Client) DescribeRestoreTables(request *DescribeRestoreTablesRequest) (_result *DescribeRestoreTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRestoreTablesResponse{}
	_body, _err := client.DescribeRestoreTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls DescribeSecurityGroups to retrieve the security group information of a cluster by cluster ID.
//
// @param request - DescribeSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityGroupsResponse
func (client *Client) DescribeSecurityGroupsWithOptions(request *DescribeSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityGroups"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeSecurityGroups to retrieve the security group information of a cluster by cluster ID.
//
// @param request - DescribeSecurityGroupsRequest
//
// @return DescribeSecurityGroupsResponse
func (client *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (_result *DescribeSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.DescribeSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a serverless cluster by calling DescribeServerlessCluster.
//
// @param request - DescribeServerlessClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServerlessClusterResponse
func (client *Client) DescribeServerlessClusterWithOptions(request *DescribeServerlessClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeServerlessClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServerlessCluster"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServerlessClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a serverless cluster by calling DescribeServerlessCluster.
//
// @param request - DescribeServerlessClusterRequest
//
// @return DescribeServerlessClusterResponse
func (client *Client) DescribeServerlessCluster(request *DescribeServerlessClusterRequest) (_result *DescribeServerlessClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServerlessClusterResponse{}
	_body, _err := client.DescribeServerlessClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls DescribeSubDomain to retrieve an available subdomain.
//
// @param request - DescribeSubDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubDomainResponse
func (client *Client) DescribeSubDomainWithOptions(request *DescribeSubDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubDomainResponse, _err error) {
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

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubDomain"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeSubDomain to retrieve an available subdomain.
//
// @param request - DescribeSubDomainRequest
//
// @return DescribeSubDomainResponse
func (client *Client) DescribeSubDomain(request *DescribeSubDomainRequest) (_result *DescribeSubDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSubDomainResponse{}
	_body, _err := client.DescribeSubDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables backup and recovery for an HBaseue cluster.
//
// @param request - EnableHBaseueBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHBaseueBackupResponse
func (client *Client) EnableHBaseueBackupWithOptions(request *EnableHBaseueBackupRequest, runtime *dara.RuntimeOptions) (_result *EnableHBaseueBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ColdStorageSize) {
		query["ColdStorageSize"] = request.ColdStorageSize
	}

	if !dara.IsNil(request.HbaseueClusterId) {
		query["HbaseueClusterId"] = request.HbaseueClusterId
	}

	if !dara.IsNil(request.NodeCount) {
		query["NodeCount"] = request.NodeCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableHBaseueBackup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableHBaseueBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables backup and recovery for an HBaseue cluster.
//
// @param request - EnableHBaseueBackupRequest
//
// @return EnableHBaseueBackupResponse
func (client *Client) EnableHBaseueBackup(request *EnableHBaseueBackupRequest) (_result *EnableHBaseueBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableHBaseueBackupResponse{}
	_body, _err := client.EnableHBaseueBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the EnableHBaseueModule operation to enable an extension service.
//
// @param request - EnableHBaseueModuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHBaseueModuleResponse
func (client *Client) EnableHBaseueModuleWithOptions(request *EnableHBaseueModuleRequest, runtime *dara.RuntimeOptions) (_result *EnableHBaseueModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.BdsId) {
		query["BdsId"] = request.BdsId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CoreInstanceType) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.HbaseueClusterId) {
		query["HbaseueClusterId"] = request.HbaseueClusterId
	}

	if !dara.IsNil(request.MasterInstanceType) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !dara.IsNil(request.ModuleClusterName) {
		query["ModuleClusterName"] = request.ModuleClusterName
	}

	if !dara.IsNil(request.ModuleTypeName) {
		query["ModuleTypeName"] = request.ModuleTypeName
	}

	if !dara.IsNil(request.NodeCount) {
		query["NodeCount"] = request.NodeCount
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("EnableHBaseueModule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableHBaseueModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the EnableHBaseueModule operation to enable an extension service.
//
// @param request - EnableHBaseueModuleRequest
//
// @return EnableHBaseueModuleResponse
func (client *Client) EnableHBaseueModule(request *EnableHBaseueModuleRequest) (_result *EnableHBaseueModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableHBaseueModuleResponse{}
	_body, _err := client.EnableHBaseueModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the EvaluateMultiZoneResource operation to evaluate whether available resources exist.
//
// @param request - EvaluateMultiZoneResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateMultiZoneResourceResponse
func (client *Client) EvaluateMultiZoneResourceWithOptions(request *EvaluateMultiZoneResourceRequest, runtime *dara.RuntimeOptions) (_result *EvaluateMultiZoneResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArbiterVSwitchId) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !dara.IsNil(request.ArbiterZoneId) {
		query["ArbiterZoneId"] = request.ArbiterZoneId
	}

	if !dara.IsNil(request.ArchVersion) {
		query["ArchVersion"] = request.ArchVersion
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.CoreDiskSize) {
		query["CoreDiskSize"] = request.CoreDiskSize
	}

	if !dara.IsNil(request.CoreDiskType) {
		query["CoreDiskType"] = request.CoreDiskType
	}

	if !dara.IsNil(request.CoreInstanceType) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !dara.IsNil(request.CoreNodeCount) {
		query["CoreNodeCount"] = request.CoreNodeCount
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.LogDiskSize) {
		query["LogDiskSize"] = request.LogDiskSize
	}

	if !dara.IsNil(request.LogDiskType) {
		query["LogDiskType"] = request.LogDiskType
	}

	if !dara.IsNil(request.LogInstanceType) {
		query["LogInstanceType"] = request.LogInstanceType
	}

	if !dara.IsNil(request.LogNodeCount) {
		query["LogNodeCount"] = request.LogNodeCount
	}

	if !dara.IsNil(request.MasterInstanceType) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	if !dara.IsNil(request.MultiZoneCombination) {
		query["MultiZoneCombination"] = request.MultiZoneCombination
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PrimaryVSwitchId) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !dara.IsNil(request.PrimaryZoneId) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !dara.IsNil(request.StandbyZoneId) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EvaluateMultiZoneResource"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EvaluateMultiZoneResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the EvaluateMultiZoneResource operation to evaluate whether available resources exist.
//
// @param request - EvaluateMultiZoneResourceRequest
//
// @return EvaluateMultiZoneResourceResponse
func (client *Client) EvaluateMultiZoneResource(request *EvaluateMultiZoneResourceRequest) (_result *EvaluateMultiZoneResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EvaluateMultiZoneResourceResponse{}
	_body, _err := client.EvaluateMultiZoneResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the monitoring URL based on the cluster ID.
//
// Description:
//
// CloudMonitor has been migrated. This operation is no longer in use. Access monitoring from the Monitoring and Alerts page in cluster management.
//
// @param request - GetMultimodeCmsUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultimodeCmsUrlResponse
func (client *Client) GetMultimodeCmsUrlWithOptions(request *GetMultimodeCmsUrlRequest, runtime *dara.RuntimeOptions) (_result *GetMultimodeCmsUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMultimodeCmsUrl"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMultimodeCmsUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the monitoring URL based on the cluster ID.
//
// Description:
//
// CloudMonitor has been migrated. This operation is no longer in use. Access monitoring from the Monitoring and Alerts page in cluster management.
//
// @param request - GetMultimodeCmsUrlRequest
//
// @return GetMultimodeCmsUrlResponse
func (client *Client) GetMultimodeCmsUrl(request *GetMultimodeCmsUrlRequest) (_result *GetMultimodeCmsUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMultimodeCmsUrlResponse{}
	_body, _err := client.GetMultimodeCmsUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions to an account.
//
// @param request - GrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantResponse
func (client *Client) GrantWithOptions(request *GrantRequest, runtime *dara.RuntimeOptions) (_result *GrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AclActions) {
		query["AclActions"] = request.AclActions
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Grant"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions to an account.
//
// @param request - GrantRequest
//
// @return GrantResponse
func (client *Client) Grant(request *GrantRequest) (_result *GrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantResponse{}
	_body, _err := client.GrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ListHBaseInstances to retrieve ApsaraDB for HBase instances within the same VPC.
//
// @param request - ListHBaseInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHBaseInstancesResponse
func (client *Client) ListHBaseInstancesWithOptions(request *ListHBaseInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListHBaseInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHBaseInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHBaseInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls ListHBaseInstances to retrieve ApsaraDB for HBase instances within the same VPC.
//
// @param request - ListHBaseInstancesRequest
//
// @return ListHBaseInstancesResponse
func (client *Client) ListHBaseInstances(request *ListHBaseInstancesRequest) (_result *ListHBaseInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHBaseInstancesResponse{}
	_body, _err := client.ListHBaseInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the modification history of cluster parameters by calling ListInstanceServiceConfigHistories.
//
// @param request - ListInstanceServiceConfigHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceServiceConfigHistoriesResponse
func (client *Client) ListInstanceServiceConfigHistoriesWithOptions(request *ListInstanceServiceConfigHistoriesRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceServiceConfigHistoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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
		Action:      dara.String("ListInstanceServiceConfigHistories"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceServiceConfigHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the modification history of cluster parameters by calling ListInstanceServiceConfigHistories.
//
// @param request - ListInstanceServiceConfigHistoriesRequest
//
// @return ListInstanceServiceConfigHistoriesResponse
func (client *Client) ListInstanceServiceConfigHistories(request *ListInstanceServiceConfigHistoriesRequest) (_result *ListInstanceServiceConfigHistoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceServiceConfigHistoriesResponse{}
	_body, _err := client.ListInstanceServiceConfigHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration items of a cluster by calling ListInstanceServiceConfigurations.
//
// @param request - ListInstanceServiceConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceServiceConfigurationsResponse
func (client *Client) ListInstanceServiceConfigurationsWithOptions(request *ListInstanceServiceConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceServiceConfigurationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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
		Action:      dara.String("ListInstanceServiceConfigurations"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceServiceConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration items of a cluster by calling ListInstanceServiceConfigurations.
//
// @param request - ListInstanceServiceConfigurationsRequest
//
// @return ListInstanceServiceConfigurationsResponse
func (client *Client) ListInstanceServiceConfigurations(request *ListInstanceServiceConfigurationsRequest) (_result *ListInstanceServiceConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceServiceConfigurationsResponse{}
	_body, _err := client.ListInstanceServiceConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of tags by ResourceId or Tag (query by Key only, or by both Key and Value).
//
// Description:
//
// You must specify at least one of the ResourceId and Tag parameters. Otherwise, an error is returned.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-01-01"),
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
// Retrieves a list of tags by ResourceId or Tag (query by Key only, or by both Key and Value).
//
// Description:
//
// You must specify at least one of the ResourceId and Tag parameters. Otherwise, an error is returned.
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
// Retrieves all labels in a specified region.
//
// @param request - ListTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithOptions(request *ListTagsRequest, runtime *dara.RuntimeOptions) (_result *ListTagsResponse, _err error) {
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
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTags"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all labels in a specified region.
//
// @param request - ListTagsRequest
//
// @return ListTagsResponse
func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the password of a database access account.
//
// @param request - ModifyAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPasswordResponse
func (client *Client) ModifyAccountPasswordWithOptions(request *ModifyAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NewAccountPassword) {
		query["NewAccountPassword"] = request.NewAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountPassword"),
		Version:     dara.String("2019-01-01"),
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
// Changes the password of a database access account.
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
// Calls the ModifyActiveOperationTasks operation to modify the scheduled switchover time of O&M tasks.
//
// Description:
//
// In addition to notifications sent by text message, phone call, email, or internal message, O&M events of ApsaraDB for HBase (such as minor engine version updates) are also displayed in the console. In addition to calling this operation to modify the scheduled switchover time, you can also modify it in the console. For more information, see [Query or manage pending events](https://help.aliyun.com/document_detail/405057.html).
//
// @param request - ModifyActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasksWithOptions(request *ModifyActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2019-01-01"),
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
// Calls the ModifyActiveOperationTasks operation to modify the scheduled switchover time of O&M tasks.
//
// Description:
//
// In addition to notifications sent by text message, phone call, email, or internal message, O&M events of ApsaraDB for HBase (such as minor engine version updates) are also displayed in the console. In addition to calling this operation to modify the scheduled switchover time, you can also modify it in the console. For more information, see [Query or manage pending events](https://help.aliyun.com/document_detail/405057.html).
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
// Modifies the backup configuration for an HBaseue cluster that has backup and recovery enabled.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster.
//
// @param request - ModifyBackupPlanConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPlanConfigResponse
func (client *Client) ModifyBackupPlanConfigWithOptions(request *ModifyBackupPlanConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPlanConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.FullBackupCycle) {
		query["FullBackupCycle"] = request.FullBackupCycle
	}

	if !dara.IsNil(request.MinHFileBackupCount) {
		query["MinHFileBackupCount"] = request.MinHFileBackupCount
	}

	if !dara.IsNil(request.NextFullBackupDate) {
		query["NextFullBackupDate"] = request.NextFullBackupDate
	}

	if !dara.IsNil(request.Tables) {
		query["Tables"] = request.Tables
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPlanConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupPlanConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the backup configuration for an HBaseue cluster that has backup and recovery enabled.
//
// Description:
//
// Before calling this operation, make sure that the backup and recovery feature is enabled for the HBaseue cluster.
//
// @param request - ModifyBackupPlanConfigRequest
//
// @return ModifyBackupPlanConfigResponse
func (client *Client) ModifyBackupPlanConfig(request *ModifyBackupPlanConfigRequest) (_result *ModifyBackupPlanConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackupPlanConfigResponse{}
	_body, _err := client.ModifyBackupPlanConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ModifyBackupPolicy to modify the backup plan of an HBase cluster.
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
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PreferredBackupEndTimeUTC) {
		query["PreferredBackupEndTimeUTC"] = request.PreferredBackupEndTimeUTC
	}

	if !dara.IsNil(request.PreferredBackupPeriod) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !dara.IsNil(request.PreferredBackupStartTimeUTC) {
		query["PreferredBackupStartTimeUTC"] = request.PreferredBackupStartTimeUTC
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2019-01-01"),
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
// Calls ModifyBackupPolicy to modify the backup plan of an HBase cluster.
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
// Sets the deletion protection attribute of an instance by calling ModifyClusterDeletionProtection.
//
// @param request - ModifyClusterDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterDeletionProtectionResponse
func (client *Client) ModifyClusterDeletionProtectionWithOptions(request *ModifyClusterDeletionProtectionRequest, runtime *dara.RuntimeOptions) (_result *ModifyClusterDeletionProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Protection) {
		query["Protection"] = request.Protection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterDeletionProtection"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterDeletionProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the deletion protection attribute of an instance by calling ModifyClusterDeletionProtection.
//
// @param request - ModifyClusterDeletionProtectionRequest
//
// @return ModifyClusterDeletionProtectionResponse
func (client *Client) ModifyClusterDeletionProtection(request *ModifyClusterDeletionProtectionRequest) (_result *ModifyClusterDeletionProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyClusterDeletionProtectionResponse{}
	_body, _err := client.ModifyClusterDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the disk capacity alert threshold for HBase Cluster Edition and HBaseue clusters. The default alert threshold is 80%.
//
// @param request - ModifyDiskWarningLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskWarningLineResponse
func (client *Client) ModifyDiskWarningLineWithOptions(request *ModifyDiskWarningLineRequest, runtime *dara.RuntimeOptions) (_result *ModifyDiskWarningLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.WarningLine) {
		query["WarningLine"] = request.WarningLine
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDiskWarningLine"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDiskWarningLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the disk capacity alert threshold for HBase Cluster Edition and HBaseue clusters. The default alert threshold is 80%.
//
// @param request - ModifyDiskWarningLineRequest
//
// @return ModifyDiskWarningLineResponse
func (client *Client) ModifyDiskWarningLine(request *ModifyDiskWarningLineRequest) (_result *ModifyDiskWarningLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDiskWarningLineResponse{}
	_body, _err := client.ModifyDiskWarningLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the ModifyInstanceMaintainTime operation to modify the O&M window start time of an instance.
//
// @param request - ModifyInstanceMaintainTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMaintainTimeResponse
func (client *Client) ModifyInstanceMaintainTimeWithOptions(request *ModifyInstanceMaintainTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMaintainTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaintainEndTime) {
		query["MaintainEndTime"] = request.MaintainEndTime
	}

	if !dara.IsNil(request.MaintainStartTime) {
		query["MaintainStartTime"] = request.MaintainStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceMaintainTime"),
		Version:     dara.String("2019-01-01"),
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
// Calls the ModifyInstanceMaintainTime operation to modify the O&M window start time of an instance.
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
// Modifies the name of an instance by calling ModifyInstanceName.
//
// @param request - ModifyInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceNameResponse
func (client *Client) ModifyInstanceNameWithOptions(request *ModifyInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceName"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of an instance by calling ModifyInstanceName.
//
// @param request - ModifyInstanceNameRequest
//
// @return ModifyInstanceNameResponse
func (client *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (_result *ModifyInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.ModifyInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ModifyInstanceServiceConfig to modify cluster configurations. You can call the ListInstanceServiceConfigurations operation to query the configuration items that can be modified.
//
// @param request - ModifyInstanceServiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceServiceConfigResponse
func (client *Client) ModifyInstanceServiceConfigWithOptions(request *ModifyInstanceServiceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceServiceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConfigureName) {
		query["ConfigureName"] = request.ConfigureName
	}

	if !dara.IsNil(request.ConfigureValue) {
		query["ConfigureValue"] = request.ConfigureValue
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceServiceConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls ModifyInstanceServiceConfig to modify cluster configurations. You can call the ListInstanceServiceConfigurations operation to query the configuration items that can be modified.
//
// @param request - ModifyInstanceServiceConfigRequest
//
// @return ModifyInstanceServiceConfigResponse
func (client *Client) ModifyInstanceServiceConfig(request *ModifyInstanceServiceConfigRequest) (_result *ModifyInstanceServiceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceServiceConfigResponse{}
	_body, _err := client.ModifyInstanceServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ModifyInstanceType to change the specifications of an instance.
//
// @param request - ModifyInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceTypeResponse
func (client *Client) ModifyInstanceTypeWithOptions(request *ModifyInstanceTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CoreInstanceType) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !dara.IsNil(request.MasterInstanceType) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceType"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls ModifyInstanceType to change the specifications of an instance.
//
// @param request - ModifyInstanceTypeRequest
//
// @return ModifyInstanceTypeResponse
func (client *Client) ModifyInstanceType(request *ModifyInstanceTypeRequest) (_result *ModifyInstanceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceTypeResponse{}
	_body, _err := client.ModifyInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP whitelist group of an instance by calling ModifyIpWhitelist.
//
// @param request - ModifyIpWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpWhitelistResponse
func (client *Client) ModifyIpWhitelistWithOptions(request *ModifyIpWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IpList) {
		query["IpList"] = request.IpList
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIpWhitelist"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP whitelist group of an instance by calling ModifyIpWhitelist.
//
// @param request - ModifyIpWhitelistRequest
//
// @return ModifyIpWhitelistResponse
func (client *Client) ModifyIpWhitelist(request *ModifyIpWhitelistRequest) (_result *ModifyIpWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyIpWhitelistResponse{}
	_body, _err := client.ModifyIpWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the node specifications of a multi-zone instance by calling ModifyMultiZoneClusterNodeType.
//
// @param request - ModifyMultiZoneClusterNodeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMultiZoneClusterNodeTypeResponse
func (client *Client) ModifyMultiZoneClusterNodeTypeWithOptions(request *ModifyMultiZoneClusterNodeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyMultiZoneClusterNodeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CoreInstanceType) {
		query["CoreInstanceType"] = request.CoreInstanceType
	}

	if !dara.IsNil(request.LogInstanceType) {
		query["LogInstanceType"] = request.LogInstanceType
	}

	if !dara.IsNil(request.MasterInstanceType) {
		query["MasterInstanceType"] = request.MasterInstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMultiZoneClusterNodeType"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMultiZoneClusterNodeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the node specifications of a multi-zone instance by calling ModifyMultiZoneClusterNodeType.
//
// @param request - ModifyMultiZoneClusterNodeTypeRequest
//
// @return ModifyMultiZoneClusterNodeTypeResponse
func (client *Client) ModifyMultiZoneClusterNodeType(request *ModifyMultiZoneClusterNodeTypeRequest) (_result *ModifyMultiZoneClusterNodeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMultiZoneClusterNodeTypeResponse{}
	_body, _err := client.ModifyMultiZoneClusterNodeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ModifySecurityGroups to modify the security groups associated with a cluster.
//
// @param request - ModifySecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityGroupsResponse
func (client *Client) ModifySecurityGroupsWithOptions(request *ModifySecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.SecurityGroupIds) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityGroups"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls ModifySecurityGroups to modify the security groups associated with a cluster.
//
// @param request - ModifySecurityGroupsRequest
//
// @return ModifySecurityGroupsResponse
func (client *Client) ModifySecurityGroups(request *ModifySecurityGroupsRequest) (_result *ModifySecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySecurityGroupsResponse{}
	_body, _err := client.ModifySecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password for accessing the cluster management UI by calling ModifyUIProxyAccountPassword.
//
// @param request - ModifyUIAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUIAccountPasswordResponse
func (client *Client) ModifyUIAccountPasswordWithOptions(request *ModifyUIAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifyUIAccountPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUIAccountPassword"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUIAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password for accessing the cluster management UI by calling ModifyUIProxyAccountPassword.
//
// @param request - ModifyUIAccountPasswordRequest
//
// @return ModifyUIAccountPasswordResponse
func (client *Client) ModifyUIAccountPassword(request *ModifyUIAccountPasswordRequest) (_result *ModifyUIAccountPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUIAccountPasswordResponse{}
	_body, _err := client.ModifyUIAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a target instance to a target resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a target instance to a target resource group.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the backup and recovery feature for an ApsaraDB for HBase cluster by calling OpenBackup.
//
// @param request - OpenBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenBackupResponse
func (client *Client) OpenBackupWithOptions(request *OpenBackupRequest, runtime *dara.RuntimeOptions) (_result *OpenBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenBackup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the backup and recovery feature for an ApsaraDB for HBase cluster by calling OpenBackup.
//
// @param request - OpenBackupRequest
//
// @return OpenBackupResponse
func (client *Client) OpenBackup(request *OpenBackupRequest) (_result *OpenBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenBackupResponse{}
	_body, _err := client.OpenBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Completely cleans up an instance that has been deleted (within the last 7 days) but not fully cleaned up.
//
// @param request - PurgeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurgeInstanceResponse
func (client *Client) PurgeInstanceWithOptions(request *PurgeInstanceRequest, runtime *dara.RuntimeOptions) (_result *PurgeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PurgeInstance"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PurgeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Completely cleans up an instance that has been deleted (within the last 7 days) but not fully cleaned up.
//
// @param request - PurgeInstanceRequest
//
// @return PurgeInstanceResponse
func (client *Client) PurgeInstance(request *PurgeInstanceRequest) (_result *PurgeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PurgeInstanceResponse{}
	_body, _err := client.PurgeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the high-availability HBase list for a BDS cluster. This operation corresponds to CreateHaCluster and CreateHbaseHaSlb. It returns the corresponding high-availability list.
//
// @param request - QueryHBaseHaDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryHBaseHaDBResponse
func (client *Client) QueryHBaseHaDBWithOptions(request *QueryHBaseHaDBRequest, runtime *dara.RuntimeOptions) (_result *QueryHBaseHaDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BdsId) {
		query["BdsId"] = request.BdsId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryHBaseHaDB"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryHBaseHaDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the high-availability HBase list for a BDS cluster. This operation corresponds to CreateHaCluster and CreateHbaseHaSlb. It returns the corresponding high-availability list.
//
// @param request - QueryHBaseHaDBRequest
//
// @return QueryHBaseHaDBResponse
func (client *Client) QueryHBaseHaDB(request *QueryHBaseHaDBRequest) (_result *QueryHBaseHaDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryHBaseHaDBResponse{}
	_body, _err := client.QueryHBaseHaDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of databases that can be associated with the current instance by calling QueryXpackRelateDB.
//
// @param request - QueryXpackRelateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryXpackRelateDBResponse
func (client *Client) QueryXpackRelateDBWithOptions(request *QueryXpackRelateDBRequest, runtime *dara.RuntimeOptions) (_result *QueryXpackRelateDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.HasSingleNode) {
		query["HasSingleNode"] = request.HasSingleNode
	}

	if !dara.IsNil(request.RelateDbType) {
		query["RelateDbType"] = request.RelateDbType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryXpackRelateDB"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryXpackRelateDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of databases that can be associated with the current instance by calling QueryXpackRelateDB.
//
// @param request - QueryXpackRelateDBRequest
//
// @return QueryXpackRelateDBResponse
func (client *Client) QueryXpackRelateDB(request *QueryXpackRelateDBRequest) (_result *QueryXpackRelateDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryXpackRelateDBResponse{}
	_body, _err := client.QueryXpackRelateDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an active-active instance pair to a BDS cluster.
//
// @param request - RelateDbForHBaseHaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RelateDbForHBaseHaResponse
func (client *Client) RelateDbForHBaseHaWithOptions(request *RelateDbForHBaseHaRequest, runtime *dara.RuntimeOptions) (_result *RelateDbForHBaseHaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.HaActive) {
		query["HaActive"] = request.HaActive
	}

	if !dara.IsNil(request.HaActiveClusterKey) {
		query["HaActiveClusterKey"] = request.HaActiveClusterKey
	}

	if !dara.IsNil(request.HaActiveDBType) {
		query["HaActiveDBType"] = request.HaActiveDBType
	}

	if !dara.IsNil(request.HaActiveHbaseFsDir) {
		query["HaActiveHbaseFsDir"] = request.HaActiveHbaseFsDir
	}

	if !dara.IsNil(request.HaActiveHdfsUri) {
		query["HaActiveHdfsUri"] = request.HaActiveHdfsUri
	}

	if !dara.IsNil(request.HaActivePassword) {
		query["HaActivePassword"] = request.HaActivePassword
	}

	if !dara.IsNil(request.HaActiveUser) {
		query["HaActiveUser"] = request.HaActiveUser
	}

	if !dara.IsNil(request.HaActiveVersion) {
		query["HaActiveVersion"] = request.HaActiveVersion
	}

	if !dara.IsNil(request.HaMigrateType) {
		query["HaMigrateType"] = request.HaMigrateType
	}

	if !dara.IsNil(request.HaStandby) {
		query["HaStandby"] = request.HaStandby
	}

	if !dara.IsNil(request.HaStandbyClusterKey) {
		query["HaStandbyClusterKey"] = request.HaStandbyClusterKey
	}

	if !dara.IsNil(request.HaStandbyDBType) {
		query["HaStandbyDBType"] = request.HaStandbyDBType
	}

	if !dara.IsNil(request.HaStandbyHbaseFsDir) {
		query["HaStandbyHbaseFsDir"] = request.HaStandbyHbaseFsDir
	}

	if !dara.IsNil(request.HaStandbyHdfsUri) {
		query["HaStandbyHdfsUri"] = request.HaStandbyHdfsUri
	}

	if !dara.IsNil(request.HaStandbyPassword) {
		query["HaStandbyPassword"] = request.HaStandbyPassword
	}

	if !dara.IsNil(request.HaStandbyUser) {
		query["HaStandbyUser"] = request.HaStandbyUser
	}

	if !dara.IsNil(request.HaStandbyVersion) {
		query["HaStandbyVersion"] = request.HaStandbyVersion
	}

	if !dara.IsNil(request.HaTables) {
		query["HaTables"] = request.HaTables
	}

	if !dara.IsNil(request.IsActiveStandard) {
		query["IsActiveStandard"] = request.IsActiveStandard
	}

	if !dara.IsNil(request.IsStandbyStandard) {
		query["IsStandbyStandard"] = request.IsStandbyStandard
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RelateDbForHBaseHa"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RelateDbForHBaseHaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an active-active instance pair to a BDS cluster.
//
// @param request - RelateDbForHBaseHaRequest
//
// @return RelateDbForHBaseHaResponse
func (client *Client) RelateDbForHBaseHa(request *RelateDbForHBaseHaRequest) (_result *RelateDbForHBaseHaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RelateDbForHBaseHaResponse{}
	_body, _err := client.RelateDbForHBaseHaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an instance.
//
// @param request - ReleasePublicNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleasePublicNetworkAddressResponse
func (client *Client) ReleasePublicNetworkAddressWithOptions(request *ReleasePublicNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleasePublicNetworkAddress"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an instance.
//
// @param request - ReleasePublicNetworkAddressRequest
//
// @return ReleasePublicNetworkAddressResponse
func (client *Client) ReleasePublicNetworkAddress(request *ReleasePublicNetworkAddressRequest) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.ReleasePublicNetworkAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a subscription for an HBase instance.
//
// Description:
//
// This operation applies only to subscription HBase instances.
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2019-01-01"),
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
// Renews a subscription for an HBase instance.
//
// Description:
//
// This operation applies only to subscription HBase instances.
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
// Changes the cold storage size.
//
// @param request - ResizeColdStorageSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeColdStorageSizeResponse
func (client *Client) ResizeColdStorageSizeWithOptions(request *ResizeColdStorageSizeRequest, runtime *dara.RuntimeOptions) (_result *ResizeColdStorageSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ColdStorageSize) {
		query["ColdStorageSize"] = request.ColdStorageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResizeColdStorageSize"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeColdStorageSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the cold storage size.
//
// @param request - ResizeColdStorageSizeRequest
//
// @return ResizeColdStorageSizeResponse
func (client *Client) ResizeColdStorageSize(request *ResizeColdStorageSizeRequest) (_result *ResizeColdStorageSizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResizeColdStorageSizeResponse{}
	_body, _err := client.ResizeColdStorageSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ResizeDiskSize to resize a disk.
//
// @param request - ResizeDiskSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeDiskSizeResponse
func (client *Client) ResizeDiskSizeWithOptions(request *ResizeDiskSizeRequest, runtime *dara.RuntimeOptions) (_result *ResizeDiskSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeDiskSize) {
		query["NodeDiskSize"] = request.NodeDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResizeDiskSize"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeDiskSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls ResizeDiskSize to resize a disk.
//
// @param request - ResizeDiskSizeRequest
//
// @return ResizeDiskSizeResponse
func (client *Client) ResizeDiskSize(request *ResizeDiskSizeRequest) (_result *ResizeDiskSizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResizeDiskSizeResponse{}
	_body, _err := client.ResizeDiskSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the ResizeMultiZoneClusterDiskSize operation to modify the disk size of a multi-zone instance.
//
// @param request - ResizeMultiZoneClusterDiskSizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeMultiZoneClusterDiskSizeResponse
func (client *Client) ResizeMultiZoneClusterDiskSizeWithOptions(request *ResizeMultiZoneClusterDiskSizeRequest, runtime *dara.RuntimeOptions) (_result *ResizeMultiZoneClusterDiskSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CoreDiskSize) {
		query["CoreDiskSize"] = request.CoreDiskSize
	}

	if !dara.IsNil(request.LogDiskSize) {
		query["LogDiskSize"] = request.LogDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResizeMultiZoneClusterDiskSize"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeMultiZoneClusterDiskSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ResizeMultiZoneClusterDiskSize operation to modify the disk size of a multi-zone instance.
//
// @param request - ResizeMultiZoneClusterDiskSizeRequest
//
// @return ResizeMultiZoneClusterDiskSizeResponse
func (client *Client) ResizeMultiZoneClusterDiskSize(request *ResizeMultiZoneClusterDiskSizeRequest) (_result *ResizeMultiZoneClusterDiskSizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResizeMultiZoneClusterDiskSizeResponse{}
	_body, _err := client.ResizeMultiZoneClusterDiskSizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales out nodes for a multi-zone instance by calling the ResizeMultiZoneClusterNodeCount operation.
//
// @param request - ResizeMultiZoneClusterNodeCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeMultiZoneClusterNodeCountResponse
func (client *Client) ResizeMultiZoneClusterNodeCountWithOptions(request *ResizeMultiZoneClusterNodeCountRequest, runtime *dara.RuntimeOptions) (_result *ResizeMultiZoneClusterNodeCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArbiterVSwitchId) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CoreNodeCount) {
		query["CoreNodeCount"] = request.CoreNodeCount
	}

	if !dara.IsNil(request.LogNodeCount) {
		query["LogNodeCount"] = request.LogNodeCount
	}

	if !dara.IsNil(request.PrimaryCoreNodeCount) {
		query["PrimaryCoreNodeCount"] = request.PrimaryCoreNodeCount
	}

	if !dara.IsNil(request.PrimaryVSwitchId) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !dara.IsNil(request.StandbyCoreNodeCount) {
		query["StandbyCoreNodeCount"] = request.StandbyCoreNodeCount
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResizeMultiZoneClusterNodeCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeMultiZoneClusterNodeCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales out nodes for a multi-zone instance by calling the ResizeMultiZoneClusterNodeCount operation.
//
// @param request - ResizeMultiZoneClusterNodeCountRequest
//
// @return ResizeMultiZoneClusterNodeCountResponse
func (client *Client) ResizeMultiZoneClusterNodeCount(request *ResizeMultiZoneClusterNodeCountRequest) (_result *ResizeMultiZoneClusterNodeCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResizeMultiZoneClusterNodeCountResponse{}
	_body, _err := client.ResizeMultiZoneClusterNodeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls ResizeNodeCount to adjust the number of nodes for an instance.
//
// Description:
//
// Scales up the number of core nodes in a cluster. You can add up to 50 nodes at a time, and the total number of nodes can be scaled up to 250. If you have additional requirements, submit a ticket.
//
// @param request - ResizeNodeCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeNodeCountResponse
func (client *Client) ResizeNodeCountWithOptions(request *ResizeNodeCountRequest, runtime *dara.RuntimeOptions) (_result *ResizeNodeCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeCount) {
		query["NodeCount"] = request.NodeCount
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
		Action:      dara.String("ResizeNodeCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeNodeCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls ResizeNodeCount to adjust the number of nodes for an instance.
//
// Description:
//
// Scales up the number of core nodes in a cluster. You can add up to 50 nodes at a time, and the total number of nodes can be scaled up to 250. If you have additional requirements, submit a ticket.
//
// @param request - ResizeNodeCountRequest
//
// @return ResizeNodeCountResponse
func (client *Client) ResizeNodeCount(request *ResizeNodeCountRequest) (_result *ResizeNodeCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResizeNodeCountResponse{}
	_body, _err := client.ResizeNodeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an HBase cluster by calling RestartInstance.
//
// @param request - RestartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Components) {
		query["Components"] = request.Components
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartInstance"),
		Version:     dara.String("2019-01-01"),
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
// Restarts an HBase cluster by calling RestartInstance.
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
// Revokes permissions from a wide table account.
//
// @param request - RevokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResponse
func (client *Client) RevokeWithOptions(request *RevokeRequest, runtime *dara.RuntimeOptions) (_result *RevokeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AclActions) {
		query["AclActions"] = request.AclActions
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Revoke"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions from a wide table account.
//
// @param request - RevokeRequest
//
// @return RevokeResponse
func (client *Client) Revoke(request *RevokeRequest) (_result *RevokeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeResponse{}
	_body, _err := client.RevokeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a primary/secondary switchover for high-availability Thrift or high-availability Phoenix. This operation corresponds to the CreateHbaseHaSlb operation.
//
// @param request - SwitchHbaseHaSlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchHbaseHaSlbResponse
func (client *Client) SwitchHbaseHaSlbWithOptions(request *SwitchHbaseHaSlbRequest, runtime *dara.RuntimeOptions) (_result *SwitchHbaseHaSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BdsId) {
		query["BdsId"] = request.BdsId
	}

	if !dara.IsNil(request.HaId) {
		query["HaId"] = request.HaId
	}

	if !dara.IsNil(request.HaTypes) {
		query["HaTypes"] = request.HaTypes
	}

	if !dara.IsNil(request.HbaseType) {
		query["HbaseType"] = request.HbaseType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchHbaseHaSlb"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchHbaseHaSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a primary/secondary switchover for high-availability Thrift or high-availability Phoenix. This operation corresponds to the CreateHbaseHaSlb operation.
//
// @param request - SwitchHbaseHaSlbRequest
//
// @return SwitchHbaseHaSlbResponse
func (client *Client) SwitchHbaseHaSlb(request *SwitchHbaseHaSlbRequest) (_result *SwitchHbaseHaSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchHbaseHaSlbResponse{}
	_body, _err := client.SwitchHbaseHaSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Enable or disable an extension service
//
// Description:
//
// # Extension service
//
// For example: the Lindorm service compatible with the HBaseProxy protocol.
//
// @param request - SwitchServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchServiceResponse
func (client *Client) SwitchServiceWithOptions(request *SwitchServiceRequest, runtime *dara.RuntimeOptions) (_result *SwitchServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Operate) {
		query["Operate"] = request.Operate
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchService"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enable or disable an extension service
//
// Description:
//
// # Extension service
//
// For example: the Lindorm service compatible with the HBaseProxy protocol.
//
// @param request - SwitchServiceRequest
//
// @return SwitchServiceResponse
func (client *Client) SwitchService(request *SwitchServiceRequest) (_result *SwitchServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchServiceResponse{}
	_body, _err := client.SwitchServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to instances.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-01-01"),
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
// Adds tags to instances.
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
// Untags resources. This operation is the counterpart of TagResources.
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnTagResources"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Untags resources. This operation is the counterpart of TagResources.
//
// @param request - UnTagResourcesRequest
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the minor version of a cluster by calling UpgradeMinorVersion.
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
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Components) {
		query["Components"] = request.Components
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeMinorVersion"),
		Version:     dara.String("2019-01-01"),
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
// Upgrades the minor version of a cluster by calling UpgradeMinorVersion.
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

// Summary:
//
// Upgrades the minor version of components for a multi-zone instance by calling the UpgradeMultiZoneCluster operation.
//
// @param request - UpgradeMultiZoneClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeMultiZoneClusterResponse
func (client *Client) UpgradeMultiZoneClusterWithOptions(request *UpgradeMultiZoneClusterRequest, runtime *dara.RuntimeOptions) (_result *UpgradeMultiZoneClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Components) {
		query["Components"] = request.Components
	}

	if !dara.IsNil(request.RestartComponents) {
		query["RestartComponents"] = request.RestartComponents
	}

	if !dara.IsNil(request.RunMode) {
		query["RunMode"] = request.RunMode
	}

	if !dara.IsNil(request.UpgradeInsName) {
		query["UpgradeInsName"] = request.UpgradeInsName
	}

	if !dara.IsNil(request.Versions) {
		query["Versions"] = request.Versions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeMultiZoneCluster"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeMultiZoneClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the minor version of components for a multi-zone instance by calling the UpgradeMultiZoneCluster operation.
//
// @param request - UpgradeMultiZoneClusterRequest
//
// @return UpgradeMultiZoneClusterResponse
func (client *Client) UpgradeMultiZoneCluster(request *UpgradeMultiZoneClusterRequest) (_result *UpgradeMultiZoneClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeMultiZoneClusterResponse{}
	_body, _err := client.UpgradeMultiZoneClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a database by calling XpackRelateDB.
//
// @param request - XpackRelateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return XpackRelateDBResponse
func (client *Client) XpackRelateDBWithOptions(request *XpackRelateDBRequest, runtime *dara.RuntimeOptions) (_result *XpackRelateDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DbClusterIds) {
		query["DbClusterIds"] = request.DbClusterIds
	}

	if !dara.IsNil(request.RelateDbType) {
		query["RelateDbType"] = request.RelateDbType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("XpackRelateDB"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &XpackRelateDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a database by calling XpackRelateDB.
//
// @param request - XpackRelateDBRequest
//
// @return XpackRelateDBResponse
func (client *Client) XpackRelateDB(request *XpackRelateDBRequest) (_result *XpackRelateDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &XpackRelateDBResponse{}
	_body, _err := client.XpackRelateDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
